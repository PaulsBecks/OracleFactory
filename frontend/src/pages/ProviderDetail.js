import React from "react";
import { useParams, Link } from "react-router-dom";
import useProvider from "../hooks/useProvider";
import { Button } from "semantic-ui-react";
import ProviderCard from "../components/ProviderCard";
export default function ProviderDetail() {
  const { providerID } = useParams();
  const [provider] = useProvider(providerID);
  if (!provider) {
    return "Loading...";
  }

  return (
    <div>
      <h1>Provider</h1>
      <ProviderCard provider={provider} />
      <div>
        <h2>Active Oracles</h2>
        <Button
          primary
          basic
          content="Create Oracle"
          icon="plus"
          as={Link}
          to={"/pubSubOracles/create?providerID=" + provider.ID}
        />
      </div>
    </div>
  );
}
