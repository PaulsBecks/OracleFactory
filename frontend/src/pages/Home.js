import useOutboundOracles from "../hooks/useOutboundOracles";
import usePubSubOracles from "../hooks/usePubSubOracles";
import { Link } from "react-router-dom";
import { PubSubOracleTable, OutboundOracleTable } from "../components";
import { Button } from "semantic-ui-react";

function Home() {
  const [outboundOracles] = useOutboundOracles();
  const [pubSubOracles] = usePubSubOracles();
  return (
    <div>
      <div style={{ display: "flex" }}>
        <Button
          basic
          primary
          icon="plus"
          content="Create Oracle"
          as={Link}
          to="/pubSubOracles/create"
        />
      </div>
      <br />
      {pubSubOracles && pubSubOracles.length > 0 && (
        <div>
          <h2>Pub-Sub Oracles</h2>
          <PubSubOracleTable pubSubOracles={pubSubOracles} />
        </div>
      )}
    </div>
  );
}

export default Home;
