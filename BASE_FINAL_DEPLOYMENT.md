# ✅ Base Network - Final Deployment

**Date**: November 25, 2025
**Network**: Base (Chain ID: 8453)
**Status**: DEPLOYED & VERIFIED

---

## 🎯 Production Contracts

### TestToken (Existing)
- **Address**: `0x239Db952bde69A15962436C6CD86FDd3b45342e4`
- **Note**: Using existing token from other repo (NOT newly deployed)
- **View**: [Basescan](https://basescan.org/address/0x239Db952bde69A15962436C6CD86FDd3b45342e4)

### PriceOracle
- **Address**: `0x4c90551763C1498aE96589202E386019655c1781`
- **Constructor Args**: [100000, 100]
- **View**: [Blockscout](https://base.blockscout.com/address/0x4c90551763C1498aE96589202E386019655c1781) | [Basescan](https://basescan.org/address/0x4c90551763C1498aE96589202E386019655c1781)

### SimpleSwapFactory ⭐
- **Address**: `0xe4620F49ebDEF146366E63B08Eb66cAe32d51c8f`
- **Token Used**: `0x239Db952bde69A15962436C6CD86FDd3b45342e4` (existing TestToken)
- **View**: [Blockscout](https://base.blockscout.com/address/0xe4620F49ebDEF146366E63B08Eb66cAe32d51c8f) | [Basescan](https://basescan.org/address/0xe4620F49ebDEF146366E63B08Eb66cAe32d51c8f)

---

## 📋 Important Notes

1. **TestToken**: We're using the EXISTING token at `0x239Db952bde69A15962436C6CD86FDd3b45342e4` that was deployed from another repo.

2. **Unused Deployments**: During setup, test tokens were deployed at these addresses but are NOT being used:
   - `0x4b87F6D09f42B2808D0E4c26584a0c966C95089B` (ignore)
   - `0x8742202C0b388e38D04Aa08218Dbd13cbB6FD9B3` (ignore)

3. **Active Factory**: The production factory is at `0xe4620F49ebDEF146366E63B08Eb66cAe32d51c8f` and correctly uses the existing token.

---

## 🚀 Quick Start

### Verify Factory Token
```bash
npx hardhat console --network base
```

```javascript
const factory = await ethers.getContractAt(
  "SimpleSwapFactory", 
  "0xe4620F49ebDEF146366E63B08Eb66cAe32d51c8f"
);
console.log("Token address:", await factory.ERC20Address());
// Should print: 0x239Db952bde69A15962436C6CD86FDd3b45342e4
```

### Deploy a SimpleSwap Instance
```javascript
const [signer] = await ethers.getSigners();
const issuer = await signer.getAddress();
const timeout = 86400; // 24 hours
const salt = ethers.utils.formatBytes32String("my-swap-1");

const tx = await factory.deploySimpleSwap(issuer, timeout, salt);
const receipt = await tx.wait();
console.log("SimpleSwap deployed!");
```

---

## 📊 Gas Costs

- PriceOracle: 311,866 gas
- SimpleSwapFactory: 1,734,210 gas
- **Total**: 2,046,076 gas

---

## 🔗 Explorer Links

### Blockscout (Primary - Has Verified Source)
- [PriceOracle](https://base.blockscout.com/address/0x4c90551763C1498aE96589202E386019655c1781)
- [SimpleSwapFactory](https://base.blockscout.com/address/0xe4620F49ebDEF146366E63B08Eb66cAe32d51c8f)

### Basescan
- [TestToken](https://basescan.org/address/0x239Db952bde69A15962436C6CD86FDd3b45342e4)
- [PriceOracle](https://basescan.org/address/0x4c90551763C1498aE96589202E386019655c1781)
- [SimpleSwapFactory](https://basescan.org/address/0xe4620F49ebDEF146366E63B08Eb66cAe32d51c8f)

---

**Configuration File**: `deploy/base/002_deploy_factory.ts` is configured to use the existing token at `0x239Db952bde69A15962436C6CD86FDd3b45342e4`
