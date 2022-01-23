import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Button, Segment, Table } from "semantic-ui-react";
import useOutboundSubscription from "../hooks/useOutboundSubscription";
import {
  ExampleRequest,
  SubscriptionOnOffRibbon,
  OutboundSubscriptionForm,
  StartStopButton,
} from "../components";
import FilterForm from "../components/FilterForm";
import { ORACLE_STATUS_STARTED } from "../config/constants";
import SmartContractProviderCard from "../components/SmartContractProviderCard";

export default function OutboundSubscriptionDetail({}) {
  const { outboundSubscriptionID } = useParams();
  const [
    outboundSubscription,
    updateOutboundSubscription,
    loading,
    startOutboundSubscription,
    stopOutboundSubscription,
  ] = useOutboundSubscription(outboundSubscriptionID);

  const [localOutboundSubscriptions, setLocalOutboundSubscription] = useState(
    {}
  );

  useEffect(() => {
    setLocalOutboundSubscription(outboundSubscription);
  }, [outboundSubscription]);
  if (!outboundSubscription) {
    return "Loading...";
  }

  let subscriptionStarted =
    outboundSubscription.Subscription.Status === ORACLE_STATUS_STARTED;
  return (
    <div>
      <h1>Outbound Subscription</h1>
      <div style={{ display: "flex", flexWrap: "wrap" }}>
        <div style={{ marginLeft: "1em", marginTop: "1em" }}>
          <Segment>
            <SubscriptionOnOffRibbon
              subscriptionStarted={subscriptionStarted}
            />
            <OutboundSubscriptionForm
              outboundSubscription={localOutboundSubscriptions}
              setOutboundSubscription={setLocalOutboundSubscription}
            />
            <br />
            <label>
              <b>Forward to:</b>
            </label>
            <p>{outboundSubscription.WebServiceConsumer.Url}</p>
            {JSON.stringify(localOutboundSubscriptions) !==
              JSON.stringify(outboundSubscription) && (
              <>
                <br />
                <Button
                  loading={loading}
                  content="Save"
                  positive
                  basic
                  onClick={async () =>
                    updateOutboundSubscription(localOutboundSubscriptions)
                  }
                />
                <br />
              </>
            )}
            <br />
            <StartStopButton
              loading={loading}
              subscriptionStarted={subscriptionStarted}
              stopSubscription={stopOutboundSubscription}
              startSubscription={startOutboundSubscription}
            />
          </Segment>
        </div>
        <div style={{ marginLeft: "1em", marginTop: "1em" }}>
          <ExampleRequest
            eventParameters={
              outboundSubscription.SmartContractProvider.ProviderConsumer
                .EventParameters
            }
          />
        </div>
      </div>
      <br />
      <FilterForm
        subscriptionID={outboundSubscription.SubscriptionID}
        providerConsumerID={
          outboundSubscription.SmartContractProvider.ProviderConsumer.ID
        }
      />
      <br />
      <div>
        <h2>Events</h2>
        {outboundSubscription.Subscription.Events.length > 0 ? (
          <Table>
            <Table.Header>
              <Table.Row>
                <Table.HeaderCell>ID</Table.HeaderCell>
                <Table.HeaderCell>At</Table.HeaderCell>
                <Table.HeaderCell>Content</Table.HeaderCell>
              </Table.Row>
            </Table.Header>
            <Table.Body>
              {outboundSubscription.Subscription.Events.map((outboundEvent) => (
                <Table.Row>
                  <Table.Cell>{outboundEvent.ID}</Table.Cell>
                  <Table.Cell>
                    {new Date(outboundEvent.CreatedAt).toLocaleString()}
                  </Table.Cell>
                  <Table.Cell>
                    {outboundEvent.EventValues.map((value) => {
                      console.log(value);
                      return (
                        <>
                          <b>{value.EventParameter.Name}:</b>
                          {value.Value}
                          <br />
                        </>
                      );
                    })}
                  </Table.Cell>
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
