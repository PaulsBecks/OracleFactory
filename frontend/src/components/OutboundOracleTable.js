import { Button, Table } from "semantic-ui-react";
import { Link } from "react-router-dom";
import Identicon from "react-identicons";

export default function OutboundOracleTable({ outboundOracles }) {
  return (
    <Table>
      <Table.Header>
        <Table.Row>
          <Table.HeaderCell></Table.HeaderCell>
          <Table.HeaderCell>Name</Table.HeaderCell>
          <Table.HeaderCell>Forward To</Table.HeaderCell>
          <Table.HeaderCell></Table.HeaderCell>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {outboundOracles.map((outboundOracle) => (
          <Table.Row>
            <Table.Cell>
              <Identicon
                string={
                  outboundOracle.OutboundOracleTemplate.OracleTemplate
                    .ContractAddress
                }
                size={50}
              />
            </Table.Cell>
            <Table.Cell>{outboundOracle.Oracle.Name || ""}</Table.Cell>
            <Table.Cell>{outboundOracle.URI}</Table.Cell>
            <Table.Cell>
              <Button
                as={Link}
                to={"/outboundOracles/" + outboundOracle.ID}
                content="Detail"
                icon="edit"
              />
            </Table.Cell>
          </Table.Row>
        ))}
      </Table.Body>
    </Table>
  );
}
