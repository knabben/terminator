# TERMINATOR

Kubernetes 12-factor backing services management, so far it supports:

* Memcache
* Redis
* RabbitMQ

The project has 3 components, an operator, a backend and a frontend client. Telemetry is offered via Websockets for realtime monitoring.


![Screenshot](https://raw.githubusercontent.com/knabben/blog/master/static/images/terminator-screen.png)

## Fast Flight

If you are eager to run the system try this:

```
$ kind create cluster
$ export KUBECONFIG="$(kind get kubeconfig-path)"
$ cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ServiceAccount
metadata:
  name: tiller
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: tiller
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
  - kind: ServiceAccount
    name: tiller
    namespace: kube-system
EOF
$ helm repo add knabben https://knabben.github.io/charts/
$ helm install --name terminator knabben/terminator
``` 

## Starting a Kubernetes cluster

Check out dev-setup/README.md

## Running local

The webserver is made up of 2 parts, the Django Channels backend that holds all 
websockets request and a React/Sagas listening that parses and show up the information,
to run the stack locally run:

### Frontend server

To run the frontend you need Node v10.x and yarn installed.
```
make local-frontend-run
```


### Backend server
To run the backend server (Daphne) start with a pipenv install inside backend folder:

```
make local-backend-run
```

### Operator server

To run the operator locally, make sure you have a default kubernetes configuration on your machine and run. This will bring up all the Kubernetes assets (including the inicial CustomResource config) for default testing

```
make local-ops-run

role "term-operator" unchanged
rolebinding "default-account-term-operator" unchanged
customresourcedefinition "terminators.app.terminator.dev" configured
kubectl apply -f term-operator/deploy/cr.yaml
```

## Production

### Building the image

First is necessary to build the docker image for all projects, for this run:

```
$ make ops-build
$ make ops-push

$ make backend-build
$ make backend-push

$ make frontend-build
$ make frontend-push
```

### Deploy 

Deploy the frontend to the cluster:

```
$ make web-deploy

$ make ops-deploy
$ make ops-deploy-op

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
    rabbitmq: false
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

### Accessing the dashboard

Access the it on port: http://localhost:8092
