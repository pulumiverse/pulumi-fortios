// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on fortios router bgp
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/router"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			sample1, err := router.LookupBgp(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("output1", sample1.Neighbors)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupBgp(ctx *pulumi.Context, args *LookupBgpArgs, opts ...pulumi.InvokeOption) (*LookupBgpResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBgpResult
	err := ctx.Invoke("fortios:router/getBgp:getBgp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBgp.
type LookupBgpArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getBgp.
type LookupBgpResult struct {
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
	// Number of additional paths to be selected for each VPNv6 NLRI.
	AdditionalPathSelectVpnv6 int `pulumi:"additionalPathSelectVpnv6"`
	// Enable/disable VPNv4 additional-path capability.
	AdditionalPathVpnv4 string `pulumi:"additionalPathVpnv4"`
	// Enable/disable VPNv6 additional-path capability.
	AdditionalPathVpnv6 string `pulumi:"additionalPathVpnv6"`
	// Administrative distance modifications. The structure of `adminDistance` block is documented below.
	AdminDistances []GetBgpAdminDistance `pulumi:"adminDistances"`
	// BGP IPv6 aggregate address table. The structure of `aggregateAddress6` block is documented below.
	AggregateAddress6s []GetBgpAggregateAddress6 `pulumi:"aggregateAddress6s"`
	// BGP aggregate address table. The structure of `aggregateAddress` block is documented below.
	AggregateAddresses []GetBgpAggregateAddress `pulumi:"aggregateAddresses"`
	// Enable/disable always compare MED.
	AlwaysCompareMed string `pulumi:"alwaysCompareMed"`
	// Router AS number, valid from 1 to 4294967295, 0 to disable BGP.
	As int `pulumi:"as"`
	// Router AS number, asplain/asdot/asdot+ format, 0 to disable BGP.
	AsString string `pulumi:"asString"`
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
	ConfederationPeers []GetBgpConfederationPeer `pulumi:"confederationPeers"`
	// Enable/disable cross address family conditional advertisement.
	CrossFamilyConditionalAdv string `pulumi:"crossFamilyConditionalAdv"`
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
	NeighborGroups []GetBgpNeighborGroup `pulumi:"neighborGroups"`
	// BGP IPv6 neighbor range table. The structure of `neighborRange6` block is documented below.
	NeighborRange6s []GetBgpNeighborRange6 `pulumi:"neighborRange6s"`
	// BGP neighbor range table. The structure of `neighborRange` block is documented below.
	NeighborRanges []GetBgpNeighborRange `pulumi:"neighborRanges"`
	// BGP neighbor table. The structure of `neighbor` block is documented below.
	Neighbors []GetBgpNeighbor `pulumi:"neighbors"`
	// BGP IPv6 network table. The structure of `network6` block is documented below.
	Network6s []GetBgpNetwork6 `pulumi:"network6s"`
	// Configure insurance of BGP network route existence in IGP.
	NetworkImportCheck string `pulumi:"networkImportCheck"`
	// BGP network table. The structure of `network` block is documented below.
	Networks []GetBgpNetwork `pulumi:"networks"`
	// Enable/disable priority inheritance for recursive resolution.
	RecursiveInheritPriority string `pulumi:"recursiveInheritPriority"`
	// Enable/disable recursive resolution of next-hop using BGP route.
	RecursiveNextHop string `pulumi:"recursiveNextHop"`
	// BGP IPv6 redistribute table. The structure of `redistribute6` block is documented below.
	Redistribute6s []GetBgpRedistribute6 `pulumi:"redistribute6s"`
	// BGP IPv4 redistribute table. The structure of `redistribute` block is documented below.
	Redistributes []GetBgpRedistribute `pulumi:"redistributes"`
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
	Vrf6s []GetBgpVrf6 `pulumi:"vrf6s"`
	// BGP IPv6 VRF leaking table. The structure of `vrfLeak6` block is documented below.
	VrfLeak6s []GetBgpVrfLeak6 `pulumi:"vrfLeak6s"`
	// BGP VRF leaking table. The structure of `vrfLeak` block is documented below.
	VrfLeaks []GetBgpVrfLeak `pulumi:"vrfLeaks"`
	// Target VRF ID <0 - 31>.
	Vrves []GetBgpVrf `pulumi:"vrves"`
}

func LookupBgpOutput(ctx *pulumi.Context, args LookupBgpOutputArgs, opts ...pulumi.InvokeOption) LookupBgpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBgpResult, error) {
			args := v.(LookupBgpArgs)
			r, err := LookupBgp(ctx, &args, opts...)
			var s LookupBgpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBgpResultOutput)
}

// A collection of arguments for invoking getBgp.
type LookupBgpOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupBgpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBgpArgs)(nil)).Elem()
}

// A collection of values returned by getBgp.
type LookupBgpResultOutput struct{ *pulumi.OutputState }

func (LookupBgpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBgpResult)(nil)).Elem()
}

func (o LookupBgpResultOutput) ToLookupBgpResultOutput() LookupBgpResultOutput {
	return o
}

func (o LookupBgpResultOutput) ToLookupBgpResultOutputWithContext(ctx context.Context) LookupBgpResultOutput {
	return o
}

// Enable/disable IPv4 additional-path capability.
func (o LookupBgpResultOutput) AdditionalPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.AdditionalPath }).(pulumi.StringOutput)
}

// Enable/disable IPv6 additional-path capability.
func (o LookupBgpResultOutput) AdditionalPath6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.AdditionalPath6 }).(pulumi.StringOutput)
}

// Number of additional paths to be selected for each IPv4 NLRI.
func (o LookupBgpResultOutput) AdditionalPathSelect() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.AdditionalPathSelect }).(pulumi.IntOutput)
}

// Number of additional paths to be selected for each IPv6 NLRI.
func (o LookupBgpResultOutput) AdditionalPathSelect6() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.AdditionalPathSelect6 }).(pulumi.IntOutput)
}

// Number of additional paths to be selected for each VPNv4 NLRI.
func (o LookupBgpResultOutput) AdditionalPathSelectVpnv4() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.AdditionalPathSelectVpnv4 }).(pulumi.IntOutput)
}

// Number of additional paths to be selected for each VPNv6 NLRI.
func (o LookupBgpResultOutput) AdditionalPathSelectVpnv6() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.AdditionalPathSelectVpnv6 }).(pulumi.IntOutput)
}

// Enable/disable VPNv4 additional-path capability.
func (o LookupBgpResultOutput) AdditionalPathVpnv4() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.AdditionalPathVpnv4 }).(pulumi.StringOutput)
}

// Enable/disable VPNv6 additional-path capability.
func (o LookupBgpResultOutput) AdditionalPathVpnv6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.AdditionalPathVpnv6 }).(pulumi.StringOutput)
}

// Administrative distance modifications. The structure of `adminDistance` block is documented below.
func (o LookupBgpResultOutput) AdminDistances() GetBgpAdminDistanceArrayOutput {
	return o.ApplyT(func(v LookupBgpResult) []GetBgpAdminDistance { return v.AdminDistances }).(GetBgpAdminDistanceArrayOutput)
}

// BGP IPv6 aggregate address table. The structure of `aggregateAddress6` block is documented below.
func (o LookupBgpResultOutput) AggregateAddress6s() GetBgpAggregateAddress6ArrayOutput {
	return o.ApplyT(func(v LookupBgpResult) []GetBgpAggregateAddress6 { return v.AggregateAddress6s }).(GetBgpAggregateAddress6ArrayOutput)
}

// BGP aggregate address table. The structure of `aggregateAddress` block is documented below.
func (o LookupBgpResultOutput) AggregateAddresses() GetBgpAggregateAddressArrayOutput {
	return o.ApplyT(func(v LookupBgpResult) []GetBgpAggregateAddress { return v.AggregateAddresses }).(GetBgpAggregateAddressArrayOutput)
}

// Enable/disable always compare MED.
func (o LookupBgpResultOutput) AlwaysCompareMed() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.AlwaysCompareMed }).(pulumi.StringOutput)
}

// Router AS number, valid from 1 to 4294967295, 0 to disable BGP.
func (o LookupBgpResultOutput) As() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.As }).(pulumi.IntOutput)
}

// Router AS number, asplain/asdot/asdot+ format, 0 to disable BGP.
func (o LookupBgpResultOutput) AsString() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.AsString }).(pulumi.StringOutput)
}

// Enable/disable ignore AS path.
func (o LookupBgpResultOutput) BestpathAsPathIgnore() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.BestpathAsPathIgnore }).(pulumi.StringOutput)
}

// Enable/disable compare federation AS path length.
func (o LookupBgpResultOutput) BestpathCmpConfedAspath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.BestpathCmpConfedAspath }).(pulumi.StringOutput)
}

// Enable/disable compare router ID for identical EBGP paths.
func (o LookupBgpResultOutput) BestpathCmpRouterid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.BestpathCmpRouterid }).(pulumi.StringOutput)
}

// Enable/disable compare MED among confederation paths.
func (o LookupBgpResultOutput) BestpathMedConfed() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.BestpathMedConfed }).(pulumi.StringOutput)
}

// Enable/disable treat missing MED as least preferred.
func (o LookupBgpResultOutput) BestpathMedMissingAsWorst() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.BestpathMedMissingAsWorst }).(pulumi.StringOutput)
}

// Enable/disable client-to-client route reflection.
func (o LookupBgpResultOutput) ClientToClientReflection() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.ClientToClientReflection }).(pulumi.StringOutput)
}

// Route reflector cluster ID.
func (o LookupBgpResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// Confederation identifier.
func (o LookupBgpResultOutput) ConfederationIdentifier() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.ConfederationIdentifier }).(pulumi.IntOutput)
}

// Confederation peers. The structure of `confederationPeers` block is documented below.
func (o LookupBgpResultOutput) ConfederationPeers() GetBgpConfederationPeerArrayOutput {
	return o.ApplyT(func(v LookupBgpResult) []GetBgpConfederationPeer { return v.ConfederationPeers }).(GetBgpConfederationPeerArrayOutput)
}

// Enable/disable cross address family conditional advertisement.
func (o LookupBgpResultOutput) CrossFamilyConditionalAdv() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.CrossFamilyConditionalAdv }).(pulumi.StringOutput)
}

// Enable/disable route-flap dampening.
func (o LookupBgpResultOutput) Dampening() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.Dampening }).(pulumi.StringOutput)
}

// Maximum minutes a route can be suppressed.
func (o LookupBgpResultOutput) DampeningMaxSuppressTime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.DampeningMaxSuppressTime }).(pulumi.IntOutput)
}

// Reachability half-life time for penalty (min).
func (o LookupBgpResultOutput) DampeningReachabilityHalfLife() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.DampeningReachabilityHalfLife }).(pulumi.IntOutput)
}

// Threshold to reuse routes.
func (o LookupBgpResultOutput) DampeningReuse() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.DampeningReuse }).(pulumi.IntOutput)
}

// Criteria for dampening.
func (o LookupBgpResultOutput) DampeningRouteMap() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.DampeningRouteMap }).(pulumi.StringOutput)
}

// Threshold to suppress routes.
func (o LookupBgpResultOutput) DampeningSuppress() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.DampeningSuppress }).(pulumi.IntOutput)
}

// Unreachability half-life time for penalty (min).
func (o LookupBgpResultOutput) DampeningUnreachabilityHalfLife() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.DampeningUnreachabilityHalfLife }).(pulumi.IntOutput)
}

// Default local preference.
func (o LookupBgpResultOutput) DefaultLocalPreference() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.DefaultLocalPreference }).(pulumi.IntOutput)
}

// Enable/disable enforce deterministic comparison of MED.
func (o LookupBgpResultOutput) DeterministicMed() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.DeterministicMed }).(pulumi.StringOutput)
}

// Distance for routes external to the AS.
func (o LookupBgpResultOutput) DistanceExternal() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.DistanceExternal }).(pulumi.IntOutput)
}

// Distance for routes internal to the AS.
func (o LookupBgpResultOutput) DistanceInternal() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.DistanceInternal }).(pulumi.IntOutput)
}

// Distance for routes local to the AS.
func (o LookupBgpResultOutput) DistanceLocal() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.DistanceLocal }).(pulumi.IntOutput)
}

// Enable/disable EBGP multi-path.
func (o LookupBgpResultOutput) EbgpMultipath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.EbgpMultipath }).(pulumi.StringOutput)
}

// Enable/disable enforce first AS for EBGP routes.
func (o LookupBgpResultOutput) EnforceFirstAs() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.EnforceFirstAs }).(pulumi.StringOutput)
}

// Enable/disable reset peer BGP session if link goes down.
func (o LookupBgpResultOutput) FastExternalFailover() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.FastExternalFailover }).(pulumi.StringOutput)
}

// Enable/disable to exit graceful restart on timer only.
func (o LookupBgpResultOutput) GracefulEndOnTimer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.GracefulEndOnTimer }).(pulumi.StringOutput)
}

// Enable/disable BGP graceful restart capabilities.
func (o LookupBgpResultOutput) GracefulRestart() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.GracefulRestart }).(pulumi.StringOutput)
}

// Time needed for neighbors to restart (sec).
func (o LookupBgpResultOutput) GracefulRestartTime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.GracefulRestartTime }).(pulumi.IntOutput)
}

// Time to hold stale paths of restarting neighbor (sec).
func (o LookupBgpResultOutput) GracefulStalepathTime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.GracefulStalepathTime }).(pulumi.IntOutput)
}

// Route advertisement/selection delay after restart (sec).
func (o LookupBgpResultOutput) GracefulUpdateDelay() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.GracefulUpdateDelay }).(pulumi.IntOutput)
}

// Interval (sec) before peer considered dead.
func (o LookupBgpResultOutput) HoldtimeTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.HoldtimeTimer }).(pulumi.IntOutput)
}

// Enable/disable IBGP multi-path.
func (o LookupBgpResultOutput) IbgpMultipath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.IbgpMultipath }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupBgpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.Id }).(pulumi.StringOutput)
}

// Don't send unknown optional capability notification message
func (o LookupBgpResultOutput) IgnoreOptionalCapability() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.IgnoreOptionalCapability }).(pulumi.StringOutput)
}

// Frequency to send keep alive requests.
func (o LookupBgpResultOutput) KeepaliveTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.KeepaliveTimer }).(pulumi.IntOutput)
}

// Enable logging of BGP neighbour's changes
func (o LookupBgpResultOutput) LogNeighbourChanges() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.LogNeighbourChanges }).(pulumi.StringOutput)
}

// Enable/disable use of recursive distance to select multipath.
func (o LookupBgpResultOutput) MultipathRecursiveDistance() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.MultipathRecursiveDistance }).(pulumi.StringOutput)
}

// Neighbor group name.
func (o LookupBgpResultOutput) NeighborGroups() GetBgpNeighborGroupArrayOutput {
	return o.ApplyT(func(v LookupBgpResult) []GetBgpNeighborGroup { return v.NeighborGroups }).(GetBgpNeighborGroupArrayOutput)
}

// BGP IPv6 neighbor range table. The structure of `neighborRange6` block is documented below.
func (o LookupBgpResultOutput) NeighborRange6s() GetBgpNeighborRange6ArrayOutput {
	return o.ApplyT(func(v LookupBgpResult) []GetBgpNeighborRange6 { return v.NeighborRange6s }).(GetBgpNeighborRange6ArrayOutput)
}

// BGP neighbor range table. The structure of `neighborRange` block is documented below.
func (o LookupBgpResultOutput) NeighborRanges() GetBgpNeighborRangeArrayOutput {
	return o.ApplyT(func(v LookupBgpResult) []GetBgpNeighborRange { return v.NeighborRanges }).(GetBgpNeighborRangeArrayOutput)
}

// BGP neighbor table. The structure of `neighbor` block is documented below.
func (o LookupBgpResultOutput) Neighbors() GetBgpNeighborArrayOutput {
	return o.ApplyT(func(v LookupBgpResult) []GetBgpNeighbor { return v.Neighbors }).(GetBgpNeighborArrayOutput)
}

// BGP IPv6 network table. The structure of `network6` block is documented below.
func (o LookupBgpResultOutput) Network6s() GetBgpNetwork6ArrayOutput {
	return o.ApplyT(func(v LookupBgpResult) []GetBgpNetwork6 { return v.Network6s }).(GetBgpNetwork6ArrayOutput)
}

// Configure insurance of BGP network route existence in IGP.
func (o LookupBgpResultOutput) NetworkImportCheck() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.NetworkImportCheck }).(pulumi.StringOutput)
}

// BGP network table. The structure of `network` block is documented below.
func (o LookupBgpResultOutput) Networks() GetBgpNetworkArrayOutput {
	return o.ApplyT(func(v LookupBgpResult) []GetBgpNetwork { return v.Networks }).(GetBgpNetworkArrayOutput)
}

// Enable/disable priority inheritance for recursive resolution.
func (o LookupBgpResultOutput) RecursiveInheritPriority() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.RecursiveInheritPriority }).(pulumi.StringOutput)
}

// Enable/disable recursive resolution of next-hop using BGP route.
func (o LookupBgpResultOutput) RecursiveNextHop() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.RecursiveNextHop }).(pulumi.StringOutput)
}

// BGP IPv6 redistribute table. The structure of `redistribute6` block is documented below.
func (o LookupBgpResultOutput) Redistribute6s() GetBgpRedistribute6ArrayOutput {
	return o.ApplyT(func(v LookupBgpResult) []GetBgpRedistribute6 { return v.Redistribute6s }).(GetBgpRedistribute6ArrayOutput)
}

// BGP IPv4 redistribute table. The structure of `redistribute` block is documented below.
func (o LookupBgpResultOutput) Redistributes() GetBgpRedistributeArrayOutput {
	return o.ApplyT(func(v LookupBgpResult) []GetBgpRedistribute { return v.Redistributes }).(GetBgpRedistributeArrayOutput)
}

// Router ID.
func (o LookupBgpResultOutput) RouterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.RouterId }).(pulumi.StringOutput)
}

// Background scanner interval (sec), 0 to disable it.
func (o LookupBgpResultOutput) ScanTime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpResult) int { return v.ScanTime }).(pulumi.IntOutput)
}

// Enable/disable only advertise routes from iBGP if routes present in an IGP.
func (o LookupBgpResultOutput) Synchronization() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.Synchronization }).(pulumi.StringOutput)
}

// Configure tag-match mode. Resolves BGP routes with other routes containing the same tag.
func (o LookupBgpResultOutput) TagResolveMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpResult) string { return v.TagResolveMode }).(pulumi.StringOutput)
}

func (o LookupBgpResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBgpResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// BGP IPv6 VRF leaking table. The structure of `vrf6` block is documented below.
func (o LookupBgpResultOutput) Vrf6s() GetBgpVrf6ArrayOutput {
	return o.ApplyT(func(v LookupBgpResult) []GetBgpVrf6 { return v.Vrf6s }).(GetBgpVrf6ArrayOutput)
}

// BGP IPv6 VRF leaking table. The structure of `vrfLeak6` block is documented below.
func (o LookupBgpResultOutput) VrfLeak6s() GetBgpVrfLeak6ArrayOutput {
	return o.ApplyT(func(v LookupBgpResult) []GetBgpVrfLeak6 { return v.VrfLeak6s }).(GetBgpVrfLeak6ArrayOutput)
}

// BGP VRF leaking table. The structure of `vrfLeak` block is documented below.
func (o LookupBgpResultOutput) VrfLeaks() GetBgpVrfLeakArrayOutput {
	return o.ApplyT(func(v LookupBgpResult) []GetBgpVrfLeak { return v.VrfLeaks }).(GetBgpVrfLeakArrayOutput)
}

// Target VRF ID <0 - 31>.
func (o LookupBgpResultOutput) Vrves() GetBgpVrfArrayOutput {
	return o.ApplyT(func(v LookupBgpResult) []GetBgpVrf { return v.Vrves }).(GetBgpVrfArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBgpResultOutput{})
}
