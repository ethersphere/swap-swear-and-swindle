# courtroom-contract-suite

This repository include a basic example of a courtroom contracts game.

Contracts and tests are under contract directory.
Though I tries to keep the contracts generic as possible ,the "reflector" game as describe below is currently part of the courtroom contract (courtroom.sol). On later phase it will be changes.

## reflector game
A very simple basic game and see how its fits to the defined swindle game and test various scenarios
We can do very simple game as followed

### Phase 1
a service ("reflector") which offer to reflect any data for a specific source ENS (e.g client.game) on a destination ENS (e.g reflector.game ).
-the service will deposit some fund in the swear contact to ensure the service in the case of a claim
The service client will update the client.game ENS .
If after X(5) blocks the reflector.game ENS does not reflect its client.game ENS the client will submit a new claim for the swear contract.
In this case the claim will include the client.game ENS hash.
If the client is right a refund + compensation will send back to him from the swindle contract.

### Phase 2: (Not implemented yet)
The service will charge X amount of tokens for each reflection operation .
Client will pay the service before each operation.
-In this case a client claim will include also the payment tx along with the client.game hash.


## test
go test ./contracts/ -v --run Test






