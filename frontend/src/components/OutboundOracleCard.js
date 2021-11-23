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
            string={outboundOracle.Subscription.SmartContract.ContractAddress}
            size={50}
          />
        </div>
        <Card.Header>
          {outboundOracle.Subscription.ListenerPublisher.Name}
        </Card.Header>
        <Card.Meta>
          {outboundOracle.IsSubscribing
            ? "Subscribtion Oracle"
            : "Unsubscription Oracle"}{" "}
          - {outboundOracle.Subscription.SmartContract.EventName} -{" "}
          {outboundOracle.Subscription.SmartContract.BlockchainName}
        </Card.Meta>
      </Card.Content>
      <Card.Content extra>
        <Card.Description>
          {outboundOracle.Subscription.ListenerPublisher.Description}
        </Card.Description>
      </Card.Content>
    </Card>
  );
}
