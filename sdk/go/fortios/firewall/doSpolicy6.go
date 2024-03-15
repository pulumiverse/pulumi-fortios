// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure IPv6 DoS policies.
//
// ## Import
//
// Firewall DosPolicy6 can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/doSpolicy6:DoSpolicy6 labelname {{policyid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/doSpolicy6:DoSpolicy6 labelname {{policyid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type DoSpolicy6 struct {
	pulumi.CustomResourceState

	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies DoSpolicy6AnomalyArrayOutput `pulumi:"anomalies"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs DoSpolicy6DstaddrArrayOutput `pulumi:"dstaddrs"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Incoming interface name from available interfaces.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Policy ID.
	Policyid pulumi.IntOutput `pulumi:"policyid"`
	// Service object from available options. The structure of `service` block is documented below.
	Services DoSpolicy6ServiceArrayOutput `pulumi:"services"`
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs DoSpolicy6SrcaddrArrayOutput `pulumi:"srcaddrs"`
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDoSpolicy6 registers a new resource with the given unique name, arguments, and options.
func NewDoSpolicy6(ctx *pulumi.Context,
	name string, args *DoSpolicy6Args, opts ...pulumi.ResourceOption) (*DoSpolicy6, error) {
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
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DoSpolicy6
	err := ctx.RegisterResource("fortios:firewall/doSpolicy6:DoSpolicy6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDoSpolicy6 gets an existing DoSpolicy6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDoSpolicy6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DoSpolicy6State, opts ...pulumi.ResourceOption) (*DoSpolicy6, error) {
	var resource DoSpolicy6
	err := ctx.ReadResource("fortios:firewall/doSpolicy6:DoSpolicy6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DoSpolicy6 resources.
type doSpolicy6State struct {
	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies []DoSpolicy6Anomaly `pulumi:"anomalies"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs []DoSpolicy6Dstaddr `pulumi:"dstaddrs"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Incoming interface name from available interfaces.
	Interface *string `pulumi:"interface"`
	// Policy name.
	Name *string `pulumi:"name"`
	// Policy ID.
	Policyid *int `pulumi:"policyid"`
	// Service object from available options. The structure of `service` block is documented below.
	Services []DoSpolicy6Service `pulumi:"services"`
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs []DoSpolicy6Srcaddr `pulumi:"srcaddrs"`
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DoSpolicy6State struct {
	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies DoSpolicy6AnomalyArrayInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs DoSpolicy6DstaddrArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Incoming interface name from available interfaces.
	Interface pulumi.StringPtrInput
	// Policy name.
	Name pulumi.StringPtrInput
	// Policy ID.
	Policyid pulumi.IntPtrInput
	// Service object from available options. The structure of `service` block is documented below.
	Services DoSpolicy6ServiceArrayInput
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs DoSpolicy6SrcaddrArrayInput
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DoSpolicy6State) ElementType() reflect.Type {
	return reflect.TypeOf((*doSpolicy6State)(nil)).Elem()
}

type doSpolicy6Args struct {
	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies []DoSpolicy6Anomaly `pulumi:"anomalies"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs []DoSpolicy6Dstaddr `pulumi:"dstaddrs"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Incoming interface name from available interfaces.
	Interface string `pulumi:"interface"`
	// Policy name.
	Name *string `pulumi:"name"`
	// Policy ID.
	Policyid *int `pulumi:"policyid"`
	// Service object from available options. The structure of `service` block is documented below.
	Services []DoSpolicy6Service `pulumi:"services"`
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs []DoSpolicy6Srcaddr `pulumi:"srcaddrs"`
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a DoSpolicy6 resource.
type DoSpolicy6Args struct {
	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies DoSpolicy6AnomalyArrayInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs DoSpolicy6DstaddrArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Incoming interface name from available interfaces.
	Interface pulumi.StringInput
	// Policy name.
	Name pulumi.StringPtrInput
	// Policy ID.
	Policyid pulumi.IntPtrInput
	// Service object from available options. The structure of `service` block is documented below.
	Services DoSpolicy6ServiceArrayInput
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs DoSpolicy6SrcaddrArrayInput
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DoSpolicy6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*doSpolicy6Args)(nil)).Elem()
}

type DoSpolicy6Input interface {
	pulumi.Input

	ToDoSpolicy6Output() DoSpolicy6Output
	ToDoSpolicy6OutputWithContext(ctx context.Context) DoSpolicy6Output
}

func (*DoSpolicy6) ElementType() reflect.Type {
	return reflect.TypeOf((**DoSpolicy6)(nil)).Elem()
}

func (i *DoSpolicy6) ToDoSpolicy6Output() DoSpolicy6Output {
	return i.ToDoSpolicy6OutputWithContext(context.Background())
}

func (i *DoSpolicy6) ToDoSpolicy6OutputWithContext(ctx context.Context) DoSpolicy6Output {
	return pulumi.ToOutputWithContext(ctx, i).(DoSpolicy6Output)
}

// DoSpolicy6ArrayInput is an input type that accepts DoSpolicy6Array and DoSpolicy6ArrayOutput values.
// You can construct a concrete instance of `DoSpolicy6ArrayInput` via:
//
//	DoSpolicy6Array{ DoSpolicy6Args{...} }
type DoSpolicy6ArrayInput interface {
	pulumi.Input

	ToDoSpolicy6ArrayOutput() DoSpolicy6ArrayOutput
	ToDoSpolicy6ArrayOutputWithContext(context.Context) DoSpolicy6ArrayOutput
}

type DoSpolicy6Array []DoSpolicy6Input

func (DoSpolicy6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DoSpolicy6)(nil)).Elem()
}

func (i DoSpolicy6Array) ToDoSpolicy6ArrayOutput() DoSpolicy6ArrayOutput {
	return i.ToDoSpolicy6ArrayOutputWithContext(context.Background())
}

func (i DoSpolicy6Array) ToDoSpolicy6ArrayOutputWithContext(ctx context.Context) DoSpolicy6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DoSpolicy6ArrayOutput)
}

// DoSpolicy6MapInput is an input type that accepts DoSpolicy6Map and DoSpolicy6MapOutput values.
// You can construct a concrete instance of `DoSpolicy6MapInput` via:
//
//	DoSpolicy6Map{ "key": DoSpolicy6Args{...} }
type DoSpolicy6MapInput interface {
	pulumi.Input

	ToDoSpolicy6MapOutput() DoSpolicy6MapOutput
	ToDoSpolicy6MapOutputWithContext(context.Context) DoSpolicy6MapOutput
}

type DoSpolicy6Map map[string]DoSpolicy6Input

func (DoSpolicy6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DoSpolicy6)(nil)).Elem()
}

func (i DoSpolicy6Map) ToDoSpolicy6MapOutput() DoSpolicy6MapOutput {
	return i.ToDoSpolicy6MapOutputWithContext(context.Background())
}

func (i DoSpolicy6Map) ToDoSpolicy6MapOutputWithContext(ctx context.Context) DoSpolicy6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DoSpolicy6MapOutput)
}

type DoSpolicy6Output struct{ *pulumi.OutputState }

func (DoSpolicy6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**DoSpolicy6)(nil)).Elem()
}

func (o DoSpolicy6Output) ToDoSpolicy6Output() DoSpolicy6Output {
	return o
}

func (o DoSpolicy6Output) ToDoSpolicy6OutputWithContext(ctx context.Context) DoSpolicy6Output {
	return o
}

// Anomaly name. The structure of `anomaly` block is documented below.
func (o DoSpolicy6Output) Anomalies() DoSpolicy6AnomalyArrayOutput {
	return o.ApplyT(func(v *DoSpolicy6) DoSpolicy6AnomalyArrayOutput { return v.Anomalies }).(DoSpolicy6AnomalyArrayOutput)
}

// Comment.
func (o DoSpolicy6Output) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DoSpolicy6) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
func (o DoSpolicy6Output) Dstaddrs() DoSpolicy6DstaddrArrayOutput {
	return o.ApplyT(func(v *DoSpolicy6) DoSpolicy6DstaddrArrayOutput { return v.Dstaddrs }).(DoSpolicy6DstaddrArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o DoSpolicy6Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DoSpolicy6) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Incoming interface name from available interfaces.
func (o DoSpolicy6Output) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *DoSpolicy6) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Policy name.
func (o DoSpolicy6Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DoSpolicy6) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Policy ID.
func (o DoSpolicy6Output) Policyid() pulumi.IntOutput {
	return o.ApplyT(func(v *DoSpolicy6) pulumi.IntOutput { return v.Policyid }).(pulumi.IntOutput)
}

// Service object from available options. The structure of `service` block is documented below.
func (o DoSpolicy6Output) Services() DoSpolicy6ServiceArrayOutput {
	return o.ApplyT(func(v *DoSpolicy6) DoSpolicy6ServiceArrayOutput { return v.Services }).(DoSpolicy6ServiceArrayOutput)
}

// Source address name from available addresses. The structure of `srcaddr` block is documented below.
func (o DoSpolicy6Output) Srcaddrs() DoSpolicy6SrcaddrArrayOutput {
	return o.ApplyT(func(v *DoSpolicy6) DoSpolicy6SrcaddrArrayOutput { return v.Srcaddrs }).(DoSpolicy6SrcaddrArrayOutput)
}

// Enable/disable this policy. Valid values: `enable`, `disable`.
func (o DoSpolicy6Output) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DoSpolicy6) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DoSpolicy6Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DoSpolicy6) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type DoSpolicy6ArrayOutput struct{ *pulumi.OutputState }

func (DoSpolicy6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DoSpolicy6)(nil)).Elem()
}

func (o DoSpolicy6ArrayOutput) ToDoSpolicy6ArrayOutput() DoSpolicy6ArrayOutput {
	return o
}

func (o DoSpolicy6ArrayOutput) ToDoSpolicy6ArrayOutputWithContext(ctx context.Context) DoSpolicy6ArrayOutput {
	return o
}

func (o DoSpolicy6ArrayOutput) Index(i pulumi.IntInput) DoSpolicy6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DoSpolicy6 {
		return vs[0].([]*DoSpolicy6)[vs[1].(int)]
	}).(DoSpolicy6Output)
}

type DoSpolicy6MapOutput struct{ *pulumi.OutputState }

func (DoSpolicy6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DoSpolicy6)(nil)).Elem()
}

func (o DoSpolicy6MapOutput) ToDoSpolicy6MapOutput() DoSpolicy6MapOutput {
	return o
}

func (o DoSpolicy6MapOutput) ToDoSpolicy6MapOutputWithContext(ctx context.Context) DoSpolicy6MapOutput {
	return o
}

func (o DoSpolicy6MapOutput) MapIndex(k pulumi.StringInput) DoSpolicy6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DoSpolicy6 {
		return vs[0].(map[string]*DoSpolicy6)[vs[1].(string)]
	}).(DoSpolicy6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DoSpolicy6Input)(nil)).Elem(), &DoSpolicy6{})
	pulumi.RegisterInputType(reflect.TypeOf((*DoSpolicy6ArrayInput)(nil)).Elem(), DoSpolicy6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*DoSpolicy6MapInput)(nil)).Elem(), DoSpolicy6Map{})
	pulumi.RegisterOutputType(DoSpolicy6Output{})
	pulumi.RegisterOutputType(DoSpolicy6ArrayOutput{})
	pulumi.RegisterOutputType(DoSpolicy6MapOutput{})
}
