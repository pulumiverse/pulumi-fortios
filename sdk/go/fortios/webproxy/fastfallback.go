// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package webproxy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Proxy destination connection fast-fallback. Applies to FortiOS Version `>= 7.4.1`.
//
// ## Import
//
// WebProxy FastFallback can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:webproxy/fastfallback:Fastfallback labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:webproxy/fastfallback:Fastfallback labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Fastfallback struct {
	pulumi.CustomResourceState

	// Connection mode for multiple destinations. Valid values: `sequentially`, `simultaneously`.
	ConnectionMode pulumi.StringOutput `pulumi:"connectionMode"`
	// Number of milliseconds to wait before starting another connection (200 - 1800000, default = 200). For sequential connection-mode only.
	ConnectionTimeout pulumi.IntOutput `pulumi:"connectionTimeout"`
	// Configure a name for the fast-fallback entry.
	Name pulumi.StringOutput `pulumi:"name"`
	// Connection protocols for multiple destinations. Valid values: `IPv4-first`, `IPv6-first`, `IPv4-only`, `IPv6-only`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Enable/disable the fast-fallback entry. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewFastfallback registers a new resource with the given unique name, arguments, and options.
func NewFastfallback(ctx *pulumi.Context,
	name string, args *FastfallbackArgs, opts ...pulumi.ResourceOption) (*Fastfallback, error) {
	if args == nil {
		args = &FastfallbackArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Fastfallback
	err := ctx.RegisterResource("fortios:webproxy/fastfallback:Fastfallback", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFastfallback gets an existing Fastfallback resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFastfallback(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FastfallbackState, opts ...pulumi.ResourceOption) (*Fastfallback, error) {
	var resource Fastfallback
	err := ctx.ReadResource("fortios:webproxy/fastfallback:Fastfallback", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fastfallback resources.
type fastfallbackState struct {
	// Connection mode for multiple destinations. Valid values: `sequentially`, `simultaneously`.
	ConnectionMode *string `pulumi:"connectionMode"`
	// Number of milliseconds to wait before starting another connection (200 - 1800000, default = 200). For sequential connection-mode only.
	ConnectionTimeout *int `pulumi:"connectionTimeout"`
	// Configure a name for the fast-fallback entry.
	Name *string `pulumi:"name"`
	// Connection protocols for multiple destinations. Valid values: `IPv4-first`, `IPv6-first`, `IPv4-only`, `IPv6-only`.
	Protocol *string `pulumi:"protocol"`
	// Enable/disable the fast-fallback entry. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FastfallbackState struct {
	// Connection mode for multiple destinations. Valid values: `sequentially`, `simultaneously`.
	ConnectionMode pulumi.StringPtrInput
	// Number of milliseconds to wait before starting another connection (200 - 1800000, default = 200). For sequential connection-mode only.
	ConnectionTimeout pulumi.IntPtrInput
	// Configure a name for the fast-fallback entry.
	Name pulumi.StringPtrInput
	// Connection protocols for multiple destinations. Valid values: `IPv4-first`, `IPv6-first`, `IPv4-only`, `IPv6-only`.
	Protocol pulumi.StringPtrInput
	// Enable/disable the fast-fallback entry. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FastfallbackState) ElementType() reflect.Type {
	return reflect.TypeOf((*fastfallbackState)(nil)).Elem()
}

type fastfallbackArgs struct {
	// Connection mode for multiple destinations. Valid values: `sequentially`, `simultaneously`.
	ConnectionMode *string `pulumi:"connectionMode"`
	// Number of milliseconds to wait before starting another connection (200 - 1800000, default = 200). For sequential connection-mode only.
	ConnectionTimeout *int `pulumi:"connectionTimeout"`
	// Configure a name for the fast-fallback entry.
	Name *string `pulumi:"name"`
	// Connection protocols for multiple destinations. Valid values: `IPv4-first`, `IPv6-first`, `IPv4-only`, `IPv6-only`.
	Protocol *string `pulumi:"protocol"`
	// Enable/disable the fast-fallback entry. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Fastfallback resource.
type FastfallbackArgs struct {
	// Connection mode for multiple destinations. Valid values: `sequentially`, `simultaneously`.
	ConnectionMode pulumi.StringPtrInput
	// Number of milliseconds to wait before starting another connection (200 - 1800000, default = 200). For sequential connection-mode only.
	ConnectionTimeout pulumi.IntPtrInput
	// Configure a name for the fast-fallback entry.
	Name pulumi.StringPtrInput
	// Connection protocols for multiple destinations. Valid values: `IPv4-first`, `IPv6-first`, `IPv4-only`, `IPv6-only`.
	Protocol pulumi.StringPtrInput
	// Enable/disable the fast-fallback entry. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FastfallbackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fastfallbackArgs)(nil)).Elem()
}

type FastfallbackInput interface {
	pulumi.Input

	ToFastfallbackOutput() FastfallbackOutput
	ToFastfallbackOutputWithContext(ctx context.Context) FastfallbackOutput
}

func (*Fastfallback) ElementType() reflect.Type {
	return reflect.TypeOf((**Fastfallback)(nil)).Elem()
}

func (i *Fastfallback) ToFastfallbackOutput() FastfallbackOutput {
	return i.ToFastfallbackOutputWithContext(context.Background())
}

func (i *Fastfallback) ToFastfallbackOutputWithContext(ctx context.Context) FastfallbackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastfallbackOutput)
}

// FastfallbackArrayInput is an input type that accepts FastfallbackArray and FastfallbackArrayOutput values.
// You can construct a concrete instance of `FastfallbackArrayInput` via:
//
//	FastfallbackArray{ FastfallbackArgs{...} }
type FastfallbackArrayInput interface {
	pulumi.Input

	ToFastfallbackArrayOutput() FastfallbackArrayOutput
	ToFastfallbackArrayOutputWithContext(context.Context) FastfallbackArrayOutput
}

type FastfallbackArray []FastfallbackInput

func (FastfallbackArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fastfallback)(nil)).Elem()
}

func (i FastfallbackArray) ToFastfallbackArrayOutput() FastfallbackArrayOutput {
	return i.ToFastfallbackArrayOutputWithContext(context.Background())
}

func (i FastfallbackArray) ToFastfallbackArrayOutputWithContext(ctx context.Context) FastfallbackArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastfallbackArrayOutput)
}

// FastfallbackMapInput is an input type that accepts FastfallbackMap and FastfallbackMapOutput values.
// You can construct a concrete instance of `FastfallbackMapInput` via:
//
//	FastfallbackMap{ "key": FastfallbackArgs{...} }
type FastfallbackMapInput interface {
	pulumi.Input

	ToFastfallbackMapOutput() FastfallbackMapOutput
	ToFastfallbackMapOutputWithContext(context.Context) FastfallbackMapOutput
}

type FastfallbackMap map[string]FastfallbackInput

func (FastfallbackMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fastfallback)(nil)).Elem()
}

func (i FastfallbackMap) ToFastfallbackMapOutput() FastfallbackMapOutput {
	return i.ToFastfallbackMapOutputWithContext(context.Background())
}

func (i FastfallbackMap) ToFastfallbackMapOutputWithContext(ctx context.Context) FastfallbackMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastfallbackMapOutput)
}

type FastfallbackOutput struct{ *pulumi.OutputState }

func (FastfallbackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fastfallback)(nil)).Elem()
}

func (o FastfallbackOutput) ToFastfallbackOutput() FastfallbackOutput {
	return o
}

func (o FastfallbackOutput) ToFastfallbackOutputWithContext(ctx context.Context) FastfallbackOutput {
	return o
}

// Connection mode for multiple destinations. Valid values: `sequentially`, `simultaneously`.
func (o FastfallbackOutput) ConnectionMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Fastfallback) pulumi.StringOutput { return v.ConnectionMode }).(pulumi.StringOutput)
}

// Number of milliseconds to wait before starting another connection (200 - 1800000, default = 200). For sequential connection-mode only.
func (o FastfallbackOutput) ConnectionTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Fastfallback) pulumi.IntOutput { return v.ConnectionTimeout }).(pulumi.IntOutput)
}

// Configure a name for the fast-fallback entry.
func (o FastfallbackOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Fastfallback) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Connection protocols for multiple destinations. Valid values: `IPv4-first`, `IPv6-first`, `IPv4-only`, `IPv6-only`.
func (o FastfallbackOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Fastfallback) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// Enable/disable the fast-fallback entry. Valid values: `enable`, `disable`.
func (o FastfallbackOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Fastfallback) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FastfallbackOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Fastfallback) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type FastfallbackArrayOutput struct{ *pulumi.OutputState }

func (FastfallbackArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fastfallback)(nil)).Elem()
}

func (o FastfallbackArrayOutput) ToFastfallbackArrayOutput() FastfallbackArrayOutput {
	return o
}

func (o FastfallbackArrayOutput) ToFastfallbackArrayOutputWithContext(ctx context.Context) FastfallbackArrayOutput {
	return o
}

func (o FastfallbackArrayOutput) Index(i pulumi.IntInput) FastfallbackOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Fastfallback {
		return vs[0].([]*Fastfallback)[vs[1].(int)]
	}).(FastfallbackOutput)
}

type FastfallbackMapOutput struct{ *pulumi.OutputState }

func (FastfallbackMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fastfallback)(nil)).Elem()
}

func (o FastfallbackMapOutput) ToFastfallbackMapOutput() FastfallbackMapOutput {
	return o
}

func (o FastfallbackMapOutput) ToFastfallbackMapOutputWithContext(ctx context.Context) FastfallbackMapOutput {
	return o
}

func (o FastfallbackMapOutput) MapIndex(k pulumi.StringInput) FastfallbackOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Fastfallback {
		return vs[0].(map[string]*Fastfallback)[vs[1].(string)]
	}).(FastfallbackOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FastfallbackInput)(nil)).Elem(), &Fastfallback{})
	pulumi.RegisterInputType(reflect.TypeOf((*FastfallbackArrayInput)(nil)).Elem(), FastfallbackArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FastfallbackMapInput)(nil)).Elem(), FastfallbackMap{})
	pulumi.RegisterOutputType(FastfallbackOutput{})
	pulumi.RegisterOutputType(FastfallbackArrayOutput{})
	pulumi.RegisterOutputType(FastfallbackMapOutput{})
}
