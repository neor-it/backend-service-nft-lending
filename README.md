# backend-service-nft-lending
Backend service for smart contract

## Start with docker
1. Pull image from docker hub
```
docker docker pull ghcr.io/neor-it/backend-service-nft-lending:latest
```
2. Run docker image
```
docker run -e API_KEY=YOUR_API_KEY -p 8080:8080 --name nftlending --rm ghcr.io/neor-it/backend-service-nft-lending
```
