// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure FortiSandbox.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/system"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := system.NewFortisandbox(ctx, "trname", &system.FortisandboxArgs{
//				EncAlgorithm:       pulumi.String("default"),
//				SslMinProtoVersion: pulumi.String("default"),
//				Status:             pulumi.String("disable"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// System Fortisandbox can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/fortisandbox:Fortisandbox labelname SystemFortisandbox
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/fortisandbox:Fortisandbox labelname SystemFortisandbox
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Fortisandbox struct {
	pulumi.CustomResourceState

	// Notifier email address.
	Email pulumi.StringOutput `pulumi:"email"`
	// Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
	EncAlgorithm pulumi.StringOutput `pulumi:"encAlgorithm"`
	// Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
	Forticloud pulumi.StringOutput `pulumi:"forticloud"`
	// Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
	InlineScan pulumi.StringOutput `pulumi:"inlineScan"`
	// Specify outgoing interface to reach server.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringOutput `pulumi:"interfaceSelectMethod"`
	// IPv4 or IPv6 address of the remote FortiSandbox.
	Server pulumi.StringOutput `pulumi:"server"`
	// Source IP address for communications to FortiSandbox.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion pulumi.StringOutput `pulumi:"sslMinProtoVersion"`
	// Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFortisandbox registers a new resource with the given unique name, arguments, and options.
func NewFortisandbox(ctx *pulumi.Context,
	name string, args *FortisandboxArgs, opts ...pulumi.ResourceOption) (*Fortisandbox, error) {
	if args == nil {
		args = &FortisandboxArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Fortisandbox
	err := ctx.RegisterResource("fortios:system/fortisandbox:Fortisandbox", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortisandbox gets an existing Fortisandbox resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortisandbox(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortisandboxState, opts ...pulumi.ResourceOption) (*Fortisandbox, error) {
	var resource Fortisandbox
	err := ctx.ReadResource("fortios:system/fortisandbox:Fortisandbox", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fortisandbox resources.
type fortisandboxState struct {
	// Notifier email address.
	Email *string `pulumi:"email"`
	// Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
	EncAlgorithm *string `pulumi:"encAlgorithm"`
	// Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
	Forticloud *string `pulumi:"forticloud"`
	// Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
	InlineScan *string `pulumi:"inlineScan"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// IPv4 or IPv6 address of the remote FortiSandbox.
	Server *string `pulumi:"server"`
	// Source IP address for communications to FortiSandbox.
	SourceIp *string `pulumi:"sourceIp"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FortisandboxState struct {
	// Notifier email address.
	Email pulumi.StringPtrInput
	// Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
	EncAlgorithm pulumi.StringPtrInput
	// Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
	Forticloud pulumi.StringPtrInput
	// Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
	InlineScan pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// IPv4 or IPv6 address of the remote FortiSandbox.
	Server pulumi.StringPtrInput
	// Source IP address for communications to FortiSandbox.
	SourceIp pulumi.StringPtrInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion pulumi.StringPtrInput
	// Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FortisandboxState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortisandboxState)(nil)).Elem()
}

type fortisandboxArgs struct {
	// Notifier email address.
	Email *string `pulumi:"email"`
	// Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
	EncAlgorithm *string `pulumi:"encAlgorithm"`
	// Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
	Forticloud *string `pulumi:"forticloud"`
	// Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
	InlineScan *string `pulumi:"inlineScan"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// IPv4 or IPv6 address of the remote FortiSandbox.
	Server *string `pulumi:"server"`
	// Source IP address for communications to FortiSandbox.
	SourceIp *string `pulumi:"sourceIp"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Fortisandbox resource.
type FortisandboxArgs struct {
	// Notifier email address.
	Email pulumi.StringPtrInput
	// Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
	EncAlgorithm pulumi.StringPtrInput
	// Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
	Forticloud pulumi.StringPtrInput
	// Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
	InlineScan pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// IPv4 or IPv6 address of the remote FortiSandbox.
	Server pulumi.StringPtrInput
	// Source IP address for communications to FortiSandbox.
	SourceIp pulumi.StringPtrInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion pulumi.StringPtrInput
	// Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FortisandboxArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortisandboxArgs)(nil)).Elem()
}

type FortisandboxInput interface {
	pulumi.Input

	ToFortisandboxOutput() FortisandboxOutput
	ToFortisandboxOutputWithContext(ctx context.Context) FortisandboxOutput
}

func (*Fortisandbox) ElementType() reflect.Type {
	return reflect.TypeOf((**Fortisandbox)(nil)).Elem()
}

func (i *Fortisandbox) ToFortisandboxOutput() FortisandboxOutput {
	return i.ToFortisandboxOutputWithContext(context.Background())
}

func (i *Fortisandbox) ToFortisandboxOutputWithContext(ctx context.Context) FortisandboxOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortisandboxOutput)
}

// FortisandboxArrayInput is an input type that accepts FortisandboxArray and FortisandboxArrayOutput values.
// You can construct a concrete instance of `FortisandboxArrayInput` via:
//
//	FortisandboxArray{ FortisandboxArgs{...} }
type FortisandboxArrayInput interface {
	pulumi.Input

	ToFortisandboxArrayOutput() FortisandboxArrayOutput
	ToFortisandboxArrayOutputWithContext(context.Context) FortisandboxArrayOutput
}

type FortisandboxArray []FortisandboxInput

func (FortisandboxArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fortisandbox)(nil)).Elem()
}

func (i FortisandboxArray) ToFortisandboxArrayOutput() FortisandboxArrayOutput {
	return i.ToFortisandboxArrayOutputWithContext(context.Background())
}

func (i FortisandboxArray) ToFortisandboxArrayOutputWithContext(ctx context.Context) FortisandboxArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortisandboxArrayOutput)
}

// FortisandboxMapInput is an input type that accepts FortisandboxMap and FortisandboxMapOutput values.
// You can construct a concrete instance of `FortisandboxMapInput` via:
//
//	FortisandboxMap{ "key": FortisandboxArgs{...} }
type FortisandboxMapInput interface {
	pulumi.Input

	ToFortisandboxMapOutput() FortisandboxMapOutput
	ToFortisandboxMapOutputWithContext(context.Context) FortisandboxMapOutput
}

type FortisandboxMap map[string]FortisandboxInput

func (FortisandboxMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fortisandbox)(nil)).Elem()
}

func (i FortisandboxMap) ToFortisandboxMapOutput() FortisandboxMapOutput {
	return i.ToFortisandboxMapOutputWithContext(context.Background())
}

func (i FortisandboxMap) ToFortisandboxMapOutputWithContext(ctx context.Context) FortisandboxMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortisandboxMapOutput)
}

type FortisandboxOutput struct{ *pulumi.OutputState }

func (FortisandboxOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fortisandbox)(nil)).Elem()
}

func (o FortisandboxOutput) ToFortisandboxOutput() FortisandboxOutput {
	return o
}

func (o FortisandboxOutput) ToFortisandboxOutputWithContext(ctx context.Context) FortisandboxOutput {
	return o
}

// Notifier email address.
func (o FortisandboxOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortisandbox) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
func (o FortisandboxOutput) EncAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortisandbox) pulumi.StringOutput { return v.EncAlgorithm }).(pulumi.StringOutput)
}

// Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
func (o FortisandboxOutput) Forticloud() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortisandbox) pulumi.StringOutput { return v.Forticloud }).(pulumi.StringOutput)
}

// Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
func (o FortisandboxOutput) InlineScan() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortisandbox) pulumi.StringOutput { return v.InlineScan }).(pulumi.StringOutput)
}

// Specify outgoing interface to reach server.
func (o FortisandboxOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortisandbox) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
func (o FortisandboxOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortisandbox) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

// IPv4 or IPv6 address of the remote FortiSandbox.
func (o FortisandboxOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortisandbox) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

// Source IP address for communications to FortiSandbox.
func (o FortisandboxOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortisandbox) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
func (o FortisandboxOutput) SslMinProtoVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortisandbox) pulumi.StringOutput { return v.SslMinProtoVersion }).(pulumi.StringOutput)
}

// Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
func (o FortisandboxOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortisandbox) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FortisandboxOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Fortisandbox) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FortisandboxArrayOutput struct{ *pulumi.OutputState }

func (FortisandboxArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fortisandbox)(nil)).Elem()
}

func (o FortisandboxArrayOutput) ToFortisandboxArrayOutput() FortisandboxArrayOutput {
	return o
}

func (o FortisandboxArrayOutput) ToFortisandboxArrayOutputWithContext(ctx context.Context) FortisandboxArrayOutput {
	return o
}

func (o FortisandboxArrayOutput) Index(i pulumi.IntInput) FortisandboxOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Fortisandbox {
		return vs[0].([]*Fortisandbox)[vs[1].(int)]
	}).(FortisandboxOutput)
}

type FortisandboxMapOutput struct{ *pulumi.OutputState }

func (FortisandboxMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fortisandbox)(nil)).Elem()
}

func (o FortisandboxMapOutput) ToFortisandboxMapOutput() FortisandboxMapOutput {
	return o
}

func (o FortisandboxMapOutput) ToFortisandboxMapOutputWithContext(ctx context.Context) FortisandboxMapOutput {
	return o
}

func (o FortisandboxMapOutput) MapIndex(k pulumi.StringInput) FortisandboxOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Fortisandbox {
		return vs[0].(map[string]*Fortisandbox)[vs[1].(string)]
	}).(FortisandboxOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortisandboxInput)(nil)).Elem(), &Fortisandbox{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortisandboxArrayInput)(nil)).Elem(), FortisandboxArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortisandboxMapInput)(nil)).Elem(), FortisandboxMap{})
	pulumi.RegisterOutputType(FortisandboxOutput{})
	pulumi.RegisterOutputType(FortisandboxArrayOutput{})
	pulumi.RegisterOutputType(FortisandboxMapOutput{})
}
