import React from "react";
import { useParams, Link } from "react-router-dom";
import useWebServicePublisher from "../hooks/useWebServicePublisher";
import { Button } from "semantic-ui-react";
import { OutboundOracleTable } from "../components";
import WebServicePublisherCard from "../components/WebServicePublisherCard";
export default function WebServicePublisherDetail() {
  const { webServicePublisherID } = useParams();
  const [webServicePublisher] = useWebServicePublisher(webServicePublisherID);
  if (!webServicePublisher) {
    return "Loading...";
  }

  console.log(webServicePublisher.OutboundOracles);

  return (
    <div>
      <h1>Web Service Publisher</h1>
      <WebServicePublisherCard webServicePublisher={webServicePublisher} />
      <div>
        <h2>Active Oracles</h2>
        <Button
          primary
          basic
          content="Create Oracle"
          icon="plus"
          as={Link}
          to={
            "/outboundOracles/create?webServicePublisherID=" +
            webServicePublisher.ID
          }
        />
        {webServicePublisher.OutboundOracles.length > 0 ? (
          <OutboundOracleTable
            outcdboundOracles={webServicePublisher.OutboundOracles}
          />
        ) : (
          <div>No oracles created yet.</div>
        )}
      </div>
    </div>
  );
}
