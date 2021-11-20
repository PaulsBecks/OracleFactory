import useBlockchainEvents from "../hooks/useBlockchainEvents";
import useConsumers from "../hooks/useConsumers";
import { Link } from "react-router-dom";
import { Button, Card } from "semantic-ui-react";
import BlockchainEventCard from "../components/BlockchainEventCard";
import ConsumerCard from "../components/ConsumerCard";
import useProviders from "../hooks/useProviders";
import ProviderCard from "../components/ProviderCard";

function SmartContracts() {
  const [blockchainEvents] = useBlockchainEvents();
  const [consumers] = useConsumers();
  const [providers] = useProviders();
  return (
    <div>
      <div>
        <Button
          basic
          primary
          icon="plus"
          content="Create Publisher or Listener"
          as={Link}
          to="/smartContracts/create"
        />
      </div>
      <br />
      <div>
        <h2>Subsribe/Unsubscribe Events</h2>
        <Card.Group>
          {blockchainEvents.map((blockchainEvent) => (
            <BlockchainEventCard blockchainEvent={blockchainEvent} />
          ))}
        </Card.Group>
      </div>
      <br />
      <br />
      <div>
        <h2>Provider</h2>
        <Card.Group>
          {providers.map((provider) => (
            <ProviderCard provider={provider} />
          ))}
        </Card.Group>
      </div>
      <br />
      <br />
      <div>
        <h2>Smart Contract Consumer</h2>
        <Card.Group>
          {consumers.map((consumer) => (
            <ConsumerCard consumer={consumer} />
          ))}
        </Card.Group>
      </div>
    </div>
  );
}

export default SmartContracts;
