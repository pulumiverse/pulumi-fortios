// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package user

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FSSO active directory servers for polling mode.
//
// ## Import
//
// # User FssoPolling can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:user/fssopolling:Fssopolling labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:user/fssopolling:Fssopolling labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Fssopolling struct {
	pulumi.CustomResourceState

	// LDAP Group Info. The structure of `adgrp` block is documented below.
	Adgrps FssopollingAdgrpArrayOutput `pulumi:"adgrps"`
	// Default domain managed by this Active Directory server.
	DefaultDomain pulumi.StringOutput `pulumi:"defaultDomain"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Active Directory server ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// LDAP server name used in LDAP connection strings.
	LdapServer pulumi.StringOutput `pulumi:"ldapServer"`
	// Number of hours of logon history to keep, 0 means keep all history.
	LogonHistory pulumi.IntOutput `pulumi:"logonHistory"`
	// Password required to log into this Active Directory server
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Polling frequency (every 1 to 30 seconds).
	PollingFrequency pulumi.IntOutput `pulumi:"pollingFrequency"`
	// Port to communicate with this Active Directory server.
	Port pulumi.IntOutput `pulumi:"port"`
	// Host name or IP address of the Active Directory server.
	Server pulumi.StringOutput `pulumi:"server"`
	// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
	SmbNtlmv1Auth pulumi.StringOutput `pulumi:"smbNtlmv1Auth"`
	// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
	Smbv1 pulumi.StringOutput `pulumi:"smbv1"`
	// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// User name required to log into this Active Directory server.
	User pulumi.StringOutput `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFssopolling registers a new resource with the given unique name, arguments, and options.
func NewFssopolling(ctx *pulumi.Context,
	name string, args *FssopollingArgs, opts ...pulumi.ResourceOption) (*Fssopolling, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LdapServer == nil {
		return nil, errors.New("invalid value for required argument 'LdapServer'")
	}
	if args.Server == nil {
		return nil, errors.New("invalid value for required argument 'Server'")
	}
	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource Fssopolling
	err := ctx.RegisterResource("fortios:user/fssopolling:Fssopolling", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFssopolling gets an existing Fssopolling resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFssopolling(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FssopollingState, opts ...pulumi.ResourceOption) (*Fssopolling, error) {
	var resource Fssopolling
	err := ctx.ReadResource("fortios:user/fssopolling:Fssopolling", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fssopolling resources.
type fssopollingState struct {
	// LDAP Group Info. The structure of `adgrp` block is documented below.
	Adgrps []FssopollingAdgrp `pulumi:"adgrps"`
	// Default domain managed by this Active Directory server.
	DefaultDomain *string `pulumi:"defaultDomain"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Active Directory server ID.
	Fosid *int `pulumi:"fosid"`
	// LDAP server name used in LDAP connection strings.
	LdapServer *string `pulumi:"ldapServer"`
	// Number of hours of logon history to keep, 0 means keep all history.
	LogonHistory *int `pulumi:"logonHistory"`
	// Password required to log into this Active Directory server
	Password *string `pulumi:"password"`
	// Polling frequency (every 1 to 30 seconds).
	PollingFrequency *int `pulumi:"pollingFrequency"`
	// Port to communicate with this Active Directory server.
	Port *int `pulumi:"port"`
	// Host name or IP address of the Active Directory server.
	Server *string `pulumi:"server"`
	// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
	SmbNtlmv1Auth *string `pulumi:"smbNtlmv1Auth"`
	// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
	Smbv1 *string `pulumi:"smbv1"`
	// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// User name required to log into this Active Directory server.
	User *string `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FssopollingState struct {
	// LDAP Group Info. The structure of `adgrp` block is documented below.
	Adgrps FssopollingAdgrpArrayInput
	// Default domain managed by this Active Directory server.
	DefaultDomain pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Active Directory server ID.
	Fosid pulumi.IntPtrInput
	// LDAP server name used in LDAP connection strings.
	LdapServer pulumi.StringPtrInput
	// Number of hours of logon history to keep, 0 means keep all history.
	LogonHistory pulumi.IntPtrInput
	// Password required to log into this Active Directory server
	Password pulumi.StringPtrInput
	// Polling frequency (every 1 to 30 seconds).
	PollingFrequency pulumi.IntPtrInput
	// Port to communicate with this Active Directory server.
	Port pulumi.IntPtrInput
	// Host name or IP address of the Active Directory server.
	Server pulumi.StringPtrInput
	// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
	SmbNtlmv1Auth pulumi.StringPtrInput
	// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
	Smbv1 pulumi.StringPtrInput
	// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// User name required to log into this Active Directory server.
	User pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FssopollingState) ElementType() reflect.Type {
	return reflect.TypeOf((*fssopollingState)(nil)).Elem()
}

type fssopollingArgs struct {
	// LDAP Group Info. The structure of `adgrp` block is documented below.
	Adgrps []FssopollingAdgrp `pulumi:"adgrps"`
	// Default domain managed by this Active Directory server.
	DefaultDomain *string `pulumi:"defaultDomain"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Active Directory server ID.
	Fosid *int `pulumi:"fosid"`
	// LDAP server name used in LDAP connection strings.
	LdapServer string `pulumi:"ldapServer"`
	// Number of hours of logon history to keep, 0 means keep all history.
	LogonHistory *int `pulumi:"logonHistory"`
	// Password required to log into this Active Directory server
	Password *string `pulumi:"password"`
	// Polling frequency (every 1 to 30 seconds).
	PollingFrequency *int `pulumi:"pollingFrequency"`
	// Port to communicate with this Active Directory server.
	Port *int `pulumi:"port"`
	// Host name or IP address of the Active Directory server.
	Server string `pulumi:"server"`
	// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
	SmbNtlmv1Auth *string `pulumi:"smbNtlmv1Auth"`
	// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
	Smbv1 *string `pulumi:"smbv1"`
	// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// User name required to log into this Active Directory server.
	User string `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Fssopolling resource.
type FssopollingArgs struct {
	// LDAP Group Info. The structure of `adgrp` block is documented below.
	Adgrps FssopollingAdgrpArrayInput
	// Default domain managed by this Active Directory server.
	DefaultDomain pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Active Directory server ID.
	Fosid pulumi.IntPtrInput
	// LDAP server name used in LDAP connection strings.
	LdapServer pulumi.StringInput
	// Number of hours of logon history to keep, 0 means keep all history.
	LogonHistory pulumi.IntPtrInput
	// Password required to log into this Active Directory server
	Password pulumi.StringPtrInput
	// Polling frequency (every 1 to 30 seconds).
	PollingFrequency pulumi.IntPtrInput
	// Port to communicate with this Active Directory server.
	Port pulumi.IntPtrInput
	// Host name or IP address of the Active Directory server.
	Server pulumi.StringInput
	// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
	SmbNtlmv1Auth pulumi.StringPtrInput
	// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
	Smbv1 pulumi.StringPtrInput
	// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// User name required to log into this Active Directory server.
	User pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FssopollingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fssopollingArgs)(nil)).Elem()
}

type FssopollingInput interface {
	pulumi.Input

	ToFssopollingOutput() FssopollingOutput
	ToFssopollingOutputWithContext(ctx context.Context) FssopollingOutput
}

func (*Fssopolling) ElementType() reflect.Type {
	return reflect.TypeOf((**Fssopolling)(nil)).Elem()
}

func (i *Fssopolling) ToFssopollingOutput() FssopollingOutput {
	return i.ToFssopollingOutputWithContext(context.Background())
}

func (i *Fssopolling) ToFssopollingOutputWithContext(ctx context.Context) FssopollingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FssopollingOutput)
}

// FssopollingArrayInput is an input type that accepts FssopollingArray and FssopollingArrayOutput values.
// You can construct a concrete instance of `FssopollingArrayInput` via:
//
//	FssopollingArray{ FssopollingArgs{...} }
type FssopollingArrayInput interface {
	pulumi.Input

	ToFssopollingArrayOutput() FssopollingArrayOutput
	ToFssopollingArrayOutputWithContext(context.Context) FssopollingArrayOutput
}

type FssopollingArray []FssopollingInput

func (FssopollingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fssopolling)(nil)).Elem()
}

func (i FssopollingArray) ToFssopollingArrayOutput() FssopollingArrayOutput {
	return i.ToFssopollingArrayOutputWithContext(context.Background())
}

func (i FssopollingArray) ToFssopollingArrayOutputWithContext(ctx context.Context) FssopollingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FssopollingArrayOutput)
}

// FssopollingMapInput is an input type that accepts FssopollingMap and FssopollingMapOutput values.
// You can construct a concrete instance of `FssopollingMapInput` via:
//
//	FssopollingMap{ "key": FssopollingArgs{...} }
type FssopollingMapInput interface {
	pulumi.Input

	ToFssopollingMapOutput() FssopollingMapOutput
	ToFssopollingMapOutputWithContext(context.Context) FssopollingMapOutput
}

type FssopollingMap map[string]FssopollingInput

func (FssopollingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fssopolling)(nil)).Elem()
}

func (i FssopollingMap) ToFssopollingMapOutput() FssopollingMapOutput {
	return i.ToFssopollingMapOutputWithContext(context.Background())
}

func (i FssopollingMap) ToFssopollingMapOutputWithContext(ctx context.Context) FssopollingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FssopollingMapOutput)
}

type FssopollingOutput struct{ *pulumi.OutputState }

func (FssopollingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fssopolling)(nil)).Elem()
}

func (o FssopollingOutput) ToFssopollingOutput() FssopollingOutput {
	return o
}

func (o FssopollingOutput) ToFssopollingOutputWithContext(ctx context.Context) FssopollingOutput {
	return o
}

// LDAP Group Info. The structure of `adgrp` block is documented below.
func (o FssopollingOutput) Adgrps() FssopollingAdgrpArrayOutput {
	return o.ApplyT(func(v *Fssopolling) FssopollingAdgrpArrayOutput { return v.Adgrps }).(FssopollingAdgrpArrayOutput)
}

// Default domain managed by this Active Directory server.
func (o FssopollingOutput) DefaultDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *Fssopolling) pulumi.StringOutput { return v.DefaultDomain }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o FssopollingOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Fssopolling) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Active Directory server ID.
func (o FssopollingOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Fssopolling) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// LDAP server name used in LDAP connection strings.
func (o FssopollingOutput) LdapServer() pulumi.StringOutput {
	return o.ApplyT(func(v *Fssopolling) pulumi.StringOutput { return v.LdapServer }).(pulumi.StringOutput)
}

// Number of hours of logon history to keep, 0 means keep all history.
func (o FssopollingOutput) LogonHistory() pulumi.IntOutput {
	return o.ApplyT(func(v *Fssopolling) pulumi.IntOutput { return v.LogonHistory }).(pulumi.IntOutput)
}

// Password required to log into this Active Directory server
func (o FssopollingOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Fssopolling) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Polling frequency (every 1 to 30 seconds).
func (o FssopollingOutput) PollingFrequency() pulumi.IntOutput {
	return o.ApplyT(func(v *Fssopolling) pulumi.IntOutput { return v.PollingFrequency }).(pulumi.IntOutput)
}

// Port to communicate with this Active Directory server.
func (o FssopollingOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Fssopolling) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Host name or IP address of the Active Directory server.
func (o FssopollingOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *Fssopolling) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
func (o FssopollingOutput) SmbNtlmv1Auth() pulumi.StringOutput {
	return o.ApplyT(func(v *Fssopolling) pulumi.StringOutput { return v.SmbNtlmv1Auth }).(pulumi.StringOutput)
}

// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
func (o FssopollingOutput) Smbv1() pulumi.StringOutput {
	return o.ApplyT(func(v *Fssopolling) pulumi.StringOutput { return v.Smbv1 }).(pulumi.StringOutput)
}

// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
func (o FssopollingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Fssopolling) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// User name required to log into this Active Directory server.
func (o FssopollingOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *Fssopolling) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FssopollingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Fssopolling) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FssopollingArrayOutput struct{ *pulumi.OutputState }

func (FssopollingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fssopolling)(nil)).Elem()
}

func (o FssopollingArrayOutput) ToFssopollingArrayOutput() FssopollingArrayOutput {
	return o
}

func (o FssopollingArrayOutput) ToFssopollingArrayOutputWithContext(ctx context.Context) FssopollingArrayOutput {
	return o
}

func (o FssopollingArrayOutput) Index(i pulumi.IntInput) FssopollingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Fssopolling {
		return vs[0].([]*Fssopolling)[vs[1].(int)]
	}).(FssopollingOutput)
}

type FssopollingMapOutput struct{ *pulumi.OutputState }

func (FssopollingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fssopolling)(nil)).Elem()
}

func (o FssopollingMapOutput) ToFssopollingMapOutput() FssopollingMapOutput {
	return o
}

func (o FssopollingMapOutput) ToFssopollingMapOutputWithContext(ctx context.Context) FssopollingMapOutput {
	return o
}

func (o FssopollingMapOutput) MapIndex(k pulumi.StringInput) FssopollingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Fssopolling {
		return vs[0].(map[string]*Fssopolling)[vs[1].(string)]
	}).(FssopollingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FssopollingInput)(nil)).Elem(), &Fssopolling{})
	pulumi.RegisterInputType(reflect.TypeOf((*FssopollingArrayInput)(nil)).Elem(), FssopollingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FssopollingMapInput)(nil)).Elem(), FssopollingMap{})
	pulumi.RegisterOutputType(FssopollingOutput{})
	pulumi.RegisterOutputType(FssopollingArrayOutput{})
	pulumi.RegisterOutputType(FssopollingMapOutput{})
}
