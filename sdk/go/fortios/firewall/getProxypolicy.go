// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall proxypolicy
func LookupProxypolicy(ctx *pulumi.Context, args *LookupProxypolicyArgs, opts ...pulumi.InvokeOption) (*LookupProxypolicyResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupProxypolicyResult
	err := ctx.Invoke("fortios:firewall/getProxypolicy:getProxypolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProxypolicy.
type LookupProxypolicyArgs struct {
	// Specify the policyid of the desired firewall proxypolicy.
	Policyid int `pulumi:"policyid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getProxypolicy.
type LookupProxypolicyResult struct {
	// IPv4 access proxy. The structure of `accessProxy` block is documented below.
	AccessProxies []GetProxypolicyAccessProxy `pulumi:"accessProxies"`
	// IPv6 access proxy. The structure of `accessProxy6` block is documented below.
	AccessProxy6s []GetProxypolicyAccessProxy6 `pulumi:"accessProxy6s"`
	// Accept or deny traffic matching the policy parameters.
	Action string `pulumi:"action"`
	// Name of an existing Application list.
	ApplicationList string `pulumi:"applicationList"`
	// Name of an existing Antivirus profile.
	AvProfile string `pulumi:"avProfile"`
	// Enable/disable block notification.
	BlockNotification string `pulumi:"blockNotification"`
	// Name of an existing CIFS profile.
	CifsProfile string `pulumi:"cifsProfile"`
	// Optional comments.
	Comments string `pulumi:"comments"`
	// Decrypted traffic mirror.
	DecryptedTrafficMirror string `pulumi:"decryptedTrafficMirror"`
	// When enabled, the ownership enforcement will be done at policy level.
	DeviceOwnership string `pulumi:"deviceOwnership"`
	// Web proxy disclaimer setting: by domain, policy, or user.
	Disclaimer string `pulumi:"disclaimer"`
	// Name of an existing DLP profile.
	DlpProfile string `pulumi:"dlpProfile"`
	// Name of an existing DLP sensor.
	DlpSensor string `pulumi:"dlpSensor"`
	// IPv6 destination address objects. The structure of `dstaddr6` block is documented below.
	Dstaddr6s []GetProxypolicyDstaddr6 `pulumi:"dstaddr6s"`
	// When enabled, destination addresses match against any address EXCEPT the specified destination addresses.
	DstaddrNegate string `pulumi:"dstaddrNegate"`
	// Destination address objects. The structure of `dstaddr` block is documented below.
	Dstaddrs []GetProxypolicyDstaddr `pulumi:"dstaddrs"`
	// Destination interface names. The structure of `dstintf` block is documented below.
	Dstintfs []GetProxypolicyDstintf `pulumi:"dstintfs"`
	// Name of an existing email filter profile.
	EmailfilterProfile string `pulumi:"emailfilterProfile"`
	// Name of an existing file-filter profile.
	FileFilterProfile string `pulumi:"fileFilterProfile"`
	// Global web-based manager visible label.
	GlobalLabel string `pulumi:"globalLabel"`
	// Names of group objects. The structure of `groups` block is documented below.
	Groups []GetProxypolicyGroup `pulumi:"groups"`
	// Enable/disable HTTP tunnel authentication.
	HttpTunnelAuth string `pulumi:"httpTunnelAuth"`
	// Name of an existing ICAP profile.
	IcapProfile string `pulumi:"icapProfile"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.
	InternetService string `pulumi:"internetService"`
	// Custom Internet Service group name. The structure of `internetServiceCustomGroup` block is documented below.
	InternetServiceCustomGroups []GetProxypolicyInternetServiceCustomGroup `pulumi:"internetServiceCustomGroups"`
	// Custom Internet Service name. The structure of `internetServiceCustom` block is documented below.
	InternetServiceCustoms []GetProxypolicyInternetServiceCustom `pulumi:"internetServiceCustoms"`
	// Internet Service group name. The structure of `internetServiceGroup` block is documented below.
	InternetServiceGroups []GetProxypolicyInternetServiceGroup `pulumi:"internetServiceGroups"`
	// Internet Service ID. The structure of `internetServiceId` block is documented below.
	InternetServiceIds []GetProxypolicyInternetServiceId `pulumi:"internetServiceIds"`
	// Internet Service name. The structure of `internetServiceName` block is documented below.
	InternetServiceNames []GetProxypolicyInternetServiceName `pulumi:"internetServiceNames"`
	// When enabled, Internet Services match against any internet service EXCEPT the selected Internet Service.
	InternetServiceNegate string `pulumi:"internetServiceNegate"`
	// Name of an existing IPS sensor.
	IpsSensor string `pulumi:"ipsSensor"`
	// VDOM-specific GUI visible label.
	Label string `pulumi:"label"`
	// Enable/disable logging traffic through the policy.
	Logtraffic string `pulumi:"logtraffic"`
	// Enable/disable policy log traffic start.
	LogtrafficStart string `pulumi:"logtrafficStart"`
	// Group name.
	Name string `pulumi:"name"`
	// Policy ID.
	Policyid int `pulumi:"policyid"`
	// Name of IP pool object. The structure of `poolname` block is documented below.
	Poolnames []GetProxypolicyPoolname `pulumi:"poolnames"`
	// Name of profile group.
	ProfileGroup string `pulumi:"profileGroup"`
	// Name of an existing Protocol options profile.
	ProfileProtocolOptions string `pulumi:"profileProtocolOptions"`
	// Determine whether the firewall policy allows security profile groups or single profiles only.
	ProfileType string `pulumi:"profileType"`
	// Type of explicit proxy.
	Proxy string `pulumi:"proxy"`
	// Redirect URL for further explicit web proxy processing.
	RedirectUrl string `pulumi:"redirectUrl"`
	// Authentication replacement message override group.
	ReplacemsgOverrideGroup string `pulumi:"replacemsgOverrideGroup"`
	// Enable/disable scanning of connections to Botnet servers.
	ScanBotnetConnections string `pulumi:"scanBotnetConnections"`
	// Name of schedule object.
	Schedule string `pulumi:"schedule"`
	// Name of an existing SCTP filter profile.
	SctpFilterProfile string `pulumi:"sctpFilterProfile"`
	// When enabled, services match against any service EXCEPT the specified destination services.
	ServiceNegate string `pulumi:"serviceNegate"`
	// Name of service objects. The structure of `service` block is documented below.
	Services []GetProxypolicyService `pulumi:"services"`
	// TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
	SessionTtl int `pulumi:"sessionTtl"`
	// Name of an existing Spam filter profile.
	SpamfilterProfile string `pulumi:"spamfilterProfile"`
	// IPv6 source address objects. The structure of `srcaddr6` block is documented below.
	Srcaddr6s []GetProxypolicySrcaddr6 `pulumi:"srcaddr6s"`
	// When enabled, source addresses match against any address EXCEPT the specified source addresses.
	SrcaddrNegate string `pulumi:"srcaddrNegate"`
	// Source address objects. The structure of `srcaddr` block is documented below.
	Srcaddrs []GetProxypolicySrcaddr `pulumi:"srcaddrs"`
	// Source interface names. The structure of `srcintf` block is documented below.
	Srcintfs []GetProxypolicySrcintf `pulumi:"srcintfs"`
	// Name of an existing SSH filter profile.
	SshFilterProfile string `pulumi:"sshFilterProfile"`
	// Redirect SSH traffic to matching transparent proxy policy.
	SshPolicyRedirect string `pulumi:"sshPolicyRedirect"`
	// Name of an existing SSL SSH profile.
	SslSshProfile string `pulumi:"sslSshProfile"`
	// Enable/disable the active status of the policy.
	Status string `pulumi:"status"`
	// Enable to use the IP address of the client to connect to the server.
	Transparent string `pulumi:"transparent"`
	// Names of user objects. The structure of `users` block is documented below.
	Users []GetProxypolicyUser `pulumi:"users"`
	// Enable the use of UTM profiles/sensors/lists.
	UtmStatus string `pulumi:"utmStatus"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid      string  `pulumi:"uuid"`
	Vdomparam *string `pulumi:"vdomparam"`
	// Name of an existing VideoFilter profile.
	VideofilterProfile string `pulumi:"videofilterProfile"`
	// Name of an existing VoIP profile.
	VoipProfile string `pulumi:"voipProfile"`
	// Name of an existing Web application firewall profile.
	WafProfile string `pulumi:"wafProfile"`
	// Enable/disable web caching.
	Webcache string `pulumi:"webcache"`
	// Enable/disable web caching for HTTPS (Requires deep-inspection enabled in ssl-ssh-profile).
	WebcacheHttps string `pulumi:"webcacheHttps"`
	// Name of an existing Web filter profile.
	WebfilterProfile string `pulumi:"webfilterProfile"`
	// Web proxy forward server name.
	WebproxyForwardServer string `pulumi:"webproxyForwardServer"`
	// Name of web proxy profile.
	WebproxyProfile string `pulumi:"webproxyProfile"`
	// ZTNA EMS Tag names. The structure of `ztnaEmsTag` block is documented below.
	ZtnaEmsTags []GetProxypolicyZtnaEmsTag `pulumi:"ztnaEmsTags"`
	// ZTNA tag matching logic.
	ZtnaTagsMatchLogic string `pulumi:"ztnaTagsMatchLogic"`
}

func LookupProxypolicyOutput(ctx *pulumi.Context, args LookupProxypolicyOutputArgs, opts ...pulumi.InvokeOption) LookupProxypolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProxypolicyResult, error) {
			args := v.(LookupProxypolicyArgs)
			r, err := LookupProxypolicy(ctx, &args, opts...)
			var s LookupProxypolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProxypolicyResultOutput)
}

// A collection of arguments for invoking getProxypolicy.
type LookupProxypolicyOutputArgs struct {
	// Specify the policyid of the desired firewall proxypolicy.
	Policyid pulumi.IntInput `pulumi:"policyid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupProxypolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProxypolicyArgs)(nil)).Elem()
}

// A collection of values returned by getProxypolicy.
type LookupProxypolicyResultOutput struct{ *pulumi.OutputState }

func (LookupProxypolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProxypolicyResult)(nil)).Elem()
}

func (o LookupProxypolicyResultOutput) ToLookupProxypolicyResultOutput() LookupProxypolicyResultOutput {
	return o
}

func (o LookupProxypolicyResultOutput) ToLookupProxypolicyResultOutputWithContext(ctx context.Context) LookupProxypolicyResultOutput {
	return o
}

// IPv4 access proxy. The structure of `accessProxy` block is documented below.
func (o LookupProxypolicyResultOutput) AccessProxies() GetProxypolicyAccessProxyArrayOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) []GetProxypolicyAccessProxy { return v.AccessProxies }).(GetProxypolicyAccessProxyArrayOutput)
}

// IPv6 access proxy. The structure of `accessProxy6` block is documented below.
func (o LookupProxypolicyResultOutput) AccessProxy6s() GetProxypolicyAccessProxy6ArrayOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) []GetProxypolicyAccessProxy6 { return v.AccessProxy6s }).(GetProxypolicyAccessProxy6ArrayOutput)
}

// Accept or deny traffic matching the policy parameters.
func (o LookupProxypolicyResultOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.Action }).(pulumi.StringOutput)
}

// Name of an existing Application list.
func (o LookupProxypolicyResultOutput) ApplicationList() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.ApplicationList }).(pulumi.StringOutput)
}

// Name of an existing Antivirus profile.
func (o LookupProxypolicyResultOutput) AvProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.AvProfile }).(pulumi.StringOutput)
}

// Enable/disable block notification.
func (o LookupProxypolicyResultOutput) BlockNotification() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.BlockNotification }).(pulumi.StringOutput)
}

// Name of an existing CIFS profile.
func (o LookupProxypolicyResultOutput) CifsProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.CifsProfile }).(pulumi.StringOutput)
}

// Optional comments.
func (o LookupProxypolicyResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.Comments }).(pulumi.StringOutput)
}

// Decrypted traffic mirror.
func (o LookupProxypolicyResultOutput) DecryptedTrafficMirror() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.DecryptedTrafficMirror }).(pulumi.StringOutput)
}

// When enabled, the ownership enforcement will be done at policy level.
func (o LookupProxypolicyResultOutput) DeviceOwnership() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.DeviceOwnership }).(pulumi.StringOutput)
}

// Web proxy disclaimer setting: by domain, policy, or user.
func (o LookupProxypolicyResultOutput) Disclaimer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.Disclaimer }).(pulumi.StringOutput)
}

// Name of an existing DLP profile.
func (o LookupProxypolicyResultOutput) DlpProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.DlpProfile }).(pulumi.StringOutput)
}

// Name of an existing DLP sensor.
func (o LookupProxypolicyResultOutput) DlpSensor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.DlpSensor }).(pulumi.StringOutput)
}

// IPv6 destination address objects. The structure of `dstaddr6` block is documented below.
func (o LookupProxypolicyResultOutput) Dstaddr6s() GetProxypolicyDstaddr6ArrayOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) []GetProxypolicyDstaddr6 { return v.Dstaddr6s }).(GetProxypolicyDstaddr6ArrayOutput)
}

// When enabled, destination addresses match against any address EXCEPT the specified destination addresses.
func (o LookupProxypolicyResultOutput) DstaddrNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.DstaddrNegate }).(pulumi.StringOutput)
}

// Destination address objects. The structure of `dstaddr` block is documented below.
func (o LookupProxypolicyResultOutput) Dstaddrs() GetProxypolicyDstaddrArrayOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) []GetProxypolicyDstaddr { return v.Dstaddrs }).(GetProxypolicyDstaddrArrayOutput)
}

// Destination interface names. The structure of `dstintf` block is documented below.
func (o LookupProxypolicyResultOutput) Dstintfs() GetProxypolicyDstintfArrayOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) []GetProxypolicyDstintf { return v.Dstintfs }).(GetProxypolicyDstintfArrayOutput)
}

// Name of an existing email filter profile.
func (o LookupProxypolicyResultOutput) EmailfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.EmailfilterProfile }).(pulumi.StringOutput)
}

// Name of an existing file-filter profile.
func (o LookupProxypolicyResultOutput) FileFilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.FileFilterProfile }).(pulumi.StringOutput)
}

// Global web-based manager visible label.
func (o LookupProxypolicyResultOutput) GlobalLabel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.GlobalLabel }).(pulumi.StringOutput)
}

// Names of group objects. The structure of `groups` block is documented below.
func (o LookupProxypolicyResultOutput) Groups() GetProxypolicyGroupArrayOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) []GetProxypolicyGroup { return v.Groups }).(GetProxypolicyGroupArrayOutput)
}

// Enable/disable HTTP tunnel authentication.
func (o LookupProxypolicyResultOutput) HttpTunnelAuth() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.HttpTunnelAuth }).(pulumi.StringOutput)
}

// Name of an existing ICAP profile.
func (o LookupProxypolicyResultOutput) IcapProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.IcapProfile }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProxypolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.
func (o LookupProxypolicyResultOutput) InternetService() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.InternetService }).(pulumi.StringOutput)
}

// Custom Internet Service group name. The structure of `internetServiceCustomGroup` block is documented below.
func (o LookupProxypolicyResultOutput) InternetServiceCustomGroups() GetProxypolicyInternetServiceCustomGroupArrayOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) []GetProxypolicyInternetServiceCustomGroup {
		return v.InternetServiceCustomGroups
	}).(GetProxypolicyInternetServiceCustomGroupArrayOutput)
}

// Custom Internet Service name. The structure of `internetServiceCustom` block is documented below.
func (o LookupProxypolicyResultOutput) InternetServiceCustoms() GetProxypolicyInternetServiceCustomArrayOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) []GetProxypolicyInternetServiceCustom { return v.InternetServiceCustoms }).(GetProxypolicyInternetServiceCustomArrayOutput)
}

// Internet Service group name. The structure of `internetServiceGroup` block is documented below.
func (o LookupProxypolicyResultOutput) InternetServiceGroups() GetProxypolicyInternetServiceGroupArrayOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) []GetProxypolicyInternetServiceGroup { return v.InternetServiceGroups }).(GetProxypolicyInternetServiceGroupArrayOutput)
}

// Internet Service ID. The structure of `internetServiceId` block is documented below.
func (o LookupProxypolicyResultOutput) InternetServiceIds() GetProxypolicyInternetServiceIdArrayOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) []GetProxypolicyInternetServiceId { return v.InternetServiceIds }).(GetProxypolicyInternetServiceIdArrayOutput)
}

// Internet Service name. The structure of `internetServiceName` block is documented below.
func (o LookupProxypolicyResultOutput) InternetServiceNames() GetProxypolicyInternetServiceNameArrayOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) []GetProxypolicyInternetServiceName { return v.InternetServiceNames }).(GetProxypolicyInternetServiceNameArrayOutput)
}

// When enabled, Internet Services match against any internet service EXCEPT the selected Internet Service.
func (o LookupProxypolicyResultOutput) InternetServiceNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.InternetServiceNegate }).(pulumi.StringOutput)
}

// Name of an existing IPS sensor.
func (o LookupProxypolicyResultOutput) IpsSensor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.IpsSensor }).(pulumi.StringOutput)
}

// VDOM-specific GUI visible label.
func (o LookupProxypolicyResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.Label }).(pulumi.StringOutput)
}

// Enable/disable logging traffic through the policy.
func (o LookupProxypolicyResultOutput) Logtraffic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.Logtraffic }).(pulumi.StringOutput)
}

// Enable/disable policy log traffic start.
func (o LookupProxypolicyResultOutput) LogtrafficStart() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.LogtrafficStart }).(pulumi.StringOutput)
}

// Group name.
func (o LookupProxypolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Policy ID.
func (o LookupProxypolicyResultOutput) Policyid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) int { return v.Policyid }).(pulumi.IntOutput)
}

// Name of IP pool object. The structure of `poolname` block is documented below.
func (o LookupProxypolicyResultOutput) Poolnames() GetProxypolicyPoolnameArrayOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) []GetProxypolicyPoolname { return v.Poolnames }).(GetProxypolicyPoolnameArrayOutput)
}

// Name of profile group.
func (o LookupProxypolicyResultOutput) ProfileGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.ProfileGroup }).(pulumi.StringOutput)
}

// Name of an existing Protocol options profile.
func (o LookupProxypolicyResultOutput) ProfileProtocolOptions() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.ProfileProtocolOptions }).(pulumi.StringOutput)
}

// Determine whether the firewall policy allows security profile groups or single profiles only.
func (o LookupProxypolicyResultOutput) ProfileType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.ProfileType }).(pulumi.StringOutput)
}

// Type of explicit proxy.
func (o LookupProxypolicyResultOutput) Proxy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.Proxy }).(pulumi.StringOutput)
}

// Redirect URL for further explicit web proxy processing.
func (o LookupProxypolicyResultOutput) RedirectUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.RedirectUrl }).(pulumi.StringOutput)
}

// Authentication replacement message override group.
func (o LookupProxypolicyResultOutput) ReplacemsgOverrideGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.ReplacemsgOverrideGroup }).(pulumi.StringOutput)
}

// Enable/disable scanning of connections to Botnet servers.
func (o LookupProxypolicyResultOutput) ScanBotnetConnections() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.ScanBotnetConnections }).(pulumi.StringOutput)
}

// Name of schedule object.
func (o LookupProxypolicyResultOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.Schedule }).(pulumi.StringOutput)
}

// Name of an existing SCTP filter profile.
func (o LookupProxypolicyResultOutput) SctpFilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.SctpFilterProfile }).(pulumi.StringOutput)
}

// When enabled, services match against any service EXCEPT the specified destination services.
func (o LookupProxypolicyResultOutput) ServiceNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.ServiceNegate }).(pulumi.StringOutput)
}

// Name of service objects. The structure of `service` block is documented below.
func (o LookupProxypolicyResultOutput) Services() GetProxypolicyServiceArrayOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) []GetProxypolicyService { return v.Services }).(GetProxypolicyServiceArrayOutput)
}

// TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
func (o LookupProxypolicyResultOutput) SessionTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) int { return v.SessionTtl }).(pulumi.IntOutput)
}

// Name of an existing Spam filter profile.
func (o LookupProxypolicyResultOutput) SpamfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.SpamfilterProfile }).(pulumi.StringOutput)
}

// IPv6 source address objects. The structure of `srcaddr6` block is documented below.
func (o LookupProxypolicyResultOutput) Srcaddr6s() GetProxypolicySrcaddr6ArrayOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) []GetProxypolicySrcaddr6 { return v.Srcaddr6s }).(GetProxypolicySrcaddr6ArrayOutput)
}

// When enabled, source addresses match against any address EXCEPT the specified source addresses.
func (o LookupProxypolicyResultOutput) SrcaddrNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.SrcaddrNegate }).(pulumi.StringOutput)
}

// Source address objects. The structure of `srcaddr` block is documented below.
func (o LookupProxypolicyResultOutput) Srcaddrs() GetProxypolicySrcaddrArrayOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) []GetProxypolicySrcaddr { return v.Srcaddrs }).(GetProxypolicySrcaddrArrayOutput)
}

// Source interface names. The structure of `srcintf` block is documented below.
func (o LookupProxypolicyResultOutput) Srcintfs() GetProxypolicySrcintfArrayOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) []GetProxypolicySrcintf { return v.Srcintfs }).(GetProxypolicySrcintfArrayOutput)
}

// Name of an existing SSH filter profile.
func (o LookupProxypolicyResultOutput) SshFilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.SshFilterProfile }).(pulumi.StringOutput)
}

// Redirect SSH traffic to matching transparent proxy policy.
func (o LookupProxypolicyResultOutput) SshPolicyRedirect() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.SshPolicyRedirect }).(pulumi.StringOutput)
}

// Name of an existing SSL SSH profile.
func (o LookupProxypolicyResultOutput) SslSshProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.SslSshProfile }).(pulumi.StringOutput)
}

// Enable/disable the active status of the policy.
func (o LookupProxypolicyResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.Status }).(pulumi.StringOutput)
}

// Enable to use the IP address of the client to connect to the server.
func (o LookupProxypolicyResultOutput) Transparent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.Transparent }).(pulumi.StringOutput)
}

// Names of user objects. The structure of `users` block is documented below.
func (o LookupProxypolicyResultOutput) Users() GetProxypolicyUserArrayOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) []GetProxypolicyUser { return v.Users }).(GetProxypolicyUserArrayOutput)
}

// Enable the use of UTM profiles/sensors/lists.
func (o LookupProxypolicyResultOutput) UtmStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.UtmStatus }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o LookupProxypolicyResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupProxypolicyResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Name of an existing VideoFilter profile.
func (o LookupProxypolicyResultOutput) VideofilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.VideofilterProfile }).(pulumi.StringOutput)
}

// Name of an existing VoIP profile.
func (o LookupProxypolicyResultOutput) VoipProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.VoipProfile }).(pulumi.StringOutput)
}

// Name of an existing Web application firewall profile.
func (o LookupProxypolicyResultOutput) WafProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.WafProfile }).(pulumi.StringOutput)
}

// Enable/disable web caching.
func (o LookupProxypolicyResultOutput) Webcache() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.Webcache }).(pulumi.StringOutput)
}

// Enable/disable web caching for HTTPS (Requires deep-inspection enabled in ssl-ssh-profile).
func (o LookupProxypolicyResultOutput) WebcacheHttps() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.WebcacheHttps }).(pulumi.StringOutput)
}

// Name of an existing Web filter profile.
func (o LookupProxypolicyResultOutput) WebfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.WebfilterProfile }).(pulumi.StringOutput)
}

// Web proxy forward server name.
func (o LookupProxypolicyResultOutput) WebproxyForwardServer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.WebproxyForwardServer }).(pulumi.StringOutput)
}

// Name of web proxy profile.
func (o LookupProxypolicyResultOutput) WebproxyProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.WebproxyProfile }).(pulumi.StringOutput)
}

// ZTNA EMS Tag names. The structure of `ztnaEmsTag` block is documented below.
func (o LookupProxypolicyResultOutput) ZtnaEmsTags() GetProxypolicyZtnaEmsTagArrayOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) []GetProxypolicyZtnaEmsTag { return v.ZtnaEmsTags }).(GetProxypolicyZtnaEmsTagArrayOutput)
}

// ZTNA tag matching logic.
func (o LookupProxypolicyResultOutput) ZtnaTagsMatchLogic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxypolicyResult) string { return v.ZtnaTagsMatchLogic }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProxypolicyResultOutput{})
}
