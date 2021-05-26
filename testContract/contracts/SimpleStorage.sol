pragma solidity >=0.4.22 <0.9.0;


contract SimpleStorage {
    uint storedData;

    event Stored(uint storedData);

    function set(uint x) public {
        storedData = x;
        emit Stored(x);
    }

    function get() public view returns (uint) {
        return storedData;
    }
}