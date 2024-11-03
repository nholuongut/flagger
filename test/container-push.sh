#!/usr/bin/env bash

set -o errexit


push () {
    echo $DOCKER_PASS | docker login -u=$DOCKER_USER --password-stdin

    if [[ -z "$CIRCLE_TAG" ]]; then
        BRANCH_COMMIT=${CIRCLE_BRANCH}-$(echo ${CIRCLE_SHA1} | head -c7);
        docker tag test/flagger:latest nholuongut/flagger:${BRANCH_COMMIT};
        docker push nholuongut/flagger:${BRANCH_COMMIT};
    else
        docker tag test/flagger:latest nholuongut/flagger:${CIRCLE_TAG};
        docker tag test/flagger-loadtester:latest nholuongut/flagger-loadtester:${CIRCLE_TAG};
        docker push nholuongut/flagger:${CIRCLE_TAG};
        docker push nholuongut/flagger-loadtester:${CIRCLE_TAG};
    fi
}

if [[ -z "$DOCKER_PASS" ]]; then
    echo "No Docker Hub credentials, skipping image push";
else
    push
fi

