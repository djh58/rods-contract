// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity ^0.8.15;

// testing libraries
import "@forge-std/Test.sol";

// contract dependencies
import {Rods} from "../Rods.sol";
import {MainnetDeployConfig} from "../../script/configs/MainnetDeployConfig.sol";

contract RodsTest is Test {
    Rods public rods;
    uint256 mainnetFork;

    function setUp() public {
        string memory MAINNET_RPC_URL = vm.envString("RPC_MAINNET_URL");
        mainnetFork = vm.createSelectFork(MAINNET_RPC_URL);
        
    }

    
}
