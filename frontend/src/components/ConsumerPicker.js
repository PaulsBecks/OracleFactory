import { Card, Message } from "semantic-ui-react";
import useConsumers from "../hooks/useConsumers";
import ConsumerCard from "./ConsumerCard";

export function ConsumerPicker({ onClick }) {
  const [consumers] = useConsumers();
  return (
    <div>
      <h1>Create a pub-sub oracle</h1>
      <Message>Choose a consumer</Message>
      <Card.Group>
        {consumers.map((consumer) => (
          <ConsumerCard
            consumer={consumer}
            onClick={() => onClick(consumer.ID)}
          />
        ))}
      </Card.Group>
    </div>
  );
}
