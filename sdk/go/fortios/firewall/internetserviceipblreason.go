// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// IP blacklist reason. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// Firewall InternetServiceIpblReason can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/internetserviceipblreason:Internetserviceipblreason labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/internetserviceipblreason:Internetserviceipblreason labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Internetserviceipblreason struct {
	pulumi.CustomResourceState

	// IP blacklist reason ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// IP blacklist reason name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewInternetserviceipblreason registers a new resource with the given unique name, arguments, and options.
func NewInternetserviceipblreason(ctx *pulumi.Context,
	name string, args *InternetserviceipblreasonArgs, opts ...pulumi.ResourceOption) (*Internetserviceipblreason, error) {
	if args == nil {
		args = &InternetserviceipblreasonArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Internetserviceipblreason
	err := ctx.RegisterResource("fortios:firewall/internetserviceipblreason:Internetserviceipblreason", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInternetserviceipblreason gets an existing Internetserviceipblreason resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInternetserviceipblreason(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InternetserviceipblreasonState, opts ...pulumi.ResourceOption) (*Internetserviceipblreason, error) {
	var resource Internetserviceipblreason
	err := ctx.ReadResource("fortios:firewall/internetserviceipblreason:Internetserviceipblreason", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Internetserviceipblreason resources.
type internetserviceipblreasonState struct {
	// IP blacklist reason ID.
	Fosid *int `pulumi:"fosid"`
	// IP blacklist reason name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type InternetserviceipblreasonState struct {
	// IP blacklist reason ID.
	Fosid pulumi.IntPtrInput
	// IP blacklist reason name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (InternetserviceipblreasonState) ElementType() reflect.Type {
	return reflect.TypeOf((*internetserviceipblreasonState)(nil)).Elem()
}

type internetserviceipblreasonArgs struct {
	// IP blacklist reason ID.
	Fosid *int `pulumi:"fosid"`
	// IP blacklist reason name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Internetserviceipblreason resource.
type InternetserviceipblreasonArgs struct {
	// IP blacklist reason ID.
	Fosid pulumi.IntPtrInput
	// IP blacklist reason name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (InternetserviceipblreasonArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*internetserviceipblreasonArgs)(nil)).Elem()
}

type InternetserviceipblreasonInput interface {
	pulumi.Input

	ToInternetserviceipblreasonOutput() InternetserviceipblreasonOutput
	ToInternetserviceipblreasonOutputWithContext(ctx context.Context) InternetserviceipblreasonOutput
}

func (*Internetserviceipblreason) ElementType() reflect.Type {
	return reflect.TypeOf((**Internetserviceipblreason)(nil)).Elem()
}

func (i *Internetserviceipblreason) ToInternetserviceipblreasonOutput() InternetserviceipblreasonOutput {
	return i.ToInternetserviceipblreasonOutputWithContext(context.Background())
}

func (i *Internetserviceipblreason) ToInternetserviceipblreasonOutputWithContext(ctx context.Context) InternetserviceipblreasonOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetserviceipblreasonOutput)
}

// InternetserviceipblreasonArrayInput is an input type that accepts InternetserviceipblreasonArray and InternetserviceipblreasonArrayOutput values.
// You can construct a concrete instance of `InternetserviceipblreasonArrayInput` via:
//
//	InternetserviceipblreasonArray{ InternetserviceipblreasonArgs{...} }
type InternetserviceipblreasonArrayInput interface {
	pulumi.Input

	ToInternetserviceipblreasonArrayOutput() InternetserviceipblreasonArrayOutput
	ToInternetserviceipblreasonArrayOutputWithContext(context.Context) InternetserviceipblreasonArrayOutput
}

type InternetserviceipblreasonArray []InternetserviceipblreasonInput

func (InternetserviceipblreasonArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Internetserviceipblreason)(nil)).Elem()
}

func (i InternetserviceipblreasonArray) ToInternetserviceipblreasonArrayOutput() InternetserviceipblreasonArrayOutput {
	return i.ToInternetserviceipblreasonArrayOutputWithContext(context.Background())
}

func (i InternetserviceipblreasonArray) ToInternetserviceipblreasonArrayOutputWithContext(ctx context.Context) InternetserviceipblreasonArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetserviceipblreasonArrayOutput)
}

// InternetserviceipblreasonMapInput is an input type that accepts InternetserviceipblreasonMap and InternetserviceipblreasonMapOutput values.
// You can construct a concrete instance of `InternetserviceipblreasonMapInput` via:
//
//	InternetserviceipblreasonMap{ "key": InternetserviceipblreasonArgs{...} }
type InternetserviceipblreasonMapInput interface {
	pulumi.Input

	ToInternetserviceipblreasonMapOutput() InternetserviceipblreasonMapOutput
	ToInternetserviceipblreasonMapOutputWithContext(context.Context) InternetserviceipblreasonMapOutput
}

type InternetserviceipblreasonMap map[string]InternetserviceipblreasonInput

func (InternetserviceipblreasonMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Internetserviceipblreason)(nil)).Elem()
}

func (i InternetserviceipblreasonMap) ToInternetserviceipblreasonMapOutput() InternetserviceipblreasonMapOutput {
	return i.ToInternetserviceipblreasonMapOutputWithContext(context.Background())
}

func (i InternetserviceipblreasonMap) ToInternetserviceipblreasonMapOutputWithContext(ctx context.Context) InternetserviceipblreasonMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetserviceipblreasonMapOutput)
}

type InternetserviceipblreasonOutput struct{ *pulumi.OutputState }

func (InternetserviceipblreasonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Internetserviceipblreason)(nil)).Elem()
}

func (o InternetserviceipblreasonOutput) ToInternetserviceipblreasonOutput() InternetserviceipblreasonOutput {
	return o
}

func (o InternetserviceipblreasonOutput) ToInternetserviceipblreasonOutputWithContext(ctx context.Context) InternetserviceipblreasonOutput {
	return o
}

// IP blacklist reason ID.
func (o InternetserviceipblreasonOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetserviceipblreason) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// IP blacklist reason name.
func (o InternetserviceipblreasonOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Internetserviceipblreason) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o InternetserviceipblreasonOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Internetserviceipblreason) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type InternetserviceipblreasonArrayOutput struct{ *pulumi.OutputState }

func (InternetserviceipblreasonArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Internetserviceipblreason)(nil)).Elem()
}

func (o InternetserviceipblreasonArrayOutput) ToInternetserviceipblreasonArrayOutput() InternetserviceipblreasonArrayOutput {
	return o
}

func (o InternetserviceipblreasonArrayOutput) ToInternetserviceipblreasonArrayOutputWithContext(ctx context.Context) InternetserviceipblreasonArrayOutput {
	return o
}

func (o InternetserviceipblreasonArrayOutput) Index(i pulumi.IntInput) InternetserviceipblreasonOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Internetserviceipblreason {
		return vs[0].([]*Internetserviceipblreason)[vs[1].(int)]
	}).(InternetserviceipblreasonOutput)
}

type InternetserviceipblreasonMapOutput struct{ *pulumi.OutputState }

func (InternetserviceipblreasonMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Internetserviceipblreason)(nil)).Elem()
}

func (o InternetserviceipblreasonMapOutput) ToInternetserviceipblreasonMapOutput() InternetserviceipblreasonMapOutput {
	return o
}

func (o InternetserviceipblreasonMapOutput) ToInternetserviceipblreasonMapOutputWithContext(ctx context.Context) InternetserviceipblreasonMapOutput {
	return o
}

func (o InternetserviceipblreasonMapOutput) MapIndex(k pulumi.StringInput) InternetserviceipblreasonOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Internetserviceipblreason {
		return vs[0].(map[string]*Internetserviceipblreason)[vs[1].(string)]
	}).(InternetserviceipblreasonOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InternetserviceipblreasonInput)(nil)).Elem(), &Internetserviceipblreason{})
	pulumi.RegisterInputType(reflect.TypeOf((*InternetserviceipblreasonArrayInput)(nil)).Elem(), InternetserviceipblreasonArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InternetserviceipblreasonMapInput)(nil)).Elem(), InternetserviceipblreasonMap{})
	pulumi.RegisterOutputType(InternetserviceipblreasonOutput{})
	pulumi.RegisterOutputType(InternetserviceipblreasonArrayOutput{})
	pulumi.RegisterOutputType(InternetserviceipblreasonMapOutput{})
}
