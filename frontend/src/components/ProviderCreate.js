import React, { useState } from "react";
import postData from "../services/postData";
import { Button, Form, Icon, Popup } from "semantic-ui-react";
import { useHistory } from "react-router-dom";

export default function WebServiceCreate() {
  const history = useHistory();
  const [provider, setProvider] = useState({
    Name: "",
    Description: "",
    Topic: "",
    Private: true,
  });
  const updateProvider = (_, { value, name }) =>
    setProvider({
      ...provider,
      [name]: value,
    });

  const [loading, setLoading] = useState(false);
  return (
    <div>
      <Form>
        <Form.Input
          label="Name"
          name="Name"
          value={provider.Name}
          onChange={updateProvider}
          placeholder="A name to recognize the oracle"
        />
        <Form.Input
          label="Description"
          name="Description"
          value={provider.Description}
          onChange={updateProvider}
          placeholder="A name to recognize the oracle"
        />
        <Form.Input
          label="Topic"
          name="Topic"
          value={provider.Topic}
          onChange={updateProvider}
          placeholder="/this/is/a/topic"
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
            checked={provider.Private}
            label={provider.Private ? "private" : "public"}
            name="Private"
            toggle
            onChange={updateProvider}
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
          await postData(`/providers`, provider);
          setLoading(false);
          history.push("/");
        }}
      />
      <br />
      <br />
    </div>
  );
}
