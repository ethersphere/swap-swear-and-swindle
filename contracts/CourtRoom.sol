pragma solidity ^0.4.0;

import "./owned.sol";
import "./sampletoken.sol";
import "./ensabstract.sol";


contract CaseContract is Owned {

  struct Promise{
		address beneficiary;
		uint256 blockNumber;
		uint8 sig_v;
		bytes32 sig_r;
		bytes32 sig_s;
	}
	struct Case {
		bytes32 id;
		address plaintiff;
    //not have method "push" on in memory array,
    //so need  numberOfEvidences as index of the evidence array.
    uint256   numberOfEvidences;
		bytes32[]  evidence;
		Promise   servicePromise;
		uint status;

	}

	//id map to _Case
  mapping(bytes32 => Case)  OpenCases;

	function CaseContract() {
	}

	function newClaim(address _plaintiff, bytes32 _evidence) returns (bytes32 id) {

	   Case memory _case;
		_case.id = sha3(_plaintiff, _evidence, now);
    if (OpenCases[_case.id].status != 0)return 0x0;
		_case.plaintiff = _plaintiff;
		_case.evidence = new bytes32[](4);
    _case.evidence[0] = _evidence;
    _case.numberOfEvidences = 1;
		_case.status = 1;
		 OpenCases[_case.id] = _case;
		return _case.id;
	}

	function submitEvidence(bytes32 _id,bytes32 _evidence) returns (uint status) {

		// Maximum amount of evidence has been submitted
    var numberOfEvidences = OpenCases[_id].numberOfEvidences;
		require(numberOfEvidences< 4);
		OpenCases[_id].evidence[numberOfEvidences] = _evidence;
    OpenCases[_id].numberOfEvidences++;

		return OpenCases[_id].status;

	}

	function submitPromise(bytes32 _id,address beneficiary, uint256 blockNumber,
			uint8 sig_v, bytes32 sig_r, bytes32 sig_s) returns (uint status) {

		Promise  memory prm;
		prm.beneficiary = beneficiary;
		prm.blockNumber = blockNumber;
		prm.sig_v = sig_v;
		prm.sig_r = sig_r;
		prm.sig_s = sig_s;

		OpenCases[_id].servicePromise = prm;

		return OpenCases[_id].status;

	}


	function getStatus(bytes32 id) constant returns (uint status) {

		return OpenCases[id].status;

	}

	function resolveClaim(bytes32 _id) {

		if (OpenCases[_id].id == 0x0 || OpenCases[_id].status == 0) throw;
		OpenCases[_id].id = 0x0;
		OpenCases[_id].plaintiff = 0;
		OpenCases[_id].evidence =new bytes32[](0);
		OpenCases[_id].status = 0;

	}
	function getClaim(bytes32 _id) returns (address plaintiff,address beneficiary, uint256 blockNumber,
			uint8 sig_v, bytes32 sig_r, bytes32 sig_s) {
        return (OpenCases[_id].plaintiff,
                OpenCases[_id].servicePromise.beneficiary,
                OpenCases[_id].servicePromise.blockNumber,
                OpenCases[_id].servicePromise.sig_v,
                OpenCases[_id].servicePromise.sig_r,
                OpenCases[_id].servicePromise.sig_s
                );
    }

}

contract SwearGame is Owned {

	uint256 public deposit;
	uint public reward;
	uint  public playerCount;
	mapping(address => bool) public players;
	mapping(address => bytes32[]) public ids;

	CaseContract public caseContract;
	SampleToken public token;
	bytes32 public id;


	function SwearGame(address _CaseContract, address _token, uint _reward) {
		caseContract = CaseContract(_CaseContract);
		token = SampleToken(_token);
		reward = _reward;
		playerCount = 0;

	}

  /* The function without name is the default function that is called whenever anyone sends funds to a contract */
    function () payable {
        uint amount = msg.value;
    		require(token.transferFrom(owner, address(this), amount));
    		deposit += amount;
    		DepositStaked(amount, deposit);
    }

	function verdict(bytes32 _id) private returns(bool) {

	  bool caseCompensated;
    bool decision = false;

    require(caseContract.getStatus(_id) != 0);
		// Somehow come to a resolution...
		var(plaintiff,beneficiary,blockNumber,sig_v,sig_r,sig_s) = caseContract.getClaim(_id);

    decision = (validatePromise(beneficiary,blockNumber,sig_v,sig_r,sig_s) && guilty());

    if (decision == true){
	    caseCompensated = compensate(plaintiff);
			caseContract.resolveClaim(_id);
    }

		ClaimResolved(_id, plaintiff, reward);
		return caseCompensated;

	}

	function compensate(address __Caseant) private returns(bool compensated) {

		compensated = token.transferFrom(address(this), __Caseant, reward);

		require(compensated);

		deposit -=reward;

		return compensated;

	}


	function register(address _player) onlyOwner public returns (bool registered) {

		require(!players[_player]);

		if (playerCount == 0){
			require(deposit >= reward);
		}else if ((deposit / playerCount) < reward) {
			AdditionalDepositRequired(deposit);
			throw;
		}

		players[_player] = true;
		playerCount++;

		NewPlayer(_player);

		return true;

	}


	function leaveGame(address _player) onlyOwner public {

		// If the player is not registered to the game throw
		require(players[_player]);

		PlayerLeftGame(_player);

		players[_player] = false;
		playerCount--;

	}
	function newCase(address beneficiary, uint256 blockNumber,
			uint8 sig_v, bytes32 sig_r, bytes32 sig_s,bytes32 _evidence) public returns (bool) {

		require(players[msg.sender]);

		id  = caseContract.newClaim(msg.sender,bytes32(_evidence));

    if (id == 0x0) return false;

		caseContract.submitPromise(id,beneficiary, blockNumber,sig_v, sig_r, sig_s);

		ids[msg.sender].push(id);

		verdict(id);

		return true;
	}

	/// @notice Cash cheque
	///
	/// @param beneficiary beneficiary address
	/// @param blockNumber the Promise is valid until this block number
	/// @param sig_v signature parameter v
	/// @param sig_r signature parameter r
	/// @param sig_s signature parameter s
	/// The digital signature is calculated on the concatenated triplet of contract address,beneficiary and blockNumber
	function validatePromise(address beneficiary, uint256 blockNumber,
			uint8 sig_v, bytes32 sig_r, bytes32 sig_s) private returns (bool){

      //check current block number is less than the Promise blocknumber
      if (block.number >= blockNumber ) return false;
      // Check the digital signature of the Promise.
			bytes32 hash = sha3(address(this), beneficiary, blockNumber);
			if(owner != ecrecover(hash, sig_v, sig_r, sig_s)) return false;
		  return true;
	}
	//This is the specific game logic "mirror"
	//specific game "mirror"
	ENSAbstract ens;
	//client.game namehash
	bytes32 constant clientENSNameHash = 0x94c4860d894e91f2df683b61455630d721209c6265d2e80c86a1f92cab14b370;
	//reflector.game namehash
	bytes32 constant reflectorENSNameHash = 0xacd7f5ed7d93b1526477b93e6c7def60c40420a868e7f694a7671413d89bb9a5;

	address public ensAddress = 0x8163bc885c2b14478b75f178ca76f31581dc967f;
	///////////////////////////////////////////////////////////////////////////////////////////////////////
	function ensResolve(bytes32 node) private constant returns(bytes32) {
			address resolverAddress = ens.resolver(node);
			ResolverAbstract resolver = ResolverAbstract(resolverAddress);
			bytes32 content = resolver.content(node);
			return content;
	}

	function guilty() private returns(bool){
		  //check if the two nodes resolved ENS are equal
			//for each specific game the the decision should be take diffrently
			ens = ENSAbstract(ensAddress);
			if (ensResolve(clientENSNameHash)!= ensResolve(reflectorENSNameHash)){
				Decision("guilty");
			}
			else {
				Decision("not guilty");
        return false;
			}
			return true;
	}
  //Important!!! This function should is here just for testing pouposes and to enable setting that for other enss.
	//the ens address for this game ("mirror") should be hard coded in the contract
	function setENSAddress(address _ensAddr) returns(bool){
    ensAddress = _ensAddr;
		return true;
	}

	event Decision(string decide);
	event DepositStaked(uint depositAmount, uint deposit);
	event Compensate(address recipient, uint reward);
	event NewPlayer(address playerId);
	event PlayerLeftGame(address playerId);
	event NewClaimOpened(bytes32 id, address plaintiff);
	event NewEvidenceSubmitted(bytes32 id, address plaintiff);
	event ClaimResolved(bytes32 id, address plaintiff, uint reward);
	event Payment(address from,address to ,uint256 value);
	event AdditionalDepositRequired(uint256 deposit);

}
