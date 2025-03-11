// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";

contract WastelandToken is ERC20, AccessControl {

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    address public owner;

    constructor() ERC20("Wasteland", "WST"){
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(ADMIN_ROLE, msg.sender);
    }

    function userBalance() public view returns (uint) {
        return address(this).balance;
    }
    

    function tradeToken(uint value) public {
       payable(owner).transfer(value);
    }
}