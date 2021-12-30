pragma solidity >=0.4.22 <0.9.0;

interface IntegerCallback {
    function integerCallback(string calldata topic, uint256 value) external;
}

interface PubSubOracle {
    function subscribeInteger(string calldata topic, address smartContract) external;
    function unsubscribeInteger(string calldata topic, address smartContract) external;
}

contract SimpleStorage2 is IntegerCallback{
    uint storedInteger;
    bool storedBool;
    uint8 storedInteger8;
    address storedAddress;
    bytes32 storedBytes32;
    bytes storedBytes;
    string storedString;

    address onChainOracleAddress = 0xe4EFfB267484Cd790143484de3Bae7fDfbE31F00;

    event StoredInteger(uint storedInteger);
    event StoredAll(uint storedInteger, bool storedBool, uint8 storedInteger8, address storedAddress, bytes32 storedBytes32, bytes storedBytes, string storedString);
    event PublishInteger(string topic, uint256 value);

    event OracleFactory(string kind, string token, string topic, string filter, string callback, address smartContractAddress);
    
    function integerCallback(string memory topic, uint256 value) public{
        storedInteger = value;
    }
    
    function subscribeOnChain(string memory topic) public{
        PubSubOracle(onChainOracleAddress).subscribeInteger(topic, address(this));
    }

    function unsubscribeOnChain(string memory topic) public{
        PubSubOracle(onChainOracleAddress).unsubscribeInteger(topic, address(this));
    }

    function subscribeOffChain(string memory topic) public {
        emit OracleFactory("subscribe", topic, "/integers", "integer > 5", "setIntegerGreaterThenFive", address(this));
    }

    function unsubscribeOffChain(string memory topic) public {
        emit OracleFactory("unsubscribe", topic, "/integers", "", "setIntegerGreaterThenFive", address(this));
    }

    function setIntegerGreaterThenFive(uint integer) public {
        storedInteger = integer;
        //emit OracleFactory("unsubscribe", "token", "/integers", "", "setIntegerGreaterThenFive", address(this));
    }

    /*function setAll(uint integer, bool boolean, uint8 integer8, address _address, bytes32 _bytes32, bytes memory _bytes, string memory _string) public {
        storedInteger = integer;
        storedBool = boolean;
        storedInteger8 = integer8;
        storedAddress = _address;
        storedBytes32 = _bytes32;
        storedBytes = _bytes;
        storedString = _string;

        emit StoredAll(integer, boolean, integer8, _address, _bytes32, _bytes, _string);
    }*/
}