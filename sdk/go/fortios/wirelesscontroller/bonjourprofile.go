// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wirelesscontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure Bonjour profiles. Bonjour is Apple's zero configuration networking protocol. Bonjour profiles allow APs and FortiAPs to connnect to networks using Bonjour.
//
// ## Import
//
// WirelessController BonjourProfile can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/bonjourprofile:Bonjourprofile labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/bonjourprofile:Bonjourprofile labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Bonjourprofile struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Bonjour profile name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Bonjour policy list. The structure of `policyList` block is documented below.
	PolicyLists BonjourprofilePolicyListArrayOutput `pulumi:"policyLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewBonjourprofile registers a new resource with the given unique name, arguments, and options.
func NewBonjourprofile(ctx *pulumi.Context,
	name string, args *BonjourprofileArgs, opts ...pulumi.ResourceOption) (*Bonjourprofile, error) {
	if args == nil {
		args = &BonjourprofileArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Bonjourprofile
	err := ctx.RegisterResource("fortios:wirelesscontroller/bonjourprofile:Bonjourprofile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBonjourprofile gets an existing Bonjourprofile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBonjourprofile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BonjourprofileState, opts ...pulumi.ResourceOption) (*Bonjourprofile, error) {
	var resource Bonjourprofile
	err := ctx.ReadResource("fortios:wirelesscontroller/bonjourprofile:Bonjourprofile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bonjourprofile resources.
type bonjourprofileState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Bonjour profile name.
	Name *string `pulumi:"name"`
	// Bonjour policy list. The structure of `policyList` block is documented below.
	PolicyLists []BonjourprofilePolicyList `pulumi:"policyLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type BonjourprofileState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Bonjour profile name.
	Name pulumi.StringPtrInput
	// Bonjour policy list. The structure of `policyList` block is documented below.
	PolicyLists BonjourprofilePolicyListArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (BonjourprofileState) ElementType() reflect.Type {
	return reflect.TypeOf((*bonjourprofileState)(nil)).Elem()
}

type bonjourprofileArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Bonjour profile name.
	Name *string `pulumi:"name"`
	// Bonjour policy list. The structure of `policyList` block is documented below.
	PolicyLists []BonjourprofilePolicyList `pulumi:"policyLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Bonjourprofile resource.
type BonjourprofileArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Bonjour profile name.
	Name pulumi.StringPtrInput
	// Bonjour policy list. The structure of `policyList` block is documented below.
	PolicyLists BonjourprofilePolicyListArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (BonjourprofileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bonjourprofileArgs)(nil)).Elem()
}

type BonjourprofileInput interface {
	pulumi.Input

	ToBonjourprofileOutput() BonjourprofileOutput
	ToBonjourprofileOutputWithContext(ctx context.Context) BonjourprofileOutput
}

func (*Bonjourprofile) ElementType() reflect.Type {
	return reflect.TypeOf((**Bonjourprofile)(nil)).Elem()
}

func (i *Bonjourprofile) ToBonjourprofileOutput() BonjourprofileOutput {
	return i.ToBonjourprofileOutputWithContext(context.Background())
}

func (i *Bonjourprofile) ToBonjourprofileOutputWithContext(ctx context.Context) BonjourprofileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BonjourprofileOutput)
}

// BonjourprofileArrayInput is an input type that accepts BonjourprofileArray and BonjourprofileArrayOutput values.
// You can construct a concrete instance of `BonjourprofileArrayInput` via:
//
//	BonjourprofileArray{ BonjourprofileArgs{...} }
type BonjourprofileArrayInput interface {
	pulumi.Input

	ToBonjourprofileArrayOutput() BonjourprofileArrayOutput
	ToBonjourprofileArrayOutputWithContext(context.Context) BonjourprofileArrayOutput
}

type BonjourprofileArray []BonjourprofileInput

func (BonjourprofileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bonjourprofile)(nil)).Elem()
}

func (i BonjourprofileArray) ToBonjourprofileArrayOutput() BonjourprofileArrayOutput {
	return i.ToBonjourprofileArrayOutputWithContext(context.Background())
}

func (i BonjourprofileArray) ToBonjourprofileArrayOutputWithContext(ctx context.Context) BonjourprofileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BonjourprofileArrayOutput)
}

// BonjourprofileMapInput is an input type that accepts BonjourprofileMap and BonjourprofileMapOutput values.
// You can construct a concrete instance of `BonjourprofileMapInput` via:
//
//	BonjourprofileMap{ "key": BonjourprofileArgs{...} }
type BonjourprofileMapInput interface {
	pulumi.Input

	ToBonjourprofileMapOutput() BonjourprofileMapOutput
	ToBonjourprofileMapOutputWithContext(context.Context) BonjourprofileMapOutput
}

type BonjourprofileMap map[string]BonjourprofileInput

func (BonjourprofileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bonjourprofile)(nil)).Elem()
}

func (i BonjourprofileMap) ToBonjourprofileMapOutput() BonjourprofileMapOutput {
	return i.ToBonjourprofileMapOutputWithContext(context.Background())
}

func (i BonjourprofileMap) ToBonjourprofileMapOutputWithContext(ctx context.Context) BonjourprofileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BonjourprofileMapOutput)
}

type BonjourprofileOutput struct{ *pulumi.OutputState }

func (BonjourprofileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bonjourprofile)(nil)).Elem()
}

func (o BonjourprofileOutput) ToBonjourprofileOutput() BonjourprofileOutput {
	return o
}

func (o BonjourprofileOutput) ToBonjourprofileOutputWithContext(ctx context.Context) BonjourprofileOutput {
	return o
}

// Comment.
func (o BonjourprofileOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v *Bonjourprofile) pulumi.StringOutput { return v.Comment }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o BonjourprofileOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bonjourprofile) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o BonjourprofileOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bonjourprofile) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Bonjour profile name.
func (o BonjourprofileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Bonjourprofile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Bonjour policy list. The structure of `policyList` block is documented below.
func (o BonjourprofileOutput) PolicyLists() BonjourprofilePolicyListArrayOutput {
	return o.ApplyT(func(v *Bonjourprofile) BonjourprofilePolicyListArrayOutput { return v.PolicyLists }).(BonjourprofilePolicyListArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o BonjourprofileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bonjourprofile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type BonjourprofileArrayOutput struct{ *pulumi.OutputState }

func (BonjourprofileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bonjourprofile)(nil)).Elem()
}

func (o BonjourprofileArrayOutput) ToBonjourprofileArrayOutput() BonjourprofileArrayOutput {
	return o
}

func (o BonjourprofileArrayOutput) ToBonjourprofileArrayOutputWithContext(ctx context.Context) BonjourprofileArrayOutput {
	return o
}

func (o BonjourprofileArrayOutput) Index(i pulumi.IntInput) BonjourprofileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Bonjourprofile {
		return vs[0].([]*Bonjourprofile)[vs[1].(int)]
	}).(BonjourprofileOutput)
}

type BonjourprofileMapOutput struct{ *pulumi.OutputState }

func (BonjourprofileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bonjourprofile)(nil)).Elem()
}

func (o BonjourprofileMapOutput) ToBonjourprofileMapOutput() BonjourprofileMapOutput {
	return o
}

func (o BonjourprofileMapOutput) ToBonjourprofileMapOutputWithContext(ctx context.Context) BonjourprofileMapOutput {
	return o
}

func (o BonjourprofileMapOutput) MapIndex(k pulumi.StringInput) BonjourprofileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Bonjourprofile {
		return vs[0].(map[string]*Bonjourprofile)[vs[1].(string)]
	}).(BonjourprofileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BonjourprofileInput)(nil)).Elem(), &Bonjourprofile{})
	pulumi.RegisterInputType(reflect.TypeOf((*BonjourprofileArrayInput)(nil)).Elem(), BonjourprofileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BonjourprofileMapInput)(nil)).Elem(), BonjourprofileMap{})
	pulumi.RegisterOutputType(BonjourprofileOutput{})
	pulumi.RegisterOutputType(BonjourprofileArrayOutput{})
	pulumi.RegisterOutputType(BonjourprofileMapOutput{})
}
