import useInboundOraclesTemplates from "../hooks/useInboundOraclesTemplates";
import { useHistory, Link } from "react-router-dom";
import { Button } from "semantic-ui-react";

function InboundOracleTemplates() {
  const history = useHistory();
  const [inboundOracleTemplates] = useInboundOraclesTemplates();
  console.log(inboundOracleTemplates);
  return (
    <div>
      <h1>Inbound Oracle Templates</h1>
      <div>
        <Button
          basic
          primary
          icon="plus"
          content="Create Template"
          as={Link}
          to="/inboundOracleTemplates/create"
        />
      </div>
      <br />
      <div>
        <div style={{ display: "flex", flexWrap: "wrap" }}>
          {inboundOracleTemplates.map((inboundOracleTemplate) => (
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
                  "/inboundOracleTemplates/" + inboundOracleTemplate.ID
                )
              }
            >
              <h3 style={{ marginTop: "0" }}>
                Method: {inboundOracleTemplate.ContractName}
              </h3>
              <div>
                <b>At:</b> {inboundOracleTemplate.ContractAddress}
              </div>
              <div>
                <b>Running Oracles:</b>{" "}
                {inboundOracleTemplate.InboundOracles.length}
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default InboundOracleTemplates;
