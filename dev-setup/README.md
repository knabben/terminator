## Kube-Dev

## Setup

### Minikube

```
$ minikube/setup.sh 
```

### Vagrant

```
$ cd machine
machine$ vagrant up
```

## Work

### Bringing services up

To bring the backing services up, run:

```
$ ./default_charts.sh
```

### Listen ports Locally

Take a look on https://github.com/knabben/forwarder

### Sniffing PostgreSQL

To sniff the traffic on PostgreSQL pod you can run:

```
$ sniffer
```
