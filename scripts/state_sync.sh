#!/bin/bash
echo -n Chain ID:
read CHAIN_ID
echo
echo -n Your Key Name:
read KEY
echo
echo -n Your Key Password:
read PASSWORD
echo
echo -n Moniker:
read MONIKER
echo
echo -n "Seid version commit (e.g. 5f3f795fc612c796a83726a0bbb46658586ca8fe):"
read REQUIRED_COMMIT
echo
echo -n State Sync RPC Endpoint:
read STATE_SYNC_RPC
echo
echo -n State Sync Peer:
read STATE_SYNC_PEER
echo

COMMIT=$(kiichaind version --long | grep commit)
COMMITPARTS=($COMMIT)
if [ ${COMMITPARTS[1]} != $REQUIRED_COMMIT ]
then
  echo "incorrect kiichaind version"
  exit 1
fi
mkdir $HOME/key_backup
printf ""$PASSWORD"\n"$PASSWORD"\n"$PASSWORD"\n" | kiichaind keys export $KEY > $HOME/key_backup/key
cp $HOME/.kiichain/config/priv_validator_key.json $HOME/key_backup
cp $HOME/.kiichain/data/priv_validator_state.json $HOME/key_backup
mkdir $HOME/.sei_backup
mv $HOME/.kiichain/config $HOME/.sei_backup
mv $HOME/.kiichain/data $HOME/.sei_backup
mv $HOME/.kiichain/wasm $HOME/.sei_backup
cd $HOME/.kiichain && ls | grep -xv "cosmovisor" | xargs rm -rf
kiichaind tendermint unsafe-reset-all
kiichaind init --chain-id $CHAIN_ID $MONIKER
printf ""$PASSWORD"\n"$PASSWORD"\n"$PASSWORD"\n" | kiichaind keys import $KEY $HOME/key_backup/key
cp $HOME/key_backup/priv_validator_key.json $HOME/.kiichain/config/
cp $HOME/key_backup/priv_validator_state.json $HOME/.kiichain/data/
LATEST_HEIGHT=$(curl -s $STATE_SYNC_RPC/block | jq -r .result.block.header.height); \
BLOCK_HEIGHT=$((LATEST_HEIGHT - 1000)); \
TRUST_HASH=$(curl -s "$STATE_SYNC_RPC/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)
sed -i.bak -E "s|^(enable[[:space:]]+=[[:space:]]+).*$|\1true| ; \
s|^(rpc_servers[[:space:]]+=[[:space:]]+).*$|\1\"$STATE_SYNC_RPC,$STATE_SYNC_RPC\"| ; \
s|^(trust_height[[:space:]]+=[[:space:]]+).*$|\1$BLOCK_HEIGHT| ; \
s|^(trust_hash[[:space:]]+=[[:space:]]+).*$|\1\"$TRUST_HASH\"|" $HOME/.kiichain/config/config.toml
sed -i.bak -e "s|^persistent_peers *=.*|persistent_peers = \"$STATE_SYNC_PEER\"|" \
  $HOME/.kiichain/config/config.toml
