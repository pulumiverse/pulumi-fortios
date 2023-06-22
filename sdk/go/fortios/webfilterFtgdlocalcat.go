// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiGuard Web Filter local categories.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fortios.NewWebfilterFtgdlocalcat(ctx, "trname", &fortios.WebfilterFtgdlocalcatArgs{
//				Desc:   pulumi.String("s1"),
//				Fosid:  pulumi.Int(188),
//				Status: pulumi.String("enable"),
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
// # Webfilter FtgdLocalCat can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/webfilterFtgdlocalcat:WebfilterFtgdlocalcat labelname {{desc}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/webfilterFtgdlocalcat:WebfilterFtgdlocalcat labelname {{desc}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type WebfilterFtgdlocalcat struct {
	pulumi.CustomResourceState

	// Local category description.
	Desc pulumi.StringOutput `pulumi:"desc"`
	// Local category ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Enable/disable the local category. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWebfilterFtgdlocalcat registers a new resource with the given unique name, arguments, and options.
func NewWebfilterFtgdlocalcat(ctx *pulumi.Context,
	name string, args *WebfilterFtgdlocalcatArgs, opts ...pulumi.ResourceOption) (*WebfilterFtgdlocalcat, error) {
	if args == nil {
		args = &WebfilterFtgdlocalcatArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WebfilterFtgdlocalcat
	err := ctx.RegisterResource("fortios:index/webfilterFtgdlocalcat:WebfilterFtgdlocalcat", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebfilterFtgdlocalcat gets an existing WebfilterFtgdlocalcat resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebfilterFtgdlocalcat(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebfilterFtgdlocalcatState, opts ...pulumi.ResourceOption) (*WebfilterFtgdlocalcat, error) {
	var resource WebfilterFtgdlocalcat
	err := ctx.ReadResource("fortios:index/webfilterFtgdlocalcat:WebfilterFtgdlocalcat", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebfilterFtgdlocalcat resources.
type webfilterFtgdlocalcatState struct {
	// Local category description.
	Desc *string `pulumi:"desc"`
	// Local category ID.
	Fosid *int `pulumi:"fosid"`
	// Enable/disable the local category. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WebfilterFtgdlocalcatState struct {
	// Local category description.
	Desc pulumi.StringPtrInput
	// Local category ID.
	Fosid pulumi.IntPtrInput
	// Enable/disable the local category. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebfilterFtgdlocalcatState) ElementType() reflect.Type {
	return reflect.TypeOf((*webfilterFtgdlocalcatState)(nil)).Elem()
}

type webfilterFtgdlocalcatArgs struct {
	// Local category description.
	Desc *string `pulumi:"desc"`
	// Local category ID.
	Fosid *int `pulumi:"fosid"`
	// Enable/disable the local category. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WebfilterFtgdlocalcat resource.
type WebfilterFtgdlocalcatArgs struct {
	// Local category description.
	Desc pulumi.StringPtrInput
	// Local category ID.
	Fosid pulumi.IntPtrInput
	// Enable/disable the local category. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebfilterFtgdlocalcatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webfilterFtgdlocalcatArgs)(nil)).Elem()
}

type WebfilterFtgdlocalcatInput interface {
	pulumi.Input

	ToWebfilterFtgdlocalcatOutput() WebfilterFtgdlocalcatOutput
	ToWebfilterFtgdlocalcatOutputWithContext(ctx context.Context) WebfilterFtgdlocalcatOutput
}

func (*WebfilterFtgdlocalcat) ElementType() reflect.Type {
	return reflect.TypeOf((**WebfilterFtgdlocalcat)(nil)).Elem()
}

func (i *WebfilterFtgdlocalcat) ToWebfilterFtgdlocalcatOutput() WebfilterFtgdlocalcatOutput {
	return i.ToWebfilterFtgdlocalcatOutputWithContext(context.Background())
}

func (i *WebfilterFtgdlocalcat) ToWebfilterFtgdlocalcatOutputWithContext(ctx context.Context) WebfilterFtgdlocalcatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterFtgdlocalcatOutput)
}

// WebfilterFtgdlocalcatArrayInput is an input type that accepts WebfilterFtgdlocalcatArray and WebfilterFtgdlocalcatArrayOutput values.
// You can construct a concrete instance of `WebfilterFtgdlocalcatArrayInput` via:
//
//	WebfilterFtgdlocalcatArray{ WebfilterFtgdlocalcatArgs{...} }
type WebfilterFtgdlocalcatArrayInput interface {
	pulumi.Input

	ToWebfilterFtgdlocalcatArrayOutput() WebfilterFtgdlocalcatArrayOutput
	ToWebfilterFtgdlocalcatArrayOutputWithContext(context.Context) WebfilterFtgdlocalcatArrayOutput
}

type WebfilterFtgdlocalcatArray []WebfilterFtgdlocalcatInput

func (WebfilterFtgdlocalcatArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebfilterFtgdlocalcat)(nil)).Elem()
}

func (i WebfilterFtgdlocalcatArray) ToWebfilterFtgdlocalcatArrayOutput() WebfilterFtgdlocalcatArrayOutput {
	return i.ToWebfilterFtgdlocalcatArrayOutputWithContext(context.Background())
}

func (i WebfilterFtgdlocalcatArray) ToWebfilterFtgdlocalcatArrayOutputWithContext(ctx context.Context) WebfilterFtgdlocalcatArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterFtgdlocalcatArrayOutput)
}

// WebfilterFtgdlocalcatMapInput is an input type that accepts WebfilterFtgdlocalcatMap and WebfilterFtgdlocalcatMapOutput values.
// You can construct a concrete instance of `WebfilterFtgdlocalcatMapInput` via:
//
//	WebfilterFtgdlocalcatMap{ "key": WebfilterFtgdlocalcatArgs{...} }
type WebfilterFtgdlocalcatMapInput interface {
	pulumi.Input

	ToWebfilterFtgdlocalcatMapOutput() WebfilterFtgdlocalcatMapOutput
	ToWebfilterFtgdlocalcatMapOutputWithContext(context.Context) WebfilterFtgdlocalcatMapOutput
}

type WebfilterFtgdlocalcatMap map[string]WebfilterFtgdlocalcatInput

func (WebfilterFtgdlocalcatMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebfilterFtgdlocalcat)(nil)).Elem()
}

func (i WebfilterFtgdlocalcatMap) ToWebfilterFtgdlocalcatMapOutput() WebfilterFtgdlocalcatMapOutput {
	return i.ToWebfilterFtgdlocalcatMapOutputWithContext(context.Background())
}

func (i WebfilterFtgdlocalcatMap) ToWebfilterFtgdlocalcatMapOutputWithContext(ctx context.Context) WebfilterFtgdlocalcatMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterFtgdlocalcatMapOutput)
}

type WebfilterFtgdlocalcatOutput struct{ *pulumi.OutputState }

func (WebfilterFtgdlocalcatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebfilterFtgdlocalcat)(nil)).Elem()
}

func (o WebfilterFtgdlocalcatOutput) ToWebfilterFtgdlocalcatOutput() WebfilterFtgdlocalcatOutput {
	return o
}

func (o WebfilterFtgdlocalcatOutput) ToWebfilterFtgdlocalcatOutputWithContext(ctx context.Context) WebfilterFtgdlocalcatOutput {
	return o
}

// Local category description.
func (o WebfilterFtgdlocalcatOutput) Desc() pulumi.StringOutput {
	return o.ApplyT(func(v *WebfilterFtgdlocalcat) pulumi.StringOutput { return v.Desc }).(pulumi.StringOutput)
}

// Local category ID.
func (o WebfilterFtgdlocalcatOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *WebfilterFtgdlocalcat) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Enable/disable the local category. Valid values: `enable`, `disable`.
func (o WebfilterFtgdlocalcatOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *WebfilterFtgdlocalcat) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o WebfilterFtgdlocalcatOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebfilterFtgdlocalcat) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WebfilterFtgdlocalcatArrayOutput struct{ *pulumi.OutputState }

func (WebfilterFtgdlocalcatArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebfilterFtgdlocalcat)(nil)).Elem()
}

func (o WebfilterFtgdlocalcatArrayOutput) ToWebfilterFtgdlocalcatArrayOutput() WebfilterFtgdlocalcatArrayOutput {
	return o
}

func (o WebfilterFtgdlocalcatArrayOutput) ToWebfilterFtgdlocalcatArrayOutputWithContext(ctx context.Context) WebfilterFtgdlocalcatArrayOutput {
	return o
}

func (o WebfilterFtgdlocalcatArrayOutput) Index(i pulumi.IntInput) WebfilterFtgdlocalcatOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebfilterFtgdlocalcat {
		return vs[0].([]*WebfilterFtgdlocalcat)[vs[1].(int)]
	}).(WebfilterFtgdlocalcatOutput)
}

type WebfilterFtgdlocalcatMapOutput struct{ *pulumi.OutputState }

func (WebfilterFtgdlocalcatMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebfilterFtgdlocalcat)(nil)).Elem()
}

func (o WebfilterFtgdlocalcatMapOutput) ToWebfilterFtgdlocalcatMapOutput() WebfilterFtgdlocalcatMapOutput {
	return o
}

func (o WebfilterFtgdlocalcatMapOutput) ToWebfilterFtgdlocalcatMapOutputWithContext(ctx context.Context) WebfilterFtgdlocalcatMapOutput {
	return o
}

func (o WebfilterFtgdlocalcatMapOutput) MapIndex(k pulumi.StringInput) WebfilterFtgdlocalcatOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebfilterFtgdlocalcat {
		return vs[0].(map[string]*WebfilterFtgdlocalcat)[vs[1].(string)]
	}).(WebfilterFtgdlocalcatOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebfilterFtgdlocalcatInput)(nil)).Elem(), &WebfilterFtgdlocalcat{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebfilterFtgdlocalcatArrayInput)(nil)).Elem(), WebfilterFtgdlocalcatArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebfilterFtgdlocalcatMapInput)(nil)).Elem(), WebfilterFtgdlocalcatMap{})
	pulumi.RegisterOutputType(WebfilterFtgdlocalcatOutput{})
	pulumi.RegisterOutputType(WebfilterFtgdlocalcatArrayOutput{})
	pulumi.RegisterOutputType(WebfilterFtgdlocalcatMapOutput{})
}
