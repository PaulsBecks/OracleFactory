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
          <Table.HeaderCell collapsing></Table.HeaderCell>
          <Table.HeaderCell>Name</Table.HeaderCell>
          <Table.HeaderCell>Endpoint</Table.HeaderCell>
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
            <Table.Cell>
              <Identicon
                string={
                  inboundOracle.InboundOracleTemplate.OracleTemplate
                    .ContractAddress
                }
                size={50}
              />
            </Table.Cell>
            <Table.Cell>{inboundOracle.Oracle.Name || ""}</Table.Cell>
            <Table.Cell>
              {"http://localhost:8080/inboundOracles/" +
                inboundOracle.ID +
                "/events"}
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
        <OracleTemplateCard
          inboundOracleTemplate={inboundOracle.InboundOracleTemplate}
        />
      ))}
    </div>
  );*/
}
