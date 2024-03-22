// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure FortiSwitch LLDP profiles.
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
//			_, err := switchcontroller.NewLldpprofile(ctx, "trname", &switchcontroller.LldpprofileArgs{
//				AutoIsl:               pulumi.String("enable"),
//				AutoIslHelloTimer:     pulumi.Int(3),
//				AutoIslPortGroup:      pulumi.Int(0),
//				AutoIslReceiveTimeout: pulumi.Int(60),
//				MedTlvs:               pulumi.String("inventory-management network-policy"),
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
// SwitchController LldpProfile can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/lldpprofile:Lldpprofile labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/lldpprofile:Lldpprofile labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Lldpprofile struct {
	pulumi.CustomResourceState

	// Enable/disable auto inter-switch LAG. Valid values: `disable`, `enable`.
	AutoIsl pulumi.StringOutput `pulumi:"autoIsl"`
	// Auto inter-switch LAG authentication mode. Valid values: `legacy`, `strict`, `relax`.
	AutoIslAuth pulumi.StringOutput `pulumi:"autoIslAuth"`
	// Auto inter-switch LAG encryption mode. Valid values: `none`, `mixed`, `must`.
	AutoIslAuthEncrypt pulumi.StringOutput `pulumi:"autoIslAuthEncrypt"`
	// Auto inter-switch LAG authentication identity.
	AutoIslAuthIdentity pulumi.StringOutput `pulumi:"autoIslAuthIdentity"`
	// Auto inter-switch LAG macsec profile for encryption.
	AutoIslAuthMacsecProfile pulumi.StringOutput `pulumi:"autoIslAuthMacsecProfile"`
	// Auto inter-switch LAG authentication reauth period in seconds(10 - 3600, default = 3600).
	AutoIslAuthReauth pulumi.IntOutput `pulumi:"autoIslAuthReauth"`
	// Auto inter-switch LAG authentication user certificate.
	AutoIslAuthUser pulumi.StringOutput `pulumi:"autoIslAuthUser"`
	// Auto inter-switch LAG hello timer duration (1 - 30 sec, default = 3).
	AutoIslHelloTimer pulumi.IntOutput `pulumi:"autoIslHelloTimer"`
	// Auto inter-switch LAG port group ID (0 - 9).
	AutoIslPortGroup pulumi.IntOutput `pulumi:"autoIslPortGroup"`
	// Auto inter-switch LAG timeout if no response is received (3 - 90 sec, default = 9).
	AutoIslReceiveTimeout pulumi.IntOutput `pulumi:"autoIslReceiveTimeout"`
	// Enable/disable MCLAG inter chassis link. Valid values: `disable`, `enable`.
	AutoMclagIcl pulumi.StringOutput `pulumi:"autoMclagIcl"`
	// Configuration method to edit custom TLV entries. The structure of `customTlvs` block is documented below.
	CustomTlvs LldpprofileCustomTlvArrayOutput `pulumi:"customTlvs"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Configuration method to edit Media Endpoint Discovery (MED) location service type-length-value (TLV) categories. The structure of `medLocationService` block is documented below.
	MedLocationServices LldpprofileMedLocationServiceArrayOutput `pulumi:"medLocationServices"`
	// Configuration method to edit Media Endpoint Discovery (MED) network policy type-length-value (TLV) categories. The structure of `medNetworkPolicy` block is documented below.
	MedNetworkPolicies LldpprofileMedNetworkPolicyArrayOutput `pulumi:"medNetworkPolicies"`
	// Transmitted LLDP-MED TLVs (type-length-value descriptions): inventory management TLV and/or network policy TLV.
	MedTlvs pulumi.StringOutput `pulumi:"medTlvs"`
	// Transmitted IEEE 802.1 TLVs. Valid values: `port-vlan-id`.
	N8021Tlvs pulumi.StringOutput `pulumi:"n8021Tlvs"`
	// Transmitted IEEE 802.3 TLVs.
	N8023Tlvs pulumi.StringOutput `pulumi:"n8023Tlvs"`
	// Profile name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewLldpprofile registers a new resource with the given unique name, arguments, and options.
func NewLldpprofile(ctx *pulumi.Context,
	name string, args *LldpprofileArgs, opts ...pulumi.ResourceOption) (*Lldpprofile, error) {
	if args == nil {
		args = &LldpprofileArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Lldpprofile
	err := ctx.RegisterResource("fortios:switchcontroller/lldpprofile:Lldpprofile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLldpprofile gets an existing Lldpprofile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLldpprofile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LldpprofileState, opts ...pulumi.ResourceOption) (*Lldpprofile, error) {
	var resource Lldpprofile
	err := ctx.ReadResource("fortios:switchcontroller/lldpprofile:Lldpprofile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Lldpprofile resources.
type lldpprofileState struct {
	// Enable/disable auto inter-switch LAG. Valid values: `disable`, `enable`.
	AutoIsl *string `pulumi:"autoIsl"`
	// Auto inter-switch LAG authentication mode. Valid values: `legacy`, `strict`, `relax`.
	AutoIslAuth *string `pulumi:"autoIslAuth"`
	// Auto inter-switch LAG encryption mode. Valid values: `none`, `mixed`, `must`.
	AutoIslAuthEncrypt *string `pulumi:"autoIslAuthEncrypt"`
	// Auto inter-switch LAG authentication identity.
	AutoIslAuthIdentity *string `pulumi:"autoIslAuthIdentity"`
	// Auto inter-switch LAG macsec profile for encryption.
	AutoIslAuthMacsecProfile *string `pulumi:"autoIslAuthMacsecProfile"`
	// Auto inter-switch LAG authentication reauth period in seconds(10 - 3600, default = 3600).
	AutoIslAuthReauth *int `pulumi:"autoIslAuthReauth"`
	// Auto inter-switch LAG authentication user certificate.
	AutoIslAuthUser *string `pulumi:"autoIslAuthUser"`
	// Auto inter-switch LAG hello timer duration (1 - 30 sec, default = 3).
	AutoIslHelloTimer *int `pulumi:"autoIslHelloTimer"`
	// Auto inter-switch LAG port group ID (0 - 9).
	AutoIslPortGroup *int `pulumi:"autoIslPortGroup"`
	// Auto inter-switch LAG timeout if no response is received (3 - 90 sec, default = 9).
	AutoIslReceiveTimeout *int `pulumi:"autoIslReceiveTimeout"`
	// Enable/disable MCLAG inter chassis link. Valid values: `disable`, `enable`.
	AutoMclagIcl *string `pulumi:"autoMclagIcl"`
	// Configuration method to edit custom TLV entries. The structure of `customTlvs` block is documented below.
	CustomTlvs []LldpprofileCustomTlv `pulumi:"customTlvs"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Configuration method to edit Media Endpoint Discovery (MED) location service type-length-value (TLV) categories. The structure of `medLocationService` block is documented below.
	MedLocationServices []LldpprofileMedLocationService `pulumi:"medLocationServices"`
	// Configuration method to edit Media Endpoint Discovery (MED) network policy type-length-value (TLV) categories. The structure of `medNetworkPolicy` block is documented below.
	MedNetworkPolicies []LldpprofileMedNetworkPolicy `pulumi:"medNetworkPolicies"`
	// Transmitted LLDP-MED TLVs (type-length-value descriptions): inventory management TLV and/or network policy TLV.
	MedTlvs *string `pulumi:"medTlvs"`
	// Transmitted IEEE 802.1 TLVs. Valid values: `port-vlan-id`.
	N8021Tlvs *string `pulumi:"n8021Tlvs"`
	// Transmitted IEEE 802.3 TLVs.
	N8023Tlvs *string `pulumi:"n8023Tlvs"`
	// Profile name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type LldpprofileState struct {
	// Enable/disable auto inter-switch LAG. Valid values: `disable`, `enable`.
	AutoIsl pulumi.StringPtrInput
	// Auto inter-switch LAG authentication mode. Valid values: `legacy`, `strict`, `relax`.
	AutoIslAuth pulumi.StringPtrInput
	// Auto inter-switch LAG encryption mode. Valid values: `none`, `mixed`, `must`.
	AutoIslAuthEncrypt pulumi.StringPtrInput
	// Auto inter-switch LAG authentication identity.
	AutoIslAuthIdentity pulumi.StringPtrInput
	// Auto inter-switch LAG macsec profile for encryption.
	AutoIslAuthMacsecProfile pulumi.StringPtrInput
	// Auto inter-switch LAG authentication reauth period in seconds(10 - 3600, default = 3600).
	AutoIslAuthReauth pulumi.IntPtrInput
	// Auto inter-switch LAG authentication user certificate.
	AutoIslAuthUser pulumi.StringPtrInput
	// Auto inter-switch LAG hello timer duration (1 - 30 sec, default = 3).
	AutoIslHelloTimer pulumi.IntPtrInput
	// Auto inter-switch LAG port group ID (0 - 9).
	AutoIslPortGroup pulumi.IntPtrInput
	// Auto inter-switch LAG timeout if no response is received (3 - 90 sec, default = 9).
	AutoIslReceiveTimeout pulumi.IntPtrInput
	// Enable/disable MCLAG inter chassis link. Valid values: `disable`, `enable`.
	AutoMclagIcl pulumi.StringPtrInput
	// Configuration method to edit custom TLV entries. The structure of `customTlvs` block is documented below.
	CustomTlvs LldpprofileCustomTlvArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Configuration method to edit Media Endpoint Discovery (MED) location service type-length-value (TLV) categories. The structure of `medLocationService` block is documented below.
	MedLocationServices LldpprofileMedLocationServiceArrayInput
	// Configuration method to edit Media Endpoint Discovery (MED) network policy type-length-value (TLV) categories. The structure of `medNetworkPolicy` block is documented below.
	MedNetworkPolicies LldpprofileMedNetworkPolicyArrayInput
	// Transmitted LLDP-MED TLVs (type-length-value descriptions): inventory management TLV and/or network policy TLV.
	MedTlvs pulumi.StringPtrInput
	// Transmitted IEEE 802.1 TLVs. Valid values: `port-vlan-id`.
	N8021Tlvs pulumi.StringPtrInput
	// Transmitted IEEE 802.3 TLVs.
	N8023Tlvs pulumi.StringPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LldpprofileState) ElementType() reflect.Type {
	return reflect.TypeOf((*lldpprofileState)(nil)).Elem()
}

type lldpprofileArgs struct {
	// Enable/disable auto inter-switch LAG. Valid values: `disable`, `enable`.
	AutoIsl *string `pulumi:"autoIsl"`
	// Auto inter-switch LAG authentication mode. Valid values: `legacy`, `strict`, `relax`.
	AutoIslAuth *string `pulumi:"autoIslAuth"`
	// Auto inter-switch LAG encryption mode. Valid values: `none`, `mixed`, `must`.
	AutoIslAuthEncrypt *string `pulumi:"autoIslAuthEncrypt"`
	// Auto inter-switch LAG authentication identity.
	AutoIslAuthIdentity *string `pulumi:"autoIslAuthIdentity"`
	// Auto inter-switch LAG macsec profile for encryption.
	AutoIslAuthMacsecProfile *string `pulumi:"autoIslAuthMacsecProfile"`
	// Auto inter-switch LAG authentication reauth period in seconds(10 - 3600, default = 3600).
	AutoIslAuthReauth *int `pulumi:"autoIslAuthReauth"`
	// Auto inter-switch LAG authentication user certificate.
	AutoIslAuthUser *string `pulumi:"autoIslAuthUser"`
	// Auto inter-switch LAG hello timer duration (1 - 30 sec, default = 3).
	AutoIslHelloTimer *int `pulumi:"autoIslHelloTimer"`
	// Auto inter-switch LAG port group ID (0 - 9).
	AutoIslPortGroup *int `pulumi:"autoIslPortGroup"`
	// Auto inter-switch LAG timeout if no response is received (3 - 90 sec, default = 9).
	AutoIslReceiveTimeout *int `pulumi:"autoIslReceiveTimeout"`
	// Enable/disable MCLAG inter chassis link. Valid values: `disable`, `enable`.
	AutoMclagIcl *string `pulumi:"autoMclagIcl"`
	// Configuration method to edit custom TLV entries. The structure of `customTlvs` block is documented below.
	CustomTlvs []LldpprofileCustomTlv `pulumi:"customTlvs"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Configuration method to edit Media Endpoint Discovery (MED) location service type-length-value (TLV) categories. The structure of `medLocationService` block is documented below.
	MedLocationServices []LldpprofileMedLocationService `pulumi:"medLocationServices"`
	// Configuration method to edit Media Endpoint Discovery (MED) network policy type-length-value (TLV) categories. The structure of `medNetworkPolicy` block is documented below.
	MedNetworkPolicies []LldpprofileMedNetworkPolicy `pulumi:"medNetworkPolicies"`
	// Transmitted LLDP-MED TLVs (type-length-value descriptions): inventory management TLV and/or network policy TLV.
	MedTlvs *string `pulumi:"medTlvs"`
	// Transmitted IEEE 802.1 TLVs. Valid values: `port-vlan-id`.
	N8021Tlvs *string `pulumi:"n8021Tlvs"`
	// Transmitted IEEE 802.3 TLVs.
	N8023Tlvs *string `pulumi:"n8023Tlvs"`
	// Profile name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Lldpprofile resource.
type LldpprofileArgs struct {
	// Enable/disable auto inter-switch LAG. Valid values: `disable`, `enable`.
	AutoIsl pulumi.StringPtrInput
	// Auto inter-switch LAG authentication mode. Valid values: `legacy`, `strict`, `relax`.
	AutoIslAuth pulumi.StringPtrInput
	// Auto inter-switch LAG encryption mode. Valid values: `none`, `mixed`, `must`.
	AutoIslAuthEncrypt pulumi.StringPtrInput
	// Auto inter-switch LAG authentication identity.
	AutoIslAuthIdentity pulumi.StringPtrInput
	// Auto inter-switch LAG macsec profile for encryption.
	AutoIslAuthMacsecProfile pulumi.StringPtrInput
	// Auto inter-switch LAG authentication reauth period in seconds(10 - 3600, default = 3600).
	AutoIslAuthReauth pulumi.IntPtrInput
	// Auto inter-switch LAG authentication user certificate.
	AutoIslAuthUser pulumi.StringPtrInput
	// Auto inter-switch LAG hello timer duration (1 - 30 sec, default = 3).
	AutoIslHelloTimer pulumi.IntPtrInput
	// Auto inter-switch LAG port group ID (0 - 9).
	AutoIslPortGroup pulumi.IntPtrInput
	// Auto inter-switch LAG timeout if no response is received (3 - 90 sec, default = 9).
	AutoIslReceiveTimeout pulumi.IntPtrInput
	// Enable/disable MCLAG inter chassis link. Valid values: `disable`, `enable`.
	AutoMclagIcl pulumi.StringPtrInput
	// Configuration method to edit custom TLV entries. The structure of `customTlvs` block is documented below.
	CustomTlvs LldpprofileCustomTlvArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Configuration method to edit Media Endpoint Discovery (MED) location service type-length-value (TLV) categories. The structure of `medLocationService` block is documented below.
	MedLocationServices LldpprofileMedLocationServiceArrayInput
	// Configuration method to edit Media Endpoint Discovery (MED) network policy type-length-value (TLV) categories. The structure of `medNetworkPolicy` block is documented below.
	MedNetworkPolicies LldpprofileMedNetworkPolicyArrayInput
	// Transmitted LLDP-MED TLVs (type-length-value descriptions): inventory management TLV and/or network policy TLV.
	MedTlvs pulumi.StringPtrInput
	// Transmitted IEEE 802.1 TLVs. Valid values: `port-vlan-id`.
	N8021Tlvs pulumi.StringPtrInput
	// Transmitted IEEE 802.3 TLVs.
	N8023Tlvs pulumi.StringPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LldpprofileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lldpprofileArgs)(nil)).Elem()
}

type LldpprofileInput interface {
	pulumi.Input

	ToLldpprofileOutput() LldpprofileOutput
	ToLldpprofileOutputWithContext(ctx context.Context) LldpprofileOutput
}

func (*Lldpprofile) ElementType() reflect.Type {
	return reflect.TypeOf((**Lldpprofile)(nil)).Elem()
}

func (i *Lldpprofile) ToLldpprofileOutput() LldpprofileOutput {
	return i.ToLldpprofileOutputWithContext(context.Background())
}

func (i *Lldpprofile) ToLldpprofileOutputWithContext(ctx context.Context) LldpprofileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LldpprofileOutput)
}

// LldpprofileArrayInput is an input type that accepts LldpprofileArray and LldpprofileArrayOutput values.
// You can construct a concrete instance of `LldpprofileArrayInput` via:
//
//	LldpprofileArray{ LldpprofileArgs{...} }
type LldpprofileArrayInput interface {
	pulumi.Input

	ToLldpprofileArrayOutput() LldpprofileArrayOutput
	ToLldpprofileArrayOutputWithContext(context.Context) LldpprofileArrayOutput
}

type LldpprofileArray []LldpprofileInput

func (LldpprofileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Lldpprofile)(nil)).Elem()
}

func (i LldpprofileArray) ToLldpprofileArrayOutput() LldpprofileArrayOutput {
	return i.ToLldpprofileArrayOutputWithContext(context.Background())
}

func (i LldpprofileArray) ToLldpprofileArrayOutputWithContext(ctx context.Context) LldpprofileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LldpprofileArrayOutput)
}

// LldpprofileMapInput is an input type that accepts LldpprofileMap and LldpprofileMapOutput values.
// You can construct a concrete instance of `LldpprofileMapInput` via:
//
//	LldpprofileMap{ "key": LldpprofileArgs{...} }
type LldpprofileMapInput interface {
	pulumi.Input

	ToLldpprofileMapOutput() LldpprofileMapOutput
	ToLldpprofileMapOutputWithContext(context.Context) LldpprofileMapOutput
}

type LldpprofileMap map[string]LldpprofileInput

func (LldpprofileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Lldpprofile)(nil)).Elem()
}

func (i LldpprofileMap) ToLldpprofileMapOutput() LldpprofileMapOutput {
	return i.ToLldpprofileMapOutputWithContext(context.Background())
}

func (i LldpprofileMap) ToLldpprofileMapOutputWithContext(ctx context.Context) LldpprofileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LldpprofileMapOutput)
}

type LldpprofileOutput struct{ *pulumi.OutputState }

func (LldpprofileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Lldpprofile)(nil)).Elem()
}

func (o LldpprofileOutput) ToLldpprofileOutput() LldpprofileOutput {
	return o
}

func (o LldpprofileOutput) ToLldpprofileOutputWithContext(ctx context.Context) LldpprofileOutput {
	return o
}

// Enable/disable auto inter-switch LAG. Valid values: `disable`, `enable`.
func (o LldpprofileOutput) AutoIsl() pulumi.StringOutput {
	return o.ApplyT(func(v *Lldpprofile) pulumi.StringOutput { return v.AutoIsl }).(pulumi.StringOutput)
}

// Auto inter-switch LAG authentication mode. Valid values: `legacy`, `strict`, `relax`.
func (o LldpprofileOutput) AutoIslAuth() pulumi.StringOutput {
	return o.ApplyT(func(v *Lldpprofile) pulumi.StringOutput { return v.AutoIslAuth }).(pulumi.StringOutput)
}

// Auto inter-switch LAG encryption mode. Valid values: `none`, `mixed`, `must`.
func (o LldpprofileOutput) AutoIslAuthEncrypt() pulumi.StringOutput {
	return o.ApplyT(func(v *Lldpprofile) pulumi.StringOutput { return v.AutoIslAuthEncrypt }).(pulumi.StringOutput)
}

// Auto inter-switch LAG authentication identity.
func (o LldpprofileOutput) AutoIslAuthIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v *Lldpprofile) pulumi.StringOutput { return v.AutoIslAuthIdentity }).(pulumi.StringOutput)
}

// Auto inter-switch LAG macsec profile for encryption.
func (o LldpprofileOutput) AutoIslAuthMacsecProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Lldpprofile) pulumi.StringOutput { return v.AutoIslAuthMacsecProfile }).(pulumi.StringOutput)
}

// Auto inter-switch LAG authentication reauth period in seconds(10 - 3600, default = 3600).
func (o LldpprofileOutput) AutoIslAuthReauth() pulumi.IntOutput {
	return o.ApplyT(func(v *Lldpprofile) pulumi.IntOutput { return v.AutoIslAuthReauth }).(pulumi.IntOutput)
}

// Auto inter-switch LAG authentication user certificate.
func (o LldpprofileOutput) AutoIslAuthUser() pulumi.StringOutput {
	return o.ApplyT(func(v *Lldpprofile) pulumi.StringOutput { return v.AutoIslAuthUser }).(pulumi.StringOutput)
}

// Auto inter-switch LAG hello timer duration (1 - 30 sec, default = 3).
func (o LldpprofileOutput) AutoIslHelloTimer() pulumi.IntOutput {
	return o.ApplyT(func(v *Lldpprofile) pulumi.IntOutput { return v.AutoIslHelloTimer }).(pulumi.IntOutput)
}

// Auto inter-switch LAG port group ID (0 - 9).
func (o LldpprofileOutput) AutoIslPortGroup() pulumi.IntOutput {
	return o.ApplyT(func(v *Lldpprofile) pulumi.IntOutput { return v.AutoIslPortGroup }).(pulumi.IntOutput)
}

// Auto inter-switch LAG timeout if no response is received (3 - 90 sec, default = 9).
func (o LldpprofileOutput) AutoIslReceiveTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Lldpprofile) pulumi.IntOutput { return v.AutoIslReceiveTimeout }).(pulumi.IntOutput)
}

// Enable/disable MCLAG inter chassis link. Valid values: `disable`, `enable`.
func (o LldpprofileOutput) AutoMclagIcl() pulumi.StringOutput {
	return o.ApplyT(func(v *Lldpprofile) pulumi.StringOutput { return v.AutoMclagIcl }).(pulumi.StringOutput)
}

// Configuration method to edit custom TLV entries. The structure of `customTlvs` block is documented below.
func (o LldpprofileOutput) CustomTlvs() LldpprofileCustomTlvArrayOutput {
	return o.ApplyT(func(v *Lldpprofile) LldpprofileCustomTlvArrayOutput { return v.CustomTlvs }).(LldpprofileCustomTlvArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o LldpprofileOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lldpprofile) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o LldpprofileOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lldpprofile) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Configuration method to edit Media Endpoint Discovery (MED) location service type-length-value (TLV) categories. The structure of `medLocationService` block is documented below.
func (o LldpprofileOutput) MedLocationServices() LldpprofileMedLocationServiceArrayOutput {
	return o.ApplyT(func(v *Lldpprofile) LldpprofileMedLocationServiceArrayOutput { return v.MedLocationServices }).(LldpprofileMedLocationServiceArrayOutput)
}

// Configuration method to edit Media Endpoint Discovery (MED) network policy type-length-value (TLV) categories. The structure of `medNetworkPolicy` block is documented below.
func (o LldpprofileOutput) MedNetworkPolicies() LldpprofileMedNetworkPolicyArrayOutput {
	return o.ApplyT(func(v *Lldpprofile) LldpprofileMedNetworkPolicyArrayOutput { return v.MedNetworkPolicies }).(LldpprofileMedNetworkPolicyArrayOutput)
}

// Transmitted LLDP-MED TLVs (type-length-value descriptions): inventory management TLV and/or network policy TLV.
func (o LldpprofileOutput) MedTlvs() pulumi.StringOutput {
	return o.ApplyT(func(v *Lldpprofile) pulumi.StringOutput { return v.MedTlvs }).(pulumi.StringOutput)
}

// Transmitted IEEE 802.1 TLVs. Valid values: `port-vlan-id`.
func (o LldpprofileOutput) N8021Tlvs() pulumi.StringOutput {
	return o.ApplyT(func(v *Lldpprofile) pulumi.StringOutput { return v.N8021Tlvs }).(pulumi.StringOutput)
}

// Transmitted IEEE 802.3 TLVs.
func (o LldpprofileOutput) N8023Tlvs() pulumi.StringOutput {
	return o.ApplyT(func(v *Lldpprofile) pulumi.StringOutput { return v.N8023Tlvs }).(pulumi.StringOutput)
}

// Profile name.
func (o LldpprofileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Lldpprofile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o LldpprofileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lldpprofile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type LldpprofileArrayOutput struct{ *pulumi.OutputState }

func (LldpprofileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Lldpprofile)(nil)).Elem()
}

func (o LldpprofileArrayOutput) ToLldpprofileArrayOutput() LldpprofileArrayOutput {
	return o
}

func (o LldpprofileArrayOutput) ToLldpprofileArrayOutputWithContext(ctx context.Context) LldpprofileArrayOutput {
	return o
}

func (o LldpprofileArrayOutput) Index(i pulumi.IntInput) LldpprofileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Lldpprofile {
		return vs[0].([]*Lldpprofile)[vs[1].(int)]
	}).(LldpprofileOutput)
}

type LldpprofileMapOutput struct{ *pulumi.OutputState }

func (LldpprofileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Lldpprofile)(nil)).Elem()
}

func (o LldpprofileMapOutput) ToLldpprofileMapOutput() LldpprofileMapOutput {
	return o
}

func (o LldpprofileMapOutput) ToLldpprofileMapOutputWithContext(ctx context.Context) LldpprofileMapOutput {
	return o
}

func (o LldpprofileMapOutput) MapIndex(k pulumi.StringInput) LldpprofileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Lldpprofile {
		return vs[0].(map[string]*Lldpprofile)[vs[1].(string)]
	}).(LldpprofileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LldpprofileInput)(nil)).Elem(), &Lldpprofile{})
	pulumi.RegisterInputType(reflect.TypeOf((*LldpprofileArrayInput)(nil)).Elem(), LldpprofileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LldpprofileMapInput)(nil)).Elem(), LldpprofileMap{})
	pulumi.RegisterOutputType(LldpprofileOutput{})
	pulumi.RegisterOutputType(LldpprofileArrayOutput{})
	pulumi.RegisterOutputType(LldpprofileMapOutput{})
}
