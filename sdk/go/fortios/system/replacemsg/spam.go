// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package replacemsg

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Replacement messages.
//
// ## Import
//
// SystemReplacemsg Spam can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/replacemsg/spam:Spam labelname {{msg_type}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/replacemsg/spam:Spam labelname {{msg_type}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Spam struct {
	pulumi.CustomResourceState

	// Message string.
	Buffer pulumi.StringPtrOutput `pulumi:"buffer"`
	// Format flag.
	Format pulumi.StringOutput `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringOutput `pulumi:"header"`
	// Message type.
	MsgType pulumi.StringOutput `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSpam registers a new resource with the given unique name, arguments, and options.
func NewSpam(ctx *pulumi.Context,
	name string, args *SpamArgs, opts ...pulumi.ResourceOption) (*Spam, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Spam
	err := ctx.RegisterResource("fortios:system/replacemsg/spam:Spam", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpam gets an existing Spam resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpamState, opts ...pulumi.ResourceOption) (*Spam, error) {
	var resource Spam
	err := ctx.ReadResource("fortios:system/replacemsg/spam:Spam", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Spam resources.
type spamState struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType *string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SpamState struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SpamState) ElementType() reflect.Type {
	return reflect.TypeOf((*spamState)(nil)).Elem()
}

type spamArgs struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Spam resource.
type SpamArgs struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SpamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spamArgs)(nil)).Elem()
}

type SpamInput interface {
	pulumi.Input

	ToSpamOutput() SpamOutput
	ToSpamOutputWithContext(ctx context.Context) SpamOutput
}

func (*Spam) ElementType() reflect.Type {
	return reflect.TypeOf((**Spam)(nil)).Elem()
}

func (i *Spam) ToSpamOutput() SpamOutput {
	return i.ToSpamOutputWithContext(context.Background())
}

func (i *Spam) ToSpamOutputWithContext(ctx context.Context) SpamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamOutput)
}

// SpamArrayInput is an input type that accepts SpamArray and SpamArrayOutput values.
// You can construct a concrete instance of `SpamArrayInput` via:
//
//	SpamArray{ SpamArgs{...} }
type SpamArrayInput interface {
	pulumi.Input

	ToSpamArrayOutput() SpamArrayOutput
	ToSpamArrayOutputWithContext(context.Context) SpamArrayOutput
}

type SpamArray []SpamInput

func (SpamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Spam)(nil)).Elem()
}

func (i SpamArray) ToSpamArrayOutput() SpamArrayOutput {
	return i.ToSpamArrayOutputWithContext(context.Background())
}

func (i SpamArray) ToSpamArrayOutputWithContext(ctx context.Context) SpamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamArrayOutput)
}

// SpamMapInput is an input type that accepts SpamMap and SpamMapOutput values.
// You can construct a concrete instance of `SpamMapInput` via:
//
//	SpamMap{ "key": SpamArgs{...} }
type SpamMapInput interface {
	pulumi.Input

	ToSpamMapOutput() SpamMapOutput
	ToSpamMapOutputWithContext(context.Context) SpamMapOutput
}

type SpamMap map[string]SpamInput

func (SpamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Spam)(nil)).Elem()
}

func (i SpamMap) ToSpamMapOutput() SpamMapOutput {
	return i.ToSpamMapOutputWithContext(context.Background())
}

func (i SpamMap) ToSpamMapOutputWithContext(ctx context.Context) SpamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamMapOutput)
}

type SpamOutput struct{ *pulumi.OutputState }

func (SpamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Spam)(nil)).Elem()
}

func (o SpamOutput) ToSpamOutput() SpamOutput {
	return o
}

func (o SpamOutput) ToSpamOutputWithContext(ctx context.Context) SpamOutput {
	return o
}

// Message string.
func (o SpamOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Spam) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

// Format flag.
func (o SpamOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *Spam) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Header flag. Valid values: `none`, `http`, `8bit`.
func (o SpamOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *Spam) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

// Message type.
func (o SpamOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *Spam) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SpamOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Spam) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SpamArrayOutput struct{ *pulumi.OutputState }

func (SpamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Spam)(nil)).Elem()
}

func (o SpamArrayOutput) ToSpamArrayOutput() SpamArrayOutput {
	return o
}

func (o SpamArrayOutput) ToSpamArrayOutputWithContext(ctx context.Context) SpamArrayOutput {
	return o
}

func (o SpamArrayOutput) Index(i pulumi.IntInput) SpamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Spam {
		return vs[0].([]*Spam)[vs[1].(int)]
	}).(SpamOutput)
}

type SpamMapOutput struct{ *pulumi.OutputState }

func (SpamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Spam)(nil)).Elem()
}

func (o SpamMapOutput) ToSpamMapOutput() SpamMapOutput {
	return o
}

func (o SpamMapOutput) ToSpamMapOutputWithContext(ctx context.Context) SpamMapOutput {
	return o
}

func (o SpamMapOutput) MapIndex(k pulumi.StringInput) SpamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Spam {
		return vs[0].(map[string]*Spam)[vs[1].(string)]
	}).(SpamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpamInput)(nil)).Elem(), &Spam{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpamArrayInput)(nil)).Elem(), SpamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpamMapInput)(nil)).Elem(), SpamMap{})
	pulumi.RegisterOutputType(SpamOutput{})
	pulumi.RegisterOutputType(SpamArrayOutput{})
	pulumi.RegisterOutputType(SpamMapOutput{})
}
