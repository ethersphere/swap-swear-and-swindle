function shouldReturnDEFAULT_HARDDEPPOSIT_DECREASE_TIMEOUT() {

}
function shouldReturnCheques(beneficiary, expectedSerial, expectedAmount, expectedPaidOut, expectedCashTimeout) {

}
function shouldReturnHarddeposits(beneficiary, expectedAmount, expectedDecreaseTimeout, expectedCanBeDecreasedAt) {

}
function shouldReturnTotalharddeposit(expectedHardDeposit) {

}
function shouldReturnIssuer(expectedIssuer) {

}
function shouldReturnLiquidBalance(expectedLiquidBalance) {

}
function shouldReturnLiquidBalanceFor(beneficiary, expectedLiquidBalance) {

}
function submitChequeInternal() {

}
function shouldSubmitChequeIssuer(beneficiary, serial, amount, cashTimeout, beneficiarySig, from) {

}
function shouldNotSubmitChequeIssuer(beneficiary, serial, amount, cashTimeout, beneficiarySig, from, revertMessage) {

}
function shouldSubmitChequeBeneficiary(serial, amount, cashTimeout, issuerSig, from) {

}
function shouldNotSubmitChequeBeneficiary(serial, amount, cashTimeout, issuerSig, from, revertMessage) {

}
function shouldSubmitCheque(beneficiary, serial, amount, cashTimeout, issuerSig, beneficiarySig, from) {

}
function shouldNotSubmitCheque(beneficiary, serial, amount, expectedCashTimeout, issuerSig, beneficiarySig, from, revertMessage) {

}
function cashChequeInternal(beneficiaryPrincipal, beneficiaryAgent, requestPayout, beneficiarySig, expiry, calleePayout, from) {

}
function shouldCashChequeBeneficiary(beneficiaryAgent, requestPayout, from) {

}
function shouldNotCashChequeBeneficiary(beneficiaryAgent, requestPayout, from, revertMessage) {

}
function shouldCashCheque(beneficiaryPrincipal, beneficiaryAgent, requestPayout, beneficiarySig, expiry, calleePayout, from) {

}
function shouldNotCashCheque(beneficiaryPrincipal, beneficiaryAgent, requestPayout, beneficiarySig, expiry, calleePayout, from, revertMessage) {

}
function shouldPrepareDecreaseHardDeposit(beneficiary, decreaseAmount, from) {

}
function shouldNotPrepareDecreaseHardDeposit(beneficiary, decreaseAmount, from, revertMessage) {

}
function shouldDecreaseHardDeposit(beneficiary, from) {

}
function shouldNotDecreaseHardDeposit(beneficiary, from, revertMessage) {

}
function shouldIncreaseHardDeposit(beneficiary, amount, from) {

}
function shouldNotIncreaseHardDeposit(beneficiary, amount, from, revertMessage) {

}
function shouldSetCustomHardDepositDecreaseTimeout(beneficiary, decreaseTimeout, beneficiarySig, from) {

}
function shouldNotSetCustomHardDepositDecreaseTimeout(beneficiary, decreaseTimeout, beneficiarySig, from, revertMessage) {

}

function shouldWithdraw(amount, from) {

}
function shouldNotWithdraw(amount, from, revertMessage) {

}

module.exports = {
  shouldReturnDEFAULT_HARDDEPPOSIT_DECREASE_TIMEOUT,
  shouldReturnCheques,
  shouldReturnHarddeposits,
  shouldReturnTotalharddeposit,
  shouldReturnIssuer,
  shouldReturnLiquidBalance,
  shouldReturnLiquidBalanceFor,
  shouldSubmitChequeIssuer,
  shouldNotSubmitChequeIssuer,
  shouldSubmitChequeBeneficiary,
  shouldNotSubmitChequeBeneficiary,
  shouldSubmitCheque,
  shouldNotSubmitCheque,
  shouldCashChequeBeneficiary,
  shouldNotCashChequeBeneficiary,
  shouldCashCheque,
  shouldNotCashCheque,
  shouldPrepareDecreaseHardDeposit,
  shouldNotPrepareDecreaseHardDeposit,
  shouldDecreaseHardDeposit,
  shouldNotDecreaseHardDeposit,
  shouldIncreaseHardDeposit,
  shouldNotIncreaseHardDeposit,
  shouldSetCustomHardDepositDecreaseTimeout,
  shouldNotSetCustomHardDepositDecreaseTimeout,
  shouldWithdraw,
  shouldNotWithdraw,
}

