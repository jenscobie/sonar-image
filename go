#!/bin/bash

set -e -u

rm sonarqube.box || true
packer build -only virtualbox-iso sonarqube.json
vagrant box remove sonarqube || true
vagrant box add sonarqube sonarqube.box