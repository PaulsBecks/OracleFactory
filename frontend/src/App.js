import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import {
  Home,
  OutboundOracleTemplateDetail,
  OutboundOracleDetail,
  OutboundOracleCreate,
  InboundOracleDetail,
  InboundOracleCreate,
  InboundOracleTemplateDetail,
  OracleTemplateCreate,
  OracleTemplates,
  InboundOracleTemplates,
  Settings,
} from "./pages";
import { Navbar } from "./components";
import { Container } from "semantic-ui-react";
import getHeaders from "./services/utils/getHeaders.js";
import Login from "./pages/Login";
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

          <Route
            exact
            path="/outboundOracleTemplates/:outboundOracleTemplateID"
          >
            <OutboundOracleTemplateDetail />
          </Route>

          <Route
            exact
            path="/outboundOracleTemplates/:outboundOracleTemplateID/outboundOracles/create"
          >
            <OutboundOracleCreate />
          </Route>

          <Route exact path="/inboundOracles/:inboundOracleID">
            <InboundOracleDetail />
          </Route>

          <Route exact path="/inboundOracleTemplates">
            <InboundOracleTemplates />
          </Route>

          <Route exact path="/oracleTemplates">
            <OracleTemplates />
          </Route>

          <Route exact path="/oracleTemplates/create">
            <OracleTemplateCreate />
          </Route>

          <Route exact path="/inboundOracleTemplates/:inboundOracleTemplateID">
            <InboundOracleTemplateDetail />
          </Route>

          <Route
            exact
            path="/inboundOracleTemplates/:inboundOracleTemplateID/inboundOracles/create"
          >
            <InboundOracleCreate />
          </Route>

          <Route exact path="/settings">
            <Settings />
          </Route>
        </Switch>
      </Container>
    </Router>
  );
}

export default App;
