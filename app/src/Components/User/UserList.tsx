import { useState, useEffect } from "react";
import { Box } from '@mantine/core';
import {UserCard} from "./UserCard"
import "../../Form";
type User={
    id : string;
    name : string;
    photo : string;
    point : number;
  }

export const UserList = () => {
    const [data, setData] = useState<User[]>([])
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
      const nowData = await response.json();
      setData(nowData);
    };
      useEffect(() => {get()},[]);

    return (
      <div>

      {data.map((user :User) => (
        <Box key={user.id}
        sx={(theme) => ({
          textAlign: 'center',
          padding: theme.spacing.xl,
          borderRadius: theme.radius.xl,
          cursor: 'pointer',
  
          // '&:hover': {
          //   backgroundColor:
          //     theme.colorScheme === 'dark' ? theme.colors.dark[5] : theme.colors.gray[1],
          // },
        })}
      >
        <UserCard  user={user}  />
      </Box>
      ))}
      </div>

    );
  };