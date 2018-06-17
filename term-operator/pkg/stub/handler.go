package stub

import (
	"context"

	"github.com/knabben/terminator/term-operator/pkg/apis/app/v1alpha1"
	"github.com/knabben/terminator/term-operator/pkg/terminator"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.Terminator:
		return terminator.Reconcile(o, event)
	}
	return nil
}
