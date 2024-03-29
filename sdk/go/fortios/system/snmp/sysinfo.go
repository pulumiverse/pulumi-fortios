// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package snmp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// SNMP system info configuration.
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
//			_, err := system.NewSysinfo(ctx, "trname", &system.SysinfoArgs{
//				Status:                 pulumi.String("disable"),
//				TrapHighCpuThreshold:   pulumi.Int(80),
//				TrapLogFullThreshold:   pulumi.Int(90),
//				TrapLowMemoryThreshold: pulumi.Int(80),
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
// SystemSnmp Sysinfo can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/snmp/sysinfo:Sysinfo labelname SystemSnmpSysinfo
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/snmp/sysinfo:Sysinfo labelname SystemSnmpSysinfo
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Sysinfo struct {
	pulumi.CustomResourceState

	// Contact information.
	ContactInfo pulumi.StringPtrOutput `pulumi:"contactInfo"`
	// System description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Local SNMP engineID string (maximum 24 characters).
	EngineId pulumi.StringOutput `pulumi:"engineId"`
	// Local SNMP engineID type (text/hex/mac). Valid values: `text`, `hex`, `mac`.
	EngineIdType pulumi.StringOutput `pulumi:"engineIdType"`
	// System location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Enable/disable SNMP. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Free memory usage when trap is sent.
	TrapFreeMemoryThreshold pulumi.IntOutput `pulumi:"trapFreeMemoryThreshold"`
	// Freeable memory usage when trap is sent.
	TrapFreeableMemoryThreshold pulumi.IntOutput `pulumi:"trapFreeableMemoryThreshold"`
	// CPU usage when trap is sent.
	TrapHighCpuThreshold pulumi.IntOutput `pulumi:"trapHighCpuThreshold"`
	// Log disk usage when trap is sent.
	TrapLogFullThreshold pulumi.IntOutput `pulumi:"trapLogFullThreshold"`
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold pulumi.IntOutput `pulumi:"trapLowMemoryThreshold"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSysinfo registers a new resource with the given unique name, arguments, and options.
func NewSysinfo(ctx *pulumi.Context,
	name string, args *SysinfoArgs, opts ...pulumi.ResourceOption) (*Sysinfo, error) {
	if args == nil {
		args = &SysinfoArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Sysinfo
	err := ctx.RegisterResource("fortios:system/snmp/sysinfo:Sysinfo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSysinfo gets an existing Sysinfo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSysinfo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SysinfoState, opts ...pulumi.ResourceOption) (*Sysinfo, error) {
	var resource Sysinfo
	err := ctx.ReadResource("fortios:system/snmp/sysinfo:Sysinfo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Sysinfo resources.
type sysinfoState struct {
	// Contact information.
	ContactInfo *string `pulumi:"contactInfo"`
	// System description.
	Description *string `pulumi:"description"`
	// Local SNMP engineID string (maximum 24 characters).
	EngineId *string `pulumi:"engineId"`
	// Local SNMP engineID type (text/hex/mac). Valid values: `text`, `hex`, `mac`.
	EngineIdType *string `pulumi:"engineIdType"`
	// System location.
	Location *string `pulumi:"location"`
	// Enable/disable SNMP. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Free memory usage when trap is sent.
	TrapFreeMemoryThreshold *int `pulumi:"trapFreeMemoryThreshold"`
	// Freeable memory usage when trap is sent.
	TrapFreeableMemoryThreshold *int `pulumi:"trapFreeableMemoryThreshold"`
	// CPU usage when trap is sent.
	TrapHighCpuThreshold *int `pulumi:"trapHighCpuThreshold"`
	// Log disk usage when trap is sent.
	TrapLogFullThreshold *int `pulumi:"trapLogFullThreshold"`
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold *int `pulumi:"trapLowMemoryThreshold"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SysinfoState struct {
	// Contact information.
	ContactInfo pulumi.StringPtrInput
	// System description.
	Description pulumi.StringPtrInput
	// Local SNMP engineID string (maximum 24 characters).
	EngineId pulumi.StringPtrInput
	// Local SNMP engineID type (text/hex/mac). Valid values: `text`, `hex`, `mac`.
	EngineIdType pulumi.StringPtrInput
	// System location.
	Location pulumi.StringPtrInput
	// Enable/disable SNMP. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Free memory usage when trap is sent.
	TrapFreeMemoryThreshold pulumi.IntPtrInput
	// Freeable memory usage when trap is sent.
	TrapFreeableMemoryThreshold pulumi.IntPtrInput
	// CPU usage when trap is sent.
	TrapHighCpuThreshold pulumi.IntPtrInput
	// Log disk usage when trap is sent.
	TrapLogFullThreshold pulumi.IntPtrInput
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SysinfoState) ElementType() reflect.Type {
	return reflect.TypeOf((*sysinfoState)(nil)).Elem()
}

type sysinfoArgs struct {
	// Contact information.
	ContactInfo *string `pulumi:"contactInfo"`
	// System description.
	Description *string `pulumi:"description"`
	// Local SNMP engineID string (maximum 24 characters).
	EngineId *string `pulumi:"engineId"`
	// Local SNMP engineID type (text/hex/mac). Valid values: `text`, `hex`, `mac`.
	EngineIdType *string `pulumi:"engineIdType"`
	// System location.
	Location *string `pulumi:"location"`
	// Enable/disable SNMP. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Free memory usage when trap is sent.
	TrapFreeMemoryThreshold *int `pulumi:"trapFreeMemoryThreshold"`
	// Freeable memory usage when trap is sent.
	TrapFreeableMemoryThreshold *int `pulumi:"trapFreeableMemoryThreshold"`
	// CPU usage when trap is sent.
	TrapHighCpuThreshold *int `pulumi:"trapHighCpuThreshold"`
	// Log disk usage when trap is sent.
	TrapLogFullThreshold *int `pulumi:"trapLogFullThreshold"`
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold *int `pulumi:"trapLowMemoryThreshold"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Sysinfo resource.
type SysinfoArgs struct {
	// Contact information.
	ContactInfo pulumi.StringPtrInput
	// System description.
	Description pulumi.StringPtrInput
	// Local SNMP engineID string (maximum 24 characters).
	EngineId pulumi.StringPtrInput
	// Local SNMP engineID type (text/hex/mac). Valid values: `text`, `hex`, `mac`.
	EngineIdType pulumi.StringPtrInput
	// System location.
	Location pulumi.StringPtrInput
	// Enable/disable SNMP. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Free memory usage when trap is sent.
	TrapFreeMemoryThreshold pulumi.IntPtrInput
	// Freeable memory usage when trap is sent.
	TrapFreeableMemoryThreshold pulumi.IntPtrInput
	// CPU usage when trap is sent.
	TrapHighCpuThreshold pulumi.IntPtrInput
	// Log disk usage when trap is sent.
	TrapLogFullThreshold pulumi.IntPtrInput
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SysinfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sysinfoArgs)(nil)).Elem()
}

type SysinfoInput interface {
	pulumi.Input

	ToSysinfoOutput() SysinfoOutput
	ToSysinfoOutputWithContext(ctx context.Context) SysinfoOutput
}

func (*Sysinfo) ElementType() reflect.Type {
	return reflect.TypeOf((**Sysinfo)(nil)).Elem()
}

func (i *Sysinfo) ToSysinfoOutput() SysinfoOutput {
	return i.ToSysinfoOutputWithContext(context.Background())
}

func (i *Sysinfo) ToSysinfoOutputWithContext(ctx context.Context) SysinfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SysinfoOutput)
}

// SysinfoArrayInput is an input type that accepts SysinfoArray and SysinfoArrayOutput values.
// You can construct a concrete instance of `SysinfoArrayInput` via:
//
//	SysinfoArray{ SysinfoArgs{...} }
type SysinfoArrayInput interface {
	pulumi.Input

	ToSysinfoArrayOutput() SysinfoArrayOutput
	ToSysinfoArrayOutputWithContext(context.Context) SysinfoArrayOutput
}

type SysinfoArray []SysinfoInput

func (SysinfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sysinfo)(nil)).Elem()
}

func (i SysinfoArray) ToSysinfoArrayOutput() SysinfoArrayOutput {
	return i.ToSysinfoArrayOutputWithContext(context.Background())
}

func (i SysinfoArray) ToSysinfoArrayOutputWithContext(ctx context.Context) SysinfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SysinfoArrayOutput)
}

// SysinfoMapInput is an input type that accepts SysinfoMap and SysinfoMapOutput values.
// You can construct a concrete instance of `SysinfoMapInput` via:
//
//	SysinfoMap{ "key": SysinfoArgs{...} }
type SysinfoMapInput interface {
	pulumi.Input

	ToSysinfoMapOutput() SysinfoMapOutput
	ToSysinfoMapOutputWithContext(context.Context) SysinfoMapOutput
}

type SysinfoMap map[string]SysinfoInput

func (SysinfoMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sysinfo)(nil)).Elem()
}

func (i SysinfoMap) ToSysinfoMapOutput() SysinfoMapOutput {
	return i.ToSysinfoMapOutputWithContext(context.Background())
}

func (i SysinfoMap) ToSysinfoMapOutputWithContext(ctx context.Context) SysinfoMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SysinfoMapOutput)
}

type SysinfoOutput struct{ *pulumi.OutputState }

func (SysinfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sysinfo)(nil)).Elem()
}

func (o SysinfoOutput) ToSysinfoOutput() SysinfoOutput {
	return o
}

func (o SysinfoOutput) ToSysinfoOutputWithContext(ctx context.Context) SysinfoOutput {
	return o
}

// Contact information.
func (o SysinfoOutput) ContactInfo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sysinfo) pulumi.StringPtrOutput { return v.ContactInfo }).(pulumi.StringPtrOutput)
}

// System description.
func (o SysinfoOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sysinfo) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Local SNMP engineID string (maximum 24 characters).
func (o SysinfoOutput) EngineId() pulumi.StringOutput {
	return o.ApplyT(func(v *Sysinfo) pulumi.StringOutput { return v.EngineId }).(pulumi.StringOutput)
}

// Local SNMP engineID type (text/hex/mac). Valid values: `text`, `hex`, `mac`.
func (o SysinfoOutput) EngineIdType() pulumi.StringOutput {
	return o.ApplyT(func(v *Sysinfo) pulumi.StringOutput { return v.EngineIdType }).(pulumi.StringOutput)
}

// System location.
func (o SysinfoOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sysinfo) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Enable/disable SNMP. Valid values: `enable`, `disable`.
func (o SysinfoOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Sysinfo) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Free memory usage when trap is sent.
func (o SysinfoOutput) TrapFreeMemoryThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *Sysinfo) pulumi.IntOutput { return v.TrapFreeMemoryThreshold }).(pulumi.IntOutput)
}

// Freeable memory usage when trap is sent.
func (o SysinfoOutput) TrapFreeableMemoryThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *Sysinfo) pulumi.IntOutput { return v.TrapFreeableMemoryThreshold }).(pulumi.IntOutput)
}

// CPU usage when trap is sent.
func (o SysinfoOutput) TrapHighCpuThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *Sysinfo) pulumi.IntOutput { return v.TrapHighCpuThreshold }).(pulumi.IntOutput)
}

// Log disk usage when trap is sent.
func (o SysinfoOutput) TrapLogFullThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *Sysinfo) pulumi.IntOutput { return v.TrapLogFullThreshold }).(pulumi.IntOutput)
}

// Memory usage when trap is sent.
func (o SysinfoOutput) TrapLowMemoryThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *Sysinfo) pulumi.IntOutput { return v.TrapLowMemoryThreshold }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SysinfoOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sysinfo) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SysinfoArrayOutput struct{ *pulumi.OutputState }

func (SysinfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sysinfo)(nil)).Elem()
}

func (o SysinfoArrayOutput) ToSysinfoArrayOutput() SysinfoArrayOutput {
	return o
}

func (o SysinfoArrayOutput) ToSysinfoArrayOutputWithContext(ctx context.Context) SysinfoArrayOutput {
	return o
}

func (o SysinfoArrayOutput) Index(i pulumi.IntInput) SysinfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Sysinfo {
		return vs[0].([]*Sysinfo)[vs[1].(int)]
	}).(SysinfoOutput)
}

type SysinfoMapOutput struct{ *pulumi.OutputState }

func (SysinfoMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sysinfo)(nil)).Elem()
}

func (o SysinfoMapOutput) ToSysinfoMapOutput() SysinfoMapOutput {
	return o
}

func (o SysinfoMapOutput) ToSysinfoMapOutputWithContext(ctx context.Context) SysinfoMapOutput {
	return o
}

func (o SysinfoMapOutput) MapIndex(k pulumi.StringInput) SysinfoOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Sysinfo {
		return vs[0].(map[string]*Sysinfo)[vs[1].(string)]
	}).(SysinfoOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SysinfoInput)(nil)).Elem(), &Sysinfo{})
	pulumi.RegisterInputType(reflect.TypeOf((*SysinfoArrayInput)(nil)).Elem(), SysinfoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SysinfoMapInput)(nil)).Elem(), SysinfoMap{})
	pulumi.RegisterOutputType(SysinfoOutput{})
	pulumi.RegisterOutputType(SysinfoArrayOutput{})
	pulumi.RegisterOutputType(SysinfoMapOutput{})
}
