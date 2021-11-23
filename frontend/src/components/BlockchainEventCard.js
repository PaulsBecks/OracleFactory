import { Card } from "semantic-ui-react";
import Identicon from "react-identicons";
import { useHistory } from "react-router-dom";
import { PlaceholderCard } from "./PlaceholderCard";

export default function SubscriptionCard({ subscription, onClick }) {
  const history = useHistory();
  if (!subscription) {
    return <PlaceholderCard />;
  }

  return (
    <Card
      onClick={() => {
        if (onClick) {
          onClick();
        } else {
          history.push("/subscriptions/" + subscription.ID);
        }
      }}
    >
      <Card.Content>
        <div style={{ float: "right" }}>
          <Identicon
            string={subscription.SmartContract.ContractAddress}
            size={50}
          />
        </div>
        <Card.Header>{subscription.ListenerPublisher.Name}</Card.Header>
        <Card.Meta>
          {subscription.SmartContract.ContractAddressSynonym} -{" "}
          {subscription.SmartContract.EventName} -{" "}
          {subscription.SmartContract.BlockchainName}
        </Card.Meta>
      </Card.Content>
      <Card.Content extra>
        <Card.Description>
          {subscription.ListenerPublisher.Description}
        </Card.Description>
      </Card.Content>
    </Card>
  );
}
