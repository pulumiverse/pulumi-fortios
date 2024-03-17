// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Define internet service names. Applies to FortiOS Version `>= 6.4.0`.
//
// ## Import
//
// Firewall InternetServiceName can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/internetservicename:Internetservicename labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/internetservicename:Internetservicename labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Internetservicename struct {
	pulumi.CustomResourceState

	// City ID.
	CityId pulumi.IntOutput `pulumi:"cityId"`
	// Country or Area ID.
	CountryId pulumi.IntOutput `pulumi:"countryId"`
	// Internet Service ID.
	InternetServiceId pulumi.IntOutput `pulumi:"internetServiceId"`
	// Internet Service name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Region ID.
	RegionId pulumi.IntOutput `pulumi:"regionId"`
	// Internet Service name type. Valid values: `default`, `location`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewInternetservicename registers a new resource with the given unique name, arguments, and options.
func NewInternetservicename(ctx *pulumi.Context,
	name string, args *InternetservicenameArgs, opts ...pulumi.ResourceOption) (*Internetservicename, error) {
	if args == nil {
		args = &InternetservicenameArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Internetservicename
	err := ctx.RegisterResource("fortios:firewall/internetservicename:Internetservicename", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInternetservicename gets an existing Internetservicename resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInternetservicename(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InternetservicenameState, opts ...pulumi.ResourceOption) (*Internetservicename, error) {
	var resource Internetservicename
	err := ctx.ReadResource("fortios:firewall/internetservicename:Internetservicename", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Internetservicename resources.
type internetservicenameState struct {
	// City ID.
	CityId *int `pulumi:"cityId"`
	// Country or Area ID.
	CountryId *int `pulumi:"countryId"`
	// Internet Service ID.
	InternetServiceId *int `pulumi:"internetServiceId"`
	// Internet Service name.
	Name *string `pulumi:"name"`
	// Region ID.
	RegionId *int `pulumi:"regionId"`
	// Internet Service name type. Valid values: `default`, `location`.
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type InternetservicenameState struct {
	// City ID.
	CityId pulumi.IntPtrInput
	// Country or Area ID.
	CountryId pulumi.IntPtrInput
	// Internet Service ID.
	InternetServiceId pulumi.IntPtrInput
	// Internet Service name.
	Name pulumi.StringPtrInput
	// Region ID.
	RegionId pulumi.IntPtrInput
	// Internet Service name type. Valid values: `default`, `location`.
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (InternetservicenameState) ElementType() reflect.Type {
	return reflect.TypeOf((*internetservicenameState)(nil)).Elem()
}

type internetservicenameArgs struct {
	// City ID.
	CityId *int `pulumi:"cityId"`
	// Country or Area ID.
	CountryId *int `pulumi:"countryId"`
	// Internet Service ID.
	InternetServiceId *int `pulumi:"internetServiceId"`
	// Internet Service name.
	Name *string `pulumi:"name"`
	// Region ID.
	RegionId *int `pulumi:"regionId"`
	// Internet Service name type. Valid values: `default`, `location`.
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Internetservicename resource.
type InternetservicenameArgs struct {
	// City ID.
	CityId pulumi.IntPtrInput
	// Country or Area ID.
	CountryId pulumi.IntPtrInput
	// Internet Service ID.
	InternetServiceId pulumi.IntPtrInput
	// Internet Service name.
	Name pulumi.StringPtrInput
	// Region ID.
	RegionId pulumi.IntPtrInput
	// Internet Service name type. Valid values: `default`, `location`.
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (InternetservicenameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*internetservicenameArgs)(nil)).Elem()
}

type InternetservicenameInput interface {
	pulumi.Input

	ToInternetservicenameOutput() InternetservicenameOutput
	ToInternetservicenameOutputWithContext(ctx context.Context) InternetservicenameOutput
}

func (*Internetservicename) ElementType() reflect.Type {
	return reflect.TypeOf((**Internetservicename)(nil)).Elem()
}

func (i *Internetservicename) ToInternetservicenameOutput() InternetservicenameOutput {
	return i.ToInternetservicenameOutputWithContext(context.Background())
}

func (i *Internetservicename) ToInternetservicenameOutputWithContext(ctx context.Context) InternetservicenameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetservicenameOutput)
}

// InternetservicenameArrayInput is an input type that accepts InternetservicenameArray and InternetservicenameArrayOutput values.
// You can construct a concrete instance of `InternetservicenameArrayInput` via:
//
//	InternetservicenameArray{ InternetservicenameArgs{...} }
type InternetservicenameArrayInput interface {
	pulumi.Input

	ToInternetservicenameArrayOutput() InternetservicenameArrayOutput
	ToInternetservicenameArrayOutputWithContext(context.Context) InternetservicenameArrayOutput
}

type InternetservicenameArray []InternetservicenameInput

func (InternetservicenameArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Internetservicename)(nil)).Elem()
}

func (i InternetservicenameArray) ToInternetservicenameArrayOutput() InternetservicenameArrayOutput {
	return i.ToInternetservicenameArrayOutputWithContext(context.Background())
}

func (i InternetservicenameArray) ToInternetservicenameArrayOutputWithContext(ctx context.Context) InternetservicenameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetservicenameArrayOutput)
}

// InternetservicenameMapInput is an input type that accepts InternetservicenameMap and InternetservicenameMapOutput values.
// You can construct a concrete instance of `InternetservicenameMapInput` via:
//
//	InternetservicenameMap{ "key": InternetservicenameArgs{...} }
type InternetservicenameMapInput interface {
	pulumi.Input

	ToInternetservicenameMapOutput() InternetservicenameMapOutput
	ToInternetservicenameMapOutputWithContext(context.Context) InternetservicenameMapOutput
}

type InternetservicenameMap map[string]InternetservicenameInput

func (InternetservicenameMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Internetservicename)(nil)).Elem()
}

func (i InternetservicenameMap) ToInternetservicenameMapOutput() InternetservicenameMapOutput {
	return i.ToInternetservicenameMapOutputWithContext(context.Background())
}

func (i InternetservicenameMap) ToInternetservicenameMapOutputWithContext(ctx context.Context) InternetservicenameMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetservicenameMapOutput)
}

type InternetservicenameOutput struct{ *pulumi.OutputState }

func (InternetservicenameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Internetservicename)(nil)).Elem()
}

func (o InternetservicenameOutput) ToInternetservicenameOutput() InternetservicenameOutput {
	return o
}

func (o InternetservicenameOutput) ToInternetservicenameOutputWithContext(ctx context.Context) InternetservicenameOutput {
	return o
}

// City ID.
func (o InternetservicenameOutput) CityId() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetservicename) pulumi.IntOutput { return v.CityId }).(pulumi.IntOutput)
}

// Country or Area ID.
func (o InternetservicenameOutput) CountryId() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetservicename) pulumi.IntOutput { return v.CountryId }).(pulumi.IntOutput)
}

// Internet Service ID.
func (o InternetservicenameOutput) InternetServiceId() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetservicename) pulumi.IntOutput { return v.InternetServiceId }).(pulumi.IntOutput)
}

// Internet Service name.
func (o InternetservicenameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Internetservicename) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Region ID.
func (o InternetservicenameOutput) RegionId() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetservicename) pulumi.IntOutput { return v.RegionId }).(pulumi.IntOutput)
}

// Internet Service name type. Valid values: `default`, `location`.
func (o InternetservicenameOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Internetservicename) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o InternetservicenameOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Internetservicename) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type InternetservicenameArrayOutput struct{ *pulumi.OutputState }

func (InternetservicenameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Internetservicename)(nil)).Elem()
}

func (o InternetservicenameArrayOutput) ToInternetservicenameArrayOutput() InternetservicenameArrayOutput {
	return o
}

func (o InternetservicenameArrayOutput) ToInternetservicenameArrayOutputWithContext(ctx context.Context) InternetservicenameArrayOutput {
	return o
}

func (o InternetservicenameArrayOutput) Index(i pulumi.IntInput) InternetservicenameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Internetservicename {
		return vs[0].([]*Internetservicename)[vs[1].(int)]
	}).(InternetservicenameOutput)
}

type InternetservicenameMapOutput struct{ *pulumi.OutputState }

func (InternetservicenameMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Internetservicename)(nil)).Elem()
}

func (o InternetservicenameMapOutput) ToInternetservicenameMapOutput() InternetservicenameMapOutput {
	return o
}

func (o InternetservicenameMapOutput) ToInternetservicenameMapOutputWithContext(ctx context.Context) InternetservicenameMapOutput {
	return o
}

func (o InternetservicenameMapOutput) MapIndex(k pulumi.StringInput) InternetservicenameOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Internetservicename {
		return vs[0].(map[string]*Internetservicename)[vs[1].(string)]
	}).(InternetservicenameOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InternetservicenameInput)(nil)).Elem(), &Internetservicename{})
	pulumi.RegisterInputType(reflect.TypeOf((*InternetservicenameArrayInput)(nil)).Elem(), InternetservicenameArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InternetservicenameMapInput)(nil)).Elem(), InternetservicenameMap{})
	pulumi.RegisterOutputType(InternetservicenameOutput{})
	pulumi.RegisterOutputType(InternetservicenameArrayOutput{})
	pulumi.RegisterOutputType(InternetservicenameMapOutput{})
}
