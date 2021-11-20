import React from "react";
import { useParams, Link } from "react-router-dom";
import useConsumer from "../hooks/useConsumer";
import { Button } from "semantic-ui-react";
import { ExampleRequest, PubSubOracleTable } from "../components";
import ConsumerCard from "../components/ConsumerCard";

export default function ConsumerDetail() {
  const { consumerID } = useParams();
  const [consumer] = useConsumer(consumerID);

  if (!consumer) {
    return "Loading...";
  }

  return (
    <div>
      <div style={{ display: "flex" }}>
        <ConsumerCard consumer={consumer} />
        <div style={{ marginLeft: "2em" }}>
          <ExampleRequest
            eventParameters={consumer.ListenerPublisher.EventParameters}
          />
        </div>
      </div>
      <br />
      <div>
        <h2>Active Oracles</h2>
        <Button
          basic
          primary
          content="Create Oracle"
          icon="plus"
          as={Link}
          to={"/pubSubOracles/create?consumerID=" + consumerID}
        />
        {consumer.PubSubOracles.length > 0 ? (
          <PubSubOracleTable pubSubOracles={consumer.PubSubOracles} />
        ) : (
          <div>No oracles created yet.</div>
        )}
      </div>
    </div>
  );
}
