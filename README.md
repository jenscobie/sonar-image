# sonar-image

> SonarQube machine image

## Motivation

Running code quality analysis is generally accepted to be a good idea, 
but integrating it into your development workflow or build pipeline always seems to be a bit of a faff.

This project aims to make this process easier by providing a machine image of a SonarQube 5.1 server that you
can spin up with Vagrant and easily incorporate into your local build process.

## Requirements

* [VirtualBox](https://www.virtualbox.org/wiki/Downloads)
* [Vagrant](https://www.vagrantup.com/)
* [ChefDK](https://downloads.chef.io/chef-dk/)
* [Packer](https://www.packer.io/)

## Installation

1. Install the requirements listed above
2. Run ```./go build```
3. Run ```./go deploy```
4. Run ```./go start```
5. Visit [http://192.168.50.5:9000](http://192.168.50.5:9000)

## Usage

    Usage: ./go <command>
    
    Available commands are:
        build        Build image
        deploy       Deploy image to local VM
        destroy      Destroy local VM
        start        Start the server
        stop         Stop the server

## Author

Jen Scobie (jenscobie@gmail.com)