// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure DNS databases.
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
//			_, err := system.NewDnsdatabase(ctx, "trname", &system.DnsdatabaseArgs{
//				Authoritative: pulumi.String("enable"),
//				Contact:       pulumi.String("hostmaster"),
//				DnsEntries: system.DnsdatabaseDnsEntryArray{
//					&system.DnsdatabaseDnsEntryArgs{
//						Hostname: pulumi.String("sghsgh.com"),
//						Ttl:      pulumi.Int(3),
//						Type:     pulumi.String("MX"),
//					},
//				},
//				Domain:      pulumi.String("s.com"),
//				Forwarder:   pulumi.String("\"9.9.9.9\" \"3.3.3.3\" "),
//				IpMaster:    pulumi.String("0.0.0.0"),
//				PrimaryName: pulumi.String("dns"),
//				SourceIp:    pulumi.String("0.0.0.0"),
//				Status:      pulumi.String("enable"),
//				Ttl:         pulumi.Int(86400),
//				Type:        pulumi.String("master"),
//				View:        pulumi.String("shadow"),
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
// System DnsDatabase can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/dnsdatabase:Dnsdatabase labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/dnsdatabase:Dnsdatabase labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Dnsdatabase struct {
	pulumi.CustomResourceState

	// DNS zone transfer IP address list.
	AllowTransfer pulumi.StringOutput `pulumi:"allowTransfer"`
	// Enable/disable authoritative zone. Valid values: `enable`, `disable`.
	Authoritative pulumi.StringOutput `pulumi:"authoritative"`
	// Email address of the administrator for this zone.
	// You can specify only the username (e.g. admin) or full email address (e.g. admin@test.com)
	// When using a simple username, the domain of the email will be this zone.
	Contact pulumi.StringOutput `pulumi:"contact"`
	// DNS entry. The structure of `dnsEntry` block is documented below.
	DnsEntries DnsdatabaseDnsEntryArrayOutput `pulumi:"dnsEntries"`
	// Domain name.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// DNS zone forwarder IP address list.
	Forwarder pulumi.StringOutput `pulumi:"forwarder"`
	// Forwarder IPv6 address.
	Forwarder6 pulumi.StringOutput `pulumi:"forwarder6"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.
	IpMaster pulumi.StringOutput `pulumi:"ipMaster"`
	// IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.
	IpPrimary pulumi.StringOutput `pulumi:"ipPrimary"`
	// Zone name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Domain name of the default DNS server for this zone.
	PrimaryName pulumi.StringOutput `pulumi:"primaryName"`
	// Maximum number of resource records (10 - 65536, 0 means infinite).
	RrMax pulumi.IntOutput `pulumi:"rrMax"`
	// Source IP for forwarding to DNS server.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// IPv6 source IP address for forwarding to DNS server.
	SourceIp6 pulumi.StringOutput `pulumi:"sourceIp6"`
	// Enable/disable this DNS zone. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Default time-to-live value for the entries of this DNS zone (0 - 2147483647 sec, default = 86400).
	Ttl pulumi.IntOutput `pulumi:"ttl"`
	// Zone type (master to manage entries directly, slave to import entries from other zones).
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Zone view (public to serve public clients, shadow to serve internal clients).
	View pulumi.StringOutput `pulumi:"view"`
}

// NewDnsdatabase registers a new resource with the given unique name, arguments, and options.
func NewDnsdatabase(ctx *pulumi.Context,
	name string, args *DnsdatabaseArgs, opts ...pulumi.ResourceOption) (*Dnsdatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Authoritative == nil {
		return nil, errors.New("invalid value for required argument 'Authoritative'")
	}
	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.Ttl == nil {
		return nil, errors.New("invalid value for required argument 'Ttl'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.View == nil {
		return nil, errors.New("invalid value for required argument 'View'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Dnsdatabase
	err := ctx.RegisterResource("fortios:system/dnsdatabase:Dnsdatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsdatabase gets an existing Dnsdatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsdatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsdatabaseState, opts ...pulumi.ResourceOption) (*Dnsdatabase, error) {
	var resource Dnsdatabase
	err := ctx.ReadResource("fortios:system/dnsdatabase:Dnsdatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dnsdatabase resources.
type dnsdatabaseState struct {
	// DNS zone transfer IP address list.
	AllowTransfer *string `pulumi:"allowTransfer"`
	// Enable/disable authoritative zone. Valid values: `enable`, `disable`.
	Authoritative *string `pulumi:"authoritative"`
	// Email address of the administrator for this zone.
	// You can specify only the username (e.g. admin) or full email address (e.g. admin@test.com)
	// When using a simple username, the domain of the email will be this zone.
	Contact *string `pulumi:"contact"`
	// DNS entry. The structure of `dnsEntry` block is documented below.
	DnsEntries []DnsdatabaseDnsEntry `pulumi:"dnsEntries"`
	// Domain name.
	Domain *string `pulumi:"domain"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// DNS zone forwarder IP address list.
	Forwarder *string `pulumi:"forwarder"`
	// Forwarder IPv6 address.
	Forwarder6 *string `pulumi:"forwarder6"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.
	IpMaster *string `pulumi:"ipMaster"`
	// IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.
	IpPrimary *string `pulumi:"ipPrimary"`
	// Zone name.
	Name *string `pulumi:"name"`
	// Domain name of the default DNS server for this zone.
	PrimaryName *string `pulumi:"primaryName"`
	// Maximum number of resource records (10 - 65536, 0 means infinite).
	RrMax *int `pulumi:"rrMax"`
	// Source IP for forwarding to DNS server.
	SourceIp *string `pulumi:"sourceIp"`
	// IPv6 source IP address for forwarding to DNS server.
	SourceIp6 *string `pulumi:"sourceIp6"`
	// Enable/disable this DNS zone. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Default time-to-live value for the entries of this DNS zone (0 - 2147483647 sec, default = 86400).
	Ttl *int `pulumi:"ttl"`
	// Zone type (master to manage entries directly, slave to import entries from other zones).
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Zone view (public to serve public clients, shadow to serve internal clients).
	View *string `pulumi:"view"`
}

type DnsdatabaseState struct {
	// DNS zone transfer IP address list.
	AllowTransfer pulumi.StringPtrInput
	// Enable/disable authoritative zone. Valid values: `enable`, `disable`.
	Authoritative pulumi.StringPtrInput
	// Email address of the administrator for this zone.
	// You can specify only the username (e.g. admin) or full email address (e.g. admin@test.com)
	// When using a simple username, the domain of the email will be this zone.
	Contact pulumi.StringPtrInput
	// DNS entry. The structure of `dnsEntry` block is documented below.
	DnsEntries DnsdatabaseDnsEntryArrayInput
	// Domain name.
	Domain pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// DNS zone forwarder IP address list.
	Forwarder pulumi.StringPtrInput
	// Forwarder IPv6 address.
	Forwarder6 pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.
	IpMaster pulumi.StringPtrInput
	// IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.
	IpPrimary pulumi.StringPtrInput
	// Zone name.
	Name pulumi.StringPtrInput
	// Domain name of the default DNS server for this zone.
	PrimaryName pulumi.StringPtrInput
	// Maximum number of resource records (10 - 65536, 0 means infinite).
	RrMax pulumi.IntPtrInput
	// Source IP for forwarding to DNS server.
	SourceIp pulumi.StringPtrInput
	// IPv6 source IP address for forwarding to DNS server.
	SourceIp6 pulumi.StringPtrInput
	// Enable/disable this DNS zone. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Default time-to-live value for the entries of this DNS zone (0 - 2147483647 sec, default = 86400).
	Ttl pulumi.IntPtrInput
	// Zone type (master to manage entries directly, slave to import entries from other zones).
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Zone view (public to serve public clients, shadow to serve internal clients).
	View pulumi.StringPtrInput
}

func (DnsdatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsdatabaseState)(nil)).Elem()
}

type dnsdatabaseArgs struct {
	// DNS zone transfer IP address list.
	AllowTransfer *string `pulumi:"allowTransfer"`
	// Enable/disable authoritative zone. Valid values: `enable`, `disable`.
	Authoritative string `pulumi:"authoritative"`
	// Email address of the administrator for this zone.
	// You can specify only the username (e.g. admin) or full email address (e.g. admin@test.com)
	// When using a simple username, the domain of the email will be this zone.
	Contact *string `pulumi:"contact"`
	// DNS entry. The structure of `dnsEntry` block is documented below.
	DnsEntries []DnsdatabaseDnsEntry `pulumi:"dnsEntries"`
	// Domain name.
	Domain string `pulumi:"domain"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// DNS zone forwarder IP address list.
	Forwarder *string `pulumi:"forwarder"`
	// Forwarder IPv6 address.
	Forwarder6 *string `pulumi:"forwarder6"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.
	IpMaster *string `pulumi:"ipMaster"`
	// IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.
	IpPrimary *string `pulumi:"ipPrimary"`
	// Zone name.
	Name *string `pulumi:"name"`
	// Domain name of the default DNS server for this zone.
	PrimaryName *string `pulumi:"primaryName"`
	// Maximum number of resource records (10 - 65536, 0 means infinite).
	RrMax *int `pulumi:"rrMax"`
	// Source IP for forwarding to DNS server.
	SourceIp *string `pulumi:"sourceIp"`
	// IPv6 source IP address for forwarding to DNS server.
	SourceIp6 *string `pulumi:"sourceIp6"`
	// Enable/disable this DNS zone. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Default time-to-live value for the entries of this DNS zone (0 - 2147483647 sec, default = 86400).
	Ttl int `pulumi:"ttl"`
	// Zone type (master to manage entries directly, slave to import entries from other zones).
	Type string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Zone view (public to serve public clients, shadow to serve internal clients).
	View string `pulumi:"view"`
}

// The set of arguments for constructing a Dnsdatabase resource.
type DnsdatabaseArgs struct {
	// DNS zone transfer IP address list.
	AllowTransfer pulumi.StringPtrInput
	// Enable/disable authoritative zone. Valid values: `enable`, `disable`.
	Authoritative pulumi.StringInput
	// Email address of the administrator for this zone.
	// You can specify only the username (e.g. admin) or full email address (e.g. admin@test.com)
	// When using a simple username, the domain of the email will be this zone.
	Contact pulumi.StringPtrInput
	// DNS entry. The structure of `dnsEntry` block is documented below.
	DnsEntries DnsdatabaseDnsEntryArrayInput
	// Domain name.
	Domain pulumi.StringInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// DNS zone forwarder IP address list.
	Forwarder pulumi.StringPtrInput
	// Forwarder IPv6 address.
	Forwarder6 pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.
	IpMaster pulumi.StringPtrInput
	// IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.
	IpPrimary pulumi.StringPtrInput
	// Zone name.
	Name pulumi.StringPtrInput
	// Domain name of the default DNS server for this zone.
	PrimaryName pulumi.StringPtrInput
	// Maximum number of resource records (10 - 65536, 0 means infinite).
	RrMax pulumi.IntPtrInput
	// Source IP for forwarding to DNS server.
	SourceIp pulumi.StringPtrInput
	// IPv6 source IP address for forwarding to DNS server.
	SourceIp6 pulumi.StringPtrInput
	// Enable/disable this DNS zone. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Default time-to-live value for the entries of this DNS zone (0 - 2147483647 sec, default = 86400).
	Ttl pulumi.IntInput
	// Zone type (master to manage entries directly, slave to import entries from other zones).
	Type pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Zone view (public to serve public clients, shadow to serve internal clients).
	View pulumi.StringInput
}

func (DnsdatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsdatabaseArgs)(nil)).Elem()
}

type DnsdatabaseInput interface {
	pulumi.Input

	ToDnsdatabaseOutput() DnsdatabaseOutput
	ToDnsdatabaseOutputWithContext(ctx context.Context) DnsdatabaseOutput
}

func (*Dnsdatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**Dnsdatabase)(nil)).Elem()
}

func (i *Dnsdatabase) ToDnsdatabaseOutput() DnsdatabaseOutput {
	return i.ToDnsdatabaseOutputWithContext(context.Background())
}

func (i *Dnsdatabase) ToDnsdatabaseOutputWithContext(ctx context.Context) DnsdatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsdatabaseOutput)
}

// DnsdatabaseArrayInput is an input type that accepts DnsdatabaseArray and DnsdatabaseArrayOutput values.
// You can construct a concrete instance of `DnsdatabaseArrayInput` via:
//
//	DnsdatabaseArray{ DnsdatabaseArgs{...} }
type DnsdatabaseArrayInput interface {
	pulumi.Input

	ToDnsdatabaseArrayOutput() DnsdatabaseArrayOutput
	ToDnsdatabaseArrayOutputWithContext(context.Context) DnsdatabaseArrayOutput
}

type DnsdatabaseArray []DnsdatabaseInput

func (DnsdatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dnsdatabase)(nil)).Elem()
}

func (i DnsdatabaseArray) ToDnsdatabaseArrayOutput() DnsdatabaseArrayOutput {
	return i.ToDnsdatabaseArrayOutputWithContext(context.Background())
}

func (i DnsdatabaseArray) ToDnsdatabaseArrayOutputWithContext(ctx context.Context) DnsdatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsdatabaseArrayOutput)
}

// DnsdatabaseMapInput is an input type that accepts DnsdatabaseMap and DnsdatabaseMapOutput values.
// You can construct a concrete instance of `DnsdatabaseMapInput` via:
//
//	DnsdatabaseMap{ "key": DnsdatabaseArgs{...} }
type DnsdatabaseMapInput interface {
	pulumi.Input

	ToDnsdatabaseMapOutput() DnsdatabaseMapOutput
	ToDnsdatabaseMapOutputWithContext(context.Context) DnsdatabaseMapOutput
}

type DnsdatabaseMap map[string]DnsdatabaseInput

func (DnsdatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dnsdatabase)(nil)).Elem()
}

func (i DnsdatabaseMap) ToDnsdatabaseMapOutput() DnsdatabaseMapOutput {
	return i.ToDnsdatabaseMapOutputWithContext(context.Background())
}

func (i DnsdatabaseMap) ToDnsdatabaseMapOutputWithContext(ctx context.Context) DnsdatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsdatabaseMapOutput)
}

type DnsdatabaseOutput struct{ *pulumi.OutputState }

func (DnsdatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dnsdatabase)(nil)).Elem()
}

func (o DnsdatabaseOutput) ToDnsdatabaseOutput() DnsdatabaseOutput {
	return o
}

func (o DnsdatabaseOutput) ToDnsdatabaseOutputWithContext(ctx context.Context) DnsdatabaseOutput {
	return o
}

// DNS zone transfer IP address list.
func (o DnsdatabaseOutput) AllowTransfer() pulumi.StringOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.StringOutput { return v.AllowTransfer }).(pulumi.StringOutput)
}

// Enable/disable authoritative zone. Valid values: `enable`, `disable`.
func (o DnsdatabaseOutput) Authoritative() pulumi.StringOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.StringOutput { return v.Authoritative }).(pulumi.StringOutput)
}

// Email address of the administrator for this zone.
// You can specify only the username (e.g. admin) or full email address (e.g. admin@test.com)
// When using a simple username, the domain of the email will be this zone.
func (o DnsdatabaseOutput) Contact() pulumi.StringOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.StringOutput { return v.Contact }).(pulumi.StringOutput)
}

// DNS entry. The structure of `dnsEntry` block is documented below.
func (o DnsdatabaseOutput) DnsEntries() DnsdatabaseDnsEntryArrayOutput {
	return o.ApplyT(func(v *Dnsdatabase) DnsdatabaseDnsEntryArrayOutput { return v.DnsEntries }).(DnsdatabaseDnsEntryArrayOutput)
}

// Domain name.
func (o DnsdatabaseOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o DnsdatabaseOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// DNS zone forwarder IP address list.
func (o DnsdatabaseOutput) Forwarder() pulumi.StringOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.StringOutput { return v.Forwarder }).(pulumi.StringOutput)
}

// Forwarder IPv6 address.
func (o DnsdatabaseOutput) Forwarder6() pulumi.StringOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.StringOutput { return v.Forwarder6 }).(pulumi.StringOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o DnsdatabaseOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.
func (o DnsdatabaseOutput) IpMaster() pulumi.StringOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.StringOutput { return v.IpMaster }).(pulumi.StringOutput)
}

// IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.
func (o DnsdatabaseOutput) IpPrimary() pulumi.StringOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.StringOutput { return v.IpPrimary }).(pulumi.StringOutput)
}

// Zone name.
func (o DnsdatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Domain name of the default DNS server for this zone.
func (o DnsdatabaseOutput) PrimaryName() pulumi.StringOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.StringOutput { return v.PrimaryName }).(pulumi.StringOutput)
}

// Maximum number of resource records (10 - 65536, 0 means infinite).
func (o DnsdatabaseOutput) RrMax() pulumi.IntOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.IntOutput { return v.RrMax }).(pulumi.IntOutput)
}

// Source IP for forwarding to DNS server.
func (o DnsdatabaseOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// IPv6 source IP address for forwarding to DNS server.
func (o DnsdatabaseOutput) SourceIp6() pulumi.StringOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.StringOutput { return v.SourceIp6 }).(pulumi.StringOutput)
}

// Enable/disable this DNS zone. Valid values: `enable`, `disable`.
func (o DnsdatabaseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Default time-to-live value for the entries of this DNS zone (0 - 2147483647 sec, default = 86400).
func (o DnsdatabaseOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.IntOutput { return v.Ttl }).(pulumi.IntOutput)
}

// Zone type (master to manage entries directly, slave to import entries from other zones).
func (o DnsdatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DnsdatabaseOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Zone view (public to serve public clients, shadow to serve internal clients).
func (o DnsdatabaseOutput) View() pulumi.StringOutput {
	return o.ApplyT(func(v *Dnsdatabase) pulumi.StringOutput { return v.View }).(pulumi.StringOutput)
}

type DnsdatabaseArrayOutput struct{ *pulumi.OutputState }

func (DnsdatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dnsdatabase)(nil)).Elem()
}

func (o DnsdatabaseArrayOutput) ToDnsdatabaseArrayOutput() DnsdatabaseArrayOutput {
	return o
}

func (o DnsdatabaseArrayOutput) ToDnsdatabaseArrayOutputWithContext(ctx context.Context) DnsdatabaseArrayOutput {
	return o
}

func (o DnsdatabaseArrayOutput) Index(i pulumi.IntInput) DnsdatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Dnsdatabase {
		return vs[0].([]*Dnsdatabase)[vs[1].(int)]
	}).(DnsdatabaseOutput)
}

type DnsdatabaseMapOutput struct{ *pulumi.OutputState }

func (DnsdatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dnsdatabase)(nil)).Elem()
}

func (o DnsdatabaseMapOutput) ToDnsdatabaseMapOutput() DnsdatabaseMapOutput {
	return o
}

func (o DnsdatabaseMapOutput) ToDnsdatabaseMapOutputWithContext(ctx context.Context) DnsdatabaseMapOutput {
	return o
}

func (o DnsdatabaseMapOutput) MapIndex(k pulumi.StringInput) DnsdatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Dnsdatabase {
		return vs[0].(map[string]*Dnsdatabase)[vs[1].(string)]
	}).(DnsdatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsdatabaseInput)(nil)).Elem(), &Dnsdatabase{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsdatabaseArrayInput)(nil)).Elem(), DnsdatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsdatabaseMapInput)(nil)).Elem(), DnsdatabaseMap{})
	pulumi.RegisterOutputType(DnsdatabaseOutput{})
	pulumi.RegisterOutputType(DnsdatabaseArrayOutput{})
	pulumi.RegisterOutputType(DnsdatabaseMapOutput{})
}
