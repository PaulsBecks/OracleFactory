import { Button, Label, Table } from "semantic-ui-react";
import { Link } from "react-router-dom";
import Identicon from "react-identicons";
import { ORACLE_STATUS_STARTED } from "../config/constants";
import OracleOnOffRibbon from "./OracleOnOffRibbon";

export default function InboundOracleTable({ inboundOracles }) {
  return (
    <Table color="orange">
      <Table.Header>
        <Table.Row>
          <Table.HeaderCell collapsing></Table.HeaderCell>
          <Table.HeaderCell>Name</Table.HeaderCell>
          <Table.HeaderCell>Listener</Table.HeaderCell>
          <Table.HeaderCell>Publisher</Table.HeaderCell>
          <Table.HeaderCell></Table.HeaderCell>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {inboundOracles.map((inboundOracle) => (
          <Table.Row>
            <Table.Cell>
              <OracleOnOffRibbon
                oracleStarted={
                  inboundOracle.Oracle.Status === ORACLE_STATUS_STARTED
                }
              />
            </Table.Cell>
            <Table.Cell>{inboundOracle.Oracle.Name || ""}</Table.Cell>
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
                    string={"Listener" + inboundOracle.WebServiceListener.ID}
                    size={50}
                  />
                </div>
                <label>
                  {inboundOracle.WebServiceListener.ListenerPublisher.Name}
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
                      inboundOracle.SmartContractPublisher.SmartContract
                        .ContractAddress
                    }
                    size={50}
                  />
                </div>
                <label>
                  {inboundOracle.SmartContractPublisher.ListenerPublisher.Name}
                </label>
              </div>
            </Table.Cell>
            <Table.Cell textAlign="right">
              <Button
                as={Link}
                to={"/inboundOracles/" + inboundOracle.ID}
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
      {inboundOracles.map((inboundOracle) => (
        <SmartContractCard
          smartContractPublisher={inboundOracle.SmartContractPublisher}
        />
      ))}
    </div>
  );*/
}
