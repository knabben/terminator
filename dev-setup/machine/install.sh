#!/bin/bash -x

apt-get update
apt-get -y upgrade

apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
add-apt-repository \
   "deb https://download.docker.com/linux/$(. /etc/os-release; echo "$ID") \
   $(lsb_release -cs) \
   stable"
apt-get update && apt-get install -y docker-ce
sudo sed -i 's/-H fd:\/\// -H fd:\/\/ -H tcp:\/\/192.168.99.252:2376/g' /lib/systemd/system/docker.service

curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF
apt-get update
apt-get install -y kubelet kubeadm kubectl

sed -i "s/cgroup-driver=systemd/cgroup-driver=cgroupfs/g" /etc/systemd/system/kubelet.service.d/10-kubeadm.conf

echo 'GRUB_CMDLINE_LINUX="cgroup_enable=memory"' >> /etc/default/grub
sed -i '10,11d' /etc/fstab

swapoff -a

systemctl daemon-reload
systemctl restart kubelet
systemctl restart docker

if [ ! -f go1.10.1.linux-amd64.tar.gz ]; then
	curl https://dl.google.com/go/go1.10.1.linux-amd64.tar.gz --output go1.10.1.linux-amd64.tar.gz
	tar -C /usr/local -xvf go1.10.1.linux-amd64.tar.gz
fi

if ! grep -Fxq GOPATH /etc/profile
then
	echo "export GOPATH=/home/vagrant/go" >> /etc/profile
	echo "export PATH=${PATH}:/usr/local/go/bin:/home/vagrant/bin" >> /etc/profile
fi

source /etc/profile

swapoff --all
go get github.com/kubernetes-incubator/cri-tools/cmd/crictl
go install github.com/kubernetes-incubator/cri-tools/cmd/crictl

kubeadm init --apiserver-advertise-address 192.168.99.252

mkdir -p /home/vagrant/.kube
cp /etc/kubernetes/admin.conf /home/vagrant/.kube/config
chown -R vagrant /home/vagrant/.kube

export kubever=$(kubectl version | base64 | tr -d '\n')

kubectl apply -f "https://cloud.weave.works/k8s/net?k8s-version=$kubever"
kubectl taint nodes --all node-role.kubernetes.io/master-
