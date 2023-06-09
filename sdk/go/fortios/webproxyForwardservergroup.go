// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure a forward server group consisting or multiple forward servers. Supports failover and load balancing.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			trname1WebproxyForwardserver, err := fortios.NewWebproxyForwardserver(ctx, "trname1WebproxyForwardserver", &fortios.WebproxyForwardserverArgs{
//				AddrType:         pulumi.String("fqdn"),
//				Healthcheck:      pulumi.String("disable"),
//				Ip:               pulumi.String("0.0.0.0"),
//				Monitor:          pulumi.String("http://www.google.com"),
//				Port:             pulumi.Int(1128),
//				ServerDownOption: pulumi.String("block"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = fortios.NewWebproxyForwardservergroup(ctx, "trname1WebproxyForwardservergroup", &fortios.WebproxyForwardservergroupArgs{
//				Affinity:        pulumi.String("disable"),
//				GroupDownOption: pulumi.String("block"),
//				LdbMethod:       pulumi.String("weighted"),
//				ServerLists: fortios.WebproxyForwardservergroupServerListArray{
//					&fortios.WebproxyForwardservergroupServerListArgs{
//						Name:   trname1WebproxyForwardserver.Name,
//						Weight: pulumi.Int(12),
//					},
//				},
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
// # WebProxy ForwardServerGroup can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/webproxyForwardservergroup:WebproxyForwardservergroup labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/webproxyForwardservergroup:WebproxyForwardservergroup labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type WebproxyForwardservergroup struct {
	pulumi.CustomResourceState

	// Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable`, `disable`.
	Affinity pulumi.StringOutput `pulumi:"affinity"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
	GroupDownOption pulumi.StringOutput `pulumi:"groupDownOption"`
	// Load balance method: weighted or least-session.
	LdbMethod pulumi.StringOutput `pulumi:"ldbMethod"`
	// Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name pulumi.StringOutput `pulumi:"name"`
	// Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists WebproxyForwardservergroupServerListArrayOutput `pulumi:"serverLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWebproxyForwardservergroup registers a new resource with the given unique name, arguments, and options.
func NewWebproxyForwardservergroup(ctx *pulumi.Context,
	name string, args *WebproxyForwardservergroupArgs, opts ...pulumi.ResourceOption) (*WebproxyForwardservergroup, error) {
	if args == nil {
		args = &WebproxyForwardservergroupArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WebproxyForwardservergroup
	err := ctx.RegisterResource("fortios:index/webproxyForwardservergroup:WebproxyForwardservergroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebproxyForwardservergroup gets an existing WebproxyForwardservergroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebproxyForwardservergroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebproxyForwardservergroupState, opts ...pulumi.ResourceOption) (*WebproxyForwardservergroup, error) {
	var resource WebproxyForwardservergroup
	err := ctx.ReadResource("fortios:index/webproxyForwardservergroup:WebproxyForwardservergroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebproxyForwardservergroup resources.
type webproxyForwardservergroupState struct {
	// Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable`, `disable`.
	Affinity *string `pulumi:"affinity"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
	GroupDownOption *string `pulumi:"groupDownOption"`
	// Load balance method: weighted or least-session.
	LdbMethod *string `pulumi:"ldbMethod"`
	// Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name *string `pulumi:"name"`
	// Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists []WebproxyForwardservergroupServerList `pulumi:"serverLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WebproxyForwardservergroupState struct {
	// Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable`, `disable`.
	Affinity pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
	GroupDownOption pulumi.StringPtrInput
	// Load balance method: weighted or least-session.
	LdbMethod pulumi.StringPtrInput
	// Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name pulumi.StringPtrInput
	// Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists WebproxyForwardservergroupServerListArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebproxyForwardservergroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*webproxyForwardservergroupState)(nil)).Elem()
}

type webproxyForwardservergroupArgs struct {
	// Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable`, `disable`.
	Affinity *string `pulumi:"affinity"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
	GroupDownOption *string `pulumi:"groupDownOption"`
	// Load balance method: weighted or least-session.
	LdbMethod *string `pulumi:"ldbMethod"`
	// Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name *string `pulumi:"name"`
	// Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists []WebproxyForwardservergroupServerList `pulumi:"serverLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WebproxyForwardservergroup resource.
type WebproxyForwardservergroupArgs struct {
	// Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable`, `disable`.
	Affinity pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
	GroupDownOption pulumi.StringPtrInput
	// Load balance method: weighted or least-session.
	LdbMethod pulumi.StringPtrInput
	// Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name pulumi.StringPtrInput
	// Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists WebproxyForwardservergroupServerListArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebproxyForwardservergroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webproxyForwardservergroupArgs)(nil)).Elem()
}

type WebproxyForwardservergroupInput interface {
	pulumi.Input

	ToWebproxyForwardservergroupOutput() WebproxyForwardservergroupOutput
	ToWebproxyForwardservergroupOutputWithContext(ctx context.Context) WebproxyForwardservergroupOutput
}

func (*WebproxyForwardservergroup) ElementType() reflect.Type {
	return reflect.TypeOf((**WebproxyForwardservergroup)(nil)).Elem()
}

func (i *WebproxyForwardservergroup) ToWebproxyForwardservergroupOutput() WebproxyForwardservergroupOutput {
	return i.ToWebproxyForwardservergroupOutputWithContext(context.Background())
}

func (i *WebproxyForwardservergroup) ToWebproxyForwardservergroupOutputWithContext(ctx context.Context) WebproxyForwardservergroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebproxyForwardservergroupOutput)
}

// WebproxyForwardservergroupArrayInput is an input type that accepts WebproxyForwardservergroupArray and WebproxyForwardservergroupArrayOutput values.
// You can construct a concrete instance of `WebproxyForwardservergroupArrayInput` via:
//
//	WebproxyForwardservergroupArray{ WebproxyForwardservergroupArgs{...} }
type WebproxyForwardservergroupArrayInput interface {
	pulumi.Input

	ToWebproxyForwardservergroupArrayOutput() WebproxyForwardservergroupArrayOutput
	ToWebproxyForwardservergroupArrayOutputWithContext(context.Context) WebproxyForwardservergroupArrayOutput
}

type WebproxyForwardservergroupArray []WebproxyForwardservergroupInput

func (WebproxyForwardservergroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebproxyForwardservergroup)(nil)).Elem()
}

func (i WebproxyForwardservergroupArray) ToWebproxyForwardservergroupArrayOutput() WebproxyForwardservergroupArrayOutput {
	return i.ToWebproxyForwardservergroupArrayOutputWithContext(context.Background())
}

func (i WebproxyForwardservergroupArray) ToWebproxyForwardservergroupArrayOutputWithContext(ctx context.Context) WebproxyForwardservergroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebproxyForwardservergroupArrayOutput)
}

// WebproxyForwardservergroupMapInput is an input type that accepts WebproxyForwardservergroupMap and WebproxyForwardservergroupMapOutput values.
// You can construct a concrete instance of `WebproxyForwardservergroupMapInput` via:
//
//	WebproxyForwardservergroupMap{ "key": WebproxyForwardservergroupArgs{...} }
type WebproxyForwardservergroupMapInput interface {
	pulumi.Input

	ToWebproxyForwardservergroupMapOutput() WebproxyForwardservergroupMapOutput
	ToWebproxyForwardservergroupMapOutputWithContext(context.Context) WebproxyForwardservergroupMapOutput
}

type WebproxyForwardservergroupMap map[string]WebproxyForwardservergroupInput

func (WebproxyForwardservergroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebproxyForwardservergroup)(nil)).Elem()
}

func (i WebproxyForwardservergroupMap) ToWebproxyForwardservergroupMapOutput() WebproxyForwardservergroupMapOutput {
	return i.ToWebproxyForwardservergroupMapOutputWithContext(context.Background())
}

func (i WebproxyForwardservergroupMap) ToWebproxyForwardservergroupMapOutputWithContext(ctx context.Context) WebproxyForwardservergroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebproxyForwardservergroupMapOutput)
}

type WebproxyForwardservergroupOutput struct{ *pulumi.OutputState }

func (WebproxyForwardservergroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebproxyForwardservergroup)(nil)).Elem()
}

func (o WebproxyForwardservergroupOutput) ToWebproxyForwardservergroupOutput() WebproxyForwardservergroupOutput {
	return o
}

func (o WebproxyForwardservergroupOutput) ToWebproxyForwardservergroupOutputWithContext(ctx context.Context) WebproxyForwardservergroupOutput {
	return o
}

// Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable`, `disable`.
func (o WebproxyForwardservergroupOutput) Affinity() pulumi.StringOutput {
	return o.ApplyT(func(v *WebproxyForwardservergroup) pulumi.StringOutput { return v.Affinity }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o WebproxyForwardservergroupOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebproxyForwardservergroup) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
func (o WebproxyForwardservergroupOutput) GroupDownOption() pulumi.StringOutput {
	return o.ApplyT(func(v *WebproxyForwardservergroup) pulumi.StringOutput { return v.GroupDownOption }).(pulumi.StringOutput)
}

// Load balance method: weighted or least-session.
func (o WebproxyForwardservergroupOutput) LdbMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *WebproxyForwardservergroup) pulumi.StringOutput { return v.LdbMethod }).(pulumi.StringOutput)
}

// Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
func (o WebproxyForwardservergroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebproxyForwardservergroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
func (o WebproxyForwardservergroupOutput) ServerLists() WebproxyForwardservergroupServerListArrayOutput {
	return o.ApplyT(func(v *WebproxyForwardservergroup) WebproxyForwardservergroupServerListArrayOutput {
		return v.ServerLists
	}).(WebproxyForwardservergroupServerListArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o WebproxyForwardservergroupOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebproxyForwardservergroup) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WebproxyForwardservergroupArrayOutput struct{ *pulumi.OutputState }

func (WebproxyForwardservergroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebproxyForwardservergroup)(nil)).Elem()
}

func (o WebproxyForwardservergroupArrayOutput) ToWebproxyForwardservergroupArrayOutput() WebproxyForwardservergroupArrayOutput {
	return o
}

func (o WebproxyForwardservergroupArrayOutput) ToWebproxyForwardservergroupArrayOutputWithContext(ctx context.Context) WebproxyForwardservergroupArrayOutput {
	return o
}

func (o WebproxyForwardservergroupArrayOutput) Index(i pulumi.IntInput) WebproxyForwardservergroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebproxyForwardservergroup {
		return vs[0].([]*WebproxyForwardservergroup)[vs[1].(int)]
	}).(WebproxyForwardservergroupOutput)
}

type WebproxyForwardservergroupMapOutput struct{ *pulumi.OutputState }

func (WebproxyForwardservergroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebproxyForwardservergroup)(nil)).Elem()
}

func (o WebproxyForwardservergroupMapOutput) ToWebproxyForwardservergroupMapOutput() WebproxyForwardservergroupMapOutput {
	return o
}

func (o WebproxyForwardservergroupMapOutput) ToWebproxyForwardservergroupMapOutputWithContext(ctx context.Context) WebproxyForwardservergroupMapOutput {
	return o
}

func (o WebproxyForwardservergroupMapOutput) MapIndex(k pulumi.StringInput) WebproxyForwardservergroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebproxyForwardservergroup {
		return vs[0].(map[string]*WebproxyForwardservergroup)[vs[1].(string)]
	}).(WebproxyForwardservergroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebproxyForwardservergroupInput)(nil)).Elem(), &WebproxyForwardservergroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebproxyForwardservergroupArrayInput)(nil)).Elem(), WebproxyForwardservergroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebproxyForwardservergroupMapInput)(nil)).Elem(), WebproxyForwardservergroupMap{})
	pulumi.RegisterOutputType(WebproxyForwardservergroupOutput{})
	pulumi.RegisterOutputType(WebproxyForwardservergroupArrayOutput{})
	pulumi.RegisterOutputType(WebproxyForwardservergroupMapOutput{})
}
