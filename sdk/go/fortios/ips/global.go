// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ips

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure IPS global parameter.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/ips"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ips.NewGlobal(ctx, "trname", &ips.GlobalArgs{
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Ips Global can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:ips/global:Global labelname IpsGlobal
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:ips/global:Global labelname IpsGlobal
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Global struct {
	pulumi.CustomResourceState

	// Global blocking mode for rate-based anomalies. Valid values: `periodical`, `continuous`.
	AnomalyMode pulumi.StringOutput `pulumi:"anomalyMode"`
	// Maximum percentage of system memory allowed for use on AV scanning (10 - 50, default = zero). To disable set to zero. When disabled, there is no limit on the AV memory usage.
	AvMemLimit pulumi.IntOutput `pulumi:"avMemLimit"`
	// IPS Pattern matching acceleration/offloading to CPx processors. Valid values: `none`, `basic`, `advanced`.
	CpAccelMode pulumi.StringOutput `pulumi:"cpAccelMode"`
	// Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks. Valid values: `regular`, `extended`.
	Database pulumi.StringOutput `pulumi:"database"`
	// Limit on number of entries in deep application inspection database (1 - 2147483647, use recommended setting = 0).
	DeepAppInspDbLimit pulumi.IntOutput `pulumi:"deepAppInspDbLimit"`
	// Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
	DeepAppInspTimeout pulumi.IntOutput `pulumi:"deepAppInspTimeout"`
	// Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
	EngineCount pulumi.IntOutput `pulumi:"engineCount"`
	// Excluded signatures.
	ExcludeSignatures pulumi.StringOutput `pulumi:"excludeSignatures"`
	// Enable to allow traffic if the IPS process crashes. Default is disable and IPS traffic is blocked when the IPS process crashes. Valid values: `enable`, `disable`.
	FailOpen pulumi.StringOutput `pulumi:"failOpen"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
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
	// IPS socket buffer size. Max and default value depend on available memory. Can be changed to tune performance.
	SocketSize pulumi.IntOutput `pulumi:"socketSize"`
	// Enable/disable use of kernel session TTL for IPS sessions. Valid values: `enable`, `disable`.
	SyncSessionTtl pulumi.StringOutput `pulumi:"syncSessionTtl"`
	// TLS active probe configuration. The structure of `tlsActiveProbe` block is documented below.
	TlsActiveProbe GlobalTlsActiveProbeOutput `pulumi:"tlsActiveProbe"`
	// Enable/disable submitting attack data found by this FortiGate to FortiGuard. Valid values: `enable`, `disable`.
	TrafficSubmit pulumi.StringOutput `pulumi:"trafficSubmit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewGlobal registers a new resource with the given unique name, arguments, and options.
func NewGlobal(ctx *pulumi.Context,
	name string, args *GlobalArgs, opts ...pulumi.ResourceOption) (*Global, error) {
	if args == nil {
		args = &GlobalArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Global
	err := ctx.RegisterResource("fortios:ips/global:Global", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobal gets an existing Global resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalState, opts ...pulumi.ResourceOption) (*Global, error) {
	var resource Global
	err := ctx.ReadResource("fortios:ips/global:Global", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Global resources.
type globalState struct {
	// Global blocking mode for rate-based anomalies. Valid values: `periodical`, `continuous`.
	AnomalyMode *string `pulumi:"anomalyMode"`
	// Maximum percentage of system memory allowed for use on AV scanning (10 - 50, default = zero). To disable set to zero. When disabled, there is no limit on the AV memory usage.
	AvMemLimit *int `pulumi:"avMemLimit"`
	// IPS Pattern matching acceleration/offloading to CPx processors. Valid values: `none`, `basic`, `advanced`.
	CpAccelMode *string `pulumi:"cpAccelMode"`
	// Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks. Valid values: `regular`, `extended`.
	Database *string `pulumi:"database"`
	// Limit on number of entries in deep application inspection database (1 - 2147483647, use recommended setting = 0).
	DeepAppInspDbLimit *int `pulumi:"deepAppInspDbLimit"`
	// Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
	DeepAppInspTimeout *int `pulumi:"deepAppInspTimeout"`
	// Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
	EngineCount *int `pulumi:"engineCount"`
	// Excluded signatures.
	ExcludeSignatures *string `pulumi:"excludeSignatures"`
	// Enable to allow traffic if the IPS process crashes. Default is disable and IPS traffic is blocked when the IPS process crashes. Valid values: `enable`, `disable`.
	FailOpen *string `pulumi:"failOpen"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
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
	// IPS socket buffer size. Max and default value depend on available memory. Can be changed to tune performance.
	SocketSize *int `pulumi:"socketSize"`
	// Enable/disable use of kernel session TTL for IPS sessions. Valid values: `enable`, `disable`.
	SyncSessionTtl *string `pulumi:"syncSessionTtl"`
	// TLS active probe configuration. The structure of `tlsActiveProbe` block is documented below.
	TlsActiveProbe *GlobalTlsActiveProbe `pulumi:"tlsActiveProbe"`
	// Enable/disable submitting attack data found by this FortiGate to FortiGuard. Valid values: `enable`, `disable`.
	TrafficSubmit *string `pulumi:"trafficSubmit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type GlobalState struct {
	// Global blocking mode for rate-based anomalies. Valid values: `periodical`, `continuous`.
	AnomalyMode pulumi.StringPtrInput
	// Maximum percentage of system memory allowed for use on AV scanning (10 - 50, default = zero). To disable set to zero. When disabled, there is no limit on the AV memory usage.
	AvMemLimit pulumi.IntPtrInput
	// IPS Pattern matching acceleration/offloading to CPx processors. Valid values: `none`, `basic`, `advanced`.
	CpAccelMode pulumi.StringPtrInput
	// Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks. Valid values: `regular`, `extended`.
	Database pulumi.StringPtrInput
	// Limit on number of entries in deep application inspection database (1 - 2147483647, use recommended setting = 0).
	DeepAppInspDbLimit pulumi.IntPtrInput
	// Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
	DeepAppInspTimeout pulumi.IntPtrInput
	// Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
	EngineCount pulumi.IntPtrInput
	// Excluded signatures.
	ExcludeSignatures pulumi.StringPtrInput
	// Enable to allow traffic if the IPS process crashes. Default is disable and IPS traffic is blocked when the IPS process crashes. Valid values: `enable`, `disable`.
	FailOpen pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
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
	// IPS socket buffer size. Max and default value depend on available memory. Can be changed to tune performance.
	SocketSize pulumi.IntPtrInput
	// Enable/disable use of kernel session TTL for IPS sessions. Valid values: `enable`, `disable`.
	SyncSessionTtl pulumi.StringPtrInput
	// TLS active probe configuration. The structure of `tlsActiveProbe` block is documented below.
	TlsActiveProbe GlobalTlsActiveProbePtrInput
	// Enable/disable submitting attack data found by this FortiGate to FortiGuard. Valid values: `enable`, `disable`.
	TrafficSubmit pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (GlobalState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalState)(nil)).Elem()
}

type globalArgs struct {
	// Global blocking mode for rate-based anomalies. Valid values: `periodical`, `continuous`.
	AnomalyMode *string `pulumi:"anomalyMode"`
	// Maximum percentage of system memory allowed for use on AV scanning (10 - 50, default = zero). To disable set to zero. When disabled, there is no limit on the AV memory usage.
	AvMemLimit *int `pulumi:"avMemLimit"`
	// IPS Pattern matching acceleration/offloading to CPx processors. Valid values: `none`, `basic`, `advanced`.
	CpAccelMode *string `pulumi:"cpAccelMode"`
	// Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks. Valid values: `regular`, `extended`.
	Database *string `pulumi:"database"`
	// Limit on number of entries in deep application inspection database (1 - 2147483647, use recommended setting = 0).
	DeepAppInspDbLimit *int `pulumi:"deepAppInspDbLimit"`
	// Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
	DeepAppInspTimeout *int `pulumi:"deepAppInspTimeout"`
	// Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
	EngineCount *int `pulumi:"engineCount"`
	// Excluded signatures.
	ExcludeSignatures *string `pulumi:"excludeSignatures"`
	// Enable to allow traffic if the IPS process crashes. Default is disable and IPS traffic is blocked when the IPS process crashes. Valid values: `enable`, `disable`.
	FailOpen *string `pulumi:"failOpen"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
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
	// IPS socket buffer size. Max and default value depend on available memory. Can be changed to tune performance.
	SocketSize *int `pulumi:"socketSize"`
	// Enable/disable use of kernel session TTL for IPS sessions. Valid values: `enable`, `disable`.
	SyncSessionTtl *string `pulumi:"syncSessionTtl"`
	// TLS active probe configuration. The structure of `tlsActiveProbe` block is documented below.
	TlsActiveProbe *GlobalTlsActiveProbe `pulumi:"tlsActiveProbe"`
	// Enable/disable submitting attack data found by this FortiGate to FortiGuard. Valid values: `enable`, `disable`.
	TrafficSubmit *string `pulumi:"trafficSubmit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Global resource.
type GlobalArgs struct {
	// Global blocking mode for rate-based anomalies. Valid values: `periodical`, `continuous`.
	AnomalyMode pulumi.StringPtrInput
	// Maximum percentage of system memory allowed for use on AV scanning (10 - 50, default = zero). To disable set to zero. When disabled, there is no limit on the AV memory usage.
	AvMemLimit pulumi.IntPtrInput
	// IPS Pattern matching acceleration/offloading to CPx processors. Valid values: `none`, `basic`, `advanced`.
	CpAccelMode pulumi.StringPtrInput
	// Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks. Valid values: `regular`, `extended`.
	Database pulumi.StringPtrInput
	// Limit on number of entries in deep application inspection database (1 - 2147483647, use recommended setting = 0).
	DeepAppInspDbLimit pulumi.IntPtrInput
	// Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
	DeepAppInspTimeout pulumi.IntPtrInput
	// Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
	EngineCount pulumi.IntPtrInput
	// Excluded signatures.
	ExcludeSignatures pulumi.StringPtrInput
	// Enable to allow traffic if the IPS process crashes. Default is disable and IPS traffic is blocked when the IPS process crashes. Valid values: `enable`, `disable`.
	FailOpen pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
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
	// IPS socket buffer size. Max and default value depend on available memory. Can be changed to tune performance.
	SocketSize pulumi.IntPtrInput
	// Enable/disable use of kernel session TTL for IPS sessions. Valid values: `enable`, `disable`.
	SyncSessionTtl pulumi.StringPtrInput
	// TLS active probe configuration. The structure of `tlsActiveProbe` block is documented below.
	TlsActiveProbe GlobalTlsActiveProbePtrInput
	// Enable/disable submitting attack data found by this FortiGate to FortiGuard. Valid values: `enable`, `disable`.
	TrafficSubmit pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (GlobalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalArgs)(nil)).Elem()
}

type GlobalInput interface {
	pulumi.Input

	ToGlobalOutput() GlobalOutput
	ToGlobalOutputWithContext(ctx context.Context) GlobalOutput
}

func (*Global) ElementType() reflect.Type {
	return reflect.TypeOf((**Global)(nil)).Elem()
}

func (i *Global) ToGlobalOutput() GlobalOutput {
	return i.ToGlobalOutputWithContext(context.Background())
}

func (i *Global) ToGlobalOutputWithContext(ctx context.Context) GlobalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalOutput)
}

// GlobalArrayInput is an input type that accepts GlobalArray and GlobalArrayOutput values.
// You can construct a concrete instance of `GlobalArrayInput` via:
//
//	GlobalArray{ GlobalArgs{...} }
type GlobalArrayInput interface {
	pulumi.Input

	ToGlobalArrayOutput() GlobalArrayOutput
	ToGlobalArrayOutputWithContext(context.Context) GlobalArrayOutput
}

type GlobalArray []GlobalInput

func (GlobalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Global)(nil)).Elem()
}

func (i GlobalArray) ToGlobalArrayOutput() GlobalArrayOutput {
	return i.ToGlobalArrayOutputWithContext(context.Background())
}

func (i GlobalArray) ToGlobalArrayOutputWithContext(ctx context.Context) GlobalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalArrayOutput)
}

// GlobalMapInput is an input type that accepts GlobalMap and GlobalMapOutput values.
// You can construct a concrete instance of `GlobalMapInput` via:
//
//	GlobalMap{ "key": GlobalArgs{...} }
type GlobalMapInput interface {
	pulumi.Input

	ToGlobalMapOutput() GlobalMapOutput
	ToGlobalMapOutputWithContext(context.Context) GlobalMapOutput
}

type GlobalMap map[string]GlobalInput

func (GlobalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Global)(nil)).Elem()
}

func (i GlobalMap) ToGlobalMapOutput() GlobalMapOutput {
	return i.ToGlobalMapOutputWithContext(context.Background())
}

func (i GlobalMap) ToGlobalMapOutputWithContext(ctx context.Context) GlobalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalMapOutput)
}

type GlobalOutput struct{ *pulumi.OutputState }

func (GlobalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Global)(nil)).Elem()
}

func (o GlobalOutput) ToGlobalOutput() GlobalOutput {
	return o
}

func (o GlobalOutput) ToGlobalOutputWithContext(ctx context.Context) GlobalOutput {
	return o
}

// Global blocking mode for rate-based anomalies. Valid values: `periodical`, `continuous`.
func (o GlobalOutput) AnomalyMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.AnomalyMode }).(pulumi.StringOutput)
}

// Maximum percentage of system memory allowed for use on AV scanning (10 - 50, default = zero). To disable set to zero. When disabled, there is no limit on the AV memory usage.
func (o GlobalOutput) AvMemLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *Global) pulumi.IntOutput { return v.AvMemLimit }).(pulumi.IntOutput)
}

// IPS Pattern matching acceleration/offloading to CPx processors. Valid values: `none`, `basic`, `advanced`.
func (o GlobalOutput) CpAccelMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.CpAccelMode }).(pulumi.StringOutput)
}

// Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks. Valid values: `regular`, `extended`.
func (o GlobalOutput) Database() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.Database }).(pulumi.StringOutput)
}

// Limit on number of entries in deep application inspection database (1 - 2147483647, use recommended setting = 0).
func (o GlobalOutput) DeepAppInspDbLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *Global) pulumi.IntOutput { return v.DeepAppInspDbLimit }).(pulumi.IntOutput)
}

// Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
func (o GlobalOutput) DeepAppInspTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Global) pulumi.IntOutput { return v.DeepAppInspTimeout }).(pulumi.IntOutput)
}

// Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
func (o GlobalOutput) EngineCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Global) pulumi.IntOutput { return v.EngineCount }).(pulumi.IntOutput)
}

// Excluded signatures.
func (o GlobalOutput) ExcludeSignatures() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.ExcludeSignatures }).(pulumi.StringOutput)
}

// Enable to allow traffic if the IPS process crashes. Default is disable and IPS traffic is blocked when the IPS process crashes. Valid values: `enable`, `disable`.
func (o GlobalOutput) FailOpen() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.FailOpen }).(pulumi.StringOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o GlobalOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Global) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Enable/disable IPS adaptive scanning (intelligent mode). Intelligent mode optimizes the scanning method for the type of traffic. Valid values: `enable`, `disable`.
func (o GlobalOutput) IntelligentMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.IntelligentMode }).(pulumi.StringOutput)
}

// Enable/disable IPS daemon's use of CPUs other than CPU 0 Valid values: `disable`, `enable`.
func (o GlobalOutput) IpsReserveCpu() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.IpsReserveCpu }).(pulumi.StringOutput)
}

// NGFW policy-mode app detection threshold.
func (o GlobalOutput) NgfwMaxScanRange() pulumi.IntOutput {
	return o.ApplyT(func(v *Global) pulumi.IntOutput { return v.NgfwMaxScanRange }).(pulumi.IntOutput)
}

// Acceleration mode for IPS processing by NPx processors. Valid values: `none`, `basic`.
func (o GlobalOutput) NpAccelMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.NpAccelMode }).(pulumi.StringOutput)
}

// Packet/pcap log queue depth per IPS engine.
func (o GlobalOutput) PacketLogQueueDepth() pulumi.IntOutput {
	return o.ApplyT(func(v *Global) pulumi.IntOutput { return v.PacketLogQueueDepth }).(pulumi.IntOutput)
}

// Method of counting concurrent sessions used by session limit anomalies. Choose between greater accuracy (accurate) or improved performance (heuristics). Valid values: `accurate`, `heuristic`.
func (o GlobalOutput) SessionLimitMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.SessionLimitMode }).(pulumi.StringOutput)
}

// Public IP addresses of your network that receive Skype sessions. Helps identify Skype sessions. Separate IP addresses with commas.
func (o GlobalOutput) SkypeClientPublicIpaddr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Global) pulumi.StringPtrOutput { return v.SkypeClientPublicIpaddr }).(pulumi.StringPtrOutput)
}

// IPS socket buffer size. Max and default value depend on available memory. Can be changed to tune performance.
func (o GlobalOutput) SocketSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Global) pulumi.IntOutput { return v.SocketSize }).(pulumi.IntOutput)
}

// Enable/disable use of kernel session TTL for IPS sessions. Valid values: `enable`, `disable`.
func (o GlobalOutput) SyncSessionTtl() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.SyncSessionTtl }).(pulumi.StringOutput)
}

// TLS active probe configuration. The structure of `tlsActiveProbe` block is documented below.
func (o GlobalOutput) TlsActiveProbe() GlobalTlsActiveProbeOutput {
	return o.ApplyT(func(v *Global) GlobalTlsActiveProbeOutput { return v.TlsActiveProbe }).(GlobalTlsActiveProbeOutput)
}

// Enable/disable submitting attack data found by this FortiGate to FortiGuard. Valid values: `enable`, `disable`.
func (o GlobalOutput) TrafficSubmit() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.TrafficSubmit }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o GlobalOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Global) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type GlobalArrayOutput struct{ *pulumi.OutputState }

func (GlobalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Global)(nil)).Elem()
}

func (o GlobalArrayOutput) ToGlobalArrayOutput() GlobalArrayOutput {
	return o
}

func (o GlobalArrayOutput) ToGlobalArrayOutputWithContext(ctx context.Context) GlobalArrayOutput {
	return o
}

func (o GlobalArrayOutput) Index(i pulumi.IntInput) GlobalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Global {
		return vs[0].([]*Global)[vs[1].(int)]
	}).(GlobalOutput)
}

type GlobalMapOutput struct{ *pulumi.OutputState }

func (GlobalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Global)(nil)).Elem()
}

func (o GlobalMapOutput) ToGlobalMapOutput() GlobalMapOutput {
	return o
}

func (o GlobalMapOutput) ToGlobalMapOutputWithContext(ctx context.Context) GlobalMapOutput {
	return o
}

func (o GlobalMapOutput) MapIndex(k pulumi.StringInput) GlobalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Global {
		return vs[0].(map[string]*Global)[vs[1].(string)]
	}).(GlobalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalInput)(nil)).Elem(), &Global{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalArrayInput)(nil)).Elem(), GlobalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalMapInput)(nil)).Elem(), GlobalMap{})
	pulumi.RegisterOutputType(GlobalOutput{})
	pulumi.RegisterOutputType(GlobalArrayOutput{})
	pulumi.RegisterOutputType(GlobalMapOutput{})
}
