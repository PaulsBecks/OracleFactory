import { Card, Image } from "semantic-ui-react";
import Identicon from "react-identicons";
import { useHistory } from "react-router-dom";

export default function SmartContractPublisherCard({
  smartContractPublisher,
  onClick,
}) {
  const history = useHistory();
  if (!smartContractPublisher) {
    return "";
  }

  return (
    <Card
      onClick={() => {
        if (onClick) {
          onClick();
        } else {
          history.push("/smartContractPublishers/" + smartContractPublisher.ID);
        }
      }}
    >
      <Card.Content>
        <div style={{ float: "right" }}>
          <Identicon
            string={smartContractPublisher.SmartContract.ContractAddress}
            size={50}
          />
        </div>
        <Card.Header>
          {smartContractPublisher.ListenerPublisher.Name}
        </Card.Header>
        <Card.Meta>
          {smartContractPublisher.SmartContract.EventName} -{" "}
          {smartContractPublisher.SmartContract.ContractAddressSynonym} -{" "}
          {smartContractPublisher.SmartContract.BlockchainName}
        </Card.Meta>
      </Card.Content>
      <Card.Content extra>
        <Card.Description>
          {smartContractPublisher.ListenerPublisher.Description}
        </Card.Description>
      </Card.Content>
    </Card>
  );
}
