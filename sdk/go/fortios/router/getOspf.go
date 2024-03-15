// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on fortios router ospf
func LookupOspf(ctx *pulumi.Context, args *LookupOspfArgs, opts ...pulumi.InvokeOption) (*LookupOspfResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOspfResult
	err := ctx.Invoke("fortios:router/getOspf:getOspf", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOspf.
type LookupOspfArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getOspf.
type LookupOspfResult struct {
	// Area border router type.
	AbrType string `pulumi:"abrType"`
	// Attach the network to area.
	Areas []GetOspfArea `pulumi:"areas"`
	// Reference bandwidth in terms of megabits per second.
	AutoCostRefBandwidth int `pulumi:"autoCostRefBandwidth"`
	// Bidirectional Forwarding Detection (BFD).
	Bfd string `pulumi:"bfd"`
	// Enable/disable database overflow.
	DatabaseOverflow string `pulumi:"databaseOverflow"`
	// Database overflow maximum LSAs.
	DatabaseOverflowMaxLsas int `pulumi:"databaseOverflowMaxLsas"`
	// Database overflow time to recover (sec).
	DatabaseOverflowTimeToRecover int `pulumi:"databaseOverflowTimeToRecover"`
	// Default information metric.
	DefaultInformationMetric int `pulumi:"defaultInformationMetric"`
	// Default information metric type.
	DefaultInformationMetricType string `pulumi:"defaultInformationMetricType"`
	// Enable/disable generation of default route.
	DefaultInformationOriginate string `pulumi:"defaultInformationOriginate"`
	// Default information route map.
	DefaultInformationRouteMap string `pulumi:"defaultInformationRouteMap"`
	// Default metric of redistribute routes.
	DefaultMetric int `pulumi:"defaultMetric"`
	// Distance of the route.
	Distance int `pulumi:"distance"`
	// Administrative external distance.
	DistanceExternal int `pulumi:"distanceExternal"`
	// Administrative inter-area distance.
	DistanceInterArea int `pulumi:"distanceInterArea"`
	// Administrative intra-area distance.
	DistanceIntraArea int `pulumi:"distanceIntraArea"`
	// Filter incoming routes.
	DistributeListIn string `pulumi:"distributeListIn"`
	// Distribute list configuration. The structure of `distributeList` block is documented below.
	DistributeLists []GetOspfDistributeList `pulumi:"distributeLists"`
	// Filter incoming external routes by route-map.
	DistributeRouteMapIn string `pulumi:"distributeRouteMapIn"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable logging of OSPF neighbour's changes
	LogNeighbourChanges string `pulumi:"logNeighbourChanges"`
	// OSPF neighbor configuration are used when OSPF runs on non-broadcast media The structure of `neighbor` block is documented below.
	Neighbors []GetOspfNeighbor `pulumi:"neighbors"`
	// OSPF network configuration. The structure of `network` block is documented below.
	Networks []GetOspfNetwork `pulumi:"networks"`
	// OSPF interface configuration. The structure of `ospfInterface` block is documented below.
	OspfInterfaces []GetOspfOspfInterface `pulumi:"ospfInterfaces"`
	// Passive interface configuration. The structure of `passiveInterface` block is documented below.
	PassiveInterfaces []GetOspfPassiveInterface `pulumi:"passiveInterfaces"`
	// Redistribute configuration. The structure of `redistribute` block is documented below.
	Redistributes []GetOspfRedistribute `pulumi:"redistributes"`
	// OSPF restart mode (graceful or LLS).
	RestartMode string `pulumi:"restartMode"`
	// Enable/disable continuing graceful restart upon topology change.
	RestartOnTopologyChange string `pulumi:"restartOnTopologyChange"`
	// Graceful restart period.
	RestartPeriod int `pulumi:"restartPeriod"`
	// Enable/disable RFC1583 compatibility.
	Rfc1583Compatible string `pulumi:"rfc1583Compatible"`
	// Router ID.
	RouterId string `pulumi:"routerId"`
	// SPF calculation frequency.
	SpfTimers string `pulumi:"spfTimers"`
	// IP address summary configuration. The structure of `summaryAddress` block is documented below.
	SummaryAddresses []GetOspfSummaryAddress `pulumi:"summaryAddresses"`
	Vdomparam        *string                 `pulumi:"vdomparam"`
}

func LookupOspfOutput(ctx *pulumi.Context, args LookupOspfOutputArgs, opts ...pulumi.InvokeOption) LookupOspfResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOspfResult, error) {
			args := v.(LookupOspfArgs)
			r, err := LookupOspf(ctx, &args, opts...)
			var s LookupOspfResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOspfResultOutput)
}

// A collection of arguments for invoking getOspf.
type LookupOspfOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupOspfOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOspfArgs)(nil)).Elem()
}

// A collection of values returned by getOspf.
type LookupOspfResultOutput struct{ *pulumi.OutputState }

func (LookupOspfResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOspfResult)(nil)).Elem()
}

func (o LookupOspfResultOutput) ToLookupOspfResultOutput() LookupOspfResultOutput {
	return o
}

func (o LookupOspfResultOutput) ToLookupOspfResultOutputWithContext(ctx context.Context) LookupOspfResultOutput {
	return o
}

// Area border router type.
func (o LookupOspfResultOutput) AbrType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfResult) string { return v.AbrType }).(pulumi.StringOutput)
}

// Attach the network to area.
func (o LookupOspfResultOutput) Areas() GetOspfAreaArrayOutput {
	return o.ApplyT(func(v LookupOspfResult) []GetOspfArea { return v.Areas }).(GetOspfAreaArrayOutput)
}

// Reference bandwidth in terms of megabits per second.
func (o LookupOspfResultOutput) AutoCostRefBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOspfResult) int { return v.AutoCostRefBandwidth }).(pulumi.IntOutput)
}

// Bidirectional Forwarding Detection (BFD).
func (o LookupOspfResultOutput) Bfd() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfResult) string { return v.Bfd }).(pulumi.StringOutput)
}

// Enable/disable database overflow.
func (o LookupOspfResultOutput) DatabaseOverflow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfResult) string { return v.DatabaseOverflow }).(pulumi.StringOutput)
}

// Database overflow maximum LSAs.
func (o LookupOspfResultOutput) DatabaseOverflowMaxLsas() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOspfResult) int { return v.DatabaseOverflowMaxLsas }).(pulumi.IntOutput)
}

// Database overflow time to recover (sec).
func (o LookupOspfResultOutput) DatabaseOverflowTimeToRecover() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOspfResult) int { return v.DatabaseOverflowTimeToRecover }).(pulumi.IntOutput)
}

// Default information metric.
func (o LookupOspfResultOutput) DefaultInformationMetric() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOspfResult) int { return v.DefaultInformationMetric }).(pulumi.IntOutput)
}

// Default information metric type.
func (o LookupOspfResultOutput) DefaultInformationMetricType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfResult) string { return v.DefaultInformationMetricType }).(pulumi.StringOutput)
}

// Enable/disable generation of default route.
func (o LookupOspfResultOutput) DefaultInformationOriginate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfResult) string { return v.DefaultInformationOriginate }).(pulumi.StringOutput)
}

// Default information route map.
func (o LookupOspfResultOutput) DefaultInformationRouteMap() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfResult) string { return v.DefaultInformationRouteMap }).(pulumi.StringOutput)
}

// Default metric of redistribute routes.
func (o LookupOspfResultOutput) DefaultMetric() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOspfResult) int { return v.DefaultMetric }).(pulumi.IntOutput)
}

// Distance of the route.
func (o LookupOspfResultOutput) Distance() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOspfResult) int { return v.Distance }).(pulumi.IntOutput)
}

// Administrative external distance.
func (o LookupOspfResultOutput) DistanceExternal() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOspfResult) int { return v.DistanceExternal }).(pulumi.IntOutput)
}

// Administrative inter-area distance.
func (o LookupOspfResultOutput) DistanceInterArea() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOspfResult) int { return v.DistanceInterArea }).(pulumi.IntOutput)
}

// Administrative intra-area distance.
func (o LookupOspfResultOutput) DistanceIntraArea() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOspfResult) int { return v.DistanceIntraArea }).(pulumi.IntOutput)
}

// Filter incoming routes.
func (o LookupOspfResultOutput) DistributeListIn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfResult) string { return v.DistributeListIn }).(pulumi.StringOutput)
}

// Distribute list configuration. The structure of `distributeList` block is documented below.
func (o LookupOspfResultOutput) DistributeLists() GetOspfDistributeListArrayOutput {
	return o.ApplyT(func(v LookupOspfResult) []GetOspfDistributeList { return v.DistributeLists }).(GetOspfDistributeListArrayOutput)
}

// Filter incoming external routes by route-map.
func (o LookupOspfResultOutput) DistributeRouteMapIn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfResult) string { return v.DistributeRouteMapIn }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupOspfResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable logging of OSPF neighbour's changes
func (o LookupOspfResultOutput) LogNeighbourChanges() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfResult) string { return v.LogNeighbourChanges }).(pulumi.StringOutput)
}

// OSPF neighbor configuration are used when OSPF runs on non-broadcast media The structure of `neighbor` block is documented below.
func (o LookupOspfResultOutput) Neighbors() GetOspfNeighborArrayOutput {
	return o.ApplyT(func(v LookupOspfResult) []GetOspfNeighbor { return v.Neighbors }).(GetOspfNeighborArrayOutput)
}

// OSPF network configuration. The structure of `network` block is documented below.
func (o LookupOspfResultOutput) Networks() GetOspfNetworkArrayOutput {
	return o.ApplyT(func(v LookupOspfResult) []GetOspfNetwork { return v.Networks }).(GetOspfNetworkArrayOutput)
}

// OSPF interface configuration. The structure of `ospfInterface` block is documented below.
func (o LookupOspfResultOutput) OspfInterfaces() GetOspfOspfInterfaceArrayOutput {
	return o.ApplyT(func(v LookupOspfResult) []GetOspfOspfInterface { return v.OspfInterfaces }).(GetOspfOspfInterfaceArrayOutput)
}

// Passive interface configuration. The structure of `passiveInterface` block is documented below.
func (o LookupOspfResultOutput) PassiveInterfaces() GetOspfPassiveInterfaceArrayOutput {
	return o.ApplyT(func(v LookupOspfResult) []GetOspfPassiveInterface { return v.PassiveInterfaces }).(GetOspfPassiveInterfaceArrayOutput)
}

// Redistribute configuration. The structure of `redistribute` block is documented below.
func (o LookupOspfResultOutput) Redistributes() GetOspfRedistributeArrayOutput {
	return o.ApplyT(func(v LookupOspfResult) []GetOspfRedistribute { return v.Redistributes }).(GetOspfRedistributeArrayOutput)
}

// OSPF restart mode (graceful or LLS).
func (o LookupOspfResultOutput) RestartMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfResult) string { return v.RestartMode }).(pulumi.StringOutput)
}

// Enable/disable continuing graceful restart upon topology change.
func (o LookupOspfResultOutput) RestartOnTopologyChange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfResult) string { return v.RestartOnTopologyChange }).(pulumi.StringOutput)
}

// Graceful restart period.
func (o LookupOspfResultOutput) RestartPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOspfResult) int { return v.RestartPeriod }).(pulumi.IntOutput)
}

// Enable/disable RFC1583 compatibility.
func (o LookupOspfResultOutput) Rfc1583Compatible() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfResult) string { return v.Rfc1583Compatible }).(pulumi.StringOutput)
}

// Router ID.
func (o LookupOspfResultOutput) RouterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfResult) string { return v.RouterId }).(pulumi.StringOutput)
}

// SPF calculation frequency.
func (o LookupOspfResultOutput) SpfTimers() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfResult) string { return v.SpfTimers }).(pulumi.StringOutput)
}

// IP address summary configuration. The structure of `summaryAddress` block is documented below.
func (o LookupOspfResultOutput) SummaryAddresses() GetOspfSummaryAddressArrayOutput {
	return o.ApplyT(func(v LookupOspfResult) []GetOspfSummaryAddress { return v.SummaryAddresses }).(GetOspfSummaryAddressArrayOutput)
}

func (o LookupOspfResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOspfResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOspfResultOutput{})
}
