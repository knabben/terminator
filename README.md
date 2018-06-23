# TERMINATOR

Kubernetes 12-factor backing services management, so far it supports:

* Memcache
* Redis

The project has 3 components, an operator, a backend and a frontend client. Telemetry is offered via Websockets for realtime monitoring.

## Starting a cluster

Check out dev-setup/README.md

## Operator

### Running local

To run the operator locally, make sure you have a default kubernetes configuration on your machine and run:

```
$make ops-run
```

### Deploy 

Deploy it to a real cluster:

```
$ make ops-deploy


# Check if ops pod is there
$ kubectl get pods
term-operator-8c8bb94f9-z9h7s         1/1       Running   0          4m


# Check the CRD created for the operator
$ kubectl get crd
terminators.app.terminator.dev   4m
```

### First flight

There's an example of CR on term-operator/deploy/cr.yaml:

```
apiVersion: "app.terminator.dev/v1alpha1"
kind: "Terminator"
metadata:
  name: "deploy"
  spec:
    memcache: true
    redis: true
```

Basically you just need to tell the spec, which services you are going to need,
the setting is a boolean flag.

To start it use:
```
$ kubectl create -f term-operator/deploy/cr.yaml

$ kubectl get pods -l hasta=la-vista
NAME                               READY     STATUS    RESTARTS   AGE
deploy-memcache-744bcb88-tzq6q     1/1       Running   0          10s
deploy-redis-84c898d79c-svwhn      1/1       Running   0          10s

$ kubectl get svc -l hasta=la-vista
NAME              TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)     AGE
deploy-memcache   ClusterIP   10.103.101.240   <none>        11211/TCP   10s
deploy-redis      ClusterIP   10.100.121.215   <none>        6379/TCP    10s
```

### Using externally

Check out https://github.com/knabben/forwarder

### Cleaning up

```
$ make ops-clean
```

## Web Server

### Running local

The webserver is made up of 2 parts, the Django Channels backend that holds all 
websockets request and a react/sagas listening that parses and show up the information,
to run the stack locally run:

```
make web-run
```
