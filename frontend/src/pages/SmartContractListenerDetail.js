import React from "react";
import { useHistory, useParams, Link } from "react-router-dom";
import useSmartContractListener from "../hooks/useSmartContractListener";
import { Button } from "semantic-ui-react";
import { OutboundOracleTable } from "../components";
import SmartContractListenerCard from "../components/SmartContractListenerCard";
export default function SmartContractListenerDetail({}) {
  const { smartContractListenerID } = useParams();
  const history = useHistory();
  const [smartContractListener] = useSmartContractListener(
    smartContractListenerID
  );

  if (!smartContractListener) {
    return "Loading...";
  }

  return (
    <div>
      <h1>Smart Contract Listener</h1>
      <SmartContractListenerCard
        smartContractListener={smartContractListener}
      />
      <div>
        <h2>Active Oracles</h2>
        <Button
          primary
          basic
          content="Create Oracle"
          icon="plus"
          as={Link}
          to={
            "/outboundOracles/create?smartContractListenerID=" +
            smartContractListenerID
          }
        />
        {smartContractListener.OutboundOracles.length > 0 ? (
          <OutboundOracleTable
            outboundOracles={smartContractListener.OutboundOracles}
          />
        ) : (
          <div>No oracles created yet.</div>
        )}
      </div>
    </div>
  );
}
