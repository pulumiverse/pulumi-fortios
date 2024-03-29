// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on an fortios system dnsserver
func LookupDnsserver(ctx *pulumi.Context, args *LookupDnsserverArgs, opts ...pulumi.InvokeOption) (*LookupDnsserverResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDnsserverResult
	err := ctx.Invoke("fortios:system/getDnsserver:getDnsserver", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDnsserver.
type LookupDnsserverArgs struct {
	// Specify the name of the desired system dnsserver.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getDnsserver.
type LookupDnsserverResult struct {
	// DNS filter profile.
	DnsfilterProfile string `pulumi:"dnsfilterProfile"`
	// DNS over HTTPS.
	Doh string `pulumi:"doh"`
	// Enable/disable DNS over QUIC/HTTP3/443 (default = disable).
	Doh3 string `pulumi:"doh3"`
	// Enable/disable DNS over QUIC/853 (default = disable).
	Doq string `pulumi:"doq"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// DNS server mode.
	Mode string `pulumi:"mode"`
	// DNS server name.
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupDnsserverOutput(ctx *pulumi.Context, args LookupDnsserverOutputArgs, opts ...pulumi.InvokeOption) LookupDnsserverResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDnsserverResult, error) {
			args := v.(LookupDnsserverArgs)
			r, err := LookupDnsserver(ctx, &args, opts...)
			var s LookupDnsserverResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDnsserverResultOutput)
}

// A collection of arguments for invoking getDnsserver.
type LookupDnsserverOutputArgs struct {
	// Specify the name of the desired system dnsserver.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupDnsserverOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsserverArgs)(nil)).Elem()
}

// A collection of values returned by getDnsserver.
type LookupDnsserverResultOutput struct{ *pulumi.OutputState }

func (LookupDnsserverResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsserverResult)(nil)).Elem()
}

func (o LookupDnsserverResultOutput) ToLookupDnsserverResultOutput() LookupDnsserverResultOutput {
	return o
}

func (o LookupDnsserverResultOutput) ToLookupDnsserverResultOutputWithContext(ctx context.Context) LookupDnsserverResultOutput {
	return o
}

// DNS filter profile.
func (o LookupDnsserverResultOutput) DnsfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsserverResult) string { return v.DnsfilterProfile }).(pulumi.StringOutput)
}

// DNS over HTTPS.
func (o LookupDnsserverResultOutput) Doh() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsserverResult) string { return v.Doh }).(pulumi.StringOutput)
}

// Enable/disable DNS over QUIC/HTTP3/443 (default = disable).
func (o LookupDnsserverResultOutput) Doh3() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsserverResult) string { return v.Doh3 }).(pulumi.StringOutput)
}

// Enable/disable DNS over QUIC/853 (default = disable).
func (o LookupDnsserverResultOutput) Doq() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsserverResult) string { return v.Doq }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDnsserverResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsserverResult) string { return v.Id }).(pulumi.StringOutput)
}

// DNS server mode.
func (o LookupDnsserverResultOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsserverResult) string { return v.Mode }).(pulumi.StringOutput)
}

// DNS server name.
func (o LookupDnsserverResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsserverResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDnsserverResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDnsserverResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDnsserverResultOutput{})
}
