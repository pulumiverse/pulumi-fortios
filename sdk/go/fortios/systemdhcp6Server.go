// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure DHCPv6 servers.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fortios.NewSystemdhcp6Server(ctx, "trname", &fortios.Systemdhcp6ServerArgs{
//				Fosid:       pulumi.Int(1),
//				Interface:   pulumi.String("port3"),
//				LeaseTime:   pulumi.Int(604800),
//				RapidCommit: pulumi.String("disable"),
//				Status:      pulumi.String("enable"),
//				Subnet:      pulumi.String("2001:db8:1234:113::/64"),
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
// # SystemDhcp6 Server can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemdhcp6Server:Systemdhcp6Server labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemdhcp6Server:Systemdhcp6Server labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Systemdhcp6Server struct {
	pulumi.CustomResourceState

	// IAID of obtained delegated-prefix from the upstream interface.
	DelegatedPrefixIaid pulumi.IntOutput `pulumi:"delegatedPrefixIaid"`
	// DNS search list options. Valid values: `delegated`, `specify`.
	DnsSearchList pulumi.StringOutput `pulumi:"dnsSearchList"`
	// DNS server 1.
	DnsServer1 pulumi.StringOutput `pulumi:"dnsServer1"`
	// DNS server 2.
	DnsServer2 pulumi.StringOutput `pulumi:"dnsServer2"`
	// DNS server 3.
	DnsServer3 pulumi.StringOutput `pulumi:"dnsServer3"`
	// DNS server 4.
	DnsServer4 pulumi.StringOutput `pulumi:"dnsServer4"`
	// Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
	DnsService pulumi.StringOutput `pulumi:"dnsService"`
	// Domain name suffix for the IP addresses that the DHCP server assigns to clients.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// DHCP server can assign IP configurations to clients connected to this interface.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Method used to assign client IP. Valid values: `range`, `delegated`.
	IpMode pulumi.StringOutput `pulumi:"ipMode"`
	// DHCP IP range configuration. The structure of `ipRange` block is documented below.
	IpRanges Systemdhcp6ServerIpRangeArrayOutput `pulumi:"ipRanges"`
	// Lease time in seconds, 0 means unlimited.
	LeaseTime pulumi.IntOutput `pulumi:"leaseTime"`
	// Option 1.
	Option1 pulumi.StringOutput `pulumi:"option1"`
	// Option 2.
	Option2 pulumi.StringOutput `pulumi:"option2"`
	// Option 3.
	Option3 pulumi.StringOutput `pulumi:"option3"`
	// Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
	PrefixMode pulumi.StringOutput `pulumi:"prefixMode"`
	// DHCP prefix configuration. The structure of `prefixRange` block is documented below.
	PrefixRanges Systemdhcp6ServerPrefixRangeArrayOutput `pulumi:"prefixRanges"`
	// Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
	RapidCommit pulumi.StringOutput `pulumi:"rapidCommit"`
	// Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Subnet or subnet-id if the IP mode is delegated.
	Subnet pulumi.StringOutput `pulumi:"subnet"`
	// Interface name from where delegated information is provided.
	UpstreamInterface pulumi.StringOutput `pulumi:"upstreamInterface"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemdhcp6Server registers a new resource with the given unique name, arguments, and options.
func NewSystemdhcp6Server(ctx *pulumi.Context,
	name string, args *Systemdhcp6ServerArgs, opts ...pulumi.ResourceOption) (*Systemdhcp6Server, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.Subnet == nil {
		return nil, errors.New("invalid value for required argument 'Subnet'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Systemdhcp6Server
	err := ctx.RegisterResource("fortios:index/systemdhcp6Server:Systemdhcp6Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemdhcp6Server gets an existing Systemdhcp6Server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemdhcp6Server(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Systemdhcp6ServerState, opts ...pulumi.ResourceOption) (*Systemdhcp6Server, error) {
	var resource Systemdhcp6Server
	err := ctx.ReadResource("fortios:index/systemdhcp6Server:Systemdhcp6Server", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Systemdhcp6Server resources.
type systemdhcp6ServerState struct {
	// IAID of obtained delegated-prefix from the upstream interface.
	DelegatedPrefixIaid *int `pulumi:"delegatedPrefixIaid"`
	// DNS search list options. Valid values: `delegated`, `specify`.
	DnsSearchList *string `pulumi:"dnsSearchList"`
	// DNS server 1.
	DnsServer1 *string `pulumi:"dnsServer1"`
	// DNS server 2.
	DnsServer2 *string `pulumi:"dnsServer2"`
	// DNS server 3.
	DnsServer3 *string `pulumi:"dnsServer3"`
	// DNS server 4.
	DnsServer4 *string `pulumi:"dnsServer4"`
	// Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
	DnsService *string `pulumi:"dnsService"`
	// Domain name suffix for the IP addresses that the DHCP server assigns to clients.
	Domain *string `pulumi:"domain"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// DHCP server can assign IP configurations to clients connected to this interface.
	Interface *string `pulumi:"interface"`
	// Method used to assign client IP. Valid values: `range`, `delegated`.
	IpMode *string `pulumi:"ipMode"`
	// DHCP IP range configuration. The structure of `ipRange` block is documented below.
	IpRanges []Systemdhcp6ServerIpRange `pulumi:"ipRanges"`
	// Lease time in seconds, 0 means unlimited.
	LeaseTime *int `pulumi:"leaseTime"`
	// Option 1.
	Option1 *string `pulumi:"option1"`
	// Option 2.
	Option2 *string `pulumi:"option2"`
	// Option 3.
	Option3 *string `pulumi:"option3"`
	// Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
	PrefixMode *string `pulumi:"prefixMode"`
	// DHCP prefix configuration. The structure of `prefixRange` block is documented below.
	PrefixRanges []Systemdhcp6ServerPrefixRange `pulumi:"prefixRanges"`
	// Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
	RapidCommit *string `pulumi:"rapidCommit"`
	// Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Subnet or subnet-id if the IP mode is delegated.
	Subnet *string `pulumi:"subnet"`
	// Interface name from where delegated information is provided.
	UpstreamInterface *string `pulumi:"upstreamInterface"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Systemdhcp6ServerState struct {
	// IAID of obtained delegated-prefix from the upstream interface.
	DelegatedPrefixIaid pulumi.IntPtrInput
	// DNS search list options. Valid values: `delegated`, `specify`.
	DnsSearchList pulumi.StringPtrInput
	// DNS server 1.
	DnsServer1 pulumi.StringPtrInput
	// DNS server 2.
	DnsServer2 pulumi.StringPtrInput
	// DNS server 3.
	DnsServer3 pulumi.StringPtrInput
	// DNS server 4.
	DnsServer4 pulumi.StringPtrInput
	// Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
	DnsService pulumi.StringPtrInput
	// Domain name suffix for the IP addresses that the DHCP server assigns to clients.
	Domain pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// ID.
	Fosid pulumi.IntPtrInput
	// DHCP server can assign IP configurations to clients connected to this interface.
	Interface pulumi.StringPtrInput
	// Method used to assign client IP. Valid values: `range`, `delegated`.
	IpMode pulumi.StringPtrInput
	// DHCP IP range configuration. The structure of `ipRange` block is documented below.
	IpRanges Systemdhcp6ServerIpRangeArrayInput
	// Lease time in seconds, 0 means unlimited.
	LeaseTime pulumi.IntPtrInput
	// Option 1.
	Option1 pulumi.StringPtrInput
	// Option 2.
	Option2 pulumi.StringPtrInput
	// Option 3.
	Option3 pulumi.StringPtrInput
	// Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
	PrefixMode pulumi.StringPtrInput
	// DHCP prefix configuration. The structure of `prefixRange` block is documented below.
	PrefixRanges Systemdhcp6ServerPrefixRangeArrayInput
	// Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
	RapidCommit pulumi.StringPtrInput
	// Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Subnet or subnet-id if the IP mode is delegated.
	Subnet pulumi.StringPtrInput
	// Interface name from where delegated information is provided.
	UpstreamInterface pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Systemdhcp6ServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemdhcp6ServerState)(nil)).Elem()
}

type systemdhcp6ServerArgs struct {
	// IAID of obtained delegated-prefix from the upstream interface.
	DelegatedPrefixIaid *int `pulumi:"delegatedPrefixIaid"`
	// DNS search list options. Valid values: `delegated`, `specify`.
	DnsSearchList *string `pulumi:"dnsSearchList"`
	// DNS server 1.
	DnsServer1 *string `pulumi:"dnsServer1"`
	// DNS server 2.
	DnsServer2 *string `pulumi:"dnsServer2"`
	// DNS server 3.
	DnsServer3 *string `pulumi:"dnsServer3"`
	// DNS server 4.
	DnsServer4 *string `pulumi:"dnsServer4"`
	// Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
	DnsService *string `pulumi:"dnsService"`
	// Domain name suffix for the IP addresses that the DHCP server assigns to clients.
	Domain *string `pulumi:"domain"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// ID.
	Fosid int `pulumi:"fosid"`
	// DHCP server can assign IP configurations to clients connected to this interface.
	Interface string `pulumi:"interface"`
	// Method used to assign client IP. Valid values: `range`, `delegated`.
	IpMode *string `pulumi:"ipMode"`
	// DHCP IP range configuration. The structure of `ipRange` block is documented below.
	IpRanges []Systemdhcp6ServerIpRange `pulumi:"ipRanges"`
	// Lease time in seconds, 0 means unlimited.
	LeaseTime *int `pulumi:"leaseTime"`
	// Option 1.
	Option1 *string `pulumi:"option1"`
	// Option 2.
	Option2 *string `pulumi:"option2"`
	// Option 3.
	Option3 *string `pulumi:"option3"`
	// Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
	PrefixMode *string `pulumi:"prefixMode"`
	// DHCP prefix configuration. The structure of `prefixRange` block is documented below.
	PrefixRanges []Systemdhcp6ServerPrefixRange `pulumi:"prefixRanges"`
	// Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
	RapidCommit *string `pulumi:"rapidCommit"`
	// Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Subnet or subnet-id if the IP mode is delegated.
	Subnet string `pulumi:"subnet"`
	// Interface name from where delegated information is provided.
	UpstreamInterface *string `pulumi:"upstreamInterface"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Systemdhcp6Server resource.
type Systemdhcp6ServerArgs struct {
	// IAID of obtained delegated-prefix from the upstream interface.
	DelegatedPrefixIaid pulumi.IntPtrInput
	// DNS search list options. Valid values: `delegated`, `specify`.
	DnsSearchList pulumi.StringPtrInput
	// DNS server 1.
	DnsServer1 pulumi.StringPtrInput
	// DNS server 2.
	DnsServer2 pulumi.StringPtrInput
	// DNS server 3.
	DnsServer3 pulumi.StringPtrInput
	// DNS server 4.
	DnsServer4 pulumi.StringPtrInput
	// Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
	DnsService pulumi.StringPtrInput
	// Domain name suffix for the IP addresses that the DHCP server assigns to clients.
	Domain pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// ID.
	Fosid pulumi.IntInput
	// DHCP server can assign IP configurations to clients connected to this interface.
	Interface pulumi.StringInput
	// Method used to assign client IP. Valid values: `range`, `delegated`.
	IpMode pulumi.StringPtrInput
	// DHCP IP range configuration. The structure of `ipRange` block is documented below.
	IpRanges Systemdhcp6ServerIpRangeArrayInput
	// Lease time in seconds, 0 means unlimited.
	LeaseTime pulumi.IntPtrInput
	// Option 1.
	Option1 pulumi.StringPtrInput
	// Option 2.
	Option2 pulumi.StringPtrInput
	// Option 3.
	Option3 pulumi.StringPtrInput
	// Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
	PrefixMode pulumi.StringPtrInput
	// DHCP prefix configuration. The structure of `prefixRange` block is documented below.
	PrefixRanges Systemdhcp6ServerPrefixRangeArrayInput
	// Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
	RapidCommit pulumi.StringPtrInput
	// Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Subnet or subnet-id if the IP mode is delegated.
	Subnet pulumi.StringInput
	// Interface name from where delegated information is provided.
	UpstreamInterface pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Systemdhcp6ServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemdhcp6ServerArgs)(nil)).Elem()
}

type Systemdhcp6ServerInput interface {
	pulumi.Input

	ToSystemdhcp6ServerOutput() Systemdhcp6ServerOutput
	ToSystemdhcp6ServerOutputWithContext(ctx context.Context) Systemdhcp6ServerOutput
}

func (*Systemdhcp6Server) ElementType() reflect.Type {
	return reflect.TypeOf((**Systemdhcp6Server)(nil)).Elem()
}

func (i *Systemdhcp6Server) ToSystemdhcp6ServerOutput() Systemdhcp6ServerOutput {
	return i.ToSystemdhcp6ServerOutputWithContext(context.Background())
}

func (i *Systemdhcp6Server) ToSystemdhcp6ServerOutputWithContext(ctx context.Context) Systemdhcp6ServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Systemdhcp6ServerOutput)
}

// Systemdhcp6ServerArrayInput is an input type that accepts Systemdhcp6ServerArray and Systemdhcp6ServerArrayOutput values.
// You can construct a concrete instance of `Systemdhcp6ServerArrayInput` via:
//
//	Systemdhcp6ServerArray{ Systemdhcp6ServerArgs{...} }
type Systemdhcp6ServerArrayInput interface {
	pulumi.Input

	ToSystemdhcp6ServerArrayOutput() Systemdhcp6ServerArrayOutput
	ToSystemdhcp6ServerArrayOutputWithContext(context.Context) Systemdhcp6ServerArrayOutput
}

type Systemdhcp6ServerArray []Systemdhcp6ServerInput

func (Systemdhcp6ServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Systemdhcp6Server)(nil)).Elem()
}

func (i Systemdhcp6ServerArray) ToSystemdhcp6ServerArrayOutput() Systemdhcp6ServerArrayOutput {
	return i.ToSystemdhcp6ServerArrayOutputWithContext(context.Background())
}

func (i Systemdhcp6ServerArray) ToSystemdhcp6ServerArrayOutputWithContext(ctx context.Context) Systemdhcp6ServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Systemdhcp6ServerArrayOutput)
}

// Systemdhcp6ServerMapInput is an input type that accepts Systemdhcp6ServerMap and Systemdhcp6ServerMapOutput values.
// You can construct a concrete instance of `Systemdhcp6ServerMapInput` via:
//
//	Systemdhcp6ServerMap{ "key": Systemdhcp6ServerArgs{...} }
type Systemdhcp6ServerMapInput interface {
	pulumi.Input

	ToSystemdhcp6ServerMapOutput() Systemdhcp6ServerMapOutput
	ToSystemdhcp6ServerMapOutputWithContext(context.Context) Systemdhcp6ServerMapOutput
}

type Systemdhcp6ServerMap map[string]Systemdhcp6ServerInput

func (Systemdhcp6ServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Systemdhcp6Server)(nil)).Elem()
}

func (i Systemdhcp6ServerMap) ToSystemdhcp6ServerMapOutput() Systemdhcp6ServerMapOutput {
	return i.ToSystemdhcp6ServerMapOutputWithContext(context.Background())
}

func (i Systemdhcp6ServerMap) ToSystemdhcp6ServerMapOutputWithContext(ctx context.Context) Systemdhcp6ServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Systemdhcp6ServerMapOutput)
}

type Systemdhcp6ServerOutput struct{ *pulumi.OutputState }

func (Systemdhcp6ServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Systemdhcp6Server)(nil)).Elem()
}

func (o Systemdhcp6ServerOutput) ToSystemdhcp6ServerOutput() Systemdhcp6ServerOutput {
	return o
}

func (o Systemdhcp6ServerOutput) ToSystemdhcp6ServerOutputWithContext(ctx context.Context) Systemdhcp6ServerOutput {
	return o
}

// IAID of obtained delegated-prefix from the upstream interface.
func (o Systemdhcp6ServerOutput) DelegatedPrefixIaid() pulumi.IntOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.IntOutput { return v.DelegatedPrefixIaid }).(pulumi.IntOutput)
}

// DNS search list options. Valid values: `delegated`, `specify`.
func (o Systemdhcp6ServerOutput) DnsSearchList() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.StringOutput { return v.DnsSearchList }).(pulumi.StringOutput)
}

// DNS server 1.
func (o Systemdhcp6ServerOutput) DnsServer1() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.StringOutput { return v.DnsServer1 }).(pulumi.StringOutput)
}

// DNS server 2.
func (o Systemdhcp6ServerOutput) DnsServer2() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.StringOutput { return v.DnsServer2 }).(pulumi.StringOutput)
}

// DNS server 3.
func (o Systemdhcp6ServerOutput) DnsServer3() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.StringOutput { return v.DnsServer3 }).(pulumi.StringOutput)
}

// DNS server 4.
func (o Systemdhcp6ServerOutput) DnsServer4() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.StringOutput { return v.DnsServer4 }).(pulumi.StringOutput)
}

// Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
func (o Systemdhcp6ServerOutput) DnsService() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.StringOutput { return v.DnsService }).(pulumi.StringOutput)
}

// Domain name suffix for the IP addresses that the DHCP server assigns to clients.
func (o Systemdhcp6ServerOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o Systemdhcp6ServerOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// ID.
func (o Systemdhcp6ServerOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// DHCP server can assign IP configurations to clients connected to this interface.
func (o Systemdhcp6ServerOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Method used to assign client IP. Valid values: `range`, `delegated`.
func (o Systemdhcp6ServerOutput) IpMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.StringOutput { return v.IpMode }).(pulumi.StringOutput)
}

// DHCP IP range configuration. The structure of `ipRange` block is documented below.
func (o Systemdhcp6ServerOutput) IpRanges() Systemdhcp6ServerIpRangeArrayOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) Systemdhcp6ServerIpRangeArrayOutput { return v.IpRanges }).(Systemdhcp6ServerIpRangeArrayOutput)
}

// Lease time in seconds, 0 means unlimited.
func (o Systemdhcp6ServerOutput) LeaseTime() pulumi.IntOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.IntOutput { return v.LeaseTime }).(pulumi.IntOutput)
}

// Option 1.
func (o Systemdhcp6ServerOutput) Option1() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.StringOutput { return v.Option1 }).(pulumi.StringOutput)
}

// Option 2.
func (o Systemdhcp6ServerOutput) Option2() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.StringOutput { return v.Option2 }).(pulumi.StringOutput)
}

// Option 3.
func (o Systemdhcp6ServerOutput) Option3() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.StringOutput { return v.Option3 }).(pulumi.StringOutput)
}

// Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
func (o Systemdhcp6ServerOutput) PrefixMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.StringOutput { return v.PrefixMode }).(pulumi.StringOutput)
}

// DHCP prefix configuration. The structure of `prefixRange` block is documented below.
func (o Systemdhcp6ServerOutput) PrefixRanges() Systemdhcp6ServerPrefixRangeArrayOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) Systemdhcp6ServerPrefixRangeArrayOutput { return v.PrefixRanges }).(Systemdhcp6ServerPrefixRangeArrayOutput)
}

// Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
func (o Systemdhcp6ServerOutput) RapidCommit() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.StringOutput { return v.RapidCommit }).(pulumi.StringOutput)
}

// Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
func (o Systemdhcp6ServerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Subnet or subnet-id if the IP mode is delegated.
func (o Systemdhcp6ServerOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

// Interface name from where delegated information is provided.
func (o Systemdhcp6ServerOutput) UpstreamInterface() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.StringOutput { return v.UpstreamInterface }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Systemdhcp6ServerOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Systemdhcp6Server) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type Systemdhcp6ServerArrayOutput struct{ *pulumi.OutputState }

func (Systemdhcp6ServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Systemdhcp6Server)(nil)).Elem()
}

func (o Systemdhcp6ServerArrayOutput) ToSystemdhcp6ServerArrayOutput() Systemdhcp6ServerArrayOutput {
	return o
}

func (o Systemdhcp6ServerArrayOutput) ToSystemdhcp6ServerArrayOutputWithContext(ctx context.Context) Systemdhcp6ServerArrayOutput {
	return o
}

func (o Systemdhcp6ServerArrayOutput) Index(i pulumi.IntInput) Systemdhcp6ServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Systemdhcp6Server {
		return vs[0].([]*Systemdhcp6Server)[vs[1].(int)]
	}).(Systemdhcp6ServerOutput)
}

type Systemdhcp6ServerMapOutput struct{ *pulumi.OutputState }

func (Systemdhcp6ServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Systemdhcp6Server)(nil)).Elem()
}

func (o Systemdhcp6ServerMapOutput) ToSystemdhcp6ServerMapOutput() Systemdhcp6ServerMapOutput {
	return o
}

func (o Systemdhcp6ServerMapOutput) ToSystemdhcp6ServerMapOutputWithContext(ctx context.Context) Systemdhcp6ServerMapOutput {
	return o
}

func (o Systemdhcp6ServerMapOutput) MapIndex(k pulumi.StringInput) Systemdhcp6ServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Systemdhcp6Server {
		return vs[0].(map[string]*Systemdhcp6Server)[vs[1].(string)]
	}).(Systemdhcp6ServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Systemdhcp6ServerInput)(nil)).Elem(), &Systemdhcp6Server{})
	pulumi.RegisterInputType(reflect.TypeOf((*Systemdhcp6ServerArrayInput)(nil)).Elem(), Systemdhcp6ServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Systemdhcp6ServerMapInput)(nil)).Elem(), Systemdhcp6ServerMap{})
	pulumi.RegisterOutputType(Systemdhcp6ServerOutput{})
	pulumi.RegisterOutputType(Systemdhcp6ServerArrayOutput{})
	pulumi.RegisterOutputType(Systemdhcp6ServerMapOutput{})
}
