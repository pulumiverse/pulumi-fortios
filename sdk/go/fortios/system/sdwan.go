// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure redundant internet connections using SD-WAN (formerly virtual WAN link). Applies to FortiOS Version `>= 6.4.1`.
//
// ## Import
//
// System Sdwan can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/sdwan:Sdwan labelname SystemSdwan
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/sdwan:Sdwan labelname SystemSdwan
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Sdwan struct {
	pulumi.CustomResourceState

	// Time interval in seconds that applicationperformance logs are generated (0 - 3600, default = 0).
	AppPerfLogPeriod pulumi.IntOutput `pulumi:"appPerfLogPeriod"`
	// Maximum number of interface members a packet is duplicated in the SD-WAN zone (2 - 4, default = 2; if set to 3, the original packet plus 2 more copies are created).
	DuplicationMaxNum pulumi.IntOutput `pulumi:"duplicationMaxNum"`
	// Create SD-WAN duplication rule. The structure of `duplication` block is documented below.
	Duplications SdwanDuplicationArrayOutput `pulumi:"duplications"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
	FailAlertInterfaces SdwanFailAlertInterfaceArrayOutput `pulumi:"failAlertInterfaces"`
	// Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
	FailDetect pulumi.StringOutput `pulumi:"failDetect"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `healthCheck` block is documented below.
	HealthChecks SdwanHealthCheckArrayOutput `pulumi:"healthChecks"`
	// Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
	LoadBalanceMode pulumi.StringOutput `pulumi:"loadBalanceMode"`
	// FortiGate interfaces added to the SD-WAN. The structure of `members` block is documented below.
	Members SdwanMemberArrayOutput `pulumi:"members"`
	// Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
	NeighborHoldBootTime pulumi.IntOutput `pulumi:"neighborHoldBootTime"`
	// Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
	NeighborHoldDown pulumi.StringOutput `pulumi:"neighborHoldDown"`
	// Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
	NeighborHoldDownTime pulumi.IntOutput `pulumi:"neighborHoldDownTime"`
	// Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
	Neighbors SdwanNeighborArrayOutput `pulumi:"neighbors"`
	// Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN. The structure of `service` block is documented below.
	Services SdwanServiceArrayOutput `pulumi:"services"`
	// Enable/disable bypass routing when speedtest on a SD-WAN member. Valid values: `disable`, `enable`.
	SpeedtestBypassRouting pulumi.StringOutput `pulumi:"speedtestBypassRouting"`
	// Enable/disable SD-WAN. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Configure SD-WAN zones. The structure of `zone` block is documented below.
	Zones SdwanZoneArrayOutput `pulumi:"zones"`
}

// NewSdwan registers a new resource with the given unique name, arguments, and options.
func NewSdwan(ctx *pulumi.Context,
	name string, args *SdwanArgs, opts ...pulumi.ResourceOption) (*Sdwan, error) {
	if args == nil {
		args = &SdwanArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Sdwan
	err := ctx.RegisterResource("fortios:system/sdwan:Sdwan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSdwan gets an existing Sdwan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSdwan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SdwanState, opts ...pulumi.ResourceOption) (*Sdwan, error) {
	var resource Sdwan
	err := ctx.ReadResource("fortios:system/sdwan:Sdwan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Sdwan resources.
type sdwanState struct {
	// Time interval in seconds that applicationperformance logs are generated (0 - 3600, default = 0).
	AppPerfLogPeriod *int `pulumi:"appPerfLogPeriod"`
	// Maximum number of interface members a packet is duplicated in the SD-WAN zone (2 - 4, default = 2; if set to 3, the original packet plus 2 more copies are created).
	DuplicationMaxNum *int `pulumi:"duplicationMaxNum"`
	// Create SD-WAN duplication rule. The structure of `duplication` block is documented below.
	Duplications []SdwanDuplication `pulumi:"duplications"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
	FailAlertInterfaces []SdwanFailAlertInterface `pulumi:"failAlertInterfaces"`
	// Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
	FailDetect *string `pulumi:"failDetect"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `healthCheck` block is documented below.
	HealthChecks []SdwanHealthCheck `pulumi:"healthChecks"`
	// Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
	LoadBalanceMode *string `pulumi:"loadBalanceMode"`
	// FortiGate interfaces added to the SD-WAN. The structure of `members` block is documented below.
	Members []SdwanMember `pulumi:"members"`
	// Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
	NeighborHoldBootTime *int `pulumi:"neighborHoldBootTime"`
	// Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
	NeighborHoldDown *string `pulumi:"neighborHoldDown"`
	// Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
	NeighborHoldDownTime *int `pulumi:"neighborHoldDownTime"`
	// Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
	Neighbors []SdwanNeighbor `pulumi:"neighbors"`
	// Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN. The structure of `service` block is documented below.
	Services []SdwanService `pulumi:"services"`
	// Enable/disable bypass routing when speedtest on a SD-WAN member. Valid values: `disable`, `enable`.
	SpeedtestBypassRouting *string `pulumi:"speedtestBypassRouting"`
	// Enable/disable SD-WAN. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Configure SD-WAN zones. The structure of `zone` block is documented below.
	Zones []SdwanZone `pulumi:"zones"`
}

type SdwanState struct {
	// Time interval in seconds that applicationperformance logs are generated (0 - 3600, default = 0).
	AppPerfLogPeriod pulumi.IntPtrInput
	// Maximum number of interface members a packet is duplicated in the SD-WAN zone (2 - 4, default = 2; if set to 3, the original packet plus 2 more copies are created).
	DuplicationMaxNum pulumi.IntPtrInput
	// Create SD-WAN duplication rule. The structure of `duplication` block is documented below.
	Duplications SdwanDuplicationArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
	FailAlertInterfaces SdwanFailAlertInterfaceArrayInput
	// Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
	FailDetect pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `healthCheck` block is documented below.
	HealthChecks SdwanHealthCheckArrayInput
	// Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
	LoadBalanceMode pulumi.StringPtrInput
	// FortiGate interfaces added to the SD-WAN. The structure of `members` block is documented below.
	Members SdwanMemberArrayInput
	// Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
	NeighborHoldBootTime pulumi.IntPtrInput
	// Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
	NeighborHoldDown pulumi.StringPtrInput
	// Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
	NeighborHoldDownTime pulumi.IntPtrInput
	// Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
	Neighbors SdwanNeighborArrayInput
	// Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN. The structure of `service` block is documented below.
	Services SdwanServiceArrayInput
	// Enable/disable bypass routing when speedtest on a SD-WAN member. Valid values: `disable`, `enable`.
	SpeedtestBypassRouting pulumi.StringPtrInput
	// Enable/disable SD-WAN. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Configure SD-WAN zones. The structure of `zone` block is documented below.
	Zones SdwanZoneArrayInput
}

func (SdwanState) ElementType() reflect.Type {
	return reflect.TypeOf((*sdwanState)(nil)).Elem()
}

type sdwanArgs struct {
	// Time interval in seconds that applicationperformance logs are generated (0 - 3600, default = 0).
	AppPerfLogPeriod *int `pulumi:"appPerfLogPeriod"`
	// Maximum number of interface members a packet is duplicated in the SD-WAN zone (2 - 4, default = 2; if set to 3, the original packet plus 2 more copies are created).
	DuplicationMaxNum *int `pulumi:"duplicationMaxNum"`
	// Create SD-WAN duplication rule. The structure of `duplication` block is documented below.
	Duplications []SdwanDuplication `pulumi:"duplications"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
	FailAlertInterfaces []SdwanFailAlertInterface `pulumi:"failAlertInterfaces"`
	// Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
	FailDetect *string `pulumi:"failDetect"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `healthCheck` block is documented below.
	HealthChecks []SdwanHealthCheck `pulumi:"healthChecks"`
	// Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
	LoadBalanceMode *string `pulumi:"loadBalanceMode"`
	// FortiGate interfaces added to the SD-WAN. The structure of `members` block is documented below.
	Members []SdwanMember `pulumi:"members"`
	// Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
	NeighborHoldBootTime *int `pulumi:"neighborHoldBootTime"`
	// Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
	NeighborHoldDown *string `pulumi:"neighborHoldDown"`
	// Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
	NeighborHoldDownTime *int `pulumi:"neighborHoldDownTime"`
	// Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
	Neighbors []SdwanNeighbor `pulumi:"neighbors"`
	// Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN. The structure of `service` block is documented below.
	Services []SdwanService `pulumi:"services"`
	// Enable/disable bypass routing when speedtest on a SD-WAN member. Valid values: `disable`, `enable`.
	SpeedtestBypassRouting *string `pulumi:"speedtestBypassRouting"`
	// Enable/disable SD-WAN. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Configure SD-WAN zones. The structure of `zone` block is documented below.
	Zones []SdwanZone `pulumi:"zones"`
}

// The set of arguments for constructing a Sdwan resource.
type SdwanArgs struct {
	// Time interval in seconds that applicationperformance logs are generated (0 - 3600, default = 0).
	AppPerfLogPeriod pulumi.IntPtrInput
	// Maximum number of interface members a packet is duplicated in the SD-WAN zone (2 - 4, default = 2; if set to 3, the original packet plus 2 more copies are created).
	DuplicationMaxNum pulumi.IntPtrInput
	// Create SD-WAN duplication rule. The structure of `duplication` block is documented below.
	Duplications SdwanDuplicationArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
	FailAlertInterfaces SdwanFailAlertInterfaceArrayInput
	// Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
	FailDetect pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `healthCheck` block is documented below.
	HealthChecks SdwanHealthCheckArrayInput
	// Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
	LoadBalanceMode pulumi.StringPtrInput
	// FortiGate interfaces added to the SD-WAN. The structure of `members` block is documented below.
	Members SdwanMemberArrayInput
	// Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
	NeighborHoldBootTime pulumi.IntPtrInput
	// Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
	NeighborHoldDown pulumi.StringPtrInput
	// Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
	NeighborHoldDownTime pulumi.IntPtrInput
	// Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
	Neighbors SdwanNeighborArrayInput
	// Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN. The structure of `service` block is documented below.
	Services SdwanServiceArrayInput
	// Enable/disable bypass routing when speedtest on a SD-WAN member. Valid values: `disable`, `enable`.
	SpeedtestBypassRouting pulumi.StringPtrInput
	// Enable/disable SD-WAN. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Configure SD-WAN zones. The structure of `zone` block is documented below.
	Zones SdwanZoneArrayInput
}

func (SdwanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sdwanArgs)(nil)).Elem()
}

type SdwanInput interface {
	pulumi.Input

	ToSdwanOutput() SdwanOutput
	ToSdwanOutputWithContext(ctx context.Context) SdwanOutput
}

func (*Sdwan) ElementType() reflect.Type {
	return reflect.TypeOf((**Sdwan)(nil)).Elem()
}

func (i *Sdwan) ToSdwanOutput() SdwanOutput {
	return i.ToSdwanOutputWithContext(context.Background())
}

func (i *Sdwan) ToSdwanOutputWithContext(ctx context.Context) SdwanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SdwanOutput)
}

// SdwanArrayInput is an input type that accepts SdwanArray and SdwanArrayOutput values.
// You can construct a concrete instance of `SdwanArrayInput` via:
//
//	SdwanArray{ SdwanArgs{...} }
type SdwanArrayInput interface {
	pulumi.Input

	ToSdwanArrayOutput() SdwanArrayOutput
	ToSdwanArrayOutputWithContext(context.Context) SdwanArrayOutput
}

type SdwanArray []SdwanInput

func (SdwanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sdwan)(nil)).Elem()
}

func (i SdwanArray) ToSdwanArrayOutput() SdwanArrayOutput {
	return i.ToSdwanArrayOutputWithContext(context.Background())
}

func (i SdwanArray) ToSdwanArrayOutputWithContext(ctx context.Context) SdwanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SdwanArrayOutput)
}

// SdwanMapInput is an input type that accepts SdwanMap and SdwanMapOutput values.
// You can construct a concrete instance of `SdwanMapInput` via:
//
//	SdwanMap{ "key": SdwanArgs{...} }
type SdwanMapInput interface {
	pulumi.Input

	ToSdwanMapOutput() SdwanMapOutput
	ToSdwanMapOutputWithContext(context.Context) SdwanMapOutput
}

type SdwanMap map[string]SdwanInput

func (SdwanMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sdwan)(nil)).Elem()
}

func (i SdwanMap) ToSdwanMapOutput() SdwanMapOutput {
	return i.ToSdwanMapOutputWithContext(context.Background())
}

func (i SdwanMap) ToSdwanMapOutputWithContext(ctx context.Context) SdwanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SdwanMapOutput)
}

type SdwanOutput struct{ *pulumi.OutputState }

func (SdwanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sdwan)(nil)).Elem()
}

func (o SdwanOutput) ToSdwanOutput() SdwanOutput {
	return o
}

func (o SdwanOutput) ToSdwanOutputWithContext(ctx context.Context) SdwanOutput {
	return o
}

// Time interval in seconds that applicationperformance logs are generated (0 - 3600, default = 0).
func (o SdwanOutput) AppPerfLogPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *Sdwan) pulumi.IntOutput { return v.AppPerfLogPeriod }).(pulumi.IntOutput)
}

// Maximum number of interface members a packet is duplicated in the SD-WAN zone (2 - 4, default = 2; if set to 3, the original packet plus 2 more copies are created).
func (o SdwanOutput) DuplicationMaxNum() pulumi.IntOutput {
	return o.ApplyT(func(v *Sdwan) pulumi.IntOutput { return v.DuplicationMaxNum }).(pulumi.IntOutput)
}

// Create SD-WAN duplication rule. The structure of `duplication` block is documented below.
func (o SdwanOutput) Duplications() SdwanDuplicationArrayOutput {
	return o.ApplyT(func(v *Sdwan) SdwanDuplicationArrayOutput { return v.Duplications }).(SdwanDuplicationArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SdwanOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sdwan) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
func (o SdwanOutput) FailAlertInterfaces() SdwanFailAlertInterfaceArrayOutput {
	return o.ApplyT(func(v *Sdwan) SdwanFailAlertInterfaceArrayOutput { return v.FailAlertInterfaces }).(SdwanFailAlertInterfaceArrayOutput)
}

// Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
func (o SdwanOutput) FailDetect() pulumi.StringOutput {
	return o.ApplyT(func(v *Sdwan) pulumi.StringOutput { return v.FailDetect }).(pulumi.StringOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o SdwanOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sdwan) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `healthCheck` block is documented below.
func (o SdwanOutput) HealthChecks() SdwanHealthCheckArrayOutput {
	return o.ApplyT(func(v *Sdwan) SdwanHealthCheckArrayOutput { return v.HealthChecks }).(SdwanHealthCheckArrayOutput)
}

// Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
func (o SdwanOutput) LoadBalanceMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Sdwan) pulumi.StringOutput { return v.LoadBalanceMode }).(pulumi.StringOutput)
}

// FortiGate interfaces added to the SD-WAN. The structure of `members` block is documented below.
func (o SdwanOutput) Members() SdwanMemberArrayOutput {
	return o.ApplyT(func(v *Sdwan) SdwanMemberArrayOutput { return v.Members }).(SdwanMemberArrayOutput)
}

// Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
func (o SdwanOutput) NeighborHoldBootTime() pulumi.IntOutput {
	return o.ApplyT(func(v *Sdwan) pulumi.IntOutput { return v.NeighborHoldBootTime }).(pulumi.IntOutput)
}

// Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
func (o SdwanOutput) NeighborHoldDown() pulumi.StringOutput {
	return o.ApplyT(func(v *Sdwan) pulumi.StringOutput { return v.NeighborHoldDown }).(pulumi.StringOutput)
}

// Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
func (o SdwanOutput) NeighborHoldDownTime() pulumi.IntOutput {
	return o.ApplyT(func(v *Sdwan) pulumi.IntOutput { return v.NeighborHoldDownTime }).(pulumi.IntOutput)
}

// Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
func (o SdwanOutput) Neighbors() SdwanNeighborArrayOutput {
	return o.ApplyT(func(v *Sdwan) SdwanNeighborArrayOutput { return v.Neighbors }).(SdwanNeighborArrayOutput)
}

// Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN. The structure of `service` block is documented below.
func (o SdwanOutput) Services() SdwanServiceArrayOutput {
	return o.ApplyT(func(v *Sdwan) SdwanServiceArrayOutput { return v.Services }).(SdwanServiceArrayOutput)
}

// Enable/disable bypass routing when speedtest on a SD-WAN member. Valid values: `disable`, `enable`.
func (o SdwanOutput) SpeedtestBypassRouting() pulumi.StringOutput {
	return o.ApplyT(func(v *Sdwan) pulumi.StringOutput { return v.SpeedtestBypassRouting }).(pulumi.StringOutput)
}

// Enable/disable SD-WAN. Valid values: `disable`, `enable`.
func (o SdwanOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Sdwan) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SdwanOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sdwan) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Configure SD-WAN zones. The structure of `zone` block is documented below.
func (o SdwanOutput) Zones() SdwanZoneArrayOutput {
	return o.ApplyT(func(v *Sdwan) SdwanZoneArrayOutput { return v.Zones }).(SdwanZoneArrayOutput)
}

type SdwanArrayOutput struct{ *pulumi.OutputState }

func (SdwanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sdwan)(nil)).Elem()
}

func (o SdwanArrayOutput) ToSdwanArrayOutput() SdwanArrayOutput {
	return o
}

func (o SdwanArrayOutput) ToSdwanArrayOutputWithContext(ctx context.Context) SdwanArrayOutput {
	return o
}

func (o SdwanArrayOutput) Index(i pulumi.IntInput) SdwanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Sdwan {
		return vs[0].([]*Sdwan)[vs[1].(int)]
	}).(SdwanOutput)
}

type SdwanMapOutput struct{ *pulumi.OutputState }

func (SdwanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sdwan)(nil)).Elem()
}

func (o SdwanMapOutput) ToSdwanMapOutput() SdwanMapOutput {
	return o
}

func (o SdwanMapOutput) ToSdwanMapOutputWithContext(ctx context.Context) SdwanMapOutput {
	return o
}

func (o SdwanMapOutput) MapIndex(k pulumi.StringInput) SdwanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Sdwan {
		return vs[0].(map[string]*Sdwan)[vs[1].(string)]
	}).(SdwanOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SdwanInput)(nil)).Elem(), &Sdwan{})
	pulumi.RegisterInputType(reflect.TypeOf((*SdwanArrayInput)(nil)).Elem(), SdwanArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SdwanMapInput)(nil)).Elem(), SdwanMap{})
	pulumi.RegisterOutputType(SdwanOutput{})
	pulumi.RegisterOutputType(SdwanArrayOutput{})
	pulumi.RegisterOutputType(SdwanMapOutput{})
}
