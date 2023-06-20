// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SSH proxy local CA.
//
// ## Import
//
// # FirewallSsh LocalCa can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallsshLocalca:FirewallsshLocalca labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallsshLocalca:FirewallsshLocalca labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallsshLocalca struct {
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
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallsshLocalca registers a new resource with the given unique name, arguments, and options.
func NewFirewallsshLocalca(ctx *pulumi.Context,
	name string, args *FirewallsshLocalcaArgs, opts ...pulumi.ResourceOption) (*FirewallsshLocalca, error) {
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
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallsshLocalca
	err := ctx.RegisterResource("fortios:index/firewallsshLocalca:FirewallsshLocalca", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallsshLocalca gets an existing FirewallsshLocalca resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallsshLocalca(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallsshLocalcaState, opts ...pulumi.ResourceOption) (*FirewallsshLocalca, error) {
	var resource FirewallsshLocalca
	err := ctx.ReadResource("fortios:index/firewallsshLocalca:FirewallsshLocalca", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallsshLocalca resources.
type firewallsshLocalcaState struct {
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

type FirewallsshLocalcaState struct {
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

func (FirewallsshLocalcaState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallsshLocalcaState)(nil)).Elem()
}

type firewallsshLocalcaArgs struct {
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

// The set of arguments for constructing a FirewallsshLocalca resource.
type FirewallsshLocalcaArgs struct {
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

func (FirewallsshLocalcaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallsshLocalcaArgs)(nil)).Elem()
}

type FirewallsshLocalcaInput interface {
	pulumi.Input

	ToFirewallsshLocalcaOutput() FirewallsshLocalcaOutput
	ToFirewallsshLocalcaOutputWithContext(ctx context.Context) FirewallsshLocalcaOutput
}

func (*FirewallsshLocalca) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallsshLocalca)(nil)).Elem()
}

func (i *FirewallsshLocalca) ToFirewallsshLocalcaOutput() FirewallsshLocalcaOutput {
	return i.ToFirewallsshLocalcaOutputWithContext(context.Background())
}

func (i *FirewallsshLocalca) ToFirewallsshLocalcaOutputWithContext(ctx context.Context) FirewallsshLocalcaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallsshLocalcaOutput)
}

// FirewallsshLocalcaArrayInput is an input type that accepts FirewallsshLocalcaArray and FirewallsshLocalcaArrayOutput values.
// You can construct a concrete instance of `FirewallsshLocalcaArrayInput` via:
//
//	FirewallsshLocalcaArray{ FirewallsshLocalcaArgs{...} }
type FirewallsshLocalcaArrayInput interface {
	pulumi.Input

	ToFirewallsshLocalcaArrayOutput() FirewallsshLocalcaArrayOutput
	ToFirewallsshLocalcaArrayOutputWithContext(context.Context) FirewallsshLocalcaArrayOutput
}

type FirewallsshLocalcaArray []FirewallsshLocalcaInput

func (FirewallsshLocalcaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallsshLocalca)(nil)).Elem()
}

func (i FirewallsshLocalcaArray) ToFirewallsshLocalcaArrayOutput() FirewallsshLocalcaArrayOutput {
	return i.ToFirewallsshLocalcaArrayOutputWithContext(context.Background())
}

func (i FirewallsshLocalcaArray) ToFirewallsshLocalcaArrayOutputWithContext(ctx context.Context) FirewallsshLocalcaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallsshLocalcaArrayOutput)
}

// FirewallsshLocalcaMapInput is an input type that accepts FirewallsshLocalcaMap and FirewallsshLocalcaMapOutput values.
// You can construct a concrete instance of `FirewallsshLocalcaMapInput` via:
//
//	FirewallsshLocalcaMap{ "key": FirewallsshLocalcaArgs{...} }
type FirewallsshLocalcaMapInput interface {
	pulumi.Input

	ToFirewallsshLocalcaMapOutput() FirewallsshLocalcaMapOutput
	ToFirewallsshLocalcaMapOutputWithContext(context.Context) FirewallsshLocalcaMapOutput
}

type FirewallsshLocalcaMap map[string]FirewallsshLocalcaInput

func (FirewallsshLocalcaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallsshLocalca)(nil)).Elem()
}

func (i FirewallsshLocalcaMap) ToFirewallsshLocalcaMapOutput() FirewallsshLocalcaMapOutput {
	return i.ToFirewallsshLocalcaMapOutputWithContext(context.Background())
}

func (i FirewallsshLocalcaMap) ToFirewallsshLocalcaMapOutputWithContext(ctx context.Context) FirewallsshLocalcaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallsshLocalcaMapOutput)
}

type FirewallsshLocalcaOutput struct{ *pulumi.OutputState }

func (FirewallsshLocalcaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallsshLocalca)(nil)).Elem()
}

func (o FirewallsshLocalcaOutput) ToFirewallsshLocalcaOutput() FirewallsshLocalcaOutput {
	return o
}

func (o FirewallsshLocalcaOutput) ToFirewallsshLocalcaOutputWithContext(ctx context.Context) FirewallsshLocalcaOutput {
	return o
}

// SSH proxy local CA name.
func (o FirewallsshLocalcaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallsshLocalca) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Password for SSH private key.
func (o FirewallsshLocalcaOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallsshLocalca) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// SSH proxy private key, encrypted with a password.
func (o FirewallsshLocalcaOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallsshLocalca) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

// SSH proxy public key.
func (o FirewallsshLocalcaOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallsshLocalca) pulumi.StringOutput { return v.PublicKey }).(pulumi.StringOutput)
}

// SSH proxy local CA source type. Valid values: `built-in`, `user`.
func (o FirewallsshLocalcaOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallsshLocalca) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallsshLocalcaOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallsshLocalca) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallsshLocalcaArrayOutput struct{ *pulumi.OutputState }

func (FirewallsshLocalcaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallsshLocalca)(nil)).Elem()
}

func (o FirewallsshLocalcaArrayOutput) ToFirewallsshLocalcaArrayOutput() FirewallsshLocalcaArrayOutput {
	return o
}

func (o FirewallsshLocalcaArrayOutput) ToFirewallsshLocalcaArrayOutputWithContext(ctx context.Context) FirewallsshLocalcaArrayOutput {
	return o
}

func (o FirewallsshLocalcaArrayOutput) Index(i pulumi.IntInput) FirewallsshLocalcaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallsshLocalca {
		return vs[0].([]*FirewallsshLocalca)[vs[1].(int)]
	}).(FirewallsshLocalcaOutput)
}

type FirewallsshLocalcaMapOutput struct{ *pulumi.OutputState }

func (FirewallsshLocalcaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallsshLocalca)(nil)).Elem()
}

func (o FirewallsshLocalcaMapOutput) ToFirewallsshLocalcaMapOutput() FirewallsshLocalcaMapOutput {
	return o
}

func (o FirewallsshLocalcaMapOutput) ToFirewallsshLocalcaMapOutputWithContext(ctx context.Context) FirewallsshLocalcaMapOutput {
	return o
}

func (o FirewallsshLocalcaMapOutput) MapIndex(k pulumi.StringInput) FirewallsshLocalcaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallsshLocalca {
		return vs[0].(map[string]*FirewallsshLocalca)[vs[1].(string)]
	}).(FirewallsshLocalcaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallsshLocalcaInput)(nil)).Elem(), &FirewallsshLocalca{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallsshLocalcaArrayInput)(nil)).Elem(), FirewallsshLocalcaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallsshLocalcaMapInput)(nil)).Elem(), FirewallsshLocalcaMap{})
	pulumi.RegisterOutputType(FirewallsshLocalcaOutput{})
	pulumi.RegisterOutputType(FirewallsshLocalcaArrayOutput{})
	pulumi.RegisterOutputType(FirewallsshLocalcaMapOutput{})
}