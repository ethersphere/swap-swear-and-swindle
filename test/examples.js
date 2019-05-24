const SwapWitness = artifacts.require("./SwapWitness");
const AckWitness = artifacts.require("./AckWitness");
const UpdateWitness = artifacts.require("./UpdateWitness");
const MerkleWitness = artifacts.require("./MerkleWitness");
const MailboxTrial = artifacts.require("./MailboxTrial");
const Swap = artifacts.require("./Swap");
const Swear = artifacts.require("./Swear");
const Swindle = artifacts.require("./Swindle");
const { keccak256, bufferToHex } = require('ethereumjs-util');

const {
  balance,
  time,
  shouldFail,
  expectEvent,
  BN,
  constants
} = require("openzeppelin-test-helpers");
const { signNote, encodeNote } = require("./swutils");
const { MerkleTree } = require("./merkleTree")

const VALID = new BN(1);
const INVALID = new BN(2);

const GUILTY = new BN(1);
const NOT_GUILTY = new BN(2);
const WITNESS_1 = new BN(3);
const WITNESS_2 = new BN(4);

const demoData = '0x68656c6c6f000000000000000000000000000000000000000000000000000000'

contract("MailboxTrial", function(accounts) {
  const [contractor, client] = accounts;

  it("should allow swap if guilty", async () => {
    let swapContractor = await Swap.new(contractor);
    let swapClient = await Swap.new(client);    
    let swindle = await Swindle.new();
    let swear = await Swear.new(swindle.address);
    let ackWitness = await AckWitness.new();
    let swapWitness = await SwapWitness.new();
    let updateWitness = await UpdateWitness.new();
    let merkleWitness = await MerkleWitness.new();
    let trial = await MailboxTrial.new(swindle.address, updateWitness.address, merkleWitness.address, swapWitness.address, ackWitness.address);

    let remark = await trial.getSwearNoteRemark(500);

    let expiry = await time.latest() + 1000

    let serviceNoteData = {
      swap: swapContractor.address,
      beneficiary: client,
      serial: 0,
      amount: 2000,
      witness: swear.address,
      validFrom: 0,
      validUntil: expiry,
      remark,
      timeout: 1000
    }
    let serviceNote = await signNote(contractor, serviceNoteData);
    let serviceNoteEncoded = await encodeNote(serviceNoteData);
    var { receipt, logs: [{ args: { caseId, trialData } }] } = await swindle.startTrial(contractor, client, trial.address, await trial.encodeInitialPayload(500))    
    
    var { receipt, logs: [{ args: { trialData } }] } = await swindle.continueTrial(caseId, trialData, await trial.encodePayloadSwearNote(serviceNoteEncoded, serviceNote.sig))
    
    let paymentNoteData = {
      swap: swapClient.address,
      beneficiary: contractor,
      serial: 0,
      amount: 500,
      witness: constants.ZERO_ADDRESS,
      validFrom: 0,
      validUntil: expiry,
      remark: serviceNote.hash,
      timeout: 1000
    };
    let paymentNote = await signNote(client, paymentNoteData);
    let paymentNoteEncoded = await encodeNote(paymentNoteData);
    var { receipt, logs: [{ args: { trialData } }] } = await swindle.continueTrial(caseId, trialData, await trial.encodePayloadSwearNote(paymentNoteEncoded, paymentNote.sig))
    
    let ackMessage = [demoData, client, expiry - 500]
    let ackEncoded = await ackWitness.encodeAckMessage(ackMessage)
    let ackHash = await ackWitness.ackHash(ackMessage)
    var { receipt, logs: [{ args: { trialData } }] } = await swindle.continueTrial(caseId, trialData, await trial.encodePayloadSwearNote(ackEncoded, await web3.eth.sign(ackHash, contractor)))
    
    const tree = new MerkleTree(['a', 'b', demoData, 'c', 'd'])
    let updateMessage = [client, expiry - 600, expiry - 400, tree.getHexRoot()]
    let updateEncoded = await updateWitness.encodeUpdateMessage(updateMessage)
    let updateHash = await updateWitness.updateHash(updateMessage)
    var { receipt, logs: [{ args: { trialData } }] } = await swindle.continueTrial(caseId, trialData, await trial.encodePayloadSwearNote(updateEncoded, await web3.eth.sign(updateHash, contractor)))
    
    await time.increase(3600*24*30)

    var { receipt, logs: [{ args: { trialData } }] } = await swindle.continueTrial(caseId, trialData, "0x")
    
    let swearData = await swear.encodePayload(trial.address, await trial.encodeInitialPayload(500), caseId)
    await swapContractor.submitNoteWithData(serviceNoteEncoded, serviceNote.sig, swearData)
    
    await time.increase(3600*24*30)
    await swapContractor.send(100000000000000)
    await swapContractor.cashNoteWithData(serviceNoteEncoded, 2000, swearData, { from: client })
  });

  it("should not allow swap if not guilty", async () => {
    let swapContractor = await Swap.new(contractor);
    let swapClient = await Swap.new(client);    
    let swindle = await Swindle.new();
    let swear = await Swear.new(swindle.address);
    let ackWitness = await AckWitness.new();
    let swapWitness = await SwapWitness.new();
    let updateWitness = await UpdateWitness.new();
    let merkleWitness = await MerkleWitness.new();
    let trial = await MailboxTrial.new(swindle.address, updateWitness.address, merkleWitness.address, swapWitness.address, ackWitness.address);

    let remark = await trial.getSwearNoteRemark(500);
    let expiry = await time.latest() + 1000

    let serviceNoteData = {
      swap: swapContractor.address,
      beneficiary: client,
      serial: 0,
      amount: 2000,
      witness: swear.address,
      validFrom: 0,
      validUntil: expiry,
      remark,
      timeout: 1000
    }
    let serviceNote = await signNote(contractor, serviceNoteData);
    let serviceNoteEncoded = await encodeNote(serviceNoteData);
    var { receipt, logs: [{ args: { caseId, trialData } }] } = await swindle.startTrial(contractor, client, trial.address, await trial.encodeInitialPayload(500))    

    var { receipt, logs: [{ args: { trialData } }] } = await swindle.continueTrial(caseId, trialData, await trial.encodePayloadSwearNote(serviceNoteEncoded, serviceNote.sig))

    let paymentNoteData = {
      swap: swapClient.address,
      beneficiary: contractor,
      serial: 0,
      amount: 500,
      witness: constants.ZERO_ADDRESS,
      validFrom: 0,
      validUntil: expiry,
      remark: serviceNote.hash,
      timeout: 1000
    };
    let paymentNote = await signNote(client, paymentNoteData);
    let paymentNoteEncoded = await encodeNote(paymentNoteData);
    var { receipt, logs: [{ args: { trialData } }] } = await swindle.continueTrial(caseId, trialData, await trial.encodePayloadSwearNote(paymentNoteEncoded, paymentNote.sig))

    let ackMessage = [demoData, client, expiry - 500]
    let ackEncoded = await ackWitness.encodeAckMessage(ackMessage)
    let ackHash = await ackWitness.ackHash(ackMessage)    
    var { receipt, logs: [{ args: { trialData } }] } = await swindle.continueTrial(caseId, trialData, await trial.encodePayloadSwearNote(ackEncoded, await web3.eth.sign(ackHash, contractor)))
    
    const tree = new MerkleTree(['a', 'b', demoData, 'c', 'd'])
    let updateMessage = [client, expiry - 600, expiry - 400, tree.getHexRoot()]
    let updateEncoded = await updateWitness.encodeUpdateMessage(updateMessage)
    let updateHash = await updateWitness.updateHash(updateMessage)
    var { receipt, logs: [{ args: { trialData } }] } = await swindle.continueTrial(caseId, trialData, await trial.encodePayloadSwearNote(updateEncoded, await web3.eth.sign(updateHash, contractor)))

    let encodedProof = await merkleWitness.encodeProof(tree.getHexProof(demoData))    
    var { receipt } = await swindle.continueTrial(caseId, trialData, encodedProof);
        
    let swearData = await swear.encodePayload(trial.address, await trial.encodeInitialPayload(500), caseId)
    await shouldFail.reverting(swapContractor.submitNoteWithData(serviceNoteEncoded, serviceNote.sig, swearData))    
  });
});