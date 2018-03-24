#!/usr/bin/env bash
# Deploy the static website to Google App Engine.

name=grafana

OLD_GOPATH=${GOPATH}
export GOPATH=/tmp/gopath_static_${name}

echo GOROOT=${GOROOT}
echo GOPATH=${GOPATH}
echo OLD_GOPATH=${OLD_GOPATH}

echo ========= ${name} ========= "rm -rf ${GOPATH}/src/*"
rm -rf ${GOPATH}/src/*

echo ========= ${name} ========= "mkdir -p ${GOPATH}/src/"
mkdir -p ${GOPATH}/src/

echo ========= ${name} ========= "./scripts/go-get-static.sh"
./scripts/go-get-static.sh

# Set GCLOUD_PROJECT in .env
export GCLOUD_PROJECT=
. ../../../.env

# Deploy to Google Cloud.
echo ========= ${name} ========= gcloud config set project ${GCLOUD_PROJECT}
gcloud config set project ${GCLOUD_PROJECT}

echo ========= ${name} ========= gcloud config list project
gcloud config list project

echo ========= ${name} ========= gcloud app deploy static.yaml --quiet --verbosity=info
gcloud app deploy static.yaml --quiet --verbosity=info

echo ========= ${name} Deployed! =========
