const HashWitness = artifacts.require('./HashWitness.sol')
const ChunkWitness = artifacts.require('./ChunkWitness.sol')
const Swear = artifacts.require('./SwearSwap.sol')
const Swindle = artifacts.require('./Swindle.sol')
const Swap = artifacts.require('./Swap.sol')
const SimpleTrial = artifacts.require('./SimpleTrial.sol')
const util = require('ethereumjs-util')

require('chai')
    .use(require('chai-as-promised'))
    .use(require('bn-chai')(web3.utils.BN))
    .should();

const { getTime, increaseTime, expectFail, matchLogs, sign, nulladdress, computeCost } = require('./testutils')
const { signNote } = require('./swutils')

/* Dockerfile from swarm repo */
const length = 'a701000000000000'

const data = '23204275696c64204765746820696e20612073746f636b20476f206275696c64657220636f6e7461696e65720a46524f4d20676f6c616e673a312e31302d616c70696e65206173206275696c6465720a0a52554e2061706b20616464202d2d6e6f2d6361636865206d616b6520676363206d75736c2d646576206c696e75782d686561646572730a0a414444202e202f676f2d657468657265756d0a52554e206364202f676f2d657468657265756d202626206d616b6520676574680a0a232050756c6c204765746820696e746f2061207365636f6e64207374616765206465706c6f7920616c70696e6520636f6e7461696e65720a46524f4d20616c70696e653a6c61746573740a0a52554e2061706b20616464202d2d6e6f2d63616368652063612d6365727469666963617465730a434f5059202d2d66726f6d3d6275696c646572202f676f2d657468657265756d2f6275696c642f62696e2f67657468202f7573722f6c6f63616c2f62696e2f0a0a4558504f5345203835343520383534362033303330332033303330332f7564700a454e545259504f494e54205b2267657468225d0a';

const swarmHash = '0x7ab16e71d5b49df10f4d387baeb83fcd5324d853940ba0a04f73ef3a42fe3239'
const poc3Hash = '0x2140ed1741bedbc646e644bcb809a075e1f280b3a4007846ee3b2374c65e29bf'

contract('Storage', (accounts) => {

  const [dataInsurer, dataOwner, carol] = accounts

  it('should take deposit if nothing is provided', async () => {
    const swap = await Swap.new(dataInsurer)
    const swindle = await Swindle.new()
    const swear = await Swear.new(swindle.address)
    const hashWitness = await HashWitness.new()
    const trial = await SimpleTrial.new(hashWitness.address)

    const amount = 10000

    const witness = swear.address

    await swap.send(web3.utils.toWei("1"))

    let expires = await getTime() + 3600 * 24 * 365

    let remark = '0x' + util.sha3(Buffer.concat([Buffer.from(trial.address.substring(2), 'hex'), Buffer.from(swarmHash.substring(2), 'hex')])).toString('hex')

    let note = await signNote(swap, dataInsurer, dataOwner, 1, amount, witness, 0, expires, remark)

    let encoded = await swap.encodeNote(swap.address, dataOwner, 1, amount, witness, 0, expires, remark);

    await expectFail(swap.submitNote(encoded, note.sig, { from: dataOwner }));

    var { logs } = await swear.startTrialFromNote(encoded, trial.address, swarmHash, note.sig)
    let { caseId, commitmentHash } = logs[0].args

    await increaseTime(3600 * 24 * 31)

    await swindle.continueTrial(caseId)

    await swindle.endTrial(caseId)

    await swap.submitNote(encoded, note.sig, { from: dataOwner });

    await increaseTime(3600 * 24 * 2)

    await swap.cashNote(encoded, 1, { from: dataOwner })
  })

  it('should accept valid POC2 chunk', async () => {
    const swap = await Swap.new(dataInsurer)
    const swindle = await Swindle.new()
    const swear = await Swear.new(swindle.address)
    const hashWitness = await HashWitness.new()
    const trial = await SimpleTrial.new(hashWitness.address)

    const amount = 10000

    const witness = swear.address

    await swap.send(web3.utils.toWei("1"))

    let expires = await getTime() + 3600 * 24 * 365

    let remark = '0x' + util.sha3(Buffer.concat([Buffer.from(trial.address.substring(2), 'hex'), Buffer.from(swarmHash.substring(2), 'hex')])).toString('hex')

    let note = await signNote(swap, dataInsurer, dataOwner, 1, amount, witness, 0, expires, remark)

    let encoded = await swap.encodeNote(swap.address, dataOwner, 1, amount, witness, 0, expires, remark);

    await expectFail(swap.submitNote(encoded, note.sig, { from: dataOwner }));

    var { logs } = await swear.startTrialFromNote(encoded, trial.address, swarmHash, note.sig)
    let { caseId, commitmentHash } = logs[0].args

    await hashWitness.testify('0x' + length + data)

    await swindle.continueTrial(caseId)

    await swindle.endTrial(caseId)

    await expectFail(swap.submitNote(encoded, note.sig, { from: dataOwner }));
  })

  it('should accept valid POC3 chunk', async () => {
    const swap = await Swap.new(dataInsurer)
    const swindle = await Swindle.new()
    const swear = await Swear.new(swindle.address)

    const chunkWitness = await ChunkWitness.new()
    const trial = await SimpleTrial.new(chunkWitness.address)

    const amount = 10000

    const witness = swear.address

    await swap.send(web3.utils.toWei("1"))

    let expires = await getTime() + 3600 * 24 * 365

    let remark = '0x' + util.sha3(Buffer.concat([Buffer.from(trial.address.substring(2), 'hex'), Buffer.from(poc3Hash.substring(2), 'hex')])).toString('hex')

    let note = await signNote(swap, dataInsurer, dataOwner, 1, amount, witness, 0, expires, remark)

    let encoded = await swap.encodeNote(swap.address, dataOwner, 1, amount, witness, 0, expires, remark);

    await expectFail(swap.submitNote(encoded, note.sig, { from: dataOwner }));

    var { logs } = await swear.startTrialFromNote(encoded, trial.address, poc3Hash, note.sig)
    let { caseId, commitmentHash } = logs[0].args

    await chunkWitness.testify('0x' + data, { gas: 1000000 })

    await swindle.continueTrial(caseId)

    await swindle.endTrial(caseId)

    await expectFail(swap.submitNote(encoded, note.sig, { from: dataOwner }));
  })

})
