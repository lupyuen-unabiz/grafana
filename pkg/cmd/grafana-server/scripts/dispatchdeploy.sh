#!/usr/bin/env bash
# Deploy the dispatch file to Google App Engine.

# Set GCLOUD_PROJECT in .env
export GCLOUD_PROJECT=
. ../../../.env

# Deploy to Google Cloud.
echo ========= ${name} ========= gcloud config set project ${GCLOUD_PROJECT}
gcloud config set project ${GCLOUD_PROJECT}

echo ========= ${name} ========= gcloud config list project
gcloud config list project

echo ========= ${name} ========= gcloud app deploy dispatch.yaml --quiet --verbosity=info
gcloud app deploy dispatch.yaml --quiet --verbosity=info

echo ========= ${name} Deployed! =========
