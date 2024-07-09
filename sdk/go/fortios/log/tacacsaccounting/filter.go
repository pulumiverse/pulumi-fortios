// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tacacsaccounting

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Settings for TACACS+ accounting events filter. Applies to FortiOS Version `>= 7.0.2`.
//
// ## Import
//
// LogTacacsAccounting Filter can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:log/tacacsaccounting/filter:Filter labelname LogTacacsAccountingFilter
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:log/tacacsaccounting/filter:Filter labelname LogTacacsAccountingFilter
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Filter struct {
	pulumi.CustomResourceState

	// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
	CliCmdAudit pulumi.StringOutput `pulumi:"cliCmdAudit"`
	// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
	ConfigChangeAudit pulumi.StringOutput `pulumi:"configChangeAudit"`
	// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
	LoginAudit pulumi.StringOutput `pulumi:"loginAudit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewFilter registers a new resource with the given unique name, arguments, and options.
func NewFilter(ctx *pulumi.Context,
	name string, args *FilterArgs, opts ...pulumi.ResourceOption) (*Filter, error) {
	if args == nil {
		args = &FilterArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Filter
	err := ctx.RegisterResource("fortios:log/tacacsaccounting/filter:Filter", name, args, &resource, opts...)
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
	err := ctx.ReadResource("fortios:log/tacacsaccounting/filter:Filter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Filter resources.
type filterState struct {
	// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
	CliCmdAudit *string `pulumi:"cliCmdAudit"`
	// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
	ConfigChangeAudit *string `pulumi:"configChangeAudit"`
	// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
	LoginAudit *string `pulumi:"loginAudit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FilterState struct {
	// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
	CliCmdAudit pulumi.StringPtrInput
	// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
	ConfigChangeAudit pulumi.StringPtrInput
	// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
	LoginAudit pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*filterState)(nil)).Elem()
}

type filterArgs struct {
	// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
	CliCmdAudit *string `pulumi:"cliCmdAudit"`
	// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
	ConfigChangeAudit *string `pulumi:"configChangeAudit"`
	// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
	LoginAudit *string `pulumi:"loginAudit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Filter resource.
type FilterArgs struct {
	// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
	CliCmdAudit pulumi.StringPtrInput
	// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
	ConfigChangeAudit pulumi.StringPtrInput
	// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
	LoginAudit pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
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

// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
func (o FilterOutput) CliCmdAudit() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.CliCmdAudit }).(pulumi.StringOutput)
}

// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
func (o FilterOutput) ConfigChangeAudit() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.ConfigChangeAudit }).(pulumi.StringOutput)
}

// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
func (o FilterOutput) LoginAudit() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.LoginAudit }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FilterOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
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
