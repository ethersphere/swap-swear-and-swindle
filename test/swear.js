const OracleTrial = artifacts.require('./OracleTrial.sol')
const OracleWitness = artifacts.require('./OracleWitness.sol')
const Swear = artifacts.require('./Swear.sol')
const Swindle = artifacts.require('./Swindle.sol')

require('chai')
    .use(require('chai-as-promised'))
    .use(require('chai-bignumber')(web3.BigNumber))
    .should();

const { getBalance, getTime, increaseTime, expectFail, matchLogs, sign, nulladdress, computeCost } = require('./testutils')

const VALID = 1
const INVALID = 2

const GUILTY = 1
const NOT_GUILTY = 2
const WITNESS_1 = 3
const WITNESS_2 = 4

contract('OracleTrial', function(accounts) {

  it('should have the right transitions', async() => {
    const oracleTrial = await OracleTrial.deployed();

    (await oracleTrial.nextStatus(VALID, WITNESS_1)).should.bignumber.equal(WITNESS_2);
    (await oracleTrial.nextStatus(INVALID, WITNESS_1)).should.bignumber.equal(NOT_GUILTY);

    (await oracleTrial.nextStatus(VALID, WITNESS_2)).should.bignumber.equal(GUILTY);
    (await oracleTrial.nextStatus(INVALID, WITNESS_2)).should.bignumber.equal(NOT_GUILTY);
  })

  it('should have the right witnesses', async() => {
    const oracleTrial = await OracleTrial.deployed();

    const witness1 = await oracleTrial.witness1();
    const witness2 = await oracleTrial.witness2();

    (await oracleTrial.getWitness(WITNESS_1)).should.deep.equal([
      witness1,
      web3.toBigNumber(2 * 24 * 3600)
    ]);

    (await oracleTrial.getWitness(WITNESS_2)).should.deep.equal([
      witness2,
      web3.toBigNumber(2 * 24 * 3600)
    ]);
  })

})

contract('swear', function(accounts) {

  const [
    bob,
    alice,
    carol
  ] = accounts

  it('should accept an on-chain commitment', async() => {
    const swear = await Swear.deployed();
    const swindle = await Swindle.deployed();
    const oracleTrial = await OracleTrial.deployed();

    var { logs } = await swear.addCommitment(oracleTrial.address, getTime() + 2 * 30 * 24 * 3600, 316, {
      value: 100
    });

    let { commitmentHash } = logs[0].args

    await increaseTime(3 * 30 * 24 * 3600)

    var { logs } = await swear.startTrial(commitmentHash)

    let { caseId } = logs[0].args

    await expectFail(swear.withdraw(commitmentHash));
    await expectFail(swindle.endTrial(caseId));

    await OracleWitness.at(await oracleTrial.witness1()).testify(316, VALID)

    var { logs } = await swindle.continueTrial(caseId)

    matchLogs(logs, [{
      event: 'StateTransition', args: { caseId, from: WITNESS_1, to: WITNESS_2 }
    }])

    await expectFail(swindle.endTrial(caseId));

    await OracleWitness.at(await oracleTrial.witness2()).testify(316, VALID)

    var { logs } = await swindle.continueTrial(caseId)

    matchLogs(logs, [{
      event: 'StateTransition', args: { caseId, from: WITNESS_2, to: GUILTY }
    }])

    await swindle.endTrial(caseId);

    (await getBalance(swear.address)).should.bignumber.equal(0)
  })

})
