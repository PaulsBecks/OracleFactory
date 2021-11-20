import { Link } from "react-router-dom";
import { Button, Card, Message } from "semantic-ui-react";
import useProviders from "../hooks/useProviders";
import ProviderCard from "./ProviderCard";

export function ProviderPicker({ onClick }) {
  const [providers] = useProviders();
  return (
    <div>
      <h1>Create Pub-Sub Oracle</h1>
      <Button content="Create Provider" as={Link} to="/smartContracts/create" />
      <Message>Choose a data provider</Message>
      <Card.Group>
        {providers.map((provider) => (
          <ProviderCard
            provider={provider}
            onClick={() => onClick(provider.ID)}
          />
        ))}
      </Card.Group>
    </div>
  );
}
