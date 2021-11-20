import { Card, Message } from "semantic-ui-react";
import useBlockchainEvents from "../hooks/useBlockchainEvents";
import BlockchainEventCard from "./BlockchainEventCard";

export function BlockchainEventPicker({ onClick }) {
  const [blockchainEvents] = useBlockchainEvents();
  return (
    <div>
      <h1>Create a pub-sub oracle</h1>
      <Message>Choose an event</Message>
      <Card.Group>
        {blockchainEvents.map((blockchainEvent) => (
          <BlockchainEventCard
            blockchainEvent={blockchainEvent}
            onClick={() => onClick(blockchainEvent.ID)}
          />
        ))}
      </Card.Group>
    </div>
  );
}
