import { Card } from "semantic-ui-react";
import Identicon from "react-identicons";
import { useHistory } from "react-router-dom";
import config from "../config";
import { PlaceholderCard } from "./PlaceholderCard";

export default function ProviderCard({ provider, onClick }) {
  const history = useHistory();
  if (!provider) {
    return <PlaceholderCard />;
  }

  return (
    <Card
      onClick={() => {
        if (onClick) {
          onClick();
        } else {
          history.push("/providers/" + provider.ID);
        }
      }}
    >
      <Card.Content>
        <div style={{ float: "right" }}>
          <Identicon string={"Listener" + provider.ID} size={50} />
        </div>
        <Card.Header>{provider.ListenerPublisher.Name}</Card.Header>
        <Card.Meta>Provider</Card.Meta>
      </Card.Content>
      <Card.Content extra>
        <Card.Description>
          {provider.ListenerPublisher.Description}
        </Card.Description>
      </Card.Content>
    </Card>
  );
}
