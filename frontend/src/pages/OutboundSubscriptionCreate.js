import React, { useState } from "react";
import { useLocation, useParams } from "react-router";
import postData from "../services/postData";
import {
  InboundSubscriptionForm,
  OutboundSubscriptionForm,
} from "../components";
import { Button, Card, Message } from "semantic-ui-react";
import { Link, useHistory } from "react-router-dom";
import useWebServiceConsumers from "../hooks/useWebServiceConsumers";
import WebServiceConsumerCard from "../components/WebServiceConsumerCard";
import SmartContractProviderCard from "../components/SmartContractProviderCard";
import useSmartContractProviders from "../hooks/useSmartContractProviders";

function useQuery() {
  return new URLSearchParams(useLocation().search);
}

export default function OutboundSubscriptionCreate() {
  const history = useHistory();
  const query = useQuery();
  const [outboundSubscription, setOutboundSubscription] = useState({
    Subscription: { Name: "" },
    URI: "",
    smartContractProviderID: parseInt(query.get("smartContractProviderID")),
    webServiceConsumerID: parseInt(query.get("webServiceConsumerID")),
  });
  const [loading, setLoading] = useState(false);
  const [webServiceConsumers] = useWebServiceConsumers();
  const [smartContractProviders] = useSmartContractProviders();

  if (!outboundSubscription.smartContractProviderID) {
    return (
      <div>
        <h1>Create Outbound Subscription</h1>
        <Message>
          Choose a smart contract provider or{" "}
          <Link to="/smartContracts/create">create one here</Link>
        </Message>
        <Card.Group>
          {smartContractProviders.map((smartContractProvider) => (
            <SmartContractProviderCard
              smartContractProvider={smartContractProvider}
              onClick={() => {
                history.push({
                  pathname: "",
                  search: `?smartContractProviderID=${
                    smartContractProvider.ID
                  }${
                    outboundSubscription.webServiceConsumerID
                      ? "&webServiceConsumerID=" +
                        outboundSubscription.webServiceConsumerID
                      : ""
                  }`,
                });
                setOutboundSubscription({
                  ...outboundSubscription,
                  smartContractProviderID: smartContractProvider.ID,
                });
              }}
            />
          ))}
        </Card.Group>
      </div>
    );
  }

  if (!outboundSubscription.webServiceConsumerID) {
    return (
      <div>
        <h1>Create Outbound Subscription</h1>
        <Message>
          Choose a web service consumer or{" "}
          <Link to="/smartContracts/create">create one here</Link>
        </Message>
        <Card.Group>
          {webServiceConsumers.map((webServiceConsumer) => (
            <WebServiceConsumerCard
              webServiceConsumer={webServiceConsumer}
              onClick={() => {
                history.push({
                  pathname: "",
                  search: `?webServiceConsumerID=${webServiceConsumer.ID}${
                    outboundSubscription.smartContractProviderID
                      ? "&smartContractProviderID=" +
                        outboundSubscription.smartContractProviderID
                      : ""
                  }`,
                });
                setOutboundSubscription({
                  ...outboundSubscription,
                  webServiceConsumerID: webServiceConsumer.ID,
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
      <h1>Create Outbound Subscription</h1>
      <InboundSubscriptionForm
        inboundSubscription={outboundSubscription}
        setInboundSubscription={setOutboundSubscription}
      />
      <br />
      <Button
        loading={loading}
        basic
        negative
        content="Cancel"
        as={Link}
        to={
          "/smartContractProviders/" +
          outboundSubscription.smartContractProviderID
        }
      />
      <Button
        loading={loading}
        basic
        positive
        content="Create"
        onClick={async () => {
          setLoading(true);
          await postData(`/outboundSubscriptions`, {
            ...outboundSubscription,
          });
          setLoading(false);
          history.push(
            "/smartContractProviders/" +
              outboundSubscription.smartContractProviderID
          );
        }}
      />
    </div>
  );
}
