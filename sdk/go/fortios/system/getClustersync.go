// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on an fortios system clustersync
func LookupClustersync(ctx *pulumi.Context, args *LookupClustersyncArgs, opts ...pulumi.InvokeOption) (*LookupClustersyncResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupClustersyncResult
	err := ctx.Invoke("fortios:system/getClustersync:getClustersync", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClustersync.
type LookupClustersyncArgs struct {
	// Specify the syncId of the desired system clustersync.
	SyncId int `pulumi:"syncId"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getClustersync.
type LookupClustersyncResult struct {
	// List of interfaces to be turned down before session synchronization is complete. The structure of `downIntfsBeforeSessSync` block is documented below.
	DownIntfsBeforeSessSyncs []GetClustersyncDownIntfsBeforeSessSync `pulumi:"downIntfsBeforeSessSyncs"`
	// Heartbeat interval (1 - 10 sec).
	HbInterval int `pulumi:"hbInterval"`
	// Lost heartbeat threshold (1 - 10).
	HbLostThreshold int `pulumi:"hbLostThreshold"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IKE heartbeat interval (1 - 60 secs).
	IkeHeartbeatInterval int `pulumi:"ikeHeartbeatInterval"`
	// Enable/disable IKE HA monitor.
	IkeMonitor string `pulumi:"ikeMonitor"`
	// IKE HA monitor interval (10 - 300 secs).
	IkeMonitorInterval int `pulumi:"ikeMonitorInterval"`
	// Enable/disable IPsec tunnel synchronization.
	IpsecTunnelSync string `pulumi:"ipsecTunnelSync"`
	// IP address of the interface on the peer unit that is used for the session synchronization link.
	Peerip string `pulumi:"peerip"`
	// VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.
	Peervd string `pulumi:"peervd"`
	// Enable/disable IKE route announcement on the backup unit.
	SecondaryAddIpsecRoutes string `pulumi:"secondaryAddIpsecRoutes"`
	// Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize. The structure of `sessionSyncFilter` block is documented below.
	SessionSyncFilters []GetClustersyncSessionSyncFilter `pulumi:"sessionSyncFilters"`
	// Enable/disable IKE route announcement on the backup unit.
	SlaveAddIkeRoutes string `pulumi:"slaveAddIkeRoutes"`
	// Sync ID.
	SyncId int `pulumi:"syncId"`
	// Sessions from these VDOMs are synchronized using this session synchronization configuration. The structure of `syncvd` block is documented below.
	Syncvds   []GetClustersyncSyncvd `pulumi:"syncvds"`
	Vdomparam *string                `pulumi:"vdomparam"`
}

func LookupClustersyncOutput(ctx *pulumi.Context, args LookupClustersyncOutputArgs, opts ...pulumi.InvokeOption) LookupClustersyncResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClustersyncResult, error) {
			args := v.(LookupClustersyncArgs)
			r, err := LookupClustersync(ctx, &args, opts...)
			var s LookupClustersyncResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClustersyncResultOutput)
}

// A collection of arguments for invoking getClustersync.
type LookupClustersyncOutputArgs struct {
	// Specify the syncId of the desired system clustersync.
	SyncId pulumi.IntInput `pulumi:"syncId"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupClustersyncOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClustersyncArgs)(nil)).Elem()
}

// A collection of values returned by getClustersync.
type LookupClustersyncResultOutput struct{ *pulumi.OutputState }

func (LookupClustersyncResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClustersyncResult)(nil)).Elem()
}

func (o LookupClustersyncResultOutput) ToLookupClustersyncResultOutput() LookupClustersyncResultOutput {
	return o
}

func (o LookupClustersyncResultOutput) ToLookupClustersyncResultOutputWithContext(ctx context.Context) LookupClustersyncResultOutput {
	return o
}

// List of interfaces to be turned down before session synchronization is complete. The structure of `downIntfsBeforeSessSync` block is documented below.
func (o LookupClustersyncResultOutput) DownIntfsBeforeSessSyncs() GetClustersyncDownIntfsBeforeSessSyncArrayOutput {
	return o.ApplyT(func(v LookupClustersyncResult) []GetClustersyncDownIntfsBeforeSessSync {
		return v.DownIntfsBeforeSessSyncs
	}).(GetClustersyncDownIntfsBeforeSessSyncArrayOutput)
}

// Heartbeat interval (1 - 10 sec).
func (o LookupClustersyncResultOutput) HbInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClustersyncResult) int { return v.HbInterval }).(pulumi.IntOutput)
}

// Lost heartbeat threshold (1 - 10).
func (o LookupClustersyncResultOutput) HbLostThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClustersyncResult) int { return v.HbLostThreshold }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupClustersyncResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClustersyncResult) string { return v.Id }).(pulumi.StringOutput)
}

// IKE heartbeat interval (1 - 60 secs).
func (o LookupClustersyncResultOutput) IkeHeartbeatInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClustersyncResult) int { return v.IkeHeartbeatInterval }).(pulumi.IntOutput)
}

// Enable/disable IKE HA monitor.
func (o LookupClustersyncResultOutput) IkeMonitor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClustersyncResult) string { return v.IkeMonitor }).(pulumi.StringOutput)
}

// IKE HA monitor interval (10 - 300 secs).
func (o LookupClustersyncResultOutput) IkeMonitorInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClustersyncResult) int { return v.IkeMonitorInterval }).(pulumi.IntOutput)
}

// Enable/disable IPsec tunnel synchronization.
func (o LookupClustersyncResultOutput) IpsecTunnelSync() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClustersyncResult) string { return v.IpsecTunnelSync }).(pulumi.StringOutput)
}

// IP address of the interface on the peer unit that is used for the session synchronization link.
func (o LookupClustersyncResultOutput) Peerip() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClustersyncResult) string { return v.Peerip }).(pulumi.StringOutput)
}

// VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.
func (o LookupClustersyncResultOutput) Peervd() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClustersyncResult) string { return v.Peervd }).(pulumi.StringOutput)
}

// Enable/disable IKE route announcement on the backup unit.
func (o LookupClustersyncResultOutput) SecondaryAddIpsecRoutes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClustersyncResult) string { return v.SecondaryAddIpsecRoutes }).(pulumi.StringOutput)
}

// Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize. The structure of `sessionSyncFilter` block is documented below.
func (o LookupClustersyncResultOutput) SessionSyncFilters() GetClustersyncSessionSyncFilterArrayOutput {
	return o.ApplyT(func(v LookupClustersyncResult) []GetClustersyncSessionSyncFilter { return v.SessionSyncFilters }).(GetClustersyncSessionSyncFilterArrayOutput)
}

// Enable/disable IKE route announcement on the backup unit.
func (o LookupClustersyncResultOutput) SlaveAddIkeRoutes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClustersyncResult) string { return v.SlaveAddIkeRoutes }).(pulumi.StringOutput)
}

// Sync ID.
func (o LookupClustersyncResultOutput) SyncId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClustersyncResult) int { return v.SyncId }).(pulumi.IntOutput)
}

// Sessions from these VDOMs are synchronized using this session synchronization configuration. The structure of `syncvd` block is documented below.
func (o LookupClustersyncResultOutput) Syncvds() GetClustersyncSyncvdArrayOutput {
	return o.ApplyT(func(v LookupClustersyncResult) []GetClustersyncSyncvd { return v.Syncvds }).(GetClustersyncSyncvdArrayOutput)
}

func (o LookupClustersyncResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClustersyncResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClustersyncResultOutput{})
}
