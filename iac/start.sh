#!/bin/bash

export CONSUL_MASTER_TOKEN=28321E98-DBC7-47DA-A988-E0E8B56BCFEF
export MONGO_PLAYER_LGN=player
export MONGO_PLAYER_PWD=player

cd infrastructure
docker-compose up -d
cd ..

######################################### Vault initialize ######################################### 

######################################### Vault unsealed ######################################### 

until nc -z -w5 localhost 8200
do
  sleep 0.5
  echo "Wait for Vault"
done

sleep 1

secrets=$(curl -X PUT -d "{\"secret_shares\":1, \"secret_threshold\":1}" http://localhost:8200/v1/sys/init -s)
# {"keys":["e5621a61c4e75ffb0376874b0cb1d14fdcd3275027fc500e6b02ad0aee3b45ea"],"keys_base64":["5WIaYcTnX/sDdodLDLHRT9zTJ1An/FAOawKtCu47Reo="],"root_token":"bccce153-63cc-bcf2-eed7-c3c5d62960f4"}

root_token=$(echo $secrets | jq -r '.root_token')
key_1=$(echo $secrets | jq -r '.keys[0]')

echo "ROOT_TOKEN: $root_token"
echo "FIRST_KEY: $key_1"

data=$(curl -X PUT -d "{\"key\": \"$key_1\"}" http://localhost:8200/v1/sys/unseal -s)

echo $data

data_sealed=$(echo $data | jq -r '.sealed')

echo "Vault sealed: $data_sealed" 


######################################### MongoDB initialize ######################################### 

echo "######### MongoDB initialize #########"
docker exec -it infrastructure_playerdb_1 mongo admin --eval "db.createUser({ user: '$MONGO_PLAYER_LGN', pwd: '$MONGO_PLAYER_PWD', roles: [ { role: 'userAdminAnyDatabase', db: 'admin' } ] });"
docker exec -it infrastructure_playerdb_1 mongo admin -u $MONGO_PLAYER_LGN -p $MONGO_PLAYER_PWD --eval "db.getSiblingDB('admin').createUser({ user: 'mongodb_exporter', pwd: 's3cr3tpassw0rd', roles: [ { role: 'clusterMonitor', db: 'admin' }, { role: 'read', db: 'local' } ]})"

######################################### Consul KV initialize ######################################### 

curl -X PUT --data 'gamename' http://localhost:8500/v1/kv/appconfig/appname?token=$CONSUL_MASTER_TOKEN
curl -X PUT --data '0.2' http://localhost:8500/v1/kv/appconfig/appversion?token=$CONSUL_MASTER_TOKEN


# activate approle
curl -X POST -H "X-Vault-Token:$root_token" -d '{"type":"approle"}' http://localhost:8200/v1/sys/auth/approle


######################################### Vault IaCRole ######################################### 

echo "######### Vault IaCRole #########"
# create iac-policy
curl -X POST -H "X-Vault-Token:$root_token" --data @policy/iac-policy.json http://localhost:8200/v1/sys/policy/iac-policy

# create iacrole
curl -X POST -H "X-Vault-Token:$root_token" -d '{"policies":"default, iac-policy"}' http://localhost:8200/v1/auth/approle/role/iacrole
iacrole=$(curl -X GET -H "X-Vault-Token:$root_token" http://localhost:8200/v1/auth/approle/role/iacrole/role-id -s)
iac_role_id=$(echo $iacrole | jq -r '.data.role_id')
echo "IaC role_id: $iac_role_id"

# create a secretid for iacrole
iaclogin=$(curl -X POST -H "X-Vault-Token:$root_token" http://localhost:8200/v1/auth/approle/role/iacrole/secret-id -s)
iac_secret_id=$(echo $iaclogin | jq -r '.data.secret_id')
echo "IaC secret_id: $iac_secret_id"

# login with playerrole and get token
iactoken=$(curl -X POST -d "{\"role_id\":\"$iac_role_id\",\"secret_id\":\"$iac_secret_id\"}" http://localhost:8200/v1/auth/approle/login -s)
iac_client_token=$(echo $iactoken | jq -r '.auth.client_token')
echo "Vault token for iacrole: $iac_client_token"

# create mongo secret connection
curl -X POST -H "X-Vault-Token:$iac_client_token" -d "{\"login\":\"$MONGO_PLAYER_LGN\", \"password\":\"$MONGO_PLAYER_PWD\"}" http://localhost:8200/v1/secret/playerdb

######################################### Vault PlayerRole ######################################### 

echo "######### Vault PlayerRole #########"
# create player-policy
curl -X POST -H "X-Vault-Token:$root_token" --data @policy/player-policy.json http://localhost:8200/v1/sys/policy/player-policy

# create playerrole
curl -X POST -H "X-Vault-Token:$root_token" -d '{"policies":"default, player-policy"}' http://localhost:8200/v1/auth/approle/role/playerrole
playerrole=$(curl -X GET -H "X-Vault-Token:$root_token" http://localhost:8200/v1/auth/approle/role/playerrole/role-id -s)
player_role_id=$(echo $playerrole | jq -r '.data.role_id')
echo "Player role_id: $player_role_id"

# create a secretid for playerrole
playerlogin=$(curl -X POST -H "X-Vault-Token:$root_token" http://localhost:8200/v1/auth/approle/role/playerrole/secret-id -s)
player_secret_id=$(echo $playerlogin | jq -r '.data.secret_id')
echo "Player secret_id: $player_secret_id"

# login with playerrole and get token
playertoken=$(curl -X POST -d "{\"role_id\":\"$player_role_id\",\"secret_id\":\"$player_secret_id\"}" http://localhost:8200/v1/auth/approle/login -s)
echo $playertoken
player_client_token=$(echo $playertoken | jq -r '.auth.client_token')
echo "Vault token for playerrole: $player_client_token"

export VAULT_PLAYER_TOKEN=$player_client_token

# test
test=$(curl -X GET -H "X-Vault-Token:$player_client_token" http://localhost:8200/v1/secret/playerdb -s)
test_data=$(echo $test | jq -r '.data')
echo "Test: $test_data"

######################################### Apps Launch ######################################### 

# cd apps

# docker-compose up -d

# cd ..