// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Internet Service list. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// Firewall InternetServiceList can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/internetservicelist:Internetservicelist labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/internetservicelist:Internetservicelist labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Internetservicelist struct {
	pulumi.CustomResourceState

	// Internet Service category ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Internet Service category name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewInternetservicelist registers a new resource with the given unique name, arguments, and options.
func NewInternetservicelist(ctx *pulumi.Context,
	name string, args *InternetservicelistArgs, opts ...pulumi.ResourceOption) (*Internetservicelist, error) {
	if args == nil {
		args = &InternetservicelistArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Internetservicelist
	err := ctx.RegisterResource("fortios:firewall/internetservicelist:Internetservicelist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInternetservicelist gets an existing Internetservicelist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInternetservicelist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InternetservicelistState, opts ...pulumi.ResourceOption) (*Internetservicelist, error) {
	var resource Internetservicelist
	err := ctx.ReadResource("fortios:firewall/internetservicelist:Internetservicelist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Internetservicelist resources.
type internetservicelistState struct {
	// Internet Service category ID.
	Fosid *int `pulumi:"fosid"`
	// Internet Service category name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type InternetservicelistState struct {
	// Internet Service category ID.
	Fosid pulumi.IntPtrInput
	// Internet Service category name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (InternetservicelistState) ElementType() reflect.Type {
	return reflect.TypeOf((*internetservicelistState)(nil)).Elem()
}

type internetservicelistArgs struct {
	// Internet Service category ID.
	Fosid *int `pulumi:"fosid"`
	// Internet Service category name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Internetservicelist resource.
type InternetservicelistArgs struct {
	// Internet Service category ID.
	Fosid pulumi.IntPtrInput
	// Internet Service category name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (InternetservicelistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*internetservicelistArgs)(nil)).Elem()
}

type InternetservicelistInput interface {
	pulumi.Input

	ToInternetservicelistOutput() InternetservicelistOutput
	ToInternetservicelistOutputWithContext(ctx context.Context) InternetservicelistOutput
}

func (*Internetservicelist) ElementType() reflect.Type {
	return reflect.TypeOf((**Internetservicelist)(nil)).Elem()
}

func (i *Internetservicelist) ToInternetservicelistOutput() InternetservicelistOutput {
	return i.ToInternetservicelistOutputWithContext(context.Background())
}

func (i *Internetservicelist) ToInternetservicelistOutputWithContext(ctx context.Context) InternetservicelistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetservicelistOutput)
}

// InternetservicelistArrayInput is an input type that accepts InternetservicelistArray and InternetservicelistArrayOutput values.
// You can construct a concrete instance of `InternetservicelistArrayInput` via:
//
//	InternetservicelistArray{ InternetservicelistArgs{...} }
type InternetservicelistArrayInput interface {
	pulumi.Input

	ToInternetservicelistArrayOutput() InternetservicelistArrayOutput
	ToInternetservicelistArrayOutputWithContext(context.Context) InternetservicelistArrayOutput
}

type InternetservicelistArray []InternetservicelistInput

func (InternetservicelistArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Internetservicelist)(nil)).Elem()
}

func (i InternetservicelistArray) ToInternetservicelistArrayOutput() InternetservicelistArrayOutput {
	return i.ToInternetservicelistArrayOutputWithContext(context.Background())
}

func (i InternetservicelistArray) ToInternetservicelistArrayOutputWithContext(ctx context.Context) InternetservicelistArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetservicelistArrayOutput)
}

// InternetservicelistMapInput is an input type that accepts InternetservicelistMap and InternetservicelistMapOutput values.
// You can construct a concrete instance of `InternetservicelistMapInput` via:
//
//	InternetservicelistMap{ "key": InternetservicelistArgs{...} }
type InternetservicelistMapInput interface {
	pulumi.Input

	ToInternetservicelistMapOutput() InternetservicelistMapOutput
	ToInternetservicelistMapOutputWithContext(context.Context) InternetservicelistMapOutput
}

type InternetservicelistMap map[string]InternetservicelistInput

func (InternetservicelistMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Internetservicelist)(nil)).Elem()
}

func (i InternetservicelistMap) ToInternetservicelistMapOutput() InternetservicelistMapOutput {
	return i.ToInternetservicelistMapOutputWithContext(context.Background())
}

func (i InternetservicelistMap) ToInternetservicelistMapOutputWithContext(ctx context.Context) InternetservicelistMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetservicelistMapOutput)
}

type InternetservicelistOutput struct{ *pulumi.OutputState }

func (InternetservicelistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Internetservicelist)(nil)).Elem()
}

func (o InternetservicelistOutput) ToInternetservicelistOutput() InternetservicelistOutput {
	return o
}

func (o InternetservicelistOutput) ToInternetservicelistOutputWithContext(ctx context.Context) InternetservicelistOutput {
	return o
}

// Internet Service category ID.
func (o InternetservicelistOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetservicelist) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Internet Service category name.
func (o InternetservicelistOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Internetservicelist) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o InternetservicelistOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Internetservicelist) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type InternetservicelistArrayOutput struct{ *pulumi.OutputState }

func (InternetservicelistArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Internetservicelist)(nil)).Elem()
}

func (o InternetservicelistArrayOutput) ToInternetservicelistArrayOutput() InternetservicelistArrayOutput {
	return o
}

func (o InternetservicelistArrayOutput) ToInternetservicelistArrayOutputWithContext(ctx context.Context) InternetservicelistArrayOutput {
	return o
}

func (o InternetservicelistArrayOutput) Index(i pulumi.IntInput) InternetservicelistOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Internetservicelist {
		return vs[0].([]*Internetservicelist)[vs[1].(int)]
	}).(InternetservicelistOutput)
}

type InternetservicelistMapOutput struct{ *pulumi.OutputState }

func (InternetservicelistMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Internetservicelist)(nil)).Elem()
}

func (o InternetservicelistMapOutput) ToInternetservicelistMapOutput() InternetservicelistMapOutput {
	return o
}

func (o InternetservicelistMapOutput) ToInternetservicelistMapOutputWithContext(ctx context.Context) InternetservicelistMapOutput {
	return o
}

func (o InternetservicelistMapOutput) MapIndex(k pulumi.StringInput) InternetservicelistOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Internetservicelist {
		return vs[0].(map[string]*Internetservicelist)[vs[1].(string)]
	}).(InternetservicelistOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InternetservicelistInput)(nil)).Elem(), &Internetservicelist{})
	pulumi.RegisterInputType(reflect.TypeOf((*InternetservicelistArrayInput)(nil)).Elem(), InternetservicelistArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InternetservicelistMapInput)(nil)).Elem(), InternetservicelistMap{})
	pulumi.RegisterOutputType(InternetservicelistOutput{})
	pulumi.RegisterOutputType(InternetservicelistArrayOutput{})
	pulumi.RegisterOutputType(InternetservicelistMapOutput{})
}
