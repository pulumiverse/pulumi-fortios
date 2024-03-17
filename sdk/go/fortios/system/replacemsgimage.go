// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure replacement message images.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/system"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := system.NewReplacemsgimage(ctx, "trname", &system.ReplacemsgimageArgs{
//				ImageBase64: pulumi.String("iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAAEWAAABFgAVshLGQAAAAMSURBVBhXY/j//z8ABf4C/qc1gYQAAAAASUVORK5CYII="),
//				ImageType:   pulumi.String("png"),
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
// System ReplacemsgImage can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/replacemsgimage:Replacemsgimage labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/replacemsgimage:Replacemsgimage labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Replacemsgimage struct {
	pulumi.CustomResourceState

	// Image data.
	ImageBase64 pulumi.StringPtrOutput `pulumi:"imageBase64"`
	// Image type. Valid values: `gif`, `jpg`, `tiff`, `png`.
	ImageType pulumi.StringOutput `pulumi:"imageType"`
	// Image name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewReplacemsgimage registers a new resource with the given unique name, arguments, and options.
func NewReplacemsgimage(ctx *pulumi.Context,
	name string, args *ReplacemsgimageArgs, opts ...pulumi.ResourceOption) (*Replacemsgimage, error) {
	if args == nil {
		args = &ReplacemsgimageArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Replacemsgimage
	err := ctx.RegisterResource("fortios:system/replacemsgimage:Replacemsgimage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplacemsgimage gets an existing Replacemsgimage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplacemsgimage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplacemsgimageState, opts ...pulumi.ResourceOption) (*Replacemsgimage, error) {
	var resource Replacemsgimage
	err := ctx.ReadResource("fortios:system/replacemsgimage:Replacemsgimage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Replacemsgimage resources.
type replacemsgimageState struct {
	// Image data.
	ImageBase64 *string `pulumi:"imageBase64"`
	// Image type. Valid values: `gif`, `jpg`, `tiff`, `png`.
	ImageType *string `pulumi:"imageType"`
	// Image name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ReplacemsgimageState struct {
	// Image data.
	ImageBase64 pulumi.StringPtrInput
	// Image type. Valid values: `gif`, `jpg`, `tiff`, `png`.
	ImageType pulumi.StringPtrInput
	// Image name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ReplacemsgimageState) ElementType() reflect.Type {
	return reflect.TypeOf((*replacemsgimageState)(nil)).Elem()
}

type replacemsgimageArgs struct {
	// Image data.
	ImageBase64 *string `pulumi:"imageBase64"`
	// Image type. Valid values: `gif`, `jpg`, `tiff`, `png`.
	ImageType *string `pulumi:"imageType"`
	// Image name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Replacemsgimage resource.
type ReplacemsgimageArgs struct {
	// Image data.
	ImageBase64 pulumi.StringPtrInput
	// Image type. Valid values: `gif`, `jpg`, `tiff`, `png`.
	ImageType pulumi.StringPtrInput
	// Image name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ReplacemsgimageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replacemsgimageArgs)(nil)).Elem()
}

type ReplacemsgimageInput interface {
	pulumi.Input

	ToReplacemsgimageOutput() ReplacemsgimageOutput
	ToReplacemsgimageOutputWithContext(ctx context.Context) ReplacemsgimageOutput
}

func (*Replacemsgimage) ElementType() reflect.Type {
	return reflect.TypeOf((**Replacemsgimage)(nil)).Elem()
}

func (i *Replacemsgimage) ToReplacemsgimageOutput() ReplacemsgimageOutput {
	return i.ToReplacemsgimageOutputWithContext(context.Background())
}

func (i *Replacemsgimage) ToReplacemsgimageOutputWithContext(ctx context.Context) ReplacemsgimageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplacemsgimageOutput)
}

// ReplacemsgimageArrayInput is an input type that accepts ReplacemsgimageArray and ReplacemsgimageArrayOutput values.
// You can construct a concrete instance of `ReplacemsgimageArrayInput` via:
//
//	ReplacemsgimageArray{ ReplacemsgimageArgs{...} }
type ReplacemsgimageArrayInput interface {
	pulumi.Input

	ToReplacemsgimageArrayOutput() ReplacemsgimageArrayOutput
	ToReplacemsgimageArrayOutputWithContext(context.Context) ReplacemsgimageArrayOutput
}

type ReplacemsgimageArray []ReplacemsgimageInput

func (ReplacemsgimageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Replacemsgimage)(nil)).Elem()
}

func (i ReplacemsgimageArray) ToReplacemsgimageArrayOutput() ReplacemsgimageArrayOutput {
	return i.ToReplacemsgimageArrayOutputWithContext(context.Background())
}

func (i ReplacemsgimageArray) ToReplacemsgimageArrayOutputWithContext(ctx context.Context) ReplacemsgimageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplacemsgimageArrayOutput)
}

// ReplacemsgimageMapInput is an input type that accepts ReplacemsgimageMap and ReplacemsgimageMapOutput values.
// You can construct a concrete instance of `ReplacemsgimageMapInput` via:
//
//	ReplacemsgimageMap{ "key": ReplacemsgimageArgs{...} }
type ReplacemsgimageMapInput interface {
	pulumi.Input

	ToReplacemsgimageMapOutput() ReplacemsgimageMapOutput
	ToReplacemsgimageMapOutputWithContext(context.Context) ReplacemsgimageMapOutput
}

type ReplacemsgimageMap map[string]ReplacemsgimageInput

func (ReplacemsgimageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Replacemsgimage)(nil)).Elem()
}

func (i ReplacemsgimageMap) ToReplacemsgimageMapOutput() ReplacemsgimageMapOutput {
	return i.ToReplacemsgimageMapOutputWithContext(context.Background())
}

func (i ReplacemsgimageMap) ToReplacemsgimageMapOutputWithContext(ctx context.Context) ReplacemsgimageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplacemsgimageMapOutput)
}

type ReplacemsgimageOutput struct{ *pulumi.OutputState }

func (ReplacemsgimageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Replacemsgimage)(nil)).Elem()
}

func (o ReplacemsgimageOutput) ToReplacemsgimageOutput() ReplacemsgimageOutput {
	return o
}

func (o ReplacemsgimageOutput) ToReplacemsgimageOutputWithContext(ctx context.Context) ReplacemsgimageOutput {
	return o
}

// Image data.
func (o ReplacemsgimageOutput) ImageBase64() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Replacemsgimage) pulumi.StringPtrOutput { return v.ImageBase64 }).(pulumi.StringPtrOutput)
}

// Image type. Valid values: `gif`, `jpg`, `tiff`, `png`.
func (o ReplacemsgimageOutput) ImageType() pulumi.StringOutput {
	return o.ApplyT(func(v *Replacemsgimage) pulumi.StringOutput { return v.ImageType }).(pulumi.StringOutput)
}

// Image name.
func (o ReplacemsgimageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Replacemsgimage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ReplacemsgimageOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Replacemsgimage) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ReplacemsgimageArrayOutput struct{ *pulumi.OutputState }

func (ReplacemsgimageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Replacemsgimage)(nil)).Elem()
}

func (o ReplacemsgimageArrayOutput) ToReplacemsgimageArrayOutput() ReplacemsgimageArrayOutput {
	return o
}

func (o ReplacemsgimageArrayOutput) ToReplacemsgimageArrayOutputWithContext(ctx context.Context) ReplacemsgimageArrayOutput {
	return o
}

func (o ReplacemsgimageArrayOutput) Index(i pulumi.IntInput) ReplacemsgimageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Replacemsgimage {
		return vs[0].([]*Replacemsgimage)[vs[1].(int)]
	}).(ReplacemsgimageOutput)
}

type ReplacemsgimageMapOutput struct{ *pulumi.OutputState }

func (ReplacemsgimageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Replacemsgimage)(nil)).Elem()
}

func (o ReplacemsgimageMapOutput) ToReplacemsgimageMapOutput() ReplacemsgimageMapOutput {
	return o
}

func (o ReplacemsgimageMapOutput) ToReplacemsgimageMapOutputWithContext(ctx context.Context) ReplacemsgimageMapOutput {
	return o
}

func (o ReplacemsgimageMapOutput) MapIndex(k pulumi.StringInput) ReplacemsgimageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Replacemsgimage {
		return vs[0].(map[string]*Replacemsgimage)[vs[1].(string)]
	}).(ReplacemsgimageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplacemsgimageInput)(nil)).Elem(), &Replacemsgimage{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplacemsgimageArrayInput)(nil)).Elem(), ReplacemsgimageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplacemsgimageMapInput)(nil)).Elem(), ReplacemsgimageMap{})
	pulumi.RegisterOutputType(ReplacemsgimageOutput{})
	pulumi.RegisterOutputType(ReplacemsgimageArrayOutput{})
	pulumi.RegisterOutputType(ReplacemsgimageMapOutput{})
}
