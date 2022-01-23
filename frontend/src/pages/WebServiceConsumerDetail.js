import React from "react";
import { useParams, Link } from "react-router-dom";
import useWebServiceConsumer from "../hooks/useWebServiceConsumer";
import { Button } from "semantic-ui-react";
import { OutboundSubscriptionTable } from "../components";
import WebServiceConsumerCard from "../components/WebServiceConsumerCard";
export default function WebServiceConsumerDetail() {
  const { webServiceConsumerID } = useParams();
  const [webServiceConsumer] = useWebServiceConsumer(webServiceConsumerID);
  if (!webServiceConsumer) {
    return "Loading...";
  }

  console.log(webServiceConsumer.OutboundSubscriptions);

  return (
    <div>
      <h1>Web Service Consumer</h1>
      <WebServiceConsumerCard webServiceConsumer={webServiceConsumer} />
      <div>
        <h2>Active Subscriptions</h2>
        <Button
          primary
          basic
          content="Create Subscription"
          icon="plus"
          as={Link}
          to={
            "/outboundSubscriptions/create?webServiceConsumerID=" +
            webServiceConsumer.ID
          }
        />
        {webServiceConsumer.OutboundSubscriptions.length > 0 ? (
          <OutboundSubscriptionTable
            outcdboundSubscriptions={webServiceConsumer.OutboundSubscriptions}
          />
        ) : (
          <div>No subscriptions created yet.</div>
        )}
      </div>
    </div>
  );
}
