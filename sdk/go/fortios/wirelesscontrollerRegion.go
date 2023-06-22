// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiAP regions (for floor plans and maps).
//
// ## Import
//
// # WirelessController Region can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerRegion:WirelesscontrollerRegion labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerRegion:WirelesscontrollerRegion labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type WirelesscontrollerRegion struct {
	pulumi.CustomResourceState

	// Comments.
	Comments pulumi.StringOutput `pulumi:"comments"`
	// Region image grayscale. Valid values: `enable`, `disable`.
	Grayscale pulumi.StringOutput `pulumi:"grayscale"`
	// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
	ImageType pulumi.StringOutput `pulumi:"imageType"`
	// FortiAP region name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Region image opacity (0 - 100).
	Opacity pulumi.IntOutput `pulumi:"opacity"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelesscontrollerRegion registers a new resource with the given unique name, arguments, and options.
func NewWirelesscontrollerRegion(ctx *pulumi.Context,
	name string, args *WirelesscontrollerRegionArgs, opts ...pulumi.ResourceOption) (*WirelesscontrollerRegion, error) {
	if args == nil {
		args = &WirelesscontrollerRegionArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WirelesscontrollerRegion
	err := ctx.RegisterResource("fortios:index/wirelesscontrollerRegion:WirelesscontrollerRegion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelesscontrollerRegion gets an existing WirelesscontrollerRegion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelesscontrollerRegion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelesscontrollerRegionState, opts ...pulumi.ResourceOption) (*WirelesscontrollerRegion, error) {
	var resource WirelesscontrollerRegion
	err := ctx.ReadResource("fortios:index/wirelesscontrollerRegion:WirelesscontrollerRegion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelesscontrollerRegion resources.
type wirelesscontrollerRegionState struct {
	// Comments.
	Comments *string `pulumi:"comments"`
	// Region image grayscale. Valid values: `enable`, `disable`.
	Grayscale *string `pulumi:"grayscale"`
	// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
	ImageType *string `pulumi:"imageType"`
	// FortiAP region name.
	Name *string `pulumi:"name"`
	// Region image opacity (0 - 100).
	Opacity *int `pulumi:"opacity"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WirelesscontrollerRegionState struct {
	// Comments.
	Comments pulumi.StringPtrInput
	// Region image grayscale. Valid values: `enable`, `disable`.
	Grayscale pulumi.StringPtrInput
	// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
	ImageType pulumi.StringPtrInput
	// FortiAP region name.
	Name pulumi.StringPtrInput
	// Region image opacity (0 - 100).
	Opacity pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelesscontrollerRegionState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerRegionState)(nil)).Elem()
}

type wirelesscontrollerRegionArgs struct {
	// Comments.
	Comments *string `pulumi:"comments"`
	// Region image grayscale. Valid values: `enable`, `disable`.
	Grayscale *string `pulumi:"grayscale"`
	// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
	ImageType *string `pulumi:"imageType"`
	// FortiAP region name.
	Name *string `pulumi:"name"`
	// Region image opacity (0 - 100).
	Opacity *int `pulumi:"opacity"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelesscontrollerRegion resource.
type WirelesscontrollerRegionArgs struct {
	// Comments.
	Comments pulumi.StringPtrInput
	// Region image grayscale. Valid values: `enable`, `disable`.
	Grayscale pulumi.StringPtrInput
	// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
	ImageType pulumi.StringPtrInput
	// FortiAP region name.
	Name pulumi.StringPtrInput
	// Region image opacity (0 - 100).
	Opacity pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelesscontrollerRegionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerRegionArgs)(nil)).Elem()
}

type WirelesscontrollerRegionInput interface {
	pulumi.Input

	ToWirelesscontrollerRegionOutput() WirelesscontrollerRegionOutput
	ToWirelesscontrollerRegionOutputWithContext(ctx context.Context) WirelesscontrollerRegionOutput
}

func (*WirelesscontrollerRegion) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelesscontrollerRegion)(nil)).Elem()
}

func (i *WirelesscontrollerRegion) ToWirelesscontrollerRegionOutput() WirelesscontrollerRegionOutput {
	return i.ToWirelesscontrollerRegionOutputWithContext(context.Background())
}

func (i *WirelesscontrollerRegion) ToWirelesscontrollerRegionOutputWithContext(ctx context.Context) WirelesscontrollerRegionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerRegionOutput)
}

// WirelesscontrollerRegionArrayInput is an input type that accepts WirelesscontrollerRegionArray and WirelesscontrollerRegionArrayOutput values.
// You can construct a concrete instance of `WirelesscontrollerRegionArrayInput` via:
//
//	WirelesscontrollerRegionArray{ WirelesscontrollerRegionArgs{...} }
type WirelesscontrollerRegionArrayInput interface {
	pulumi.Input

	ToWirelesscontrollerRegionArrayOutput() WirelesscontrollerRegionArrayOutput
	ToWirelesscontrollerRegionArrayOutputWithContext(context.Context) WirelesscontrollerRegionArrayOutput
}

type WirelesscontrollerRegionArray []WirelesscontrollerRegionInput

func (WirelesscontrollerRegionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelesscontrollerRegion)(nil)).Elem()
}

func (i WirelesscontrollerRegionArray) ToWirelesscontrollerRegionArrayOutput() WirelesscontrollerRegionArrayOutput {
	return i.ToWirelesscontrollerRegionArrayOutputWithContext(context.Background())
}

func (i WirelesscontrollerRegionArray) ToWirelesscontrollerRegionArrayOutputWithContext(ctx context.Context) WirelesscontrollerRegionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerRegionArrayOutput)
}

// WirelesscontrollerRegionMapInput is an input type that accepts WirelesscontrollerRegionMap and WirelesscontrollerRegionMapOutput values.
// You can construct a concrete instance of `WirelesscontrollerRegionMapInput` via:
//
//	WirelesscontrollerRegionMap{ "key": WirelesscontrollerRegionArgs{...} }
type WirelesscontrollerRegionMapInput interface {
	pulumi.Input

	ToWirelesscontrollerRegionMapOutput() WirelesscontrollerRegionMapOutput
	ToWirelesscontrollerRegionMapOutputWithContext(context.Context) WirelesscontrollerRegionMapOutput
}

type WirelesscontrollerRegionMap map[string]WirelesscontrollerRegionInput

func (WirelesscontrollerRegionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelesscontrollerRegion)(nil)).Elem()
}

func (i WirelesscontrollerRegionMap) ToWirelesscontrollerRegionMapOutput() WirelesscontrollerRegionMapOutput {
	return i.ToWirelesscontrollerRegionMapOutputWithContext(context.Background())
}

func (i WirelesscontrollerRegionMap) ToWirelesscontrollerRegionMapOutputWithContext(ctx context.Context) WirelesscontrollerRegionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerRegionMapOutput)
}

type WirelesscontrollerRegionOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerRegionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelesscontrollerRegion)(nil)).Elem()
}

func (o WirelesscontrollerRegionOutput) ToWirelesscontrollerRegionOutput() WirelesscontrollerRegionOutput {
	return o
}

func (o WirelesscontrollerRegionOutput) ToWirelesscontrollerRegionOutputWithContext(ctx context.Context) WirelesscontrollerRegionOutput {
	return o
}

// Comments.
func (o WirelesscontrollerRegionOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerRegion) pulumi.StringOutput { return v.Comments }).(pulumi.StringOutput)
}

// Region image grayscale. Valid values: `enable`, `disable`.
func (o WirelesscontrollerRegionOutput) Grayscale() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerRegion) pulumi.StringOutput { return v.Grayscale }).(pulumi.StringOutput)
}

// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
func (o WirelesscontrollerRegionOutput) ImageType() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerRegion) pulumi.StringOutput { return v.ImageType }).(pulumi.StringOutput)
}

// FortiAP region name.
func (o WirelesscontrollerRegionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerRegion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Region image opacity (0 - 100).
func (o WirelesscontrollerRegionOutput) Opacity() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelesscontrollerRegion) pulumi.IntOutput { return v.Opacity }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o WirelesscontrollerRegionOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelesscontrollerRegion) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WirelesscontrollerRegionArrayOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerRegionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelesscontrollerRegion)(nil)).Elem()
}

func (o WirelesscontrollerRegionArrayOutput) ToWirelesscontrollerRegionArrayOutput() WirelesscontrollerRegionArrayOutput {
	return o
}

func (o WirelesscontrollerRegionArrayOutput) ToWirelesscontrollerRegionArrayOutputWithContext(ctx context.Context) WirelesscontrollerRegionArrayOutput {
	return o
}

func (o WirelesscontrollerRegionArrayOutput) Index(i pulumi.IntInput) WirelesscontrollerRegionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelesscontrollerRegion {
		return vs[0].([]*WirelesscontrollerRegion)[vs[1].(int)]
	}).(WirelesscontrollerRegionOutput)
}

type WirelesscontrollerRegionMapOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerRegionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelesscontrollerRegion)(nil)).Elem()
}

func (o WirelesscontrollerRegionMapOutput) ToWirelesscontrollerRegionMapOutput() WirelesscontrollerRegionMapOutput {
	return o
}

func (o WirelesscontrollerRegionMapOutput) ToWirelesscontrollerRegionMapOutputWithContext(ctx context.Context) WirelesscontrollerRegionMapOutput {
	return o
}

func (o WirelesscontrollerRegionMapOutput) MapIndex(k pulumi.StringInput) WirelesscontrollerRegionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelesscontrollerRegion {
		return vs[0].(map[string]*WirelesscontrollerRegion)[vs[1].(string)]
	}).(WirelesscontrollerRegionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerRegionInput)(nil)).Elem(), &WirelesscontrollerRegion{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerRegionArrayInput)(nil)).Elem(), WirelesscontrollerRegionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerRegionMapInput)(nil)).Elem(), WirelesscontrollerRegionMap{})
	pulumi.RegisterOutputType(WirelesscontrollerRegionOutput{})
	pulumi.RegisterOutputType(WirelesscontrollerRegionArrayOutput{})
	pulumi.RegisterOutputType(WirelesscontrollerRegionMapOutput{})
}
