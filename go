#!/bin/bash

set -e -u

[[ $(vagrant plugin list) == *vagrant-vbguest* ]] || {
    vagrant plugin install vagrant-vbguest
}

function helptext {
    echo "usage: ./go <command>"
    echo ""
    echo "Available commands are:"
    echo "    build        build image"
    echo "    deploy       deploy image to local VM"
    echo "    destroy      destroy local VM"
}

function buildvm {
    rm sonarqube.box || true
    packer build -only virtualbox-iso sonarqube.json
    vagrant box remove sonarqube || true
    vagrant box add sonarqube sonarqube.box
}

[[ $@ ]] || {
    helptext; exit 1;
}

case "$1" in
    build) buildvm
    ;;
    deploy)  vagrant up --no-provision
    ;;
    destroy) vagrant destroy -f
    ;;
esac