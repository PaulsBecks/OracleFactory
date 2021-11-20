import { Card, Placeholder } from "semantic-ui-react";
import Identicon from "react-identicons";
import { useHistory } from "react-router-dom";
import { PlaceholderCard } from "./PlaceholderCard";

export default function PubSubOracleCard({ pubSubOracle, onClick }) {
  const history = useHistory();
  if (!pubSubOracle) {
    return <PlaceholderCard />;
  }

  return (
    <Card
      onClick={() => {
        if (onClick) {
          onClick();
        } else {
          history.push("/pubSubOracles/" + pubSubOracle.ID);
        }
      }}
    >
      <Card.Content>
        <Card.Header>{pubSubOracle.Oracle.Name}</Card.Header>
        <Card.Meta>Pub-Sub Oracle</Card.Meta>
      </Card.Content>
    </Card>
  );
}
