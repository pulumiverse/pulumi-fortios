// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure access lists.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/router"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := router.NewAccesslist(ctx, "trname", &router.AccesslistArgs{
//				Comments: pulumi.String("test accesslist"),
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
// ## Note
//
// The feature can only be correctly supported when FortiOS Version >= 6.2.4, for FortiOS Version < 6.2.4, please use the following resource configuration as an alternative.
//
// ### Example
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
//			_, err := system.NewAutoscript(ctx, "trname1", &system.AutoscriptArgs{
//				Interval:   pulumi.Int(1),
//				OutputSize: pulumi.Int(10),
//				Repeat:     pulumi.Int(1),
//				Script: pulumi.String(`config router access-list
//
// edit "static-redistribution"
// config rule
// edit 10
// set prefix 10.0.0.0 255.255.255.0
// set action permit
// set exact-match enable
// end
// end
//
// `),
//
//				Start: pulumi.String("auto"),
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
// Router AccessList can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:router/accesslist:Accesslist labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
// $ pulumi import fortios:router/accesslist:Accesslist labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Accesslist struct {
	pulumi.CustomResourceState

	// Comment.
	Comments            pulumi.StringOutput    `pulumi:"comments"`
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules     AccesslistRuleArrayOutput `pulumi:"rules"`
	Vdomparam pulumi.StringPtrOutput    `pulumi:"vdomparam"`
}

// NewAccesslist registers a new resource with the given unique name, arguments, and options.
func NewAccesslist(ctx *pulumi.Context,
	name string, args *AccesslistArgs, opts ...pulumi.ResourceOption) (*Accesslist, error) {
	if args == nil {
		args = &AccesslistArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Accesslist
	err := ctx.RegisterResource("fortios:router/accesslist:Accesslist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccesslist gets an existing Accesslist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccesslist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccesslistState, opts ...pulumi.ResourceOption) (*Accesslist, error) {
	var resource Accesslist
	err := ctx.ReadResource("fortios:router/accesslist:Accesslist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Accesslist resources.
type accesslistState struct {
	// Comment.
	Comments            *string `pulumi:"comments"`
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name.
	Name *string `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules     []AccesslistRule `pulumi:"rules"`
	Vdomparam *string          `pulumi:"vdomparam"`
}

type AccesslistState struct {
	// Comment.
	Comments            pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Rule. The structure of `rule` block is documented below.
	Rules     AccesslistRuleArrayInput
	Vdomparam pulumi.StringPtrInput
}

func (AccesslistState) ElementType() reflect.Type {
	return reflect.TypeOf((*accesslistState)(nil)).Elem()
}

type accesslistArgs struct {
	// Comment.
	Comments            *string `pulumi:"comments"`
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name.
	Name *string `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules     []AccesslistRule `pulumi:"rules"`
	Vdomparam *string          `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Accesslist resource.
type AccesslistArgs struct {
	// Comment.
	Comments            pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Rule. The structure of `rule` block is documented below.
	Rules     AccesslistRuleArrayInput
	Vdomparam pulumi.StringPtrInput
}

func (AccesslistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accesslistArgs)(nil)).Elem()
}

type AccesslistInput interface {
	pulumi.Input

	ToAccesslistOutput() AccesslistOutput
	ToAccesslistOutputWithContext(ctx context.Context) AccesslistOutput
}

func (*Accesslist) ElementType() reflect.Type {
	return reflect.TypeOf((**Accesslist)(nil)).Elem()
}

func (i *Accesslist) ToAccesslistOutput() AccesslistOutput {
	return i.ToAccesslistOutputWithContext(context.Background())
}

func (i *Accesslist) ToAccesslistOutputWithContext(ctx context.Context) AccesslistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccesslistOutput)
}

// AccesslistArrayInput is an input type that accepts AccesslistArray and AccesslistArrayOutput values.
// You can construct a concrete instance of `AccesslistArrayInput` via:
//
//	AccesslistArray{ AccesslistArgs{...} }
type AccesslistArrayInput interface {
	pulumi.Input

	ToAccesslistArrayOutput() AccesslistArrayOutput
	ToAccesslistArrayOutputWithContext(context.Context) AccesslistArrayOutput
}

type AccesslistArray []AccesslistInput

func (AccesslistArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Accesslist)(nil)).Elem()
}

func (i AccesslistArray) ToAccesslistArrayOutput() AccesslistArrayOutput {
	return i.ToAccesslistArrayOutputWithContext(context.Background())
}

func (i AccesslistArray) ToAccesslistArrayOutputWithContext(ctx context.Context) AccesslistArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccesslistArrayOutput)
}

// AccesslistMapInput is an input type that accepts AccesslistMap and AccesslistMapOutput values.
// You can construct a concrete instance of `AccesslistMapInput` via:
//
//	AccesslistMap{ "key": AccesslistArgs{...} }
type AccesslistMapInput interface {
	pulumi.Input

	ToAccesslistMapOutput() AccesslistMapOutput
	ToAccesslistMapOutputWithContext(context.Context) AccesslistMapOutput
}

type AccesslistMap map[string]AccesslistInput

func (AccesslistMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Accesslist)(nil)).Elem()
}

func (i AccesslistMap) ToAccesslistMapOutput() AccesslistMapOutput {
	return i.ToAccesslistMapOutputWithContext(context.Background())
}

func (i AccesslistMap) ToAccesslistMapOutputWithContext(ctx context.Context) AccesslistMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccesslistMapOutput)
}

type AccesslistOutput struct{ *pulumi.OutputState }

func (AccesslistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Accesslist)(nil)).Elem()
}

func (o AccesslistOutput) ToAccesslistOutput() AccesslistOutput {
	return o
}

func (o AccesslistOutput) ToAccesslistOutputWithContext(ctx context.Context) AccesslistOutput {
	return o
}

// Comment.
func (o AccesslistOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v *Accesslist) pulumi.StringOutput { return v.Comments }).(pulumi.StringOutput)
}

func (o AccesslistOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Accesslist) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Name.
func (o AccesslistOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Accesslist) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Rule. The structure of `rule` block is documented below.
func (o AccesslistOutput) Rules() AccesslistRuleArrayOutput {
	return o.ApplyT(func(v *Accesslist) AccesslistRuleArrayOutput { return v.Rules }).(AccesslistRuleArrayOutput)
}

func (o AccesslistOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Accesslist) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type AccesslistArrayOutput struct{ *pulumi.OutputState }

func (AccesslistArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Accesslist)(nil)).Elem()
}

func (o AccesslistArrayOutput) ToAccesslistArrayOutput() AccesslistArrayOutput {
	return o
}

func (o AccesslistArrayOutput) ToAccesslistArrayOutputWithContext(ctx context.Context) AccesslistArrayOutput {
	return o
}

func (o AccesslistArrayOutput) Index(i pulumi.IntInput) AccesslistOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Accesslist {
		return vs[0].([]*Accesslist)[vs[1].(int)]
	}).(AccesslistOutput)
}

type AccesslistMapOutput struct{ *pulumi.OutputState }

func (AccesslistMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Accesslist)(nil)).Elem()
}

func (o AccesslistMapOutput) ToAccesslistMapOutput() AccesslistMapOutput {
	return o
}

func (o AccesslistMapOutput) ToAccesslistMapOutputWithContext(ctx context.Context) AccesslistMapOutput {
	return o
}

func (o AccesslistMapOutput) MapIndex(k pulumi.StringInput) AccesslistOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Accesslist {
		return vs[0].(map[string]*Accesslist)[vs[1].(string)]
	}).(AccesslistOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccesslistInput)(nil)).Elem(), &Accesslist{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccesslistArrayInput)(nil)).Elem(), AccesslistArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccesslistMapInput)(nil)).Elem(), AccesslistMap{})
	pulumi.RegisterOutputType(AccesslistOutput{})
	pulumi.RegisterOutputType(AccesslistArrayOutput{})
	pulumi.RegisterOutputType(AccesslistMapOutput{})
}
