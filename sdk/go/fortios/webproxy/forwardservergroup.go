// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package webproxy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
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
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/webproxy"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			trname1Forwardserver, err := webproxy.NewForwardserver(ctx, "trname1Forwardserver", &webproxy.ForwardserverArgs{
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
//			_, err = webproxy.NewForwardservergroup(ctx, "trname1Forwardservergroup", &webproxy.ForwardservergroupArgs{
//				Affinity:        pulumi.String("disable"),
//				GroupDownOption: pulumi.String("block"),
//				LdbMethod:       pulumi.String("weighted"),
//				ServerLists: webproxy.ForwardservergroupServerListArray{
//					&webproxy.ForwardservergroupServerListArgs{
//						Name:   trname1Forwardserver.Name,
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
// WebProxy ForwardServerGroup can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:webproxy/forwardservergroup:Forwardservergroup labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:webproxy/forwardservergroup:Forwardservergroup labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Forwardservergroup struct {
	pulumi.CustomResourceState

	// Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable`, `disable`.
	Affinity pulumi.StringOutput `pulumi:"affinity"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
	GroupDownOption pulumi.StringOutput `pulumi:"groupDownOption"`
	// Load balance method: weighted or least-session.
	LdbMethod pulumi.StringOutput `pulumi:"ldbMethod"`
	// Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name pulumi.StringOutput `pulumi:"name"`
	// Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists ForwardservergroupServerListArrayOutput `pulumi:"serverLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewForwardservergroup registers a new resource with the given unique name, arguments, and options.
func NewForwardservergroup(ctx *pulumi.Context,
	name string, args *ForwardservergroupArgs, opts ...pulumi.ResourceOption) (*Forwardservergroup, error) {
	if args == nil {
		args = &ForwardservergroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Forwardservergroup
	err := ctx.RegisterResource("fortios:webproxy/forwardservergroup:Forwardservergroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetForwardservergroup gets an existing Forwardservergroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetForwardservergroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ForwardservergroupState, opts ...pulumi.ResourceOption) (*Forwardservergroup, error) {
	var resource Forwardservergroup
	err := ctx.ReadResource("fortios:webproxy/forwardservergroup:Forwardservergroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Forwardservergroup resources.
type forwardservergroupState struct {
	// Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable`, `disable`.
	Affinity *string `pulumi:"affinity"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
	GroupDownOption *string `pulumi:"groupDownOption"`
	// Load balance method: weighted or least-session.
	LdbMethod *string `pulumi:"ldbMethod"`
	// Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name *string `pulumi:"name"`
	// Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists []ForwardservergroupServerList `pulumi:"serverLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ForwardservergroupState struct {
	// Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable`, `disable`.
	Affinity pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
	GroupDownOption pulumi.StringPtrInput
	// Load balance method: weighted or least-session.
	LdbMethod pulumi.StringPtrInput
	// Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name pulumi.StringPtrInput
	// Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists ForwardservergroupServerListArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ForwardservergroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*forwardservergroupState)(nil)).Elem()
}

type forwardservergroupArgs struct {
	// Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable`, `disable`.
	Affinity *string `pulumi:"affinity"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
	GroupDownOption *string `pulumi:"groupDownOption"`
	// Load balance method: weighted or least-session.
	LdbMethod *string `pulumi:"ldbMethod"`
	// Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name *string `pulumi:"name"`
	// Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists []ForwardservergroupServerList `pulumi:"serverLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Forwardservergroup resource.
type ForwardservergroupArgs struct {
	// Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable`, `disable`.
	Affinity pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
	GroupDownOption pulumi.StringPtrInput
	// Load balance method: weighted or least-session.
	LdbMethod pulumi.StringPtrInput
	// Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name pulumi.StringPtrInput
	// Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists ForwardservergroupServerListArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ForwardservergroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*forwardservergroupArgs)(nil)).Elem()
}

type ForwardservergroupInput interface {
	pulumi.Input

	ToForwardservergroupOutput() ForwardservergroupOutput
	ToForwardservergroupOutputWithContext(ctx context.Context) ForwardservergroupOutput
}

func (*Forwardservergroup) ElementType() reflect.Type {
	return reflect.TypeOf((**Forwardservergroup)(nil)).Elem()
}

func (i *Forwardservergroup) ToForwardservergroupOutput() ForwardservergroupOutput {
	return i.ToForwardservergroupOutputWithContext(context.Background())
}

func (i *Forwardservergroup) ToForwardservergroupOutputWithContext(ctx context.Context) ForwardservergroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardservergroupOutput)
}

// ForwardservergroupArrayInput is an input type that accepts ForwardservergroupArray and ForwardservergroupArrayOutput values.
// You can construct a concrete instance of `ForwardservergroupArrayInput` via:
//
//	ForwardservergroupArray{ ForwardservergroupArgs{...} }
type ForwardservergroupArrayInput interface {
	pulumi.Input

	ToForwardservergroupArrayOutput() ForwardservergroupArrayOutput
	ToForwardservergroupArrayOutputWithContext(context.Context) ForwardservergroupArrayOutput
}

type ForwardservergroupArray []ForwardservergroupInput

func (ForwardservergroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Forwardservergroup)(nil)).Elem()
}

func (i ForwardservergroupArray) ToForwardservergroupArrayOutput() ForwardservergroupArrayOutput {
	return i.ToForwardservergroupArrayOutputWithContext(context.Background())
}

func (i ForwardservergroupArray) ToForwardservergroupArrayOutputWithContext(ctx context.Context) ForwardservergroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardservergroupArrayOutput)
}

// ForwardservergroupMapInput is an input type that accepts ForwardservergroupMap and ForwardservergroupMapOutput values.
// You can construct a concrete instance of `ForwardservergroupMapInput` via:
//
//	ForwardservergroupMap{ "key": ForwardservergroupArgs{...} }
type ForwardservergroupMapInput interface {
	pulumi.Input

	ToForwardservergroupMapOutput() ForwardservergroupMapOutput
	ToForwardservergroupMapOutputWithContext(context.Context) ForwardservergroupMapOutput
}

type ForwardservergroupMap map[string]ForwardservergroupInput

func (ForwardservergroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Forwardservergroup)(nil)).Elem()
}

func (i ForwardservergroupMap) ToForwardservergroupMapOutput() ForwardservergroupMapOutput {
	return i.ToForwardservergroupMapOutputWithContext(context.Background())
}

func (i ForwardservergroupMap) ToForwardservergroupMapOutputWithContext(ctx context.Context) ForwardservergroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardservergroupMapOutput)
}

type ForwardservergroupOutput struct{ *pulumi.OutputState }

func (ForwardservergroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Forwardservergroup)(nil)).Elem()
}

func (o ForwardservergroupOutput) ToForwardservergroupOutput() ForwardservergroupOutput {
	return o
}

func (o ForwardservergroupOutput) ToForwardservergroupOutputWithContext(ctx context.Context) ForwardservergroupOutput {
	return o
}

// Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable`, `disable`.
func (o ForwardservergroupOutput) Affinity() pulumi.StringOutput {
	return o.ApplyT(func(v *Forwardservergroup) pulumi.StringOutput { return v.Affinity }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o ForwardservergroupOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Forwardservergroup) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o ForwardservergroupOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Forwardservergroup) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
func (o ForwardservergroupOutput) GroupDownOption() pulumi.StringOutput {
	return o.ApplyT(func(v *Forwardservergroup) pulumi.StringOutput { return v.GroupDownOption }).(pulumi.StringOutput)
}

// Load balance method: weighted or least-session.
func (o ForwardservergroupOutput) LdbMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *Forwardservergroup) pulumi.StringOutput { return v.LdbMethod }).(pulumi.StringOutput)
}

// Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
func (o ForwardservergroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Forwardservergroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
func (o ForwardservergroupOutput) ServerLists() ForwardservergroupServerListArrayOutput {
	return o.ApplyT(func(v *Forwardservergroup) ForwardservergroupServerListArrayOutput { return v.ServerLists }).(ForwardservergroupServerListArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ForwardservergroupOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Forwardservergroup) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ForwardservergroupArrayOutput struct{ *pulumi.OutputState }

func (ForwardservergroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Forwardservergroup)(nil)).Elem()
}

func (o ForwardservergroupArrayOutput) ToForwardservergroupArrayOutput() ForwardservergroupArrayOutput {
	return o
}

func (o ForwardservergroupArrayOutput) ToForwardservergroupArrayOutputWithContext(ctx context.Context) ForwardservergroupArrayOutput {
	return o
}

func (o ForwardservergroupArrayOutput) Index(i pulumi.IntInput) ForwardservergroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Forwardservergroup {
		return vs[0].([]*Forwardservergroup)[vs[1].(int)]
	}).(ForwardservergroupOutput)
}

type ForwardservergroupMapOutput struct{ *pulumi.OutputState }

func (ForwardservergroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Forwardservergroup)(nil)).Elem()
}

func (o ForwardservergroupMapOutput) ToForwardservergroupMapOutput() ForwardservergroupMapOutput {
	return o
}

func (o ForwardservergroupMapOutput) ToForwardservergroupMapOutputWithContext(ctx context.Context) ForwardservergroupMapOutput {
	return o
}

func (o ForwardservergroupMapOutput) MapIndex(k pulumi.StringInput) ForwardservergroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Forwardservergroup {
		return vs[0].(map[string]*Forwardservergroup)[vs[1].(string)]
	}).(ForwardservergroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ForwardservergroupInput)(nil)).Elem(), &Forwardservergroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ForwardservergroupArrayInput)(nil)).Elem(), ForwardservergroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ForwardservergroupMapInput)(nil)).Elem(), ForwardservergroupMap{})
	pulumi.RegisterOutputType(ForwardservergroupOutput{})
	pulumi.RegisterOutputType(ForwardservergroupArrayOutput{})
	pulumi.RegisterOutputType(ForwardservergroupMapOutput{})
}
