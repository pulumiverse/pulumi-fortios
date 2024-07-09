// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package extensioncontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// FortiGate connector profile configuration. Applies to FortiOS Version `>= 7.2.1`.
//
// ## Import
//
// ExtensionController FortigateProfile can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:extensioncontroller/fortigateprofile:Fortigateprofile labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:extensioncontroller/fortigateprofile:Fortigateprofile labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Fortigateprofile struct {
	pulumi.CustomResourceState

	// Extension option. Valid values: `lan-extension`.
	Extension pulumi.StringOutput `pulumi:"extension"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// FortiGate connector LAN extension configuration. The structure of `lanExtension` block is documented below.
	LanExtension FortigateprofileLanExtensionOutput `pulumi:"lanExtension"`
	// FortiGate connector profile name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewFortigateprofile registers a new resource with the given unique name, arguments, and options.
func NewFortigateprofile(ctx *pulumi.Context,
	name string, args *FortigateprofileArgs, opts ...pulumi.ResourceOption) (*Fortigateprofile, error) {
	if args == nil {
		args = &FortigateprofileArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Fortigateprofile
	err := ctx.RegisterResource("fortios:extensioncontroller/fortigateprofile:Fortigateprofile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortigateprofile gets an existing Fortigateprofile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortigateprofile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortigateprofileState, opts ...pulumi.ResourceOption) (*Fortigateprofile, error) {
	var resource Fortigateprofile
	err := ctx.ReadResource("fortios:extensioncontroller/fortigateprofile:Fortigateprofile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fortigateprofile resources.
type fortigateprofileState struct {
	// Extension option. Valid values: `lan-extension`.
	Extension *string `pulumi:"extension"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// FortiGate connector LAN extension configuration. The structure of `lanExtension` block is documented below.
	LanExtension *FortigateprofileLanExtension `pulumi:"lanExtension"`
	// FortiGate connector profile name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FortigateprofileState struct {
	// Extension option. Valid values: `lan-extension`.
	Extension pulumi.StringPtrInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// FortiGate connector LAN extension configuration. The structure of `lanExtension` block is documented below.
	LanExtension FortigateprofileLanExtensionPtrInput
	// FortiGate connector profile name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FortigateprofileState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortigateprofileState)(nil)).Elem()
}

type fortigateprofileArgs struct {
	// Extension option. Valid values: `lan-extension`.
	Extension *string `pulumi:"extension"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// FortiGate connector LAN extension configuration. The structure of `lanExtension` block is documented below.
	LanExtension *FortigateprofileLanExtension `pulumi:"lanExtension"`
	// FortiGate connector profile name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Fortigateprofile resource.
type FortigateprofileArgs struct {
	// Extension option. Valid values: `lan-extension`.
	Extension pulumi.StringPtrInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// FortiGate connector LAN extension configuration. The structure of `lanExtension` block is documented below.
	LanExtension FortigateprofileLanExtensionPtrInput
	// FortiGate connector profile name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FortigateprofileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortigateprofileArgs)(nil)).Elem()
}

type FortigateprofileInput interface {
	pulumi.Input

	ToFortigateprofileOutput() FortigateprofileOutput
	ToFortigateprofileOutputWithContext(ctx context.Context) FortigateprofileOutput
}

func (*Fortigateprofile) ElementType() reflect.Type {
	return reflect.TypeOf((**Fortigateprofile)(nil)).Elem()
}

func (i *Fortigateprofile) ToFortigateprofileOutput() FortigateprofileOutput {
	return i.ToFortigateprofileOutputWithContext(context.Background())
}

func (i *Fortigateprofile) ToFortigateprofileOutputWithContext(ctx context.Context) FortigateprofileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortigateprofileOutput)
}

// FortigateprofileArrayInput is an input type that accepts FortigateprofileArray and FortigateprofileArrayOutput values.
// You can construct a concrete instance of `FortigateprofileArrayInput` via:
//
//	FortigateprofileArray{ FortigateprofileArgs{...} }
type FortigateprofileArrayInput interface {
	pulumi.Input

	ToFortigateprofileArrayOutput() FortigateprofileArrayOutput
	ToFortigateprofileArrayOutputWithContext(context.Context) FortigateprofileArrayOutput
}

type FortigateprofileArray []FortigateprofileInput

func (FortigateprofileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fortigateprofile)(nil)).Elem()
}

func (i FortigateprofileArray) ToFortigateprofileArrayOutput() FortigateprofileArrayOutput {
	return i.ToFortigateprofileArrayOutputWithContext(context.Background())
}

func (i FortigateprofileArray) ToFortigateprofileArrayOutputWithContext(ctx context.Context) FortigateprofileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortigateprofileArrayOutput)
}

// FortigateprofileMapInput is an input type that accepts FortigateprofileMap and FortigateprofileMapOutput values.
// You can construct a concrete instance of `FortigateprofileMapInput` via:
//
//	FortigateprofileMap{ "key": FortigateprofileArgs{...} }
type FortigateprofileMapInput interface {
	pulumi.Input

	ToFortigateprofileMapOutput() FortigateprofileMapOutput
	ToFortigateprofileMapOutputWithContext(context.Context) FortigateprofileMapOutput
}

type FortigateprofileMap map[string]FortigateprofileInput

func (FortigateprofileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fortigateprofile)(nil)).Elem()
}

func (i FortigateprofileMap) ToFortigateprofileMapOutput() FortigateprofileMapOutput {
	return i.ToFortigateprofileMapOutputWithContext(context.Background())
}

func (i FortigateprofileMap) ToFortigateprofileMapOutputWithContext(ctx context.Context) FortigateprofileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortigateprofileMapOutput)
}

type FortigateprofileOutput struct{ *pulumi.OutputState }

func (FortigateprofileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fortigateprofile)(nil)).Elem()
}

func (o FortigateprofileOutput) ToFortigateprofileOutput() FortigateprofileOutput {
	return o
}

func (o FortigateprofileOutput) ToFortigateprofileOutputWithContext(ctx context.Context) FortigateprofileOutput {
	return o
}

// Extension option. Valid values: `lan-extension`.
func (o FortigateprofileOutput) Extension() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortigateprofile) pulumi.StringOutput { return v.Extension }).(pulumi.StringOutput)
}

// ID.
func (o FortigateprofileOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Fortigateprofile) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o FortigateprofileOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Fortigateprofile) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// FortiGate connector LAN extension configuration. The structure of `lanExtension` block is documented below.
func (o FortigateprofileOutput) LanExtension() FortigateprofileLanExtensionOutput {
	return o.ApplyT(func(v *Fortigateprofile) FortigateprofileLanExtensionOutput { return v.LanExtension }).(FortigateprofileLanExtensionOutput)
}

// FortiGate connector profile name.
func (o FortigateprofileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortigateprofile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FortigateprofileOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortigateprofile) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type FortigateprofileArrayOutput struct{ *pulumi.OutputState }

func (FortigateprofileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fortigateprofile)(nil)).Elem()
}

func (o FortigateprofileArrayOutput) ToFortigateprofileArrayOutput() FortigateprofileArrayOutput {
	return o
}

func (o FortigateprofileArrayOutput) ToFortigateprofileArrayOutputWithContext(ctx context.Context) FortigateprofileArrayOutput {
	return o
}

func (o FortigateprofileArrayOutput) Index(i pulumi.IntInput) FortigateprofileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Fortigateprofile {
		return vs[0].([]*Fortigateprofile)[vs[1].(int)]
	}).(FortigateprofileOutput)
}

type FortigateprofileMapOutput struct{ *pulumi.OutputState }

func (FortigateprofileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fortigateprofile)(nil)).Elem()
}

func (o FortigateprofileMapOutput) ToFortigateprofileMapOutput() FortigateprofileMapOutput {
	return o
}

func (o FortigateprofileMapOutput) ToFortigateprofileMapOutputWithContext(ctx context.Context) FortigateprofileMapOutput {
	return o
}

func (o FortigateprofileMapOutput) MapIndex(k pulumi.StringInput) FortigateprofileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Fortigateprofile {
		return vs[0].(map[string]*Fortigateprofile)[vs[1].(string)]
	}).(FortigateprofileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortigateprofileInput)(nil)).Elem(), &Fortigateprofile{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortigateprofileArrayInput)(nil)).Elem(), FortigateprofileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortigateprofileMapInput)(nil)).Elem(), FortigateprofileMap{})
	pulumi.RegisterOutputType(FortigateprofileOutput{})
	pulumi.RegisterOutputType(FortigateprofileArrayOutput{})
	pulumi.RegisterOutputType(FortigateprofileMapOutput{})
}
