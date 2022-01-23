import { Card, Image } from "semantic-ui-react";
import Identicon from "react-identicons";
import { useHistory } from "react-router-dom";

export default function SmartContractConsumerCard({
  smartContractConsumer,
  onClick,
}) {
  const history = useHistory();
  if (!smartContractConsumer) {
    return "";
  }

  return (
    <Card
      onClick={() => {
        if (onClick) {
          onClick();
        } else {
          history.push("/smartContractConsumers/" + smartContractConsumer.ID);
        }
      }}
    >
      <Card.Content>
        <div style={{ float: "right" }}>
          <Identicon
            string={smartContractConsumer.SmartContract.ContractAddress}
            size={50}
          />
        </div>
        <Card.Header>{smartContractConsumer.ProviderConsumer.Name}</Card.Header>
        <Card.Meta>
          {smartContractConsumer.SmartContract.EventName} -{" "}
          {smartContractConsumer.SmartContract.ContractAddressSynonym} -{" "}
          {smartContractConsumer.SmartContract.BlockchainName}
        </Card.Meta>
      </Card.Content>
      <Card.Content extra>
        <Card.Description>
          {smartContractConsumer.ProviderConsumer.Description}
        </Card.Description>
      </Card.Content>
    </Card>
  );
}
