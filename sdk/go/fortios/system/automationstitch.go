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

// Automation stitches.
//
// ## Import
//
// System AutomationStitch can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/automationstitch:Automationstitch labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/automationstitch:Automationstitch labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Automationstitch struct {
	pulumi.CustomResourceState

	// Action names. The structure of `action` block is documented below.
	Action AutomationstitchActionArrayOutput `pulumi:"action"`
	// Configure stitch actions. The structure of `actions` block is documented below.
	Actions AutomationstitchActionArrayOutput `pulumi:"actions"`
	// Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Serial number/HA group-name of destination devices. The structure of `destination` block is documented below.
	Destinations AutomationstitchDestinationArrayOutput `pulumi:"destinations"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable this stitch. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Trigger name.
	Trigger pulumi.StringOutput `pulumi:"trigger"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewAutomationstitch registers a new resource with the given unique name, arguments, and options.
func NewAutomationstitch(ctx *pulumi.Context,
	name string, args *AutomationstitchArgs, opts ...pulumi.ResourceOption) (*Automationstitch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	if args.Trigger == nil {
		return nil, errors.New("invalid value for required argument 'Trigger'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Automationstitch
	err := ctx.RegisterResource("fortios:system/automationstitch:Automationstitch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutomationstitch gets an existing Automationstitch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutomationstitch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutomationstitchState, opts ...pulumi.ResourceOption) (*Automationstitch, error) {
	var resource Automationstitch
	err := ctx.ReadResource("fortios:system/automationstitch:Automationstitch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Automationstitch resources.
type automationstitchState struct {
	// Action names. The structure of `action` block is documented below.
	Action []AutomationstitchAction `pulumi:"action"`
	// Configure stitch actions. The structure of `actions` block is documented below.
	Actions []AutomationstitchAction `pulumi:"actions"`
	// Description.
	Description *string `pulumi:"description"`
	// Serial number/HA group-name of destination devices. The structure of `destination` block is documented below.
	Destinations []AutomationstitchDestination `pulumi:"destinations"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name.
	Name *string `pulumi:"name"`
	// Enable/disable this stitch. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Trigger name.
	Trigger *string `pulumi:"trigger"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type AutomationstitchState struct {
	// Action names. The structure of `action` block is documented below.
	Action AutomationstitchActionArrayInput
	// Configure stitch actions. The structure of `actions` block is documented below.
	Actions AutomationstitchActionArrayInput
	// Description.
	Description pulumi.StringPtrInput
	// Serial number/HA group-name of destination devices. The structure of `destination` block is documented below.
	Destinations AutomationstitchDestinationArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Enable/disable this stitch. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Trigger name.
	Trigger pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AutomationstitchState) ElementType() reflect.Type {
	return reflect.TypeOf((*automationstitchState)(nil)).Elem()
}

type automationstitchArgs struct {
	// Action names. The structure of `action` block is documented below.
	Action []AutomationstitchAction `pulumi:"action"`
	// Configure stitch actions. The structure of `actions` block is documented below.
	Actions []AutomationstitchAction `pulumi:"actions"`
	// Description.
	Description *string `pulumi:"description"`
	// Serial number/HA group-name of destination devices. The structure of `destination` block is documented below.
	Destinations []AutomationstitchDestination `pulumi:"destinations"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name.
	Name *string `pulumi:"name"`
	// Enable/disable this stitch. Valid values: `enable`, `disable`.
	Status string `pulumi:"status"`
	// Trigger name.
	Trigger string `pulumi:"trigger"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Automationstitch resource.
type AutomationstitchArgs struct {
	// Action names. The structure of `action` block is documented below.
	Action AutomationstitchActionArrayInput
	// Configure stitch actions. The structure of `actions` block is documented below.
	Actions AutomationstitchActionArrayInput
	// Description.
	Description pulumi.StringPtrInput
	// Serial number/HA group-name of destination devices. The structure of `destination` block is documented below.
	Destinations AutomationstitchDestinationArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Enable/disable this stitch. Valid values: `enable`, `disable`.
	Status pulumi.StringInput
	// Trigger name.
	Trigger pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AutomationstitchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*automationstitchArgs)(nil)).Elem()
}

type AutomationstitchInput interface {
	pulumi.Input

	ToAutomationstitchOutput() AutomationstitchOutput
	ToAutomationstitchOutputWithContext(ctx context.Context) AutomationstitchOutput
}

func (*Automationstitch) ElementType() reflect.Type {
	return reflect.TypeOf((**Automationstitch)(nil)).Elem()
}

func (i *Automationstitch) ToAutomationstitchOutput() AutomationstitchOutput {
	return i.ToAutomationstitchOutputWithContext(context.Background())
}

func (i *Automationstitch) ToAutomationstitchOutputWithContext(ctx context.Context) AutomationstitchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationstitchOutput)
}

// AutomationstitchArrayInput is an input type that accepts AutomationstitchArray and AutomationstitchArrayOutput values.
// You can construct a concrete instance of `AutomationstitchArrayInput` via:
//
//	AutomationstitchArray{ AutomationstitchArgs{...} }
type AutomationstitchArrayInput interface {
	pulumi.Input

	ToAutomationstitchArrayOutput() AutomationstitchArrayOutput
	ToAutomationstitchArrayOutputWithContext(context.Context) AutomationstitchArrayOutput
}

type AutomationstitchArray []AutomationstitchInput

func (AutomationstitchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Automationstitch)(nil)).Elem()
}

func (i AutomationstitchArray) ToAutomationstitchArrayOutput() AutomationstitchArrayOutput {
	return i.ToAutomationstitchArrayOutputWithContext(context.Background())
}

func (i AutomationstitchArray) ToAutomationstitchArrayOutputWithContext(ctx context.Context) AutomationstitchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationstitchArrayOutput)
}

// AutomationstitchMapInput is an input type that accepts AutomationstitchMap and AutomationstitchMapOutput values.
// You can construct a concrete instance of `AutomationstitchMapInput` via:
//
//	AutomationstitchMap{ "key": AutomationstitchArgs{...} }
type AutomationstitchMapInput interface {
	pulumi.Input

	ToAutomationstitchMapOutput() AutomationstitchMapOutput
	ToAutomationstitchMapOutputWithContext(context.Context) AutomationstitchMapOutput
}

type AutomationstitchMap map[string]AutomationstitchInput

func (AutomationstitchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Automationstitch)(nil)).Elem()
}

func (i AutomationstitchMap) ToAutomationstitchMapOutput() AutomationstitchMapOutput {
	return i.ToAutomationstitchMapOutputWithContext(context.Background())
}

func (i AutomationstitchMap) ToAutomationstitchMapOutputWithContext(ctx context.Context) AutomationstitchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationstitchMapOutput)
}

type AutomationstitchOutput struct{ *pulumi.OutputState }

func (AutomationstitchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Automationstitch)(nil)).Elem()
}

func (o AutomationstitchOutput) ToAutomationstitchOutput() AutomationstitchOutput {
	return o
}

func (o AutomationstitchOutput) ToAutomationstitchOutputWithContext(ctx context.Context) AutomationstitchOutput {
	return o
}

// Action names. The structure of `action` block is documented below.
func (o AutomationstitchOutput) Action() AutomationstitchActionArrayOutput {
	return o.ApplyT(func(v *Automationstitch) AutomationstitchActionArrayOutput { return v.Action }).(AutomationstitchActionArrayOutput)
}

// Configure stitch actions. The structure of `actions` block is documented below.
func (o AutomationstitchOutput) Actions() AutomationstitchActionArrayOutput {
	return o.ApplyT(func(v *Automationstitch) AutomationstitchActionArrayOutput { return v.Actions }).(AutomationstitchActionArrayOutput)
}

// Description.
func (o AutomationstitchOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Automationstitch) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Serial number/HA group-name of destination devices. The structure of `destination` block is documented below.
func (o AutomationstitchOutput) Destinations() AutomationstitchDestinationArrayOutput {
	return o.ApplyT(func(v *Automationstitch) AutomationstitchDestinationArrayOutput { return v.Destinations }).(AutomationstitchDestinationArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o AutomationstitchOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Automationstitch) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Name.
func (o AutomationstitchOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Automationstitch) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable this stitch. Valid values: `enable`, `disable`.
func (o AutomationstitchOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Automationstitch) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Trigger name.
func (o AutomationstitchOutput) Trigger() pulumi.StringOutput {
	return o.ApplyT(func(v *Automationstitch) pulumi.StringOutput { return v.Trigger }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o AutomationstitchOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Automationstitch) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type AutomationstitchArrayOutput struct{ *pulumi.OutputState }

func (AutomationstitchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Automationstitch)(nil)).Elem()
}

func (o AutomationstitchArrayOutput) ToAutomationstitchArrayOutput() AutomationstitchArrayOutput {
	return o
}

func (o AutomationstitchArrayOutput) ToAutomationstitchArrayOutputWithContext(ctx context.Context) AutomationstitchArrayOutput {
	return o
}

func (o AutomationstitchArrayOutput) Index(i pulumi.IntInput) AutomationstitchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Automationstitch {
		return vs[0].([]*Automationstitch)[vs[1].(int)]
	}).(AutomationstitchOutput)
}

type AutomationstitchMapOutput struct{ *pulumi.OutputState }

func (AutomationstitchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Automationstitch)(nil)).Elem()
}

func (o AutomationstitchMapOutput) ToAutomationstitchMapOutput() AutomationstitchMapOutput {
	return o
}

func (o AutomationstitchMapOutput) ToAutomationstitchMapOutputWithContext(ctx context.Context) AutomationstitchMapOutput {
	return o
}

func (o AutomationstitchMapOutput) MapIndex(k pulumi.StringInput) AutomationstitchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Automationstitch {
		return vs[0].(map[string]*Automationstitch)[vs[1].(string)]
	}).(AutomationstitchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationstitchInput)(nil)).Elem(), &Automationstitch{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationstitchArrayInput)(nil)).Elem(), AutomationstitchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationstitchMapInput)(nil)).Elem(), AutomationstitchMap{})
	pulumi.RegisterOutputType(AutomationstitchOutput{})
	pulumi.RegisterOutputType(AutomationstitchArrayOutput{})
	pulumi.RegisterOutputType(AutomationstitchMapOutput{})
}