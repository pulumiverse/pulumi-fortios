// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Global configuration objects that can be configured independently for all VDOMs or for the defined VDOM scope.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/system"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := system.NewVdomexception(ctx, "trname", &system.VdomexceptionArgs{
//				Fosid:  pulumi.Int(1),
//				Object: pulumi.String("log.fortianalyzer.setting"),
//				Oid:    pulumi.Int(7150),
//				Scope:  pulumi.String("all"),
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
// System VdomException can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/vdomexception:Vdomexception labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/vdomexception:Vdomexception labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Vdomexception struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Index <1-4096>.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Name of the configuration object that can be configured independently for all VDOMs.
	Object pulumi.StringOutput `pulumi:"object"`
	// Object ID.
	Oid pulumi.IntOutput `pulumi:"oid"`
	// Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration. Valid values: `all`, `inclusive`, `exclusive`.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Names of the VDOMs. The structure of `vdom` block is documented below.
	Vdoms VdomexceptionVdomArrayOutput `pulumi:"vdoms"`
}

// NewVdomexception registers a new resource with the given unique name, arguments, and options.
func NewVdomexception(ctx *pulumi.Context,
	name string, args *VdomexceptionArgs, opts ...pulumi.ResourceOption) (*Vdomexception, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Object == nil {
		return nil, errors.New("invalid value for required argument 'Object'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Vdomexception
	err := ctx.RegisterResource("fortios:system/vdomexception:Vdomexception", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVdomexception gets an existing Vdomexception resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVdomexception(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VdomexceptionState, opts ...pulumi.ResourceOption) (*Vdomexception, error) {
	var resource Vdomexception
	err := ctx.ReadResource("fortios:system/vdomexception:Vdomexception", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vdomexception resources.
type vdomexceptionState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Index <1-4096>.
	Fosid *int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Name of the configuration object that can be configured independently for all VDOMs.
	Object *string `pulumi:"object"`
	// Object ID.
	Oid *int `pulumi:"oid"`
	// Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration. Valid values: `all`, `inclusive`, `exclusive`.
	Scope *string `pulumi:"scope"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Names of the VDOMs. The structure of `vdom` block is documented below.
	Vdoms []VdomexceptionVdom `pulumi:"vdoms"`
}

type VdomexceptionState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Index <1-4096>.
	Fosid pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Name of the configuration object that can be configured independently for all VDOMs.
	Object pulumi.StringPtrInput
	// Object ID.
	Oid pulumi.IntPtrInput
	// Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration. Valid values: `all`, `inclusive`, `exclusive`.
	Scope pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Names of the VDOMs. The structure of `vdom` block is documented below.
	Vdoms VdomexceptionVdomArrayInput
}

func (VdomexceptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*vdomexceptionState)(nil)).Elem()
}

type vdomexceptionArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Index <1-4096>.
	Fosid *int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Name of the configuration object that can be configured independently for all VDOMs.
	Object string `pulumi:"object"`
	// Object ID.
	Oid *int `pulumi:"oid"`
	// Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration. Valid values: `all`, `inclusive`, `exclusive`.
	Scope *string `pulumi:"scope"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Names of the VDOMs. The structure of `vdom` block is documented below.
	Vdoms []VdomexceptionVdom `pulumi:"vdoms"`
}

// The set of arguments for constructing a Vdomexception resource.
type VdomexceptionArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Index <1-4096>.
	Fosid pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Name of the configuration object that can be configured independently for all VDOMs.
	Object pulumi.StringInput
	// Object ID.
	Oid pulumi.IntPtrInput
	// Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration. Valid values: `all`, `inclusive`, `exclusive`.
	Scope pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Names of the VDOMs. The structure of `vdom` block is documented below.
	Vdoms VdomexceptionVdomArrayInput
}

func (VdomexceptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vdomexceptionArgs)(nil)).Elem()
}

type VdomexceptionInput interface {
	pulumi.Input

	ToVdomexceptionOutput() VdomexceptionOutput
	ToVdomexceptionOutputWithContext(ctx context.Context) VdomexceptionOutput
}

func (*Vdomexception) ElementType() reflect.Type {
	return reflect.TypeOf((**Vdomexception)(nil)).Elem()
}

func (i *Vdomexception) ToVdomexceptionOutput() VdomexceptionOutput {
	return i.ToVdomexceptionOutputWithContext(context.Background())
}

func (i *Vdomexception) ToVdomexceptionOutputWithContext(ctx context.Context) VdomexceptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VdomexceptionOutput)
}

// VdomexceptionArrayInput is an input type that accepts VdomexceptionArray and VdomexceptionArrayOutput values.
// You can construct a concrete instance of `VdomexceptionArrayInput` via:
//
//	VdomexceptionArray{ VdomexceptionArgs{...} }
type VdomexceptionArrayInput interface {
	pulumi.Input

	ToVdomexceptionArrayOutput() VdomexceptionArrayOutput
	ToVdomexceptionArrayOutputWithContext(context.Context) VdomexceptionArrayOutput
}

type VdomexceptionArray []VdomexceptionInput

func (VdomexceptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vdomexception)(nil)).Elem()
}

func (i VdomexceptionArray) ToVdomexceptionArrayOutput() VdomexceptionArrayOutput {
	return i.ToVdomexceptionArrayOutputWithContext(context.Background())
}

func (i VdomexceptionArray) ToVdomexceptionArrayOutputWithContext(ctx context.Context) VdomexceptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VdomexceptionArrayOutput)
}

// VdomexceptionMapInput is an input type that accepts VdomexceptionMap and VdomexceptionMapOutput values.
// You can construct a concrete instance of `VdomexceptionMapInput` via:
//
//	VdomexceptionMap{ "key": VdomexceptionArgs{...} }
type VdomexceptionMapInput interface {
	pulumi.Input

	ToVdomexceptionMapOutput() VdomexceptionMapOutput
	ToVdomexceptionMapOutputWithContext(context.Context) VdomexceptionMapOutput
}

type VdomexceptionMap map[string]VdomexceptionInput

func (VdomexceptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vdomexception)(nil)).Elem()
}

func (i VdomexceptionMap) ToVdomexceptionMapOutput() VdomexceptionMapOutput {
	return i.ToVdomexceptionMapOutputWithContext(context.Background())
}

func (i VdomexceptionMap) ToVdomexceptionMapOutputWithContext(ctx context.Context) VdomexceptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VdomexceptionMapOutput)
}

type VdomexceptionOutput struct{ *pulumi.OutputState }

func (VdomexceptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vdomexception)(nil)).Elem()
}

func (o VdomexceptionOutput) ToVdomexceptionOutput() VdomexceptionOutput {
	return o
}

func (o VdomexceptionOutput) ToVdomexceptionOutputWithContext(ctx context.Context) VdomexceptionOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o VdomexceptionOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vdomexception) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Index <1-4096>.
func (o VdomexceptionOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Vdomexception) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o VdomexceptionOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vdomexception) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Name of the configuration object that can be configured independently for all VDOMs.
func (o VdomexceptionOutput) Object() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomexception) pulumi.StringOutput { return v.Object }).(pulumi.StringOutput)
}

// Object ID.
func (o VdomexceptionOutput) Oid() pulumi.IntOutput {
	return o.ApplyT(func(v *Vdomexception) pulumi.IntOutput { return v.Oid }).(pulumi.IntOutput)
}

// Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration. Valid values: `all`, `inclusive`, `exclusive`.
func (o VdomexceptionOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomexception) pulumi.StringOutput { return v.Scope }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o VdomexceptionOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vdomexception) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Names of the VDOMs. The structure of `vdom` block is documented below.
func (o VdomexceptionOutput) Vdoms() VdomexceptionVdomArrayOutput {
	return o.ApplyT(func(v *Vdomexception) VdomexceptionVdomArrayOutput { return v.Vdoms }).(VdomexceptionVdomArrayOutput)
}

type VdomexceptionArrayOutput struct{ *pulumi.OutputState }

func (VdomexceptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vdomexception)(nil)).Elem()
}

func (o VdomexceptionArrayOutput) ToVdomexceptionArrayOutput() VdomexceptionArrayOutput {
	return o
}

func (o VdomexceptionArrayOutput) ToVdomexceptionArrayOutputWithContext(ctx context.Context) VdomexceptionArrayOutput {
	return o
}

func (o VdomexceptionArrayOutput) Index(i pulumi.IntInput) VdomexceptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vdomexception {
		return vs[0].([]*Vdomexception)[vs[1].(int)]
	}).(VdomexceptionOutput)
}

type VdomexceptionMapOutput struct{ *pulumi.OutputState }

func (VdomexceptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vdomexception)(nil)).Elem()
}

func (o VdomexceptionMapOutput) ToVdomexceptionMapOutput() VdomexceptionMapOutput {
	return o
}

func (o VdomexceptionMapOutput) ToVdomexceptionMapOutputWithContext(ctx context.Context) VdomexceptionMapOutput {
	return o
}

func (o VdomexceptionMapOutput) MapIndex(k pulumi.StringInput) VdomexceptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vdomexception {
		return vs[0].(map[string]*Vdomexception)[vs[1].(string)]
	}).(VdomexceptionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VdomexceptionInput)(nil)).Elem(), &Vdomexception{})
	pulumi.RegisterInputType(reflect.TypeOf((*VdomexceptionArrayInput)(nil)).Elem(), VdomexceptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VdomexceptionMapInput)(nil)).Elem(), VdomexceptionMap{})
	pulumi.RegisterOutputType(VdomexceptionOutput{})
	pulumi.RegisterOutputType(VdomexceptionArrayOutput{})
	pulumi.RegisterOutputType(VdomexceptionMapOutput{})
}
