import { Label } from "semantic-ui-react";

export default function SubscriptionOnOffRibbon({ subscriptionStarted }) {
  return (
    <Label
      ribbon
      content={subscriptionStarted ? "ON" : "OFF"}
      color={subscriptionStarted ? "green" : "red"}
    />
  );
}
