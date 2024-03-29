// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dhcp6

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure DHCPv6 servers.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/system"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := system.NewServer(ctx, "trname", &system.ServerArgs{
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// SystemDhcp6 Server can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/dhcp6/server:Server labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/dhcp6/server:Server labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Server struct {
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
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// DHCP server can assign IP configurations to clients connected to this interface.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Method used to assign client IP. Valid values: `range`, `delegated`.
	IpMode pulumi.StringOutput `pulumi:"ipMode"`
	// DHCP IP range configuration. The structure of `ipRange` block is documented below.
	IpRanges ServerIpRangeArrayOutput `pulumi:"ipRanges"`
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
	PrefixRanges ServerPrefixRangeArrayOutput `pulumi:"prefixRanges"`
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

// NewServer registers a new resource with the given unique name, arguments, and options.
func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOption) (*Server, error) {
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
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Server
	err := ctx.RegisterResource("fortios:system/dhcp6/server:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServer gets an existing Server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("fortios:system/dhcp6/server:Server", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Server resources.
type serverState struct {
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
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// DHCP server can assign IP configurations to clients connected to this interface.
	Interface *string `pulumi:"interface"`
	// Method used to assign client IP. Valid values: `range`, `delegated`.
	IpMode *string `pulumi:"ipMode"`
	// DHCP IP range configuration. The structure of `ipRange` block is documented below.
	IpRanges []ServerIpRange `pulumi:"ipRanges"`
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
	PrefixRanges []ServerPrefixRange `pulumi:"prefixRanges"`
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

type ServerState struct {
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
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// DHCP server can assign IP configurations to clients connected to this interface.
	Interface pulumi.StringPtrInput
	// Method used to assign client IP. Valid values: `range`, `delegated`.
	IpMode pulumi.StringPtrInput
	// DHCP IP range configuration. The structure of `ipRange` block is documented below.
	IpRanges ServerIpRangeArrayInput
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
	PrefixRanges ServerPrefixRangeArrayInput
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

func (ServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverState)(nil)).Elem()
}

type serverArgs struct {
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
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// DHCP server can assign IP configurations to clients connected to this interface.
	Interface string `pulumi:"interface"`
	// Method used to assign client IP. Valid values: `range`, `delegated`.
	IpMode *string `pulumi:"ipMode"`
	// DHCP IP range configuration. The structure of `ipRange` block is documented below.
	IpRanges []ServerIpRange `pulumi:"ipRanges"`
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
	PrefixRanges []ServerPrefixRange `pulumi:"prefixRanges"`
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

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
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
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// DHCP server can assign IP configurations to clients connected to this interface.
	Interface pulumi.StringInput
	// Method used to assign client IP. Valid values: `range`, `delegated`.
	IpMode pulumi.StringPtrInput
	// DHCP IP range configuration. The structure of `ipRange` block is documented below.
	IpRanges ServerIpRangeArrayInput
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
	PrefixRanges ServerPrefixRangeArrayInput
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

func (ServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverArgs)(nil)).Elem()
}

type ServerInput interface {
	pulumi.Input

	ToServerOutput() ServerOutput
	ToServerOutputWithContext(ctx context.Context) ServerOutput
}

func (*Server) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (i *Server) ToServerOutput() ServerOutput {
	return i.ToServerOutputWithContext(context.Background())
}

func (i *Server) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerOutput)
}

// ServerArrayInput is an input type that accepts ServerArray and ServerArrayOutput values.
// You can construct a concrete instance of `ServerArrayInput` via:
//
//	ServerArray{ ServerArgs{...} }
type ServerArrayInput interface {
	pulumi.Input

	ToServerArrayOutput() ServerArrayOutput
	ToServerArrayOutputWithContext(context.Context) ServerArrayOutput
}

type ServerArray []ServerInput

func (ServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Server)(nil)).Elem()
}

func (i ServerArray) ToServerArrayOutput() ServerArrayOutput {
	return i.ToServerArrayOutputWithContext(context.Background())
}

func (i ServerArray) ToServerArrayOutputWithContext(ctx context.Context) ServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerArrayOutput)
}

// ServerMapInput is an input type that accepts ServerMap and ServerMapOutput values.
// You can construct a concrete instance of `ServerMapInput` via:
//
//	ServerMap{ "key": ServerArgs{...} }
type ServerMapInput interface {
	pulumi.Input

	ToServerMapOutput() ServerMapOutput
	ToServerMapOutputWithContext(context.Context) ServerMapOutput
}

type ServerMap map[string]ServerInput

func (ServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Server)(nil)).Elem()
}

func (i ServerMap) ToServerMapOutput() ServerMapOutput {
	return i.ToServerMapOutputWithContext(context.Background())
}

func (i ServerMap) ToServerMapOutputWithContext(ctx context.Context) ServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerMapOutput)
}

type ServerOutput struct{ *pulumi.OutputState }

func (ServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (o ServerOutput) ToServerOutput() ServerOutput {
	return o
}

func (o ServerOutput) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return o
}

// IAID of obtained delegated-prefix from the upstream interface.
func (o ServerOutput) DelegatedPrefixIaid() pulumi.IntOutput {
	return o.ApplyT(func(v *Server) pulumi.IntOutput { return v.DelegatedPrefixIaid }).(pulumi.IntOutput)
}

// DNS search list options. Valid values: `delegated`, `specify`.
func (o ServerOutput) DnsSearchList() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.DnsSearchList }).(pulumi.StringOutput)
}

// DNS server 1.
func (o ServerOutput) DnsServer1() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.DnsServer1 }).(pulumi.StringOutput)
}

// DNS server 2.
func (o ServerOutput) DnsServer2() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.DnsServer2 }).(pulumi.StringOutput)
}

// DNS server 3.
func (o ServerOutput) DnsServer3() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.DnsServer3 }).(pulumi.StringOutput)
}

// DNS server 4.
func (o ServerOutput) DnsServer4() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.DnsServer4 }).(pulumi.StringOutput)
}

// Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
func (o ServerOutput) DnsService() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.DnsService }).(pulumi.StringOutput)
}

// Domain name suffix for the IP addresses that the DHCP server assigns to clients.
func (o ServerOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o ServerOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// ID.
func (o ServerOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Server) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o ServerOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// DHCP server can assign IP configurations to clients connected to this interface.
func (o ServerOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Method used to assign client IP. Valid values: `range`, `delegated`.
func (o ServerOutput) IpMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.IpMode }).(pulumi.StringOutput)
}

// DHCP IP range configuration. The structure of `ipRange` block is documented below.
func (o ServerOutput) IpRanges() ServerIpRangeArrayOutput {
	return o.ApplyT(func(v *Server) ServerIpRangeArrayOutput { return v.IpRanges }).(ServerIpRangeArrayOutput)
}

// Lease time in seconds, 0 means unlimited.
func (o ServerOutput) LeaseTime() pulumi.IntOutput {
	return o.ApplyT(func(v *Server) pulumi.IntOutput { return v.LeaseTime }).(pulumi.IntOutput)
}

// Option 1.
func (o ServerOutput) Option1() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Option1 }).(pulumi.StringOutput)
}

// Option 2.
func (o ServerOutput) Option2() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Option2 }).(pulumi.StringOutput)
}

// Option 3.
func (o ServerOutput) Option3() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Option3 }).(pulumi.StringOutput)
}

// Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
func (o ServerOutput) PrefixMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.PrefixMode }).(pulumi.StringOutput)
}

// DHCP prefix configuration. The structure of `prefixRange` block is documented below.
func (o ServerOutput) PrefixRanges() ServerPrefixRangeArrayOutput {
	return o.ApplyT(func(v *Server) ServerPrefixRangeArrayOutput { return v.PrefixRanges }).(ServerPrefixRangeArrayOutput)
}

// Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
func (o ServerOutput) RapidCommit() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.RapidCommit }).(pulumi.StringOutput)
}

// Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
func (o ServerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Subnet or subnet-id if the IP mode is delegated.
func (o ServerOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

// Interface name from where delegated information is provided.
func (o ServerOutput) UpstreamInterface() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.UpstreamInterface }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ServerOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ServerArrayOutput struct{ *pulumi.OutputState }

func (ServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Server)(nil)).Elem()
}

func (o ServerArrayOutput) ToServerArrayOutput() ServerArrayOutput {
	return o
}

func (o ServerArrayOutput) ToServerArrayOutputWithContext(ctx context.Context) ServerArrayOutput {
	return o
}

func (o ServerArrayOutput) Index(i pulumi.IntInput) ServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Server {
		return vs[0].([]*Server)[vs[1].(int)]
	}).(ServerOutput)
}

type ServerMapOutput struct{ *pulumi.OutputState }

func (ServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Server)(nil)).Elem()
}

func (o ServerMapOutput) ToServerMapOutput() ServerMapOutput {
	return o
}

func (o ServerMapOutput) ToServerMapOutputWithContext(ctx context.Context) ServerMapOutput {
	return o
}

func (o ServerMapOutput) MapIndex(k pulumi.StringInput) ServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Server {
		return vs[0].(map[string]*Server)[vs[1].(string)]
	}).(ServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerInput)(nil)).Elem(), &Server{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerArrayInput)(nil)).Elem(), ServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerMapInput)(nil)).Elem(), ServerMap{})
	pulumi.RegisterOutputType(ServerOutput{})
	pulumi.RegisterOutputType(ServerArrayOutput{})
	pulumi.RegisterOutputType(ServerMapOutput{})
}
