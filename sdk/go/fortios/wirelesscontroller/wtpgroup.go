// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wirelesscontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure WTP groups.
//
// ## Import
//
// WirelessController WtpGroup can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/wtpgroup:Wtpgroup labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/wtpgroup:Wtpgroup labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Wtpgroup struct {
	pulumi.CustomResourceState

	// Override BLE Major ID.
	BleMajorId pulumi.IntOutput `pulumi:"bleMajorId"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// WTP group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// FortiAP models to define the WTP group platform type.
	PlatformType pulumi.StringOutput `pulumi:"platformType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// WTP list. The structure of `wtps` block is documented below.
	Wtps WtpgroupWtpArrayOutput `pulumi:"wtps"`
}

// NewWtpgroup registers a new resource with the given unique name, arguments, and options.
func NewWtpgroup(ctx *pulumi.Context,
	name string, args *WtpgroupArgs, opts ...pulumi.ResourceOption) (*Wtpgroup, error) {
	if args == nil {
		args = &WtpgroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Wtpgroup
	err := ctx.RegisterResource("fortios:wirelesscontroller/wtpgroup:Wtpgroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWtpgroup gets an existing Wtpgroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWtpgroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WtpgroupState, opts ...pulumi.ResourceOption) (*Wtpgroup, error) {
	var resource Wtpgroup
	err := ctx.ReadResource("fortios:wirelesscontroller/wtpgroup:Wtpgroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Wtpgroup resources.
type wtpgroupState struct {
	// Override BLE Major ID.
	BleMajorId *int `pulumi:"bleMajorId"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// WTP group name.
	Name *string `pulumi:"name"`
	// FortiAP models to define the WTP group platform type.
	PlatformType *string `pulumi:"platformType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// WTP list. The structure of `wtps` block is documented below.
	Wtps []WtpgroupWtp `pulumi:"wtps"`
}

type WtpgroupState struct {
	// Override BLE Major ID.
	BleMajorId pulumi.IntPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// WTP group name.
	Name pulumi.StringPtrInput
	// FortiAP models to define the WTP group platform type.
	PlatformType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// WTP list. The structure of `wtps` block is documented below.
	Wtps WtpgroupWtpArrayInput
}

func (WtpgroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*wtpgroupState)(nil)).Elem()
}

type wtpgroupArgs struct {
	// Override BLE Major ID.
	BleMajorId *int `pulumi:"bleMajorId"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// WTP group name.
	Name *string `pulumi:"name"`
	// FortiAP models to define the WTP group platform type.
	PlatformType *string `pulumi:"platformType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// WTP list. The structure of `wtps` block is documented below.
	Wtps []WtpgroupWtp `pulumi:"wtps"`
}

// The set of arguments for constructing a Wtpgroup resource.
type WtpgroupArgs struct {
	// Override BLE Major ID.
	BleMajorId pulumi.IntPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// WTP group name.
	Name pulumi.StringPtrInput
	// FortiAP models to define the WTP group platform type.
	PlatformType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// WTP list. The structure of `wtps` block is documented below.
	Wtps WtpgroupWtpArrayInput
}

func (WtpgroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wtpgroupArgs)(nil)).Elem()
}

type WtpgroupInput interface {
	pulumi.Input

	ToWtpgroupOutput() WtpgroupOutput
	ToWtpgroupOutputWithContext(ctx context.Context) WtpgroupOutput
}

func (*Wtpgroup) ElementType() reflect.Type {
	return reflect.TypeOf((**Wtpgroup)(nil)).Elem()
}

func (i *Wtpgroup) ToWtpgroupOutput() WtpgroupOutput {
	return i.ToWtpgroupOutputWithContext(context.Background())
}

func (i *Wtpgroup) ToWtpgroupOutputWithContext(ctx context.Context) WtpgroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WtpgroupOutput)
}

// WtpgroupArrayInput is an input type that accepts WtpgroupArray and WtpgroupArrayOutput values.
// You can construct a concrete instance of `WtpgroupArrayInput` via:
//
//	WtpgroupArray{ WtpgroupArgs{...} }
type WtpgroupArrayInput interface {
	pulumi.Input

	ToWtpgroupArrayOutput() WtpgroupArrayOutput
	ToWtpgroupArrayOutputWithContext(context.Context) WtpgroupArrayOutput
}

type WtpgroupArray []WtpgroupInput

func (WtpgroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wtpgroup)(nil)).Elem()
}

func (i WtpgroupArray) ToWtpgroupArrayOutput() WtpgroupArrayOutput {
	return i.ToWtpgroupArrayOutputWithContext(context.Background())
}

func (i WtpgroupArray) ToWtpgroupArrayOutputWithContext(ctx context.Context) WtpgroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WtpgroupArrayOutput)
}

// WtpgroupMapInput is an input type that accepts WtpgroupMap and WtpgroupMapOutput values.
// You can construct a concrete instance of `WtpgroupMapInput` via:
//
//	WtpgroupMap{ "key": WtpgroupArgs{...} }
type WtpgroupMapInput interface {
	pulumi.Input

	ToWtpgroupMapOutput() WtpgroupMapOutput
	ToWtpgroupMapOutputWithContext(context.Context) WtpgroupMapOutput
}

type WtpgroupMap map[string]WtpgroupInput

func (WtpgroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wtpgroup)(nil)).Elem()
}

func (i WtpgroupMap) ToWtpgroupMapOutput() WtpgroupMapOutput {
	return i.ToWtpgroupMapOutputWithContext(context.Background())
}

func (i WtpgroupMap) ToWtpgroupMapOutputWithContext(ctx context.Context) WtpgroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WtpgroupMapOutput)
}

type WtpgroupOutput struct{ *pulumi.OutputState }

func (WtpgroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Wtpgroup)(nil)).Elem()
}

func (o WtpgroupOutput) ToWtpgroupOutput() WtpgroupOutput {
	return o
}

func (o WtpgroupOutput) ToWtpgroupOutputWithContext(ctx context.Context) WtpgroupOutput {
	return o
}

// Override BLE Major ID.
func (o WtpgroupOutput) BleMajorId() pulumi.IntOutput {
	return o.ApplyT(func(v *Wtpgroup) pulumi.IntOutput { return v.BleMajorId }).(pulumi.IntOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o WtpgroupOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wtpgroup) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o WtpgroupOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wtpgroup) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// WTP group name.
func (o WtpgroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Wtpgroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// FortiAP models to define the WTP group platform type.
func (o WtpgroupOutput) PlatformType() pulumi.StringOutput {
	return o.ApplyT(func(v *Wtpgroup) pulumi.StringOutput { return v.PlatformType }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o WtpgroupOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wtpgroup) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// WTP list. The structure of `wtps` block is documented below.
func (o WtpgroupOutput) Wtps() WtpgroupWtpArrayOutput {
	return o.ApplyT(func(v *Wtpgroup) WtpgroupWtpArrayOutput { return v.Wtps }).(WtpgroupWtpArrayOutput)
}

type WtpgroupArrayOutput struct{ *pulumi.OutputState }

func (WtpgroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wtpgroup)(nil)).Elem()
}

func (o WtpgroupArrayOutput) ToWtpgroupArrayOutput() WtpgroupArrayOutput {
	return o
}

func (o WtpgroupArrayOutput) ToWtpgroupArrayOutputWithContext(ctx context.Context) WtpgroupArrayOutput {
	return o
}

func (o WtpgroupArrayOutput) Index(i pulumi.IntInput) WtpgroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Wtpgroup {
		return vs[0].([]*Wtpgroup)[vs[1].(int)]
	}).(WtpgroupOutput)
}

type WtpgroupMapOutput struct{ *pulumi.OutputState }

func (WtpgroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wtpgroup)(nil)).Elem()
}

func (o WtpgroupMapOutput) ToWtpgroupMapOutput() WtpgroupMapOutput {
	return o
}

func (o WtpgroupMapOutput) ToWtpgroupMapOutputWithContext(ctx context.Context) WtpgroupMapOutput {
	return o
}

func (o WtpgroupMapOutput) MapIndex(k pulumi.StringInput) WtpgroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Wtpgroup {
		return vs[0].(map[string]*Wtpgroup)[vs[1].(string)]
	}).(WtpgroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WtpgroupInput)(nil)).Elem(), &Wtpgroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*WtpgroupArrayInput)(nil)).Elem(), WtpgroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WtpgroupMapInput)(nil)).Elem(), WtpgroupMap{})
	pulumi.RegisterOutputType(WtpgroupOutput{})
	pulumi.RegisterOutputType(WtpgroupArrayOutput{})
	pulumi.RegisterOutputType(WtpgroupMapOutput{})
}
