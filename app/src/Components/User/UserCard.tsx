import { Group, Avatar, Text } from '@mantine/core';

type User={
    name : string;
    id : string;
    photo : string;
    point : number;
  };

type Props = {
  user: User;
};

export const UserCard=(props: Props) => {
  const { user } = props; // ユーザー情報受け取り

  return (
    <>
      <Group>
        <Avatar size={50} src={user.photo}></Avatar>
        <div>
          <Text size="xl">{user.name}</Text>
        </div>
      </Group>
    </>
  );
};