// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Internet Service list. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// # Firewall InternetServiceList can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallInternetservicelist:FirewallInternetservicelist labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallInternetservicelist:FirewallInternetservicelist labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallInternetservicelist struct {
	pulumi.CustomResourceState

	// Internet Service category ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Internet Service category name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallInternetservicelist registers a new resource with the given unique name, arguments, and options.
func NewFirewallInternetservicelist(ctx *pulumi.Context,
	name string, args *FirewallInternetservicelistArgs, opts ...pulumi.ResourceOption) (*FirewallInternetservicelist, error) {
	if args == nil {
		args = &FirewallInternetservicelistArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallInternetservicelist
	err := ctx.RegisterResource("fortios:index/firewallInternetservicelist:FirewallInternetservicelist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallInternetservicelist gets an existing FirewallInternetservicelist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallInternetservicelist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallInternetservicelistState, opts ...pulumi.ResourceOption) (*FirewallInternetservicelist, error) {
	var resource FirewallInternetservicelist
	err := ctx.ReadResource("fortios:index/firewallInternetservicelist:FirewallInternetservicelist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallInternetservicelist resources.
type firewallInternetservicelistState struct {
	// Internet Service category ID.
	Fosid *int `pulumi:"fosid"`
	// Internet Service category name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallInternetservicelistState struct {
	// Internet Service category ID.
	Fosid pulumi.IntPtrInput
	// Internet Service category name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetservicelistState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetservicelistState)(nil)).Elem()
}

type firewallInternetservicelistArgs struct {
	// Internet Service category ID.
	Fosid *int `pulumi:"fosid"`
	// Internet Service category name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallInternetservicelist resource.
type FirewallInternetservicelistArgs struct {
	// Internet Service category ID.
	Fosid pulumi.IntPtrInput
	// Internet Service category name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetservicelistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetservicelistArgs)(nil)).Elem()
}

type FirewallInternetservicelistInput interface {
	pulumi.Input

	ToFirewallInternetservicelistOutput() FirewallInternetservicelistOutput
	ToFirewallInternetservicelistOutputWithContext(ctx context.Context) FirewallInternetservicelistOutput
}

func (*FirewallInternetservicelist) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetservicelist)(nil)).Elem()
}

func (i *FirewallInternetservicelist) ToFirewallInternetservicelistOutput() FirewallInternetservicelistOutput {
	return i.ToFirewallInternetservicelistOutputWithContext(context.Background())
}

func (i *FirewallInternetservicelist) ToFirewallInternetservicelistOutputWithContext(ctx context.Context) FirewallInternetservicelistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetservicelistOutput)
}

// FirewallInternetservicelistArrayInput is an input type that accepts FirewallInternetservicelistArray and FirewallInternetservicelistArrayOutput values.
// You can construct a concrete instance of `FirewallInternetservicelistArrayInput` via:
//
//	FirewallInternetservicelistArray{ FirewallInternetservicelistArgs{...} }
type FirewallInternetservicelistArrayInput interface {
	pulumi.Input

	ToFirewallInternetservicelistArrayOutput() FirewallInternetservicelistArrayOutput
	ToFirewallInternetservicelistArrayOutputWithContext(context.Context) FirewallInternetservicelistArrayOutput
}

type FirewallInternetservicelistArray []FirewallInternetservicelistInput

func (FirewallInternetservicelistArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetservicelist)(nil)).Elem()
}

func (i FirewallInternetservicelistArray) ToFirewallInternetservicelistArrayOutput() FirewallInternetservicelistArrayOutput {
	return i.ToFirewallInternetservicelistArrayOutputWithContext(context.Background())
}

func (i FirewallInternetservicelistArray) ToFirewallInternetservicelistArrayOutputWithContext(ctx context.Context) FirewallInternetservicelistArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetservicelistArrayOutput)
}

// FirewallInternetservicelistMapInput is an input type that accepts FirewallInternetservicelistMap and FirewallInternetservicelistMapOutput values.
// You can construct a concrete instance of `FirewallInternetservicelistMapInput` via:
//
//	FirewallInternetservicelistMap{ "key": FirewallInternetservicelistArgs{...} }
type FirewallInternetservicelistMapInput interface {
	pulumi.Input

	ToFirewallInternetservicelistMapOutput() FirewallInternetservicelistMapOutput
	ToFirewallInternetservicelistMapOutputWithContext(context.Context) FirewallInternetservicelistMapOutput
}

type FirewallInternetservicelistMap map[string]FirewallInternetservicelistInput

func (FirewallInternetservicelistMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetservicelist)(nil)).Elem()
}

func (i FirewallInternetservicelistMap) ToFirewallInternetservicelistMapOutput() FirewallInternetservicelistMapOutput {
	return i.ToFirewallInternetservicelistMapOutputWithContext(context.Background())
}

func (i FirewallInternetservicelistMap) ToFirewallInternetservicelistMapOutputWithContext(ctx context.Context) FirewallInternetservicelistMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetservicelistMapOutput)
}

type FirewallInternetservicelistOutput struct{ *pulumi.OutputState }

func (FirewallInternetservicelistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetservicelist)(nil)).Elem()
}

func (o FirewallInternetservicelistOutput) ToFirewallInternetservicelistOutput() FirewallInternetservicelistOutput {
	return o
}

func (o FirewallInternetservicelistOutput) ToFirewallInternetservicelistOutputWithContext(ctx context.Context) FirewallInternetservicelistOutput {
	return o
}

// Internet Service category ID.
func (o FirewallInternetservicelistOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallInternetservicelist) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Internet Service category name.
func (o FirewallInternetservicelistOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInternetservicelist) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallInternetservicelistOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallInternetservicelist) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallInternetservicelistArrayOutput struct{ *pulumi.OutputState }

func (FirewallInternetservicelistArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetservicelist)(nil)).Elem()
}

func (o FirewallInternetservicelistArrayOutput) ToFirewallInternetservicelistArrayOutput() FirewallInternetservicelistArrayOutput {
	return o
}

func (o FirewallInternetservicelistArrayOutput) ToFirewallInternetservicelistArrayOutputWithContext(ctx context.Context) FirewallInternetservicelistArrayOutput {
	return o
}

func (o FirewallInternetservicelistArrayOutput) Index(i pulumi.IntInput) FirewallInternetservicelistOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallInternetservicelist {
		return vs[0].([]*FirewallInternetservicelist)[vs[1].(int)]
	}).(FirewallInternetservicelistOutput)
}

type FirewallInternetservicelistMapOutput struct{ *pulumi.OutputState }

func (FirewallInternetservicelistMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetservicelist)(nil)).Elem()
}

func (o FirewallInternetservicelistMapOutput) ToFirewallInternetservicelistMapOutput() FirewallInternetservicelistMapOutput {
	return o
}

func (o FirewallInternetservicelistMapOutput) ToFirewallInternetservicelistMapOutputWithContext(ctx context.Context) FirewallInternetservicelistMapOutput {
	return o
}

func (o FirewallInternetservicelistMapOutput) MapIndex(k pulumi.StringInput) FirewallInternetservicelistOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallInternetservicelist {
		return vs[0].(map[string]*FirewallInternetservicelist)[vs[1].(string)]
	}).(FirewallInternetservicelistOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetservicelistInput)(nil)).Elem(), &FirewallInternetservicelist{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetservicelistArrayInput)(nil)).Elem(), FirewallInternetservicelistArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetservicelistMapInput)(nil)).Elem(), FirewallInternetservicelistMap{})
	pulumi.RegisterOutputType(FirewallInternetservicelistOutput{})
	pulumi.RegisterOutputType(FirewallInternetservicelistArrayOutput{})
	pulumi.RegisterOutputType(FirewallInternetservicelistMapOutput{})
}
