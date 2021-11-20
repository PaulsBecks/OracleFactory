import { Card } from "semantic-ui-react";
import Identicon from "react-identicons";
import { useHistory } from "react-router-dom";
import { PlaceholderCard } from "./PlaceholderCard";

export default function OutboundOracleCard({ outboundOracle, onClick }) {
  const history = useHistory();
  if (!outboundOracle) {
    return <PlaceholderCard />;
  }

  return (
    <Card
      onClick={() => {
        if (onClick) {
          onClick();
        } else {
          history.push("/outboundOracles/" + outboundOracle.ID);
        }
      }}
    >
      <Card.Content>
        <div style={{ float: "right" }}>
          <Identicon
            string={
              outboundOracle.BlockchainEvent.SmartContract.ContractAddress
            }
            size={50}
          />
        </div>
        <Card.Header>
          {outboundOracle.BlockchainEvent.ListenerPublisher.Name}
        </Card.Header>
        <Card.Meta>
          {outboundOracle.IsSubscribing
            ? "Subscribtion Oracle"
            : "Unsubscription Oracle"}{" "}
          - {outboundOracle.BlockchainEvent.SmartContract.EventName} -{" "}
          {outboundOracle.BlockchainEvent.SmartContract.BlockchainName}
        </Card.Meta>
      </Card.Content>
      <Card.Content extra>
        <Card.Description>
          {outboundOracle.BlockchainEvent.ListenerPublisher.Description}
        </Card.Description>
      </Card.Content>
    </Card>
  );
}
