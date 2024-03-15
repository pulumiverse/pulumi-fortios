// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure FortiSwitch global settings.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/switchcontroller"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := switchcontroller.NewGlobal(ctx, "trname", &switchcontroller.GlobalArgs{
//				AllowMultipleInterfaces: pulumi.String("disable"),
//				HttpsImagePush:          pulumi.String("disable"),
//				LogMacLimitViolations:   pulumi.String("disable"),
//				MacAgingInterval:        pulumi.Int(332),
//				MacRetentionPeriod:      pulumi.Int(24),
//				MacViolationTimer:       pulumi.Int(0),
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
// SwitchController Global can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/global:Global labelname SwitchControllerGlobal
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/global:Global labelname SwitchControllerGlobal
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Global struct {
	pulumi.CustomResourceState

	// Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. Valid values: `enable`, `disable`.
	AllowMultipleInterfaces pulumi.StringOutput `pulumi:"allowMultipleInterfaces"`
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceQuarantinedLink pulumi.StringOutput `pulumi:"bounceQuarantinedLink"`
	// List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `customCommand` block is documented below.
	CustomCommands GlobalCustomCommandArrayOutput `pulumi:"customCommands"`
	// Default VLAN for ports when added to the virtual-switch.
	DefaultVirtualSwitchVlan pulumi.StringOutput `pulumi:"defaultVirtualSwitchVlan"`
	// Enable/disable DHCP snooping server access list. Valid values: `enable`, `disable`.
	DhcpServerAccessList pulumi.StringOutput `pulumi:"dhcpServerAccessList"`
	// Prevent this FortiSwitch from discovering. The structure of `disableDiscovery` block is documented below.
	DisableDiscoveries GlobalDisableDiscoveryArrayOutput `pulumi:"disableDiscoveries"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable enforcement of FIPS on managed FortiSwitch devices. Valid values: `disable`, `enable`.
	FipsEnforce pulumi.StringOutput `pulumi:"fipsEnforce"`
	// Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
	FirmwareProvisionOnAuthorization pulumi.StringOutput `pulumi:"firmwareProvisionOnAuthorization"`
	// Enable/disable image push to FortiSwitch using HTTPS. Valid values: `enable`, `disable`.
	HttpsImagePush pulumi.StringOutput `pulumi:"httpsImagePush"`
	// Enable/disable logs for Learning Limit Violations. Valid values: `enable`, `disable`.
	LogMacLimitViolations pulumi.StringOutput `pulumi:"logMacLimitViolations"`
	// Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
	MacAgingInterval pulumi.IntOutput `pulumi:"macAgingInterval"`
	// Enable/disable MAC address event logging. Valid values: `enable`, `disable`.
	MacEventLogging pulumi.StringOutput `pulumi:"macEventLogging"`
	// Time in hours after which an inactive MAC is removed from client DB.
	MacRetentionPeriod pulumi.IntOutput `pulumi:"macRetentionPeriod"`
	// Set timeout for Learning Limit Violations (0 = disabled).
	MacViolationTimer pulumi.IntOutput `pulumi:"macViolationTimer"`
	// Quarantine mode. Valid values: `by-vlan`, `by-redirect`.
	QuarantineMode pulumi.StringOutput `pulumi:"quarantineMode"`
	// Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number. Valid values: `enable`, `disable`.
	SnDnsResolution pulumi.StringOutput `pulumi:"snDnsResolution"`
	// Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.
	UpdateUserDevice pulumi.StringOutput `pulumi:"updateUserDevice"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `all`, `defined`.
	VlanAllMode pulumi.StringOutput `pulumi:"vlanAllMode"`
	// FortiLink VLAN optimization. Valid values: `enable`, `disable`.
	VlanOptimization pulumi.StringOutput `pulumi:"vlanOptimization"`
}

// NewGlobal registers a new resource with the given unique name, arguments, and options.
func NewGlobal(ctx *pulumi.Context,
	name string, args *GlobalArgs, opts ...pulumi.ResourceOption) (*Global, error) {
	if args == nil {
		args = &GlobalArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Global
	err := ctx.RegisterResource("fortios:switchcontroller/global:Global", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobal gets an existing Global resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalState, opts ...pulumi.ResourceOption) (*Global, error) {
	var resource Global
	err := ctx.ReadResource("fortios:switchcontroller/global:Global", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Global resources.
type globalState struct {
	// Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. Valid values: `enable`, `disable`.
	AllowMultipleInterfaces *string `pulumi:"allowMultipleInterfaces"`
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceQuarantinedLink *string `pulumi:"bounceQuarantinedLink"`
	// List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `customCommand` block is documented below.
	CustomCommands []GlobalCustomCommand `pulumi:"customCommands"`
	// Default VLAN for ports when added to the virtual-switch.
	DefaultVirtualSwitchVlan *string `pulumi:"defaultVirtualSwitchVlan"`
	// Enable/disable DHCP snooping server access list. Valid values: `enable`, `disable`.
	DhcpServerAccessList *string `pulumi:"dhcpServerAccessList"`
	// Prevent this FortiSwitch from discovering. The structure of `disableDiscovery` block is documented below.
	DisableDiscoveries []GlobalDisableDiscovery `pulumi:"disableDiscoveries"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable enforcement of FIPS on managed FortiSwitch devices. Valid values: `disable`, `enable`.
	FipsEnforce *string `pulumi:"fipsEnforce"`
	// Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
	FirmwareProvisionOnAuthorization *string `pulumi:"firmwareProvisionOnAuthorization"`
	// Enable/disable image push to FortiSwitch using HTTPS. Valid values: `enable`, `disable`.
	HttpsImagePush *string `pulumi:"httpsImagePush"`
	// Enable/disable logs for Learning Limit Violations. Valid values: `enable`, `disable`.
	LogMacLimitViolations *string `pulumi:"logMacLimitViolations"`
	// Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
	MacAgingInterval *int `pulumi:"macAgingInterval"`
	// Enable/disable MAC address event logging. Valid values: `enable`, `disable`.
	MacEventLogging *string `pulumi:"macEventLogging"`
	// Time in hours after which an inactive MAC is removed from client DB.
	MacRetentionPeriod *int `pulumi:"macRetentionPeriod"`
	// Set timeout for Learning Limit Violations (0 = disabled).
	MacViolationTimer *int `pulumi:"macViolationTimer"`
	// Quarantine mode. Valid values: `by-vlan`, `by-redirect`.
	QuarantineMode *string `pulumi:"quarantineMode"`
	// Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number. Valid values: `enable`, `disable`.
	SnDnsResolution *string `pulumi:"snDnsResolution"`
	// Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.
	UpdateUserDevice *string `pulumi:"updateUserDevice"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `all`, `defined`.
	VlanAllMode *string `pulumi:"vlanAllMode"`
	// FortiLink VLAN optimization. Valid values: `enable`, `disable`.
	VlanOptimization *string `pulumi:"vlanOptimization"`
}

type GlobalState struct {
	// Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. Valid values: `enable`, `disable`.
	AllowMultipleInterfaces pulumi.StringPtrInput
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceQuarantinedLink pulumi.StringPtrInput
	// List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `customCommand` block is documented below.
	CustomCommands GlobalCustomCommandArrayInput
	// Default VLAN for ports when added to the virtual-switch.
	DefaultVirtualSwitchVlan pulumi.StringPtrInput
	// Enable/disable DHCP snooping server access list. Valid values: `enable`, `disable`.
	DhcpServerAccessList pulumi.StringPtrInput
	// Prevent this FortiSwitch from discovering. The structure of `disableDiscovery` block is documented below.
	DisableDiscoveries GlobalDisableDiscoveryArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable enforcement of FIPS on managed FortiSwitch devices. Valid values: `disable`, `enable`.
	FipsEnforce pulumi.StringPtrInput
	// Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
	FirmwareProvisionOnAuthorization pulumi.StringPtrInput
	// Enable/disable image push to FortiSwitch using HTTPS. Valid values: `enable`, `disable`.
	HttpsImagePush pulumi.StringPtrInput
	// Enable/disable logs for Learning Limit Violations. Valid values: `enable`, `disable`.
	LogMacLimitViolations pulumi.StringPtrInput
	// Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
	MacAgingInterval pulumi.IntPtrInput
	// Enable/disable MAC address event logging. Valid values: `enable`, `disable`.
	MacEventLogging pulumi.StringPtrInput
	// Time in hours after which an inactive MAC is removed from client DB.
	MacRetentionPeriod pulumi.IntPtrInput
	// Set timeout for Learning Limit Violations (0 = disabled).
	MacViolationTimer pulumi.IntPtrInput
	// Quarantine mode. Valid values: `by-vlan`, `by-redirect`.
	QuarantineMode pulumi.StringPtrInput
	// Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number. Valid values: `enable`, `disable`.
	SnDnsResolution pulumi.StringPtrInput
	// Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.
	UpdateUserDevice pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `all`, `defined`.
	VlanAllMode pulumi.StringPtrInput
	// FortiLink VLAN optimization. Valid values: `enable`, `disable`.
	VlanOptimization pulumi.StringPtrInput
}

func (GlobalState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalState)(nil)).Elem()
}

type globalArgs struct {
	// Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. Valid values: `enable`, `disable`.
	AllowMultipleInterfaces *string `pulumi:"allowMultipleInterfaces"`
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceQuarantinedLink *string `pulumi:"bounceQuarantinedLink"`
	// List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `customCommand` block is documented below.
	CustomCommands []GlobalCustomCommand `pulumi:"customCommands"`
	// Default VLAN for ports when added to the virtual-switch.
	DefaultVirtualSwitchVlan *string `pulumi:"defaultVirtualSwitchVlan"`
	// Enable/disable DHCP snooping server access list. Valid values: `enable`, `disable`.
	DhcpServerAccessList *string `pulumi:"dhcpServerAccessList"`
	// Prevent this FortiSwitch from discovering. The structure of `disableDiscovery` block is documented below.
	DisableDiscoveries []GlobalDisableDiscovery `pulumi:"disableDiscoveries"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable enforcement of FIPS on managed FortiSwitch devices. Valid values: `disable`, `enable`.
	FipsEnforce *string `pulumi:"fipsEnforce"`
	// Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
	FirmwareProvisionOnAuthorization *string `pulumi:"firmwareProvisionOnAuthorization"`
	// Enable/disable image push to FortiSwitch using HTTPS. Valid values: `enable`, `disable`.
	HttpsImagePush *string `pulumi:"httpsImagePush"`
	// Enable/disable logs for Learning Limit Violations. Valid values: `enable`, `disable`.
	LogMacLimitViolations *string `pulumi:"logMacLimitViolations"`
	// Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
	MacAgingInterval *int `pulumi:"macAgingInterval"`
	// Enable/disable MAC address event logging. Valid values: `enable`, `disable`.
	MacEventLogging *string `pulumi:"macEventLogging"`
	// Time in hours after which an inactive MAC is removed from client DB.
	MacRetentionPeriod *int `pulumi:"macRetentionPeriod"`
	// Set timeout for Learning Limit Violations (0 = disabled).
	MacViolationTimer *int `pulumi:"macViolationTimer"`
	// Quarantine mode. Valid values: `by-vlan`, `by-redirect`.
	QuarantineMode *string `pulumi:"quarantineMode"`
	// Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number. Valid values: `enable`, `disable`.
	SnDnsResolution *string `pulumi:"snDnsResolution"`
	// Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.
	UpdateUserDevice *string `pulumi:"updateUserDevice"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `all`, `defined`.
	VlanAllMode *string `pulumi:"vlanAllMode"`
	// FortiLink VLAN optimization. Valid values: `enable`, `disable`.
	VlanOptimization *string `pulumi:"vlanOptimization"`
}

// The set of arguments for constructing a Global resource.
type GlobalArgs struct {
	// Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. Valid values: `enable`, `disable`.
	AllowMultipleInterfaces pulumi.StringPtrInput
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceQuarantinedLink pulumi.StringPtrInput
	// List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `customCommand` block is documented below.
	CustomCommands GlobalCustomCommandArrayInput
	// Default VLAN for ports when added to the virtual-switch.
	DefaultVirtualSwitchVlan pulumi.StringPtrInput
	// Enable/disable DHCP snooping server access list. Valid values: `enable`, `disable`.
	DhcpServerAccessList pulumi.StringPtrInput
	// Prevent this FortiSwitch from discovering. The structure of `disableDiscovery` block is documented below.
	DisableDiscoveries GlobalDisableDiscoveryArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable enforcement of FIPS on managed FortiSwitch devices. Valid values: `disable`, `enable`.
	FipsEnforce pulumi.StringPtrInput
	// Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
	FirmwareProvisionOnAuthorization pulumi.StringPtrInput
	// Enable/disable image push to FortiSwitch using HTTPS. Valid values: `enable`, `disable`.
	HttpsImagePush pulumi.StringPtrInput
	// Enable/disable logs for Learning Limit Violations. Valid values: `enable`, `disable`.
	LogMacLimitViolations pulumi.StringPtrInput
	// Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
	MacAgingInterval pulumi.IntPtrInput
	// Enable/disable MAC address event logging. Valid values: `enable`, `disable`.
	MacEventLogging pulumi.StringPtrInput
	// Time in hours after which an inactive MAC is removed from client DB.
	MacRetentionPeriod pulumi.IntPtrInput
	// Set timeout for Learning Limit Violations (0 = disabled).
	MacViolationTimer pulumi.IntPtrInput
	// Quarantine mode. Valid values: `by-vlan`, `by-redirect`.
	QuarantineMode pulumi.StringPtrInput
	// Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number. Valid values: `enable`, `disable`.
	SnDnsResolution pulumi.StringPtrInput
	// Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.
	UpdateUserDevice pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `all`, `defined`.
	VlanAllMode pulumi.StringPtrInput
	// FortiLink VLAN optimization. Valid values: `enable`, `disable`.
	VlanOptimization pulumi.StringPtrInput
}

func (GlobalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalArgs)(nil)).Elem()
}

type GlobalInput interface {
	pulumi.Input

	ToGlobalOutput() GlobalOutput
	ToGlobalOutputWithContext(ctx context.Context) GlobalOutput
}

func (*Global) ElementType() reflect.Type {
	return reflect.TypeOf((**Global)(nil)).Elem()
}

func (i *Global) ToGlobalOutput() GlobalOutput {
	return i.ToGlobalOutputWithContext(context.Background())
}

func (i *Global) ToGlobalOutputWithContext(ctx context.Context) GlobalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalOutput)
}

// GlobalArrayInput is an input type that accepts GlobalArray and GlobalArrayOutput values.
// You can construct a concrete instance of `GlobalArrayInput` via:
//
//	GlobalArray{ GlobalArgs{...} }
type GlobalArrayInput interface {
	pulumi.Input

	ToGlobalArrayOutput() GlobalArrayOutput
	ToGlobalArrayOutputWithContext(context.Context) GlobalArrayOutput
}

type GlobalArray []GlobalInput

func (GlobalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Global)(nil)).Elem()
}

func (i GlobalArray) ToGlobalArrayOutput() GlobalArrayOutput {
	return i.ToGlobalArrayOutputWithContext(context.Background())
}

func (i GlobalArray) ToGlobalArrayOutputWithContext(ctx context.Context) GlobalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalArrayOutput)
}

// GlobalMapInput is an input type that accepts GlobalMap and GlobalMapOutput values.
// You can construct a concrete instance of `GlobalMapInput` via:
//
//	GlobalMap{ "key": GlobalArgs{...} }
type GlobalMapInput interface {
	pulumi.Input

	ToGlobalMapOutput() GlobalMapOutput
	ToGlobalMapOutputWithContext(context.Context) GlobalMapOutput
}

type GlobalMap map[string]GlobalInput

func (GlobalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Global)(nil)).Elem()
}

func (i GlobalMap) ToGlobalMapOutput() GlobalMapOutput {
	return i.ToGlobalMapOutputWithContext(context.Background())
}

func (i GlobalMap) ToGlobalMapOutputWithContext(ctx context.Context) GlobalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalMapOutput)
}

type GlobalOutput struct{ *pulumi.OutputState }

func (GlobalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Global)(nil)).Elem()
}

func (o GlobalOutput) ToGlobalOutput() GlobalOutput {
	return o
}

func (o GlobalOutput) ToGlobalOutputWithContext(ctx context.Context) GlobalOutput {
	return o
}

// Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. Valid values: `enable`, `disable`.
func (o GlobalOutput) AllowMultipleInterfaces() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.AllowMultipleInterfaces }).(pulumi.StringOutput)
}

// Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
func (o GlobalOutput) BounceQuarantinedLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.BounceQuarantinedLink }).(pulumi.StringOutput)
}

// List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `customCommand` block is documented below.
func (o GlobalOutput) CustomCommands() GlobalCustomCommandArrayOutput {
	return o.ApplyT(func(v *Global) GlobalCustomCommandArrayOutput { return v.CustomCommands }).(GlobalCustomCommandArrayOutput)
}

// Default VLAN for ports when added to the virtual-switch.
func (o GlobalOutput) DefaultVirtualSwitchVlan() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.DefaultVirtualSwitchVlan }).(pulumi.StringOutput)
}

// Enable/disable DHCP snooping server access list. Valid values: `enable`, `disable`.
func (o GlobalOutput) DhcpServerAccessList() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.DhcpServerAccessList }).(pulumi.StringOutput)
}

// Prevent this FortiSwitch from discovering. The structure of `disableDiscovery` block is documented below.
func (o GlobalOutput) DisableDiscoveries() GlobalDisableDiscoveryArrayOutput {
	return o.ApplyT(func(v *Global) GlobalDisableDiscoveryArrayOutput { return v.DisableDiscoveries }).(GlobalDisableDiscoveryArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o GlobalOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Global) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Enable/disable enforcement of FIPS on managed FortiSwitch devices. Valid values: `disable`, `enable`.
func (o GlobalOutput) FipsEnforce() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.FipsEnforce }).(pulumi.StringOutput)
}

// Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
func (o GlobalOutput) FirmwareProvisionOnAuthorization() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.FirmwareProvisionOnAuthorization }).(pulumi.StringOutput)
}

// Enable/disable image push to FortiSwitch using HTTPS. Valid values: `enable`, `disable`.
func (o GlobalOutput) HttpsImagePush() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.HttpsImagePush }).(pulumi.StringOutput)
}

// Enable/disable logs for Learning Limit Violations. Valid values: `enable`, `disable`.
func (o GlobalOutput) LogMacLimitViolations() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.LogMacLimitViolations }).(pulumi.StringOutput)
}

// Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
func (o GlobalOutput) MacAgingInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Global) pulumi.IntOutput { return v.MacAgingInterval }).(pulumi.IntOutput)
}

// Enable/disable MAC address event logging. Valid values: `enable`, `disable`.
func (o GlobalOutput) MacEventLogging() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.MacEventLogging }).(pulumi.StringOutput)
}

// Time in hours after which an inactive MAC is removed from client DB.
func (o GlobalOutput) MacRetentionPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *Global) pulumi.IntOutput { return v.MacRetentionPeriod }).(pulumi.IntOutput)
}

// Set timeout for Learning Limit Violations (0 = disabled).
func (o GlobalOutput) MacViolationTimer() pulumi.IntOutput {
	return o.ApplyT(func(v *Global) pulumi.IntOutput { return v.MacViolationTimer }).(pulumi.IntOutput)
}

// Quarantine mode. Valid values: `by-vlan`, `by-redirect`.
func (o GlobalOutput) QuarantineMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.QuarantineMode }).(pulumi.StringOutput)
}

// Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number. Valid values: `enable`, `disable`.
func (o GlobalOutput) SnDnsResolution() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.SnDnsResolution }).(pulumi.StringOutput)
}

// Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.
func (o GlobalOutput) UpdateUserDevice() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.UpdateUserDevice }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o GlobalOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Global) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `all`, `defined`.
func (o GlobalOutput) VlanAllMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.VlanAllMode }).(pulumi.StringOutput)
}

// FortiLink VLAN optimization. Valid values: `enable`, `disable`.
func (o GlobalOutput) VlanOptimization() pulumi.StringOutput {
	return o.ApplyT(func(v *Global) pulumi.StringOutput { return v.VlanOptimization }).(pulumi.StringOutput)
}

type GlobalArrayOutput struct{ *pulumi.OutputState }

func (GlobalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Global)(nil)).Elem()
}

func (o GlobalArrayOutput) ToGlobalArrayOutput() GlobalArrayOutput {
	return o
}

func (o GlobalArrayOutput) ToGlobalArrayOutputWithContext(ctx context.Context) GlobalArrayOutput {
	return o
}

func (o GlobalArrayOutput) Index(i pulumi.IntInput) GlobalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Global {
		return vs[0].([]*Global)[vs[1].(int)]
	}).(GlobalOutput)
}

type GlobalMapOutput struct{ *pulumi.OutputState }

func (GlobalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Global)(nil)).Elem()
}

func (o GlobalMapOutput) ToGlobalMapOutput() GlobalMapOutput {
	return o
}

func (o GlobalMapOutput) ToGlobalMapOutputWithContext(ctx context.Context) GlobalMapOutput {
	return o
}

func (o GlobalMapOutput) MapIndex(k pulumi.StringInput) GlobalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Global {
		return vs[0].(map[string]*Global)[vs[1].(string)]
	}).(GlobalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalInput)(nil)).Elem(), &Global{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalArrayInput)(nil)).Elem(), GlobalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalMapInput)(nil)).Elem(), GlobalMap{})
	pulumi.RegisterOutputType(GlobalOutput{})
	pulumi.RegisterOutputType(GlobalArrayOutput{})
	pulumi.RegisterOutputType(GlobalMapOutput{})
}
