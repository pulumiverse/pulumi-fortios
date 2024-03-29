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

// Configure replacement message groups.
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
//			_, err := system.NewReplacemsggroup(ctx, "trname", &system.ReplacemsggroupArgs{
//				Comment:   pulumi.String("sgh"),
//				GroupType: pulumi.String("utm"),
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
// System ReplacemsgGroup can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/replacemsggroup:Replacemsggroup labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/replacemsggroup:Replacemsggroup labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Replacemsggroup struct {
	pulumi.CustomResourceState

	// Replacement message table entries. The structure of `admin` block is documented below.
	Admins ReplacemsggroupAdminArrayOutput `pulumi:"admins"`
	// Replacement message table entries. The structure of `alertmail` block is documented below.
	Alertmails ReplacemsggroupAlertmailArrayOutput `pulumi:"alertmails"`
	// Replacement message table entries. The structure of `auth` block is documented below.
	Auths ReplacemsggroupAuthArrayOutput `pulumi:"auths"`
	// Replacement message table entries. The structure of `automation` block is documented below.
	Automations ReplacemsggroupAutomationArrayOutput `pulumi:"automations"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Replacement message table entries. The structure of `customMessage` block is documented below.
	CustomMessages ReplacemsggroupCustomMessageArrayOutput `pulumi:"customMessages"`
	// Replacement message table entries. The structure of `deviceDetectionPortal` block is documented below.
	DeviceDetectionPortals ReplacemsggroupDeviceDetectionPortalArrayOutput `pulumi:"deviceDetectionPortals"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Replacement message table entries. The structure of `ec` block is documented below.
	Ecs ReplacemsggroupEcArrayOutput `pulumi:"ecs"`
	// Replacement message table entries. The structure of `fortiguardWf` block is documented below.
	FortiguardWfs ReplacemsggroupFortiguardWfArrayOutput `pulumi:"fortiguardWfs"`
	// Replacement message table entries. The structure of `ftp` block is documented below.
	Ftps ReplacemsggroupFtpArrayOutput `pulumi:"ftps"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Group type.
	GroupType pulumi.StringOutput `pulumi:"groupType"`
	// Replacement message table entries. The structure of `http` block is documented below.
	Https ReplacemsggroupHttpArrayOutput `pulumi:"https"`
	// Replacement message table entries. The structure of `icap` block is documented below.
	Icaps ReplacemsggroupIcapArrayOutput `pulumi:"icaps"`
	// Replacement message table entries. The structure of `mail` block is documented below.
	Mails ReplacemsggroupMailArrayOutput `pulumi:"mails"`
	// Replacement message table entries. The structure of `nacQuar` block is documented below.
	NacQuars ReplacemsggroupNacQuarArrayOutput `pulumi:"nacQuars"`
	// Group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Replacement message table entries. The structure of `nntp` block is documented below.
	Nntps ReplacemsggroupNntpArrayOutput `pulumi:"nntps"`
	// Replacement message table entries. The structure of `spam` block is documented below.
	Spams ReplacemsggroupSpamArrayOutput `pulumi:"spams"`
	// Replacement message table entries. The structure of `sslvpn` block is documented below.
	Sslvpns ReplacemsggroupSslvpnArrayOutput `pulumi:"sslvpns"`
	// Replacement message table entries. The structure of `trafficQuota` block is documented below.
	TrafficQuotas ReplacemsggroupTrafficQuotaArrayOutput `pulumi:"trafficQuotas"`
	// Replacement message table entries. The structure of `utm` block is documented below.
	Utms ReplacemsggroupUtmArrayOutput `pulumi:"utms"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Replacement message table entries. The structure of `webproxy` block is documented below.
	Webproxies ReplacemsggroupWebproxyArrayOutput `pulumi:"webproxies"`
}

// NewReplacemsggroup registers a new resource with the given unique name, arguments, and options.
func NewReplacemsggroup(ctx *pulumi.Context,
	name string, args *ReplacemsggroupArgs, opts ...pulumi.ResourceOption) (*Replacemsggroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupType == nil {
		return nil, errors.New("invalid value for required argument 'GroupType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Replacemsggroup
	err := ctx.RegisterResource("fortios:system/replacemsggroup:Replacemsggroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplacemsggroup gets an existing Replacemsggroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplacemsggroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplacemsggroupState, opts ...pulumi.ResourceOption) (*Replacemsggroup, error) {
	var resource Replacemsggroup
	err := ctx.ReadResource("fortios:system/replacemsggroup:Replacemsggroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Replacemsggroup resources.
type replacemsggroupState struct {
	// Replacement message table entries. The structure of `admin` block is documented below.
	Admins []ReplacemsggroupAdmin `pulumi:"admins"`
	// Replacement message table entries. The structure of `alertmail` block is documented below.
	Alertmails []ReplacemsggroupAlertmail `pulumi:"alertmails"`
	// Replacement message table entries. The structure of `auth` block is documented below.
	Auths []ReplacemsggroupAuth `pulumi:"auths"`
	// Replacement message table entries. The structure of `automation` block is documented below.
	Automations []ReplacemsggroupAutomation `pulumi:"automations"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Replacement message table entries. The structure of `customMessage` block is documented below.
	CustomMessages []ReplacemsggroupCustomMessage `pulumi:"customMessages"`
	// Replacement message table entries. The structure of `deviceDetectionPortal` block is documented below.
	DeviceDetectionPortals []ReplacemsggroupDeviceDetectionPortal `pulumi:"deviceDetectionPortals"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Replacement message table entries. The structure of `ec` block is documented below.
	Ecs []ReplacemsggroupEc `pulumi:"ecs"`
	// Replacement message table entries. The structure of `fortiguardWf` block is documented below.
	FortiguardWfs []ReplacemsggroupFortiguardWf `pulumi:"fortiguardWfs"`
	// Replacement message table entries. The structure of `ftp` block is documented below.
	Ftps []ReplacemsggroupFtp `pulumi:"ftps"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Group type.
	GroupType *string `pulumi:"groupType"`
	// Replacement message table entries. The structure of `http` block is documented below.
	Https []ReplacemsggroupHttp `pulumi:"https"`
	// Replacement message table entries. The structure of `icap` block is documented below.
	Icaps []ReplacemsggroupIcap `pulumi:"icaps"`
	// Replacement message table entries. The structure of `mail` block is documented below.
	Mails []ReplacemsggroupMail `pulumi:"mails"`
	// Replacement message table entries. The structure of `nacQuar` block is documented below.
	NacQuars []ReplacemsggroupNacQuar `pulumi:"nacQuars"`
	// Group name.
	Name *string `pulumi:"name"`
	// Replacement message table entries. The structure of `nntp` block is documented below.
	Nntps []ReplacemsggroupNntp `pulumi:"nntps"`
	// Replacement message table entries. The structure of `spam` block is documented below.
	Spams []ReplacemsggroupSpam `pulumi:"spams"`
	// Replacement message table entries. The structure of `sslvpn` block is documented below.
	Sslvpns []ReplacemsggroupSslvpn `pulumi:"sslvpns"`
	// Replacement message table entries. The structure of `trafficQuota` block is documented below.
	TrafficQuotas []ReplacemsggroupTrafficQuota `pulumi:"trafficQuotas"`
	// Replacement message table entries. The structure of `utm` block is documented below.
	Utms []ReplacemsggroupUtm `pulumi:"utms"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Replacement message table entries. The structure of `webproxy` block is documented below.
	Webproxies []ReplacemsggroupWebproxy `pulumi:"webproxies"`
}

type ReplacemsggroupState struct {
	// Replacement message table entries. The structure of `admin` block is documented below.
	Admins ReplacemsggroupAdminArrayInput
	// Replacement message table entries. The structure of `alertmail` block is documented below.
	Alertmails ReplacemsggroupAlertmailArrayInput
	// Replacement message table entries. The structure of `auth` block is documented below.
	Auths ReplacemsggroupAuthArrayInput
	// Replacement message table entries. The structure of `automation` block is documented below.
	Automations ReplacemsggroupAutomationArrayInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Replacement message table entries. The structure of `customMessage` block is documented below.
	CustomMessages ReplacemsggroupCustomMessageArrayInput
	// Replacement message table entries. The structure of `deviceDetectionPortal` block is documented below.
	DeviceDetectionPortals ReplacemsggroupDeviceDetectionPortalArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Replacement message table entries. The structure of `ec` block is documented below.
	Ecs ReplacemsggroupEcArrayInput
	// Replacement message table entries. The structure of `fortiguardWf` block is documented below.
	FortiguardWfs ReplacemsggroupFortiguardWfArrayInput
	// Replacement message table entries. The structure of `ftp` block is documented below.
	Ftps ReplacemsggroupFtpArrayInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Group type.
	GroupType pulumi.StringPtrInput
	// Replacement message table entries. The structure of `http` block is documented below.
	Https ReplacemsggroupHttpArrayInput
	// Replacement message table entries. The structure of `icap` block is documented below.
	Icaps ReplacemsggroupIcapArrayInput
	// Replacement message table entries. The structure of `mail` block is documented below.
	Mails ReplacemsggroupMailArrayInput
	// Replacement message table entries. The structure of `nacQuar` block is documented below.
	NacQuars ReplacemsggroupNacQuarArrayInput
	// Group name.
	Name pulumi.StringPtrInput
	// Replacement message table entries. The structure of `nntp` block is documented below.
	Nntps ReplacemsggroupNntpArrayInput
	// Replacement message table entries. The structure of `spam` block is documented below.
	Spams ReplacemsggroupSpamArrayInput
	// Replacement message table entries. The structure of `sslvpn` block is documented below.
	Sslvpns ReplacemsggroupSslvpnArrayInput
	// Replacement message table entries. The structure of `trafficQuota` block is documented below.
	TrafficQuotas ReplacemsggroupTrafficQuotaArrayInput
	// Replacement message table entries. The structure of `utm` block is documented below.
	Utms ReplacemsggroupUtmArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Replacement message table entries. The structure of `webproxy` block is documented below.
	Webproxies ReplacemsggroupWebproxyArrayInput
}

func (ReplacemsggroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*replacemsggroupState)(nil)).Elem()
}

type replacemsggroupArgs struct {
	// Replacement message table entries. The structure of `admin` block is documented below.
	Admins []ReplacemsggroupAdmin `pulumi:"admins"`
	// Replacement message table entries. The structure of `alertmail` block is documented below.
	Alertmails []ReplacemsggroupAlertmail `pulumi:"alertmails"`
	// Replacement message table entries. The structure of `auth` block is documented below.
	Auths []ReplacemsggroupAuth `pulumi:"auths"`
	// Replacement message table entries. The structure of `automation` block is documented below.
	Automations []ReplacemsggroupAutomation `pulumi:"automations"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Replacement message table entries. The structure of `customMessage` block is documented below.
	CustomMessages []ReplacemsggroupCustomMessage `pulumi:"customMessages"`
	// Replacement message table entries. The structure of `deviceDetectionPortal` block is documented below.
	DeviceDetectionPortals []ReplacemsggroupDeviceDetectionPortal `pulumi:"deviceDetectionPortals"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Replacement message table entries. The structure of `ec` block is documented below.
	Ecs []ReplacemsggroupEc `pulumi:"ecs"`
	// Replacement message table entries. The structure of `fortiguardWf` block is documented below.
	FortiguardWfs []ReplacemsggroupFortiguardWf `pulumi:"fortiguardWfs"`
	// Replacement message table entries. The structure of `ftp` block is documented below.
	Ftps []ReplacemsggroupFtp `pulumi:"ftps"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Group type.
	GroupType string `pulumi:"groupType"`
	// Replacement message table entries. The structure of `http` block is documented below.
	Https []ReplacemsggroupHttp `pulumi:"https"`
	// Replacement message table entries. The structure of `icap` block is documented below.
	Icaps []ReplacemsggroupIcap `pulumi:"icaps"`
	// Replacement message table entries. The structure of `mail` block is documented below.
	Mails []ReplacemsggroupMail `pulumi:"mails"`
	// Replacement message table entries. The structure of `nacQuar` block is documented below.
	NacQuars []ReplacemsggroupNacQuar `pulumi:"nacQuars"`
	// Group name.
	Name *string `pulumi:"name"`
	// Replacement message table entries. The structure of `nntp` block is documented below.
	Nntps []ReplacemsggroupNntp `pulumi:"nntps"`
	// Replacement message table entries. The structure of `spam` block is documented below.
	Spams []ReplacemsggroupSpam `pulumi:"spams"`
	// Replacement message table entries. The structure of `sslvpn` block is documented below.
	Sslvpns []ReplacemsggroupSslvpn `pulumi:"sslvpns"`
	// Replacement message table entries. The structure of `trafficQuota` block is documented below.
	TrafficQuotas []ReplacemsggroupTrafficQuota `pulumi:"trafficQuotas"`
	// Replacement message table entries. The structure of `utm` block is documented below.
	Utms []ReplacemsggroupUtm `pulumi:"utms"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Replacement message table entries. The structure of `webproxy` block is documented below.
	Webproxies []ReplacemsggroupWebproxy `pulumi:"webproxies"`
}

// The set of arguments for constructing a Replacemsggroup resource.
type ReplacemsggroupArgs struct {
	// Replacement message table entries. The structure of `admin` block is documented below.
	Admins ReplacemsggroupAdminArrayInput
	// Replacement message table entries. The structure of `alertmail` block is documented below.
	Alertmails ReplacemsggroupAlertmailArrayInput
	// Replacement message table entries. The structure of `auth` block is documented below.
	Auths ReplacemsggroupAuthArrayInput
	// Replacement message table entries. The structure of `automation` block is documented below.
	Automations ReplacemsggroupAutomationArrayInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Replacement message table entries. The structure of `customMessage` block is documented below.
	CustomMessages ReplacemsggroupCustomMessageArrayInput
	// Replacement message table entries. The structure of `deviceDetectionPortal` block is documented below.
	DeviceDetectionPortals ReplacemsggroupDeviceDetectionPortalArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Replacement message table entries. The structure of `ec` block is documented below.
	Ecs ReplacemsggroupEcArrayInput
	// Replacement message table entries. The structure of `fortiguardWf` block is documented below.
	FortiguardWfs ReplacemsggroupFortiguardWfArrayInput
	// Replacement message table entries. The structure of `ftp` block is documented below.
	Ftps ReplacemsggroupFtpArrayInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Group type.
	GroupType pulumi.StringInput
	// Replacement message table entries. The structure of `http` block is documented below.
	Https ReplacemsggroupHttpArrayInput
	// Replacement message table entries. The structure of `icap` block is documented below.
	Icaps ReplacemsggroupIcapArrayInput
	// Replacement message table entries. The structure of `mail` block is documented below.
	Mails ReplacemsggroupMailArrayInput
	// Replacement message table entries. The structure of `nacQuar` block is documented below.
	NacQuars ReplacemsggroupNacQuarArrayInput
	// Group name.
	Name pulumi.StringPtrInput
	// Replacement message table entries. The structure of `nntp` block is documented below.
	Nntps ReplacemsggroupNntpArrayInput
	// Replacement message table entries. The structure of `spam` block is documented below.
	Spams ReplacemsggroupSpamArrayInput
	// Replacement message table entries. The structure of `sslvpn` block is documented below.
	Sslvpns ReplacemsggroupSslvpnArrayInput
	// Replacement message table entries. The structure of `trafficQuota` block is documented below.
	TrafficQuotas ReplacemsggroupTrafficQuotaArrayInput
	// Replacement message table entries. The structure of `utm` block is documented below.
	Utms ReplacemsggroupUtmArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Replacement message table entries. The structure of `webproxy` block is documented below.
	Webproxies ReplacemsggroupWebproxyArrayInput
}

func (ReplacemsggroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replacemsggroupArgs)(nil)).Elem()
}

type ReplacemsggroupInput interface {
	pulumi.Input

	ToReplacemsggroupOutput() ReplacemsggroupOutput
	ToReplacemsggroupOutputWithContext(ctx context.Context) ReplacemsggroupOutput
}

func (*Replacemsggroup) ElementType() reflect.Type {
	return reflect.TypeOf((**Replacemsggroup)(nil)).Elem()
}

func (i *Replacemsggroup) ToReplacemsggroupOutput() ReplacemsggroupOutput {
	return i.ToReplacemsggroupOutputWithContext(context.Background())
}

func (i *Replacemsggroup) ToReplacemsggroupOutputWithContext(ctx context.Context) ReplacemsggroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplacemsggroupOutput)
}

// ReplacemsggroupArrayInput is an input type that accepts ReplacemsggroupArray and ReplacemsggroupArrayOutput values.
// You can construct a concrete instance of `ReplacemsggroupArrayInput` via:
//
//	ReplacemsggroupArray{ ReplacemsggroupArgs{...} }
type ReplacemsggroupArrayInput interface {
	pulumi.Input

	ToReplacemsggroupArrayOutput() ReplacemsggroupArrayOutput
	ToReplacemsggroupArrayOutputWithContext(context.Context) ReplacemsggroupArrayOutput
}

type ReplacemsggroupArray []ReplacemsggroupInput

func (ReplacemsggroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Replacemsggroup)(nil)).Elem()
}

func (i ReplacemsggroupArray) ToReplacemsggroupArrayOutput() ReplacemsggroupArrayOutput {
	return i.ToReplacemsggroupArrayOutputWithContext(context.Background())
}

func (i ReplacemsggroupArray) ToReplacemsggroupArrayOutputWithContext(ctx context.Context) ReplacemsggroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplacemsggroupArrayOutput)
}

// ReplacemsggroupMapInput is an input type that accepts ReplacemsggroupMap and ReplacemsggroupMapOutput values.
// You can construct a concrete instance of `ReplacemsggroupMapInput` via:
//
//	ReplacemsggroupMap{ "key": ReplacemsggroupArgs{...} }
type ReplacemsggroupMapInput interface {
	pulumi.Input

	ToReplacemsggroupMapOutput() ReplacemsggroupMapOutput
	ToReplacemsggroupMapOutputWithContext(context.Context) ReplacemsggroupMapOutput
}

type ReplacemsggroupMap map[string]ReplacemsggroupInput

func (ReplacemsggroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Replacemsggroup)(nil)).Elem()
}

func (i ReplacemsggroupMap) ToReplacemsggroupMapOutput() ReplacemsggroupMapOutput {
	return i.ToReplacemsggroupMapOutputWithContext(context.Background())
}

func (i ReplacemsggroupMap) ToReplacemsggroupMapOutputWithContext(ctx context.Context) ReplacemsggroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplacemsggroupMapOutput)
}

type ReplacemsggroupOutput struct{ *pulumi.OutputState }

func (ReplacemsggroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Replacemsggroup)(nil)).Elem()
}

func (o ReplacemsggroupOutput) ToReplacemsggroupOutput() ReplacemsggroupOutput {
	return o
}

func (o ReplacemsggroupOutput) ToReplacemsggroupOutputWithContext(ctx context.Context) ReplacemsggroupOutput {
	return o
}

// Replacement message table entries. The structure of `admin` block is documented below.
func (o ReplacemsggroupOutput) Admins() ReplacemsggroupAdminArrayOutput {
	return o.ApplyT(func(v *Replacemsggroup) ReplacemsggroupAdminArrayOutput { return v.Admins }).(ReplacemsggroupAdminArrayOutput)
}

// Replacement message table entries. The structure of `alertmail` block is documented below.
func (o ReplacemsggroupOutput) Alertmails() ReplacemsggroupAlertmailArrayOutput {
	return o.ApplyT(func(v *Replacemsggroup) ReplacemsggroupAlertmailArrayOutput { return v.Alertmails }).(ReplacemsggroupAlertmailArrayOutput)
}

// Replacement message table entries. The structure of `auth` block is documented below.
func (o ReplacemsggroupOutput) Auths() ReplacemsggroupAuthArrayOutput {
	return o.ApplyT(func(v *Replacemsggroup) ReplacemsggroupAuthArrayOutput { return v.Auths }).(ReplacemsggroupAuthArrayOutput)
}

// Replacement message table entries. The structure of `automation` block is documented below.
func (o ReplacemsggroupOutput) Automations() ReplacemsggroupAutomationArrayOutput {
	return o.ApplyT(func(v *Replacemsggroup) ReplacemsggroupAutomationArrayOutput { return v.Automations }).(ReplacemsggroupAutomationArrayOutput)
}

// Comment.
func (o ReplacemsggroupOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Replacemsggroup) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Replacement message table entries. The structure of `customMessage` block is documented below.
func (o ReplacemsggroupOutput) CustomMessages() ReplacemsggroupCustomMessageArrayOutput {
	return o.ApplyT(func(v *Replacemsggroup) ReplacemsggroupCustomMessageArrayOutput { return v.CustomMessages }).(ReplacemsggroupCustomMessageArrayOutput)
}

// Replacement message table entries. The structure of `deviceDetectionPortal` block is documented below.
func (o ReplacemsggroupOutput) DeviceDetectionPortals() ReplacemsggroupDeviceDetectionPortalArrayOutput {
	return o.ApplyT(func(v *Replacemsggroup) ReplacemsggroupDeviceDetectionPortalArrayOutput {
		return v.DeviceDetectionPortals
	}).(ReplacemsggroupDeviceDetectionPortalArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o ReplacemsggroupOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Replacemsggroup) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Replacement message table entries. The structure of `ec` block is documented below.
func (o ReplacemsggroupOutput) Ecs() ReplacemsggroupEcArrayOutput {
	return o.ApplyT(func(v *Replacemsggroup) ReplacemsggroupEcArrayOutput { return v.Ecs }).(ReplacemsggroupEcArrayOutput)
}

// Replacement message table entries. The structure of `fortiguardWf` block is documented below.
func (o ReplacemsggroupOutput) FortiguardWfs() ReplacemsggroupFortiguardWfArrayOutput {
	return o.ApplyT(func(v *Replacemsggroup) ReplacemsggroupFortiguardWfArrayOutput { return v.FortiguardWfs }).(ReplacemsggroupFortiguardWfArrayOutput)
}

// Replacement message table entries. The structure of `ftp` block is documented below.
func (o ReplacemsggroupOutput) Ftps() ReplacemsggroupFtpArrayOutput {
	return o.ApplyT(func(v *Replacemsggroup) ReplacemsggroupFtpArrayOutput { return v.Ftps }).(ReplacemsggroupFtpArrayOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o ReplacemsggroupOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Replacemsggroup) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Group type.
func (o ReplacemsggroupOutput) GroupType() pulumi.StringOutput {
	return o.ApplyT(func(v *Replacemsggroup) pulumi.StringOutput { return v.GroupType }).(pulumi.StringOutput)
}

// Replacement message table entries. The structure of `http` block is documented below.
func (o ReplacemsggroupOutput) Https() ReplacemsggroupHttpArrayOutput {
	return o.ApplyT(func(v *Replacemsggroup) ReplacemsggroupHttpArrayOutput { return v.Https }).(ReplacemsggroupHttpArrayOutput)
}

// Replacement message table entries. The structure of `icap` block is documented below.
func (o ReplacemsggroupOutput) Icaps() ReplacemsggroupIcapArrayOutput {
	return o.ApplyT(func(v *Replacemsggroup) ReplacemsggroupIcapArrayOutput { return v.Icaps }).(ReplacemsggroupIcapArrayOutput)
}

// Replacement message table entries. The structure of `mail` block is documented below.
func (o ReplacemsggroupOutput) Mails() ReplacemsggroupMailArrayOutput {
	return o.ApplyT(func(v *Replacemsggroup) ReplacemsggroupMailArrayOutput { return v.Mails }).(ReplacemsggroupMailArrayOutput)
}

// Replacement message table entries. The structure of `nacQuar` block is documented below.
func (o ReplacemsggroupOutput) NacQuars() ReplacemsggroupNacQuarArrayOutput {
	return o.ApplyT(func(v *Replacemsggroup) ReplacemsggroupNacQuarArrayOutput { return v.NacQuars }).(ReplacemsggroupNacQuarArrayOutput)
}

// Group name.
func (o ReplacemsggroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Replacemsggroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Replacement message table entries. The structure of `nntp` block is documented below.
func (o ReplacemsggroupOutput) Nntps() ReplacemsggroupNntpArrayOutput {
	return o.ApplyT(func(v *Replacemsggroup) ReplacemsggroupNntpArrayOutput { return v.Nntps }).(ReplacemsggroupNntpArrayOutput)
}

// Replacement message table entries. The structure of `spam` block is documented below.
func (o ReplacemsggroupOutput) Spams() ReplacemsggroupSpamArrayOutput {
	return o.ApplyT(func(v *Replacemsggroup) ReplacemsggroupSpamArrayOutput { return v.Spams }).(ReplacemsggroupSpamArrayOutput)
}

// Replacement message table entries. The structure of `sslvpn` block is documented below.
func (o ReplacemsggroupOutput) Sslvpns() ReplacemsggroupSslvpnArrayOutput {
	return o.ApplyT(func(v *Replacemsggroup) ReplacemsggroupSslvpnArrayOutput { return v.Sslvpns }).(ReplacemsggroupSslvpnArrayOutput)
}

// Replacement message table entries. The structure of `trafficQuota` block is documented below.
func (o ReplacemsggroupOutput) TrafficQuotas() ReplacemsggroupTrafficQuotaArrayOutput {
	return o.ApplyT(func(v *Replacemsggroup) ReplacemsggroupTrafficQuotaArrayOutput { return v.TrafficQuotas }).(ReplacemsggroupTrafficQuotaArrayOutput)
}

// Replacement message table entries. The structure of `utm` block is documented below.
func (o ReplacemsggroupOutput) Utms() ReplacemsggroupUtmArrayOutput {
	return o.ApplyT(func(v *Replacemsggroup) ReplacemsggroupUtmArrayOutput { return v.Utms }).(ReplacemsggroupUtmArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ReplacemsggroupOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Replacemsggroup) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Replacement message table entries. The structure of `webproxy` block is documented below.
func (o ReplacemsggroupOutput) Webproxies() ReplacemsggroupWebproxyArrayOutput {
	return o.ApplyT(func(v *Replacemsggroup) ReplacemsggroupWebproxyArrayOutput { return v.Webproxies }).(ReplacemsggroupWebproxyArrayOutput)
}

type ReplacemsggroupArrayOutput struct{ *pulumi.OutputState }

func (ReplacemsggroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Replacemsggroup)(nil)).Elem()
}

func (o ReplacemsggroupArrayOutput) ToReplacemsggroupArrayOutput() ReplacemsggroupArrayOutput {
	return o
}

func (o ReplacemsggroupArrayOutput) ToReplacemsggroupArrayOutputWithContext(ctx context.Context) ReplacemsggroupArrayOutput {
	return o
}

func (o ReplacemsggroupArrayOutput) Index(i pulumi.IntInput) ReplacemsggroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Replacemsggroup {
		return vs[0].([]*Replacemsggroup)[vs[1].(int)]
	}).(ReplacemsggroupOutput)
}

type ReplacemsggroupMapOutput struct{ *pulumi.OutputState }

func (ReplacemsggroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Replacemsggroup)(nil)).Elem()
}

func (o ReplacemsggroupMapOutput) ToReplacemsggroupMapOutput() ReplacemsggroupMapOutput {
	return o
}

func (o ReplacemsggroupMapOutput) ToReplacemsggroupMapOutputWithContext(ctx context.Context) ReplacemsggroupMapOutput {
	return o
}

func (o ReplacemsggroupMapOutput) MapIndex(k pulumi.StringInput) ReplacemsggroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Replacemsggroup {
		return vs[0].(map[string]*Replacemsggroup)[vs[1].(string)]
	}).(ReplacemsggroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplacemsggroupInput)(nil)).Elem(), &Replacemsggroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplacemsggroupArrayInput)(nil)).Elem(), ReplacemsggroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplacemsggroupMapInput)(nil)).Elem(), ReplacemsggroupMap{})
	pulumi.RegisterOutputType(ReplacemsggroupOutput{})
	pulumi.RegisterOutputType(ReplacemsggroupArrayOutput{})
	pulumi.RegisterOutputType(ReplacemsggroupMapOutput{})
}
