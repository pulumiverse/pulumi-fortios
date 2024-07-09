// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure PPTP.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/vpn"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpn.NewPptp(ctx, "trname", &vpn.PptpArgs{
//				Eip:     pulumi.String("1.1.1.22"),
//				IpMode:  pulumi.String("range"),
//				LocalIp: pulumi.String("0.0.0.0"),
//				Sip:     pulumi.String("1.1.1.1"),
//				Status:  pulumi.String("enable"),
//				Usrgrp:  pulumi.String("Guest-group"),
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
// Vpn Pptp can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:vpn/pptp:Pptp labelname VpnPptp
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:vpn/pptp:Pptp labelname VpnPptp
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Pptp struct {
	pulumi.CustomResourceState

	// End IP.
	Eip pulumi.StringOutput `pulumi:"eip"`
	// IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
	IpMode pulumi.StringOutput `pulumi:"ipMode"`
	// Local IP to be used for peer's remote IP.
	LocalIp pulumi.StringOutput `pulumi:"localIp"`
	// Start IP.
	Sip pulumi.StringOutput `pulumi:"sip"`
	// Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// User group.
	Usrgrp pulumi.StringOutput `pulumi:"usrgrp"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewPptp registers a new resource with the given unique name, arguments, and options.
func NewPptp(ctx *pulumi.Context,
	name string, args *PptpArgs, opts ...pulumi.ResourceOption) (*Pptp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Pptp
	err := ctx.RegisterResource("fortios:vpn/pptp:Pptp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPptp gets an existing Pptp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPptp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PptpState, opts ...pulumi.ResourceOption) (*Pptp, error) {
	var resource Pptp
	err := ctx.ReadResource("fortios:vpn/pptp:Pptp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pptp resources.
type pptpState struct {
	// End IP.
	Eip *string `pulumi:"eip"`
	// IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
	IpMode *string `pulumi:"ipMode"`
	// Local IP to be used for peer's remote IP.
	LocalIp *string `pulumi:"localIp"`
	// Start IP.
	Sip *string `pulumi:"sip"`
	// Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// User group.
	Usrgrp *string `pulumi:"usrgrp"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type PptpState struct {
	// End IP.
	Eip pulumi.StringPtrInput
	// IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
	IpMode pulumi.StringPtrInput
	// Local IP to be used for peer's remote IP.
	LocalIp pulumi.StringPtrInput
	// Start IP.
	Sip pulumi.StringPtrInput
	// Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// User group.
	Usrgrp pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (PptpState) ElementType() reflect.Type {
	return reflect.TypeOf((*pptpState)(nil)).Elem()
}

type pptpArgs struct {
	// End IP.
	Eip *string `pulumi:"eip"`
	// IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
	IpMode *string `pulumi:"ipMode"`
	// Local IP to be used for peer's remote IP.
	LocalIp *string `pulumi:"localIp"`
	// Start IP.
	Sip *string `pulumi:"sip"`
	// Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
	Status string `pulumi:"status"`
	// User group.
	Usrgrp *string `pulumi:"usrgrp"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Pptp resource.
type PptpArgs struct {
	// End IP.
	Eip pulumi.StringPtrInput
	// IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
	IpMode pulumi.StringPtrInput
	// Local IP to be used for peer's remote IP.
	LocalIp pulumi.StringPtrInput
	// Start IP.
	Sip pulumi.StringPtrInput
	// Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
	Status pulumi.StringInput
	// User group.
	Usrgrp pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (PptpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pptpArgs)(nil)).Elem()
}

type PptpInput interface {
	pulumi.Input

	ToPptpOutput() PptpOutput
	ToPptpOutputWithContext(ctx context.Context) PptpOutput
}

func (*Pptp) ElementType() reflect.Type {
	return reflect.TypeOf((**Pptp)(nil)).Elem()
}

func (i *Pptp) ToPptpOutput() PptpOutput {
	return i.ToPptpOutputWithContext(context.Background())
}

func (i *Pptp) ToPptpOutputWithContext(ctx context.Context) PptpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PptpOutput)
}

// PptpArrayInput is an input type that accepts PptpArray and PptpArrayOutput values.
// You can construct a concrete instance of `PptpArrayInput` via:
//
//	PptpArray{ PptpArgs{...} }
type PptpArrayInput interface {
	pulumi.Input

	ToPptpArrayOutput() PptpArrayOutput
	ToPptpArrayOutputWithContext(context.Context) PptpArrayOutput
}

type PptpArray []PptpInput

func (PptpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pptp)(nil)).Elem()
}

func (i PptpArray) ToPptpArrayOutput() PptpArrayOutput {
	return i.ToPptpArrayOutputWithContext(context.Background())
}

func (i PptpArray) ToPptpArrayOutputWithContext(ctx context.Context) PptpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PptpArrayOutput)
}

// PptpMapInput is an input type that accepts PptpMap and PptpMapOutput values.
// You can construct a concrete instance of `PptpMapInput` via:
//
//	PptpMap{ "key": PptpArgs{...} }
type PptpMapInput interface {
	pulumi.Input

	ToPptpMapOutput() PptpMapOutput
	ToPptpMapOutputWithContext(context.Context) PptpMapOutput
}

type PptpMap map[string]PptpInput

func (PptpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pptp)(nil)).Elem()
}

func (i PptpMap) ToPptpMapOutput() PptpMapOutput {
	return i.ToPptpMapOutputWithContext(context.Background())
}

func (i PptpMap) ToPptpMapOutputWithContext(ctx context.Context) PptpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PptpMapOutput)
}

type PptpOutput struct{ *pulumi.OutputState }

func (PptpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pptp)(nil)).Elem()
}

func (o PptpOutput) ToPptpOutput() PptpOutput {
	return o
}

func (o PptpOutput) ToPptpOutputWithContext(ctx context.Context) PptpOutput {
	return o
}

// End IP.
func (o PptpOutput) Eip() pulumi.StringOutput {
	return o.ApplyT(func(v *Pptp) pulumi.StringOutput { return v.Eip }).(pulumi.StringOutput)
}

// IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
func (o PptpOutput) IpMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Pptp) pulumi.StringOutput { return v.IpMode }).(pulumi.StringOutput)
}

// Local IP to be used for peer's remote IP.
func (o PptpOutput) LocalIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Pptp) pulumi.StringOutput { return v.LocalIp }).(pulumi.StringOutput)
}

// Start IP.
func (o PptpOutput) Sip() pulumi.StringOutput {
	return o.ApplyT(func(v *Pptp) pulumi.StringOutput { return v.Sip }).(pulumi.StringOutput)
}

// Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
func (o PptpOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Pptp) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// User group.
func (o PptpOutput) Usrgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *Pptp) pulumi.StringOutput { return v.Usrgrp }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o PptpOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pptp) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type PptpArrayOutput struct{ *pulumi.OutputState }

func (PptpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pptp)(nil)).Elem()
}

func (o PptpArrayOutput) ToPptpArrayOutput() PptpArrayOutput {
	return o
}

func (o PptpArrayOutput) ToPptpArrayOutputWithContext(ctx context.Context) PptpArrayOutput {
	return o
}

func (o PptpArrayOutput) Index(i pulumi.IntInput) PptpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Pptp {
		return vs[0].([]*Pptp)[vs[1].(int)]
	}).(PptpOutput)
}

type PptpMapOutput struct{ *pulumi.OutputState }

func (PptpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pptp)(nil)).Elem()
}

func (o PptpMapOutput) ToPptpMapOutput() PptpMapOutput {
	return o
}

func (o PptpMapOutput) ToPptpMapOutputWithContext(ctx context.Context) PptpMapOutput {
	return o
}

func (o PptpMapOutput) MapIndex(k pulumi.StringInput) PptpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Pptp {
		return vs[0].(map[string]*Pptp)[vs[1].(string)]
	}).(PptpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PptpInput)(nil)).Elem(), &Pptp{})
	pulumi.RegisterInputType(reflect.TypeOf((*PptpArrayInput)(nil)).Elem(), PptpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PptpMapInput)(nil)).Elem(), PptpMap{})
	pulumi.RegisterOutputType(PptpOutput{})
	pulumi.RegisterOutputType(PptpArrayOutput{})
	pulumi.RegisterOutputType(PptpMapOutput{})
}
