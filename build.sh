#!/bin/bash

set -e

make clean
cd rust
rm -f Cargo.lock
cd ..
make rockxcrypto
go mod tidy
