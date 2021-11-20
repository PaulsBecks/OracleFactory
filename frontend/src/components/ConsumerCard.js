import { Card, Image } from "semantic-ui-react";
import Identicon from "react-identicons";
import { useHistory } from "react-router-dom";

export default function ConsumerCard({ consumer, onClick }) {
  const history = useHistory();
  if (!consumer) {
    return "";
  }

  return (
    <Card
      onClick={() => {
        if (onClick) {
          onClick();
        } else {
          history.push("/consumers/" + consumer.ID);
        }
      }}
    >
      <Card.Content>
        <div style={{ float: "right" }}>
          <Identicon
            string={consumer.SmartContract.ContractAddress}
            size={50}
          />
        </div>
        <Card.Header>{consumer.ListenerPublisher.Name}</Card.Header>
        <Card.Meta>
          Consumer - {consumer.SmartContract.EventName} -{" "}
          {consumer.SmartContract.BlockchainName}
        </Card.Meta>
      </Card.Content>
      <Card.Content extra>
        <Card.Description>
          {consumer.ListenerPublisher.Description}
        </Card.Description>
      </Card.Content>
    </Card>
  );
}
