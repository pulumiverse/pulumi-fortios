// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wirelesscontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure inter wireless controller operation.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/wirelesscontroller"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := wirelesscontroller.NewIntercontroller(ctx, "trname", &wirelesscontroller.IntercontrollerArgs{
//				FastFailoverMax:     pulumi.Int(10),
//				FastFailoverWait:    pulumi.Int(10),
//				InterControllerKey:  pulumi.String("ENC XXXX"),
//				InterControllerMode: pulumi.String("disable"),
//				InterControllerPri:  pulumi.String("primary"),
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
// WirelessController InterController can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/intercontroller:Intercontroller labelname WirelessControllerInterController
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/intercontroller:Intercontroller labelname WirelessControllerInterController
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Intercontroller struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
	FastFailoverMax pulumi.IntOutput `pulumi:"fastFailoverMax"`
	// Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
	FastFailoverWait pulumi.IntOutput `pulumi:"fastFailoverWait"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Secret key for inter-controller communications.
	InterControllerKey pulumi.StringPtrOutput `pulumi:"interControllerKey"`
	// Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
	InterControllerMode pulumi.StringOutput `pulumi:"interControllerMode"`
	// Fast failover peer wireless controller list. The structure of `interControllerPeer` block is documented below.
	InterControllerPeers IntercontrollerInterControllerPeerArrayOutput `pulumi:"interControllerPeers"`
	// Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
	InterControllerPri pulumi.StringOutput `pulumi:"interControllerPri"`
	// Enable/disable layer 3 roaming (default = disable). Valid values: `enable`, `disable`.
	L3Roaming pulumi.StringOutput `pulumi:"l3Roaming"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewIntercontroller registers a new resource with the given unique name, arguments, and options.
func NewIntercontroller(ctx *pulumi.Context,
	name string, args *IntercontrollerArgs, opts ...pulumi.ResourceOption) (*Intercontroller, error) {
	if args == nil {
		args = &IntercontrollerArgs{}
	}

	if args.InterControllerKey != nil {
		args.InterControllerKey = pulumi.ToSecret(args.InterControllerKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"interControllerKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Intercontroller
	err := ctx.RegisterResource("fortios:wirelesscontroller/intercontroller:Intercontroller", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntercontroller gets an existing Intercontroller resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntercontroller(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntercontrollerState, opts ...pulumi.ResourceOption) (*Intercontroller, error) {
	var resource Intercontroller
	err := ctx.ReadResource("fortios:wirelesscontroller/intercontroller:Intercontroller", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Intercontroller resources.
type intercontrollerState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
	FastFailoverMax *int `pulumi:"fastFailoverMax"`
	// Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
	FastFailoverWait *int `pulumi:"fastFailoverWait"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Secret key for inter-controller communications.
	InterControllerKey *string `pulumi:"interControllerKey"`
	// Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
	InterControllerMode *string `pulumi:"interControllerMode"`
	// Fast failover peer wireless controller list. The structure of `interControllerPeer` block is documented below.
	InterControllerPeers []IntercontrollerInterControllerPeer `pulumi:"interControllerPeers"`
	// Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
	InterControllerPri *string `pulumi:"interControllerPri"`
	// Enable/disable layer 3 roaming (default = disable). Valid values: `enable`, `disable`.
	L3Roaming *string `pulumi:"l3Roaming"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type IntercontrollerState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
	FastFailoverMax pulumi.IntPtrInput
	// Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
	FastFailoverWait pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Secret key for inter-controller communications.
	InterControllerKey pulumi.StringPtrInput
	// Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
	InterControllerMode pulumi.StringPtrInput
	// Fast failover peer wireless controller list. The structure of `interControllerPeer` block is documented below.
	InterControllerPeers IntercontrollerInterControllerPeerArrayInput
	// Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
	InterControllerPri pulumi.StringPtrInput
	// Enable/disable layer 3 roaming (default = disable). Valid values: `enable`, `disable`.
	L3Roaming pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IntercontrollerState) ElementType() reflect.Type {
	return reflect.TypeOf((*intercontrollerState)(nil)).Elem()
}

type intercontrollerArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
	FastFailoverMax *int `pulumi:"fastFailoverMax"`
	// Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
	FastFailoverWait *int `pulumi:"fastFailoverWait"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Secret key for inter-controller communications.
	InterControllerKey *string `pulumi:"interControllerKey"`
	// Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
	InterControllerMode *string `pulumi:"interControllerMode"`
	// Fast failover peer wireless controller list. The structure of `interControllerPeer` block is documented below.
	InterControllerPeers []IntercontrollerInterControllerPeer `pulumi:"interControllerPeers"`
	// Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
	InterControllerPri *string `pulumi:"interControllerPri"`
	// Enable/disable layer 3 roaming (default = disable). Valid values: `enable`, `disable`.
	L3Roaming *string `pulumi:"l3Roaming"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Intercontroller resource.
type IntercontrollerArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
	FastFailoverMax pulumi.IntPtrInput
	// Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
	FastFailoverWait pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Secret key for inter-controller communications.
	InterControllerKey pulumi.StringPtrInput
	// Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
	InterControllerMode pulumi.StringPtrInput
	// Fast failover peer wireless controller list. The structure of `interControllerPeer` block is documented below.
	InterControllerPeers IntercontrollerInterControllerPeerArrayInput
	// Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
	InterControllerPri pulumi.StringPtrInput
	// Enable/disable layer 3 roaming (default = disable). Valid values: `enable`, `disable`.
	L3Roaming pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IntercontrollerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*intercontrollerArgs)(nil)).Elem()
}

type IntercontrollerInput interface {
	pulumi.Input

	ToIntercontrollerOutput() IntercontrollerOutput
	ToIntercontrollerOutputWithContext(ctx context.Context) IntercontrollerOutput
}

func (*Intercontroller) ElementType() reflect.Type {
	return reflect.TypeOf((**Intercontroller)(nil)).Elem()
}

func (i *Intercontroller) ToIntercontrollerOutput() IntercontrollerOutput {
	return i.ToIntercontrollerOutputWithContext(context.Background())
}

func (i *Intercontroller) ToIntercontrollerOutputWithContext(ctx context.Context) IntercontrollerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntercontrollerOutput)
}

// IntercontrollerArrayInput is an input type that accepts IntercontrollerArray and IntercontrollerArrayOutput values.
// You can construct a concrete instance of `IntercontrollerArrayInput` via:
//
//	IntercontrollerArray{ IntercontrollerArgs{...} }
type IntercontrollerArrayInput interface {
	pulumi.Input

	ToIntercontrollerArrayOutput() IntercontrollerArrayOutput
	ToIntercontrollerArrayOutputWithContext(context.Context) IntercontrollerArrayOutput
}

type IntercontrollerArray []IntercontrollerInput

func (IntercontrollerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Intercontroller)(nil)).Elem()
}

func (i IntercontrollerArray) ToIntercontrollerArrayOutput() IntercontrollerArrayOutput {
	return i.ToIntercontrollerArrayOutputWithContext(context.Background())
}

func (i IntercontrollerArray) ToIntercontrollerArrayOutputWithContext(ctx context.Context) IntercontrollerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntercontrollerArrayOutput)
}

// IntercontrollerMapInput is an input type that accepts IntercontrollerMap and IntercontrollerMapOutput values.
// You can construct a concrete instance of `IntercontrollerMapInput` via:
//
//	IntercontrollerMap{ "key": IntercontrollerArgs{...} }
type IntercontrollerMapInput interface {
	pulumi.Input

	ToIntercontrollerMapOutput() IntercontrollerMapOutput
	ToIntercontrollerMapOutputWithContext(context.Context) IntercontrollerMapOutput
}

type IntercontrollerMap map[string]IntercontrollerInput

func (IntercontrollerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Intercontroller)(nil)).Elem()
}

func (i IntercontrollerMap) ToIntercontrollerMapOutput() IntercontrollerMapOutput {
	return i.ToIntercontrollerMapOutputWithContext(context.Background())
}

func (i IntercontrollerMap) ToIntercontrollerMapOutputWithContext(ctx context.Context) IntercontrollerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntercontrollerMapOutput)
}

type IntercontrollerOutput struct{ *pulumi.OutputState }

func (IntercontrollerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Intercontroller)(nil)).Elem()
}

func (o IntercontrollerOutput) ToIntercontrollerOutput() IntercontrollerOutput {
	return o
}

func (o IntercontrollerOutput) ToIntercontrollerOutputWithContext(ctx context.Context) IntercontrollerOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o IntercontrollerOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Intercontroller) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
func (o IntercontrollerOutput) FastFailoverMax() pulumi.IntOutput {
	return o.ApplyT(func(v *Intercontroller) pulumi.IntOutput { return v.FastFailoverMax }).(pulumi.IntOutput)
}

// Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
func (o IntercontrollerOutput) FastFailoverWait() pulumi.IntOutput {
	return o.ApplyT(func(v *Intercontroller) pulumi.IntOutput { return v.FastFailoverWait }).(pulumi.IntOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o IntercontrollerOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Intercontroller) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Secret key for inter-controller communications.
func (o IntercontrollerOutput) InterControllerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Intercontroller) pulumi.StringPtrOutput { return v.InterControllerKey }).(pulumi.StringPtrOutput)
}

// Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
func (o IntercontrollerOutput) InterControllerMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Intercontroller) pulumi.StringOutput { return v.InterControllerMode }).(pulumi.StringOutput)
}

// Fast failover peer wireless controller list. The structure of `interControllerPeer` block is documented below.
func (o IntercontrollerOutput) InterControllerPeers() IntercontrollerInterControllerPeerArrayOutput {
	return o.ApplyT(func(v *Intercontroller) IntercontrollerInterControllerPeerArrayOutput { return v.InterControllerPeers }).(IntercontrollerInterControllerPeerArrayOutput)
}

// Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
func (o IntercontrollerOutput) InterControllerPri() pulumi.StringOutput {
	return o.ApplyT(func(v *Intercontroller) pulumi.StringOutput { return v.InterControllerPri }).(pulumi.StringOutput)
}

// Enable/disable layer 3 roaming (default = disable). Valid values: `enable`, `disable`.
func (o IntercontrollerOutput) L3Roaming() pulumi.StringOutput {
	return o.ApplyT(func(v *Intercontroller) pulumi.StringOutput { return v.L3Roaming }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o IntercontrollerOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Intercontroller) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type IntercontrollerArrayOutput struct{ *pulumi.OutputState }

func (IntercontrollerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Intercontroller)(nil)).Elem()
}

func (o IntercontrollerArrayOutput) ToIntercontrollerArrayOutput() IntercontrollerArrayOutput {
	return o
}

func (o IntercontrollerArrayOutput) ToIntercontrollerArrayOutputWithContext(ctx context.Context) IntercontrollerArrayOutput {
	return o
}

func (o IntercontrollerArrayOutput) Index(i pulumi.IntInput) IntercontrollerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Intercontroller {
		return vs[0].([]*Intercontroller)[vs[1].(int)]
	}).(IntercontrollerOutput)
}

type IntercontrollerMapOutput struct{ *pulumi.OutputState }

func (IntercontrollerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Intercontroller)(nil)).Elem()
}

func (o IntercontrollerMapOutput) ToIntercontrollerMapOutput() IntercontrollerMapOutput {
	return o
}

func (o IntercontrollerMapOutput) ToIntercontrollerMapOutputWithContext(ctx context.Context) IntercontrollerMapOutput {
	return o
}

func (o IntercontrollerMapOutput) MapIndex(k pulumi.StringInput) IntercontrollerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Intercontroller {
		return vs[0].(map[string]*Intercontroller)[vs[1].(string)]
	}).(IntercontrollerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntercontrollerInput)(nil)).Elem(), &Intercontroller{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntercontrollerArrayInput)(nil)).Elem(), IntercontrollerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntercontrollerMapInput)(nil)).Elem(), IntercontrollerMap{})
	pulumi.RegisterOutputType(IntercontrollerOutput{})
	pulumi.RegisterOutputType(IntercontrollerArrayOutput{})
	pulumi.RegisterOutputType(IntercontrollerMapOutput{})
}
