import useOutboundOracleTemplates from "../hooks/useOutboundOracleTemplates";
import useInboundOraclesTemplates from "../hooks/useInboundOraclesTemplates";
import { useHistory } from "react-router-dom";

function Home() {
  const history = useHistory();
  const [outboundOracleTemplates] = useOutboundOracleTemplates();
  const [inboundOracleTemplates] = useInboundOraclesTemplates();
  return (
    <div>
      <h1>Home</h1>
      <div>
        <h2>Outbound Oracle Templates</h2>
        <div style={{ display: "flex", justifyContent: "space-between" }}>
          {outboundOracleTemplates.map((outboundOracleTemplate) => (
            <div
              style={{
                border: "1px solid black",
                borderRadius: "1em",
                padding: "2em",
                cursor: "pointer",
              }}
              onClick={() =>
                history.push(
                  "/outboundOracleTemplates/" + outboundOracleTemplate.ID
                )
              }
            >
              <h3 style={{ marginTop: "0" }}>
                Event: {outboundOracleTemplate.EventName}
              </h3>
              <div>
                <b>At:</b> {outboundOracleTemplate.Address}
              </div>
              <div>
                <b>Running Oracles:</b>{" "}
                {outboundOracleTemplate.OutboundOracles.length}
              </div>
            </div>
          ))}
        </div>
      </div>
      <div>
        <h2>Inbound Oracle Templates</h2>
        <div style={{ display: "flex", justifyContent: "space-between" }}>
          {inboundOracleTemplates.map((inboundOracleTemplate) => (
            <div
              style={{
                border: "1px solid black",
                borderRadius: "1em",
                padding: "2em",
                cursor: "pointer",
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

export default Home;
