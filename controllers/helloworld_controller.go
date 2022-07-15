/*
Copyright 2022.

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
	demov1alpha1 "github.com/slintes/demo-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/util/patch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"time"
)

// HelloWorldReconciler reconciles a HelloWorld object
type HelloWorldReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=demo.example.com,resources=helloworlds,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=demo.example.com,resources=helloworlds/status,verbs=get;update;patch
// + k ubebuilder:rbac:groups=demo.example.com,resources=helloworlds/finalizers,verbs=update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the HelloWorld object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *HelloWorldReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// TODO(user): your logic here
	hw := &demov1alpha1.HelloWorld{}
	err := r.Get(ctx, client.ObjectKey{Namespace: req.Namespace, Name: req.Name}, hw)
	if err != nil {
		if errors.IsNotFound(err) {
			// CR deleted, do cleanup if needed
			logger.Info("HelloWorld CR was deleted", "name", req.Name)
			return ctrl.Result{}, nil
		} else {
			logger.Error(err, "Failed to get CR", "name", req.Name)
			// returning the error will trigger Reconcile again later on
			return ctrl.Result{}, err
		}
	}

	// CR was created or updated
	// Compare with cluster state and take appropriate action
	logger.Info("Reconciling HelloWorld CR", "name", hw.GetName(), "message", hw.Spec.Message)

	helper, err := patch.NewHelper(hw, r.Client)
	if err != nil {
		return ctrl.Result{}, err
	}

	defer func() {
		// Always attempt to Patch the Remediation object and status after each reconciliation.
		// Patch ObservedGeneration only if the reconciliation completed successfully
		patchOpts := []patch.Option{}
		patchOpts = append(patchOpts, patch.WithStatusObservedGeneration{})

		err := helper.Patch(ctx, hw, patchOpts...)
		if err != nil {
			logger.Error(err, "failed to Patch metal3Remediation")
		}
	}()

	// TEST finalizer
	// add and remove finalizer every few seconds...

	testFinalizer := "example.com/test"
	if !controllerutil.ContainsFinalizer(hw, testFinalizer) {
		controllerutil.AddFinalizer(hw, testFinalizer)
		hw.Annotations = map[string]string{
			"test1": "test1",
			"test2": "test2",
		}
		hw.Labels = map[string]string{
			"test1": "test1",
			"test2": "test2",
		}
	} else {
		controllerutil.RemoveFinalizer(hw, testFinalizer)
		hw.Annotations = map[string]string{
			"test1": "test1",
		}
		hw.Labels = map[string]string{
			"test1": "test1",
		}
	}

	time.Sleep(5 * time.Second)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *HelloWorldReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1alpha1.HelloWorld{}).
		Complete(r)
}
