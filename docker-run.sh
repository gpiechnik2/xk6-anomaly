#!/usr/bin/env zsh

# taken from https://github.com/grafana/xk6-sql/

set -e

if [ $# -lt 1 ]; then
    echo "Usage: ./docker-run.sh <SCRIPT_NAME> [additional k6 args]"
    exit 1
fi

IMAGE_NAME=${IMAGE_NAME:="k6-for-anomaly:latest"}

docker run -v $PWD:/scripts -it --rm $IMAGE_NAME run /scripts/$1 ${@:2}
