import { Card, Message } from "semantic-ui-react";
import useSubscriptions from "../hooks/useSubscriptions";
import SubscriptionCard from "./SubscriptionCard";

export function SubscriptionPicker({ onClick }) {
  const [subscriptions] = useSubscriptions();
  return (
    <div>
      <h1>Create a pub-sub oracle</h1>
      <Message>Choose an event</Message>
      <Card.Group>
        {subscriptions.map((subscription) => (
          <SubscriptionCard
            subscription={subscription}
            onClick={() => onClick(subscription.ID)}
          />
        ))}
      </Card.Group>
    </div>
  );
}
