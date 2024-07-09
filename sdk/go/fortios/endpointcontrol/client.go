// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package endpointcontrol

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure endpoint control client lists. Applies to FortiOS Version `<= 6.2.0`.
//
// ## Import
//
// EndpointControl Client can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:endpointcontrol/client:Client labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:endpointcontrol/client:Client labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Client struct {
	pulumi.CustomResourceState

	// Endpoint client AD logon groups.
	AdGroups pulumi.StringPtrOutput `pulumi:"adGroups"`
	// Endpoint client ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Endpoint FortiClient UID.
	FtclUid pulumi.StringOutput `pulumi:"ftclUid"`
	// Endpoint client information.
	Info pulumi.StringOutput `pulumi:"info"`
	// Endpoint client IP address.
	SrcIp pulumi.StringOutput `pulumi:"srcIp"`
	// Endpoint client MAC address.
	SrcMac pulumi.StringOutput `pulumi:"srcMac"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewClient registers a new resource with the given unique name, arguments, and options.
func NewClient(ctx *pulumi.Context,
	name string, args *ClientArgs, opts ...pulumi.ResourceOption) (*Client, error) {
	if args == nil {
		args = &ClientArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Client
	err := ctx.RegisterResource("fortios:endpointcontrol/client:Client", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClient gets an existing Client resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientState, opts ...pulumi.ResourceOption) (*Client, error) {
	var resource Client
	err := ctx.ReadResource("fortios:endpointcontrol/client:Client", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Client resources.
type clientState struct {
	// Endpoint client AD logon groups.
	AdGroups *string `pulumi:"adGroups"`
	// Endpoint client ID.
	Fosid *int `pulumi:"fosid"`
	// Endpoint FortiClient UID.
	FtclUid *string `pulumi:"ftclUid"`
	// Endpoint client information.
	Info *string `pulumi:"info"`
	// Endpoint client IP address.
	SrcIp *string `pulumi:"srcIp"`
	// Endpoint client MAC address.
	SrcMac *string `pulumi:"srcMac"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ClientState struct {
	// Endpoint client AD logon groups.
	AdGroups pulumi.StringPtrInput
	// Endpoint client ID.
	Fosid pulumi.IntPtrInput
	// Endpoint FortiClient UID.
	FtclUid pulumi.StringPtrInput
	// Endpoint client information.
	Info pulumi.StringPtrInput
	// Endpoint client IP address.
	SrcIp pulumi.StringPtrInput
	// Endpoint client MAC address.
	SrcMac pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ClientState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientState)(nil)).Elem()
}

type clientArgs struct {
	// Endpoint client AD logon groups.
	AdGroups *string `pulumi:"adGroups"`
	// Endpoint client ID.
	Fosid *int `pulumi:"fosid"`
	// Endpoint FortiClient UID.
	FtclUid *string `pulumi:"ftclUid"`
	// Endpoint client information.
	Info *string `pulumi:"info"`
	// Endpoint client IP address.
	SrcIp *string `pulumi:"srcIp"`
	// Endpoint client MAC address.
	SrcMac *string `pulumi:"srcMac"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Client resource.
type ClientArgs struct {
	// Endpoint client AD logon groups.
	AdGroups pulumi.StringPtrInput
	// Endpoint client ID.
	Fosid pulumi.IntPtrInput
	// Endpoint FortiClient UID.
	FtclUid pulumi.StringPtrInput
	// Endpoint client information.
	Info pulumi.StringPtrInput
	// Endpoint client IP address.
	SrcIp pulumi.StringPtrInput
	// Endpoint client MAC address.
	SrcMac pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientArgs)(nil)).Elem()
}

type ClientInput interface {
	pulumi.Input

	ToClientOutput() ClientOutput
	ToClientOutputWithContext(ctx context.Context) ClientOutput
}

func (*Client) ElementType() reflect.Type {
	return reflect.TypeOf((**Client)(nil)).Elem()
}

func (i *Client) ToClientOutput() ClientOutput {
	return i.ToClientOutputWithContext(context.Background())
}

func (i *Client) ToClientOutputWithContext(ctx context.Context) ClientOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientOutput)
}

// ClientArrayInput is an input type that accepts ClientArray and ClientArrayOutput values.
// You can construct a concrete instance of `ClientArrayInput` via:
//
//	ClientArray{ ClientArgs{...} }
type ClientArrayInput interface {
	pulumi.Input

	ToClientArrayOutput() ClientArrayOutput
	ToClientArrayOutputWithContext(context.Context) ClientArrayOutput
}

type ClientArray []ClientInput

func (ClientArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Client)(nil)).Elem()
}

func (i ClientArray) ToClientArrayOutput() ClientArrayOutput {
	return i.ToClientArrayOutputWithContext(context.Background())
}

func (i ClientArray) ToClientArrayOutputWithContext(ctx context.Context) ClientArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientArrayOutput)
}

// ClientMapInput is an input type that accepts ClientMap and ClientMapOutput values.
// You can construct a concrete instance of `ClientMapInput` via:
//
//	ClientMap{ "key": ClientArgs{...} }
type ClientMapInput interface {
	pulumi.Input

	ToClientMapOutput() ClientMapOutput
	ToClientMapOutputWithContext(context.Context) ClientMapOutput
}

type ClientMap map[string]ClientInput

func (ClientMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Client)(nil)).Elem()
}

func (i ClientMap) ToClientMapOutput() ClientMapOutput {
	return i.ToClientMapOutputWithContext(context.Background())
}

func (i ClientMap) ToClientMapOutputWithContext(ctx context.Context) ClientMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientMapOutput)
}

type ClientOutput struct{ *pulumi.OutputState }

func (ClientOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Client)(nil)).Elem()
}

func (o ClientOutput) ToClientOutput() ClientOutput {
	return o
}

func (o ClientOutput) ToClientOutputWithContext(ctx context.Context) ClientOutput {
	return o
}

// Endpoint client AD logon groups.
func (o ClientOutput) AdGroups() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Client) pulumi.StringPtrOutput { return v.AdGroups }).(pulumi.StringPtrOutput)
}

// Endpoint client ID.
func (o ClientOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Client) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Endpoint FortiClient UID.
func (o ClientOutput) FtclUid() pulumi.StringOutput {
	return o.ApplyT(func(v *Client) pulumi.StringOutput { return v.FtclUid }).(pulumi.StringOutput)
}

// Endpoint client information.
func (o ClientOutput) Info() pulumi.StringOutput {
	return o.ApplyT(func(v *Client) pulumi.StringOutput { return v.Info }).(pulumi.StringOutput)
}

// Endpoint client IP address.
func (o ClientOutput) SrcIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Client) pulumi.StringOutput { return v.SrcIp }).(pulumi.StringOutput)
}

// Endpoint client MAC address.
func (o ClientOutput) SrcMac() pulumi.StringOutput {
	return o.ApplyT(func(v *Client) pulumi.StringOutput { return v.SrcMac }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ClientOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Client) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type ClientArrayOutput struct{ *pulumi.OutputState }

func (ClientArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Client)(nil)).Elem()
}

func (o ClientArrayOutput) ToClientArrayOutput() ClientArrayOutput {
	return o
}

func (o ClientArrayOutput) ToClientArrayOutputWithContext(ctx context.Context) ClientArrayOutput {
	return o
}

func (o ClientArrayOutput) Index(i pulumi.IntInput) ClientOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Client {
		return vs[0].([]*Client)[vs[1].(int)]
	}).(ClientOutput)
}

type ClientMapOutput struct{ *pulumi.OutputState }

func (ClientMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Client)(nil)).Elem()
}

func (o ClientMapOutput) ToClientMapOutput() ClientMapOutput {
	return o
}

func (o ClientMapOutput) ToClientMapOutputWithContext(ctx context.Context) ClientMapOutput {
	return o
}

func (o ClientMapOutput) MapIndex(k pulumi.StringInput) ClientOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Client {
		return vs[0].(map[string]*Client)[vs[1].(string)]
	}).(ClientOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientInput)(nil)).Elem(), &Client{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientArrayInput)(nil)).Elem(), ClientArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientMapInput)(nil)).Elem(), ClientMap{})
	pulumi.RegisterOutputType(ClientOutput{})
	pulumi.RegisterOutputType(ClientArrayOutput{})
	pulumi.RegisterOutputType(ClientMapOutput{})
}
