/*
Copyright (c) 2019 the Octant contributors. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	_ "k8s.io/apimachinery/pkg/runtime/schema"
	"log"
	"strings"

	_ "github.com/knabben/terminator/plugin/api/v1"
	"github.com/vmware-tanzu/octant/pkg/navigation"
	"github.com/vmware-tanzu/octant/pkg/plugin"
	"github.com/vmware-tanzu/octant/pkg/plugin/service"
	"github.com/vmware-tanzu/octant/pkg/view/component"
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

// HandleActions - TODO - create a handler for the action
func handleActions(request *service.ActionRequest) error {
	fmt.Println("request", request)
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
