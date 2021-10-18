import useSmartContractPublishers from "../hooks/useSmartContractPublishers";
import { useHistory, Link } from "react-router-dom";
import { Button } from "semantic-ui-react";

function SmartContractPublishers() {
  const history = useHistory();
  const [smartContractPublishers] = useSmartContractPublishers();
  return (
    <div>
      <h1>Smart Contract Publisher</h1>
      <div>
        <Button
          basic
          primary
          icon="plus"
          content="Create Publisher"
          as={Link}
          to="/smartContractPublishers/create"
        />
      </div>
      <br />
      <div>
        <div style={{ display: "flex", flexWrap: "wrap" }}>
          {smartContractPublishers.map((smartContractPublisher) => (
            <div
              style={{
                border: "1px solid black",
                borderRadius: "1em",
                padding: "2em",
                cursor: "pointer",
                marginRight: "2em",
              }}
              onClick={() =>
                history.push(
                  "/smartContractPublishers/" + smartContractPublisher.ID
                )
              }
            >
              <h3 style={{ marginTop: "0" }}>
                Method: {smartContractPublisher.ContractName}
              </h3>
              <div>
                <b>At:</b> {smartContractPublisher.ContractAddress}
              </div>
              <div>
                <b>Running Oracles:</b>{" "}
                {smartContractPublisher.InboundOracles.length}
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default SmartContractPublishers;
