import { useState , useContext, ReactNode} from 'react';
import {
  AppShell,
  Avatar,
  Box,
  Flex,
  Navbar,
  NavLink,
  Header,
  Divider,
  Text,
  MediaQuery,
  Burger,
  useMantineTheme,
  Image,
} from '@mantine/core';
import {UserContext} from "./Context";
import UserList from './UserList';
import { useHistory } from 'react-router-dom';
import { IconAward , IconSend, IconChevronRight, IconStar} from '@tabler/icons';
import { ActiveContext } from './ActiveProvider';
import pic from "./Logo.png";

interface FrameProps {
  children: ReactNode;
}

export const Frame =({children} :FrameProps) => {
  const {set, active} =useContext(ActiveContext);
  const history = useHistory();
  const theme = useMantineTheme();
  const [opened, setOpened] = useState(false);

  const rinkTop =()=>{
    set(-1) ;
    history.push("/");
  };
  const rinkMyPage =()=>{
    set(0) ;
    history.push("/mypage");
  };
  const rinkUser =()=>{
    set(1) ;
    history.push("/ranking");
  };
  const rinkThank =()=>{
    set(2) ;
    history.push("/thank");
  };
  
  const data = [
    {icon: IconStar, label: 'My Page',   rink :rinkMyPage,
    },
    { icon: IconAward, label: 'Ranking',   rink :  rinkUser},
    {
      // icon: IconSend,
      label: 'Thank!',
      //  
      rink :rinkThank,
    },
  ];
  const items = data.map((item, index) => (
      <NavLink
      key={item.label}
      active={index === active}
      label={item.label}
      // rightSection={item.rightSection}
      // icon={<item.icon size={16} stroke={1.5} />}
      onClick={item.rink}
      color='#8ED1F4'
    />
    
  ));
  return (
    <AppShell
      navbarOffsetBreakpoint="sm"
      asideOffsetBreakpoint="sm"
      navbar={
        <Navbar p="md" hiddenBreakpoint="sm" hidden={!opened} width={{ sm: 200, lg: 300 }}>
        <Navbar.Section>
          <Flex justify="center" align="center" gap="xl"> 
           <Avatar src={useContext(UserContext).photo} radius="xl" size="xl" />
           <Text fz="xl">{useContext(UserContext).name}</Text> 
          </Flex>
          <p></p>
          <Flex justify="center" align="center" gap="xs">
            <Text>Point Got：</Text>
            <Text  fz="lg" c="cyan">{useContext(UserContext).point} </Text>
            <Text> point</Text>
          </Flex>
            
            </Navbar.Section>
            <Divider my="sm" color='#8ED1F4' />
          <Box >{items}</Box>
        </Navbar>
      }
      header={
        <Header height={{ base: 80 }} p="md" style={{ display: 'flex', alignItems: 'center' ,justifyContent:'space-between'}} >
          <Flex gap="md" align="center">
          <div style={{ display: 'flex', alignItems: 'center', height: '100%' }}>
            <MediaQuery largerThan="sm" styles={{ display: 'none' }}>   
              <Burger
                opened={opened}
                onClick={() => setOpened((o) => !o)}
                size="sm"
                color={theme.colors.gray[1]}
                mr="xl"
              />
            </MediaQuery>

          </div>
          <button onClick={rinkTop} >home</button>
          
          </Flex>
          <UserList/>
          </Header>
      }
    >
      <div className='props'>{children}</div>
    </AppShell>
  );
}