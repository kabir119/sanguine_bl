// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { Attestation } from "./libs/Attestation.sol";
import { TypedMemView } from "./libs/TypedMemView.sol";
import { GlobalNotaryRegistry } from "./registry/GlobalNotaryRegistry.sol";

import {
    OwnableUpgradeable
} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract AttestationCollector is GlobalNotaryRegistry, OwnableUpgradeable {
    using Attestation for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                EVENTS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    event AttestationSubmitted(address indexed updater, bytes attestation);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev All submitted Notary Attestations are stored.
     * As different Notaries might sign attestations with the same nonce,
     * but different root (meaning one of the attestations is fraudulent),
     * we need a system so store all such attestations.
     *
     * `attestationRoots` stores a list of attested roots for every (domain, nonce) pair
     * `signatures` stores a signature for every submitted (domain, nonce, root) attestation.
     * We only store the first submitted signature for such attestation.
     */
    // [homeDomain => [nonce => [roots]]]
    mapping(uint32 => mapping(uint32 => bytes32[])) internal attestationRoots;
    // [homeDomain => [nonce => [root => signature]]]
    mapping(uint32 => mapping(uint32 => mapping(bytes32 => bytes))) internal signatures;

    /// @dev We are also storing last submitted (nonce, root) attestation for every Notary.
    // [homeDomain => [notary => latestNonce]]
    mapping(uint32 => mapping(address => uint32)) public latestNonce;
    // [homeDomain => [notary => latestRoot]]
    mapping(uint32 => mapping(address => bytes32)) public latestRoot;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             UPGRADE GAP                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    uint256[46] private __GAP;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             INITIALIZER                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function initialize() external initializer {
        __Ownable_init_unchained();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Get i-th attestation for given (domain, nonce), if exists.
     * Assuming no fraud is committed, index = 0 should be used.
     * If fraud was committed, there might be more than one attestation for given (domain, nonce).
     */
    function getAttestation(
        uint32 _domain,
        uint32 _nonce,
        uint256 _index
    ) external view returns (bytes memory) {
        bytes32 root = getRoot(_domain, _nonce, _index);
        // signature always exists for a stored root
        return _formatAttestation(_domain, _nonce, root);
    }

    /**
     * @notice Get attestation for (domain, nonce, root), if exists.
     */
    function getAttestation(
        uint32 _domain,
        uint32 _nonce,
        bytes32 _root
    ) external view returns (bytes memory) {
        require(_signatureExists(_domain, _nonce, _root), "!signature");
        return _formatAttestation(_domain, _nonce, _root);
    }

    /**
     * @notice Get latest attestation for the domain.
     */
    function getLatestAttestation(uint32 _domain) external view returns (bytes memory) {
        uint256 notariesAmount = domainNotaries[_domain].length;
        require(notariesAmount != 0, "!notaries");
        uint32 _latestNonce = 0;
        bytes32 _latestRoot;
        for (uint256 i = 0; i < notariesAmount; ) {
            address notary = domainNotaries[_domain][i];
            uint32 nonce = latestNonce[_domain][notary];
            // Check latest Notary's nonce against current latest nonce
            if (nonce > _latestNonce) {
                _latestRoot = latestRoot[_domain][notary];
                _latestNonce = nonce;
            }
            unchecked {
                ++i;
            }
        }
        // Check if we found anything
        require(_latestNonce != 0, "No attestations found");
        return _formatAttestation(_domain, _latestNonce, _latestRoot);
    }

    /**
     * @notice Get latest attestation for the domain signed by given Notary.
     */
    function getLatestAttestation(uint32 _domain, address _notary)
        external
        view
        returns (bytes memory)
    {
        uint32 nonce = latestNonce[_domain][_notary];
        require(nonce != 0, "No attestations found");
        bytes32 root = latestRoot[_domain][_notary];
        return _formatAttestation(_domain, nonce, root);
    }

    /**
     * @notice Get i-th root for given (domain, nonce), if exists.
     * Assuming no fraud is committed, index = 0 should be used.
     * If fraud was committed, there might be more than one root for given (domain, nonce).
     */
    function getRoot(
        uint32 _domain,
        uint32 _nonce,
        uint256 _index
    ) public view returns (bytes32) {
        require(_index < attestationRoots[_domain][_nonce].length, "!index");
        return attestationRoots[_domain][_nonce][_index];
    }

    /**
     * @notice Get amount of attested roots for given (domain, nonce).
     * Assuming no fraud is committed, amount <= 1.
     * If amount > 1, fraud was committed.
     */
    function rootsAmount(uint32 _domain, uint32 _nonce) external view returns (uint256) {
        return attestationRoots[_domain][_nonce].length;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              OWNER ONLY                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // TODO: add/remove notaries upon bonding/unbonding

    function addNotary(uint32 _domain, address _notary) external onlyOwner {
        _addNotary(_domain, _notary);
    }

    function removeNotary(uint32 _domain, address _notary) external onlyOwner {
        _removeNotary(_domain, _notary);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          EXTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function submitAttestation(bytes memory _attestation)
        external
        returns (bool attestationStored)
    {
        (address _notary, bytes29 _view) = _checkNotaryAuth(_attestation);
        attestationStored = _storeAttestation(_notary, _view);
        if (attestationStored) {
            // Emit Event only if the Attestation was stored
            emit AttestationSubmitted(_notary, _attestation);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _formatAttestation(
        uint32 _domain,
        uint32 _nonce,
        bytes32 _root
    ) internal view returns (bytes memory) {
        return
            Attestation.formatAttestation(
                Attestation.formatAttestationData(_domain, _nonce, _root),
                signatures[_domain][_nonce][_root]
            );
    }

    function _signatureExists(
        uint32 _domain,
        uint32 _nonce,
        bytes32 _root
    ) internal view returns (bool) {
        return signatures[_domain][_nonce][_root].length > 0;
    }

    function _storeAttestation(address _notary, bytes29 _view) internal returns (bool) {
        uint32 domain = _view.attestationDomain();
        uint32 nonce = _view.attestationNonce();
        bytes32 root = _view.attestationRoot();
        require(nonce > latestNonce[domain][_notary], "Outdated attestation");
        // Don't store Attestation, if another Notary
        // have submitted the same (domain, nonce, root) before.
        if (_signatureExists(domain, nonce, root)) return false;
        latestNonce[domain][_notary] = nonce;
        latestRoot[domain][_notary] = root;
        signatures[domain][nonce][root] = _view.attestationSignature().clone();
        attestationRoots[domain][nonce].push(root);
        return true;
    }
}
