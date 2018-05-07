const OracleTrial = artifacts.require('./OracleTrial.sol')
const Swear = artifacts.require('./Swear.sol')
const Swindle = artifacts.require('./Swindle.sol')

contract('OracleTrial', function(accounts) {

  const VALID = 1
  const INVALID = 2

  const GUILTY = 1
  const NOT_GUILTY = 2
  const WITNESS_1 = 3
  const WITNESS_2 = 4

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
