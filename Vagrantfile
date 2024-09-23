# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  # The most common configuration options are documented and commented below.
  # For a complete reference, please see the online documentation at
  # https://docs.vagrantup.com.

  config.vm.network :private_network, type: "dhcp"

    config.ssh.username = "root"
    config.vm.allow_fstab_modification = false
    config.vm.provider "docker" do |d|
      d.image = "kevbob/docker-vagrant:jammy-1.0.1-1"
      d.has_ssh = true
      d.remains_running = true
      d.name = "godemo"
    end
    config.ssh.extra_args = ["-t", "cd /vagrant; bash --login"]
    config.vm.network "forwarded_port", guest: 3000, host: 3000

    config.vm.provision "shell", privileged: false, inline: <<-SHELL
      #os
      apt-get update
      apt-get install -y \
          redis-tools \
          iputils-ping
      #golang
      cd /tmp
      curl -L -O "https://go.dev/dl/go1.23.1.linux-arm64.tar.gz"
      rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.1.linux-arm64.tar.gz


      cd ~/
      sed -i '/local\/go\/bin/d' ~/.bashrc
      echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
    SHELL

end
