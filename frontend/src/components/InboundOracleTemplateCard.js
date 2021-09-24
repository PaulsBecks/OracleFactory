import { Card } from "semantic-ui-react";
import Identicon from "react-identicons";
import { useHistory } from "react-router-dom";

export default function InboundOracleTemplateCard({ inboundOracleTemplate }) {
  const history = useHistory();
  if (!inboundOracleTemplate) {
    return "";
  }

  console.log(inboundOracleTemplate);

  return (
    <Card
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
          {inboundOracleTemplate.OracleTemplate.ContractAddressSynonym} -{" "}
          {inboundOracleTemplate.OracleTemplate.EventName}
        </Card.Header>
        <Card.Meta>
          {inboundOracleTemplate.OracleTemplate.BlockchainName}
        </Card.Meta>
      </Card.Content>
    </Card>
  );
}