// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package user

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// POP3 server entry configuration.
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
//			_, err := user.NewPop3(ctx, "trname", &user.Pop3Args{
//				Port:               pulumi.Int(0),
//				Secure:             pulumi.String("pop3s"),
//				Server:             pulumi.String("1.1.1.1"),
//				SslMinProtoVersion: pulumi.String("default"),
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
// User Pop3 can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:user/pop3:Pop3 labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:user/pop3:Pop3 labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Pop3 struct {
	pulumi.CustomResourceState

	// POP3 server entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// POP3 service port number.
	Port pulumi.IntOutput `pulumi:"port"`
	// SSL connection. Valid values: `none`, `starttls`, `pop3s`.
	Secure pulumi.StringOutput `pulumi:"secure"`
	// {<name_str|ip_str>} server domain name or IP.
	Server pulumi.StringOutput `pulumi:"server"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion pulumi.StringOutput `pulumi:"sslMinProtoVersion"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewPop3 registers a new resource with the given unique name, arguments, and options.
func NewPop3(ctx *pulumi.Context,
	name string, args *Pop3Args, opts ...pulumi.ResourceOption) (*Pop3, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Server == nil {
		return nil, errors.New("invalid value for required argument 'Server'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Pop3
	err := ctx.RegisterResource("fortios:user/pop3:Pop3", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPop3 gets an existing Pop3 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPop3(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Pop3State, opts ...pulumi.ResourceOption) (*Pop3, error) {
	var resource Pop3
	err := ctx.ReadResource("fortios:user/pop3:Pop3", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pop3 resources.
type pop3State struct {
	// POP3 server entry name.
	Name *string `pulumi:"name"`
	// POP3 service port number.
	Port *int `pulumi:"port"`
	// SSL connection. Valid values: `none`, `starttls`, `pop3s`.
	Secure *string `pulumi:"secure"`
	// {<name_str|ip_str>} server domain name or IP.
	Server *string `pulumi:"server"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Pop3State struct {
	// POP3 server entry name.
	Name pulumi.StringPtrInput
	// POP3 service port number.
	Port pulumi.IntPtrInput
	// SSL connection. Valid values: `none`, `starttls`, `pop3s`.
	Secure pulumi.StringPtrInput
	// {<name_str|ip_str>} server domain name or IP.
	Server pulumi.StringPtrInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Pop3State) ElementType() reflect.Type {
	return reflect.TypeOf((*pop3State)(nil)).Elem()
}

type pop3Args struct {
	// POP3 server entry name.
	Name *string `pulumi:"name"`
	// POP3 service port number.
	Port *int `pulumi:"port"`
	// SSL connection. Valid values: `none`, `starttls`, `pop3s`.
	Secure *string `pulumi:"secure"`
	// {<name_str|ip_str>} server domain name or IP.
	Server string `pulumi:"server"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Pop3 resource.
type Pop3Args struct {
	// POP3 server entry name.
	Name pulumi.StringPtrInput
	// POP3 service port number.
	Port pulumi.IntPtrInput
	// SSL connection. Valid values: `none`, `starttls`, `pop3s`.
	Secure pulumi.StringPtrInput
	// {<name_str|ip_str>} server domain name or IP.
	Server pulumi.StringInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Pop3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*pop3Args)(nil)).Elem()
}

type Pop3Input interface {
	pulumi.Input

	ToPop3Output() Pop3Output
	ToPop3OutputWithContext(ctx context.Context) Pop3Output
}

func (*Pop3) ElementType() reflect.Type {
	return reflect.TypeOf((**Pop3)(nil)).Elem()
}

func (i *Pop3) ToPop3Output() Pop3Output {
	return i.ToPop3OutputWithContext(context.Background())
}

func (i *Pop3) ToPop3OutputWithContext(ctx context.Context) Pop3Output {
	return pulumi.ToOutputWithContext(ctx, i).(Pop3Output)
}

// Pop3ArrayInput is an input type that accepts Pop3Array and Pop3ArrayOutput values.
// You can construct a concrete instance of `Pop3ArrayInput` via:
//
//	Pop3Array{ Pop3Args{...} }
type Pop3ArrayInput interface {
	pulumi.Input

	ToPop3ArrayOutput() Pop3ArrayOutput
	ToPop3ArrayOutputWithContext(context.Context) Pop3ArrayOutput
}

type Pop3Array []Pop3Input

func (Pop3Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pop3)(nil)).Elem()
}

func (i Pop3Array) ToPop3ArrayOutput() Pop3ArrayOutput {
	return i.ToPop3ArrayOutputWithContext(context.Background())
}

func (i Pop3Array) ToPop3ArrayOutputWithContext(ctx context.Context) Pop3ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Pop3ArrayOutput)
}

// Pop3MapInput is an input type that accepts Pop3Map and Pop3MapOutput values.
// You can construct a concrete instance of `Pop3MapInput` via:
//
//	Pop3Map{ "key": Pop3Args{...} }
type Pop3MapInput interface {
	pulumi.Input

	ToPop3MapOutput() Pop3MapOutput
	ToPop3MapOutputWithContext(context.Context) Pop3MapOutput
}

type Pop3Map map[string]Pop3Input

func (Pop3Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pop3)(nil)).Elem()
}

func (i Pop3Map) ToPop3MapOutput() Pop3MapOutput {
	return i.ToPop3MapOutputWithContext(context.Background())
}

func (i Pop3Map) ToPop3MapOutputWithContext(ctx context.Context) Pop3MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Pop3MapOutput)
}

type Pop3Output struct{ *pulumi.OutputState }

func (Pop3Output) ElementType() reflect.Type {
	return reflect.TypeOf((**Pop3)(nil)).Elem()
}

func (o Pop3Output) ToPop3Output() Pop3Output {
	return o
}

func (o Pop3Output) ToPop3OutputWithContext(ctx context.Context) Pop3Output {
	return o
}

// POP3 server entry name.
func (o Pop3Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Pop3) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// POP3 service port number.
func (o Pop3Output) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Pop3) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// SSL connection. Valid values: `none`, `starttls`, `pop3s`.
func (o Pop3Output) Secure() pulumi.StringOutput {
	return o.ApplyT(func(v *Pop3) pulumi.StringOutput { return v.Secure }).(pulumi.StringOutput)
}

// {<name_str|ip_str>} server domain name or IP.
func (o Pop3Output) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *Pop3) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
func (o Pop3Output) SslMinProtoVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Pop3) pulumi.StringOutput { return v.SslMinProtoVersion }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Pop3Output) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Pop3) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type Pop3ArrayOutput struct{ *pulumi.OutputState }

func (Pop3ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pop3)(nil)).Elem()
}

func (o Pop3ArrayOutput) ToPop3ArrayOutput() Pop3ArrayOutput {
	return o
}

func (o Pop3ArrayOutput) ToPop3ArrayOutputWithContext(ctx context.Context) Pop3ArrayOutput {
	return o
}

func (o Pop3ArrayOutput) Index(i pulumi.IntInput) Pop3Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Pop3 {
		return vs[0].([]*Pop3)[vs[1].(int)]
	}).(Pop3Output)
}

type Pop3MapOutput struct{ *pulumi.OutputState }

func (Pop3MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pop3)(nil)).Elem()
}

func (o Pop3MapOutput) ToPop3MapOutput() Pop3MapOutput {
	return o
}

func (o Pop3MapOutput) ToPop3MapOutputWithContext(ctx context.Context) Pop3MapOutput {
	return o
}

func (o Pop3MapOutput) MapIndex(k pulumi.StringInput) Pop3Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Pop3 {
		return vs[0].(map[string]*Pop3)[vs[1].(string)]
	}).(Pop3Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Pop3Input)(nil)).Elem(), &Pop3{})
	pulumi.RegisterInputType(reflect.TypeOf((*Pop3ArrayInput)(nil)).Elem(), Pop3Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*Pop3MapInput)(nil)).Elem(), Pop3Map{})
	pulumi.RegisterOutputType(Pop3Output{})
	pulumi.RegisterOutputType(Pop3ArrayOutput{})
	pulumi.RegisterOutputType(Pop3MapOutput{})
}
