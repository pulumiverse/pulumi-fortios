// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ipsec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Concentrator configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/vpn"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpn.NewConcentrator(ctx, "trname", &vpn.ConcentratorArgs{
//				SrcCheck: pulumi.String("disable"),
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
// VpnIpsec Concentrator can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:vpn/ipsec/concentrator:Concentrator labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:vpn/ipsec/concentrator:Concentrator labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Concentrator struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Concentrator ID. (1-65535)
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Names of up to 3 VPN tunnels to add to the concentrator. The structure of `member` block is documented below.
	Members ConcentratorMemberArrayOutput `pulumi:"members"`
	// Concentrator name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable to check source address of phase 2 selector. Disable to check only the destination selector. Valid values: `disable`, `enable`.
	SrcCheck pulumi.StringOutput `pulumi:"srcCheck"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewConcentrator registers a new resource with the given unique name, arguments, and options.
func NewConcentrator(ctx *pulumi.Context,
	name string, args *ConcentratorArgs, opts ...pulumi.ResourceOption) (*Concentrator, error) {
	if args == nil {
		args = &ConcentratorArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Concentrator
	err := ctx.RegisterResource("fortios:vpn/ipsec/concentrator:Concentrator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConcentrator gets an existing Concentrator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConcentrator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConcentratorState, opts ...pulumi.ResourceOption) (*Concentrator, error) {
	var resource Concentrator
	err := ctx.ReadResource("fortios:vpn/ipsec/concentrator:Concentrator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Concentrator resources.
type concentratorState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Concentrator ID. (1-65535)
	Fosid *int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Names of up to 3 VPN tunnels to add to the concentrator. The structure of `member` block is documented below.
	Members []ConcentratorMember `pulumi:"members"`
	// Concentrator name.
	Name *string `pulumi:"name"`
	// Enable to check source address of phase 2 selector. Disable to check only the destination selector. Valid values: `disable`, `enable`.
	SrcCheck *string `pulumi:"srcCheck"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ConcentratorState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Concentrator ID. (1-65535)
	Fosid pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Names of up to 3 VPN tunnels to add to the concentrator. The structure of `member` block is documented below.
	Members ConcentratorMemberArrayInput
	// Concentrator name.
	Name pulumi.StringPtrInput
	// Enable to check source address of phase 2 selector. Disable to check only the destination selector. Valid values: `disable`, `enable`.
	SrcCheck pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ConcentratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*concentratorState)(nil)).Elem()
}

type concentratorArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Concentrator ID. (1-65535)
	Fosid *int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Names of up to 3 VPN tunnels to add to the concentrator. The structure of `member` block is documented below.
	Members []ConcentratorMember `pulumi:"members"`
	// Concentrator name.
	Name *string `pulumi:"name"`
	// Enable to check source address of phase 2 selector. Disable to check only the destination selector. Valid values: `disable`, `enable`.
	SrcCheck *string `pulumi:"srcCheck"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Concentrator resource.
type ConcentratorArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Concentrator ID. (1-65535)
	Fosid pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Names of up to 3 VPN tunnels to add to the concentrator. The structure of `member` block is documented below.
	Members ConcentratorMemberArrayInput
	// Concentrator name.
	Name pulumi.StringPtrInput
	// Enable to check source address of phase 2 selector. Disable to check only the destination selector. Valid values: `disable`, `enable`.
	SrcCheck pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ConcentratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*concentratorArgs)(nil)).Elem()
}

type ConcentratorInput interface {
	pulumi.Input

	ToConcentratorOutput() ConcentratorOutput
	ToConcentratorOutputWithContext(ctx context.Context) ConcentratorOutput
}

func (*Concentrator) ElementType() reflect.Type {
	return reflect.TypeOf((**Concentrator)(nil)).Elem()
}

func (i *Concentrator) ToConcentratorOutput() ConcentratorOutput {
	return i.ToConcentratorOutputWithContext(context.Background())
}

func (i *Concentrator) ToConcentratorOutputWithContext(ctx context.Context) ConcentratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConcentratorOutput)
}

// ConcentratorArrayInput is an input type that accepts ConcentratorArray and ConcentratorArrayOutput values.
// You can construct a concrete instance of `ConcentratorArrayInput` via:
//
//	ConcentratorArray{ ConcentratorArgs{...} }
type ConcentratorArrayInput interface {
	pulumi.Input

	ToConcentratorArrayOutput() ConcentratorArrayOutput
	ToConcentratorArrayOutputWithContext(context.Context) ConcentratorArrayOutput
}

type ConcentratorArray []ConcentratorInput

func (ConcentratorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Concentrator)(nil)).Elem()
}

func (i ConcentratorArray) ToConcentratorArrayOutput() ConcentratorArrayOutput {
	return i.ToConcentratorArrayOutputWithContext(context.Background())
}

func (i ConcentratorArray) ToConcentratorArrayOutputWithContext(ctx context.Context) ConcentratorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConcentratorArrayOutput)
}

// ConcentratorMapInput is an input type that accepts ConcentratorMap and ConcentratorMapOutput values.
// You can construct a concrete instance of `ConcentratorMapInput` via:
//
//	ConcentratorMap{ "key": ConcentratorArgs{...} }
type ConcentratorMapInput interface {
	pulumi.Input

	ToConcentratorMapOutput() ConcentratorMapOutput
	ToConcentratorMapOutputWithContext(context.Context) ConcentratorMapOutput
}

type ConcentratorMap map[string]ConcentratorInput

func (ConcentratorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Concentrator)(nil)).Elem()
}

func (i ConcentratorMap) ToConcentratorMapOutput() ConcentratorMapOutput {
	return i.ToConcentratorMapOutputWithContext(context.Background())
}

func (i ConcentratorMap) ToConcentratorMapOutputWithContext(ctx context.Context) ConcentratorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConcentratorMapOutput)
}

type ConcentratorOutput struct{ *pulumi.OutputState }

func (ConcentratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Concentrator)(nil)).Elem()
}

func (o ConcentratorOutput) ToConcentratorOutput() ConcentratorOutput {
	return o
}

func (o ConcentratorOutput) ToConcentratorOutputWithContext(ctx context.Context) ConcentratorOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o ConcentratorOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Concentrator) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Concentrator ID. (1-65535)
func (o ConcentratorOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Concentrator) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o ConcentratorOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Concentrator) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Names of up to 3 VPN tunnels to add to the concentrator. The structure of `member` block is documented below.
func (o ConcentratorOutput) Members() ConcentratorMemberArrayOutput {
	return o.ApplyT(func(v *Concentrator) ConcentratorMemberArrayOutput { return v.Members }).(ConcentratorMemberArrayOutput)
}

// Concentrator name.
func (o ConcentratorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Concentrator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable to check source address of phase 2 selector. Disable to check only the destination selector. Valid values: `disable`, `enable`.
func (o ConcentratorOutput) SrcCheck() pulumi.StringOutput {
	return o.ApplyT(func(v *Concentrator) pulumi.StringOutput { return v.SrcCheck }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ConcentratorOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Concentrator) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ConcentratorArrayOutput struct{ *pulumi.OutputState }

func (ConcentratorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Concentrator)(nil)).Elem()
}

func (o ConcentratorArrayOutput) ToConcentratorArrayOutput() ConcentratorArrayOutput {
	return o
}

func (o ConcentratorArrayOutput) ToConcentratorArrayOutputWithContext(ctx context.Context) ConcentratorArrayOutput {
	return o
}

func (o ConcentratorArrayOutput) Index(i pulumi.IntInput) ConcentratorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Concentrator {
		return vs[0].([]*Concentrator)[vs[1].(int)]
	}).(ConcentratorOutput)
}

type ConcentratorMapOutput struct{ *pulumi.OutputState }

func (ConcentratorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Concentrator)(nil)).Elem()
}

func (o ConcentratorMapOutput) ToConcentratorMapOutput() ConcentratorMapOutput {
	return o
}

func (o ConcentratorMapOutput) ToConcentratorMapOutputWithContext(ctx context.Context) ConcentratorMapOutput {
	return o
}

func (o ConcentratorMapOutput) MapIndex(k pulumi.StringInput) ConcentratorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Concentrator {
		return vs[0].(map[string]*Concentrator)[vs[1].(string)]
	}).(ConcentratorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConcentratorInput)(nil)).Elem(), &Concentrator{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConcentratorArrayInput)(nil)).Elem(), ConcentratorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConcentratorMapInput)(nil)).Elem(), ConcentratorMap{})
	pulumi.RegisterOutputType(ConcentratorOutput{})
	pulumi.RegisterOutputType(ConcentratorArrayOutput{})
	pulumi.RegisterOutputType(ConcentratorMapOutput{})
}
