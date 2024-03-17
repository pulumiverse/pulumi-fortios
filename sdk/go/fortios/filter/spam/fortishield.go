// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure FortiGuard - AntiSpam. Applies to FortiOS Version `<= 6.2.0`.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/filter"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := filter.NewFortishield(ctx, "trname", &filter.FortishieldArgs{
//				SpamSubmitForce:   pulumi.String("enable"),
//				SpamSubmitSrv:     pulumi.String("www.nospammer.net"),
//				SpamSubmitTxt2htm: pulumi.String("enable"),
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
// Spamfilter Fortishield can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:filter/spam/fortishield:Fortishield labelname SpamfilterFortishield
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:filter/spam/fortishield:Fortishield labelname SpamfilterFortishield
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Fortishield struct {
	pulumi.CustomResourceState

	// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
	SpamSubmitForce pulumi.StringOutput `pulumi:"spamSubmitForce"`
	// Hostname of the spam submission server.
	SpamSubmitSrv pulumi.StringOutput `pulumi:"spamSubmitSrv"`
	// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
	SpamSubmitTxt2htm pulumi.StringOutput `pulumi:"spamSubmitTxt2htm"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFortishield registers a new resource with the given unique name, arguments, and options.
func NewFortishield(ctx *pulumi.Context,
	name string, args *FortishieldArgs, opts ...pulumi.ResourceOption) (*Fortishield, error) {
	if args == nil {
		args = &FortishieldArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Fortishield
	err := ctx.RegisterResource("fortios:filter/spam/fortishield:Fortishield", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortishield gets an existing Fortishield resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortishield(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortishieldState, opts ...pulumi.ResourceOption) (*Fortishield, error) {
	var resource Fortishield
	err := ctx.ReadResource("fortios:filter/spam/fortishield:Fortishield", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fortishield resources.
type fortishieldState struct {
	// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
	SpamSubmitForce *string `pulumi:"spamSubmitForce"`
	// Hostname of the spam submission server.
	SpamSubmitSrv *string `pulumi:"spamSubmitSrv"`
	// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
	SpamSubmitTxt2htm *string `pulumi:"spamSubmitTxt2htm"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FortishieldState struct {
	// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
	SpamSubmitForce pulumi.StringPtrInput
	// Hostname of the spam submission server.
	SpamSubmitSrv pulumi.StringPtrInput
	// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
	SpamSubmitTxt2htm pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FortishieldState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortishieldState)(nil)).Elem()
}

type fortishieldArgs struct {
	// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
	SpamSubmitForce *string `pulumi:"spamSubmitForce"`
	// Hostname of the spam submission server.
	SpamSubmitSrv *string `pulumi:"spamSubmitSrv"`
	// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
	SpamSubmitTxt2htm *string `pulumi:"spamSubmitTxt2htm"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Fortishield resource.
type FortishieldArgs struct {
	// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
	SpamSubmitForce pulumi.StringPtrInput
	// Hostname of the spam submission server.
	SpamSubmitSrv pulumi.StringPtrInput
	// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
	SpamSubmitTxt2htm pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FortishieldArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortishieldArgs)(nil)).Elem()
}

type FortishieldInput interface {
	pulumi.Input

	ToFortishieldOutput() FortishieldOutput
	ToFortishieldOutputWithContext(ctx context.Context) FortishieldOutput
}

func (*Fortishield) ElementType() reflect.Type {
	return reflect.TypeOf((**Fortishield)(nil)).Elem()
}

func (i *Fortishield) ToFortishieldOutput() FortishieldOutput {
	return i.ToFortishieldOutputWithContext(context.Background())
}

func (i *Fortishield) ToFortishieldOutputWithContext(ctx context.Context) FortishieldOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortishieldOutput)
}

// FortishieldArrayInput is an input type that accepts FortishieldArray and FortishieldArrayOutput values.
// You can construct a concrete instance of `FortishieldArrayInput` via:
//
//	FortishieldArray{ FortishieldArgs{...} }
type FortishieldArrayInput interface {
	pulumi.Input

	ToFortishieldArrayOutput() FortishieldArrayOutput
	ToFortishieldArrayOutputWithContext(context.Context) FortishieldArrayOutput
}

type FortishieldArray []FortishieldInput

func (FortishieldArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fortishield)(nil)).Elem()
}

func (i FortishieldArray) ToFortishieldArrayOutput() FortishieldArrayOutput {
	return i.ToFortishieldArrayOutputWithContext(context.Background())
}

func (i FortishieldArray) ToFortishieldArrayOutputWithContext(ctx context.Context) FortishieldArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortishieldArrayOutput)
}

// FortishieldMapInput is an input type that accepts FortishieldMap and FortishieldMapOutput values.
// You can construct a concrete instance of `FortishieldMapInput` via:
//
//	FortishieldMap{ "key": FortishieldArgs{...} }
type FortishieldMapInput interface {
	pulumi.Input

	ToFortishieldMapOutput() FortishieldMapOutput
	ToFortishieldMapOutputWithContext(context.Context) FortishieldMapOutput
}

type FortishieldMap map[string]FortishieldInput

func (FortishieldMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fortishield)(nil)).Elem()
}

func (i FortishieldMap) ToFortishieldMapOutput() FortishieldMapOutput {
	return i.ToFortishieldMapOutputWithContext(context.Background())
}

func (i FortishieldMap) ToFortishieldMapOutputWithContext(ctx context.Context) FortishieldMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortishieldMapOutput)
}

type FortishieldOutput struct{ *pulumi.OutputState }

func (FortishieldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fortishield)(nil)).Elem()
}

func (o FortishieldOutput) ToFortishieldOutput() FortishieldOutput {
	return o
}

func (o FortishieldOutput) ToFortishieldOutputWithContext(ctx context.Context) FortishieldOutput {
	return o
}

// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
func (o FortishieldOutput) SpamSubmitForce() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortishield) pulumi.StringOutput { return v.SpamSubmitForce }).(pulumi.StringOutput)
}

// Hostname of the spam submission server.
func (o FortishieldOutput) SpamSubmitSrv() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortishield) pulumi.StringOutput { return v.SpamSubmitSrv }).(pulumi.StringOutput)
}

// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
func (o FortishieldOutput) SpamSubmitTxt2htm() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortishield) pulumi.StringOutput { return v.SpamSubmitTxt2htm }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FortishieldOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Fortishield) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FortishieldArrayOutput struct{ *pulumi.OutputState }

func (FortishieldArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fortishield)(nil)).Elem()
}

func (o FortishieldArrayOutput) ToFortishieldArrayOutput() FortishieldArrayOutput {
	return o
}

func (o FortishieldArrayOutput) ToFortishieldArrayOutputWithContext(ctx context.Context) FortishieldArrayOutput {
	return o
}

func (o FortishieldArrayOutput) Index(i pulumi.IntInput) FortishieldOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Fortishield {
		return vs[0].([]*Fortishield)[vs[1].(int)]
	}).(FortishieldOutput)
}

type FortishieldMapOutput struct{ *pulumi.OutputState }

func (FortishieldMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fortishield)(nil)).Elem()
}

func (o FortishieldMapOutput) ToFortishieldMapOutput() FortishieldMapOutput {
	return o
}

func (o FortishieldMapOutput) ToFortishieldMapOutputWithContext(ctx context.Context) FortishieldMapOutput {
	return o
}

func (o FortishieldMapOutput) MapIndex(k pulumi.StringInput) FortishieldOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Fortishield {
		return vs[0].(map[string]*Fortishield)[vs[1].(string)]
	}).(FortishieldOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortishieldInput)(nil)).Elem(), &Fortishield{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortishieldArrayInput)(nil)).Elem(), FortishieldArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortishieldMapInput)(nil)).Elem(), FortishieldMap{})
	pulumi.RegisterOutputType(FortishieldOutput{})
	pulumi.RegisterOutputType(FortishieldArrayOutput{})
	pulumi.RegisterOutputType(FortishieldMapOutput{})
}
