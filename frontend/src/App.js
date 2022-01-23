import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import {
  Home,
  SmartContractProviderDetail,
  OutboundSubscriptionDetail,
  OutboundSubscriptionCreate,
  InboundSubscriptionDetail,
  InboundSubscriptionCreate,
  SmartContractConsumerDetail,
  SmartContractCreate,
  SmartContracts,
  SmartContractConsumers,
  Settings,
} from "./pages";
import { Footer, Navbar } from "./components";
import { Container } from "semantic-ui-react";
import getHeaders from "./services/utils/getHeaders.js";
import Login from "./pages/Login";
import WebServiceProviderDetail from "./pages/WebServiceProviderDetail";
import WebServiceConsumerDetail from "./pages/WebServiceConsumerDetail";
function App() {
  const loggedIn = getHeaders();
  if (!loggedIn) {
    return <Login />;
  }

  return (
    <Router>
      <Navbar />
      <Container>
        <Switch>
          <Route exact path="/">
            <Home />
          </Route>

          <Route exact path="/outboundSubscriptions/create">
            <OutboundSubscriptionCreate />
          </Route>

          <Route exact path="/outboundSubscriptions/:outboundSubscriptionID">
            <OutboundSubscriptionDetail />
          </Route>

          <Route exact path="/smartContractProviders/:smartContractProviderID">
            <SmartContractProviderDetail />
          </Route>

          <Route exact path="/webServiceProviders/:webServiceProviderID">
            <WebServiceProviderDetail />
          </Route>

          <Route exact path="/webServiceConsumers/:webServiceConsumerID">
            <WebServiceConsumerDetail />
          </Route>

          <Route exact path="/inboundSubscriptions/create">
            <InboundSubscriptionCreate />
          </Route>

          <Route exact path="/inboundSubscriptions/:inboundSubscriptionID">
            <InboundSubscriptionDetail />
          </Route>

          <Route exact path="/smartContractConsumers">
            <SmartContractConsumers />
          </Route>

          <Route exact path="/smartContracts">
            <SmartContracts />
          </Route>

          <Route exact path="/smartContracts/create">
            <SmartContractCreate />
          </Route>

          <Route exact path="/smartContractConsumers/:smartContractConsumerID">
            <SmartContractConsumerDetail />
          </Route>

          <Route exact path="/settings">
            <Settings />
          </Route>
        </Switch>
      </Container>
      <Footer />
    </Router>
  );
}

export default App;
