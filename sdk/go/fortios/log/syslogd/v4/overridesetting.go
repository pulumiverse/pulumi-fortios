// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v4

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Override settings for remote syslog server.
//
// ## Import
//
// LogSyslogd4 OverrideSetting can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:log/syslogd/v4/overridesetting:Overridesetting labelname LogSyslogd4OverrideSetting
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:log/syslogd/v4/overridesetting:Overridesetting labelname LogSyslogd4OverrideSetting
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Overridesetting struct {
	pulumi.CustomResourceState

	// Certificate used to communicate with Syslog server.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames OverridesettingCustomFieldNameArrayOutput `pulumi:"customFieldNames"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium`, `high`, `low`, `disable`.
	EncAlgorithm pulumi.StringOutput `pulumi:"encAlgorithm"`
	// Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
	Facility pulumi.StringOutput `pulumi:"facility"`
	// Log format.
	Format pulumi.StringOutput `pulumi:"format"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Specify outgoing interface to reach server.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringOutput `pulumi:"interfaceSelectMethod"`
	// Syslog maximum log rate in MBps (0 = unlimited).
	MaxLogRate pulumi.IntOutput `pulumi:"maxLogRate"`
	// Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// Enable/disable override syslog settings. Valid values: `enable`, `disable`.
	Override pulumi.StringOutput `pulumi:"override"`
	// Server listen port.
	Port pulumi.IntOutput `pulumi:"port"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringOutput `pulumi:"priority"`
	// Address of remote syslog server.
	Server pulumi.StringOutput `pulumi:"server"`
	// Source IP address of syslog.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion pulumi.StringOutput `pulumi:"sslMinProtoVersion"`
	// Enable/disable remote syslog logging. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Hidden setting index of Syslog.
	SyslogType pulumi.IntOutput `pulumi:"syslogType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewOverridesetting registers a new resource with the given unique name, arguments, and options.
func NewOverridesetting(ctx *pulumi.Context,
	name string, args *OverridesettingArgs, opts ...pulumi.ResourceOption) (*Overridesetting, error) {
	if args == nil {
		args = &OverridesettingArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Overridesetting
	err := ctx.RegisterResource("fortios:log/syslogd/v4/overridesetting:Overridesetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOverridesetting gets an existing Overridesetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOverridesetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OverridesettingState, opts ...pulumi.ResourceOption) (*Overridesetting, error) {
	var resource Overridesetting
	err := ctx.ReadResource("fortios:log/syslogd/v4/overridesetting:Overridesetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Overridesetting resources.
type overridesettingState struct {
	// Certificate used to communicate with Syslog server.
	Certificate *string `pulumi:"certificate"`
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames []OverridesettingCustomFieldName `pulumi:"customFieldNames"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium`, `high`, `low`, `disable`.
	EncAlgorithm *string `pulumi:"encAlgorithm"`
	// Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
	Facility *string `pulumi:"facility"`
	// Log format.
	Format *string `pulumi:"format"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Syslog maximum log rate in MBps (0 = unlimited).
	MaxLogRate *int `pulumi:"maxLogRate"`
	// Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.
	Mode *string `pulumi:"mode"`
	// Enable/disable override syslog settings. Valid values: `enable`, `disable`.
	Override *string `pulumi:"override"`
	// Server listen port.
	Port *int `pulumi:"port"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority *string `pulumi:"priority"`
	// Address of remote syslog server.
	Server *string `pulumi:"server"`
	// Source IP address of syslog.
	SourceIp *string `pulumi:"sourceIp"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Enable/disable remote syslog logging. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Hidden setting index of Syslog.
	SyslogType *int `pulumi:"syslogType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type OverridesettingState struct {
	// Certificate used to communicate with Syslog server.
	Certificate pulumi.StringPtrInput
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames OverridesettingCustomFieldNameArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium`, `high`, `low`, `disable`.
	EncAlgorithm pulumi.StringPtrInput
	// Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
	Facility pulumi.StringPtrInput
	// Log format.
	Format pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Syslog maximum log rate in MBps (0 = unlimited).
	MaxLogRate pulumi.IntPtrInput
	// Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.
	Mode pulumi.StringPtrInput
	// Enable/disable override syslog settings. Valid values: `enable`, `disable`.
	Override pulumi.StringPtrInput
	// Server listen port.
	Port pulumi.IntPtrInput
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringPtrInput
	// Address of remote syslog server.
	Server pulumi.StringPtrInput
	// Source IP address of syslog.
	SourceIp pulumi.StringPtrInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion pulumi.StringPtrInput
	// Enable/disable remote syslog logging. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Hidden setting index of Syslog.
	SyslogType pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (OverridesettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*overridesettingState)(nil)).Elem()
}

type overridesettingArgs struct {
	// Certificate used to communicate with Syslog server.
	Certificate *string `pulumi:"certificate"`
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames []OverridesettingCustomFieldName `pulumi:"customFieldNames"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium`, `high`, `low`, `disable`.
	EncAlgorithm *string `pulumi:"encAlgorithm"`
	// Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
	Facility *string `pulumi:"facility"`
	// Log format.
	Format *string `pulumi:"format"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Syslog maximum log rate in MBps (0 = unlimited).
	MaxLogRate *int `pulumi:"maxLogRate"`
	// Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.
	Mode *string `pulumi:"mode"`
	// Enable/disable override syslog settings. Valid values: `enable`, `disable`.
	Override *string `pulumi:"override"`
	// Server listen port.
	Port *int `pulumi:"port"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority *string `pulumi:"priority"`
	// Address of remote syslog server.
	Server *string `pulumi:"server"`
	// Source IP address of syslog.
	SourceIp *string `pulumi:"sourceIp"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Enable/disable remote syslog logging. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Hidden setting index of Syslog.
	SyslogType *int `pulumi:"syslogType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Overridesetting resource.
type OverridesettingArgs struct {
	// Certificate used to communicate with Syslog server.
	Certificate pulumi.StringPtrInput
	// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
	CustomFieldNames OverridesettingCustomFieldNameArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium`, `high`, `low`, `disable`.
	EncAlgorithm pulumi.StringPtrInput
	// Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
	Facility pulumi.StringPtrInput
	// Log format.
	Format pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Syslog maximum log rate in MBps (0 = unlimited).
	MaxLogRate pulumi.IntPtrInput
	// Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.
	Mode pulumi.StringPtrInput
	// Enable/disable override syslog settings. Valid values: `enable`, `disable`.
	Override pulumi.StringPtrInput
	// Server listen port.
	Port pulumi.IntPtrInput
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringPtrInput
	// Address of remote syslog server.
	Server pulumi.StringPtrInput
	// Source IP address of syslog.
	SourceIp pulumi.StringPtrInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion pulumi.StringPtrInput
	// Enable/disable remote syslog logging. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Hidden setting index of Syslog.
	SyslogType pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (OverridesettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*overridesettingArgs)(nil)).Elem()
}

type OverridesettingInput interface {
	pulumi.Input

	ToOverridesettingOutput() OverridesettingOutput
	ToOverridesettingOutputWithContext(ctx context.Context) OverridesettingOutput
}

func (*Overridesetting) ElementType() reflect.Type {
	return reflect.TypeOf((**Overridesetting)(nil)).Elem()
}

func (i *Overridesetting) ToOverridesettingOutput() OverridesettingOutput {
	return i.ToOverridesettingOutputWithContext(context.Background())
}

func (i *Overridesetting) ToOverridesettingOutputWithContext(ctx context.Context) OverridesettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverridesettingOutput)
}

// OverridesettingArrayInput is an input type that accepts OverridesettingArray and OverridesettingArrayOutput values.
// You can construct a concrete instance of `OverridesettingArrayInput` via:
//
//	OverridesettingArray{ OverridesettingArgs{...} }
type OverridesettingArrayInput interface {
	pulumi.Input

	ToOverridesettingArrayOutput() OverridesettingArrayOutput
	ToOverridesettingArrayOutputWithContext(context.Context) OverridesettingArrayOutput
}

type OverridesettingArray []OverridesettingInput

func (OverridesettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Overridesetting)(nil)).Elem()
}

func (i OverridesettingArray) ToOverridesettingArrayOutput() OverridesettingArrayOutput {
	return i.ToOverridesettingArrayOutputWithContext(context.Background())
}

func (i OverridesettingArray) ToOverridesettingArrayOutputWithContext(ctx context.Context) OverridesettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverridesettingArrayOutput)
}

// OverridesettingMapInput is an input type that accepts OverridesettingMap and OverridesettingMapOutput values.
// You can construct a concrete instance of `OverridesettingMapInput` via:
//
//	OverridesettingMap{ "key": OverridesettingArgs{...} }
type OverridesettingMapInput interface {
	pulumi.Input

	ToOverridesettingMapOutput() OverridesettingMapOutput
	ToOverridesettingMapOutputWithContext(context.Context) OverridesettingMapOutput
}

type OverridesettingMap map[string]OverridesettingInput

func (OverridesettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Overridesetting)(nil)).Elem()
}

func (i OverridesettingMap) ToOverridesettingMapOutput() OverridesettingMapOutput {
	return i.ToOverridesettingMapOutputWithContext(context.Background())
}

func (i OverridesettingMap) ToOverridesettingMapOutputWithContext(ctx context.Context) OverridesettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverridesettingMapOutput)
}

type OverridesettingOutput struct{ *pulumi.OutputState }

func (OverridesettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Overridesetting)(nil)).Elem()
}

func (o OverridesettingOutput) ToOverridesettingOutput() OverridesettingOutput {
	return o
}

func (o OverridesettingOutput) ToOverridesettingOutputWithContext(ctx context.Context) OverridesettingOutput {
	return o
}

// Certificate used to communicate with Syslog server.
func (o OverridesettingOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridesetting) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
func (o OverridesettingOutput) CustomFieldNames() OverridesettingCustomFieldNameArrayOutput {
	return o.ApplyT(func(v *Overridesetting) OverridesettingCustomFieldNameArrayOutput { return v.CustomFieldNames }).(OverridesettingCustomFieldNameArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o OverridesettingOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Overridesetting) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium`, `high`, `low`, `disable`.
func (o OverridesettingOutput) EncAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridesetting) pulumi.StringOutput { return v.EncAlgorithm }).(pulumi.StringOutput)
}

// Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
func (o OverridesettingOutput) Facility() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridesetting) pulumi.StringOutput { return v.Facility }).(pulumi.StringOutput)
}

// Log format.
func (o OverridesettingOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridesetting) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o OverridesettingOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Overridesetting) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Specify outgoing interface to reach server.
func (o OverridesettingOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridesetting) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
func (o OverridesettingOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridesetting) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

// Syslog maximum log rate in MBps (0 = unlimited).
func (o OverridesettingOutput) MaxLogRate() pulumi.IntOutput {
	return o.ApplyT(func(v *Overridesetting) pulumi.IntOutput { return v.MaxLogRate }).(pulumi.IntOutput)
}

// Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.
func (o OverridesettingOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridesetting) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// Enable/disable override syslog settings. Valid values: `enable`, `disable`.
func (o OverridesettingOutput) Override() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridesetting) pulumi.StringOutput { return v.Override }).(pulumi.StringOutput)
}

// Server listen port.
func (o OverridesettingOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Overridesetting) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Set log transmission priority. Valid values: `default`, `low`.
func (o OverridesettingOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridesetting) pulumi.StringOutput { return v.Priority }).(pulumi.StringOutput)
}

// Address of remote syslog server.
func (o OverridesettingOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridesetting) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

// Source IP address of syslog.
func (o OverridesettingOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridesetting) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
func (o OverridesettingOutput) SslMinProtoVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridesetting) pulumi.StringOutput { return v.SslMinProtoVersion }).(pulumi.StringOutput)
}

// Enable/disable remote syslog logging. Valid values: `enable`, `disable`.
func (o OverridesettingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridesetting) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Hidden setting index of Syslog.
func (o OverridesettingOutput) SyslogType() pulumi.IntOutput {
	return o.ApplyT(func(v *Overridesetting) pulumi.IntOutput { return v.SyslogType }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o OverridesettingOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Overridesetting) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type OverridesettingArrayOutput struct{ *pulumi.OutputState }

func (OverridesettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Overridesetting)(nil)).Elem()
}

func (o OverridesettingArrayOutput) ToOverridesettingArrayOutput() OverridesettingArrayOutput {
	return o
}

func (o OverridesettingArrayOutput) ToOverridesettingArrayOutputWithContext(ctx context.Context) OverridesettingArrayOutput {
	return o
}

func (o OverridesettingArrayOutput) Index(i pulumi.IntInput) OverridesettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Overridesetting {
		return vs[0].([]*Overridesetting)[vs[1].(int)]
	}).(OverridesettingOutput)
}

type OverridesettingMapOutput struct{ *pulumi.OutputState }

func (OverridesettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Overridesetting)(nil)).Elem()
}

func (o OverridesettingMapOutput) ToOverridesettingMapOutput() OverridesettingMapOutput {
	return o
}

func (o OverridesettingMapOutput) ToOverridesettingMapOutputWithContext(ctx context.Context) OverridesettingMapOutput {
	return o
}

func (o OverridesettingMapOutput) MapIndex(k pulumi.StringInput) OverridesettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Overridesetting {
		return vs[0].(map[string]*Overridesetting)[vs[1].(string)]
	}).(OverridesettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OverridesettingInput)(nil)).Elem(), &Overridesetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*OverridesettingArrayInput)(nil)).Elem(), OverridesettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OverridesettingMapInput)(nil)).Elem(), OverridesettingMap{})
	pulumi.RegisterOutputType(OverridesettingOutput{})
	pulumi.RegisterOutputType(OverridesettingArrayOutput{})
	pulumi.RegisterOutputType(OverridesettingMapOutput{})
}
