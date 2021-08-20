import { Card } from "semantic-ui-react";
import Identicon from "react-identicons";
import { useHistory } from "react-router-dom";

export default function OutboundOracleTemplateCard({ outboundOracleTemplate }) {
  const history = useHistory();
  if (!outboundOracleTemplate) {
    return "";
  }

  return (
    <Card
      style={{ margin: "0 1em 0 1em" }}
      onClick={() =>
        history.push("/outboundOracleTemplates/" + outboundOracleTemplate.ID)
      }
    >
      <Identicon
        string={outboundOracleTemplate.OracleTemplate.ContractAddress}
        size={290}
      />
      <Card.Content>
        <Card.Header>
          {outboundOracleTemplate.OracleTemplate.ContractAddress} -{" "}
          {outboundOracleTemplate.OracleTemplate.EventName}
        </Card.Header>
        <Card.Meta>
          {outboundOracleTemplate.OracleTemplate.BlockchainName}
        </Card.Meta>
      </Card.Content>
    </Card>
  );
}
