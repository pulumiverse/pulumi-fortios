// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package log

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure custom log fields.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/log"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := log.NewCustomfield(ctx, "trname", &log.CustomfieldArgs{
//				Fosid: pulumi.String("1"),
//				Value: pulumi.String("logteststr"),
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
// Log CustomField can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:log/customfield:Customfield labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:log/customfield:Customfield labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Customfield struct {
	pulumi.CustomResourceState

	// field ID <string>.
	Fosid pulumi.StringOutput `pulumi:"fosid"`
	// Field name (max: 15 characters).
	Name pulumi.StringOutput `pulumi:"name"`
	// Field value (max: 15 characters).
	Value pulumi.StringOutput `pulumi:"value"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewCustomfield registers a new resource with the given unique name, arguments, and options.
func NewCustomfield(ctx *pulumi.Context,
	name string, args *CustomfieldArgs, opts ...pulumi.ResourceOption) (*Customfield, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Customfield
	err := ctx.RegisterResource("fortios:log/customfield:Customfield", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomfield gets an existing Customfield resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomfield(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomfieldState, opts ...pulumi.ResourceOption) (*Customfield, error) {
	var resource Customfield
	err := ctx.ReadResource("fortios:log/customfield:Customfield", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Customfield resources.
type customfieldState struct {
	// field ID <string>.
	Fosid *string `pulumi:"fosid"`
	// Field name (max: 15 characters).
	Name *string `pulumi:"name"`
	// Field value (max: 15 characters).
	Value *string `pulumi:"value"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type CustomfieldState struct {
	// field ID <string>.
	Fosid pulumi.StringPtrInput
	// Field name (max: 15 characters).
	Name pulumi.StringPtrInput
	// Field value (max: 15 characters).
	Value pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (CustomfieldState) ElementType() reflect.Type {
	return reflect.TypeOf((*customfieldState)(nil)).Elem()
}

type customfieldArgs struct {
	// field ID <string>.
	Fosid *string `pulumi:"fosid"`
	// Field name (max: 15 characters).
	Name *string `pulumi:"name"`
	// Field value (max: 15 characters).
	Value string `pulumi:"value"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Customfield resource.
type CustomfieldArgs struct {
	// field ID <string>.
	Fosid pulumi.StringPtrInput
	// Field name (max: 15 characters).
	Name pulumi.StringPtrInput
	// Field value (max: 15 characters).
	Value pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (CustomfieldArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customfieldArgs)(nil)).Elem()
}

type CustomfieldInput interface {
	pulumi.Input

	ToCustomfieldOutput() CustomfieldOutput
	ToCustomfieldOutputWithContext(ctx context.Context) CustomfieldOutput
}

func (*Customfield) ElementType() reflect.Type {
	return reflect.TypeOf((**Customfield)(nil)).Elem()
}

func (i *Customfield) ToCustomfieldOutput() CustomfieldOutput {
	return i.ToCustomfieldOutputWithContext(context.Background())
}

func (i *Customfield) ToCustomfieldOutputWithContext(ctx context.Context) CustomfieldOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomfieldOutput)
}

// CustomfieldArrayInput is an input type that accepts CustomfieldArray and CustomfieldArrayOutput values.
// You can construct a concrete instance of `CustomfieldArrayInput` via:
//
//	CustomfieldArray{ CustomfieldArgs{...} }
type CustomfieldArrayInput interface {
	pulumi.Input

	ToCustomfieldArrayOutput() CustomfieldArrayOutput
	ToCustomfieldArrayOutputWithContext(context.Context) CustomfieldArrayOutput
}

type CustomfieldArray []CustomfieldInput

func (CustomfieldArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Customfield)(nil)).Elem()
}

func (i CustomfieldArray) ToCustomfieldArrayOutput() CustomfieldArrayOutput {
	return i.ToCustomfieldArrayOutputWithContext(context.Background())
}

func (i CustomfieldArray) ToCustomfieldArrayOutputWithContext(ctx context.Context) CustomfieldArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomfieldArrayOutput)
}

// CustomfieldMapInput is an input type that accepts CustomfieldMap and CustomfieldMapOutput values.
// You can construct a concrete instance of `CustomfieldMapInput` via:
//
//	CustomfieldMap{ "key": CustomfieldArgs{...} }
type CustomfieldMapInput interface {
	pulumi.Input

	ToCustomfieldMapOutput() CustomfieldMapOutput
	ToCustomfieldMapOutputWithContext(context.Context) CustomfieldMapOutput
}

type CustomfieldMap map[string]CustomfieldInput

func (CustomfieldMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Customfield)(nil)).Elem()
}

func (i CustomfieldMap) ToCustomfieldMapOutput() CustomfieldMapOutput {
	return i.ToCustomfieldMapOutputWithContext(context.Background())
}

func (i CustomfieldMap) ToCustomfieldMapOutputWithContext(ctx context.Context) CustomfieldMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomfieldMapOutput)
}

type CustomfieldOutput struct{ *pulumi.OutputState }

func (CustomfieldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Customfield)(nil)).Elem()
}

func (o CustomfieldOutput) ToCustomfieldOutput() CustomfieldOutput {
	return o
}

func (o CustomfieldOutput) ToCustomfieldOutputWithContext(ctx context.Context) CustomfieldOutput {
	return o
}

// field ID <string>.
func (o CustomfieldOutput) Fosid() pulumi.StringOutput {
	return o.ApplyT(func(v *Customfield) pulumi.StringOutput { return v.Fosid }).(pulumi.StringOutput)
}

// Field name (max: 15 characters).
func (o CustomfieldOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Customfield) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Field value (max: 15 characters).
func (o CustomfieldOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *Customfield) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o CustomfieldOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Customfield) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type CustomfieldArrayOutput struct{ *pulumi.OutputState }

func (CustomfieldArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Customfield)(nil)).Elem()
}

func (o CustomfieldArrayOutput) ToCustomfieldArrayOutput() CustomfieldArrayOutput {
	return o
}

func (o CustomfieldArrayOutput) ToCustomfieldArrayOutputWithContext(ctx context.Context) CustomfieldArrayOutput {
	return o
}

func (o CustomfieldArrayOutput) Index(i pulumi.IntInput) CustomfieldOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Customfield {
		return vs[0].([]*Customfield)[vs[1].(int)]
	}).(CustomfieldOutput)
}

type CustomfieldMapOutput struct{ *pulumi.OutputState }

func (CustomfieldMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Customfield)(nil)).Elem()
}

func (o CustomfieldMapOutput) ToCustomfieldMapOutput() CustomfieldMapOutput {
	return o
}

func (o CustomfieldMapOutput) ToCustomfieldMapOutputWithContext(ctx context.Context) CustomfieldMapOutput {
	return o
}

func (o CustomfieldMapOutput) MapIndex(k pulumi.StringInput) CustomfieldOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Customfield {
		return vs[0].(map[string]*Customfield)[vs[1].(string)]
	}).(CustomfieldOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomfieldInput)(nil)).Elem(), &Customfield{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomfieldArrayInput)(nil)).Elem(), CustomfieldArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomfieldMapInput)(nil)).Elem(), CustomfieldMap{})
	pulumi.RegisterOutputType(CustomfieldOutput{})
	pulumi.RegisterOutputType(CustomfieldArrayOutput{})
	pulumi.RegisterOutputType(CustomfieldMapOutput{})
}
