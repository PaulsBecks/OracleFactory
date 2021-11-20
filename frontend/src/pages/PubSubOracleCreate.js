import React, { useState } from "react";
import { useLocation, useParams } from "react-router";
import postData from "../services/postData";
import { PubSubOracleForm } from "../components";
import { Button, Card, Message } from "semantic-ui-react";
import { Link, useHistory } from "react-router-dom";
import useProviders from "../hooks/useProviders";
import ProviderCard from "../components/ProviderCard";
import ConsumerCard from "../components/ConsumerCard";
import useConsumers from "../hooks/useConsumers";
import useBlockchainEvents from "../hooks/useBlockchainEvents";
import { ProviderPicker } from "../components/ProviderPicker";
import { ConsumerPicker } from "../components/ConsumerPicker";
import { BlockchainEventPicker } from "../components/BlockchainEventPicker";

function useQuery() {
  return new URLSearchParams(useLocation().search);
}

export default function PubSubOracleCreate() {
  const history = useHistory();
  const query = useQuery();
  const [pubSubOracle, setPubSubOracle] = useState({
    Oracle: { Name: "" },
    URI: "",
    consumerID: parseInt(query.get("consumerID")),
    providerID: parseInt(query.get("providerID")),
    subBlockchainEventID: null,
    unsubBlockchainEventID: null,
  });
  const [loading, setLoading] = useState(false);
  const updatePubSubOracle = (name, value) => {
    setPubSubOracle({
      ...pubSubOracle,
      [name]: value,
    });
  };

  if (!pubSubOracle.providerID) {
    return (
      <ProviderPicker
        onClick={(providerID) => {
          updatePubSubOracle("providerID", providerID);
        }}
      />
    );
  }

  if (!pubSubOracle.consumerID) {
    return (
      <ConsumerPicker
        onClick={(consumerID) => updatePubSubOracle("consumerID", consumerID)}
      />
    );
  }

  if (!pubSubOracle.subBlockchainEventID) {
    return (
      <BlockchainEventPicker
        onClick={(blockchainEventID) => {
          console.log("lol");
          updatePubSubOracle("subBlockchainEventID", blockchainEventID);
        }}
      />
    );
  }

  if (!pubSubOracle.unsubBlockchainEventID) {
    return (
      <BlockchainEventPicker
        onClick={(blockchainEventID) =>
          updatePubSubOracle("unsubBlockchainEventID", blockchainEventID)
        }
      />
    );
  }

  return (
    <div>
      <h1>Create Pub-Sub Oracle</h1>
      <PubSubOracleForm
        pubSubOracle={pubSubOracle}
        setPubSubOracle={setPubSubOracle}
      />
      <br />
      <Button
        loading={loading}
        basic
        negative
        content="Cancel"
        as={Link}
        to={"/consumers/" + pubSubOracle.consumerID}
      />
      <Button
        loading={loading}
        basic
        positive
        content="Create"
        onClick={async () => {
          setLoading(true);
          await postData(`/pubSubOracles`, {
            ...pubSubOracle,
          });
          setLoading(false);
          history.push("/");
        }}
      />
    </div>
  );
}
