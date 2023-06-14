FROM golang:1.19.1

WORKDIR /app

# Install PostgreSQL client
RUN apt-get update && apt-get install -y postgresql-client

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8080

ENTRYPOINT [ "go", "run",  "main.go", "NFTHistory.go", "eventTracker.go", "getAllNFTs.go"]
