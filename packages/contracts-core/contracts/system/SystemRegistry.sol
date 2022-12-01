// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ISystemRouter } from "../interfaces/ISystemRouter.sol";
import { SystemContract } from "./SystemContract.sol";
import { AbstractGuardRegistry } from "../registry/AbstractGuardRegistry.sol";
import { AbstractNotaryRegistry } from "../registry/AbstractNotaryRegistry.sol";

/**
 * @notice Shared agents registry utilities for Origin, Destination.
 * Agents are added/removed via a system call from a local BondingManager.
 */
abstract contract SystemRegistry is AbstractGuardRegistry, AbstractNotaryRegistry, SystemContract {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SYSTEM ROUTER ONLY                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Receive a system call indicating the off-chain agent needs to be slashed.
     * @dev Must be called from a local BondingManager. Therefore
     * `uint256 _rootSubmittedAt` is ignored.
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _info             Information about agent to slash
     */
    function slashAgent(
        uint256,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        AgentInfo memory _info
    ) external override onlySystemRouter onlyLocalBondingManager(_callOrigin, _caller) {
        // TODO: decide if we need to store anything, as the slashing occurred on another chain
        _beforeAgentSlashed(_info);
        if (_info.agent == Agent.Guard) {
            _removeGuard({ _guard: _info.account });
        } else if (_info.agent == Agent.Notary) {
            _removeNotary({ _origin: _info.domain, _notary: _info.account });
        } else {
            // Sanity check
            assert(false);
        }
    }

    /**
     * @notice Receive a system call indicating the list of off-chain agents needs to be synced.
     * @dev Must be called from a local BondingManager. Therefore
     * `uint256 _rootSubmittedAt` is ignored.
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _requestID        Unique ID of the sync request
     * @param _removeExisting   Whether the existing agents need to be removed first
     * @param _infos            Information about a list of agents to sync
     */
    function syncAgents(
        uint256,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256 _requestID,
        bool _removeExisting,
        AgentInfo[] memory _infos
    ) external override onlySystemRouter onlyLocalBondingManager(_callOrigin, _caller) {
        // TODO: do we need to store this in any way?
        _requestID;
        // TODO: implement removeAllGuards(), removeAllNotaries()
        _removeExisting;
        // Sync every agent status one by one
        uint256 amount = _infos.length;
        for (uint256 i = 0; i < amount; ++i) {
            _updateAgentStatus(_infos[i]);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           INTERNAL HELPERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Perform a System Call to a local BondingManager with the given `_data`.
     */
    function _callLocalBondingManager(bytes memory _data) internal {
        systemRouter.systemCall({
            _destination: _localDomain(),
            _optimisticSeconds: 0,
            _recipient: ISystemRouter.SystemEntity.BondingManager,
            _data: _data
        });
    }

    function _updateAgentStatus(AgentInfo memory _info) internal {
        if (_info.agent == Agent.Guard) {
            address guard = _info.account;
            (_info.bonded ? _addGuard : _removeGuard)(guard);
        } else if (_info.agent == Agent.Notary) {
            uint32 domain = _info.domain;
            address notary = _info.account;
            (_info.bonded ? _addNotary : _removeNotary)(domain, notary);
        } else {
            // Sanity check: unreachable code
            assert(false);
        }
    }

    // solhint-disable no-empty-blocks
    /**
     * @notice Hook that is called before the specified agent was slashed via a system call.
     */
    function _beforeAgentSlashed(AgentInfo memory _info) internal virtual {}
}