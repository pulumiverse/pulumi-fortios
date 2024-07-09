// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure community lists.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/router"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := router.NewCommunitylist(ctx, "trname", &router.CommunitylistArgs{
//				Rules: router.CommunitylistRuleArray{
//					&router.CommunitylistRuleArgs{
//						Action: pulumi.String("deny"),
//						Match:  pulumi.String("123:234 345:456"),
//						Regexp: pulumi.String("123:234 345:456"),
//					},
//				},
//				Type: pulumi.String("standard"),
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
// Router CommunityList can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:router/communitylist:Communitylist labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:router/communitylist:Communitylist labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Communitylist struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Community list name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Community list rule. The structure of `rule` block is documented below.
	Rules CommunitylistRuleArrayOutput `pulumi:"rules"`
	// Community list type (standard or expanded). Valid values: `standard`, `expanded`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewCommunitylist registers a new resource with the given unique name, arguments, and options.
func NewCommunitylist(ctx *pulumi.Context,
	name string, args *CommunitylistArgs, opts ...pulumi.ResourceOption) (*Communitylist, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Communitylist
	err := ctx.RegisterResource("fortios:router/communitylist:Communitylist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCommunitylist gets an existing Communitylist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCommunitylist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommunitylistState, opts ...pulumi.ResourceOption) (*Communitylist, error) {
	var resource Communitylist
	err := ctx.ReadResource("fortios:router/communitylist:Communitylist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Communitylist resources.
type communitylistState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Community list name.
	Name *string `pulumi:"name"`
	// Community list rule. The structure of `rule` block is documented below.
	Rules []CommunitylistRule `pulumi:"rules"`
	// Community list type (standard or expanded). Valid values: `standard`, `expanded`.
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type CommunitylistState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Community list name.
	Name pulumi.StringPtrInput
	// Community list rule. The structure of `rule` block is documented below.
	Rules CommunitylistRuleArrayInput
	// Community list type (standard or expanded). Valid values: `standard`, `expanded`.
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (CommunitylistState) ElementType() reflect.Type {
	return reflect.TypeOf((*communitylistState)(nil)).Elem()
}

type communitylistArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Community list name.
	Name *string `pulumi:"name"`
	// Community list rule. The structure of `rule` block is documented below.
	Rules []CommunitylistRule `pulumi:"rules"`
	// Community list type (standard or expanded). Valid values: `standard`, `expanded`.
	Type string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Communitylist resource.
type CommunitylistArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Community list name.
	Name pulumi.StringPtrInput
	// Community list rule. The structure of `rule` block is documented below.
	Rules CommunitylistRuleArrayInput
	// Community list type (standard or expanded). Valid values: `standard`, `expanded`.
	Type pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (CommunitylistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*communitylistArgs)(nil)).Elem()
}

type CommunitylistInput interface {
	pulumi.Input

	ToCommunitylistOutput() CommunitylistOutput
	ToCommunitylistOutputWithContext(ctx context.Context) CommunitylistOutput
}

func (*Communitylist) ElementType() reflect.Type {
	return reflect.TypeOf((**Communitylist)(nil)).Elem()
}

func (i *Communitylist) ToCommunitylistOutput() CommunitylistOutput {
	return i.ToCommunitylistOutputWithContext(context.Background())
}

func (i *Communitylist) ToCommunitylistOutputWithContext(ctx context.Context) CommunitylistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommunitylistOutput)
}

// CommunitylistArrayInput is an input type that accepts CommunitylistArray and CommunitylistArrayOutput values.
// You can construct a concrete instance of `CommunitylistArrayInput` via:
//
//	CommunitylistArray{ CommunitylistArgs{...} }
type CommunitylistArrayInput interface {
	pulumi.Input

	ToCommunitylistArrayOutput() CommunitylistArrayOutput
	ToCommunitylistArrayOutputWithContext(context.Context) CommunitylistArrayOutput
}

type CommunitylistArray []CommunitylistInput

func (CommunitylistArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Communitylist)(nil)).Elem()
}

func (i CommunitylistArray) ToCommunitylistArrayOutput() CommunitylistArrayOutput {
	return i.ToCommunitylistArrayOutputWithContext(context.Background())
}

func (i CommunitylistArray) ToCommunitylistArrayOutputWithContext(ctx context.Context) CommunitylistArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommunitylistArrayOutput)
}

// CommunitylistMapInput is an input type that accepts CommunitylistMap and CommunitylistMapOutput values.
// You can construct a concrete instance of `CommunitylistMapInput` via:
//
//	CommunitylistMap{ "key": CommunitylistArgs{...} }
type CommunitylistMapInput interface {
	pulumi.Input

	ToCommunitylistMapOutput() CommunitylistMapOutput
	ToCommunitylistMapOutputWithContext(context.Context) CommunitylistMapOutput
}

type CommunitylistMap map[string]CommunitylistInput

func (CommunitylistMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Communitylist)(nil)).Elem()
}

func (i CommunitylistMap) ToCommunitylistMapOutput() CommunitylistMapOutput {
	return i.ToCommunitylistMapOutputWithContext(context.Background())
}

func (i CommunitylistMap) ToCommunitylistMapOutputWithContext(ctx context.Context) CommunitylistMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommunitylistMapOutput)
}

type CommunitylistOutput struct{ *pulumi.OutputState }

func (CommunitylistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Communitylist)(nil)).Elem()
}

func (o CommunitylistOutput) ToCommunitylistOutput() CommunitylistOutput {
	return o
}

func (o CommunitylistOutput) ToCommunitylistOutputWithContext(ctx context.Context) CommunitylistOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o CommunitylistOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Communitylist) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o CommunitylistOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Communitylist) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Community list name.
func (o CommunitylistOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Communitylist) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Community list rule. The structure of `rule` block is documented below.
func (o CommunitylistOutput) Rules() CommunitylistRuleArrayOutput {
	return o.ApplyT(func(v *Communitylist) CommunitylistRuleArrayOutput { return v.Rules }).(CommunitylistRuleArrayOutput)
}

// Community list type (standard or expanded). Valid values: `standard`, `expanded`.
func (o CommunitylistOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Communitylist) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o CommunitylistOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Communitylist) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type CommunitylistArrayOutput struct{ *pulumi.OutputState }

func (CommunitylistArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Communitylist)(nil)).Elem()
}

func (o CommunitylistArrayOutput) ToCommunitylistArrayOutput() CommunitylistArrayOutput {
	return o
}

func (o CommunitylistArrayOutput) ToCommunitylistArrayOutputWithContext(ctx context.Context) CommunitylistArrayOutput {
	return o
}

func (o CommunitylistArrayOutput) Index(i pulumi.IntInput) CommunitylistOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Communitylist {
		return vs[0].([]*Communitylist)[vs[1].(int)]
	}).(CommunitylistOutput)
}

type CommunitylistMapOutput struct{ *pulumi.OutputState }

func (CommunitylistMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Communitylist)(nil)).Elem()
}

func (o CommunitylistMapOutput) ToCommunitylistMapOutput() CommunitylistMapOutput {
	return o
}

func (o CommunitylistMapOutput) ToCommunitylistMapOutputWithContext(ctx context.Context) CommunitylistMapOutput {
	return o
}

func (o CommunitylistMapOutput) MapIndex(k pulumi.StringInput) CommunitylistOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Communitylist {
		return vs[0].(map[string]*Communitylist)[vs[1].(string)]
	}).(CommunitylistOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CommunitylistInput)(nil)).Elem(), &Communitylist{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommunitylistArrayInput)(nil)).Elem(), CommunitylistArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommunitylistMapInput)(nil)).Elem(), CommunitylistMap{})
	pulumi.RegisterOutputType(CommunitylistOutput{})
	pulumi.RegisterOutputType(CommunitylistArrayOutput{})
	pulumi.RegisterOutputType(CommunitylistMapOutput{})
}
