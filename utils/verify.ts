import { run } from 'hardhat';

const verify = async (contractAddress: string, args: unknown[]): Promise<void> => {
  console.log('Verifying contract...');
  try {
    await run('verify:verify', {
      address: contractAddress,
      constructorArguments: args,
    });
  } catch (e: unknown) {
    if (e instanceof Error) {
      if (e.message.toLowerCase().includes('already verified')) {
        console.log('Already verified!');
      } else {
        console.log(e);
      }
    }
  }
};

export default verify;
