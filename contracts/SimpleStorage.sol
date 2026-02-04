// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleStorage {
    uint256 private storedData;
    address public owner;
    
    event DataChanged(uint256 oldValue, uint256 newValue);
    
    modifier onlyOwner() {
        require(msg.sender == owner, "Not owner");
        _;
    }
    
    constructor() {
        owner = msg.sender;
    }
    
    function set(uint256 x) public {
        uint256 oldValue = storedData;
        storedData = x;
        emit DataChanged(oldValue, x);
    }
    
    function get() public view returns (uint256) {
        return storedData;
    }
    
    function getOwner() public view returns (address) {
        return owner;
    }
}
