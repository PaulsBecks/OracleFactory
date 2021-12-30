import { useParams } from "react-router";
import { Button } from "semantic-ui-react";
import useEthereumBlockchainConnector from "../hooks/useEthereumBlockchainConnector";
import SyntaxHighlighter from "react-syntax-highlighter";
import { atomOneDark } from "react-syntax-highlighter/dist/esm/styles/hljs";

export function EthereumBlockchainConnectorDetail() {
  const { ethereumBlockchainConnectorID } = useParams();
  const [ethereumConnector, start, stop] = useEthereumBlockchainConnector(
    ethereumBlockchainConnectorID
  );
  if (!ethereumConnector) {
    return "";
  }
  return (
    <div>
      <h2>Ethereum Connector</h2>
      <p>Ethereum Node URL: {ethereumConnector.EthereumAddress}</p>
      <p>Ethereum Private Key: {ethereumConnector.EthereumPrivateKey}</p>
      <p>
        Is Active:{" "}
        {ethereumConnector.OutboundOracle.IsActive ? "true" : "false"}
        {"     "}
        <Button
          size={"tiny"}
          content={ethereumConnector.OutboundOracle.IsActive ? "stop" : "start"}
          positive={!ethereumConnector.OutboundOracle.IsActive}
          basic
          negative={ethereumConnector.OutboundOracle.IsActive}
          onClick={() =>
            ethereumConnector.OutboundOracle.IsActive ? stop() : start()
          }
        />
      </p>
      <p>
        Is On Chain:{" "}
        {ethereumConnector.OutboundOracle.IsOnChain ? "true" : "false"}
      </p>
      {ethereumConnector.OutboundOracle.IsOnChain && (
        <p>
          Smart Contract Address:
          {ethereumConnector.OutboundOracle.PubSubOracleAddress}
        </p>
      )}

      <div>
        <h3>How to subscribe to a topic</h3>
        {ethereumConnector.IsOnChain ? (
          <SyntaxHighlighter language="javascript" style={atomOneDark}>
            {`interface IntegerCallback { 
    function integerCallback(string calldata topic, uint256 value) external; 
}

interface PubSubOracle {
    function subscribeInteger(string memory topic, address smartContract) external;
}

contract ExampleContract is IntegerCallback{

    address onChainOracle = 0x7d4C98D3edC6aa64Ff7F5cd92fC9cb640073Bc9a;

    function integerCallback(string memory topic, uint256 value) public {
            // This method will be called when a new value is published the topic
    }

    function subscribe(string memory topic) private{
        PubSubOracle(0x7d4C98D3edC6aa64Ff7F5cd92fC9cb640073Bc9a).subscribeInteger(topic, address(this));
    }

    function unsubscribe(string memory topic) private{
        PubSubOracle(0x7d4C98D3edC6aa64Ff7F5cd92fC9cb640073Bc9a).unsubscribeInteger(topic, address(this));
    }
}`}
          </SyntaxHighlighter>
        ) : (
          <SyntaxHighlighter language="javascript" style={atomOneDark}>
            {`interface IntegerCallback { 
    function integerCallback(string calldata topic, uint256 value) external; 
}
  
interface PubSubOracle {
    function subscribeInteger(string memory topic, address smartContract) external;
}

contract ExampleContract is IntegerCallback{
    event OracleFactory(string kind, string token, string topic, string filter, string callback, address smartContractAddress);

    function subscribe(string memory topic) private {
        // This method can be used to subscribe to a topic and assign filters
        emit OracleFactory("subscribe", "token", topic, "integer > 5", "setIntegerGreaterThenFive", address(this));
    }

    function unsubscribe(string memory topic) private {
        // This method can be used to unsubscribe from a topic
        emit OracleFactory("subscribe", "token", topic, "", "setIntegerGreaterThenFive", address(this));
    }

    function setIntegerGreaterThenFive(uint integer) public {
        // This method will be called when a new value is published for the topic
    }
}`}
          </SyntaxHighlighter>
        )}
      </div>
    </div>
  );
}
