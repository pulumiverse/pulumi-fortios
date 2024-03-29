// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure sFlow.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/system"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := system.NewSflow(ctx, "trname", &system.SflowArgs{
//				CollectorIp:   pulumi.String("0.0.0.0"),
//				CollectorPort: pulumi.Int(6343),
//				SourceIp:      pulumi.String("0.0.0.0"),
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
// System Sflow can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/sflow:Sflow labelname SystemSflow
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/sflow:Sflow labelname SystemSflow
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Sflow struct {
	pulumi.CustomResourceState

	// IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
	CollectorIp pulumi.StringOutput `pulumi:"collectorIp"`
	// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
	CollectorPort pulumi.IntOutput `pulumi:"collectorPort"`
	// sFlow collectors. The structure of `collectors` block is documented below.
	Collectors SflowCollectorArrayOutput `pulumi:"collectors"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Specify outgoing interface to reach server.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringOutput `pulumi:"interfaceSelectMethod"`
	// Source IP address for sFlow agent.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSflow registers a new resource with the given unique name, arguments, and options.
func NewSflow(ctx *pulumi.Context,
	name string, args *SflowArgs, opts ...pulumi.ResourceOption) (*Sflow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CollectorIp == nil {
		return nil, errors.New("invalid value for required argument 'CollectorIp'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Sflow
	err := ctx.RegisterResource("fortios:system/sflow:Sflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSflow gets an existing Sflow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SflowState, opts ...pulumi.ResourceOption) (*Sflow, error) {
	var resource Sflow
	err := ctx.ReadResource("fortios:system/sflow:Sflow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Sflow resources.
type sflowState struct {
	// IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
	CollectorIp *string `pulumi:"collectorIp"`
	// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
	CollectorPort *int `pulumi:"collectorPort"`
	// sFlow collectors. The structure of `collectors` block is documented below.
	Collectors []SflowCollector `pulumi:"collectors"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Source IP address for sFlow agent.
	SourceIp *string `pulumi:"sourceIp"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SflowState struct {
	// IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
	CollectorIp pulumi.StringPtrInput
	// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
	CollectorPort pulumi.IntPtrInput
	// sFlow collectors. The structure of `collectors` block is documented below.
	Collectors SflowCollectorArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Source IP address for sFlow agent.
	SourceIp pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SflowState) ElementType() reflect.Type {
	return reflect.TypeOf((*sflowState)(nil)).Elem()
}

type sflowArgs struct {
	// IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
	CollectorIp string `pulumi:"collectorIp"`
	// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
	CollectorPort *int `pulumi:"collectorPort"`
	// sFlow collectors. The structure of `collectors` block is documented below.
	Collectors []SflowCollector `pulumi:"collectors"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Source IP address for sFlow agent.
	SourceIp *string `pulumi:"sourceIp"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Sflow resource.
type SflowArgs struct {
	// IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
	CollectorIp pulumi.StringInput
	// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
	CollectorPort pulumi.IntPtrInput
	// sFlow collectors. The structure of `collectors` block is documented below.
	Collectors SflowCollectorArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Source IP address for sFlow agent.
	SourceIp pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SflowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sflowArgs)(nil)).Elem()
}

type SflowInput interface {
	pulumi.Input

	ToSflowOutput() SflowOutput
	ToSflowOutputWithContext(ctx context.Context) SflowOutput
}

func (*Sflow) ElementType() reflect.Type {
	return reflect.TypeOf((**Sflow)(nil)).Elem()
}

func (i *Sflow) ToSflowOutput() SflowOutput {
	return i.ToSflowOutputWithContext(context.Background())
}

func (i *Sflow) ToSflowOutputWithContext(ctx context.Context) SflowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SflowOutput)
}

// SflowArrayInput is an input type that accepts SflowArray and SflowArrayOutput values.
// You can construct a concrete instance of `SflowArrayInput` via:
//
//	SflowArray{ SflowArgs{...} }
type SflowArrayInput interface {
	pulumi.Input

	ToSflowArrayOutput() SflowArrayOutput
	ToSflowArrayOutputWithContext(context.Context) SflowArrayOutput
}

type SflowArray []SflowInput

func (SflowArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sflow)(nil)).Elem()
}

func (i SflowArray) ToSflowArrayOutput() SflowArrayOutput {
	return i.ToSflowArrayOutputWithContext(context.Background())
}

func (i SflowArray) ToSflowArrayOutputWithContext(ctx context.Context) SflowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SflowArrayOutput)
}

// SflowMapInput is an input type that accepts SflowMap and SflowMapOutput values.
// You can construct a concrete instance of `SflowMapInput` via:
//
//	SflowMap{ "key": SflowArgs{...} }
type SflowMapInput interface {
	pulumi.Input

	ToSflowMapOutput() SflowMapOutput
	ToSflowMapOutputWithContext(context.Context) SflowMapOutput
}

type SflowMap map[string]SflowInput

func (SflowMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sflow)(nil)).Elem()
}

func (i SflowMap) ToSflowMapOutput() SflowMapOutput {
	return i.ToSflowMapOutputWithContext(context.Background())
}

func (i SflowMap) ToSflowMapOutputWithContext(ctx context.Context) SflowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SflowMapOutput)
}

type SflowOutput struct{ *pulumi.OutputState }

func (SflowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sflow)(nil)).Elem()
}

func (o SflowOutput) ToSflowOutput() SflowOutput {
	return o
}

func (o SflowOutput) ToSflowOutputWithContext(ctx context.Context) SflowOutput {
	return o
}

// IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
func (o SflowOutput) CollectorIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Sflow) pulumi.StringOutput { return v.CollectorIp }).(pulumi.StringOutput)
}

// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
func (o SflowOutput) CollectorPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Sflow) pulumi.IntOutput { return v.CollectorPort }).(pulumi.IntOutput)
}

// sFlow collectors. The structure of `collectors` block is documented below.
func (o SflowOutput) Collectors() SflowCollectorArrayOutput {
	return o.ApplyT(func(v *Sflow) SflowCollectorArrayOutput { return v.Collectors }).(SflowCollectorArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SflowOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sflow) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o SflowOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sflow) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Specify outgoing interface to reach server.
func (o SflowOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Sflow) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
func (o SflowOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *Sflow) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

// Source IP address for sFlow agent.
func (o SflowOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Sflow) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SflowOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sflow) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SflowArrayOutput struct{ *pulumi.OutputState }

func (SflowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sflow)(nil)).Elem()
}

func (o SflowArrayOutput) ToSflowArrayOutput() SflowArrayOutput {
	return o
}

func (o SflowArrayOutput) ToSflowArrayOutputWithContext(ctx context.Context) SflowArrayOutput {
	return o
}

func (o SflowArrayOutput) Index(i pulumi.IntInput) SflowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Sflow {
		return vs[0].([]*Sflow)[vs[1].(int)]
	}).(SflowOutput)
}

type SflowMapOutput struct{ *pulumi.OutputState }

func (SflowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sflow)(nil)).Elem()
}

func (o SflowMapOutput) ToSflowMapOutput() SflowMapOutput {
	return o
}

func (o SflowMapOutput) ToSflowMapOutputWithContext(ctx context.Context) SflowMapOutput {
	return o
}

func (o SflowMapOutput) MapIndex(k pulumi.StringInput) SflowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Sflow {
		return vs[0].(map[string]*Sflow)[vs[1].(string)]
	}).(SflowOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SflowInput)(nil)).Elem(), &Sflow{})
	pulumi.RegisterInputType(reflect.TypeOf((*SflowArrayInput)(nil)).Elem(), SflowArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SflowMapInput)(nil)).Elem(), SflowMap{})
	pulumi.RegisterOutputType(SflowOutput{})
	pulumi.RegisterOutputType(SflowArrayOutput{})
	pulumi.RegisterOutputType(SflowMapOutput{})
}
