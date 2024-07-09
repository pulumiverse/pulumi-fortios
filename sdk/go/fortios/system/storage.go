// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure logical storage.
//
// ## Import
//
// System Storage can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/storage:Storage labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/storage:Storage labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Storage struct {
	pulumi.CustomResourceState

	// Partition device.
	Device pulumi.StringOutput `pulumi:"device"`
	// The physical status of current media. Valid values: `enable`, `disable`, `fail`.
	MediaStatus pulumi.StringOutput `pulumi:"mediaStatus"`
	// Storage name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Set storage order.
	Order pulumi.IntOutput `pulumi:"order"`
	// Label of underlying partition.
	Partition pulumi.StringOutput `pulumi:"partition"`
	// Partition size.
	Size pulumi.IntOutput `pulumi:"size"`
	// Enable/disable storage. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Use hard disk for logging or WAN Optimization (default = log).
	Usage pulumi.StringOutput `pulumi:"usage"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
	// WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
	WanoptMode pulumi.StringOutput `pulumi:"wanoptMode"`
}

// NewStorage registers a new resource with the given unique name, arguments, and options.
func NewStorage(ctx *pulumi.Context,
	name string, args *StorageArgs, opts ...pulumi.ResourceOption) (*Storage, error) {
	if args == nil {
		args = &StorageArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Storage
	err := ctx.RegisterResource("fortios:system/storage:Storage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorage gets an existing Storage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageState, opts ...pulumi.ResourceOption) (*Storage, error) {
	var resource Storage
	err := ctx.ReadResource("fortios:system/storage:Storage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Storage resources.
type storageState struct {
	// Partition device.
	Device *string `pulumi:"device"`
	// The physical status of current media. Valid values: `enable`, `disable`, `fail`.
	MediaStatus *string `pulumi:"mediaStatus"`
	// Storage name.
	Name *string `pulumi:"name"`
	// Set storage order.
	Order *int `pulumi:"order"`
	// Label of underlying partition.
	Partition *string `pulumi:"partition"`
	// Partition size.
	Size *int `pulumi:"size"`
	// Enable/disable storage. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Use hard disk for logging or WAN Optimization (default = log).
	Usage *string `pulumi:"usage"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
	WanoptMode *string `pulumi:"wanoptMode"`
}

type StorageState struct {
	// Partition device.
	Device pulumi.StringPtrInput
	// The physical status of current media. Valid values: `enable`, `disable`, `fail`.
	MediaStatus pulumi.StringPtrInput
	// Storage name.
	Name pulumi.StringPtrInput
	// Set storage order.
	Order pulumi.IntPtrInput
	// Label of underlying partition.
	Partition pulumi.StringPtrInput
	// Partition size.
	Size pulumi.IntPtrInput
	// Enable/disable storage. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Use hard disk for logging or WAN Optimization (default = log).
	Usage pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
	WanoptMode pulumi.StringPtrInput
}

func (StorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageState)(nil)).Elem()
}

type storageArgs struct {
	// Partition device.
	Device *string `pulumi:"device"`
	// The physical status of current media. Valid values: `enable`, `disable`, `fail`.
	MediaStatus *string `pulumi:"mediaStatus"`
	// Storage name.
	Name *string `pulumi:"name"`
	// Set storage order.
	Order *int `pulumi:"order"`
	// Label of underlying partition.
	Partition *string `pulumi:"partition"`
	// Partition size.
	Size *int `pulumi:"size"`
	// Enable/disable storage. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Use hard disk for logging or WAN Optimization (default = log).
	Usage *string `pulumi:"usage"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
	WanoptMode *string `pulumi:"wanoptMode"`
}

// The set of arguments for constructing a Storage resource.
type StorageArgs struct {
	// Partition device.
	Device pulumi.StringPtrInput
	// The physical status of current media. Valid values: `enable`, `disable`, `fail`.
	MediaStatus pulumi.StringPtrInput
	// Storage name.
	Name pulumi.StringPtrInput
	// Set storage order.
	Order pulumi.IntPtrInput
	// Label of underlying partition.
	Partition pulumi.StringPtrInput
	// Partition size.
	Size pulumi.IntPtrInput
	// Enable/disable storage. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Use hard disk for logging or WAN Optimization (default = log).
	Usage pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
	WanoptMode pulumi.StringPtrInput
}

func (StorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageArgs)(nil)).Elem()
}

type StorageInput interface {
	pulumi.Input

	ToStorageOutput() StorageOutput
	ToStorageOutputWithContext(ctx context.Context) StorageOutput
}

func (*Storage) ElementType() reflect.Type {
	return reflect.TypeOf((**Storage)(nil)).Elem()
}

func (i *Storage) ToStorageOutput() StorageOutput {
	return i.ToStorageOutputWithContext(context.Background())
}

func (i *Storage) ToStorageOutputWithContext(ctx context.Context) StorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageOutput)
}

// StorageArrayInput is an input type that accepts StorageArray and StorageArrayOutput values.
// You can construct a concrete instance of `StorageArrayInput` via:
//
//	StorageArray{ StorageArgs{...} }
type StorageArrayInput interface {
	pulumi.Input

	ToStorageArrayOutput() StorageArrayOutput
	ToStorageArrayOutputWithContext(context.Context) StorageArrayOutput
}

type StorageArray []StorageInput

func (StorageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Storage)(nil)).Elem()
}

func (i StorageArray) ToStorageArrayOutput() StorageArrayOutput {
	return i.ToStorageArrayOutputWithContext(context.Background())
}

func (i StorageArray) ToStorageArrayOutputWithContext(ctx context.Context) StorageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageArrayOutput)
}

// StorageMapInput is an input type that accepts StorageMap and StorageMapOutput values.
// You can construct a concrete instance of `StorageMapInput` via:
//
//	StorageMap{ "key": StorageArgs{...} }
type StorageMapInput interface {
	pulumi.Input

	ToStorageMapOutput() StorageMapOutput
	ToStorageMapOutputWithContext(context.Context) StorageMapOutput
}

type StorageMap map[string]StorageInput

func (StorageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Storage)(nil)).Elem()
}

func (i StorageMap) ToStorageMapOutput() StorageMapOutput {
	return i.ToStorageMapOutputWithContext(context.Background())
}

func (i StorageMap) ToStorageMapOutputWithContext(ctx context.Context) StorageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageMapOutput)
}

type StorageOutput struct{ *pulumi.OutputState }

func (StorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Storage)(nil)).Elem()
}

func (o StorageOutput) ToStorageOutput() StorageOutput {
	return o
}

func (o StorageOutput) ToStorageOutputWithContext(ctx context.Context) StorageOutput {
	return o
}

// Partition device.
func (o StorageOutput) Device() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.Device }).(pulumi.StringOutput)
}

// The physical status of current media. Valid values: `enable`, `disable`, `fail`.
func (o StorageOutput) MediaStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.MediaStatus }).(pulumi.StringOutput)
}

// Storage name.
func (o StorageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Set storage order.
func (o StorageOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v *Storage) pulumi.IntOutput { return v.Order }).(pulumi.IntOutput)
}

// Label of underlying partition.
func (o StorageOutput) Partition() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.Partition }).(pulumi.StringOutput)
}

// Partition size.
func (o StorageOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *Storage) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// Enable/disable storage. Valid values: `enable`, `disable`.
func (o StorageOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Use hard disk for logging or WAN Optimization (default = log).
func (o StorageOutput) Usage() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.Usage }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o StorageOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

// WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
func (o StorageOutput) WanoptMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.WanoptMode }).(pulumi.StringOutput)
}

type StorageArrayOutput struct{ *pulumi.OutputState }

func (StorageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Storage)(nil)).Elem()
}

func (o StorageArrayOutput) ToStorageArrayOutput() StorageArrayOutput {
	return o
}

func (o StorageArrayOutput) ToStorageArrayOutputWithContext(ctx context.Context) StorageArrayOutput {
	return o
}

func (o StorageArrayOutput) Index(i pulumi.IntInput) StorageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Storage {
		return vs[0].([]*Storage)[vs[1].(int)]
	}).(StorageOutput)
}

type StorageMapOutput struct{ *pulumi.OutputState }

func (StorageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Storage)(nil)).Elem()
}

func (o StorageMapOutput) ToStorageMapOutput() StorageMapOutput {
	return o
}

func (o StorageMapOutput) ToStorageMapOutputWithContext(ctx context.Context) StorageMapOutput {
	return o
}

func (o StorageMapOutput) MapIndex(k pulumi.StringInput) StorageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Storage {
		return vs[0].(map[string]*Storage)[vs[1].(string)]
	}).(StorageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageInput)(nil)).Elem(), &Storage{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageArrayInput)(nil)).Elem(), StorageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageMapInput)(nil)).Elem(), StorageMap{})
	pulumi.RegisterOutputType(StorageOutput{})
	pulumi.RegisterOutputType(StorageArrayOutput{})
	pulumi.RegisterOutputType(StorageMapOutput{})
}
