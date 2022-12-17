import { useState } from 'react';
import { Tabs } from '@mantine/core';
import {ThankTo} from "./ToUser"
import { ThankFrom } from './FromUser';
import { EditProfile } from './EditProfile';

export const Tab =()=> {
  // console.log("Tab");
  const [activeTab, setActiveTab] = useState<string | null>('first');

  return (
  <div >
    <Tabs value={activeTab} onTabChange={setActiveTab}>
      <Tabs.List>
        <Tabs.Tab value="first">Thanks Sent</Tabs.Tab>
        <Tabs.Tab value="second">Thanks Got</Tabs.Tab>
        <Tabs.Tab value="third">Edit Profile</Tabs.Tab>
      </Tabs.List>
      <Tabs.Panel value="first">{ThankFrom()}</Tabs.Panel>
      <Tabs.Panel value="second">{ThankTo()}</Tabs.Panel>
      <Tabs.Panel value="third">{EditProfile()}</Tabs.Panel>
    </Tabs>
    </div>
  );
}