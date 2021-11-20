import React from "react";
import { useHistory, useParams, Link } from "react-router-dom";
import useBlockchainEvent from "../hooks/useBlockchainEvent";
import { Button } from "semantic-ui-react";
import { OutboundOracleTable } from "../components";
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
      <h1>Smart Contract Listener</h1>
      <BlockchainEventCard blockchainEvent={blockchainEvent} />
      <div>
        <h2>Active Oracles</h2>
        <Button
          primary
          basic
          content="Create Oracle"
          icon="plus"
          as={Link}
          to={"/outboundOracles/create?blockchainEventID=" + blockchainEventID}
        />
        {blockchainEvent.OutboundOracles.length > 0 ? (
          <OutboundOracleTable
            outboundOracles={blockchainEvent.OutboundOracles}
          />
        ) : (
          <div>No oracles created yet.</div>
        )}
      </div>
    </div>
  );
}
