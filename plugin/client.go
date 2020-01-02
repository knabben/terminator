package main

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/tools/clientcmd"

	"k8s.io/client-go/dynamic"
)

var clientset dynamic.Interface

// SetConfiguration return configuration from kubeconfig
func SetConfiguration() (*rest.Config, error) {
	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	return clientcmd.BuildConfigFromFlags("", *kubeconfig)
}

// Connect return the new clientset
func Connect(config *rest.Config) (dynamic.Interface, error) {
	return dynamic.NewForConfig(config)
}
