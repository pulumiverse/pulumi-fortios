// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure IPv6 routing policies.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/router"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := router.NewPolicy6(ctx, "trname", &router.Policy6Args{
//				Dst:          pulumi.String("::/0"),
//				EndPort:      pulumi.Int(65535),
//				Gateway:      pulumi.String("::"),
//				InputDevice:  pulumi.String("port1"),
//				OutputDevice: pulumi.String("port3"),
//				Protocol:     pulumi.Int(33),
//				SeqNum:       pulumi.Int(1),
//				Src:          pulumi.String("2001:db8:85a3::8a2e:370:7334/128"),
//				StartPort:    pulumi.Int(1),
//				Status:       pulumi.String("enable"),
//				Tos:          pulumi.String("0x00"),
//				TosMask:      pulumi.String("0x00"),
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
// Router Policy6 can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:router/policy6:Policy6 labelname {{seq_num}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:router/policy6:Policy6 labelname {{seq_num}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Policy6 struct {
	pulumi.CustomResourceState

	// Action of the policy route. Valid values: `deny`, `permit`.
	Action pulumi.StringOutput `pulumi:"action"`
	// Optional comments.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Destination IPv6 prefix.
	Dst pulumi.StringOutput `pulumi:"dst"`
	// Enable/disable negating destination address match. Valid values: `enable`, `disable`.
	DstNegate pulumi.StringOutput `pulumi:"dstNegate"`
	// Destination address name. The structure of `dstaddr` block is documented below.
	Dstaddrs Policy6DstaddrArrayOutput `pulumi:"dstaddrs"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// End destination port number (1 - 65535).
	EndPort pulumi.IntOutput `pulumi:"endPort"`
	// End source port number (1 - 65535).
	EndSourcePort pulumi.IntOutput `pulumi:"endSourcePort"`
	// IPv6 address of the gateway.
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Incoming interface name.
	InputDevice pulumi.StringOutput `pulumi:"inputDevice"`
	// Enable/disable negation of input device match. Valid values: `enable`, `disable`.
	InputDeviceNegate pulumi.StringOutput `pulumi:"inputDeviceNegate"`
	// Custom Destination Internet Service name. The structure of `internetServiceCustom` block is documented below.
	InternetServiceCustoms Policy6InternetServiceCustomArrayOutput `pulumi:"internetServiceCustoms"`
	// Destination Internet Service ID. The structure of `internetServiceId` block is documented below.
	InternetServiceIds Policy6InternetServiceIdArrayOutput `pulumi:"internetServiceIds"`
	// Outgoing interface name.
	OutputDevice pulumi.StringOutput `pulumi:"outputDevice"`
	// Protocol number (0 - 255).
	Protocol pulumi.IntOutput `pulumi:"protocol"`
	// Sequence number.
	SeqNum pulumi.IntOutput `pulumi:"seqNum"`
	// Source IPv6 prefix.
	Src pulumi.StringOutput `pulumi:"src"`
	// Enable/disable negating source address match. Valid values: `enable`, `disable`.
	SrcNegate pulumi.StringOutput `pulumi:"srcNegate"`
	// Source address name. The structure of `srcaddr` block is documented below.
	Srcaddrs Policy6SrcaddrArrayOutput `pulumi:"srcaddrs"`
	// Start destination port number (1 - 65535).
	StartPort pulumi.IntOutput `pulumi:"startPort"`
	// Start source port number (1 - 65535).
	StartSourcePort pulumi.IntOutput `pulumi:"startSourcePort"`
	// Enable/disable this policy route. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Type of service bit pattern.
	Tos pulumi.StringOutput `pulumi:"tos"`
	// Type of service evaluated bits.
	TosMask pulumi.StringOutput `pulumi:"tosMask"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewPolicy6 registers a new resource with the given unique name, arguments, and options.
func NewPolicy6(ctx *pulumi.Context,
	name string, args *Policy6Args, opts ...pulumi.ResourceOption) (*Policy6, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InputDevice == nil {
		return nil, errors.New("invalid value for required argument 'InputDevice'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Policy6
	err := ctx.RegisterResource("fortios:router/policy6:Policy6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy6 gets an existing Policy6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Policy6State, opts ...pulumi.ResourceOption) (*Policy6, error) {
	var resource Policy6
	err := ctx.ReadResource("fortios:router/policy6:Policy6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy6 resources.
type policy6State struct {
	// Action of the policy route. Valid values: `deny`, `permit`.
	Action *string `pulumi:"action"`
	// Optional comments.
	Comments *string `pulumi:"comments"`
	// Destination IPv6 prefix.
	Dst *string `pulumi:"dst"`
	// Enable/disable negating destination address match. Valid values: `enable`, `disable`.
	DstNegate *string `pulumi:"dstNegate"`
	// Destination address name. The structure of `dstaddr` block is documented below.
	Dstaddrs []Policy6Dstaddr `pulumi:"dstaddrs"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// End destination port number (1 - 65535).
	EndPort *int `pulumi:"endPort"`
	// End source port number (1 - 65535).
	EndSourcePort *int `pulumi:"endSourcePort"`
	// IPv6 address of the gateway.
	Gateway *string `pulumi:"gateway"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Incoming interface name.
	InputDevice *string `pulumi:"inputDevice"`
	// Enable/disable negation of input device match. Valid values: `enable`, `disable`.
	InputDeviceNegate *string `pulumi:"inputDeviceNegate"`
	// Custom Destination Internet Service name. The structure of `internetServiceCustom` block is documented below.
	InternetServiceCustoms []Policy6InternetServiceCustom `pulumi:"internetServiceCustoms"`
	// Destination Internet Service ID. The structure of `internetServiceId` block is documented below.
	InternetServiceIds []Policy6InternetServiceId `pulumi:"internetServiceIds"`
	// Outgoing interface name.
	OutputDevice *string `pulumi:"outputDevice"`
	// Protocol number (0 - 255).
	Protocol *int `pulumi:"protocol"`
	// Sequence number.
	SeqNum *int `pulumi:"seqNum"`
	// Source IPv6 prefix.
	Src *string `pulumi:"src"`
	// Enable/disable negating source address match. Valid values: `enable`, `disable`.
	SrcNegate *string `pulumi:"srcNegate"`
	// Source address name. The structure of `srcaddr` block is documented below.
	Srcaddrs []Policy6Srcaddr `pulumi:"srcaddrs"`
	// Start destination port number (1 - 65535).
	StartPort *int `pulumi:"startPort"`
	// Start source port number (1 - 65535).
	StartSourcePort *int `pulumi:"startSourcePort"`
	// Enable/disable this policy route. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Type of service bit pattern.
	Tos *string `pulumi:"tos"`
	// Type of service evaluated bits.
	TosMask *string `pulumi:"tosMask"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Policy6State struct {
	// Action of the policy route. Valid values: `deny`, `permit`.
	Action pulumi.StringPtrInput
	// Optional comments.
	Comments pulumi.StringPtrInput
	// Destination IPv6 prefix.
	Dst pulumi.StringPtrInput
	// Enable/disable negating destination address match. Valid values: `enable`, `disable`.
	DstNegate pulumi.StringPtrInput
	// Destination address name. The structure of `dstaddr` block is documented below.
	Dstaddrs Policy6DstaddrArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// End destination port number (1 - 65535).
	EndPort pulumi.IntPtrInput
	// End source port number (1 - 65535).
	EndSourcePort pulumi.IntPtrInput
	// IPv6 address of the gateway.
	Gateway pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Incoming interface name.
	InputDevice pulumi.StringPtrInput
	// Enable/disable negation of input device match. Valid values: `enable`, `disable`.
	InputDeviceNegate pulumi.StringPtrInput
	// Custom Destination Internet Service name. The structure of `internetServiceCustom` block is documented below.
	InternetServiceCustoms Policy6InternetServiceCustomArrayInput
	// Destination Internet Service ID. The structure of `internetServiceId` block is documented below.
	InternetServiceIds Policy6InternetServiceIdArrayInput
	// Outgoing interface name.
	OutputDevice pulumi.StringPtrInput
	// Protocol number (0 - 255).
	Protocol pulumi.IntPtrInput
	// Sequence number.
	SeqNum pulumi.IntPtrInput
	// Source IPv6 prefix.
	Src pulumi.StringPtrInput
	// Enable/disable negating source address match. Valid values: `enable`, `disable`.
	SrcNegate pulumi.StringPtrInput
	// Source address name. The structure of `srcaddr` block is documented below.
	Srcaddrs Policy6SrcaddrArrayInput
	// Start destination port number (1 - 65535).
	StartPort pulumi.IntPtrInput
	// Start source port number (1 - 65535).
	StartSourcePort pulumi.IntPtrInput
	// Enable/disable this policy route. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Type of service bit pattern.
	Tos pulumi.StringPtrInput
	// Type of service evaluated bits.
	TosMask pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Policy6State) ElementType() reflect.Type {
	return reflect.TypeOf((*policy6State)(nil)).Elem()
}

type policy6Args struct {
	// Action of the policy route. Valid values: `deny`, `permit`.
	Action *string `pulumi:"action"`
	// Optional comments.
	Comments *string `pulumi:"comments"`
	// Destination IPv6 prefix.
	Dst *string `pulumi:"dst"`
	// Enable/disable negating destination address match. Valid values: `enable`, `disable`.
	DstNegate *string `pulumi:"dstNegate"`
	// Destination address name. The structure of `dstaddr` block is documented below.
	Dstaddrs []Policy6Dstaddr `pulumi:"dstaddrs"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// End destination port number (1 - 65535).
	EndPort *int `pulumi:"endPort"`
	// End source port number (1 - 65535).
	EndSourcePort *int `pulumi:"endSourcePort"`
	// IPv6 address of the gateway.
	Gateway *string `pulumi:"gateway"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Incoming interface name.
	InputDevice string `pulumi:"inputDevice"`
	// Enable/disable negation of input device match. Valid values: `enable`, `disable`.
	InputDeviceNegate *string `pulumi:"inputDeviceNegate"`
	// Custom Destination Internet Service name. The structure of `internetServiceCustom` block is documented below.
	InternetServiceCustoms []Policy6InternetServiceCustom `pulumi:"internetServiceCustoms"`
	// Destination Internet Service ID. The structure of `internetServiceId` block is documented below.
	InternetServiceIds []Policy6InternetServiceId `pulumi:"internetServiceIds"`
	// Outgoing interface name.
	OutputDevice *string `pulumi:"outputDevice"`
	// Protocol number (0 - 255).
	Protocol *int `pulumi:"protocol"`
	// Sequence number.
	SeqNum *int `pulumi:"seqNum"`
	// Source IPv6 prefix.
	Src *string `pulumi:"src"`
	// Enable/disable negating source address match. Valid values: `enable`, `disable`.
	SrcNegate *string `pulumi:"srcNegate"`
	// Source address name. The structure of `srcaddr` block is documented below.
	Srcaddrs []Policy6Srcaddr `pulumi:"srcaddrs"`
	// Start destination port number (1 - 65535).
	StartPort *int `pulumi:"startPort"`
	// Start source port number (1 - 65535).
	StartSourcePort *int `pulumi:"startSourcePort"`
	// Enable/disable this policy route. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Type of service bit pattern.
	Tos *string `pulumi:"tos"`
	// Type of service evaluated bits.
	TosMask *string `pulumi:"tosMask"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Policy6 resource.
type Policy6Args struct {
	// Action of the policy route. Valid values: `deny`, `permit`.
	Action pulumi.StringPtrInput
	// Optional comments.
	Comments pulumi.StringPtrInput
	// Destination IPv6 prefix.
	Dst pulumi.StringPtrInput
	// Enable/disable negating destination address match. Valid values: `enable`, `disable`.
	DstNegate pulumi.StringPtrInput
	// Destination address name. The structure of `dstaddr` block is documented below.
	Dstaddrs Policy6DstaddrArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// End destination port number (1 - 65535).
	EndPort pulumi.IntPtrInput
	// End source port number (1 - 65535).
	EndSourcePort pulumi.IntPtrInput
	// IPv6 address of the gateway.
	Gateway pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Incoming interface name.
	InputDevice pulumi.StringInput
	// Enable/disable negation of input device match. Valid values: `enable`, `disable`.
	InputDeviceNegate pulumi.StringPtrInput
	// Custom Destination Internet Service name. The structure of `internetServiceCustom` block is documented below.
	InternetServiceCustoms Policy6InternetServiceCustomArrayInput
	// Destination Internet Service ID. The structure of `internetServiceId` block is documented below.
	InternetServiceIds Policy6InternetServiceIdArrayInput
	// Outgoing interface name.
	OutputDevice pulumi.StringPtrInput
	// Protocol number (0 - 255).
	Protocol pulumi.IntPtrInput
	// Sequence number.
	SeqNum pulumi.IntPtrInput
	// Source IPv6 prefix.
	Src pulumi.StringPtrInput
	// Enable/disable negating source address match. Valid values: `enable`, `disable`.
	SrcNegate pulumi.StringPtrInput
	// Source address name. The structure of `srcaddr` block is documented below.
	Srcaddrs Policy6SrcaddrArrayInput
	// Start destination port number (1 - 65535).
	StartPort pulumi.IntPtrInput
	// Start source port number (1 - 65535).
	StartSourcePort pulumi.IntPtrInput
	// Enable/disable this policy route. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Type of service bit pattern.
	Tos pulumi.StringPtrInput
	// Type of service evaluated bits.
	TosMask pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Policy6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*policy6Args)(nil)).Elem()
}

type Policy6Input interface {
	pulumi.Input

	ToPolicy6Output() Policy6Output
	ToPolicy6OutputWithContext(ctx context.Context) Policy6Output
}

func (*Policy6) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy6)(nil)).Elem()
}

func (i *Policy6) ToPolicy6Output() Policy6Output {
	return i.ToPolicy6OutputWithContext(context.Background())
}

func (i *Policy6) ToPolicy6OutputWithContext(ctx context.Context) Policy6Output {
	return pulumi.ToOutputWithContext(ctx, i).(Policy6Output)
}

// Policy6ArrayInput is an input type that accepts Policy6Array and Policy6ArrayOutput values.
// You can construct a concrete instance of `Policy6ArrayInput` via:
//
//	Policy6Array{ Policy6Args{...} }
type Policy6ArrayInput interface {
	pulumi.Input

	ToPolicy6ArrayOutput() Policy6ArrayOutput
	ToPolicy6ArrayOutputWithContext(context.Context) Policy6ArrayOutput
}

type Policy6Array []Policy6Input

func (Policy6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy6)(nil)).Elem()
}

func (i Policy6Array) ToPolicy6ArrayOutput() Policy6ArrayOutput {
	return i.ToPolicy6ArrayOutputWithContext(context.Background())
}

func (i Policy6Array) ToPolicy6ArrayOutputWithContext(ctx context.Context) Policy6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Policy6ArrayOutput)
}

// Policy6MapInput is an input type that accepts Policy6Map and Policy6MapOutput values.
// You can construct a concrete instance of `Policy6MapInput` via:
//
//	Policy6Map{ "key": Policy6Args{...} }
type Policy6MapInput interface {
	pulumi.Input

	ToPolicy6MapOutput() Policy6MapOutput
	ToPolicy6MapOutputWithContext(context.Context) Policy6MapOutput
}

type Policy6Map map[string]Policy6Input

func (Policy6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy6)(nil)).Elem()
}

func (i Policy6Map) ToPolicy6MapOutput() Policy6MapOutput {
	return i.ToPolicy6MapOutputWithContext(context.Background())
}

func (i Policy6Map) ToPolicy6MapOutputWithContext(ctx context.Context) Policy6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Policy6MapOutput)
}

type Policy6Output struct{ *pulumi.OutputState }

func (Policy6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy6)(nil)).Elem()
}

func (o Policy6Output) ToPolicy6Output() Policy6Output {
	return o
}

func (o Policy6Output) ToPolicy6OutputWithContext(ctx context.Context) Policy6Output {
	return o
}

// Action of the policy route. Valid values: `deny`, `permit`.
func (o Policy6Output) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy6) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// Optional comments.
func (o Policy6Output) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy6) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// Destination IPv6 prefix.
func (o Policy6Output) Dst() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy6) pulumi.StringOutput { return v.Dst }).(pulumi.StringOutput)
}

// Enable/disable negating destination address match. Valid values: `enable`, `disable`.
func (o Policy6Output) DstNegate() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy6) pulumi.StringOutput { return v.DstNegate }).(pulumi.StringOutput)
}

// Destination address name. The structure of `dstaddr` block is documented below.
func (o Policy6Output) Dstaddrs() Policy6DstaddrArrayOutput {
	return o.ApplyT(func(v *Policy6) Policy6DstaddrArrayOutput { return v.Dstaddrs }).(Policy6DstaddrArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o Policy6Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy6) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// End destination port number (1 - 65535).
func (o Policy6Output) EndPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Policy6) pulumi.IntOutput { return v.EndPort }).(pulumi.IntOutput)
}

// End source port number (1 - 65535).
func (o Policy6Output) EndSourcePort() pulumi.IntOutput {
	return o.ApplyT(func(v *Policy6) pulumi.IntOutput { return v.EndSourcePort }).(pulumi.IntOutput)
}

// IPv6 address of the gateway.
func (o Policy6Output) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy6) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o Policy6Output) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy6) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Incoming interface name.
func (o Policy6Output) InputDevice() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy6) pulumi.StringOutput { return v.InputDevice }).(pulumi.StringOutput)
}

// Enable/disable negation of input device match. Valid values: `enable`, `disable`.
func (o Policy6Output) InputDeviceNegate() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy6) pulumi.StringOutput { return v.InputDeviceNegate }).(pulumi.StringOutput)
}

// Custom Destination Internet Service name. The structure of `internetServiceCustom` block is documented below.
func (o Policy6Output) InternetServiceCustoms() Policy6InternetServiceCustomArrayOutput {
	return o.ApplyT(func(v *Policy6) Policy6InternetServiceCustomArrayOutput { return v.InternetServiceCustoms }).(Policy6InternetServiceCustomArrayOutput)
}

// Destination Internet Service ID. The structure of `internetServiceId` block is documented below.
func (o Policy6Output) InternetServiceIds() Policy6InternetServiceIdArrayOutput {
	return o.ApplyT(func(v *Policy6) Policy6InternetServiceIdArrayOutput { return v.InternetServiceIds }).(Policy6InternetServiceIdArrayOutput)
}

// Outgoing interface name.
func (o Policy6Output) OutputDevice() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy6) pulumi.StringOutput { return v.OutputDevice }).(pulumi.StringOutput)
}

// Protocol number (0 - 255).
func (o Policy6Output) Protocol() pulumi.IntOutput {
	return o.ApplyT(func(v *Policy6) pulumi.IntOutput { return v.Protocol }).(pulumi.IntOutput)
}

// Sequence number.
func (o Policy6Output) SeqNum() pulumi.IntOutput {
	return o.ApplyT(func(v *Policy6) pulumi.IntOutput { return v.SeqNum }).(pulumi.IntOutput)
}

// Source IPv6 prefix.
func (o Policy6Output) Src() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy6) pulumi.StringOutput { return v.Src }).(pulumi.StringOutput)
}

// Enable/disable negating source address match. Valid values: `enable`, `disable`.
func (o Policy6Output) SrcNegate() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy6) pulumi.StringOutput { return v.SrcNegate }).(pulumi.StringOutput)
}

// Source address name. The structure of `srcaddr` block is documented below.
func (o Policy6Output) Srcaddrs() Policy6SrcaddrArrayOutput {
	return o.ApplyT(func(v *Policy6) Policy6SrcaddrArrayOutput { return v.Srcaddrs }).(Policy6SrcaddrArrayOutput)
}

// Start destination port number (1 - 65535).
func (o Policy6Output) StartPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Policy6) pulumi.IntOutput { return v.StartPort }).(pulumi.IntOutput)
}

// Start source port number (1 - 65535).
func (o Policy6Output) StartSourcePort() pulumi.IntOutput {
	return o.ApplyT(func(v *Policy6) pulumi.IntOutput { return v.StartSourcePort }).(pulumi.IntOutput)
}

// Enable/disable this policy route. Valid values: `enable`, `disable`.
func (o Policy6Output) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy6) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Type of service bit pattern.
func (o Policy6Output) Tos() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy6) pulumi.StringOutput { return v.Tos }).(pulumi.StringOutput)
}

// Type of service evaluated bits.
func (o Policy6Output) TosMask() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy6) pulumi.StringOutput { return v.TosMask }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Policy6Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy6) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type Policy6ArrayOutput struct{ *pulumi.OutputState }

func (Policy6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy6)(nil)).Elem()
}

func (o Policy6ArrayOutput) ToPolicy6ArrayOutput() Policy6ArrayOutput {
	return o
}

func (o Policy6ArrayOutput) ToPolicy6ArrayOutputWithContext(ctx context.Context) Policy6ArrayOutput {
	return o
}

func (o Policy6ArrayOutput) Index(i pulumi.IntInput) Policy6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Policy6 {
		return vs[0].([]*Policy6)[vs[1].(int)]
	}).(Policy6Output)
}

type Policy6MapOutput struct{ *pulumi.OutputState }

func (Policy6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy6)(nil)).Elem()
}

func (o Policy6MapOutput) ToPolicy6MapOutput() Policy6MapOutput {
	return o
}

func (o Policy6MapOutput) ToPolicy6MapOutputWithContext(ctx context.Context) Policy6MapOutput {
	return o
}

func (o Policy6MapOutput) MapIndex(k pulumi.StringInput) Policy6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Policy6 {
		return vs[0].(map[string]*Policy6)[vs[1].(string)]
	}).(Policy6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Policy6Input)(nil)).Elem(), &Policy6{})
	pulumi.RegisterInputType(reflect.TypeOf((*Policy6ArrayInput)(nil)).Elem(), Policy6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*Policy6MapInput)(nil)).Elem(), Policy6Map{})
	pulumi.RegisterOutputType(Policy6Output{})
	pulumi.RegisterOutputType(Policy6ArrayOutput{})
	pulumi.RegisterOutputType(Policy6MapOutput{})
}
