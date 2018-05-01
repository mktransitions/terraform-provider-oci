#!/bin/bash
## Node Tuning Script
## by Zachary Smith (Zachary.Smith@oracle.com)
## Last Update - March 2018


## Install and start NTP
sudo yum install ntp.x86_64 java-1.8.0-openjdk.x86_64 -y
sudo cat /etc/ntp.conf | grep -iv rhel >> /tmp/ntp.conf
echo "server 169.254.169.254 iburst" >> /tmp/ntp.conf
sudo cp /tmp/ntp.conf /etc/ntp.conf
sudo rm -f /tmp/ntp.conf
systemctl stop ntpd
sudo ntpdate 169.254.169.254
systemctl start ntpd
systemctl enable ntpd
systemctl stop chronyd
systemctl disable chronyd

## Disable Transparent Huge Pages
echo never | tee -a /sys/kernel/mm/transparent_hugepage/enabled
echo "echo never | tee -a /sys/kernel/mm/transparent_hugepage/enabled" | tee -a /etc/rc.local

## Set vm.swappiness to 1
echo vm.swappiness=1 | tee -a /etc/sysctl.conf
echo 1 | tee /proc/sys/vm/swappiness

## Tune system network performance
echo net.ipv4.tcp_timestamps=0 >> /etc/sysctl.conf
echo net.ipv4.tcp_sack=1 >> /etc/sysctl.conf
echo net.core.rmem_max=4194304 >> /etc/sysctl.conf
echo net.core.wmem_max=4194304 >> /etc/sysctl.conf
echo net.core.rmem_default=4194304 >> /etc/sysctl.conf
echo net.core.wmem_default=4194304 >> /etc/sysctl.conf
echo net.core.optmem_max=4194304 >> /etc/sysctl.conf
echo net.ipv4.tcp_rmem="4096 87380 4194304" >> /etc/sysctl.conf
echo net.ipv4.tcp_wmem="4096 65536 4194304" >> /etc/sysctl.conf
echo net.ipv4.tcp_low_latency=1 >> /etc/sysctl.conf

## Tune File System options
sed -i "s/defaults        1 1/defaults,noatime        0 0/" /etc/fstab

## Enable root login via SSH key
sudo cp /root/.ssh/authorized_keys /root/.ssh/authorized_keys.bak
sudo cp /home/opc/.ssh/authorized_keys /root/.ssh/authorized_keys

