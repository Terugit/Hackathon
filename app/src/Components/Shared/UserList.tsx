import { useState, useEffect, useContext, memo, FC} from 'react';
import { Drawer, Button, Group, Box, Flex } from '@mantine/core';
import { UserContext } from './Context';
import { UserCard } from './UserCard';
import {AddUser} from "./AddUser";
import { useHistory } from 'react-router-dom';
import { ActiveContext } from './ActiveProvider';
import ReactLoading from "react-loading";

type User={
  name : string;
  id : string;
  photo : string;
  point : number;
};

export const UserList=() => {
  // console.log("UserLIst")
  const [isLoading ,setLoading]= useState<boolean>(false); 
  const history = useHistory();
  const [data, setData] = useState<User[]>([])
  const [opened, setOpened] = useState(false);
  const get = async () => {
    // setLoading(true);
      const response = await fetch("https://hackathon14-qftu2uez4a-uc.a.run.app/user",
        // "http://localhost:8080/user",
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      },
    );
    const nowData : User[] = await response.json();
    setData(nowData);
    setLoading(false);
  };
  useEffect(() => {setLoading(true); get()},[]);
  const {set} =useContext(ActiveContext);
  const {setUser} =useContext(UserContext);
  const onSubmit = ( member : User)  => {
    setUser(member.id, member.photo, member.name, member.point);
    set(-1) ;
    history.push("/");
   }


return (
<>
      <Drawer
        opened={opened}
        onClose={() => setOpened(false)}
        // title="Select user"
        padding="sm"
        size="sm"
        position="top" 
      >
        {isLoading &&    <Flex justify="center" align="center"> 
    <section className="flex justify-center items-center h-screen">
      <div>
        <p></p>
        <ReactLoading
          type="spin"
          color='#8ED1F4'
          height="50px"
          width="50px"
          className="mx-auto"
        />
        
      </div>
    </section>
    </Flex>}
    {isLoading===false &&
        <>
        <AddUser reload={get} ></AddUser>
        <p></p>
        {data.map((user :User) => (
              <Box onClick={()=>onSubmit(user)} key={user.id} className="usercard"
              sx={(theme) => ({
                // backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[6] : theme.colors.gray[0],
                textAlign: 'center',
                padding: theme.spacing.xs,
                // borderRadius: theme.radius.xs,
                cursor: 'pointer',
        
                '&:hover': {
                  backgroundColor:
                    theme.colorScheme === 'dark' ? theme.colors.dark[5] : theme.colors.gray[1],
                },
              })}
            >
    
              
              <UserCard  user={user}  />
            </Box>
            
            ))}
            </>
}
          <Box>
        
            
          </Box>

        
      </Drawer>
{/* <Button  onClick={() => setOpened(true)} color='#EB94E2' */
}
      <Group position="center">
      <Button style={{backgroundColor: 'blue', width: 200, marginLeft: 'auto', marginRight: 'auto' }} variant="gradient" gradient={{ from: '#000000', to: '#000001', deg: 35 }} onClick={() => setOpened(true)} >Select a user</Button>
      </Group>

    

  
  </>
);
}

export default UserList;
