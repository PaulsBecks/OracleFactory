import { Card } from "semantic-ui-react";
import Identicon from "react-identicons";
import { useHistory } from "react-router-dom";

export default function SmartContractProviderCard({ smartContractProvider }) {
  const history = useHistory();
  if (!smartContractProvider) {
    return "";
  }

  return (
    <Card
      onClick={() =>
        history.push("/smartContractProviders/" + smartContractProvider.ID)
      }
    >
      <Card.Content>
        <div style={{ float: "right" }}>
          <Identicon
            string={smartContractProvider.SmartContract.ContractAddress}
            size={50}
          />
        </div>
        <Card.Header>{smartContractProvider.ProviderConsumer.Name}</Card.Header>
        <Card.Meta>
          {smartContractProvider.SmartContract.ContractAddressSynonym} -{" "}
          {smartContractProvider.SmartContract.EventName} -{" "}
          {smartContractProvider.SmartContract.BlockchainName}
        </Card.Meta>
      </Card.Content>
      <Card.Content extra>
        <Card.Description>
          {smartContractProvider.ProviderConsumer.Description}
        </Card.Description>
      </Card.Content>
    </Card>
  );
}
