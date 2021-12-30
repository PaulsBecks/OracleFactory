import React from "react";
import { useHistory } from "react-router";
import { Button, Table } from "semantic-ui-react";

export default function HyperledgerBlockchainConnectorsTable({
  hyperledgerBlockchainConnectors,
}) {
  const history = useHistory();
  if (!hyperledgerBlockchainConnectors) {
    return "";
  }
  return (
    <Table celled>
      <Table.Header>
        <Table.Row>
          <Table.HeaderCell>Hyperledger Organization Name</Table.HeaderCell>
          <Table.HeaderCell>Hyperledger Channel</Table.HeaderCell>
          <Table.HeaderCell>On Chain</Table.HeaderCell>
          <Table.HeaderCell>Is Active</Table.HeaderCell>
          <Table.HeaderCell></Table.HeaderCell>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {hyperledgerBlockchainConnectors.map(
          (hyperledgerBlockchainConnector) => (
            <Table.Row>
              <Table.Cell>
                {hyperledgerBlockchainConnector.HyperledgerOrganizationName}
              </Table.Cell>
              <Table.Cell>
                {hyperledgerBlockchainConnector.HyperledgerChannel}
              </Table.Cell>
              <Table.Cell>
                {hyperledgerBlockchainConnector.OutboundOracle.IsOnChain
                  ? "true"
                  : "false"}
              </Table.Cell>
              <Table.Cell>
                {hyperledgerBlockchainConnector.OutboundOracle.IsActive
                  ? "true"
                  : "false"}
              </Table.Cell>
              <Table.Cell>
                <Button
                  primary
                  icon="pencil"
                  content="Details"
                  onClick={() => {
                    history.push(
                      "/blockchainConnectors/hyperledger/" +
                        hyperledgerBlockchainConnector.ID
                    );
                  }}
                />
              </Table.Cell>
            </Table.Row>
          )
        )}
      </Table.Body>
    </Table>
  );
}
