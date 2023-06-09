// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Hidden table for datasource.
//
// ## Import
//
// # Waf Signature can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/wafSignature:WafSignature labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/wafSignature:WafSignature labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type WafSignature struct {
	pulumi.CustomResourceState

	// Signature description.
	Desc pulumi.StringOutput `pulumi:"desc"`
	// Signature ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWafSignature registers a new resource with the given unique name, arguments, and options.
func NewWafSignature(ctx *pulumi.Context,
	name string, args *WafSignatureArgs, opts ...pulumi.ResourceOption) (*WafSignature, error) {
	if args == nil {
		args = &WafSignatureArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WafSignature
	err := ctx.RegisterResource("fortios:index/wafSignature:WafSignature", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWafSignature gets an existing WafSignature resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWafSignature(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WafSignatureState, opts ...pulumi.ResourceOption) (*WafSignature, error) {
	var resource WafSignature
	err := ctx.ReadResource("fortios:index/wafSignature:WafSignature", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WafSignature resources.
type wafSignatureState struct {
	// Signature description.
	Desc *string `pulumi:"desc"`
	// Signature ID.
	Fosid *int `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WafSignatureState struct {
	// Signature description.
	Desc pulumi.StringPtrInput
	// Signature ID.
	Fosid pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WafSignatureState) ElementType() reflect.Type {
	return reflect.TypeOf((*wafSignatureState)(nil)).Elem()
}

type wafSignatureArgs struct {
	// Signature description.
	Desc *string `pulumi:"desc"`
	// Signature ID.
	Fosid *int `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WafSignature resource.
type WafSignatureArgs struct {
	// Signature description.
	Desc pulumi.StringPtrInput
	// Signature ID.
	Fosid pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WafSignatureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wafSignatureArgs)(nil)).Elem()
}

type WafSignatureInput interface {
	pulumi.Input

	ToWafSignatureOutput() WafSignatureOutput
	ToWafSignatureOutputWithContext(ctx context.Context) WafSignatureOutput
}

func (*WafSignature) ElementType() reflect.Type {
	return reflect.TypeOf((**WafSignature)(nil)).Elem()
}

func (i *WafSignature) ToWafSignatureOutput() WafSignatureOutput {
	return i.ToWafSignatureOutputWithContext(context.Background())
}

func (i *WafSignature) ToWafSignatureOutputWithContext(ctx context.Context) WafSignatureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WafSignatureOutput)
}

// WafSignatureArrayInput is an input type that accepts WafSignatureArray and WafSignatureArrayOutput values.
// You can construct a concrete instance of `WafSignatureArrayInput` via:
//
//	WafSignatureArray{ WafSignatureArgs{...} }
type WafSignatureArrayInput interface {
	pulumi.Input

	ToWafSignatureArrayOutput() WafSignatureArrayOutput
	ToWafSignatureArrayOutputWithContext(context.Context) WafSignatureArrayOutput
}

type WafSignatureArray []WafSignatureInput

func (WafSignatureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WafSignature)(nil)).Elem()
}

func (i WafSignatureArray) ToWafSignatureArrayOutput() WafSignatureArrayOutput {
	return i.ToWafSignatureArrayOutputWithContext(context.Background())
}

func (i WafSignatureArray) ToWafSignatureArrayOutputWithContext(ctx context.Context) WafSignatureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WafSignatureArrayOutput)
}

// WafSignatureMapInput is an input type that accepts WafSignatureMap and WafSignatureMapOutput values.
// You can construct a concrete instance of `WafSignatureMapInput` via:
//
//	WafSignatureMap{ "key": WafSignatureArgs{...} }
type WafSignatureMapInput interface {
	pulumi.Input

	ToWafSignatureMapOutput() WafSignatureMapOutput
	ToWafSignatureMapOutputWithContext(context.Context) WafSignatureMapOutput
}

type WafSignatureMap map[string]WafSignatureInput

func (WafSignatureMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WafSignature)(nil)).Elem()
}

func (i WafSignatureMap) ToWafSignatureMapOutput() WafSignatureMapOutput {
	return i.ToWafSignatureMapOutputWithContext(context.Background())
}

func (i WafSignatureMap) ToWafSignatureMapOutputWithContext(ctx context.Context) WafSignatureMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WafSignatureMapOutput)
}

type WafSignatureOutput struct{ *pulumi.OutputState }

func (WafSignatureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WafSignature)(nil)).Elem()
}

func (o WafSignatureOutput) ToWafSignatureOutput() WafSignatureOutput {
	return o
}

func (o WafSignatureOutput) ToWafSignatureOutputWithContext(ctx context.Context) WafSignatureOutput {
	return o
}

// Signature description.
func (o WafSignatureOutput) Desc() pulumi.StringOutput {
	return o.ApplyT(func(v *WafSignature) pulumi.StringOutput { return v.Desc }).(pulumi.StringOutput)
}

// Signature ID.
func (o WafSignatureOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *WafSignature) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o WafSignatureOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WafSignature) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WafSignatureArrayOutput struct{ *pulumi.OutputState }

func (WafSignatureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WafSignature)(nil)).Elem()
}

func (o WafSignatureArrayOutput) ToWafSignatureArrayOutput() WafSignatureArrayOutput {
	return o
}

func (o WafSignatureArrayOutput) ToWafSignatureArrayOutputWithContext(ctx context.Context) WafSignatureArrayOutput {
	return o
}

func (o WafSignatureArrayOutput) Index(i pulumi.IntInput) WafSignatureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WafSignature {
		return vs[0].([]*WafSignature)[vs[1].(int)]
	}).(WafSignatureOutput)
}

type WafSignatureMapOutput struct{ *pulumi.OutputState }

func (WafSignatureMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WafSignature)(nil)).Elem()
}

func (o WafSignatureMapOutput) ToWafSignatureMapOutput() WafSignatureMapOutput {
	return o
}

func (o WafSignatureMapOutput) ToWafSignatureMapOutputWithContext(ctx context.Context) WafSignatureMapOutput {
	return o
}

func (o WafSignatureMapOutput) MapIndex(k pulumi.StringInput) WafSignatureOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WafSignature {
		return vs[0].(map[string]*WafSignature)[vs[1].(string)]
	}).(WafSignatureOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WafSignatureInput)(nil)).Elem(), &WafSignature{})
	pulumi.RegisterInputType(reflect.TypeOf((*WafSignatureArrayInput)(nil)).Elem(), WafSignatureArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WafSignatureMapInput)(nil)).Elem(), WafSignatureMap{})
	pulumi.RegisterOutputType(WafSignatureOutput{})
	pulumi.RegisterOutputType(WafSignatureArrayOutput{})
	pulumi.RegisterOutputType(WafSignatureMapOutput{})
}
