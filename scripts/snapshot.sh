#!/bin/bash

mkdir $HOME/kiichain-snapshot
mkdir $HOME/key_backup
# Move priv_validator_state out so it isn't used by anyone else
mv $HOME/.kiichain/data/priv_validator_state.json $HOME/key_backup
# Create backups
cd $HOME/kiichain-snapshot
tar -czf data.tar.gz -C $HOME/.kiichain data/
tar -czf wasm.tar.gz -C $HOME/.kiichain wasm/
echo "Data and Wasm snapshots created in $HOME/kiichain-snapshot"