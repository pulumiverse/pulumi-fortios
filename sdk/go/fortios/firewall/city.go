// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Define city table. Applies to FortiOS Version `>= 6.4.0`.
//
// ## Import
//
// Firewall City can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/city:City labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/city:City labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type City struct {
	pulumi.CustomResourceState

	// City ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// City name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewCity registers a new resource with the given unique name, arguments, and options.
func NewCity(ctx *pulumi.Context,
	name string, args *CityArgs, opts ...pulumi.ResourceOption) (*City, error) {
	if args == nil {
		args = &CityArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource City
	err := ctx.RegisterResource("fortios:firewall/city:City", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCity gets an existing City resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CityState, opts ...pulumi.ResourceOption) (*City, error) {
	var resource City
	err := ctx.ReadResource("fortios:firewall/city:City", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering City resources.
type cityState struct {
	// City ID.
	Fosid *int `pulumi:"fosid"`
	// City name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type CityState struct {
	// City ID.
	Fosid pulumi.IntPtrInput
	// City name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (CityState) ElementType() reflect.Type {
	return reflect.TypeOf((*cityState)(nil)).Elem()
}

type cityArgs struct {
	// City ID.
	Fosid *int `pulumi:"fosid"`
	// City name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a City resource.
type CityArgs struct {
	// City ID.
	Fosid pulumi.IntPtrInput
	// City name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (CityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cityArgs)(nil)).Elem()
}

type CityInput interface {
	pulumi.Input

	ToCityOutput() CityOutput
	ToCityOutputWithContext(ctx context.Context) CityOutput
}

func (*City) ElementType() reflect.Type {
	return reflect.TypeOf((**City)(nil)).Elem()
}

func (i *City) ToCityOutput() CityOutput {
	return i.ToCityOutputWithContext(context.Background())
}

func (i *City) ToCityOutputWithContext(ctx context.Context) CityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CityOutput)
}

// CityArrayInput is an input type that accepts CityArray and CityArrayOutput values.
// You can construct a concrete instance of `CityArrayInput` via:
//
//	CityArray{ CityArgs{...} }
type CityArrayInput interface {
	pulumi.Input

	ToCityArrayOutput() CityArrayOutput
	ToCityArrayOutputWithContext(context.Context) CityArrayOutput
}

type CityArray []CityInput

func (CityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*City)(nil)).Elem()
}

func (i CityArray) ToCityArrayOutput() CityArrayOutput {
	return i.ToCityArrayOutputWithContext(context.Background())
}

func (i CityArray) ToCityArrayOutputWithContext(ctx context.Context) CityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CityArrayOutput)
}

// CityMapInput is an input type that accepts CityMap and CityMapOutput values.
// You can construct a concrete instance of `CityMapInput` via:
//
//	CityMap{ "key": CityArgs{...} }
type CityMapInput interface {
	pulumi.Input

	ToCityMapOutput() CityMapOutput
	ToCityMapOutputWithContext(context.Context) CityMapOutput
}

type CityMap map[string]CityInput

func (CityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*City)(nil)).Elem()
}

func (i CityMap) ToCityMapOutput() CityMapOutput {
	return i.ToCityMapOutputWithContext(context.Background())
}

func (i CityMap) ToCityMapOutputWithContext(ctx context.Context) CityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CityMapOutput)
}

type CityOutput struct{ *pulumi.OutputState }

func (CityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**City)(nil)).Elem()
}

func (o CityOutput) ToCityOutput() CityOutput {
	return o
}

func (o CityOutput) ToCityOutputWithContext(ctx context.Context) CityOutput {
	return o
}

// City ID.
func (o CityOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *City) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// City name.
func (o CityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *City) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o CityOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *City) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type CityArrayOutput struct{ *pulumi.OutputState }

func (CityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*City)(nil)).Elem()
}

func (o CityArrayOutput) ToCityArrayOutput() CityArrayOutput {
	return o
}

func (o CityArrayOutput) ToCityArrayOutputWithContext(ctx context.Context) CityArrayOutput {
	return o
}

func (o CityArrayOutput) Index(i pulumi.IntInput) CityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *City {
		return vs[0].([]*City)[vs[1].(int)]
	}).(CityOutput)
}

type CityMapOutput struct{ *pulumi.OutputState }

func (CityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*City)(nil)).Elem()
}

func (o CityMapOutput) ToCityMapOutput() CityMapOutput {
	return o
}

func (o CityMapOutput) ToCityMapOutputWithContext(ctx context.Context) CityMapOutput {
	return o
}

func (o CityMapOutput) MapIndex(k pulumi.StringInput) CityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *City {
		return vs[0].(map[string]*City)[vs[1].(string)]
	}).(CityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CityInput)(nil)).Elem(), &City{})
	pulumi.RegisterInputType(reflect.TypeOf((*CityArrayInput)(nil)).Elem(), CityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CityMapInput)(nil)).Elem(), CityMap{})
	pulumi.RegisterOutputType(CityOutput{})
	pulumi.RegisterOutputType(CityArrayOutput{})
	pulumi.RegisterOutputType(CityMapOutput{})
}
