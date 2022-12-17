import {useState, useEffect, useContext} from "react";
import { Group, Avatar, Text, Box , Flex, Divider} from '@mantine/core';
import {MdOutlineDoubleArrow} from "react-icons/md";
import { IconContext } from 'react-icons';
import ReactLoading from "react-loading";

type Thanks= {
    id : string
    fromWho : string;
    toWho : string;
    from_photo : string;
    to_photo : string;
    point : number;
    message : string;
    postAt: string;
    editAt : string;
  }

  export const Thanks=() =>  {  
    const [thank, setThank] = useState([]);
    const [isLoading ,setLoading]= useState<boolean>(false);
    const getconst = async () => {
      // setLoading(true);
      const response = await fetch("https://hackathon14-qftu2uez4a-uc.a.run.app/main",
         {
           method: "GET",
           headers: {"Content-Type": "application/json",},
                },
              );
      const nowThank = await response.json();
      setThank(nowThank);
      setLoading(false);
      // console.log("false")  
        };
        useEffect(() => {setLoading(true); getconst()},[]);
  
interface AccordionLabelProps {
    id : string;
    fromWho : string;
    toWho : string;
    from_photo : string;
    to_photo : string;
    point : number;
    message : string;
    postAt : string;
    editAt : string;
}

const AccordionLabel = (item: AccordionLabelProps) =>{

    return <>
    <div style={{margin: "auto"}}>
    <Group noWrap>
    <Flex gap="xl" align="center">
    <div>
        <Avatar src={item.from_photo} radius="xl" size="sm" />
        <Text>{item.fromWho}</Text>
    </div>
    <div>
      <Flex gap="md" align="center">
      <Text >{item.point} Pt</Text>
        <IconContext.Provider value={{ color: '#8ED1F4', size: '10px' }}>
            <MdOutlineDoubleArrow/>
        </IconContext.Provider>
        
      </Flex>
    </div>
    <div>
        <Avatar src={item.to_photo} radius="xl" size="sm" /> 
        <Text>{item.toWho}</Text>
    </div>
    
      </Flex>
    </Group>
    <div>
        
        <Text size='xl' weight={4000}>
          {item.message}
        </Text>
        <Text size="sm" color="dimmed" weight={400}>{item.postAt}{item.postAt!==item.editAt && (<> (編集済み)</>)}</Text>
      </div>
    </div>
    </>
    ;
  }

    const items = thank.map((item : Thanks) => (

      <div key={item.id}>
      <Flex align="center">
      <Box color="black">
            
        <AccordionLabel {...item} />

    </Box>
    </Flex>
    <Divider my="sm" color='#8ED1F4'/>
    </div>
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
      return <>
      <Flex >
      <Text 
      variant="gradient"
      gradient={{ from: '#ffffff', to: '#EB94E2',deg: 35}}
      fz="xl" fw={3000} >All Thanks</Text>
      </Flex>
      <Divider my="sm" color='#8ED1F4' />
      {items}
    </>;
  }
}
