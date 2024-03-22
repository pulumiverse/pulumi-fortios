// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package endpointcontrol

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure FortiClient Enterprise Management Server (EMS) entries. Applies to FortiOS Version `>= 7.4.0`.
//
// ## Import
//
// EndpointControl FctemsOverride can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:endpointcontrol/fctemsoverride:Fctemsoverride labelname {{ems_id}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:endpointcontrol/fctemsoverride:Fctemsoverride labelname {{ems_id}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Fctemsoverride struct {
	pulumi.CustomResourceState

	// FortiClient EMS call timeout in seconds (1 - 180 seconds, default = 30).
	CallTimeout pulumi.IntOutput `pulumi:"callTimeout"`
	// List of EMS capabilities.
	Capabilities pulumi.StringOutput `pulumi:"capabilities"`
	// Cloud server type. Valid values: `production`, `alpha`, `beta`.
	CloudServerType pulumi.StringOutput `pulumi:"cloudServerType"`
	// Dirty Reason for FortiClient EMS. Valid values: `none`, `mismatched-ems-sn`.
	DirtyReason pulumi.StringOutput `pulumi:"dirtyReason"`
	// EMS ID in order (1 - 7).
	EmsId pulumi.IntOutput `pulumi:"emsId"`
	// Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `enable`, `disable`.
	FortinetoneCloudAuthentication pulumi.StringOutput `pulumi:"fortinetoneCloudAuthentication"`
	// FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
	HttpsPort pulumi.IntOutput `pulumi:"httpsPort"`
	// Specify outgoing interface to reach server.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringOutput `pulumi:"interfaceSelectMethod"`
	// FortiClient Enterprise Management Server (EMS) name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Outdated resource threshold in seconds (10 - 3600, default = 180).
	OutOfSyncThreshold pulumi.IntOutput `pulumi:"outOfSyncThreshold"`
	// Enable/disable preservation of EMS SSL session connection. Warning, most users should not touch this setting. Valid values: `enable`, `disable`.
	PreserveSslSession pulumi.StringOutput `pulumi:"preserveSslSession"`
	// Enable/disable pulling avatars from EMS. Valid values: `enable`, `disable`.
	PullAvatars pulumi.StringOutput `pulumi:"pullAvatars"`
	// Enable/disable pulling FortiClient malware hash from EMS. Valid values: `enable`, `disable`.
	PullMalwareHash pulumi.StringOutput `pulumi:"pullMalwareHash"`
	// Enable/disable pulling SysInfo from EMS. Valid values: `enable`, `disable`.
	PullSysinfo pulumi.StringOutput `pulumi:"pullSysinfo"`
	// Enable/disable pulling FortiClient user tags from EMS. Valid values: `enable`, `disable`.
	PullTags pulumi.StringOutput `pulumi:"pullTags"`
	// Enable/disable pulling vulnerabilities from EMS. Valid values: `enable`, `disable`.
	PullVulnerabilities pulumi.StringOutput `pulumi:"pullVulnerabilities"`
	// Relax restrictions on tags to send all EMS tags to all VDOMs Valid values: `enable`, `disable`.
	SendTagsToAllVdoms pulumi.StringOutput `pulumi:"sendTagsToAllVdoms"`
	// EMS Serial Number.
	SerialNumber pulumi.StringOutput `pulumi:"serialNumber"`
	// FortiClient EMS FQDN or IPv4 address.
	Server pulumi.StringOutput `pulumi:"server"`
	// REST API call source IP.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Enable or disable this EMS configuration. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// EMS Tenant ID.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Enable/disable trust of the EMS certificate issuer(CA) and common name(CN) for certificate auto-renewal. Valid values: `enable`, `disable`.
	TrustCaCn pulumi.StringOutput `pulumi:"trustCaCn"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Lowest CA cert on Fortigate in verified EMS cert chain.
	VerifyingCa pulumi.StringOutput `pulumi:"verifyingCa"`
	// Enable/disable override behavior for how this FortiGate unit connects to EMS using a WebSocket connection. Valid values: `disable`, `enable`.
	WebsocketOverride pulumi.StringOutput `pulumi:"websocketOverride"`
}

// NewFctemsoverride registers a new resource with the given unique name, arguments, and options.
func NewFctemsoverride(ctx *pulumi.Context,
	name string, args *FctemsoverrideArgs, opts ...pulumi.ResourceOption) (*Fctemsoverride, error) {
	if args == nil {
		args = &FctemsoverrideArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Fctemsoverride
	err := ctx.RegisterResource("fortios:endpointcontrol/fctemsoverride:Fctemsoverride", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFctemsoverride gets an existing Fctemsoverride resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFctemsoverride(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FctemsoverrideState, opts ...pulumi.ResourceOption) (*Fctemsoverride, error) {
	var resource Fctemsoverride
	err := ctx.ReadResource("fortios:endpointcontrol/fctemsoverride:Fctemsoverride", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fctemsoverride resources.
type fctemsoverrideState struct {
	// FortiClient EMS call timeout in seconds (1 - 180 seconds, default = 30).
	CallTimeout *int `pulumi:"callTimeout"`
	// List of EMS capabilities.
	Capabilities *string `pulumi:"capabilities"`
	// Cloud server type. Valid values: `production`, `alpha`, `beta`.
	CloudServerType *string `pulumi:"cloudServerType"`
	// Dirty Reason for FortiClient EMS. Valid values: `none`, `mismatched-ems-sn`.
	DirtyReason *string `pulumi:"dirtyReason"`
	// EMS ID in order (1 - 7).
	EmsId *int `pulumi:"emsId"`
	// Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `enable`, `disable`.
	FortinetoneCloudAuthentication *string `pulumi:"fortinetoneCloudAuthentication"`
	// FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
	HttpsPort *int `pulumi:"httpsPort"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// FortiClient Enterprise Management Server (EMS) name.
	Name *string `pulumi:"name"`
	// Outdated resource threshold in seconds (10 - 3600, default = 180).
	OutOfSyncThreshold *int `pulumi:"outOfSyncThreshold"`
	// Enable/disable preservation of EMS SSL session connection. Warning, most users should not touch this setting. Valid values: `enable`, `disable`.
	PreserveSslSession *string `pulumi:"preserveSslSession"`
	// Enable/disable pulling avatars from EMS. Valid values: `enable`, `disable`.
	PullAvatars *string `pulumi:"pullAvatars"`
	// Enable/disable pulling FortiClient malware hash from EMS. Valid values: `enable`, `disable`.
	PullMalwareHash *string `pulumi:"pullMalwareHash"`
	// Enable/disable pulling SysInfo from EMS. Valid values: `enable`, `disable`.
	PullSysinfo *string `pulumi:"pullSysinfo"`
	// Enable/disable pulling FortiClient user tags from EMS. Valid values: `enable`, `disable`.
	PullTags *string `pulumi:"pullTags"`
	// Enable/disable pulling vulnerabilities from EMS. Valid values: `enable`, `disable`.
	PullVulnerabilities *string `pulumi:"pullVulnerabilities"`
	// Relax restrictions on tags to send all EMS tags to all VDOMs Valid values: `enable`, `disable`.
	SendTagsToAllVdoms *string `pulumi:"sendTagsToAllVdoms"`
	// EMS Serial Number.
	SerialNumber *string `pulumi:"serialNumber"`
	// FortiClient EMS FQDN or IPv4 address.
	Server *string `pulumi:"server"`
	// REST API call source IP.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable or disable this EMS configuration. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// EMS Tenant ID.
	TenantId *string `pulumi:"tenantId"`
	// Enable/disable trust of the EMS certificate issuer(CA) and common name(CN) for certificate auto-renewal. Valid values: `enable`, `disable`.
	TrustCaCn *string `pulumi:"trustCaCn"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Lowest CA cert on Fortigate in verified EMS cert chain.
	VerifyingCa *string `pulumi:"verifyingCa"`
	// Enable/disable override behavior for how this FortiGate unit connects to EMS using a WebSocket connection. Valid values: `disable`, `enable`.
	WebsocketOverride *string `pulumi:"websocketOverride"`
}

type FctemsoverrideState struct {
	// FortiClient EMS call timeout in seconds (1 - 180 seconds, default = 30).
	CallTimeout pulumi.IntPtrInput
	// List of EMS capabilities.
	Capabilities pulumi.StringPtrInput
	// Cloud server type. Valid values: `production`, `alpha`, `beta`.
	CloudServerType pulumi.StringPtrInput
	// Dirty Reason for FortiClient EMS. Valid values: `none`, `mismatched-ems-sn`.
	DirtyReason pulumi.StringPtrInput
	// EMS ID in order (1 - 7).
	EmsId pulumi.IntPtrInput
	// Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `enable`, `disable`.
	FortinetoneCloudAuthentication pulumi.StringPtrInput
	// FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
	HttpsPort pulumi.IntPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// FortiClient Enterprise Management Server (EMS) name.
	Name pulumi.StringPtrInput
	// Outdated resource threshold in seconds (10 - 3600, default = 180).
	OutOfSyncThreshold pulumi.IntPtrInput
	// Enable/disable preservation of EMS SSL session connection. Warning, most users should not touch this setting. Valid values: `enable`, `disable`.
	PreserveSslSession pulumi.StringPtrInput
	// Enable/disable pulling avatars from EMS. Valid values: `enable`, `disable`.
	PullAvatars pulumi.StringPtrInput
	// Enable/disable pulling FortiClient malware hash from EMS. Valid values: `enable`, `disable`.
	PullMalwareHash pulumi.StringPtrInput
	// Enable/disable pulling SysInfo from EMS. Valid values: `enable`, `disable`.
	PullSysinfo pulumi.StringPtrInput
	// Enable/disable pulling FortiClient user tags from EMS. Valid values: `enable`, `disable`.
	PullTags pulumi.StringPtrInput
	// Enable/disable pulling vulnerabilities from EMS. Valid values: `enable`, `disable`.
	PullVulnerabilities pulumi.StringPtrInput
	// Relax restrictions on tags to send all EMS tags to all VDOMs Valid values: `enable`, `disable`.
	SendTagsToAllVdoms pulumi.StringPtrInput
	// EMS Serial Number.
	SerialNumber pulumi.StringPtrInput
	// FortiClient EMS FQDN or IPv4 address.
	Server pulumi.StringPtrInput
	// REST API call source IP.
	SourceIp pulumi.StringPtrInput
	// Enable or disable this EMS configuration. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// EMS Tenant ID.
	TenantId pulumi.StringPtrInput
	// Enable/disable trust of the EMS certificate issuer(CA) and common name(CN) for certificate auto-renewal. Valid values: `enable`, `disable`.
	TrustCaCn pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Lowest CA cert on Fortigate in verified EMS cert chain.
	VerifyingCa pulumi.StringPtrInput
	// Enable/disable override behavior for how this FortiGate unit connects to EMS using a WebSocket connection. Valid values: `disable`, `enable`.
	WebsocketOverride pulumi.StringPtrInput
}

func (FctemsoverrideState) ElementType() reflect.Type {
	return reflect.TypeOf((*fctemsoverrideState)(nil)).Elem()
}

type fctemsoverrideArgs struct {
	// FortiClient EMS call timeout in seconds (1 - 180 seconds, default = 30).
	CallTimeout *int `pulumi:"callTimeout"`
	// List of EMS capabilities.
	Capabilities *string `pulumi:"capabilities"`
	// Cloud server type. Valid values: `production`, `alpha`, `beta`.
	CloudServerType *string `pulumi:"cloudServerType"`
	// Dirty Reason for FortiClient EMS. Valid values: `none`, `mismatched-ems-sn`.
	DirtyReason *string `pulumi:"dirtyReason"`
	// EMS ID in order (1 - 7).
	EmsId *int `pulumi:"emsId"`
	// Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `enable`, `disable`.
	FortinetoneCloudAuthentication *string `pulumi:"fortinetoneCloudAuthentication"`
	// FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
	HttpsPort *int `pulumi:"httpsPort"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// FortiClient Enterprise Management Server (EMS) name.
	Name *string `pulumi:"name"`
	// Outdated resource threshold in seconds (10 - 3600, default = 180).
	OutOfSyncThreshold *int `pulumi:"outOfSyncThreshold"`
	// Enable/disable preservation of EMS SSL session connection. Warning, most users should not touch this setting. Valid values: `enable`, `disable`.
	PreserveSslSession *string `pulumi:"preserveSslSession"`
	// Enable/disable pulling avatars from EMS. Valid values: `enable`, `disable`.
	PullAvatars *string `pulumi:"pullAvatars"`
	// Enable/disable pulling FortiClient malware hash from EMS. Valid values: `enable`, `disable`.
	PullMalwareHash *string `pulumi:"pullMalwareHash"`
	// Enable/disable pulling SysInfo from EMS. Valid values: `enable`, `disable`.
	PullSysinfo *string `pulumi:"pullSysinfo"`
	// Enable/disable pulling FortiClient user tags from EMS. Valid values: `enable`, `disable`.
	PullTags *string `pulumi:"pullTags"`
	// Enable/disable pulling vulnerabilities from EMS. Valid values: `enable`, `disable`.
	PullVulnerabilities *string `pulumi:"pullVulnerabilities"`
	// Relax restrictions on tags to send all EMS tags to all VDOMs Valid values: `enable`, `disable`.
	SendTagsToAllVdoms *string `pulumi:"sendTagsToAllVdoms"`
	// EMS Serial Number.
	SerialNumber *string `pulumi:"serialNumber"`
	// FortiClient EMS FQDN or IPv4 address.
	Server *string `pulumi:"server"`
	// REST API call source IP.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable or disable this EMS configuration. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// EMS Tenant ID.
	TenantId *string `pulumi:"tenantId"`
	// Enable/disable trust of the EMS certificate issuer(CA) and common name(CN) for certificate auto-renewal. Valid values: `enable`, `disable`.
	TrustCaCn *string `pulumi:"trustCaCn"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Lowest CA cert on Fortigate in verified EMS cert chain.
	VerifyingCa *string `pulumi:"verifyingCa"`
	// Enable/disable override behavior for how this FortiGate unit connects to EMS using a WebSocket connection. Valid values: `disable`, `enable`.
	WebsocketOverride *string `pulumi:"websocketOverride"`
}

// The set of arguments for constructing a Fctemsoverride resource.
type FctemsoverrideArgs struct {
	// FortiClient EMS call timeout in seconds (1 - 180 seconds, default = 30).
	CallTimeout pulumi.IntPtrInput
	// List of EMS capabilities.
	Capabilities pulumi.StringPtrInput
	// Cloud server type. Valid values: `production`, `alpha`, `beta`.
	CloudServerType pulumi.StringPtrInput
	// Dirty Reason for FortiClient EMS. Valid values: `none`, `mismatched-ems-sn`.
	DirtyReason pulumi.StringPtrInput
	// EMS ID in order (1 - 7).
	EmsId pulumi.IntPtrInput
	// Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `enable`, `disable`.
	FortinetoneCloudAuthentication pulumi.StringPtrInput
	// FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
	HttpsPort pulumi.IntPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// FortiClient Enterprise Management Server (EMS) name.
	Name pulumi.StringPtrInput
	// Outdated resource threshold in seconds (10 - 3600, default = 180).
	OutOfSyncThreshold pulumi.IntPtrInput
	// Enable/disable preservation of EMS SSL session connection. Warning, most users should not touch this setting. Valid values: `enable`, `disable`.
	PreserveSslSession pulumi.StringPtrInput
	// Enable/disable pulling avatars from EMS. Valid values: `enable`, `disable`.
	PullAvatars pulumi.StringPtrInput
	// Enable/disable pulling FortiClient malware hash from EMS. Valid values: `enable`, `disable`.
	PullMalwareHash pulumi.StringPtrInput
	// Enable/disable pulling SysInfo from EMS. Valid values: `enable`, `disable`.
	PullSysinfo pulumi.StringPtrInput
	// Enable/disable pulling FortiClient user tags from EMS. Valid values: `enable`, `disable`.
	PullTags pulumi.StringPtrInput
	// Enable/disable pulling vulnerabilities from EMS. Valid values: `enable`, `disable`.
	PullVulnerabilities pulumi.StringPtrInput
	// Relax restrictions on tags to send all EMS tags to all VDOMs Valid values: `enable`, `disable`.
	SendTagsToAllVdoms pulumi.StringPtrInput
	// EMS Serial Number.
	SerialNumber pulumi.StringPtrInput
	// FortiClient EMS FQDN or IPv4 address.
	Server pulumi.StringPtrInput
	// REST API call source IP.
	SourceIp pulumi.StringPtrInput
	// Enable or disable this EMS configuration. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// EMS Tenant ID.
	TenantId pulumi.StringPtrInput
	// Enable/disable trust of the EMS certificate issuer(CA) and common name(CN) for certificate auto-renewal. Valid values: `enable`, `disable`.
	TrustCaCn pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Lowest CA cert on Fortigate in verified EMS cert chain.
	VerifyingCa pulumi.StringPtrInput
	// Enable/disable override behavior for how this FortiGate unit connects to EMS using a WebSocket connection. Valid values: `disable`, `enable`.
	WebsocketOverride pulumi.StringPtrInput
}

func (FctemsoverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fctemsoverrideArgs)(nil)).Elem()
}

type FctemsoverrideInput interface {
	pulumi.Input

	ToFctemsoverrideOutput() FctemsoverrideOutput
	ToFctemsoverrideOutputWithContext(ctx context.Context) FctemsoverrideOutput
}

func (*Fctemsoverride) ElementType() reflect.Type {
	return reflect.TypeOf((**Fctemsoverride)(nil)).Elem()
}

func (i *Fctemsoverride) ToFctemsoverrideOutput() FctemsoverrideOutput {
	return i.ToFctemsoverrideOutputWithContext(context.Background())
}

func (i *Fctemsoverride) ToFctemsoverrideOutputWithContext(ctx context.Context) FctemsoverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FctemsoverrideOutput)
}

// FctemsoverrideArrayInput is an input type that accepts FctemsoverrideArray and FctemsoverrideArrayOutput values.
// You can construct a concrete instance of `FctemsoverrideArrayInput` via:
//
//	FctemsoverrideArray{ FctemsoverrideArgs{...} }
type FctemsoverrideArrayInput interface {
	pulumi.Input

	ToFctemsoverrideArrayOutput() FctemsoverrideArrayOutput
	ToFctemsoverrideArrayOutputWithContext(context.Context) FctemsoverrideArrayOutput
}

type FctemsoverrideArray []FctemsoverrideInput

func (FctemsoverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fctemsoverride)(nil)).Elem()
}

func (i FctemsoverrideArray) ToFctemsoverrideArrayOutput() FctemsoverrideArrayOutput {
	return i.ToFctemsoverrideArrayOutputWithContext(context.Background())
}

func (i FctemsoverrideArray) ToFctemsoverrideArrayOutputWithContext(ctx context.Context) FctemsoverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FctemsoverrideArrayOutput)
}

// FctemsoverrideMapInput is an input type that accepts FctemsoverrideMap and FctemsoverrideMapOutput values.
// You can construct a concrete instance of `FctemsoverrideMapInput` via:
//
//	FctemsoverrideMap{ "key": FctemsoverrideArgs{...} }
type FctemsoverrideMapInput interface {
	pulumi.Input

	ToFctemsoverrideMapOutput() FctemsoverrideMapOutput
	ToFctemsoverrideMapOutputWithContext(context.Context) FctemsoverrideMapOutput
}

type FctemsoverrideMap map[string]FctemsoverrideInput

func (FctemsoverrideMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fctemsoverride)(nil)).Elem()
}

func (i FctemsoverrideMap) ToFctemsoverrideMapOutput() FctemsoverrideMapOutput {
	return i.ToFctemsoverrideMapOutputWithContext(context.Background())
}

func (i FctemsoverrideMap) ToFctemsoverrideMapOutputWithContext(ctx context.Context) FctemsoverrideMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FctemsoverrideMapOutput)
}

type FctemsoverrideOutput struct{ *pulumi.OutputState }

func (FctemsoverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fctemsoverride)(nil)).Elem()
}

func (o FctemsoverrideOutput) ToFctemsoverrideOutput() FctemsoverrideOutput {
	return o
}

func (o FctemsoverrideOutput) ToFctemsoverrideOutputWithContext(ctx context.Context) FctemsoverrideOutput {
	return o
}

// FortiClient EMS call timeout in seconds (1 - 180 seconds, default = 30).
func (o FctemsoverrideOutput) CallTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.IntOutput { return v.CallTimeout }).(pulumi.IntOutput)
}

// List of EMS capabilities.
func (o FctemsoverrideOutput) Capabilities() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.Capabilities }).(pulumi.StringOutput)
}

// Cloud server type. Valid values: `production`, `alpha`, `beta`.
func (o FctemsoverrideOutput) CloudServerType() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.CloudServerType }).(pulumi.StringOutput)
}

// Dirty Reason for FortiClient EMS. Valid values: `none`, `mismatched-ems-sn`.
func (o FctemsoverrideOutput) DirtyReason() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.DirtyReason }).(pulumi.StringOutput)
}

// EMS ID in order (1 - 7).
func (o FctemsoverrideOutput) EmsId() pulumi.IntOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.IntOutput { return v.EmsId }).(pulumi.IntOutput)
}

// Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `enable`, `disable`.
func (o FctemsoverrideOutput) FortinetoneCloudAuthentication() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.FortinetoneCloudAuthentication }).(pulumi.StringOutput)
}

// FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
func (o FctemsoverrideOutput) HttpsPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.IntOutput { return v.HttpsPort }).(pulumi.IntOutput)
}

// Specify outgoing interface to reach server.
func (o FctemsoverrideOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
func (o FctemsoverrideOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

// FortiClient Enterprise Management Server (EMS) name.
func (o FctemsoverrideOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Outdated resource threshold in seconds (10 - 3600, default = 180).
func (o FctemsoverrideOutput) OutOfSyncThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.IntOutput { return v.OutOfSyncThreshold }).(pulumi.IntOutput)
}

// Enable/disable preservation of EMS SSL session connection. Warning, most users should not touch this setting. Valid values: `enable`, `disable`.
func (o FctemsoverrideOutput) PreserveSslSession() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.PreserveSslSession }).(pulumi.StringOutput)
}

// Enable/disable pulling avatars from EMS. Valid values: `enable`, `disable`.
func (o FctemsoverrideOutput) PullAvatars() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.PullAvatars }).(pulumi.StringOutput)
}

// Enable/disable pulling FortiClient malware hash from EMS. Valid values: `enable`, `disable`.
func (o FctemsoverrideOutput) PullMalwareHash() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.PullMalwareHash }).(pulumi.StringOutput)
}

// Enable/disable pulling SysInfo from EMS. Valid values: `enable`, `disable`.
func (o FctemsoverrideOutput) PullSysinfo() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.PullSysinfo }).(pulumi.StringOutput)
}

// Enable/disable pulling FortiClient user tags from EMS. Valid values: `enable`, `disable`.
func (o FctemsoverrideOutput) PullTags() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.PullTags }).(pulumi.StringOutput)
}

// Enable/disable pulling vulnerabilities from EMS. Valid values: `enable`, `disable`.
func (o FctemsoverrideOutput) PullVulnerabilities() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.PullVulnerabilities }).(pulumi.StringOutput)
}

// Relax restrictions on tags to send all EMS tags to all VDOMs Valid values: `enable`, `disable`.
func (o FctemsoverrideOutput) SendTagsToAllVdoms() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.SendTagsToAllVdoms }).(pulumi.StringOutput)
}

// EMS Serial Number.
func (o FctemsoverrideOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

// FortiClient EMS FQDN or IPv4 address.
func (o FctemsoverrideOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

// REST API call source IP.
func (o FctemsoverrideOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// Enable or disable this EMS configuration. Valid values: `enable`, `disable`.
func (o FctemsoverrideOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// EMS Tenant ID.
func (o FctemsoverrideOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Enable/disable trust of the EMS certificate issuer(CA) and common name(CN) for certificate auto-renewal. Valid values: `enable`, `disable`.
func (o FctemsoverrideOutput) TrustCaCn() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.TrustCaCn }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FctemsoverrideOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Lowest CA cert on Fortigate in verified EMS cert chain.
func (o FctemsoverrideOutput) VerifyingCa() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.VerifyingCa }).(pulumi.StringOutput)
}

// Enable/disable override behavior for how this FortiGate unit connects to EMS using a WebSocket connection. Valid values: `disable`, `enable`.
func (o FctemsoverrideOutput) WebsocketOverride() pulumi.StringOutput {
	return o.ApplyT(func(v *Fctemsoverride) pulumi.StringOutput { return v.WebsocketOverride }).(pulumi.StringOutput)
}

type FctemsoverrideArrayOutput struct{ *pulumi.OutputState }

func (FctemsoverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fctemsoverride)(nil)).Elem()
}

func (o FctemsoverrideArrayOutput) ToFctemsoverrideArrayOutput() FctemsoverrideArrayOutput {
	return o
}

func (o FctemsoverrideArrayOutput) ToFctemsoverrideArrayOutputWithContext(ctx context.Context) FctemsoverrideArrayOutput {
	return o
}

func (o FctemsoverrideArrayOutput) Index(i pulumi.IntInput) FctemsoverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Fctemsoverride {
		return vs[0].([]*Fctemsoverride)[vs[1].(int)]
	}).(FctemsoverrideOutput)
}

type FctemsoverrideMapOutput struct{ *pulumi.OutputState }

func (FctemsoverrideMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fctemsoverride)(nil)).Elem()
}

func (o FctemsoverrideMapOutput) ToFctemsoverrideMapOutput() FctemsoverrideMapOutput {
	return o
}

func (o FctemsoverrideMapOutput) ToFctemsoverrideMapOutputWithContext(ctx context.Context) FctemsoverrideMapOutput {
	return o
}

func (o FctemsoverrideMapOutput) MapIndex(k pulumi.StringInput) FctemsoverrideOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Fctemsoverride {
		return vs[0].(map[string]*Fctemsoverride)[vs[1].(string)]
	}).(FctemsoverrideOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FctemsoverrideInput)(nil)).Elem(), &Fctemsoverride{})
	pulumi.RegisterInputType(reflect.TypeOf((*FctemsoverrideArrayInput)(nil)).Elem(), FctemsoverrideArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FctemsoverrideMapInput)(nil)).Elem(), FctemsoverrideMap{})
	pulumi.RegisterOutputType(FctemsoverrideOutput{})
	pulumi.RegisterOutputType(FctemsoverrideArrayOutput{})
	pulumi.RegisterOutputType(FctemsoverrideMapOutput{})
}
