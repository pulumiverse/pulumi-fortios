// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securitypolicy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure 802.1x MAC Authentication Bypass (MAB) policies.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/switchcontroller"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := switchcontroller.NewPolicy8021X(ctx, "trname", &switchcontroller.Policy8021XArgs{
//				AuthFailVlan:           pulumi.String("disable"),
//				AuthFailVlanid:         pulumi.Int(0),
//				EapPassthru:            pulumi.String("disable"),
//				FramevidApply:          pulumi.String("enable"),
//				GuestAuthDelay:         pulumi.Int(30),
//				GuestVlan:              pulumi.String("disable"),
//				GuestVlanid:            pulumi.Int(100),
//				MacAuthBypass:          pulumi.String("disable"),
//				OpenAuth:               pulumi.String("disable"),
//				PolicyType:             pulumi.String("802.1X"),
//				RadiusTimeoutOverwrite: pulumi.String("disable"),
//				SecurityMode:           pulumi.String("802.1X"),
//				UserGroups: securitypolicy.Policy8021XUserGroupArray{
//					&securitypolicy.Policy8021XUserGroupArgs{
//						Name: pulumi.String("Guest-group"),
//					},
//				},
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
// SwitchControllerSecurityPolicy 8021X can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/securitypolicy/policy8021X:Policy8021X labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/securitypolicy/policy8021X:Policy8021X labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Policy8021X struct {
	pulumi.CustomResourceState

	// Enable to allow limited access to clients that cannot authenticate. Valid values: `disable`, `enable`.
	AuthFailVlan pulumi.StringOutput `pulumi:"authFailVlan"`
	// VLAN ID on which authentication failed.
	AuthFailVlanId pulumi.StringOutput `pulumi:"authFailVlanId"`
	// VLAN ID on which authentication failed.
	AuthFailVlanid pulumi.IntOutput `pulumi:"authFailVlanid"`
	// Authentication server timeout period (3 - 15 sec, default = 3).
	AuthserverTimeoutPeriod pulumi.IntOutput `pulumi:"authserverTimeoutPeriod"`
	// Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable.  Valid values: `disable`, `enable`.
	AuthserverTimeoutVlan pulumi.StringOutput `pulumi:"authserverTimeoutVlan"`
	// Authentication server timeout VLAN name.
	AuthserverTimeoutVlanid pulumi.StringOutput `pulumi:"authserverTimeoutVlanid"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable`, `enable`.
	EapAutoUntaggedVlans pulumi.StringOutput `pulumi:"eapAutoUntaggedVlans"`
	// Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable`, `enable`.
	EapPassthru pulumi.StringOutput `pulumi:"eapPassthru"`
	// Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable`, `enable`.
	FramevidApply pulumi.StringOutput `pulumi:"framevidApply"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Guest authentication delay (1 - 900  sec, default = 30).
	GuestAuthDelay pulumi.IntOutput `pulumi:"guestAuthDelay"`
	// Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable`, `enable`.
	GuestVlan pulumi.StringOutput `pulumi:"guestVlan"`
	// Guest VLAN name.
	GuestVlanId pulumi.StringOutput `pulumi:"guestVlanId"`
	// Guest VLAN ID.
	GuestVlanid pulumi.IntOutput `pulumi:"guestVlanid"`
	// Enable/disable MAB for this policy. Valid values: `disable`, `enable`.
	MacAuthBypass pulumi.StringOutput `pulumi:"macAuthBypass"`
	// Policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable open authentication for this policy. Valid values: `disable`, `enable`.
	OpenAuth pulumi.StringOutput `pulumi:"openAuth"`
	// Policy type. Valid values: `802.1X`.
	PolicyType pulumi.StringOutput `pulumi:"policyType"`
	// Enable to override the global RADIUS session timeout. Valid values: `disable`, `enable`.
	RadiusTimeoutOverwrite pulumi.StringOutput `pulumi:"radiusTimeoutOverwrite"`
	// Port or MAC based 802.1X security mode. Valid values: `802.1X`, `802.1X-mac-based`.
	SecurityMode pulumi.StringOutput `pulumi:"securityMode"`
	// Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `userGroup` block is documented below.
	UserGroups Policy8021XUserGroupArrayOutput `pulumi:"userGroups"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewPolicy8021X registers a new resource with the given unique name, arguments, and options.
func NewPolicy8021X(ctx *pulumi.Context,
	name string, args *Policy8021XArgs, opts ...pulumi.ResourceOption) (*Policy8021X, error) {
	if args == nil {
		args = &Policy8021XArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Policy8021X
	err := ctx.RegisterResource("fortios:switchcontroller/securitypolicy/policy8021X:Policy8021X", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy8021X gets an existing Policy8021X resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy8021X(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Policy8021XState, opts ...pulumi.ResourceOption) (*Policy8021X, error) {
	var resource Policy8021X
	err := ctx.ReadResource("fortios:switchcontroller/securitypolicy/policy8021X:Policy8021X", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy8021X resources.
type policy8021XState struct {
	// Enable to allow limited access to clients that cannot authenticate. Valid values: `disable`, `enable`.
	AuthFailVlan *string `pulumi:"authFailVlan"`
	// VLAN ID on which authentication failed.
	AuthFailVlanId *string `pulumi:"authFailVlanId"`
	// VLAN ID on which authentication failed.
	AuthFailVlanid *int `pulumi:"authFailVlanid"`
	// Authentication server timeout period (3 - 15 sec, default = 3).
	AuthserverTimeoutPeriod *int `pulumi:"authserverTimeoutPeriod"`
	// Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable.  Valid values: `disable`, `enable`.
	AuthserverTimeoutVlan *string `pulumi:"authserverTimeoutVlan"`
	// Authentication server timeout VLAN name.
	AuthserverTimeoutVlanid *string `pulumi:"authserverTimeoutVlanid"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable`, `enable`.
	EapAutoUntaggedVlans *string `pulumi:"eapAutoUntaggedVlans"`
	// Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable`, `enable`.
	EapPassthru *string `pulumi:"eapPassthru"`
	// Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable`, `enable`.
	FramevidApply *string `pulumi:"framevidApply"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Guest authentication delay (1 - 900  sec, default = 30).
	GuestAuthDelay *int `pulumi:"guestAuthDelay"`
	// Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable`, `enable`.
	GuestVlan *string `pulumi:"guestVlan"`
	// Guest VLAN name.
	GuestVlanId *string `pulumi:"guestVlanId"`
	// Guest VLAN ID.
	GuestVlanid *int `pulumi:"guestVlanid"`
	// Enable/disable MAB for this policy. Valid values: `disable`, `enable`.
	MacAuthBypass *string `pulumi:"macAuthBypass"`
	// Policy name.
	Name *string `pulumi:"name"`
	// Enable/disable open authentication for this policy. Valid values: `disable`, `enable`.
	OpenAuth *string `pulumi:"openAuth"`
	// Policy type. Valid values: `802.1X`.
	PolicyType *string `pulumi:"policyType"`
	// Enable to override the global RADIUS session timeout. Valid values: `disable`, `enable`.
	RadiusTimeoutOverwrite *string `pulumi:"radiusTimeoutOverwrite"`
	// Port or MAC based 802.1X security mode. Valid values: `802.1X`, `802.1X-mac-based`.
	SecurityMode *string `pulumi:"securityMode"`
	// Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `userGroup` block is documented below.
	UserGroups []Policy8021XUserGroup `pulumi:"userGroups"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Policy8021XState struct {
	// Enable to allow limited access to clients that cannot authenticate. Valid values: `disable`, `enable`.
	AuthFailVlan pulumi.StringPtrInput
	// VLAN ID on which authentication failed.
	AuthFailVlanId pulumi.StringPtrInput
	// VLAN ID on which authentication failed.
	AuthFailVlanid pulumi.IntPtrInput
	// Authentication server timeout period (3 - 15 sec, default = 3).
	AuthserverTimeoutPeriod pulumi.IntPtrInput
	// Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable.  Valid values: `disable`, `enable`.
	AuthserverTimeoutVlan pulumi.StringPtrInput
	// Authentication server timeout VLAN name.
	AuthserverTimeoutVlanid pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable`, `enable`.
	EapAutoUntaggedVlans pulumi.StringPtrInput
	// Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable`, `enable`.
	EapPassthru pulumi.StringPtrInput
	// Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable`, `enable`.
	FramevidApply pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Guest authentication delay (1 - 900  sec, default = 30).
	GuestAuthDelay pulumi.IntPtrInput
	// Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable`, `enable`.
	GuestVlan pulumi.StringPtrInput
	// Guest VLAN name.
	GuestVlanId pulumi.StringPtrInput
	// Guest VLAN ID.
	GuestVlanid pulumi.IntPtrInput
	// Enable/disable MAB for this policy. Valid values: `disable`, `enable`.
	MacAuthBypass pulumi.StringPtrInput
	// Policy name.
	Name pulumi.StringPtrInput
	// Enable/disable open authentication for this policy. Valid values: `disable`, `enable`.
	OpenAuth pulumi.StringPtrInput
	// Policy type. Valid values: `802.1X`.
	PolicyType pulumi.StringPtrInput
	// Enable to override the global RADIUS session timeout. Valid values: `disable`, `enable`.
	RadiusTimeoutOverwrite pulumi.StringPtrInput
	// Port or MAC based 802.1X security mode. Valid values: `802.1X`, `802.1X-mac-based`.
	SecurityMode pulumi.StringPtrInput
	// Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `userGroup` block is documented below.
	UserGroups Policy8021XUserGroupArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Policy8021XState) ElementType() reflect.Type {
	return reflect.TypeOf((*policy8021XState)(nil)).Elem()
}

type policy8021XArgs struct {
	// Enable to allow limited access to clients that cannot authenticate. Valid values: `disable`, `enable`.
	AuthFailVlan *string `pulumi:"authFailVlan"`
	// VLAN ID on which authentication failed.
	AuthFailVlanId *string `pulumi:"authFailVlanId"`
	// VLAN ID on which authentication failed.
	AuthFailVlanid *int `pulumi:"authFailVlanid"`
	// Authentication server timeout period (3 - 15 sec, default = 3).
	AuthserverTimeoutPeriod *int `pulumi:"authserverTimeoutPeriod"`
	// Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable.  Valid values: `disable`, `enable`.
	AuthserverTimeoutVlan *string `pulumi:"authserverTimeoutVlan"`
	// Authentication server timeout VLAN name.
	AuthserverTimeoutVlanid *string `pulumi:"authserverTimeoutVlanid"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable`, `enable`.
	EapAutoUntaggedVlans *string `pulumi:"eapAutoUntaggedVlans"`
	// Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable`, `enable`.
	EapPassthru *string `pulumi:"eapPassthru"`
	// Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable`, `enable`.
	FramevidApply *string `pulumi:"framevidApply"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Guest authentication delay (1 - 900  sec, default = 30).
	GuestAuthDelay *int `pulumi:"guestAuthDelay"`
	// Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable`, `enable`.
	GuestVlan *string `pulumi:"guestVlan"`
	// Guest VLAN name.
	GuestVlanId *string `pulumi:"guestVlanId"`
	// Guest VLAN ID.
	GuestVlanid *int `pulumi:"guestVlanid"`
	// Enable/disable MAB for this policy. Valid values: `disable`, `enable`.
	MacAuthBypass *string `pulumi:"macAuthBypass"`
	// Policy name.
	Name *string `pulumi:"name"`
	// Enable/disable open authentication for this policy. Valid values: `disable`, `enable`.
	OpenAuth *string `pulumi:"openAuth"`
	// Policy type. Valid values: `802.1X`.
	PolicyType *string `pulumi:"policyType"`
	// Enable to override the global RADIUS session timeout. Valid values: `disable`, `enable`.
	RadiusTimeoutOverwrite *string `pulumi:"radiusTimeoutOverwrite"`
	// Port or MAC based 802.1X security mode. Valid values: `802.1X`, `802.1X-mac-based`.
	SecurityMode *string `pulumi:"securityMode"`
	// Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `userGroup` block is documented below.
	UserGroups []Policy8021XUserGroup `pulumi:"userGroups"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Policy8021X resource.
type Policy8021XArgs struct {
	// Enable to allow limited access to clients that cannot authenticate. Valid values: `disable`, `enable`.
	AuthFailVlan pulumi.StringPtrInput
	// VLAN ID on which authentication failed.
	AuthFailVlanId pulumi.StringPtrInput
	// VLAN ID on which authentication failed.
	AuthFailVlanid pulumi.IntPtrInput
	// Authentication server timeout period (3 - 15 sec, default = 3).
	AuthserverTimeoutPeriod pulumi.IntPtrInput
	// Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable.  Valid values: `disable`, `enable`.
	AuthserverTimeoutVlan pulumi.StringPtrInput
	// Authentication server timeout VLAN name.
	AuthserverTimeoutVlanid pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable`, `enable`.
	EapAutoUntaggedVlans pulumi.StringPtrInput
	// Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable`, `enable`.
	EapPassthru pulumi.StringPtrInput
	// Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable`, `enable`.
	FramevidApply pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Guest authentication delay (1 - 900  sec, default = 30).
	GuestAuthDelay pulumi.IntPtrInput
	// Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable`, `enable`.
	GuestVlan pulumi.StringPtrInput
	// Guest VLAN name.
	GuestVlanId pulumi.StringPtrInput
	// Guest VLAN ID.
	GuestVlanid pulumi.IntPtrInput
	// Enable/disable MAB for this policy. Valid values: `disable`, `enable`.
	MacAuthBypass pulumi.StringPtrInput
	// Policy name.
	Name pulumi.StringPtrInput
	// Enable/disable open authentication for this policy. Valid values: `disable`, `enable`.
	OpenAuth pulumi.StringPtrInput
	// Policy type. Valid values: `802.1X`.
	PolicyType pulumi.StringPtrInput
	// Enable to override the global RADIUS session timeout. Valid values: `disable`, `enable`.
	RadiusTimeoutOverwrite pulumi.StringPtrInput
	// Port or MAC based 802.1X security mode. Valid values: `802.1X`, `802.1X-mac-based`.
	SecurityMode pulumi.StringPtrInput
	// Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `userGroup` block is documented below.
	UserGroups Policy8021XUserGroupArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Policy8021XArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policy8021XArgs)(nil)).Elem()
}

type Policy8021XInput interface {
	pulumi.Input

	ToPolicy8021XOutput() Policy8021XOutput
	ToPolicy8021XOutputWithContext(ctx context.Context) Policy8021XOutput
}

func (*Policy8021X) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy8021X)(nil)).Elem()
}

func (i *Policy8021X) ToPolicy8021XOutput() Policy8021XOutput {
	return i.ToPolicy8021XOutputWithContext(context.Background())
}

func (i *Policy8021X) ToPolicy8021XOutputWithContext(ctx context.Context) Policy8021XOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Policy8021XOutput)
}

// Policy8021XArrayInput is an input type that accepts Policy8021XArray and Policy8021XArrayOutput values.
// You can construct a concrete instance of `Policy8021XArrayInput` via:
//
//	Policy8021XArray{ Policy8021XArgs{...} }
type Policy8021XArrayInput interface {
	pulumi.Input

	ToPolicy8021XArrayOutput() Policy8021XArrayOutput
	ToPolicy8021XArrayOutputWithContext(context.Context) Policy8021XArrayOutput
}

type Policy8021XArray []Policy8021XInput

func (Policy8021XArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy8021X)(nil)).Elem()
}

func (i Policy8021XArray) ToPolicy8021XArrayOutput() Policy8021XArrayOutput {
	return i.ToPolicy8021XArrayOutputWithContext(context.Background())
}

func (i Policy8021XArray) ToPolicy8021XArrayOutputWithContext(ctx context.Context) Policy8021XArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Policy8021XArrayOutput)
}

// Policy8021XMapInput is an input type that accepts Policy8021XMap and Policy8021XMapOutput values.
// You can construct a concrete instance of `Policy8021XMapInput` via:
//
//	Policy8021XMap{ "key": Policy8021XArgs{...} }
type Policy8021XMapInput interface {
	pulumi.Input

	ToPolicy8021XMapOutput() Policy8021XMapOutput
	ToPolicy8021XMapOutputWithContext(context.Context) Policy8021XMapOutput
}

type Policy8021XMap map[string]Policy8021XInput

func (Policy8021XMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy8021X)(nil)).Elem()
}

func (i Policy8021XMap) ToPolicy8021XMapOutput() Policy8021XMapOutput {
	return i.ToPolicy8021XMapOutputWithContext(context.Background())
}

func (i Policy8021XMap) ToPolicy8021XMapOutputWithContext(ctx context.Context) Policy8021XMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Policy8021XMapOutput)
}

type Policy8021XOutput struct{ *pulumi.OutputState }

func (Policy8021XOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy8021X)(nil)).Elem()
}

func (o Policy8021XOutput) ToPolicy8021XOutput() Policy8021XOutput {
	return o
}

func (o Policy8021XOutput) ToPolicy8021XOutputWithContext(ctx context.Context) Policy8021XOutput {
	return o
}

// Enable to allow limited access to clients that cannot authenticate. Valid values: `disable`, `enable`.
func (o Policy8021XOutput) AuthFailVlan() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.StringOutput { return v.AuthFailVlan }).(pulumi.StringOutput)
}

// VLAN ID on which authentication failed.
func (o Policy8021XOutput) AuthFailVlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.StringOutput { return v.AuthFailVlanId }).(pulumi.StringOutput)
}

// VLAN ID on which authentication failed.
func (o Policy8021XOutput) AuthFailVlanid() pulumi.IntOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.IntOutput { return v.AuthFailVlanid }).(pulumi.IntOutput)
}

// Authentication server timeout period (3 - 15 sec, default = 3).
func (o Policy8021XOutput) AuthserverTimeoutPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.IntOutput { return v.AuthserverTimeoutPeriod }).(pulumi.IntOutput)
}

// Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable.  Valid values: `disable`, `enable`.
func (o Policy8021XOutput) AuthserverTimeoutVlan() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.StringOutput { return v.AuthserverTimeoutVlan }).(pulumi.StringOutput)
}

// Authentication server timeout VLAN name.
func (o Policy8021XOutput) AuthserverTimeoutVlanid() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.StringOutput { return v.AuthserverTimeoutVlanid }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o Policy8021XOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable`, `enable`.
func (o Policy8021XOutput) EapAutoUntaggedVlans() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.StringOutput { return v.EapAutoUntaggedVlans }).(pulumi.StringOutput)
}

// Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable`, `enable`.
func (o Policy8021XOutput) EapPassthru() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.StringOutput { return v.EapPassthru }).(pulumi.StringOutput)
}

// Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable`, `enable`.
func (o Policy8021XOutput) FramevidApply() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.StringOutput { return v.FramevidApply }).(pulumi.StringOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o Policy8021XOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Guest authentication delay (1 - 900  sec, default = 30).
func (o Policy8021XOutput) GuestAuthDelay() pulumi.IntOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.IntOutput { return v.GuestAuthDelay }).(pulumi.IntOutput)
}

// Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable`, `enable`.
func (o Policy8021XOutput) GuestVlan() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.StringOutput { return v.GuestVlan }).(pulumi.StringOutput)
}

// Guest VLAN name.
func (o Policy8021XOutput) GuestVlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.StringOutput { return v.GuestVlanId }).(pulumi.StringOutput)
}

// Guest VLAN ID.
func (o Policy8021XOutput) GuestVlanid() pulumi.IntOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.IntOutput { return v.GuestVlanid }).(pulumi.IntOutput)
}

// Enable/disable MAB for this policy. Valid values: `disable`, `enable`.
func (o Policy8021XOutput) MacAuthBypass() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.StringOutput { return v.MacAuthBypass }).(pulumi.StringOutput)
}

// Policy name.
func (o Policy8021XOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable open authentication for this policy. Valid values: `disable`, `enable`.
func (o Policy8021XOutput) OpenAuth() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.StringOutput { return v.OpenAuth }).(pulumi.StringOutput)
}

// Policy type. Valid values: `802.1X`.
func (o Policy8021XOutput) PolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.StringOutput { return v.PolicyType }).(pulumi.StringOutput)
}

// Enable to override the global RADIUS session timeout. Valid values: `disable`, `enable`.
func (o Policy8021XOutput) RadiusTimeoutOverwrite() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.StringOutput { return v.RadiusTimeoutOverwrite }).(pulumi.StringOutput)
}

// Port or MAC based 802.1X security mode. Valid values: `802.1X`, `802.1X-mac-based`.
func (o Policy8021XOutput) SecurityMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.StringOutput { return v.SecurityMode }).(pulumi.StringOutput)
}

// Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `userGroup` block is documented below.
func (o Policy8021XOutput) UserGroups() Policy8021XUserGroupArrayOutput {
	return o.ApplyT(func(v *Policy8021X) Policy8021XUserGroupArrayOutput { return v.UserGroups }).(Policy8021XUserGroupArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Policy8021XOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy8021X) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type Policy8021XArrayOutput struct{ *pulumi.OutputState }

func (Policy8021XArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy8021X)(nil)).Elem()
}

func (o Policy8021XArrayOutput) ToPolicy8021XArrayOutput() Policy8021XArrayOutput {
	return o
}

func (o Policy8021XArrayOutput) ToPolicy8021XArrayOutputWithContext(ctx context.Context) Policy8021XArrayOutput {
	return o
}

func (o Policy8021XArrayOutput) Index(i pulumi.IntInput) Policy8021XOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Policy8021X {
		return vs[0].([]*Policy8021X)[vs[1].(int)]
	}).(Policy8021XOutput)
}

type Policy8021XMapOutput struct{ *pulumi.OutputState }

func (Policy8021XMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy8021X)(nil)).Elem()
}

func (o Policy8021XMapOutput) ToPolicy8021XMapOutput() Policy8021XMapOutput {
	return o
}

func (o Policy8021XMapOutput) ToPolicy8021XMapOutputWithContext(ctx context.Context) Policy8021XMapOutput {
	return o
}

func (o Policy8021XMapOutput) MapIndex(k pulumi.StringInput) Policy8021XOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Policy8021X {
		return vs[0].(map[string]*Policy8021X)[vs[1].(string)]
	}).(Policy8021XOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Policy8021XInput)(nil)).Elem(), &Policy8021X{})
	pulumi.RegisterInputType(reflect.TypeOf((*Policy8021XArrayInput)(nil)).Elem(), Policy8021XArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Policy8021XMapInput)(nil)).Elem(), Policy8021XMap{})
	pulumi.RegisterOutputType(Policy8021XOutput{})
	pulumi.RegisterOutputType(Policy8021XArrayOutput{})
	pulumi.RegisterOutputType(Policy8021XMapOutput{})
}
