// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package endpointcontrol

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Registered FortiClient list. Applies to FortiOS Version `<= 6.2.0`.
//
// ## Import
//
// EndpointControl RegisteredForticlient can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:endpointcontrol/registeredforticlient:Registeredforticlient labelname {{uid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:endpointcontrol/registeredforticlient:Registeredforticlient labelname {{uid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Registeredforticlient struct {
	pulumi.CustomResourceState

	// FortiClient registration flag.
	Flag pulumi.IntOutput `pulumi:"flag"`
	// Endpoint IP address.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// Endpoint MAC address.
	Mac pulumi.StringOutput `pulumi:"mac"`
	// Registering FortiGate SN.
	RegFortigate pulumi.StringOutput `pulumi:"regFortigate"`
	// FortiClient registration status.
	Status pulumi.IntOutput `pulumi:"status"`
	// FortiClient UID.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Registering vdom.
	Vdom pulumi.StringOutput `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewRegisteredforticlient registers a new resource with the given unique name, arguments, and options.
func NewRegisteredforticlient(ctx *pulumi.Context,
	name string, args *RegisteredforticlientArgs, opts ...pulumi.ResourceOption) (*Registeredforticlient, error) {
	if args == nil {
		args = &RegisteredforticlientArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Registeredforticlient
	err := ctx.RegisterResource("fortios:endpointcontrol/registeredforticlient:Registeredforticlient", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegisteredforticlient gets an existing Registeredforticlient resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegisteredforticlient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegisteredforticlientState, opts ...pulumi.ResourceOption) (*Registeredforticlient, error) {
	var resource Registeredforticlient
	err := ctx.ReadResource("fortios:endpointcontrol/registeredforticlient:Registeredforticlient", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Registeredforticlient resources.
type registeredforticlientState struct {
	// FortiClient registration flag.
	Flag *int `pulumi:"flag"`
	// Endpoint IP address.
	Ip *string `pulumi:"ip"`
	// Endpoint MAC address.
	Mac *string `pulumi:"mac"`
	// Registering FortiGate SN.
	RegFortigate *string `pulumi:"regFortigate"`
	// FortiClient registration status.
	Status *int `pulumi:"status"`
	// FortiClient UID.
	Uid *string `pulumi:"uid"`
	// Registering vdom.
	Vdom *string `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RegisteredforticlientState struct {
	// FortiClient registration flag.
	Flag pulumi.IntPtrInput
	// Endpoint IP address.
	Ip pulumi.StringPtrInput
	// Endpoint MAC address.
	Mac pulumi.StringPtrInput
	// Registering FortiGate SN.
	RegFortigate pulumi.StringPtrInput
	// FortiClient registration status.
	Status pulumi.IntPtrInput
	// FortiClient UID.
	Uid pulumi.StringPtrInput
	// Registering vdom.
	Vdom pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RegisteredforticlientState) ElementType() reflect.Type {
	return reflect.TypeOf((*registeredforticlientState)(nil)).Elem()
}

type registeredforticlientArgs struct {
	// FortiClient registration flag.
	Flag *int `pulumi:"flag"`
	// Endpoint IP address.
	Ip *string `pulumi:"ip"`
	// Endpoint MAC address.
	Mac *string `pulumi:"mac"`
	// Registering FortiGate SN.
	RegFortigate *string `pulumi:"regFortigate"`
	// FortiClient registration status.
	Status *int `pulumi:"status"`
	// FortiClient UID.
	Uid *string `pulumi:"uid"`
	// Registering vdom.
	Vdom *string `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Registeredforticlient resource.
type RegisteredforticlientArgs struct {
	// FortiClient registration flag.
	Flag pulumi.IntPtrInput
	// Endpoint IP address.
	Ip pulumi.StringPtrInput
	// Endpoint MAC address.
	Mac pulumi.StringPtrInput
	// Registering FortiGate SN.
	RegFortigate pulumi.StringPtrInput
	// FortiClient registration status.
	Status pulumi.IntPtrInput
	// FortiClient UID.
	Uid pulumi.StringPtrInput
	// Registering vdom.
	Vdom pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RegisteredforticlientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registeredforticlientArgs)(nil)).Elem()
}

type RegisteredforticlientInput interface {
	pulumi.Input

	ToRegisteredforticlientOutput() RegisteredforticlientOutput
	ToRegisteredforticlientOutputWithContext(ctx context.Context) RegisteredforticlientOutput
}

func (*Registeredforticlient) ElementType() reflect.Type {
	return reflect.TypeOf((**Registeredforticlient)(nil)).Elem()
}

func (i *Registeredforticlient) ToRegisteredforticlientOutput() RegisteredforticlientOutput {
	return i.ToRegisteredforticlientOutputWithContext(context.Background())
}

func (i *Registeredforticlient) ToRegisteredforticlientOutputWithContext(ctx context.Context) RegisteredforticlientOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegisteredforticlientOutput)
}

// RegisteredforticlientArrayInput is an input type that accepts RegisteredforticlientArray and RegisteredforticlientArrayOutput values.
// You can construct a concrete instance of `RegisteredforticlientArrayInput` via:
//
//	RegisteredforticlientArray{ RegisteredforticlientArgs{...} }
type RegisteredforticlientArrayInput interface {
	pulumi.Input

	ToRegisteredforticlientArrayOutput() RegisteredforticlientArrayOutput
	ToRegisteredforticlientArrayOutputWithContext(context.Context) RegisteredforticlientArrayOutput
}

type RegisteredforticlientArray []RegisteredforticlientInput

func (RegisteredforticlientArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Registeredforticlient)(nil)).Elem()
}

func (i RegisteredforticlientArray) ToRegisteredforticlientArrayOutput() RegisteredforticlientArrayOutput {
	return i.ToRegisteredforticlientArrayOutputWithContext(context.Background())
}

func (i RegisteredforticlientArray) ToRegisteredforticlientArrayOutputWithContext(ctx context.Context) RegisteredforticlientArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegisteredforticlientArrayOutput)
}

// RegisteredforticlientMapInput is an input type that accepts RegisteredforticlientMap and RegisteredforticlientMapOutput values.
// You can construct a concrete instance of `RegisteredforticlientMapInput` via:
//
//	RegisteredforticlientMap{ "key": RegisteredforticlientArgs{...} }
type RegisteredforticlientMapInput interface {
	pulumi.Input

	ToRegisteredforticlientMapOutput() RegisteredforticlientMapOutput
	ToRegisteredforticlientMapOutputWithContext(context.Context) RegisteredforticlientMapOutput
}

type RegisteredforticlientMap map[string]RegisteredforticlientInput

func (RegisteredforticlientMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Registeredforticlient)(nil)).Elem()
}

func (i RegisteredforticlientMap) ToRegisteredforticlientMapOutput() RegisteredforticlientMapOutput {
	return i.ToRegisteredforticlientMapOutputWithContext(context.Background())
}

func (i RegisteredforticlientMap) ToRegisteredforticlientMapOutputWithContext(ctx context.Context) RegisteredforticlientMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegisteredforticlientMapOutput)
}

type RegisteredforticlientOutput struct{ *pulumi.OutputState }

func (RegisteredforticlientOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Registeredforticlient)(nil)).Elem()
}

func (o RegisteredforticlientOutput) ToRegisteredforticlientOutput() RegisteredforticlientOutput {
	return o
}

func (o RegisteredforticlientOutput) ToRegisteredforticlientOutputWithContext(ctx context.Context) RegisteredforticlientOutput {
	return o
}

// FortiClient registration flag.
func (o RegisteredforticlientOutput) Flag() pulumi.IntOutput {
	return o.ApplyT(func(v *Registeredforticlient) pulumi.IntOutput { return v.Flag }).(pulumi.IntOutput)
}

// Endpoint IP address.
func (o RegisteredforticlientOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *Registeredforticlient) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// Endpoint MAC address.
func (o RegisteredforticlientOutput) Mac() pulumi.StringOutput {
	return o.ApplyT(func(v *Registeredforticlient) pulumi.StringOutput { return v.Mac }).(pulumi.StringOutput)
}

// Registering FortiGate SN.
func (o RegisteredforticlientOutput) RegFortigate() pulumi.StringOutput {
	return o.ApplyT(func(v *Registeredforticlient) pulumi.StringOutput { return v.RegFortigate }).(pulumi.StringOutput)
}

// FortiClient registration status.
func (o RegisteredforticlientOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *Registeredforticlient) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// FortiClient UID.
func (o RegisteredforticlientOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Registeredforticlient) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Registering vdom.
func (o RegisteredforticlientOutput) Vdom() pulumi.StringOutput {
	return o.ApplyT(func(v *Registeredforticlient) pulumi.StringOutput { return v.Vdom }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o RegisteredforticlientOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Registeredforticlient) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type RegisteredforticlientArrayOutput struct{ *pulumi.OutputState }

func (RegisteredforticlientArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Registeredforticlient)(nil)).Elem()
}

func (o RegisteredforticlientArrayOutput) ToRegisteredforticlientArrayOutput() RegisteredforticlientArrayOutput {
	return o
}

func (o RegisteredforticlientArrayOutput) ToRegisteredforticlientArrayOutputWithContext(ctx context.Context) RegisteredforticlientArrayOutput {
	return o
}

func (o RegisteredforticlientArrayOutput) Index(i pulumi.IntInput) RegisteredforticlientOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Registeredforticlient {
		return vs[0].([]*Registeredforticlient)[vs[1].(int)]
	}).(RegisteredforticlientOutput)
}

type RegisteredforticlientMapOutput struct{ *pulumi.OutputState }

func (RegisteredforticlientMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Registeredforticlient)(nil)).Elem()
}

func (o RegisteredforticlientMapOutput) ToRegisteredforticlientMapOutput() RegisteredforticlientMapOutput {
	return o
}

func (o RegisteredforticlientMapOutput) ToRegisteredforticlientMapOutputWithContext(ctx context.Context) RegisteredforticlientMapOutput {
	return o
}

func (o RegisteredforticlientMapOutput) MapIndex(k pulumi.StringInput) RegisteredforticlientOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Registeredforticlient {
		return vs[0].(map[string]*Registeredforticlient)[vs[1].(string)]
	}).(RegisteredforticlientOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegisteredforticlientInput)(nil)).Elem(), &Registeredforticlient{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegisteredforticlientArrayInput)(nil)).Elem(), RegisteredforticlientArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegisteredforticlientMapInput)(nil)).Elem(), RegisteredforticlientMap{})
	pulumi.RegisterOutputType(RegisteredforticlientOutput{})
	pulumi.RegisterOutputType(RegisteredforticlientArrayOutput{})
	pulumi.RegisterOutputType(RegisteredforticlientMapOutput{})
}
