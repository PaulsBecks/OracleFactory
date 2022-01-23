import React from "react";
import { Form } from "semantic-ui-react";

export default function InboundSubscriptionForm({
  inboundSubscription,
  setInboundSubscription,
}) {
  function updateInboundSubscription(_, { name, value }) {
    setInboundSubscription({ ...inboundSubscription, [name]: value });
  }
  if (!inboundSubscription) return "";

  return (
    <Form>
      <Form.Input
        label="Name"
        name="Name"
        value={inboundSubscription.Subscription.Name}
        onChange={(_, { value }) =>
          setInboundSubscription({
            ...inboundSubscription,
            Subscription: { ...inboundSubscription.Subscription, Name: value },
          })
        }
        placeholder="A name to recognize the subscription"
      />
    </Form>
  );
}
