import React from "react";
import { useParams, Link } from "react-router-dom";
import useWebServiceProvider from "../hooks/useWebServiceProvider";
import { Button } from "semantic-ui-react";
import { InboundSubscriptionTable } from "../components";
import WebServiceProviderCard from "../components/WebServiceProviderCard";
export default function WebServiceProviderDetail() {
  const { webServiceProviderID } = useParams();
  const [webServiceProvider] = useWebServiceProvider(webServiceProviderID);
  if (!webServiceProvider) {
    return "Loading...";
  }

  return (
    <div>
      <h1>Web Service Provider</h1>
      <WebServiceProviderCard webServiceProvider={webServiceProvider} />
      <div>
        <h2>Active Subscriptions</h2>
        <Button
          primary
          basic
          content="Create Subscription"
          icon="plus"
          as={Link}
          to={
            "/inboundSubscriptions/create?webServiceProviderID=" +
            webServiceProvider.ID
          }
        />
        {webServiceProvider.InboundSubscriptions.length > 0 ? (
          <InboundSubscriptionTable
            inboundSubscriptions={webServiceProvider.InboundSubscriptions}
          />
        ) : (
          <div>No subscriptions created yet.</div>
        )}
      </div>
    </div>
  );
}
