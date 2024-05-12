// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity ^0.8.15;

import "@forge-std/console.sol";
import {Script} from "@forge-std/Script.sol";
import {Rods} from "../src/Rods.sol";
import {MainnetDeployConfig} from "./configs/MainnetDeployConfig.sol";
import {TestnetDeployConfig} from "./configs/TestnetDeployConfig.sol";

contract DeployContract is Script {
    Rods rods;

    function run() external {
        uint256 chainId = vm.envUint("CHAIN_TO_DEPLOY_ON");
        vm.startBroadcast();
        if (chainId == 8453) {
            rods = new Rods(
                MainnetDeployConfig.admin,
                MainnetDeployConfig.custodial
            );
        } else if (chainId == 11155111) {
            rods = new Rods(
                TestnetDeployConfig.admin,
                TestnetDeployConfig.custodial
            );
        } else {
            revert("Unsupported chain");
        }
        vm.stopBroadcast();
    }
}
