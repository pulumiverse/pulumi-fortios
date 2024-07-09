// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package application

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure custom application signatures.
//
// ## Import
//
// Application Custom can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:application/custom:Custom labelname {{tag}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:application/custom:Custom labelname {{tag}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Custom struct {
	pulumi.CustomResourceState

	// Custom application signature behavior.
	Behavior pulumi.StringOutput `pulumi:"behavior"`
	// Custom application category ID (use ? to view available options).
	Category pulumi.IntOutput `pulumi:"category"`
	// Comment.
	Comment pulumi.StringOutput `pulumi:"comment"`
	// Custom application category ID (use ? to view available options).
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name of this custom application signature.
	Name pulumi.StringOutput `pulumi:"name"`
	// Custom application signature protocol.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The text that makes up the actual custom application signature.
	Signature pulumi.StringOutput `pulumi:"signature"`
	// Signature tag.
	Tag pulumi.StringOutput `pulumi:"tag"`
	// Custom application signature technology.
	Technology pulumi.StringOutput `pulumi:"technology"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
	// Custom application signature vendor.
	Vendor pulumi.StringOutput `pulumi:"vendor"`
}

// NewCustom registers a new resource with the given unique name, arguments, and options.
func NewCustom(ctx *pulumi.Context,
	name string, args *CustomArgs, opts ...pulumi.ResourceOption) (*Custom, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Category == nil {
		return nil, errors.New("invalid value for required argument 'Category'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Custom
	err := ctx.RegisterResource("fortios:application/custom:Custom", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustom gets an existing Custom resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomState, opts ...pulumi.ResourceOption) (*Custom, error) {
	var resource Custom
	err := ctx.ReadResource("fortios:application/custom:Custom", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Custom resources.
type customState struct {
	// Custom application signature behavior.
	Behavior *string `pulumi:"behavior"`
	// Custom application category ID (use ? to view available options).
	Category *int `pulumi:"category"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Custom application category ID (use ? to view available options).
	Fosid *int `pulumi:"fosid"`
	// Name of this custom application signature.
	Name *string `pulumi:"name"`
	// Custom application signature protocol.
	Protocol *string `pulumi:"protocol"`
	// The text that makes up the actual custom application signature.
	Signature *string `pulumi:"signature"`
	// Signature tag.
	Tag *string `pulumi:"tag"`
	// Custom application signature technology.
	Technology *string `pulumi:"technology"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Custom application signature vendor.
	Vendor *string `pulumi:"vendor"`
}

type CustomState struct {
	// Custom application signature behavior.
	Behavior pulumi.StringPtrInput
	// Custom application category ID (use ? to view available options).
	Category pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Custom application category ID (use ? to view available options).
	Fosid pulumi.IntPtrInput
	// Name of this custom application signature.
	Name pulumi.StringPtrInput
	// Custom application signature protocol.
	Protocol pulumi.StringPtrInput
	// The text that makes up the actual custom application signature.
	Signature pulumi.StringPtrInput
	// Signature tag.
	Tag pulumi.StringPtrInput
	// Custom application signature technology.
	Technology pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Custom application signature vendor.
	Vendor pulumi.StringPtrInput
}

func (CustomState) ElementType() reflect.Type {
	return reflect.TypeOf((*customState)(nil)).Elem()
}

type customArgs struct {
	// Custom application signature behavior.
	Behavior *string `pulumi:"behavior"`
	// Custom application category ID (use ? to view available options).
	Category int `pulumi:"category"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Custom application category ID (use ? to view available options).
	Fosid *int `pulumi:"fosid"`
	// Name of this custom application signature.
	Name *string `pulumi:"name"`
	// Custom application signature protocol.
	Protocol *string `pulumi:"protocol"`
	// The text that makes up the actual custom application signature.
	Signature *string `pulumi:"signature"`
	// Signature tag.
	Tag *string `pulumi:"tag"`
	// Custom application signature technology.
	Technology *string `pulumi:"technology"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Custom application signature vendor.
	Vendor *string `pulumi:"vendor"`
}

// The set of arguments for constructing a Custom resource.
type CustomArgs struct {
	// Custom application signature behavior.
	Behavior pulumi.StringPtrInput
	// Custom application category ID (use ? to view available options).
	Category pulumi.IntInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Custom application category ID (use ? to view available options).
	Fosid pulumi.IntPtrInput
	// Name of this custom application signature.
	Name pulumi.StringPtrInput
	// Custom application signature protocol.
	Protocol pulumi.StringPtrInput
	// The text that makes up the actual custom application signature.
	Signature pulumi.StringPtrInput
	// Signature tag.
	Tag pulumi.StringPtrInput
	// Custom application signature technology.
	Technology pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Custom application signature vendor.
	Vendor pulumi.StringPtrInput
}

func (CustomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customArgs)(nil)).Elem()
}

type CustomInput interface {
	pulumi.Input

	ToCustomOutput() CustomOutput
	ToCustomOutputWithContext(ctx context.Context) CustomOutput
}

func (*Custom) ElementType() reflect.Type {
	return reflect.TypeOf((**Custom)(nil)).Elem()
}

func (i *Custom) ToCustomOutput() CustomOutput {
	return i.ToCustomOutputWithContext(context.Background())
}

func (i *Custom) ToCustomOutputWithContext(ctx context.Context) CustomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomOutput)
}

// CustomArrayInput is an input type that accepts CustomArray and CustomArrayOutput values.
// You can construct a concrete instance of `CustomArrayInput` via:
//
//	CustomArray{ CustomArgs{...} }
type CustomArrayInput interface {
	pulumi.Input

	ToCustomArrayOutput() CustomArrayOutput
	ToCustomArrayOutputWithContext(context.Context) CustomArrayOutput
}

type CustomArray []CustomInput

func (CustomArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Custom)(nil)).Elem()
}

func (i CustomArray) ToCustomArrayOutput() CustomArrayOutput {
	return i.ToCustomArrayOutputWithContext(context.Background())
}

func (i CustomArray) ToCustomArrayOutputWithContext(ctx context.Context) CustomArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomArrayOutput)
}

// CustomMapInput is an input type that accepts CustomMap and CustomMapOutput values.
// You can construct a concrete instance of `CustomMapInput` via:
//
//	CustomMap{ "key": CustomArgs{...} }
type CustomMapInput interface {
	pulumi.Input

	ToCustomMapOutput() CustomMapOutput
	ToCustomMapOutputWithContext(context.Context) CustomMapOutput
}

type CustomMap map[string]CustomInput

func (CustomMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Custom)(nil)).Elem()
}

func (i CustomMap) ToCustomMapOutput() CustomMapOutput {
	return i.ToCustomMapOutputWithContext(context.Background())
}

func (i CustomMap) ToCustomMapOutputWithContext(ctx context.Context) CustomMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomMapOutput)
}

type CustomOutput struct{ *pulumi.OutputState }

func (CustomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Custom)(nil)).Elem()
}

func (o CustomOutput) ToCustomOutput() CustomOutput {
	return o
}

func (o CustomOutput) ToCustomOutputWithContext(ctx context.Context) CustomOutput {
	return o
}

// Custom application signature behavior.
func (o CustomOutput) Behavior() pulumi.StringOutput {
	return o.ApplyT(func(v *Custom) pulumi.StringOutput { return v.Behavior }).(pulumi.StringOutput)
}

// Custom application category ID (use ? to view available options).
func (o CustomOutput) Category() pulumi.IntOutput {
	return o.ApplyT(func(v *Custom) pulumi.IntOutput { return v.Category }).(pulumi.IntOutput)
}

// Comment.
func (o CustomOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v *Custom) pulumi.StringOutput { return v.Comment }).(pulumi.StringOutput)
}

// Custom application category ID (use ? to view available options).
func (o CustomOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Custom) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Name of this custom application signature.
func (o CustomOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Custom) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Custom application signature protocol.
func (o CustomOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Custom) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The text that makes up the actual custom application signature.
func (o CustomOutput) Signature() pulumi.StringOutput {
	return o.ApplyT(func(v *Custom) pulumi.StringOutput { return v.Signature }).(pulumi.StringOutput)
}

// Signature tag.
func (o CustomOutput) Tag() pulumi.StringOutput {
	return o.ApplyT(func(v *Custom) pulumi.StringOutput { return v.Tag }).(pulumi.StringOutput)
}

// Custom application signature technology.
func (o CustomOutput) Technology() pulumi.StringOutput {
	return o.ApplyT(func(v *Custom) pulumi.StringOutput { return v.Technology }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o CustomOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Custom) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

// Custom application signature vendor.
func (o CustomOutput) Vendor() pulumi.StringOutput {
	return o.ApplyT(func(v *Custom) pulumi.StringOutput { return v.Vendor }).(pulumi.StringOutput)
}

type CustomArrayOutput struct{ *pulumi.OutputState }

func (CustomArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Custom)(nil)).Elem()
}

func (o CustomArrayOutput) ToCustomArrayOutput() CustomArrayOutput {
	return o
}

func (o CustomArrayOutput) ToCustomArrayOutputWithContext(ctx context.Context) CustomArrayOutput {
	return o
}

func (o CustomArrayOutput) Index(i pulumi.IntInput) CustomOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Custom {
		return vs[0].([]*Custom)[vs[1].(int)]
	}).(CustomOutput)
}

type CustomMapOutput struct{ *pulumi.OutputState }

func (CustomMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Custom)(nil)).Elem()
}

func (o CustomMapOutput) ToCustomMapOutput() CustomMapOutput {
	return o
}

func (o CustomMapOutput) ToCustomMapOutputWithContext(ctx context.Context) CustomMapOutput {
	return o
}

func (o CustomMapOutput) MapIndex(k pulumi.StringInput) CustomOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Custom {
		return vs[0].(map[string]*Custom)[vs[1].(string)]
	}).(CustomOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomInput)(nil)).Elem(), &Custom{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomArrayInput)(nil)).Elem(), CustomArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomMapInput)(nil)).Elem(), CustomMap{})
	pulumi.RegisterOutputType(CustomOutput{})
	pulumi.RegisterOutputType(CustomArrayOutput{})
	pulumi.RegisterOutputType(CustomMapOutput{})
}
