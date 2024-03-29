// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package snmp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// SNMP community configuration.
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
//			_, err := system.NewCommunity(ctx, "trname", &system.CommunityArgs{
//				Events:         pulumi.String("cpu-high mem-low log-full intf-ip vpn-tun-up vpn-tun-down ha-switch ha-hb-failure ips-signature ips-anomaly av-virus av-oversize av-pattern av-fragmented fm-if-change bgp-established bgp-backward-transition ha-member-up ha-member-down ent-conf-change av-conserve av-bypass av-oversize-passed av-oversize-blocked ips-pkg-update ips-fail-open faz-disconnect wc-ap-up wc-ap-down fswctl-session-up fswctl-session-down load-balance-real-server-down per-cpu-high"),
//				Fosid:          pulumi.Int(1),
//				QueryV1Port:    pulumi.Int(161),
//				QueryV1Status:  pulumi.String("enable"),
//				QueryV2cPort:   pulumi.Int(161),
//				QueryV2cStatus: pulumi.String("enable"),
//				Status:         pulumi.String("enable"),
//				TrapV1Lport:    pulumi.Int(162),
//				TrapV1Rport:    pulumi.Int(162),
//				TrapV1Status:   pulumi.String("enable"),
//				TrapV2cLport:   pulumi.Int(162),
//				TrapV2cRport:   pulumi.Int(162),
//				TrapV2cStatus:  pulumi.String("enable"),
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
// SystemSnmp Community can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/snmp/community:Community labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/snmp/community:Community labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Community struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// SNMP trap events.
	Events pulumi.StringOutput `pulumi:"events"`
	// Community ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
	Hosts CommunityHostArrayOutput `pulumi:"hosts"`
	// Configure IPv6 SNMP managers. The structure of `hosts6` block is documented below.
	Hosts6s CommunityHosts6ArrayOutput `pulumi:"hosts6s"`
	// SNMP access control MIB view.
	MibView pulumi.StringOutput `pulumi:"mibView"`
	// Community name.
	Name pulumi.StringOutput `pulumi:"name"`
	// SNMP v1 query port (default = 161).
	QueryV1Port pulumi.IntOutput `pulumi:"queryV1Port"`
	// Enable/disable SNMP v1 queries. Valid values: `enable`, `disable`.
	QueryV1Status pulumi.StringOutput `pulumi:"queryV1Status"`
	// SNMP v2c query port (default = 161).
	QueryV2cPort pulumi.IntOutput `pulumi:"queryV2cPort"`
	// Enable/disable SNMP v2c queries. Valid values: `enable`, `disable`.
	QueryV2cStatus pulumi.StringOutput `pulumi:"queryV2cStatus"`
	// Enable/disable this SNMP community. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// SNMP v1 trap local port (default = 162).
	TrapV1Lport pulumi.IntOutput `pulumi:"trapV1Lport"`
	// SNMP v1 trap remote port (default = 162).
	TrapV1Rport pulumi.IntOutput `pulumi:"trapV1Rport"`
	// Enable/disable SNMP v1 traps. Valid values: `enable`, `disable`.
	TrapV1Status pulumi.StringOutput `pulumi:"trapV1Status"`
	// SNMP v2c trap local port (default = 162).
	TrapV2cLport pulumi.IntOutput `pulumi:"trapV2cLport"`
	// SNMP v2c trap remote port (default = 162).
	TrapV2cRport pulumi.IntOutput `pulumi:"trapV2cRport"`
	// Enable/disable SNMP v2c traps. Valid values: `enable`, `disable`.
	TrapV2cStatus pulumi.StringOutput `pulumi:"trapV2cStatus"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// SNMP access control VDOMs. The structure of `vdoms` block is documented below.
	Vdoms CommunityVdomArrayOutput `pulumi:"vdoms"`
}

// NewCommunity registers a new resource with the given unique name, arguments, and options.
func NewCommunity(ctx *pulumi.Context,
	name string, args *CommunityArgs, opts ...pulumi.ResourceOption) (*Community, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Community
	err := ctx.RegisterResource("fortios:system/snmp/community:Community", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCommunity gets an existing Community resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCommunity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommunityState, opts ...pulumi.ResourceOption) (*Community, error) {
	var resource Community
	err := ctx.ReadResource("fortios:system/snmp/community:Community", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Community resources.
type communityState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// SNMP trap events.
	Events *string `pulumi:"events"`
	// Community ID.
	Fosid *int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
	Hosts []CommunityHost `pulumi:"hosts"`
	// Configure IPv6 SNMP managers. The structure of `hosts6` block is documented below.
	Hosts6s []CommunityHosts6 `pulumi:"hosts6s"`
	// SNMP access control MIB view.
	MibView *string `pulumi:"mibView"`
	// Community name.
	Name *string `pulumi:"name"`
	// SNMP v1 query port (default = 161).
	QueryV1Port *int `pulumi:"queryV1Port"`
	// Enable/disable SNMP v1 queries. Valid values: `enable`, `disable`.
	QueryV1Status *string `pulumi:"queryV1Status"`
	// SNMP v2c query port (default = 161).
	QueryV2cPort *int `pulumi:"queryV2cPort"`
	// Enable/disable SNMP v2c queries. Valid values: `enable`, `disable`.
	QueryV2cStatus *string `pulumi:"queryV2cStatus"`
	// Enable/disable this SNMP community. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// SNMP v1 trap local port (default = 162).
	TrapV1Lport *int `pulumi:"trapV1Lport"`
	// SNMP v1 trap remote port (default = 162).
	TrapV1Rport *int `pulumi:"trapV1Rport"`
	// Enable/disable SNMP v1 traps. Valid values: `enable`, `disable`.
	TrapV1Status *string `pulumi:"trapV1Status"`
	// SNMP v2c trap local port (default = 162).
	TrapV2cLport *int `pulumi:"trapV2cLport"`
	// SNMP v2c trap remote port (default = 162).
	TrapV2cRport *int `pulumi:"trapV2cRport"`
	// Enable/disable SNMP v2c traps. Valid values: `enable`, `disable`.
	TrapV2cStatus *string `pulumi:"trapV2cStatus"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// SNMP access control VDOMs. The structure of `vdoms` block is documented below.
	Vdoms []CommunityVdom `pulumi:"vdoms"`
}

type CommunityState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// SNMP trap events.
	Events pulumi.StringPtrInput
	// Community ID.
	Fosid pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
	Hosts CommunityHostArrayInput
	// Configure IPv6 SNMP managers. The structure of `hosts6` block is documented below.
	Hosts6s CommunityHosts6ArrayInput
	// SNMP access control MIB view.
	MibView pulumi.StringPtrInput
	// Community name.
	Name pulumi.StringPtrInput
	// SNMP v1 query port (default = 161).
	QueryV1Port pulumi.IntPtrInput
	// Enable/disable SNMP v1 queries. Valid values: `enable`, `disable`.
	QueryV1Status pulumi.StringPtrInput
	// SNMP v2c query port (default = 161).
	QueryV2cPort pulumi.IntPtrInput
	// Enable/disable SNMP v2c queries. Valid values: `enable`, `disable`.
	QueryV2cStatus pulumi.StringPtrInput
	// Enable/disable this SNMP community. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// SNMP v1 trap local port (default = 162).
	TrapV1Lport pulumi.IntPtrInput
	// SNMP v1 trap remote port (default = 162).
	TrapV1Rport pulumi.IntPtrInput
	// Enable/disable SNMP v1 traps. Valid values: `enable`, `disable`.
	TrapV1Status pulumi.StringPtrInput
	// SNMP v2c trap local port (default = 162).
	TrapV2cLport pulumi.IntPtrInput
	// SNMP v2c trap remote port (default = 162).
	TrapV2cRport pulumi.IntPtrInput
	// Enable/disable SNMP v2c traps. Valid values: `enable`, `disable`.
	TrapV2cStatus pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// SNMP access control VDOMs. The structure of `vdoms` block is documented below.
	Vdoms CommunityVdomArrayInput
}

func (CommunityState) ElementType() reflect.Type {
	return reflect.TypeOf((*communityState)(nil)).Elem()
}

type communityArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// SNMP trap events.
	Events *string `pulumi:"events"`
	// Community ID.
	Fosid int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
	Hosts []CommunityHost `pulumi:"hosts"`
	// Configure IPv6 SNMP managers. The structure of `hosts6` block is documented below.
	Hosts6s []CommunityHosts6 `pulumi:"hosts6s"`
	// SNMP access control MIB view.
	MibView *string `pulumi:"mibView"`
	// Community name.
	Name *string `pulumi:"name"`
	// SNMP v1 query port (default = 161).
	QueryV1Port *int `pulumi:"queryV1Port"`
	// Enable/disable SNMP v1 queries. Valid values: `enable`, `disable`.
	QueryV1Status *string `pulumi:"queryV1Status"`
	// SNMP v2c query port (default = 161).
	QueryV2cPort *int `pulumi:"queryV2cPort"`
	// Enable/disable SNMP v2c queries. Valid values: `enable`, `disable`.
	QueryV2cStatus *string `pulumi:"queryV2cStatus"`
	// Enable/disable this SNMP community. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// SNMP v1 trap local port (default = 162).
	TrapV1Lport *int `pulumi:"trapV1Lport"`
	// SNMP v1 trap remote port (default = 162).
	TrapV1Rport *int `pulumi:"trapV1Rport"`
	// Enable/disable SNMP v1 traps. Valid values: `enable`, `disable`.
	TrapV1Status *string `pulumi:"trapV1Status"`
	// SNMP v2c trap local port (default = 162).
	TrapV2cLport *int `pulumi:"trapV2cLport"`
	// SNMP v2c trap remote port (default = 162).
	TrapV2cRport *int `pulumi:"trapV2cRport"`
	// Enable/disable SNMP v2c traps. Valid values: `enable`, `disable`.
	TrapV2cStatus *string `pulumi:"trapV2cStatus"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// SNMP access control VDOMs. The structure of `vdoms` block is documented below.
	Vdoms []CommunityVdom `pulumi:"vdoms"`
}

// The set of arguments for constructing a Community resource.
type CommunityArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// SNMP trap events.
	Events pulumi.StringPtrInput
	// Community ID.
	Fosid pulumi.IntInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
	Hosts CommunityHostArrayInput
	// Configure IPv6 SNMP managers. The structure of `hosts6` block is documented below.
	Hosts6s CommunityHosts6ArrayInput
	// SNMP access control MIB view.
	MibView pulumi.StringPtrInput
	// Community name.
	Name pulumi.StringPtrInput
	// SNMP v1 query port (default = 161).
	QueryV1Port pulumi.IntPtrInput
	// Enable/disable SNMP v1 queries. Valid values: `enable`, `disable`.
	QueryV1Status pulumi.StringPtrInput
	// SNMP v2c query port (default = 161).
	QueryV2cPort pulumi.IntPtrInput
	// Enable/disable SNMP v2c queries. Valid values: `enable`, `disable`.
	QueryV2cStatus pulumi.StringPtrInput
	// Enable/disable this SNMP community. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// SNMP v1 trap local port (default = 162).
	TrapV1Lport pulumi.IntPtrInput
	// SNMP v1 trap remote port (default = 162).
	TrapV1Rport pulumi.IntPtrInput
	// Enable/disable SNMP v1 traps. Valid values: `enable`, `disable`.
	TrapV1Status pulumi.StringPtrInput
	// SNMP v2c trap local port (default = 162).
	TrapV2cLport pulumi.IntPtrInput
	// SNMP v2c trap remote port (default = 162).
	TrapV2cRport pulumi.IntPtrInput
	// Enable/disable SNMP v2c traps. Valid values: `enable`, `disable`.
	TrapV2cStatus pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// SNMP access control VDOMs. The structure of `vdoms` block is documented below.
	Vdoms CommunityVdomArrayInput
}

func (CommunityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*communityArgs)(nil)).Elem()
}

type CommunityInput interface {
	pulumi.Input

	ToCommunityOutput() CommunityOutput
	ToCommunityOutputWithContext(ctx context.Context) CommunityOutput
}

func (*Community) ElementType() reflect.Type {
	return reflect.TypeOf((**Community)(nil)).Elem()
}

func (i *Community) ToCommunityOutput() CommunityOutput {
	return i.ToCommunityOutputWithContext(context.Background())
}

func (i *Community) ToCommunityOutputWithContext(ctx context.Context) CommunityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommunityOutput)
}

// CommunityArrayInput is an input type that accepts CommunityArray and CommunityArrayOutput values.
// You can construct a concrete instance of `CommunityArrayInput` via:
//
//	CommunityArray{ CommunityArgs{...} }
type CommunityArrayInput interface {
	pulumi.Input

	ToCommunityArrayOutput() CommunityArrayOutput
	ToCommunityArrayOutputWithContext(context.Context) CommunityArrayOutput
}

type CommunityArray []CommunityInput

func (CommunityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Community)(nil)).Elem()
}

func (i CommunityArray) ToCommunityArrayOutput() CommunityArrayOutput {
	return i.ToCommunityArrayOutputWithContext(context.Background())
}

func (i CommunityArray) ToCommunityArrayOutputWithContext(ctx context.Context) CommunityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommunityArrayOutput)
}

// CommunityMapInput is an input type that accepts CommunityMap and CommunityMapOutput values.
// You can construct a concrete instance of `CommunityMapInput` via:
//
//	CommunityMap{ "key": CommunityArgs{...} }
type CommunityMapInput interface {
	pulumi.Input

	ToCommunityMapOutput() CommunityMapOutput
	ToCommunityMapOutputWithContext(context.Context) CommunityMapOutput
}

type CommunityMap map[string]CommunityInput

func (CommunityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Community)(nil)).Elem()
}

func (i CommunityMap) ToCommunityMapOutput() CommunityMapOutput {
	return i.ToCommunityMapOutputWithContext(context.Background())
}

func (i CommunityMap) ToCommunityMapOutputWithContext(ctx context.Context) CommunityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommunityMapOutput)
}

type CommunityOutput struct{ *pulumi.OutputState }

func (CommunityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Community)(nil)).Elem()
}

func (o CommunityOutput) ToCommunityOutput() CommunityOutput {
	return o
}

func (o CommunityOutput) ToCommunityOutputWithContext(ctx context.Context) CommunityOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o CommunityOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Community) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// SNMP trap events.
func (o CommunityOutput) Events() pulumi.StringOutput {
	return o.ApplyT(func(v *Community) pulumi.StringOutput { return v.Events }).(pulumi.StringOutput)
}

// Community ID.
func (o CommunityOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Community) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o CommunityOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Community) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
func (o CommunityOutput) Hosts() CommunityHostArrayOutput {
	return o.ApplyT(func(v *Community) CommunityHostArrayOutput { return v.Hosts }).(CommunityHostArrayOutput)
}

// Configure IPv6 SNMP managers. The structure of `hosts6` block is documented below.
func (o CommunityOutput) Hosts6s() CommunityHosts6ArrayOutput {
	return o.ApplyT(func(v *Community) CommunityHosts6ArrayOutput { return v.Hosts6s }).(CommunityHosts6ArrayOutput)
}

// SNMP access control MIB view.
func (o CommunityOutput) MibView() pulumi.StringOutput {
	return o.ApplyT(func(v *Community) pulumi.StringOutput { return v.MibView }).(pulumi.StringOutput)
}

// Community name.
func (o CommunityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Community) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// SNMP v1 query port (default = 161).
func (o CommunityOutput) QueryV1Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Community) pulumi.IntOutput { return v.QueryV1Port }).(pulumi.IntOutput)
}

// Enable/disable SNMP v1 queries. Valid values: `enable`, `disable`.
func (o CommunityOutput) QueryV1Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Community) pulumi.StringOutput { return v.QueryV1Status }).(pulumi.StringOutput)
}

// SNMP v2c query port (default = 161).
func (o CommunityOutput) QueryV2cPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Community) pulumi.IntOutput { return v.QueryV2cPort }).(pulumi.IntOutput)
}

// Enable/disable SNMP v2c queries. Valid values: `enable`, `disable`.
func (o CommunityOutput) QueryV2cStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Community) pulumi.StringOutput { return v.QueryV2cStatus }).(pulumi.StringOutput)
}

// Enable/disable this SNMP community. Valid values: `enable`, `disable`.
func (o CommunityOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Community) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// SNMP v1 trap local port (default = 162).
func (o CommunityOutput) TrapV1Lport() pulumi.IntOutput {
	return o.ApplyT(func(v *Community) pulumi.IntOutput { return v.TrapV1Lport }).(pulumi.IntOutput)
}

// SNMP v1 trap remote port (default = 162).
func (o CommunityOutput) TrapV1Rport() pulumi.IntOutput {
	return o.ApplyT(func(v *Community) pulumi.IntOutput { return v.TrapV1Rport }).(pulumi.IntOutput)
}

// Enable/disable SNMP v1 traps. Valid values: `enable`, `disable`.
func (o CommunityOutput) TrapV1Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Community) pulumi.StringOutput { return v.TrapV1Status }).(pulumi.StringOutput)
}

// SNMP v2c trap local port (default = 162).
func (o CommunityOutput) TrapV2cLport() pulumi.IntOutput {
	return o.ApplyT(func(v *Community) pulumi.IntOutput { return v.TrapV2cLport }).(pulumi.IntOutput)
}

// SNMP v2c trap remote port (default = 162).
func (o CommunityOutput) TrapV2cRport() pulumi.IntOutput {
	return o.ApplyT(func(v *Community) pulumi.IntOutput { return v.TrapV2cRport }).(pulumi.IntOutput)
}

// Enable/disable SNMP v2c traps. Valid values: `enable`, `disable`.
func (o CommunityOutput) TrapV2cStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Community) pulumi.StringOutput { return v.TrapV2cStatus }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o CommunityOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Community) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// SNMP access control VDOMs. The structure of `vdoms` block is documented below.
func (o CommunityOutput) Vdoms() CommunityVdomArrayOutput {
	return o.ApplyT(func(v *Community) CommunityVdomArrayOutput { return v.Vdoms }).(CommunityVdomArrayOutput)
}

type CommunityArrayOutput struct{ *pulumi.OutputState }

func (CommunityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Community)(nil)).Elem()
}

func (o CommunityArrayOutput) ToCommunityArrayOutput() CommunityArrayOutput {
	return o
}

func (o CommunityArrayOutput) ToCommunityArrayOutputWithContext(ctx context.Context) CommunityArrayOutput {
	return o
}

func (o CommunityArrayOutput) Index(i pulumi.IntInput) CommunityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Community {
		return vs[0].([]*Community)[vs[1].(int)]
	}).(CommunityOutput)
}

type CommunityMapOutput struct{ *pulumi.OutputState }

func (CommunityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Community)(nil)).Elem()
}

func (o CommunityMapOutput) ToCommunityMapOutput() CommunityMapOutput {
	return o
}

func (o CommunityMapOutput) ToCommunityMapOutputWithContext(ctx context.Context) CommunityMapOutput {
	return o
}

func (o CommunityMapOutput) MapIndex(k pulumi.StringInput) CommunityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Community {
		return vs[0].(map[string]*Community)[vs[1].(string)]
	}).(CommunityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CommunityInput)(nil)).Elem(), &Community{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommunityArrayInput)(nil)).Elem(), CommunityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommunityMapInput)(nil)).Elem(), CommunityMap{})
	pulumi.RegisterOutputType(CommunityOutput{})
	pulumi.RegisterOutputType(CommunityArrayOutput{})
	pulumi.RegisterOutputType(CommunityMapOutput{})
}
