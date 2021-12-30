import React from "react";
import { Table } from "semantic-ui-react";

export default function SubscriptionsTable({ subscriptions }) {
  if (!subscriptions) {
    return "";
  }
  return (
    <Table>
      <Table.Header celled>
        <Table.Row>
          <Table.HeaderCell>Topic</Table.HeaderCell>
          <Table.HeaderCell>Filter</Table.HeaderCell>
          <Table.HeaderCell>Callback Method</Table.HeaderCell>
          <Table.HeaderCell>Smart Contract Address</Table.HeaderCell>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {subscriptions.map((subscription) => (
          <Table.Row>
            <Table.Cell>{subscription.Topic}</Table.Cell>
            <Table.Cell>{subscription.Filter}</Table.Cell>
            <Table.Cell>{subscription.Callback}</Table.Cell>
            <Table.Cell>{subscription.SmartContractAddress}</Table.Cell>
          </Table.Row>
        ))}
      </Table.Body>
    </Table>
  );
}
