import React, { useState, useEffect, useContext,forwardRef} from "react";
import { UserContext } from "../Shared/Context";
import { Select, Group, Avatar, Text, SelectItem, Flex} from '@mantine/core';
import "./Thank.css";
import ReactLoading from "react-loading";

type User={
  name : string;
  id : string;
  photo : string;
  point : number;
}
type AddUser={
  name : string;
  id : string;
  photo : string;
  label : string;
  value : string;
}


const Thank = () => {
  const useId = useContext(UserContext).id;
  const [message, Message]  = useState<string>("");
  const [point, Point]  = useState<number>(0);
  const [value, Value ]  =useState<string | null>(null);
  const [isLoading ,setLoading]= useState<boolean>(false); 
  const onSubmit = async(e: React.FormEvent<HTMLFormElement>)=> {
    // setLoading(true);
    //console.log(value)
    e.preventDefault();
    const time = new Date().toLocaleString();
    // console.log(time)
    if (!useId){
      alert("ログインするアカウントを選択してください");
      return;
    };
    if (!value){
      alert("送信先を選択してください。");
      return;
    };
    if(value===useId){
      alert("自分以外の人を選んでください")
      return
    };
    if (point<=0){
        alert ("Point: 0より大きい整数値を入力してください。");
        return;
      };
    if (point%1!==0){
        alert("Point: 0より大きい整数値を入力してください。");
        return;
    };
    if (message ===""){
        alert("メッセージを入力してください。")
        return;
    };
    if (message.length >5000){
      alert("5000字以内で入力してください。");
      return;
    };
    if (time ===""){
        alert("もう一度送信してください。");
        return;
      };
    try{
      const result = 
        await fetch("https://hackathon14-qftu2uez4a-uc.a.run.app/main",{
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
        "from_" : useId,
        "to_" : value,
        "point" : point,
        "message" : message,
        "postAt": time,
        "editAt": time,
            }), 
      }
    );
    if (!result.ok){
      throw Error('Failed to create user : ${result.status}');
    }
    Value("");
    Message("");
    Point(0);
  }catch (err){
    console.error(err);
  }
  };




  const [data, setData] = useState<User[]>([])
  const [addData, setAddData] = useState<AddUser[]>([])
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
    const nowData: User[] = await response.json();
    setData(nowData);
    const users = nowData.map((user : User)=>{
      return {...user,
      label : user.name,
      value: user.id}
    });
    setAddData(users);
    setLoading(false);
  }

  useEffect(() => {setLoading(true); get()},[]);
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
}else{

  return (
    <div>
    <form onSubmit={onSubmit} style={{ display: "flex", flexDirection: "column" }}>
      <label>To:</label>
    <Select
      // label="To:"
      placeholder="送りたい相手を選択してください"
      itemComponent={SelectItem}
      data={addData}
      searchable
      searchValue={searchValue}
      onSearchChange={onSearchChange}
      nothingFound="Nothing found"
      filter={(value : string, item: SelectItem) =>
        (item.id!==useId)&&(item.name.toLowerCase().includes(value.toLowerCase().trim()))
      }
      value={value}
      onChange={Value}
    />
    <label>Point: </label>
      <input
        type={"number"}
        value={point}
        onChange={(e) => Point(e.target.valueAsNumber)}
      ></input>

      <label>Message: </label>
      <textarea
        value={message}
        onChange={(e) => Message(e.target.value)}
        className="message"
      ></textarea>
       <button>Thank</button>
    </form>
    </div>
  );
};}

export default Thank;