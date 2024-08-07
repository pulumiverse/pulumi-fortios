// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fmg

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// This resource supports Create/Read/Update/Delete system adom for FortiManager.
//
// ## Example Usage
//
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
//			_, err := fmg.NewSystemAdom(ctx, "test1", &fmg.SystemAdomArgs{
//				ActionWhenConflictsOccurDuringPolicyCheck:  pulumi.String("Continue"),
//				AutoPushPolicyPackagesWhenDeviceBackOnline: pulumi.String("Enable"),
//				CentralManagementFortiap:                   pulumi.Bool(true),
//				CentralManagementSdwan:                     pulumi.Bool(false),
//				CentralManagementVpn:                       pulumi.Bool(false),
//				Mode:                                       pulumi.String("Normal"),
//				PerformPolicyCheckBeforeEveryInstall:       pulumi.Bool(true),
//				Status:                                     pulumi.Int(1),
//				Type:                                       pulumi.String("FortiCarrier"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SystemAdom struct {
	pulumi.CustomResourceState

	// True or False.
	ActionWhenConflictsOccurDuringPolicyCheck pulumi.StringPtrOutput `pulumi:"actionWhenConflictsOccurDuringPolicyCheck"`
	// True or False.
	AutoPushPolicyPackagesWhenDeviceBackOnline pulumi.StringPtrOutput `pulumi:"autoPushPolicyPackagesWhenDeviceBackOnline"`
	// True or False.
	CentralManagementFortiap pulumi.BoolPtrOutput `pulumi:"centralManagementFortiap"`
	// True or False.
	CentralManagementSdwan pulumi.BoolPtrOutput `pulumi:"centralManagementSdwan"`
	// True or False.
	CentralManagementVpn pulumi.BoolPtrOutput `pulumi:"centralManagementVpn"`
	// Adom mode: Normal or Backup.
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// Administrative Domain name.
	Name pulumi.StringOutput `pulumi:"name"`
	// True or False.
	PerformPolicyCheckBeforeEveryInstall pulumi.BoolPtrOutput `pulumi:"performPolicyCheckBeforeEveryInstall"`
	// Adom status. 0 means off and 1 means on.
	Status pulumi.IntPtrOutput `pulumi:"status"`
	// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewSystemAdom registers a new resource with the given unique name, arguments, and options.
func NewSystemAdom(ctx *pulumi.Context,
	name string, args *SystemAdomArgs, opts ...pulumi.ResourceOption) (*SystemAdom, error) {
	if args == nil {
		args = &SystemAdomArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemAdom
	err := ctx.RegisterResource("fortios:fmg/systemAdom:SystemAdom", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemAdom gets an existing SystemAdom resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemAdom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemAdomState, opts ...pulumi.ResourceOption) (*SystemAdom, error) {
	var resource SystemAdom
	err := ctx.ReadResource("fortios:fmg/systemAdom:SystemAdom", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemAdom resources.
type systemAdomState struct {
	// True or False.
	ActionWhenConflictsOccurDuringPolicyCheck *string `pulumi:"actionWhenConflictsOccurDuringPolicyCheck"`
	// True or False.
	AutoPushPolicyPackagesWhenDeviceBackOnline *string `pulumi:"autoPushPolicyPackagesWhenDeviceBackOnline"`
	// True or False.
	CentralManagementFortiap *bool `pulumi:"centralManagementFortiap"`
	// True or False.
	CentralManagementSdwan *bool `pulumi:"centralManagementSdwan"`
	// True or False.
	CentralManagementVpn *bool `pulumi:"centralManagementVpn"`
	// Adom mode: Normal or Backup.
	Mode *string `pulumi:"mode"`
	// Administrative Domain name.
	Name *string `pulumi:"name"`
	// True or False.
	PerformPolicyCheckBeforeEveryInstall *bool `pulumi:"performPolicyCheckBeforeEveryInstall"`
	// Adom status. 0 means off and 1 means on.
	Status *int `pulumi:"status"`
	// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
	Type *string `pulumi:"type"`
}

type SystemAdomState struct {
	// True or False.
	ActionWhenConflictsOccurDuringPolicyCheck pulumi.StringPtrInput
	// True or False.
	AutoPushPolicyPackagesWhenDeviceBackOnline pulumi.StringPtrInput
	// True or False.
	CentralManagementFortiap pulumi.BoolPtrInput
	// True or False.
	CentralManagementSdwan pulumi.BoolPtrInput
	// True or False.
	CentralManagementVpn pulumi.BoolPtrInput
	// Adom mode: Normal or Backup.
	Mode pulumi.StringPtrInput
	// Administrative Domain name.
	Name pulumi.StringPtrInput
	// True or False.
	PerformPolicyCheckBeforeEveryInstall pulumi.BoolPtrInput
	// Adom status. 0 means off and 1 means on.
	Status pulumi.IntPtrInput
	// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
	Type pulumi.StringPtrInput
}

func (SystemAdomState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAdomState)(nil)).Elem()
}

type systemAdomArgs struct {
	// True or False.
	ActionWhenConflictsOccurDuringPolicyCheck *string `pulumi:"actionWhenConflictsOccurDuringPolicyCheck"`
	// True or False.
	AutoPushPolicyPackagesWhenDeviceBackOnline *string `pulumi:"autoPushPolicyPackagesWhenDeviceBackOnline"`
	// True or False.
	CentralManagementFortiap *bool `pulumi:"centralManagementFortiap"`
	// True or False.
	CentralManagementSdwan *bool `pulumi:"centralManagementSdwan"`
	// True or False.
	CentralManagementVpn *bool `pulumi:"centralManagementVpn"`
	// Adom mode: Normal or Backup.
	Mode *string `pulumi:"mode"`
	// Administrative Domain name.
	Name *string `pulumi:"name"`
	// True or False.
	PerformPolicyCheckBeforeEveryInstall *bool `pulumi:"performPolicyCheckBeforeEveryInstall"`
	// Adom status. 0 means off and 1 means on.
	Status *int `pulumi:"status"`
	// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a SystemAdom resource.
type SystemAdomArgs struct {
	// True or False.
	ActionWhenConflictsOccurDuringPolicyCheck pulumi.StringPtrInput
	// True or False.
	AutoPushPolicyPackagesWhenDeviceBackOnline pulumi.StringPtrInput
	// True or False.
	CentralManagementFortiap pulumi.BoolPtrInput
	// True or False.
	CentralManagementSdwan pulumi.BoolPtrInput
	// True or False.
	CentralManagementVpn pulumi.BoolPtrInput
	// Adom mode: Normal or Backup.
	Mode pulumi.StringPtrInput
	// Administrative Domain name.
	Name pulumi.StringPtrInput
	// True or False.
	PerformPolicyCheckBeforeEveryInstall pulumi.BoolPtrInput
	// Adom status. 0 means off and 1 means on.
	Status pulumi.IntPtrInput
	// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
	Type pulumi.StringPtrInput
}

func (SystemAdomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAdomArgs)(nil)).Elem()
}

type SystemAdomInput interface {
	pulumi.Input

	ToSystemAdomOutput() SystemAdomOutput
	ToSystemAdomOutputWithContext(ctx context.Context) SystemAdomOutput
}

func (*SystemAdom) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAdom)(nil)).Elem()
}

func (i *SystemAdom) ToSystemAdomOutput() SystemAdomOutput {
	return i.ToSystemAdomOutputWithContext(context.Background())
}

func (i *SystemAdom) ToSystemAdomOutputWithContext(ctx context.Context) SystemAdomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAdomOutput)
}

// SystemAdomArrayInput is an input type that accepts SystemAdomArray and SystemAdomArrayOutput values.
// You can construct a concrete instance of `SystemAdomArrayInput` via:
//
//	SystemAdomArray{ SystemAdomArgs{...} }
type SystemAdomArrayInput interface {
	pulumi.Input

	ToSystemAdomArrayOutput() SystemAdomArrayOutput
	ToSystemAdomArrayOutputWithContext(context.Context) SystemAdomArrayOutput
}

type SystemAdomArray []SystemAdomInput

func (SystemAdomArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAdom)(nil)).Elem()
}

func (i SystemAdomArray) ToSystemAdomArrayOutput() SystemAdomArrayOutput {
	return i.ToSystemAdomArrayOutputWithContext(context.Background())
}

func (i SystemAdomArray) ToSystemAdomArrayOutputWithContext(ctx context.Context) SystemAdomArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAdomArrayOutput)
}

// SystemAdomMapInput is an input type that accepts SystemAdomMap and SystemAdomMapOutput values.
// You can construct a concrete instance of `SystemAdomMapInput` via:
//
//	SystemAdomMap{ "key": SystemAdomArgs{...} }
type SystemAdomMapInput interface {
	pulumi.Input

	ToSystemAdomMapOutput() SystemAdomMapOutput
	ToSystemAdomMapOutputWithContext(context.Context) SystemAdomMapOutput
}

type SystemAdomMap map[string]SystemAdomInput

func (SystemAdomMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAdom)(nil)).Elem()
}

func (i SystemAdomMap) ToSystemAdomMapOutput() SystemAdomMapOutput {
	return i.ToSystemAdomMapOutputWithContext(context.Background())
}

func (i SystemAdomMap) ToSystemAdomMapOutputWithContext(ctx context.Context) SystemAdomMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAdomMapOutput)
}

type SystemAdomOutput struct{ *pulumi.OutputState }

func (SystemAdomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAdom)(nil)).Elem()
}

func (o SystemAdomOutput) ToSystemAdomOutput() SystemAdomOutput {
	return o
}

func (o SystemAdomOutput) ToSystemAdomOutputWithContext(ctx context.Context) SystemAdomOutput {
	return o
}

// True or False.
func (o SystemAdomOutput) ActionWhenConflictsOccurDuringPolicyCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAdom) pulumi.StringPtrOutput { return v.ActionWhenConflictsOccurDuringPolicyCheck }).(pulumi.StringPtrOutput)
}

// True or False.
func (o SystemAdomOutput) AutoPushPolicyPackagesWhenDeviceBackOnline() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAdom) pulumi.StringPtrOutput { return v.AutoPushPolicyPackagesWhenDeviceBackOnline }).(pulumi.StringPtrOutput)
}

// True or False.
func (o SystemAdomOutput) CentralManagementFortiap() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SystemAdom) pulumi.BoolPtrOutput { return v.CentralManagementFortiap }).(pulumi.BoolPtrOutput)
}

// True or False.
func (o SystemAdomOutput) CentralManagementSdwan() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SystemAdom) pulumi.BoolPtrOutput { return v.CentralManagementSdwan }).(pulumi.BoolPtrOutput)
}

// True or False.
func (o SystemAdomOutput) CentralManagementVpn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SystemAdom) pulumi.BoolPtrOutput { return v.CentralManagementVpn }).(pulumi.BoolPtrOutput)
}

// Adom mode: Normal or Backup.
func (o SystemAdomOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAdom) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// Administrative Domain name.
func (o SystemAdomOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAdom) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// True or False.
func (o SystemAdomOutput) PerformPolicyCheckBeforeEveryInstall() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SystemAdom) pulumi.BoolPtrOutput { return v.PerformPolicyCheckBeforeEveryInstall }).(pulumi.BoolPtrOutput)
}

// Adom status. 0 means off and 1 means on.
func (o SystemAdomOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SystemAdom) pulumi.IntPtrOutput { return v.Status }).(pulumi.IntPtrOutput)
}

// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
func (o SystemAdomOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAdom) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type SystemAdomArrayOutput struct{ *pulumi.OutputState }

func (SystemAdomArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAdom)(nil)).Elem()
}

func (o SystemAdomArrayOutput) ToSystemAdomArrayOutput() SystemAdomArrayOutput {
	return o
}

func (o SystemAdomArrayOutput) ToSystemAdomArrayOutputWithContext(ctx context.Context) SystemAdomArrayOutput {
	return o
}

func (o SystemAdomArrayOutput) Index(i pulumi.IntInput) SystemAdomOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemAdom {
		return vs[0].([]*SystemAdom)[vs[1].(int)]
	}).(SystemAdomOutput)
}

type SystemAdomMapOutput struct{ *pulumi.OutputState }

func (SystemAdomMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAdom)(nil)).Elem()
}

func (o SystemAdomMapOutput) ToSystemAdomMapOutput() SystemAdomMapOutput {
	return o
}

func (o SystemAdomMapOutput) ToSystemAdomMapOutputWithContext(ctx context.Context) SystemAdomMapOutput {
	return o
}

func (o SystemAdomMapOutput) MapIndex(k pulumi.StringInput) SystemAdomOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemAdom {
		return vs[0].(map[string]*SystemAdom)[vs[1].(string)]
	}).(SystemAdomOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAdomInput)(nil)).Elem(), &SystemAdom{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAdomArrayInput)(nil)).Elem(), SystemAdomArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAdomMapInput)(nil)).Elem(), SystemAdomMap{})
	pulumi.RegisterOutputType(SystemAdomOutput{})
	pulumi.RegisterOutputType(SystemAdomArrayOutput{})
	pulumi.RegisterOutputType(SystemAdomMapOutput{})
}
