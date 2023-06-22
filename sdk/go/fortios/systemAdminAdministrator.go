// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to configure administrator accounts of FortiOS.
//
// !> **Warning:** The resource will be deprecated and replaced by new resource `SystemAdmin`, we recommend that you use the new resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fortios.NewSystemAdminAdministrator(ctx, "admintest", &fortios.SystemAdminAdministratorArgs{
//				Accprofile: pulumi.String("3d3"),
//				Comments:   pulumi.String("comments"),
//				Password:   pulumi.String("cc37331AC1"),
//				Trusthost1: pulumi.String("1.1.1.0 255.255.255.0"),
//				Trusthost2: pulumi.String("2.2.2.0 255.255.255.0"),
//				Vdoms: pulumi.StringArray{
//					pulumi.String("root"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SystemAdminAdministrator struct {
	pulumi.CustomResourceState

	// Access profile for this administrator. Access profiles control administrator access to FortiGate features.
	Accprofile pulumi.StringOutput `pulumi:"accprofile"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// User name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Admin user password.
	Password    pulumi.StringOutput `pulumi:"password"`
	Trusthost1  pulumi.StringOutput `pulumi:"trusthost1"`
	Trusthost10 pulumi.StringOutput `pulumi:"trusthost10"`
	Trusthost2  pulumi.StringOutput `pulumi:"trusthost2"`
	Trusthost3  pulumi.StringOutput `pulumi:"trusthost3"`
	Trusthost4  pulumi.StringOutput `pulumi:"trusthost4"`
	Trusthost5  pulumi.StringOutput `pulumi:"trusthost5"`
	Trusthost6  pulumi.StringOutput `pulumi:"trusthost6"`
	Trusthost7  pulumi.StringOutput `pulumi:"trusthost7"`
	Trusthost8  pulumi.StringOutput `pulumi:"trusthost8"`
	Trusthost9  pulumi.StringOutput `pulumi:"trusthost9"`
	// Virtual domain(s) that the administrator can access.
	Vdoms pulumi.StringArrayOutput `pulumi:"vdoms"`
}

// NewSystemAdminAdministrator registers a new resource with the given unique name, arguments, and options.
func NewSystemAdminAdministrator(ctx *pulumi.Context,
	name string, args *SystemAdminAdministratorArgs, opts ...pulumi.ResourceOption) (*SystemAdminAdministrator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Accprofile == nil {
		return nil, errors.New("invalid value for required argument 'Accprofile'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemAdminAdministrator
	err := ctx.RegisterResource("fortios:index/systemAdminAdministrator:SystemAdminAdministrator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemAdminAdministrator gets an existing SystemAdminAdministrator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemAdminAdministrator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemAdminAdministratorState, opts ...pulumi.ResourceOption) (*SystemAdminAdministrator, error) {
	var resource SystemAdminAdministrator
	err := ctx.ReadResource("fortios:index/systemAdminAdministrator:SystemAdminAdministrator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemAdminAdministrator resources.
type systemAdminAdministratorState struct {
	// Access profile for this administrator. Access profiles control administrator access to FortiGate features.
	Accprofile *string `pulumi:"accprofile"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// User name.
	Name *string `pulumi:"name"`
	// Admin user password.
	Password    *string `pulumi:"password"`
	Trusthost1  *string `pulumi:"trusthost1"`
	Trusthost10 *string `pulumi:"trusthost10"`
	Trusthost2  *string `pulumi:"trusthost2"`
	Trusthost3  *string `pulumi:"trusthost3"`
	Trusthost4  *string `pulumi:"trusthost4"`
	Trusthost5  *string `pulumi:"trusthost5"`
	Trusthost6  *string `pulumi:"trusthost6"`
	Trusthost7  *string `pulumi:"trusthost7"`
	Trusthost8  *string `pulumi:"trusthost8"`
	Trusthost9  *string `pulumi:"trusthost9"`
	// Virtual domain(s) that the administrator can access.
	Vdoms []string `pulumi:"vdoms"`
}

type SystemAdminAdministratorState struct {
	// Access profile for this administrator. Access profiles control administrator access to FortiGate features.
	Accprofile pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// User name.
	Name pulumi.StringPtrInput
	// Admin user password.
	Password    pulumi.StringPtrInput
	Trusthost1  pulumi.StringPtrInput
	Trusthost10 pulumi.StringPtrInput
	Trusthost2  pulumi.StringPtrInput
	Trusthost3  pulumi.StringPtrInput
	Trusthost4  pulumi.StringPtrInput
	Trusthost5  pulumi.StringPtrInput
	Trusthost6  pulumi.StringPtrInput
	Trusthost7  pulumi.StringPtrInput
	Trusthost8  pulumi.StringPtrInput
	Trusthost9  pulumi.StringPtrInput
	// Virtual domain(s) that the administrator can access.
	Vdoms pulumi.StringArrayInput
}

func (SystemAdminAdministratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAdminAdministratorState)(nil)).Elem()
}

type systemAdminAdministratorArgs struct {
	// Access profile for this administrator. Access profiles control administrator access to FortiGate features.
	Accprofile string `pulumi:"accprofile"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// User name.
	Name *string `pulumi:"name"`
	// Admin user password.
	Password    string  `pulumi:"password"`
	Trusthost1  *string `pulumi:"trusthost1"`
	Trusthost10 *string `pulumi:"trusthost10"`
	Trusthost2  *string `pulumi:"trusthost2"`
	Trusthost3  *string `pulumi:"trusthost3"`
	Trusthost4  *string `pulumi:"trusthost4"`
	Trusthost5  *string `pulumi:"trusthost5"`
	Trusthost6  *string `pulumi:"trusthost6"`
	Trusthost7  *string `pulumi:"trusthost7"`
	Trusthost8  *string `pulumi:"trusthost8"`
	Trusthost9  *string `pulumi:"trusthost9"`
	// Virtual domain(s) that the administrator can access.
	Vdoms []string `pulumi:"vdoms"`
}

// The set of arguments for constructing a SystemAdminAdministrator resource.
type SystemAdminAdministratorArgs struct {
	// Access profile for this administrator. Access profiles control administrator access to FortiGate features.
	Accprofile pulumi.StringInput
	// Comment.
	Comments pulumi.StringPtrInput
	// User name.
	Name pulumi.StringPtrInput
	// Admin user password.
	Password    pulumi.StringInput
	Trusthost1  pulumi.StringPtrInput
	Trusthost10 pulumi.StringPtrInput
	Trusthost2  pulumi.StringPtrInput
	Trusthost3  pulumi.StringPtrInput
	Trusthost4  pulumi.StringPtrInput
	Trusthost5  pulumi.StringPtrInput
	Trusthost6  pulumi.StringPtrInput
	Trusthost7  pulumi.StringPtrInput
	Trusthost8  pulumi.StringPtrInput
	Trusthost9  pulumi.StringPtrInput
	// Virtual domain(s) that the administrator can access.
	Vdoms pulumi.StringArrayInput
}

func (SystemAdminAdministratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAdminAdministratorArgs)(nil)).Elem()
}

type SystemAdminAdministratorInput interface {
	pulumi.Input

	ToSystemAdminAdministratorOutput() SystemAdminAdministratorOutput
	ToSystemAdminAdministratorOutputWithContext(ctx context.Context) SystemAdminAdministratorOutput
}

func (*SystemAdminAdministrator) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAdminAdministrator)(nil)).Elem()
}

func (i *SystemAdminAdministrator) ToSystemAdminAdministratorOutput() SystemAdminAdministratorOutput {
	return i.ToSystemAdminAdministratorOutputWithContext(context.Background())
}

func (i *SystemAdminAdministrator) ToSystemAdminAdministratorOutputWithContext(ctx context.Context) SystemAdminAdministratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAdminAdministratorOutput)
}

// SystemAdminAdministratorArrayInput is an input type that accepts SystemAdminAdministratorArray and SystemAdminAdministratorArrayOutput values.
// You can construct a concrete instance of `SystemAdminAdministratorArrayInput` via:
//
//	SystemAdminAdministratorArray{ SystemAdminAdministratorArgs{...} }
type SystemAdminAdministratorArrayInput interface {
	pulumi.Input

	ToSystemAdminAdministratorArrayOutput() SystemAdminAdministratorArrayOutput
	ToSystemAdminAdministratorArrayOutputWithContext(context.Context) SystemAdminAdministratorArrayOutput
}

type SystemAdminAdministratorArray []SystemAdminAdministratorInput

func (SystemAdminAdministratorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAdminAdministrator)(nil)).Elem()
}

func (i SystemAdminAdministratorArray) ToSystemAdminAdministratorArrayOutput() SystemAdminAdministratorArrayOutput {
	return i.ToSystemAdminAdministratorArrayOutputWithContext(context.Background())
}

func (i SystemAdminAdministratorArray) ToSystemAdminAdministratorArrayOutputWithContext(ctx context.Context) SystemAdminAdministratorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAdminAdministratorArrayOutput)
}

// SystemAdminAdministratorMapInput is an input type that accepts SystemAdminAdministratorMap and SystemAdminAdministratorMapOutput values.
// You can construct a concrete instance of `SystemAdminAdministratorMapInput` via:
//
//	SystemAdminAdministratorMap{ "key": SystemAdminAdministratorArgs{...} }
type SystemAdminAdministratorMapInput interface {
	pulumi.Input

	ToSystemAdminAdministratorMapOutput() SystemAdminAdministratorMapOutput
	ToSystemAdminAdministratorMapOutputWithContext(context.Context) SystemAdminAdministratorMapOutput
}

type SystemAdminAdministratorMap map[string]SystemAdminAdministratorInput

func (SystemAdminAdministratorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAdminAdministrator)(nil)).Elem()
}

func (i SystemAdminAdministratorMap) ToSystemAdminAdministratorMapOutput() SystemAdminAdministratorMapOutput {
	return i.ToSystemAdminAdministratorMapOutputWithContext(context.Background())
}

func (i SystemAdminAdministratorMap) ToSystemAdminAdministratorMapOutputWithContext(ctx context.Context) SystemAdminAdministratorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAdminAdministratorMapOutput)
}

type SystemAdminAdministratorOutput struct{ *pulumi.OutputState }

func (SystemAdminAdministratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAdminAdministrator)(nil)).Elem()
}

func (o SystemAdminAdministratorOutput) ToSystemAdminAdministratorOutput() SystemAdminAdministratorOutput {
	return o
}

func (o SystemAdminAdministratorOutput) ToSystemAdminAdministratorOutputWithContext(ctx context.Context) SystemAdminAdministratorOutput {
	return o
}

// Access profile for this administrator. Access profiles control administrator access to FortiGate features.
func (o SystemAdminAdministratorOutput) Accprofile() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAdminAdministrator) pulumi.StringOutput { return v.Accprofile }).(pulumi.StringOutput)
}

// Comment.
func (o SystemAdminAdministratorOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAdminAdministrator) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// User name.
func (o SystemAdminAdministratorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAdminAdministrator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Admin user password.
func (o SystemAdminAdministratorOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAdminAdministrator) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

func (o SystemAdminAdministratorOutput) Trusthost1() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAdminAdministrator) pulumi.StringOutput { return v.Trusthost1 }).(pulumi.StringOutput)
}

func (o SystemAdminAdministratorOutput) Trusthost10() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAdminAdministrator) pulumi.StringOutput { return v.Trusthost10 }).(pulumi.StringOutput)
}

func (o SystemAdminAdministratorOutput) Trusthost2() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAdminAdministrator) pulumi.StringOutput { return v.Trusthost2 }).(pulumi.StringOutput)
}

func (o SystemAdminAdministratorOutput) Trusthost3() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAdminAdministrator) pulumi.StringOutput { return v.Trusthost3 }).(pulumi.StringOutput)
}

func (o SystemAdminAdministratorOutput) Trusthost4() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAdminAdministrator) pulumi.StringOutput { return v.Trusthost4 }).(pulumi.StringOutput)
}

func (o SystemAdminAdministratorOutput) Trusthost5() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAdminAdministrator) pulumi.StringOutput { return v.Trusthost5 }).(pulumi.StringOutput)
}

func (o SystemAdminAdministratorOutput) Trusthost6() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAdminAdministrator) pulumi.StringOutput { return v.Trusthost6 }).(pulumi.StringOutput)
}

func (o SystemAdminAdministratorOutput) Trusthost7() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAdminAdministrator) pulumi.StringOutput { return v.Trusthost7 }).(pulumi.StringOutput)
}

func (o SystemAdminAdministratorOutput) Trusthost8() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAdminAdministrator) pulumi.StringOutput { return v.Trusthost8 }).(pulumi.StringOutput)
}

func (o SystemAdminAdministratorOutput) Trusthost9() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAdminAdministrator) pulumi.StringOutput { return v.Trusthost9 }).(pulumi.StringOutput)
}

// Virtual domain(s) that the administrator can access.
func (o SystemAdminAdministratorOutput) Vdoms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SystemAdminAdministrator) pulumi.StringArrayOutput { return v.Vdoms }).(pulumi.StringArrayOutput)
}

type SystemAdminAdministratorArrayOutput struct{ *pulumi.OutputState }

func (SystemAdminAdministratorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAdminAdministrator)(nil)).Elem()
}

func (o SystemAdminAdministratorArrayOutput) ToSystemAdminAdministratorArrayOutput() SystemAdminAdministratorArrayOutput {
	return o
}

func (o SystemAdminAdministratorArrayOutput) ToSystemAdminAdministratorArrayOutputWithContext(ctx context.Context) SystemAdminAdministratorArrayOutput {
	return o
}

func (o SystemAdminAdministratorArrayOutput) Index(i pulumi.IntInput) SystemAdminAdministratorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemAdminAdministrator {
		return vs[0].([]*SystemAdminAdministrator)[vs[1].(int)]
	}).(SystemAdminAdministratorOutput)
}

type SystemAdminAdministratorMapOutput struct{ *pulumi.OutputState }

func (SystemAdminAdministratorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAdminAdministrator)(nil)).Elem()
}

func (o SystemAdminAdministratorMapOutput) ToSystemAdminAdministratorMapOutput() SystemAdminAdministratorMapOutput {
	return o
}

func (o SystemAdminAdministratorMapOutput) ToSystemAdminAdministratorMapOutputWithContext(ctx context.Context) SystemAdminAdministratorMapOutput {
	return o
}

func (o SystemAdminAdministratorMapOutput) MapIndex(k pulumi.StringInput) SystemAdminAdministratorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemAdminAdministrator {
		return vs[0].(map[string]*SystemAdminAdministrator)[vs[1].(string)]
	}).(SystemAdminAdministratorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAdminAdministratorInput)(nil)).Elem(), &SystemAdminAdministrator{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAdminAdministratorArrayInput)(nil)).Elem(), SystemAdminAdministratorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAdminAdministratorMapInput)(nil)).Elem(), SystemAdminAdministratorMap{})
	pulumi.RegisterOutputType(SystemAdminAdministratorOutput{})
	pulumi.RegisterOutputType(SystemAdminAdministratorArrayOutput{})
	pulumi.RegisterOutputType(SystemAdminAdministratorMapOutput{})
}
