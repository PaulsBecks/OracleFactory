import React from "react";
import { useParams, Link } from "react-router-dom";
import useInboundOracleTemplate from "../hooks/useInboundOracleTemplate";
import { Button, Table } from "semantic-ui-react";

export default function InboundOracleTemplateDetail({}) {
  const { inboundOracleTemplateID } = useParams();
  const [inboundOracleTemplate] = useInboundOracleTemplate(
    inboundOracleTemplateID
  );

  if (!inboundOracleTemplate) {
    return "Loading...";
  }

  return (
    <div>
      <h1>Inbound Oracle Template</h1>
      <div>
        <p>
          <b>Event Name:</b> {inboundOracleTemplate.ContractName}
        </p>
        <p>
          <b>Contract Address:</b> {inboundOracleTemplate.ContractAddress}
        </p>
        <p>
          <b>Blockchain Name:</b> {inboundOracleTemplate.BlockchainName}
        </p>
        <p>
          <b>Blockchain Address:</b> {inboundOracleTemplate.BlockchainAddress}
        </p>
      </div>
      <br />
      <div>
        <h2>Inbound Oracles</h2>
        <Button
          content="Create Oracle"
          icon="plus"
          as={Link}
          to={
            "/inboundOracleTemplates/" +
            inboundOracleTemplateID +
            "/inboundOracles/create"
          }
        />
        {inboundOracleTemplate.InboundOracles.length > 0 ? (
          <Table>
            <Table.Header>
              <Table.Row>
                <Table.HeaderCell>Name</Table.HeaderCell>
                <Table.HeaderCell>Endpoint</Table.HeaderCell>
                <Table.HeaderCell></Table.HeaderCell>
              </Table.Row>
            </Table.Header>
            <Table.Body>
              {inboundOracleTemplate.InboundOracles.map((inboundOracle) => (
                <Table.Row>
                  <Table.Cell>{inboundOracle.Name || ""}</Table.Cell>
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
        ) : (
          <div>No oracles created yet.</div>
        )}
      </div>
    </div>
  );
}
