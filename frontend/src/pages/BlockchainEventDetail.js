import React from "react";
import { useHistory, useParams, Link } from "react-router-dom";
import useSubscription from "../hooks/useSubscription";
import { Button } from "semantic-ui-react";
import { ExampleRequest, OutboundOracleTable } from "../components";
import SubscriptionCard from "../components/SubscriptionCard";
export default function SubscriptionDetail({}) {
  const { subscriptionID } = useParams();
  const history = useHistory();
  const [subscription] = useSubscription(subscriptionID);

  if (!subscription) {
    return "Loading...";
  }

  return (
    <div>
      <h1>Blockchain Event</h1>
      <SubscriptionCard subscription={subscription} />
      <ExampleRequest
        eventParameters={subscription.ListenerPublisher.EventParameters}
      />
      {/*<div>
        <h2>Active Oracles</h2>
        <Button
          primary
          basic
          content="Create Oracle"
          icon="plus"
          as={Link}
          to={"/outboundOracles/create?subscriptionID=" + subscriptionID}
        />
        {/*subscription.OutboundOracles.length > 0 ? (
          <OutboundOracleTable
            outboundOracles={subscription.OutboundOracles}
          />
        ) : (
          <div>No oracles created yet.</div>
        )
      </div>*/}
    </div>
  );
}
