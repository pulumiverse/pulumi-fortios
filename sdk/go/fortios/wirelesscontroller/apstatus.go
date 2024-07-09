// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wirelesscontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure access point status (rogue | accepted | suppressed).
//
// ## Import
//
// WirelessController ApStatus can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/apstatus:Apstatus labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/apstatus:Apstatus labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Apstatus struct {
	pulumi.CustomResourceState

	// Access Point's (AP's) BSSID.
	Bssid pulumi.StringOutput `pulumi:"bssid"`
	// AP ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Access Point's (AP's) SSID.
	Ssid pulumi.StringOutput `pulumi:"ssid"`
	// Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewApstatus registers a new resource with the given unique name, arguments, and options.
func NewApstatus(ctx *pulumi.Context,
	name string, args *ApstatusArgs, opts ...pulumi.ResourceOption) (*Apstatus, error) {
	if args == nil {
		args = &ApstatusArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Apstatus
	err := ctx.RegisterResource("fortios:wirelesscontroller/apstatus:Apstatus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApstatus gets an existing Apstatus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApstatus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApstatusState, opts ...pulumi.ResourceOption) (*Apstatus, error) {
	var resource Apstatus
	err := ctx.ReadResource("fortios:wirelesscontroller/apstatus:Apstatus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Apstatus resources.
type apstatusState struct {
	// Access Point's (AP's) BSSID.
	Bssid *string `pulumi:"bssid"`
	// AP ID.
	Fosid *int `pulumi:"fosid"`
	// Access Point's (AP's) SSID.
	Ssid *string `pulumi:"ssid"`
	// Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ApstatusState struct {
	// Access Point's (AP's) BSSID.
	Bssid pulumi.StringPtrInput
	// AP ID.
	Fosid pulumi.IntPtrInput
	// Access Point's (AP's) SSID.
	Ssid pulumi.StringPtrInput
	// Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ApstatusState) ElementType() reflect.Type {
	return reflect.TypeOf((*apstatusState)(nil)).Elem()
}

type apstatusArgs struct {
	// Access Point's (AP's) BSSID.
	Bssid *string `pulumi:"bssid"`
	// AP ID.
	Fosid *int `pulumi:"fosid"`
	// Access Point's (AP's) SSID.
	Ssid *string `pulumi:"ssid"`
	// Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Apstatus resource.
type ApstatusArgs struct {
	// Access Point's (AP's) BSSID.
	Bssid pulumi.StringPtrInput
	// AP ID.
	Fosid pulumi.IntPtrInput
	// Access Point's (AP's) SSID.
	Ssid pulumi.StringPtrInput
	// Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ApstatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apstatusArgs)(nil)).Elem()
}

type ApstatusInput interface {
	pulumi.Input

	ToApstatusOutput() ApstatusOutput
	ToApstatusOutputWithContext(ctx context.Context) ApstatusOutput
}

func (*Apstatus) ElementType() reflect.Type {
	return reflect.TypeOf((**Apstatus)(nil)).Elem()
}

func (i *Apstatus) ToApstatusOutput() ApstatusOutput {
	return i.ToApstatusOutputWithContext(context.Background())
}

func (i *Apstatus) ToApstatusOutputWithContext(ctx context.Context) ApstatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApstatusOutput)
}

// ApstatusArrayInput is an input type that accepts ApstatusArray and ApstatusArrayOutput values.
// You can construct a concrete instance of `ApstatusArrayInput` via:
//
//	ApstatusArray{ ApstatusArgs{...} }
type ApstatusArrayInput interface {
	pulumi.Input

	ToApstatusArrayOutput() ApstatusArrayOutput
	ToApstatusArrayOutputWithContext(context.Context) ApstatusArrayOutput
}

type ApstatusArray []ApstatusInput

func (ApstatusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Apstatus)(nil)).Elem()
}

func (i ApstatusArray) ToApstatusArrayOutput() ApstatusArrayOutput {
	return i.ToApstatusArrayOutputWithContext(context.Background())
}

func (i ApstatusArray) ToApstatusArrayOutputWithContext(ctx context.Context) ApstatusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApstatusArrayOutput)
}

// ApstatusMapInput is an input type that accepts ApstatusMap and ApstatusMapOutput values.
// You can construct a concrete instance of `ApstatusMapInput` via:
//
//	ApstatusMap{ "key": ApstatusArgs{...} }
type ApstatusMapInput interface {
	pulumi.Input

	ToApstatusMapOutput() ApstatusMapOutput
	ToApstatusMapOutputWithContext(context.Context) ApstatusMapOutput
}

type ApstatusMap map[string]ApstatusInput

func (ApstatusMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Apstatus)(nil)).Elem()
}

func (i ApstatusMap) ToApstatusMapOutput() ApstatusMapOutput {
	return i.ToApstatusMapOutputWithContext(context.Background())
}

func (i ApstatusMap) ToApstatusMapOutputWithContext(ctx context.Context) ApstatusMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApstatusMapOutput)
}

type ApstatusOutput struct{ *pulumi.OutputState }

func (ApstatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Apstatus)(nil)).Elem()
}

func (o ApstatusOutput) ToApstatusOutput() ApstatusOutput {
	return o
}

func (o ApstatusOutput) ToApstatusOutputWithContext(ctx context.Context) ApstatusOutput {
	return o
}

// Access Point's (AP's) BSSID.
func (o ApstatusOutput) Bssid() pulumi.StringOutput {
	return o.ApplyT(func(v *Apstatus) pulumi.StringOutput { return v.Bssid }).(pulumi.StringOutput)
}

// AP ID.
func (o ApstatusOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Apstatus) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Access Point's (AP's) SSID.
func (o ApstatusOutput) Ssid() pulumi.StringOutput {
	return o.ApplyT(func(v *Apstatus) pulumi.StringOutput { return v.Ssid }).(pulumi.StringOutput)
}

// Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
func (o ApstatusOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Apstatus) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ApstatusOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Apstatus) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type ApstatusArrayOutput struct{ *pulumi.OutputState }

func (ApstatusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Apstatus)(nil)).Elem()
}

func (o ApstatusArrayOutput) ToApstatusArrayOutput() ApstatusArrayOutput {
	return o
}

func (o ApstatusArrayOutput) ToApstatusArrayOutputWithContext(ctx context.Context) ApstatusArrayOutput {
	return o
}

func (o ApstatusArrayOutput) Index(i pulumi.IntInput) ApstatusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Apstatus {
		return vs[0].([]*Apstatus)[vs[1].(int)]
	}).(ApstatusOutput)
}

type ApstatusMapOutput struct{ *pulumi.OutputState }

func (ApstatusMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Apstatus)(nil)).Elem()
}

func (o ApstatusMapOutput) ToApstatusMapOutput() ApstatusMapOutput {
	return o
}

func (o ApstatusMapOutput) ToApstatusMapOutputWithContext(ctx context.Context) ApstatusMapOutput {
	return o
}

func (o ApstatusMapOutput) MapIndex(k pulumi.StringInput) ApstatusOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Apstatus {
		return vs[0].(map[string]*Apstatus)[vs[1].(string)]
	}).(ApstatusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApstatusInput)(nil)).Elem(), &Apstatus{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApstatusArrayInput)(nil)).Elem(), ApstatusArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApstatusMapInput)(nil)).Elem(), ApstatusMap{})
	pulumi.RegisterOutputType(ApstatusOutput{})
	pulumi.RegisterOutputType(ApstatusArrayOutput{})
	pulumi.RegisterOutputType(ApstatusMapOutput{})
}
