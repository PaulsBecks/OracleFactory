import { Button, Label, Table } from "semantic-ui-react";
import { Link } from "react-router-dom";
import Identicon from "react-identicons";
import { ORACLE_STATUS_STARTED } from "../config/constants";
import OracleOnOffRibbon from "./OracleOnOffRibbon";

export default function PubSubOracleTable({ pubSubOracles }) {
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
        {pubSubOracles.map((pubSubOracle) => (
          <Table.Row>
            <Table.Cell>
              <OracleOnOffRibbon
                oracleStarted={
                  pubSubOracle.Oracle.Status === ORACLE_STATUS_STARTED
                }
              />
            </Table.Cell>
            <Table.Cell>{pubSubOracle.Oracle.Name || ""}</Table.Cell>
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
                    string={"Listener" + pubSubOracle.Provider.ID}
                    size={50}
                  />
                </div>
                <label>{pubSubOracle.Provider.ListenerPublisher.Name}</label>
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
                    string={pubSubOracle.Consumer.SmartContract.ContractAddress}
                    size={50}
                  />
                </div>
                <label>{pubSubOracle.Consumer.ListenerPublisher.Name}</label>
              </div>
            </Table.Cell>
            <Table.Cell textAlign="right">
              <Button
                as={Link}
                to={"/pubSubOracles/" + pubSubOracle.ID}
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
      {pubSubOracles.map((pubSubOracle) => (
        <SmartContractCard
          consumer={pubSubOracle.Consumer}
        />
      ))}
    </div>
  );*/
}
