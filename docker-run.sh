#!/usr/bin/env zsh

# taken from https://github.com/grafana/xk6-sql/

set -e

if [ $# -lt 1 ]; then
    echo "Usage: ./docker-run.sh <SCRIPT_NAME> [additional k6 args]"
    exit 1
fi

IMAGE_NAME=${IMAGE_NAME:="k6-for-anomaly:latest"}

docker run -v $PWD:/examples -it --rm $IMAGE_NAME run /examples/$1


# sudo docker build -t k6-for-anomaly
# sudo docker run -v /home/figaro/Desktop/k6-extensions/xk6-anomaly/examples:/examples -it --rm k6-for-anomaly:latest run /examples/lof1.js
