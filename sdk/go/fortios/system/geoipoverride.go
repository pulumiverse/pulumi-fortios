// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.
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
//			_, err := system.NewGeoipoverride(ctx, "trname", &system.GeoipoverrideArgs{
//				Description: pulumi.String("TEST for country"),
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
// System GeoipOverride can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/geoipoverride:Geoipoverride labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/geoipoverride:Geoipoverride labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Geoipoverride struct {
	pulumi.CustomResourceState

	// Two character Country ID code.
	CountryId pulumi.StringOutput `pulumi:"countryId"`
	// Description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Table of IPv6 ranges assigned to country. The structure of `ip6Range` block is documented below.
	Ip6Ranges GeoipoverrideIp6RangeArrayOutput `pulumi:"ip6Ranges"`
	// Table of IP ranges assigned to country. The structure of `ipRange` block is documented below.
	IpRanges GeoipoverrideIpRangeArrayOutput `pulumi:"ipRanges"`
	// Location name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewGeoipoverride registers a new resource with the given unique name, arguments, and options.
func NewGeoipoverride(ctx *pulumi.Context,
	name string, args *GeoipoverrideArgs, opts ...pulumi.ResourceOption) (*Geoipoverride, error) {
	if args == nil {
		args = &GeoipoverrideArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Geoipoverride
	err := ctx.RegisterResource("fortios:system/geoipoverride:Geoipoverride", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGeoipoverride gets an existing Geoipoverride resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGeoipoverride(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GeoipoverrideState, opts ...pulumi.ResourceOption) (*Geoipoverride, error) {
	var resource Geoipoverride
	err := ctx.ReadResource("fortios:system/geoipoverride:Geoipoverride", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Geoipoverride resources.
type geoipoverrideState struct {
	// Two character Country ID code.
	CountryId *string `pulumi:"countryId"`
	// Description.
	Description *string `pulumi:"description"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Table of IPv6 ranges assigned to country. The structure of `ip6Range` block is documented below.
	Ip6Ranges []GeoipoverrideIp6Range `pulumi:"ip6Ranges"`
	// Table of IP ranges assigned to country. The structure of `ipRange` block is documented below.
	IpRanges []GeoipoverrideIpRange `pulumi:"ipRanges"`
	// Location name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type GeoipoverrideState struct {
	// Two character Country ID code.
	CountryId pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Table of IPv6 ranges assigned to country. The structure of `ip6Range` block is documented below.
	Ip6Ranges GeoipoverrideIp6RangeArrayInput
	// Table of IP ranges assigned to country. The structure of `ipRange` block is documented below.
	IpRanges GeoipoverrideIpRangeArrayInput
	// Location name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (GeoipoverrideState) ElementType() reflect.Type {
	return reflect.TypeOf((*geoipoverrideState)(nil)).Elem()
}

type geoipoverrideArgs struct {
	// Two character Country ID code.
	CountryId *string `pulumi:"countryId"`
	// Description.
	Description *string `pulumi:"description"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Table of IPv6 ranges assigned to country. The structure of `ip6Range` block is documented below.
	Ip6Ranges []GeoipoverrideIp6Range `pulumi:"ip6Ranges"`
	// Table of IP ranges assigned to country. The structure of `ipRange` block is documented below.
	IpRanges []GeoipoverrideIpRange `pulumi:"ipRanges"`
	// Location name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Geoipoverride resource.
type GeoipoverrideArgs struct {
	// Two character Country ID code.
	CountryId pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Table of IPv6 ranges assigned to country. The structure of `ip6Range` block is documented below.
	Ip6Ranges GeoipoverrideIp6RangeArrayInput
	// Table of IP ranges assigned to country. The structure of `ipRange` block is documented below.
	IpRanges GeoipoverrideIpRangeArrayInput
	// Location name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (GeoipoverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*geoipoverrideArgs)(nil)).Elem()
}

type GeoipoverrideInput interface {
	pulumi.Input

	ToGeoipoverrideOutput() GeoipoverrideOutput
	ToGeoipoverrideOutputWithContext(ctx context.Context) GeoipoverrideOutput
}

func (*Geoipoverride) ElementType() reflect.Type {
	return reflect.TypeOf((**Geoipoverride)(nil)).Elem()
}

func (i *Geoipoverride) ToGeoipoverrideOutput() GeoipoverrideOutput {
	return i.ToGeoipoverrideOutputWithContext(context.Background())
}

func (i *Geoipoverride) ToGeoipoverrideOutputWithContext(ctx context.Context) GeoipoverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoipoverrideOutput)
}

// GeoipoverrideArrayInput is an input type that accepts GeoipoverrideArray and GeoipoverrideArrayOutput values.
// You can construct a concrete instance of `GeoipoverrideArrayInput` via:
//
//	GeoipoverrideArray{ GeoipoverrideArgs{...} }
type GeoipoverrideArrayInput interface {
	pulumi.Input

	ToGeoipoverrideArrayOutput() GeoipoverrideArrayOutput
	ToGeoipoverrideArrayOutputWithContext(context.Context) GeoipoverrideArrayOutput
}

type GeoipoverrideArray []GeoipoverrideInput

func (GeoipoverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Geoipoverride)(nil)).Elem()
}

func (i GeoipoverrideArray) ToGeoipoverrideArrayOutput() GeoipoverrideArrayOutput {
	return i.ToGeoipoverrideArrayOutputWithContext(context.Background())
}

func (i GeoipoverrideArray) ToGeoipoverrideArrayOutputWithContext(ctx context.Context) GeoipoverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoipoverrideArrayOutput)
}

// GeoipoverrideMapInput is an input type that accepts GeoipoverrideMap and GeoipoverrideMapOutput values.
// You can construct a concrete instance of `GeoipoverrideMapInput` via:
//
//	GeoipoverrideMap{ "key": GeoipoverrideArgs{...} }
type GeoipoverrideMapInput interface {
	pulumi.Input

	ToGeoipoverrideMapOutput() GeoipoverrideMapOutput
	ToGeoipoverrideMapOutputWithContext(context.Context) GeoipoverrideMapOutput
}

type GeoipoverrideMap map[string]GeoipoverrideInput

func (GeoipoverrideMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Geoipoverride)(nil)).Elem()
}

func (i GeoipoverrideMap) ToGeoipoverrideMapOutput() GeoipoverrideMapOutput {
	return i.ToGeoipoverrideMapOutputWithContext(context.Background())
}

func (i GeoipoverrideMap) ToGeoipoverrideMapOutputWithContext(ctx context.Context) GeoipoverrideMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoipoverrideMapOutput)
}

type GeoipoverrideOutput struct{ *pulumi.OutputState }

func (GeoipoverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Geoipoverride)(nil)).Elem()
}

func (o GeoipoverrideOutput) ToGeoipoverrideOutput() GeoipoverrideOutput {
	return o
}

func (o GeoipoverrideOutput) ToGeoipoverrideOutputWithContext(ctx context.Context) GeoipoverrideOutput {
	return o
}

// Two character Country ID code.
func (o GeoipoverrideOutput) CountryId() pulumi.StringOutput {
	return o.ApplyT(func(v *Geoipoverride) pulumi.StringOutput { return v.CountryId }).(pulumi.StringOutput)
}

// Description.
func (o GeoipoverrideOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Geoipoverride) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o GeoipoverrideOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Geoipoverride) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o GeoipoverrideOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Geoipoverride) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Table of IPv6 ranges assigned to country. The structure of `ip6Range` block is documented below.
func (o GeoipoverrideOutput) Ip6Ranges() GeoipoverrideIp6RangeArrayOutput {
	return o.ApplyT(func(v *Geoipoverride) GeoipoverrideIp6RangeArrayOutput { return v.Ip6Ranges }).(GeoipoverrideIp6RangeArrayOutput)
}

// Table of IP ranges assigned to country. The structure of `ipRange` block is documented below.
func (o GeoipoverrideOutput) IpRanges() GeoipoverrideIpRangeArrayOutput {
	return o.ApplyT(func(v *Geoipoverride) GeoipoverrideIpRangeArrayOutput { return v.IpRanges }).(GeoipoverrideIpRangeArrayOutput)
}

// Location name.
func (o GeoipoverrideOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Geoipoverride) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o GeoipoverrideOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Geoipoverride) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type GeoipoverrideArrayOutput struct{ *pulumi.OutputState }

func (GeoipoverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Geoipoverride)(nil)).Elem()
}

func (o GeoipoverrideArrayOutput) ToGeoipoverrideArrayOutput() GeoipoverrideArrayOutput {
	return o
}

func (o GeoipoverrideArrayOutput) ToGeoipoverrideArrayOutputWithContext(ctx context.Context) GeoipoverrideArrayOutput {
	return o
}

func (o GeoipoverrideArrayOutput) Index(i pulumi.IntInput) GeoipoverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Geoipoverride {
		return vs[0].([]*Geoipoverride)[vs[1].(int)]
	}).(GeoipoverrideOutput)
}

type GeoipoverrideMapOutput struct{ *pulumi.OutputState }

func (GeoipoverrideMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Geoipoverride)(nil)).Elem()
}

func (o GeoipoverrideMapOutput) ToGeoipoverrideMapOutput() GeoipoverrideMapOutput {
	return o
}

func (o GeoipoverrideMapOutput) ToGeoipoverrideMapOutputWithContext(ctx context.Context) GeoipoverrideMapOutput {
	return o
}

func (o GeoipoverrideMapOutput) MapIndex(k pulumi.StringInput) GeoipoverrideOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Geoipoverride {
		return vs[0].(map[string]*Geoipoverride)[vs[1].(string)]
	}).(GeoipoverrideOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GeoipoverrideInput)(nil)).Elem(), &Geoipoverride{})
	pulumi.RegisterInputType(reflect.TypeOf((*GeoipoverrideArrayInput)(nil)).Elem(), GeoipoverrideArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GeoipoverrideMapInput)(nil)).Elem(), GeoipoverrideMap{})
	pulumi.RegisterOutputType(GeoipoverrideOutput{})
	pulumi.RegisterOutputType(GeoipoverrideArrayOutput{})
	pulumi.RegisterOutputType(GeoipoverrideMapOutput{})
}
