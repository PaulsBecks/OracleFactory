import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import {
  Home,
  OutboundOracleTemplateDetail,
  OutboundOracleDetail,
  OutboundOracleCreate,
  InboundOracleDetail,
  InboundOracleCreate,
  InboundOracleTemplateDetail,
} from "./pages";
import { Navbar } from "./components";
import { Container } from "semantic-ui-react";

function App() {
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
          <Route exact path="/inboundOracleTemplates/:inboundOracleTemplateID">
            <InboundOracleTemplateDetail />
          </Route>
          <Route
            exact
            path="/inboundOracleTemplates/:inboundOracleTemplateID/inboundOracles/create"
          >
            <InboundOracleCreate />
          </Route>
        </Switch>
      </Container>
    </Router>
  );
}

export default App;
