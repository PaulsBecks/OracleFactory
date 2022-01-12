import React, { useState } from "react";
import postData from "../services/postData";
import { Button, Form, Icon, Popup } from "semantic-ui-react";
import { useHistory } from "react-router-dom";

export default function WebServicePublisherCreate() {
  const history = useHistory();
  const [webServicePublisher, setWebServicePublisher] = useState({
    Name: "",
    Description: "",
    URL: "",
    Private: true,
  });
  const updateWebServicePublisher = (_, { value, name }) =>
    setWebServicePublisher({
      ...webServicePublisher,
      [name]: value,
    });

  const [loading, setLoading] = useState(false);
  return (
    <div>
      <Form>
        <Form.Input
          label="Name"
          name="Name"
          value={webServicePublisher.Name}
          onChange={updateWebServicePublisher}
          placeholder="A name to recognize the publisher"
        />
        <Form.Input
          label="Description"
          name="Description"
          value={webServicePublisher.Description}
          onChange={updateWebServicePublisher}
          placeholder="Describe what the publisher does"
        />
        <Form.Input
          label="URL"
          name="URL"
          value={webServicePublisher.URL}
          onChange={updateWebServicePublisher}
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
            checked={webServicePublisher.Private}
            label={webServicePublisher.Private ? "private" : "public"}
            name="Private"
            toggle
            onChange={updateWebServicePublisher}
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
          await postData(`/webServicePublishers`, webServicePublisher);
          setLoading(false);
          history.push("/smartContracts");
        }}
      />
      <br />
      <br />
    </div>
  );
}
