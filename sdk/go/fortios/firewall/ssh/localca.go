// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssh

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// SSH proxy local CA.
//
// ## Import
//
// FirewallSsh LocalCa can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/ssh/localca:Localca labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/ssh/localca:Localca labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Localca struct {
	pulumi.CustomResourceState

	// SSH proxy local CA name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password for SSH private key.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// SSH proxy private key, encrypted with a password.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// SSH proxy public key.
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// SSH proxy local CA source type. Valid values: `built-in`, `user`.
	Source pulumi.StringOutput `pulumi:"source"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewLocalca registers a new resource with the given unique name, arguments, and options.
func NewLocalca(ctx *pulumi.Context,
	name string, args *LocalcaArgs, opts ...pulumi.ResourceOption) (*Localca, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateKey == nil {
		return nil, errors.New("invalid value for required argument 'PrivateKey'")
	}
	if args.PublicKey == nil {
		return nil, errors.New("invalid value for required argument 'PublicKey'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	if args.PrivateKey != nil {
		args.PrivateKey = pulumi.ToSecret(args.PrivateKey).(pulumi.StringInput)
	}
	if args.PublicKey != nil {
		args.PublicKey = pulumi.ToSecret(args.PublicKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
		"privateKey",
		"publicKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Localca
	err := ctx.RegisterResource("fortios:firewall/ssh/localca:Localca", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalca gets an existing Localca resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalca(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalcaState, opts ...pulumi.ResourceOption) (*Localca, error) {
	var resource Localca
	err := ctx.ReadResource("fortios:firewall/ssh/localca:Localca", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Localca resources.
type localcaState struct {
	// SSH proxy local CA name.
	Name *string `pulumi:"name"`
	// Password for SSH private key.
	Password *string `pulumi:"password"`
	// SSH proxy private key, encrypted with a password.
	PrivateKey *string `pulumi:"privateKey"`
	// SSH proxy public key.
	PublicKey *string `pulumi:"publicKey"`
	// SSH proxy local CA source type. Valid values: `built-in`, `user`.
	Source *string `pulumi:"source"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type LocalcaState struct {
	// SSH proxy local CA name.
	Name pulumi.StringPtrInput
	// Password for SSH private key.
	Password pulumi.StringPtrInput
	// SSH proxy private key, encrypted with a password.
	PrivateKey pulumi.StringPtrInput
	// SSH proxy public key.
	PublicKey pulumi.StringPtrInput
	// SSH proxy local CA source type. Valid values: `built-in`, `user`.
	Source pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LocalcaState) ElementType() reflect.Type {
	return reflect.TypeOf((*localcaState)(nil)).Elem()
}

type localcaArgs struct {
	// SSH proxy local CA name.
	Name *string `pulumi:"name"`
	// Password for SSH private key.
	Password *string `pulumi:"password"`
	// SSH proxy private key, encrypted with a password.
	PrivateKey string `pulumi:"privateKey"`
	// SSH proxy public key.
	PublicKey string `pulumi:"publicKey"`
	// SSH proxy local CA source type. Valid values: `built-in`, `user`.
	Source *string `pulumi:"source"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Localca resource.
type LocalcaArgs struct {
	// SSH proxy local CA name.
	Name pulumi.StringPtrInput
	// Password for SSH private key.
	Password pulumi.StringPtrInput
	// SSH proxy private key, encrypted with a password.
	PrivateKey pulumi.StringInput
	// SSH proxy public key.
	PublicKey pulumi.StringInput
	// SSH proxy local CA source type. Valid values: `built-in`, `user`.
	Source pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LocalcaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localcaArgs)(nil)).Elem()
}

type LocalcaInput interface {
	pulumi.Input

	ToLocalcaOutput() LocalcaOutput
	ToLocalcaOutputWithContext(ctx context.Context) LocalcaOutput
}

func (*Localca) ElementType() reflect.Type {
	return reflect.TypeOf((**Localca)(nil)).Elem()
}

func (i *Localca) ToLocalcaOutput() LocalcaOutput {
	return i.ToLocalcaOutputWithContext(context.Background())
}

func (i *Localca) ToLocalcaOutputWithContext(ctx context.Context) LocalcaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalcaOutput)
}

// LocalcaArrayInput is an input type that accepts LocalcaArray and LocalcaArrayOutput values.
// You can construct a concrete instance of `LocalcaArrayInput` via:
//
//	LocalcaArray{ LocalcaArgs{...} }
type LocalcaArrayInput interface {
	pulumi.Input

	ToLocalcaArrayOutput() LocalcaArrayOutput
	ToLocalcaArrayOutputWithContext(context.Context) LocalcaArrayOutput
}

type LocalcaArray []LocalcaInput

func (LocalcaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Localca)(nil)).Elem()
}

func (i LocalcaArray) ToLocalcaArrayOutput() LocalcaArrayOutput {
	return i.ToLocalcaArrayOutputWithContext(context.Background())
}

func (i LocalcaArray) ToLocalcaArrayOutputWithContext(ctx context.Context) LocalcaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalcaArrayOutput)
}

// LocalcaMapInput is an input type that accepts LocalcaMap and LocalcaMapOutput values.
// You can construct a concrete instance of `LocalcaMapInput` via:
//
//	LocalcaMap{ "key": LocalcaArgs{...} }
type LocalcaMapInput interface {
	pulumi.Input

	ToLocalcaMapOutput() LocalcaMapOutput
	ToLocalcaMapOutputWithContext(context.Context) LocalcaMapOutput
}

type LocalcaMap map[string]LocalcaInput

func (LocalcaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Localca)(nil)).Elem()
}

func (i LocalcaMap) ToLocalcaMapOutput() LocalcaMapOutput {
	return i.ToLocalcaMapOutputWithContext(context.Background())
}

func (i LocalcaMap) ToLocalcaMapOutputWithContext(ctx context.Context) LocalcaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalcaMapOutput)
}

type LocalcaOutput struct{ *pulumi.OutputState }

func (LocalcaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Localca)(nil)).Elem()
}

func (o LocalcaOutput) ToLocalcaOutput() LocalcaOutput {
	return o
}

func (o LocalcaOutput) ToLocalcaOutputWithContext(ctx context.Context) LocalcaOutput {
	return o
}

// SSH proxy local CA name.
func (o LocalcaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Localca) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Password for SSH private key.
func (o LocalcaOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Localca) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// SSH proxy private key, encrypted with a password.
func (o LocalcaOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Localca) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

// SSH proxy public key.
func (o LocalcaOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Localca) pulumi.StringOutput { return v.PublicKey }).(pulumi.StringOutput)
}

// SSH proxy local CA source type. Valid values: `built-in`, `user`.
func (o LocalcaOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *Localca) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o LocalcaOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Localca) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type LocalcaArrayOutput struct{ *pulumi.OutputState }

func (LocalcaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Localca)(nil)).Elem()
}

func (o LocalcaArrayOutput) ToLocalcaArrayOutput() LocalcaArrayOutput {
	return o
}

func (o LocalcaArrayOutput) ToLocalcaArrayOutputWithContext(ctx context.Context) LocalcaArrayOutput {
	return o
}

func (o LocalcaArrayOutput) Index(i pulumi.IntInput) LocalcaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Localca {
		return vs[0].([]*Localca)[vs[1].(int)]
	}).(LocalcaOutput)
}

type LocalcaMapOutput struct{ *pulumi.OutputState }

func (LocalcaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Localca)(nil)).Elem()
}

func (o LocalcaMapOutput) ToLocalcaMapOutput() LocalcaMapOutput {
	return o
}

func (o LocalcaMapOutput) ToLocalcaMapOutputWithContext(ctx context.Context) LocalcaMapOutput {
	return o
}

func (o LocalcaMapOutput) MapIndex(k pulumi.StringInput) LocalcaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Localca {
		return vs[0].(map[string]*Localca)[vs[1].(string)]
	}).(LocalcaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalcaInput)(nil)).Elem(), &Localca{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalcaArrayInput)(nil)).Elem(), LocalcaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalcaMapInput)(nil)).Elem(), LocalcaMap{})
	pulumi.RegisterOutputType(LocalcaOutput{})
	pulumi.RegisterOutputType(LocalcaArrayOutput{})
	pulumi.RegisterOutputType(LocalcaMapOutput{})
}
