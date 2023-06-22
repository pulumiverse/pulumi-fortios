// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPS URL filter DNS servers.
//
// ## Import
//
// # System IpsUrlfilterDns can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemIpsurlfilterdns:SystemIpsurlfilterdns labelname {{address}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemIpsurlfilterdns:SystemIpsurlfilterdns labelname {{address}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemIpsurlfilterdns struct {
	pulumi.CustomResourceState

	// DNS server IP address.
	Address pulumi.StringOutput `pulumi:"address"`
	// Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
	Ipv6Capability pulumi.StringOutput `pulumi:"ipv6Capability"`
	// Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemIpsurlfilterdns registers a new resource with the given unique name, arguments, and options.
func NewSystemIpsurlfilterdns(ctx *pulumi.Context,
	name string, args *SystemIpsurlfilterdnsArgs, opts ...pulumi.ResourceOption) (*SystemIpsurlfilterdns, error) {
	if args == nil {
		args = &SystemIpsurlfilterdnsArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemIpsurlfilterdns
	err := ctx.RegisterResource("fortios:index/systemIpsurlfilterdns:SystemIpsurlfilterdns", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemIpsurlfilterdns gets an existing SystemIpsurlfilterdns resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemIpsurlfilterdns(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemIpsurlfilterdnsState, opts ...pulumi.ResourceOption) (*SystemIpsurlfilterdns, error) {
	var resource SystemIpsurlfilterdns
	err := ctx.ReadResource("fortios:index/systemIpsurlfilterdns:SystemIpsurlfilterdns", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemIpsurlfilterdns resources.
type systemIpsurlfilterdnsState struct {
	// DNS server IP address.
	Address *string `pulumi:"address"`
	// Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
	Ipv6Capability *string `pulumi:"ipv6Capability"`
	// Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemIpsurlfilterdnsState struct {
	// DNS server IP address.
	Address pulumi.StringPtrInput
	// Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
	Ipv6Capability pulumi.StringPtrInput
	// Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemIpsurlfilterdnsState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemIpsurlfilterdnsState)(nil)).Elem()
}

type systemIpsurlfilterdnsArgs struct {
	// DNS server IP address.
	Address *string `pulumi:"address"`
	// Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
	Ipv6Capability *string `pulumi:"ipv6Capability"`
	// Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemIpsurlfilterdns resource.
type SystemIpsurlfilterdnsArgs struct {
	// DNS server IP address.
	Address pulumi.StringPtrInput
	// Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
	Ipv6Capability pulumi.StringPtrInput
	// Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemIpsurlfilterdnsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemIpsurlfilterdnsArgs)(nil)).Elem()
}

type SystemIpsurlfilterdnsInput interface {
	pulumi.Input

	ToSystemIpsurlfilterdnsOutput() SystemIpsurlfilterdnsOutput
	ToSystemIpsurlfilterdnsOutputWithContext(ctx context.Context) SystemIpsurlfilterdnsOutput
}

func (*SystemIpsurlfilterdns) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemIpsurlfilterdns)(nil)).Elem()
}

func (i *SystemIpsurlfilterdns) ToSystemIpsurlfilterdnsOutput() SystemIpsurlfilterdnsOutput {
	return i.ToSystemIpsurlfilterdnsOutputWithContext(context.Background())
}

func (i *SystemIpsurlfilterdns) ToSystemIpsurlfilterdnsOutputWithContext(ctx context.Context) SystemIpsurlfilterdnsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemIpsurlfilterdnsOutput)
}

// SystemIpsurlfilterdnsArrayInput is an input type that accepts SystemIpsurlfilterdnsArray and SystemIpsurlfilterdnsArrayOutput values.
// You can construct a concrete instance of `SystemIpsurlfilterdnsArrayInput` via:
//
//	SystemIpsurlfilterdnsArray{ SystemIpsurlfilterdnsArgs{...} }
type SystemIpsurlfilterdnsArrayInput interface {
	pulumi.Input

	ToSystemIpsurlfilterdnsArrayOutput() SystemIpsurlfilterdnsArrayOutput
	ToSystemIpsurlfilterdnsArrayOutputWithContext(context.Context) SystemIpsurlfilterdnsArrayOutput
}

type SystemIpsurlfilterdnsArray []SystemIpsurlfilterdnsInput

func (SystemIpsurlfilterdnsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemIpsurlfilterdns)(nil)).Elem()
}

func (i SystemIpsurlfilterdnsArray) ToSystemIpsurlfilterdnsArrayOutput() SystemIpsurlfilterdnsArrayOutput {
	return i.ToSystemIpsurlfilterdnsArrayOutputWithContext(context.Background())
}

func (i SystemIpsurlfilterdnsArray) ToSystemIpsurlfilterdnsArrayOutputWithContext(ctx context.Context) SystemIpsurlfilterdnsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemIpsurlfilterdnsArrayOutput)
}

// SystemIpsurlfilterdnsMapInput is an input type that accepts SystemIpsurlfilterdnsMap and SystemIpsurlfilterdnsMapOutput values.
// You can construct a concrete instance of `SystemIpsurlfilterdnsMapInput` via:
//
//	SystemIpsurlfilterdnsMap{ "key": SystemIpsurlfilterdnsArgs{...} }
type SystemIpsurlfilterdnsMapInput interface {
	pulumi.Input

	ToSystemIpsurlfilterdnsMapOutput() SystemIpsurlfilterdnsMapOutput
	ToSystemIpsurlfilterdnsMapOutputWithContext(context.Context) SystemIpsurlfilterdnsMapOutput
}

type SystemIpsurlfilterdnsMap map[string]SystemIpsurlfilterdnsInput

func (SystemIpsurlfilterdnsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemIpsurlfilterdns)(nil)).Elem()
}

func (i SystemIpsurlfilterdnsMap) ToSystemIpsurlfilterdnsMapOutput() SystemIpsurlfilterdnsMapOutput {
	return i.ToSystemIpsurlfilterdnsMapOutputWithContext(context.Background())
}

func (i SystemIpsurlfilterdnsMap) ToSystemIpsurlfilterdnsMapOutputWithContext(ctx context.Context) SystemIpsurlfilterdnsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemIpsurlfilterdnsMapOutput)
}

type SystemIpsurlfilterdnsOutput struct{ *pulumi.OutputState }

func (SystemIpsurlfilterdnsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemIpsurlfilterdns)(nil)).Elem()
}

func (o SystemIpsurlfilterdnsOutput) ToSystemIpsurlfilterdnsOutput() SystemIpsurlfilterdnsOutput {
	return o
}

func (o SystemIpsurlfilterdnsOutput) ToSystemIpsurlfilterdnsOutputWithContext(ctx context.Context) SystemIpsurlfilterdnsOutput {
	return o
}

// DNS server IP address.
func (o SystemIpsurlfilterdnsOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemIpsurlfilterdns) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
func (o SystemIpsurlfilterdnsOutput) Ipv6Capability() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemIpsurlfilterdns) pulumi.StringOutput { return v.Ipv6Capability }).(pulumi.StringOutput)
}

// Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
func (o SystemIpsurlfilterdnsOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemIpsurlfilterdns) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemIpsurlfilterdnsOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemIpsurlfilterdns) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemIpsurlfilterdnsArrayOutput struct{ *pulumi.OutputState }

func (SystemIpsurlfilterdnsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemIpsurlfilterdns)(nil)).Elem()
}

func (o SystemIpsurlfilterdnsArrayOutput) ToSystemIpsurlfilterdnsArrayOutput() SystemIpsurlfilterdnsArrayOutput {
	return o
}

func (o SystemIpsurlfilterdnsArrayOutput) ToSystemIpsurlfilterdnsArrayOutputWithContext(ctx context.Context) SystemIpsurlfilterdnsArrayOutput {
	return o
}

func (o SystemIpsurlfilterdnsArrayOutput) Index(i pulumi.IntInput) SystemIpsurlfilterdnsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemIpsurlfilterdns {
		return vs[0].([]*SystemIpsurlfilterdns)[vs[1].(int)]
	}).(SystemIpsurlfilterdnsOutput)
}

type SystemIpsurlfilterdnsMapOutput struct{ *pulumi.OutputState }

func (SystemIpsurlfilterdnsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemIpsurlfilterdns)(nil)).Elem()
}

func (o SystemIpsurlfilterdnsMapOutput) ToSystemIpsurlfilterdnsMapOutput() SystemIpsurlfilterdnsMapOutput {
	return o
}

func (o SystemIpsurlfilterdnsMapOutput) ToSystemIpsurlfilterdnsMapOutputWithContext(ctx context.Context) SystemIpsurlfilterdnsMapOutput {
	return o
}

func (o SystemIpsurlfilterdnsMapOutput) MapIndex(k pulumi.StringInput) SystemIpsurlfilterdnsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemIpsurlfilterdns {
		return vs[0].(map[string]*SystemIpsurlfilterdns)[vs[1].(string)]
	}).(SystemIpsurlfilterdnsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemIpsurlfilterdnsInput)(nil)).Elem(), &SystemIpsurlfilterdns{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemIpsurlfilterdnsArrayInput)(nil)).Elem(), SystemIpsurlfilterdnsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemIpsurlfilterdnsMapInput)(nil)).Elem(), SystemIpsurlfilterdnsMap{})
	pulumi.RegisterOutputType(SystemIpsurlfilterdnsOutput{})
	pulumi.RegisterOutputType(SystemIpsurlfilterdnsArrayOutput{})
	pulumi.RegisterOutputType(SystemIpsurlfilterdnsMapOutput{})
}
