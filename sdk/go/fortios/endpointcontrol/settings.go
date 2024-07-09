// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package endpointcontrol

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure endpoint control settings. Applies to FortiOS Version `6.2.0,6.2.4,6.2.6,7.4.0,7.4.1,7.4.2`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/endpointcontrol"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := endpointcontrol.NewSettings(ctx, "trname", &endpointcontrol.SettingsArgs{
//				DownloadLocation:                  pulumi.String("fortiguard"),
//				ForticlientAvdbUpdateInterval:     pulumi.Int(8),
//				ForticlientDeregUnsupportedClient: pulumi.String("enable"),
//				ForticlientEmsRestApiCallTimeout:  pulumi.Int(5000),
//				ForticlientKeepaliveInterval:      pulumi.Int(60),
//				ForticlientOfflineGrace:           pulumi.String("disable"),
//				ForticlientOfflineGraceInterval:   pulumi.Int(120),
//				ForticlientRegKeyEnforce:          pulumi.String("disable"),
//				ForticlientRegTimeout:             pulumi.Int(7),
//				ForticlientSysUpdateInterval:      pulumi.Int(720),
//				ForticlientUserAvatar:             pulumi.String("enable"),
//				ForticlientWarningInterval:        pulumi.Int(1),
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
// EndpointControl Settings can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:endpointcontrol/settings:Settings labelname EndpointControlSettings
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:endpointcontrol/settings:Settings labelname EndpointControlSettings
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Settings struct {
	pulumi.CustomResourceState

	// Customized URL for downloading FortiClient.
	DownloadCustomLink pulumi.StringOutput `pulumi:"downloadCustomLink"`
	// FortiClient download location (FortiGuard or custom). Valid values: `fortiguard`, `custom`.
	DownloadLocation pulumi.StringOutput `pulumi:"downloadLocation"`
	// Period of time between FortiClient AntiVirus database updates (0 - 24 hours, default = 8).
	ForticlientAvdbUpdateInterval pulumi.IntOutput `pulumi:"forticlientAvdbUpdateInterval"`
	// Enable/disable deregistering unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDeregUnsupportedClient pulumi.StringOutput `pulumi:"forticlientDeregUnsupportedClient"`
	// Enable/disable disconnecting of unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDisconnectUnsupportedClient pulumi.StringOutput `pulumi:"forticlientDisconnectUnsupportedClient"`
	// FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
	ForticlientEmsRestApiCallTimeout pulumi.IntOutput `pulumi:"forticlientEmsRestApiCallTimeout"`
	// Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).
	ForticlientKeepaliveInterval pulumi.IntOutput `pulumi:"forticlientKeepaliveInterval"`
	// Enable/disable grace period for offline registered clients. Valid values: `enable`, `disable`.
	ForticlientOfflineGrace pulumi.StringOutput `pulumi:"forticlientOfflineGrace"`
	// Grace period for offline registered FortiClient (60 - 600 sec, default = 120).
	ForticlientOfflineGraceInterval pulumi.IntOutput `pulumi:"forticlientOfflineGraceInterval"`
	// FortiClient registration key.
	ForticlientRegKey pulumi.StringPtrOutput `pulumi:"forticlientRegKey"`
	// Enable/disable requiring or enforcing FortiClient registration keys. Valid values: `enable`, `disable`.
	ForticlientRegKeyEnforce pulumi.StringOutput `pulumi:"forticlientRegKeyEnforce"`
	// FortiClient registration license timeout (days, min = 1, max = 180, 0 means unlimited).
	ForticlientRegTimeout pulumi.IntOutput `pulumi:"forticlientRegTimeout"`
	// Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).
	ForticlientSysUpdateInterval pulumi.IntOutput `pulumi:"forticlientSysUpdateInterval"`
	// Enable/disable uploading FortiClient user avatars. Valid values: `enable`, `disable`.
	ForticlientUserAvatar pulumi.StringOutput `pulumi:"forticlientUserAvatar"`
	// Period of time between FortiClient portal warnings (0 - 24 hours, default = 1).
	ForticlientWarningInterval pulumi.IntOutput `pulumi:"forticlientWarningInterval"`
	// Override global EMS table for this VDOM. Valid values: `enable`, `disable`.
	Override pulumi.StringOutput `pulumi:"override"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSettings registers a new resource with the given unique name, arguments, and options.
func NewSettings(ctx *pulumi.Context,
	name string, args *SettingsArgs, opts ...pulumi.ResourceOption) (*Settings, error) {
	if args == nil {
		args = &SettingsArgs{}
	}

	if args.ForticlientRegKey != nil {
		args.ForticlientRegKey = pulumi.ToSecret(args.ForticlientRegKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"forticlientRegKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Settings
	err := ctx.RegisterResource("fortios:endpointcontrol/settings:Settings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSettings gets an existing Settings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SettingsState, opts ...pulumi.ResourceOption) (*Settings, error) {
	var resource Settings
	err := ctx.ReadResource("fortios:endpointcontrol/settings:Settings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Settings resources.
type settingsState struct {
	// Customized URL for downloading FortiClient.
	DownloadCustomLink *string `pulumi:"downloadCustomLink"`
	// FortiClient download location (FortiGuard or custom). Valid values: `fortiguard`, `custom`.
	DownloadLocation *string `pulumi:"downloadLocation"`
	// Period of time between FortiClient AntiVirus database updates (0 - 24 hours, default = 8).
	ForticlientAvdbUpdateInterval *int `pulumi:"forticlientAvdbUpdateInterval"`
	// Enable/disable deregistering unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDeregUnsupportedClient *string `pulumi:"forticlientDeregUnsupportedClient"`
	// Enable/disable disconnecting of unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDisconnectUnsupportedClient *string `pulumi:"forticlientDisconnectUnsupportedClient"`
	// FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
	ForticlientEmsRestApiCallTimeout *int `pulumi:"forticlientEmsRestApiCallTimeout"`
	// Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).
	ForticlientKeepaliveInterval *int `pulumi:"forticlientKeepaliveInterval"`
	// Enable/disable grace period for offline registered clients. Valid values: `enable`, `disable`.
	ForticlientOfflineGrace *string `pulumi:"forticlientOfflineGrace"`
	// Grace period for offline registered FortiClient (60 - 600 sec, default = 120).
	ForticlientOfflineGraceInterval *int `pulumi:"forticlientOfflineGraceInterval"`
	// FortiClient registration key.
	ForticlientRegKey *string `pulumi:"forticlientRegKey"`
	// Enable/disable requiring or enforcing FortiClient registration keys. Valid values: `enable`, `disable`.
	ForticlientRegKeyEnforce *string `pulumi:"forticlientRegKeyEnforce"`
	// FortiClient registration license timeout (days, min = 1, max = 180, 0 means unlimited).
	ForticlientRegTimeout *int `pulumi:"forticlientRegTimeout"`
	// Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).
	ForticlientSysUpdateInterval *int `pulumi:"forticlientSysUpdateInterval"`
	// Enable/disable uploading FortiClient user avatars. Valid values: `enable`, `disable`.
	ForticlientUserAvatar *string `pulumi:"forticlientUserAvatar"`
	// Period of time between FortiClient portal warnings (0 - 24 hours, default = 1).
	ForticlientWarningInterval *int `pulumi:"forticlientWarningInterval"`
	// Override global EMS table for this VDOM. Valid values: `enable`, `disable`.
	Override *string `pulumi:"override"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SettingsState struct {
	// Customized URL for downloading FortiClient.
	DownloadCustomLink pulumi.StringPtrInput
	// FortiClient download location (FortiGuard or custom). Valid values: `fortiguard`, `custom`.
	DownloadLocation pulumi.StringPtrInput
	// Period of time between FortiClient AntiVirus database updates (0 - 24 hours, default = 8).
	ForticlientAvdbUpdateInterval pulumi.IntPtrInput
	// Enable/disable deregistering unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDeregUnsupportedClient pulumi.StringPtrInput
	// Enable/disable disconnecting of unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDisconnectUnsupportedClient pulumi.StringPtrInput
	// FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
	ForticlientEmsRestApiCallTimeout pulumi.IntPtrInput
	// Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).
	ForticlientKeepaliveInterval pulumi.IntPtrInput
	// Enable/disable grace period for offline registered clients. Valid values: `enable`, `disable`.
	ForticlientOfflineGrace pulumi.StringPtrInput
	// Grace period for offline registered FortiClient (60 - 600 sec, default = 120).
	ForticlientOfflineGraceInterval pulumi.IntPtrInput
	// FortiClient registration key.
	ForticlientRegKey pulumi.StringPtrInput
	// Enable/disable requiring or enforcing FortiClient registration keys. Valid values: `enable`, `disable`.
	ForticlientRegKeyEnforce pulumi.StringPtrInput
	// FortiClient registration license timeout (days, min = 1, max = 180, 0 means unlimited).
	ForticlientRegTimeout pulumi.IntPtrInput
	// Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).
	ForticlientSysUpdateInterval pulumi.IntPtrInput
	// Enable/disable uploading FortiClient user avatars. Valid values: `enable`, `disable`.
	ForticlientUserAvatar pulumi.StringPtrInput
	// Period of time between FortiClient portal warnings (0 - 24 hours, default = 1).
	ForticlientWarningInterval pulumi.IntPtrInput
	// Override global EMS table for this VDOM. Valid values: `enable`, `disable`.
	Override pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*settingsState)(nil)).Elem()
}

type settingsArgs struct {
	// Customized URL for downloading FortiClient.
	DownloadCustomLink *string `pulumi:"downloadCustomLink"`
	// FortiClient download location (FortiGuard or custom). Valid values: `fortiguard`, `custom`.
	DownloadLocation *string `pulumi:"downloadLocation"`
	// Period of time between FortiClient AntiVirus database updates (0 - 24 hours, default = 8).
	ForticlientAvdbUpdateInterval *int `pulumi:"forticlientAvdbUpdateInterval"`
	// Enable/disable deregistering unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDeregUnsupportedClient *string `pulumi:"forticlientDeregUnsupportedClient"`
	// Enable/disable disconnecting of unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDisconnectUnsupportedClient *string `pulumi:"forticlientDisconnectUnsupportedClient"`
	// FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
	ForticlientEmsRestApiCallTimeout *int `pulumi:"forticlientEmsRestApiCallTimeout"`
	// Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).
	ForticlientKeepaliveInterval *int `pulumi:"forticlientKeepaliveInterval"`
	// Enable/disable grace period for offline registered clients. Valid values: `enable`, `disable`.
	ForticlientOfflineGrace *string `pulumi:"forticlientOfflineGrace"`
	// Grace period for offline registered FortiClient (60 - 600 sec, default = 120).
	ForticlientOfflineGraceInterval *int `pulumi:"forticlientOfflineGraceInterval"`
	// FortiClient registration key.
	ForticlientRegKey *string `pulumi:"forticlientRegKey"`
	// Enable/disable requiring or enforcing FortiClient registration keys. Valid values: `enable`, `disable`.
	ForticlientRegKeyEnforce *string `pulumi:"forticlientRegKeyEnforce"`
	// FortiClient registration license timeout (days, min = 1, max = 180, 0 means unlimited).
	ForticlientRegTimeout *int `pulumi:"forticlientRegTimeout"`
	// Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).
	ForticlientSysUpdateInterval *int `pulumi:"forticlientSysUpdateInterval"`
	// Enable/disable uploading FortiClient user avatars. Valid values: `enable`, `disable`.
	ForticlientUserAvatar *string `pulumi:"forticlientUserAvatar"`
	// Period of time between FortiClient portal warnings (0 - 24 hours, default = 1).
	ForticlientWarningInterval *int `pulumi:"forticlientWarningInterval"`
	// Override global EMS table for this VDOM. Valid values: `enable`, `disable`.
	Override *string `pulumi:"override"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Settings resource.
type SettingsArgs struct {
	// Customized URL for downloading FortiClient.
	DownloadCustomLink pulumi.StringPtrInput
	// FortiClient download location (FortiGuard or custom). Valid values: `fortiguard`, `custom`.
	DownloadLocation pulumi.StringPtrInput
	// Period of time between FortiClient AntiVirus database updates (0 - 24 hours, default = 8).
	ForticlientAvdbUpdateInterval pulumi.IntPtrInput
	// Enable/disable deregistering unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDeregUnsupportedClient pulumi.StringPtrInput
	// Enable/disable disconnecting of unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDisconnectUnsupportedClient pulumi.StringPtrInput
	// FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
	ForticlientEmsRestApiCallTimeout pulumi.IntPtrInput
	// Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).
	ForticlientKeepaliveInterval pulumi.IntPtrInput
	// Enable/disable grace period for offline registered clients. Valid values: `enable`, `disable`.
	ForticlientOfflineGrace pulumi.StringPtrInput
	// Grace period for offline registered FortiClient (60 - 600 sec, default = 120).
	ForticlientOfflineGraceInterval pulumi.IntPtrInput
	// FortiClient registration key.
	ForticlientRegKey pulumi.StringPtrInput
	// Enable/disable requiring or enforcing FortiClient registration keys. Valid values: `enable`, `disable`.
	ForticlientRegKeyEnforce pulumi.StringPtrInput
	// FortiClient registration license timeout (days, min = 1, max = 180, 0 means unlimited).
	ForticlientRegTimeout pulumi.IntPtrInput
	// Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).
	ForticlientSysUpdateInterval pulumi.IntPtrInput
	// Enable/disable uploading FortiClient user avatars. Valid values: `enable`, `disable`.
	ForticlientUserAvatar pulumi.StringPtrInput
	// Period of time between FortiClient portal warnings (0 - 24 hours, default = 1).
	ForticlientWarningInterval pulumi.IntPtrInput
	// Override global EMS table for this VDOM. Valid values: `enable`, `disable`.
	Override pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*settingsArgs)(nil)).Elem()
}

type SettingsInput interface {
	pulumi.Input

	ToSettingsOutput() SettingsOutput
	ToSettingsOutputWithContext(ctx context.Context) SettingsOutput
}

func (*Settings) ElementType() reflect.Type {
	return reflect.TypeOf((**Settings)(nil)).Elem()
}

func (i *Settings) ToSettingsOutput() SettingsOutput {
	return i.ToSettingsOutputWithContext(context.Background())
}

func (i *Settings) ToSettingsOutputWithContext(ctx context.Context) SettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsOutput)
}

// SettingsArrayInput is an input type that accepts SettingsArray and SettingsArrayOutput values.
// You can construct a concrete instance of `SettingsArrayInput` via:
//
//	SettingsArray{ SettingsArgs{...} }
type SettingsArrayInput interface {
	pulumi.Input

	ToSettingsArrayOutput() SettingsArrayOutput
	ToSettingsArrayOutputWithContext(context.Context) SettingsArrayOutput
}

type SettingsArray []SettingsInput

func (SettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Settings)(nil)).Elem()
}

func (i SettingsArray) ToSettingsArrayOutput() SettingsArrayOutput {
	return i.ToSettingsArrayOutputWithContext(context.Background())
}

func (i SettingsArray) ToSettingsArrayOutputWithContext(ctx context.Context) SettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsArrayOutput)
}

// SettingsMapInput is an input type that accepts SettingsMap and SettingsMapOutput values.
// You can construct a concrete instance of `SettingsMapInput` via:
//
//	SettingsMap{ "key": SettingsArgs{...} }
type SettingsMapInput interface {
	pulumi.Input

	ToSettingsMapOutput() SettingsMapOutput
	ToSettingsMapOutputWithContext(context.Context) SettingsMapOutput
}

type SettingsMap map[string]SettingsInput

func (SettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Settings)(nil)).Elem()
}

func (i SettingsMap) ToSettingsMapOutput() SettingsMapOutput {
	return i.ToSettingsMapOutputWithContext(context.Background())
}

func (i SettingsMap) ToSettingsMapOutputWithContext(ctx context.Context) SettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsMapOutput)
}

type SettingsOutput struct{ *pulumi.OutputState }

func (SettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Settings)(nil)).Elem()
}

func (o SettingsOutput) ToSettingsOutput() SettingsOutput {
	return o
}

func (o SettingsOutput) ToSettingsOutputWithContext(ctx context.Context) SettingsOutput {
	return o
}

// Customized URL for downloading FortiClient.
func (o SettingsOutput) DownloadCustomLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Settings) pulumi.StringOutput { return v.DownloadCustomLink }).(pulumi.StringOutput)
}

// FortiClient download location (FortiGuard or custom). Valid values: `fortiguard`, `custom`.
func (o SettingsOutput) DownloadLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *Settings) pulumi.StringOutput { return v.DownloadLocation }).(pulumi.StringOutput)
}

// Period of time between FortiClient AntiVirus database updates (0 - 24 hours, default = 8).
func (o SettingsOutput) ForticlientAvdbUpdateInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Settings) pulumi.IntOutput { return v.ForticlientAvdbUpdateInterval }).(pulumi.IntOutput)
}

// Enable/disable deregistering unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
func (o SettingsOutput) ForticlientDeregUnsupportedClient() pulumi.StringOutput {
	return o.ApplyT(func(v *Settings) pulumi.StringOutput { return v.ForticlientDeregUnsupportedClient }).(pulumi.StringOutput)
}

// Enable/disable disconnecting of unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
func (o SettingsOutput) ForticlientDisconnectUnsupportedClient() pulumi.StringOutput {
	return o.ApplyT(func(v *Settings) pulumi.StringOutput { return v.ForticlientDisconnectUnsupportedClient }).(pulumi.StringOutput)
}

// FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
func (o SettingsOutput) ForticlientEmsRestApiCallTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Settings) pulumi.IntOutput { return v.ForticlientEmsRestApiCallTimeout }).(pulumi.IntOutput)
}

// Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).
func (o SettingsOutput) ForticlientKeepaliveInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Settings) pulumi.IntOutput { return v.ForticlientKeepaliveInterval }).(pulumi.IntOutput)
}

// Enable/disable grace period for offline registered clients. Valid values: `enable`, `disable`.
func (o SettingsOutput) ForticlientOfflineGrace() pulumi.StringOutput {
	return o.ApplyT(func(v *Settings) pulumi.StringOutput { return v.ForticlientOfflineGrace }).(pulumi.StringOutput)
}

// Grace period for offline registered FortiClient (60 - 600 sec, default = 120).
func (o SettingsOutput) ForticlientOfflineGraceInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Settings) pulumi.IntOutput { return v.ForticlientOfflineGraceInterval }).(pulumi.IntOutput)
}

// FortiClient registration key.
func (o SettingsOutput) ForticlientRegKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Settings) pulumi.StringPtrOutput { return v.ForticlientRegKey }).(pulumi.StringPtrOutput)
}

// Enable/disable requiring or enforcing FortiClient registration keys. Valid values: `enable`, `disable`.
func (o SettingsOutput) ForticlientRegKeyEnforce() pulumi.StringOutput {
	return o.ApplyT(func(v *Settings) pulumi.StringOutput { return v.ForticlientRegKeyEnforce }).(pulumi.StringOutput)
}

// FortiClient registration license timeout (days, min = 1, max = 180, 0 means unlimited).
func (o SettingsOutput) ForticlientRegTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Settings) pulumi.IntOutput { return v.ForticlientRegTimeout }).(pulumi.IntOutput)
}

// Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).
func (o SettingsOutput) ForticlientSysUpdateInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Settings) pulumi.IntOutput { return v.ForticlientSysUpdateInterval }).(pulumi.IntOutput)
}

// Enable/disable uploading FortiClient user avatars. Valid values: `enable`, `disable`.
func (o SettingsOutput) ForticlientUserAvatar() pulumi.StringOutput {
	return o.ApplyT(func(v *Settings) pulumi.StringOutput { return v.ForticlientUserAvatar }).(pulumi.StringOutput)
}

// Period of time between FortiClient portal warnings (0 - 24 hours, default = 1).
func (o SettingsOutput) ForticlientWarningInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Settings) pulumi.IntOutput { return v.ForticlientWarningInterval }).(pulumi.IntOutput)
}

// Override global EMS table for this VDOM. Valid values: `enable`, `disable`.
func (o SettingsOutput) Override() pulumi.StringOutput {
	return o.ApplyT(func(v *Settings) pulumi.StringOutput { return v.Override }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SettingsOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Settings) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SettingsArrayOutput struct{ *pulumi.OutputState }

func (SettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Settings)(nil)).Elem()
}

func (o SettingsArrayOutput) ToSettingsArrayOutput() SettingsArrayOutput {
	return o
}

func (o SettingsArrayOutput) ToSettingsArrayOutputWithContext(ctx context.Context) SettingsArrayOutput {
	return o
}

func (o SettingsArrayOutput) Index(i pulumi.IntInput) SettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Settings {
		return vs[0].([]*Settings)[vs[1].(int)]
	}).(SettingsOutput)
}

type SettingsMapOutput struct{ *pulumi.OutputState }

func (SettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Settings)(nil)).Elem()
}

func (o SettingsMapOutput) ToSettingsMapOutput() SettingsMapOutput {
	return o
}

func (o SettingsMapOutput) ToSettingsMapOutputWithContext(ctx context.Context) SettingsMapOutput {
	return o
}

func (o SettingsMapOutput) MapIndex(k pulumi.StringInput) SettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Settings {
		return vs[0].(map[string]*Settings)[vs[1].(string)]
	}).(SettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SettingsInput)(nil)).Elem(), &Settings{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingsArrayInput)(nil)).Elem(), SettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingsMapInput)(nil)).Elem(), SettingsMap{})
	pulumi.RegisterOutputType(SettingsOutput{})
	pulumi.RegisterOutputType(SettingsArrayOutput{})
	pulumi.RegisterOutputType(SettingsMapOutput{})
}
