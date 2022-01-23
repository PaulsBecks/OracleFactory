import { Button, Label, Table } from "semantic-ui-react";
import { Link } from "react-router-dom";
import Identicon from "react-identicons";
import { ORACLE_STATUS_STARTED } from "../config/constants";
import SubscriptionOnOffRibbon from "./SubscriptionOnOffRibbon";

export default function OutboundSubscriptionTable({ outboundSubscriptions }) {
  console.log(outboundSubscriptions);
  if (!outboundSubscriptions) {
    return "";
  }
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
        {outboundSubscriptions.map((outboundSubscription) => (
          <Table.Row>
            <Table.Cell>
              <SubscriptionOnOffRibbon
                subscriptionStarted={
                  outboundSubscription.Subscription.Status ===
                  ORACLE_STATUS_STARTED
                }
              />
            </Table.Cell>
            <Table.Cell>
              {outboundSubscription.Subscription.Name || ""}
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
                      outboundSubscription.SmartContractProvider.SmartContract
                        .ContractAddress
                    }
                    size={50}
                  />
                </div>
                <label>
                  {
                    outboundSubscription.SmartContractProvider.ProviderConsumer
                      .Name
                  }
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
                      "Consumer" + outboundSubscription.WebServiceConsumer.ID
                    }
                    size={50}
                  />
                </div>
                <label>
                  {
                    outboundSubscription.WebServiceConsumer.ProviderConsumer
                      .Name
                  }
                </label>
              </div>
            </Table.Cell>
            <Table.Cell textAlign="right">
              <Button
                as={Link}
                to={"/outboundSubscriptions/" + outboundSubscription.ID}
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
}
