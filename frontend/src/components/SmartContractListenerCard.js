import { Card } from "semantic-ui-react";
import Identicon from "react-identicons";
import { useHistory } from "react-router-dom";

export default function SmartContractListenerCard({ smartContractListener }) {
  const history = useHistory();
  if (!smartContractListener) {
    return "";
  }

  return (
    <Card
      onClick={() =>
        history.push("/smartContractListeners/" + smartContractListener.ID)
      }
    >
      <Card.Content>
        <div style={{ float: "right" }}>
          <Identicon
            string={smartContractListener.SmartContract.ContractAddress}
            size={50}
          />
        </div>
        <Card.Header>
          {smartContractListener.ListenerPublisher.Name}
        </Card.Header>
        <Card.Meta>
          {smartContractListener.SmartContract.ContractAddressSynonym} -{" "}
          {smartContractListener.SmartContract.EventName} -{" "}
          {smartContractListener.SmartContract.BlockchainName}
        </Card.Meta>
      </Card.Content>
      <Card.Content extra>
        <Card.Description>
          {smartContractListener.ListenerPublisher.Description}
        </Card.Description>
      </Card.Content>
    </Card>
  );
}
