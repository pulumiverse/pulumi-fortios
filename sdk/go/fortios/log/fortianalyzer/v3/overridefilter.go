// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Override filters for FortiAnalyzer.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := log.NewOverridefilter(ctx, "trname", &log.OverridefilterArgs{
//				Anomaly:          pulumi.String("enable"),
//				DlpArchive:       pulumi.String("enable"),
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// LogFortianalyzer3 OverrideFilter can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:log/fortianalyzer/v3/overridefilter:Overridefilter labelname LogFortianalyzer3OverrideFilter
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:log/fortianalyzer/v3/overridefilter:Overridefilter labelname LogFortianalyzer3OverrideFilter
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Overridefilter struct {
	pulumi.CustomResourceState

	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringOutput `pulumi:"anomaly"`
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive pulumi.StringOutput `pulumi:"dlpArchive"`
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns pulumi.StringOutput `pulumi:"dns"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// FortiAnalyzer 3 log filter.
	Filter pulumi.StringOutput `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringOutput `pulumi:"filterType"`
	// Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
	FortiSwitch pulumi.StringOutput `pulumi:"fortiSwitch"`
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic pulumi.StringOutput `pulumi:"forwardTraffic"`
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles OverridefilterFreeStyleArrayOutput `pulumi:"freeStyles"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
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
	// Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity pulumi.StringOutput `pulumi:"severity"`
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic pulumi.StringOutput `pulumi:"snifferTraffic"`
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh pulumi.StringOutput `pulumi:"ssh"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip pulumi.StringOutput `pulumi:"voip"`
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic pulumi.StringOutput `pulumi:"ztnaTraffic"`
}

// NewOverridefilter registers a new resource with the given unique name, arguments, and options.
func NewOverridefilter(ctx *pulumi.Context,
	name string, args *OverridefilterArgs, opts ...pulumi.ResourceOption) (*Overridefilter, error) {
	if args == nil {
		args = &OverridefilterArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Overridefilter
	err := ctx.RegisterResource("fortios:log/fortianalyzer/v3/overridefilter:Overridefilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOverridefilter gets an existing Overridefilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOverridefilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OverridefilterState, opts ...pulumi.ResourceOption) (*Overridefilter, error) {
	var resource Overridefilter
	err := ctx.ReadResource("fortios:log/fortianalyzer/v3/overridefilter:Overridefilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Overridefilter resources.
type overridefilterState struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly *string `pulumi:"anomaly"`
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive *string `pulumi:"dlpArchive"`
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns *string `pulumi:"dns"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// FortiAnalyzer 3 log filter.
	Filter *string `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType *string `pulumi:"filterType"`
	// Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
	FortiSwitch *string `pulumi:"fortiSwitch"`
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic *string `pulumi:"forwardTraffic"`
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles []OverridefilterFreeStyle `pulumi:"freeStyles"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
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
	// Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
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

type OverridefilterState struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringPtrInput
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive pulumi.StringPtrInput
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// FortiAnalyzer 3 log filter.
	Filter pulumi.StringPtrInput
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringPtrInput
	// Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
	FortiSwitch pulumi.StringPtrInput
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic pulumi.StringPtrInput
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles OverridefilterFreeStyleArrayInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
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
	// Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
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

func (OverridefilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*overridefilterState)(nil)).Elem()
}

type overridefilterArgs struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly *string `pulumi:"anomaly"`
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive *string `pulumi:"dlpArchive"`
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns *string `pulumi:"dns"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// FortiAnalyzer 3 log filter.
	Filter *string `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType *string `pulumi:"filterType"`
	// Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
	FortiSwitch *string `pulumi:"fortiSwitch"`
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic *string `pulumi:"forwardTraffic"`
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles []OverridefilterFreeStyle `pulumi:"freeStyles"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
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
	// Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
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

// The set of arguments for constructing a Overridefilter resource.
type OverridefilterArgs struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringPtrInput
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive pulumi.StringPtrInput
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// FortiAnalyzer 3 log filter.
	Filter pulumi.StringPtrInput
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringPtrInput
	// Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
	FortiSwitch pulumi.StringPtrInput
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic pulumi.StringPtrInput
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles OverridefilterFreeStyleArrayInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
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
	// Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
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

func (OverridefilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*overridefilterArgs)(nil)).Elem()
}

type OverridefilterInput interface {
	pulumi.Input

	ToOverridefilterOutput() OverridefilterOutput
	ToOverridefilterOutputWithContext(ctx context.Context) OverridefilterOutput
}

func (*Overridefilter) ElementType() reflect.Type {
	return reflect.TypeOf((**Overridefilter)(nil)).Elem()
}

func (i *Overridefilter) ToOverridefilterOutput() OverridefilterOutput {
	return i.ToOverridefilterOutputWithContext(context.Background())
}

func (i *Overridefilter) ToOverridefilterOutputWithContext(ctx context.Context) OverridefilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverridefilterOutput)
}

// OverridefilterArrayInput is an input type that accepts OverridefilterArray and OverridefilterArrayOutput values.
// You can construct a concrete instance of `OverridefilterArrayInput` via:
//
//	OverridefilterArray{ OverridefilterArgs{...} }
type OverridefilterArrayInput interface {
	pulumi.Input

	ToOverridefilterArrayOutput() OverridefilterArrayOutput
	ToOverridefilterArrayOutputWithContext(context.Context) OverridefilterArrayOutput
}

type OverridefilterArray []OverridefilterInput

func (OverridefilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Overridefilter)(nil)).Elem()
}

func (i OverridefilterArray) ToOverridefilterArrayOutput() OverridefilterArrayOutput {
	return i.ToOverridefilterArrayOutputWithContext(context.Background())
}

func (i OverridefilterArray) ToOverridefilterArrayOutputWithContext(ctx context.Context) OverridefilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverridefilterArrayOutput)
}

// OverridefilterMapInput is an input type that accepts OverridefilterMap and OverridefilterMapOutput values.
// You can construct a concrete instance of `OverridefilterMapInput` via:
//
//	OverridefilterMap{ "key": OverridefilterArgs{...} }
type OverridefilterMapInput interface {
	pulumi.Input

	ToOverridefilterMapOutput() OverridefilterMapOutput
	ToOverridefilterMapOutputWithContext(context.Context) OverridefilterMapOutput
}

type OverridefilterMap map[string]OverridefilterInput

func (OverridefilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Overridefilter)(nil)).Elem()
}

func (i OverridefilterMap) ToOverridefilterMapOutput() OverridefilterMapOutput {
	return i.ToOverridefilterMapOutputWithContext(context.Background())
}

func (i OverridefilterMap) ToOverridefilterMapOutputWithContext(ctx context.Context) OverridefilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverridefilterMapOutput)
}

type OverridefilterOutput struct{ *pulumi.OutputState }

func (OverridefilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Overridefilter)(nil)).Elem()
}

func (o OverridefilterOutput) ToOverridefilterOutput() OverridefilterOutput {
	return o
}

func (o OverridefilterOutput) ToOverridefilterOutputWithContext(ctx context.Context) OverridefilterOutput {
	return o
}

// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
func (o OverridefilterOutput) Anomaly() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringOutput { return v.Anomaly }).(pulumi.StringOutput)
}

// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
func (o OverridefilterOutput) DlpArchive() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringOutput { return v.DlpArchive }).(pulumi.StringOutput)
}

// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
func (o OverridefilterOutput) Dns() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringOutput { return v.Dns }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o OverridefilterOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// FortiAnalyzer 3 log filter.
func (o OverridefilterOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
func (o OverridefilterOutput) FilterType() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringOutput { return v.FilterType }).(pulumi.StringOutput)
}

// Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
func (o OverridefilterOutput) FortiSwitch() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringOutput { return v.FortiSwitch }).(pulumi.StringOutput)
}

// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
func (o OverridefilterOutput) ForwardTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringOutput { return v.ForwardTraffic }).(pulumi.StringOutput)
}

// Free Style Filters The structure of `freeStyle` block is documented below.
func (o OverridefilterOutput) FreeStyles() OverridefilterFreeStyleArrayOutput {
	return o.ApplyT(func(v *Overridefilter) OverridefilterFreeStyleArrayOutput { return v.FreeStyles }).(OverridefilterFreeStyleArrayOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o OverridefilterOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
func (o OverridefilterOutput) Gtp() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringOutput { return v.Gtp }).(pulumi.StringOutput)
}

// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
func (o OverridefilterOutput) LocalTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringOutput { return v.LocalTraffic }).(pulumi.StringOutput)
}

// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
func (o OverridefilterOutput) MulticastTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringOutput { return v.MulticastTraffic }).(pulumi.StringOutput)
}

// Enable/disable netscan discovery event logging.
func (o OverridefilterOutput) NetscanDiscovery() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringOutput { return v.NetscanDiscovery }).(pulumi.StringOutput)
}

// Enable/disable netscan vulnerability event logging.
func (o OverridefilterOutput) NetscanVulnerability() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringOutput { return v.NetscanVulnerability }).(pulumi.StringOutput)
}

// Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
func (o OverridefilterOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
func (o OverridefilterOutput) SnifferTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringOutput { return v.SnifferTraffic }).(pulumi.StringOutput)
}

// Enable/disable SSH logging. Valid values: `enable`, `disable`.
func (o OverridefilterOutput) Ssh() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringOutput { return v.Ssh }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o OverridefilterOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
func (o OverridefilterOutput) Voip() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringOutput { return v.Voip }).(pulumi.StringOutput)
}

// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
func (o OverridefilterOutput) ZtnaTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridefilter) pulumi.StringOutput { return v.ZtnaTraffic }).(pulumi.StringOutput)
}

type OverridefilterArrayOutput struct{ *pulumi.OutputState }

func (OverridefilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Overridefilter)(nil)).Elem()
}

func (o OverridefilterArrayOutput) ToOverridefilterArrayOutput() OverridefilterArrayOutput {
	return o
}

func (o OverridefilterArrayOutput) ToOverridefilterArrayOutputWithContext(ctx context.Context) OverridefilterArrayOutput {
	return o
}

func (o OverridefilterArrayOutput) Index(i pulumi.IntInput) OverridefilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Overridefilter {
		return vs[0].([]*Overridefilter)[vs[1].(int)]
	}).(OverridefilterOutput)
}

type OverridefilterMapOutput struct{ *pulumi.OutputState }

func (OverridefilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Overridefilter)(nil)).Elem()
}

func (o OverridefilterMapOutput) ToOverridefilterMapOutput() OverridefilterMapOutput {
	return o
}

func (o OverridefilterMapOutput) ToOverridefilterMapOutputWithContext(ctx context.Context) OverridefilterMapOutput {
	return o
}

func (o OverridefilterMapOutput) MapIndex(k pulumi.StringInput) OverridefilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Overridefilter {
		return vs[0].(map[string]*Overridefilter)[vs[1].(string)]
	}).(OverridefilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OverridefilterInput)(nil)).Elem(), &Overridefilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*OverridefilterArrayInput)(nil)).Elem(), OverridefilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OverridefilterMapInput)(nil)).Elem(), OverridefilterMap{})
	pulumi.RegisterOutputType(OverridefilterOutput{})
	pulumi.RegisterOutputType(OverridefilterArrayOutput{})
	pulumi.RegisterOutputType(OverridefilterMapOutput{})
}
