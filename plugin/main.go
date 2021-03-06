/*
Copyright (c) 2019 the Octant contributors. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/vmware-tanzu/octant/pkg/navigation"
	"github.com/vmware-tanzu/octant/pkg/plugin"
	"github.com/vmware-tanzu/octant/pkg/plugin/service"
	"github.com/vmware-tanzu/octant/pkg/view/component"

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

	config, err := SetConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	clientset, err = Connect(config)
	if err != nil {
		log.Fatal(err)
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
	port := request.Payload["port"]
	log.Printf(fmt.Sprintf("-- PORT %s --", port))

	// Create a new CRD
	gvr, resource := crdBackingService(port.(string))
	gvrClient := clientset.Resource(gvr).Namespace("default")
	if _, err := gvrClient.Create(resource, metav1.CreateOptions{}); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// initRoutes - initialize the routes for the tabs
func initRoutes(router *service.Router) {
	router.HandleFunc("*", func(request *service.Request) (component.ContentResponse, error) {
		contentResponse := component.NewContentResponse(component.TitleFromString("Backing Services"))
		contentResponse.Add(
			componentGenerator("Redis", "tab1", request.Path),
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
				component.NewFormFieldNumber("Port", "port", "6379"),
				component.NewFormFieldHidden("action", "crds/create-new"),
				component.NewFormFieldHidden("object", strings.ToLower(name)),
			},
		},
	}
	card.AddAction(action)
	card.SetAccessor(accessor)
	return card
}
