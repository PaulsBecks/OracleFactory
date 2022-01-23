import useOutboundSubscriptions from "../hooks/useOutboundSubscriptions";
import useInboundSubscriptions from "../hooks/useInboundSubscriptions";
import { Link } from "react-router-dom";
import {
  InboundSubscriptionTable,
  OutboundSubscriptionTable,
} from "../components";
import { Button, Message } from "semantic-ui-react";
import useUser from "../hooks/useUser";

function Home() {
  const [outboundSubscriptions] = useOutboundSubscriptions();
  const [inboundSubscriptions] = useInboundSubscriptions();
  const [user] = useUser();
  console.log(user);
  return (
    <div>
      <div style={{ display: "flex" }}>
        <Button
          basic
          primary
          icon="plus"
          content="Create Subscription"
          as={Link}
          to="/smartContracts"
        />
      </div>
      <br />
      {(!user || !user.EthereumPrivateKey || !user.EthereumAddress) && (
        <Message warning>
          You did not yet complete your{" "}
          <Link to="/settings">ethereum connection settings</Link>. You have to
          enter your credentials before you can create subscriptions.
        </Message>
      )}
      {outboundSubscriptions && outboundSubscriptions.length > 0 && (
        <div>
          <h2>Outbound Subscriptions</h2>
          <OutboundSubscriptionTable
            outboundSubscriptions={outboundSubscriptions}
          />
        </div>
      )}
      <br />
      {inboundSubscriptions && inboundSubscriptions.length > 0 && (
        <div>
          <h2>Inbound Subscriptions</h2>
          <InboundSubscriptionTable
            inboundSubscriptions={inboundSubscriptions}
          />
        </div>
      )}
    </div>
  );
}

export default Home;
