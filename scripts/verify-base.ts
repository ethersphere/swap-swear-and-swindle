import { run } from 'hardhat';

async function main() {
  console.log('Starting Base contract verification...\n');

  const contracts = [
    {
      name: 'TestToken',
      address: '0x4b87F6D09f42B2808D0E4c26584a0c966C95089B',
      constructorArguments: [],
      contract: 'contracts/TestToken.sol:TestToken',
    },
    {
      name: 'PriceOracle',
      address: '0x4c90551763C1498aE96589202E386019655c1781',
      constructorArguments: [100000, 100],
      contract: 'contracts/PriceOracle.sol:PriceOracle',
    },
    {
      name: 'SimpleSwapFactory',
      address: '0xe9339c7dbccBB0C48c1BBFFD07B3534719000160',
      constructorArguments: ['0x4b87F6D09f42B2808D0E4c26584a0c966C95089B'],
      contract: 'contracts/SimpleSwapFactory.sol:SimpleSwapFactory',
    },
  ];

  for (const contractInfo of contracts) {
    console.log(`\n${'='.repeat(60)}`);
    console.log(`Verifying ${contractInfo.name}...`);
    console.log(`Address: ${contractInfo.address}`);
    console.log(`${'='.repeat(60)}\n`);

    try {
      await run('verify:verify', {
        address: contractInfo.address,
        constructorArguments: contractInfo.constructorArguments,
        contract: contractInfo.contract,
      });
      console.log(`✅ ${contractInfo.name} verified successfully!`);
    } catch (error: any) {
      if (error.message.toLowerCase().includes('already verified')) {
        console.log(`✅ ${contractInfo.name} is already verified!`);
      } else {
        console.log(`❌ Error verifying ${contractInfo.name}:`);
        console.log(error.message);
      }
    }
  }

  console.log('\n' + '='.repeat(60));
  console.log('Verification process completed!');
  console.log('='.repeat(60));
  console.log('\nView contracts on Basescan:');
  contracts.forEach((c) => {
    console.log(`${c.name}: https://basescan.org/address/${c.address}#code`);
  });
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
