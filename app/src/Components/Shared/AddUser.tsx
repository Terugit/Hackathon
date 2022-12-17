import { useState } from 'react';
import { Modal, Button, Group } from '@mantine/core';

type Props={
    reload :() => Promise<void>
}
export const AddUser = (props:  Props) => {
    const [opened, setOpened] = useState(false);
    const [name, setName]  = useState<string>("");
    const [photo, setPhoto]  = useState<string>("");   
  
    const onSubmit = async(e: React.FormEvent<HTMLFormElement>)=> {
        e.preventDefault();
        if (name===''){
            alert ("名前を入力してください");
            return;
          }
        // if (photo===''){
        //     alert("プロフィール写真のリンクを入力してください。");
        //     return;
        // }
        try{
          const result = 
            await fetch("https://hackathon14-qftu2uez4a-uc.a.run.app/user/add",
                // "http://localhost:8080/registermember",
                {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
            "name" : name,
            "photo" : photo
                }), 
          }
        );
        if (!result.ok){
          throw Error('Failed to edit profile : ${result.status}');
        }
        setOpened(false);
        props.reload();
      }catch (err){
        console.error(err);
      };

      };

  return (
    <>
      <Modal
        opened={opened}
        onClose={() => setOpened(false)}
        // title="This is fullscreen modal!"
        fullScreen
      >
        
        <form onSubmit={onSubmit} style={{ display: "flex", flexDirection: "column" }}>
      <label>Name: </label>
     <input
       type={"string"}
       value={name}
       onChange={(e) => setName(e.target.value)}
     ></input>
     {/* <label>Photo URL: </label>
     <input
       type={"string"}
       value={photo}
       onChange={(e) => setPhoto(e.target.value)}
     ></input> */}
     
   
      <button>Register</button>
   </form>

      </Modal>

      <Group position="center">
        <Button onClick={() => setOpened(true)} variant="gradient" gradient={{ from: '#000000', to: '#000001', deg: 35 }}>Add a user</Button>
      </Group>
    </>
  );
}