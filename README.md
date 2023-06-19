# Backend service
### Backend service for NFT Lending Protocol
<p/> main.go - main file for run service on localhost
<h4>Package contracthandler</h4>
<p/> 1. ethUtil.go -  has the GetEthClientAndAddress() and ReadAbi() functions.
<p/> 2. eventTracker.go - has the function GetTransfersByAddress() which is called from httphandler\httphandlers.go
<p/> 3. getAllNFTs.go - has the function GetNFTs() which returns a list of NFTInfo structs.
<p/> 4. NFTHistory.go - has a function that returns a slice of structure.Event
<h4>Package database</h4>
<p/> 1. dbManagement.go - has functions to connect to the database and create the table if it doesn't exist.
<p/> 2.  getLastBlocknumber.go - has function GetLastProcessedBlockNumber which returns the last block number processed by the database.
<p/> 3. trackEvents.go - has a function TrackEvents() that takes in an ethclient.Client, a common.Address, a []byte, and a *sql.DB and returns a []structure.Event and an error.
<h4>Package httphandler</h4>
<p/> 1. eventHandlers.go - has all the event handlers for the events emitted by the contract.
<p/> 2. httpHandlers.go - has the handlers for the routes of the server.
<h4>Package structure</h4>
<p/> structures.go - has all the structures used in the project.
<h4>Web</h4>
<p/> 1. index.html - main page for service with form for getting stats and list of NFTs

# Start with docker
1. Pull database image
```
docker pull ghcr.io/neor-it/backend-service-nft-lending/postgres:latest
```
2. Pull backend service image
```
docker pull ghcr.io/neor-it/backend-service-nft-lending/backend:latest
```
3. Create network
```
docker network create backend-network
```
4. Run database
```
docker run -d --name=postgres --net=backend-network -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=root -p 5432:5432 --rm ghcr.io/neor-it/backend-service-nft-lending/postgres:latest
```
5. Run backend service (Get API_KEY from <a href="https://www.infura.io/">infura.io/</a>)
```
docker run -d --name=backend --net=backend-network -p 8080:8080 -e API_KEY=YOUR_API_KEY -e POSTGRES_HOST=postgres -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=root -e POSTGRES_DB=postgres ghcr.io/neor-it/backend-service-nft-lending/backend:latest
```

6. Open in browser <a href="http://localhost:8080">http://localhost:8080</a>

## API
1. <a href="https://www.infura.io/">infura.io/</a> - API_KEY for get access to Sepolia Testnet

# Smart Contract
Smart Contract is located in the contracts folder and is written in Solidity. It is compiled using Remix IDE and deploy to the Sepolia Testnet.

<p/>Lending Protocol Contract: https://sepolia.etherscan.io/address/0x7ed82e52689d7c542c3f8ca255cd921c6fc24e27
<p/>FakeUSDT Contract: https://sepolia.etherscan.io/address/0x45942dd3a289bf7c088b8ebe2c61465437616cad

## Smart Contract Functions
### NFTLending
<p/> setFee - a function for setting the fee for using NFTs.
<p/> getNFTs - a function that returns a list of registered NFTs.
<p/> registerNFT - a private function for adding a new NFT to the list of registered ones.
<p/> deleteNFT - a function for deleting an NFT from the list of registered ones.
<p/> purposeNFT/purposeNFTWithUSDT - a function for offering an NFT for rent to another user at a specified price and for a certain period of time.
<p/> purchaseNFT/purchaseNFTWithUSDT - a function for sending an NFT to a user after they have paid the rental fee.
<p/> cancelPurposeNFT - a function for cancelling an NFT rental offer and returning it to the owner.
<p/> returnNFT - function for returning the NFT to the owner and sending funds to the temporary owner.
<p/> withdrawAll - a function for withdrawing funds from the NFT to the owner.

### FakeUSDT
<p/> transfer - a function for transferring funds from one user to another.
<p/> balanceOf - a function for checking the balance of a user.
<p/> approve - a function for approving the transfer of funds from one user to another.
<p/> transferFrom - a function for transferring funds from one user to another after approval.
<p/> allowance - a function for checking the amount of funds approved for transfer.