#!/usr/bin/env bash

set -Exeo pipefail

main() {
    if [[ -z "$1" ]]
    then
        (>&2 echo '[build-release/main] Error: script requires a library name')
        exit 1
    fi

    if [[ -z "$2" ]]
    then
        (>&2 echo '[build-release/main] Error: script requires a build action, e.g. ./build-release.sh [build|lipo]')
        exit 1
    fi

    cargo build --release  2>&1

    # generate rockxcrypto.h
    HEADER_DIR="." \
        cargo test --locked build_headers --features c-headers

    # ensure header file was built
    #
    find -L . -type f -name "$1.h" | read

    # ensure the archive file was built
    #
    find -L . -type f -name "lib$1.a" | read
}

main "$@"; exit
