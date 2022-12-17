import { Route, Switch, BrowserRouter } from "react-router-dom";
import { Home } from "./Pages/Home";
import { User } from "./Pages/User";
import { ThankPage } from "./Pages/ThankPage";
import { MyPage } from "./Pages/MyPage";

export const Router = () => {
  // console.log("Router")
  return (
      
      <BrowserRouter>
      <Switch>
      <Route exact path="/">
        <Home />
      </Route>
      <Route path="/ranking">
        <User />
      </Route>
      <Route path="/mypage">
        <MyPage/>
      </Route>
      <Route path="/thank">
        <ThankPage/>
      </Route>
      </Switch>
      </BrowserRouter>

  );
};
