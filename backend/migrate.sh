#!/bin/bash

echo "Running database migration..."
cd "$(dirname "$0")"
go run migrate/migrate.go
