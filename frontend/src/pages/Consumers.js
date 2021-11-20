import useConsumers from "../hooks/useConsumers";
import { useHistory, Link } from "react-router-dom";
import { Button } from "semantic-ui-react";

function Consumers() {
  const history = useHistory();
  const [consumers] = useConsumers();
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
          to="/consumers/create"
        />
      </div>
      <br />
      <div>
        <div style={{ display: "flex", flexWrap: "wrap" }}>
          {consumers.map((consumer) => (
            <div
              style={{
                border: "1px solid black",
                borderRadius: "1em",
                padding: "2em",
                cursor: "pointer",
                marginRight: "2em",
              }}
              onClick={() => history.push("/consumers/" + consumer.ID)}
            >
              <h3 style={{ marginTop: "0" }}>
                Method: {consumer.ContractName}
              </h3>
              <div>
                <b>At:</b> {consumer.ContractAddress}
              </div>
              <div>
                <b>Running Oracles:</b> {consumer.PubSubOracles.length}
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default Consumers;
