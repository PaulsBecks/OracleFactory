import useOutboundOracles from "../hooks/useOutboundOracles";
import useInboundOracles from "../hooks/useInboundOracles";
import { Link } from "react-router-dom";
import { InboundOracleTable, OutboundOracleTable } from "../components";
import { Button } from "semantic-ui-react";
function Home() {
  const [outboundOracles] = useOutboundOracles();
  const [inboundOracles] = useInboundOracles();
  return (
    <div>
      <h1>Home</h1>
      <div style={{ display: "flex" }}>
        <Button
          basic
          primary
          icon="plus"
          content="Create Oracle"
          as={Link}
          to="/oracleTemplates"
        />
        <Button
          basic
          icon="settings"
          content="Settings"
          as={Link}
          to="/settings"
        />
      </div>
      <br />
      <div>
        <h2>Outbound Oracles</h2>
        <OutboundOracleTable outboundOracles={outboundOracles} />
      </div>
      <br />
      <div>
        <h2>Inbound Oracles</h2>
        <InboundOracleTable inboundOracles={inboundOracles} />
      </div>
    </div>
  );
}

export default Home;
