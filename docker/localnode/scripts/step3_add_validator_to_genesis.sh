#!/bin/bash

jq '.validators = []' ~/.kiichain/config/genesis.json > ~/.kiichain/config/tmp_genesis.json
cd build/generated/gentx
IDX=0
for FILE in *
do
    jq '.validators['$IDX'] |= .+ {}' ~/.kiichain/config/tmp_genesis.json > ~/.kiichain/config/tmp_genesis_step_1.json && rm ~/.kiichain/config/tmp_genesis.json
    KEY=$(jq '.body.messages[0].pubkey.key' $FILE -c)
    DELEGATION=$(jq -r '.body.messages[0].value.amount' $FILE)
    POWER=$(($DELEGATION / 1000000))
    jq '.validators['$IDX'] += {"power":"'$POWER'"}' ~/.kiichain/config/tmp_genesis_step_1.json > ~/.kiichain/config/tmp_genesis_step_2.json && rm ~/.kiichain/config/tmp_genesis_step_1.json
    jq '.validators['$IDX'] += {"pub_key":{"type":"tendermint/PubKeyEd25519","value":'$KEY'}}' ~/.kiichain/config/tmp_genesis_step_2.json > ~/.kiichain/config/tmp_genesis_step_3.json && rm ~/.kiichain/config/tmp_genesis_step_2.json
    mv ~/.kiichain/config/tmp_genesis_step_3.json ~/.kiichain/config/tmp_genesis.json
    IDX=$(($IDX+1))
done

mv ~/.kiichain/config/tmp_genesis.json ~/.kiichain/config/genesis.json
