// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package user

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure device access control lists. Applies to FortiOS Version `<= 6.2.0`.
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
//			_, err := user.NewDeviceaccesslist(ctx, "trname", &user.DeviceaccesslistArgs{
//				DefaultAction: pulumi.String("accept"),
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
// User DeviceAccessList can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:user/deviceaccesslist:Deviceaccesslist labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:user/deviceaccesslist:Deviceaccesslist labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Deviceaccesslist struct {
	pulumi.CustomResourceState

	// Accept or deny unknown/unspecified devices. Valid values: `accept`, `deny`.
	DefaultAction pulumi.StringOutput `pulumi:"defaultAction"`
	// Device list. The structure of `deviceList` block is documented below.
	DeviceLists DeviceaccesslistDeviceListArrayOutput `pulumi:"deviceLists"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Device access list name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewDeviceaccesslist registers a new resource with the given unique name, arguments, and options.
func NewDeviceaccesslist(ctx *pulumi.Context,
	name string, args *DeviceaccesslistArgs, opts ...pulumi.ResourceOption) (*Deviceaccesslist, error) {
	if args == nil {
		args = &DeviceaccesslistArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Deviceaccesslist
	err := ctx.RegisterResource("fortios:user/deviceaccesslist:Deviceaccesslist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeviceaccesslist gets an existing Deviceaccesslist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeviceaccesslist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceaccesslistState, opts ...pulumi.ResourceOption) (*Deviceaccesslist, error) {
	var resource Deviceaccesslist
	err := ctx.ReadResource("fortios:user/deviceaccesslist:Deviceaccesslist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Deviceaccesslist resources.
type deviceaccesslistState struct {
	// Accept or deny unknown/unspecified devices. Valid values: `accept`, `deny`.
	DefaultAction *string `pulumi:"defaultAction"`
	// Device list. The structure of `deviceList` block is documented below.
	DeviceLists []DeviceaccesslistDeviceList `pulumi:"deviceLists"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Device access list name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DeviceaccesslistState struct {
	// Accept or deny unknown/unspecified devices. Valid values: `accept`, `deny`.
	DefaultAction pulumi.StringPtrInput
	// Device list. The structure of `deviceList` block is documented below.
	DeviceLists DeviceaccesslistDeviceListArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Device access list name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DeviceaccesslistState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceaccesslistState)(nil)).Elem()
}

type deviceaccesslistArgs struct {
	// Accept or deny unknown/unspecified devices. Valid values: `accept`, `deny`.
	DefaultAction *string `pulumi:"defaultAction"`
	// Device list. The structure of `deviceList` block is documented below.
	DeviceLists []DeviceaccesslistDeviceList `pulumi:"deviceLists"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Device access list name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Deviceaccesslist resource.
type DeviceaccesslistArgs struct {
	// Accept or deny unknown/unspecified devices. Valid values: `accept`, `deny`.
	DefaultAction pulumi.StringPtrInput
	// Device list. The structure of `deviceList` block is documented below.
	DeviceLists DeviceaccesslistDeviceListArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Device access list name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DeviceaccesslistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceaccesslistArgs)(nil)).Elem()
}

type DeviceaccesslistInput interface {
	pulumi.Input

	ToDeviceaccesslistOutput() DeviceaccesslistOutput
	ToDeviceaccesslistOutputWithContext(ctx context.Context) DeviceaccesslistOutput
}

func (*Deviceaccesslist) ElementType() reflect.Type {
	return reflect.TypeOf((**Deviceaccesslist)(nil)).Elem()
}

func (i *Deviceaccesslist) ToDeviceaccesslistOutput() DeviceaccesslistOutput {
	return i.ToDeviceaccesslistOutputWithContext(context.Background())
}

func (i *Deviceaccesslist) ToDeviceaccesslistOutputWithContext(ctx context.Context) DeviceaccesslistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceaccesslistOutput)
}

// DeviceaccesslistArrayInput is an input type that accepts DeviceaccesslistArray and DeviceaccesslistArrayOutput values.
// You can construct a concrete instance of `DeviceaccesslistArrayInput` via:
//
//	DeviceaccesslistArray{ DeviceaccesslistArgs{...} }
type DeviceaccesslistArrayInput interface {
	pulumi.Input

	ToDeviceaccesslistArrayOutput() DeviceaccesslistArrayOutput
	ToDeviceaccesslistArrayOutputWithContext(context.Context) DeviceaccesslistArrayOutput
}

type DeviceaccesslistArray []DeviceaccesslistInput

func (DeviceaccesslistArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Deviceaccesslist)(nil)).Elem()
}

func (i DeviceaccesslistArray) ToDeviceaccesslistArrayOutput() DeviceaccesslistArrayOutput {
	return i.ToDeviceaccesslistArrayOutputWithContext(context.Background())
}

func (i DeviceaccesslistArray) ToDeviceaccesslistArrayOutputWithContext(ctx context.Context) DeviceaccesslistArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceaccesslistArrayOutput)
}

// DeviceaccesslistMapInput is an input type that accepts DeviceaccesslistMap and DeviceaccesslistMapOutput values.
// You can construct a concrete instance of `DeviceaccesslistMapInput` via:
//
//	DeviceaccesslistMap{ "key": DeviceaccesslistArgs{...} }
type DeviceaccesslistMapInput interface {
	pulumi.Input

	ToDeviceaccesslistMapOutput() DeviceaccesslistMapOutput
	ToDeviceaccesslistMapOutputWithContext(context.Context) DeviceaccesslistMapOutput
}

type DeviceaccesslistMap map[string]DeviceaccesslistInput

func (DeviceaccesslistMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Deviceaccesslist)(nil)).Elem()
}

func (i DeviceaccesslistMap) ToDeviceaccesslistMapOutput() DeviceaccesslistMapOutput {
	return i.ToDeviceaccesslistMapOutputWithContext(context.Background())
}

func (i DeviceaccesslistMap) ToDeviceaccesslistMapOutputWithContext(ctx context.Context) DeviceaccesslistMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceaccesslistMapOutput)
}

type DeviceaccesslistOutput struct{ *pulumi.OutputState }

func (DeviceaccesslistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Deviceaccesslist)(nil)).Elem()
}

func (o DeviceaccesslistOutput) ToDeviceaccesslistOutput() DeviceaccesslistOutput {
	return o
}

func (o DeviceaccesslistOutput) ToDeviceaccesslistOutputWithContext(ctx context.Context) DeviceaccesslistOutput {
	return o
}

// Accept or deny unknown/unspecified devices. Valid values: `accept`, `deny`.
func (o DeviceaccesslistOutput) DefaultAction() pulumi.StringOutput {
	return o.ApplyT(func(v *Deviceaccesslist) pulumi.StringOutput { return v.DefaultAction }).(pulumi.StringOutput)
}

// Device list. The structure of `deviceList` block is documented below.
func (o DeviceaccesslistOutput) DeviceLists() DeviceaccesslistDeviceListArrayOutput {
	return o.ApplyT(func(v *Deviceaccesslist) DeviceaccesslistDeviceListArrayOutput { return v.DeviceLists }).(DeviceaccesslistDeviceListArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o DeviceaccesslistOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deviceaccesslist) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o DeviceaccesslistOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deviceaccesslist) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Device access list name.
func (o DeviceaccesslistOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Deviceaccesslist) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DeviceaccesslistOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Deviceaccesslist) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type DeviceaccesslistArrayOutput struct{ *pulumi.OutputState }

func (DeviceaccesslistArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Deviceaccesslist)(nil)).Elem()
}

func (o DeviceaccesslistArrayOutput) ToDeviceaccesslistArrayOutput() DeviceaccesslistArrayOutput {
	return o
}

func (o DeviceaccesslistArrayOutput) ToDeviceaccesslistArrayOutputWithContext(ctx context.Context) DeviceaccesslistArrayOutput {
	return o
}

func (o DeviceaccesslistArrayOutput) Index(i pulumi.IntInput) DeviceaccesslistOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Deviceaccesslist {
		return vs[0].([]*Deviceaccesslist)[vs[1].(int)]
	}).(DeviceaccesslistOutput)
}

type DeviceaccesslistMapOutput struct{ *pulumi.OutputState }

func (DeviceaccesslistMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Deviceaccesslist)(nil)).Elem()
}

func (o DeviceaccesslistMapOutput) ToDeviceaccesslistMapOutput() DeviceaccesslistMapOutput {
	return o
}

func (o DeviceaccesslistMapOutput) ToDeviceaccesslistMapOutputWithContext(ctx context.Context) DeviceaccesslistMapOutput {
	return o
}

func (o DeviceaccesslistMapOutput) MapIndex(k pulumi.StringInput) DeviceaccesslistOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Deviceaccesslist {
		return vs[0].(map[string]*Deviceaccesslist)[vs[1].(string)]
	}).(DeviceaccesslistOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceaccesslistInput)(nil)).Elem(), &Deviceaccesslist{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceaccesslistArrayInput)(nil)).Elem(), DeviceaccesslistArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceaccesslistMapInput)(nil)).Elem(), DeviceaccesslistMap{})
	pulumi.RegisterOutputType(DeviceaccesslistOutput{})
	pulumi.RegisterOutputType(DeviceaccesslistArrayOutput{})
	pulumi.RegisterOutputType(DeviceaccesslistMapOutput{})
}
