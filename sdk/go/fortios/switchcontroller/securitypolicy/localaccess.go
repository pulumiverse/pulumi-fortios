// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securitypolicy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure allowaccess list for mgmt and internal interfaces on managed FortiSwitch. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// SwitchControllerSecurityPolicy LocalAccess can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/securitypolicy/localaccess:Localaccess labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/securitypolicy/localaccess:Localaccess labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Localaccess struct {
	pulumi.CustomResourceState

	// Allowed access on the switch internal interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
	InternalAllowaccess pulumi.StringOutput `pulumi:"internalAllowaccess"`
	// Allowed access on the switch management interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
	MgmtAllowaccess pulumi.StringOutput `pulumi:"mgmtAllowaccess"`
	// Policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewLocalaccess registers a new resource with the given unique name, arguments, and options.
func NewLocalaccess(ctx *pulumi.Context,
	name string, args *LocalaccessArgs, opts ...pulumi.ResourceOption) (*Localaccess, error) {
	if args == nil {
		args = &LocalaccessArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Localaccess
	err := ctx.RegisterResource("fortios:switchcontroller/securitypolicy/localaccess:Localaccess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalaccess gets an existing Localaccess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalaccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalaccessState, opts ...pulumi.ResourceOption) (*Localaccess, error) {
	var resource Localaccess
	err := ctx.ReadResource("fortios:switchcontroller/securitypolicy/localaccess:Localaccess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Localaccess resources.
type localaccessState struct {
	// Allowed access on the switch internal interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
	InternalAllowaccess *string `pulumi:"internalAllowaccess"`
	// Allowed access on the switch management interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
	MgmtAllowaccess *string `pulumi:"mgmtAllowaccess"`
	// Policy name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type LocalaccessState struct {
	// Allowed access on the switch internal interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
	InternalAllowaccess pulumi.StringPtrInput
	// Allowed access on the switch management interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
	MgmtAllowaccess pulumi.StringPtrInput
	// Policy name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LocalaccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*localaccessState)(nil)).Elem()
}

type localaccessArgs struct {
	// Allowed access on the switch internal interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
	InternalAllowaccess *string `pulumi:"internalAllowaccess"`
	// Allowed access on the switch management interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
	MgmtAllowaccess *string `pulumi:"mgmtAllowaccess"`
	// Policy name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Localaccess resource.
type LocalaccessArgs struct {
	// Allowed access on the switch internal interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
	InternalAllowaccess pulumi.StringPtrInput
	// Allowed access on the switch management interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
	MgmtAllowaccess pulumi.StringPtrInput
	// Policy name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LocalaccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localaccessArgs)(nil)).Elem()
}

type LocalaccessInput interface {
	pulumi.Input

	ToLocalaccessOutput() LocalaccessOutput
	ToLocalaccessOutputWithContext(ctx context.Context) LocalaccessOutput
}

func (*Localaccess) ElementType() reflect.Type {
	return reflect.TypeOf((**Localaccess)(nil)).Elem()
}

func (i *Localaccess) ToLocalaccessOutput() LocalaccessOutput {
	return i.ToLocalaccessOutputWithContext(context.Background())
}

func (i *Localaccess) ToLocalaccessOutputWithContext(ctx context.Context) LocalaccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalaccessOutput)
}

// LocalaccessArrayInput is an input type that accepts LocalaccessArray and LocalaccessArrayOutput values.
// You can construct a concrete instance of `LocalaccessArrayInput` via:
//
//	LocalaccessArray{ LocalaccessArgs{...} }
type LocalaccessArrayInput interface {
	pulumi.Input

	ToLocalaccessArrayOutput() LocalaccessArrayOutput
	ToLocalaccessArrayOutputWithContext(context.Context) LocalaccessArrayOutput
}

type LocalaccessArray []LocalaccessInput

func (LocalaccessArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Localaccess)(nil)).Elem()
}

func (i LocalaccessArray) ToLocalaccessArrayOutput() LocalaccessArrayOutput {
	return i.ToLocalaccessArrayOutputWithContext(context.Background())
}

func (i LocalaccessArray) ToLocalaccessArrayOutputWithContext(ctx context.Context) LocalaccessArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalaccessArrayOutput)
}

// LocalaccessMapInput is an input type that accepts LocalaccessMap and LocalaccessMapOutput values.
// You can construct a concrete instance of `LocalaccessMapInput` via:
//
//	LocalaccessMap{ "key": LocalaccessArgs{...} }
type LocalaccessMapInput interface {
	pulumi.Input

	ToLocalaccessMapOutput() LocalaccessMapOutput
	ToLocalaccessMapOutputWithContext(context.Context) LocalaccessMapOutput
}

type LocalaccessMap map[string]LocalaccessInput

func (LocalaccessMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Localaccess)(nil)).Elem()
}

func (i LocalaccessMap) ToLocalaccessMapOutput() LocalaccessMapOutput {
	return i.ToLocalaccessMapOutputWithContext(context.Background())
}

func (i LocalaccessMap) ToLocalaccessMapOutputWithContext(ctx context.Context) LocalaccessMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalaccessMapOutput)
}

type LocalaccessOutput struct{ *pulumi.OutputState }

func (LocalaccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Localaccess)(nil)).Elem()
}

func (o LocalaccessOutput) ToLocalaccessOutput() LocalaccessOutput {
	return o
}

func (o LocalaccessOutput) ToLocalaccessOutputWithContext(ctx context.Context) LocalaccessOutput {
	return o
}

// Allowed access on the switch internal interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
func (o LocalaccessOutput) InternalAllowaccess() pulumi.StringOutput {
	return o.ApplyT(func(v *Localaccess) pulumi.StringOutput { return v.InternalAllowaccess }).(pulumi.StringOutput)
}

// Allowed access on the switch management interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
func (o LocalaccessOutput) MgmtAllowaccess() pulumi.StringOutput {
	return o.ApplyT(func(v *Localaccess) pulumi.StringOutput { return v.MgmtAllowaccess }).(pulumi.StringOutput)
}

// Policy name.
func (o LocalaccessOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Localaccess) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o LocalaccessOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Localaccess) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type LocalaccessArrayOutput struct{ *pulumi.OutputState }

func (LocalaccessArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Localaccess)(nil)).Elem()
}

func (o LocalaccessArrayOutput) ToLocalaccessArrayOutput() LocalaccessArrayOutput {
	return o
}

func (o LocalaccessArrayOutput) ToLocalaccessArrayOutputWithContext(ctx context.Context) LocalaccessArrayOutput {
	return o
}

func (o LocalaccessArrayOutput) Index(i pulumi.IntInput) LocalaccessOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Localaccess {
		return vs[0].([]*Localaccess)[vs[1].(int)]
	}).(LocalaccessOutput)
}

type LocalaccessMapOutput struct{ *pulumi.OutputState }

func (LocalaccessMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Localaccess)(nil)).Elem()
}

func (o LocalaccessMapOutput) ToLocalaccessMapOutput() LocalaccessMapOutput {
	return o
}

func (o LocalaccessMapOutput) ToLocalaccessMapOutputWithContext(ctx context.Context) LocalaccessMapOutput {
	return o
}

func (o LocalaccessMapOutput) MapIndex(k pulumi.StringInput) LocalaccessOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Localaccess {
		return vs[0].(map[string]*Localaccess)[vs[1].(string)]
	}).(LocalaccessOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalaccessInput)(nil)).Elem(), &Localaccess{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalaccessArrayInput)(nil)).Elem(), LocalaccessArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalaccessMapInput)(nil)).Elem(), LocalaccessMap{})
	pulumi.RegisterOutputType(LocalaccessOutput{})
	pulumi.RegisterOutputType(LocalaccessArrayOutput{})
	pulumi.RegisterOutputType(LocalaccessMapOutput{})
}
