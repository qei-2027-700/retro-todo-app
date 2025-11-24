#!/bin/bash

echo "Running database migration..."
cd "$(dirname "$0")"
go run cmd/migrate/main.go
