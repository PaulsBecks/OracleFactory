import React from "react";
import { useHistory, useParams, Link } from "react-router-dom";
import useBlockchainEvent from "../hooks/useBlockchainEvent";
import { Button } from "semantic-ui-react";
import { ExampleRequest, OutboundOracleTable } from "../components";
import BlockchainEventCard from "../components/BlockchainEventCard";
export default function BlockchainEventDetail({}) {
  const { blockchainEventID } = useParams();
  const history = useHistory();
  const [blockchainEvent] = useBlockchainEvent(blockchainEventID);

  if (!blockchainEvent) {
    return "Loading...";
  }

  return (
    <div>
      <h1>Blockchain Event</h1>
      <BlockchainEventCard blockchainEvent={blockchainEvent} />
      <ExampleRequest
        eventParameters={blockchainEvent.ListenerPublisher.EventParameters}
      />
      {/*<div>
        <h2>Active Oracles</h2>
        <Button
          primary
          basic
          content="Create Oracle"
          icon="plus"
          as={Link}
          to={"/outboundOracles/create?blockchainEventID=" + blockchainEventID}
        />
        {/*blockchainEvent.OutboundOracles.length > 0 ? (
          <OutboundOracleTable
            outboundOracles={blockchainEvent.OutboundOracles}
          />
        ) : (
          <div>No oracles created yet.</div>
        )
      </div>*/}
    </div>
  );
}
