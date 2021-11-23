pragma solidity >=0.4.22 <0.9.0;


contract SimpleStorage {
    uint storedInteger;
    bool storedBool;
    uint8 storedInteger8;
    address storedAddress;
    bytes32 storedBytes32;
    bytes storedBytes;
    string storedString;

    event StoredInteger(uint storedInteger);
    event StoredAll(uint storedInteger, bool storedBool, uint8 storedInteger8, address storedAddress, bytes32 storedBytes32, bytes storedBytes, string storedString);

    event OracleFactory(string kind, string token, string topic, string filter, string callback, address smartContractAddress);
    
    function start() public {
        emit OracleFactory("subscribe", "token", "/integers", "integer > 5", "setIntegerGreaterThenFive", address(this));
    }

    function setIntegerGreaterThenFive(uint integer) public {
        storedInteger = integer;
        emit OracleFactory("unsubscribe", "token", "/integers", "", "setIntegerGreaterThenFive", address(this));
    }

    function setAll(uint integer, bool boolean, uint8 integer8, address _address, bytes32 _bytes32, bytes memory _bytes, string memory _string) public {
        storedInteger = integer;
        storedBool = boolean;
        storedInteger8 = integer8;
        storedAddress = _address;
        storedBytes32 = _bytes32;
        storedBytes = _bytes;
        storedString = _string;

        emit StoredAll(integer, boolean, integer8, _address, _bytes32, _bytes, _string);
    }
}