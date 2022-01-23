import { Button } from "semantic-ui-react";

export default function StartStopButton({
  loading,
  subscriptionStarted,
  stopSubscription,
  startSubscription,
}) {
  return (
    <Button
      basic
      fluid
      loading={loading}
      content={subscriptionStarted ? "Stop Subscription" : "Start Subscription"}
      color={subscriptionStarted ? "negative" : "positive"}
      onClick={subscriptionStarted ? stopSubscription : startSubscription}
    />
  );
}
