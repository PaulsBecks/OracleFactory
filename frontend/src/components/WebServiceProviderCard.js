import { Card } from "semantic-ui-react";
import Identicon from "react-identicons";
import { useHistory } from "react-router-dom";
import config from "../config";

export default function WebServiceProviderCard({
  webServiceProvider,
  onClick,
}) {
  const history = useHistory();
  if (!webServiceProvider) {
    return "";
  }

  return (
    <Card
      onClick={() => {
        if (onClick) {
          onClick();
        } else {
          history.push("/webServiceProviders/" + webServiceProvider.ID);
        }
      }}
    >
      <Card.Content>
        <div style={{ float: "right" }}>
          <Identicon string={"Provider" + webServiceProvider.ID} size={50} />
        </div>
        <Card.Header>{webServiceProvider.ProviderConsumer.Name}</Card.Header>
        <Card.Meta>
          {config.BASE_URL +
            "/webServiceProviders/" +
            webServiceProvider.ID +
            "/events"}
        </Card.Meta>
      </Card.Content>
      <Card.Content extra>
        <Card.Description>
          {webServiceProvider.ProviderConsumer.Description}
        </Card.Description>
      </Card.Content>
    </Card>
  );
}
