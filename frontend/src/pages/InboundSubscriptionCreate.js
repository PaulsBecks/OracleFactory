import React, { useState } from "react";
import { useLocation, useParams } from "react-router";
import postData from "../services/postData";
import { InboundSubscriptionForm } from "../components";
import { Button, Card, Message } from "semantic-ui-react";
import { Link, useHistory } from "react-router-dom";
import useWebServiceProviders from "../hooks/useWebServiceProviders";
import WebServiceProviderCard from "../components/WebServiceProviderCard";
import SmartContractConsumerCard from "../components/SmartContractConsumerCard";
import useSmartContractConsumers from "../hooks/useSmartContractConsumers";

function useQuery() {
  return new URLSearchParams(useLocation().search);
}

export default function InboundSubscriptionCreate() {
  const history = useHistory();
  const query = useQuery();
  const [inboundSubscription, setInboundSubscription] = useState({
    Subscription: { Name: "" },
    URI: "",
    smartContractConsumerID: parseInt(query.get("smartContractConsumerID")),
    webServiceProviderID: parseInt(query.get("webServiceProviderID")),
  });
  const [loading, setLoading] = useState(false);
  const [webServiceProviders] = useWebServiceProviders();
  const [smartContractConsumers] = useSmartContractConsumers();
  if (!inboundSubscription.webServiceProviderID) {
    return (
      <div>
        <h1>Create Inbound Subscription</h1>
        <Message>
          Choose a web service provider or{" "}
          <Link to="/smartContracts/create">create one here</Link>
        </Message>
        <Card.Group>
          {webServiceProviders.map((webServiceProvider) => (
            <WebServiceProviderCard
              webServiceProvider={webServiceProvider}
              onClick={() => {
                history.push({
                  pathname: "",
                  search: `?webServiceProviderID=${webServiceProvider.ID}${
                    inboundSubscription.smartContractConsumerID
                      ? "&smartContractProviderID=" +
                        inboundSubscription.smartContractConsumerID
                      : ""
                  }`,
                });
                setInboundSubscription({
                  ...inboundSubscription,
                  webServiceProviderID: webServiceProvider.ID,
                });
              }}
            />
          ))}
        </Card.Group>
      </div>
    );
  }

  if (!inboundSubscription.smartContractConsumerID) {
    return (
      <div>
        <h1>Create Inbound Subscription</h1>
        <Message>
          Choose a smart contract consumer or{" "}
          <Link to="/smartContracts/create">create one here</Link>
        </Message>
        <Card.Group>
          {smartContractConsumers.map((smartContractConsumer) => (
            <SmartContractConsumerCard
              smartContractConsumer={smartContractConsumer}
              onClick={() => {
                history.push({
                  pathname: "",
                  search: `?smartContractConsumerID=${
                    smartContractConsumer.ID
                  }${
                    inboundSubscription.webServiceProviderID
                      ? "&webServiceProviderID=" +
                        inboundSubscription.webServiceProviderID
                      : ""
                  }`,
                });
                setInboundSubscription({
                  ...inboundSubscription,
                  smartContractConsumerID: smartContractConsumer.ID,
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
      <h1>Create Inbound Subscription</h1>
      <InboundSubscriptionForm
        inboundSubscription={inboundSubscription}
        setInboundSubscription={setInboundSubscription}
      />
      <br />
      <Button
        loading={loading}
        basic
        negative
        content="Cancel"
        as={Link}
        to={
          "/smartContractConsumers/" +
          inboundSubscription.smartContractConsumerID
        }
      />
      <Button
        loading={loading}
        basic
        positive
        content="Create"
        onClick={async () => {
          setLoading(true);
          await postData(`/inboundSubscriptions`, {
            ...inboundSubscription,
          });
          setLoading(false);
          history.push(
            "/smartContractConsumers/" +
              inboundSubscription.smartContractConsumerID
          );
        }}
      />
    </div>
  );
}
