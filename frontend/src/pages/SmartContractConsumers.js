import useSmartContractConsumers from "../hooks/useSmartContractConsumers";
import { useHistory, Link } from "react-router-dom";
import { Button } from "semantic-ui-react";

function SmartContractConsumers() {
  const history = useHistory();
  const [smartContractConsumers] = useSmartContractConsumers();
  return (
    <div>
      <h1>Smart Contract Consumer</h1>
      <div>
        <Button
          basic
          primary
          icon="plus"
          content="Create Consumer"
          as={Link}
          to="/smartContractConsumers/create"
        />
      </div>
      <br />
      <div>
        <div style={{ display: "flex", flexWrap: "wrap" }}>
          {smartContractConsumers.map((smartContractConsumer) => (
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
                  "/smartContractConsumers/" + smartContractConsumer.ID
                )
              }
            >
              <h3 style={{ marginTop: "0" }}>
                Method: {smartContractConsumer.ContractName}
              </h3>
              <div>
                <b>At:</b> {smartContractConsumer.ContractAddress}
              </div>
              <div>
                <b>Running Subscriptions:</b>{" "}
                {smartContractConsumer.InboundSubscriptions.length}
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default SmartContractConsumers;
