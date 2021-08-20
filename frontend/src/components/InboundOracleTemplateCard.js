import { Card } from "semantic-ui-react";
import Identicon from "react-identicons";
import { useHistory } from "react-router-dom";

export default function InboundOracleTemplateCard({ inboundOracleTemplate }) {
  const history = useHistory();
  if (!inboundOracleTemplate) {
    return "";
  }

  return (
    <Card
      style={{ margin: "0 1em 1em 1em" }}
      onClick={() =>
        history.push("/inboundOracleTemplates/" + inboundOracleTemplate.ID)
      }
      raised
    >
      <Identicon
        string={inboundOracleTemplate.OracleTemplate.ContractAddress}
        size={290}
      />
      <Card.Content>
        <Card.Header>
          {inboundOracleTemplate.OracleTemplate.ContractAddress} -{" "}
          {inboundOracleTemplate.OracleTemplate.EventName}
        </Card.Header>
        <Card.Meta>
          {inboundOracleTemplate.OracleTemplate.BlockchainName}
        </Card.Meta>
      </Card.Content>
    </Card>
  );
}
