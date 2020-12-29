# Doru client

Doru client provides orchestration for an fully decentralised application. It provides an endpoint for:
- cosigning transactions
- hot and cold storage on IPFS and Filecoin respectively (using Powergate)
- querying state (over The Graph node from network)

authorization happens over UCAN, issued from user;
user must request service from Doru client.

Doru client has policy for accepting to service a did (or delegated dids). The obvious one is the opening of a funded payment channel, from which the did must pay for services rendered. Alternatively, apps with a web2 authentication can connect a known user to their did, and render service without additional charge as part of their application.
