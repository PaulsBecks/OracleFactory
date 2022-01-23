import React, { useState } from "react";
import postData from "../services/postData";
import { Button, Form, Icon, Popup } from "semantic-ui-react";
import { useHistory } from "react-router-dom";

export default function WebServiceCreate() {
  const history = useHistory();
  const [webServiceProvider, setWebServiceProvider] = useState({
    Name: "",
    Description: "",
    Private: true,
  });
  const updateWebServiceProvider = (_, { value, name }) =>
    setWebServiceProvider({
      ...webServiceProvider,
      [name]: value,
    });

  const [loading, setLoading] = useState(false);
  return (
    <div>
      <Form>
        <Form.Input
          label="Name"
          name="Name"
          value={webServiceProvider.Name}
          onChange={updateWebServiceProvider}
          placeholder="A name to recognize the provider"
        />
        <Form.Input
          label="Description"
          name="Description"
          value={webServiceProvider.Description}
          onChange={updateWebServiceProvider}
          placeholder="Explain what the provider does"
        />
        <Form.Field>
          <label>
            Visibility{" "}
            <Popup
              content="Deactivate the toggle to share the subscription template with other users."
              trigger={<Icon name="info circle" />}
            />
          </label>
          <Form.Checkbox
            checked={webServiceProvider.Private}
            label={webServiceProvider.Private ? "private" : "public"}
            name="Private"
            toggle
            onChange={updateWebServiceProvider}
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
          await postData(`/webServiceProviders`, webServiceProvider);
          setLoading(false);
          history.push("/smartContracts");
        }}
      />
      <br />
      <br />
    </div>
  );
}
