// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPS global parameter.
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
//			_, err := fortios.NewIpsGlobal(ctx, "trname", &fortios.IpsGlobalArgs{
//				AnomalyMode:        pulumi.String("continuous"),
//				Database:           pulumi.String("regular"),
//				DeepAppInspDbLimit: pulumi.Int(0),
//				DeepAppInspTimeout: pulumi.Int(0),
//				EngineCount:        pulumi.Int(0),
//				ExcludeSignatures:  pulumi.String("industrial"),
//				FailOpen:           pulumi.String("disable"),
//				IntelligentMode:    pulumi.String("enable"),
//				SessionLimitMode:   pulumi.String("heuristic"),
//				SocketSize:         pulumi.Int(0),
//				SyncSessionTtl:     pulumi.String("enable"),
//				TrafficSubmit:      pulumi.String("disable"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # Ips Global can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/ipsGlobal:IpsGlobal labelname IpsGlobal
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/ipsGlobal:IpsGlobal labelname IpsGlobal
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type IpsGlobal struct {
	pulumi.CustomResourceState

	// Global blocking mode for rate-based anomalies. Valid values: `periodical`, `continuous`.
	AnomalyMode pulumi.StringOutput `pulumi:"anomalyMode"`
	// IPS Pattern matching acceleration/offloading to CPx processors. Valid values: `none`, `basic`, `advanced`.
	CpAccelMode pulumi.StringOutput `pulumi:"cpAccelMode"`
	// Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks. Valid values: `regular`, `extended`.
	Database pulumi.StringOutput `pulumi:"database"`
	// Limit on number of entries in deep application inspection database (1 - 2147483647, 0 = use recommended setting)
	DeepAppInspDbLimit pulumi.IntOutput `pulumi:"deepAppInspDbLimit"`
	// Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
	DeepAppInspTimeout pulumi.IntOutput `pulumi:"deepAppInspTimeout"`
	// Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
	EngineCount pulumi.IntOutput `pulumi:"engineCount"`
	// Excluded signatures. Valid values: `none`, `industrial`.
	ExcludeSignatures pulumi.StringOutput `pulumi:"excludeSignatures"`
	// Enable to allow traffic if the IPS process crashes. Default is disable and IPS traffic is blocked when the IPS process crashes. Valid values: `enable`, `disable`.
	FailOpen pulumi.StringOutput `pulumi:"failOpen"`
	// Enable/disable IPS adaptive scanning (intelligent mode). Intelligent mode optimizes the scanning method for the type of traffic. Valid values: `enable`, `disable`.
	IntelligentMode pulumi.StringOutput `pulumi:"intelligentMode"`
	// Enable/disable IPS daemon's use of CPUs other than CPU 0 Valid values: `disable`, `enable`.
	IpsReserveCpu pulumi.StringOutput `pulumi:"ipsReserveCpu"`
	// NGFW policy-mode app detection threshold.
	NgfwMaxScanRange pulumi.IntOutput `pulumi:"ngfwMaxScanRange"`
	// Acceleration mode for IPS processing by NPx processors. Valid values: `none`, `basic`.
	NpAccelMode pulumi.StringOutput `pulumi:"npAccelMode"`
	// Packet/pcap log queue depth per IPS engine.
	PacketLogQueueDepth pulumi.IntOutput `pulumi:"packetLogQueueDepth"`
	// Method of counting concurrent sessions used by session limit anomalies. Choose between greater accuracy (accurate) or improved performance (heuristics). Valid values: `accurate`, `heuristic`.
	SessionLimitMode pulumi.StringOutput `pulumi:"sessionLimitMode"`
	// Public IP addresses of your network that receive Skype sessions. Helps identify Skype sessions. Separate IP addresses with commas.
	SkypeClientPublicIpaddr pulumi.StringPtrOutput `pulumi:"skypeClientPublicIpaddr"`
	// IPS socket buffer size (0 - 256 MB). Default depends on available memory. Can be changed to tune performance.
	SocketSize pulumi.IntOutput `pulumi:"socketSize"`
	// Enable/disable use of kernel session TTL for IPS sessions. Valid values: `enable`, `disable`.
	SyncSessionTtl pulumi.StringOutput `pulumi:"syncSessionTtl"`
	// TLS active probe configuration. The structure of `tlsActiveProbe` block is documented below.
	TlsActiveProbe IpsGlobalTlsActiveProbeOutput `pulumi:"tlsActiveProbe"`
	// Enable/disable submitting attack data found by this FortiGate to FortiGuard. Valid values: `enable`, `disable`.
	TrafficSubmit pulumi.StringOutput `pulumi:"trafficSubmit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewIpsGlobal registers a new resource with the given unique name, arguments, and options.
func NewIpsGlobal(ctx *pulumi.Context,
	name string, args *IpsGlobalArgs, opts ...pulumi.ResourceOption) (*IpsGlobal, error) {
	if args == nil {
		args = &IpsGlobalArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource IpsGlobal
	err := ctx.RegisterResource("fortios:index/ipsGlobal:IpsGlobal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpsGlobal gets an existing IpsGlobal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpsGlobal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpsGlobalState, opts ...pulumi.ResourceOption) (*IpsGlobal, error) {
	var resource IpsGlobal
	err := ctx.ReadResource("fortios:index/ipsGlobal:IpsGlobal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpsGlobal resources.
type ipsGlobalState struct {
	// Global blocking mode for rate-based anomalies. Valid values: `periodical`, `continuous`.
	AnomalyMode *string `pulumi:"anomalyMode"`
	// IPS Pattern matching acceleration/offloading to CPx processors. Valid values: `none`, `basic`, `advanced`.
	CpAccelMode *string `pulumi:"cpAccelMode"`
	// Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks. Valid values: `regular`, `extended`.
	Database *string `pulumi:"database"`
	// Limit on number of entries in deep application inspection database (1 - 2147483647, 0 = use recommended setting)
	DeepAppInspDbLimit *int `pulumi:"deepAppInspDbLimit"`
	// Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
	DeepAppInspTimeout *int `pulumi:"deepAppInspTimeout"`
	// Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
	EngineCount *int `pulumi:"engineCount"`
	// Excluded signatures. Valid values: `none`, `industrial`.
	ExcludeSignatures *string `pulumi:"excludeSignatures"`
	// Enable to allow traffic if the IPS process crashes. Default is disable and IPS traffic is blocked when the IPS process crashes. Valid values: `enable`, `disable`.
	FailOpen *string `pulumi:"failOpen"`
	// Enable/disable IPS adaptive scanning (intelligent mode). Intelligent mode optimizes the scanning method for the type of traffic. Valid values: `enable`, `disable`.
	IntelligentMode *string `pulumi:"intelligentMode"`
	// Enable/disable IPS daemon's use of CPUs other than CPU 0 Valid values: `disable`, `enable`.
	IpsReserveCpu *string `pulumi:"ipsReserveCpu"`
	// NGFW policy-mode app detection threshold.
	NgfwMaxScanRange *int `pulumi:"ngfwMaxScanRange"`
	// Acceleration mode for IPS processing by NPx processors. Valid values: `none`, `basic`.
	NpAccelMode *string `pulumi:"npAccelMode"`
	// Packet/pcap log queue depth per IPS engine.
	PacketLogQueueDepth *int `pulumi:"packetLogQueueDepth"`
	// Method of counting concurrent sessions used by session limit anomalies. Choose between greater accuracy (accurate) or improved performance (heuristics). Valid values: `accurate`, `heuristic`.
	SessionLimitMode *string `pulumi:"sessionLimitMode"`
	// Public IP addresses of your network that receive Skype sessions. Helps identify Skype sessions. Separate IP addresses with commas.
	SkypeClientPublicIpaddr *string `pulumi:"skypeClientPublicIpaddr"`
	// IPS socket buffer size (0 - 256 MB). Default depends on available memory. Can be changed to tune performance.
	SocketSize *int `pulumi:"socketSize"`
	// Enable/disable use of kernel session TTL for IPS sessions. Valid values: `enable`, `disable`.
	SyncSessionTtl *string `pulumi:"syncSessionTtl"`
	// TLS active probe configuration. The structure of `tlsActiveProbe` block is documented below.
	TlsActiveProbe *IpsGlobalTlsActiveProbe `pulumi:"tlsActiveProbe"`
	// Enable/disable submitting attack data found by this FortiGate to FortiGuard. Valid values: `enable`, `disable`.
	TrafficSubmit *string `pulumi:"trafficSubmit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type IpsGlobalState struct {
	// Global blocking mode for rate-based anomalies. Valid values: `periodical`, `continuous`.
	AnomalyMode pulumi.StringPtrInput
	// IPS Pattern matching acceleration/offloading to CPx processors. Valid values: `none`, `basic`, `advanced`.
	CpAccelMode pulumi.StringPtrInput
	// Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks. Valid values: `regular`, `extended`.
	Database pulumi.StringPtrInput
	// Limit on number of entries in deep application inspection database (1 - 2147483647, 0 = use recommended setting)
	DeepAppInspDbLimit pulumi.IntPtrInput
	// Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
	DeepAppInspTimeout pulumi.IntPtrInput
	// Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
	EngineCount pulumi.IntPtrInput
	// Excluded signatures. Valid values: `none`, `industrial`.
	ExcludeSignatures pulumi.StringPtrInput
	// Enable to allow traffic if the IPS process crashes. Default is disable and IPS traffic is blocked when the IPS process crashes. Valid values: `enable`, `disable`.
	FailOpen pulumi.StringPtrInput
	// Enable/disable IPS adaptive scanning (intelligent mode). Intelligent mode optimizes the scanning method for the type of traffic. Valid values: `enable`, `disable`.
	IntelligentMode pulumi.StringPtrInput
	// Enable/disable IPS daemon's use of CPUs other than CPU 0 Valid values: `disable`, `enable`.
	IpsReserveCpu pulumi.StringPtrInput
	// NGFW policy-mode app detection threshold.
	NgfwMaxScanRange pulumi.IntPtrInput
	// Acceleration mode for IPS processing by NPx processors. Valid values: `none`, `basic`.
	NpAccelMode pulumi.StringPtrInput
	// Packet/pcap log queue depth per IPS engine.
	PacketLogQueueDepth pulumi.IntPtrInput
	// Method of counting concurrent sessions used by session limit anomalies. Choose between greater accuracy (accurate) or improved performance (heuristics). Valid values: `accurate`, `heuristic`.
	SessionLimitMode pulumi.StringPtrInput
	// Public IP addresses of your network that receive Skype sessions. Helps identify Skype sessions. Separate IP addresses with commas.
	SkypeClientPublicIpaddr pulumi.StringPtrInput
	// IPS socket buffer size (0 - 256 MB). Default depends on available memory. Can be changed to tune performance.
	SocketSize pulumi.IntPtrInput
	// Enable/disable use of kernel session TTL for IPS sessions. Valid values: `enable`, `disable`.
	SyncSessionTtl pulumi.StringPtrInput
	// TLS active probe configuration. The structure of `tlsActiveProbe` block is documented below.
	TlsActiveProbe IpsGlobalTlsActiveProbePtrInput
	// Enable/disable submitting attack data found by this FortiGate to FortiGuard. Valid values: `enable`, `disable`.
	TrafficSubmit pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IpsGlobalState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsGlobalState)(nil)).Elem()
}

type ipsGlobalArgs struct {
	// Global blocking mode for rate-based anomalies. Valid values: `periodical`, `continuous`.
	AnomalyMode *string `pulumi:"anomalyMode"`
	// IPS Pattern matching acceleration/offloading to CPx processors. Valid values: `none`, `basic`, `advanced`.
	CpAccelMode *string `pulumi:"cpAccelMode"`
	// Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks. Valid values: `regular`, `extended`.
	Database *string `pulumi:"database"`
	// Limit on number of entries in deep application inspection database (1 - 2147483647, 0 = use recommended setting)
	DeepAppInspDbLimit *int `pulumi:"deepAppInspDbLimit"`
	// Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
	DeepAppInspTimeout *int `pulumi:"deepAppInspTimeout"`
	// Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
	EngineCount *int `pulumi:"engineCount"`
	// Excluded signatures. Valid values: `none`, `industrial`.
	ExcludeSignatures *string `pulumi:"excludeSignatures"`
	// Enable to allow traffic if the IPS process crashes. Default is disable and IPS traffic is blocked when the IPS process crashes. Valid values: `enable`, `disable`.
	FailOpen *string `pulumi:"failOpen"`
	// Enable/disable IPS adaptive scanning (intelligent mode). Intelligent mode optimizes the scanning method for the type of traffic. Valid values: `enable`, `disable`.
	IntelligentMode *string `pulumi:"intelligentMode"`
	// Enable/disable IPS daemon's use of CPUs other than CPU 0 Valid values: `disable`, `enable`.
	IpsReserveCpu *string `pulumi:"ipsReserveCpu"`
	// NGFW policy-mode app detection threshold.
	NgfwMaxScanRange *int `pulumi:"ngfwMaxScanRange"`
	// Acceleration mode for IPS processing by NPx processors. Valid values: `none`, `basic`.
	NpAccelMode *string `pulumi:"npAccelMode"`
	// Packet/pcap log queue depth per IPS engine.
	PacketLogQueueDepth *int `pulumi:"packetLogQueueDepth"`
	// Method of counting concurrent sessions used by session limit anomalies. Choose between greater accuracy (accurate) or improved performance (heuristics). Valid values: `accurate`, `heuristic`.
	SessionLimitMode *string `pulumi:"sessionLimitMode"`
	// Public IP addresses of your network that receive Skype sessions. Helps identify Skype sessions. Separate IP addresses with commas.
	SkypeClientPublicIpaddr *string `pulumi:"skypeClientPublicIpaddr"`
	// IPS socket buffer size (0 - 256 MB). Default depends on available memory. Can be changed to tune performance.
	SocketSize *int `pulumi:"socketSize"`
	// Enable/disable use of kernel session TTL for IPS sessions. Valid values: `enable`, `disable`.
	SyncSessionTtl *string `pulumi:"syncSessionTtl"`
	// TLS active probe configuration. The structure of `tlsActiveProbe` block is documented below.
	TlsActiveProbe *IpsGlobalTlsActiveProbe `pulumi:"tlsActiveProbe"`
	// Enable/disable submitting attack data found by this FortiGate to FortiGuard. Valid values: `enable`, `disable`.
	TrafficSubmit *string `pulumi:"trafficSubmit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a IpsGlobal resource.
type IpsGlobalArgs struct {
	// Global blocking mode for rate-based anomalies. Valid values: `periodical`, `continuous`.
	AnomalyMode pulumi.StringPtrInput
	// IPS Pattern matching acceleration/offloading to CPx processors. Valid values: `none`, `basic`, `advanced`.
	CpAccelMode pulumi.StringPtrInput
	// Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks. Valid values: `regular`, `extended`.
	Database pulumi.StringPtrInput
	// Limit on number of entries in deep application inspection database (1 - 2147483647, 0 = use recommended setting)
	DeepAppInspDbLimit pulumi.IntPtrInput
	// Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
	DeepAppInspTimeout pulumi.IntPtrInput
	// Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
	EngineCount pulumi.IntPtrInput
	// Excluded signatures. Valid values: `none`, `industrial`.
	ExcludeSignatures pulumi.StringPtrInput
	// Enable to allow traffic if the IPS process crashes. Default is disable and IPS traffic is blocked when the IPS process crashes. Valid values: `enable`, `disable`.
	FailOpen pulumi.StringPtrInput
	// Enable/disable IPS adaptive scanning (intelligent mode). Intelligent mode optimizes the scanning method for the type of traffic. Valid values: `enable`, `disable`.
	IntelligentMode pulumi.StringPtrInput
	// Enable/disable IPS daemon's use of CPUs other than CPU 0 Valid values: `disable`, `enable`.
	IpsReserveCpu pulumi.StringPtrInput
	// NGFW policy-mode app detection threshold.
	NgfwMaxScanRange pulumi.IntPtrInput
	// Acceleration mode for IPS processing by NPx processors. Valid values: `none`, `basic`.
	NpAccelMode pulumi.StringPtrInput
	// Packet/pcap log queue depth per IPS engine.
	PacketLogQueueDepth pulumi.IntPtrInput
	// Method of counting concurrent sessions used by session limit anomalies. Choose between greater accuracy (accurate) or improved performance (heuristics). Valid values: `accurate`, `heuristic`.
	SessionLimitMode pulumi.StringPtrInput
	// Public IP addresses of your network that receive Skype sessions. Helps identify Skype sessions. Separate IP addresses with commas.
	SkypeClientPublicIpaddr pulumi.StringPtrInput
	// IPS socket buffer size (0 - 256 MB). Default depends on available memory. Can be changed to tune performance.
	SocketSize pulumi.IntPtrInput
	// Enable/disable use of kernel session TTL for IPS sessions. Valid values: `enable`, `disable`.
	SyncSessionTtl pulumi.StringPtrInput
	// TLS active probe configuration. The structure of `tlsActiveProbe` block is documented below.
	TlsActiveProbe IpsGlobalTlsActiveProbePtrInput
	// Enable/disable submitting attack data found by this FortiGate to FortiGuard. Valid values: `enable`, `disable`.
	TrafficSubmit pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IpsGlobalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsGlobalArgs)(nil)).Elem()
}

type IpsGlobalInput interface {
	pulumi.Input

	ToIpsGlobalOutput() IpsGlobalOutput
	ToIpsGlobalOutputWithContext(ctx context.Context) IpsGlobalOutput
}

func (*IpsGlobal) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsGlobal)(nil)).Elem()
}

func (i *IpsGlobal) ToIpsGlobalOutput() IpsGlobalOutput {
	return i.ToIpsGlobalOutputWithContext(context.Background())
}

func (i *IpsGlobal) ToIpsGlobalOutputWithContext(ctx context.Context) IpsGlobalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsGlobalOutput)
}

// IpsGlobalArrayInput is an input type that accepts IpsGlobalArray and IpsGlobalArrayOutput values.
// You can construct a concrete instance of `IpsGlobalArrayInput` via:
//
//	IpsGlobalArray{ IpsGlobalArgs{...} }
type IpsGlobalArrayInput interface {
	pulumi.Input

	ToIpsGlobalArrayOutput() IpsGlobalArrayOutput
	ToIpsGlobalArrayOutputWithContext(context.Context) IpsGlobalArrayOutput
}

type IpsGlobalArray []IpsGlobalInput

func (IpsGlobalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsGlobal)(nil)).Elem()
}

func (i IpsGlobalArray) ToIpsGlobalArrayOutput() IpsGlobalArrayOutput {
	return i.ToIpsGlobalArrayOutputWithContext(context.Background())
}

func (i IpsGlobalArray) ToIpsGlobalArrayOutputWithContext(ctx context.Context) IpsGlobalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsGlobalArrayOutput)
}

// IpsGlobalMapInput is an input type that accepts IpsGlobalMap and IpsGlobalMapOutput values.
// You can construct a concrete instance of `IpsGlobalMapInput` via:
//
//	IpsGlobalMap{ "key": IpsGlobalArgs{...} }
type IpsGlobalMapInput interface {
	pulumi.Input

	ToIpsGlobalMapOutput() IpsGlobalMapOutput
	ToIpsGlobalMapOutputWithContext(context.Context) IpsGlobalMapOutput
}

type IpsGlobalMap map[string]IpsGlobalInput

func (IpsGlobalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsGlobal)(nil)).Elem()
}

func (i IpsGlobalMap) ToIpsGlobalMapOutput() IpsGlobalMapOutput {
	return i.ToIpsGlobalMapOutputWithContext(context.Background())
}

func (i IpsGlobalMap) ToIpsGlobalMapOutputWithContext(ctx context.Context) IpsGlobalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsGlobalMapOutput)
}

type IpsGlobalOutput struct{ *pulumi.OutputState }

func (IpsGlobalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsGlobal)(nil)).Elem()
}

func (o IpsGlobalOutput) ToIpsGlobalOutput() IpsGlobalOutput {
	return o
}

func (o IpsGlobalOutput) ToIpsGlobalOutputWithContext(ctx context.Context) IpsGlobalOutput {
	return o
}

// Global blocking mode for rate-based anomalies. Valid values: `periodical`, `continuous`.
func (o IpsGlobalOutput) AnomalyMode() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsGlobal) pulumi.StringOutput { return v.AnomalyMode }).(pulumi.StringOutput)
}

// IPS Pattern matching acceleration/offloading to CPx processors. Valid values: `none`, `basic`, `advanced`.
func (o IpsGlobalOutput) CpAccelMode() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsGlobal) pulumi.StringOutput { return v.CpAccelMode }).(pulumi.StringOutput)
}

// Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks. Valid values: `regular`, `extended`.
func (o IpsGlobalOutput) Database() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsGlobal) pulumi.StringOutput { return v.Database }).(pulumi.StringOutput)
}

// Limit on number of entries in deep application inspection database (1 - 2147483647, 0 = use recommended setting)
func (o IpsGlobalOutput) DeepAppInspDbLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *IpsGlobal) pulumi.IntOutput { return v.DeepAppInspDbLimit }).(pulumi.IntOutput)
}

// Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
func (o IpsGlobalOutput) DeepAppInspTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *IpsGlobal) pulumi.IntOutput { return v.DeepAppInspTimeout }).(pulumi.IntOutput)
}

// Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
func (o IpsGlobalOutput) EngineCount() pulumi.IntOutput {
	return o.ApplyT(func(v *IpsGlobal) pulumi.IntOutput { return v.EngineCount }).(pulumi.IntOutput)
}

// Excluded signatures. Valid values: `none`, `industrial`.
func (o IpsGlobalOutput) ExcludeSignatures() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsGlobal) pulumi.StringOutput { return v.ExcludeSignatures }).(pulumi.StringOutput)
}

// Enable to allow traffic if the IPS process crashes. Default is disable and IPS traffic is blocked when the IPS process crashes. Valid values: `enable`, `disable`.
func (o IpsGlobalOutput) FailOpen() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsGlobal) pulumi.StringOutput { return v.FailOpen }).(pulumi.StringOutput)
}

// Enable/disable IPS adaptive scanning (intelligent mode). Intelligent mode optimizes the scanning method for the type of traffic. Valid values: `enable`, `disable`.
func (o IpsGlobalOutput) IntelligentMode() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsGlobal) pulumi.StringOutput { return v.IntelligentMode }).(pulumi.StringOutput)
}

// Enable/disable IPS daemon's use of CPUs other than CPU 0 Valid values: `disable`, `enable`.
func (o IpsGlobalOutput) IpsReserveCpu() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsGlobal) pulumi.StringOutput { return v.IpsReserveCpu }).(pulumi.StringOutput)
}

// NGFW policy-mode app detection threshold.
func (o IpsGlobalOutput) NgfwMaxScanRange() pulumi.IntOutput {
	return o.ApplyT(func(v *IpsGlobal) pulumi.IntOutput { return v.NgfwMaxScanRange }).(pulumi.IntOutput)
}

// Acceleration mode for IPS processing by NPx processors. Valid values: `none`, `basic`.
func (o IpsGlobalOutput) NpAccelMode() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsGlobal) pulumi.StringOutput { return v.NpAccelMode }).(pulumi.StringOutput)
}

// Packet/pcap log queue depth per IPS engine.
func (o IpsGlobalOutput) PacketLogQueueDepth() pulumi.IntOutput {
	return o.ApplyT(func(v *IpsGlobal) pulumi.IntOutput { return v.PacketLogQueueDepth }).(pulumi.IntOutput)
}

// Method of counting concurrent sessions used by session limit anomalies. Choose between greater accuracy (accurate) or improved performance (heuristics). Valid values: `accurate`, `heuristic`.
func (o IpsGlobalOutput) SessionLimitMode() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsGlobal) pulumi.StringOutput { return v.SessionLimitMode }).(pulumi.StringOutput)
}

// Public IP addresses of your network that receive Skype sessions. Helps identify Skype sessions. Separate IP addresses with commas.
func (o IpsGlobalOutput) SkypeClientPublicIpaddr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpsGlobal) pulumi.StringPtrOutput { return v.SkypeClientPublicIpaddr }).(pulumi.StringPtrOutput)
}

// IPS socket buffer size (0 - 256 MB). Default depends on available memory. Can be changed to tune performance.
func (o IpsGlobalOutput) SocketSize() pulumi.IntOutput {
	return o.ApplyT(func(v *IpsGlobal) pulumi.IntOutput { return v.SocketSize }).(pulumi.IntOutput)
}

// Enable/disable use of kernel session TTL for IPS sessions. Valid values: `enable`, `disable`.
func (o IpsGlobalOutput) SyncSessionTtl() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsGlobal) pulumi.StringOutput { return v.SyncSessionTtl }).(pulumi.StringOutput)
}

// TLS active probe configuration. The structure of `tlsActiveProbe` block is documented below.
func (o IpsGlobalOutput) TlsActiveProbe() IpsGlobalTlsActiveProbeOutput {
	return o.ApplyT(func(v *IpsGlobal) IpsGlobalTlsActiveProbeOutput { return v.TlsActiveProbe }).(IpsGlobalTlsActiveProbeOutput)
}

// Enable/disable submitting attack data found by this FortiGate to FortiGuard. Valid values: `enable`, `disable`.
func (o IpsGlobalOutput) TrafficSubmit() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsGlobal) pulumi.StringOutput { return v.TrafficSubmit }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o IpsGlobalOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpsGlobal) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type IpsGlobalArrayOutput struct{ *pulumi.OutputState }

func (IpsGlobalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsGlobal)(nil)).Elem()
}

func (o IpsGlobalArrayOutput) ToIpsGlobalArrayOutput() IpsGlobalArrayOutput {
	return o
}

func (o IpsGlobalArrayOutput) ToIpsGlobalArrayOutputWithContext(ctx context.Context) IpsGlobalArrayOutput {
	return o
}

func (o IpsGlobalArrayOutput) Index(i pulumi.IntInput) IpsGlobalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpsGlobal {
		return vs[0].([]*IpsGlobal)[vs[1].(int)]
	}).(IpsGlobalOutput)
}

type IpsGlobalMapOutput struct{ *pulumi.OutputState }

func (IpsGlobalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsGlobal)(nil)).Elem()
}

func (o IpsGlobalMapOutput) ToIpsGlobalMapOutput() IpsGlobalMapOutput {
	return o
}

func (o IpsGlobalMapOutput) ToIpsGlobalMapOutputWithContext(ctx context.Context) IpsGlobalMapOutput {
	return o
}

func (o IpsGlobalMapOutput) MapIndex(k pulumi.StringInput) IpsGlobalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpsGlobal {
		return vs[0].(map[string]*IpsGlobal)[vs[1].(string)]
	}).(IpsGlobalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpsGlobalInput)(nil)).Elem(), &IpsGlobal{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsGlobalArrayInput)(nil)).Elem(), IpsGlobalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsGlobalMapInput)(nil)).Elem(), IpsGlobalMap{})
	pulumi.RegisterOutputType(IpsGlobalOutput{})
	pulumi.RegisterOutputType(IpsGlobalArrayOutput{})
	pulumi.RegisterOutputType(IpsGlobalMapOutput{})
}
