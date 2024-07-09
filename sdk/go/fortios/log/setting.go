// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package log

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure general log settings.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/log"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := log.NewSetting(ctx, "trname", &log.SettingArgs{
//				BriefTrafficFormat:   pulumi.String("disable"),
//				DaemonLog:            pulumi.String("disable"),
//				ExpolicyImplicitLog:  pulumi.String("disable"),
//				FazOverride:          pulumi.String("disable"),
//				Fwpolicy6ImplicitLog: pulumi.String("disable"),
//				FwpolicyImplicitLog:  pulumi.String("disable"),
//				LocalInAllow:         pulumi.String("disable"),
//				LocalInDenyBroadcast: pulumi.String("disable"),
//				LocalInDenyUnicast:   pulumi.String("disable"),
//				LocalOut:             pulumi.String("disable"),
//				LogInvalidPacket:     pulumi.String("disable"),
//				LogPolicyComment:     pulumi.String("disable"),
//				LogPolicyName:        pulumi.String("disable"),
//				LogUserInUpper:       pulumi.String("disable"),
//				NeighborEvent:        pulumi.String("disable"),
//				ResolveIp:            pulumi.String("disable"),
//				ResolvePort:          pulumi.String("enable"),
//				SyslogOverride:       pulumi.String("disable"),
//				UserAnonymize:        pulumi.String("disable"),
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
// Log Setting can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:log/setting:Setting labelname LogSetting
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:log/setting:Setting labelname LogSetting
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Setting struct {
	pulumi.CustomResourceState

	// User name anonymization hash salt.
	AnonymizationHash pulumi.StringOutput `pulumi:"anonymizationHash"`
	// Enable/disable brief format traffic logging. Valid values: `enable`, `disable`.
	BriefTrafficFormat pulumi.StringOutput `pulumi:"briefTrafficFormat"`
	// Custom fields to append to all log messages. The structure of `customLogFields` block is documented below.
	CustomLogFields SettingCustomLogFieldArrayOutput `pulumi:"customLogFields"`
	// Enable/disable daemon logging. Valid values: `enable`, `disable`.
	DaemonLog pulumi.StringOutput `pulumi:"daemonLog"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable explicit proxy firewall implicit policy logging. Valid values: `enable`, `disable`.
	ExpolicyImplicitLog pulumi.StringOutput `pulumi:"expolicyImplicitLog"`
	// Enable/disable extended traffic logging. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringOutput `pulumi:"extendedLog"`
	// Enable/disable override FortiAnalyzer settings. Valid values: `enable`, `disable`.
	FazOverride pulumi.StringOutput `pulumi:"fazOverride"`
	// Enable/disable implicit firewall policy6 logging. Valid values: `enable`, `disable`.
	Fwpolicy6ImplicitLog pulumi.StringOutput `pulumi:"fwpolicy6ImplicitLog"`
	// Enable/disable implicit firewall policy logging. Valid values: `enable`, `disable`.
	FwpolicyImplicitLog pulumi.StringOutput `pulumi:"fwpolicyImplicitLog"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Enable/disable local-in-allow logging. Valid values: `enable`, `disable`.
	LocalInAllow pulumi.StringOutput `pulumi:"localInAllow"`
	// Enable/disable local-in-deny-broadcast logging. Valid values: `enable`, `disable`.
	LocalInDenyBroadcast pulumi.StringOutput `pulumi:"localInDenyBroadcast"`
	// Enable/disable local-in-deny-unicast logging. Valid values: `enable`, `disable`.
	LocalInDenyUnicast pulumi.StringOutput `pulumi:"localInDenyUnicast"`
	// Enable/disable local-out logging. Valid values: `enable`, `disable`.
	LocalOut pulumi.StringOutput `pulumi:"localOut"`
	// Enable/disable local-out traffic IoC detection. Requires local-out to be enabled. Valid values: `enable`, `disable`.
	LocalOutIocDetection pulumi.StringOutput `pulumi:"localOutIocDetection"`
	// Enable/disable invalid packet traffic logging. Valid values: `enable`, `disable`.
	LogInvalidPacket pulumi.StringOutput `pulumi:"logInvalidPacket"`
	// Enable/disable inserting policy comments into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyComment pulumi.StringOutput `pulumi:"logPolicyComment"`
	// Enable/disable inserting policy name into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyName pulumi.StringOutput `pulumi:"logPolicyName"`
	// Enable/disable logs with user-in-upper. Valid values: `enable`, `disable`.
	LogUserInUpper pulumi.StringOutput `pulumi:"logUserInUpper"`
	// Enable/disable long-live-session statistics logging. Valid values: `enable`, `disable`.
	LongLiveSessionStat pulumi.StringOutput `pulumi:"longLiveSessionStat"`
	// Enable/disable neighbor event logging. Valid values: `enable`, `disable`.
	NeighborEvent pulumi.StringOutput `pulumi:"neighborEvent"`
	// Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `enable`, `disable`.
	ResolveIp pulumi.StringOutput `pulumi:"resolveIp"`
	// Enable/disable adding resolved service names to traffic logs. Valid values: `enable`, `disable`.
	ResolvePort pulumi.StringOutput `pulumi:"resolvePort"`
	// Enable/disable REST API GET request logging. Valid values: `enable`, `disable`.
	RestApiGet pulumi.StringOutput `pulumi:"restApiGet"`
	// Enable/disable REST API POST/PUT/DELETE request logging. Valid values: `enable`, `disable`.
	RestApiSet pulumi.StringOutput `pulumi:"restApiSet"`
	// Enable/disable override Syslog settings. Valid values: `enable`, `disable`.
	SyslogOverride pulumi.StringOutput `pulumi:"syslogOverride"`
	// Enable/disable anonymizing user names in log messages. Valid values: `enable`, `disable`.
	UserAnonymize pulumi.StringOutput `pulumi:"userAnonymize"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewSetting registers a new resource with the given unique name, arguments, and options.
func NewSetting(ctx *pulumi.Context,
	name string, args *SettingArgs, opts ...pulumi.ResourceOption) (*Setting, error) {
	if args == nil {
		args = &SettingArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Setting
	err := ctx.RegisterResource("fortios:log/setting:Setting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSetting gets an existing Setting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SettingState, opts ...pulumi.ResourceOption) (*Setting, error) {
	var resource Setting
	err := ctx.ReadResource("fortios:log/setting:Setting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Setting resources.
type settingState struct {
	// User name anonymization hash salt.
	AnonymizationHash *string `pulumi:"anonymizationHash"`
	// Enable/disable brief format traffic logging. Valid values: `enable`, `disable`.
	BriefTrafficFormat *string `pulumi:"briefTrafficFormat"`
	// Custom fields to append to all log messages. The structure of `customLogFields` block is documented below.
	CustomLogFields []SettingCustomLogField `pulumi:"customLogFields"`
	// Enable/disable daemon logging. Valid values: `enable`, `disable`.
	DaemonLog *string `pulumi:"daemonLog"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable explicit proxy firewall implicit policy logging. Valid values: `enable`, `disable`.
	ExpolicyImplicitLog *string `pulumi:"expolicyImplicitLog"`
	// Enable/disable extended traffic logging. Valid values: `enable`, `disable`.
	ExtendedLog *string `pulumi:"extendedLog"`
	// Enable/disable override FortiAnalyzer settings. Valid values: `enable`, `disable`.
	FazOverride *string `pulumi:"fazOverride"`
	// Enable/disable implicit firewall policy6 logging. Valid values: `enable`, `disable`.
	Fwpolicy6ImplicitLog *string `pulumi:"fwpolicy6ImplicitLog"`
	// Enable/disable implicit firewall policy logging. Valid values: `enable`, `disable`.
	FwpolicyImplicitLog *string `pulumi:"fwpolicyImplicitLog"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Enable/disable local-in-allow logging. Valid values: `enable`, `disable`.
	LocalInAllow *string `pulumi:"localInAllow"`
	// Enable/disable local-in-deny-broadcast logging. Valid values: `enable`, `disable`.
	LocalInDenyBroadcast *string `pulumi:"localInDenyBroadcast"`
	// Enable/disable local-in-deny-unicast logging. Valid values: `enable`, `disable`.
	LocalInDenyUnicast *string `pulumi:"localInDenyUnicast"`
	// Enable/disable local-out logging. Valid values: `enable`, `disable`.
	LocalOut *string `pulumi:"localOut"`
	// Enable/disable local-out traffic IoC detection. Requires local-out to be enabled. Valid values: `enable`, `disable`.
	LocalOutIocDetection *string `pulumi:"localOutIocDetection"`
	// Enable/disable invalid packet traffic logging. Valid values: `enable`, `disable`.
	LogInvalidPacket *string `pulumi:"logInvalidPacket"`
	// Enable/disable inserting policy comments into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyComment *string `pulumi:"logPolicyComment"`
	// Enable/disable inserting policy name into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyName *string `pulumi:"logPolicyName"`
	// Enable/disable logs with user-in-upper. Valid values: `enable`, `disable`.
	LogUserInUpper *string `pulumi:"logUserInUpper"`
	// Enable/disable long-live-session statistics logging. Valid values: `enable`, `disable`.
	LongLiveSessionStat *string `pulumi:"longLiveSessionStat"`
	// Enable/disable neighbor event logging. Valid values: `enable`, `disable`.
	NeighborEvent *string `pulumi:"neighborEvent"`
	// Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `enable`, `disable`.
	ResolveIp *string `pulumi:"resolveIp"`
	// Enable/disable adding resolved service names to traffic logs. Valid values: `enable`, `disable`.
	ResolvePort *string `pulumi:"resolvePort"`
	// Enable/disable REST API GET request logging. Valid values: `enable`, `disable`.
	RestApiGet *string `pulumi:"restApiGet"`
	// Enable/disable REST API POST/PUT/DELETE request logging. Valid values: `enable`, `disable`.
	RestApiSet *string `pulumi:"restApiSet"`
	// Enable/disable override Syslog settings. Valid values: `enable`, `disable`.
	SyslogOverride *string `pulumi:"syslogOverride"`
	// Enable/disable anonymizing user names in log messages. Valid values: `enable`, `disable`.
	UserAnonymize *string `pulumi:"userAnonymize"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SettingState struct {
	// User name anonymization hash salt.
	AnonymizationHash pulumi.StringPtrInput
	// Enable/disable brief format traffic logging. Valid values: `enable`, `disable`.
	BriefTrafficFormat pulumi.StringPtrInput
	// Custom fields to append to all log messages. The structure of `customLogFields` block is documented below.
	CustomLogFields SettingCustomLogFieldArrayInput
	// Enable/disable daemon logging. Valid values: `enable`, `disable`.
	DaemonLog pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable explicit proxy firewall implicit policy logging. Valid values: `enable`, `disable`.
	ExpolicyImplicitLog pulumi.StringPtrInput
	// Enable/disable extended traffic logging. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringPtrInput
	// Enable/disable override FortiAnalyzer settings. Valid values: `enable`, `disable`.
	FazOverride pulumi.StringPtrInput
	// Enable/disable implicit firewall policy6 logging. Valid values: `enable`, `disable`.
	Fwpolicy6ImplicitLog pulumi.StringPtrInput
	// Enable/disable implicit firewall policy logging. Valid values: `enable`, `disable`.
	FwpolicyImplicitLog pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Enable/disable local-in-allow logging. Valid values: `enable`, `disable`.
	LocalInAllow pulumi.StringPtrInput
	// Enable/disable local-in-deny-broadcast logging. Valid values: `enable`, `disable`.
	LocalInDenyBroadcast pulumi.StringPtrInput
	// Enable/disable local-in-deny-unicast logging. Valid values: `enable`, `disable`.
	LocalInDenyUnicast pulumi.StringPtrInput
	// Enable/disable local-out logging. Valid values: `enable`, `disable`.
	LocalOut pulumi.StringPtrInput
	// Enable/disable local-out traffic IoC detection. Requires local-out to be enabled. Valid values: `enable`, `disable`.
	LocalOutIocDetection pulumi.StringPtrInput
	// Enable/disable invalid packet traffic logging. Valid values: `enable`, `disable`.
	LogInvalidPacket pulumi.StringPtrInput
	// Enable/disable inserting policy comments into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyComment pulumi.StringPtrInput
	// Enable/disable inserting policy name into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyName pulumi.StringPtrInput
	// Enable/disable logs with user-in-upper. Valid values: `enable`, `disable`.
	LogUserInUpper pulumi.StringPtrInput
	// Enable/disable long-live-session statistics logging. Valid values: `enable`, `disable`.
	LongLiveSessionStat pulumi.StringPtrInput
	// Enable/disable neighbor event logging. Valid values: `enable`, `disable`.
	NeighborEvent pulumi.StringPtrInput
	// Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `enable`, `disable`.
	ResolveIp pulumi.StringPtrInput
	// Enable/disable adding resolved service names to traffic logs. Valid values: `enable`, `disable`.
	ResolvePort pulumi.StringPtrInput
	// Enable/disable REST API GET request logging. Valid values: `enable`, `disable`.
	RestApiGet pulumi.StringPtrInput
	// Enable/disable REST API POST/PUT/DELETE request logging. Valid values: `enable`, `disable`.
	RestApiSet pulumi.StringPtrInput
	// Enable/disable override Syslog settings. Valid values: `enable`, `disable`.
	SyslogOverride pulumi.StringPtrInput
	// Enable/disable anonymizing user names in log messages. Valid values: `enable`, `disable`.
	UserAnonymize pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*settingState)(nil)).Elem()
}

type settingArgs struct {
	// User name anonymization hash salt.
	AnonymizationHash *string `pulumi:"anonymizationHash"`
	// Enable/disable brief format traffic logging. Valid values: `enable`, `disable`.
	BriefTrafficFormat *string `pulumi:"briefTrafficFormat"`
	// Custom fields to append to all log messages. The structure of `customLogFields` block is documented below.
	CustomLogFields []SettingCustomLogField `pulumi:"customLogFields"`
	// Enable/disable daemon logging. Valid values: `enable`, `disable`.
	DaemonLog *string `pulumi:"daemonLog"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable explicit proxy firewall implicit policy logging. Valid values: `enable`, `disable`.
	ExpolicyImplicitLog *string `pulumi:"expolicyImplicitLog"`
	// Enable/disable extended traffic logging. Valid values: `enable`, `disable`.
	ExtendedLog *string `pulumi:"extendedLog"`
	// Enable/disable override FortiAnalyzer settings. Valid values: `enable`, `disable`.
	FazOverride *string `pulumi:"fazOverride"`
	// Enable/disable implicit firewall policy6 logging. Valid values: `enable`, `disable`.
	Fwpolicy6ImplicitLog *string `pulumi:"fwpolicy6ImplicitLog"`
	// Enable/disable implicit firewall policy logging. Valid values: `enable`, `disable`.
	FwpolicyImplicitLog *string `pulumi:"fwpolicyImplicitLog"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Enable/disable local-in-allow logging. Valid values: `enable`, `disable`.
	LocalInAllow *string `pulumi:"localInAllow"`
	// Enable/disable local-in-deny-broadcast logging. Valid values: `enable`, `disable`.
	LocalInDenyBroadcast *string `pulumi:"localInDenyBroadcast"`
	// Enable/disable local-in-deny-unicast logging. Valid values: `enable`, `disable`.
	LocalInDenyUnicast *string `pulumi:"localInDenyUnicast"`
	// Enable/disable local-out logging. Valid values: `enable`, `disable`.
	LocalOut *string `pulumi:"localOut"`
	// Enable/disable local-out traffic IoC detection. Requires local-out to be enabled. Valid values: `enable`, `disable`.
	LocalOutIocDetection *string `pulumi:"localOutIocDetection"`
	// Enable/disable invalid packet traffic logging. Valid values: `enable`, `disable`.
	LogInvalidPacket *string `pulumi:"logInvalidPacket"`
	// Enable/disable inserting policy comments into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyComment *string `pulumi:"logPolicyComment"`
	// Enable/disable inserting policy name into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyName *string `pulumi:"logPolicyName"`
	// Enable/disable logs with user-in-upper. Valid values: `enable`, `disable`.
	LogUserInUpper *string `pulumi:"logUserInUpper"`
	// Enable/disable long-live-session statistics logging. Valid values: `enable`, `disable`.
	LongLiveSessionStat *string `pulumi:"longLiveSessionStat"`
	// Enable/disable neighbor event logging. Valid values: `enable`, `disable`.
	NeighborEvent *string `pulumi:"neighborEvent"`
	// Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `enable`, `disable`.
	ResolveIp *string `pulumi:"resolveIp"`
	// Enable/disable adding resolved service names to traffic logs. Valid values: `enable`, `disable`.
	ResolvePort *string `pulumi:"resolvePort"`
	// Enable/disable REST API GET request logging. Valid values: `enable`, `disable`.
	RestApiGet *string `pulumi:"restApiGet"`
	// Enable/disable REST API POST/PUT/DELETE request logging. Valid values: `enable`, `disable`.
	RestApiSet *string `pulumi:"restApiSet"`
	// Enable/disable override Syslog settings. Valid values: `enable`, `disable`.
	SyslogOverride *string `pulumi:"syslogOverride"`
	// Enable/disable anonymizing user names in log messages. Valid values: `enable`, `disable`.
	UserAnonymize *string `pulumi:"userAnonymize"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Setting resource.
type SettingArgs struct {
	// User name anonymization hash salt.
	AnonymizationHash pulumi.StringPtrInput
	// Enable/disable brief format traffic logging. Valid values: `enable`, `disable`.
	BriefTrafficFormat pulumi.StringPtrInput
	// Custom fields to append to all log messages. The structure of `customLogFields` block is documented below.
	CustomLogFields SettingCustomLogFieldArrayInput
	// Enable/disable daemon logging. Valid values: `enable`, `disable`.
	DaemonLog pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable explicit proxy firewall implicit policy logging. Valid values: `enable`, `disable`.
	ExpolicyImplicitLog pulumi.StringPtrInput
	// Enable/disable extended traffic logging. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringPtrInput
	// Enable/disable override FortiAnalyzer settings. Valid values: `enable`, `disable`.
	FazOverride pulumi.StringPtrInput
	// Enable/disable implicit firewall policy6 logging. Valid values: `enable`, `disable`.
	Fwpolicy6ImplicitLog pulumi.StringPtrInput
	// Enable/disable implicit firewall policy logging. Valid values: `enable`, `disable`.
	FwpolicyImplicitLog pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Enable/disable local-in-allow logging. Valid values: `enable`, `disable`.
	LocalInAllow pulumi.StringPtrInput
	// Enable/disable local-in-deny-broadcast logging. Valid values: `enable`, `disable`.
	LocalInDenyBroadcast pulumi.StringPtrInput
	// Enable/disable local-in-deny-unicast logging. Valid values: `enable`, `disable`.
	LocalInDenyUnicast pulumi.StringPtrInput
	// Enable/disable local-out logging. Valid values: `enable`, `disable`.
	LocalOut pulumi.StringPtrInput
	// Enable/disable local-out traffic IoC detection. Requires local-out to be enabled. Valid values: `enable`, `disable`.
	LocalOutIocDetection pulumi.StringPtrInput
	// Enable/disable invalid packet traffic logging. Valid values: `enable`, `disable`.
	LogInvalidPacket pulumi.StringPtrInput
	// Enable/disable inserting policy comments into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyComment pulumi.StringPtrInput
	// Enable/disable inserting policy name into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyName pulumi.StringPtrInput
	// Enable/disable logs with user-in-upper. Valid values: `enable`, `disable`.
	LogUserInUpper pulumi.StringPtrInput
	// Enable/disable long-live-session statistics logging. Valid values: `enable`, `disable`.
	LongLiveSessionStat pulumi.StringPtrInput
	// Enable/disable neighbor event logging. Valid values: `enable`, `disable`.
	NeighborEvent pulumi.StringPtrInput
	// Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `enable`, `disable`.
	ResolveIp pulumi.StringPtrInput
	// Enable/disable adding resolved service names to traffic logs. Valid values: `enable`, `disable`.
	ResolvePort pulumi.StringPtrInput
	// Enable/disable REST API GET request logging. Valid values: `enable`, `disable`.
	RestApiGet pulumi.StringPtrInput
	// Enable/disable REST API POST/PUT/DELETE request logging. Valid values: `enable`, `disable`.
	RestApiSet pulumi.StringPtrInput
	// Enable/disable override Syslog settings. Valid values: `enable`, `disable`.
	SyslogOverride pulumi.StringPtrInput
	// Enable/disable anonymizing user names in log messages. Valid values: `enable`, `disable`.
	UserAnonymize pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*settingArgs)(nil)).Elem()
}

type SettingInput interface {
	pulumi.Input

	ToSettingOutput() SettingOutput
	ToSettingOutputWithContext(ctx context.Context) SettingOutput
}

func (*Setting) ElementType() reflect.Type {
	return reflect.TypeOf((**Setting)(nil)).Elem()
}

func (i *Setting) ToSettingOutput() SettingOutput {
	return i.ToSettingOutputWithContext(context.Background())
}

func (i *Setting) ToSettingOutputWithContext(ctx context.Context) SettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingOutput)
}

// SettingArrayInput is an input type that accepts SettingArray and SettingArrayOutput values.
// You can construct a concrete instance of `SettingArrayInput` via:
//
//	SettingArray{ SettingArgs{...} }
type SettingArrayInput interface {
	pulumi.Input

	ToSettingArrayOutput() SettingArrayOutput
	ToSettingArrayOutputWithContext(context.Context) SettingArrayOutput
}

type SettingArray []SettingInput

func (SettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Setting)(nil)).Elem()
}

func (i SettingArray) ToSettingArrayOutput() SettingArrayOutput {
	return i.ToSettingArrayOutputWithContext(context.Background())
}

func (i SettingArray) ToSettingArrayOutputWithContext(ctx context.Context) SettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingArrayOutput)
}

// SettingMapInput is an input type that accepts SettingMap and SettingMapOutput values.
// You can construct a concrete instance of `SettingMapInput` via:
//
//	SettingMap{ "key": SettingArgs{...} }
type SettingMapInput interface {
	pulumi.Input

	ToSettingMapOutput() SettingMapOutput
	ToSettingMapOutputWithContext(context.Context) SettingMapOutput
}

type SettingMap map[string]SettingInput

func (SettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Setting)(nil)).Elem()
}

func (i SettingMap) ToSettingMapOutput() SettingMapOutput {
	return i.ToSettingMapOutputWithContext(context.Background())
}

func (i SettingMap) ToSettingMapOutputWithContext(ctx context.Context) SettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingMapOutput)
}

type SettingOutput struct{ *pulumi.OutputState }

func (SettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Setting)(nil)).Elem()
}

func (o SettingOutput) ToSettingOutput() SettingOutput {
	return o
}

func (o SettingOutput) ToSettingOutputWithContext(ctx context.Context) SettingOutput {
	return o
}

// User name anonymization hash salt.
func (o SettingOutput) AnonymizationHash() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.AnonymizationHash }).(pulumi.StringOutput)
}

// Enable/disable brief format traffic logging. Valid values: `enable`, `disable`.
func (o SettingOutput) BriefTrafficFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.BriefTrafficFormat }).(pulumi.StringOutput)
}

// Custom fields to append to all log messages. The structure of `customLogFields` block is documented below.
func (o SettingOutput) CustomLogFields() SettingCustomLogFieldArrayOutput {
	return o.ApplyT(func(v *Setting) SettingCustomLogFieldArrayOutput { return v.CustomLogFields }).(SettingCustomLogFieldArrayOutput)
}

// Enable/disable daemon logging. Valid values: `enable`, `disable`.
func (o SettingOutput) DaemonLog() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.DaemonLog }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SettingOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Enable/disable explicit proxy firewall implicit policy logging. Valid values: `enable`, `disable`.
func (o SettingOutput) ExpolicyImplicitLog() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.ExpolicyImplicitLog }).(pulumi.StringOutput)
}

// Enable/disable extended traffic logging. Valid values: `enable`, `disable`.
func (o SettingOutput) ExtendedLog() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.ExtendedLog }).(pulumi.StringOutput)
}

// Enable/disable override FortiAnalyzer settings. Valid values: `enable`, `disable`.
func (o SettingOutput) FazOverride() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.FazOverride }).(pulumi.StringOutput)
}

// Enable/disable implicit firewall policy6 logging. Valid values: `enable`, `disable`.
func (o SettingOutput) Fwpolicy6ImplicitLog() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.Fwpolicy6ImplicitLog }).(pulumi.StringOutput)
}

// Enable/disable implicit firewall policy logging. Valid values: `enable`, `disable`.
func (o SettingOutput) FwpolicyImplicitLog() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.FwpolicyImplicitLog }).(pulumi.StringOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o SettingOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Enable/disable local-in-allow logging. Valid values: `enable`, `disable`.
func (o SettingOutput) LocalInAllow() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.LocalInAllow }).(pulumi.StringOutput)
}

// Enable/disable local-in-deny-broadcast logging. Valid values: `enable`, `disable`.
func (o SettingOutput) LocalInDenyBroadcast() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.LocalInDenyBroadcast }).(pulumi.StringOutput)
}

// Enable/disable local-in-deny-unicast logging. Valid values: `enable`, `disable`.
func (o SettingOutput) LocalInDenyUnicast() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.LocalInDenyUnicast }).(pulumi.StringOutput)
}

// Enable/disable local-out logging. Valid values: `enable`, `disable`.
func (o SettingOutput) LocalOut() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.LocalOut }).(pulumi.StringOutput)
}

// Enable/disable local-out traffic IoC detection. Requires local-out to be enabled. Valid values: `enable`, `disable`.
func (o SettingOutput) LocalOutIocDetection() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.LocalOutIocDetection }).(pulumi.StringOutput)
}

// Enable/disable invalid packet traffic logging. Valid values: `enable`, `disable`.
func (o SettingOutput) LogInvalidPacket() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.LogInvalidPacket }).(pulumi.StringOutput)
}

// Enable/disable inserting policy comments into traffic logs. Valid values: `enable`, `disable`.
func (o SettingOutput) LogPolicyComment() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.LogPolicyComment }).(pulumi.StringOutput)
}

// Enable/disable inserting policy name into traffic logs. Valid values: `enable`, `disable`.
func (o SettingOutput) LogPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.LogPolicyName }).(pulumi.StringOutput)
}

// Enable/disable logs with user-in-upper. Valid values: `enable`, `disable`.
func (o SettingOutput) LogUserInUpper() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.LogUserInUpper }).(pulumi.StringOutput)
}

// Enable/disable long-live-session statistics logging. Valid values: `enable`, `disable`.
func (o SettingOutput) LongLiveSessionStat() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.LongLiveSessionStat }).(pulumi.StringOutput)
}

// Enable/disable neighbor event logging. Valid values: `enable`, `disable`.
func (o SettingOutput) NeighborEvent() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.NeighborEvent }).(pulumi.StringOutput)
}

// Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `enable`, `disable`.
func (o SettingOutput) ResolveIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.ResolveIp }).(pulumi.StringOutput)
}

// Enable/disable adding resolved service names to traffic logs. Valid values: `enable`, `disable`.
func (o SettingOutput) ResolvePort() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.ResolvePort }).(pulumi.StringOutput)
}

// Enable/disable REST API GET request logging. Valid values: `enable`, `disable`.
func (o SettingOutput) RestApiGet() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.RestApiGet }).(pulumi.StringOutput)
}

// Enable/disable REST API POST/PUT/DELETE request logging. Valid values: `enable`, `disable`.
func (o SettingOutput) RestApiSet() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.RestApiSet }).(pulumi.StringOutput)
}

// Enable/disable override Syslog settings. Valid values: `enable`, `disable`.
func (o SettingOutput) SyslogOverride() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.SyslogOverride }).(pulumi.StringOutput)
}

// Enable/disable anonymizing user names in log messages. Valid values: `enable`, `disable`.
func (o SettingOutput) UserAnonymize() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.UserAnonymize }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SettingOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type SettingArrayOutput struct{ *pulumi.OutputState }

func (SettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Setting)(nil)).Elem()
}

func (o SettingArrayOutput) ToSettingArrayOutput() SettingArrayOutput {
	return o
}

func (o SettingArrayOutput) ToSettingArrayOutputWithContext(ctx context.Context) SettingArrayOutput {
	return o
}

func (o SettingArrayOutput) Index(i pulumi.IntInput) SettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Setting {
		return vs[0].([]*Setting)[vs[1].(int)]
	}).(SettingOutput)
}

type SettingMapOutput struct{ *pulumi.OutputState }

func (SettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Setting)(nil)).Elem()
}

func (o SettingMapOutput) ToSettingMapOutput() SettingMapOutput {
	return o
}

func (o SettingMapOutput) ToSettingMapOutputWithContext(ctx context.Context) SettingMapOutput {
	return o
}

func (o SettingMapOutput) MapIndex(k pulumi.StringInput) SettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Setting {
		return vs[0].(map[string]*Setting)[vs[1].(string)]
	}).(SettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SettingInput)(nil)).Elem(), &Setting{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingArrayInput)(nil)).Elem(), SettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingMapInput)(nil)).Elem(), SettingMap{})
	pulumi.RegisterOutputType(SettingOutput{})
	pulumi.RegisterOutputType(SettingArrayOutput{})
	pulumi.RegisterOutputType(SettingMapOutput{})
}
