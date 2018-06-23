## Kube-Dev - setup

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

### Bringing SOME services up

To bring the backing services up, run:

```
$ helm init
$ ./machine/install-charts.sh
```

### Sniffing PostgreSQL

To sniff the traffic on PostgreSQL pod you can run:

```
$ sniffer
```
