import React from "react";
import { useHistory, useParams, Link } from "react-router-dom";
import useOutboundOracleTemplate from "../hooks/useOutboundOracleTemplate";
import { Button } from "semantic-ui-react";
import { OutboundOracleTable } from "../components";
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
        <h2>Active Oracles</h2>
        <Button
          primary
          basic
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
          <OutboundOracleTable
            outboundOracles={outboundOracleTemplate.OutboundOracles}
          />
        ) : (
          <div>No oracles created yet.</div>
        )}
      </div>
    </div>
  );
}
