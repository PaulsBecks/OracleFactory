import React from "react";
import { useHistory, useParams, Link } from "react-router-dom";
import useSmartContractProvider from "../hooks/useSmartContractProvider";
import { Button } from "semantic-ui-react";
import { OutboundSubscriptionTable } from "../components";
import SmartContractProviderCard from "../components/SmartContractProviderCard";
export default function SmartContractProviderDetail({}) {
  const { smartContractProviderID } = useParams();
  const history = useHistory();
  const [smartContractProvider] = useSmartContractProvider(
    smartContractProviderID
  );

  if (!smartContractProvider) {
    return "Loading...";
  }

  return (
    <div>
      <h1>Smart Contract Provider</h1>
      <SmartContractProviderCard
        smartContractProvider={smartContractProvider}
      />
      <div>
        <h2>Active Subscriptions</h2>
        <Button
          primary
          basic
          content="Create Subscription"
          icon="plus"
          as={Link}
          to={
            "/outboundSubscriptions/create?smartContractProviderID=" +
            smartContractProviderID
          }
        />
        {smartContractProvider.OutboundSubscriptions.length > 0 ? (
          <OutboundSubscriptionTable
            outboundSubscriptions={smartContractProvider.OutboundSubscriptions}
          />
        ) : (
          <div>No subscriptions created yet.</div>
        )}
      </div>
    </div>
  );
}
