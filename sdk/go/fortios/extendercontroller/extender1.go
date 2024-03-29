// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package extendercontroller

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Extender controller configuration.
// This resource will be deprecated. For FortiOS Version >= 7.2.1, using `extensioncontroller.Extender`. For FortiOS version < 7.2.1, see `extendercontroller.Extender`
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/extendercontroller"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := extendercontroller.NewExtender1(ctx, "trname", &extendercontroller.Extender1Args{
//				Authorized: pulumi.String("disable"),
//				ControllerReport: &extendercontroller.Extender1ControllerReportArgs{
//					Interval:        pulumi.Int(300),
//					SignalThreshold: pulumi.Int(10),
//					Status:          pulumi.String("disable"),
//				},
//				ExtName: pulumi.String("2932"),
//				Fosid:   pulumi.String("FX201E5919004031"),
//				Modem1: &extendercontroller.Extender1Modem1Args{
//					AutoSwitch: &extendercontroller.Extender1Modem1AutoSwitchArgs{
//						Dataplan:            pulumi.String("disable"),
//						Disconnect:          pulumi.String("disable"),
//						DisconnectPeriod:    pulumi.Int(600),
//						DisconnectThreshold: pulumi.Int(3),
//						Signal:              pulumi.String("disable"),
//						SwitchBack:          pulumi.String("timer"),
//						SwitchBackTime:      pulumi.String("00:01"),
//						SwitchBackTimer:     pulumi.Int(86400),
//					},
//					ConnStatus:    pulumi.Int(0),
//					DefaultSim:    pulumi.String("sim2"),
//					Gps:           pulumi.String("enable"),
//					RedundantIntf: pulumi.String("s1"),
//					RedundantMode: pulumi.String("enable"),
//					Sim1Pin:       pulumi.String("disable"),
//					Sim1PinCode:   pulumi.String("testpincode"),
//					Sim2Pin:       pulumi.String("disable"),
//				},
//				Modem2: &extendercontroller.Extender1Modem2Args{
//					AutoSwitch: &extendercontroller.Extender1Modem2AutoSwitchArgs{
//						Dataplan:            pulumi.String("disable"),
//						Disconnect:          pulumi.String("disable"),
//						DisconnectPeriod:    pulumi.Int(600),
//						DisconnectThreshold: pulumi.Int(3),
//						Signal:              pulumi.String("disable"),
//						SwitchBackTime:      pulumi.String("00:01"),
//						SwitchBackTimer:     pulumi.Int(86400),
//					},
//					ConnStatus:    pulumi.Int(0),
//					DefaultSim:    pulumi.String("sim1"),
//					Gps:           pulumi.String("enable"),
//					RedundantMode: pulumi.String("disable"),
//					Sim1Pin:       pulumi.String("disable"),
//					Sim2Pin:       pulumi.String("disable"),
//				},
//				Vdom: pulumi.Int(0),
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
// ExtenderController Extender1 can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:extendercontroller/extender1:Extender1 labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:extendercontroller/extender1:Extender1 labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Extender1 struct {
	pulumi.CustomResourceState

	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized pulumi.StringOutput `pulumi:"authorized"`
	// FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
	ControllerReport Extender1ControllerReportOutput `pulumi:"controllerReport"`
	// Description.
	Description pulumi.StringOutput `pulumi:"description"`
	// FortiExtender name.
	ExtName pulumi.StringOutput `pulumi:"extName"`
	// FortiExtender serial number.
	Fosid pulumi.StringOutput `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// FortiExtender login password.
	LoginPassword pulumi.StringPtrOutput `pulumi:"loginPassword"`
	// Configuration options for modem 1. The structure of `modem1` block is documented below.
	Modem1 Extender1Modem1Output `pulumi:"modem1"`
	// Configuration options for modem 2. The structure of `modem2` block is documented below.
	Modem2 Extender1Modem2Output `pulumi:"modem2"`
	// FortiExtender entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// VDOM
	Vdom pulumi.IntOutput `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewExtender1 registers a new resource with the given unique name, arguments, and options.
func NewExtender1(ctx *pulumi.Context,
	name string, args *Extender1Args, opts ...pulumi.ResourceOption) (*Extender1, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Authorized == nil {
		return nil, errors.New("invalid value for required argument 'Authorized'")
	}
	if args.LoginPassword != nil {
		args.LoginPassword = pulumi.ToSecret(args.LoginPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"loginPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Extender1
	err := ctx.RegisterResource("fortios:extendercontroller/extender1:Extender1", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExtender1 gets an existing Extender1 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExtender1(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Extender1State, opts ...pulumi.ResourceOption) (*Extender1, error) {
	var resource Extender1
	err := ctx.ReadResource("fortios:extendercontroller/extender1:Extender1", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Extender1 resources.
type extender1State struct {
	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized *string `pulumi:"authorized"`
	// FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
	ControllerReport *Extender1ControllerReport `pulumi:"controllerReport"`
	// Description.
	Description *string `pulumi:"description"`
	// FortiExtender name.
	ExtName *string `pulumi:"extName"`
	// FortiExtender serial number.
	Fosid *string `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// FortiExtender login password.
	LoginPassword *string `pulumi:"loginPassword"`
	// Configuration options for modem 1. The structure of `modem1` block is documented below.
	Modem1 *Extender1Modem1 `pulumi:"modem1"`
	// Configuration options for modem 2. The structure of `modem2` block is documented below.
	Modem2 *Extender1Modem2 `pulumi:"modem2"`
	// FortiExtender entry name.
	Name *string `pulumi:"name"`
	// VDOM
	Vdom *int `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Extender1State struct {
	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized pulumi.StringPtrInput
	// FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
	ControllerReport Extender1ControllerReportPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// FortiExtender name.
	ExtName pulumi.StringPtrInput
	// FortiExtender serial number.
	Fosid pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// FortiExtender login password.
	LoginPassword pulumi.StringPtrInput
	// Configuration options for modem 1. The structure of `modem1` block is documented below.
	Modem1 Extender1Modem1PtrInput
	// Configuration options for modem 2. The structure of `modem2` block is documented below.
	Modem2 Extender1Modem2PtrInput
	// FortiExtender entry name.
	Name pulumi.StringPtrInput
	// VDOM
	Vdom pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Extender1State) ElementType() reflect.Type {
	return reflect.TypeOf((*extender1State)(nil)).Elem()
}

type extender1Args struct {
	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized string `pulumi:"authorized"`
	// FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
	ControllerReport *Extender1ControllerReport `pulumi:"controllerReport"`
	// Description.
	Description *string `pulumi:"description"`
	// FortiExtender name.
	ExtName *string `pulumi:"extName"`
	// FortiExtender serial number.
	Fosid *string `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// FortiExtender login password.
	LoginPassword *string `pulumi:"loginPassword"`
	// Configuration options for modem 1. The structure of `modem1` block is documented below.
	Modem1 *Extender1Modem1 `pulumi:"modem1"`
	// Configuration options for modem 2. The structure of `modem2` block is documented below.
	Modem2 *Extender1Modem2 `pulumi:"modem2"`
	// FortiExtender entry name.
	Name *string `pulumi:"name"`
	// VDOM
	Vdom *int `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Extender1 resource.
type Extender1Args struct {
	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized pulumi.StringInput
	// FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
	ControllerReport Extender1ControllerReportPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// FortiExtender name.
	ExtName pulumi.StringPtrInput
	// FortiExtender serial number.
	Fosid pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// FortiExtender login password.
	LoginPassword pulumi.StringPtrInput
	// Configuration options for modem 1. The structure of `modem1` block is documented below.
	Modem1 Extender1Modem1PtrInput
	// Configuration options for modem 2. The structure of `modem2` block is documented below.
	Modem2 Extender1Modem2PtrInput
	// FortiExtender entry name.
	Name pulumi.StringPtrInput
	// VDOM
	Vdom pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Extender1Args) ElementType() reflect.Type {
	return reflect.TypeOf((*extender1Args)(nil)).Elem()
}

type Extender1Input interface {
	pulumi.Input

	ToExtender1Output() Extender1Output
	ToExtender1OutputWithContext(ctx context.Context) Extender1Output
}

func (*Extender1) ElementType() reflect.Type {
	return reflect.TypeOf((**Extender1)(nil)).Elem()
}

func (i *Extender1) ToExtender1Output() Extender1Output {
	return i.ToExtender1OutputWithContext(context.Background())
}

func (i *Extender1) ToExtender1OutputWithContext(ctx context.Context) Extender1Output {
	return pulumi.ToOutputWithContext(ctx, i).(Extender1Output)
}

// Extender1ArrayInput is an input type that accepts Extender1Array and Extender1ArrayOutput values.
// You can construct a concrete instance of `Extender1ArrayInput` via:
//
//	Extender1Array{ Extender1Args{...} }
type Extender1ArrayInput interface {
	pulumi.Input

	ToExtender1ArrayOutput() Extender1ArrayOutput
	ToExtender1ArrayOutputWithContext(context.Context) Extender1ArrayOutput
}

type Extender1Array []Extender1Input

func (Extender1Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Extender1)(nil)).Elem()
}

func (i Extender1Array) ToExtender1ArrayOutput() Extender1ArrayOutput {
	return i.ToExtender1ArrayOutputWithContext(context.Background())
}

func (i Extender1Array) ToExtender1ArrayOutputWithContext(ctx context.Context) Extender1ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Extender1ArrayOutput)
}

// Extender1MapInput is an input type that accepts Extender1Map and Extender1MapOutput values.
// You can construct a concrete instance of `Extender1MapInput` via:
//
//	Extender1Map{ "key": Extender1Args{...} }
type Extender1MapInput interface {
	pulumi.Input

	ToExtender1MapOutput() Extender1MapOutput
	ToExtender1MapOutputWithContext(context.Context) Extender1MapOutput
}

type Extender1Map map[string]Extender1Input

func (Extender1Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Extender1)(nil)).Elem()
}

func (i Extender1Map) ToExtender1MapOutput() Extender1MapOutput {
	return i.ToExtender1MapOutputWithContext(context.Background())
}

func (i Extender1Map) ToExtender1MapOutputWithContext(ctx context.Context) Extender1MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Extender1MapOutput)
}

type Extender1Output struct{ *pulumi.OutputState }

func (Extender1Output) ElementType() reflect.Type {
	return reflect.TypeOf((**Extender1)(nil)).Elem()
}

func (o Extender1Output) ToExtender1Output() Extender1Output {
	return o
}

func (o Extender1Output) ToExtender1OutputWithContext(ctx context.Context) Extender1Output {
	return o
}

// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
func (o Extender1Output) Authorized() pulumi.StringOutput {
	return o.ApplyT(func(v *Extender1) pulumi.StringOutput { return v.Authorized }).(pulumi.StringOutput)
}

// FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
func (o Extender1Output) ControllerReport() Extender1ControllerReportOutput {
	return o.ApplyT(func(v *Extender1) Extender1ControllerReportOutput { return v.ControllerReport }).(Extender1ControllerReportOutput)
}

// Description.
func (o Extender1Output) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Extender1) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// FortiExtender name.
func (o Extender1Output) ExtName() pulumi.StringOutput {
	return o.ApplyT(func(v *Extender1) pulumi.StringOutput { return v.ExtName }).(pulumi.StringOutput)
}

// FortiExtender serial number.
func (o Extender1Output) Fosid() pulumi.StringOutput {
	return o.ApplyT(func(v *Extender1) pulumi.StringOutput { return v.Fosid }).(pulumi.StringOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o Extender1Output) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extender1) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// FortiExtender login password.
func (o Extender1Output) LoginPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extender1) pulumi.StringPtrOutput { return v.LoginPassword }).(pulumi.StringPtrOutput)
}

// Configuration options for modem 1. The structure of `modem1` block is documented below.
func (o Extender1Output) Modem1() Extender1Modem1Output {
	return o.ApplyT(func(v *Extender1) Extender1Modem1Output { return v.Modem1 }).(Extender1Modem1Output)
}

// Configuration options for modem 2. The structure of `modem2` block is documented below.
func (o Extender1Output) Modem2() Extender1Modem2Output {
	return o.ApplyT(func(v *Extender1) Extender1Modem2Output { return v.Modem2 }).(Extender1Modem2Output)
}

// FortiExtender entry name.
func (o Extender1Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Extender1) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// VDOM
func (o Extender1Output) Vdom() pulumi.IntOutput {
	return o.ApplyT(func(v *Extender1) pulumi.IntOutput { return v.Vdom }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Extender1Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extender1) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type Extender1ArrayOutput struct{ *pulumi.OutputState }

func (Extender1ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Extender1)(nil)).Elem()
}

func (o Extender1ArrayOutput) ToExtender1ArrayOutput() Extender1ArrayOutput {
	return o
}

func (o Extender1ArrayOutput) ToExtender1ArrayOutputWithContext(ctx context.Context) Extender1ArrayOutput {
	return o
}

func (o Extender1ArrayOutput) Index(i pulumi.IntInput) Extender1Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Extender1 {
		return vs[0].([]*Extender1)[vs[1].(int)]
	}).(Extender1Output)
}

type Extender1MapOutput struct{ *pulumi.OutputState }

func (Extender1MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Extender1)(nil)).Elem()
}

func (o Extender1MapOutput) ToExtender1MapOutput() Extender1MapOutput {
	return o
}

func (o Extender1MapOutput) ToExtender1MapOutputWithContext(ctx context.Context) Extender1MapOutput {
	return o
}

func (o Extender1MapOutput) MapIndex(k pulumi.StringInput) Extender1Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Extender1 {
		return vs[0].(map[string]*Extender1)[vs[1].(string)]
	}).(Extender1Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Extender1Input)(nil)).Elem(), &Extender1{})
	pulumi.RegisterInputType(reflect.TypeOf((*Extender1ArrayInput)(nil)).Elem(), Extender1Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*Extender1MapInput)(nil)).Elem(), Extender1Map{})
	pulumi.RegisterOutputType(Extender1Output{})
	pulumi.RegisterOutputType(Extender1ArrayOutput{})
	pulumi.RegisterOutputType(Extender1MapOutput{})
}
