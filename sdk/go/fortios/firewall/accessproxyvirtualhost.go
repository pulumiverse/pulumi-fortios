// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure Access Proxy virtual hosts. Applies to FortiOS Version `>= 7.0.1`.
//
// ## Import
//
// Firewall AccessProxyVirtualHost can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/accessproxyvirtualhost:Accessproxyvirtualhost labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/accessproxyvirtualhost:Accessproxyvirtualhost labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Accessproxyvirtualhost struct {
	pulumi.CustomResourceState

	// The host name.
	Host pulumi.StringOutput `pulumi:"host"`
	// Type of host pattern. Valid values: `sub-string`, `wildcard`.
	HostType pulumi.StringOutput `pulumi:"hostType"`
	// Virtual host name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Access-proxy-virtual-host replacement message override group.
	ReplacemsgGroup pulumi.StringOutput `pulumi:"replacemsgGroup"`
	// SSL certificate for this host.
	SslCertificate pulumi.StringOutput `pulumi:"sslCertificate"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewAccessproxyvirtualhost registers a new resource with the given unique name, arguments, and options.
func NewAccessproxyvirtualhost(ctx *pulumi.Context,
	name string, args *AccessproxyvirtualhostArgs, opts ...pulumi.ResourceOption) (*Accessproxyvirtualhost, error) {
	if args == nil {
		args = &AccessproxyvirtualhostArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Accessproxyvirtualhost
	err := ctx.RegisterResource("fortios:firewall/accessproxyvirtualhost:Accessproxyvirtualhost", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessproxyvirtualhost gets an existing Accessproxyvirtualhost resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessproxyvirtualhost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessproxyvirtualhostState, opts ...pulumi.ResourceOption) (*Accessproxyvirtualhost, error) {
	var resource Accessproxyvirtualhost
	err := ctx.ReadResource("fortios:firewall/accessproxyvirtualhost:Accessproxyvirtualhost", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Accessproxyvirtualhost resources.
type accessproxyvirtualhostState struct {
	// The host name.
	Host *string `pulumi:"host"`
	// Type of host pattern. Valid values: `sub-string`, `wildcard`.
	HostType *string `pulumi:"hostType"`
	// Virtual host name.
	Name *string `pulumi:"name"`
	// Access-proxy-virtual-host replacement message override group.
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// SSL certificate for this host.
	SslCertificate *string `pulumi:"sslCertificate"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type AccessproxyvirtualhostState struct {
	// The host name.
	Host pulumi.StringPtrInput
	// Type of host pattern. Valid values: `sub-string`, `wildcard`.
	HostType pulumi.StringPtrInput
	// Virtual host name.
	Name pulumi.StringPtrInput
	// Access-proxy-virtual-host replacement message override group.
	ReplacemsgGroup pulumi.StringPtrInput
	// SSL certificate for this host.
	SslCertificate pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AccessproxyvirtualhostState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessproxyvirtualhostState)(nil)).Elem()
}

type accessproxyvirtualhostArgs struct {
	// The host name.
	Host *string `pulumi:"host"`
	// Type of host pattern. Valid values: `sub-string`, `wildcard`.
	HostType *string `pulumi:"hostType"`
	// Virtual host name.
	Name *string `pulumi:"name"`
	// Access-proxy-virtual-host replacement message override group.
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// SSL certificate for this host.
	SslCertificate *string `pulumi:"sslCertificate"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Accessproxyvirtualhost resource.
type AccessproxyvirtualhostArgs struct {
	// The host name.
	Host pulumi.StringPtrInput
	// Type of host pattern. Valid values: `sub-string`, `wildcard`.
	HostType pulumi.StringPtrInput
	// Virtual host name.
	Name pulumi.StringPtrInput
	// Access-proxy-virtual-host replacement message override group.
	ReplacemsgGroup pulumi.StringPtrInput
	// SSL certificate for this host.
	SslCertificate pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AccessproxyvirtualhostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessproxyvirtualhostArgs)(nil)).Elem()
}

type AccessproxyvirtualhostInput interface {
	pulumi.Input

	ToAccessproxyvirtualhostOutput() AccessproxyvirtualhostOutput
	ToAccessproxyvirtualhostOutputWithContext(ctx context.Context) AccessproxyvirtualhostOutput
}

func (*Accessproxyvirtualhost) ElementType() reflect.Type {
	return reflect.TypeOf((**Accessproxyvirtualhost)(nil)).Elem()
}

func (i *Accessproxyvirtualhost) ToAccessproxyvirtualhostOutput() AccessproxyvirtualhostOutput {
	return i.ToAccessproxyvirtualhostOutputWithContext(context.Background())
}

func (i *Accessproxyvirtualhost) ToAccessproxyvirtualhostOutputWithContext(ctx context.Context) AccessproxyvirtualhostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessproxyvirtualhostOutput)
}

// AccessproxyvirtualhostArrayInput is an input type that accepts AccessproxyvirtualhostArray and AccessproxyvirtualhostArrayOutput values.
// You can construct a concrete instance of `AccessproxyvirtualhostArrayInput` via:
//
//	AccessproxyvirtualhostArray{ AccessproxyvirtualhostArgs{...} }
type AccessproxyvirtualhostArrayInput interface {
	pulumi.Input

	ToAccessproxyvirtualhostArrayOutput() AccessproxyvirtualhostArrayOutput
	ToAccessproxyvirtualhostArrayOutputWithContext(context.Context) AccessproxyvirtualhostArrayOutput
}

type AccessproxyvirtualhostArray []AccessproxyvirtualhostInput

func (AccessproxyvirtualhostArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Accessproxyvirtualhost)(nil)).Elem()
}

func (i AccessproxyvirtualhostArray) ToAccessproxyvirtualhostArrayOutput() AccessproxyvirtualhostArrayOutput {
	return i.ToAccessproxyvirtualhostArrayOutputWithContext(context.Background())
}

func (i AccessproxyvirtualhostArray) ToAccessproxyvirtualhostArrayOutputWithContext(ctx context.Context) AccessproxyvirtualhostArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessproxyvirtualhostArrayOutput)
}

// AccessproxyvirtualhostMapInput is an input type that accepts AccessproxyvirtualhostMap and AccessproxyvirtualhostMapOutput values.
// You can construct a concrete instance of `AccessproxyvirtualhostMapInput` via:
//
//	AccessproxyvirtualhostMap{ "key": AccessproxyvirtualhostArgs{...} }
type AccessproxyvirtualhostMapInput interface {
	pulumi.Input

	ToAccessproxyvirtualhostMapOutput() AccessproxyvirtualhostMapOutput
	ToAccessproxyvirtualhostMapOutputWithContext(context.Context) AccessproxyvirtualhostMapOutput
}

type AccessproxyvirtualhostMap map[string]AccessproxyvirtualhostInput

func (AccessproxyvirtualhostMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Accessproxyvirtualhost)(nil)).Elem()
}

func (i AccessproxyvirtualhostMap) ToAccessproxyvirtualhostMapOutput() AccessproxyvirtualhostMapOutput {
	return i.ToAccessproxyvirtualhostMapOutputWithContext(context.Background())
}

func (i AccessproxyvirtualhostMap) ToAccessproxyvirtualhostMapOutputWithContext(ctx context.Context) AccessproxyvirtualhostMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessproxyvirtualhostMapOutput)
}

type AccessproxyvirtualhostOutput struct{ *pulumi.OutputState }

func (AccessproxyvirtualhostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Accessproxyvirtualhost)(nil)).Elem()
}

func (o AccessproxyvirtualhostOutput) ToAccessproxyvirtualhostOutput() AccessproxyvirtualhostOutput {
	return o
}

func (o AccessproxyvirtualhostOutput) ToAccessproxyvirtualhostOutputWithContext(ctx context.Context) AccessproxyvirtualhostOutput {
	return o
}

// The host name.
func (o AccessproxyvirtualhostOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *Accessproxyvirtualhost) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// Type of host pattern. Valid values: `sub-string`, `wildcard`.
func (o AccessproxyvirtualhostOutput) HostType() pulumi.StringOutput {
	return o.ApplyT(func(v *Accessproxyvirtualhost) pulumi.StringOutput { return v.HostType }).(pulumi.StringOutput)
}

// Virtual host name.
func (o AccessproxyvirtualhostOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Accessproxyvirtualhost) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Access-proxy-virtual-host replacement message override group.
func (o AccessproxyvirtualhostOutput) ReplacemsgGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Accessproxyvirtualhost) pulumi.StringOutput { return v.ReplacemsgGroup }).(pulumi.StringOutput)
}

// SSL certificate for this host.
func (o AccessproxyvirtualhostOutput) SslCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v *Accessproxyvirtualhost) pulumi.StringOutput { return v.SslCertificate }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o AccessproxyvirtualhostOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Accessproxyvirtualhost) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type AccessproxyvirtualhostArrayOutput struct{ *pulumi.OutputState }

func (AccessproxyvirtualhostArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Accessproxyvirtualhost)(nil)).Elem()
}

func (o AccessproxyvirtualhostArrayOutput) ToAccessproxyvirtualhostArrayOutput() AccessproxyvirtualhostArrayOutput {
	return o
}

func (o AccessproxyvirtualhostArrayOutput) ToAccessproxyvirtualhostArrayOutputWithContext(ctx context.Context) AccessproxyvirtualhostArrayOutput {
	return o
}

func (o AccessproxyvirtualhostArrayOutput) Index(i pulumi.IntInput) AccessproxyvirtualhostOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Accessproxyvirtualhost {
		return vs[0].([]*Accessproxyvirtualhost)[vs[1].(int)]
	}).(AccessproxyvirtualhostOutput)
}

type AccessproxyvirtualhostMapOutput struct{ *pulumi.OutputState }

func (AccessproxyvirtualhostMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Accessproxyvirtualhost)(nil)).Elem()
}

func (o AccessproxyvirtualhostMapOutput) ToAccessproxyvirtualhostMapOutput() AccessproxyvirtualhostMapOutput {
	return o
}

func (o AccessproxyvirtualhostMapOutput) ToAccessproxyvirtualhostMapOutputWithContext(ctx context.Context) AccessproxyvirtualhostMapOutput {
	return o
}

func (o AccessproxyvirtualhostMapOutput) MapIndex(k pulumi.StringInput) AccessproxyvirtualhostOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Accessproxyvirtualhost {
		return vs[0].(map[string]*Accessproxyvirtualhost)[vs[1].(string)]
	}).(AccessproxyvirtualhostOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessproxyvirtualhostInput)(nil)).Elem(), &Accessproxyvirtualhost{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessproxyvirtualhostArrayInput)(nil)).Elem(), AccessproxyvirtualhostArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessproxyvirtualhostMapInput)(nil)).Elem(), AccessproxyvirtualhostMap{})
	pulumi.RegisterOutputType(AccessproxyvirtualhostOutput{})
	pulumi.RegisterOutputType(AccessproxyvirtualhostArrayOutput{})
	pulumi.RegisterOutputType(AccessproxyvirtualhostMapOutput{})
}
