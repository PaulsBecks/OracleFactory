import { Card } from "semantic-ui-react";
import Identicon from "react-identicons";
import { useHistory } from "react-router-dom";
import config from "../config";

export default function WebServicePublisherCard({
  webServicePublisher,
  onClick,
}) {
  const history = useHistory();
  if (!webServicePublisher) {
    return "";
  }

  return (
    <Card
      onClick={() => {
        if (onClick) {
          onClick();
        } else {
          history.push("/webServicePublishers/" + webServicePublisher.ID);
        }
      }}
    >
      <Card.Content>
        <div style={{ float: "right" }}>
          <Identicon string={"Publisher" + webServicePublisher.ID} size={50} />
        </div>
        <Card.Header>{webServicePublisher.ListenerPublisher.Name}</Card.Header>
        <Card.Meta>{webServicePublisher.Url}</Card.Meta>
      </Card.Content>
      <Card.Content extra>
        <Card.Description>
          {webServicePublisher.ListenerPublisher.Description}
        </Card.Description>
      </Card.Content>
    </Card>
  );
}
