// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv4 DoS policies.
//
// ## Import
//
// # Firewall DosPolicy can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallDoSpolicy:FirewallDoSpolicy labelname {{policyid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallDoSpolicy:FirewallDoSpolicy labelname {{policyid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallDoSpolicy struct {
	pulumi.CustomResourceState

	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies FirewallDoSpolicyAnomalyArrayOutput `pulumi:"anomalies"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs FirewallDoSpolicyDstaddrArrayOutput `pulumi:"dstaddrs"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Incoming interface name from available interfaces.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Policy ID.
	Policyid pulumi.IntOutput `pulumi:"policyid"`
	// Service object from available options. The structure of `service` block is documented below.
	Services FirewallDoSpolicyServiceArrayOutput `pulumi:"services"`
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs FirewallDoSpolicySrcaddrArrayOutput `pulumi:"srcaddrs"`
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallDoSpolicy registers a new resource with the given unique name, arguments, and options.
func NewFirewallDoSpolicy(ctx *pulumi.Context,
	name string, args *FirewallDoSpolicyArgs, opts ...pulumi.ResourceOption) (*FirewallDoSpolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dstaddrs == nil {
		return nil, errors.New("invalid value for required argument 'Dstaddrs'")
	}
	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.Srcaddrs == nil {
		return nil, errors.New("invalid value for required argument 'Srcaddrs'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallDoSpolicy
	err := ctx.RegisterResource("fortios:index/firewallDoSpolicy:FirewallDoSpolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallDoSpolicy gets an existing FirewallDoSpolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallDoSpolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallDoSpolicyState, opts ...pulumi.ResourceOption) (*FirewallDoSpolicy, error) {
	var resource FirewallDoSpolicy
	err := ctx.ReadResource("fortios:index/firewallDoSpolicy:FirewallDoSpolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallDoSpolicy resources.
type firewallDoSpolicyState struct {
	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies []FirewallDoSpolicyAnomaly `pulumi:"anomalies"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs []FirewallDoSpolicyDstaddr `pulumi:"dstaddrs"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Incoming interface name from available interfaces.
	Interface *string `pulumi:"interface"`
	// Policy name.
	Name *string `pulumi:"name"`
	// Policy ID.
	Policyid *int `pulumi:"policyid"`
	// Service object from available options. The structure of `service` block is documented below.
	Services []FirewallDoSpolicyService `pulumi:"services"`
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs []FirewallDoSpolicySrcaddr `pulumi:"srcaddrs"`
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallDoSpolicyState struct {
	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies FirewallDoSpolicyAnomalyArrayInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs FirewallDoSpolicyDstaddrArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Incoming interface name from available interfaces.
	Interface pulumi.StringPtrInput
	// Policy name.
	Name pulumi.StringPtrInput
	// Policy ID.
	Policyid pulumi.IntPtrInput
	// Service object from available options. The structure of `service` block is documented below.
	Services FirewallDoSpolicyServiceArrayInput
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs FirewallDoSpolicySrcaddrArrayInput
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallDoSpolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallDoSpolicyState)(nil)).Elem()
}

type firewallDoSpolicyArgs struct {
	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies []FirewallDoSpolicyAnomaly `pulumi:"anomalies"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs []FirewallDoSpolicyDstaddr `pulumi:"dstaddrs"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Incoming interface name from available interfaces.
	Interface string `pulumi:"interface"`
	// Policy name.
	Name *string `pulumi:"name"`
	// Policy ID.
	Policyid *int `pulumi:"policyid"`
	// Service object from available options. The structure of `service` block is documented below.
	Services []FirewallDoSpolicyService `pulumi:"services"`
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs []FirewallDoSpolicySrcaddr `pulumi:"srcaddrs"`
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallDoSpolicy resource.
type FirewallDoSpolicyArgs struct {
	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies FirewallDoSpolicyAnomalyArrayInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs FirewallDoSpolicyDstaddrArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Incoming interface name from available interfaces.
	Interface pulumi.StringInput
	// Policy name.
	Name pulumi.StringPtrInput
	// Policy ID.
	Policyid pulumi.IntPtrInput
	// Service object from available options. The structure of `service` block is documented below.
	Services FirewallDoSpolicyServiceArrayInput
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs FirewallDoSpolicySrcaddrArrayInput
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallDoSpolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallDoSpolicyArgs)(nil)).Elem()
}

type FirewallDoSpolicyInput interface {
	pulumi.Input

	ToFirewallDoSpolicyOutput() FirewallDoSpolicyOutput
	ToFirewallDoSpolicyOutputWithContext(ctx context.Context) FirewallDoSpolicyOutput
}

func (*FirewallDoSpolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallDoSpolicy)(nil)).Elem()
}

func (i *FirewallDoSpolicy) ToFirewallDoSpolicyOutput() FirewallDoSpolicyOutput {
	return i.ToFirewallDoSpolicyOutputWithContext(context.Background())
}

func (i *FirewallDoSpolicy) ToFirewallDoSpolicyOutputWithContext(ctx context.Context) FirewallDoSpolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallDoSpolicyOutput)
}

// FirewallDoSpolicyArrayInput is an input type that accepts FirewallDoSpolicyArray and FirewallDoSpolicyArrayOutput values.
// You can construct a concrete instance of `FirewallDoSpolicyArrayInput` via:
//
//	FirewallDoSpolicyArray{ FirewallDoSpolicyArgs{...} }
type FirewallDoSpolicyArrayInput interface {
	pulumi.Input

	ToFirewallDoSpolicyArrayOutput() FirewallDoSpolicyArrayOutput
	ToFirewallDoSpolicyArrayOutputWithContext(context.Context) FirewallDoSpolicyArrayOutput
}

type FirewallDoSpolicyArray []FirewallDoSpolicyInput

func (FirewallDoSpolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallDoSpolicy)(nil)).Elem()
}

func (i FirewallDoSpolicyArray) ToFirewallDoSpolicyArrayOutput() FirewallDoSpolicyArrayOutput {
	return i.ToFirewallDoSpolicyArrayOutputWithContext(context.Background())
}

func (i FirewallDoSpolicyArray) ToFirewallDoSpolicyArrayOutputWithContext(ctx context.Context) FirewallDoSpolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallDoSpolicyArrayOutput)
}

// FirewallDoSpolicyMapInput is an input type that accepts FirewallDoSpolicyMap and FirewallDoSpolicyMapOutput values.
// You can construct a concrete instance of `FirewallDoSpolicyMapInput` via:
//
//	FirewallDoSpolicyMap{ "key": FirewallDoSpolicyArgs{...} }
type FirewallDoSpolicyMapInput interface {
	pulumi.Input

	ToFirewallDoSpolicyMapOutput() FirewallDoSpolicyMapOutput
	ToFirewallDoSpolicyMapOutputWithContext(context.Context) FirewallDoSpolicyMapOutput
}

type FirewallDoSpolicyMap map[string]FirewallDoSpolicyInput

func (FirewallDoSpolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallDoSpolicy)(nil)).Elem()
}

func (i FirewallDoSpolicyMap) ToFirewallDoSpolicyMapOutput() FirewallDoSpolicyMapOutput {
	return i.ToFirewallDoSpolicyMapOutputWithContext(context.Background())
}

func (i FirewallDoSpolicyMap) ToFirewallDoSpolicyMapOutputWithContext(ctx context.Context) FirewallDoSpolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallDoSpolicyMapOutput)
}

type FirewallDoSpolicyOutput struct{ *pulumi.OutputState }

func (FirewallDoSpolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallDoSpolicy)(nil)).Elem()
}

func (o FirewallDoSpolicyOutput) ToFirewallDoSpolicyOutput() FirewallDoSpolicyOutput {
	return o
}

func (o FirewallDoSpolicyOutput) ToFirewallDoSpolicyOutputWithContext(ctx context.Context) FirewallDoSpolicyOutput {
	return o
}

// Anomaly name. The structure of `anomaly` block is documented below.
func (o FirewallDoSpolicyOutput) Anomalies() FirewallDoSpolicyAnomalyArrayOutput {
	return o.ApplyT(func(v *FirewallDoSpolicy) FirewallDoSpolicyAnomalyArrayOutput { return v.Anomalies }).(FirewallDoSpolicyAnomalyArrayOutput)
}

// Comment.
func (o FirewallDoSpolicyOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallDoSpolicy) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
func (o FirewallDoSpolicyOutput) Dstaddrs() FirewallDoSpolicyDstaddrArrayOutput {
	return o.ApplyT(func(v *FirewallDoSpolicy) FirewallDoSpolicyDstaddrArrayOutput { return v.Dstaddrs }).(FirewallDoSpolicyDstaddrArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o FirewallDoSpolicyOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallDoSpolicy) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Incoming interface name from available interfaces.
func (o FirewallDoSpolicyOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallDoSpolicy) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Policy name.
func (o FirewallDoSpolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallDoSpolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Policy ID.
func (o FirewallDoSpolicyOutput) Policyid() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallDoSpolicy) pulumi.IntOutput { return v.Policyid }).(pulumi.IntOutput)
}

// Service object from available options. The structure of `service` block is documented below.
func (o FirewallDoSpolicyOutput) Services() FirewallDoSpolicyServiceArrayOutput {
	return o.ApplyT(func(v *FirewallDoSpolicy) FirewallDoSpolicyServiceArrayOutput { return v.Services }).(FirewallDoSpolicyServiceArrayOutput)
}

// Source address name from available addresses. The structure of `srcaddr` block is documented below.
func (o FirewallDoSpolicyOutput) Srcaddrs() FirewallDoSpolicySrcaddrArrayOutput {
	return o.ApplyT(func(v *FirewallDoSpolicy) FirewallDoSpolicySrcaddrArrayOutput { return v.Srcaddrs }).(FirewallDoSpolicySrcaddrArrayOutput)
}

// Enable/disable this policy. Valid values: `enable`, `disable`.
func (o FirewallDoSpolicyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallDoSpolicy) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallDoSpolicyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallDoSpolicy) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallDoSpolicyArrayOutput struct{ *pulumi.OutputState }

func (FirewallDoSpolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallDoSpolicy)(nil)).Elem()
}

func (o FirewallDoSpolicyArrayOutput) ToFirewallDoSpolicyArrayOutput() FirewallDoSpolicyArrayOutput {
	return o
}

func (o FirewallDoSpolicyArrayOutput) ToFirewallDoSpolicyArrayOutputWithContext(ctx context.Context) FirewallDoSpolicyArrayOutput {
	return o
}

func (o FirewallDoSpolicyArrayOutput) Index(i pulumi.IntInput) FirewallDoSpolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallDoSpolicy {
		return vs[0].([]*FirewallDoSpolicy)[vs[1].(int)]
	}).(FirewallDoSpolicyOutput)
}

type FirewallDoSpolicyMapOutput struct{ *pulumi.OutputState }

func (FirewallDoSpolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallDoSpolicy)(nil)).Elem()
}

func (o FirewallDoSpolicyMapOutput) ToFirewallDoSpolicyMapOutput() FirewallDoSpolicyMapOutput {
	return o
}

func (o FirewallDoSpolicyMapOutput) ToFirewallDoSpolicyMapOutputWithContext(ctx context.Context) FirewallDoSpolicyMapOutput {
	return o
}

func (o FirewallDoSpolicyMapOutput) MapIndex(k pulumi.StringInput) FirewallDoSpolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallDoSpolicy {
		return vs[0].(map[string]*FirewallDoSpolicy)[vs[1].(string)]
	}).(FirewallDoSpolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallDoSpolicyInput)(nil)).Elem(), &FirewallDoSpolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallDoSpolicyArrayInput)(nil)).Elem(), FirewallDoSpolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallDoSpolicyMapInput)(nil)).Elem(), FirewallDoSpolicyMap{})
	pulumi.RegisterOutputType(FirewallDoSpolicyOutput{})
	pulumi.RegisterOutputType(FirewallDoSpolicyArrayOutput{})
	pulumi.RegisterOutputType(FirewallDoSpolicyMapOutput{})
}
