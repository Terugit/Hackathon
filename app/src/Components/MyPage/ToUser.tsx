import { useState, useEffect, useContext} from "react";
import {UserContext} from "../Shared/Context";
import { Group, Avatar, Text, Accordion, Flex} from '@mantine/core';
import {MdOutlineDoubleArrow} from "react-icons/md";
import { IconContext } from 'react-icons';
import ReactLoading from "react-loading";

type Thanks= {
    id : string
    fromWho : string;
    from_photo : string;
    toWho : string;
    to_photo : string;
    point : number;
    message : string;
    postAt : string;
    editAt : string;
  
  }

export const ThankTo = () => {
  const [isLoading ,setLoading]= useState<boolean>(false); 
    const [thank, setThank] = useState<Thanks[]>([])
    const url = "https://hackathon14-qftu2uez4a-uc.a.run.app/user/thanks/sent?id="+useContext(UserContext).id;
    const getconst = async () => {
      // setLoading(true);
              const response = await fetch(url,
                {
                  method: "GET",
                  headers: {
                    "Content-Type": "application/json",
                  },
                },
              );
              const nowThank = await response.json();
            setThank(nowThank);
            setLoading(false);
            }
            useEffect(() =>  {setLoading(true); getconst()},[]);
                    
  
                    interface AccordionLabelProps {
                      id : string;
                      from_photo : string;
                      fromWho : string;
                      to_photo : string;
                      toWho : string;
                      point : number;
                      message : string;
                      postAt : string;
                      editAt : string;
                  }
//                   const [searchValue, onSearchChange] = useState('');
// interface ItemProps extends React.ComponentPropsWithoutRef<'div'> {
//   id: string;
//   name: string;
//   photo: string;
//   label : string;
// }

const AccordionLabel = (item: AccordionLabelProps) =>{
  return (
    <>
    <Group noWrap>
      <div>
        <Avatar src={item.from_photo} radius="xl" size="lg" />
        <Text>{item.fromWho}</Text>
      </div>
      <div>
        <IconContext.Provider value={{ color: '#8ED1F4', size: '30px' }}>
          <MdOutlineDoubleArrow/>
        </IconContext.Provider>
        <Text >{item.point} Pt</Text>
      </div>
                  
    <div>
        <Avatar src={item.to_photo} radius="xl" size="lg" /> 
        <Text>{item.toWho}</Text>
    </div>
      <div>
        <Text size="sm" color="dimmed" weight={400}>{item.postAt}{item.postAt!==item.editAt && (<> (編集済み)</>)}</Text>
        <Text>
          {item.message}
        </Text>
        
      </div>
    </Group>
    
    </>
      )
       }

const items = thank.map((item : Thanks) => (
  <Accordion.Item value={item.id} key={item.id}>
    <Accordion.Control>
    <AccordionLabel {...item} />
    </Accordion.Control>
    </Accordion.Item>
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
      } else { return <Accordion chevronPosition="right" variant="contained" chevron="" mx="auto">{items}</Accordion>;
    }}
  

                  