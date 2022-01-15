import React, { useState } from "react";
import { useLocation, useParams } from "react-router";
import postData from "../services/postData";
import { InboundOracleForm } from "../components";
import { Button, Card, Message } from "semantic-ui-react";
import { Link, useHistory } from "react-router-dom";
import useWebServiceListeners from "../hooks/useWebServiceListeners";
import WebServiceListenerCard from "../components/WebServiceListenerCard";
import SmartContractPublisherCard from "../components/SmartContractPublisherCard";
import useSmartContractPublishers from "../hooks/useSmartContractPublishers";

function useQuery() {
  return new URLSearchParams(useLocation().search);
}

export default function InboundOracleCreate() {
  const history = useHistory();
  const query = useQuery();
  const [inboundOracle, setInboundOracle] = useState({
    Oracle: { Name: "" },
    URI: "",
    smartContractPublisherID: parseInt(query.get("smartContractPublisherID")),
    webServiceListenerID: parseInt(query.get("webServiceListenerID")),
  });
  const [loading, setLoading] = useState(false);
  const [webServiceListeners] = useWebServiceListeners();
  const [smartContractPublishers] = useSmartContractPublishers();
  if (!inboundOracle.webServiceListenerID) {
    return (
      <div>
        <h1>Create Inbound Oracle</h1>
        <Message>
          Choose a web service listener or{" "}
          <Link to="/smartContracts/create">create one here</Link>
        </Message>
        <Card.Group>
          {webServiceListeners.map((webServiceListener) => (
            <WebServiceListenerCard
              webServiceListener={webServiceListener}
              onClick={() => {
                history.push({
                  pathname: "",
                  search: `?webServiceListenerID=${webServiceListener.ID}${
                    inboundOracle.smartContractPublisherID
                      ? "&smartContractListenerID=" +
                        inboundOracle.smartContractPublisherID
                      : ""
                  }`,
                });
                setInboundOracle({
                  ...inboundOracle,
                  webServiceListenerID: webServiceListener.ID,
                });
              }}
            />
          ))}
        </Card.Group>
      </div>
    );
  }

  if (!inboundOracle.smartContractPublisherID) {
    return (
      <div>
        <h1>Create Inbound Oracle</h1>
        <Message>
          Choose a smart contract publisher or{" "}
          <Link to="/smartContracts/create">create one here</Link>
        </Message>
        <Card.Group>
          {smartContractPublishers.map((smartContractPublisher) => (
            <SmartContractPublisherCard
              smartContractPublisher={smartContractPublisher}
              onClick={() => {
                history.push({
                  pathname: "",
                  search: `?smartContractPublisherID=${
                    smartContractPublisher.ID
                  }${
                    inboundOracle.webServiceListenerID
                      ? "&webServiceListenerID=" +
                        inboundOracle.webServiceListenerID
                      : ""
                  }`,
                });
                setInboundOracle({
                  ...inboundOracle,
                  smartContractPublisherID: smartContractPublisher.ID,
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
      <h1>Create Inbound Oracle</h1>
      <InboundOracleForm
        inboundOracle={inboundOracle}
        setInboundOracle={setInboundOracle}
      />
      <br />
      <Button
        loading={loading}
        basic
        negative
        content="Cancel"
        as={Link}
        to={
          "/smartContractPublishers/" + inboundOracle.smartContractPublisherID
        }
      />
      <Button
        loading={loading}
        basic
        positive
        content="Create"
        onClick={async () => {
          setLoading(true);
          await postData(`/inboundOracles`, {
            ...inboundOracle,
          });
          setLoading(false);
          history.push(
            "/smartContractPublishers/" + inboundOracle.smartContractPublisherID
          );
        }}
      />
    </div>
  );
}
