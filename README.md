# Backend service
### Backend service for NFT Lending Protocol
<p/> 1. main.go - main file for run service on localhost or docker
<p/> 2. index.html - main page for service with form for getting stats and list of NFTs

# Start with docker
1. Pull image from docker hub
```
docker pull ghcr.io/neor-it/backend-service-nft-lending:latest
```
2. Run docker image
```
docker run -e API_KEY=YOUR_API_KEY -p 8080:8080 --name nftlending --rm ghcr.io/neor-it/backend-service-nft-lending
```
3. Open in browser <a href="http://localhost:8080">http://localhost:8080</a>

## API
1. <a href="https://www.infura.io/">infura.io/</a> - API_KEY for get access to Sepolia Testnet

# Smart Contract
Smart Contract is located in the contracts folder and is written in Solidity. It is compiled using Remix IDE and deploy to the Sepolia Testnet.

Lending Protocol Contract: https://sepolia.etherscan.io/address/0x22b63f333dB05DC4ead6c781349893378ed77F70

FakeUSDT Contract: https://sepolia.etherscan.io/address/0x45942dd3a289bf7c088b8ebe2c61465437616cad

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