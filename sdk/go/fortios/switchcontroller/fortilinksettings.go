// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure integrated FortiLink settings for FortiSwitch. Applies to FortiOS Version `>= 7.0.1`.
//
// ## Import
//
// SwitchController FortilinkSettings can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/fortilinksettings:Fortilinksettings labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/fortilinksettings:Fortilinksettings labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Fortilinksettings struct {
	pulumi.CustomResourceState

	// Intra VLAN traffic behavior with loss of connection to the FortiGate. Valid values: `legacy`, `fail-open`, `fail-close`.
	AccessVlanMode pulumi.StringOutput `pulumi:"accessVlanMode"`
	// FortiLink interface to which this fortilink-setting belongs.
	Fortilink pulumi.StringOutput `pulumi:"fortilink"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
	InactiveTimer pulumi.IntOutput `pulumi:"inactiveTimer"`
	// Clear NAC and dynamic devices on switch ports on link down event. Valid values: `disable`, `enable`.
	LinkDownFlush pulumi.StringOutput `pulumi:"linkDownFlush"`
	// NAC specific configuration. The structure of `nacPorts` block is documented below.
	NacPorts FortilinksettingsNacPortsOutput `pulumi:"nacPorts"`
	// FortiLink settings name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewFortilinksettings registers a new resource with the given unique name, arguments, and options.
func NewFortilinksettings(ctx *pulumi.Context,
	name string, args *FortilinksettingsArgs, opts ...pulumi.ResourceOption) (*Fortilinksettings, error) {
	if args == nil {
		args = &FortilinksettingsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Fortilinksettings
	err := ctx.RegisterResource("fortios:switchcontroller/fortilinksettings:Fortilinksettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortilinksettings gets an existing Fortilinksettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortilinksettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortilinksettingsState, opts ...pulumi.ResourceOption) (*Fortilinksettings, error) {
	var resource Fortilinksettings
	err := ctx.ReadResource("fortios:switchcontroller/fortilinksettings:Fortilinksettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fortilinksettings resources.
type fortilinksettingsState struct {
	// Intra VLAN traffic behavior with loss of connection to the FortiGate. Valid values: `legacy`, `fail-open`, `fail-close`.
	AccessVlanMode *string `pulumi:"accessVlanMode"`
	// FortiLink interface to which this fortilink-setting belongs.
	Fortilink *string `pulumi:"fortilink"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
	InactiveTimer *int `pulumi:"inactiveTimer"`
	// Clear NAC and dynamic devices on switch ports on link down event. Valid values: `disable`, `enable`.
	LinkDownFlush *string `pulumi:"linkDownFlush"`
	// NAC specific configuration. The structure of `nacPorts` block is documented below.
	NacPorts *FortilinksettingsNacPorts `pulumi:"nacPorts"`
	// FortiLink settings name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FortilinksettingsState struct {
	// Intra VLAN traffic behavior with loss of connection to the FortiGate. Valid values: `legacy`, `fail-open`, `fail-close`.
	AccessVlanMode pulumi.StringPtrInput
	// FortiLink interface to which this fortilink-setting belongs.
	Fortilink pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
	InactiveTimer pulumi.IntPtrInput
	// Clear NAC and dynamic devices on switch ports on link down event. Valid values: `disable`, `enable`.
	LinkDownFlush pulumi.StringPtrInput
	// NAC specific configuration. The structure of `nacPorts` block is documented below.
	NacPorts FortilinksettingsNacPortsPtrInput
	// FortiLink settings name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FortilinksettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortilinksettingsState)(nil)).Elem()
}

type fortilinksettingsArgs struct {
	// Intra VLAN traffic behavior with loss of connection to the FortiGate. Valid values: `legacy`, `fail-open`, `fail-close`.
	AccessVlanMode *string `pulumi:"accessVlanMode"`
	// FortiLink interface to which this fortilink-setting belongs.
	Fortilink *string `pulumi:"fortilink"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
	InactiveTimer *int `pulumi:"inactiveTimer"`
	// Clear NAC and dynamic devices on switch ports on link down event. Valid values: `disable`, `enable`.
	LinkDownFlush *string `pulumi:"linkDownFlush"`
	// NAC specific configuration. The structure of `nacPorts` block is documented below.
	NacPorts *FortilinksettingsNacPorts `pulumi:"nacPorts"`
	// FortiLink settings name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Fortilinksettings resource.
type FortilinksettingsArgs struct {
	// Intra VLAN traffic behavior with loss of connection to the FortiGate. Valid values: `legacy`, `fail-open`, `fail-close`.
	AccessVlanMode pulumi.StringPtrInput
	// FortiLink interface to which this fortilink-setting belongs.
	Fortilink pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
	InactiveTimer pulumi.IntPtrInput
	// Clear NAC and dynamic devices on switch ports on link down event. Valid values: `disable`, `enable`.
	LinkDownFlush pulumi.StringPtrInput
	// NAC specific configuration. The structure of `nacPorts` block is documented below.
	NacPorts FortilinksettingsNacPortsPtrInput
	// FortiLink settings name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FortilinksettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortilinksettingsArgs)(nil)).Elem()
}

type FortilinksettingsInput interface {
	pulumi.Input

	ToFortilinksettingsOutput() FortilinksettingsOutput
	ToFortilinksettingsOutputWithContext(ctx context.Context) FortilinksettingsOutput
}

func (*Fortilinksettings) ElementType() reflect.Type {
	return reflect.TypeOf((**Fortilinksettings)(nil)).Elem()
}

func (i *Fortilinksettings) ToFortilinksettingsOutput() FortilinksettingsOutput {
	return i.ToFortilinksettingsOutputWithContext(context.Background())
}

func (i *Fortilinksettings) ToFortilinksettingsOutputWithContext(ctx context.Context) FortilinksettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortilinksettingsOutput)
}

// FortilinksettingsArrayInput is an input type that accepts FortilinksettingsArray and FortilinksettingsArrayOutput values.
// You can construct a concrete instance of `FortilinksettingsArrayInput` via:
//
//	FortilinksettingsArray{ FortilinksettingsArgs{...} }
type FortilinksettingsArrayInput interface {
	pulumi.Input

	ToFortilinksettingsArrayOutput() FortilinksettingsArrayOutput
	ToFortilinksettingsArrayOutputWithContext(context.Context) FortilinksettingsArrayOutput
}

type FortilinksettingsArray []FortilinksettingsInput

func (FortilinksettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fortilinksettings)(nil)).Elem()
}

func (i FortilinksettingsArray) ToFortilinksettingsArrayOutput() FortilinksettingsArrayOutput {
	return i.ToFortilinksettingsArrayOutputWithContext(context.Background())
}

func (i FortilinksettingsArray) ToFortilinksettingsArrayOutputWithContext(ctx context.Context) FortilinksettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortilinksettingsArrayOutput)
}

// FortilinksettingsMapInput is an input type that accepts FortilinksettingsMap and FortilinksettingsMapOutput values.
// You can construct a concrete instance of `FortilinksettingsMapInput` via:
//
//	FortilinksettingsMap{ "key": FortilinksettingsArgs{...} }
type FortilinksettingsMapInput interface {
	pulumi.Input

	ToFortilinksettingsMapOutput() FortilinksettingsMapOutput
	ToFortilinksettingsMapOutputWithContext(context.Context) FortilinksettingsMapOutput
}

type FortilinksettingsMap map[string]FortilinksettingsInput

func (FortilinksettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fortilinksettings)(nil)).Elem()
}

func (i FortilinksettingsMap) ToFortilinksettingsMapOutput() FortilinksettingsMapOutput {
	return i.ToFortilinksettingsMapOutputWithContext(context.Background())
}

func (i FortilinksettingsMap) ToFortilinksettingsMapOutputWithContext(ctx context.Context) FortilinksettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortilinksettingsMapOutput)
}

type FortilinksettingsOutput struct{ *pulumi.OutputState }

func (FortilinksettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fortilinksettings)(nil)).Elem()
}

func (o FortilinksettingsOutput) ToFortilinksettingsOutput() FortilinksettingsOutput {
	return o
}

func (o FortilinksettingsOutput) ToFortilinksettingsOutputWithContext(ctx context.Context) FortilinksettingsOutput {
	return o
}

// Intra VLAN traffic behavior with loss of connection to the FortiGate. Valid values: `legacy`, `fail-open`, `fail-close`.
func (o FortilinksettingsOutput) AccessVlanMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortilinksettings) pulumi.StringOutput { return v.AccessVlanMode }).(pulumi.StringOutput)
}

// FortiLink interface to which this fortilink-setting belongs.
func (o FortilinksettingsOutput) Fortilink() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortilinksettings) pulumi.StringOutput { return v.Fortilink }).(pulumi.StringOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o FortilinksettingsOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Fortilinksettings) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
func (o FortilinksettingsOutput) InactiveTimer() pulumi.IntOutput {
	return o.ApplyT(func(v *Fortilinksettings) pulumi.IntOutput { return v.InactiveTimer }).(pulumi.IntOutput)
}

// Clear NAC and dynamic devices on switch ports on link down event. Valid values: `disable`, `enable`.
func (o FortilinksettingsOutput) LinkDownFlush() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortilinksettings) pulumi.StringOutput { return v.LinkDownFlush }).(pulumi.StringOutput)
}

// NAC specific configuration. The structure of `nacPorts` block is documented below.
func (o FortilinksettingsOutput) NacPorts() FortilinksettingsNacPortsOutput {
	return o.ApplyT(func(v *Fortilinksettings) FortilinksettingsNacPortsOutput { return v.NacPorts }).(FortilinksettingsNacPortsOutput)
}

// FortiLink settings name.
func (o FortilinksettingsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortilinksettings) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FortilinksettingsOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortilinksettings) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type FortilinksettingsArrayOutput struct{ *pulumi.OutputState }

func (FortilinksettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fortilinksettings)(nil)).Elem()
}

func (o FortilinksettingsArrayOutput) ToFortilinksettingsArrayOutput() FortilinksettingsArrayOutput {
	return o
}

func (o FortilinksettingsArrayOutput) ToFortilinksettingsArrayOutputWithContext(ctx context.Context) FortilinksettingsArrayOutput {
	return o
}

func (o FortilinksettingsArrayOutput) Index(i pulumi.IntInput) FortilinksettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Fortilinksettings {
		return vs[0].([]*Fortilinksettings)[vs[1].(int)]
	}).(FortilinksettingsOutput)
}

type FortilinksettingsMapOutput struct{ *pulumi.OutputState }

func (FortilinksettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fortilinksettings)(nil)).Elem()
}

func (o FortilinksettingsMapOutput) ToFortilinksettingsMapOutput() FortilinksettingsMapOutput {
	return o
}

func (o FortilinksettingsMapOutput) ToFortilinksettingsMapOutputWithContext(ctx context.Context) FortilinksettingsMapOutput {
	return o
}

func (o FortilinksettingsMapOutput) MapIndex(k pulumi.StringInput) FortilinksettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Fortilinksettings {
		return vs[0].(map[string]*Fortilinksettings)[vs[1].(string)]
	}).(FortilinksettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortilinksettingsInput)(nil)).Elem(), &Fortilinksettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortilinksettingsArrayInput)(nil)).Elem(), FortilinksettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortilinksettingsMapInput)(nil)).Elem(), FortilinksettingsMap{})
	pulumi.RegisterOutputType(FortilinksettingsOutput{})
	pulumi.RegisterOutputType(FortilinksettingsArrayOutput{})
	pulumi.RegisterOutputType(FortilinksettingsMapOutput{})
}
