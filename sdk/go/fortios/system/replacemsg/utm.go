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
// SystemReplacemsg Utm can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/replacemsg/utm:Utm labelname {{msg_type}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/replacemsg/utm:Utm labelname {{msg_type}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Utm struct {
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
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewUtm registers a new resource with the given unique name, arguments, and options.
func NewUtm(ctx *pulumi.Context,
	name string, args *UtmArgs, opts ...pulumi.ResourceOption) (*Utm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Utm
	err := ctx.RegisterResource("fortios:system/replacemsg/utm:Utm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUtm gets an existing Utm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUtm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UtmState, opts ...pulumi.ResourceOption) (*Utm, error) {
	var resource Utm
	err := ctx.ReadResource("fortios:system/replacemsg/utm:Utm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Utm resources.
type utmState struct {
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

type UtmState struct {
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

func (UtmState) ElementType() reflect.Type {
	return reflect.TypeOf((*utmState)(nil)).Elem()
}

type utmArgs struct {
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

// The set of arguments for constructing a Utm resource.
type UtmArgs struct {
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

func (UtmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*utmArgs)(nil)).Elem()
}

type UtmInput interface {
	pulumi.Input

	ToUtmOutput() UtmOutput
	ToUtmOutputWithContext(ctx context.Context) UtmOutput
}

func (*Utm) ElementType() reflect.Type {
	return reflect.TypeOf((**Utm)(nil)).Elem()
}

func (i *Utm) ToUtmOutput() UtmOutput {
	return i.ToUtmOutputWithContext(context.Background())
}

func (i *Utm) ToUtmOutputWithContext(ctx context.Context) UtmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UtmOutput)
}

// UtmArrayInput is an input type that accepts UtmArray and UtmArrayOutput values.
// You can construct a concrete instance of `UtmArrayInput` via:
//
//	UtmArray{ UtmArgs{...} }
type UtmArrayInput interface {
	pulumi.Input

	ToUtmArrayOutput() UtmArrayOutput
	ToUtmArrayOutputWithContext(context.Context) UtmArrayOutput
}

type UtmArray []UtmInput

func (UtmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Utm)(nil)).Elem()
}

func (i UtmArray) ToUtmArrayOutput() UtmArrayOutput {
	return i.ToUtmArrayOutputWithContext(context.Background())
}

func (i UtmArray) ToUtmArrayOutputWithContext(ctx context.Context) UtmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UtmArrayOutput)
}

// UtmMapInput is an input type that accepts UtmMap and UtmMapOutput values.
// You can construct a concrete instance of `UtmMapInput` via:
//
//	UtmMap{ "key": UtmArgs{...} }
type UtmMapInput interface {
	pulumi.Input

	ToUtmMapOutput() UtmMapOutput
	ToUtmMapOutputWithContext(context.Context) UtmMapOutput
}

type UtmMap map[string]UtmInput

func (UtmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Utm)(nil)).Elem()
}

func (i UtmMap) ToUtmMapOutput() UtmMapOutput {
	return i.ToUtmMapOutputWithContext(context.Background())
}

func (i UtmMap) ToUtmMapOutputWithContext(ctx context.Context) UtmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UtmMapOutput)
}

type UtmOutput struct{ *pulumi.OutputState }

func (UtmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Utm)(nil)).Elem()
}

func (o UtmOutput) ToUtmOutput() UtmOutput {
	return o
}

func (o UtmOutput) ToUtmOutputWithContext(ctx context.Context) UtmOutput {
	return o
}

// Message string.
func (o UtmOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Utm) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

// Format flag.
func (o UtmOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *Utm) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Header flag. Valid values: `none`, `http`, `8bit`.
func (o UtmOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *Utm) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

// Message type.
func (o UtmOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *Utm) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o UtmOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Utm) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type UtmArrayOutput struct{ *pulumi.OutputState }

func (UtmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Utm)(nil)).Elem()
}

func (o UtmArrayOutput) ToUtmArrayOutput() UtmArrayOutput {
	return o
}

func (o UtmArrayOutput) ToUtmArrayOutputWithContext(ctx context.Context) UtmArrayOutput {
	return o
}

func (o UtmArrayOutput) Index(i pulumi.IntInput) UtmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Utm {
		return vs[0].([]*Utm)[vs[1].(int)]
	}).(UtmOutput)
}

type UtmMapOutput struct{ *pulumi.OutputState }

func (UtmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Utm)(nil)).Elem()
}

func (o UtmMapOutput) ToUtmMapOutput() UtmMapOutput {
	return o
}

func (o UtmMapOutput) ToUtmMapOutputWithContext(ctx context.Context) UtmMapOutput {
	return o
}

func (o UtmMapOutput) MapIndex(k pulumi.StringInput) UtmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Utm {
		return vs[0].(map[string]*Utm)[vs[1].(string)]
	}).(UtmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UtmInput)(nil)).Elem(), &Utm{})
	pulumi.RegisterInputType(reflect.TypeOf((*UtmArrayInput)(nil)).Elem(), UtmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UtmMapInput)(nil)).Elem(), UtmMap{})
	pulumi.RegisterOutputType(UtmOutput{})
	pulumi.RegisterOutputType(UtmArrayOutput{})
	pulumi.RegisterOutputType(UtmMapOutput{})
}
