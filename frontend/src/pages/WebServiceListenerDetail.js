import React from "react";
import { useParams, Link } from "react-router-dom";
import useWebServiceListener from "../hooks/useWebServiceListener";
import { Button } from "semantic-ui-react";
import { InboundOracleTable } from "../components";
import WebServiceListenerCard from "../components/WebServiceListenerCard";
export default function WebServiceListenerDetail() {
  const { webServiceListenerID } = useParams();
  const [webServiceListener] = useWebServiceListener(webServiceListenerID);
  if (!webServiceListener) {
    return "Loading...";
  }

  return (
    <div>
      <h1>Web Service Listener</h1>
      <WebServiceListenerCard webServiceListener={webServiceListener} />
      <div>
        <h2>Active Oracles</h2>
        <Button
          primary
          basic
          content="Create Oracle"
          icon="plus"
          as={Link}
          to={
            "/inboundOracles/create?webServiceListenerID=" +
            webServiceListener.ID
          }
        />
        {webServiceListener.InboundOracles.length > 0 ? (
          <InboundOracleTable
            inboundOracles={webServiceListener.InboundOracles}
          />
        ) : (
          <div>No oracles created yet.</div>
        )}
      </div>
    </div>
  );
}
