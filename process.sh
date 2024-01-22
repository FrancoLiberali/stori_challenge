#!/bin/bash

set -Eeuo pipefail

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

docker compose -f "$SCRIPT_DIR/docker/docker-compose.yml" up db -d --wait --wait-timeout 30
FILE=$1 EMAIL=$2 docker compose -f "$SCRIPT_DIR/docker/docker-compose.yml" up main --build

exit $(docker inspect stori-challenge-test-main --format='{{.State.ExitCode}}')