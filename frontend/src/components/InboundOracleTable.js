import { Button, Table } from "semantic-ui-react";
import { Link } from "react-router-dom";
import Identicon from "react-identicons";

export default function InboundOracleTable({ inboundOracles }) {
  return (
    <Table>
      <Table.Header>
        <Table.Row>
          <Table.HeaderCell></Table.HeaderCell>
          <Table.HeaderCell>Name</Table.HeaderCell>
          <Table.HeaderCell>Endpoint</Table.HeaderCell>
          <Table.HeaderCell></Table.HeaderCell>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {inboundOracles.map((inboundOracle) => (
          <Table.Row>
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
            <Table.Cell>
              <Button
                as={Link}
                to={"/inboundOracles/" + inboundOracle.ID}
                content="Detail"
                icon="edit"
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
