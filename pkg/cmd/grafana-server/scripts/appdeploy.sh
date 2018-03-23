#!/usr/bin/env bash
# Deploy the web app to Google App Engine.

# sudo apt install golang
# sudo apt install google-cloud-sdk-app-engine-go
# echo 'export GOROOT=$HOME/go' >>~/.zshrc
# echo 'export PATH=$PATH:$GOROOT/bin' >>~/.zshrc
# echo 'export GOPATH=$GOROOT/bin' >>~/.zshrc
# go get -u google.golang.org/appengine

# Set GCLOUD_PROJECT in .env
export GCLOUD_PROJECT=
. ../../../.env

echo GOROOT=${GOROOT}
echo GOPATH=${GOPATH}

# Deploy to Google Cloud.
gcloud config set project ${GCLOUD_PROJECT}
gcloud config list project

echo ========= ${name} ========= gcloud app deploy app.yaml --quiet --verbosity=info
gcloud app deploy app.yaml --quiet --verbosity=info

echo ========= ${name} Deployed! =========
