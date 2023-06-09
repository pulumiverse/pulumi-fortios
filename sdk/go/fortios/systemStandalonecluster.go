// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes. Applies to FortiOS Version `>= 6.4.0`.
//
// ## Import
//
// # System StandaloneCluster can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemStandalonecluster:SystemStandalonecluster labelname SystemStandaloneCluster
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemStandalonecluster:SystemStandalonecluster labelname SystemStandaloneCluster
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemStandalonecluster struct {
	pulumi.CustomResourceState

	// Configure FortiGate Session Life Support Protocol (FGSP) session synchronization. The structure of `clusterPeer` block is documented below.
	ClusterPeers SystemStandaloneclusterClusterPeerArrayOutput `pulumi:"clusterPeers"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
	Encryption pulumi.StringOutput `pulumi:"encryption"`
	// Cluster member ID (0 - 3).
	GroupMemberId pulumi.IntOutput `pulumi:"groupMemberId"`
	// Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
	Layer2Connection pulumi.StringOutput `pulumi:"layer2Connection"`
	// Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
	Psksecret pulumi.StringPtrOutput `pulumi:"psksecret"`
	// Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
	SessionSyncDev pulumi.StringOutput `pulumi:"sessionSyncDev"`
	// Cluster group ID (0 - 255). Must be the same for all members.
	StandaloneGroupId pulumi.IntOutput `pulumi:"standaloneGroupId"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemStandalonecluster registers a new resource with the given unique name, arguments, and options.
func NewSystemStandalonecluster(ctx *pulumi.Context,
	name string, args *SystemStandaloneclusterArgs, opts ...pulumi.ResourceOption) (*SystemStandalonecluster, error) {
	if args == nil {
		args = &SystemStandaloneclusterArgs{}
	}

	if args.Psksecret != nil {
		args.Psksecret = pulumi.ToSecret(args.Psksecret).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"psksecret",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemStandalonecluster
	err := ctx.RegisterResource("fortios:index/systemStandalonecluster:SystemStandalonecluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemStandalonecluster gets an existing SystemStandalonecluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemStandalonecluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemStandaloneclusterState, opts ...pulumi.ResourceOption) (*SystemStandalonecluster, error) {
	var resource SystemStandalonecluster
	err := ctx.ReadResource("fortios:index/systemStandalonecluster:SystemStandalonecluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemStandalonecluster resources.
type systemStandaloneclusterState struct {
	// Configure FortiGate Session Life Support Protocol (FGSP) session synchronization. The structure of `clusterPeer` block is documented below.
	ClusterPeers []SystemStandaloneclusterClusterPeer `pulumi:"clusterPeers"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
	Encryption *string `pulumi:"encryption"`
	// Cluster member ID (0 - 3).
	GroupMemberId *int `pulumi:"groupMemberId"`
	// Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
	Layer2Connection *string `pulumi:"layer2Connection"`
	// Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
	Psksecret *string `pulumi:"psksecret"`
	// Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
	SessionSyncDev *string `pulumi:"sessionSyncDev"`
	// Cluster group ID (0 - 255). Must be the same for all members.
	StandaloneGroupId *int `pulumi:"standaloneGroupId"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemStandaloneclusterState struct {
	// Configure FortiGate Session Life Support Protocol (FGSP) session synchronization. The structure of `clusterPeer` block is documented below.
	ClusterPeers SystemStandaloneclusterClusterPeerArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
	Encryption pulumi.StringPtrInput
	// Cluster member ID (0 - 3).
	GroupMemberId pulumi.IntPtrInput
	// Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
	Layer2Connection pulumi.StringPtrInput
	// Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
	Psksecret pulumi.StringPtrInput
	// Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
	SessionSyncDev pulumi.StringPtrInput
	// Cluster group ID (0 - 255). Must be the same for all members.
	StandaloneGroupId pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemStandaloneclusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemStandaloneclusterState)(nil)).Elem()
}

type systemStandaloneclusterArgs struct {
	// Configure FortiGate Session Life Support Protocol (FGSP) session synchronization. The structure of `clusterPeer` block is documented below.
	ClusterPeers []SystemStandaloneclusterClusterPeer `pulumi:"clusterPeers"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
	Encryption *string `pulumi:"encryption"`
	// Cluster member ID (0 - 3).
	GroupMemberId *int `pulumi:"groupMemberId"`
	// Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
	Layer2Connection *string `pulumi:"layer2Connection"`
	// Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
	Psksecret *string `pulumi:"psksecret"`
	// Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
	SessionSyncDev *string `pulumi:"sessionSyncDev"`
	// Cluster group ID (0 - 255). Must be the same for all members.
	StandaloneGroupId *int `pulumi:"standaloneGroupId"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemStandalonecluster resource.
type SystemStandaloneclusterArgs struct {
	// Configure FortiGate Session Life Support Protocol (FGSP) session synchronization. The structure of `clusterPeer` block is documented below.
	ClusterPeers SystemStandaloneclusterClusterPeerArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
	Encryption pulumi.StringPtrInput
	// Cluster member ID (0 - 3).
	GroupMemberId pulumi.IntPtrInput
	// Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
	Layer2Connection pulumi.StringPtrInput
	// Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
	Psksecret pulumi.StringPtrInput
	// Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
	SessionSyncDev pulumi.StringPtrInput
	// Cluster group ID (0 - 255). Must be the same for all members.
	StandaloneGroupId pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemStandaloneclusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemStandaloneclusterArgs)(nil)).Elem()
}

type SystemStandaloneclusterInput interface {
	pulumi.Input

	ToSystemStandaloneclusterOutput() SystemStandaloneclusterOutput
	ToSystemStandaloneclusterOutputWithContext(ctx context.Context) SystemStandaloneclusterOutput
}

func (*SystemStandalonecluster) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemStandalonecluster)(nil)).Elem()
}

func (i *SystemStandalonecluster) ToSystemStandaloneclusterOutput() SystemStandaloneclusterOutput {
	return i.ToSystemStandaloneclusterOutputWithContext(context.Background())
}

func (i *SystemStandalonecluster) ToSystemStandaloneclusterOutputWithContext(ctx context.Context) SystemStandaloneclusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemStandaloneclusterOutput)
}

// SystemStandaloneclusterArrayInput is an input type that accepts SystemStandaloneclusterArray and SystemStandaloneclusterArrayOutput values.
// You can construct a concrete instance of `SystemStandaloneclusterArrayInput` via:
//
//	SystemStandaloneclusterArray{ SystemStandaloneclusterArgs{...} }
type SystemStandaloneclusterArrayInput interface {
	pulumi.Input

	ToSystemStandaloneclusterArrayOutput() SystemStandaloneclusterArrayOutput
	ToSystemStandaloneclusterArrayOutputWithContext(context.Context) SystemStandaloneclusterArrayOutput
}

type SystemStandaloneclusterArray []SystemStandaloneclusterInput

func (SystemStandaloneclusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemStandalonecluster)(nil)).Elem()
}

func (i SystemStandaloneclusterArray) ToSystemStandaloneclusterArrayOutput() SystemStandaloneclusterArrayOutput {
	return i.ToSystemStandaloneclusterArrayOutputWithContext(context.Background())
}

func (i SystemStandaloneclusterArray) ToSystemStandaloneclusterArrayOutputWithContext(ctx context.Context) SystemStandaloneclusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemStandaloneclusterArrayOutput)
}

// SystemStandaloneclusterMapInput is an input type that accepts SystemStandaloneclusterMap and SystemStandaloneclusterMapOutput values.
// You can construct a concrete instance of `SystemStandaloneclusterMapInput` via:
//
//	SystemStandaloneclusterMap{ "key": SystemStandaloneclusterArgs{...} }
type SystemStandaloneclusterMapInput interface {
	pulumi.Input

	ToSystemStandaloneclusterMapOutput() SystemStandaloneclusterMapOutput
	ToSystemStandaloneclusterMapOutputWithContext(context.Context) SystemStandaloneclusterMapOutput
}

type SystemStandaloneclusterMap map[string]SystemStandaloneclusterInput

func (SystemStandaloneclusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemStandalonecluster)(nil)).Elem()
}

func (i SystemStandaloneclusterMap) ToSystemStandaloneclusterMapOutput() SystemStandaloneclusterMapOutput {
	return i.ToSystemStandaloneclusterMapOutputWithContext(context.Background())
}

func (i SystemStandaloneclusterMap) ToSystemStandaloneclusterMapOutputWithContext(ctx context.Context) SystemStandaloneclusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemStandaloneclusterMapOutput)
}

type SystemStandaloneclusterOutput struct{ *pulumi.OutputState }

func (SystemStandaloneclusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemStandalonecluster)(nil)).Elem()
}

func (o SystemStandaloneclusterOutput) ToSystemStandaloneclusterOutput() SystemStandaloneclusterOutput {
	return o
}

func (o SystemStandaloneclusterOutput) ToSystemStandaloneclusterOutputWithContext(ctx context.Context) SystemStandaloneclusterOutput {
	return o
}

// Configure FortiGate Session Life Support Protocol (FGSP) session synchronization. The structure of `clusterPeer` block is documented below.
func (o SystemStandaloneclusterOutput) ClusterPeers() SystemStandaloneclusterClusterPeerArrayOutput {
	return o.ApplyT(func(v *SystemStandalonecluster) SystemStandaloneclusterClusterPeerArrayOutput { return v.ClusterPeers }).(SystemStandaloneclusterClusterPeerArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SystemStandaloneclusterOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemStandalonecluster) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
func (o SystemStandaloneclusterOutput) Encryption() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemStandalonecluster) pulumi.StringOutput { return v.Encryption }).(pulumi.StringOutput)
}

// Cluster member ID (0 - 3).
func (o SystemStandaloneclusterOutput) GroupMemberId() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemStandalonecluster) pulumi.IntOutput { return v.GroupMemberId }).(pulumi.IntOutput)
}

// Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
func (o SystemStandaloneclusterOutput) Layer2Connection() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemStandalonecluster) pulumi.StringOutput { return v.Layer2Connection }).(pulumi.StringOutput)
}

// Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
func (o SystemStandaloneclusterOutput) Psksecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemStandalonecluster) pulumi.StringPtrOutput { return v.Psksecret }).(pulumi.StringPtrOutput)
}

// Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
func (o SystemStandaloneclusterOutput) SessionSyncDev() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemStandalonecluster) pulumi.StringOutput { return v.SessionSyncDev }).(pulumi.StringOutput)
}

// Cluster group ID (0 - 255). Must be the same for all members.
func (o SystemStandaloneclusterOutput) StandaloneGroupId() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemStandalonecluster) pulumi.IntOutput { return v.StandaloneGroupId }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemStandaloneclusterOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemStandalonecluster) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemStandaloneclusterArrayOutput struct{ *pulumi.OutputState }

func (SystemStandaloneclusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemStandalonecluster)(nil)).Elem()
}

func (o SystemStandaloneclusterArrayOutput) ToSystemStandaloneclusterArrayOutput() SystemStandaloneclusterArrayOutput {
	return o
}

func (o SystemStandaloneclusterArrayOutput) ToSystemStandaloneclusterArrayOutputWithContext(ctx context.Context) SystemStandaloneclusterArrayOutput {
	return o
}

func (o SystemStandaloneclusterArrayOutput) Index(i pulumi.IntInput) SystemStandaloneclusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemStandalonecluster {
		return vs[0].([]*SystemStandalonecluster)[vs[1].(int)]
	}).(SystemStandaloneclusterOutput)
}

type SystemStandaloneclusterMapOutput struct{ *pulumi.OutputState }

func (SystemStandaloneclusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemStandalonecluster)(nil)).Elem()
}

func (o SystemStandaloneclusterMapOutput) ToSystemStandaloneclusterMapOutput() SystemStandaloneclusterMapOutput {
	return o
}

func (o SystemStandaloneclusterMapOutput) ToSystemStandaloneclusterMapOutputWithContext(ctx context.Context) SystemStandaloneclusterMapOutput {
	return o
}

func (o SystemStandaloneclusterMapOutput) MapIndex(k pulumi.StringInput) SystemStandaloneclusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemStandalonecluster {
		return vs[0].(map[string]*SystemStandalonecluster)[vs[1].(string)]
	}).(SystemStandaloneclusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemStandaloneclusterInput)(nil)).Elem(), &SystemStandalonecluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemStandaloneclusterArrayInput)(nil)).Elem(), SystemStandaloneclusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemStandaloneclusterMapInput)(nil)).Elem(), SystemStandaloneclusterMap{})
	pulumi.RegisterOutputType(SystemStandaloneclusterOutput{})
	pulumi.RegisterOutputType(SystemStandaloneclusterArrayOutput{})
	pulumi.RegisterOutputType(SystemStandaloneclusterMapOutput{})
}
