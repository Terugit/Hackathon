import { Table, Flex } from '@mantine/core';
import { useState, useEffect } from 'react';
import { UserCard } from './UserCard';
import ReactLoading from "react-loading";

type UserRank={
    name : string;
    id : string;
    photo : string;
    point : number;
    rank :number;
  };

export const UserTable =()=>{
  const [data, setData] = useState<UserRank[]>([]);
  const [isLoading ,setLoading]= useState<boolean>(false);
const get = async () => {
  // setLoading(true);
    const response = await fetch("https://hackathon14-qftu2uez4a-uc.a.run.app/ranking",
      // "http://localhost:8080/ranking",
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    },
  );
  const nowData = await response.json();
  setData(nowData);
  // console.log(nowData);
  setLoading(false);
};
useEffect(() => {setLoading(true); get()},[]);

  const rows = data.map((user : UserRank) => (
    <tr key={user.id}>
      <td>{user.rank}</td>
      <td>
    <UserCard user={user}></UserCard>
        </td>
      <td>{user.point} Pt</td>

    </tr>
  ));

  if (isLoading) {
    return (
      <Flex justify="center" align="center"> 
      <section className="flex justify-center items-center h-screen">
        <div>
          <p></p>
          <ReactLoading
            type="spin"
            color='#8ED1F4'
            height="100px"
            width="100px"
            className="mx-auto"
          />
          
        </div>
      </section>
      </Flex>
    );
  } else {
  return (
    <Table fontSize="sm">
      <thead>
        <tr>
        <th>Rank</th>
          <th>Name</th>
          <th>Point Got</th>
        </tr>
      </thead>
      <tbody>{rows}</tbody>
    </Table>
  );
}
}