import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import useInboundSubscription from "../hooks/useInboundSubscription";
import { Button, Icon, Segment, Table, TableCell } from "semantic-ui-react";
import {
  ExampleRequest,
  InboundSubscriptionForm,
  SubscriptionOnOffRibbon,
  StartStopButton,
} from "../components";
import FilterForm from "../components/FilterForm";
import { ORACLE_STATUS_STARTED } from "../config/constants";
import config from "../config";

export default function InboundSubscriptionDetail({}) {
  const { inboundSubscriptionID } = useParams();
  const [
    inboundSubscription,
    updateInboundSubscription,
    loading,
    startInboundSubscription,
    stopInboundSubscription,
  ] = useInboundSubscription(inboundSubscriptionID);

  const [localInboundSubscription, setLocalInboundSubscription] = useState();
  useEffect(() => {
    setLocalInboundSubscription(inboundSubscription);
  }, [inboundSubscription]);

  const createFilter = () => {
    console.log(createFilter);
  };

  if (!inboundSubscription) {
    return "Loading...";
  }

  return (
    <div>
      <h1>Inbound Subscription</h1>
      <div
        style={{
          display: "flex",
          flexWrap: "wrap",
        }}
      >
        <div style={{ marginLeft: "1em", marginTop: "1em" }}>
          <Segment>
            <SubscriptionOnOffRibbon
              subscriptionStarted={
                inboundSubscription.Subscription.Status ===
                ORACLE_STATUS_STARTED
              }
            />
            <InboundSubscriptionForm
              inboundSubscription={localInboundSubscription}
              setInboundSubscription={setLocalInboundSubscription}
            />
            {JSON.stringify(localInboundSubscription) !==
              JSON.stringify(inboundSubscription) && (
              <>
                <br />
                <Button
                  positive
                  basic
                  loading={loading}
                  content="Save"
                  onClick={() =>
                    updateInboundSubscription(localInboundSubscription)
                  }
                />
                <br />
              </>
            )}
            <br />
            <p>
              <b>Webhook:</b> {config.BASE_URL}/inboundSubscriptions/
              {inboundSubscription.ID}/events
            </p>
            <StartStopButton
              loading={loading}
              subscriptionStarted={
                inboundSubscription.Subscription.Status ===
                ORACLE_STATUS_STARTED
              }
              stopSubscription={stopInboundSubscription}
              startSubscription={startInboundSubscription}
            />
          </Segment>
        </div>
        <div style={{ marginLeft: "1em", marginTop: "1em" }}>
          <ExampleRequest
            eventParameters={
              inboundSubscription.SmartContractConsumer.ProviderConsumer
                .EventParameters
            }
          />
        </div>
      </div>
      <br />
      <FilterForm
        providerConsumerID={
          inboundSubscription.SmartContractConsumer.ProviderConsumer.ID
        }
        subscriptionID={inboundSubscription.Subscription.ID}
      />
      <br />
      <div>
        <h2>Events</h2>
        {inboundSubscription.Subscription.Events.length > 0 ? (
          <Table unstackable>
            <Table.Header>
              <Table.Row>
                <Table.HeaderCell>ID</Table.HeaderCell>
                <Table.HeaderCell>At</Table.HeaderCell>
                <Table.HeaderCell>Content</Table.HeaderCell>
                <Table.HeaderCell>Success</Table.HeaderCell>
              </Table.Row>
            </Table.Header>
            <Table.Body>
              {inboundSubscription.Subscription.Events.map((inboundEvent) => (
                <Table.Row>
                  <Table.Cell>{inboundEvent.ID}</Table.Cell>
                  <Table.Cell>
                    {new Date(inboundEvent.CreatedAt).toLocaleString()}
                  </Table.Cell>
                  <Table.Cell>
                    {inboundEvent.EventValues.map((value) => (
                      <>
                        <b>{value.EventParameter.Name}:</b>
                        {value.Value}
                        <br />
                      </>
                    ))}
                  </Table.Cell>
                  <TableCell>
                    {inboundEvent.Success ? (
                      <Icon name="check" />
                    ) : (
                      <Icon name="times" />
                    )}
                  </TableCell>
                </Table.Row>
              ))}
            </Table.Body>
          </Table>
        ) : (
          <div>No events registered so far</div>
        )}
      </div>
    </div>
  );
}
