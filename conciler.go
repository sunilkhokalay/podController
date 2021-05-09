package main

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type PodReconciler struct {
	client.Client
	scheme *runtime.Scheme
}

func (p *PodReconciler) Reconcile(context context.Context, req ctrl.Request) (ctrl.Result, error) {
	println(fmt.Sprintf("Reconcile: \n\tkind: Pod \n\tName: \"%s\"\n\tNameSpace: %s\n\n",
		req.Name, req.Namespace))
	return ctrl.Result{}, nil
}

func (p *PodReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.Pod{}).
		WithEventFilter(NewPodPredicate()).
		Complete(p)
}
