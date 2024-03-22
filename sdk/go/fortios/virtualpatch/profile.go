// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package virtualpatch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure virtual-patch profile. Applies to FortiOS Version `>= 7.4.1`.
//
// ## Import
//
// VirtualPatch Profile can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:virtualpatch/profile:Profile labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:virtualpatch/profile:Profile labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Profile struct {
	pulumi.CustomResourceState

	// Action (pass/block). Valid values: `pass`, `block`.
	Action pulumi.StringOutput `pulumi:"action"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Exempt devices or rules. The structure of `exemption` block is documented below.
	Exemptions ProfileExemptionArrayOutput `pulumi:"exemptions"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Enable/disable logging of detection. Valid values: `enable`, `disable`.
	Log pulumi.StringOutput `pulumi:"log"`
	// Profile name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Relative severity of the signature (low, medium, high, critical). Valid values: `low`, `medium`, `high`, `critical`.
	Severity pulumi.StringOutput `pulumi:"severity"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewProfile registers a new resource with the given unique name, arguments, and options.
func NewProfile(ctx *pulumi.Context,
	name string, args *ProfileArgs, opts ...pulumi.ResourceOption) (*Profile, error) {
	if args == nil {
		args = &ProfileArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Profile
	err := ctx.RegisterResource("fortios:virtualpatch/profile:Profile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfile gets an existing Profile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileState, opts ...pulumi.ResourceOption) (*Profile, error) {
	var resource Profile
	err := ctx.ReadResource("fortios:virtualpatch/profile:Profile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Profile resources.
type profileState struct {
	// Action (pass/block). Valid values: `pass`, `block`.
	Action *string `pulumi:"action"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Exempt devices or rules. The structure of `exemption` block is documented below.
	Exemptions []ProfileExemption `pulumi:"exemptions"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Enable/disable logging of detection. Valid values: `enable`, `disable`.
	Log *string `pulumi:"log"`
	// Profile name.
	Name *string `pulumi:"name"`
	// Relative severity of the signature (low, medium, high, critical). Valid values: `low`, `medium`, `high`, `critical`.
	Severity *string `pulumi:"severity"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ProfileState struct {
	// Action (pass/block). Valid values: `pass`, `block`.
	Action pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Exempt devices or rules. The structure of `exemption` block is documented below.
	Exemptions ProfileExemptionArrayInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Enable/disable logging of detection. Valid values: `enable`, `disable`.
	Log pulumi.StringPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// Relative severity of the signature (low, medium, high, critical). Valid values: `low`, `medium`, `high`, `critical`.
	Severity pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileState)(nil)).Elem()
}

type profileArgs struct {
	// Action (pass/block). Valid values: `pass`, `block`.
	Action *string `pulumi:"action"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Exempt devices or rules. The structure of `exemption` block is documented below.
	Exemptions []ProfileExemption `pulumi:"exemptions"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Enable/disable logging of detection. Valid values: `enable`, `disable`.
	Log *string `pulumi:"log"`
	// Profile name.
	Name *string `pulumi:"name"`
	// Relative severity of the signature (low, medium, high, critical). Valid values: `low`, `medium`, `high`, `critical`.
	Severity *string `pulumi:"severity"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Profile resource.
type ProfileArgs struct {
	// Action (pass/block). Valid values: `pass`, `block`.
	Action pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Exempt devices or rules. The structure of `exemption` block is documented below.
	Exemptions ProfileExemptionArrayInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Enable/disable logging of detection. Valid values: `enable`, `disable`.
	Log pulumi.StringPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// Relative severity of the signature (low, medium, high, critical). Valid values: `low`, `medium`, `high`, `critical`.
	Severity pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileArgs)(nil)).Elem()
}

type ProfileInput interface {
	pulumi.Input

	ToProfileOutput() ProfileOutput
	ToProfileOutputWithContext(ctx context.Context) ProfileOutput
}

func (*Profile) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (i *Profile) ToProfileOutput() ProfileOutput {
	return i.ToProfileOutputWithContext(context.Background())
}

func (i *Profile) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileOutput)
}

// ProfileArrayInput is an input type that accepts ProfileArray and ProfileArrayOutput values.
// You can construct a concrete instance of `ProfileArrayInput` via:
//
//	ProfileArray{ ProfileArgs{...} }
type ProfileArrayInput interface {
	pulumi.Input

	ToProfileArrayOutput() ProfileArrayOutput
	ToProfileArrayOutputWithContext(context.Context) ProfileArrayOutput
}

type ProfileArray []ProfileInput

func (ProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Profile)(nil)).Elem()
}

func (i ProfileArray) ToProfileArrayOutput() ProfileArrayOutput {
	return i.ToProfileArrayOutputWithContext(context.Background())
}

func (i ProfileArray) ToProfileArrayOutputWithContext(ctx context.Context) ProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileArrayOutput)
}

// ProfileMapInput is an input type that accepts ProfileMap and ProfileMapOutput values.
// You can construct a concrete instance of `ProfileMapInput` via:
//
//	ProfileMap{ "key": ProfileArgs{...} }
type ProfileMapInput interface {
	pulumi.Input

	ToProfileMapOutput() ProfileMapOutput
	ToProfileMapOutputWithContext(context.Context) ProfileMapOutput
}

type ProfileMap map[string]ProfileInput

func (ProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Profile)(nil)).Elem()
}

func (i ProfileMap) ToProfileMapOutput() ProfileMapOutput {
	return i.ToProfileMapOutputWithContext(context.Background())
}

func (i ProfileMap) ToProfileMapOutputWithContext(ctx context.Context) ProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileMapOutput)
}

type ProfileOutput struct{ *pulumi.OutputState }

func (ProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (o ProfileOutput) ToProfileOutput() ProfileOutput {
	return o
}

func (o ProfileOutput) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return o
}

// Action (pass/block). Valid values: `pass`, `block`.
func (o ProfileOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// Comment.
func (o ProfileOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o ProfileOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Exempt devices or rules. The structure of `exemption` block is documented below.
func (o ProfileOutput) Exemptions() ProfileExemptionArrayOutput {
	return o.ApplyT(func(v *Profile) ProfileExemptionArrayOutput { return v.Exemptions }).(ProfileExemptionArrayOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o ProfileOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Enable/disable logging of detection. Valid values: `enable`, `disable`.
func (o ProfileOutput) Log() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Log }).(pulumi.StringOutput)
}

// Profile name.
func (o ProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Relative severity of the signature (low, medium, high, critical). Valid values: `low`, `medium`, `high`, `critical`.
func (o ProfileOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ProfileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ProfileArrayOutput struct{ *pulumi.OutputState }

func (ProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Profile)(nil)).Elem()
}

func (o ProfileArrayOutput) ToProfileArrayOutput() ProfileArrayOutput {
	return o
}

func (o ProfileArrayOutput) ToProfileArrayOutputWithContext(ctx context.Context) ProfileArrayOutput {
	return o
}

func (o ProfileArrayOutput) Index(i pulumi.IntInput) ProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Profile {
		return vs[0].([]*Profile)[vs[1].(int)]
	}).(ProfileOutput)
}

type ProfileMapOutput struct{ *pulumi.OutputState }

func (ProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Profile)(nil)).Elem()
}

func (o ProfileMapOutput) ToProfileMapOutput() ProfileMapOutput {
	return o
}

func (o ProfileMapOutput) ToProfileMapOutputWithContext(ctx context.Context) ProfileMapOutput {
	return o
}

func (o ProfileMapOutput) MapIndex(k pulumi.StringInput) ProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Profile {
		return vs[0].(map[string]*Profile)[vs[1].(string)]
	}).(ProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileInput)(nil)).Elem(), &Profile{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileArrayInput)(nil)).Elem(), ProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileMapInput)(nil)).Elem(), ProfileMap{})
	pulumi.RegisterOutputType(ProfileOutput{})
	pulumi.RegisterOutputType(ProfileArrayOutput{})
	pulumi.RegisterOutputType(ProfileMapOutput{})
}
