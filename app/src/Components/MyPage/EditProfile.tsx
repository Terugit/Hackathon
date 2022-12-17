import {UserContext} from "../Shared/Context";
import { useState , useContext} from 'react';
import "./MyPage.css"

export const EditProfile=()=>{
  // console.log("Edit");
    const id = useContext(UserContext).id
    const [name, setName]  = useState<string>(useContext(UserContext).name);
    const [photo, setPhoto]  = useState<string>(useContext(UserContext).photo);
    const point = useContext(UserContext).point;
    const {setUser} =useContext(UserContext);     
  
    const onSubmit = async(e: React.FormEvent<HTMLFormElement>)=> {
        e.preventDefault();
        if (name===''){
            alert ("名前を入力してください");
            return;
          }
        try{
          const result = 
            await fetch("https://hackathon14-qftu2uez4a-uc.a.run.app/user/edit",{
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
            "id" : id,
            "name" : name,
            "photo" : photo
                }), 
          }
        );
        if (!result.ok){
          throw Error('Failed to edit profile : ${result.status}');
        }
        setUser(id, photo, name, point);
      }catch (err){
        console.error(err);
      };
      }; 
  
    return (
      <>     
      <form onSubmit={onSubmit} style={{ display: "flex", flexDirection: "column" }}>
        <p></p>
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
    
      <button>Edit</button>
   </form>
      </>)
    ;

}