// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dlp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure predefined data type used by DLP blocking. Applies to FortiOS Version `>= 7.2.0`.
//
// ## Import
//
// Dlp DataType can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:dlp/datatype:Datatype labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:dlp/datatype:Datatype labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Datatype struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Number of characters to obtain in advance for verification (1 - 255, default = 1).
	LookAhead pulumi.IntOutput `pulumi:"lookAhead"`
	// Number of characters required to save for verification (1 - 255, default = 1).
	LookBack pulumi.IntOutput `pulumi:"lookBack"`
	// Number of characters behind for match-around (1 - 4096, default = 1).
	MatchAhead pulumi.IntOutput `pulumi:"matchAhead"`
	// Dictionary to check whether it has a match around (Only support match-any and basic types, no repeat supported).
	MatchAround pulumi.StringOutput `pulumi:"matchAround"`
	// Number of characters in front for match-around (1 - 4096, default = 1).
	MatchBack pulumi.IntOutput `pulumi:"matchBack"`
	// Name of table containing the data type.
	Name pulumi.StringOutput `pulumi:"name"`
	// Regular expression pattern string without look around.
	Pattern pulumi.StringOutput `pulumi:"pattern"`
	// Template to transform user input to a pattern using capture group from 'pattern'.
	Transform pulumi.StringOutput `pulumi:"transform"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Regular expression pattern string used to verify the data type.
	Verify pulumi.StringOutput `pulumi:"verify"`
	// Extra regular expression pattern string used to verify the data type.
	Verify2 pulumi.StringOutput `pulumi:"verify2"`
	// Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
	VerifyTransformedPattern pulumi.StringOutput `pulumi:"verifyTransformedPattern"`
}

// NewDatatype registers a new resource with the given unique name, arguments, and options.
func NewDatatype(ctx *pulumi.Context,
	name string, args *DatatypeArgs, opts ...pulumi.ResourceOption) (*Datatype, error) {
	if args == nil {
		args = &DatatypeArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Datatype
	err := ctx.RegisterResource("fortios:dlp/datatype:Datatype", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatatype gets an existing Datatype resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatatype(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatatypeState, opts ...pulumi.ResourceOption) (*Datatype, error) {
	var resource Datatype
	err := ctx.ReadResource("fortios:dlp/datatype:Datatype", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Datatype resources.
type datatypeState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Number of characters to obtain in advance for verification (1 - 255, default = 1).
	LookAhead *int `pulumi:"lookAhead"`
	// Number of characters required to save for verification (1 - 255, default = 1).
	LookBack *int `pulumi:"lookBack"`
	// Number of characters behind for match-around (1 - 4096, default = 1).
	MatchAhead *int `pulumi:"matchAhead"`
	// Dictionary to check whether it has a match around (Only support match-any and basic types, no repeat supported).
	MatchAround *string `pulumi:"matchAround"`
	// Number of characters in front for match-around (1 - 4096, default = 1).
	MatchBack *int `pulumi:"matchBack"`
	// Name of table containing the data type.
	Name *string `pulumi:"name"`
	// Regular expression pattern string without look around.
	Pattern *string `pulumi:"pattern"`
	// Template to transform user input to a pattern using capture group from 'pattern'.
	Transform *string `pulumi:"transform"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Regular expression pattern string used to verify the data type.
	Verify *string `pulumi:"verify"`
	// Extra regular expression pattern string used to verify the data type.
	Verify2 *string `pulumi:"verify2"`
	// Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
	VerifyTransformedPattern *string `pulumi:"verifyTransformedPattern"`
}

type DatatypeState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Number of characters to obtain in advance for verification (1 - 255, default = 1).
	LookAhead pulumi.IntPtrInput
	// Number of characters required to save for verification (1 - 255, default = 1).
	LookBack pulumi.IntPtrInput
	// Number of characters behind for match-around (1 - 4096, default = 1).
	MatchAhead pulumi.IntPtrInput
	// Dictionary to check whether it has a match around (Only support match-any and basic types, no repeat supported).
	MatchAround pulumi.StringPtrInput
	// Number of characters in front for match-around (1 - 4096, default = 1).
	MatchBack pulumi.IntPtrInput
	// Name of table containing the data type.
	Name pulumi.StringPtrInput
	// Regular expression pattern string without look around.
	Pattern pulumi.StringPtrInput
	// Template to transform user input to a pattern using capture group from 'pattern'.
	Transform pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Regular expression pattern string used to verify the data type.
	Verify pulumi.StringPtrInput
	// Extra regular expression pattern string used to verify the data type.
	Verify2 pulumi.StringPtrInput
	// Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
	VerifyTransformedPattern pulumi.StringPtrInput
}

func (DatatypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*datatypeState)(nil)).Elem()
}

type datatypeArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Number of characters to obtain in advance for verification (1 - 255, default = 1).
	LookAhead *int `pulumi:"lookAhead"`
	// Number of characters required to save for verification (1 - 255, default = 1).
	LookBack *int `pulumi:"lookBack"`
	// Number of characters behind for match-around (1 - 4096, default = 1).
	MatchAhead *int `pulumi:"matchAhead"`
	// Dictionary to check whether it has a match around (Only support match-any and basic types, no repeat supported).
	MatchAround *string `pulumi:"matchAround"`
	// Number of characters in front for match-around (1 - 4096, default = 1).
	MatchBack *int `pulumi:"matchBack"`
	// Name of table containing the data type.
	Name *string `pulumi:"name"`
	// Regular expression pattern string without look around.
	Pattern *string `pulumi:"pattern"`
	// Template to transform user input to a pattern using capture group from 'pattern'.
	Transform *string `pulumi:"transform"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Regular expression pattern string used to verify the data type.
	Verify *string `pulumi:"verify"`
	// Extra regular expression pattern string used to verify the data type.
	Verify2 *string `pulumi:"verify2"`
	// Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
	VerifyTransformedPattern *string `pulumi:"verifyTransformedPattern"`
}

// The set of arguments for constructing a Datatype resource.
type DatatypeArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Number of characters to obtain in advance for verification (1 - 255, default = 1).
	LookAhead pulumi.IntPtrInput
	// Number of characters required to save for verification (1 - 255, default = 1).
	LookBack pulumi.IntPtrInput
	// Number of characters behind for match-around (1 - 4096, default = 1).
	MatchAhead pulumi.IntPtrInput
	// Dictionary to check whether it has a match around (Only support match-any and basic types, no repeat supported).
	MatchAround pulumi.StringPtrInput
	// Number of characters in front for match-around (1 - 4096, default = 1).
	MatchBack pulumi.IntPtrInput
	// Name of table containing the data type.
	Name pulumi.StringPtrInput
	// Regular expression pattern string without look around.
	Pattern pulumi.StringPtrInput
	// Template to transform user input to a pattern using capture group from 'pattern'.
	Transform pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Regular expression pattern string used to verify the data type.
	Verify pulumi.StringPtrInput
	// Extra regular expression pattern string used to verify the data type.
	Verify2 pulumi.StringPtrInput
	// Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
	VerifyTransformedPattern pulumi.StringPtrInput
}

func (DatatypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datatypeArgs)(nil)).Elem()
}

type DatatypeInput interface {
	pulumi.Input

	ToDatatypeOutput() DatatypeOutput
	ToDatatypeOutputWithContext(ctx context.Context) DatatypeOutput
}

func (*Datatype) ElementType() reflect.Type {
	return reflect.TypeOf((**Datatype)(nil)).Elem()
}

func (i *Datatype) ToDatatypeOutput() DatatypeOutput {
	return i.ToDatatypeOutputWithContext(context.Background())
}

func (i *Datatype) ToDatatypeOutputWithContext(ctx context.Context) DatatypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatatypeOutput)
}

// DatatypeArrayInput is an input type that accepts DatatypeArray and DatatypeArrayOutput values.
// You can construct a concrete instance of `DatatypeArrayInput` via:
//
//	DatatypeArray{ DatatypeArgs{...} }
type DatatypeArrayInput interface {
	pulumi.Input

	ToDatatypeArrayOutput() DatatypeArrayOutput
	ToDatatypeArrayOutputWithContext(context.Context) DatatypeArrayOutput
}

type DatatypeArray []DatatypeInput

func (DatatypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Datatype)(nil)).Elem()
}

func (i DatatypeArray) ToDatatypeArrayOutput() DatatypeArrayOutput {
	return i.ToDatatypeArrayOutputWithContext(context.Background())
}

func (i DatatypeArray) ToDatatypeArrayOutputWithContext(ctx context.Context) DatatypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatatypeArrayOutput)
}

// DatatypeMapInput is an input type that accepts DatatypeMap and DatatypeMapOutput values.
// You can construct a concrete instance of `DatatypeMapInput` via:
//
//	DatatypeMap{ "key": DatatypeArgs{...} }
type DatatypeMapInput interface {
	pulumi.Input

	ToDatatypeMapOutput() DatatypeMapOutput
	ToDatatypeMapOutputWithContext(context.Context) DatatypeMapOutput
}

type DatatypeMap map[string]DatatypeInput

func (DatatypeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Datatype)(nil)).Elem()
}

func (i DatatypeMap) ToDatatypeMapOutput() DatatypeMapOutput {
	return i.ToDatatypeMapOutputWithContext(context.Background())
}

func (i DatatypeMap) ToDatatypeMapOutputWithContext(ctx context.Context) DatatypeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatatypeMapOutput)
}

type DatatypeOutput struct{ *pulumi.OutputState }

func (DatatypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Datatype)(nil)).Elem()
}

func (o DatatypeOutput) ToDatatypeOutput() DatatypeOutput {
	return o
}

func (o DatatypeOutput) ToDatatypeOutputWithContext(ctx context.Context) DatatypeOutput {
	return o
}

// Optional comments.
func (o DatatypeOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datatype) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Number of characters to obtain in advance for verification (1 - 255, default = 1).
func (o DatatypeOutput) LookAhead() pulumi.IntOutput {
	return o.ApplyT(func(v *Datatype) pulumi.IntOutput { return v.LookAhead }).(pulumi.IntOutput)
}

// Number of characters required to save for verification (1 - 255, default = 1).
func (o DatatypeOutput) LookBack() pulumi.IntOutput {
	return o.ApplyT(func(v *Datatype) pulumi.IntOutput { return v.LookBack }).(pulumi.IntOutput)
}

// Number of characters behind for match-around (1 - 4096, default = 1).
func (o DatatypeOutput) MatchAhead() pulumi.IntOutput {
	return o.ApplyT(func(v *Datatype) pulumi.IntOutput { return v.MatchAhead }).(pulumi.IntOutput)
}

// Dictionary to check whether it has a match around (Only support match-any and basic types, no repeat supported).
func (o DatatypeOutput) MatchAround() pulumi.StringOutput {
	return o.ApplyT(func(v *Datatype) pulumi.StringOutput { return v.MatchAround }).(pulumi.StringOutput)
}

// Number of characters in front for match-around (1 - 4096, default = 1).
func (o DatatypeOutput) MatchBack() pulumi.IntOutput {
	return o.ApplyT(func(v *Datatype) pulumi.IntOutput { return v.MatchBack }).(pulumi.IntOutput)
}

// Name of table containing the data type.
func (o DatatypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Datatype) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Regular expression pattern string without look around.
func (o DatatypeOutput) Pattern() pulumi.StringOutput {
	return o.ApplyT(func(v *Datatype) pulumi.StringOutput { return v.Pattern }).(pulumi.StringOutput)
}

// Template to transform user input to a pattern using capture group from 'pattern'.
func (o DatatypeOutput) Transform() pulumi.StringOutput {
	return o.ApplyT(func(v *Datatype) pulumi.StringOutput { return v.Transform }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DatatypeOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datatype) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Regular expression pattern string used to verify the data type.
func (o DatatypeOutput) Verify() pulumi.StringOutput {
	return o.ApplyT(func(v *Datatype) pulumi.StringOutput { return v.Verify }).(pulumi.StringOutput)
}

// Extra regular expression pattern string used to verify the data type.
func (o DatatypeOutput) Verify2() pulumi.StringOutput {
	return o.ApplyT(func(v *Datatype) pulumi.StringOutput { return v.Verify2 }).(pulumi.StringOutput)
}

// Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
func (o DatatypeOutput) VerifyTransformedPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *Datatype) pulumi.StringOutput { return v.VerifyTransformedPattern }).(pulumi.StringOutput)
}

type DatatypeArrayOutput struct{ *pulumi.OutputState }

func (DatatypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Datatype)(nil)).Elem()
}

func (o DatatypeArrayOutput) ToDatatypeArrayOutput() DatatypeArrayOutput {
	return o
}

func (o DatatypeArrayOutput) ToDatatypeArrayOutputWithContext(ctx context.Context) DatatypeArrayOutput {
	return o
}

func (o DatatypeArrayOutput) Index(i pulumi.IntInput) DatatypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Datatype {
		return vs[0].([]*Datatype)[vs[1].(int)]
	}).(DatatypeOutput)
}

type DatatypeMapOutput struct{ *pulumi.OutputState }

func (DatatypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Datatype)(nil)).Elem()
}

func (o DatatypeMapOutput) ToDatatypeMapOutput() DatatypeMapOutput {
	return o
}

func (o DatatypeMapOutput) ToDatatypeMapOutputWithContext(ctx context.Context) DatatypeMapOutput {
	return o
}

func (o DatatypeMapOutput) MapIndex(k pulumi.StringInput) DatatypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Datatype {
		return vs[0].(map[string]*Datatype)[vs[1].(string)]
	}).(DatatypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatatypeInput)(nil)).Elem(), &Datatype{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatatypeArrayInput)(nil)).Elem(), DatatypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatatypeMapInput)(nil)).Elem(), DatatypeMap{})
	pulumi.RegisterOutputType(DatatypeOutput{})
	pulumi.RegisterOutputType(DatatypeArrayOutput{})
	pulumi.RegisterOutputType(DatatypeMapOutput{})
}
