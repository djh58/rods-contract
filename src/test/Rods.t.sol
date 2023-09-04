// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity ^0.8.15;

// testing libraries
import "@forge-std/Test.sol";

// contract dependencies
import {Rods} from "../Rods.sol";
import {MainnetDeployConfig} from "../../script/configs/MainnetDeployConfig.sol";
import {IERC721} from "@openzeppelin/token/ERC721/IERC721.sol";

contract RodsTest is Test {
    Rods public rods;
    address admin;
    address custodial;

    function setUp() public {
        rods = new Rods(
            MainnetDeployConfig.admin,
            MainnetDeployConfig.custodial
        );
        admin = MainnetDeployConfig.admin;
        custodial = MainnetDeployConfig.custodial;
    }

    function testSetCustodialWallet() public {
        vm.startPrank(admin);
        rods.setCustodialWallet(admin);
        vm.stopPrank();
        assertEq(admin, rods.CUSTODIAL_WALLET());
    }

    function testMint() public {
        uint256 tokenId = 0;
        string memory hash = "bizzbuzz";
        vm.startPrank(admin);
        rods.mint(tokenId, hash);
        vm.stopPrank();
        assertEq(rods.ownerOf(tokenId), rods.CUSTODIAL_WALLET());
        assertEq(rods.totalCount(), 1);
        assertEq(rods.tokenURI(tokenId), hash);
    }

    function testSetUri() public {
        uint256 tokenId = 0;
        string memory hash = "bizzbuzz";
        string memory newHash = "foobar";
        vm.startPrank(admin);
        rods.mint(tokenId, hash);
        rods.setTokenUri(tokenId, newHash);
        vm.stopPrank();
        assertEq(rods.tokenURI(tokenId), newHash);
    }

}
