// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure dedicated management.
//
// ## Import
//
// System DedicatedMgmt can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/dedicatedmgmt:Dedicatedmgmt labelname SystemDedicatedMgmt
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/dedicatedmgmt:Dedicatedmgmt labelname SystemDedicatedMgmt
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Dedicatedmgmt struct {
	pulumi.CustomResourceState

	// Default gateway for dedicated management interface.
	DefaultGateway pulumi.StringOutput `pulumi:"defaultGateway"`
	// DHCP end IP for dedicated management.
	DhcpEndIp pulumi.StringOutput `pulumi:"dhcpEndIp"`
	// DHCP netmask.
	DhcpNetmask pulumi.StringOutput `pulumi:"dhcpNetmask"`
	// Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
	DhcpServer pulumi.StringOutput `pulumi:"dhcpServer"`
	// DHCP start IP for dedicated management.
	DhcpStartIp pulumi.StringOutput `pulumi:"dhcpStartIp"`
	// Dedicated management interface.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Enable/disable dedicated management. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDedicatedmgmt registers a new resource with the given unique name, arguments, and options.
func NewDedicatedmgmt(ctx *pulumi.Context,
	name string, args *DedicatedmgmtArgs, opts ...pulumi.ResourceOption) (*Dedicatedmgmt, error) {
	if args == nil {
		args = &DedicatedmgmtArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Dedicatedmgmt
	err := ctx.RegisterResource("fortios:system/dedicatedmgmt:Dedicatedmgmt", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedmgmt gets an existing Dedicatedmgmt resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedmgmt(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedmgmtState, opts ...pulumi.ResourceOption) (*Dedicatedmgmt, error) {
	var resource Dedicatedmgmt
	err := ctx.ReadResource("fortios:system/dedicatedmgmt:Dedicatedmgmt", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dedicatedmgmt resources.
type dedicatedmgmtState struct {
	// Default gateway for dedicated management interface.
	DefaultGateway *string `pulumi:"defaultGateway"`
	// DHCP end IP for dedicated management.
	DhcpEndIp *string `pulumi:"dhcpEndIp"`
	// DHCP netmask.
	DhcpNetmask *string `pulumi:"dhcpNetmask"`
	// Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
	DhcpServer *string `pulumi:"dhcpServer"`
	// DHCP start IP for dedicated management.
	DhcpStartIp *string `pulumi:"dhcpStartIp"`
	// Dedicated management interface.
	Interface *string `pulumi:"interface"`
	// Enable/disable dedicated management. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DedicatedmgmtState struct {
	// Default gateway for dedicated management interface.
	DefaultGateway pulumi.StringPtrInput
	// DHCP end IP for dedicated management.
	DhcpEndIp pulumi.StringPtrInput
	// DHCP netmask.
	DhcpNetmask pulumi.StringPtrInput
	// Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
	DhcpServer pulumi.StringPtrInput
	// DHCP start IP for dedicated management.
	DhcpStartIp pulumi.StringPtrInput
	// Dedicated management interface.
	Interface pulumi.StringPtrInput
	// Enable/disable dedicated management. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DedicatedmgmtState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedmgmtState)(nil)).Elem()
}

type dedicatedmgmtArgs struct {
	// Default gateway for dedicated management interface.
	DefaultGateway *string `pulumi:"defaultGateway"`
	// DHCP end IP for dedicated management.
	DhcpEndIp *string `pulumi:"dhcpEndIp"`
	// DHCP netmask.
	DhcpNetmask *string `pulumi:"dhcpNetmask"`
	// Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
	DhcpServer *string `pulumi:"dhcpServer"`
	// DHCP start IP for dedicated management.
	DhcpStartIp *string `pulumi:"dhcpStartIp"`
	// Dedicated management interface.
	Interface *string `pulumi:"interface"`
	// Enable/disable dedicated management. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Dedicatedmgmt resource.
type DedicatedmgmtArgs struct {
	// Default gateway for dedicated management interface.
	DefaultGateway pulumi.StringPtrInput
	// DHCP end IP for dedicated management.
	DhcpEndIp pulumi.StringPtrInput
	// DHCP netmask.
	DhcpNetmask pulumi.StringPtrInput
	// Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
	DhcpServer pulumi.StringPtrInput
	// DHCP start IP for dedicated management.
	DhcpStartIp pulumi.StringPtrInput
	// Dedicated management interface.
	Interface pulumi.StringPtrInput
	// Enable/disable dedicated management. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DedicatedmgmtArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedmgmtArgs)(nil)).Elem()
}

type DedicatedmgmtInput interface {
	pulumi.Input

	ToDedicatedmgmtOutput() DedicatedmgmtOutput
	ToDedicatedmgmtOutputWithContext(ctx context.Context) DedicatedmgmtOutput
}

func (*Dedicatedmgmt) ElementType() reflect.Type {
	return reflect.TypeOf((**Dedicatedmgmt)(nil)).Elem()
}

func (i *Dedicatedmgmt) ToDedicatedmgmtOutput() DedicatedmgmtOutput {
	return i.ToDedicatedmgmtOutputWithContext(context.Background())
}

func (i *Dedicatedmgmt) ToDedicatedmgmtOutputWithContext(ctx context.Context) DedicatedmgmtOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedmgmtOutput)
}

// DedicatedmgmtArrayInput is an input type that accepts DedicatedmgmtArray and DedicatedmgmtArrayOutput values.
// You can construct a concrete instance of `DedicatedmgmtArrayInput` via:
//
//	DedicatedmgmtArray{ DedicatedmgmtArgs{...} }
type DedicatedmgmtArrayInput interface {
	pulumi.Input

	ToDedicatedmgmtArrayOutput() DedicatedmgmtArrayOutput
	ToDedicatedmgmtArrayOutputWithContext(context.Context) DedicatedmgmtArrayOutput
}

type DedicatedmgmtArray []DedicatedmgmtInput

func (DedicatedmgmtArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dedicatedmgmt)(nil)).Elem()
}

func (i DedicatedmgmtArray) ToDedicatedmgmtArrayOutput() DedicatedmgmtArrayOutput {
	return i.ToDedicatedmgmtArrayOutputWithContext(context.Background())
}

func (i DedicatedmgmtArray) ToDedicatedmgmtArrayOutputWithContext(ctx context.Context) DedicatedmgmtArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedmgmtArrayOutput)
}

// DedicatedmgmtMapInput is an input type that accepts DedicatedmgmtMap and DedicatedmgmtMapOutput values.
// You can construct a concrete instance of `DedicatedmgmtMapInput` via:
//
//	DedicatedmgmtMap{ "key": DedicatedmgmtArgs{...} }
type DedicatedmgmtMapInput interface {
	pulumi.Input

	ToDedicatedmgmtMapOutput() DedicatedmgmtMapOutput
	ToDedicatedmgmtMapOutputWithContext(context.Context) DedicatedmgmtMapOutput
}

type DedicatedmgmtMap map[string]DedicatedmgmtInput

func (DedicatedmgmtMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dedicatedmgmt)(nil)).Elem()
}

func (i DedicatedmgmtMap) ToDedicatedmgmtMapOutput() DedicatedmgmtMapOutput {
	return i.ToDedicatedmgmtMapOutputWithContext(context.Background())
}

func (i DedicatedmgmtMap) ToDedicatedmgmtMapOutputWithContext(ctx context.Context) DedicatedmgmtMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedmgmtMapOutput)
}

type DedicatedmgmtOutput struct{ *pulumi.OutputState }

func (DedicatedmgmtOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dedicatedmgmt)(nil)).Elem()
}

func (o DedicatedmgmtOutput) ToDedicatedmgmtOutput() DedicatedmgmtOutput {
	return o
}

func (o DedicatedmgmtOutput) ToDedicatedmgmtOutputWithContext(ctx context.Context) DedicatedmgmtOutput {
	return o
}

// Default gateway for dedicated management interface.
func (o DedicatedmgmtOutput) DefaultGateway() pulumi.StringOutput {
	return o.ApplyT(func(v *Dedicatedmgmt) pulumi.StringOutput { return v.DefaultGateway }).(pulumi.StringOutput)
}

// DHCP end IP for dedicated management.
func (o DedicatedmgmtOutput) DhcpEndIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Dedicatedmgmt) pulumi.StringOutput { return v.DhcpEndIp }).(pulumi.StringOutput)
}

// DHCP netmask.
func (o DedicatedmgmtOutput) DhcpNetmask() pulumi.StringOutput {
	return o.ApplyT(func(v *Dedicatedmgmt) pulumi.StringOutput { return v.DhcpNetmask }).(pulumi.StringOutput)
}

// Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
func (o DedicatedmgmtOutput) DhcpServer() pulumi.StringOutput {
	return o.ApplyT(func(v *Dedicatedmgmt) pulumi.StringOutput { return v.DhcpServer }).(pulumi.StringOutput)
}

// DHCP start IP for dedicated management.
func (o DedicatedmgmtOutput) DhcpStartIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Dedicatedmgmt) pulumi.StringOutput { return v.DhcpStartIp }).(pulumi.StringOutput)
}

// Dedicated management interface.
func (o DedicatedmgmtOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Dedicatedmgmt) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Enable/disable dedicated management. Valid values: `enable`, `disable`.
func (o DedicatedmgmtOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Dedicatedmgmt) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DedicatedmgmtOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dedicatedmgmt) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type DedicatedmgmtArrayOutput struct{ *pulumi.OutputState }

func (DedicatedmgmtArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dedicatedmgmt)(nil)).Elem()
}

func (o DedicatedmgmtArrayOutput) ToDedicatedmgmtArrayOutput() DedicatedmgmtArrayOutput {
	return o
}

func (o DedicatedmgmtArrayOutput) ToDedicatedmgmtArrayOutputWithContext(ctx context.Context) DedicatedmgmtArrayOutput {
	return o
}

func (o DedicatedmgmtArrayOutput) Index(i pulumi.IntInput) DedicatedmgmtOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Dedicatedmgmt {
		return vs[0].([]*Dedicatedmgmt)[vs[1].(int)]
	}).(DedicatedmgmtOutput)
}

type DedicatedmgmtMapOutput struct{ *pulumi.OutputState }

func (DedicatedmgmtMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dedicatedmgmt)(nil)).Elem()
}

func (o DedicatedmgmtMapOutput) ToDedicatedmgmtMapOutput() DedicatedmgmtMapOutput {
	return o
}

func (o DedicatedmgmtMapOutput) ToDedicatedmgmtMapOutputWithContext(ctx context.Context) DedicatedmgmtMapOutput {
	return o
}

func (o DedicatedmgmtMapOutput) MapIndex(k pulumi.StringInput) DedicatedmgmtOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Dedicatedmgmt {
		return vs[0].(map[string]*Dedicatedmgmt)[vs[1].(string)]
	}).(DedicatedmgmtOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedmgmtInput)(nil)).Elem(), &Dedicatedmgmt{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedmgmtArrayInput)(nil)).Elem(), DedicatedmgmtArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedmgmtMapInput)(nil)).Elem(), DedicatedmgmtMap{})
	pulumi.RegisterOutputType(DedicatedmgmtOutput{})
	pulumi.RegisterOutputType(DedicatedmgmtArrayOutput{})
	pulumi.RegisterOutputType(DedicatedmgmtMapOutput{})
}
