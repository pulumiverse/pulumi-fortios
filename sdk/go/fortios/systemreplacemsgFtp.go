// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Replacement messages.
//
// ## Import
//
// # SystemReplacemsg Ftp can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemreplacemsgFtp:SystemreplacemsgFtp labelname {{msg_type}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemreplacemsgFtp:SystemreplacemsgFtp labelname {{msg_type}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemreplacemsgFtp struct {
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

// NewSystemreplacemsgFtp registers a new resource with the given unique name, arguments, and options.
func NewSystemreplacemsgFtp(ctx *pulumi.Context,
	name string, args *SystemreplacemsgFtpArgs, opts ...pulumi.ResourceOption) (*SystemreplacemsgFtp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemreplacemsgFtp
	err := ctx.RegisterResource("fortios:index/systemreplacemsgFtp:SystemreplacemsgFtp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemreplacemsgFtp gets an existing SystemreplacemsgFtp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemreplacemsgFtp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemreplacemsgFtpState, opts ...pulumi.ResourceOption) (*SystemreplacemsgFtp, error) {
	var resource SystemreplacemsgFtp
	err := ctx.ReadResource("fortios:index/systemreplacemsgFtp:SystemreplacemsgFtp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemreplacemsgFtp resources.
type systemreplacemsgFtpState struct {
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

type SystemreplacemsgFtpState struct {
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

func (SystemreplacemsgFtpState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemreplacemsgFtpState)(nil)).Elem()
}

type systemreplacemsgFtpArgs struct {
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

// The set of arguments for constructing a SystemreplacemsgFtp resource.
type SystemreplacemsgFtpArgs struct {
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

func (SystemreplacemsgFtpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemreplacemsgFtpArgs)(nil)).Elem()
}

type SystemreplacemsgFtpInput interface {
	pulumi.Input

	ToSystemreplacemsgFtpOutput() SystemreplacemsgFtpOutput
	ToSystemreplacemsgFtpOutputWithContext(ctx context.Context) SystemreplacemsgFtpOutput
}

func (*SystemreplacemsgFtp) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemreplacemsgFtp)(nil)).Elem()
}

func (i *SystemreplacemsgFtp) ToSystemreplacemsgFtpOutput() SystemreplacemsgFtpOutput {
	return i.ToSystemreplacemsgFtpOutputWithContext(context.Background())
}

func (i *SystemreplacemsgFtp) ToSystemreplacemsgFtpOutputWithContext(ctx context.Context) SystemreplacemsgFtpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgFtpOutput)
}

// SystemreplacemsgFtpArrayInput is an input type that accepts SystemreplacemsgFtpArray and SystemreplacemsgFtpArrayOutput values.
// You can construct a concrete instance of `SystemreplacemsgFtpArrayInput` via:
//
//	SystemreplacemsgFtpArray{ SystemreplacemsgFtpArgs{...} }
type SystemreplacemsgFtpArrayInput interface {
	pulumi.Input

	ToSystemreplacemsgFtpArrayOutput() SystemreplacemsgFtpArrayOutput
	ToSystemreplacemsgFtpArrayOutputWithContext(context.Context) SystemreplacemsgFtpArrayOutput
}

type SystemreplacemsgFtpArray []SystemreplacemsgFtpInput

func (SystemreplacemsgFtpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemreplacemsgFtp)(nil)).Elem()
}

func (i SystemreplacemsgFtpArray) ToSystemreplacemsgFtpArrayOutput() SystemreplacemsgFtpArrayOutput {
	return i.ToSystemreplacemsgFtpArrayOutputWithContext(context.Background())
}

func (i SystemreplacemsgFtpArray) ToSystemreplacemsgFtpArrayOutputWithContext(ctx context.Context) SystemreplacemsgFtpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgFtpArrayOutput)
}

// SystemreplacemsgFtpMapInput is an input type that accepts SystemreplacemsgFtpMap and SystemreplacemsgFtpMapOutput values.
// You can construct a concrete instance of `SystemreplacemsgFtpMapInput` via:
//
//	SystemreplacemsgFtpMap{ "key": SystemreplacemsgFtpArgs{...} }
type SystemreplacemsgFtpMapInput interface {
	pulumi.Input

	ToSystemreplacemsgFtpMapOutput() SystemreplacemsgFtpMapOutput
	ToSystemreplacemsgFtpMapOutputWithContext(context.Context) SystemreplacemsgFtpMapOutput
}

type SystemreplacemsgFtpMap map[string]SystemreplacemsgFtpInput

func (SystemreplacemsgFtpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemreplacemsgFtp)(nil)).Elem()
}

func (i SystemreplacemsgFtpMap) ToSystemreplacemsgFtpMapOutput() SystemreplacemsgFtpMapOutput {
	return i.ToSystemreplacemsgFtpMapOutputWithContext(context.Background())
}

func (i SystemreplacemsgFtpMap) ToSystemreplacemsgFtpMapOutputWithContext(ctx context.Context) SystemreplacemsgFtpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgFtpMapOutput)
}

type SystemreplacemsgFtpOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgFtpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemreplacemsgFtp)(nil)).Elem()
}

func (o SystemreplacemsgFtpOutput) ToSystemreplacemsgFtpOutput() SystemreplacemsgFtpOutput {
	return o
}

func (o SystemreplacemsgFtpOutput) ToSystemreplacemsgFtpOutputWithContext(ctx context.Context) SystemreplacemsgFtpOutput {
	return o
}

// Message string.
func (o SystemreplacemsgFtpOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemreplacemsgFtp) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

// Format flag.
func (o SystemreplacemsgFtpOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgFtp) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Header flag. Valid values: `none`, `http`, `8bit`.
func (o SystemreplacemsgFtpOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgFtp) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

// Message type.
func (o SystemreplacemsgFtpOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgFtp) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemreplacemsgFtpOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemreplacemsgFtp) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemreplacemsgFtpArrayOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgFtpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemreplacemsgFtp)(nil)).Elem()
}

func (o SystemreplacemsgFtpArrayOutput) ToSystemreplacemsgFtpArrayOutput() SystemreplacemsgFtpArrayOutput {
	return o
}

func (o SystemreplacemsgFtpArrayOutput) ToSystemreplacemsgFtpArrayOutputWithContext(ctx context.Context) SystemreplacemsgFtpArrayOutput {
	return o
}

func (o SystemreplacemsgFtpArrayOutput) Index(i pulumi.IntInput) SystemreplacemsgFtpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemreplacemsgFtp {
		return vs[0].([]*SystemreplacemsgFtp)[vs[1].(int)]
	}).(SystemreplacemsgFtpOutput)
}

type SystemreplacemsgFtpMapOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgFtpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemreplacemsgFtp)(nil)).Elem()
}

func (o SystemreplacemsgFtpMapOutput) ToSystemreplacemsgFtpMapOutput() SystemreplacemsgFtpMapOutput {
	return o
}

func (o SystemreplacemsgFtpMapOutput) ToSystemreplacemsgFtpMapOutputWithContext(ctx context.Context) SystemreplacemsgFtpMapOutput {
	return o
}

func (o SystemreplacemsgFtpMapOutput) MapIndex(k pulumi.StringInput) SystemreplacemsgFtpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemreplacemsgFtp {
		return vs[0].(map[string]*SystemreplacemsgFtp)[vs[1].(string)]
	}).(SystemreplacemsgFtpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgFtpInput)(nil)).Elem(), &SystemreplacemsgFtp{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgFtpArrayInput)(nil)).Elem(), SystemreplacemsgFtpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgFtpMapInput)(nil)).Elem(), SystemreplacemsgFtpMap{})
	pulumi.RegisterOutputType(SystemreplacemsgFtpOutput{})
	pulumi.RegisterOutputType(SystemreplacemsgFtpArrayOutput{})
	pulumi.RegisterOutputType(SystemreplacemsgFtpMapOutput{})
}
