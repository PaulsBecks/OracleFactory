import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import {
  Home,
  SmartContractListenerDetail,
  OutboundOracleDetail,
  OutboundOracleCreate,
  InboundOracleDetail,
  InboundOracleCreate,
  SmartContractPublisherDetail,
  SmartContractCreate,
  SmartContracts,
  SmartContractPublishers,
  Settings,
} from "./pages";
import { Footer, Navbar } from "./components";
import { Container } from "semantic-ui-react";
import getHeaders from "./services/utils/getHeaders.js";
import Login from "./pages/Login";
import WebServiceListenerDetail from "./pages/WebServiceListenerDetail";
import WebServicePublisherDetail from "./pages/WebServicePublisherDetail";
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

          <Route exact path="/outboundOracles/create">
            <OutboundOracleCreate />
          </Route>

          <Route exact path="/outboundOracles/:outboundOracleID">
            <OutboundOracleDetail />
          </Route>

          <Route exact path="/smartContractListeners/:smartContractListenerID">
            <SmartContractListenerDetail />
          </Route>

          <Route exact path="/webServiceListeners/:webServiceListenerID">
            <WebServiceListenerDetail />
          </Route>

          <Route exact path="/webServicePublishers/:webServicePublisherID">
            <WebServicePublisherDetail />
          </Route>

          <Route exact path="/inboundOracles/create">
            <InboundOracleCreate />
          </Route>

          <Route exact path="/inboundOracles/:inboundOracleID">
            <InboundOracleDetail />
          </Route>

          <Route exact path="/smartContractPublishers">
            <SmartContractPublishers />
          </Route>

          <Route exact path="/smartContracts">
            <SmartContracts />
          </Route>

          <Route exact path="/smartContracts/create">
            <SmartContractCreate />
          </Route>

          <Route
            exact
            path="/smartContractPublishers/:smartContractPublisherID"
          >
            <SmartContractPublisherDetail />
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
