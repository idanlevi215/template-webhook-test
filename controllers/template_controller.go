/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	//"sigs.k8s.io/controller-runtime/pkg/log"
	"fmt"
	"github.com/go-logr/logr"
	tem "github.com/openshift/api/template/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TemplateReconciler reconciles a Template object
type TemplateReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=template.openshift.io.idan.test,resources=templates,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=template.openshift.io.idan.test,resources=templates/status,verbs=get;update;patch

func (r *TemplateReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("template", req.NamespacedName)
	log := r.Log.WithValues("template", req.NamespacedName)
	log.Info("starting to reconcile")
	var (
		template tem.Template
		name string

	)

	if err := r.Get(ctx, req.NamespacedName, &template); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, client.IgnoreNotFound(err)
		}
		return ctrl.Result{}, err
	}

	name =template.ObjectMeta.Name
	if name == "createnamesapcewolt" {
		fmt.Print("THis is your Tempalte")
		return ctrl.Result{}, nil
	}


	return ctrl.Result{}, nil
}

func (r *TemplateReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&tem.Template{}).
		Complete(r)
}
