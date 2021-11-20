import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import {
  Home,
  BlockchainEventDetail,
  OutboundOracleDetail,
  PubSubOracleDetail,
  PubSubOracleCreate,
  ConsumerDetail,
  SmartContractCreate,
  SmartContracts,
  Consumers,
  Settings,
} from "./pages";
import { Footer, Navbar } from "./components";
import { Container } from "semantic-ui-react";
import getHeaders from "./services/utils/getHeaders.js";
import Login from "./pages/Login";
import ProviderDetail from "./pages/ProviderDetail";
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

          <Route exact path="/outboundOracles/:outboundOracleID">
            <OutboundOracleDetail />
          </Route>

          <Route exact path="/blockchainEvents/:blockchainEventID">
            <BlockchainEventDetail />
          </Route>

          <Route exact path="/providers/:providerID">
            <ProviderDetail />
          </Route>

          <Route exact path="/pubSubOracles/create">
            <PubSubOracleCreate />
          </Route>

          <Route exact path="/pubSubOracles/:pubSubOracleID">
            <PubSubOracleDetail />
          </Route>

          <Route exact path="/consumers">
            <Consumers />
          </Route>

          <Route exact path="/smartContracts">
            <SmartContracts />
          </Route>

          <Route exact path="/smartContracts/create">
            <SmartContractCreate />
          </Route>

          <Route exact path="/consumers/:consumerID">
            <ConsumerDetail />
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
