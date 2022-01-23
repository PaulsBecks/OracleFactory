import React from "react";
import { Form } from "semantic-ui-react";

export default function OutboundSubscriptionForm({
  outboundSubscription,
  setOutboundSubscription,
}) {
  function updateOutboundSubscription(_, { name, value }) {
    setOutboundSubscription({ ...outboundSubscription, [name]: value });
  }
  if (!outboundSubscription) return "";

  return (
    <Form>
      <Form.Input
        label="Name"
        name="Name"
        value={outboundSubscription.Subscription.Name}
        onChange={(_, { value }) =>
          setOutboundSubscription({
            ...outboundSubscription,
            Subscription: { ...outboundSubscription.Subscription, Name: value },
          })
        }
        placeholder="A name to recognize the subscription"
      />
    </Form>
  );
}
