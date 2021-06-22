import React from "react";
import { useParams, Link } from "react-router-dom";
import useInboundOracleTemplate from "../hooks/useInboundOracleTemplate";
import { Button } from "semantic-ui-react";
import { InboundOracleTable } from "../components";

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
        <p>
          <b>Parameters:</b>
          {JSON.stringify(
            inboundOracleTemplate.EventParameters.map((parameter) => ({
              name: parameter.Name,
              type: parameter.Type,
            }))
          )}
        </p>
      </div>
      <br />
      <div>
        <h2>Active Oracles</h2>
        <Button
          basic
          primary
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
          <InboundOracleTable
            inboundOracles={inboundOracleTemplate.InboundOracles}
          />
        ) : (
          <div>No oracles created yet.</div>
        )}
      </div>
    </div>
  );
}
