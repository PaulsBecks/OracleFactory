import React, { useState } from "react";
import postData from "../services/postData";
import { Button, Form, Icon, Popup } from "semantic-ui-react";
import { useHistory } from "react-router-dom";

export default function WebServiceCreate() {
  const history = useHistory();
  const [webServiceListener, setWebServiceListener] = useState({
    Name: "",
    Description: "",
    Private: true,
  });
  const updateWebServiceListener = (_, { value, name }) =>
    setWebServiceListener({
      ...webServiceListener,
      [name]: value,
    });

  const [loading, setLoading] = useState(false);
  return (
    <div>
      <Form>
        <Form.Input
          label="Name"
          name="Name"
          value={webServiceListener.Name}
          onChange={updateWebServiceListener}
          placeholder="A name to recognize the oracle"
        />
        <Form.Input
          label="Description"
          name="Description"
          value={webServiceListener.Description}
          onChange={updateWebServiceListener}
          placeholder="A name to recognize the oracle"
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
            checked={webServiceListener.Private}
            label={webServiceListener.Private ? "private" : "public"}
            name="Private"
            toggle
            onChange={updateWebServiceListener}
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
          await postData(`/webServiceListeners`, webServiceListener);
          setLoading(false);
          history.push("/smartContracts");
        }}
      />
      <br />
      <br />
    </div>
  );
}
