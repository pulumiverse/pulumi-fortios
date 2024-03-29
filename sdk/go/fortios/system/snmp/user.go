// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package snmp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// SNMP user configuration.
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
//			_, err := system.NewUser(ctx, "trname", &system.UserArgs{
//				AuthProto:     pulumi.String("sha"),
//				Events:        pulumi.String("cpu-high mem-low log-full intf-ip vpn-tun-up vpn-tun-down ha-switch ha-hb-failure ips-signature ips-anomaly av-virus av-oversize av-pattern av-fragmented fm-if-change bgp-established bgp-backward-transition ha-member-up ha-member-down ent-conf-change av-conserve av-bypass av-oversize-passed av-oversize-blocked ips-pkg-update ips-fail-open faz-disconnect wc-ap-up wc-ap-down fswctl-session-up fswctl-session-down load-balance-real-server-down per-cpu-high"),
//				HaDirect:      pulumi.String("disable"),
//				PrivProto:     pulumi.String("aes"),
//				Queries:       pulumi.String("disable"),
//				QueryPort:     pulumi.Int(161),
//				SecurityLevel: pulumi.String("no-auth-no-priv"),
//				SourceIp:      pulumi.String("0.0.0.0"),
//				SourceIpv6:    pulumi.String("::"),
//				Status:        pulumi.String("disable"),
//				TrapLport:     pulumi.Int(162),
//				TrapRport:     pulumi.Int(162),
//				TrapStatus:    pulumi.String("enable"),
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
// SystemSnmp User can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/snmp/user:User labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/snmp/user:User labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type User struct {
	pulumi.CustomResourceState

	// Authentication protocol.
	AuthProto pulumi.StringOutput `pulumi:"authProto"`
	// Password for authentication protocol.
	AuthPwd pulumi.StringPtrOutput `pulumi:"authPwd"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// SNMP notifications (traps) to send.
	Events pulumi.StringOutput `pulumi:"events"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Enable/disable direct management of HA cluster members. Valid values: `enable`, `disable`.
	HaDirect pulumi.StringOutput `pulumi:"haDirect"`
	// SNMP access control MIB view.
	MibView pulumi.StringOutput `pulumi:"mibView"`
	// SNMP user name.
	Name pulumi.StringOutput `pulumi:"name"`
	// SNMP managers to send notifications (traps) to.
	NotifyHosts pulumi.StringOutput `pulumi:"notifyHosts"`
	// IPv6 SNMP managers to send notifications (traps) to.
	NotifyHosts6 pulumi.StringOutput `pulumi:"notifyHosts6"`
	// Privacy (encryption) protocol. Valid values: `aes`, `des`, `aes256`, `aes256cisco`.
	PrivProto pulumi.StringOutput `pulumi:"privProto"`
	// Password for privacy (encryption) protocol.
	PrivPwd pulumi.StringPtrOutput `pulumi:"privPwd"`
	// Enable/disable SNMP queries for this user. Valid values: `enable`, `disable`.
	Queries pulumi.StringOutput `pulumi:"queries"`
	// SNMPv3 query port (default = 161).
	QueryPort pulumi.IntOutput `pulumi:"queryPort"`
	// Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.
	SecurityLevel pulumi.StringOutput `pulumi:"securityLevel"`
	// Source IP for SNMP trap.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Source IPv6 for SNMP trap.
	SourceIpv6 pulumi.StringOutput `pulumi:"sourceIpv6"`
	// Enable/disable this SNMP user. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// SNMPv3 local trap port (default = 162).
	TrapLport pulumi.IntOutput `pulumi:"trapLport"`
	// SNMPv3 trap remote port (default = 162).
	TrapRport pulumi.IntOutput `pulumi:"trapRport"`
	// Enable/disable traps for this SNMP user. Valid values: `enable`, `disable`.
	TrapStatus pulumi.StringOutput `pulumi:"trapStatus"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// SNMP access control VDOMs. The structure of `vdoms` block is documented below.
	Vdoms UserVdomArrayOutput `pulumi:"vdoms"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		args = &UserArgs{}
	}

	if args.AuthPwd != nil {
		args.AuthPwd = pulumi.ToSecret(args.AuthPwd).(pulumi.StringPtrInput)
	}
	if args.PrivPwd != nil {
		args.PrivPwd = pulumi.ToSecret(args.PrivPwd).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"authPwd",
		"privPwd",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource User
	err := ctx.RegisterResource("fortios:system/snmp/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("fortios:system/snmp/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// Authentication protocol.
	AuthProto *string `pulumi:"authProto"`
	// Password for authentication protocol.
	AuthPwd *string `pulumi:"authPwd"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// SNMP notifications (traps) to send.
	Events *string `pulumi:"events"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Enable/disable direct management of HA cluster members. Valid values: `enable`, `disable`.
	HaDirect *string `pulumi:"haDirect"`
	// SNMP access control MIB view.
	MibView *string `pulumi:"mibView"`
	// SNMP user name.
	Name *string `pulumi:"name"`
	// SNMP managers to send notifications (traps) to.
	NotifyHosts *string `pulumi:"notifyHosts"`
	// IPv6 SNMP managers to send notifications (traps) to.
	NotifyHosts6 *string `pulumi:"notifyHosts6"`
	// Privacy (encryption) protocol. Valid values: `aes`, `des`, `aes256`, `aes256cisco`.
	PrivProto *string `pulumi:"privProto"`
	// Password for privacy (encryption) protocol.
	PrivPwd *string `pulumi:"privPwd"`
	// Enable/disable SNMP queries for this user. Valid values: `enable`, `disable`.
	Queries *string `pulumi:"queries"`
	// SNMPv3 query port (default = 161).
	QueryPort *int `pulumi:"queryPort"`
	// Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.
	SecurityLevel *string `pulumi:"securityLevel"`
	// Source IP for SNMP trap.
	SourceIp *string `pulumi:"sourceIp"`
	// Source IPv6 for SNMP trap.
	SourceIpv6 *string `pulumi:"sourceIpv6"`
	// Enable/disable this SNMP user. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// SNMPv3 local trap port (default = 162).
	TrapLport *int `pulumi:"trapLport"`
	// SNMPv3 trap remote port (default = 162).
	TrapRport *int `pulumi:"trapRport"`
	// Enable/disable traps for this SNMP user. Valid values: `enable`, `disable`.
	TrapStatus *string `pulumi:"trapStatus"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// SNMP access control VDOMs. The structure of `vdoms` block is documented below.
	Vdoms []UserVdom `pulumi:"vdoms"`
}

type UserState struct {
	// Authentication protocol.
	AuthProto pulumi.StringPtrInput
	// Password for authentication protocol.
	AuthPwd pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// SNMP notifications (traps) to send.
	Events pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Enable/disable direct management of HA cluster members. Valid values: `enable`, `disable`.
	HaDirect pulumi.StringPtrInput
	// SNMP access control MIB view.
	MibView pulumi.StringPtrInput
	// SNMP user name.
	Name pulumi.StringPtrInput
	// SNMP managers to send notifications (traps) to.
	NotifyHosts pulumi.StringPtrInput
	// IPv6 SNMP managers to send notifications (traps) to.
	NotifyHosts6 pulumi.StringPtrInput
	// Privacy (encryption) protocol. Valid values: `aes`, `des`, `aes256`, `aes256cisco`.
	PrivProto pulumi.StringPtrInput
	// Password for privacy (encryption) protocol.
	PrivPwd pulumi.StringPtrInput
	// Enable/disable SNMP queries for this user. Valid values: `enable`, `disable`.
	Queries pulumi.StringPtrInput
	// SNMPv3 query port (default = 161).
	QueryPort pulumi.IntPtrInput
	// Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.
	SecurityLevel pulumi.StringPtrInput
	// Source IP for SNMP trap.
	SourceIp pulumi.StringPtrInput
	// Source IPv6 for SNMP trap.
	SourceIpv6 pulumi.StringPtrInput
	// Enable/disable this SNMP user. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// SNMPv3 local trap port (default = 162).
	TrapLport pulumi.IntPtrInput
	// SNMPv3 trap remote port (default = 162).
	TrapRport pulumi.IntPtrInput
	// Enable/disable traps for this SNMP user. Valid values: `enable`, `disable`.
	TrapStatus pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// SNMP access control VDOMs. The structure of `vdoms` block is documented below.
	Vdoms UserVdomArrayInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// Authentication protocol.
	AuthProto *string `pulumi:"authProto"`
	// Password for authentication protocol.
	AuthPwd *string `pulumi:"authPwd"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// SNMP notifications (traps) to send.
	Events *string `pulumi:"events"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Enable/disable direct management of HA cluster members. Valid values: `enable`, `disable`.
	HaDirect *string `pulumi:"haDirect"`
	// SNMP access control MIB view.
	MibView *string `pulumi:"mibView"`
	// SNMP user name.
	Name *string `pulumi:"name"`
	// SNMP managers to send notifications (traps) to.
	NotifyHosts *string `pulumi:"notifyHosts"`
	// IPv6 SNMP managers to send notifications (traps) to.
	NotifyHosts6 *string `pulumi:"notifyHosts6"`
	// Privacy (encryption) protocol. Valid values: `aes`, `des`, `aes256`, `aes256cisco`.
	PrivProto *string `pulumi:"privProto"`
	// Password for privacy (encryption) protocol.
	PrivPwd *string `pulumi:"privPwd"`
	// Enable/disable SNMP queries for this user. Valid values: `enable`, `disable`.
	Queries *string `pulumi:"queries"`
	// SNMPv3 query port (default = 161).
	QueryPort *int `pulumi:"queryPort"`
	// Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.
	SecurityLevel *string `pulumi:"securityLevel"`
	// Source IP for SNMP trap.
	SourceIp *string `pulumi:"sourceIp"`
	// Source IPv6 for SNMP trap.
	SourceIpv6 *string `pulumi:"sourceIpv6"`
	// Enable/disable this SNMP user. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// SNMPv3 local trap port (default = 162).
	TrapLport *int `pulumi:"trapLport"`
	// SNMPv3 trap remote port (default = 162).
	TrapRport *int `pulumi:"trapRport"`
	// Enable/disable traps for this SNMP user. Valid values: `enable`, `disable`.
	TrapStatus *string `pulumi:"trapStatus"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// SNMP access control VDOMs. The structure of `vdoms` block is documented below.
	Vdoms []UserVdom `pulumi:"vdoms"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// Authentication protocol.
	AuthProto pulumi.StringPtrInput
	// Password for authentication protocol.
	AuthPwd pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// SNMP notifications (traps) to send.
	Events pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Enable/disable direct management of HA cluster members. Valid values: `enable`, `disable`.
	HaDirect pulumi.StringPtrInput
	// SNMP access control MIB view.
	MibView pulumi.StringPtrInput
	// SNMP user name.
	Name pulumi.StringPtrInput
	// SNMP managers to send notifications (traps) to.
	NotifyHosts pulumi.StringPtrInput
	// IPv6 SNMP managers to send notifications (traps) to.
	NotifyHosts6 pulumi.StringPtrInput
	// Privacy (encryption) protocol. Valid values: `aes`, `des`, `aes256`, `aes256cisco`.
	PrivProto pulumi.StringPtrInput
	// Password for privacy (encryption) protocol.
	PrivPwd pulumi.StringPtrInput
	// Enable/disable SNMP queries for this user. Valid values: `enable`, `disable`.
	Queries pulumi.StringPtrInput
	// SNMPv3 query port (default = 161).
	QueryPort pulumi.IntPtrInput
	// Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.
	SecurityLevel pulumi.StringPtrInput
	// Source IP for SNMP trap.
	SourceIp pulumi.StringPtrInput
	// Source IPv6 for SNMP trap.
	SourceIpv6 pulumi.StringPtrInput
	// Enable/disable this SNMP user. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// SNMPv3 local trap port (default = 162).
	TrapLport pulumi.IntPtrInput
	// SNMPv3 trap remote port (default = 162).
	TrapRport pulumi.IntPtrInput
	// Enable/disable traps for this SNMP user. Valid values: `enable`, `disable`.
	TrapStatus pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// SNMP access control VDOMs. The structure of `vdoms` block is documented below.
	Vdoms UserVdomArrayInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//	UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (i UserArray) ToUserArrayOutput() UserArrayOutput {
	return i.ToUserArrayOutputWithContext(context.Background())
}

func (i UserArray) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArrayOutput)
}

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//	UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

// Authentication protocol.
func (o UserOutput) AuthProto() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.AuthProto }).(pulumi.StringOutput)
}

// Password for authentication protocol.
func (o UserOutput) AuthPwd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.AuthPwd }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o UserOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// SNMP notifications (traps) to send.
func (o UserOutput) Events() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Events }).(pulumi.StringOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o UserOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Enable/disable direct management of HA cluster members. Valid values: `enable`, `disable`.
func (o UserOutput) HaDirect() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.HaDirect }).(pulumi.StringOutput)
}

// SNMP access control MIB view.
func (o UserOutput) MibView() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.MibView }).(pulumi.StringOutput)
}

// SNMP user name.
func (o UserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// SNMP managers to send notifications (traps) to.
func (o UserOutput) NotifyHosts() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.NotifyHosts }).(pulumi.StringOutput)
}

// IPv6 SNMP managers to send notifications (traps) to.
func (o UserOutput) NotifyHosts6() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.NotifyHosts6 }).(pulumi.StringOutput)
}

// Privacy (encryption) protocol. Valid values: `aes`, `des`, `aes256`, `aes256cisco`.
func (o UserOutput) PrivProto() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.PrivProto }).(pulumi.StringOutput)
}

// Password for privacy (encryption) protocol.
func (o UserOutput) PrivPwd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.PrivPwd }).(pulumi.StringPtrOutput)
}

// Enable/disable SNMP queries for this user. Valid values: `enable`, `disable`.
func (o UserOutput) Queries() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Queries }).(pulumi.StringOutput)
}

// SNMPv3 query port (default = 161).
func (o UserOutput) QueryPort() pulumi.IntOutput {
	return o.ApplyT(func(v *User) pulumi.IntOutput { return v.QueryPort }).(pulumi.IntOutput)
}

// Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.
func (o UserOutput) SecurityLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.SecurityLevel }).(pulumi.StringOutput)
}

// Source IP for SNMP trap.
func (o UserOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// Source IPv6 for SNMP trap.
func (o UserOutput) SourceIpv6() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.SourceIpv6 }).(pulumi.StringOutput)
}

// Enable/disable this SNMP user. Valid values: `enable`, `disable`.
func (o UserOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// SNMPv3 local trap port (default = 162).
func (o UserOutput) TrapLport() pulumi.IntOutput {
	return o.ApplyT(func(v *User) pulumi.IntOutput { return v.TrapLport }).(pulumi.IntOutput)
}

// SNMPv3 trap remote port (default = 162).
func (o UserOutput) TrapRport() pulumi.IntOutput {
	return o.ApplyT(func(v *User) pulumi.IntOutput { return v.TrapRport }).(pulumi.IntOutput)
}

// Enable/disable traps for this SNMP user. Valid values: `enable`, `disable`.
func (o UserOutput) TrapStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.TrapStatus }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o UserOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// SNMP access control VDOMs. The structure of `vdoms` block is documented below.
func (o UserOutput) Vdoms() UserVdomArrayOutput {
	return o.ApplyT(func(v *User) UserVdomArrayOutput { return v.Vdoms }).(UserVdomArrayOutput)
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *User {
		return vs[0].([]*User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *User {
		return vs[0].(map[string]*User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserArrayInput)(nil)).Elem(), UserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMapInput)(nil)).Elem(), UserMap{})
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}
