import useOutboundOracles from "../hooks/useOutboundOracles";
import useInboundOracles from "../hooks/useInboundOracles";
import { Link } from "react-router-dom";
import { InboundOracleTable, OutboundOracleTable } from "../components";
import { Button, Message } from "semantic-ui-react";
import useUser from "../hooks/useUser";

function Home() {
  const [outboundOracles] = useOutboundOracles();
  const [inboundOracles] = useInboundOracles();
  const [user] = useUser();
  console.log(user);
  return (
    <div>
      <div style={{ display: "flex" }}>
        <Button
          basic
          primary
          icon="plus"
          content="Create Oracle"
          as={Link}
          to="/smartContracts"
        />
      </div>
      <br />
      {(!user || !user.EthereumPrivateKey || !user.EthereumAddress) && (
        <Message warning>
          You did not yet complete your ethereum connection settings. You have
          to enter your credentials before you can create oracles.
        </Message>
      )}
      {outboundOracles && outboundOracles.length > 0 && (
        <div>
          <h2>Outbound Oracles</h2>
          <OutboundOracleTable outboundOracles={outboundOracles} />
        </div>
      )}
      <br />
      {inboundOracles && inboundOracles.length > 0 && (
        <div>
          <h2>Inbound Oracles</h2>
          <InboundOracleTable inboundOracles={inboundOracles} />
        </div>
      )}
    </div>
  );
}

export default Home;
