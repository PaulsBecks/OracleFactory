import useSmartContractProviders from "../hooks/useSmartContractProviders";
import useSmartContractConsumers from "../hooks/useSmartContractConsumers";
import { Link } from "react-router-dom";
import { Button, Card, Message } from "semantic-ui-react";
import SmartContractProviderCard from "../components/SmartContractProviderCard";
import SmartContractConsumerCard from "../components/SmartContractConsumerCard";
import useWebServiceProviders from "../hooks/useWebServiceProviders";
import WebServiceProviderCard from "../components/WebServiceProviderCard";
import WebServiceConsumerCard from "../components/WebServiceConsumerCard";
import useWebServiceConsumers from "../hooks/useWebServiceConsumers";

function SmartContracts() {
  const [smartContractProviders] = useSmartContractProviders();
  const [smartContractConsumers] = useSmartContractConsumers();
  const [webServiceProviders] = useWebServiceProviders();
  const [webServiceConsumers] = useWebServiceConsumers();
  return (
    <div>
      <div>
        <Button
          basic
          primary
          icon="plus"
          content="Create Consumer or Provider"
          as={Link}
          to="/smartContracts/create"
        />
        <Button
          basic
          primary
          icon="plus"
          content="Create Outbound Subscription"
          as={Link}
          to="/outboundSubscriptions/create"
        />
        <Button
          basic
          primary
          icon="plus"
          content="Create Inbound Subscription"
          as={Link}
          to="/inboundSubscriptions/create"
        />
      </div>
      <br />
      <Message info size="huge">
        {" "}
        Click on a provider or a consumer you want to create an subscription
        with.
      </Message>
      <div>
        <h2>Smart Contract Provider</h2>
        <Card.Group>
          {smartContractProviders.map((smartContractProvider) => (
            <SmartContractProviderCard
              smartContractProvider={smartContractProvider}
            />
          ))}
          {(!webServiceProviders || webServiceProviders.length === 0) &&
            (!smartContractProviders ||
              smartContractProviders.length === 0) && (
              <Message info>
                No providers available. You have to{" "}
                <Link to="/smartContracts/create">create them first here</Link>!
              </Message>
            )}
        </Card.Group>

        <br />
        <br />
        <h2>Web Service Provider</h2>
        <Card.Group>
          {webServiceProviders.map((webServiceProvider) => (
            <WebServiceProviderCard webServiceProvider={webServiceProvider} />
          ))}
        </Card.Group>
        {(!webServiceProviders || webServiceProviders.length === 0) && (
          <Message info>
            No providers available. You have to{" "}
            <Link to="/smartContracts/create">create them first here</Link>!
          </Message>
        )}
      </div>
      <br />
      <br />
      <div>
        <h2>Smart Contract Consumer</h2>
        <Card.Group>
          {smartContractConsumers.map((smartContractConsumer) => (
            <SmartContractConsumerCard
              smartContractConsumer={smartContractConsumer}
            />
          ))}
        </Card.Group>
        {(!smartContractConsumers || smartContractConsumers.length === 0) && (
          <Message info>
            No smart contract consumers available. You have to{" "}
            <Link to="/smartContracts/create">create them first here</Link>!
          </Message>
        )}
        <br />
        <br />
        <h2>Web Service Consumer</h2>
        <Card.Group>
          {webServiceConsumers.map((webServiceConsumer) => (
            <WebServiceConsumerCard webServiceConsumer={webServiceConsumer} />
          ))}
          {(!webServiceConsumers || webServiceConsumers.length === 0) && (
            <Message info>
              No service consumers available. You have to{" "}
              <Link to="/smartContracts/create">create them first here</Link>!
            </Message>
          )}
        </Card.Group>
      </div>
    </div>
  );
}

export default SmartContracts;
