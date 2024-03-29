// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on fortios system dns
func LookupDns(ctx *pulumi.Context, args *LookupDnsArgs, opts ...pulumi.InvokeOption) (*LookupDnsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDnsResult
	err := ctx.Invoke("fortios:system/getDns:getDns", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDns.
type LookupDnsArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getDns.
type LookupDnsResult struct {
	// Alternate primary DNS server. (This is not used as a failover DNS server.)
	AltPrimary string `pulumi:"altPrimary"`
	// Alternate secondary DNS server. (This is not used as a failover DNS server.)
	AltSecondary string `pulumi:"altSecondary"`
	// Enable/disable response from the DNS server when a record is not in cache.
	CacheNotfoundResponses string `pulumi:"cacheNotfoundResponses"`
	// Maximum number of records in the DNS cache.
	DnsCacheLimit int `pulumi:"dnsCacheLimit"`
	// Duration in seconds that the DNS cache retains information.
	DnsCacheTtl int `pulumi:"dnsCacheTtl"`
	// Enable/disable/enforce DNS over TLS.
	DnsOverTls string `pulumi:"dnsOverTls"`
	// DNS search domain list separated by space (maximum 8 domains)
	Domains []GetDnsDomain `pulumi:"domains"`
	// FQDN cache time to live in seconds (0 - 86400, default = 0).
	FqdnCacheTtl int `pulumi:"fqdnCacheTtl"`
	// FQDN cache maximum refresh time in seconds (3600 - 86400, default = 3600).
	FqdnMaxRefresh int `pulumi:"fqdnMaxRefresh"`
	// FQDN cache minimum refresh time in seconds (10 - 3600, default = 60).
	FqdnMinRefresh int `pulumi:"fqdnMinRefresh"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Specify outgoing interface to reach server.
	Interface string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server.
	InterfaceSelectMethod string `pulumi:"interfaceSelectMethod"`
	// Primary DNS server IPv6 address.
	Ip6Primary string `pulumi:"ip6Primary"`
	// Secondary DNS server IPv6 address.
	Ip6Secondary string `pulumi:"ip6Secondary"`
	// Local DNS log setting.
	Log string `pulumi:"log"`
	// Primary DNS server IP address.
	Primary string `pulumi:"primary"`
	// DNS protocols.
	Protocol string `pulumi:"protocol"`
	// Number of times to retry (0 - 5).
	Retry int `pulumi:"retry"`
	// Secondary DNS server IP address.
	Secondary string `pulumi:"secondary"`
	// DNS server host name list. The structure of `serverHostname` block is documented below.
	ServerHostnames []GetDnsServerHostname `pulumi:"serverHostnames"`
	// Specify how configured servers are prioritized.
	ServerSelectMethod string `pulumi:"serverSelectMethod"`
	// IP address used by the DNS server as its source IP.
	SourceIp string `pulumi:"sourceIp"`
	// Name of local certificate for SSL connections.
	SslCertificate string `pulumi:"sslCertificate"`
	// DNS query timeout interval in seconds (1 - 10).
	Timeout   int     `pulumi:"timeout"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupDnsOutput(ctx *pulumi.Context, args LookupDnsOutputArgs, opts ...pulumi.InvokeOption) LookupDnsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDnsResult, error) {
			args := v.(LookupDnsArgs)
			r, err := LookupDns(ctx, &args, opts...)
			var s LookupDnsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDnsResultOutput)
}

// A collection of arguments for invoking getDns.
type LookupDnsOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupDnsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsArgs)(nil)).Elem()
}

// A collection of values returned by getDns.
type LookupDnsResultOutput struct{ *pulumi.OutputState }

func (LookupDnsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsResult)(nil)).Elem()
}

func (o LookupDnsResultOutput) ToLookupDnsResultOutput() LookupDnsResultOutput {
	return o
}

func (o LookupDnsResultOutput) ToLookupDnsResultOutputWithContext(ctx context.Context) LookupDnsResultOutput {
	return o
}

// Alternate primary DNS server. (This is not used as a failover DNS server.)
func (o LookupDnsResultOutput) AltPrimary() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResult) string { return v.AltPrimary }).(pulumi.StringOutput)
}

// Alternate secondary DNS server. (This is not used as a failover DNS server.)
func (o LookupDnsResultOutput) AltSecondary() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResult) string { return v.AltSecondary }).(pulumi.StringOutput)
}

// Enable/disable response from the DNS server when a record is not in cache.
func (o LookupDnsResultOutput) CacheNotfoundResponses() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResult) string { return v.CacheNotfoundResponses }).(pulumi.StringOutput)
}

// Maximum number of records in the DNS cache.
func (o LookupDnsResultOutput) DnsCacheLimit() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDnsResult) int { return v.DnsCacheLimit }).(pulumi.IntOutput)
}

// Duration in seconds that the DNS cache retains information.
func (o LookupDnsResultOutput) DnsCacheTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDnsResult) int { return v.DnsCacheTtl }).(pulumi.IntOutput)
}

// Enable/disable/enforce DNS over TLS.
func (o LookupDnsResultOutput) DnsOverTls() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResult) string { return v.DnsOverTls }).(pulumi.StringOutput)
}

// DNS search domain list separated by space (maximum 8 domains)
func (o LookupDnsResultOutput) Domains() GetDnsDomainArrayOutput {
	return o.ApplyT(func(v LookupDnsResult) []GetDnsDomain { return v.Domains }).(GetDnsDomainArrayOutput)
}

// FQDN cache time to live in seconds (0 - 86400, default = 0).
func (o LookupDnsResultOutput) FqdnCacheTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDnsResult) int { return v.FqdnCacheTtl }).(pulumi.IntOutput)
}

// FQDN cache maximum refresh time in seconds (3600 - 86400, default = 3600).
func (o LookupDnsResultOutput) FqdnMaxRefresh() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDnsResult) int { return v.FqdnMaxRefresh }).(pulumi.IntOutput)
}

// FQDN cache minimum refresh time in seconds (10 - 3600, default = 60).
func (o LookupDnsResultOutput) FqdnMinRefresh() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDnsResult) int { return v.FqdnMinRefresh }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDnsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specify outgoing interface to reach server.
func (o LookupDnsResultOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResult) string { return v.Interface }).(pulumi.StringOutput)
}

// Specify how to select outgoing interface to reach server.
func (o LookupDnsResultOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResult) string { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

// Primary DNS server IPv6 address.
func (o LookupDnsResultOutput) Ip6Primary() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResult) string { return v.Ip6Primary }).(pulumi.StringOutput)
}

// Secondary DNS server IPv6 address.
func (o LookupDnsResultOutput) Ip6Secondary() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResult) string { return v.Ip6Secondary }).(pulumi.StringOutput)
}

// Local DNS log setting.
func (o LookupDnsResultOutput) Log() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResult) string { return v.Log }).(pulumi.StringOutput)
}

// Primary DNS server IP address.
func (o LookupDnsResultOutput) Primary() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResult) string { return v.Primary }).(pulumi.StringOutput)
}

// DNS protocols.
func (o LookupDnsResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResult) string { return v.Protocol }).(pulumi.StringOutput)
}

// Number of times to retry (0 - 5).
func (o LookupDnsResultOutput) Retry() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDnsResult) int { return v.Retry }).(pulumi.IntOutput)
}

// Secondary DNS server IP address.
func (o LookupDnsResultOutput) Secondary() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResult) string { return v.Secondary }).(pulumi.StringOutput)
}

// DNS server host name list. The structure of `serverHostname` block is documented below.
func (o LookupDnsResultOutput) ServerHostnames() GetDnsServerHostnameArrayOutput {
	return o.ApplyT(func(v LookupDnsResult) []GetDnsServerHostname { return v.ServerHostnames }).(GetDnsServerHostnameArrayOutput)
}

// Specify how configured servers are prioritized.
func (o LookupDnsResultOutput) ServerSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResult) string { return v.ServerSelectMethod }).(pulumi.StringOutput)
}

// IP address used by the DNS server as its source IP.
func (o LookupDnsResultOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResult) string { return v.SourceIp }).(pulumi.StringOutput)
}

// Name of local certificate for SSL connections.
func (o LookupDnsResultOutput) SslCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResult) string { return v.SslCertificate }).(pulumi.StringOutput)
}

// DNS query timeout interval in seconds (1 - 10).
func (o LookupDnsResultOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDnsResult) int { return v.Timeout }).(pulumi.IntOutput)
}

func (o LookupDnsResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDnsResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDnsResultOutput{})
}
