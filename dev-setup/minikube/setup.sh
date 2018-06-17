#!/bin/bash
#
# Initial minikube setup
#

mkdir -p ~/.minikube/config
cat > ~/.minikube/config/config.json <<EOF
 {
    "WantReportError": true,
    "cpus": 4,
    "dashboard": false,
    "disk-size": "30G",
    "vm-driver": "xhyve"
 }
EOF
minikube start

minikube addons enable registry-creds
minikube addons disable dashboard
minikube addons configure registry-creds

echo 'eval $(minikube docker-env)' >> ~/.zshrc
echo 'alias sniffer="kubectl exec -i -t $(kubectl get pods -o json -l app=postgis-postgresql | jq ".items[0].metadata.name" -r) ./vc-pgsql-sniffer"' >> ~/.zshrc

helm init
