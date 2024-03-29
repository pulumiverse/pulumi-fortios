// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure FortiSwitch SNMP system information globally. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// SwitchController SnmpSysinfo can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/snmpsysinfo:Snmpsysinfo labelname SwitchControllerSnmpSysinfo
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/snmpsysinfo:Snmpsysinfo labelname SwitchControllerSnmpSysinfo
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Snmpsysinfo struct {
	pulumi.CustomResourceState

	// Contact information.
	ContactInfo pulumi.StringOutput `pulumi:"contactInfo"`
	// System description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Local SNMP engine ID string (max 24 char).
	EngineId pulumi.StringOutput `pulumi:"engineId"`
	// System location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Enable/disable SNMP. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSnmpsysinfo registers a new resource with the given unique name, arguments, and options.
func NewSnmpsysinfo(ctx *pulumi.Context,
	name string, args *SnmpsysinfoArgs, opts ...pulumi.ResourceOption) (*Snmpsysinfo, error) {
	if args == nil {
		args = &SnmpsysinfoArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Snmpsysinfo
	err := ctx.RegisterResource("fortios:switchcontroller/snmpsysinfo:Snmpsysinfo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnmpsysinfo gets an existing Snmpsysinfo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnmpsysinfo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnmpsysinfoState, opts ...pulumi.ResourceOption) (*Snmpsysinfo, error) {
	var resource Snmpsysinfo
	err := ctx.ReadResource("fortios:switchcontroller/snmpsysinfo:Snmpsysinfo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snmpsysinfo resources.
type snmpsysinfoState struct {
	// Contact information.
	ContactInfo *string `pulumi:"contactInfo"`
	// System description.
	Description *string `pulumi:"description"`
	// Local SNMP engine ID string (max 24 char).
	EngineId *string `pulumi:"engineId"`
	// System location.
	Location *string `pulumi:"location"`
	// Enable/disable SNMP. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SnmpsysinfoState struct {
	// Contact information.
	ContactInfo pulumi.StringPtrInput
	// System description.
	Description pulumi.StringPtrInput
	// Local SNMP engine ID string (max 24 char).
	EngineId pulumi.StringPtrInput
	// System location.
	Location pulumi.StringPtrInput
	// Enable/disable SNMP. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SnmpsysinfoState) ElementType() reflect.Type {
	return reflect.TypeOf((*snmpsysinfoState)(nil)).Elem()
}

type snmpsysinfoArgs struct {
	// Contact information.
	ContactInfo *string `pulumi:"contactInfo"`
	// System description.
	Description *string `pulumi:"description"`
	// Local SNMP engine ID string (max 24 char).
	EngineId *string `pulumi:"engineId"`
	// System location.
	Location *string `pulumi:"location"`
	// Enable/disable SNMP. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Snmpsysinfo resource.
type SnmpsysinfoArgs struct {
	// Contact information.
	ContactInfo pulumi.StringPtrInput
	// System description.
	Description pulumi.StringPtrInput
	// Local SNMP engine ID string (max 24 char).
	EngineId pulumi.StringPtrInput
	// System location.
	Location pulumi.StringPtrInput
	// Enable/disable SNMP. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SnmpsysinfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snmpsysinfoArgs)(nil)).Elem()
}

type SnmpsysinfoInput interface {
	pulumi.Input

	ToSnmpsysinfoOutput() SnmpsysinfoOutput
	ToSnmpsysinfoOutputWithContext(ctx context.Context) SnmpsysinfoOutput
}

func (*Snmpsysinfo) ElementType() reflect.Type {
	return reflect.TypeOf((**Snmpsysinfo)(nil)).Elem()
}

func (i *Snmpsysinfo) ToSnmpsysinfoOutput() SnmpsysinfoOutput {
	return i.ToSnmpsysinfoOutputWithContext(context.Background())
}

func (i *Snmpsysinfo) ToSnmpsysinfoOutputWithContext(ctx context.Context) SnmpsysinfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnmpsysinfoOutput)
}

// SnmpsysinfoArrayInput is an input type that accepts SnmpsysinfoArray and SnmpsysinfoArrayOutput values.
// You can construct a concrete instance of `SnmpsysinfoArrayInput` via:
//
//	SnmpsysinfoArray{ SnmpsysinfoArgs{...} }
type SnmpsysinfoArrayInput interface {
	pulumi.Input

	ToSnmpsysinfoArrayOutput() SnmpsysinfoArrayOutput
	ToSnmpsysinfoArrayOutputWithContext(context.Context) SnmpsysinfoArrayOutput
}

type SnmpsysinfoArray []SnmpsysinfoInput

func (SnmpsysinfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snmpsysinfo)(nil)).Elem()
}

func (i SnmpsysinfoArray) ToSnmpsysinfoArrayOutput() SnmpsysinfoArrayOutput {
	return i.ToSnmpsysinfoArrayOutputWithContext(context.Background())
}

func (i SnmpsysinfoArray) ToSnmpsysinfoArrayOutputWithContext(ctx context.Context) SnmpsysinfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnmpsysinfoArrayOutput)
}

// SnmpsysinfoMapInput is an input type that accepts SnmpsysinfoMap and SnmpsysinfoMapOutput values.
// You can construct a concrete instance of `SnmpsysinfoMapInput` via:
//
//	SnmpsysinfoMap{ "key": SnmpsysinfoArgs{...} }
type SnmpsysinfoMapInput interface {
	pulumi.Input

	ToSnmpsysinfoMapOutput() SnmpsysinfoMapOutput
	ToSnmpsysinfoMapOutputWithContext(context.Context) SnmpsysinfoMapOutput
}

type SnmpsysinfoMap map[string]SnmpsysinfoInput

func (SnmpsysinfoMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snmpsysinfo)(nil)).Elem()
}

func (i SnmpsysinfoMap) ToSnmpsysinfoMapOutput() SnmpsysinfoMapOutput {
	return i.ToSnmpsysinfoMapOutputWithContext(context.Background())
}

func (i SnmpsysinfoMap) ToSnmpsysinfoMapOutputWithContext(ctx context.Context) SnmpsysinfoMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnmpsysinfoMapOutput)
}

type SnmpsysinfoOutput struct{ *pulumi.OutputState }

func (SnmpsysinfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snmpsysinfo)(nil)).Elem()
}

func (o SnmpsysinfoOutput) ToSnmpsysinfoOutput() SnmpsysinfoOutput {
	return o
}

func (o SnmpsysinfoOutput) ToSnmpsysinfoOutputWithContext(ctx context.Context) SnmpsysinfoOutput {
	return o
}

// Contact information.
func (o SnmpsysinfoOutput) ContactInfo() pulumi.StringOutput {
	return o.ApplyT(func(v *Snmpsysinfo) pulumi.StringOutput { return v.ContactInfo }).(pulumi.StringOutput)
}

// System description.
func (o SnmpsysinfoOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Snmpsysinfo) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Local SNMP engine ID string (max 24 char).
func (o SnmpsysinfoOutput) EngineId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snmpsysinfo) pulumi.StringOutput { return v.EngineId }).(pulumi.StringOutput)
}

// System location.
func (o SnmpsysinfoOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Snmpsysinfo) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Enable/disable SNMP. Valid values: `disable`, `enable`.
func (o SnmpsysinfoOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Snmpsysinfo) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SnmpsysinfoOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snmpsysinfo) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SnmpsysinfoArrayOutput struct{ *pulumi.OutputState }

func (SnmpsysinfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snmpsysinfo)(nil)).Elem()
}

func (o SnmpsysinfoArrayOutput) ToSnmpsysinfoArrayOutput() SnmpsysinfoArrayOutput {
	return o
}

func (o SnmpsysinfoArrayOutput) ToSnmpsysinfoArrayOutputWithContext(ctx context.Context) SnmpsysinfoArrayOutput {
	return o
}

func (o SnmpsysinfoArrayOutput) Index(i pulumi.IntInput) SnmpsysinfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Snmpsysinfo {
		return vs[0].([]*Snmpsysinfo)[vs[1].(int)]
	}).(SnmpsysinfoOutput)
}

type SnmpsysinfoMapOutput struct{ *pulumi.OutputState }

func (SnmpsysinfoMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snmpsysinfo)(nil)).Elem()
}

func (o SnmpsysinfoMapOutput) ToSnmpsysinfoMapOutput() SnmpsysinfoMapOutput {
	return o
}

func (o SnmpsysinfoMapOutput) ToSnmpsysinfoMapOutputWithContext(ctx context.Context) SnmpsysinfoMapOutput {
	return o
}

func (o SnmpsysinfoMapOutput) MapIndex(k pulumi.StringInput) SnmpsysinfoOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Snmpsysinfo {
		return vs[0].(map[string]*Snmpsysinfo)[vs[1].(string)]
	}).(SnmpsysinfoOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnmpsysinfoInput)(nil)).Elem(), &Snmpsysinfo{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnmpsysinfoArrayInput)(nil)).Elem(), SnmpsysinfoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnmpsysinfoMapInput)(nil)).Elem(), SnmpsysinfoMap{})
	pulumi.RegisterOutputType(SnmpsysinfoOutput{})
	pulumi.RegisterOutputType(SnmpsysinfoArrayOutput{})
	pulumi.RegisterOutputType(SnmpsysinfoMapOutput{})
}
