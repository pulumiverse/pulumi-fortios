// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dpdk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure CPUs enabled to run engines in each DPDK stage. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// Dpdk Cpus can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:dpdk/cpus:Cpus labelname DpdkCpus
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:dpdk/cpus:Cpus labelname DpdkCpus
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Cpus struct {
	pulumi.CustomResourceState

	// CPUs enabled to run DPDK IPS engines.
	IpsCpus pulumi.StringOutput `pulumi:"ipsCpus"`
	// CPUs isolated to run only the DPDK engines with the exception of processes that have affinity explicitly set by either a user configuration or by their implementation.
	IsolatedCpus pulumi.StringOutput `pulumi:"isolatedCpus"`
	// CPUs enabled to run DPDK RX engines.
	RxCpus pulumi.StringOutput `pulumi:"rxCpus"`
	// CPUs enabled to run DPDK TX engines.
	TxCpus pulumi.StringOutput `pulumi:"txCpus"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// CPUs enabled to run DPDK VNP engines.
	VnpCpus pulumi.StringOutput `pulumi:"vnpCpus"`
}

// NewCpus registers a new resource with the given unique name, arguments, and options.
func NewCpus(ctx *pulumi.Context,
	name string, args *CpusArgs, opts ...pulumi.ResourceOption) (*Cpus, error) {
	if args == nil {
		args = &CpusArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Cpus
	err := ctx.RegisterResource("fortios:dpdk/cpus:Cpus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCpus gets an existing Cpus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCpus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CpusState, opts ...pulumi.ResourceOption) (*Cpus, error) {
	var resource Cpus
	err := ctx.ReadResource("fortios:dpdk/cpus:Cpus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cpus resources.
type cpusState struct {
	// CPUs enabled to run DPDK IPS engines.
	IpsCpus *string `pulumi:"ipsCpus"`
	// CPUs isolated to run only the DPDK engines with the exception of processes that have affinity explicitly set by either a user configuration or by their implementation.
	IsolatedCpus *string `pulumi:"isolatedCpus"`
	// CPUs enabled to run DPDK RX engines.
	RxCpus *string `pulumi:"rxCpus"`
	// CPUs enabled to run DPDK TX engines.
	TxCpus *string `pulumi:"txCpus"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// CPUs enabled to run DPDK VNP engines.
	VnpCpus *string `pulumi:"vnpCpus"`
}

type CpusState struct {
	// CPUs enabled to run DPDK IPS engines.
	IpsCpus pulumi.StringPtrInput
	// CPUs isolated to run only the DPDK engines with the exception of processes that have affinity explicitly set by either a user configuration or by their implementation.
	IsolatedCpus pulumi.StringPtrInput
	// CPUs enabled to run DPDK RX engines.
	RxCpus pulumi.StringPtrInput
	// CPUs enabled to run DPDK TX engines.
	TxCpus pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// CPUs enabled to run DPDK VNP engines.
	VnpCpus pulumi.StringPtrInput
}

func (CpusState) ElementType() reflect.Type {
	return reflect.TypeOf((*cpusState)(nil)).Elem()
}

type cpusArgs struct {
	// CPUs enabled to run DPDK IPS engines.
	IpsCpus *string `pulumi:"ipsCpus"`
	// CPUs isolated to run only the DPDK engines with the exception of processes that have affinity explicitly set by either a user configuration or by their implementation.
	IsolatedCpus *string `pulumi:"isolatedCpus"`
	// CPUs enabled to run DPDK RX engines.
	RxCpus *string `pulumi:"rxCpus"`
	// CPUs enabled to run DPDK TX engines.
	TxCpus *string `pulumi:"txCpus"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// CPUs enabled to run DPDK VNP engines.
	VnpCpus *string `pulumi:"vnpCpus"`
}

// The set of arguments for constructing a Cpus resource.
type CpusArgs struct {
	// CPUs enabled to run DPDK IPS engines.
	IpsCpus pulumi.StringPtrInput
	// CPUs isolated to run only the DPDK engines with the exception of processes that have affinity explicitly set by either a user configuration or by their implementation.
	IsolatedCpus pulumi.StringPtrInput
	// CPUs enabled to run DPDK RX engines.
	RxCpus pulumi.StringPtrInput
	// CPUs enabled to run DPDK TX engines.
	TxCpus pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// CPUs enabled to run DPDK VNP engines.
	VnpCpus pulumi.StringPtrInput
}

func (CpusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cpusArgs)(nil)).Elem()
}

type CpusInput interface {
	pulumi.Input

	ToCpusOutput() CpusOutput
	ToCpusOutputWithContext(ctx context.Context) CpusOutput
}

func (*Cpus) ElementType() reflect.Type {
	return reflect.TypeOf((**Cpus)(nil)).Elem()
}

func (i *Cpus) ToCpusOutput() CpusOutput {
	return i.ToCpusOutputWithContext(context.Background())
}

func (i *Cpus) ToCpusOutputWithContext(ctx context.Context) CpusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CpusOutput)
}

// CpusArrayInput is an input type that accepts CpusArray and CpusArrayOutput values.
// You can construct a concrete instance of `CpusArrayInput` via:
//
//	CpusArray{ CpusArgs{...} }
type CpusArrayInput interface {
	pulumi.Input

	ToCpusArrayOutput() CpusArrayOutput
	ToCpusArrayOutputWithContext(context.Context) CpusArrayOutput
}

type CpusArray []CpusInput

func (CpusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cpus)(nil)).Elem()
}

func (i CpusArray) ToCpusArrayOutput() CpusArrayOutput {
	return i.ToCpusArrayOutputWithContext(context.Background())
}

func (i CpusArray) ToCpusArrayOutputWithContext(ctx context.Context) CpusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CpusArrayOutput)
}

// CpusMapInput is an input type that accepts CpusMap and CpusMapOutput values.
// You can construct a concrete instance of `CpusMapInput` via:
//
//	CpusMap{ "key": CpusArgs{...} }
type CpusMapInput interface {
	pulumi.Input

	ToCpusMapOutput() CpusMapOutput
	ToCpusMapOutputWithContext(context.Context) CpusMapOutput
}

type CpusMap map[string]CpusInput

func (CpusMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cpus)(nil)).Elem()
}

func (i CpusMap) ToCpusMapOutput() CpusMapOutput {
	return i.ToCpusMapOutputWithContext(context.Background())
}

func (i CpusMap) ToCpusMapOutputWithContext(ctx context.Context) CpusMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CpusMapOutput)
}

type CpusOutput struct{ *pulumi.OutputState }

func (CpusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cpus)(nil)).Elem()
}

func (o CpusOutput) ToCpusOutput() CpusOutput {
	return o
}

func (o CpusOutput) ToCpusOutputWithContext(ctx context.Context) CpusOutput {
	return o
}

// CPUs enabled to run DPDK IPS engines.
func (o CpusOutput) IpsCpus() pulumi.StringOutput {
	return o.ApplyT(func(v *Cpus) pulumi.StringOutput { return v.IpsCpus }).(pulumi.StringOutput)
}

// CPUs isolated to run only the DPDK engines with the exception of processes that have affinity explicitly set by either a user configuration or by their implementation.
func (o CpusOutput) IsolatedCpus() pulumi.StringOutput {
	return o.ApplyT(func(v *Cpus) pulumi.StringOutput { return v.IsolatedCpus }).(pulumi.StringOutput)
}

// CPUs enabled to run DPDK RX engines.
func (o CpusOutput) RxCpus() pulumi.StringOutput {
	return o.ApplyT(func(v *Cpus) pulumi.StringOutput { return v.RxCpus }).(pulumi.StringOutput)
}

// CPUs enabled to run DPDK TX engines.
func (o CpusOutput) TxCpus() pulumi.StringOutput {
	return o.ApplyT(func(v *Cpus) pulumi.StringOutput { return v.TxCpus }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o CpusOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cpus) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// CPUs enabled to run DPDK VNP engines.
func (o CpusOutput) VnpCpus() pulumi.StringOutput {
	return o.ApplyT(func(v *Cpus) pulumi.StringOutput { return v.VnpCpus }).(pulumi.StringOutput)
}

type CpusArrayOutput struct{ *pulumi.OutputState }

func (CpusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cpus)(nil)).Elem()
}

func (o CpusArrayOutput) ToCpusArrayOutput() CpusArrayOutput {
	return o
}

func (o CpusArrayOutput) ToCpusArrayOutputWithContext(ctx context.Context) CpusArrayOutput {
	return o
}

func (o CpusArrayOutput) Index(i pulumi.IntInput) CpusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Cpus {
		return vs[0].([]*Cpus)[vs[1].(int)]
	}).(CpusOutput)
}

type CpusMapOutput struct{ *pulumi.OutputState }

func (CpusMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cpus)(nil)).Elem()
}

func (o CpusMapOutput) ToCpusMapOutput() CpusMapOutput {
	return o
}

func (o CpusMapOutput) ToCpusMapOutputWithContext(ctx context.Context) CpusMapOutput {
	return o
}

func (o CpusMapOutput) MapIndex(k pulumi.StringInput) CpusOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Cpus {
		return vs[0].(map[string]*Cpus)[vs[1].(string)]
	}).(CpusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CpusInput)(nil)).Elem(), &Cpus{})
	pulumi.RegisterInputType(reflect.TypeOf((*CpusArrayInput)(nil)).Elem(), CpusArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CpusMapInput)(nil)).Elem(), CpusMap{})
	pulumi.RegisterOutputType(CpusOutput{})
	pulumi.RegisterOutputType(CpusArrayOutput{})
	pulumi.RegisterOutputType(CpusMapOutput{})
}
