#!/bin/bash

export VAULT_MASTER_TOKEN=28321E98-DBC7-47DA-A988-E0E8B56BCFEF

cd infrastructure

docker-compose up -d

until nc -z -w5 localhost 8200
do
  sleep 0.5
  echo "Wait for Vault"
done

sleep 1

secrets=$(curl -X PUT -d "{\"secret_shares\":1, \"secret_threshold\":1}" http://localhost:8200/v1/sys/init)

echo $secrets

# {"keys":["e5621a61c4e75ffb0376874b0cb1d14fdcd3275027fc500e6b02ad0aee3b45ea"],"keys_base64":["5WIaYcTnX/sDdodLDLHRT9zTJ1An/FAOawKtCu47Reo="],"root_token":"bccce153-63cc-bcf2-eed7-c3c5d62960f4"}

root_token=$(echo $secrets | jq -r '.root_token')
key_1=$(echo $secrets | jq -r '.keys[0]')

echo "ROOT_TOKEN: $root_token"
echo "FIRST_KEY: $key_1"

data=$(curl -X PUT -d "{\"key\": \"$key_1\"}" http://localhost:8200/v1/sys/unseal)

echo $data

data_sealed=$(echo $data | jq -r '.sealed')

echo "Vault sealed: $data_sealed" 

curl -X POST -H "X-Vault-Token:$root_token" -d '{"type":"approle"}' http://localhost:8200/v1/sys/auth/approle