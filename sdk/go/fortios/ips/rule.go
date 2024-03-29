// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ips

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure IPS rules.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/ips"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// import first and then modify
//			_, err := ips.NewRule(ctx, "trname", &ips.RuleArgs{
//				Action:      pulumi.String("block"),
//				Application: pulumi.String("All"),
//				Date:        pulumi.Int(1462435200),
//				Group:       pulumi.String("backdoor"),
//				Location:    pulumi.String("server"),
//				Log:         pulumi.String("enable"),
//				LogPacket:   pulumi.String("disable"),
//				Os:          pulumi.String("All"),
//				Rev:         pulumi.Int(6637),
//				RuleId:      pulumi.Int(40473),
//				Service:     pulumi.String("UDP, DNS"),
//				Severity:    pulumi.String("critical"),
//				Status:      pulumi.String("enable"),
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
// Ips Rule can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:ips/rule:Rule labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:ips/rule:Rule labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Rule struct {
	pulumi.CustomResourceState

	// Action. Valid values: `pass`, `block`.
	Action pulumi.StringOutput `pulumi:"action"`
	// Vulnerable applications.
	Application pulumi.StringOutput `pulumi:"application"`
	// Date.
	Date pulumi.IntOutput `pulumi:"date"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Group.
	Group pulumi.StringOutput `pulumi:"group"`
	// Vulnerable location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Enable/disable logging. Valid values: `disable`, `enable`.
	Log pulumi.StringOutput `pulumi:"log"`
	// Enable/disable packet logging. Valid values: `disable`, `enable`.
	LogPacket pulumi.StringOutput `pulumi:"logPacket"`
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas RuleMetadataArrayOutput `pulumi:"metadatas"`
	// Rule name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Vulnerable operation systems.
	Os pulumi.StringOutput `pulumi:"os"`
	// Revision.
	Rev pulumi.IntOutput `pulumi:"rev"`
	// Rule ID.
	RuleId pulumi.IntOutput `pulumi:"ruleId"`
	// Vulnerable service.
	Service pulumi.StringOutput `pulumi:"service"`
	// Severity.
	Severity pulumi.StringOutput `pulumi:"severity"`
	// Enable/disable status. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil {
		args = &RuleArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Rule
	err := ctx.RegisterResource("fortios:ips/rule:Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleState, opts ...pulumi.ResourceOption) (*Rule, error) {
	var resource Rule
	err := ctx.ReadResource("fortios:ips/rule:Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rule resources.
type ruleState struct {
	// Action. Valid values: `pass`, `block`.
	Action *string `pulumi:"action"`
	// Vulnerable applications.
	Application *string `pulumi:"application"`
	// Date.
	Date *int `pulumi:"date"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Group.
	Group *string `pulumi:"group"`
	// Vulnerable location.
	Location *string `pulumi:"location"`
	// Enable/disable logging. Valid values: `disable`, `enable`.
	Log *string `pulumi:"log"`
	// Enable/disable packet logging. Valid values: `disable`, `enable`.
	LogPacket *string `pulumi:"logPacket"`
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas []RuleMetadata `pulumi:"metadatas"`
	// Rule name.
	Name *string `pulumi:"name"`
	// Vulnerable operation systems.
	Os *string `pulumi:"os"`
	// Revision.
	Rev *int `pulumi:"rev"`
	// Rule ID.
	RuleId *int `pulumi:"ruleId"`
	// Vulnerable service.
	Service *string `pulumi:"service"`
	// Severity.
	Severity *string `pulumi:"severity"`
	// Enable/disable status. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RuleState struct {
	// Action. Valid values: `pass`, `block`.
	Action pulumi.StringPtrInput
	// Vulnerable applications.
	Application pulumi.StringPtrInput
	// Date.
	Date pulumi.IntPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Group.
	Group pulumi.StringPtrInput
	// Vulnerable location.
	Location pulumi.StringPtrInput
	// Enable/disable logging. Valid values: `disable`, `enable`.
	Log pulumi.StringPtrInput
	// Enable/disable packet logging. Valid values: `disable`, `enable`.
	LogPacket pulumi.StringPtrInput
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas RuleMetadataArrayInput
	// Rule name.
	Name pulumi.StringPtrInput
	// Vulnerable operation systems.
	Os pulumi.StringPtrInput
	// Revision.
	Rev pulumi.IntPtrInput
	// Rule ID.
	RuleId pulumi.IntPtrInput
	// Vulnerable service.
	Service pulumi.StringPtrInput
	// Severity.
	Severity pulumi.StringPtrInput
	// Enable/disable status. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	// Action. Valid values: `pass`, `block`.
	Action *string `pulumi:"action"`
	// Vulnerable applications.
	Application *string `pulumi:"application"`
	// Date.
	Date *int `pulumi:"date"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Group.
	Group *string `pulumi:"group"`
	// Vulnerable location.
	Location *string `pulumi:"location"`
	// Enable/disable logging. Valid values: `disable`, `enable`.
	Log *string `pulumi:"log"`
	// Enable/disable packet logging. Valid values: `disable`, `enable`.
	LogPacket *string `pulumi:"logPacket"`
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas []RuleMetadata `pulumi:"metadatas"`
	// Rule name.
	Name *string `pulumi:"name"`
	// Vulnerable operation systems.
	Os *string `pulumi:"os"`
	// Revision.
	Rev *int `pulumi:"rev"`
	// Rule ID.
	RuleId *int `pulumi:"ruleId"`
	// Vulnerable service.
	Service *string `pulumi:"service"`
	// Severity.
	Severity *string `pulumi:"severity"`
	// Enable/disable status. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// Action. Valid values: `pass`, `block`.
	Action pulumi.StringPtrInput
	// Vulnerable applications.
	Application pulumi.StringPtrInput
	// Date.
	Date pulumi.IntPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Group.
	Group pulumi.StringPtrInput
	// Vulnerable location.
	Location pulumi.StringPtrInput
	// Enable/disable logging. Valid values: `disable`, `enable`.
	Log pulumi.StringPtrInput
	// Enable/disable packet logging. Valid values: `disable`, `enable`.
	LogPacket pulumi.StringPtrInput
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas RuleMetadataArrayInput
	// Rule name.
	Name pulumi.StringPtrInput
	// Vulnerable operation systems.
	Os pulumi.StringPtrInput
	// Revision.
	Rev pulumi.IntPtrInput
	// Rule ID.
	RuleId pulumi.IntPtrInput
	// Vulnerable service.
	Service pulumi.StringPtrInput
	// Severity.
	Severity pulumi.StringPtrInput
	// Enable/disable status. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleArgs)(nil)).Elem()
}

type RuleInput interface {
	pulumi.Input

	ToRuleOutput() RuleOutput
	ToRuleOutputWithContext(ctx context.Context) RuleOutput
}

func (*Rule) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil)).Elem()
}

func (i *Rule) ToRuleOutput() RuleOutput {
	return i.ToRuleOutputWithContext(context.Background())
}

func (i *Rule) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleOutput)
}

// RuleArrayInput is an input type that accepts RuleArray and RuleArrayOutput values.
// You can construct a concrete instance of `RuleArrayInput` via:
//
//	RuleArray{ RuleArgs{...} }
type RuleArrayInput interface {
	pulumi.Input

	ToRuleArrayOutput() RuleArrayOutput
	ToRuleArrayOutputWithContext(context.Context) RuleArrayOutput
}

type RuleArray []RuleInput

func (RuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rule)(nil)).Elem()
}

func (i RuleArray) ToRuleArrayOutput() RuleArrayOutput {
	return i.ToRuleArrayOutputWithContext(context.Background())
}

func (i RuleArray) ToRuleArrayOutputWithContext(ctx context.Context) RuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleArrayOutput)
}

// RuleMapInput is an input type that accepts RuleMap and RuleMapOutput values.
// You can construct a concrete instance of `RuleMapInput` via:
//
//	RuleMap{ "key": RuleArgs{...} }
type RuleMapInput interface {
	pulumi.Input

	ToRuleMapOutput() RuleMapOutput
	ToRuleMapOutputWithContext(context.Context) RuleMapOutput
}

type RuleMap map[string]RuleInput

func (RuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rule)(nil)).Elem()
}

func (i RuleMap) ToRuleMapOutput() RuleMapOutput {
	return i.ToRuleMapOutputWithContext(context.Background())
}

func (i RuleMap) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleMapOutput)
}

type RuleOutput struct{ *pulumi.OutputState }

func (RuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil)).Elem()
}

func (o RuleOutput) ToRuleOutput() RuleOutput {
	return o
}

func (o RuleOutput) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return o
}

// Action. Valid values: `pass`, `block`.
func (o RuleOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// Vulnerable applications.
func (o RuleOutput) Application() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Application }).(pulumi.StringOutput)
}

// Date.
func (o RuleOutput) Date() pulumi.IntOutput {
	return o.ApplyT(func(v *Rule) pulumi.IntOutput { return v.Date }).(pulumi.IntOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o RuleOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o RuleOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Group.
func (o RuleOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// Vulnerable location.
func (o RuleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Enable/disable logging. Valid values: `disable`, `enable`.
func (o RuleOutput) Log() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Log }).(pulumi.StringOutput)
}

// Enable/disable packet logging. Valid values: `disable`, `enable`.
func (o RuleOutput) LogPacket() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.LogPacket }).(pulumi.StringOutput)
}

// Meta data. The structure of `metadata` block is documented below.
func (o RuleOutput) Metadatas() RuleMetadataArrayOutput {
	return o.ApplyT(func(v *Rule) RuleMetadataArrayOutput { return v.Metadatas }).(RuleMetadataArrayOutput)
}

// Rule name.
func (o RuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Vulnerable operation systems.
func (o RuleOutput) Os() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Os }).(pulumi.StringOutput)
}

// Revision.
func (o RuleOutput) Rev() pulumi.IntOutput {
	return o.ApplyT(func(v *Rule) pulumi.IntOutput { return v.Rev }).(pulumi.IntOutput)
}

// Rule ID.
func (o RuleOutput) RuleId() pulumi.IntOutput {
	return o.ApplyT(func(v *Rule) pulumi.IntOutput { return v.RuleId }).(pulumi.IntOutput)
}

// Vulnerable service.
func (o RuleOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

// Severity.
func (o RuleOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

// Enable/disable status. Valid values: `disable`, `enable`.
func (o RuleOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o RuleOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RuleArrayOutput struct{ *pulumi.OutputState }

func (RuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rule)(nil)).Elem()
}

func (o RuleArrayOutput) ToRuleArrayOutput() RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) ToRuleArrayOutputWithContext(ctx context.Context) RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) Index(i pulumi.IntInput) RuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Rule {
		return vs[0].([]*Rule)[vs[1].(int)]
	}).(RuleOutput)
}

type RuleMapOutput struct{ *pulumi.OutputState }

func (RuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rule)(nil)).Elem()
}

func (o RuleMapOutput) ToRuleMapOutput() RuleMapOutput {
	return o
}

func (o RuleMapOutput) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return o
}

func (o RuleMapOutput) MapIndex(k pulumi.StringInput) RuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Rule {
		return vs[0].(map[string]*Rule)[vs[1].(string)]
	}).(RuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleInput)(nil)).Elem(), &Rule{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleArrayInput)(nil)).Elem(), RuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleMapInput)(nil)).Elem(), RuleMap{})
	pulumi.RegisterOutputType(RuleOutput{})
	pulumi.RegisterOutputType(RuleArrayOutput{})
	pulumi.RegisterOutputType(RuleMapOutput{})
}
