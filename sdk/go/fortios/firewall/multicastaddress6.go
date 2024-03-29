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

// Configure IPv6 multicast address.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/firewall"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firewall.NewMulticastaddress6(ctx, "trname", &firewall.Multicastaddress6Args{
//				Color:      pulumi.Int(0),
//				Ip6:        pulumi.String("ff02::1:ff0e:8c6c/128"),
//				Visibility: pulumi.String("enable"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Firewall MulticastAddress6 can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/multicastaddress6:Multicastaddress6 labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/multicastaddress6:Multicastaddress6 labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Multicastaddress6 struct {
	pulumi.CustomResourceState

	// Color of icon on the GUI.
	Color pulumi.IntOutput `pulumi:"color"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
	Ip6 pulumi.StringOutput `pulumi:"ip6"`
	// IPv6 multicast address name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings Multicastaddress6TaggingArrayOutput `pulumi:"taggings"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable visibility of the IPv6 multicast address on the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewMulticastaddress6 registers a new resource with the given unique name, arguments, and options.
func NewMulticastaddress6(ctx *pulumi.Context,
	name string, args *Multicastaddress6Args, opts ...pulumi.ResourceOption) (*Multicastaddress6, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ip6 == nil {
		return nil, errors.New("invalid value for required argument 'Ip6'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Multicastaddress6
	err := ctx.RegisterResource("fortios:firewall/multicastaddress6:Multicastaddress6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMulticastaddress6 gets an existing Multicastaddress6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMulticastaddress6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Multicastaddress6State, opts ...pulumi.ResourceOption) (*Multicastaddress6, error) {
	var resource Multicastaddress6
	err := ctx.ReadResource("fortios:firewall/multicastaddress6:Multicastaddress6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Multicastaddress6 resources.
type multicastaddress6State struct {
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
	Ip6 *string `pulumi:"ip6"`
	// IPv6 multicast address name.
	Name *string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []Multicastaddress6Tagging `pulumi:"taggings"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable visibility of the IPv6 multicast address on the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

type Multicastaddress6State struct {
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
	Ip6 pulumi.StringPtrInput
	// IPv6 multicast address name.
	Name pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings Multicastaddress6TaggingArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable visibility of the IPv6 multicast address on the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (Multicastaddress6State) ElementType() reflect.Type {
	return reflect.TypeOf((*multicastaddress6State)(nil)).Elem()
}

type multicastaddress6Args struct {
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
	Ip6 string `pulumi:"ip6"`
	// IPv6 multicast address name.
	Name *string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []Multicastaddress6Tagging `pulumi:"taggings"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable visibility of the IPv6 multicast address on the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

// The set of arguments for constructing a Multicastaddress6 resource.
type Multicastaddress6Args struct {
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
	Ip6 pulumi.StringInput
	// IPv6 multicast address name.
	Name pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings Multicastaddress6TaggingArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable visibility of the IPv6 multicast address on the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (Multicastaddress6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*multicastaddress6Args)(nil)).Elem()
}

type Multicastaddress6Input interface {
	pulumi.Input

	ToMulticastaddress6Output() Multicastaddress6Output
	ToMulticastaddress6OutputWithContext(ctx context.Context) Multicastaddress6Output
}

func (*Multicastaddress6) ElementType() reflect.Type {
	return reflect.TypeOf((**Multicastaddress6)(nil)).Elem()
}

func (i *Multicastaddress6) ToMulticastaddress6Output() Multicastaddress6Output {
	return i.ToMulticastaddress6OutputWithContext(context.Background())
}

func (i *Multicastaddress6) ToMulticastaddress6OutputWithContext(ctx context.Context) Multicastaddress6Output {
	return pulumi.ToOutputWithContext(ctx, i).(Multicastaddress6Output)
}

// Multicastaddress6ArrayInput is an input type that accepts Multicastaddress6Array and Multicastaddress6ArrayOutput values.
// You can construct a concrete instance of `Multicastaddress6ArrayInput` via:
//
//	Multicastaddress6Array{ Multicastaddress6Args{...} }
type Multicastaddress6ArrayInput interface {
	pulumi.Input

	ToMulticastaddress6ArrayOutput() Multicastaddress6ArrayOutput
	ToMulticastaddress6ArrayOutputWithContext(context.Context) Multicastaddress6ArrayOutput
}

type Multicastaddress6Array []Multicastaddress6Input

func (Multicastaddress6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Multicastaddress6)(nil)).Elem()
}

func (i Multicastaddress6Array) ToMulticastaddress6ArrayOutput() Multicastaddress6ArrayOutput {
	return i.ToMulticastaddress6ArrayOutputWithContext(context.Background())
}

func (i Multicastaddress6Array) ToMulticastaddress6ArrayOutputWithContext(ctx context.Context) Multicastaddress6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Multicastaddress6ArrayOutput)
}

// Multicastaddress6MapInput is an input type that accepts Multicastaddress6Map and Multicastaddress6MapOutput values.
// You can construct a concrete instance of `Multicastaddress6MapInput` via:
//
//	Multicastaddress6Map{ "key": Multicastaddress6Args{...} }
type Multicastaddress6MapInput interface {
	pulumi.Input

	ToMulticastaddress6MapOutput() Multicastaddress6MapOutput
	ToMulticastaddress6MapOutputWithContext(context.Context) Multicastaddress6MapOutput
}

type Multicastaddress6Map map[string]Multicastaddress6Input

func (Multicastaddress6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Multicastaddress6)(nil)).Elem()
}

func (i Multicastaddress6Map) ToMulticastaddress6MapOutput() Multicastaddress6MapOutput {
	return i.ToMulticastaddress6MapOutputWithContext(context.Background())
}

func (i Multicastaddress6Map) ToMulticastaddress6MapOutputWithContext(ctx context.Context) Multicastaddress6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Multicastaddress6MapOutput)
}

type Multicastaddress6Output struct{ *pulumi.OutputState }

func (Multicastaddress6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**Multicastaddress6)(nil)).Elem()
}

func (o Multicastaddress6Output) ToMulticastaddress6Output() Multicastaddress6Output {
	return o
}

func (o Multicastaddress6Output) ToMulticastaddress6OutputWithContext(ctx context.Context) Multicastaddress6Output {
	return o
}

// Color of icon on the GUI.
func (o Multicastaddress6Output) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *Multicastaddress6) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o Multicastaddress6Output) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Multicastaddress6) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o Multicastaddress6Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Multicastaddress6) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o Multicastaddress6Output) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Multicastaddress6) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
func (o Multicastaddress6Output) Ip6() pulumi.StringOutput {
	return o.ApplyT(func(v *Multicastaddress6) pulumi.StringOutput { return v.Ip6 }).(pulumi.StringOutput)
}

// IPv6 multicast address name.
func (o Multicastaddress6Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Multicastaddress6) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Config object tagging. The structure of `tagging` block is documented below.
func (o Multicastaddress6Output) Taggings() Multicastaddress6TaggingArrayOutput {
	return o.ApplyT(func(v *Multicastaddress6) Multicastaddress6TaggingArrayOutput { return v.Taggings }).(Multicastaddress6TaggingArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Multicastaddress6Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Multicastaddress6) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable visibility of the IPv6 multicast address on the GUI. Valid values: `enable`, `disable`.
func (o Multicastaddress6Output) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *Multicastaddress6) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

type Multicastaddress6ArrayOutput struct{ *pulumi.OutputState }

func (Multicastaddress6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Multicastaddress6)(nil)).Elem()
}

func (o Multicastaddress6ArrayOutput) ToMulticastaddress6ArrayOutput() Multicastaddress6ArrayOutput {
	return o
}

func (o Multicastaddress6ArrayOutput) ToMulticastaddress6ArrayOutputWithContext(ctx context.Context) Multicastaddress6ArrayOutput {
	return o
}

func (o Multicastaddress6ArrayOutput) Index(i pulumi.IntInput) Multicastaddress6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Multicastaddress6 {
		return vs[0].([]*Multicastaddress6)[vs[1].(int)]
	}).(Multicastaddress6Output)
}

type Multicastaddress6MapOutput struct{ *pulumi.OutputState }

func (Multicastaddress6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Multicastaddress6)(nil)).Elem()
}

func (o Multicastaddress6MapOutput) ToMulticastaddress6MapOutput() Multicastaddress6MapOutput {
	return o
}

func (o Multicastaddress6MapOutput) ToMulticastaddress6MapOutputWithContext(ctx context.Context) Multicastaddress6MapOutput {
	return o
}

func (o Multicastaddress6MapOutput) MapIndex(k pulumi.StringInput) Multicastaddress6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Multicastaddress6 {
		return vs[0].(map[string]*Multicastaddress6)[vs[1].(string)]
	}).(Multicastaddress6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Multicastaddress6Input)(nil)).Elem(), &Multicastaddress6{})
	pulumi.RegisterInputType(reflect.TypeOf((*Multicastaddress6ArrayInput)(nil)).Elem(), Multicastaddress6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*Multicastaddress6MapInput)(nil)).Elem(), Multicastaddress6Map{})
	pulumi.RegisterOutputType(Multicastaddress6Output{})
	pulumi.RegisterOutputType(Multicastaddress6ArrayOutput{})
	pulumi.RegisterOutputType(Multicastaddress6MapOutput{})
}
