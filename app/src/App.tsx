import { BrowserRouter } from "react-router-dom";
import {Router} from "./Router"
import { UserProvider } from "./Components/Shared/Context";
import { ActiveProvider } from "./Components/Shared/ActiveProvider";
const App = ()=> {
  // console.log("App")
return (
  <UserProvider>
    <ActiveProvider>
  <div className="App">
    <header className="App-header">
      <> 

      {/* <BrowserRouter> */}
      <div className="App">

      <Router />
      </div>
    {/* </BrowserRouter> */}
      </>
     
    </header>
  </div> 
  </ActiveProvider>
  </UserProvider>
); }
export default App;
