// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package user

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure TACACS+ server entries.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/user"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := user.NewTacacs(ctx, "trname", &user.TacacsArgs{
//				AuthenType:    pulumi.String("auto"),
//				Authorization: pulumi.String("disable"),
//				Port:          pulumi.Int(2342),
//				Server:        pulumi.String("1.1.1.1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # User Tacacs can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:user/tacacs:Tacacs labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:user/tacacs:Tacacs labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Tacacs struct {
	pulumi.CustomResourceState

	// Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
	AuthenType pulumi.StringOutput `pulumi:"authenType"`
	// Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
	Authorization pulumi.StringOutput `pulumi:"authorization"`
	// Specify outgoing interface to reach server.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringOutput `pulumi:"interfaceSelectMethod"`
	// Key to access the primary server.
	Key pulumi.StringPtrOutput `pulumi:"key"`
	// TACACS+ server entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Port number of the TACACS+ server.
	Port pulumi.IntOutput `pulumi:"port"`
	// Key to access the secondary server.
	SecondaryKey pulumi.StringPtrOutput `pulumi:"secondaryKey"`
	// Secondary TACACS+ server CN domain name or IP address.
	SecondaryServer pulumi.StringOutput `pulumi:"secondaryServer"`
	// Primary TACACS+ server CN domain name or IP address.
	Server pulumi.StringOutput `pulumi:"server"`
	// source IP for communications to TACACS+ server.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Key to access the tertiary server.
	TertiaryKey pulumi.StringPtrOutput `pulumi:"tertiaryKey"`
	// Tertiary TACACS+ server CN domain name or IP address.
	TertiaryServer pulumi.StringOutput `pulumi:"tertiaryServer"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewTacacs registers a new resource with the given unique name, arguments, and options.
func NewTacacs(ctx *pulumi.Context,
	name string, args *TacacsArgs, opts ...pulumi.ResourceOption) (*Tacacs, error) {
	if args == nil {
		args = &TacacsArgs{}
	}

	if args.Key != nil {
		args.Key = pulumi.ToSecret(args.Key).(pulumi.StringPtrInput)
	}
	if args.SecondaryKey != nil {
		args.SecondaryKey = pulumi.ToSecret(args.SecondaryKey).(pulumi.StringPtrInput)
	}
	if args.TertiaryKey != nil {
		args.TertiaryKey = pulumi.ToSecret(args.TertiaryKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"key",
		"secondaryKey",
		"tertiaryKey",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource Tacacs
	err := ctx.RegisterResource("fortios:user/tacacs:Tacacs", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTacacs gets an existing Tacacs resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTacacs(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TacacsState, opts ...pulumi.ResourceOption) (*Tacacs, error) {
	var resource Tacacs
	err := ctx.ReadResource("fortios:user/tacacs:Tacacs", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tacacs resources.
type tacacsState struct {
	// Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
	AuthenType *string `pulumi:"authenType"`
	// Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
	Authorization *string `pulumi:"authorization"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Key to access the primary server.
	Key *string `pulumi:"key"`
	// TACACS+ server entry name.
	Name *string `pulumi:"name"`
	// Port number of the TACACS+ server.
	Port *int `pulumi:"port"`
	// Key to access the secondary server.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// Secondary TACACS+ server CN domain name or IP address.
	SecondaryServer *string `pulumi:"secondaryServer"`
	// Primary TACACS+ server CN domain name or IP address.
	Server *string `pulumi:"server"`
	// source IP for communications to TACACS+ server.
	SourceIp *string `pulumi:"sourceIp"`
	// Key to access the tertiary server.
	TertiaryKey *string `pulumi:"tertiaryKey"`
	// Tertiary TACACS+ server CN domain name or IP address.
	TertiaryServer *string `pulumi:"tertiaryServer"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type TacacsState struct {
	// Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
	AuthenType pulumi.StringPtrInput
	// Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
	Authorization pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Key to access the primary server.
	Key pulumi.StringPtrInput
	// TACACS+ server entry name.
	Name pulumi.StringPtrInput
	// Port number of the TACACS+ server.
	Port pulumi.IntPtrInput
	// Key to access the secondary server.
	SecondaryKey pulumi.StringPtrInput
	// Secondary TACACS+ server CN domain name or IP address.
	SecondaryServer pulumi.StringPtrInput
	// Primary TACACS+ server CN domain name or IP address.
	Server pulumi.StringPtrInput
	// source IP for communications to TACACS+ server.
	SourceIp pulumi.StringPtrInput
	// Key to access the tertiary server.
	TertiaryKey pulumi.StringPtrInput
	// Tertiary TACACS+ server CN domain name or IP address.
	TertiaryServer pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (TacacsState) ElementType() reflect.Type {
	return reflect.TypeOf((*tacacsState)(nil)).Elem()
}

type tacacsArgs struct {
	// Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
	AuthenType *string `pulumi:"authenType"`
	// Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
	Authorization *string `pulumi:"authorization"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Key to access the primary server.
	Key *string `pulumi:"key"`
	// TACACS+ server entry name.
	Name *string `pulumi:"name"`
	// Port number of the TACACS+ server.
	Port *int `pulumi:"port"`
	// Key to access the secondary server.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// Secondary TACACS+ server CN domain name or IP address.
	SecondaryServer *string `pulumi:"secondaryServer"`
	// Primary TACACS+ server CN domain name or IP address.
	Server *string `pulumi:"server"`
	// source IP for communications to TACACS+ server.
	SourceIp *string `pulumi:"sourceIp"`
	// Key to access the tertiary server.
	TertiaryKey *string `pulumi:"tertiaryKey"`
	// Tertiary TACACS+ server CN domain name or IP address.
	TertiaryServer *string `pulumi:"tertiaryServer"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Tacacs resource.
type TacacsArgs struct {
	// Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
	AuthenType pulumi.StringPtrInput
	// Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
	Authorization pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Key to access the primary server.
	Key pulumi.StringPtrInput
	// TACACS+ server entry name.
	Name pulumi.StringPtrInput
	// Port number of the TACACS+ server.
	Port pulumi.IntPtrInput
	// Key to access the secondary server.
	SecondaryKey pulumi.StringPtrInput
	// Secondary TACACS+ server CN domain name or IP address.
	SecondaryServer pulumi.StringPtrInput
	// Primary TACACS+ server CN domain name or IP address.
	Server pulumi.StringPtrInput
	// source IP for communications to TACACS+ server.
	SourceIp pulumi.StringPtrInput
	// Key to access the tertiary server.
	TertiaryKey pulumi.StringPtrInput
	// Tertiary TACACS+ server CN domain name or IP address.
	TertiaryServer pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (TacacsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tacacsArgs)(nil)).Elem()
}

type TacacsInput interface {
	pulumi.Input

	ToTacacsOutput() TacacsOutput
	ToTacacsOutputWithContext(ctx context.Context) TacacsOutput
}

func (*Tacacs) ElementType() reflect.Type {
	return reflect.TypeOf((**Tacacs)(nil)).Elem()
}

func (i *Tacacs) ToTacacsOutput() TacacsOutput {
	return i.ToTacacsOutputWithContext(context.Background())
}

func (i *Tacacs) ToTacacsOutputWithContext(ctx context.Context) TacacsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TacacsOutput)
}

// TacacsArrayInput is an input type that accepts TacacsArray and TacacsArrayOutput values.
// You can construct a concrete instance of `TacacsArrayInput` via:
//
//	TacacsArray{ TacacsArgs{...} }
type TacacsArrayInput interface {
	pulumi.Input

	ToTacacsArrayOutput() TacacsArrayOutput
	ToTacacsArrayOutputWithContext(context.Context) TacacsArrayOutput
}

type TacacsArray []TacacsInput

func (TacacsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Tacacs)(nil)).Elem()
}

func (i TacacsArray) ToTacacsArrayOutput() TacacsArrayOutput {
	return i.ToTacacsArrayOutputWithContext(context.Background())
}

func (i TacacsArray) ToTacacsArrayOutputWithContext(ctx context.Context) TacacsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TacacsArrayOutput)
}

// TacacsMapInput is an input type that accepts TacacsMap and TacacsMapOutput values.
// You can construct a concrete instance of `TacacsMapInput` via:
//
//	TacacsMap{ "key": TacacsArgs{...} }
type TacacsMapInput interface {
	pulumi.Input

	ToTacacsMapOutput() TacacsMapOutput
	ToTacacsMapOutputWithContext(context.Context) TacacsMapOutput
}

type TacacsMap map[string]TacacsInput

func (TacacsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Tacacs)(nil)).Elem()
}

func (i TacacsMap) ToTacacsMapOutput() TacacsMapOutput {
	return i.ToTacacsMapOutputWithContext(context.Background())
}

func (i TacacsMap) ToTacacsMapOutputWithContext(ctx context.Context) TacacsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TacacsMapOutput)
}

type TacacsOutput struct{ *pulumi.OutputState }

func (TacacsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Tacacs)(nil)).Elem()
}

func (o TacacsOutput) ToTacacsOutput() TacacsOutput {
	return o
}

func (o TacacsOutput) ToTacacsOutputWithContext(ctx context.Context) TacacsOutput {
	return o
}

// Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
func (o TacacsOutput) AuthenType() pulumi.StringOutput {
	return o.ApplyT(func(v *Tacacs) pulumi.StringOutput { return v.AuthenType }).(pulumi.StringOutput)
}

// Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
func (o TacacsOutput) Authorization() pulumi.StringOutput {
	return o.ApplyT(func(v *Tacacs) pulumi.StringOutput { return v.Authorization }).(pulumi.StringOutput)
}

// Specify outgoing interface to reach server.
func (o TacacsOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Tacacs) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
func (o TacacsOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *Tacacs) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

// Key to access the primary server.
func (o TacacsOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Tacacs) pulumi.StringPtrOutput { return v.Key }).(pulumi.StringPtrOutput)
}

// TACACS+ server entry name.
func (o TacacsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Tacacs) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Port number of the TACACS+ server.
func (o TacacsOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Tacacs) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Key to access the secondary server.
func (o TacacsOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Tacacs) pulumi.StringPtrOutput { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

// Secondary TACACS+ server CN domain name or IP address.
func (o TacacsOutput) SecondaryServer() pulumi.StringOutput {
	return o.ApplyT(func(v *Tacacs) pulumi.StringOutput { return v.SecondaryServer }).(pulumi.StringOutput)
}

// Primary TACACS+ server CN domain name or IP address.
func (o TacacsOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *Tacacs) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

// source IP for communications to TACACS+ server.
func (o TacacsOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Tacacs) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// Key to access the tertiary server.
func (o TacacsOutput) TertiaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Tacacs) pulumi.StringPtrOutput { return v.TertiaryKey }).(pulumi.StringPtrOutput)
}

// Tertiary TACACS+ server CN domain name or IP address.
func (o TacacsOutput) TertiaryServer() pulumi.StringOutput {
	return o.ApplyT(func(v *Tacacs) pulumi.StringOutput { return v.TertiaryServer }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o TacacsOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Tacacs) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type TacacsArrayOutput struct{ *pulumi.OutputState }

func (TacacsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Tacacs)(nil)).Elem()
}

func (o TacacsArrayOutput) ToTacacsArrayOutput() TacacsArrayOutput {
	return o
}

func (o TacacsArrayOutput) ToTacacsArrayOutputWithContext(ctx context.Context) TacacsArrayOutput {
	return o
}

func (o TacacsArrayOutput) Index(i pulumi.IntInput) TacacsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Tacacs {
		return vs[0].([]*Tacacs)[vs[1].(int)]
	}).(TacacsOutput)
}

type TacacsMapOutput struct{ *pulumi.OutputState }

func (TacacsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Tacacs)(nil)).Elem()
}

func (o TacacsMapOutput) ToTacacsMapOutput() TacacsMapOutput {
	return o
}

func (o TacacsMapOutput) ToTacacsMapOutputWithContext(ctx context.Context) TacacsMapOutput {
	return o
}

func (o TacacsMapOutput) MapIndex(k pulumi.StringInput) TacacsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Tacacs {
		return vs[0].(map[string]*Tacacs)[vs[1].(string)]
	}).(TacacsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TacacsInput)(nil)).Elem(), &Tacacs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TacacsArrayInput)(nil)).Elem(), TacacsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TacacsMapInput)(nil)).Elem(), TacacsMap{})
	pulumi.RegisterOutputType(TacacsOutput{})
	pulumi.RegisterOutputType(TacacsArrayOutput{})
	pulumi.RegisterOutputType(TacacsMapOutput{})
}
