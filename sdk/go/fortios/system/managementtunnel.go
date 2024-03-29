// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Management tunnel configuration.
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
//			_, err := system.NewManagementtunnel(ctx, "trname", &system.ManagementtunnelArgs{
//				AllowCollectStatistics: pulumi.String("enable"),
//				AllowConfigRestore:     pulumi.String("enable"),
//				AllowPushConfiguration: pulumi.String("enable"),
//				AllowPushFirmware:      pulumi.String("enable"),
//				AuthorizedManagerOnly:  pulumi.String("enable"),
//				Status:                 pulumi.String("enable"),
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
// System ManagementTunnel can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/managementtunnel:Managementtunnel labelname SystemManagementTunnel
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/managementtunnel:Managementtunnel labelname SystemManagementTunnel
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Managementtunnel struct {
	pulumi.CustomResourceState

	// Enable/disable collection of run time statistics. Valid values: `enable`, `disable`.
	AllowCollectStatistics pulumi.StringOutput `pulumi:"allowCollectStatistics"`
	// Enable/disable allow config restore. Valid values: `enable`, `disable`.
	AllowConfigRestore pulumi.StringOutput `pulumi:"allowConfigRestore"`
	// Enable/disable push configuration. Valid values: `enable`, `disable`.
	AllowPushConfiguration pulumi.StringOutput `pulumi:"allowPushConfiguration"`
	// Enable/disable push firmware. Valid values: `enable`, `disable`.
	AllowPushFirmware pulumi.StringOutput `pulumi:"allowPushFirmware"`
	// Enable/disable restriction of authorized manager only. Valid values: `enable`, `disable`.
	AuthorizedManagerOnly pulumi.StringOutput `pulumi:"authorizedManagerOnly"`
	// Serial number.
	SerialNumber pulumi.StringOutput `pulumi:"serialNumber"`
	// Enable/disable FGFM tunnel. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewManagementtunnel registers a new resource with the given unique name, arguments, and options.
func NewManagementtunnel(ctx *pulumi.Context,
	name string, args *ManagementtunnelArgs, opts ...pulumi.ResourceOption) (*Managementtunnel, error) {
	if args == nil {
		args = &ManagementtunnelArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Managementtunnel
	err := ctx.RegisterResource("fortios:system/managementtunnel:Managementtunnel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagementtunnel gets an existing Managementtunnel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagementtunnel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementtunnelState, opts ...pulumi.ResourceOption) (*Managementtunnel, error) {
	var resource Managementtunnel
	err := ctx.ReadResource("fortios:system/managementtunnel:Managementtunnel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Managementtunnel resources.
type managementtunnelState struct {
	// Enable/disable collection of run time statistics. Valid values: `enable`, `disable`.
	AllowCollectStatistics *string `pulumi:"allowCollectStatistics"`
	// Enable/disable allow config restore. Valid values: `enable`, `disable`.
	AllowConfigRestore *string `pulumi:"allowConfigRestore"`
	// Enable/disable push configuration. Valid values: `enable`, `disable`.
	AllowPushConfiguration *string `pulumi:"allowPushConfiguration"`
	// Enable/disable push firmware. Valid values: `enable`, `disable`.
	AllowPushFirmware *string `pulumi:"allowPushFirmware"`
	// Enable/disable restriction of authorized manager only. Valid values: `enable`, `disable`.
	AuthorizedManagerOnly *string `pulumi:"authorizedManagerOnly"`
	// Serial number.
	SerialNumber *string `pulumi:"serialNumber"`
	// Enable/disable FGFM tunnel. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ManagementtunnelState struct {
	// Enable/disable collection of run time statistics. Valid values: `enable`, `disable`.
	AllowCollectStatistics pulumi.StringPtrInput
	// Enable/disable allow config restore. Valid values: `enable`, `disable`.
	AllowConfigRestore pulumi.StringPtrInput
	// Enable/disable push configuration. Valid values: `enable`, `disable`.
	AllowPushConfiguration pulumi.StringPtrInput
	// Enable/disable push firmware. Valid values: `enable`, `disable`.
	AllowPushFirmware pulumi.StringPtrInput
	// Enable/disable restriction of authorized manager only. Valid values: `enable`, `disable`.
	AuthorizedManagerOnly pulumi.StringPtrInput
	// Serial number.
	SerialNumber pulumi.StringPtrInput
	// Enable/disable FGFM tunnel. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ManagementtunnelState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementtunnelState)(nil)).Elem()
}

type managementtunnelArgs struct {
	// Enable/disable collection of run time statistics. Valid values: `enable`, `disable`.
	AllowCollectStatistics *string `pulumi:"allowCollectStatistics"`
	// Enable/disable allow config restore. Valid values: `enable`, `disable`.
	AllowConfigRestore *string `pulumi:"allowConfigRestore"`
	// Enable/disable push configuration. Valid values: `enable`, `disable`.
	AllowPushConfiguration *string `pulumi:"allowPushConfiguration"`
	// Enable/disable push firmware. Valid values: `enable`, `disable`.
	AllowPushFirmware *string `pulumi:"allowPushFirmware"`
	// Enable/disable restriction of authorized manager only. Valid values: `enable`, `disable`.
	AuthorizedManagerOnly *string `pulumi:"authorizedManagerOnly"`
	// Serial number.
	SerialNumber *string `pulumi:"serialNumber"`
	// Enable/disable FGFM tunnel. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Managementtunnel resource.
type ManagementtunnelArgs struct {
	// Enable/disable collection of run time statistics. Valid values: `enable`, `disable`.
	AllowCollectStatistics pulumi.StringPtrInput
	// Enable/disable allow config restore. Valid values: `enable`, `disable`.
	AllowConfigRestore pulumi.StringPtrInput
	// Enable/disable push configuration. Valid values: `enable`, `disable`.
	AllowPushConfiguration pulumi.StringPtrInput
	// Enable/disable push firmware. Valid values: `enable`, `disable`.
	AllowPushFirmware pulumi.StringPtrInput
	// Enable/disable restriction of authorized manager only. Valid values: `enable`, `disable`.
	AuthorizedManagerOnly pulumi.StringPtrInput
	// Serial number.
	SerialNumber pulumi.StringPtrInput
	// Enable/disable FGFM tunnel. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ManagementtunnelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementtunnelArgs)(nil)).Elem()
}

type ManagementtunnelInput interface {
	pulumi.Input

	ToManagementtunnelOutput() ManagementtunnelOutput
	ToManagementtunnelOutputWithContext(ctx context.Context) ManagementtunnelOutput
}

func (*Managementtunnel) ElementType() reflect.Type {
	return reflect.TypeOf((**Managementtunnel)(nil)).Elem()
}

func (i *Managementtunnel) ToManagementtunnelOutput() ManagementtunnelOutput {
	return i.ToManagementtunnelOutputWithContext(context.Background())
}

func (i *Managementtunnel) ToManagementtunnelOutputWithContext(ctx context.Context) ManagementtunnelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementtunnelOutput)
}

// ManagementtunnelArrayInput is an input type that accepts ManagementtunnelArray and ManagementtunnelArrayOutput values.
// You can construct a concrete instance of `ManagementtunnelArrayInput` via:
//
//	ManagementtunnelArray{ ManagementtunnelArgs{...} }
type ManagementtunnelArrayInput interface {
	pulumi.Input

	ToManagementtunnelArrayOutput() ManagementtunnelArrayOutput
	ToManagementtunnelArrayOutputWithContext(context.Context) ManagementtunnelArrayOutput
}

type ManagementtunnelArray []ManagementtunnelInput

func (ManagementtunnelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Managementtunnel)(nil)).Elem()
}

func (i ManagementtunnelArray) ToManagementtunnelArrayOutput() ManagementtunnelArrayOutput {
	return i.ToManagementtunnelArrayOutputWithContext(context.Background())
}

func (i ManagementtunnelArray) ToManagementtunnelArrayOutputWithContext(ctx context.Context) ManagementtunnelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementtunnelArrayOutput)
}

// ManagementtunnelMapInput is an input type that accepts ManagementtunnelMap and ManagementtunnelMapOutput values.
// You can construct a concrete instance of `ManagementtunnelMapInput` via:
//
//	ManagementtunnelMap{ "key": ManagementtunnelArgs{...} }
type ManagementtunnelMapInput interface {
	pulumi.Input

	ToManagementtunnelMapOutput() ManagementtunnelMapOutput
	ToManagementtunnelMapOutputWithContext(context.Context) ManagementtunnelMapOutput
}

type ManagementtunnelMap map[string]ManagementtunnelInput

func (ManagementtunnelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Managementtunnel)(nil)).Elem()
}

func (i ManagementtunnelMap) ToManagementtunnelMapOutput() ManagementtunnelMapOutput {
	return i.ToManagementtunnelMapOutputWithContext(context.Background())
}

func (i ManagementtunnelMap) ToManagementtunnelMapOutputWithContext(ctx context.Context) ManagementtunnelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementtunnelMapOutput)
}

type ManagementtunnelOutput struct{ *pulumi.OutputState }

func (ManagementtunnelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Managementtunnel)(nil)).Elem()
}

func (o ManagementtunnelOutput) ToManagementtunnelOutput() ManagementtunnelOutput {
	return o
}

func (o ManagementtunnelOutput) ToManagementtunnelOutputWithContext(ctx context.Context) ManagementtunnelOutput {
	return o
}

// Enable/disable collection of run time statistics. Valid values: `enable`, `disable`.
func (o ManagementtunnelOutput) AllowCollectStatistics() pulumi.StringOutput {
	return o.ApplyT(func(v *Managementtunnel) pulumi.StringOutput { return v.AllowCollectStatistics }).(pulumi.StringOutput)
}

// Enable/disable allow config restore. Valid values: `enable`, `disable`.
func (o ManagementtunnelOutput) AllowConfigRestore() pulumi.StringOutput {
	return o.ApplyT(func(v *Managementtunnel) pulumi.StringOutput { return v.AllowConfigRestore }).(pulumi.StringOutput)
}

// Enable/disable push configuration. Valid values: `enable`, `disable`.
func (o ManagementtunnelOutput) AllowPushConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v *Managementtunnel) pulumi.StringOutput { return v.AllowPushConfiguration }).(pulumi.StringOutput)
}

// Enable/disable push firmware. Valid values: `enable`, `disable`.
func (o ManagementtunnelOutput) AllowPushFirmware() pulumi.StringOutput {
	return o.ApplyT(func(v *Managementtunnel) pulumi.StringOutput { return v.AllowPushFirmware }).(pulumi.StringOutput)
}

// Enable/disable restriction of authorized manager only. Valid values: `enable`, `disable`.
func (o ManagementtunnelOutput) AuthorizedManagerOnly() pulumi.StringOutput {
	return o.ApplyT(func(v *Managementtunnel) pulumi.StringOutput { return v.AuthorizedManagerOnly }).(pulumi.StringOutput)
}

// Serial number.
func (o ManagementtunnelOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *Managementtunnel) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

// Enable/disable FGFM tunnel. Valid values: `enable`, `disable`.
func (o ManagementtunnelOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Managementtunnel) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ManagementtunnelOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Managementtunnel) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ManagementtunnelArrayOutput struct{ *pulumi.OutputState }

func (ManagementtunnelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Managementtunnel)(nil)).Elem()
}

func (o ManagementtunnelArrayOutput) ToManagementtunnelArrayOutput() ManagementtunnelArrayOutput {
	return o
}

func (o ManagementtunnelArrayOutput) ToManagementtunnelArrayOutputWithContext(ctx context.Context) ManagementtunnelArrayOutput {
	return o
}

func (o ManagementtunnelArrayOutput) Index(i pulumi.IntInput) ManagementtunnelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Managementtunnel {
		return vs[0].([]*Managementtunnel)[vs[1].(int)]
	}).(ManagementtunnelOutput)
}

type ManagementtunnelMapOutput struct{ *pulumi.OutputState }

func (ManagementtunnelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Managementtunnel)(nil)).Elem()
}

func (o ManagementtunnelMapOutput) ToManagementtunnelMapOutput() ManagementtunnelMapOutput {
	return o
}

func (o ManagementtunnelMapOutput) ToManagementtunnelMapOutputWithContext(ctx context.Context) ManagementtunnelMapOutput {
	return o
}

func (o ManagementtunnelMapOutput) MapIndex(k pulumi.StringInput) ManagementtunnelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Managementtunnel {
		return vs[0].(map[string]*Managementtunnel)[vs[1].(string)]
	}).(ManagementtunnelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManagementtunnelInput)(nil)).Elem(), &Managementtunnel{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagementtunnelArrayInput)(nil)).Elem(), ManagementtunnelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagementtunnelMapInput)(nil)).Elem(), ManagementtunnelMap{})
	pulumi.RegisterOutputType(ManagementtunnelOutput{})
	pulumi.RegisterOutputType(ManagementtunnelArrayOutput{})
	pulumi.RegisterOutputType(ManagementtunnelMapOutput{})
}
