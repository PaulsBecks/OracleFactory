import { Button, Label, Table } from "semantic-ui-react";
import { Link } from "react-router-dom";
import Identicon from "react-identicons";
import { ORACLE_STATUS_STARTED } from "../config/constants";
import SubscriptionOnOffRibbon from "./SubscriptionOnOffRibbon";

export default function InboundSubscriptionTable({ inboundSubscriptions }) {
  return (
    <Table color="orange">
      <Table.Header>
        <Table.Row>
          <Table.HeaderCell collapsing></Table.HeaderCell>
          <Table.HeaderCell>Name</Table.HeaderCell>
          <Table.HeaderCell>Provider</Table.HeaderCell>
          <Table.HeaderCell>Consumer</Table.HeaderCell>
          <Table.HeaderCell></Table.HeaderCell>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {inboundSubscriptions.map((inboundSubscription) => (
          <Table.Row>
            <Table.Cell>
              <SubscriptionOnOffRibbon
                subscriptionStarted={
                  inboundSubscription.Subscription.Status ===
                  ORACLE_STATUS_STARTED
                }
              />
            </Table.Cell>
            <Table.Cell>
              {inboundSubscription.Subscription.Name || ""}
            </Table.Cell>
            <Table.Cell>
              <div
                style={{
                  display: "flex",
                  flexWrap: "wrap",
                  alignItems: "center",
                }}
              >
                <div style={{ marginRight: "1em" }}>
                  <Identicon
                    string={
                      "Provider" + inboundSubscription.WebServiceProvider.ID
                    }
                    size={50}
                  />
                </div>
                <label>
                  {inboundSubscription.WebServiceProvider.ProviderConsumer.Name}
                </label>
              </div>
            </Table.Cell>
            <Table.Cell>
              <div
                style={{
                  display: "flex",
                  flexWrap: "wrap",
                  alignItems: "center",
                }}
              >
                <div style={{ marginRight: "1em" }}>
                  <Identicon
                    string={
                      inboundSubscription.SmartContractConsumer.SmartContract
                        .ContractAddress
                    }
                    size={50}
                  />
                </div>
                <label>
                  {
                    inboundSubscription.SmartContractConsumer.ProviderConsumer
                      .Name
                  }
                </label>
              </div>
            </Table.Cell>
            <Table.Cell textAlign="right">
              <Button
                as={Link}
                to={"/inboundSubscriptions/" + inboundSubscription.ID}
                content="Detail"
                icon="edit"
                primary
                basic
              />
            </Table.Cell>
          </Table.Row>
        ))}
      </Table.Body>
    </Table>
  );
  /*return (
    <div>
      {inboundSubscriptions.map((inboundSubscription) => (
        <SmartContractCard
          smartContractConsumer={inboundSubscription.SmartContractConsumer}
        />
      ))}
    </div>
  );*/
}
