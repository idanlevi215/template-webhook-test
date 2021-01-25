package webhooks

import (
	"fmt"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"context"
	//"k8s.io/api/rbac/v1alpha1"
	"net/http"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)



type NamepsaceAnnotator struct {
	Client  client.Client
	Decoder *admission.Decoder
	Log     logr.Logger
}

// +kubebuilder:webhook:path=/validate-v1-namespace,mutating=false,failurePolicy=fail,groups="core",resources=namespaces,verbs=delete,versions=v1,name=dana.nhs.io

func (a *NamepsaceAnnotator) Handle(ctx context.Context, req admission.Request) admission.Response {
	log := a.Log.WithValues("webhook", "Namespace Webhook")
	log.V(1).Info("webhook request received")
	namespace := &corev1.Namespace{}

	if err := a.Decoder.DecodeRaw(req.OldObject, namespace); err != nil {
		log.V(2).Error(err, "could not decode object")
		return admission.Errored(http.StatusBadRequest, err)
	}
	fmt.Print("\nNamesapce", namespace)


	log.V(1).Info("annotation matched")
	return admission.Allowed("OK")
}

func (a *NamepsaceAnnotator) InjectDecoder(d *admission.Decoder) error {
	a.Decoder = d
	return nil
}