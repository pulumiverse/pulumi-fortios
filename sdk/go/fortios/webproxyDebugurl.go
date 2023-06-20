// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure debug URL addresses.
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
//			_, err := fortios.NewWebproxyDebugurl(ctx, "trname", &fortios.WebproxyDebugurlArgs{
//				Exact:      pulumi.String("enable"),
//				Status:     pulumi.String("enable"),
//				UrlPattern: pulumi.String("/examples/servlet/*Servlet"),
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
// # WebProxy DebugUrl can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/webproxyDebugurl:WebproxyDebugurl labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/webproxyDebugurl:WebproxyDebugurl labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type WebproxyDebugurl struct {
	pulumi.CustomResourceState

	// Enable/disable matching the exact path. Valid values: `enable`, `disable`.
	Exact pulumi.StringOutput `pulumi:"exact"`
	// Debug URL name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable this URL exemption. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// URL exemption pattern.
	UrlPattern pulumi.StringOutput `pulumi:"urlPattern"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWebproxyDebugurl registers a new resource with the given unique name, arguments, and options.
func NewWebproxyDebugurl(ctx *pulumi.Context,
	name string, args *WebproxyDebugurlArgs, opts ...pulumi.ResourceOption) (*WebproxyDebugurl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.UrlPattern == nil {
		return nil, errors.New("invalid value for required argument 'UrlPattern'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource WebproxyDebugurl
	err := ctx.RegisterResource("fortios:index/webproxyDebugurl:WebproxyDebugurl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebproxyDebugurl gets an existing WebproxyDebugurl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebproxyDebugurl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebproxyDebugurlState, opts ...pulumi.ResourceOption) (*WebproxyDebugurl, error) {
	var resource WebproxyDebugurl
	err := ctx.ReadResource("fortios:index/webproxyDebugurl:WebproxyDebugurl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebproxyDebugurl resources.
type webproxyDebugurlState struct {
	// Enable/disable matching the exact path. Valid values: `enable`, `disable`.
	Exact *string `pulumi:"exact"`
	// Debug URL name.
	Name *string `pulumi:"name"`
	// Enable/disable this URL exemption. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// URL exemption pattern.
	UrlPattern *string `pulumi:"urlPattern"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WebproxyDebugurlState struct {
	// Enable/disable matching the exact path. Valid values: `enable`, `disable`.
	Exact pulumi.StringPtrInput
	// Debug URL name.
	Name pulumi.StringPtrInput
	// Enable/disable this URL exemption. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// URL exemption pattern.
	UrlPattern pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebproxyDebugurlState) ElementType() reflect.Type {
	return reflect.TypeOf((*webproxyDebugurlState)(nil)).Elem()
}

type webproxyDebugurlArgs struct {
	// Enable/disable matching the exact path. Valid values: `enable`, `disable`.
	Exact *string `pulumi:"exact"`
	// Debug URL name.
	Name *string `pulumi:"name"`
	// Enable/disable this URL exemption. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// URL exemption pattern.
	UrlPattern string `pulumi:"urlPattern"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WebproxyDebugurl resource.
type WebproxyDebugurlArgs struct {
	// Enable/disable matching the exact path. Valid values: `enable`, `disable`.
	Exact pulumi.StringPtrInput
	// Debug URL name.
	Name pulumi.StringPtrInput
	// Enable/disable this URL exemption. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// URL exemption pattern.
	UrlPattern pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebproxyDebugurlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webproxyDebugurlArgs)(nil)).Elem()
}

type WebproxyDebugurlInput interface {
	pulumi.Input

	ToWebproxyDebugurlOutput() WebproxyDebugurlOutput
	ToWebproxyDebugurlOutputWithContext(ctx context.Context) WebproxyDebugurlOutput
}

func (*WebproxyDebugurl) ElementType() reflect.Type {
	return reflect.TypeOf((**WebproxyDebugurl)(nil)).Elem()
}

func (i *WebproxyDebugurl) ToWebproxyDebugurlOutput() WebproxyDebugurlOutput {
	return i.ToWebproxyDebugurlOutputWithContext(context.Background())
}

func (i *WebproxyDebugurl) ToWebproxyDebugurlOutputWithContext(ctx context.Context) WebproxyDebugurlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebproxyDebugurlOutput)
}

// WebproxyDebugurlArrayInput is an input type that accepts WebproxyDebugurlArray and WebproxyDebugurlArrayOutput values.
// You can construct a concrete instance of `WebproxyDebugurlArrayInput` via:
//
//	WebproxyDebugurlArray{ WebproxyDebugurlArgs{...} }
type WebproxyDebugurlArrayInput interface {
	pulumi.Input

	ToWebproxyDebugurlArrayOutput() WebproxyDebugurlArrayOutput
	ToWebproxyDebugurlArrayOutputWithContext(context.Context) WebproxyDebugurlArrayOutput
}

type WebproxyDebugurlArray []WebproxyDebugurlInput

func (WebproxyDebugurlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebproxyDebugurl)(nil)).Elem()
}

func (i WebproxyDebugurlArray) ToWebproxyDebugurlArrayOutput() WebproxyDebugurlArrayOutput {
	return i.ToWebproxyDebugurlArrayOutputWithContext(context.Background())
}

func (i WebproxyDebugurlArray) ToWebproxyDebugurlArrayOutputWithContext(ctx context.Context) WebproxyDebugurlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebproxyDebugurlArrayOutput)
}

// WebproxyDebugurlMapInput is an input type that accepts WebproxyDebugurlMap and WebproxyDebugurlMapOutput values.
// You can construct a concrete instance of `WebproxyDebugurlMapInput` via:
//
//	WebproxyDebugurlMap{ "key": WebproxyDebugurlArgs{...} }
type WebproxyDebugurlMapInput interface {
	pulumi.Input

	ToWebproxyDebugurlMapOutput() WebproxyDebugurlMapOutput
	ToWebproxyDebugurlMapOutputWithContext(context.Context) WebproxyDebugurlMapOutput
}

type WebproxyDebugurlMap map[string]WebproxyDebugurlInput

func (WebproxyDebugurlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebproxyDebugurl)(nil)).Elem()
}

func (i WebproxyDebugurlMap) ToWebproxyDebugurlMapOutput() WebproxyDebugurlMapOutput {
	return i.ToWebproxyDebugurlMapOutputWithContext(context.Background())
}

func (i WebproxyDebugurlMap) ToWebproxyDebugurlMapOutputWithContext(ctx context.Context) WebproxyDebugurlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebproxyDebugurlMapOutput)
}

type WebproxyDebugurlOutput struct{ *pulumi.OutputState }

func (WebproxyDebugurlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebproxyDebugurl)(nil)).Elem()
}

func (o WebproxyDebugurlOutput) ToWebproxyDebugurlOutput() WebproxyDebugurlOutput {
	return o
}

func (o WebproxyDebugurlOutput) ToWebproxyDebugurlOutputWithContext(ctx context.Context) WebproxyDebugurlOutput {
	return o
}

// Enable/disable matching the exact path. Valid values: `enable`, `disable`.
func (o WebproxyDebugurlOutput) Exact() pulumi.StringOutput {
	return o.ApplyT(func(v *WebproxyDebugurl) pulumi.StringOutput { return v.Exact }).(pulumi.StringOutput)
}

// Debug URL name.
func (o WebproxyDebugurlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebproxyDebugurl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable this URL exemption. Valid values: `enable`, `disable`.
func (o WebproxyDebugurlOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *WebproxyDebugurl) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// URL exemption pattern.
func (o WebproxyDebugurlOutput) UrlPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *WebproxyDebugurl) pulumi.StringOutput { return v.UrlPattern }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o WebproxyDebugurlOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebproxyDebugurl) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WebproxyDebugurlArrayOutput struct{ *pulumi.OutputState }

func (WebproxyDebugurlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebproxyDebugurl)(nil)).Elem()
}

func (o WebproxyDebugurlArrayOutput) ToWebproxyDebugurlArrayOutput() WebproxyDebugurlArrayOutput {
	return o
}

func (o WebproxyDebugurlArrayOutput) ToWebproxyDebugurlArrayOutputWithContext(ctx context.Context) WebproxyDebugurlArrayOutput {
	return o
}

func (o WebproxyDebugurlArrayOutput) Index(i pulumi.IntInput) WebproxyDebugurlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebproxyDebugurl {
		return vs[0].([]*WebproxyDebugurl)[vs[1].(int)]
	}).(WebproxyDebugurlOutput)
}

type WebproxyDebugurlMapOutput struct{ *pulumi.OutputState }

func (WebproxyDebugurlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebproxyDebugurl)(nil)).Elem()
}

func (o WebproxyDebugurlMapOutput) ToWebproxyDebugurlMapOutput() WebproxyDebugurlMapOutput {
	return o
}

func (o WebproxyDebugurlMapOutput) ToWebproxyDebugurlMapOutputWithContext(ctx context.Context) WebproxyDebugurlMapOutput {
	return o
}

func (o WebproxyDebugurlMapOutput) MapIndex(k pulumi.StringInput) WebproxyDebugurlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebproxyDebugurl {
		return vs[0].(map[string]*WebproxyDebugurl)[vs[1].(string)]
	}).(WebproxyDebugurlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebproxyDebugurlInput)(nil)).Elem(), &WebproxyDebugurl{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebproxyDebugurlArrayInput)(nil)).Elem(), WebproxyDebugurlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebproxyDebugurlMapInput)(nil)).Elem(), WebproxyDebugurlMap{})
	pulumi.RegisterOutputType(WebproxyDebugurlOutput{})
	pulumi.RegisterOutputType(WebproxyDebugurlArrayOutput{})
	pulumi.RegisterOutputType(WebproxyDebugurlMapOutput{})
}