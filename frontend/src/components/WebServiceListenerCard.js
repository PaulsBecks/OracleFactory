import { Card } from "semantic-ui-react";
import Identicon from "react-identicons";
import { useHistory } from "react-router-dom";
import config from "../config";

export default function WebServiceListenerCard({
  webServiceListener,
  onClick,
}) {
  const history = useHistory();
  if (!webServiceListener) {
    return "";
  }

  return (
    <Card
      onClick={() => {
        if (onClick) {
          onClick();
        } else {
          history.push("/webServiceListeners/" + webServiceListener.ID);
        }
      }}
    >
      <Card.Content>
        <div style={{ float: "right" }}>
          <Identicon string={"Listener" + webServiceListener.ID} size={50} />
        </div>
        <Card.Header>{webServiceListener.ListenerPublisher.Name}</Card.Header>
        <Card.Meta>
          {config.BASE_URL +
            "/webServiceListeners/" +
            webServiceListener.ID +
            "/events"}
        </Card.Meta>
      </Card.Content>
      <Card.Content extra>
        <Card.Description>
          {webServiceListener.ListenerPublisher.Description}
        </Card.Description>
      </Card.Content>
    </Card>
  );
}
