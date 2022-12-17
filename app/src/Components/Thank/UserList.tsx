import { Select, Group, Avatar, Text, SelectItem } from '@mantine/core';
import {useState, useEffect, forwardRef, useContext} from "react";
import { UserContext } from '../Shared/Context';

type User={
  name : string;
  id : number;
  photo : string;
}
type AddUser={
  name : string;
  id : number;
  photo : string;
  label : string;
  value : string;
}

export const UserList = () => {
    const [data, setData] = useState<User[]>([])
    const [addData, setAddData] = useState<AddUser[]>([])
    const useId = useContext(UserContext).id
    const get = async () => {
        const response = await fetch("https://hackathon14-qftu2uez4a-uc.a.run.app/user",
          // "http://localhost:8080/user",
        {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
        },
      );
      const nowData: User[] = await response.json();
      setData(nowData);
      const users = nowData.map((user : User)=>{
        return {...user,
        label : user.name,
        value: user.name}
      });
      setAddData(users);
    };

    useEffect(() => {get()},[]);
  const [searchValue, onSearchChange] = useState('');

  interface ItemProps extends React.ComponentPropsWithoutRef<'div'> {
    id: string;
    name: string;
    photo: string;
    label : string;
  }
    
  const SelectItem = forwardRef<HTMLDivElement, ItemProps>(
    ({ id, name, photo, label, ...others }: ItemProps, ref) => (
      <div ref={ref} {...others} key={id}>
        <Group noWrap>
          <Avatar src={photo} />
          <div>
            <Text>{name}</Text>
          </div>
        </Group>
      </div>
    )
  );

  // console.log(addData)

  return (
    <>
    <Select
      label="誰に送りますか？"
      placeholder="Pick all that you like"
      itemComponent={SelectItem}
      data={addData}
      searchable
      searchValue={searchValue}
      onSearchChange={onSearchChange}
      nothingFound="Nothing found"
      filter={(value : string, item: SelectItem) =>
        (item.id!==useId)&&(item.name.toLowerCase().includes(value.toLowerCase().trim()))
      }
    />
    </>

  );
}