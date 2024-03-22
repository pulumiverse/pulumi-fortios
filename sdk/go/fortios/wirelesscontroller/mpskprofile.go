// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wirelesscontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure MPSK profile. Applies to FortiOS Version `>= 6.4.2`.
//
// ## Import
//
// WirelessController MpskProfile can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/mpskprofile:Mpskprofile labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/mpskprofile:Mpskprofile labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Mpskprofile struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
	MpskConcurrentClients pulumi.IntOutput `pulumi:"mpskConcurrentClients"`
	// List of multiple PSK groups. The structure of `mpskGroup` block is documented below.
	MpskGroups MpskprofileMpskGroupArrayOutput `pulumi:"mpskGroups"`
	// MPSK profile name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewMpskprofile registers a new resource with the given unique name, arguments, and options.
func NewMpskprofile(ctx *pulumi.Context,
	name string, args *MpskprofileArgs, opts ...pulumi.ResourceOption) (*Mpskprofile, error) {
	if args == nil {
		args = &MpskprofileArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Mpskprofile
	err := ctx.RegisterResource("fortios:wirelesscontroller/mpskprofile:Mpskprofile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMpskprofile gets an existing Mpskprofile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMpskprofile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MpskprofileState, opts ...pulumi.ResourceOption) (*Mpskprofile, error) {
	var resource Mpskprofile
	err := ctx.ReadResource("fortios:wirelesscontroller/mpskprofile:Mpskprofile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Mpskprofile resources.
type mpskprofileState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
	MpskConcurrentClients *int `pulumi:"mpskConcurrentClients"`
	// List of multiple PSK groups. The structure of `mpskGroup` block is documented below.
	MpskGroups []MpskprofileMpskGroup `pulumi:"mpskGroups"`
	// MPSK profile name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type MpskprofileState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
	MpskConcurrentClients pulumi.IntPtrInput
	// List of multiple PSK groups. The structure of `mpskGroup` block is documented below.
	MpskGroups MpskprofileMpskGroupArrayInput
	// MPSK profile name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (MpskprofileState) ElementType() reflect.Type {
	return reflect.TypeOf((*mpskprofileState)(nil)).Elem()
}

type mpskprofileArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
	MpskConcurrentClients *int `pulumi:"mpskConcurrentClients"`
	// List of multiple PSK groups. The structure of `mpskGroup` block is documented below.
	MpskGroups []MpskprofileMpskGroup `pulumi:"mpskGroups"`
	// MPSK profile name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Mpskprofile resource.
type MpskprofileArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
	MpskConcurrentClients pulumi.IntPtrInput
	// List of multiple PSK groups. The structure of `mpskGroup` block is documented below.
	MpskGroups MpskprofileMpskGroupArrayInput
	// MPSK profile name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (MpskprofileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mpskprofileArgs)(nil)).Elem()
}

type MpskprofileInput interface {
	pulumi.Input

	ToMpskprofileOutput() MpskprofileOutput
	ToMpskprofileOutputWithContext(ctx context.Context) MpskprofileOutput
}

func (*Mpskprofile) ElementType() reflect.Type {
	return reflect.TypeOf((**Mpskprofile)(nil)).Elem()
}

func (i *Mpskprofile) ToMpskprofileOutput() MpskprofileOutput {
	return i.ToMpskprofileOutputWithContext(context.Background())
}

func (i *Mpskprofile) ToMpskprofileOutputWithContext(ctx context.Context) MpskprofileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MpskprofileOutput)
}

// MpskprofileArrayInput is an input type that accepts MpskprofileArray and MpskprofileArrayOutput values.
// You can construct a concrete instance of `MpskprofileArrayInput` via:
//
//	MpskprofileArray{ MpskprofileArgs{...} }
type MpskprofileArrayInput interface {
	pulumi.Input

	ToMpskprofileArrayOutput() MpskprofileArrayOutput
	ToMpskprofileArrayOutputWithContext(context.Context) MpskprofileArrayOutput
}

type MpskprofileArray []MpskprofileInput

func (MpskprofileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Mpskprofile)(nil)).Elem()
}

func (i MpskprofileArray) ToMpskprofileArrayOutput() MpskprofileArrayOutput {
	return i.ToMpskprofileArrayOutputWithContext(context.Background())
}

func (i MpskprofileArray) ToMpskprofileArrayOutputWithContext(ctx context.Context) MpskprofileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MpskprofileArrayOutput)
}

// MpskprofileMapInput is an input type that accepts MpskprofileMap and MpskprofileMapOutput values.
// You can construct a concrete instance of `MpskprofileMapInput` via:
//
//	MpskprofileMap{ "key": MpskprofileArgs{...} }
type MpskprofileMapInput interface {
	pulumi.Input

	ToMpskprofileMapOutput() MpskprofileMapOutput
	ToMpskprofileMapOutputWithContext(context.Context) MpskprofileMapOutput
}

type MpskprofileMap map[string]MpskprofileInput

func (MpskprofileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Mpskprofile)(nil)).Elem()
}

func (i MpskprofileMap) ToMpskprofileMapOutput() MpskprofileMapOutput {
	return i.ToMpskprofileMapOutputWithContext(context.Background())
}

func (i MpskprofileMap) ToMpskprofileMapOutputWithContext(ctx context.Context) MpskprofileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MpskprofileMapOutput)
}

type MpskprofileOutput struct{ *pulumi.OutputState }

func (MpskprofileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Mpskprofile)(nil)).Elem()
}

func (o MpskprofileOutput) ToMpskprofileOutput() MpskprofileOutput {
	return o
}

func (o MpskprofileOutput) ToMpskprofileOutputWithContext(ctx context.Context) MpskprofileOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o MpskprofileOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mpskprofile) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o MpskprofileOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mpskprofile) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
func (o MpskprofileOutput) MpskConcurrentClients() pulumi.IntOutput {
	return o.ApplyT(func(v *Mpskprofile) pulumi.IntOutput { return v.MpskConcurrentClients }).(pulumi.IntOutput)
}

// List of multiple PSK groups. The structure of `mpskGroup` block is documented below.
func (o MpskprofileOutput) MpskGroups() MpskprofileMpskGroupArrayOutput {
	return o.ApplyT(func(v *Mpskprofile) MpskprofileMpskGroupArrayOutput { return v.MpskGroups }).(MpskprofileMpskGroupArrayOutput)
}

// MPSK profile name.
func (o MpskprofileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Mpskprofile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o MpskprofileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mpskprofile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type MpskprofileArrayOutput struct{ *pulumi.OutputState }

func (MpskprofileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Mpskprofile)(nil)).Elem()
}

func (o MpskprofileArrayOutput) ToMpskprofileArrayOutput() MpskprofileArrayOutput {
	return o
}

func (o MpskprofileArrayOutput) ToMpskprofileArrayOutputWithContext(ctx context.Context) MpskprofileArrayOutput {
	return o
}

func (o MpskprofileArrayOutput) Index(i pulumi.IntInput) MpskprofileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Mpskprofile {
		return vs[0].([]*Mpskprofile)[vs[1].(int)]
	}).(MpskprofileOutput)
}

type MpskprofileMapOutput struct{ *pulumi.OutputState }

func (MpskprofileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Mpskprofile)(nil)).Elem()
}

func (o MpskprofileMapOutput) ToMpskprofileMapOutput() MpskprofileMapOutput {
	return o
}

func (o MpskprofileMapOutput) ToMpskprofileMapOutputWithContext(ctx context.Context) MpskprofileMapOutput {
	return o
}

func (o MpskprofileMapOutput) MapIndex(k pulumi.StringInput) MpskprofileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Mpskprofile {
		return vs[0].(map[string]*Mpskprofile)[vs[1].(string)]
	}).(MpskprofileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MpskprofileInput)(nil)).Elem(), &Mpskprofile{})
	pulumi.RegisterInputType(reflect.TypeOf((*MpskprofileArrayInput)(nil)).Elem(), MpskprofileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MpskprofileMapInput)(nil)).Elem(), MpskprofileMap{})
	pulumi.RegisterOutputType(MpskprofileOutput{})
	pulumi.RegisterOutputType(MpskprofileArrayOutput{})
	pulumi.RegisterOutputType(MpskprofileMapOutput{})
}
