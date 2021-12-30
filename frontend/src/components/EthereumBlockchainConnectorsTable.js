import React from "react";
import { useHistory } from "react-router";
import { Button, Table } from "semantic-ui-react";

export default function EthereumBlockchainConnectorsTable({
  ethereumBlockchainConnectors,
}) {
  const history = useHistory();
  if (!ethereumBlockchainConnectors) {
    return "";
  }
  return (
    <Table celled>
      <Table.Header>
        <Table.Row>
          <Table.HeaderCell>Ethereum Node Address</Table.HeaderCell>
          <Table.HeaderCell>Private Key</Table.HeaderCell>
          <Table.HeaderCell>On Chain</Table.HeaderCell>
          <Table.HeaderCell>Is Active</Table.HeaderCell>
          <Table.HeaderCell></Table.HeaderCell>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {ethereumBlockchainConnectors.map((ethereumBlockchainConnector) => (
          <Table.Row>
            <Table.Cell>
              {ethereumBlockchainConnector.EthereumAddress}
            </Table.Cell>
            <Table.Cell>
              {ethereumBlockchainConnector.EthereumPrivateKey}
            </Table.Cell>
            <Table.Cell>
              {ethereumBlockchainConnector.OutboundOracle.IsOnChain
                ? "true"
                : "false"}
            </Table.Cell>
            <Table.Cell>
              {ethereumBlockchainConnector.OutboundOracle.IsActive
                ? "true"
                : "false"}
            </Table.Cell>
            <Table.Cell>
              <Button
                content={"Details"}
                icon={"pencil"}
                primary
                onClick={() => {
                  history.push(
                    "/blockchainConnectors/ethereum/" +
                      ethereumBlockchainConnector.ID
                  );
                }}
              />
            </Table.Cell>
          </Table.Row>
        ))}
      </Table.Body>
    </Table>
  );
}
