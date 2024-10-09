#!/usr/bin/env sh

# Input parameters
ARCH=$(uname -m)

# Build kiichaind
echo "Building kiichaind from local branch"
git config --global --add safe.directory /kiichain-protocol/kiichain-chain
LEDGER_ENABLED=false
make install
mkdir -p build/generated
echo "DONE" > build/generated/build.complete
