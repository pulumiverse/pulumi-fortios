// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create self-explanatory DLP sensitivity levels to be used when setting sensitivity under config fp-doc-source. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// # Dlp Sensitivity can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/dlpSensitivity:DlpSensitivity labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/dlpSensitivity:DlpSensitivity labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type DlpSensitivity struct {
	pulumi.CustomResourceState

	// DLP Sensitivity Levels.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDlpSensitivity registers a new resource with the given unique name, arguments, and options.
func NewDlpSensitivity(ctx *pulumi.Context,
	name string, args *DlpSensitivityArgs, opts ...pulumi.ResourceOption) (*DlpSensitivity, error) {
	if args == nil {
		args = &DlpSensitivityArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource DlpSensitivity
	err := ctx.RegisterResource("fortios:index/dlpSensitivity:DlpSensitivity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDlpSensitivity gets an existing DlpSensitivity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDlpSensitivity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DlpSensitivityState, opts ...pulumi.ResourceOption) (*DlpSensitivity, error) {
	var resource DlpSensitivity
	err := ctx.ReadResource("fortios:index/dlpSensitivity:DlpSensitivity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DlpSensitivity resources.
type dlpSensitivityState struct {
	// DLP Sensitivity Levels.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DlpSensitivityState struct {
	// DLP Sensitivity Levels.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DlpSensitivityState) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpSensitivityState)(nil)).Elem()
}

type dlpSensitivityArgs struct {
	// DLP Sensitivity Levels.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a DlpSensitivity resource.
type DlpSensitivityArgs struct {
	// DLP Sensitivity Levels.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DlpSensitivityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpSensitivityArgs)(nil)).Elem()
}

type DlpSensitivityInput interface {
	pulumi.Input

	ToDlpSensitivityOutput() DlpSensitivityOutput
	ToDlpSensitivityOutputWithContext(ctx context.Context) DlpSensitivityOutput
}

func (*DlpSensitivity) ElementType() reflect.Type {
	return reflect.TypeOf((**DlpSensitivity)(nil)).Elem()
}

func (i *DlpSensitivity) ToDlpSensitivityOutput() DlpSensitivityOutput {
	return i.ToDlpSensitivityOutputWithContext(context.Background())
}

func (i *DlpSensitivity) ToDlpSensitivityOutputWithContext(ctx context.Context) DlpSensitivityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpSensitivityOutput)
}

// DlpSensitivityArrayInput is an input type that accepts DlpSensitivityArray and DlpSensitivityArrayOutput values.
// You can construct a concrete instance of `DlpSensitivityArrayInput` via:
//
//	DlpSensitivityArray{ DlpSensitivityArgs{...} }
type DlpSensitivityArrayInput interface {
	pulumi.Input

	ToDlpSensitivityArrayOutput() DlpSensitivityArrayOutput
	ToDlpSensitivityArrayOutputWithContext(context.Context) DlpSensitivityArrayOutput
}

type DlpSensitivityArray []DlpSensitivityInput

func (DlpSensitivityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DlpSensitivity)(nil)).Elem()
}

func (i DlpSensitivityArray) ToDlpSensitivityArrayOutput() DlpSensitivityArrayOutput {
	return i.ToDlpSensitivityArrayOutputWithContext(context.Background())
}

func (i DlpSensitivityArray) ToDlpSensitivityArrayOutputWithContext(ctx context.Context) DlpSensitivityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpSensitivityArrayOutput)
}

// DlpSensitivityMapInput is an input type that accepts DlpSensitivityMap and DlpSensitivityMapOutput values.
// You can construct a concrete instance of `DlpSensitivityMapInput` via:
//
//	DlpSensitivityMap{ "key": DlpSensitivityArgs{...} }
type DlpSensitivityMapInput interface {
	pulumi.Input

	ToDlpSensitivityMapOutput() DlpSensitivityMapOutput
	ToDlpSensitivityMapOutputWithContext(context.Context) DlpSensitivityMapOutput
}

type DlpSensitivityMap map[string]DlpSensitivityInput

func (DlpSensitivityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DlpSensitivity)(nil)).Elem()
}

func (i DlpSensitivityMap) ToDlpSensitivityMapOutput() DlpSensitivityMapOutput {
	return i.ToDlpSensitivityMapOutputWithContext(context.Background())
}

func (i DlpSensitivityMap) ToDlpSensitivityMapOutputWithContext(ctx context.Context) DlpSensitivityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpSensitivityMapOutput)
}

type DlpSensitivityOutput struct{ *pulumi.OutputState }

func (DlpSensitivityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DlpSensitivity)(nil)).Elem()
}

func (o DlpSensitivityOutput) ToDlpSensitivityOutput() DlpSensitivityOutput {
	return o
}

func (o DlpSensitivityOutput) ToDlpSensitivityOutputWithContext(ctx context.Context) DlpSensitivityOutput {
	return o
}

// DLP Sensitivity Levels.
func (o DlpSensitivityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DlpSensitivity) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DlpSensitivityOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DlpSensitivity) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type DlpSensitivityArrayOutput struct{ *pulumi.OutputState }

func (DlpSensitivityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DlpSensitivity)(nil)).Elem()
}

func (o DlpSensitivityArrayOutput) ToDlpSensitivityArrayOutput() DlpSensitivityArrayOutput {
	return o
}

func (o DlpSensitivityArrayOutput) ToDlpSensitivityArrayOutputWithContext(ctx context.Context) DlpSensitivityArrayOutput {
	return o
}

func (o DlpSensitivityArrayOutput) Index(i pulumi.IntInput) DlpSensitivityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DlpSensitivity {
		return vs[0].([]*DlpSensitivity)[vs[1].(int)]
	}).(DlpSensitivityOutput)
}

type DlpSensitivityMapOutput struct{ *pulumi.OutputState }

func (DlpSensitivityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DlpSensitivity)(nil)).Elem()
}

func (o DlpSensitivityMapOutput) ToDlpSensitivityMapOutput() DlpSensitivityMapOutput {
	return o
}

func (o DlpSensitivityMapOutput) ToDlpSensitivityMapOutputWithContext(ctx context.Context) DlpSensitivityMapOutput {
	return o
}

func (o DlpSensitivityMapOutput) MapIndex(k pulumi.StringInput) DlpSensitivityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DlpSensitivity {
		return vs[0].(map[string]*DlpSensitivity)[vs[1].(string)]
	}).(DlpSensitivityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DlpSensitivityInput)(nil)).Elem(), &DlpSensitivity{})
	pulumi.RegisterInputType(reflect.TypeOf((*DlpSensitivityArrayInput)(nil)).Elem(), DlpSensitivityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DlpSensitivityMapInput)(nil)).Elem(), DlpSensitivityMap{})
	pulumi.RegisterOutputType(DlpSensitivityOutput{})
	pulumi.RegisterOutputType(DlpSensitivityArrayOutput{})
	pulumi.RegisterOutputType(DlpSensitivityMapOutput{})
}
