import React, { useState } from "react";
import { useLocation, useParams } from "react-router";
import postData from "../services/postData";
import { InboundOracleForm, OutboundOracleForm } from "../components";
import { Button, Card, Message } from "semantic-ui-react";
import { Link, useHistory } from "react-router-dom";
import useWebServicePublishers from "../hooks/useWebServicePublishers";
import WebServicePublisherCard from "../components/WebServicePublisherCard";
import SmartContractListenerCard from "../components/SmartContractListenerCard";
import useSmartContractListeners from "../hooks/useSmartContractListeners";

function useQuery() {
  return new URLSearchParams(useLocation().search);
}

export default function OutboundOracleCreate() {
  const history = useHistory();
  const query = useQuery();
  const [outboundOracle, setOutboundOracle] = useState({
    Oracle: { Name: "" },
    URI: "",
    smartContractListenerID: parseInt(query.get("smartContractListenerID")),
    webServicePublisherID: parseInt(query.get("webServicePublisherID")),
  });
  const [loading, setLoading] = useState(false);
  const [webServicePublishers] = useWebServicePublishers();
  const [smartContractListeners] = useSmartContractListeners();

  if (!outboundOracle.smartContractListenerID) {
    return (
      <div>
        <h1>Create Outbound Oracle</h1>
        <Message>
          Choose a smart contract listener or{" "}
          <Link to="/smartContracts/create">create one here</Link>
        </Message>
        <Card.Group>
          {smartContractListeners.map((smartContractListener) => (
            <SmartContractListenerCard
              smartContractListener={smartContractListener}
              onClick={() => {
                history.push({
                  pathname: "",
                  search: `?smartContractListenerID=${
                    smartContractListener.ID
                  }${
                    outboundOracle.webServicePublisherID
                      ? "&webServicePublisherID=" +
                        outboundOracle.webServicePublisherID
                      : ""
                  }`,
                });
                setOutboundOracle({
                  ...outboundOracle,
                  smartContractListenerID: smartContractListener.ID,
                });
              }}
            />
          ))}
        </Card.Group>
      </div>
    );
  }

  if (!outboundOracle.webServicePublisherID) {
    return (
      <div>
        <h1>Create Outbound Oracle</h1>
        <Message>
          Choose a web service publisher or{" "}
          <Link to="/smartContracts/create">create one here</Link>
        </Message>
        <Card.Group>
          {webServicePublishers.map((webServicePublisher) => (
            <WebServicePublisherCard
              webServicePublisher={webServicePublisher}
              onClick={() => {
                history.push({
                  pathname: "",
                  search: `?webServicePublisherID=${webServicePublisher.ID}${
                    outboundOracle.smartContractListenerID
                      ? "&smartContractListenerID=" +
                        outboundOracle.smartContractListenerID
                      : ""
                  }`,
                });
                setOutboundOracle({
                  ...outboundOracle,
                  webServicePublisherID: webServicePublisher.ID,
                });
              }}
            />
          ))}
        </Card.Group>
      </div>
    );
  }

  return (
    <div>
      <h1>Create Outbound Oracle</h1>
      <InboundOracleForm
        inboundOracle={outboundOracle}
        setInboundOracle={setOutboundOracle}
      />
      <br />
      <Button
        loading={loading}
        basic
        negative
        content="Cancel"
        as={Link}
        to={"/smartContractListeners/" + outboundOracle.smartContractListenerID}
      />
      <Button
        loading={loading}
        basic
        positive
        content="Create"
        onClick={async () => {
          setLoading(true);
          await postData(`/outboundOracles`, {
            ...outboundOracle,
          });
          setLoading(false);
          history.push(
            "/smartContractListeners/" + outboundOracle.smartContractListenerID
          );
        }}
      />
    </div>
  );
}
