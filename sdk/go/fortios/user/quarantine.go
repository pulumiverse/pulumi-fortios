// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package user

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure quarantine support.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/user"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := user.NewQuarantine(ctx, "trname", &user.QuarantineArgs{
//				Quarantine: pulumi.String("enable"),
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
// User Quarantine can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:user/quarantine:Quarantine labelname UserQuarantine
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:user/quarantine:Quarantine labelname UserQuarantine
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Quarantine struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Firewall address group which includes all quarantine MAC address.
	FirewallGroups pulumi.StringOutput `pulumi:"firewallGroups"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Enable/disable quarantine. Valid values: `enable`, `disable`.
	Quarantine pulumi.StringOutput `pulumi:"quarantine"`
	// Quarantine entry to hold multiple MACs. The structure of `targets` block is documented below.
	Targets QuarantineTargetArrayOutput `pulumi:"targets"`
	// Traffic policy for quarantined MACs.
	TrafficPolicy pulumi.StringOutput `pulumi:"trafficPolicy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewQuarantine registers a new resource with the given unique name, arguments, and options.
func NewQuarantine(ctx *pulumi.Context,
	name string, args *QuarantineArgs, opts ...pulumi.ResourceOption) (*Quarantine, error) {
	if args == nil {
		args = &QuarantineArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Quarantine
	err := ctx.RegisterResource("fortios:user/quarantine:Quarantine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQuarantine gets an existing Quarantine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQuarantine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QuarantineState, opts ...pulumi.ResourceOption) (*Quarantine, error) {
	var resource Quarantine
	err := ctx.ReadResource("fortios:user/quarantine:Quarantine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Quarantine resources.
type quarantineState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Firewall address group which includes all quarantine MAC address.
	FirewallGroups *string `pulumi:"firewallGroups"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Enable/disable quarantine. Valid values: `enable`, `disable`.
	Quarantine *string `pulumi:"quarantine"`
	// Quarantine entry to hold multiple MACs. The structure of `targets` block is documented below.
	Targets []QuarantineTarget `pulumi:"targets"`
	// Traffic policy for quarantined MACs.
	TrafficPolicy *string `pulumi:"trafficPolicy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type QuarantineState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Firewall address group which includes all quarantine MAC address.
	FirewallGroups pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Enable/disable quarantine. Valid values: `enable`, `disable`.
	Quarantine pulumi.StringPtrInput
	// Quarantine entry to hold multiple MACs. The structure of `targets` block is documented below.
	Targets QuarantineTargetArrayInput
	// Traffic policy for quarantined MACs.
	TrafficPolicy pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (QuarantineState) ElementType() reflect.Type {
	return reflect.TypeOf((*quarantineState)(nil)).Elem()
}

type quarantineArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Firewall address group which includes all quarantine MAC address.
	FirewallGroups *string `pulumi:"firewallGroups"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Enable/disable quarantine. Valid values: `enable`, `disable`.
	Quarantine *string `pulumi:"quarantine"`
	// Quarantine entry to hold multiple MACs. The structure of `targets` block is documented below.
	Targets []QuarantineTarget `pulumi:"targets"`
	// Traffic policy for quarantined MACs.
	TrafficPolicy *string `pulumi:"trafficPolicy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Quarantine resource.
type QuarantineArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Firewall address group which includes all quarantine MAC address.
	FirewallGroups pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Enable/disable quarantine. Valid values: `enable`, `disable`.
	Quarantine pulumi.StringPtrInput
	// Quarantine entry to hold multiple MACs. The structure of `targets` block is documented below.
	Targets QuarantineTargetArrayInput
	// Traffic policy for quarantined MACs.
	TrafficPolicy pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (QuarantineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*quarantineArgs)(nil)).Elem()
}

type QuarantineInput interface {
	pulumi.Input

	ToQuarantineOutput() QuarantineOutput
	ToQuarantineOutputWithContext(ctx context.Context) QuarantineOutput
}

func (*Quarantine) ElementType() reflect.Type {
	return reflect.TypeOf((**Quarantine)(nil)).Elem()
}

func (i *Quarantine) ToQuarantineOutput() QuarantineOutput {
	return i.ToQuarantineOutputWithContext(context.Background())
}

func (i *Quarantine) ToQuarantineOutputWithContext(ctx context.Context) QuarantineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuarantineOutput)
}

// QuarantineArrayInput is an input type that accepts QuarantineArray and QuarantineArrayOutput values.
// You can construct a concrete instance of `QuarantineArrayInput` via:
//
//	QuarantineArray{ QuarantineArgs{...} }
type QuarantineArrayInput interface {
	pulumi.Input

	ToQuarantineArrayOutput() QuarantineArrayOutput
	ToQuarantineArrayOutputWithContext(context.Context) QuarantineArrayOutput
}

type QuarantineArray []QuarantineInput

func (QuarantineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Quarantine)(nil)).Elem()
}

func (i QuarantineArray) ToQuarantineArrayOutput() QuarantineArrayOutput {
	return i.ToQuarantineArrayOutputWithContext(context.Background())
}

func (i QuarantineArray) ToQuarantineArrayOutputWithContext(ctx context.Context) QuarantineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuarantineArrayOutput)
}

// QuarantineMapInput is an input type that accepts QuarantineMap and QuarantineMapOutput values.
// You can construct a concrete instance of `QuarantineMapInput` via:
//
//	QuarantineMap{ "key": QuarantineArgs{...} }
type QuarantineMapInput interface {
	pulumi.Input

	ToQuarantineMapOutput() QuarantineMapOutput
	ToQuarantineMapOutputWithContext(context.Context) QuarantineMapOutput
}

type QuarantineMap map[string]QuarantineInput

func (QuarantineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Quarantine)(nil)).Elem()
}

func (i QuarantineMap) ToQuarantineMapOutput() QuarantineMapOutput {
	return i.ToQuarantineMapOutputWithContext(context.Background())
}

func (i QuarantineMap) ToQuarantineMapOutputWithContext(ctx context.Context) QuarantineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuarantineMapOutput)
}

type QuarantineOutput struct{ *pulumi.OutputState }

func (QuarantineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Quarantine)(nil)).Elem()
}

func (o QuarantineOutput) ToQuarantineOutput() QuarantineOutput {
	return o
}

func (o QuarantineOutput) ToQuarantineOutputWithContext(ctx context.Context) QuarantineOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o QuarantineOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Firewall address group which includes all quarantine MAC address.
func (o QuarantineOutput) FirewallGroups() pulumi.StringOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.StringOutput { return v.FirewallGroups }).(pulumi.StringOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o QuarantineOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Enable/disable quarantine. Valid values: `enable`, `disable`.
func (o QuarantineOutput) Quarantine() pulumi.StringOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.StringOutput { return v.Quarantine }).(pulumi.StringOutput)
}

// Quarantine entry to hold multiple MACs. The structure of `targets` block is documented below.
func (o QuarantineOutput) Targets() QuarantineTargetArrayOutput {
	return o.ApplyT(func(v *Quarantine) QuarantineTargetArrayOutput { return v.Targets }).(QuarantineTargetArrayOutput)
}

// Traffic policy for quarantined MACs.
func (o QuarantineOutput) TrafficPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.StringOutput { return v.TrafficPolicy }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o QuarantineOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type QuarantineArrayOutput struct{ *pulumi.OutputState }

func (QuarantineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Quarantine)(nil)).Elem()
}

func (o QuarantineArrayOutput) ToQuarantineArrayOutput() QuarantineArrayOutput {
	return o
}

func (o QuarantineArrayOutput) ToQuarantineArrayOutputWithContext(ctx context.Context) QuarantineArrayOutput {
	return o
}

func (o QuarantineArrayOutput) Index(i pulumi.IntInput) QuarantineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Quarantine {
		return vs[0].([]*Quarantine)[vs[1].(int)]
	}).(QuarantineOutput)
}

type QuarantineMapOutput struct{ *pulumi.OutputState }

func (QuarantineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Quarantine)(nil)).Elem()
}

func (o QuarantineMapOutput) ToQuarantineMapOutput() QuarantineMapOutput {
	return o
}

func (o QuarantineMapOutput) ToQuarantineMapOutputWithContext(ctx context.Context) QuarantineMapOutput {
	return o
}

func (o QuarantineMapOutput) MapIndex(k pulumi.StringInput) QuarantineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Quarantine {
		return vs[0].(map[string]*Quarantine)[vs[1].(string)]
	}).(QuarantineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QuarantineInput)(nil)).Elem(), &Quarantine{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuarantineArrayInput)(nil)).Elem(), QuarantineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuarantineMapInput)(nil)).Elem(), QuarantineMap{})
	pulumi.RegisterOutputType(QuarantineOutput{})
	pulumi.RegisterOutputType(QuarantineArrayOutput{})
	pulumi.RegisterOutputType(QuarantineMapOutput{})
}
