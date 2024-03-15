// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure IPv6 firewall addresses.
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
//			_, err := firewall.NewAddress6(ctx, "trname", &firewall.Address6Args{
//				CacheTtl:   pulumi.Int(0),
//				Color:      pulumi.Int(0),
//				EndIp:      pulumi.String("::"),
//				Host:       pulumi.String("fdff:ffff::"),
//				HostType:   pulumi.String("any"),
//				Ip6:        pulumi.String("fdff:ffff::/120"),
//				StartIp:    pulumi.String("fdff:ffff::"),
//				Type:       pulumi.String("ipprefix"),
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
// Firewall Address6 can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/address6:Address6 labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/address6:Address6 labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Address6 struct {
	pulumi.CustomResourceState

	// Minimal TTL of individual IPv6 addresses in FQDN cache.
	CacheTtl pulumi.IntOutput `pulumi:"cacheTtl"`
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color pulumi.IntOutput `pulumi:"color"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// IPv6 addresses associated to a specific country.
	Country pulumi.StringOutput `pulumi:"country"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
	EndIp pulumi.StringOutput `pulumi:"endIp"`
	// Last MAC address in the range.
	EndMac pulumi.StringOutput `pulumi:"endMac"`
	// Endpoint group name.
	EpgName pulumi.StringOutput `pulumi:"epgName"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringOutput `pulumi:"fabricObject"`
	// Fully qualified domain name.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// Host Address.
	Host pulumi.StringOutput `pulumi:"host"`
	// Host type. Valid values: `any`, `specific`.
	HostType pulumi.StringOutput `pulumi:"hostType"`
	// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
	Ip6 pulumi.StringOutput `pulumi:"ip6"`
	// IP address list. The structure of `list` block is documented below.
	Lists Address6ListArrayOutput `pulumi:"lists"`
	// Multiple MAC address ranges. The structure of `macaddr` block is documented below.
	Macaddrs Address6MacaddrArrayOutput `pulumi:"macaddrs"`
	// Address name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Object ID for NSX.
	ObjId pulumi.StringPtrOutput `pulumi:"objId"`
	// SDN.
	Sdn pulumi.StringOutput `pulumi:"sdn"`
	// SDN Tag.
	SdnTag pulumi.StringOutput `pulumi:"sdnTag"`
	// First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
	StartIp pulumi.StringOutput `pulumi:"startIp"`
	// First MAC address in the range.
	StartMac pulumi.StringOutput `pulumi:"startMac"`
	// IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
	SubnetSegments Address6SubnetSegmentArrayOutput `pulumi:"subnetSegments"`
	// Config object tagging The structure of `tagging` block is documented below.
	Taggings Address6TaggingArrayOutput `pulumi:"taggings"`
	// IPv6 address template.
	Template pulumi.StringOutput `pulumi:"template"`
	// Tenant.
	Tenant pulumi.StringOutput `pulumi:"tenant"`
	// Type of IPv6 address object (default = ipprefix).
	Type pulumi.StringOutput `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable the visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewAddress6 registers a new resource with the given unique name, arguments, and options.
func NewAddress6(ctx *pulumi.Context,
	name string, args *Address6Args, opts ...pulumi.ResourceOption) (*Address6, error) {
	if args == nil {
		args = &Address6Args{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Address6
	err := ctx.RegisterResource("fortios:firewall/address6:Address6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAddress6 gets an existing Address6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAddress6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Address6State, opts ...pulumi.ResourceOption) (*Address6, error) {
	var resource Address6
	err := ctx.ReadResource("fortios:firewall/address6:Address6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Address6 resources.
type address6State struct {
	// Minimal TTL of individual IPv6 addresses in FQDN cache.
	CacheTtl *int `pulumi:"cacheTtl"`
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// IPv6 addresses associated to a specific country.
	Country *string `pulumi:"country"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
	EndIp *string `pulumi:"endIp"`
	// Last MAC address in the range.
	EndMac *string `pulumi:"endMac"`
	// Endpoint group name.
	EpgName *string `pulumi:"epgName"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// Fully qualified domain name.
	Fqdn *string `pulumi:"fqdn"`
	// Host Address.
	Host *string `pulumi:"host"`
	// Host type. Valid values: `any`, `specific`.
	HostType *string `pulumi:"hostType"`
	// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
	Ip6 *string `pulumi:"ip6"`
	// IP address list. The structure of `list` block is documented below.
	Lists []Address6List `pulumi:"lists"`
	// Multiple MAC address ranges. The structure of `macaddr` block is documented below.
	Macaddrs []Address6Macaddr `pulumi:"macaddrs"`
	// Address name.
	Name *string `pulumi:"name"`
	// Object ID for NSX.
	ObjId *string `pulumi:"objId"`
	// SDN.
	Sdn *string `pulumi:"sdn"`
	// SDN Tag.
	SdnTag *string `pulumi:"sdnTag"`
	// First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
	StartIp *string `pulumi:"startIp"`
	// First MAC address in the range.
	StartMac *string `pulumi:"startMac"`
	// IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
	SubnetSegments []Address6SubnetSegment `pulumi:"subnetSegments"`
	// Config object tagging The structure of `tagging` block is documented below.
	Taggings []Address6Tagging `pulumi:"taggings"`
	// IPv6 address template.
	Template *string `pulumi:"template"`
	// Tenant.
	Tenant *string `pulumi:"tenant"`
	// Type of IPv6 address object (default = ipprefix).
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable the visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

type Address6State struct {
	// Minimal TTL of individual IPv6 addresses in FQDN cache.
	CacheTtl pulumi.IntPtrInput
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// IPv6 addresses associated to a specific country.
	Country pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
	EndIp pulumi.StringPtrInput
	// Last MAC address in the range.
	EndMac pulumi.StringPtrInput
	// Endpoint group name.
	EpgName pulumi.StringPtrInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// Fully qualified domain name.
	Fqdn pulumi.StringPtrInput
	// Host Address.
	Host pulumi.StringPtrInput
	// Host type. Valid values: `any`, `specific`.
	HostType pulumi.StringPtrInput
	// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
	Ip6 pulumi.StringPtrInput
	// IP address list. The structure of `list` block is documented below.
	Lists Address6ListArrayInput
	// Multiple MAC address ranges. The structure of `macaddr` block is documented below.
	Macaddrs Address6MacaddrArrayInput
	// Address name.
	Name pulumi.StringPtrInput
	// Object ID for NSX.
	ObjId pulumi.StringPtrInput
	// SDN.
	Sdn pulumi.StringPtrInput
	// SDN Tag.
	SdnTag pulumi.StringPtrInput
	// First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
	StartIp pulumi.StringPtrInput
	// First MAC address in the range.
	StartMac pulumi.StringPtrInput
	// IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
	SubnetSegments Address6SubnetSegmentArrayInput
	// Config object tagging The structure of `tagging` block is documented below.
	Taggings Address6TaggingArrayInput
	// IPv6 address template.
	Template pulumi.StringPtrInput
	// Tenant.
	Tenant pulumi.StringPtrInput
	// Type of IPv6 address object (default = ipprefix).
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable the visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (Address6State) ElementType() reflect.Type {
	return reflect.TypeOf((*address6State)(nil)).Elem()
}

type address6Args struct {
	// Minimal TTL of individual IPv6 addresses in FQDN cache.
	CacheTtl *int `pulumi:"cacheTtl"`
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// IPv6 addresses associated to a specific country.
	Country *string `pulumi:"country"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
	EndIp *string `pulumi:"endIp"`
	// Last MAC address in the range.
	EndMac *string `pulumi:"endMac"`
	// Endpoint group name.
	EpgName *string `pulumi:"epgName"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// Fully qualified domain name.
	Fqdn *string `pulumi:"fqdn"`
	// Host Address.
	Host *string `pulumi:"host"`
	// Host type. Valid values: `any`, `specific`.
	HostType *string `pulumi:"hostType"`
	// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
	Ip6 *string `pulumi:"ip6"`
	// IP address list. The structure of `list` block is documented below.
	Lists []Address6List `pulumi:"lists"`
	// Multiple MAC address ranges. The structure of `macaddr` block is documented below.
	Macaddrs []Address6Macaddr `pulumi:"macaddrs"`
	// Address name.
	Name *string `pulumi:"name"`
	// Object ID for NSX.
	ObjId *string `pulumi:"objId"`
	// SDN.
	Sdn *string `pulumi:"sdn"`
	// SDN Tag.
	SdnTag *string `pulumi:"sdnTag"`
	// First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
	StartIp *string `pulumi:"startIp"`
	// First MAC address in the range.
	StartMac *string `pulumi:"startMac"`
	// IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
	SubnetSegments []Address6SubnetSegment `pulumi:"subnetSegments"`
	// Config object tagging The structure of `tagging` block is documented below.
	Taggings []Address6Tagging `pulumi:"taggings"`
	// IPv6 address template.
	Template *string `pulumi:"template"`
	// Tenant.
	Tenant *string `pulumi:"tenant"`
	// Type of IPv6 address object (default = ipprefix).
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable the visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

// The set of arguments for constructing a Address6 resource.
type Address6Args struct {
	// Minimal TTL of individual IPv6 addresses in FQDN cache.
	CacheTtl pulumi.IntPtrInput
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// IPv6 addresses associated to a specific country.
	Country pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
	EndIp pulumi.StringPtrInput
	// Last MAC address in the range.
	EndMac pulumi.StringPtrInput
	// Endpoint group name.
	EpgName pulumi.StringPtrInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// Fully qualified domain name.
	Fqdn pulumi.StringPtrInput
	// Host Address.
	Host pulumi.StringPtrInput
	// Host type. Valid values: `any`, `specific`.
	HostType pulumi.StringPtrInput
	// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
	Ip6 pulumi.StringPtrInput
	// IP address list. The structure of `list` block is documented below.
	Lists Address6ListArrayInput
	// Multiple MAC address ranges. The structure of `macaddr` block is documented below.
	Macaddrs Address6MacaddrArrayInput
	// Address name.
	Name pulumi.StringPtrInput
	// Object ID for NSX.
	ObjId pulumi.StringPtrInput
	// SDN.
	Sdn pulumi.StringPtrInput
	// SDN Tag.
	SdnTag pulumi.StringPtrInput
	// First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
	StartIp pulumi.StringPtrInput
	// First MAC address in the range.
	StartMac pulumi.StringPtrInput
	// IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
	SubnetSegments Address6SubnetSegmentArrayInput
	// Config object tagging The structure of `tagging` block is documented below.
	Taggings Address6TaggingArrayInput
	// IPv6 address template.
	Template pulumi.StringPtrInput
	// Tenant.
	Tenant pulumi.StringPtrInput
	// Type of IPv6 address object (default = ipprefix).
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable the visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (Address6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*address6Args)(nil)).Elem()
}

type Address6Input interface {
	pulumi.Input

	ToAddress6Output() Address6Output
	ToAddress6OutputWithContext(ctx context.Context) Address6Output
}

func (*Address6) ElementType() reflect.Type {
	return reflect.TypeOf((**Address6)(nil)).Elem()
}

func (i *Address6) ToAddress6Output() Address6Output {
	return i.ToAddress6OutputWithContext(context.Background())
}

func (i *Address6) ToAddress6OutputWithContext(ctx context.Context) Address6Output {
	return pulumi.ToOutputWithContext(ctx, i).(Address6Output)
}

// Address6ArrayInput is an input type that accepts Address6Array and Address6ArrayOutput values.
// You can construct a concrete instance of `Address6ArrayInput` via:
//
//	Address6Array{ Address6Args{...} }
type Address6ArrayInput interface {
	pulumi.Input

	ToAddress6ArrayOutput() Address6ArrayOutput
	ToAddress6ArrayOutputWithContext(context.Context) Address6ArrayOutput
}

type Address6Array []Address6Input

func (Address6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Address6)(nil)).Elem()
}

func (i Address6Array) ToAddress6ArrayOutput() Address6ArrayOutput {
	return i.ToAddress6ArrayOutputWithContext(context.Background())
}

func (i Address6Array) ToAddress6ArrayOutputWithContext(ctx context.Context) Address6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Address6ArrayOutput)
}

// Address6MapInput is an input type that accepts Address6Map and Address6MapOutput values.
// You can construct a concrete instance of `Address6MapInput` via:
//
//	Address6Map{ "key": Address6Args{...} }
type Address6MapInput interface {
	pulumi.Input

	ToAddress6MapOutput() Address6MapOutput
	ToAddress6MapOutputWithContext(context.Context) Address6MapOutput
}

type Address6Map map[string]Address6Input

func (Address6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Address6)(nil)).Elem()
}

func (i Address6Map) ToAddress6MapOutput() Address6MapOutput {
	return i.ToAddress6MapOutputWithContext(context.Background())
}

func (i Address6Map) ToAddress6MapOutputWithContext(ctx context.Context) Address6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Address6MapOutput)
}

type Address6Output struct{ *pulumi.OutputState }

func (Address6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**Address6)(nil)).Elem()
}

func (o Address6Output) ToAddress6Output() Address6Output {
	return o
}

func (o Address6Output) ToAddress6OutputWithContext(ctx context.Context) Address6Output {
	return o
}

// Minimal TTL of individual IPv6 addresses in FQDN cache.
func (o Address6Output) CacheTtl() pulumi.IntOutput {
	return o.ApplyT(func(v *Address6) pulumi.IntOutput { return v.CacheTtl }).(pulumi.IntOutput)
}

// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
func (o Address6Output) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *Address6) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o Address6Output) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// IPv6 addresses associated to a specific country.
func (o Address6Output) Country() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringOutput { return v.Country }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o Address6Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
func (o Address6Output) EndIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringOutput { return v.EndIp }).(pulumi.StringOutput)
}

// Last MAC address in the range.
func (o Address6Output) EndMac() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringOutput { return v.EndMac }).(pulumi.StringOutput)
}

// Endpoint group name.
func (o Address6Output) EpgName() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringOutput { return v.EpgName }).(pulumi.StringOutput)
}

// Security Fabric global object setting. Valid values: `enable`, `disable`.
func (o Address6Output) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringOutput { return v.FabricObject }).(pulumi.StringOutput)
}

// Fully qualified domain name.
func (o Address6Output) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

// Host Address.
func (o Address6Output) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// Host type. Valid values: `any`, `specific`.
func (o Address6Output) HostType() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringOutput { return v.HostType }).(pulumi.StringOutput)
}

// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
func (o Address6Output) Ip6() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringOutput { return v.Ip6 }).(pulumi.StringOutput)
}

// IP address list. The structure of `list` block is documented below.
func (o Address6Output) Lists() Address6ListArrayOutput {
	return o.ApplyT(func(v *Address6) Address6ListArrayOutput { return v.Lists }).(Address6ListArrayOutput)
}

// Multiple MAC address ranges. The structure of `macaddr` block is documented below.
func (o Address6Output) Macaddrs() Address6MacaddrArrayOutput {
	return o.ApplyT(func(v *Address6) Address6MacaddrArrayOutput { return v.Macaddrs }).(Address6MacaddrArrayOutput)
}

// Address name.
func (o Address6Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Object ID for NSX.
func (o Address6Output) ObjId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringPtrOutput { return v.ObjId }).(pulumi.StringPtrOutput)
}

// SDN.
func (o Address6Output) Sdn() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringOutput { return v.Sdn }).(pulumi.StringOutput)
}

// SDN Tag.
func (o Address6Output) SdnTag() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringOutput { return v.SdnTag }).(pulumi.StringOutput)
}

// First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
func (o Address6Output) StartIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringOutput { return v.StartIp }).(pulumi.StringOutput)
}

// First MAC address in the range.
func (o Address6Output) StartMac() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringOutput { return v.StartMac }).(pulumi.StringOutput)
}

// IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
func (o Address6Output) SubnetSegments() Address6SubnetSegmentArrayOutput {
	return o.ApplyT(func(v *Address6) Address6SubnetSegmentArrayOutput { return v.SubnetSegments }).(Address6SubnetSegmentArrayOutput)
}

// Config object tagging The structure of `tagging` block is documented below.
func (o Address6Output) Taggings() Address6TaggingArrayOutput {
	return o.ApplyT(func(v *Address6) Address6TaggingArrayOutput { return v.Taggings }).(Address6TaggingArrayOutput)
}

// IPv6 address template.
func (o Address6Output) Template() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringOutput { return v.Template }).(pulumi.StringOutput)
}

// Tenant.
func (o Address6Output) Tenant() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringOutput { return v.Tenant }).(pulumi.StringOutput)
}

// Type of IPv6 address object (default = ipprefix).
func (o Address6Output) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o Address6Output) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Address6Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable the visibility of the object in the GUI. Valid values: `enable`, `disable`.
func (o Address6Output) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

type Address6ArrayOutput struct{ *pulumi.OutputState }

func (Address6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Address6)(nil)).Elem()
}

func (o Address6ArrayOutput) ToAddress6ArrayOutput() Address6ArrayOutput {
	return o
}

func (o Address6ArrayOutput) ToAddress6ArrayOutputWithContext(ctx context.Context) Address6ArrayOutput {
	return o
}

func (o Address6ArrayOutput) Index(i pulumi.IntInput) Address6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Address6 {
		return vs[0].([]*Address6)[vs[1].(int)]
	}).(Address6Output)
}

type Address6MapOutput struct{ *pulumi.OutputState }

func (Address6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Address6)(nil)).Elem()
}

func (o Address6MapOutput) ToAddress6MapOutput() Address6MapOutput {
	return o
}

func (o Address6MapOutput) ToAddress6MapOutputWithContext(ctx context.Context) Address6MapOutput {
	return o
}

func (o Address6MapOutput) MapIndex(k pulumi.StringInput) Address6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Address6 {
		return vs[0].(map[string]*Address6)[vs[1].(string)]
	}).(Address6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Address6Input)(nil)).Elem(), &Address6{})
	pulumi.RegisterInputType(reflect.TypeOf((*Address6ArrayInput)(nil)).Elem(), Address6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*Address6MapInput)(nil)).Elem(), Address6Map{})
	pulumi.RegisterOutputType(Address6Output{})
	pulumi.RegisterOutputType(Address6ArrayOutput{})
	pulumi.RegisterOutputType(Address6MapOutput{})
}
