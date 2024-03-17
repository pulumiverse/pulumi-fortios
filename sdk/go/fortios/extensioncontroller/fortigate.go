// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package extensioncontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// FortiGate controller configuration. Applies to FortiOS Version `>= 7.2.1`.
//
// ## Import
//
// ExtensionController Fortigate can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:extensioncontroller/fortigate:Fortigate labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:extensioncontroller/fortigate:Fortigate labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Fortigate struct {
	pulumi.CustomResourceState

	// Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
	Authorized pulumi.StringOutput `pulumi:"authorized"`
	// Description.
	Description pulumi.StringOutput `pulumi:"description"`
	// device-id
	DeviceId pulumi.IntOutput `pulumi:"deviceId"`
	// FortiGate serial number.
	Fosid pulumi.StringOutput `pulumi:"fosid"`
	// FortiGate hostname.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// FortiGate entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// FortiGate profile configuration.
	Profile pulumi.StringOutput `pulumi:"profile"`
	// VDOM.
	Vdom pulumi.IntOutput `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFortigate registers a new resource with the given unique name, arguments, and options.
func NewFortigate(ctx *pulumi.Context,
	name string, args *FortigateArgs, opts ...pulumi.ResourceOption) (*Fortigate, error) {
	if args == nil {
		args = &FortigateArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Fortigate
	err := ctx.RegisterResource("fortios:extensioncontroller/fortigate:Fortigate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortigate gets an existing Fortigate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortigate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortigateState, opts ...pulumi.ResourceOption) (*Fortigate, error) {
	var resource Fortigate
	err := ctx.ReadResource("fortios:extensioncontroller/fortigate:Fortigate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fortigate resources.
type fortigateState struct {
	// Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
	Authorized *string `pulumi:"authorized"`
	// Description.
	Description *string `pulumi:"description"`
	// device-id
	DeviceId *int `pulumi:"deviceId"`
	// FortiGate serial number.
	Fosid *string `pulumi:"fosid"`
	// FortiGate hostname.
	Hostname *string `pulumi:"hostname"`
	// FortiGate entry name.
	Name *string `pulumi:"name"`
	// FortiGate profile configuration.
	Profile *string `pulumi:"profile"`
	// VDOM.
	Vdom *int `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FortigateState struct {
	// Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
	Authorized pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// device-id
	DeviceId pulumi.IntPtrInput
	// FortiGate serial number.
	Fosid pulumi.StringPtrInput
	// FortiGate hostname.
	Hostname pulumi.StringPtrInput
	// FortiGate entry name.
	Name pulumi.StringPtrInput
	// FortiGate profile configuration.
	Profile pulumi.StringPtrInput
	// VDOM.
	Vdom pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FortigateState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortigateState)(nil)).Elem()
}

type fortigateArgs struct {
	// Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
	Authorized *string `pulumi:"authorized"`
	// Description.
	Description *string `pulumi:"description"`
	// device-id
	DeviceId *int `pulumi:"deviceId"`
	// FortiGate serial number.
	Fosid *string `pulumi:"fosid"`
	// FortiGate hostname.
	Hostname *string `pulumi:"hostname"`
	// FortiGate entry name.
	Name *string `pulumi:"name"`
	// FortiGate profile configuration.
	Profile *string `pulumi:"profile"`
	// VDOM.
	Vdom *int `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Fortigate resource.
type FortigateArgs struct {
	// Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
	Authorized pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// device-id
	DeviceId pulumi.IntPtrInput
	// FortiGate serial number.
	Fosid pulumi.StringPtrInput
	// FortiGate hostname.
	Hostname pulumi.StringPtrInput
	// FortiGate entry name.
	Name pulumi.StringPtrInput
	// FortiGate profile configuration.
	Profile pulumi.StringPtrInput
	// VDOM.
	Vdom pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FortigateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortigateArgs)(nil)).Elem()
}

type FortigateInput interface {
	pulumi.Input

	ToFortigateOutput() FortigateOutput
	ToFortigateOutputWithContext(ctx context.Context) FortigateOutput
}

func (*Fortigate) ElementType() reflect.Type {
	return reflect.TypeOf((**Fortigate)(nil)).Elem()
}

func (i *Fortigate) ToFortigateOutput() FortigateOutput {
	return i.ToFortigateOutputWithContext(context.Background())
}

func (i *Fortigate) ToFortigateOutputWithContext(ctx context.Context) FortigateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortigateOutput)
}

// FortigateArrayInput is an input type that accepts FortigateArray and FortigateArrayOutput values.
// You can construct a concrete instance of `FortigateArrayInput` via:
//
//	FortigateArray{ FortigateArgs{...} }
type FortigateArrayInput interface {
	pulumi.Input

	ToFortigateArrayOutput() FortigateArrayOutput
	ToFortigateArrayOutputWithContext(context.Context) FortigateArrayOutput
}

type FortigateArray []FortigateInput

func (FortigateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fortigate)(nil)).Elem()
}

func (i FortigateArray) ToFortigateArrayOutput() FortigateArrayOutput {
	return i.ToFortigateArrayOutputWithContext(context.Background())
}

func (i FortigateArray) ToFortigateArrayOutputWithContext(ctx context.Context) FortigateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortigateArrayOutput)
}

// FortigateMapInput is an input type that accepts FortigateMap and FortigateMapOutput values.
// You can construct a concrete instance of `FortigateMapInput` via:
//
//	FortigateMap{ "key": FortigateArgs{...} }
type FortigateMapInput interface {
	pulumi.Input

	ToFortigateMapOutput() FortigateMapOutput
	ToFortigateMapOutputWithContext(context.Context) FortigateMapOutput
}

type FortigateMap map[string]FortigateInput

func (FortigateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fortigate)(nil)).Elem()
}

func (i FortigateMap) ToFortigateMapOutput() FortigateMapOutput {
	return i.ToFortigateMapOutputWithContext(context.Background())
}

func (i FortigateMap) ToFortigateMapOutputWithContext(ctx context.Context) FortigateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortigateMapOutput)
}

type FortigateOutput struct{ *pulumi.OutputState }

func (FortigateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fortigate)(nil)).Elem()
}

func (o FortigateOutput) ToFortigateOutput() FortigateOutput {
	return o
}

func (o FortigateOutput) ToFortigateOutputWithContext(ctx context.Context) FortigateOutput {
	return o
}

// Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
func (o FortigateOutput) Authorized() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortigate) pulumi.StringOutput { return v.Authorized }).(pulumi.StringOutput)
}

// Description.
func (o FortigateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortigate) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// device-id
func (o FortigateOutput) DeviceId() pulumi.IntOutput {
	return o.ApplyT(func(v *Fortigate) pulumi.IntOutput { return v.DeviceId }).(pulumi.IntOutput)
}

// FortiGate serial number.
func (o FortigateOutput) Fosid() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortigate) pulumi.StringOutput { return v.Fosid }).(pulumi.StringOutput)
}

// FortiGate hostname.
func (o FortigateOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortigate) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// FortiGate entry name.
func (o FortigateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortigate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// FortiGate profile configuration.
func (o FortigateOutput) Profile() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortigate) pulumi.StringOutput { return v.Profile }).(pulumi.StringOutput)
}

// VDOM.
func (o FortigateOutput) Vdom() pulumi.IntOutput {
	return o.ApplyT(func(v *Fortigate) pulumi.IntOutput { return v.Vdom }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FortigateOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Fortigate) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FortigateArrayOutput struct{ *pulumi.OutputState }

func (FortigateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fortigate)(nil)).Elem()
}

func (o FortigateArrayOutput) ToFortigateArrayOutput() FortigateArrayOutput {
	return o
}

func (o FortigateArrayOutput) ToFortigateArrayOutputWithContext(ctx context.Context) FortigateArrayOutput {
	return o
}

func (o FortigateArrayOutput) Index(i pulumi.IntInput) FortigateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Fortigate {
		return vs[0].([]*Fortigate)[vs[1].(int)]
	}).(FortigateOutput)
}

type FortigateMapOutput struct{ *pulumi.OutputState }

func (FortigateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fortigate)(nil)).Elem()
}

func (o FortigateMapOutput) ToFortigateMapOutput() FortigateMapOutput {
	return o
}

func (o FortigateMapOutput) ToFortigateMapOutputWithContext(ctx context.Context) FortigateMapOutput {
	return o
}

func (o FortigateMapOutput) MapIndex(k pulumi.StringInput) FortigateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Fortigate {
		return vs[0].(map[string]*Fortigate)[vs[1].(string)]
	}).(FortigateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortigateInput)(nil)).Elem(), &Fortigate{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortigateArrayInput)(nil)).Elem(), FortigateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortigateMapInput)(nil)).Elem(), FortigateMap{})
	pulumi.RegisterOutputType(FortigateOutput{})
	pulumi.RegisterOutputType(FortigateArrayOutput{})
	pulumi.RegisterOutputType(FortigateMapOutput{})
}
