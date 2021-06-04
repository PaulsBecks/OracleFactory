import React from "react";
import { useHistory, useParams, Link } from "react-router-dom";
import useOutboundOracleTemplate from "../hooks/useOutboundOracleTemplate";
import { Button, Form, Table } from "semantic-ui-react";

export default function OutboundOracleTemplateDetail({}) {
  const { outboundOracleTemplateID } = useParams();
  const history = useHistory();
  const [outboundOracleTemplate] = useOutboundOracleTemplate(
    outboundOracleTemplateID
  );

  if (!outboundOracleTemplate) {
    return "Loading...";
  }

  return (
    <div>
      <h1>Outbound Oracle Template</h1>
      <div>
        <p>
          <b>Event Name:</b> {outboundOracleTemplate.EventName}
        </p>
        <p>
          <b>Contract Address:</b> {outboundOracleTemplate.Address}
        </p>
        <p>
          <b>Blockchain Name:</b> {outboundOracleTemplate.Blockchain}
        </p>
        <p>
          <b>Blockchain Address:</b> {outboundOracleTemplate.BlockchainAddress}
        </p>
        <br />
      </div>
      <div>
        <h2>Oracles</h2>
        <Button
          content="Create Oracle"
          icon="plus"
          as={Link}
          to={
            "/outboundOracleTemplates/" +
            outboundOracleTemplateID +
            "/outboundOracles/create"
          }
        />
        {outboundOracleTemplate.OutboundOracles.length > 0 ? (
          <Table>
            <Table.Header>
              <Table.Row>
                <Table.HeaderCell>Name</Table.HeaderCell>
                <Table.HeaderCell>Forward To</Table.HeaderCell>
                <Table.HeaderCell></Table.HeaderCell>
              </Table.Row>
            </Table.Header>
            <Table.Body>
              {outboundOracleTemplate.OutboundOracles.map((outboundOracle) => (
                <Table.Row>
                  <Table.Cell>{outboundOracle.Name || ""}</Table.Cell>
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
        ) : (
          <div>No oracles created yet.</div>
        )}
      </div>
    </div>
  );
}
