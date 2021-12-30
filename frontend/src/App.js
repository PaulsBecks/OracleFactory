import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import { Home } from "./pages";
import { Footer, Navbar } from "./components";
import { Container } from "semantic-ui-react";
import getHeaders from "./services/utils/getHeaders.js";
import Login from "./pages/Login";
import ProviderDetail from "./pages/ProviderDetail";
import { ProviderCreate } from "./pages/ProviderCreate";
import { HyperledgerBlockchainConnectorDetail } from "./pages/HyperledgerBlockchainConnectorDetail";
import { EthereumBlockchainConnectorDetail } from "./pages/EthereumBlockchainConnectorDetail";
import { EthereumBlockchainConnectorCreate } from "./pages/EthereumBlockchainConnectorCreate";
import { HyperledgerBlockchainConnectorCreate } from "./pages/HyperledgerBlockchainConnectorCreate";
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

          <Route exact path="/blockchainConnectors/ethereum/create">
            <EthereumBlockchainConnectorCreate />
          </Route>

          <Route exact path="/blockchainConnectors/hyperledger/create">
            <HyperledgerBlockchainConnectorCreate />
          </Route>

          <Route
            exact
            path="/blockchainConnectors/ethereum/:ethereumBlockchainConnectorID"
          >
            <EthereumBlockchainConnectorDetail />
          </Route>

          <Route
            exact
            path="/blockchainConnectors/hyperledger/:hyperledgerBlockchainConnectorID"
          >
            <HyperledgerBlockchainConnectorDetail />
          </Route>

          <Route exact path="/providers/create">
            <ProviderCreate />
          </Route>

          <Route exact path="/providers/:providerID">
            <ProviderDetail />
          </Route>
        </Switch>
      </Container>
      <Footer />
    </Router>
  );
}

export default App;
