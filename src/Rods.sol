// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity ^0.8.15;

/**
 * @title Rods contract NFT
 * @author Dan Hepworth
 */

import "@openzeppelin/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/access/AccessControl.sol";


contract Rods is ERC721URIStorage, AccessControl {
    
    /**
     * @dev this role can mint NFTs, update URIs, and the custodial wallet
     * @notice set to deployer and address in constructor
     */
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    /// @dev total number of NFTs minted, as the tokenIds are set by minter
    uint256 public totalCount;

    /// @dev holds all minted assets for now
    address public CUSTODIAL_WALLET;

    /**
     * @dev admin and custodial wallet must be specified at deployment time
     * @notice the motivation behind granting the deployer and admin wallet all permissions is that the admin is likely not techincal, giving the deployer the ability to intervene. this can and should be revoked though eventually
     */
    constructor(
        address _adminWallet,
        address _custodialWallet
    ) ERC721("Rods by Dru NFTs", "RODZ") {
        require(address(0) != _custodialWallet, "Rods::constructor: Cannot set custodial wallet to zero address");
        require(address(0) != _adminWallet, "Rods::constructor: Cannot set admin wallet to zero address");
        CUSTODIAL_WALLET = _custodialWallet;
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _setupRole(DEFAULT_ADMIN_ROLE, _adminWallet);
        _setupRole(ADMIN_ROLE, msg.sender);
        _setupRole(ADMIN_ROLE, _adminWallet);
    }

    // @dev this is just needed for it to compile
    function supportsInterface(bytes4 interfaceId) public view virtual override(ERC721, AccessControl) returns (bool) {
        return super.supportsInterface(interfaceId);
    }

    /// @dev updates custodial wallet
    function setCustodialWallet(address _custodialWallet) external onlyRole(ADMIN_ROLE) {
        require(CUSTODIAL_WALLET != _custodialWallet, "Rods::setCustodialWallet: This is already the custodial wallet");
        require(address(0) != _custodialWallet, "Rods::setCustodialWallet: Cannot set to zero address");
        CUSTODIAL_WALLET = _custodialWallet;
    }

    /// @dev updates token uri. doesn't check if minted. should only be used if there is an ipfs / backend error
    function setTokenUri(uint256 tokenId, string memory uri) external onlyRole(ADMIN_ROLE) {
        _setTokenURI(tokenId, uri);
    }

    /// @dev mints to custodial wallet. when we want to go noncustodial would be best to use a fresh contract and airdrop/migrate/wrap assets from this contract
    function mint(uint256 tokenId, string memory uri) external onlyRole(ADMIN_ROLE) {
        _safeMint(CUSTODIAL_WALLET, tokenId);
        _setTokenURI(tokenId, uri);
        totalCount++;
    }

}
