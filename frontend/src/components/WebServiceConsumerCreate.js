import React, { useState } from "react";
import postData from "../services/postData";
import { Button, Form, Icon, Popup } from "semantic-ui-react";
import { useHistory } from "react-router-dom";

export default function WebServiceConsumerCreate() {
  const history = useHistory();
  const [webServiceConsumer, setWebServiceConsumer] = useState({
    Name: "",
    Description: "",
    URL: "",
    Private: true,
  });
  const updateWebServiceConsumer = (_, { value, name }) =>
    setWebServiceConsumer({
      ...webServiceConsumer,
      [name]: value,
    });

  const [loading, setLoading] = useState(false);
  return (
    <div>
      <Form>
        <Form.Input
          label="Name"
          name="Name"
          value={webServiceConsumer.Name}
          onChange={updateWebServiceConsumer}
          placeholder="A name to recognize the consumer"
        />
        <Form.Input
          label="Description"
          name="Description"
          value={webServiceConsumer.Description}
          onChange={updateWebServiceConsumer}
          placeholder="Describe what the consumer does"
        />
        <Form.Input
          label="URL"
          name="URL"
          value={webServiceConsumer.URL}
          onChange={updateWebServiceConsumer}
          placeholder="The URL you want to publish to"
        />
        <Form.Field>
          <label>
            Visibility{" "}
            <Popup
              content="Deactivate the toggle to share the oracle template with other users."
              trigger={<Icon name="info circle" />}
            />
          </label>
          <Form.Checkbox
            checked={webServiceConsumer.Private}
            label={webServiceConsumer.Private ? "private" : "public"}
            name="Private"
            toggle
            onChange={updateWebServiceConsumer}
          />
        </Form.Field>
      </Form>
      <br />
      <Button
        loading={loading}
        basic
        positive
        fluid
        floated="right"
        content="Create"
        onClick={async () => {
          setLoading(true);
          await postData(`/webServiceConsumers`, webServiceConsumer);
          setLoading(false);
          history.push("/smartContracts");
        }}
      />
      <br />
      <br />
    </div>
  );
}
