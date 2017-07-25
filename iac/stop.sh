#!/bin/bash

cd monitoring/

docker-compose down -v

cd ../infrastructure

docker-compose down -v

unset VAULT_MASTER_TOKEN
