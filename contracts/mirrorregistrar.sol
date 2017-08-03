pragma solidity ^0.4.0;

import "./abstracts/registrarabstract.sol";
import "./abstracts/trialrulesabstract.sol";
import "./abstracts/token.sol";
import "./sampletoken.sol";


contract MirrorRegistrar is RegistrarAbstract {

    mapping(address=>Deposit) public deposits;
    mapping(address => bool) public players;

    uint  public playerCount;
    SampleToken public token;
    TrialRulesAbstract public trialRules;
    address swearGame;
    mapping(address => uint256)  OpenCases;

    struct Deposit {
        bool inDepositPeriod;
        uint vestingPeriod;
        uint depositedAmount;
    }

    function MirrorRegistrar(address _trialRules,address _token) {
        token = SampleToken(_token);
        trialRules = TrialRulesAbstract(_trialRules);
    }

    function incrementOpenCases(address _address) {
        require(msg.sender == swearGame);
        OpenCases[_address]++;
    }

    function decrementOpenCases(address _address) {
        require(msg.sender == swearGame);
        OpenCases[_address]--;
    }

    /// @notice register - register a player to the game
    ///
    /// The function will throw if the player is already register or there is not
    /// enough deposit in the contract to ensure the player could be compensated for the
    /// case of a valid case.
    /// @param _player  - the player address
    /// @return bool registered - true for success registration.
    function register(address _player) onlyOwner public returns (bool) {

        require(!players[_player]);
        uint reward = trialRules.getReward();
        if (playerCount == 0) {
            require(deposits[owner].depositedAmount >= reward);
        }else if ((deposits[owner].depositedAmount / playerCount) < reward) {
            AdditionalDepositRequired(deposits[owner].depositedAmount);
            throw;
        }
        players[_player] = true;
        playerCount++;
        NewPlayer(_player);
        return true;
    }

    function deposit(uint epochs) payable returns (bool) {

        //A client must be register before deposit
        if (msg.sender != owner){
           require(isRegistered(msg.sender));
        }
        require(token.transferFrom(msg.sender, address(this), msg.value));
        if (deposits[msg.sender].inDepositPeriod) {
            deposits[msg.sender].depositedAmount += msg.value;
        }else {
            deposits[msg.sender] = Deposit({inDepositPeriod: true, vestingPeriod: block.number + epochs * trialRules.getEpoch(), depositedAmount: msg.value});
        }
        DepositStaked(msg.value, deposits[msg.sender].depositedAmount);
        return true;
    }

    function collectDeposit() external returns (bool) {

        require(OpenCases[msg.sender] == 0);//check that there is no open case for the specific caller.
        //Client which collect deposit is beeing un register for the game.
        if ((msg.sender !=owner)&& (!_unRegister(msg.sender))){
             return false;
        }

        if (playerCount > 0) {
            //check that there is enough deposit left in the contract for the case of a valid case compensation.
            uint reward = trialRules.getReward();
            if ((deposits[owner].depositedAmount / playerCount) < reward)
                 return false;
        }

        Deposit storage depositInfo = deposits[msg.sender];
        if (depositInfo.inDepositPeriod && depositInfo.vestingPeriod <= block.number) {
            uint toTransfer = depositInfo.depositedAmount;
            deposits[msg.sender] = Deposit(false, 0, 0);
            token.transferFrom(address(this),msg.sender, toTransfer);
            return true;
          }
        return false;
    }

    function isRegistered(address player) returns (bool) {
        return players[player];
    }

    function compensate(address _beneficiary,uint reward) returns(bool compensated) {
        require(msg.sender == swearGame);
        compensated = token.transferFrom(address(this), _beneficiary, reward);
        require(compensated);
        deposits[owner].depositedAmount -= reward;
        Compensate(_beneficiary,reward);
        return compensated;
    }

    function unRegister(address _player) {

        require(swearGame == msg.sender);
        _unRegister(_player);
    }

    function _unRegister(address _player) private returns (bool){

        require(players[_player]);
        PlayerLeftGame(_player);
        players[_player] = false;
        playerCount--;
        return true;
    }

    function setSwearContractAddress(address _swearAddress) returns(bool) {
        require(swearGame == address(0x0));
        swearGame = _swearAddress;
        return true;
    }

    event NewPlayer(address playerId);
    event AdditionalDepositRequired(uint256 deposit);
    event DepositStaked(uint depositAmount, uint deposit);
    event Compensate(address recipient, uint reward);
    event PlayerLeftGame(address playerId);
}
