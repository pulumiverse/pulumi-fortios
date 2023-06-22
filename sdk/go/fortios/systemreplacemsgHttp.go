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
// # SystemReplacemsg Http can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemreplacemsgHttp:SystemreplacemsgHttp labelname {{msg_type}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemreplacemsgHttp:SystemreplacemsgHttp labelname {{msg_type}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemreplacemsgHttp struct {
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

// NewSystemreplacemsgHttp registers a new resource with the given unique name, arguments, and options.
func NewSystemreplacemsgHttp(ctx *pulumi.Context,
	name string, args *SystemreplacemsgHttpArgs, opts ...pulumi.ResourceOption) (*SystemreplacemsgHttp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemreplacemsgHttp
	err := ctx.RegisterResource("fortios:index/systemreplacemsgHttp:SystemreplacemsgHttp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemreplacemsgHttp gets an existing SystemreplacemsgHttp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemreplacemsgHttp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemreplacemsgHttpState, opts ...pulumi.ResourceOption) (*SystemreplacemsgHttp, error) {
	var resource SystemreplacemsgHttp
	err := ctx.ReadResource("fortios:index/systemreplacemsgHttp:SystemreplacemsgHttp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemreplacemsgHttp resources.
type systemreplacemsgHttpState struct {
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

type SystemreplacemsgHttpState struct {
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

func (SystemreplacemsgHttpState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemreplacemsgHttpState)(nil)).Elem()
}

type systemreplacemsgHttpArgs struct {
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

// The set of arguments for constructing a SystemreplacemsgHttp resource.
type SystemreplacemsgHttpArgs struct {
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

func (SystemreplacemsgHttpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemreplacemsgHttpArgs)(nil)).Elem()
}

type SystemreplacemsgHttpInput interface {
	pulumi.Input

	ToSystemreplacemsgHttpOutput() SystemreplacemsgHttpOutput
	ToSystemreplacemsgHttpOutputWithContext(ctx context.Context) SystemreplacemsgHttpOutput
}

func (*SystemreplacemsgHttp) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemreplacemsgHttp)(nil)).Elem()
}

func (i *SystemreplacemsgHttp) ToSystemreplacemsgHttpOutput() SystemreplacemsgHttpOutput {
	return i.ToSystemreplacemsgHttpOutputWithContext(context.Background())
}

func (i *SystemreplacemsgHttp) ToSystemreplacemsgHttpOutputWithContext(ctx context.Context) SystemreplacemsgHttpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgHttpOutput)
}

// SystemreplacemsgHttpArrayInput is an input type that accepts SystemreplacemsgHttpArray and SystemreplacemsgHttpArrayOutput values.
// You can construct a concrete instance of `SystemreplacemsgHttpArrayInput` via:
//
//	SystemreplacemsgHttpArray{ SystemreplacemsgHttpArgs{...} }
type SystemreplacemsgHttpArrayInput interface {
	pulumi.Input

	ToSystemreplacemsgHttpArrayOutput() SystemreplacemsgHttpArrayOutput
	ToSystemreplacemsgHttpArrayOutputWithContext(context.Context) SystemreplacemsgHttpArrayOutput
}

type SystemreplacemsgHttpArray []SystemreplacemsgHttpInput

func (SystemreplacemsgHttpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemreplacemsgHttp)(nil)).Elem()
}

func (i SystemreplacemsgHttpArray) ToSystemreplacemsgHttpArrayOutput() SystemreplacemsgHttpArrayOutput {
	return i.ToSystemreplacemsgHttpArrayOutputWithContext(context.Background())
}

func (i SystemreplacemsgHttpArray) ToSystemreplacemsgHttpArrayOutputWithContext(ctx context.Context) SystemreplacemsgHttpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgHttpArrayOutput)
}

// SystemreplacemsgHttpMapInput is an input type that accepts SystemreplacemsgHttpMap and SystemreplacemsgHttpMapOutput values.
// You can construct a concrete instance of `SystemreplacemsgHttpMapInput` via:
//
//	SystemreplacemsgHttpMap{ "key": SystemreplacemsgHttpArgs{...} }
type SystemreplacemsgHttpMapInput interface {
	pulumi.Input

	ToSystemreplacemsgHttpMapOutput() SystemreplacemsgHttpMapOutput
	ToSystemreplacemsgHttpMapOutputWithContext(context.Context) SystemreplacemsgHttpMapOutput
}

type SystemreplacemsgHttpMap map[string]SystemreplacemsgHttpInput

func (SystemreplacemsgHttpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemreplacemsgHttp)(nil)).Elem()
}

func (i SystemreplacemsgHttpMap) ToSystemreplacemsgHttpMapOutput() SystemreplacemsgHttpMapOutput {
	return i.ToSystemreplacemsgHttpMapOutputWithContext(context.Background())
}

func (i SystemreplacemsgHttpMap) ToSystemreplacemsgHttpMapOutputWithContext(ctx context.Context) SystemreplacemsgHttpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgHttpMapOutput)
}

type SystemreplacemsgHttpOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgHttpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemreplacemsgHttp)(nil)).Elem()
}

func (o SystemreplacemsgHttpOutput) ToSystemreplacemsgHttpOutput() SystemreplacemsgHttpOutput {
	return o
}

func (o SystemreplacemsgHttpOutput) ToSystemreplacemsgHttpOutputWithContext(ctx context.Context) SystemreplacemsgHttpOutput {
	return o
}

// Message string.
func (o SystemreplacemsgHttpOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemreplacemsgHttp) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

// Format flag.
func (o SystemreplacemsgHttpOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgHttp) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Header flag. Valid values: `none`, `http`, `8bit`.
func (o SystemreplacemsgHttpOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgHttp) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

// Message type.
func (o SystemreplacemsgHttpOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgHttp) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemreplacemsgHttpOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemreplacemsgHttp) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemreplacemsgHttpArrayOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgHttpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemreplacemsgHttp)(nil)).Elem()
}

func (o SystemreplacemsgHttpArrayOutput) ToSystemreplacemsgHttpArrayOutput() SystemreplacemsgHttpArrayOutput {
	return o
}

func (o SystemreplacemsgHttpArrayOutput) ToSystemreplacemsgHttpArrayOutputWithContext(ctx context.Context) SystemreplacemsgHttpArrayOutput {
	return o
}

func (o SystemreplacemsgHttpArrayOutput) Index(i pulumi.IntInput) SystemreplacemsgHttpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemreplacemsgHttp {
		return vs[0].([]*SystemreplacemsgHttp)[vs[1].(int)]
	}).(SystemreplacemsgHttpOutput)
}

type SystemreplacemsgHttpMapOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgHttpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemreplacemsgHttp)(nil)).Elem()
}

func (o SystemreplacemsgHttpMapOutput) ToSystemreplacemsgHttpMapOutput() SystemreplacemsgHttpMapOutput {
	return o
}

func (o SystemreplacemsgHttpMapOutput) ToSystemreplacemsgHttpMapOutputWithContext(ctx context.Context) SystemreplacemsgHttpMapOutput {
	return o
}

func (o SystemreplacemsgHttpMapOutput) MapIndex(k pulumi.StringInput) SystemreplacemsgHttpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemreplacemsgHttp {
		return vs[0].(map[string]*SystemreplacemsgHttp)[vs[1].(string)]
	}).(SystemreplacemsgHttpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgHttpInput)(nil)).Elem(), &SystemreplacemsgHttp{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgHttpArrayInput)(nil)).Elem(), SystemreplacemsgHttpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgHttpMapInput)(nil)).Elem(), SystemreplacemsgHttpMap{})
	pulumi.RegisterOutputType(SystemreplacemsgHttpOutput{})
	pulumi.RegisterOutputType(SystemreplacemsgHttpArrayOutput{})
	pulumi.RegisterOutputType(SystemreplacemsgHttpMapOutput{})
}
