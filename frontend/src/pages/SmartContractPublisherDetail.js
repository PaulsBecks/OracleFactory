import React from "react";
import { useParams, Link } from "react-router-dom";
import useSmartContractPublisher from "../hooks/useSmartContractPublisher";
import { Button } from "semantic-ui-react";
import { ExampleRequest, InboundOracleTable } from "../components";
import SmartContractPublisherCard from "../components/SmartContractPublisherCard";

export default function SmartContractPublisherDetail() {
  const { smartContractPublisherID } = useParams();
  const [smartContractPublisher] = useSmartContractPublisher(
    smartContractPublisherID
  );

  if (!smartContractPublisher) {
    return "Loading...";
  }

  return (
    <div>
      <div style={{ display: "flex" }}>
        <SmartContractPublisherCard
          smartContractPublisher={smartContractPublisher}
        />
        <div style={{ marginLeft: "2em" }}>
          <ExampleRequest
            eventParameters={
              smartContractPublisher.ListenerPublisher.EventParameters
            }
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
          to={
            "/inboundOracles/create?smartContractPublisherID=" +
            smartContractPublisherID
          }
        />
        {smartContractPublisher.InboundOracles.length > 0 ? (
          <InboundOracleTable
            inboundOracles={smartContractPublisher.InboundOracles}
          />
        ) : (
          <div>No oracles created yet.</div>
        )}
      </div>
    </div>
  );
}
