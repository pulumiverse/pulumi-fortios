// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package webtrends

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Filters for WebTrends.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/log"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := log.NewFilter(ctx, "trname", &log.FilterArgs{
//				Anomaly:          pulumi.String("enable"),
//				Dns:              pulumi.String("enable"),
//				FilterType:       pulumi.String("include"),
//				ForwardTraffic:   pulumi.String("enable"),
//				Gtp:              pulumi.String("enable"),
//				LocalTraffic:     pulumi.String("enable"),
//				MulticastTraffic: pulumi.String("enable"),
//				Severity:         pulumi.String("information"),
//				SnifferTraffic:   pulumi.String("enable"),
//				Ssh:              pulumi.String("enable"),
//				Voip:             pulumi.String("enable"),
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
// LogWebtrends Filter can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:log/webtrends/filter:Filter labelname LogWebtrendsFilter
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:log/webtrends/filter:Filter labelname LogWebtrendsFilter
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Filter struct {
	pulumi.CustomResourceState

	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringOutput `pulumi:"anomaly"`
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns pulumi.StringOutput `pulumi:"dns"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Webtrends log filter.
	Filter pulumi.StringOutput `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringOutput `pulumi:"filterType"`
	// Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
	FortiSwitch pulumi.StringOutput `pulumi:"fortiSwitch"`
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic pulumi.StringOutput `pulumi:"forwardTraffic"`
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles FilterFreeStyleArrayOutput `pulumi:"freeStyles"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp pulumi.StringOutput `pulumi:"gtp"`
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic pulumi.StringOutput `pulumi:"localTraffic"`
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic pulumi.StringOutput `pulumi:"multicastTraffic"`
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery pulumi.StringOutput `pulumi:"netscanDiscovery"`
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability pulumi.StringOutput `pulumi:"netscanVulnerability"`
	// Lowest severity level to log to WebTrends. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity pulumi.StringOutput `pulumi:"severity"`
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic pulumi.StringOutput `pulumi:"snifferTraffic"`
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh pulumi.StringOutput `pulumi:"ssh"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip pulumi.StringOutput `pulumi:"voip"`
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic pulumi.StringOutput `pulumi:"ztnaTraffic"`
}

// NewFilter registers a new resource with the given unique name, arguments, and options.
func NewFilter(ctx *pulumi.Context,
	name string, args *FilterArgs, opts ...pulumi.ResourceOption) (*Filter, error) {
	if args == nil {
		args = &FilterArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Filter
	err := ctx.RegisterResource("fortios:log/webtrends/filter:Filter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFilter gets an existing Filter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FilterState, opts ...pulumi.ResourceOption) (*Filter, error) {
	var resource Filter
	err := ctx.ReadResource("fortios:log/webtrends/filter:Filter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Filter resources.
type filterState struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly *string `pulumi:"anomaly"`
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns *string `pulumi:"dns"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Webtrends log filter.
	Filter *string `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType *string `pulumi:"filterType"`
	// Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
	FortiSwitch *string `pulumi:"fortiSwitch"`
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic *string `pulumi:"forwardTraffic"`
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles []FilterFreeStyle `pulumi:"freeStyles"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp *string `pulumi:"gtp"`
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic *string `pulumi:"localTraffic"`
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic *string `pulumi:"multicastTraffic"`
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery *string `pulumi:"netscanDiscovery"`
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability *string `pulumi:"netscanVulnerability"`
	// Lowest severity level to log to WebTrends. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity *string `pulumi:"severity"`
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic *string `pulumi:"snifferTraffic"`
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh *string `pulumi:"ssh"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip *string `pulumi:"voip"`
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic *string `pulumi:"ztnaTraffic"`
}

type FilterState struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringPtrInput
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Webtrends log filter.
	Filter pulumi.StringPtrInput
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringPtrInput
	// Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
	FortiSwitch pulumi.StringPtrInput
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic pulumi.StringPtrInput
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles FilterFreeStyleArrayInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp pulumi.StringPtrInput
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic pulumi.StringPtrInput
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic pulumi.StringPtrInput
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery pulumi.StringPtrInput
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability pulumi.StringPtrInput
	// Lowest severity level to log to WebTrends. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity pulumi.StringPtrInput
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic pulumi.StringPtrInput
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip pulumi.StringPtrInput
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic pulumi.StringPtrInput
}

func (FilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*filterState)(nil)).Elem()
}

type filterArgs struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly *string `pulumi:"anomaly"`
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns *string `pulumi:"dns"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Webtrends log filter.
	Filter *string `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType *string `pulumi:"filterType"`
	// Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
	FortiSwitch *string `pulumi:"fortiSwitch"`
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic *string `pulumi:"forwardTraffic"`
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles []FilterFreeStyle `pulumi:"freeStyles"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp *string `pulumi:"gtp"`
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic *string `pulumi:"localTraffic"`
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic *string `pulumi:"multicastTraffic"`
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery *string `pulumi:"netscanDiscovery"`
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability *string `pulumi:"netscanVulnerability"`
	// Lowest severity level to log to WebTrends. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity *string `pulumi:"severity"`
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic *string `pulumi:"snifferTraffic"`
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh *string `pulumi:"ssh"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip *string `pulumi:"voip"`
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic *string `pulumi:"ztnaTraffic"`
}

// The set of arguments for constructing a Filter resource.
type FilterArgs struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringPtrInput
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Webtrends log filter.
	Filter pulumi.StringPtrInput
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringPtrInput
	// Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
	FortiSwitch pulumi.StringPtrInput
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic pulumi.StringPtrInput
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles FilterFreeStyleArrayInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp pulumi.StringPtrInput
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic pulumi.StringPtrInput
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic pulumi.StringPtrInput
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery pulumi.StringPtrInput
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability pulumi.StringPtrInput
	// Lowest severity level to log to WebTrends. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity pulumi.StringPtrInput
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic pulumi.StringPtrInput
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip pulumi.StringPtrInput
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic pulumi.StringPtrInput
}

func (FilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*filterArgs)(nil)).Elem()
}

type FilterInput interface {
	pulumi.Input

	ToFilterOutput() FilterOutput
	ToFilterOutputWithContext(ctx context.Context) FilterOutput
}

func (*Filter) ElementType() reflect.Type {
	return reflect.TypeOf((**Filter)(nil)).Elem()
}

func (i *Filter) ToFilterOutput() FilterOutput {
	return i.ToFilterOutputWithContext(context.Background())
}

func (i *Filter) ToFilterOutputWithContext(ctx context.Context) FilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterOutput)
}

// FilterArrayInput is an input type that accepts FilterArray and FilterArrayOutput values.
// You can construct a concrete instance of `FilterArrayInput` via:
//
//	FilterArray{ FilterArgs{...} }
type FilterArrayInput interface {
	pulumi.Input

	ToFilterArrayOutput() FilterArrayOutput
	ToFilterArrayOutputWithContext(context.Context) FilterArrayOutput
}

type FilterArray []FilterInput

func (FilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Filter)(nil)).Elem()
}

func (i FilterArray) ToFilterArrayOutput() FilterArrayOutput {
	return i.ToFilterArrayOutputWithContext(context.Background())
}

func (i FilterArray) ToFilterArrayOutputWithContext(ctx context.Context) FilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterArrayOutput)
}

// FilterMapInput is an input type that accepts FilterMap and FilterMapOutput values.
// You can construct a concrete instance of `FilterMapInput` via:
//
//	FilterMap{ "key": FilterArgs{...} }
type FilterMapInput interface {
	pulumi.Input

	ToFilterMapOutput() FilterMapOutput
	ToFilterMapOutputWithContext(context.Context) FilterMapOutput
}

type FilterMap map[string]FilterInput

func (FilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Filter)(nil)).Elem()
}

func (i FilterMap) ToFilterMapOutput() FilterMapOutput {
	return i.ToFilterMapOutputWithContext(context.Background())
}

func (i FilterMap) ToFilterMapOutputWithContext(ctx context.Context) FilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterMapOutput)
}

type FilterOutput struct{ *pulumi.OutputState }

func (FilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Filter)(nil)).Elem()
}

func (o FilterOutput) ToFilterOutput() FilterOutput {
	return o
}

func (o FilterOutput) ToFilterOutputWithContext(ctx context.Context) FilterOutput {
	return o
}

// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
func (o FilterOutput) Anomaly() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.Anomaly }).(pulumi.StringOutput)
}

// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
func (o FilterOutput) Dns() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.Dns }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o FilterOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Webtrends log filter.
func (o FilterOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
func (o FilterOutput) FilterType() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.FilterType }).(pulumi.StringOutput)
}

// Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
func (o FilterOutput) FortiSwitch() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.FortiSwitch }).(pulumi.StringOutput)
}

// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
func (o FilterOutput) ForwardTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.ForwardTraffic }).(pulumi.StringOutput)
}

// Free Style Filters The structure of `freeStyle` block is documented below.
func (o FilterOutput) FreeStyles() FilterFreeStyleArrayOutput {
	return o.ApplyT(func(v *Filter) FilterFreeStyleArrayOutput { return v.FreeStyles }).(FilterFreeStyleArrayOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o FilterOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
func (o FilterOutput) Gtp() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.Gtp }).(pulumi.StringOutput)
}

// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
func (o FilterOutput) LocalTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.LocalTraffic }).(pulumi.StringOutput)
}

// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
func (o FilterOutput) MulticastTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.MulticastTraffic }).(pulumi.StringOutput)
}

// Enable/disable netscan discovery event logging.
func (o FilterOutput) NetscanDiscovery() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.NetscanDiscovery }).(pulumi.StringOutput)
}

// Enable/disable netscan vulnerability event logging.
func (o FilterOutput) NetscanVulnerability() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.NetscanVulnerability }).(pulumi.StringOutput)
}

// Lowest severity level to log to WebTrends. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
func (o FilterOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
func (o FilterOutput) SnifferTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.SnifferTraffic }).(pulumi.StringOutput)
}

// Enable/disable SSH logging. Valid values: `enable`, `disable`.
func (o FilterOutput) Ssh() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.Ssh }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FilterOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
func (o FilterOutput) Voip() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.Voip }).(pulumi.StringOutput)
}

// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
func (o FilterOutput) ZtnaTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.ZtnaTraffic }).(pulumi.StringOutput)
}

type FilterArrayOutput struct{ *pulumi.OutputState }

func (FilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Filter)(nil)).Elem()
}

func (o FilterArrayOutput) ToFilterArrayOutput() FilterArrayOutput {
	return o
}

func (o FilterArrayOutput) ToFilterArrayOutputWithContext(ctx context.Context) FilterArrayOutput {
	return o
}

func (o FilterArrayOutput) Index(i pulumi.IntInput) FilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Filter {
		return vs[0].([]*Filter)[vs[1].(int)]
	}).(FilterOutput)
}

type FilterMapOutput struct{ *pulumi.OutputState }

func (FilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Filter)(nil)).Elem()
}

func (o FilterMapOutput) ToFilterMapOutput() FilterMapOutput {
	return o
}

func (o FilterMapOutput) ToFilterMapOutputWithContext(ctx context.Context) FilterMapOutput {
	return o
}

func (o FilterMapOutput) MapIndex(k pulumi.StringInput) FilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Filter {
		return vs[0].(map[string]*Filter)[vs[1].(string)]
	}).(FilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FilterInput)(nil)).Elem(), &Filter{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilterArrayInput)(nil)).Elem(), FilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilterMapInput)(nil)).Elem(), FilterMap{})
	pulumi.RegisterOutputType(FilterOutput{})
	pulumi.RegisterOutputType(FilterArrayOutput{})
	pulumi.RegisterOutputType(FilterMapOutput{})
}
