import { Button, Label, Table } from "semantic-ui-react";
import { Link } from "react-router-dom";
import Identicon from "react-identicons";
import { ORACLE_STATUS_STARTED } from "../config/constants";
import OracleOnOffRibbon from "./OracleOnOffRibbon";

export default function OutboundOracleTable({ outboundOracles }) {
  if (!outboundOracles) {
    return "";
  }
  return (
    <Table color="orange">
      <Table.Header>
        <Table.Row>
          <Table.HeaderCell collapsing></Table.HeaderCell>
          <Table.HeaderCell>Name</Table.HeaderCell>
          <Table.HeaderCell>Sub/Unsub In</Table.HeaderCell>
          <Table.HeaderCell>Publisher</Table.HeaderCell>
          <Table.HeaderCell></Table.HeaderCell>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {outboundOracles.map((outboundOracle) => (
          <Table.Row>
            <Table.Cell>
              <OracleOnOffRibbon
                oracleStarted={
                  outboundOracle.Oracle.Status === ORACLE_STATUS_STARTED
                }
              />
            </Table.Cell>
            <Table.Cell>{outboundOracle.Oracle.Name || ""}</Table.Cell>
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
                      outboundOracle.BlockchainEvent.SmartContract
                        .ContractAddress
                    }
                    size={50}
                  />
                </div>
                <label>
                  {outboundOracle.BlockchainEvent.ListenerPublisher.Name}
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
                    string={"Publisher" + outboundOracle.PubSubOracle.ID}
                    size={50}
                  />
                </div>
                <label>{outboundOracle.PubSubOracle.Name}</label>
              </div>
            </Table.Cell>
            <Table.Cell textAlign="right">
              <Button
                as={Link}
                to={"/outboundOracles/" + outboundOracle.ID}
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
