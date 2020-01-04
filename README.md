Terminator
==========

Backing services operator, fast third party services creation for development and
general testing.

![Screenshot](https://raw.githubusercontent.com/knabben/terminator/master/screen/screenshot.png)


## Install

To deloy the operator in a cluster:

```
$ export IMG=knabben/operator

cd operator
operator$ make docker-build
operator$ make docker-push

# we have the operator set with the knabben/operator image
 
operator$ make install
operator$ make deploy
```

### Development

**NOTE:** This is a [WIP], for the older version go to https://github.com/knabben/terminator/tree/v1.5.0

### Operator

For development enter the *operator* folder and run:

```
$ cd operator

operator$ make install
operator$ make run
```

### Plugin

Compile the Octant plugin and install in the *~/.config/octant/plugins/*

```
make plugin
```

Run Octant after installing the plugin: 
  
```
vmware-tanzu/octant/web$ npm run start  # ui 
vmware-tanzu/octant$ go run cmd/octant/main.go --proxy-frontend http://localhost:4200 -v=7
```