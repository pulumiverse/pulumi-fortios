// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios router bgp
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			sample1, err := fortios.LookupRouterBgp(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("output1", sample1.Neighbors)
//			return nil
//		})
//	}
//
// ```
func LookupRouterBgp(ctx *pulumi.Context, args *LookupRouterBgpArgs, opts ...pulumi.InvokeOption) (*LookupRouterBgpResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterBgpResult
	err := ctx.Invoke("fortios:index/getRouterBgp:getRouterBgp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouterBgp.
type LookupRouterBgpArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getRouterBgp.
type LookupRouterBgpResult struct {
	// Enable/disable IPv4 additional-path capability.
	AdditionalPath string `pulumi:"additionalPath"`
	// Enable/disable IPv6 additional-path capability.
	AdditionalPath6 string `pulumi:"additionalPath6"`
	// Number of additional paths to be selected for each IPv4 NLRI.
	AdditionalPathSelect int `pulumi:"additionalPathSelect"`
	// Number of additional paths to be selected for each IPv6 NLRI.
	AdditionalPathSelect6 int `pulumi:"additionalPathSelect6"`
	// Number of additional paths to be selected for each VPNv4 NLRI.
	AdditionalPathSelectVpnv4 int `pulumi:"additionalPathSelectVpnv4"`
	// Enable/disable VPNv4 additional-path capability.
	AdditionalPathVpnv4 string `pulumi:"additionalPathVpnv4"`
	// Administrative distance modifications. The structure of `adminDistance` block is documented below.
	AdminDistances []GetRouterBgpAdminDistance `pulumi:"adminDistances"`
	// BGP IPv6 aggregate address table. The structure of `aggregateAddress6` block is documented below.
	AggregateAddress6s []GetRouterBgpAggregateAddress6 `pulumi:"aggregateAddress6s"`
	// BGP aggregate address table. The structure of `aggregateAddress` block is documented below.
	AggregateAddresses []GetRouterBgpAggregateAddress `pulumi:"aggregateAddresses"`
	// Enable/disable always compare MED.
	AlwaysCompareMed string `pulumi:"alwaysCompareMed"`
	// Router AS number, valid from 1 to 4294967295, 0 to disable BGP.
	As int `pulumi:"as"`
	// Enable/disable ignore AS path.
	BestpathAsPathIgnore string `pulumi:"bestpathAsPathIgnore"`
	// Enable/disable compare federation AS path length.
	BestpathCmpConfedAspath string `pulumi:"bestpathCmpConfedAspath"`
	// Enable/disable compare router ID for identical EBGP paths.
	BestpathCmpRouterid string `pulumi:"bestpathCmpRouterid"`
	// Enable/disable compare MED among confederation paths.
	BestpathMedConfed string `pulumi:"bestpathMedConfed"`
	// Enable/disable treat missing MED as least preferred.
	BestpathMedMissingAsWorst string `pulumi:"bestpathMedMissingAsWorst"`
	// Enable/disable client-to-client route reflection.
	ClientToClientReflection string `pulumi:"clientToClientReflection"`
	// Route reflector cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Confederation identifier.
	ConfederationIdentifier int `pulumi:"confederationIdentifier"`
	// Confederation peers. The structure of `confederationPeers` block is documented below.
	ConfederationPeers []GetRouterBgpConfederationPeer `pulumi:"confederationPeers"`
	// Enable/disable route-flap dampening.
	Dampening string `pulumi:"dampening"`
	// Maximum minutes a route can be suppressed.
	DampeningMaxSuppressTime int `pulumi:"dampeningMaxSuppressTime"`
	// Reachability half-life time for penalty (min).
	DampeningReachabilityHalfLife int `pulumi:"dampeningReachabilityHalfLife"`
	// Threshold to reuse routes.
	DampeningReuse int `pulumi:"dampeningReuse"`
	// Criteria for dampening.
	DampeningRouteMap string `pulumi:"dampeningRouteMap"`
	// Threshold to suppress routes.
	DampeningSuppress int `pulumi:"dampeningSuppress"`
	// Unreachability half-life time for penalty (min).
	DampeningUnreachabilityHalfLife int `pulumi:"dampeningUnreachabilityHalfLife"`
	// Default local preference.
	DefaultLocalPreference int `pulumi:"defaultLocalPreference"`
	// Enable/disable enforce deterministic comparison of MED.
	DeterministicMed string `pulumi:"deterministicMed"`
	// Distance for routes external to the AS.
	DistanceExternal int `pulumi:"distanceExternal"`
	// Distance for routes internal to the AS.
	DistanceInternal int `pulumi:"distanceInternal"`
	// Distance for routes local to the AS.
	DistanceLocal int `pulumi:"distanceLocal"`
	// Enable/disable EBGP multi-path.
	EbgpMultipath string `pulumi:"ebgpMultipath"`
	// Enable/disable enforce first AS for EBGP routes.
	EnforceFirstAs string `pulumi:"enforceFirstAs"`
	// Enable/disable reset peer BGP session if link goes down.
	FastExternalFailover string `pulumi:"fastExternalFailover"`
	// Enable/disable to exit graceful restart on timer only.
	GracefulEndOnTimer string `pulumi:"gracefulEndOnTimer"`
	// Enable/disable BGP graceful restart capabilities.
	GracefulRestart string `pulumi:"gracefulRestart"`
	// Time needed for neighbors to restart (sec).
	GracefulRestartTime int `pulumi:"gracefulRestartTime"`
	// Time to hold stale paths of restarting neighbor (sec).
	GracefulStalepathTime int `pulumi:"gracefulStalepathTime"`
	// Route advertisement/selection delay after restart (sec).
	GracefulUpdateDelay int `pulumi:"gracefulUpdateDelay"`
	// Interval (sec) before peer considered dead.
	HoldtimeTimer int `pulumi:"holdtimeTimer"`
	// Enable/disable IBGP multi-path.
	IbgpMultipath string `pulumi:"ibgpMultipath"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Don't send unknown optional capability notification message
	IgnoreOptionalCapability string `pulumi:"ignoreOptionalCapability"`
	// Frequency to send keep alive requests.
	KeepaliveTimer int `pulumi:"keepaliveTimer"`
	// Enable logging of BGP neighbour's changes
	LogNeighbourChanges string `pulumi:"logNeighbourChanges"`
	// Enable/disable use of recursive distance to select multipath.
	MultipathRecursiveDistance string `pulumi:"multipathRecursiveDistance"`
	// Neighbor group name.
	NeighborGroups []GetRouterBgpNeighborGroup `pulumi:"neighborGroups"`
	// BGP IPv6 neighbor range table. The structure of `neighborRange6` block is documented below.
	NeighborRange6s []GetRouterBgpNeighborRange6 `pulumi:"neighborRange6s"`
	// BGP neighbor range table. The structure of `neighborRange` block is documented below.
	NeighborRanges []GetRouterBgpNeighborRange `pulumi:"neighborRanges"`
	// BGP neighbor table. The structure of `neighbor` block is documented below.
	Neighbors []GetRouterBgpNeighbor `pulumi:"neighbors"`
	// BGP IPv6 network table. The structure of `network6` block is documented below.
	Network6s []GetRouterBgpNetwork6 `pulumi:"network6s"`
	// Configure insurance of BGP network route existence in IGP.
	NetworkImportCheck string `pulumi:"networkImportCheck"`
	// BGP network table. The structure of `network` block is documented below.
	Networks []GetRouterBgpNetwork `pulumi:"networks"`
	// Enable/disable priority inheritance for recursive resolution.
	RecursiveInheritPriority string `pulumi:"recursiveInheritPriority"`
	// Enable/disable recursive resolution of next-hop using BGP route.
	RecursiveNextHop string `pulumi:"recursiveNextHop"`
	// BGP IPv6 redistribute table. The structure of `redistribute6` block is documented below.
	Redistribute6s []GetRouterBgpRedistribute6 `pulumi:"redistribute6s"`
	// BGP IPv4 redistribute table. The structure of `redistribute` block is documented below.
	Redistributes []GetRouterBgpRedistribute `pulumi:"redistributes"`
	// Router ID.
	RouterId string `pulumi:"routerId"`
	// Background scanner interval (sec), 0 to disable it.
	ScanTime int `pulumi:"scanTime"`
	// Enable/disable only advertise routes from iBGP if routes present in an IGP.
	Synchronization string `pulumi:"synchronization"`
	// Configure tag-match mode. Resolves BGP routes with other routes containing the same tag.
	TagResolveMode string  `pulumi:"tagResolveMode"`
	Vdomparam      *string `pulumi:"vdomparam"`
	// BGP IPv6 VRF leaking table. The structure of `vrf6` block is documented below.
	Vrf6s []GetRouterBgpVrf6 `pulumi:"vrf6s"`
	// BGP IPv6 VRF leaking table. The structure of `vrfLeak6` block is documented below.
	VrfLeak6s []GetRouterBgpVrfLeak6 `pulumi:"vrfLeak6s"`
	// BGP VRF leaking table. The structure of `vrfLeak` block is documented below.
	VrfLeaks []GetRouterBgpVrfLeak `pulumi:"vrfLeaks"`
	// Target VRF ID <0 - 31>.
	Vrves []GetRouterBgpVrf `pulumi:"vrves"`
}

func LookupRouterBgpOutput(ctx *pulumi.Context, args LookupRouterBgpOutputArgs, opts ...pulumi.InvokeOption) LookupRouterBgpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterBgpResult, error) {
			args := v.(LookupRouterBgpArgs)
			r, err := LookupRouterBgp(ctx, &args, opts...)
			var s LookupRouterBgpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouterBgpResultOutput)
}

// A collection of arguments for invoking getRouterBgp.
type LookupRouterBgpOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterBgpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterBgpArgs)(nil)).Elem()
}

// A collection of values returned by getRouterBgp.
type LookupRouterBgpResultOutput struct{ *pulumi.OutputState }

func (LookupRouterBgpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterBgpResult)(nil)).Elem()
}

func (o LookupRouterBgpResultOutput) ToLookupRouterBgpResultOutput() LookupRouterBgpResultOutput {
	return o
}

func (o LookupRouterBgpResultOutput) ToLookupRouterBgpResultOutputWithContext(ctx context.Context) LookupRouterBgpResultOutput {
	return o
}

// Enable/disable IPv4 additional-path capability.
func (o LookupRouterBgpResultOutput) AdditionalPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.AdditionalPath }).(pulumi.StringOutput)
}

// Enable/disable IPv6 additional-path capability.
func (o LookupRouterBgpResultOutput) AdditionalPath6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.AdditionalPath6 }).(pulumi.StringOutput)
}

// Number of additional paths to be selected for each IPv4 NLRI.
func (o LookupRouterBgpResultOutput) AdditionalPathSelect() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.AdditionalPathSelect }).(pulumi.IntOutput)
}

// Number of additional paths to be selected for each IPv6 NLRI.
func (o LookupRouterBgpResultOutput) AdditionalPathSelect6() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.AdditionalPathSelect6 }).(pulumi.IntOutput)
}

// Number of additional paths to be selected for each VPNv4 NLRI.
func (o LookupRouterBgpResultOutput) AdditionalPathSelectVpnv4() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.AdditionalPathSelectVpnv4 }).(pulumi.IntOutput)
}

// Enable/disable VPNv4 additional-path capability.
func (o LookupRouterBgpResultOutput) AdditionalPathVpnv4() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.AdditionalPathVpnv4 }).(pulumi.StringOutput)
}

// Administrative distance modifications. The structure of `adminDistance` block is documented below.
func (o LookupRouterBgpResultOutput) AdminDistances() GetRouterBgpAdminDistanceArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpAdminDistance { return v.AdminDistances }).(GetRouterBgpAdminDistanceArrayOutput)
}

// BGP IPv6 aggregate address table. The structure of `aggregateAddress6` block is documented below.
func (o LookupRouterBgpResultOutput) AggregateAddress6s() GetRouterBgpAggregateAddress6ArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpAggregateAddress6 { return v.AggregateAddress6s }).(GetRouterBgpAggregateAddress6ArrayOutput)
}

// BGP aggregate address table. The structure of `aggregateAddress` block is documented below.
func (o LookupRouterBgpResultOutput) AggregateAddresses() GetRouterBgpAggregateAddressArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpAggregateAddress { return v.AggregateAddresses }).(GetRouterBgpAggregateAddressArrayOutput)
}

// Enable/disable always compare MED.
func (o LookupRouterBgpResultOutput) AlwaysCompareMed() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.AlwaysCompareMed }).(pulumi.StringOutput)
}

// Router AS number, valid from 1 to 4294967295, 0 to disable BGP.
func (o LookupRouterBgpResultOutput) As() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.As }).(pulumi.IntOutput)
}

// Enable/disable ignore AS path.
func (o LookupRouterBgpResultOutput) BestpathAsPathIgnore() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.BestpathAsPathIgnore }).(pulumi.StringOutput)
}

// Enable/disable compare federation AS path length.
func (o LookupRouterBgpResultOutput) BestpathCmpConfedAspath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.BestpathCmpConfedAspath }).(pulumi.StringOutput)
}

// Enable/disable compare router ID for identical EBGP paths.
func (o LookupRouterBgpResultOutput) BestpathCmpRouterid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.BestpathCmpRouterid }).(pulumi.StringOutput)
}

// Enable/disable compare MED among confederation paths.
func (o LookupRouterBgpResultOutput) BestpathMedConfed() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.BestpathMedConfed }).(pulumi.StringOutput)
}

// Enable/disable treat missing MED as least preferred.
func (o LookupRouterBgpResultOutput) BestpathMedMissingAsWorst() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.BestpathMedMissingAsWorst }).(pulumi.StringOutput)
}

// Enable/disable client-to-client route reflection.
func (o LookupRouterBgpResultOutput) ClientToClientReflection() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.ClientToClientReflection }).(pulumi.StringOutput)
}

// Route reflector cluster ID.
func (o LookupRouterBgpResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// Confederation identifier.
func (o LookupRouterBgpResultOutput) ConfederationIdentifier() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.ConfederationIdentifier }).(pulumi.IntOutput)
}

// Confederation peers. The structure of `confederationPeers` block is documented below.
func (o LookupRouterBgpResultOutput) ConfederationPeers() GetRouterBgpConfederationPeerArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpConfederationPeer { return v.ConfederationPeers }).(GetRouterBgpConfederationPeerArrayOutput)
}

// Enable/disable route-flap dampening.
func (o LookupRouterBgpResultOutput) Dampening() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.Dampening }).(pulumi.StringOutput)
}

// Maximum minutes a route can be suppressed.
func (o LookupRouterBgpResultOutput) DampeningMaxSuppressTime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.DampeningMaxSuppressTime }).(pulumi.IntOutput)
}

// Reachability half-life time for penalty (min).
func (o LookupRouterBgpResultOutput) DampeningReachabilityHalfLife() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.DampeningReachabilityHalfLife }).(pulumi.IntOutput)
}

// Threshold to reuse routes.
func (o LookupRouterBgpResultOutput) DampeningReuse() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.DampeningReuse }).(pulumi.IntOutput)
}

// Criteria for dampening.
func (o LookupRouterBgpResultOutput) DampeningRouteMap() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.DampeningRouteMap }).(pulumi.StringOutput)
}

// Threshold to suppress routes.
func (o LookupRouterBgpResultOutput) DampeningSuppress() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.DampeningSuppress }).(pulumi.IntOutput)
}

// Unreachability half-life time for penalty (min).
func (o LookupRouterBgpResultOutput) DampeningUnreachabilityHalfLife() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.DampeningUnreachabilityHalfLife }).(pulumi.IntOutput)
}

// Default local preference.
func (o LookupRouterBgpResultOutput) DefaultLocalPreference() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.DefaultLocalPreference }).(pulumi.IntOutput)
}

// Enable/disable enforce deterministic comparison of MED.
func (o LookupRouterBgpResultOutput) DeterministicMed() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.DeterministicMed }).(pulumi.StringOutput)
}

// Distance for routes external to the AS.
func (o LookupRouterBgpResultOutput) DistanceExternal() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.DistanceExternal }).(pulumi.IntOutput)
}

// Distance for routes internal to the AS.
func (o LookupRouterBgpResultOutput) DistanceInternal() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.DistanceInternal }).(pulumi.IntOutput)
}

// Distance for routes local to the AS.
func (o LookupRouterBgpResultOutput) DistanceLocal() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.DistanceLocal }).(pulumi.IntOutput)
}

// Enable/disable EBGP multi-path.
func (o LookupRouterBgpResultOutput) EbgpMultipath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.EbgpMultipath }).(pulumi.StringOutput)
}

// Enable/disable enforce first AS for EBGP routes.
func (o LookupRouterBgpResultOutput) EnforceFirstAs() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.EnforceFirstAs }).(pulumi.StringOutput)
}

// Enable/disable reset peer BGP session if link goes down.
func (o LookupRouterBgpResultOutput) FastExternalFailover() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.FastExternalFailover }).(pulumi.StringOutput)
}

// Enable/disable to exit graceful restart on timer only.
func (o LookupRouterBgpResultOutput) GracefulEndOnTimer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.GracefulEndOnTimer }).(pulumi.StringOutput)
}

// Enable/disable BGP graceful restart capabilities.
func (o LookupRouterBgpResultOutput) GracefulRestart() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.GracefulRestart }).(pulumi.StringOutput)
}

// Time needed for neighbors to restart (sec).
func (o LookupRouterBgpResultOutput) GracefulRestartTime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.GracefulRestartTime }).(pulumi.IntOutput)
}

// Time to hold stale paths of restarting neighbor (sec).
func (o LookupRouterBgpResultOutput) GracefulStalepathTime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.GracefulStalepathTime }).(pulumi.IntOutput)
}

// Route advertisement/selection delay after restart (sec).
func (o LookupRouterBgpResultOutput) GracefulUpdateDelay() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.GracefulUpdateDelay }).(pulumi.IntOutput)
}

// Interval (sec) before peer considered dead.
func (o LookupRouterBgpResultOutput) HoldtimeTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.HoldtimeTimer }).(pulumi.IntOutput)
}

// Enable/disable IBGP multi-path.
func (o LookupRouterBgpResultOutput) IbgpMultipath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.IbgpMultipath }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterBgpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.Id }).(pulumi.StringOutput)
}

// Don't send unknown optional capability notification message
func (o LookupRouterBgpResultOutput) IgnoreOptionalCapability() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.IgnoreOptionalCapability }).(pulumi.StringOutput)
}

// Frequency to send keep alive requests.
func (o LookupRouterBgpResultOutput) KeepaliveTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.KeepaliveTimer }).(pulumi.IntOutput)
}

// Enable logging of BGP neighbour's changes
func (o LookupRouterBgpResultOutput) LogNeighbourChanges() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.LogNeighbourChanges }).(pulumi.StringOutput)
}

// Enable/disable use of recursive distance to select multipath.
func (o LookupRouterBgpResultOutput) MultipathRecursiveDistance() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.MultipathRecursiveDistance }).(pulumi.StringOutput)
}

// Neighbor group name.
func (o LookupRouterBgpResultOutput) NeighborGroups() GetRouterBgpNeighborGroupArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpNeighborGroup { return v.NeighborGroups }).(GetRouterBgpNeighborGroupArrayOutput)
}

// BGP IPv6 neighbor range table. The structure of `neighborRange6` block is documented below.
func (o LookupRouterBgpResultOutput) NeighborRange6s() GetRouterBgpNeighborRange6ArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpNeighborRange6 { return v.NeighborRange6s }).(GetRouterBgpNeighborRange6ArrayOutput)
}

// BGP neighbor range table. The structure of `neighborRange` block is documented below.
func (o LookupRouterBgpResultOutput) NeighborRanges() GetRouterBgpNeighborRangeArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpNeighborRange { return v.NeighborRanges }).(GetRouterBgpNeighborRangeArrayOutput)
}

// BGP neighbor table. The structure of `neighbor` block is documented below.
func (o LookupRouterBgpResultOutput) Neighbors() GetRouterBgpNeighborArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpNeighbor { return v.Neighbors }).(GetRouterBgpNeighborArrayOutput)
}

// BGP IPv6 network table. The structure of `network6` block is documented below.
func (o LookupRouterBgpResultOutput) Network6s() GetRouterBgpNetwork6ArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpNetwork6 { return v.Network6s }).(GetRouterBgpNetwork6ArrayOutput)
}

// Configure insurance of BGP network route existence in IGP.
func (o LookupRouterBgpResultOutput) NetworkImportCheck() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.NetworkImportCheck }).(pulumi.StringOutput)
}

// BGP network table. The structure of `network` block is documented below.
func (o LookupRouterBgpResultOutput) Networks() GetRouterBgpNetworkArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpNetwork { return v.Networks }).(GetRouterBgpNetworkArrayOutput)
}

// Enable/disable priority inheritance for recursive resolution.
func (o LookupRouterBgpResultOutput) RecursiveInheritPriority() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.RecursiveInheritPriority }).(pulumi.StringOutput)
}

// Enable/disable recursive resolution of next-hop using BGP route.
func (o LookupRouterBgpResultOutput) RecursiveNextHop() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.RecursiveNextHop }).(pulumi.StringOutput)
}

// BGP IPv6 redistribute table. The structure of `redistribute6` block is documented below.
func (o LookupRouterBgpResultOutput) Redistribute6s() GetRouterBgpRedistribute6ArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpRedistribute6 { return v.Redistribute6s }).(GetRouterBgpRedistribute6ArrayOutput)
}

// BGP IPv4 redistribute table. The structure of `redistribute` block is documented below.
func (o LookupRouterBgpResultOutput) Redistributes() GetRouterBgpRedistributeArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpRedistribute { return v.Redistributes }).(GetRouterBgpRedistributeArrayOutput)
}

// Router ID.
func (o LookupRouterBgpResultOutput) RouterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.RouterId }).(pulumi.StringOutput)
}

// Background scanner interval (sec), 0 to disable it.
func (o LookupRouterBgpResultOutput) ScanTime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.ScanTime }).(pulumi.IntOutput)
}

// Enable/disable only advertise routes from iBGP if routes present in an IGP.
func (o LookupRouterBgpResultOutput) Synchronization() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.Synchronization }).(pulumi.StringOutput)
}

// Configure tag-match mode. Resolves BGP routes with other routes containing the same tag.
func (o LookupRouterBgpResultOutput) TagResolveMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.TagResolveMode }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// BGP IPv6 VRF leaking table. The structure of `vrf6` block is documented below.
func (o LookupRouterBgpResultOutput) Vrf6s() GetRouterBgpVrf6ArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpVrf6 { return v.Vrf6s }).(GetRouterBgpVrf6ArrayOutput)
}

// BGP IPv6 VRF leaking table. The structure of `vrfLeak6` block is documented below.
func (o LookupRouterBgpResultOutput) VrfLeak6s() GetRouterBgpVrfLeak6ArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpVrfLeak6 { return v.VrfLeak6s }).(GetRouterBgpVrfLeak6ArrayOutput)
}

// BGP VRF leaking table. The structure of `vrfLeak` block is documented below.
func (o LookupRouterBgpResultOutput) VrfLeaks() GetRouterBgpVrfLeakArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpVrfLeak { return v.VrfLeaks }).(GetRouterBgpVrfLeakArrayOutput)
}

// Target VRF ID <0 - 31>.
func (o LookupRouterBgpResultOutput) Vrves() GetRouterBgpVrfArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpVrf { return v.Vrves }).(GetRouterBgpVrfArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterBgpResultOutput{})
}
