package query

import (
	"errors"

	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/kube/apis/gloo.solo.io/v1"

	"github.com/solo-io/gloo/projects/gateway2/reports"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// ProcessBackendRef is meant to take the result of a call to `GetBackendForRef` as well as a reporter and the original ref.
// The return value is a pointer to a string which is the cluster_name of the upstream that the ref resolved to.
// This function will return nil if the ref is not valid.
// This function will also set the appropriate condition on the parent via the reporter.
func ProcessBackendRef(obj client.Object, err error, reporter reports.ParentRefReporter, backendRef gwv1.BackendObjectReference) *string {
	if err != nil {
		switch {
		case errors.Is(err, ErrUnknownKind):
			reporter.SetCondition(reports.HTTPRouteCondition{
				Type:   gwv1.RouteConditionResolvedRefs,
				Status: metav1.ConditionFalse,
				Reason: gwv1.RouteReasonInvalidKind,
			})
		case errors.Is(err, ErrMissingReferenceGrant):
			reporter.SetCondition(reports.HTTPRouteCondition{
				Type:    gwv1.RouteConditionResolvedRefs,
				Status:  metav1.ConditionFalse,
				Reason:  gwv1.RouteReasonRefNotPermitted,
				Message: err.Error(),
			})
		case apierrors.IsNotFound(err):
			reporter.SetCondition(reports.HTTPRouteCondition{
				Type:   gwv1.RouteConditionResolvedRefs,
				Status: metav1.ConditionFalse,
				Reason: gwv1.RouteReasonBackendNotFound,
			})
		default:
			// setting other errors to not found. not sure if there's a better option.
			reporter.SetCondition(reports.HTTPRouteCondition{
				Type:   gwv1.RouteConditionResolvedRefs,
				Status: metav1.ConditionFalse,
				Reason: gwv1.RouteReasonBackendNotFound,
			})
		}
	} else {
		var port uint32
		if backendRef.Port != nil {
			port = uint32(*backendRef.Port)
		}
		switch cli := obj.(type) {
		// TODO(ilackarms): consider converging all backend ref handling to a single package and remove the various switches. Or at least document all locations where we have multiple switches.
		case *gloov1.Upstream:
			name := cli.GetName()
			return &name
		case *corev1.Service:
			if port == 0 {
				reporter.SetCondition(reports.HTTPRouteCondition{
					Type:   gwv1.RouteConditionResolvedRefs,
					Status: metav1.ConditionFalse,
					Reason: gwv1.RouteReasonUnsupportedValue,
				})
			} else {
				name := cli.GetName()
				return &name
			}
		default:
			if backendRef.Kind != nil && *backendRef.Kind == "LLMProvider" {
				name := cli.GetName()
				return &name
			}
			reporter.SetCondition(reports.HTTPRouteCondition{
				Type:   gwv1.RouteConditionResolvedRefs,
				Status: metav1.ConditionFalse,
				Reason: gwv1.RouteReasonInvalidKind,
			})
		}
	}
	return nil
}
