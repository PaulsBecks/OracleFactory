import React from "react";
import { Table } from "semantic-ui-react";

export default function ProvidersTable({ providers }) {
  console.log(providers);
  if (!providers) {
    return "";
  }
  return (
    <Table>
      <Table.Header celled>
        <Table.Row>
          <Table.HeaderCell>Name</Table.HeaderCell>
          <Table.HeaderCell>Topic</Table.HeaderCell>
          <Table.HeaderCell>Endpoint</Table.HeaderCell>
          <Table.HeaderCell>Private</Table.HeaderCell>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {providers.map((provider) => (
          <Table.Row>
            <Table.Cell>{provider.Name}</Table.Cell>
            <Table.Cell>{provider.Topic}</Table.Cell>
            <Table.Cell>{provider.ID}</Table.Cell>
            <Table.Cell>{provider.Private ? "True" : "False"}</Table.Cell>
          </Table.Row>
        ))}
      </Table.Body>
    </Table>
  );
}
