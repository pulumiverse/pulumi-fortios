// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Show Internet Service application.
//
// ## Import
//
// Firewall InternetService can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/internetservice:Internetservice labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/internetservice:Internetservice labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Internetservice struct {
	pulumi.CustomResourceState

	// Database name this Internet Service belongs to. Valid values: `isdb`, `irdb`.
	Database pulumi.StringOutput `pulumi:"database"`
	// How this service may be used in a firewall policy (source, destination or both). Valid values: `src`, `dst`, `both`.
	Direction pulumi.StringOutput `pulumi:"direction"`
	// Extra number of IPv6 ranges.
	ExtraIp6RangeNumber pulumi.IntOutput `pulumi:"extraIp6RangeNumber"`
	// Extra number of IP ranges.
	ExtraIpRangeNumber pulumi.IntOutput `pulumi:"extraIpRangeNumber"`
	// Internet Service ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Icon ID of Internet Service.
	IconId pulumi.IntOutput `pulumi:"iconId"`
	// Number of IPv6 ranges.
	Ip6RangeNumber pulumi.IntOutput `pulumi:"ip6RangeNumber"`
	// Total number of IP addresses.
	IpNumber pulumi.IntOutput `pulumi:"ipNumber"`
	// Total number of IP ranges.
	IpRangeNumber pulumi.IntOutput `pulumi:"ipRangeNumber"`
	// Internet Service name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates whether the Internet Service can be used.
	Obsolete pulumi.IntOutput `pulumi:"obsolete"`
	// Reputation level of the Internet Service.
	Reputation pulumi.IntOutput `pulumi:"reputation"`
	// Singular level of the Internet Service.
	Singularity pulumi.IntOutput `pulumi:"singularity"`
	// Second Level Domain.
	SldId pulumi.IntOutput `pulumi:"sldId"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewInternetservice registers a new resource with the given unique name, arguments, and options.
func NewInternetservice(ctx *pulumi.Context,
	name string, args *InternetserviceArgs, opts ...pulumi.ResourceOption) (*Internetservice, error) {
	if args == nil {
		args = &InternetserviceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Internetservice
	err := ctx.RegisterResource("fortios:firewall/internetservice:Internetservice", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInternetservice gets an existing Internetservice resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInternetservice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InternetserviceState, opts ...pulumi.ResourceOption) (*Internetservice, error) {
	var resource Internetservice
	err := ctx.ReadResource("fortios:firewall/internetservice:Internetservice", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Internetservice resources.
type internetserviceState struct {
	// Database name this Internet Service belongs to. Valid values: `isdb`, `irdb`.
	Database *string `pulumi:"database"`
	// How this service may be used in a firewall policy (source, destination or both). Valid values: `src`, `dst`, `both`.
	Direction *string `pulumi:"direction"`
	// Extra number of IPv6 ranges.
	ExtraIp6RangeNumber *int `pulumi:"extraIp6RangeNumber"`
	// Extra number of IP ranges.
	ExtraIpRangeNumber *int `pulumi:"extraIpRangeNumber"`
	// Internet Service ID.
	Fosid *int `pulumi:"fosid"`
	// Icon ID of Internet Service.
	IconId *int `pulumi:"iconId"`
	// Number of IPv6 ranges.
	Ip6RangeNumber *int `pulumi:"ip6RangeNumber"`
	// Total number of IP addresses.
	IpNumber *int `pulumi:"ipNumber"`
	// Total number of IP ranges.
	IpRangeNumber *int `pulumi:"ipRangeNumber"`
	// Internet Service name.
	Name *string `pulumi:"name"`
	// Indicates whether the Internet Service can be used.
	Obsolete *int `pulumi:"obsolete"`
	// Reputation level of the Internet Service.
	Reputation *int `pulumi:"reputation"`
	// Singular level of the Internet Service.
	Singularity *int `pulumi:"singularity"`
	// Second Level Domain.
	SldId *int `pulumi:"sldId"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type InternetserviceState struct {
	// Database name this Internet Service belongs to. Valid values: `isdb`, `irdb`.
	Database pulumi.StringPtrInput
	// How this service may be used in a firewall policy (source, destination or both). Valid values: `src`, `dst`, `both`.
	Direction pulumi.StringPtrInput
	// Extra number of IPv6 ranges.
	ExtraIp6RangeNumber pulumi.IntPtrInput
	// Extra number of IP ranges.
	ExtraIpRangeNumber pulumi.IntPtrInput
	// Internet Service ID.
	Fosid pulumi.IntPtrInput
	// Icon ID of Internet Service.
	IconId pulumi.IntPtrInput
	// Number of IPv6 ranges.
	Ip6RangeNumber pulumi.IntPtrInput
	// Total number of IP addresses.
	IpNumber pulumi.IntPtrInput
	// Total number of IP ranges.
	IpRangeNumber pulumi.IntPtrInput
	// Internet Service name.
	Name pulumi.StringPtrInput
	// Indicates whether the Internet Service can be used.
	Obsolete pulumi.IntPtrInput
	// Reputation level of the Internet Service.
	Reputation pulumi.IntPtrInput
	// Singular level of the Internet Service.
	Singularity pulumi.IntPtrInput
	// Second Level Domain.
	SldId pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (InternetserviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*internetserviceState)(nil)).Elem()
}

type internetserviceArgs struct {
	// Database name this Internet Service belongs to. Valid values: `isdb`, `irdb`.
	Database *string `pulumi:"database"`
	// How this service may be used in a firewall policy (source, destination or both). Valid values: `src`, `dst`, `both`.
	Direction *string `pulumi:"direction"`
	// Extra number of IPv6 ranges.
	ExtraIp6RangeNumber *int `pulumi:"extraIp6RangeNumber"`
	// Extra number of IP ranges.
	ExtraIpRangeNumber *int `pulumi:"extraIpRangeNumber"`
	// Internet Service ID.
	Fosid *int `pulumi:"fosid"`
	// Icon ID of Internet Service.
	IconId *int `pulumi:"iconId"`
	// Number of IPv6 ranges.
	Ip6RangeNumber *int `pulumi:"ip6RangeNumber"`
	// Total number of IP addresses.
	IpNumber *int `pulumi:"ipNumber"`
	// Total number of IP ranges.
	IpRangeNumber *int `pulumi:"ipRangeNumber"`
	// Internet Service name.
	Name *string `pulumi:"name"`
	// Indicates whether the Internet Service can be used.
	Obsolete *int `pulumi:"obsolete"`
	// Reputation level of the Internet Service.
	Reputation *int `pulumi:"reputation"`
	// Singular level of the Internet Service.
	Singularity *int `pulumi:"singularity"`
	// Second Level Domain.
	SldId *int `pulumi:"sldId"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Internetservice resource.
type InternetserviceArgs struct {
	// Database name this Internet Service belongs to. Valid values: `isdb`, `irdb`.
	Database pulumi.StringPtrInput
	// How this service may be used in a firewall policy (source, destination or both). Valid values: `src`, `dst`, `both`.
	Direction pulumi.StringPtrInput
	// Extra number of IPv6 ranges.
	ExtraIp6RangeNumber pulumi.IntPtrInput
	// Extra number of IP ranges.
	ExtraIpRangeNumber pulumi.IntPtrInput
	// Internet Service ID.
	Fosid pulumi.IntPtrInput
	// Icon ID of Internet Service.
	IconId pulumi.IntPtrInput
	// Number of IPv6 ranges.
	Ip6RangeNumber pulumi.IntPtrInput
	// Total number of IP addresses.
	IpNumber pulumi.IntPtrInput
	// Total number of IP ranges.
	IpRangeNumber pulumi.IntPtrInput
	// Internet Service name.
	Name pulumi.StringPtrInput
	// Indicates whether the Internet Service can be used.
	Obsolete pulumi.IntPtrInput
	// Reputation level of the Internet Service.
	Reputation pulumi.IntPtrInput
	// Singular level of the Internet Service.
	Singularity pulumi.IntPtrInput
	// Second Level Domain.
	SldId pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (InternetserviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*internetserviceArgs)(nil)).Elem()
}

type InternetserviceInput interface {
	pulumi.Input

	ToInternetserviceOutput() InternetserviceOutput
	ToInternetserviceOutputWithContext(ctx context.Context) InternetserviceOutput
}

func (*Internetservice) ElementType() reflect.Type {
	return reflect.TypeOf((**Internetservice)(nil)).Elem()
}

func (i *Internetservice) ToInternetserviceOutput() InternetserviceOutput {
	return i.ToInternetserviceOutputWithContext(context.Background())
}

func (i *Internetservice) ToInternetserviceOutputWithContext(ctx context.Context) InternetserviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetserviceOutput)
}

// InternetserviceArrayInput is an input type that accepts InternetserviceArray and InternetserviceArrayOutput values.
// You can construct a concrete instance of `InternetserviceArrayInput` via:
//
//	InternetserviceArray{ InternetserviceArgs{...} }
type InternetserviceArrayInput interface {
	pulumi.Input

	ToInternetserviceArrayOutput() InternetserviceArrayOutput
	ToInternetserviceArrayOutputWithContext(context.Context) InternetserviceArrayOutput
}

type InternetserviceArray []InternetserviceInput

func (InternetserviceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Internetservice)(nil)).Elem()
}

func (i InternetserviceArray) ToInternetserviceArrayOutput() InternetserviceArrayOutput {
	return i.ToInternetserviceArrayOutputWithContext(context.Background())
}

func (i InternetserviceArray) ToInternetserviceArrayOutputWithContext(ctx context.Context) InternetserviceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetserviceArrayOutput)
}

// InternetserviceMapInput is an input type that accepts InternetserviceMap and InternetserviceMapOutput values.
// You can construct a concrete instance of `InternetserviceMapInput` via:
//
//	InternetserviceMap{ "key": InternetserviceArgs{...} }
type InternetserviceMapInput interface {
	pulumi.Input

	ToInternetserviceMapOutput() InternetserviceMapOutput
	ToInternetserviceMapOutputWithContext(context.Context) InternetserviceMapOutput
}

type InternetserviceMap map[string]InternetserviceInput

func (InternetserviceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Internetservice)(nil)).Elem()
}

func (i InternetserviceMap) ToInternetserviceMapOutput() InternetserviceMapOutput {
	return i.ToInternetserviceMapOutputWithContext(context.Background())
}

func (i InternetserviceMap) ToInternetserviceMapOutputWithContext(ctx context.Context) InternetserviceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetserviceMapOutput)
}

type InternetserviceOutput struct{ *pulumi.OutputState }

func (InternetserviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Internetservice)(nil)).Elem()
}

func (o InternetserviceOutput) ToInternetserviceOutput() InternetserviceOutput {
	return o
}

func (o InternetserviceOutput) ToInternetserviceOutputWithContext(ctx context.Context) InternetserviceOutput {
	return o
}

// Database name this Internet Service belongs to. Valid values: `isdb`, `irdb`.
func (o InternetserviceOutput) Database() pulumi.StringOutput {
	return o.ApplyT(func(v *Internetservice) pulumi.StringOutput { return v.Database }).(pulumi.StringOutput)
}

// How this service may be used in a firewall policy (source, destination or both). Valid values: `src`, `dst`, `both`.
func (o InternetserviceOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v *Internetservice) pulumi.StringOutput { return v.Direction }).(pulumi.StringOutput)
}

// Extra number of IPv6 ranges.
func (o InternetserviceOutput) ExtraIp6RangeNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetservice) pulumi.IntOutput { return v.ExtraIp6RangeNumber }).(pulumi.IntOutput)
}

// Extra number of IP ranges.
func (o InternetserviceOutput) ExtraIpRangeNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetservice) pulumi.IntOutput { return v.ExtraIpRangeNumber }).(pulumi.IntOutput)
}

// Internet Service ID.
func (o InternetserviceOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetservice) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Icon ID of Internet Service.
func (o InternetserviceOutput) IconId() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetservice) pulumi.IntOutput { return v.IconId }).(pulumi.IntOutput)
}

// Number of IPv6 ranges.
func (o InternetserviceOutput) Ip6RangeNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetservice) pulumi.IntOutput { return v.Ip6RangeNumber }).(pulumi.IntOutput)
}

// Total number of IP addresses.
func (o InternetserviceOutput) IpNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetservice) pulumi.IntOutput { return v.IpNumber }).(pulumi.IntOutput)
}

// Total number of IP ranges.
func (o InternetserviceOutput) IpRangeNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetservice) pulumi.IntOutput { return v.IpRangeNumber }).(pulumi.IntOutput)
}

// Internet Service name.
func (o InternetserviceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Internetservice) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Indicates whether the Internet Service can be used.
func (o InternetserviceOutput) Obsolete() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetservice) pulumi.IntOutput { return v.Obsolete }).(pulumi.IntOutput)
}

// Reputation level of the Internet Service.
func (o InternetserviceOutput) Reputation() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetservice) pulumi.IntOutput { return v.Reputation }).(pulumi.IntOutput)
}

// Singular level of the Internet Service.
func (o InternetserviceOutput) Singularity() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetservice) pulumi.IntOutput { return v.Singularity }).(pulumi.IntOutput)
}

// Second Level Domain.
func (o InternetserviceOutput) SldId() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetservice) pulumi.IntOutput { return v.SldId }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o InternetserviceOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Internetservice) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type InternetserviceArrayOutput struct{ *pulumi.OutputState }

func (InternetserviceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Internetservice)(nil)).Elem()
}

func (o InternetserviceArrayOutput) ToInternetserviceArrayOutput() InternetserviceArrayOutput {
	return o
}

func (o InternetserviceArrayOutput) ToInternetserviceArrayOutputWithContext(ctx context.Context) InternetserviceArrayOutput {
	return o
}

func (o InternetserviceArrayOutput) Index(i pulumi.IntInput) InternetserviceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Internetservice {
		return vs[0].([]*Internetservice)[vs[1].(int)]
	}).(InternetserviceOutput)
}

type InternetserviceMapOutput struct{ *pulumi.OutputState }

func (InternetserviceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Internetservice)(nil)).Elem()
}

func (o InternetserviceMapOutput) ToInternetserviceMapOutput() InternetserviceMapOutput {
	return o
}

func (o InternetserviceMapOutput) ToInternetserviceMapOutputWithContext(ctx context.Context) InternetserviceMapOutput {
	return o
}

func (o InternetserviceMapOutput) MapIndex(k pulumi.StringInput) InternetserviceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Internetservice {
		return vs[0].(map[string]*Internetservice)[vs[1].(string)]
	}).(InternetserviceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InternetserviceInput)(nil)).Elem(), &Internetservice{})
	pulumi.RegisterInputType(reflect.TypeOf((*InternetserviceArrayInput)(nil)).Elem(), InternetserviceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InternetserviceMapInput)(nil)).Elem(), InternetserviceMap{})
	pulumi.RegisterOutputType(InternetserviceOutput{})
	pulumi.RegisterOutputType(InternetserviceArrayOutput{})
	pulumi.RegisterOutputType(InternetserviceMapOutput{})
}
