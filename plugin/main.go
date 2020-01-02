/*
Copyright (c) 2019 the Octant contributors. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	"log"
	"path/filepath"
	"strings"

	"github.com/vmware-tanzu/octant/pkg/navigation"
	"github.com/vmware-tanzu/octant/pkg/plugin"
	"github.com/vmware-tanzu/octant/pkg/plugin/service"
	"github.com/vmware-tanzu/octant/pkg/view/component"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var pluginName = "bs"

func main() {
	log.SetPrefix("bs")

	capabilities := &plugin.Capabilities{
		ActionNames: []string{"crds/create-new"},
		IsModule:    true,
	}

	options := []service.PluginOption{
		service.WithNavigation(handleNavigation, initRoutes),
		service.WithActionHandler(handleActions),
	}

	p, err := service.Register(pluginName, "Backing services plugin", capabilities, options...)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(fmt.Sprintf("%s is starting", pluginName))
	p.Serve()
}

// handleActions - set actions handler
func handleActions(request *service.ActionRequest) error {
	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	c, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
		return err
	}

	res := schema.GroupVersionResource{
		Group: "backing.bluebird.io",
		Version: "v1",
		Resource: "services",
	}

	grvClient := c.Resource(res).Namespace("default")

	list, err := grvClient.List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	for _, r := range list.Items {
		log.Printf(fmt.Sprintf("Found one. %s", r))
	}

	svc := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "backing.bluebird.io/v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name": "name-1",
			},
			"spec": map[string]interface{}{
				"name": "redis",
			},
		},
	}

	if _, err := grvClient.Create(svc, metav1.CreateOptions{}); err != nil {
		log.Fatal(err)
		return err
	}

	log.Printf(fmt.Sprintf("%s", request.Payload))
	return nil
}

// initRoutes - initialize the routes for the tabs
func initRoutes(router *service.Router) {
	router.HandleFunc("*", func(request *service.Request) (component.ContentResponse, error) {
		contentResponse := component.NewContentResponse(component.TitleFromString("Backing Services"))
		contentResponse.Add(
			componentGenerator("Redis", "tab1", request.Path),
			componentGenerator("RabbitMQ", "tab2", request.Path),
		)
		return *contentResponse, nil
	})
}

// handleNavigation - build the lateral menu
func handleNavigation(request *service.NavigationRequest) (navigation.Navigation, error) {
	return navigation.Navigation{
		Title: "Terminator",
		Path:  request.GeneratePath(),
		Children: []navigation.Navigation{
			{
				Title:    "Backing services",
				Path:     request.GeneratePath("bs"),
				IconName: "folder",
			},
		},
		IconName: "cloud",
	}, nil
}

// componentGenerator - Create a generic component for launch the action
func componentGenerator(name, accessor, requestPath string) component.Component {
	cardBody := component.NewText(fmt.Sprintf("This will generate a new service of type: %s", name))
	card := component.NewCard(fmt.Sprintf("%s", name))
	card.SetBody(cardBody)

	action := component.Action{
		Name:  "Create",
		Title: "Create a new replicaset",
		Form: component.Form{
			Fields: []component.FormField{
				component.NewFormFieldNumber("Replicas", "replicas", "1"),
				component.NewFormFieldHidden("action", "crds/create-new"),
				component.NewFormFieldHidden("object", strings.ToLower(name)),
			},
		},
	}
	card.AddAction(action)
	card.SetAccessor(accessor)
	return card
}
