import React from "react";
import { useParams, Link } from "react-router-dom";
import useSmartContractConsumer from "../hooks/useSmartContractConsumer";
import { Button } from "semantic-ui-react";
import { ExampleRequest, InboundSubscriptionTable } from "../components";
import SmartContractConsumerCard from "../components/SmartContractConsumerCard";

export default function SmartContractConsumerDetail() {
  const { smartContractConsumerID } = useParams();
  const [smartContractConsumer] = useSmartContractConsumer(
    smartContractConsumerID
  );

  if (!smartContractConsumer) {
    return "Loading...";
  }

  return (
    <div>
      <div style={{ display: "flex" }}>
        <SmartContractConsumerCard
          smartContractConsumer={smartContractConsumer}
        />
        <div style={{ marginLeft: "2em" }}>
          <ExampleRequest
            eventParameters={
              smartContractConsumer.ProviderConsumer.EventParameters
            }
          />
        </div>
      </div>
      <br />
      <div>
        <h2>Active Subscriptions</h2>
        <Button
          basic
          primary
          content="Create Subscription"
          icon="plus"
          as={Link}
          to={
            "/inboundSubscriptions/create?smartContractConsumerID=" +
            smartContractConsumerID
          }
        />
        {smartContractConsumer.InboundSubscriptions.length > 0 ? (
          <InboundSubscriptionTable
            inboundSubscriptions={smartContractConsumer.InboundSubscriptions}
          />
        ) : (
          <div>No subscriptions created yet.</div>
        )}
      </div>
    </div>
  );
}
