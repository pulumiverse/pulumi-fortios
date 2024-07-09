// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package schedule

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Onetime schedule configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/firewall"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firewall.NewOnetime(ctx, "trname", &firewall.OnetimeArgs{
//				Color:          pulumi.Int(0),
//				End:            pulumi.String("00:00 2020/12/12"),
//				ExpirationDays: pulumi.Int(2),
//				Start:          pulumi.String("00:00 2010/12/12"),
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
// FirewallSchedule Onetime can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/schedule/onetime:Onetime labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/schedule/onetime:Onetime labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Onetime struct {
	pulumi.CustomResourceState

	// Color of icon on the GUI.
	Color pulumi.IntOutput `pulumi:"color"`
	// Schedule end date and time, format hh:mm yyyy/mm/dd.
	End pulumi.StringOutput `pulumi:"end"`
	// Schedule end date and time, in epoch format.
	EndUtc pulumi.StringOutput `pulumi:"endUtc"`
	// Write an event log message this many days before the schedule expires.
	ExpirationDays pulumi.IntOutput `pulumi:"expirationDays"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringOutput `pulumi:"fabricObject"`
	// Onetime schedule name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Schedule start date and time, format hh:mm yyyy/mm/dd.
	Start pulumi.StringOutput `pulumi:"start"`
	// Schedule start date and time, in epoch format.
	StartUtc pulumi.StringOutput `pulumi:"startUtc"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewOnetime registers a new resource with the given unique name, arguments, and options.
func NewOnetime(ctx *pulumi.Context,
	name string, args *OnetimeArgs, opts ...pulumi.ResourceOption) (*Onetime, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.End == nil {
		return nil, errors.New("invalid value for required argument 'End'")
	}
	if args.Start == nil {
		return nil, errors.New("invalid value for required argument 'Start'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Onetime
	err := ctx.RegisterResource("fortios:firewall/schedule/onetime:Onetime", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOnetime gets an existing Onetime resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOnetime(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OnetimeState, opts ...pulumi.ResourceOption) (*Onetime, error) {
	var resource Onetime
	err := ctx.ReadResource("fortios:firewall/schedule/onetime:Onetime", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Onetime resources.
type onetimeState struct {
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Schedule end date and time, format hh:mm yyyy/mm/dd.
	End *string `pulumi:"end"`
	// Schedule end date and time, in epoch format.
	EndUtc *string `pulumi:"endUtc"`
	// Write an event log message this many days before the schedule expires.
	ExpirationDays *int `pulumi:"expirationDays"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// Onetime schedule name.
	Name *string `pulumi:"name"`
	// Schedule start date and time, format hh:mm yyyy/mm/dd.
	Start *string `pulumi:"start"`
	// Schedule start date and time, in epoch format.
	StartUtc *string `pulumi:"startUtc"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type OnetimeState struct {
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Schedule end date and time, format hh:mm yyyy/mm/dd.
	End pulumi.StringPtrInput
	// Schedule end date and time, in epoch format.
	EndUtc pulumi.StringPtrInput
	// Write an event log message this many days before the schedule expires.
	ExpirationDays pulumi.IntPtrInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// Onetime schedule name.
	Name pulumi.StringPtrInput
	// Schedule start date and time, format hh:mm yyyy/mm/dd.
	Start pulumi.StringPtrInput
	// Schedule start date and time, in epoch format.
	StartUtc pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (OnetimeState) ElementType() reflect.Type {
	return reflect.TypeOf((*onetimeState)(nil)).Elem()
}

type onetimeArgs struct {
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Schedule end date and time, format hh:mm yyyy/mm/dd.
	End string `pulumi:"end"`
	// Schedule end date and time, in epoch format.
	EndUtc *string `pulumi:"endUtc"`
	// Write an event log message this many days before the schedule expires.
	ExpirationDays *int `pulumi:"expirationDays"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// Onetime schedule name.
	Name *string `pulumi:"name"`
	// Schedule start date and time, format hh:mm yyyy/mm/dd.
	Start string `pulumi:"start"`
	// Schedule start date and time, in epoch format.
	StartUtc *string `pulumi:"startUtc"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Onetime resource.
type OnetimeArgs struct {
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Schedule end date and time, format hh:mm yyyy/mm/dd.
	End pulumi.StringInput
	// Schedule end date and time, in epoch format.
	EndUtc pulumi.StringPtrInput
	// Write an event log message this many days before the schedule expires.
	ExpirationDays pulumi.IntPtrInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// Onetime schedule name.
	Name pulumi.StringPtrInput
	// Schedule start date and time, format hh:mm yyyy/mm/dd.
	Start pulumi.StringInput
	// Schedule start date and time, in epoch format.
	StartUtc pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (OnetimeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*onetimeArgs)(nil)).Elem()
}

type OnetimeInput interface {
	pulumi.Input

	ToOnetimeOutput() OnetimeOutput
	ToOnetimeOutputWithContext(ctx context.Context) OnetimeOutput
}

func (*Onetime) ElementType() reflect.Type {
	return reflect.TypeOf((**Onetime)(nil)).Elem()
}

func (i *Onetime) ToOnetimeOutput() OnetimeOutput {
	return i.ToOnetimeOutputWithContext(context.Background())
}

func (i *Onetime) ToOnetimeOutputWithContext(ctx context.Context) OnetimeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnetimeOutput)
}

// OnetimeArrayInput is an input type that accepts OnetimeArray and OnetimeArrayOutput values.
// You can construct a concrete instance of `OnetimeArrayInput` via:
//
//	OnetimeArray{ OnetimeArgs{...} }
type OnetimeArrayInput interface {
	pulumi.Input

	ToOnetimeArrayOutput() OnetimeArrayOutput
	ToOnetimeArrayOutputWithContext(context.Context) OnetimeArrayOutput
}

type OnetimeArray []OnetimeInput

func (OnetimeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Onetime)(nil)).Elem()
}

func (i OnetimeArray) ToOnetimeArrayOutput() OnetimeArrayOutput {
	return i.ToOnetimeArrayOutputWithContext(context.Background())
}

func (i OnetimeArray) ToOnetimeArrayOutputWithContext(ctx context.Context) OnetimeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnetimeArrayOutput)
}

// OnetimeMapInput is an input type that accepts OnetimeMap and OnetimeMapOutput values.
// You can construct a concrete instance of `OnetimeMapInput` via:
//
//	OnetimeMap{ "key": OnetimeArgs{...} }
type OnetimeMapInput interface {
	pulumi.Input

	ToOnetimeMapOutput() OnetimeMapOutput
	ToOnetimeMapOutputWithContext(context.Context) OnetimeMapOutput
}

type OnetimeMap map[string]OnetimeInput

func (OnetimeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Onetime)(nil)).Elem()
}

func (i OnetimeMap) ToOnetimeMapOutput() OnetimeMapOutput {
	return i.ToOnetimeMapOutputWithContext(context.Background())
}

func (i OnetimeMap) ToOnetimeMapOutputWithContext(ctx context.Context) OnetimeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnetimeMapOutput)
}

type OnetimeOutput struct{ *pulumi.OutputState }

func (OnetimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Onetime)(nil)).Elem()
}

func (o OnetimeOutput) ToOnetimeOutput() OnetimeOutput {
	return o
}

func (o OnetimeOutput) ToOnetimeOutputWithContext(ctx context.Context) OnetimeOutput {
	return o
}

// Color of icon on the GUI.
func (o OnetimeOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *Onetime) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

// Schedule end date and time, format hh:mm yyyy/mm/dd.
func (o OnetimeOutput) End() pulumi.StringOutput {
	return o.ApplyT(func(v *Onetime) pulumi.StringOutput { return v.End }).(pulumi.StringOutput)
}

// Schedule end date and time, in epoch format.
func (o OnetimeOutput) EndUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *Onetime) pulumi.StringOutput { return v.EndUtc }).(pulumi.StringOutput)
}

// Write an event log message this many days before the schedule expires.
func (o OnetimeOutput) ExpirationDays() pulumi.IntOutput {
	return o.ApplyT(func(v *Onetime) pulumi.IntOutput { return v.ExpirationDays }).(pulumi.IntOutput)
}

// Security Fabric global object setting. Valid values: `enable`, `disable`.
func (o OnetimeOutput) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v *Onetime) pulumi.StringOutput { return v.FabricObject }).(pulumi.StringOutput)
}

// Onetime schedule name.
func (o OnetimeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Onetime) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Schedule start date and time, format hh:mm yyyy/mm/dd.
func (o OnetimeOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v *Onetime) pulumi.StringOutput { return v.Start }).(pulumi.StringOutput)
}

// Schedule start date and time, in epoch format.
func (o OnetimeOutput) StartUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *Onetime) pulumi.StringOutput { return v.StartUtc }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o OnetimeOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Onetime) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type OnetimeArrayOutput struct{ *pulumi.OutputState }

func (OnetimeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Onetime)(nil)).Elem()
}

func (o OnetimeArrayOutput) ToOnetimeArrayOutput() OnetimeArrayOutput {
	return o
}

func (o OnetimeArrayOutput) ToOnetimeArrayOutputWithContext(ctx context.Context) OnetimeArrayOutput {
	return o
}

func (o OnetimeArrayOutput) Index(i pulumi.IntInput) OnetimeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Onetime {
		return vs[0].([]*Onetime)[vs[1].(int)]
	}).(OnetimeOutput)
}

type OnetimeMapOutput struct{ *pulumi.OutputState }

func (OnetimeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Onetime)(nil)).Elem()
}

func (o OnetimeMapOutput) ToOnetimeMapOutput() OnetimeMapOutput {
	return o
}

func (o OnetimeMapOutput) ToOnetimeMapOutputWithContext(ctx context.Context) OnetimeMapOutput {
	return o
}

func (o OnetimeMapOutput) MapIndex(k pulumi.StringInput) OnetimeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Onetime {
		return vs[0].(map[string]*Onetime)[vs[1].(string)]
	}).(OnetimeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OnetimeInput)(nil)).Elem(), &Onetime{})
	pulumi.RegisterInputType(reflect.TypeOf((*OnetimeArrayInput)(nil)).Elem(), OnetimeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OnetimeMapInput)(nil)).Elem(), OnetimeMap{})
	pulumi.RegisterOutputType(OnetimeOutput{})
	pulumi.RegisterOutputType(OnetimeArrayOutput{})
	pulumi.RegisterOutputType(OnetimeMapOutput{})
}
