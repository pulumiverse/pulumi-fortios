// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fmg

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// This resource supports uploading VM license to FortiGate through FortiManager.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/fmg"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fmg.NewSystemLicenseVm(ctx, "test1", &fmg.SystemLicenseVmArgs{
//				FileContent: pulumi.String("XXX"),
//				Target:      pulumi.String("fortigate-test"),
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
type SystemLicenseVm struct {
	pulumi.CustomResourceState

	// ADOM that the target device belongs to. default is 'root'.
	Adom pulumi.StringPtrOutput `pulumi:"adom"`
	// The license file, it needs to be base64 encoded.
	FileContent pulumi.StringOutput `pulumi:"fileContent"`
	// Target name, which is managed by FortiManager.
	Target pulumi.StringOutput `pulumi:"target"`
}

// NewSystemLicenseVm registers a new resource with the given unique name, arguments, and options.
func NewSystemLicenseVm(ctx *pulumi.Context,
	name string, args *SystemLicenseVmArgs, opts ...pulumi.ResourceOption) (*SystemLicenseVm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FileContent == nil {
		return nil, errors.New("invalid value for required argument 'FileContent'")
	}
	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemLicenseVm
	err := ctx.RegisterResource("fortios:fmg/systemLicenseVm:SystemLicenseVm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemLicenseVm gets an existing SystemLicenseVm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemLicenseVm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemLicenseVmState, opts ...pulumi.ResourceOption) (*SystemLicenseVm, error) {
	var resource SystemLicenseVm
	err := ctx.ReadResource("fortios:fmg/systemLicenseVm:SystemLicenseVm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemLicenseVm resources.
type systemLicenseVmState struct {
	// ADOM that the target device belongs to. default is 'root'.
	Adom *string `pulumi:"adom"`
	// The license file, it needs to be base64 encoded.
	FileContent *string `pulumi:"fileContent"`
	// Target name, which is managed by FortiManager.
	Target *string `pulumi:"target"`
}

type SystemLicenseVmState struct {
	// ADOM that the target device belongs to. default is 'root'.
	Adom pulumi.StringPtrInput
	// The license file, it needs to be base64 encoded.
	FileContent pulumi.StringPtrInput
	// Target name, which is managed by FortiManager.
	Target pulumi.StringPtrInput
}

func (SystemLicenseVmState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemLicenseVmState)(nil)).Elem()
}

type systemLicenseVmArgs struct {
	// ADOM that the target device belongs to. default is 'root'.
	Adom *string `pulumi:"adom"`
	// The license file, it needs to be base64 encoded.
	FileContent string `pulumi:"fileContent"`
	// Target name, which is managed by FortiManager.
	Target string `pulumi:"target"`
}

// The set of arguments for constructing a SystemLicenseVm resource.
type SystemLicenseVmArgs struct {
	// ADOM that the target device belongs to. default is 'root'.
	Adom pulumi.StringPtrInput
	// The license file, it needs to be base64 encoded.
	FileContent pulumi.StringInput
	// Target name, which is managed by FortiManager.
	Target pulumi.StringInput
}

func (SystemLicenseVmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemLicenseVmArgs)(nil)).Elem()
}

type SystemLicenseVmInput interface {
	pulumi.Input

	ToSystemLicenseVmOutput() SystemLicenseVmOutput
	ToSystemLicenseVmOutputWithContext(ctx context.Context) SystemLicenseVmOutput
}

func (*SystemLicenseVm) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemLicenseVm)(nil)).Elem()
}

func (i *SystemLicenseVm) ToSystemLicenseVmOutput() SystemLicenseVmOutput {
	return i.ToSystemLicenseVmOutputWithContext(context.Background())
}

func (i *SystemLicenseVm) ToSystemLicenseVmOutputWithContext(ctx context.Context) SystemLicenseVmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemLicenseVmOutput)
}

// SystemLicenseVmArrayInput is an input type that accepts SystemLicenseVmArray and SystemLicenseVmArrayOutput values.
// You can construct a concrete instance of `SystemLicenseVmArrayInput` via:
//
//	SystemLicenseVmArray{ SystemLicenseVmArgs{...} }
type SystemLicenseVmArrayInput interface {
	pulumi.Input

	ToSystemLicenseVmArrayOutput() SystemLicenseVmArrayOutput
	ToSystemLicenseVmArrayOutputWithContext(context.Context) SystemLicenseVmArrayOutput
}

type SystemLicenseVmArray []SystemLicenseVmInput

func (SystemLicenseVmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemLicenseVm)(nil)).Elem()
}

func (i SystemLicenseVmArray) ToSystemLicenseVmArrayOutput() SystemLicenseVmArrayOutput {
	return i.ToSystemLicenseVmArrayOutputWithContext(context.Background())
}

func (i SystemLicenseVmArray) ToSystemLicenseVmArrayOutputWithContext(ctx context.Context) SystemLicenseVmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemLicenseVmArrayOutput)
}

// SystemLicenseVmMapInput is an input type that accepts SystemLicenseVmMap and SystemLicenseVmMapOutput values.
// You can construct a concrete instance of `SystemLicenseVmMapInput` via:
//
//	SystemLicenseVmMap{ "key": SystemLicenseVmArgs{...} }
type SystemLicenseVmMapInput interface {
	pulumi.Input

	ToSystemLicenseVmMapOutput() SystemLicenseVmMapOutput
	ToSystemLicenseVmMapOutputWithContext(context.Context) SystemLicenseVmMapOutput
}

type SystemLicenseVmMap map[string]SystemLicenseVmInput

func (SystemLicenseVmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemLicenseVm)(nil)).Elem()
}

func (i SystemLicenseVmMap) ToSystemLicenseVmMapOutput() SystemLicenseVmMapOutput {
	return i.ToSystemLicenseVmMapOutputWithContext(context.Background())
}

func (i SystemLicenseVmMap) ToSystemLicenseVmMapOutputWithContext(ctx context.Context) SystemLicenseVmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemLicenseVmMapOutput)
}

type SystemLicenseVmOutput struct{ *pulumi.OutputState }

func (SystemLicenseVmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemLicenseVm)(nil)).Elem()
}

func (o SystemLicenseVmOutput) ToSystemLicenseVmOutput() SystemLicenseVmOutput {
	return o
}

func (o SystemLicenseVmOutput) ToSystemLicenseVmOutputWithContext(ctx context.Context) SystemLicenseVmOutput {
	return o
}

// ADOM that the target device belongs to. default is 'root'.
func (o SystemLicenseVmOutput) Adom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemLicenseVm) pulumi.StringPtrOutput { return v.Adom }).(pulumi.StringPtrOutput)
}

// The license file, it needs to be base64 encoded.
func (o SystemLicenseVmOutput) FileContent() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemLicenseVm) pulumi.StringOutput { return v.FileContent }).(pulumi.StringOutput)
}

// Target name, which is managed by FortiManager.
func (o SystemLicenseVmOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemLicenseVm) pulumi.StringOutput { return v.Target }).(pulumi.StringOutput)
}

type SystemLicenseVmArrayOutput struct{ *pulumi.OutputState }

func (SystemLicenseVmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemLicenseVm)(nil)).Elem()
}

func (o SystemLicenseVmArrayOutput) ToSystemLicenseVmArrayOutput() SystemLicenseVmArrayOutput {
	return o
}

func (o SystemLicenseVmArrayOutput) ToSystemLicenseVmArrayOutputWithContext(ctx context.Context) SystemLicenseVmArrayOutput {
	return o
}

func (o SystemLicenseVmArrayOutput) Index(i pulumi.IntInput) SystemLicenseVmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemLicenseVm {
		return vs[0].([]*SystemLicenseVm)[vs[1].(int)]
	}).(SystemLicenseVmOutput)
}

type SystemLicenseVmMapOutput struct{ *pulumi.OutputState }

func (SystemLicenseVmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemLicenseVm)(nil)).Elem()
}

func (o SystemLicenseVmMapOutput) ToSystemLicenseVmMapOutput() SystemLicenseVmMapOutput {
	return o
}

func (o SystemLicenseVmMapOutput) ToSystemLicenseVmMapOutputWithContext(ctx context.Context) SystemLicenseVmMapOutput {
	return o
}

func (o SystemLicenseVmMapOutput) MapIndex(k pulumi.StringInput) SystemLicenseVmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemLicenseVm {
		return vs[0].(map[string]*SystemLicenseVm)[vs[1].(string)]
	}).(SystemLicenseVmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemLicenseVmInput)(nil)).Elem(), &SystemLicenseVm{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemLicenseVmArrayInput)(nil)).Elem(), SystemLicenseVmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemLicenseVmMapInput)(nil)).Elem(), SystemLicenseVmMap{})
	pulumi.RegisterOutputType(SystemLicenseVmOutput{})
	pulumi.RegisterOutputType(SystemLicenseVmArrayOutput{})
	pulumi.RegisterOutputType(SystemLicenseVmMapOutput{})
}
