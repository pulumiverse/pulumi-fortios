// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on an fortios system dnsdatabase
func LookupDnsdatabase(ctx *pulumi.Context, args *LookupDnsdatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDnsdatabaseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDnsdatabaseResult
	err := ctx.Invoke("fortios:system/getDnsdatabase:getDnsdatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDnsdatabase.
type LookupDnsdatabaseArgs struct {
	// Specify the name of the desired system dnsdatabase.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getDnsdatabase.
type LookupDnsdatabaseResult struct {
	// DNS zone transfer IP address list.
	AllowTransfer string `pulumi:"allowTransfer"`
	// Enable/disable authoritative zone.
	Authoritative string `pulumi:"authoritative"`
	// Email address of the administrator for this zone.
	// You can specify only the username (e.g. admin) or full email address (e.g. admin@test.com)
	// When using a simple username, the domain of the email will be this zone.
	Contact string `pulumi:"contact"`
	// DNS entry. The structure of `dnsEntry` block is documented below.
	DnsEntries []GetDnsdatabaseDnsEntry `pulumi:"dnsEntries"`
	// Domain name.
	Domain string `pulumi:"domain"`
	// DNS zone forwarder IP address list.
	Forwarder string `pulumi:"forwarder"`
	// Forwarder IPv6 address.
	Forwarder6 string `pulumi:"forwarder6"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.
	IpMaster string `pulumi:"ipMaster"`
	// IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.
	IpPrimary string `pulumi:"ipPrimary"`
	// Zone name.
	Name string `pulumi:"name"`
	// Domain name of the default DNS server for this zone.
	PrimaryName string `pulumi:"primaryName"`
	// Maximum number of resource records (10 - 65536, 0 means infinite).
	RrMax int `pulumi:"rrMax"`
	// Source IP for forwarding to DNS server.
	SourceIp string `pulumi:"sourceIp"`
	// IPv6 source IP address for forwarding to DNS server.
	SourceIp6 string `pulumi:"sourceIp6"`
	// Enable/disable resource record status.
	Status string `pulumi:"status"`
	// Time-to-live for this entry (0 to 2147483647 sec, default = 0).
	Ttl int `pulumi:"ttl"`
	// Resource record type.
	Type      string  `pulumi:"type"`
	Vdomparam *string `pulumi:"vdomparam"`
	// Zone view (public to serve public clients, shadow to serve internal clients).
	View string `pulumi:"view"`
}

func LookupDnsdatabaseOutput(ctx *pulumi.Context, args LookupDnsdatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupDnsdatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDnsdatabaseResult, error) {
			args := v.(LookupDnsdatabaseArgs)
			r, err := LookupDnsdatabase(ctx, &args, opts...)
			var s LookupDnsdatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDnsdatabaseResultOutput)
}

// A collection of arguments for invoking getDnsdatabase.
type LookupDnsdatabaseOutputArgs struct {
	// Specify the name of the desired system dnsdatabase.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupDnsdatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsdatabaseArgs)(nil)).Elem()
}

// A collection of values returned by getDnsdatabase.
type LookupDnsdatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupDnsdatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsdatabaseResult)(nil)).Elem()
}

func (o LookupDnsdatabaseResultOutput) ToLookupDnsdatabaseResultOutput() LookupDnsdatabaseResultOutput {
	return o
}

func (o LookupDnsdatabaseResultOutput) ToLookupDnsdatabaseResultOutputWithContext(ctx context.Context) LookupDnsdatabaseResultOutput {
	return o
}

// DNS zone transfer IP address list.
func (o LookupDnsdatabaseResultOutput) AllowTransfer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) string { return v.AllowTransfer }).(pulumi.StringOutput)
}

// Enable/disable authoritative zone.
func (o LookupDnsdatabaseResultOutput) Authoritative() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) string { return v.Authoritative }).(pulumi.StringOutput)
}

// Email address of the administrator for this zone.
// You can specify only the username (e.g. admin) or full email address (e.g. admin@test.com)
// When using a simple username, the domain of the email will be this zone.
func (o LookupDnsdatabaseResultOutput) Contact() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) string { return v.Contact }).(pulumi.StringOutput)
}

// DNS entry. The structure of `dnsEntry` block is documented below.
func (o LookupDnsdatabaseResultOutput) DnsEntries() GetDnsdatabaseDnsEntryArrayOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) []GetDnsdatabaseDnsEntry { return v.DnsEntries }).(GetDnsdatabaseDnsEntryArrayOutput)
}

// Domain name.
func (o LookupDnsdatabaseResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) string { return v.Domain }).(pulumi.StringOutput)
}

// DNS zone forwarder IP address list.
func (o LookupDnsdatabaseResultOutput) Forwarder() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) string { return v.Forwarder }).(pulumi.StringOutput)
}

// Forwarder IPv6 address.
func (o LookupDnsdatabaseResultOutput) Forwarder6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) string { return v.Forwarder6 }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDnsdatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

// IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.
func (o LookupDnsdatabaseResultOutput) IpMaster() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) string { return v.IpMaster }).(pulumi.StringOutput)
}

// IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.
func (o LookupDnsdatabaseResultOutput) IpPrimary() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) string { return v.IpPrimary }).(pulumi.StringOutput)
}

// Zone name.
func (o LookupDnsdatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

// Domain name of the default DNS server for this zone.
func (o LookupDnsdatabaseResultOutput) PrimaryName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) string { return v.PrimaryName }).(pulumi.StringOutput)
}

// Maximum number of resource records (10 - 65536, 0 means infinite).
func (o LookupDnsdatabaseResultOutput) RrMax() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) int { return v.RrMax }).(pulumi.IntOutput)
}

// Source IP for forwarding to DNS server.
func (o LookupDnsdatabaseResultOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) string { return v.SourceIp }).(pulumi.StringOutput)
}

// IPv6 source IP address for forwarding to DNS server.
func (o LookupDnsdatabaseResultOutput) SourceIp6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) string { return v.SourceIp6 }).(pulumi.StringOutput)
}

// Enable/disable resource record status.
func (o LookupDnsdatabaseResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) string { return v.Status }).(pulumi.StringOutput)
}

// Time-to-live for this entry (0 to 2147483647 sec, default = 0).
func (o LookupDnsdatabaseResultOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) int { return v.Ttl }).(pulumi.IntOutput)
}

// Resource record type.
func (o LookupDnsdatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDnsdatabaseResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Zone view (public to serve public clients, shadow to serve internal clients).
func (o LookupDnsdatabaseResultOutput) View() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsdatabaseResult) string { return v.View }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDnsdatabaseResultOutput{})
}
