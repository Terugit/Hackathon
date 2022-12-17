import {FC, useState, useContext } from 'react';
import { useHistory } from 'react-router-dom';
import {UserContext} from "../Shared/Context";
// import {ThankFrom} from "./FromUser";

type Thanks ={
    id : string;
    to_ : string;
    from_photo : string;
    fromWho : string;
    to_photo : string;
    toWho : string;
    point : number;
    message : string;
    postAt : string;
    editAt : string;
}

type EditThankProps = Thanks & {
  reload: () =>Promise<void>
} 
export const EditThank: FC<EditThankProps> =(props) =>{
  const [message, setMessage]  = useState<string>(props.message);
  const [point, setPoint]  = useState<number>(props.point);
  const history = useHistory();
  const [thank, setThank] = useState<Thanks[]>([])
  const url = "https://hackathon14-qftu2uez4a-uc.a.run.app/user/thanks/sent?id="+useContext(UserContext).id;
  
  const onSubmit = async(e: React.FormEvent<HTMLFormElement>)=> {
      e.preventDefault();
      const time = new Date().toLocaleString();
      if (point<=0){
          alert ("Point: 0より大きい整数値を入力してください。");
          return;
        }
      if (point%1!==0){
          alert("Point: 0より大きい整数値を入力してください。");
          return;
      }
      if (message ===""){
          alert("メッセージを入力してください。")
          return;
      }
      if (message.length >5000){
        alert("5000字以内で入力してください。");
        return;
      }
      if (time ===""){
          alert("もう一度送信してください。");
          return;
        }
      try{
        const result = 
          await fetch("https://hackathon14-qftu2uez4a-uc.a.run.app/thank/edit",{
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
          "id" : props.id,
          "to_" : props.to_,
          "point" : point,
          "delete_point" : props.point,
          "message" : message,
          "editAt": time,
              }), 
        }
      );
      if (!result.ok){
        throw Error('Failed to edit thank : ${result.status}');
      }
      await props.reload()
    }catch (err){
      console.error(err);
    };
    // console.log(props.point)
    };

    const onDelete = async()=> {
      try{
        const result = 
          await fetch("https://hackathon14-qftu2uez4a-uc.a.run.app/thank/delete",{
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
          "id" : props.id,
          "to_" : props.to_,
          "point": props.point,
              }), 
        }
      );
      if (!result.ok){
        throw Error('Failed to delete thank : ${result.status}');
      }
      await props.reload()
    }catch (err){
      console.error(err);
    };
    };
 

  return (
    <>
       
    <form onSubmit={onSubmit} style={{ display: "flex", flexDirection: "column" }}>
    <label>Point: </label>
   <input
     type={"number"}
     value={point}
     onChange={(e) => setPoint(e.target.valueAsNumber)}
   ></input>
   <label>Message: </label>
   <input
     type={"text"}
     value={message}
     onChange={(e) => setMessage(e.target.value)}
   ></input>
   
 
    <button>Edit</button>
 </form>
 <button onClick={onDelete}>Delete</button>
    </>)
  ;
}