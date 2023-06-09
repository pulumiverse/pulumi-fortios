// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SNMP user configuration.
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
//			_, err := fortios.NewSystemsnmpUser(ctx, "trname", &fortios.SystemsnmpUserArgs{
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
//
// ## Import
//
// # SystemSnmp User can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemsnmpUser:SystemsnmpUser labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemsnmpUser:SystemsnmpUser labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemsnmpUser struct {
	pulumi.CustomResourceState

	// Authentication protocol.
	AuthProto pulumi.StringOutput `pulumi:"authProto"`
	// Password for authentication protocol.
	AuthPwd pulumi.StringPtrOutput `pulumi:"authPwd"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// SNMP notifications (traps) to send.
	Events pulumi.StringOutput `pulumi:"events"`
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
	Vdoms SystemsnmpUserVdomArrayOutput `pulumi:"vdoms"`
}

// NewSystemsnmpUser registers a new resource with the given unique name, arguments, and options.
func NewSystemsnmpUser(ctx *pulumi.Context,
	name string, args *SystemsnmpUserArgs, opts ...pulumi.ResourceOption) (*SystemsnmpUser, error) {
	if args == nil {
		args = &SystemsnmpUserArgs{}
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
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemsnmpUser
	err := ctx.RegisterResource("fortios:index/systemsnmpUser:SystemsnmpUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemsnmpUser gets an existing SystemsnmpUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemsnmpUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemsnmpUserState, opts ...pulumi.ResourceOption) (*SystemsnmpUser, error) {
	var resource SystemsnmpUser
	err := ctx.ReadResource("fortios:index/systemsnmpUser:SystemsnmpUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemsnmpUser resources.
type systemsnmpUserState struct {
	// Authentication protocol.
	AuthProto *string `pulumi:"authProto"`
	// Password for authentication protocol.
	AuthPwd *string `pulumi:"authPwd"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// SNMP notifications (traps) to send.
	Events *string `pulumi:"events"`
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
	Vdoms []SystemsnmpUserVdom `pulumi:"vdoms"`
}

type SystemsnmpUserState struct {
	// Authentication protocol.
	AuthProto pulumi.StringPtrInput
	// Password for authentication protocol.
	AuthPwd pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// SNMP notifications (traps) to send.
	Events pulumi.StringPtrInput
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
	Vdoms SystemsnmpUserVdomArrayInput
}

func (SystemsnmpUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemsnmpUserState)(nil)).Elem()
}

type systemsnmpUserArgs struct {
	// Authentication protocol.
	AuthProto *string `pulumi:"authProto"`
	// Password for authentication protocol.
	AuthPwd *string `pulumi:"authPwd"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// SNMP notifications (traps) to send.
	Events *string `pulumi:"events"`
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
	Vdoms []SystemsnmpUserVdom `pulumi:"vdoms"`
}

// The set of arguments for constructing a SystemsnmpUser resource.
type SystemsnmpUserArgs struct {
	// Authentication protocol.
	AuthProto pulumi.StringPtrInput
	// Password for authentication protocol.
	AuthPwd pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// SNMP notifications (traps) to send.
	Events pulumi.StringPtrInput
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
	Vdoms SystemsnmpUserVdomArrayInput
}

func (SystemsnmpUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemsnmpUserArgs)(nil)).Elem()
}

type SystemsnmpUserInput interface {
	pulumi.Input

	ToSystemsnmpUserOutput() SystemsnmpUserOutput
	ToSystemsnmpUserOutputWithContext(ctx context.Context) SystemsnmpUserOutput
}

func (*SystemsnmpUser) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemsnmpUser)(nil)).Elem()
}

func (i *SystemsnmpUser) ToSystemsnmpUserOutput() SystemsnmpUserOutput {
	return i.ToSystemsnmpUserOutputWithContext(context.Background())
}

func (i *SystemsnmpUser) ToSystemsnmpUserOutputWithContext(ctx context.Context) SystemsnmpUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemsnmpUserOutput)
}

// SystemsnmpUserArrayInput is an input type that accepts SystemsnmpUserArray and SystemsnmpUserArrayOutput values.
// You can construct a concrete instance of `SystemsnmpUserArrayInput` via:
//
//	SystemsnmpUserArray{ SystemsnmpUserArgs{...} }
type SystemsnmpUserArrayInput interface {
	pulumi.Input

	ToSystemsnmpUserArrayOutput() SystemsnmpUserArrayOutput
	ToSystemsnmpUserArrayOutputWithContext(context.Context) SystemsnmpUserArrayOutput
}

type SystemsnmpUserArray []SystemsnmpUserInput

func (SystemsnmpUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemsnmpUser)(nil)).Elem()
}

func (i SystemsnmpUserArray) ToSystemsnmpUserArrayOutput() SystemsnmpUserArrayOutput {
	return i.ToSystemsnmpUserArrayOutputWithContext(context.Background())
}

func (i SystemsnmpUserArray) ToSystemsnmpUserArrayOutputWithContext(ctx context.Context) SystemsnmpUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemsnmpUserArrayOutput)
}

// SystemsnmpUserMapInput is an input type that accepts SystemsnmpUserMap and SystemsnmpUserMapOutput values.
// You can construct a concrete instance of `SystemsnmpUserMapInput` via:
//
//	SystemsnmpUserMap{ "key": SystemsnmpUserArgs{...} }
type SystemsnmpUserMapInput interface {
	pulumi.Input

	ToSystemsnmpUserMapOutput() SystemsnmpUserMapOutput
	ToSystemsnmpUserMapOutputWithContext(context.Context) SystemsnmpUserMapOutput
}

type SystemsnmpUserMap map[string]SystemsnmpUserInput

func (SystemsnmpUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemsnmpUser)(nil)).Elem()
}

func (i SystemsnmpUserMap) ToSystemsnmpUserMapOutput() SystemsnmpUserMapOutput {
	return i.ToSystemsnmpUserMapOutputWithContext(context.Background())
}

func (i SystemsnmpUserMap) ToSystemsnmpUserMapOutputWithContext(ctx context.Context) SystemsnmpUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemsnmpUserMapOutput)
}

type SystemsnmpUserOutput struct{ *pulumi.OutputState }

func (SystemsnmpUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemsnmpUser)(nil)).Elem()
}

func (o SystemsnmpUserOutput) ToSystemsnmpUserOutput() SystemsnmpUserOutput {
	return o
}

func (o SystemsnmpUserOutput) ToSystemsnmpUserOutputWithContext(ctx context.Context) SystemsnmpUserOutput {
	return o
}

// Authentication protocol.
func (o SystemsnmpUserOutput) AuthProto() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.StringOutput { return v.AuthProto }).(pulumi.StringOutput)
}

// Password for authentication protocol.
func (o SystemsnmpUserOutput) AuthPwd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.StringPtrOutput { return v.AuthPwd }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SystemsnmpUserOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// SNMP notifications (traps) to send.
func (o SystemsnmpUserOutput) Events() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.StringOutput { return v.Events }).(pulumi.StringOutput)
}

// Enable/disable direct management of HA cluster members. Valid values: `enable`, `disable`.
func (o SystemsnmpUserOutput) HaDirect() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.StringOutput { return v.HaDirect }).(pulumi.StringOutput)
}

// SNMP access control MIB view.
func (o SystemsnmpUserOutput) MibView() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.StringOutput { return v.MibView }).(pulumi.StringOutput)
}

// SNMP user name.
func (o SystemsnmpUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// SNMP managers to send notifications (traps) to.
func (o SystemsnmpUserOutput) NotifyHosts() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.StringOutput { return v.NotifyHosts }).(pulumi.StringOutput)
}

// IPv6 SNMP managers to send notifications (traps) to.
func (o SystemsnmpUserOutput) NotifyHosts6() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.StringOutput { return v.NotifyHosts6 }).(pulumi.StringOutput)
}

// Privacy (encryption) protocol. Valid values: `aes`, `des`, `aes256`, `aes256cisco`.
func (o SystemsnmpUserOutput) PrivProto() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.StringOutput { return v.PrivProto }).(pulumi.StringOutput)
}

// Password for privacy (encryption) protocol.
func (o SystemsnmpUserOutput) PrivPwd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.StringPtrOutput { return v.PrivPwd }).(pulumi.StringPtrOutput)
}

// Enable/disable SNMP queries for this user. Valid values: `enable`, `disable`.
func (o SystemsnmpUserOutput) Queries() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.StringOutput { return v.Queries }).(pulumi.StringOutput)
}

// SNMPv3 query port (default = 161).
func (o SystemsnmpUserOutput) QueryPort() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.IntOutput { return v.QueryPort }).(pulumi.IntOutput)
}

// Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.
func (o SystemsnmpUserOutput) SecurityLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.StringOutput { return v.SecurityLevel }).(pulumi.StringOutput)
}

// Source IP for SNMP trap.
func (o SystemsnmpUserOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// Source IPv6 for SNMP trap.
func (o SystemsnmpUserOutput) SourceIpv6() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.StringOutput { return v.SourceIpv6 }).(pulumi.StringOutput)
}

// Enable/disable this SNMP user. Valid values: `enable`, `disable`.
func (o SystemsnmpUserOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// SNMPv3 local trap port (default = 162).
func (o SystemsnmpUserOutput) TrapLport() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.IntOutput { return v.TrapLport }).(pulumi.IntOutput)
}

// SNMPv3 trap remote port (default = 162).
func (o SystemsnmpUserOutput) TrapRport() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.IntOutput { return v.TrapRport }).(pulumi.IntOutput)
}

// Enable/disable traps for this SNMP user. Valid values: `enable`, `disable`.
func (o SystemsnmpUserOutput) TrapStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.StringOutput { return v.TrapStatus }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemsnmpUserOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemsnmpUser) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// SNMP access control VDOMs. The structure of `vdoms` block is documented below.
func (o SystemsnmpUserOutput) Vdoms() SystemsnmpUserVdomArrayOutput {
	return o.ApplyT(func(v *SystemsnmpUser) SystemsnmpUserVdomArrayOutput { return v.Vdoms }).(SystemsnmpUserVdomArrayOutput)
}

type SystemsnmpUserArrayOutput struct{ *pulumi.OutputState }

func (SystemsnmpUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemsnmpUser)(nil)).Elem()
}

func (o SystemsnmpUserArrayOutput) ToSystemsnmpUserArrayOutput() SystemsnmpUserArrayOutput {
	return o
}

func (o SystemsnmpUserArrayOutput) ToSystemsnmpUserArrayOutputWithContext(ctx context.Context) SystemsnmpUserArrayOutput {
	return o
}

func (o SystemsnmpUserArrayOutput) Index(i pulumi.IntInput) SystemsnmpUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemsnmpUser {
		return vs[0].([]*SystemsnmpUser)[vs[1].(int)]
	}).(SystemsnmpUserOutput)
}

type SystemsnmpUserMapOutput struct{ *pulumi.OutputState }

func (SystemsnmpUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemsnmpUser)(nil)).Elem()
}

func (o SystemsnmpUserMapOutput) ToSystemsnmpUserMapOutput() SystemsnmpUserMapOutput {
	return o
}

func (o SystemsnmpUserMapOutput) ToSystemsnmpUserMapOutputWithContext(ctx context.Context) SystemsnmpUserMapOutput {
	return o
}

func (o SystemsnmpUserMapOutput) MapIndex(k pulumi.StringInput) SystemsnmpUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemsnmpUser {
		return vs[0].(map[string]*SystemsnmpUser)[vs[1].(string)]
	}).(SystemsnmpUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemsnmpUserInput)(nil)).Elem(), &SystemsnmpUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemsnmpUserArrayInput)(nil)).Elem(), SystemsnmpUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemsnmpUserMapInput)(nil)).Elem(), SystemsnmpUserMap{})
	pulumi.RegisterOutputType(SystemsnmpUserOutput{})
	pulumi.RegisterOutputType(SystemsnmpUserArrayOutput{})
	pulumi.RegisterOutputType(SystemsnmpUserMapOutput{})
}
