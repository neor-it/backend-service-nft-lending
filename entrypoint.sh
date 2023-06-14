#!/bin/bash

# Wait for PostgreSQL to be ready
until PGPASSWORD=$POSTGRES_PASSWORD psql -h $POSTGRES_HOST -U $POSTGRES_USER -c '\q' &> /dev/null; do
    echo "Waiting for PostgreSQL to be ready..."
    sleep 1
done

# Run the Go application
go run main.go NFTHistory.go eventTracker.go getAllNFTs.go
