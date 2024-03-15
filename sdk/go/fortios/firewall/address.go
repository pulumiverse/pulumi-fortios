// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure IPv4 addresses.
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
//			_, err := firewall.NewAddress(ctx, "trname", &firewall.AddressArgs{
//				AllowRouting:        pulumi.String("disable"),
//				AssociatedInterface: pulumi.String("port2"),
//				Color:               pulumi.Int(3),
//				EndIp:               pulumi.String("255.255.255.0"),
//				StartIp:             pulumi.String("22.1.1.0"),
//				Subnet:              pulumi.String("22.1.1.0 255.255.255.0"),
//				Type:                pulumi.String("ipmask"),
//				Visibility:          pulumi.String("enable"),
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
// Firewall Address can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/address:Address labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/address:Address labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Address struct {
	pulumi.CustomResourceState

	// Enable/disable use of this address in the static route configuration. Valid values: `enable`, `disable`.
	AllowRouting pulumi.StringOutput `pulumi:"allowRouting"`
	// Network interface associated with address.
	AssociatedInterface pulumi.StringPtrOutput `pulumi:"associatedInterface"`
	// Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
	CacheTtl pulumi.IntPtrOutput `pulumi:"cacheTtl"`
	// SPT (System Posture Token) value. Valid values: `unknown`, `healthy`, `quarantine`, `checkup`, `transient`, `infected`.
	ClearpassSpt pulumi.StringOutput `pulumi:"clearpassSpt"`
	// Color of icon on the GUI.
	Color pulumi.IntPtrOutput `pulumi:"color"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// IP addresses associated to a specific country.
	Country pulumi.StringPtrOutput `pulumi:"country"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Final IP address (inclusive) in the range for the address.
	EndIp pulumi.StringOutput `pulumi:"endIp"`
	// Last MAC address in the range.
	EndMac pulumi.StringOutput `pulumi:"endMac"`
	// Endpoint group name.
	EpgName pulumi.StringPtrOutput `pulumi:"epgName"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringOutput `pulumi:"fabricObject"`
	// Match criteria filter.
	Filter pulumi.StringPtrOutput `pulumi:"filter"`
	// Fully Qualified Domain Name address.
	Fqdn pulumi.StringPtrOutput `pulumi:"fqdn"`
	// FSSO group(s). The structure of `fssoGroup` block is documented below.
	FssoGroups AddressFssoGroupArrayOutput `pulumi:"fssoGroups"`
	// Name of interface whose IP address is to be used.
	Interface pulumi.StringPtrOutput `pulumi:"interface"`
	// IP address list. The structure of `list` block is documented below.
	Lists AddressListArrayOutput `pulumi:"lists"`
	// Multiple MAC address ranges. The structure of `macaddr` block is documented below.
	Macaddrs AddressMacaddrArrayOutput `pulumi:"macaddrs"`
	// Address name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable collection of node addresses only in Kubernetes. Valid values: `enable`, `disable`.
	NodeIpOnly pulumi.StringOutput `pulumi:"nodeIpOnly"`
	// Object ID for NSX.
	ObjId pulumi.StringPtrOutput `pulumi:"objId"`
	// Tag of dynamic address object.
	ObjTag pulumi.StringPtrOutput `pulumi:"objTag"`
	// Object type. Valid values: `ip`, `mac`.
	ObjType pulumi.StringOutput `pulumi:"objType"`
	// Organization domain name (Syntax: organization/domain).
	Organization pulumi.StringPtrOutput `pulumi:"organization"`
	// Policy group name.
	PolicyGroup pulumi.StringPtrOutput `pulumi:"policyGroup"`
	// SDN.
	Sdn pulumi.StringPtrOutput `pulumi:"sdn"`
	// Type of addresses to collect. Valid values: `private`, `public`, `all`.
	SdnAddrType pulumi.StringOutput `pulumi:"sdnAddrType"`
	// SDN Tag.
	SdnTag pulumi.StringPtrOutput `pulumi:"sdnTag"`
	// First IP address (inclusive) in the range for the address.
	StartIp pulumi.StringOutput `pulumi:"startIp"`
	// First MAC address in the range.
	StartMac pulumi.StringOutput `pulumi:"startMac"`
	// Sub-type of address.
	SubType pulumi.StringOutput `pulumi:"subType"`
	// IP address and subnet mask of address.
	Subnet pulumi.StringOutput `pulumi:"subnet"`
	// Subnet name.
	SubnetName pulumi.StringPtrOutput `pulumi:"subnetName"`
	// Tag detection level of dynamic address object.
	TagDetectionLevel pulumi.StringPtrOutput `pulumi:"tagDetectionLevel"`
	// Tag type of dynamic address object.
	TagType pulumi.StringPtrOutput `pulumi:"tagType"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings AddressTaggingArrayOutput `pulumi:"taggings"`
	// Tenant.
	Tenant pulumi.StringPtrOutput `pulumi:"tenant"`
	// Type of address.
	Type pulumi.StringOutput `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrOutput `pulumi:"visibility"`
	// IP address and wildcard netmask.
	Wildcard pulumi.StringOutput `pulumi:"wildcard"`
	// Fully Qualified Domain Name with wildcard characters.
	WildcardFqdn pulumi.StringPtrOutput `pulumi:"wildcardFqdn"`
}

// NewAddress registers a new resource with the given unique name, arguments, and options.
func NewAddress(ctx *pulumi.Context,
	name string, args *AddressArgs, opts ...pulumi.ResourceOption) (*Address, error) {
	if args == nil {
		args = &AddressArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Address
	err := ctx.RegisterResource("fortios:firewall/address:Address", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAddress gets an existing Address resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddressState, opts ...pulumi.ResourceOption) (*Address, error) {
	var resource Address
	err := ctx.ReadResource("fortios:firewall/address:Address", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Address resources.
type addressState struct {
	// Enable/disable use of this address in the static route configuration. Valid values: `enable`, `disable`.
	AllowRouting *string `pulumi:"allowRouting"`
	// Network interface associated with address.
	AssociatedInterface *string `pulumi:"associatedInterface"`
	// Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
	CacheTtl *int `pulumi:"cacheTtl"`
	// SPT (System Posture Token) value. Valid values: `unknown`, `healthy`, `quarantine`, `checkup`, `transient`, `infected`.
	ClearpassSpt *string `pulumi:"clearpassSpt"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// IP addresses associated to a specific country.
	Country *string `pulumi:"country"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Final IP address (inclusive) in the range for the address.
	EndIp *string `pulumi:"endIp"`
	// Last MAC address in the range.
	EndMac *string `pulumi:"endMac"`
	// Endpoint group name.
	EpgName *string `pulumi:"epgName"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// Match criteria filter.
	Filter *string `pulumi:"filter"`
	// Fully Qualified Domain Name address.
	Fqdn *string `pulumi:"fqdn"`
	// FSSO group(s). The structure of `fssoGroup` block is documented below.
	FssoGroups []AddressFssoGroup `pulumi:"fssoGroups"`
	// Name of interface whose IP address is to be used.
	Interface *string `pulumi:"interface"`
	// IP address list. The structure of `list` block is documented below.
	Lists []AddressList `pulumi:"lists"`
	// Multiple MAC address ranges. The structure of `macaddr` block is documented below.
	Macaddrs []AddressMacaddr `pulumi:"macaddrs"`
	// Address name.
	Name *string `pulumi:"name"`
	// Enable/disable collection of node addresses only in Kubernetes. Valid values: `enable`, `disable`.
	NodeIpOnly *string `pulumi:"nodeIpOnly"`
	// Object ID for NSX.
	ObjId *string `pulumi:"objId"`
	// Tag of dynamic address object.
	ObjTag *string `pulumi:"objTag"`
	// Object type. Valid values: `ip`, `mac`.
	ObjType *string `pulumi:"objType"`
	// Organization domain name (Syntax: organization/domain).
	Organization *string `pulumi:"organization"`
	// Policy group name.
	PolicyGroup *string `pulumi:"policyGroup"`
	// SDN.
	Sdn *string `pulumi:"sdn"`
	// Type of addresses to collect. Valid values: `private`, `public`, `all`.
	SdnAddrType *string `pulumi:"sdnAddrType"`
	// SDN Tag.
	SdnTag *string `pulumi:"sdnTag"`
	// First IP address (inclusive) in the range for the address.
	StartIp *string `pulumi:"startIp"`
	// First MAC address in the range.
	StartMac *string `pulumi:"startMac"`
	// Sub-type of address.
	SubType *string `pulumi:"subType"`
	// IP address and subnet mask of address.
	Subnet *string `pulumi:"subnet"`
	// Subnet name.
	SubnetName *string `pulumi:"subnetName"`
	// Tag detection level of dynamic address object.
	TagDetectionLevel *string `pulumi:"tagDetectionLevel"`
	// Tag type of dynamic address object.
	TagType *string `pulumi:"tagType"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []AddressTagging `pulumi:"taggings"`
	// Tenant.
	Tenant *string `pulumi:"tenant"`
	// Type of address.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
	// IP address and wildcard netmask.
	Wildcard *string `pulumi:"wildcard"`
	// Fully Qualified Domain Name with wildcard characters.
	WildcardFqdn *string `pulumi:"wildcardFqdn"`
}

type AddressState struct {
	// Enable/disable use of this address in the static route configuration. Valid values: `enable`, `disable`.
	AllowRouting pulumi.StringPtrInput
	// Network interface associated with address.
	AssociatedInterface pulumi.StringPtrInput
	// Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
	CacheTtl pulumi.IntPtrInput
	// SPT (System Posture Token) value. Valid values: `unknown`, `healthy`, `quarantine`, `checkup`, `transient`, `infected`.
	ClearpassSpt pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// IP addresses associated to a specific country.
	Country pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Final IP address (inclusive) in the range for the address.
	EndIp pulumi.StringPtrInput
	// Last MAC address in the range.
	EndMac pulumi.StringPtrInput
	// Endpoint group name.
	EpgName pulumi.StringPtrInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// Match criteria filter.
	Filter pulumi.StringPtrInput
	// Fully Qualified Domain Name address.
	Fqdn pulumi.StringPtrInput
	// FSSO group(s). The structure of `fssoGroup` block is documented below.
	FssoGroups AddressFssoGroupArrayInput
	// Name of interface whose IP address is to be used.
	Interface pulumi.StringPtrInput
	// IP address list. The structure of `list` block is documented below.
	Lists AddressListArrayInput
	// Multiple MAC address ranges. The structure of `macaddr` block is documented below.
	Macaddrs AddressMacaddrArrayInput
	// Address name.
	Name pulumi.StringPtrInput
	// Enable/disable collection of node addresses only in Kubernetes. Valid values: `enable`, `disable`.
	NodeIpOnly pulumi.StringPtrInput
	// Object ID for NSX.
	ObjId pulumi.StringPtrInput
	// Tag of dynamic address object.
	ObjTag pulumi.StringPtrInput
	// Object type. Valid values: `ip`, `mac`.
	ObjType pulumi.StringPtrInput
	// Organization domain name (Syntax: organization/domain).
	Organization pulumi.StringPtrInput
	// Policy group name.
	PolicyGroup pulumi.StringPtrInput
	// SDN.
	Sdn pulumi.StringPtrInput
	// Type of addresses to collect. Valid values: `private`, `public`, `all`.
	SdnAddrType pulumi.StringPtrInput
	// SDN Tag.
	SdnTag pulumi.StringPtrInput
	// First IP address (inclusive) in the range for the address.
	StartIp pulumi.StringPtrInput
	// First MAC address in the range.
	StartMac pulumi.StringPtrInput
	// Sub-type of address.
	SubType pulumi.StringPtrInput
	// IP address and subnet mask of address.
	Subnet pulumi.StringPtrInput
	// Subnet name.
	SubnetName pulumi.StringPtrInput
	// Tag detection level of dynamic address object.
	TagDetectionLevel pulumi.StringPtrInput
	// Tag type of dynamic address object.
	TagType pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings AddressTaggingArrayInput
	// Tenant.
	Tenant pulumi.StringPtrInput
	// Type of address.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
	// IP address and wildcard netmask.
	Wildcard pulumi.StringPtrInput
	// Fully Qualified Domain Name with wildcard characters.
	WildcardFqdn pulumi.StringPtrInput
}

func (AddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*addressState)(nil)).Elem()
}

type addressArgs struct {
	// Enable/disable use of this address in the static route configuration. Valid values: `enable`, `disable`.
	AllowRouting *string `pulumi:"allowRouting"`
	// Network interface associated with address.
	AssociatedInterface *string `pulumi:"associatedInterface"`
	// Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
	CacheTtl *int `pulumi:"cacheTtl"`
	// SPT (System Posture Token) value. Valid values: `unknown`, `healthy`, `quarantine`, `checkup`, `transient`, `infected`.
	ClearpassSpt *string `pulumi:"clearpassSpt"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// IP addresses associated to a specific country.
	Country *string `pulumi:"country"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Final IP address (inclusive) in the range for the address.
	EndIp *string `pulumi:"endIp"`
	// Last MAC address in the range.
	EndMac *string `pulumi:"endMac"`
	// Endpoint group name.
	EpgName *string `pulumi:"epgName"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// Match criteria filter.
	Filter *string `pulumi:"filter"`
	// Fully Qualified Domain Name address.
	Fqdn *string `pulumi:"fqdn"`
	// FSSO group(s). The structure of `fssoGroup` block is documented below.
	FssoGroups []AddressFssoGroup `pulumi:"fssoGroups"`
	// Name of interface whose IP address is to be used.
	Interface *string `pulumi:"interface"`
	// IP address list. The structure of `list` block is documented below.
	Lists []AddressList `pulumi:"lists"`
	// Multiple MAC address ranges. The structure of `macaddr` block is documented below.
	Macaddrs []AddressMacaddr `pulumi:"macaddrs"`
	// Address name.
	Name *string `pulumi:"name"`
	// Enable/disable collection of node addresses only in Kubernetes. Valid values: `enable`, `disable`.
	NodeIpOnly *string `pulumi:"nodeIpOnly"`
	// Object ID for NSX.
	ObjId *string `pulumi:"objId"`
	// Tag of dynamic address object.
	ObjTag *string `pulumi:"objTag"`
	// Object type. Valid values: `ip`, `mac`.
	ObjType *string `pulumi:"objType"`
	// Organization domain name (Syntax: organization/domain).
	Organization *string `pulumi:"organization"`
	// Policy group name.
	PolicyGroup *string `pulumi:"policyGroup"`
	// SDN.
	Sdn *string `pulumi:"sdn"`
	// Type of addresses to collect. Valid values: `private`, `public`, `all`.
	SdnAddrType *string `pulumi:"sdnAddrType"`
	// SDN Tag.
	SdnTag *string `pulumi:"sdnTag"`
	// First IP address (inclusive) in the range for the address.
	StartIp *string `pulumi:"startIp"`
	// First MAC address in the range.
	StartMac *string `pulumi:"startMac"`
	// Sub-type of address.
	SubType *string `pulumi:"subType"`
	// IP address and subnet mask of address.
	Subnet *string `pulumi:"subnet"`
	// Subnet name.
	SubnetName *string `pulumi:"subnetName"`
	// Tag detection level of dynamic address object.
	TagDetectionLevel *string `pulumi:"tagDetectionLevel"`
	// Tag type of dynamic address object.
	TagType *string `pulumi:"tagType"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []AddressTagging `pulumi:"taggings"`
	// Tenant.
	Tenant *string `pulumi:"tenant"`
	// Type of address.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
	// IP address and wildcard netmask.
	Wildcard *string `pulumi:"wildcard"`
	// Fully Qualified Domain Name with wildcard characters.
	WildcardFqdn *string `pulumi:"wildcardFqdn"`
}

// The set of arguments for constructing a Address resource.
type AddressArgs struct {
	// Enable/disable use of this address in the static route configuration. Valid values: `enable`, `disable`.
	AllowRouting pulumi.StringPtrInput
	// Network interface associated with address.
	AssociatedInterface pulumi.StringPtrInput
	// Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
	CacheTtl pulumi.IntPtrInput
	// SPT (System Posture Token) value. Valid values: `unknown`, `healthy`, `quarantine`, `checkup`, `transient`, `infected`.
	ClearpassSpt pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// IP addresses associated to a specific country.
	Country pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Final IP address (inclusive) in the range for the address.
	EndIp pulumi.StringPtrInput
	// Last MAC address in the range.
	EndMac pulumi.StringPtrInput
	// Endpoint group name.
	EpgName pulumi.StringPtrInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// Match criteria filter.
	Filter pulumi.StringPtrInput
	// Fully Qualified Domain Name address.
	Fqdn pulumi.StringPtrInput
	// FSSO group(s). The structure of `fssoGroup` block is documented below.
	FssoGroups AddressFssoGroupArrayInput
	// Name of interface whose IP address is to be used.
	Interface pulumi.StringPtrInput
	// IP address list. The structure of `list` block is documented below.
	Lists AddressListArrayInput
	// Multiple MAC address ranges. The structure of `macaddr` block is documented below.
	Macaddrs AddressMacaddrArrayInput
	// Address name.
	Name pulumi.StringPtrInput
	// Enable/disable collection of node addresses only in Kubernetes. Valid values: `enable`, `disable`.
	NodeIpOnly pulumi.StringPtrInput
	// Object ID for NSX.
	ObjId pulumi.StringPtrInput
	// Tag of dynamic address object.
	ObjTag pulumi.StringPtrInput
	// Object type. Valid values: `ip`, `mac`.
	ObjType pulumi.StringPtrInput
	// Organization domain name (Syntax: organization/domain).
	Organization pulumi.StringPtrInput
	// Policy group name.
	PolicyGroup pulumi.StringPtrInput
	// SDN.
	Sdn pulumi.StringPtrInput
	// Type of addresses to collect. Valid values: `private`, `public`, `all`.
	SdnAddrType pulumi.StringPtrInput
	// SDN Tag.
	SdnTag pulumi.StringPtrInput
	// First IP address (inclusive) in the range for the address.
	StartIp pulumi.StringPtrInput
	// First MAC address in the range.
	StartMac pulumi.StringPtrInput
	// Sub-type of address.
	SubType pulumi.StringPtrInput
	// IP address and subnet mask of address.
	Subnet pulumi.StringPtrInput
	// Subnet name.
	SubnetName pulumi.StringPtrInput
	// Tag detection level of dynamic address object.
	TagDetectionLevel pulumi.StringPtrInput
	// Tag type of dynamic address object.
	TagType pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings AddressTaggingArrayInput
	// Tenant.
	Tenant pulumi.StringPtrInput
	// Type of address.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
	// IP address and wildcard netmask.
	Wildcard pulumi.StringPtrInput
	// Fully Qualified Domain Name with wildcard characters.
	WildcardFqdn pulumi.StringPtrInput
}

func (AddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addressArgs)(nil)).Elem()
}

type AddressInput interface {
	pulumi.Input

	ToAddressOutput() AddressOutput
	ToAddressOutputWithContext(ctx context.Context) AddressOutput
}

func (*Address) ElementType() reflect.Type {
	return reflect.TypeOf((**Address)(nil)).Elem()
}

func (i *Address) ToAddressOutput() AddressOutput {
	return i.ToAddressOutputWithContext(context.Background())
}

func (i *Address) ToAddressOutputWithContext(ctx context.Context) AddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressOutput)
}

// AddressArrayInput is an input type that accepts AddressArray and AddressArrayOutput values.
// You can construct a concrete instance of `AddressArrayInput` via:
//
//	AddressArray{ AddressArgs{...} }
type AddressArrayInput interface {
	pulumi.Input

	ToAddressArrayOutput() AddressArrayOutput
	ToAddressArrayOutputWithContext(context.Context) AddressArrayOutput
}

type AddressArray []AddressInput

func (AddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Address)(nil)).Elem()
}

func (i AddressArray) ToAddressArrayOutput() AddressArrayOutput {
	return i.ToAddressArrayOutputWithContext(context.Background())
}

func (i AddressArray) ToAddressArrayOutputWithContext(ctx context.Context) AddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressArrayOutput)
}

// AddressMapInput is an input type that accepts AddressMap and AddressMapOutput values.
// You can construct a concrete instance of `AddressMapInput` via:
//
//	AddressMap{ "key": AddressArgs{...} }
type AddressMapInput interface {
	pulumi.Input

	ToAddressMapOutput() AddressMapOutput
	ToAddressMapOutputWithContext(context.Context) AddressMapOutput
}

type AddressMap map[string]AddressInput

func (AddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Address)(nil)).Elem()
}

func (i AddressMap) ToAddressMapOutput() AddressMapOutput {
	return i.ToAddressMapOutputWithContext(context.Background())
}

func (i AddressMap) ToAddressMapOutputWithContext(ctx context.Context) AddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressMapOutput)
}

type AddressOutput struct{ *pulumi.OutputState }

func (AddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Address)(nil)).Elem()
}

func (o AddressOutput) ToAddressOutput() AddressOutput {
	return o
}

func (o AddressOutput) ToAddressOutputWithContext(ctx context.Context) AddressOutput {
	return o
}

// Enable/disable use of this address in the static route configuration. Valid values: `enable`, `disable`.
func (o AddressOutput) AllowRouting() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.AllowRouting }).(pulumi.StringOutput)
}

// Network interface associated with address.
func (o AddressOutput) AssociatedInterface() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.AssociatedInterface }).(pulumi.StringPtrOutput)
}

// Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
func (o AddressOutput) CacheTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.IntPtrOutput { return v.CacheTtl }).(pulumi.IntPtrOutput)
}

// SPT (System Posture Token) value. Valid values: `unknown`, `healthy`, `quarantine`, `checkup`, `transient`, `infected`.
func (o AddressOutput) ClearpassSpt() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.ClearpassSpt }).(pulumi.StringOutput)
}

// Color of icon on the GUI.
func (o AddressOutput) Color() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.IntPtrOutput { return v.Color }).(pulumi.IntPtrOutput)
}

// Comment.
func (o AddressOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// IP addresses associated to a specific country.
func (o AddressOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.Country }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o AddressOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Final IP address (inclusive) in the range for the address.
func (o AddressOutput) EndIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.EndIp }).(pulumi.StringOutput)
}

// Last MAC address in the range.
func (o AddressOutput) EndMac() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.EndMac }).(pulumi.StringOutput)
}

// Endpoint group name.
func (o AddressOutput) EpgName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.EpgName }).(pulumi.StringPtrOutput)
}

// Security Fabric global object setting. Valid values: `enable`, `disable`.
func (o AddressOutput) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.FabricObject }).(pulumi.StringOutput)
}

// Match criteria filter.
func (o AddressOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.Filter }).(pulumi.StringPtrOutput)
}

// Fully Qualified Domain Name address.
func (o AddressOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.Fqdn }).(pulumi.StringPtrOutput)
}

// FSSO group(s). The structure of `fssoGroup` block is documented below.
func (o AddressOutput) FssoGroups() AddressFssoGroupArrayOutput {
	return o.ApplyT(func(v *Address) AddressFssoGroupArrayOutput { return v.FssoGroups }).(AddressFssoGroupArrayOutput)
}

// Name of interface whose IP address is to be used.
func (o AddressOutput) Interface() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.Interface }).(pulumi.StringPtrOutput)
}

// IP address list. The structure of `list` block is documented below.
func (o AddressOutput) Lists() AddressListArrayOutput {
	return o.ApplyT(func(v *Address) AddressListArrayOutput { return v.Lists }).(AddressListArrayOutput)
}

// Multiple MAC address ranges. The structure of `macaddr` block is documented below.
func (o AddressOutput) Macaddrs() AddressMacaddrArrayOutput {
	return o.ApplyT(func(v *Address) AddressMacaddrArrayOutput { return v.Macaddrs }).(AddressMacaddrArrayOutput)
}

// Address name.
func (o AddressOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable collection of node addresses only in Kubernetes. Valid values: `enable`, `disable`.
func (o AddressOutput) NodeIpOnly() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.NodeIpOnly }).(pulumi.StringOutput)
}

// Object ID for NSX.
func (o AddressOutput) ObjId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.ObjId }).(pulumi.StringPtrOutput)
}

// Tag of dynamic address object.
func (o AddressOutput) ObjTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.ObjTag }).(pulumi.StringPtrOutput)
}

// Object type. Valid values: `ip`, `mac`.
func (o AddressOutput) ObjType() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.ObjType }).(pulumi.StringOutput)
}

// Organization domain name (Syntax: organization/domain).
func (o AddressOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.Organization }).(pulumi.StringPtrOutput)
}

// Policy group name.
func (o AddressOutput) PolicyGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.PolicyGroup }).(pulumi.StringPtrOutput)
}

// SDN.
func (o AddressOutput) Sdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.Sdn }).(pulumi.StringPtrOutput)
}

// Type of addresses to collect. Valid values: `private`, `public`, `all`.
func (o AddressOutput) SdnAddrType() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.SdnAddrType }).(pulumi.StringOutput)
}

// SDN Tag.
func (o AddressOutput) SdnTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.SdnTag }).(pulumi.StringPtrOutput)
}

// First IP address (inclusive) in the range for the address.
func (o AddressOutput) StartIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.StartIp }).(pulumi.StringOutput)
}

// First MAC address in the range.
func (o AddressOutput) StartMac() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.StartMac }).(pulumi.StringOutput)
}

// Sub-type of address.
func (o AddressOutput) SubType() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.SubType }).(pulumi.StringOutput)
}

// IP address and subnet mask of address.
func (o AddressOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

// Subnet name.
func (o AddressOutput) SubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.SubnetName }).(pulumi.StringPtrOutput)
}

// Tag detection level of dynamic address object.
func (o AddressOutput) TagDetectionLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.TagDetectionLevel }).(pulumi.StringPtrOutput)
}

// Tag type of dynamic address object.
func (o AddressOutput) TagType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.TagType }).(pulumi.StringPtrOutput)
}

// Config object tagging. The structure of `tagging` block is documented below.
func (o AddressOutput) Taggings() AddressTaggingArrayOutput {
	return o.ApplyT(func(v *Address) AddressTaggingArrayOutput { return v.Taggings }).(AddressTaggingArrayOutput)
}

// Tenant.
func (o AddressOutput) Tenant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.Tenant }).(pulumi.StringPtrOutput)
}

// Type of address.
func (o AddressOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o AddressOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o AddressOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
func (o AddressOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.Visibility }).(pulumi.StringPtrOutput)
}

// IP address and wildcard netmask.
func (o AddressOutput) Wildcard() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Wildcard }).(pulumi.StringOutput)
}

// Fully Qualified Domain Name with wildcard characters.
func (o AddressOutput) WildcardFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.WildcardFqdn }).(pulumi.StringPtrOutput)
}

type AddressArrayOutput struct{ *pulumi.OutputState }

func (AddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Address)(nil)).Elem()
}

func (o AddressArrayOutput) ToAddressArrayOutput() AddressArrayOutput {
	return o
}

func (o AddressArrayOutput) ToAddressArrayOutputWithContext(ctx context.Context) AddressArrayOutput {
	return o
}

func (o AddressArrayOutput) Index(i pulumi.IntInput) AddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Address {
		return vs[0].([]*Address)[vs[1].(int)]
	}).(AddressOutput)
}

type AddressMapOutput struct{ *pulumi.OutputState }

func (AddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Address)(nil)).Elem()
}

func (o AddressMapOutput) ToAddressMapOutput() AddressMapOutput {
	return o
}

func (o AddressMapOutput) ToAddressMapOutputWithContext(ctx context.Context) AddressMapOutput {
	return o
}

func (o AddressMapOutput) MapIndex(k pulumi.StringInput) AddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Address {
		return vs[0].(map[string]*Address)[vs[1].(string)]
	}).(AddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AddressInput)(nil)).Elem(), &Address{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressArrayInput)(nil)).Elem(), AddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressMapInput)(nil)).Elem(), AddressMap{})
	pulumi.RegisterOutputType(AddressOutput{})
	pulumi.RegisterOutputType(AddressArrayOutput{})
	pulumi.RegisterOutputType(AddressMapOutput{})
}
