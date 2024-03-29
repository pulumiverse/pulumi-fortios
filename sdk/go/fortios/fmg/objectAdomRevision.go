// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fmg

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// This resource supports Create/Read/Update/Delete object adom revision for FortiManager.
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
//			_, err := fmg.NewObjectAdomRevision(ctx, "test1", &fmg.ObjectAdomRevisionArgs{
//				CreatedBy:   pulumi.String("fortinet"),
//				Description: pulumi.String("adom revision"),
//				Locked:      pulumi.Int(0),
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
type ObjectAdomRevision struct {
	pulumi.CustomResourceState

	// ADOM name. default is 'root'.
	Adom pulumi.StringPtrOutput `pulumi:"adom"`
	// Who created this adom revision.
	CreatedBy pulumi.StringPtrOutput `pulumi:"createdBy"`
	// Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// lock. 0 means unlock and 1 means locked.
	Locked pulumi.IntPtrOutput `pulumi:"locked"`
	// Adom revision name.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewObjectAdomRevision registers a new resource with the given unique name, arguments, and options.
func NewObjectAdomRevision(ctx *pulumi.Context,
	name string, args *ObjectAdomRevisionArgs, opts ...pulumi.ResourceOption) (*ObjectAdomRevision, error) {
	if args == nil {
		args = &ObjectAdomRevisionArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ObjectAdomRevision
	err := ctx.RegisterResource("fortios:fmg/objectAdomRevision:ObjectAdomRevision", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectAdomRevision gets an existing ObjectAdomRevision resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectAdomRevision(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectAdomRevisionState, opts ...pulumi.ResourceOption) (*ObjectAdomRevision, error) {
	var resource ObjectAdomRevision
	err := ctx.ReadResource("fortios:fmg/objectAdomRevision:ObjectAdomRevision", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectAdomRevision resources.
type objectAdomRevisionState struct {
	// ADOM name. default is 'root'.
	Adom *string `pulumi:"adom"`
	// Who created this adom revision.
	CreatedBy *string `pulumi:"createdBy"`
	// Description.
	Description *string `pulumi:"description"`
	// lock. 0 means unlock and 1 means locked.
	Locked *int `pulumi:"locked"`
	// Adom revision name.
	Name *string `pulumi:"name"`
}

type ObjectAdomRevisionState struct {
	// ADOM name. default is 'root'.
	Adom pulumi.StringPtrInput
	// Who created this adom revision.
	CreatedBy pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// lock. 0 means unlock and 1 means locked.
	Locked pulumi.IntPtrInput
	// Adom revision name.
	Name pulumi.StringPtrInput
}

func (ObjectAdomRevisionState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectAdomRevisionState)(nil)).Elem()
}

type objectAdomRevisionArgs struct {
	// ADOM name. default is 'root'.
	Adom *string `pulumi:"adom"`
	// Who created this adom revision.
	CreatedBy *string `pulumi:"createdBy"`
	// Description.
	Description *string `pulumi:"description"`
	// lock. 0 means unlock and 1 means locked.
	Locked *int `pulumi:"locked"`
	// Adom revision name.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ObjectAdomRevision resource.
type ObjectAdomRevisionArgs struct {
	// ADOM name. default is 'root'.
	Adom pulumi.StringPtrInput
	// Who created this adom revision.
	CreatedBy pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// lock. 0 means unlock and 1 means locked.
	Locked pulumi.IntPtrInput
	// Adom revision name.
	Name pulumi.StringPtrInput
}

func (ObjectAdomRevisionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectAdomRevisionArgs)(nil)).Elem()
}

type ObjectAdomRevisionInput interface {
	pulumi.Input

	ToObjectAdomRevisionOutput() ObjectAdomRevisionOutput
	ToObjectAdomRevisionOutputWithContext(ctx context.Context) ObjectAdomRevisionOutput
}

func (*ObjectAdomRevision) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectAdomRevision)(nil)).Elem()
}

func (i *ObjectAdomRevision) ToObjectAdomRevisionOutput() ObjectAdomRevisionOutput {
	return i.ToObjectAdomRevisionOutputWithContext(context.Background())
}

func (i *ObjectAdomRevision) ToObjectAdomRevisionOutputWithContext(ctx context.Context) ObjectAdomRevisionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectAdomRevisionOutput)
}

// ObjectAdomRevisionArrayInput is an input type that accepts ObjectAdomRevisionArray and ObjectAdomRevisionArrayOutput values.
// You can construct a concrete instance of `ObjectAdomRevisionArrayInput` via:
//
//	ObjectAdomRevisionArray{ ObjectAdomRevisionArgs{...} }
type ObjectAdomRevisionArrayInput interface {
	pulumi.Input

	ToObjectAdomRevisionArrayOutput() ObjectAdomRevisionArrayOutput
	ToObjectAdomRevisionArrayOutputWithContext(context.Context) ObjectAdomRevisionArrayOutput
}

type ObjectAdomRevisionArray []ObjectAdomRevisionInput

func (ObjectAdomRevisionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectAdomRevision)(nil)).Elem()
}

func (i ObjectAdomRevisionArray) ToObjectAdomRevisionArrayOutput() ObjectAdomRevisionArrayOutput {
	return i.ToObjectAdomRevisionArrayOutputWithContext(context.Background())
}

func (i ObjectAdomRevisionArray) ToObjectAdomRevisionArrayOutputWithContext(ctx context.Context) ObjectAdomRevisionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectAdomRevisionArrayOutput)
}

// ObjectAdomRevisionMapInput is an input type that accepts ObjectAdomRevisionMap and ObjectAdomRevisionMapOutput values.
// You can construct a concrete instance of `ObjectAdomRevisionMapInput` via:
//
//	ObjectAdomRevisionMap{ "key": ObjectAdomRevisionArgs{...} }
type ObjectAdomRevisionMapInput interface {
	pulumi.Input

	ToObjectAdomRevisionMapOutput() ObjectAdomRevisionMapOutput
	ToObjectAdomRevisionMapOutputWithContext(context.Context) ObjectAdomRevisionMapOutput
}

type ObjectAdomRevisionMap map[string]ObjectAdomRevisionInput

func (ObjectAdomRevisionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectAdomRevision)(nil)).Elem()
}

func (i ObjectAdomRevisionMap) ToObjectAdomRevisionMapOutput() ObjectAdomRevisionMapOutput {
	return i.ToObjectAdomRevisionMapOutputWithContext(context.Background())
}

func (i ObjectAdomRevisionMap) ToObjectAdomRevisionMapOutputWithContext(ctx context.Context) ObjectAdomRevisionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectAdomRevisionMapOutput)
}

type ObjectAdomRevisionOutput struct{ *pulumi.OutputState }

func (ObjectAdomRevisionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectAdomRevision)(nil)).Elem()
}

func (o ObjectAdomRevisionOutput) ToObjectAdomRevisionOutput() ObjectAdomRevisionOutput {
	return o
}

func (o ObjectAdomRevisionOutput) ToObjectAdomRevisionOutputWithContext(ctx context.Context) ObjectAdomRevisionOutput {
	return o
}

// ADOM name. default is 'root'.
func (o ObjectAdomRevisionOutput) Adom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectAdomRevision) pulumi.StringPtrOutput { return v.Adom }).(pulumi.StringPtrOutput)
}

// Who created this adom revision.
func (o ObjectAdomRevisionOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectAdomRevision) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// Description.
func (o ObjectAdomRevisionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectAdomRevision) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// lock. 0 means unlock and 1 means locked.
func (o ObjectAdomRevisionOutput) Locked() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ObjectAdomRevision) pulumi.IntPtrOutput { return v.Locked }).(pulumi.IntPtrOutput)
}

// Adom revision name.
func (o ObjectAdomRevisionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAdomRevision) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type ObjectAdomRevisionArrayOutput struct{ *pulumi.OutputState }

func (ObjectAdomRevisionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectAdomRevision)(nil)).Elem()
}

func (o ObjectAdomRevisionArrayOutput) ToObjectAdomRevisionArrayOutput() ObjectAdomRevisionArrayOutput {
	return o
}

func (o ObjectAdomRevisionArrayOutput) ToObjectAdomRevisionArrayOutputWithContext(ctx context.Context) ObjectAdomRevisionArrayOutput {
	return o
}

func (o ObjectAdomRevisionArrayOutput) Index(i pulumi.IntInput) ObjectAdomRevisionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ObjectAdomRevision {
		return vs[0].([]*ObjectAdomRevision)[vs[1].(int)]
	}).(ObjectAdomRevisionOutput)
}

type ObjectAdomRevisionMapOutput struct{ *pulumi.OutputState }

func (ObjectAdomRevisionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectAdomRevision)(nil)).Elem()
}

func (o ObjectAdomRevisionMapOutput) ToObjectAdomRevisionMapOutput() ObjectAdomRevisionMapOutput {
	return o
}

func (o ObjectAdomRevisionMapOutput) ToObjectAdomRevisionMapOutputWithContext(ctx context.Context) ObjectAdomRevisionMapOutput {
	return o
}

func (o ObjectAdomRevisionMapOutput) MapIndex(k pulumi.StringInput) ObjectAdomRevisionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ObjectAdomRevision {
		return vs[0].(map[string]*ObjectAdomRevision)[vs[1].(string)]
	}).(ObjectAdomRevisionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectAdomRevisionInput)(nil)).Elem(), &ObjectAdomRevision{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectAdomRevisionArrayInput)(nil)).Elem(), ObjectAdomRevisionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectAdomRevisionMapInput)(nil)).Elem(), ObjectAdomRevisionMap{})
	pulumi.RegisterOutputType(ObjectAdomRevisionOutput{})
	pulumi.RegisterOutputType(ObjectAdomRevisionArrayOutput{})
	pulumi.RegisterOutputType(ObjectAdomRevisionMapOutput{})
}
