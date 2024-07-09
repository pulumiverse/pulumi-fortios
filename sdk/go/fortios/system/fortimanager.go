// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure FortiManager. Applies to FortiOS Version `<= 7.0.1`.
//
// By design considerations, the feature is using the system.Centralmanagement resource as documented below.
//
// ## Example
//
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
//			_, err := system.NewCentralmanagement(ctx, "trname", &system.CentralmanagementArgs{
//				AllowMonitor:               pulumi.String("enable"),
//				AllowPushConfiguration:     pulumi.String("enable"),
//				AllowPushFirmware:          pulumi.String("enable"),
//				AllowRemoteFirmwareUpgrade: pulumi.String("enable"),
//				EncAlgorithm:               pulumi.String("high"),
//				Fmg:                        pulumi.String("\"192.168.52.177\""),
//				IncludeDefaultServers:      pulumi.String("enable"),
//				Mode:                       pulumi.String("normal"),
//				Type:                       pulumi.String("fortimanager"),
//				Vdom:                       pulumi.String("root"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Fortimanager struct {
	pulumi.CustomResourceState

	CentralManagement                pulumi.StringOutput `pulumi:"centralManagement"`
	CentralMgmtAutoBackup            pulumi.StringOutput `pulumi:"centralMgmtAutoBackup"`
	CentralMgmtScheduleConfigRestore pulumi.StringOutput `pulumi:"centralMgmtScheduleConfigRestore"`
	CentralMgmtScheduleScriptRestore pulumi.StringOutput `pulumi:"centralMgmtScheduleScriptRestore"`
	Ip                               pulumi.StringOutput `pulumi:"ip"`
	Ipsec                            pulumi.StringOutput `pulumi:"ipsec"`
	Vdom                             pulumi.StringOutput `pulumi:"vdom"`
	Vdomparam                        pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewFortimanager registers a new resource with the given unique name, arguments, and options.
func NewFortimanager(ctx *pulumi.Context,
	name string, args *FortimanagerArgs, opts ...pulumi.ResourceOption) (*Fortimanager, error) {
	if args == nil {
		args = &FortimanagerArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Fortimanager
	err := ctx.RegisterResource("fortios:system/fortimanager:Fortimanager", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortimanager gets an existing Fortimanager resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortimanager(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortimanagerState, opts ...pulumi.ResourceOption) (*Fortimanager, error) {
	var resource Fortimanager
	err := ctx.ReadResource("fortios:system/fortimanager:Fortimanager", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fortimanager resources.
type fortimanagerState struct {
	CentralManagement                *string `pulumi:"centralManagement"`
	CentralMgmtAutoBackup            *string `pulumi:"centralMgmtAutoBackup"`
	CentralMgmtScheduleConfigRestore *string `pulumi:"centralMgmtScheduleConfigRestore"`
	CentralMgmtScheduleScriptRestore *string `pulumi:"centralMgmtScheduleScriptRestore"`
	Ip                               *string `pulumi:"ip"`
	Ipsec                            *string `pulumi:"ipsec"`
	Vdom                             *string `pulumi:"vdom"`
	Vdomparam                        *string `pulumi:"vdomparam"`
}

type FortimanagerState struct {
	CentralManagement                pulumi.StringPtrInput
	CentralMgmtAutoBackup            pulumi.StringPtrInput
	CentralMgmtScheduleConfigRestore pulumi.StringPtrInput
	CentralMgmtScheduleScriptRestore pulumi.StringPtrInput
	Ip                               pulumi.StringPtrInput
	Ipsec                            pulumi.StringPtrInput
	Vdom                             pulumi.StringPtrInput
	Vdomparam                        pulumi.StringPtrInput
}

func (FortimanagerState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerState)(nil)).Elem()
}

type fortimanagerArgs struct {
	CentralManagement                *string `pulumi:"centralManagement"`
	CentralMgmtAutoBackup            *string `pulumi:"centralMgmtAutoBackup"`
	CentralMgmtScheduleConfigRestore *string `pulumi:"centralMgmtScheduleConfigRestore"`
	CentralMgmtScheduleScriptRestore *string `pulumi:"centralMgmtScheduleScriptRestore"`
	Ip                               *string `pulumi:"ip"`
	Ipsec                            *string `pulumi:"ipsec"`
	Vdom                             *string `pulumi:"vdom"`
	Vdomparam                        *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Fortimanager resource.
type FortimanagerArgs struct {
	CentralManagement                pulumi.StringPtrInput
	CentralMgmtAutoBackup            pulumi.StringPtrInput
	CentralMgmtScheduleConfigRestore pulumi.StringPtrInput
	CentralMgmtScheduleScriptRestore pulumi.StringPtrInput
	Ip                               pulumi.StringPtrInput
	Ipsec                            pulumi.StringPtrInput
	Vdom                             pulumi.StringPtrInput
	Vdomparam                        pulumi.StringPtrInput
}

func (FortimanagerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerArgs)(nil)).Elem()
}

type FortimanagerInput interface {
	pulumi.Input

	ToFortimanagerOutput() FortimanagerOutput
	ToFortimanagerOutputWithContext(ctx context.Context) FortimanagerOutput
}

func (*Fortimanager) ElementType() reflect.Type {
	return reflect.TypeOf((**Fortimanager)(nil)).Elem()
}

func (i *Fortimanager) ToFortimanagerOutput() FortimanagerOutput {
	return i.ToFortimanagerOutputWithContext(context.Background())
}

func (i *Fortimanager) ToFortimanagerOutputWithContext(ctx context.Context) FortimanagerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerOutput)
}

// FortimanagerArrayInput is an input type that accepts FortimanagerArray and FortimanagerArrayOutput values.
// You can construct a concrete instance of `FortimanagerArrayInput` via:
//
//	FortimanagerArray{ FortimanagerArgs{...} }
type FortimanagerArrayInput interface {
	pulumi.Input

	ToFortimanagerArrayOutput() FortimanagerArrayOutput
	ToFortimanagerArrayOutputWithContext(context.Context) FortimanagerArrayOutput
}

type FortimanagerArray []FortimanagerInput

func (FortimanagerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fortimanager)(nil)).Elem()
}

func (i FortimanagerArray) ToFortimanagerArrayOutput() FortimanagerArrayOutput {
	return i.ToFortimanagerArrayOutputWithContext(context.Background())
}

func (i FortimanagerArray) ToFortimanagerArrayOutputWithContext(ctx context.Context) FortimanagerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerArrayOutput)
}

// FortimanagerMapInput is an input type that accepts FortimanagerMap and FortimanagerMapOutput values.
// You can construct a concrete instance of `FortimanagerMapInput` via:
//
//	FortimanagerMap{ "key": FortimanagerArgs{...} }
type FortimanagerMapInput interface {
	pulumi.Input

	ToFortimanagerMapOutput() FortimanagerMapOutput
	ToFortimanagerMapOutputWithContext(context.Context) FortimanagerMapOutput
}

type FortimanagerMap map[string]FortimanagerInput

func (FortimanagerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fortimanager)(nil)).Elem()
}

func (i FortimanagerMap) ToFortimanagerMapOutput() FortimanagerMapOutput {
	return i.ToFortimanagerMapOutputWithContext(context.Background())
}

func (i FortimanagerMap) ToFortimanagerMapOutputWithContext(ctx context.Context) FortimanagerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerMapOutput)
}

type FortimanagerOutput struct{ *pulumi.OutputState }

func (FortimanagerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fortimanager)(nil)).Elem()
}

func (o FortimanagerOutput) ToFortimanagerOutput() FortimanagerOutput {
	return o
}

func (o FortimanagerOutput) ToFortimanagerOutputWithContext(ctx context.Context) FortimanagerOutput {
	return o
}

func (o FortimanagerOutput) CentralManagement() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortimanager) pulumi.StringOutput { return v.CentralManagement }).(pulumi.StringOutput)
}

func (o FortimanagerOutput) CentralMgmtAutoBackup() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortimanager) pulumi.StringOutput { return v.CentralMgmtAutoBackup }).(pulumi.StringOutput)
}

func (o FortimanagerOutput) CentralMgmtScheduleConfigRestore() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortimanager) pulumi.StringOutput { return v.CentralMgmtScheduleConfigRestore }).(pulumi.StringOutput)
}

func (o FortimanagerOutput) CentralMgmtScheduleScriptRestore() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortimanager) pulumi.StringOutput { return v.CentralMgmtScheduleScriptRestore }).(pulumi.StringOutput)
}

func (o FortimanagerOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortimanager) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

func (o FortimanagerOutput) Ipsec() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortimanager) pulumi.StringOutput { return v.Ipsec }).(pulumi.StringOutput)
}

func (o FortimanagerOutput) Vdom() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortimanager) pulumi.StringOutput { return v.Vdom }).(pulumi.StringOutput)
}

func (o FortimanagerOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortimanager) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type FortimanagerArrayOutput struct{ *pulumi.OutputState }

func (FortimanagerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fortimanager)(nil)).Elem()
}

func (o FortimanagerArrayOutput) ToFortimanagerArrayOutput() FortimanagerArrayOutput {
	return o
}

func (o FortimanagerArrayOutput) ToFortimanagerArrayOutputWithContext(ctx context.Context) FortimanagerArrayOutput {
	return o
}

func (o FortimanagerArrayOutput) Index(i pulumi.IntInput) FortimanagerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Fortimanager {
		return vs[0].([]*Fortimanager)[vs[1].(int)]
	}).(FortimanagerOutput)
}

type FortimanagerMapOutput struct{ *pulumi.OutputState }

func (FortimanagerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fortimanager)(nil)).Elem()
}

func (o FortimanagerMapOutput) ToFortimanagerMapOutput() FortimanagerMapOutput {
	return o
}

func (o FortimanagerMapOutput) ToFortimanagerMapOutputWithContext(ctx context.Context) FortimanagerMapOutput {
	return o
}

func (o FortimanagerMapOutput) MapIndex(k pulumi.StringInput) FortimanagerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Fortimanager {
		return vs[0].(map[string]*Fortimanager)[vs[1].(string)]
	}).(FortimanagerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerInput)(nil)).Elem(), &Fortimanager{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerArrayInput)(nil)).Elem(), FortimanagerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerMapInput)(nil)).Elem(), FortimanagerMap{})
	pulumi.RegisterOutputType(FortimanagerOutput{})
	pulumi.RegisterOutputType(FortimanagerArrayOutput{})
	pulumi.RegisterOutputType(FortimanagerMapOutput{})
}
