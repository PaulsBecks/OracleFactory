import { Card } from "semantic-ui-react";
import Identicon from "react-identicons";
import { useHistory } from "react-router-dom";
import config from "../config";

export default function WebServiceConsumerCard({
  webServiceConsumer,
  onClick,
}) {
  const history = useHistory();
  if (!webServiceConsumer) {
    return "";
  }

  return (
    <Card
      onClick={() => {
        if (onClick) {
          onClick();
        } else {
          history.push("/webServiceConsumers/" + webServiceConsumer.ID);
        }
      }}
    >
      <Card.Content>
        <div style={{ float: "right" }}>
          <Identicon string={"Consumer" + webServiceConsumer.ID} size={50} />
        </div>
        <Card.Header>{webServiceConsumer.ProviderConsumer.Name}</Card.Header>
        <Card.Meta>{webServiceConsumer.Url}</Card.Meta>
      </Card.Content>
      <Card.Content extra>
        <Card.Description>
          {webServiceConsumer.ProviderConsumer.Description}
        </Card.Description>
      </Card.Content>
    </Card>
  );
}
