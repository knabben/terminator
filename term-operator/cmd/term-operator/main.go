package main

import (
	"context"
	"github.com/getsentry/raven-go"
	"os"
	"runtime"

	stub "github.com/knabben/terminator/term-operator/pkg/stub"
	sdk "github.com/operator-framework/operator-sdk/pkg/sdk"
	k8sutil "github.com/operator-framework/operator-sdk/pkg/util/k8sutil"
	sdkVersion "github.com/operator-framework/operator-sdk/version"

	"github.com/sirupsen/logrus"
)

func printVersion() {
	logrus.Infof("Go Version: %s", runtime.Version())
	logrus.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	logrus.Infof("operator-sdk Version: %v", sdkVersion.Version)
}

func installRaven() {
	raven.SetDSN(os.Getenv("SENTRY_DSN"))
}

func main() {
	printVersion()
	installRaven()

	resource := "app.terminator.dev/v1alpha1"
	kind := "Terminator"
	namespace, err := k8sutil.GetWatchNamespace()
	if err != nil {
		raven.CaptureError(err, nil)
		logrus.Fatalf("Failed to get watch namespace: %v", err)
	}

	resyncPeriod := 5
	logrus.Infof("Watching %s, %s, %s, %d", resource, kind, namespace, resyncPeriod)

	raven.CapturePanic(func() {
		sdk.Watch(resource, kind, namespace, resyncPeriod)
	}, nil)

	sdk.Handle(stub.NewHandler())
	sdk.Run(context.TODO())
}
