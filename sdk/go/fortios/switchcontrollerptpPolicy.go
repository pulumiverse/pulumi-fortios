// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// PTP policy configuration. Applies to FortiOS Version `>= 6.4.2`.
//
// ## Import
//
// # SwitchControllerPtp Policy can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerptpPolicy:SwitchcontrollerptpPolicy labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerptpPolicy:SwitchcontrollerptpPolicy labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SwitchcontrollerptpPolicy struct {
	pulumi.CustomResourceState

	// Policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable PTP policy. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchcontrollerptpPolicy registers a new resource with the given unique name, arguments, and options.
func NewSwitchcontrollerptpPolicy(ctx *pulumi.Context,
	name string, args *SwitchcontrollerptpPolicyArgs, opts ...pulumi.ResourceOption) (*SwitchcontrollerptpPolicy, error) {
	if args == nil {
		args = &SwitchcontrollerptpPolicyArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchcontrollerptpPolicy
	err := ctx.RegisterResource("fortios:index/switchcontrollerptpPolicy:SwitchcontrollerptpPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchcontrollerptpPolicy gets an existing SwitchcontrollerptpPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchcontrollerptpPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchcontrollerptpPolicyState, opts ...pulumi.ResourceOption) (*SwitchcontrollerptpPolicy, error) {
	var resource SwitchcontrollerptpPolicy
	err := ctx.ReadResource("fortios:index/switchcontrollerptpPolicy:SwitchcontrollerptpPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchcontrollerptpPolicy resources.
type switchcontrollerptpPolicyState struct {
	// Policy name.
	Name *string `pulumi:"name"`
	// Enable/disable PTP policy. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchcontrollerptpPolicyState struct {
	// Policy name.
	Name pulumi.StringPtrInput
	// Enable/disable PTP policy. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchcontrollerptpPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerptpPolicyState)(nil)).Elem()
}

type switchcontrollerptpPolicyArgs struct {
	// Policy name.
	Name *string `pulumi:"name"`
	// Enable/disable PTP policy. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchcontrollerptpPolicy resource.
type SwitchcontrollerptpPolicyArgs struct {
	// Policy name.
	Name pulumi.StringPtrInput
	// Enable/disable PTP policy. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchcontrollerptpPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerptpPolicyArgs)(nil)).Elem()
}

type SwitchcontrollerptpPolicyInput interface {
	pulumi.Input

	ToSwitchcontrollerptpPolicyOutput() SwitchcontrollerptpPolicyOutput
	ToSwitchcontrollerptpPolicyOutputWithContext(ctx context.Context) SwitchcontrollerptpPolicyOutput
}

func (*SwitchcontrollerptpPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerptpPolicy)(nil)).Elem()
}

func (i *SwitchcontrollerptpPolicy) ToSwitchcontrollerptpPolicyOutput() SwitchcontrollerptpPolicyOutput {
	return i.ToSwitchcontrollerptpPolicyOutputWithContext(context.Background())
}

func (i *SwitchcontrollerptpPolicy) ToSwitchcontrollerptpPolicyOutputWithContext(ctx context.Context) SwitchcontrollerptpPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerptpPolicyOutput)
}

// SwitchcontrollerptpPolicyArrayInput is an input type that accepts SwitchcontrollerptpPolicyArray and SwitchcontrollerptpPolicyArrayOutput values.
// You can construct a concrete instance of `SwitchcontrollerptpPolicyArrayInput` via:
//
//	SwitchcontrollerptpPolicyArray{ SwitchcontrollerptpPolicyArgs{...} }
type SwitchcontrollerptpPolicyArrayInput interface {
	pulumi.Input

	ToSwitchcontrollerptpPolicyArrayOutput() SwitchcontrollerptpPolicyArrayOutput
	ToSwitchcontrollerptpPolicyArrayOutputWithContext(context.Context) SwitchcontrollerptpPolicyArrayOutput
}

type SwitchcontrollerptpPolicyArray []SwitchcontrollerptpPolicyInput

func (SwitchcontrollerptpPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerptpPolicy)(nil)).Elem()
}

func (i SwitchcontrollerptpPolicyArray) ToSwitchcontrollerptpPolicyArrayOutput() SwitchcontrollerptpPolicyArrayOutput {
	return i.ToSwitchcontrollerptpPolicyArrayOutputWithContext(context.Background())
}

func (i SwitchcontrollerptpPolicyArray) ToSwitchcontrollerptpPolicyArrayOutputWithContext(ctx context.Context) SwitchcontrollerptpPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerptpPolicyArrayOutput)
}

// SwitchcontrollerptpPolicyMapInput is an input type that accepts SwitchcontrollerptpPolicyMap and SwitchcontrollerptpPolicyMapOutput values.
// You can construct a concrete instance of `SwitchcontrollerptpPolicyMapInput` via:
//
//	SwitchcontrollerptpPolicyMap{ "key": SwitchcontrollerptpPolicyArgs{...} }
type SwitchcontrollerptpPolicyMapInput interface {
	pulumi.Input

	ToSwitchcontrollerptpPolicyMapOutput() SwitchcontrollerptpPolicyMapOutput
	ToSwitchcontrollerptpPolicyMapOutputWithContext(context.Context) SwitchcontrollerptpPolicyMapOutput
}

type SwitchcontrollerptpPolicyMap map[string]SwitchcontrollerptpPolicyInput

func (SwitchcontrollerptpPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerptpPolicy)(nil)).Elem()
}

func (i SwitchcontrollerptpPolicyMap) ToSwitchcontrollerptpPolicyMapOutput() SwitchcontrollerptpPolicyMapOutput {
	return i.ToSwitchcontrollerptpPolicyMapOutputWithContext(context.Background())
}

func (i SwitchcontrollerptpPolicyMap) ToSwitchcontrollerptpPolicyMapOutputWithContext(ctx context.Context) SwitchcontrollerptpPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerptpPolicyMapOutput)
}

type SwitchcontrollerptpPolicyOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerptpPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerptpPolicy)(nil)).Elem()
}

func (o SwitchcontrollerptpPolicyOutput) ToSwitchcontrollerptpPolicyOutput() SwitchcontrollerptpPolicyOutput {
	return o
}

func (o SwitchcontrollerptpPolicyOutput) ToSwitchcontrollerptpPolicyOutputWithContext(ctx context.Context) SwitchcontrollerptpPolicyOutput {
	return o
}

// Policy name.
func (o SwitchcontrollerptpPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerptpPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable PTP policy. Valid values: `disable`, `enable`.
func (o SwitchcontrollerptpPolicyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerptpPolicy) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SwitchcontrollerptpPolicyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchcontrollerptpPolicy) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SwitchcontrollerptpPolicyArrayOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerptpPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerptpPolicy)(nil)).Elem()
}

func (o SwitchcontrollerptpPolicyArrayOutput) ToSwitchcontrollerptpPolicyArrayOutput() SwitchcontrollerptpPolicyArrayOutput {
	return o
}

func (o SwitchcontrollerptpPolicyArrayOutput) ToSwitchcontrollerptpPolicyArrayOutputWithContext(ctx context.Context) SwitchcontrollerptpPolicyArrayOutput {
	return o
}

func (o SwitchcontrollerptpPolicyArrayOutput) Index(i pulumi.IntInput) SwitchcontrollerptpPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchcontrollerptpPolicy {
		return vs[0].([]*SwitchcontrollerptpPolicy)[vs[1].(int)]
	}).(SwitchcontrollerptpPolicyOutput)
}

type SwitchcontrollerptpPolicyMapOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerptpPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerptpPolicy)(nil)).Elem()
}

func (o SwitchcontrollerptpPolicyMapOutput) ToSwitchcontrollerptpPolicyMapOutput() SwitchcontrollerptpPolicyMapOutput {
	return o
}

func (o SwitchcontrollerptpPolicyMapOutput) ToSwitchcontrollerptpPolicyMapOutputWithContext(ctx context.Context) SwitchcontrollerptpPolicyMapOutput {
	return o
}

func (o SwitchcontrollerptpPolicyMapOutput) MapIndex(k pulumi.StringInput) SwitchcontrollerptpPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchcontrollerptpPolicy {
		return vs[0].(map[string]*SwitchcontrollerptpPolicy)[vs[1].(string)]
	}).(SwitchcontrollerptpPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerptpPolicyInput)(nil)).Elem(), &SwitchcontrollerptpPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerptpPolicyArrayInput)(nil)).Elem(), SwitchcontrollerptpPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerptpPolicyMapInput)(nil)).Elem(), SwitchcontrollerptpPolicyMap{})
	pulumi.RegisterOutputType(SwitchcontrollerptpPolicyOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerptpPolicyArrayOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerptpPolicyMapOutput{})
}
