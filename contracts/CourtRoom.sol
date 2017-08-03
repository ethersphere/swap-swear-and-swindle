pragma solidity ^0.4.0;

import "./abstracts/CourtroomAbstract.sol";
import "./sampletoken.sol";
import "./abstracts/trialrulesabstract.sol";
import "./abstracts/registrarabstract.sol";


contract Swear is SwearAbstract {

    TrialRulesAbstract public trialRules;
    RegistrarAbstract public registrar;

    struct Case {
        address plaintiff;
        bytes32 serviceId;
        uint8 status;
        uint8 valid;
    }
    //id map to Case
    mapping(bytes32 => Case)  OpenCases;
    mapping(address => bytes32[]) public ids;

    /// @notice Swear - Swear game constructor this function is called along with
    /// the contract deployment time.
    /// @param _registrar - address of the registrar contract
    /// @param _trialRules - address of the trial specific rules contract
    /// @return WitnessAbstract - return a witness contract instance
    function Swear(address _registrar,address _trialRules) {
        registrar = RegistrarAbstract(_registrar);
        require(registrar.setSwearContractAddress(address(this)));
        trialRules = TrialRulesAbstract(_trialRules);
    }

    /// @notice _newCase - open a new case and add it to OpenCases
    ///
    /// @param _plaintiff  - the plaintiff address for the case
    /// @param _serviceId - service id related to the case
    /// @return _status - the status of the case
    function _newCase(address _plaintiff,bytes32 _serviceId,uint8 _status) private returns (bytes32 id) {
        id = sha3(_plaintiff,_serviceId, now);
        if (OpenCases[id].valid != 0)
           return 0x0;
        OpenCases[id] = Case(
            _plaintiff,
            _serviceId,
            _status,
            1
        );
        return id;
    }

    function isValid(bytes32 id) private constant returns (bool) {
        return (OpenCases[id].valid != 0);
    }

    function setStatus(bytes32 id,uint8 status) private constant  returns (bool) {
        if (msg.sender == owner)
            throw;
        OpenCases[id].status = status;
        return true;
    }

    function resolveCase(bytes32 _id) private {

        if (OpenCases[_id].status == uint8(TrialRulesAbstract.Status.UNCHALLENGED))
            throw;
        OpenCases[_id].plaintiff = 0;
        OpenCases[_id].valid = 0;
        OpenCases[_id].status = uint8(TrialRulesAbstract.Status.UNCHALLENGED);
    }

    function getCase(bytes32 _id) private returns (address plaintiff,bytes32 serviceId) {
        return (OpenCases[_id].plaintiff,OpenCases[_id].serviceId);
    }

    function verdict(bytes32 _id,uint8 status,address plaintiff) private returns(bool) {

        if (status == uint8(TrialRulesAbstract.Status.NOT_GUILTY)) {
            CaseResolved(
                _id,
                plaintiff,
                0,
                status
            );
            return false;
        }
        uint reward = trialRules.getReward();
        bool caseCompensated = registrar.compensate(plaintiff,reward);
        resolveCase(_id);
        registrar.unRegister(plaintiff);
        CaseResolved(
            _id,
            plaintiff,
            reward,
            status
        );
        return caseCompensated;
    }

    /// @notice getStatus - return the trial status of a case
    ///
    /// @param id  - case id
    /// @return  status  - the status of a case
    function getStatus(bytes32 id) public constant returns (uint8 status) {
        return OpenCases[id].status;
    }

    /// @notice newCase - open a new case for a service id
    ///
    /// the function require that the msg sender is already register to the game.
    /// @param serviceId  - service id
    /// @return bool - true for successful operation.
    function newCase(bytes32 serviceId) public returns (bool) {

        require(registrar.isRegistered(msg.sender));
        bytes32 id = _newCase(msg.sender,serviceId,uint8(trialRules.getInitialStatus()));
        if (id == 0x0)
            return false;
        registrar.incrementOpenCases(msg.sender);
        registrar.incrementOpenCases(owner);
        ids[msg.sender].push(id);

        return true;
    }

    /// @notice trial - initiate or restart a trial proccess for a certain case
    ///
    /// the function requiere that the case is a valid one.
    /// @param id  - case id
    /// @return bool - true for successful operation.
    function trial(bytes32 id) public returns (bool) {

        require(registrar.isRegistered(msg.sender));
        require(isValid(id));
        _trial(id);
        return true;
    }

    function proceed() private returns (WitnessAbstract.Status) {
        return WitnessAbstract.Status.PENDING;
    }

    function _trial(bytes32 id) private {

        uint8 status = getStatus(id);
        var(plaintiff,serviceId) = getCase(id);
        while (status != uint8(TrialRulesAbstract.Status.UNCHALLENGED)) {
            WitnessAbstract witness = trialRules.getWitness(status);
            WitnessAbstract.Status outcome;
            if (witness == WitnessAbstract(0x0)) {
                outcome = proceed();
                return;
                } else {
                bool expired = trialRules.expired(id,status);
                if (witness.isEvidenceSubmitted(id,serviceId,plaintiff) && !expired) {
                    outcome = witness.testimonyFor(id,serviceId,plaintiff);
                    }else {
                    if (trialRules.startGracePeriod(id,status)||(!expired)) {
                        outcome = WitnessAbstract.Status.PENDING;
                        }else {
                        outcome = WitnessAbstract.Status.INVALID;
                        }
                    }
                }
            if (outcome == WitnessAbstract.Status.PENDING) {
                return;
                }
            status = trialRules.getStatus(uint8(outcome),status);
            setStatus(id,status);
            if ((status == uint8(TrialRulesAbstract.Status.GUILTY))||
                (status == uint8(TrialRulesAbstract.Status.NOT_GUILTY))){
                verdict(id,status,plaintiff);
                status = uint8(TrialRulesAbstract.Status.UNCHALLENGED);
                registrar.decrementOpenCases(plaintiff);
                registrar.decrementOpenCases(owner);
                setStatus(id,status);
                }
        }
    }

    event Decision(string decide);
    event NewCaseOpened(bytes32 id, address plaintiff);
    event CaseResolved(bytes32 id, address plaintiff, uint reward,uint8 status);

}
