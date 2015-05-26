# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "sonarqube"

  config.vm.network :private_network, ip: "192.168.50.5"
  config.vm.network :forwarded_port, guest: 9000, host: 9000
  config.vm.network :forwarded_port, guest: 22, host: 2258

  config.ssh.forward_x11 = true
end