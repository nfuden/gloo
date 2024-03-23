package routeoptions

import (
	"context"
	"errors"
	"fmt"

	sologatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	solokubev1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1/kube/apis/gateway.solo.io/v1"
	"github.com/solo-io/gloo/projects/gateway2/query"
	"github.com/solo-io/gloo/projects/gateway2/reports"
	"github.com/solo-io/gloo/projects/gateway2/translator/plugins"
	"github.com/solo-io/gloo/projects/gateway2/translator/plugins/utils"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/contextutils"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

var _ plugins.RoutePlugin = &plugin{}
var _ plugins.RouteRulePlugin = &plugin{}

var gk = schema.GroupKind{
	Group: sologatewayv1.RouteOptionGVK.Group,
	Kind:  sologatewayv1.RouteOptionGVK.Kind,
}

type plugin struct {
	queries query.GatewayQueries
}

func NewPlugin(queries query.GatewayQueries) *plugin {
	return &plugin{
		queries,
	}
}

func (p *plugin) GetName() string {
	return "RouteOptionsPlugin"
}

func (p *plugin) ApplyRoutePlugin(
	ctx context.Context,
	routeCtx *plugins.RouteContext,
	routeOptions *v1.RouteOptions,
) error {
	routeOptObjs := getRouteOptions(ctx, routeCtx, p.queries)
	attachedOptions := findAttachedRouteOptions(routeCtx, routeOptObjs)
	if len(attachedOptions) == 0 {
		return nil
	}

	// sort attached options and apply only the earliest
	utils.SortByCreationTime(attachedOptions)
	earliestOption := attachedOptions[0]
	if earliestOption.Spec.GetOptions() != nil {
		*routeOptions = *earliestOption.Spec.GetOptions()
	}
	return nil
}

func (p *plugin) ApplyRouteRulePlugin(
	ctx context.Context,
	routeCtx *plugins.RouteContext,
	outputRoute *v1.Route,
) error {
	filter := utils.FindExtensionRefFilter(routeCtx, gk)
	if filter == nil {
		return nil
	}

	routeOption := &solokubev1.RouteOption{}
	err := utils.GetExtensionRefObj(context.Background(), routeCtx, p.queries, filter.ExtensionRef, routeOption)
	if err != nil {
		switch {
		case apierrors.IsNotFound(err):
			notFoundMsg := formatNotFoundMessage(routeCtx, filter)
			routeCtx.Reporter.SetCondition(reports.HTTPRouteCondition{
				Type:    gwv1.RouteConditionResolvedRefs,
				Status:  metav1.ConditionFalse,
				Reason:  gwv1.RouteReasonBackendNotFound,
				Message: notFoundMsg,
			})
			return errors.New(notFoundMsg)
		case errors.Is(err, utils.ErrTypesNotEqual):
		case errors.Is(err, utils.ErrNotSettable):
			devErr := fmt.Errorf("developer error while getting RouteOptions as ExtensionRef: %w", err)
			contextutils.LoggerFrom(ctx).DPanic(devErr)
			return devErr
		default:
			return fmt.Errorf("error while getting RouteOptions as ExtensionRef: %w", err)
		}
		return err
	}

	if routeOption.Spec.GetOptions() != nil {
		// clobber the existing RouteOptions; merge semantics may be desired later
		outputRoute.Options = routeOption.Spec.GetOptions()
	}
	return nil
}

func getRouteOptions(ctx context.Context, routeCtx *plugins.RouteContext, queries query.GatewayQueries) []*solokubev1.RouteOption {
	var routeList solokubev1.RouteOptionList
	err := queries.List(ctx, &routeList, client.InNamespace(routeCtx.Route.Namespace))
	if err != nil {
		contextutils.LoggerFrom(ctx).Errorf("error while Listing RouteOptions: %v", err)
		// TODO: handle correctly
		return nil
	}

	// as the RouteOptionList does not contain pointers, and RouteOption is a concrete proto message,
	// we need to turn it into a pointer slice to avoid copying proto message state around, copying locks, etc.
	// while we perform operations on the RouteOptionList
	ptrSlice := []*solokubev1.RouteOption{}
	items := routeList.Items
	for i := range items {
		ptrSlice = append(ptrSlice, &items[i])
	}
	return ptrSlice
}

func findAttachedRouteOptions(routeCtx *plugins.RouteContext, routeOptions []*solokubev1.RouteOption) []*solokubev1.RouteOption {
	attachedOptions := []*solokubev1.RouteOption{}
	for _, roObj := range routeOptions {
		targetRef := roObj.Spec.GetTargetRef()
		if !utils.IsPolicyAttachedToRoute(targetRef, routeCtx) {
			continue
		}
		attachedOptions = append(attachedOptions, roObj)
	}
	return attachedOptions
}

func formatNotFoundMessage(routeCtx *plugins.RouteContext, filter *gwv1.HTTPRouteFilter) string {
	return fmt.Sprintf(
		"extensionRef '%s' of type %s.%s in namespace '%s' not found",
		filter.ExtensionRef.Group,
		filter.ExtensionRef.Kind,
		filter.ExtensionRef.Name,
		routeCtx.Route.GetNamespace(),
	)
}
