package main

import (
	"sigs.k8s.io/controller-runtime/pkg/event"
)

// PodPredicate is default predicate for reconciles.
type PodPredicate struct{}

// NewPodPredicate returns a new.
func NewPodPredicate() *PodPredicate {
	return &PodPredicate{}
}

// Create returns true if the Create event should be processed.
func (c *PodPredicate) Create(e event.CreateEvent) bool {
	// Dynamic creates not reconciled. Status updates comes through report CRs
	// Reports creates are reconciled.
	return true
}

// Delete returns true if the Delete event should be processed.
func (c *PodPredicate) Delete(event.DeleteEvent) bool {
	return true
}

// Update returns true if the Update event should be processed.
func (c *PodPredicate) Update(event.UpdateEvent) bool {
	return true
}

// Generic returns true if the Generic event should be processed.
func (c *PodPredicate) Generic(event.GenericEvent) bool {
	return true
}
