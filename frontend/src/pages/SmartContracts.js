import useSmartContractListeners from "../hooks/useSmartContractListeners";
import useSmartContractPublishers from "../hooks/useSmartContractPublishers";
import { Link } from "react-router-dom";
import { Button, Card, Message } from "semantic-ui-react";
import SmartContractListenerCard from "../components/SmartContractListenerCard";
import SmartContractPublisherCard from "../components/SmartContractPublisherCard";
import useWebServiceListeners from "../hooks/useWebServiceListeners";
import WebServiceListenerCard from "../components/WebServiceListenerCard";
import WebServicePublisherCard from "../components/WebServicePublisherCard";
import useWebServicePublishers from "../hooks/useWebServicePublishers";
import listenerPublisher from "../images/Listener-To-Publisher.drawio.png";

function SmartContracts() {
  const [smartContractListeners] = useSmartContractListeners();
  const [smartContractPublishers] = useSmartContractPublishers();
  const [webServiceListeners] = useWebServiceListeners();
  const [webServicePublishers] = useWebServicePublishers();
  return (
    <div>
      <div>
        <Button
          basic
          primary
          icon="plus"
          content="Create Publisher or Listener"
          as={Link}
          to="/smartContracts/create"
        />
      </div>
      <br />
      <img src={listenerPublisher} alt="listener to publisher" />
      <Message info>
        {" "}
        Pick a listener or a publisher you want to create an oracle with.
      </Message>
      <div>
        <h2>Listener</h2>
        <Card.Group>
          {smartContractListeners.map((smartContractListener) => (
            <SmartContractListenerCard
              smartContractListener={smartContractListener}
            />
          ))}
          {webServiceListeners.map((webServiceListener) => (
            <WebServiceListenerCard webServiceListener={webServiceListener} />
          ))}
          {(!webServiceListeners || webServiceListeners.length === 0) &&
            (!smartContractListeners ||
              smartContractListeners.length === 0) && (
              <Message info>
                No listeners available. You have to create them first!
              </Message>
            )}
        </Card.Group>
      </div>
      <br />
      <br />
      <div>
        <h2>Publisher</h2>
        <Card.Group>
          {smartContractPublishers.map((smartContractPublisher) => (
            <SmartContractPublisherCard
              smartContractPublisher={smartContractPublisher}
            />
          ))}
          {webServicePublishers.map((webServicePublisher) => (
            <WebServicePublisherCard
              webServicePublisher={webServicePublisher}
            />
          ))}
          {(!webServicePublishers || webServicePublishers.length === 0) &&
            (!smartContractPublishers || webServicePublishers.length === 0) && (
              <Message info>
                No publishers available. You have to create them first!
              </Message>
            )}
        </Card.Group>
      </div>
    </div>
  );
}

export default SmartContracts;
