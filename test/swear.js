const OracleTrial = artifacts.require('./OracleTrial.sol')
const OracleWitness = artifacts.require('./OracleWitness.sol')
const Swear = artifacts.require('./Swear.sol')
const Swindle = artifacts.require('./Swindle.sol')

require('chai')
    .use(require('chai-as-promised'))
    .use(require('bn-chai')(web3.utils.BN))
    .should();

const { increaseTime, matchLogs, matchStruct, nulladdress, computeCost } = require('./testutils')
const { signCheque, signNote, signInvoice } = require('./swutils')
const { balance, time, shouldFail } = require('openzeppelin-test-helpers')

const VALID = 1
const INVALID = 2

const GUILTY = 1
const NOT_GUILTY = 2
const WITNESS_1 = 3
const WITNESS_2 = 4

contract('OracleTrial', function(accounts) {

  it('should have the right transitions', async() => {
    const oracleTrial = await OracleTrial.new();

    (await oracleTrial.nextStatus(VALID, WITNESS_1)).should.eq.BN(WITNESS_2);
    (await oracleTrial.nextStatus(INVALID, WITNESS_1)).should.eq.BN(NOT_GUILTY);

    (await oracleTrial.nextStatus(VALID, WITNESS_2)).should.eq.BN(GUILTY);
    (await oracleTrial.nextStatus(INVALID, WITNESS_2)).should.eq.BN(NOT_GUILTY);
  })

  it('should have the right witnesses', async() => {
    const oracleTrial = await OracleTrial.new();

    const witness1 = await oracleTrial.witness1();
    const witness2 = await oracleTrial.witness2();
    const w1 = await oracleTrial.getWitness(WITNESS_1)
    const w2 = await oracleTrial.getWitness(WITNESS_2)

    w1[0].should.be.equal(witness1)
    w1[1].should.eq.BN(2 * 24 * 3600)

    w2[0].should.be.equal(witness2)
    w2[1].should.eq.BN(2 * 24 * 3600)
  })

})

contract('swear', function(accounts) {

  const [
    bob,
    alice,
    carol
  ] = accounts

  it('should accept an on-chain commitment', async() => {
    const swindle = await Swindle.new();
    const swear = await Swear.new(swindle.address);
    const oracleTrial = await OracleTrial.new();

    var { logs } = await swear.addCommitment(oracleTrial.address, (await time.latest()).addn(2 * 30 * 24 * 3600), "0xff", {
      value: 100
    });

    let { commitmentHash } = logs[0].args

    await increaseTime(3 * 30 * 24 * 3600)

    var { logs } = await swear.startTrial(commitmentHash, { from: alice })

    let { caseId } = logs[0].args

    await shouldFail.reverting(swear.withdraw(commitmentHash));
    await shouldFail.reverting(swindle.endTrial(caseId));

    await (await OracleWitness.at(await oracleTrial.witness1())).testify("0xff", VALID)

    var { logs } = await swindle.continueTrial(caseId)

    matchLogs(logs, [{
      event: 'StateTransition', args: { caseId, from: WITNESS_1, to: WITNESS_2 }
    }])

    await shouldFail.reverting(swindle.endTrial(caseId));

    await (await OracleWitness.at(await oracleTrial.witness2())).testify("0xff", VALID)

    var { logs } = await swindle.continueTrial(caseId)

    matchLogs(logs, [{
      event: 'StateTransition', args: { caseId, from: WITNESS_2, to: GUILTY }
    }])

    const expectedBalanceAlice = (await balance.current(alice)).addn(100)

    await swindle.endTrial(caseId);

    (await balance.current(swear.address)).should.eq.BN(0);
    (await balance.current(alice)).should.eq.BN(expectedBalanceAlice);
  })

})
