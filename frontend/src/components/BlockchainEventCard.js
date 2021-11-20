import { Card } from "semantic-ui-react";
import Identicon from "react-identicons";
import { useHistory } from "react-router-dom";
import { PlaceholderCard } from "./PlaceholderCard";

export default function BlockchainEventCard({ blockchainEvent, onClick }) {
  const history = useHistory();
  if (!blockchainEvent) {
    return <PlaceholderCard />;
  }

  return (
    <Card
      onClick={() => {
        if (onClick) {
          onClick();
        } else {
          history.push("/blockchainEvents/" + blockchainEvent.ID);
        }
      }}
    >
      <Card.Content>
        <div style={{ float: "right" }}>
          <Identicon
            string={blockchainEvent.SmartContract.ContractAddress}
            size={50}
          />
        </div>
        <Card.Header>{blockchainEvent.ListenerPublisher.Name}</Card.Header>
        <Card.Meta>
          {blockchainEvent.SmartContract.ContractAddressSynonym} -{" "}
          {blockchainEvent.SmartContract.EventName} -{" "}
          {blockchainEvent.SmartContract.BlockchainName}
        </Card.Meta>
      </Card.Content>
      <Card.Content extra>
        <Card.Description>
          {blockchainEvent.ListenerPublisher.Description}
        </Card.Description>
      </Card.Content>
    </Card>
  );
}
