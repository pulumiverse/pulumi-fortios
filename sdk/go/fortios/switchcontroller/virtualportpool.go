// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure virtual pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/switchcontroller"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := switchcontroller.NewVirtualportpool(ctx, "trname", &switchcontroller.VirtualportpoolArgs{
//				Description: pulumi.String("virtualport"),
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
// SwitchController VirtualPortPool can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/virtualportpool:Virtualportpool labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/virtualportpool:Virtualportpool labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Virtualportpool struct {
	pulumi.CustomResourceState

	// Virtual switch pool description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Virtual switch pool name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewVirtualportpool registers a new resource with the given unique name, arguments, and options.
func NewVirtualportpool(ctx *pulumi.Context,
	name string, args *VirtualportpoolArgs, opts ...pulumi.ResourceOption) (*Virtualportpool, error) {
	if args == nil {
		args = &VirtualportpoolArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Virtualportpool
	err := ctx.RegisterResource("fortios:switchcontroller/virtualportpool:Virtualportpool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualportpool gets an existing Virtualportpool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualportpool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualportpoolState, opts ...pulumi.ResourceOption) (*Virtualportpool, error) {
	var resource Virtualportpool
	err := ctx.ReadResource("fortios:switchcontroller/virtualportpool:Virtualportpool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Virtualportpool resources.
type virtualportpoolState struct {
	// Virtual switch pool description.
	Description *string `pulumi:"description"`
	// Virtual switch pool name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type VirtualportpoolState struct {
	// Virtual switch pool description.
	Description pulumi.StringPtrInput
	// Virtual switch pool name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VirtualportpoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualportpoolState)(nil)).Elem()
}

type virtualportpoolArgs struct {
	// Virtual switch pool description.
	Description *string `pulumi:"description"`
	// Virtual switch pool name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Virtualportpool resource.
type VirtualportpoolArgs struct {
	// Virtual switch pool description.
	Description pulumi.StringPtrInput
	// Virtual switch pool name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VirtualportpoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualportpoolArgs)(nil)).Elem()
}

type VirtualportpoolInput interface {
	pulumi.Input

	ToVirtualportpoolOutput() VirtualportpoolOutput
	ToVirtualportpoolOutputWithContext(ctx context.Context) VirtualportpoolOutput
}

func (*Virtualportpool) ElementType() reflect.Type {
	return reflect.TypeOf((**Virtualportpool)(nil)).Elem()
}

func (i *Virtualportpool) ToVirtualportpoolOutput() VirtualportpoolOutput {
	return i.ToVirtualportpoolOutputWithContext(context.Background())
}

func (i *Virtualportpool) ToVirtualportpoolOutputWithContext(ctx context.Context) VirtualportpoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualportpoolOutput)
}

// VirtualportpoolArrayInput is an input type that accepts VirtualportpoolArray and VirtualportpoolArrayOutput values.
// You can construct a concrete instance of `VirtualportpoolArrayInput` via:
//
//	VirtualportpoolArray{ VirtualportpoolArgs{...} }
type VirtualportpoolArrayInput interface {
	pulumi.Input

	ToVirtualportpoolArrayOutput() VirtualportpoolArrayOutput
	ToVirtualportpoolArrayOutputWithContext(context.Context) VirtualportpoolArrayOutput
}

type VirtualportpoolArray []VirtualportpoolInput

func (VirtualportpoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Virtualportpool)(nil)).Elem()
}

func (i VirtualportpoolArray) ToVirtualportpoolArrayOutput() VirtualportpoolArrayOutput {
	return i.ToVirtualportpoolArrayOutputWithContext(context.Background())
}

func (i VirtualportpoolArray) ToVirtualportpoolArrayOutputWithContext(ctx context.Context) VirtualportpoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualportpoolArrayOutput)
}

// VirtualportpoolMapInput is an input type that accepts VirtualportpoolMap and VirtualportpoolMapOutput values.
// You can construct a concrete instance of `VirtualportpoolMapInput` via:
//
//	VirtualportpoolMap{ "key": VirtualportpoolArgs{...} }
type VirtualportpoolMapInput interface {
	pulumi.Input

	ToVirtualportpoolMapOutput() VirtualportpoolMapOutput
	ToVirtualportpoolMapOutputWithContext(context.Context) VirtualportpoolMapOutput
}

type VirtualportpoolMap map[string]VirtualportpoolInput

func (VirtualportpoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Virtualportpool)(nil)).Elem()
}

func (i VirtualportpoolMap) ToVirtualportpoolMapOutput() VirtualportpoolMapOutput {
	return i.ToVirtualportpoolMapOutputWithContext(context.Background())
}

func (i VirtualportpoolMap) ToVirtualportpoolMapOutputWithContext(ctx context.Context) VirtualportpoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualportpoolMapOutput)
}

type VirtualportpoolOutput struct{ *pulumi.OutputState }

func (VirtualportpoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Virtualportpool)(nil)).Elem()
}

func (o VirtualportpoolOutput) ToVirtualportpoolOutput() VirtualportpoolOutput {
	return o
}

func (o VirtualportpoolOutput) ToVirtualportpoolOutputWithContext(ctx context.Context) VirtualportpoolOutput {
	return o
}

// Virtual switch pool description.
func (o VirtualportpoolOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Virtualportpool) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Virtual switch pool name.
func (o VirtualportpoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Virtualportpool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o VirtualportpoolOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Virtualportpool) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type VirtualportpoolArrayOutput struct{ *pulumi.OutputState }

func (VirtualportpoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Virtualportpool)(nil)).Elem()
}

func (o VirtualportpoolArrayOutput) ToVirtualportpoolArrayOutput() VirtualportpoolArrayOutput {
	return o
}

func (o VirtualportpoolArrayOutput) ToVirtualportpoolArrayOutputWithContext(ctx context.Context) VirtualportpoolArrayOutput {
	return o
}

func (o VirtualportpoolArrayOutput) Index(i pulumi.IntInput) VirtualportpoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Virtualportpool {
		return vs[0].([]*Virtualportpool)[vs[1].(int)]
	}).(VirtualportpoolOutput)
}

type VirtualportpoolMapOutput struct{ *pulumi.OutputState }

func (VirtualportpoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Virtualportpool)(nil)).Elem()
}

func (o VirtualportpoolMapOutput) ToVirtualportpoolMapOutput() VirtualportpoolMapOutput {
	return o
}

func (o VirtualportpoolMapOutput) ToVirtualportpoolMapOutputWithContext(ctx context.Context) VirtualportpoolMapOutput {
	return o
}

func (o VirtualportpoolMapOutput) MapIndex(k pulumi.StringInput) VirtualportpoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Virtualportpool {
		return vs[0].(map[string]*Virtualportpool)[vs[1].(string)]
	}).(VirtualportpoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualportpoolInput)(nil)).Elem(), &Virtualportpool{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualportpoolArrayInput)(nil)).Elem(), VirtualportpoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualportpoolMapInput)(nil)).Elem(), VirtualportpoolMap{})
	pulumi.RegisterOutputType(VirtualportpoolOutput{})
	pulumi.RegisterOutputType(VirtualportpoolArrayOutput{})
	pulumi.RegisterOutputType(VirtualportpoolMapOutput{})
}
