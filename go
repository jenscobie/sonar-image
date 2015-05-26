#!/bin/bash

set -e -u

VBoxManage -v >/dev/null 2>&1 || { echo >&2 "VirtualBox is required. Please install the latest version."; exit 1; }
vagrant -v >/dev/null 2>&1 || { echo >&2 "Vagrant is required. Please install the latest version."; exit 1; }
chef -v >/dev/null 2>&1 || { echo >&2 "Chef Development Kit is required. Please install the latest version."; exit 1; }

[[ $(vagrant plugin list) == *vagrant-vbguest* ]] || { vagrant plugin install vagrant-vbguest; }

[[ -d "provisioners/chef/berks-cookbooks" ]] || {
    chef exec berks vendor;
    mv berks-cookbooks provisioners/chef;
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

[[ $@ ]] || { helptext; exit 1; }

case "$1" in
    build) buildvm
    ;;
    deploy)  vagrant up --no-provision
    ;;
    destroy) vagrant destroy -f
    ;;
esac