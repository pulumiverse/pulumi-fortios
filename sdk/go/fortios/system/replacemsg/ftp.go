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
// SystemReplacemsg Ftp can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/replacemsg/ftp:Ftp labelname {{msg_type}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/replacemsg/ftp:Ftp labelname {{msg_type}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Ftp struct {
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

// NewFtp registers a new resource with the given unique name, arguments, and options.
func NewFtp(ctx *pulumi.Context,
	name string, args *FtpArgs, opts ...pulumi.ResourceOption) (*Ftp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ftp
	err := ctx.RegisterResource("fortios:system/replacemsg/ftp:Ftp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFtp gets an existing Ftp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFtp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FtpState, opts ...pulumi.ResourceOption) (*Ftp, error) {
	var resource Ftp
	err := ctx.ReadResource("fortios:system/replacemsg/ftp:Ftp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ftp resources.
type ftpState struct {
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

type FtpState struct {
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

func (FtpState) ElementType() reflect.Type {
	return reflect.TypeOf((*ftpState)(nil)).Elem()
}

type ftpArgs struct {
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

// The set of arguments for constructing a Ftp resource.
type FtpArgs struct {
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

func (FtpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ftpArgs)(nil)).Elem()
}

type FtpInput interface {
	pulumi.Input

	ToFtpOutput() FtpOutput
	ToFtpOutputWithContext(ctx context.Context) FtpOutput
}

func (*Ftp) ElementType() reflect.Type {
	return reflect.TypeOf((**Ftp)(nil)).Elem()
}

func (i *Ftp) ToFtpOutput() FtpOutput {
	return i.ToFtpOutputWithContext(context.Background())
}

func (i *Ftp) ToFtpOutputWithContext(ctx context.Context) FtpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FtpOutput)
}

// FtpArrayInput is an input type that accepts FtpArray and FtpArrayOutput values.
// You can construct a concrete instance of `FtpArrayInput` via:
//
//	FtpArray{ FtpArgs{...} }
type FtpArrayInput interface {
	pulumi.Input

	ToFtpArrayOutput() FtpArrayOutput
	ToFtpArrayOutputWithContext(context.Context) FtpArrayOutput
}

type FtpArray []FtpInput

func (FtpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ftp)(nil)).Elem()
}

func (i FtpArray) ToFtpArrayOutput() FtpArrayOutput {
	return i.ToFtpArrayOutputWithContext(context.Background())
}

func (i FtpArray) ToFtpArrayOutputWithContext(ctx context.Context) FtpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FtpArrayOutput)
}

// FtpMapInput is an input type that accepts FtpMap and FtpMapOutput values.
// You can construct a concrete instance of `FtpMapInput` via:
//
//	FtpMap{ "key": FtpArgs{...} }
type FtpMapInput interface {
	pulumi.Input

	ToFtpMapOutput() FtpMapOutput
	ToFtpMapOutputWithContext(context.Context) FtpMapOutput
}

type FtpMap map[string]FtpInput

func (FtpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ftp)(nil)).Elem()
}

func (i FtpMap) ToFtpMapOutput() FtpMapOutput {
	return i.ToFtpMapOutputWithContext(context.Background())
}

func (i FtpMap) ToFtpMapOutputWithContext(ctx context.Context) FtpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FtpMapOutput)
}

type FtpOutput struct{ *pulumi.OutputState }

func (FtpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ftp)(nil)).Elem()
}

func (o FtpOutput) ToFtpOutput() FtpOutput {
	return o
}

func (o FtpOutput) ToFtpOutputWithContext(ctx context.Context) FtpOutput {
	return o
}

// Message string.
func (o FtpOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ftp) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

// Format flag.
func (o FtpOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *Ftp) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Header flag. Valid values: `none`, `http`, `8bit`.
func (o FtpOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *Ftp) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

// Message type.
func (o FtpOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *Ftp) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FtpOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ftp) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FtpArrayOutput struct{ *pulumi.OutputState }

func (FtpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ftp)(nil)).Elem()
}

func (o FtpArrayOutput) ToFtpArrayOutput() FtpArrayOutput {
	return o
}

func (o FtpArrayOutput) ToFtpArrayOutputWithContext(ctx context.Context) FtpArrayOutput {
	return o
}

func (o FtpArrayOutput) Index(i pulumi.IntInput) FtpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ftp {
		return vs[0].([]*Ftp)[vs[1].(int)]
	}).(FtpOutput)
}

type FtpMapOutput struct{ *pulumi.OutputState }

func (FtpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ftp)(nil)).Elem()
}

func (o FtpMapOutput) ToFtpMapOutput() FtpMapOutput {
	return o
}

func (o FtpMapOutput) ToFtpMapOutputWithContext(ctx context.Context) FtpMapOutput {
	return o
}

func (o FtpMapOutput) MapIndex(k pulumi.StringInput) FtpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ftp {
		return vs[0].(map[string]*Ftp)[vs[1].(string)]
	}).(FtpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FtpInput)(nil)).Elem(), &Ftp{})
	pulumi.RegisterInputType(reflect.TypeOf((*FtpArrayInput)(nil)).Elem(), FtpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FtpMapInput)(nil)).Elem(), FtpMap{})
	pulumi.RegisterOutputType(FtpOutput{})
	pulumi.RegisterOutputType(FtpArrayOutput{})
	pulumi.RegisterOutputType(FtpMapOutput{})
}
