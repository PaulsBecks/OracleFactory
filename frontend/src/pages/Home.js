import { Link } from "react-router-dom";
import { Button } from "semantic-ui-react";
import EthereumBlockchainConnectorsTable from "../components/EthereumBlockchainConnectorsTable";
import HyperledgerBlockchainConnectorsTable from "../components/HyperledgerBlockchainConnectorsTable";
import ProvidersTable from "../components/ProvidersTable";
import SubscriptionsTable from "../components/SubscriptionsTable";
import useEthereumBlockchainConnectors from "../hooks/useEthereumBlockchainConnectors";
import useHyperledgerBlockchainConnectors from "../hooks/useHyperledgerBlockchainConnectors";
import useProviders from "../hooks/useProviders";
import useSubscriptions from "../hooks/useSubscriptions";

function Home() {
  const [ethereumBlockchainConnectors] = useEthereumBlockchainConnectors();
  const [hyperledgerBlockchainConnectors] =
    useHyperledgerBlockchainConnectors();
  const [providers] = useProviders();
  const [subscriptions] = useSubscriptions();
  return (
    <div>
      <div style={{ display: "flex" }}>
        <Button
          basic
          primary
          icon="plus"
          content="Create Ethereum Connector"
          as={Link}
          to="/blockchainConnectors/ethereum/create"
        />
        <Button
          basic
          primary
          icon="plus"
          content="Create Hyperledger Connector"
          as={Link}
          to="/blockchainConnectors/hyperledger/create"
        />
        <Button
          basic
          primary
          icon="plus"
          content="Create Provider"
          as={Link}
          to="/providers/create"
        />
      </div>
      <br />
      <div>
        <h2>Ethereum Connectors</h2>
        <EthereumBlockchainConnectorsTable
          ethereumBlockchainConnectors={ethereumBlockchainConnectors}
        />
      </div>
      <br />
      <div>
        <h2>Hyperledger Connectors</h2>
        <HyperledgerBlockchainConnectorsTable
          hyperledgerBlockchainConnectors={hyperledgerBlockchainConnectors}
        />
      </div>
      <br />

      <div>
        <h2>Providers</h2>
        <ProvidersTable providers={providers} />
      </div>
      <br />
      <div>
        <h2>Off-Chain Subscriptions</h2>
        <SubscriptionsTable subscriptions={subscriptions} />
      </div>
    </div>
  );
}

export default Home;
