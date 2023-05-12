// Code generated by skv2. DO NOT EDIT.

// Generated status setting functions

package types

import (
	v1 "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.solo.io/core/v1"
)

// SetPlacementStatus assigns the PlacementStatus for FederatedGatewayStatus
func (x *FederatedGatewayStatus) SetPlacementStatus(status *v1.PlacementStatus) {
	if x != nil {
		x.PlacementStatus = status
	}
}

// SetNamespacedPlacementStatuses assigns the PlacementStatuses for FederatedGatewayStatus
func (x *FederatedGatewayStatus) SetNamespacedPlacementStatuses(statuses map[string]*v1.PlacementStatus) {
	if x != nil {
		x.NamespacedPlacementStatuses = statuses
	}
}

// SetPlacementStatus assigns the PlacementStatus for FederatedMatchableHttpGatewayStatus
func (x *FederatedMatchableHttpGatewayStatus) SetPlacementStatus(status *v1.PlacementStatus) {
	if x != nil {
		x.PlacementStatus = status
	}
}

// SetNamespacedPlacementStatuses assigns the PlacementStatuses for FederatedMatchableHttpGatewayStatus
func (x *FederatedMatchableHttpGatewayStatus) SetNamespacedPlacementStatuses(statuses map[string]*v1.PlacementStatus) {
	if x != nil {
		x.NamespacedPlacementStatuses = statuses
	}
}

// SetPlacementStatus assigns the PlacementStatus for FederatedVirtualServiceStatus
func (x *FederatedVirtualServiceStatus) SetPlacementStatus(status *v1.PlacementStatus) {
	if x != nil {
		x.PlacementStatus = status
	}
}

// SetNamespacedPlacementStatuses assigns the PlacementStatuses for FederatedVirtualServiceStatus
func (x *FederatedVirtualServiceStatus) SetNamespacedPlacementStatuses(statuses map[string]*v1.PlacementStatus) {
	if x != nil {
		x.NamespacedPlacementStatuses = statuses
	}
}

// SetPlacementStatus assigns the PlacementStatus for FederatedRouteTableStatus
func (x *FederatedRouteTableStatus) SetPlacementStatus(status *v1.PlacementStatus) {
	if x != nil {
		x.PlacementStatus = status
	}
}

// SetNamespacedPlacementStatuses assigns the PlacementStatuses for FederatedRouteTableStatus
func (x *FederatedRouteTableStatus) SetNamespacedPlacementStatuses(statuses map[string]*v1.PlacementStatus) {
	if x != nil {
		x.NamespacedPlacementStatuses = statuses
	}
}