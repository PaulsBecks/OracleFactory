import React from "react";
import { useParams, Link } from "react-router-dom";
import useInboundOracleTemplate from "../hooks/useInboundOracleTemplate";
import { Button } from "semantic-ui-react";
import { ExampleRequest, InboundOracleTable } from "../components";
import InboundOracleTemplateCard from "../components/InboundOracleTemplateCard";

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
      <div style={{ display: "flex" }}>
        <InboundOracleTemplateCard
          inboundOracleTemplate={inboundOracleTemplate}
        />
        <div style={{ marginLeft: "2em" }}>
          <ExampleRequest
            eventParameters={
              inboundOracleTemplate.OracleTemplate.EventParameters
            }
          />
        </div>
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
