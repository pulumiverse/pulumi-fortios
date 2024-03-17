// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure decrypted traffic mirror. Applies to FortiOS Version `>= 6.4.0`.
//
// ## Import
//
// Firewall DecryptedTrafficMirror can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/decryptedtrafficmirror:Decryptedtrafficmirror labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/decryptedtrafficmirror:Decryptedtrafficmirror labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Decryptedtrafficmirror struct {
	pulumi.CustomResourceState

	// Set destination MAC address for mirrored traffic.
	Dstmac pulumi.StringOutput `pulumi:"dstmac"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Decrypted traffic mirror interface The structure of `interface` block is documented below.
	Interfaces DecryptedtrafficmirrorInterfaceArrayOutput `pulumi:"interfaces"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Source of decrypted traffic to be mirrored. Valid values: `client`, `server`, `both`.
	TrafficSource pulumi.StringOutput `pulumi:"trafficSource"`
	// Types of decrypted traffic to be mirrored. Valid values: `ssl`, `ssh`.
	TrafficType pulumi.StringOutput `pulumi:"trafficType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDecryptedtrafficmirror registers a new resource with the given unique name, arguments, and options.
func NewDecryptedtrafficmirror(ctx *pulumi.Context,
	name string, args *DecryptedtrafficmirrorArgs, opts ...pulumi.ResourceOption) (*Decryptedtrafficmirror, error) {
	if args == nil {
		args = &DecryptedtrafficmirrorArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Decryptedtrafficmirror
	err := ctx.RegisterResource("fortios:firewall/decryptedtrafficmirror:Decryptedtrafficmirror", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDecryptedtrafficmirror gets an existing Decryptedtrafficmirror resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDecryptedtrafficmirror(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DecryptedtrafficmirrorState, opts ...pulumi.ResourceOption) (*Decryptedtrafficmirror, error) {
	var resource Decryptedtrafficmirror
	err := ctx.ReadResource("fortios:firewall/decryptedtrafficmirror:Decryptedtrafficmirror", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Decryptedtrafficmirror resources.
type decryptedtrafficmirrorState struct {
	// Set destination MAC address for mirrored traffic.
	Dstmac *string `pulumi:"dstmac"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Decrypted traffic mirror interface The structure of `interface` block is documented below.
	Interfaces []DecryptedtrafficmirrorInterface `pulumi:"interfaces"`
	// Name.
	Name *string `pulumi:"name"`
	// Source of decrypted traffic to be mirrored. Valid values: `client`, `server`, `both`.
	TrafficSource *string `pulumi:"trafficSource"`
	// Types of decrypted traffic to be mirrored. Valid values: `ssl`, `ssh`.
	TrafficType *string `pulumi:"trafficType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DecryptedtrafficmirrorState struct {
	// Set destination MAC address for mirrored traffic.
	Dstmac pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Decrypted traffic mirror interface The structure of `interface` block is documented below.
	Interfaces DecryptedtrafficmirrorInterfaceArrayInput
	// Name.
	Name pulumi.StringPtrInput
	// Source of decrypted traffic to be mirrored. Valid values: `client`, `server`, `both`.
	TrafficSource pulumi.StringPtrInput
	// Types of decrypted traffic to be mirrored. Valid values: `ssl`, `ssh`.
	TrafficType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DecryptedtrafficmirrorState) ElementType() reflect.Type {
	return reflect.TypeOf((*decryptedtrafficmirrorState)(nil)).Elem()
}

type decryptedtrafficmirrorArgs struct {
	// Set destination MAC address for mirrored traffic.
	Dstmac *string `pulumi:"dstmac"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Decrypted traffic mirror interface The structure of `interface` block is documented below.
	Interfaces []DecryptedtrafficmirrorInterface `pulumi:"interfaces"`
	// Name.
	Name *string `pulumi:"name"`
	// Source of decrypted traffic to be mirrored. Valid values: `client`, `server`, `both`.
	TrafficSource *string `pulumi:"trafficSource"`
	// Types of decrypted traffic to be mirrored. Valid values: `ssl`, `ssh`.
	TrafficType *string `pulumi:"trafficType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Decryptedtrafficmirror resource.
type DecryptedtrafficmirrorArgs struct {
	// Set destination MAC address for mirrored traffic.
	Dstmac pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Decrypted traffic mirror interface The structure of `interface` block is documented below.
	Interfaces DecryptedtrafficmirrorInterfaceArrayInput
	// Name.
	Name pulumi.StringPtrInput
	// Source of decrypted traffic to be mirrored. Valid values: `client`, `server`, `both`.
	TrafficSource pulumi.StringPtrInput
	// Types of decrypted traffic to be mirrored. Valid values: `ssl`, `ssh`.
	TrafficType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DecryptedtrafficmirrorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*decryptedtrafficmirrorArgs)(nil)).Elem()
}

type DecryptedtrafficmirrorInput interface {
	pulumi.Input

	ToDecryptedtrafficmirrorOutput() DecryptedtrafficmirrorOutput
	ToDecryptedtrafficmirrorOutputWithContext(ctx context.Context) DecryptedtrafficmirrorOutput
}

func (*Decryptedtrafficmirror) ElementType() reflect.Type {
	return reflect.TypeOf((**Decryptedtrafficmirror)(nil)).Elem()
}

func (i *Decryptedtrafficmirror) ToDecryptedtrafficmirrorOutput() DecryptedtrafficmirrorOutput {
	return i.ToDecryptedtrafficmirrorOutputWithContext(context.Background())
}

func (i *Decryptedtrafficmirror) ToDecryptedtrafficmirrorOutputWithContext(ctx context.Context) DecryptedtrafficmirrorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DecryptedtrafficmirrorOutput)
}

// DecryptedtrafficmirrorArrayInput is an input type that accepts DecryptedtrafficmirrorArray and DecryptedtrafficmirrorArrayOutput values.
// You can construct a concrete instance of `DecryptedtrafficmirrorArrayInput` via:
//
//	DecryptedtrafficmirrorArray{ DecryptedtrafficmirrorArgs{...} }
type DecryptedtrafficmirrorArrayInput interface {
	pulumi.Input

	ToDecryptedtrafficmirrorArrayOutput() DecryptedtrafficmirrorArrayOutput
	ToDecryptedtrafficmirrorArrayOutputWithContext(context.Context) DecryptedtrafficmirrorArrayOutput
}

type DecryptedtrafficmirrorArray []DecryptedtrafficmirrorInput

func (DecryptedtrafficmirrorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Decryptedtrafficmirror)(nil)).Elem()
}

func (i DecryptedtrafficmirrorArray) ToDecryptedtrafficmirrorArrayOutput() DecryptedtrafficmirrorArrayOutput {
	return i.ToDecryptedtrafficmirrorArrayOutputWithContext(context.Background())
}

func (i DecryptedtrafficmirrorArray) ToDecryptedtrafficmirrorArrayOutputWithContext(ctx context.Context) DecryptedtrafficmirrorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DecryptedtrafficmirrorArrayOutput)
}

// DecryptedtrafficmirrorMapInput is an input type that accepts DecryptedtrafficmirrorMap and DecryptedtrafficmirrorMapOutput values.
// You can construct a concrete instance of `DecryptedtrafficmirrorMapInput` via:
//
//	DecryptedtrafficmirrorMap{ "key": DecryptedtrafficmirrorArgs{...} }
type DecryptedtrafficmirrorMapInput interface {
	pulumi.Input

	ToDecryptedtrafficmirrorMapOutput() DecryptedtrafficmirrorMapOutput
	ToDecryptedtrafficmirrorMapOutputWithContext(context.Context) DecryptedtrafficmirrorMapOutput
}

type DecryptedtrafficmirrorMap map[string]DecryptedtrafficmirrorInput

func (DecryptedtrafficmirrorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Decryptedtrafficmirror)(nil)).Elem()
}

func (i DecryptedtrafficmirrorMap) ToDecryptedtrafficmirrorMapOutput() DecryptedtrafficmirrorMapOutput {
	return i.ToDecryptedtrafficmirrorMapOutputWithContext(context.Background())
}

func (i DecryptedtrafficmirrorMap) ToDecryptedtrafficmirrorMapOutputWithContext(ctx context.Context) DecryptedtrafficmirrorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DecryptedtrafficmirrorMapOutput)
}

type DecryptedtrafficmirrorOutput struct{ *pulumi.OutputState }

func (DecryptedtrafficmirrorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Decryptedtrafficmirror)(nil)).Elem()
}

func (o DecryptedtrafficmirrorOutput) ToDecryptedtrafficmirrorOutput() DecryptedtrafficmirrorOutput {
	return o
}

func (o DecryptedtrafficmirrorOutput) ToDecryptedtrafficmirrorOutputWithContext(ctx context.Context) DecryptedtrafficmirrorOutput {
	return o
}

// Set destination MAC address for mirrored traffic.
func (o DecryptedtrafficmirrorOutput) Dstmac() pulumi.StringOutput {
	return o.ApplyT(func(v *Decryptedtrafficmirror) pulumi.StringOutput { return v.Dstmac }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o DecryptedtrafficmirrorOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Decryptedtrafficmirror) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Decrypted traffic mirror interface The structure of `interface` block is documented below.
func (o DecryptedtrafficmirrorOutput) Interfaces() DecryptedtrafficmirrorInterfaceArrayOutput {
	return o.ApplyT(func(v *Decryptedtrafficmirror) DecryptedtrafficmirrorInterfaceArrayOutput { return v.Interfaces }).(DecryptedtrafficmirrorInterfaceArrayOutput)
}

// Name.
func (o DecryptedtrafficmirrorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Decryptedtrafficmirror) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Source of decrypted traffic to be mirrored. Valid values: `client`, `server`, `both`.
func (o DecryptedtrafficmirrorOutput) TrafficSource() pulumi.StringOutput {
	return o.ApplyT(func(v *Decryptedtrafficmirror) pulumi.StringOutput { return v.TrafficSource }).(pulumi.StringOutput)
}

// Types of decrypted traffic to be mirrored. Valid values: `ssl`, `ssh`.
func (o DecryptedtrafficmirrorOutput) TrafficType() pulumi.StringOutput {
	return o.ApplyT(func(v *Decryptedtrafficmirror) pulumi.StringOutput { return v.TrafficType }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DecryptedtrafficmirrorOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Decryptedtrafficmirror) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type DecryptedtrafficmirrorArrayOutput struct{ *pulumi.OutputState }

func (DecryptedtrafficmirrorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Decryptedtrafficmirror)(nil)).Elem()
}

func (o DecryptedtrafficmirrorArrayOutput) ToDecryptedtrafficmirrorArrayOutput() DecryptedtrafficmirrorArrayOutput {
	return o
}

func (o DecryptedtrafficmirrorArrayOutput) ToDecryptedtrafficmirrorArrayOutputWithContext(ctx context.Context) DecryptedtrafficmirrorArrayOutput {
	return o
}

func (o DecryptedtrafficmirrorArrayOutput) Index(i pulumi.IntInput) DecryptedtrafficmirrorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Decryptedtrafficmirror {
		return vs[0].([]*Decryptedtrafficmirror)[vs[1].(int)]
	}).(DecryptedtrafficmirrorOutput)
}

type DecryptedtrafficmirrorMapOutput struct{ *pulumi.OutputState }

func (DecryptedtrafficmirrorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Decryptedtrafficmirror)(nil)).Elem()
}

func (o DecryptedtrafficmirrorMapOutput) ToDecryptedtrafficmirrorMapOutput() DecryptedtrafficmirrorMapOutput {
	return o
}

func (o DecryptedtrafficmirrorMapOutput) ToDecryptedtrafficmirrorMapOutputWithContext(ctx context.Context) DecryptedtrafficmirrorMapOutput {
	return o
}

func (o DecryptedtrafficmirrorMapOutput) MapIndex(k pulumi.StringInput) DecryptedtrafficmirrorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Decryptedtrafficmirror {
		return vs[0].(map[string]*Decryptedtrafficmirror)[vs[1].(string)]
	}).(DecryptedtrafficmirrorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DecryptedtrafficmirrorInput)(nil)).Elem(), &Decryptedtrafficmirror{})
	pulumi.RegisterInputType(reflect.TypeOf((*DecryptedtrafficmirrorArrayInput)(nil)).Elem(), DecryptedtrafficmirrorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DecryptedtrafficmirrorMapInput)(nil)).Elem(), DecryptedtrafficmirrorMap{})
	pulumi.RegisterOutputType(DecryptedtrafficmirrorOutput{})
	pulumi.RegisterOutputType(DecryptedtrafficmirrorArrayOutput{})
	pulumi.RegisterOutputType(DecryptedtrafficmirrorMapOutput{})
}
