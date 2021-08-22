import useOutboundOracleTemplates from "../hooks/useOutboundOracleTemplates";
import useInboundOraclesTemplates from "../hooks/useInboundOraclesTemplates";
import { useHistory, Link } from "react-router-dom";
import { Button, Card, Table } from "semantic-ui-react";
import Identicon from "react-identicons";
import OutboundOracleTemplateCard from "../components/OutboundOracleTemplateCard";
import InboundOracleTemplateCard from "../components/InboundOracleTemplateCard";

function Home() {
  const [outboundOracleTemplates] = useOutboundOracleTemplates();
  const [inboundOracleTemplates] = useInboundOraclesTemplates();
  return (
    <div>
      <div>
        <Button
          basic
          primary
          icon="plus"
          content="Create Template"
          as={Link}
          to="/oracleTemplates/create"
        />
      </div>
      <br />
      <div>
        <h2>Outbound Oracle Templates</h2>
        <Card.Group>
          {outboundOracleTemplates.map((outboundOracleTemplate) => (
            <OutboundOracleTemplateCard
              outboundOracleTemplate={outboundOracleTemplate}
            />
          ))}
        </Card.Group>
      </div>
      <br />
      <br />
      <div>
        <h2>Inbound Oracle Templates</h2>
        <Card.Group>
          {inboundOracleTemplates.map((inboundOracleTemplate) => (
            <InboundOracleTemplateCard
              inboundOracleTemplate={inboundOracleTemplate}
            />
          ))}
        </Card.Group>
      </div>
    </div>
  );
}

export default Home;
