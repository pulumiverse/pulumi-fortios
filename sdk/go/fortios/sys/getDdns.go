// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system ddns
func LookupDdns(ctx *pulumi.Context, args *LookupDdnsArgs, opts ...pulumi.InvokeOption) (*LookupDdnsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupDdnsResult
	err := ctx.Invoke("fortios:sys/getDdns:getDdns", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDdns.
type LookupDdnsArgs struct {
	// Specify the ddnsid of the desired system ddns.
	Ddnsid int `pulumi:"ddnsid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getDdns.
type LookupDdnsResult struct {
	// Address type of interface address in DDNS update.
	AddrType string `pulumi:"addrType"`
	// Bound IP address.
	BoundIp string `pulumi:"boundIp"`
	// Enable/disable use of clear text connections.
	ClearText string `pulumi:"clearText"`
	// Enable/disable TSIG authentication for your DDNS server.
	DdnsAuth string `pulumi:"ddnsAuth"`
	// Your fully qualified domain name (for example, yourname.DDNS.com).
	DdnsDomain string `pulumi:"ddnsDomain"`
	// DDNS update key (base 64 encoding).
	DdnsKey string `pulumi:"ddnsKey"`
	// DDNS update key name.
	DdnsKeyname string `pulumi:"ddnsKeyname"`
	// DDNS password.
	DdnsPassword string `pulumi:"ddnsPassword"`
	// Select a DDNS service provider.
	DdnsServer string `pulumi:"ddnsServer"`
	// Generic DDNS server IP/FQDN list. The structure of `ddnsServerAddr` block is documented below.
	DdnsServerAddrs []GetDdnsDdnsServerAddr `pulumi:"ddnsServerAddrs"`
	// Generic DDNS server IP.
	DdnsServerIp string `pulumi:"ddnsServerIp"`
	// DDNS Serial Number.
	DdnsSn string `pulumi:"ddnsSn"`
	// Time-to-live for DDNS packets.
	DdnsTtl int `pulumi:"ddnsTtl"`
	// DDNS user name.
	DdnsUsername string `pulumi:"ddnsUsername"`
	// Zone of your domain name (for example, DDNS.com).
	DdnsZone string `pulumi:"ddnsZone"`
	// DDNS ID.
	Ddnsid int `pulumi:"ddnsid"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Monitored interface. The structure of `monitorInterface` block is documented below.
	MonitorInterfaces []GetDdnsMonitorInterface `pulumi:"monitorInterfaces"`
	// Address type of the DDNS server.
	ServerType string `pulumi:"serverType"`
	// Name of local certificate for SSL connections.
	SslCertificate string `pulumi:"sslCertificate"`
	// DDNS update interval (60 - 2592000 sec, default = 300).
	UpdateInterval int `pulumi:"updateInterval"`
	// Enable/disable use of public IP address.
	UsePublicIp string  `pulumi:"usePublicIp"`
	Vdomparam   *string `pulumi:"vdomparam"`
}

func LookupDdnsOutput(ctx *pulumi.Context, args LookupDdnsOutputArgs, opts ...pulumi.InvokeOption) LookupDdnsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDdnsResult, error) {
			args := v.(LookupDdnsArgs)
			r, err := LookupDdns(ctx, &args, opts...)
			var s LookupDdnsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDdnsResultOutput)
}

// A collection of arguments for invoking getDdns.
type LookupDdnsOutputArgs struct {
	// Specify the ddnsid of the desired system ddns.
	Ddnsid pulumi.IntInput `pulumi:"ddnsid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupDdnsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDdnsArgs)(nil)).Elem()
}

// A collection of values returned by getDdns.
type LookupDdnsResultOutput struct{ *pulumi.OutputState }

func (LookupDdnsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDdnsResult)(nil)).Elem()
}

func (o LookupDdnsResultOutput) ToLookupDdnsResultOutput() LookupDdnsResultOutput {
	return o
}

func (o LookupDdnsResultOutput) ToLookupDdnsResultOutputWithContext(ctx context.Context) LookupDdnsResultOutput {
	return o
}

// Address type of interface address in DDNS update.
func (o LookupDdnsResultOutput) AddrType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdnsResult) string { return v.AddrType }).(pulumi.StringOutput)
}

// Bound IP address.
func (o LookupDdnsResultOutput) BoundIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdnsResult) string { return v.BoundIp }).(pulumi.StringOutput)
}

// Enable/disable use of clear text connections.
func (o LookupDdnsResultOutput) ClearText() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdnsResult) string { return v.ClearText }).(pulumi.StringOutput)
}

// Enable/disable TSIG authentication for your DDNS server.
func (o LookupDdnsResultOutput) DdnsAuth() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdnsResult) string { return v.DdnsAuth }).(pulumi.StringOutput)
}

// Your fully qualified domain name (for example, yourname.DDNS.com).
func (o LookupDdnsResultOutput) DdnsDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdnsResult) string { return v.DdnsDomain }).(pulumi.StringOutput)
}

// DDNS update key (base 64 encoding).
func (o LookupDdnsResultOutput) DdnsKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdnsResult) string { return v.DdnsKey }).(pulumi.StringOutput)
}

// DDNS update key name.
func (o LookupDdnsResultOutput) DdnsKeyname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdnsResult) string { return v.DdnsKeyname }).(pulumi.StringOutput)
}

// DDNS password.
func (o LookupDdnsResultOutput) DdnsPassword() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdnsResult) string { return v.DdnsPassword }).(pulumi.StringOutput)
}

// Select a DDNS service provider.
func (o LookupDdnsResultOutput) DdnsServer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdnsResult) string { return v.DdnsServer }).(pulumi.StringOutput)
}

// Generic DDNS server IP/FQDN list. The structure of `ddnsServerAddr` block is documented below.
func (o LookupDdnsResultOutput) DdnsServerAddrs() GetDdnsDdnsServerAddrArrayOutput {
	return o.ApplyT(func(v LookupDdnsResult) []GetDdnsDdnsServerAddr { return v.DdnsServerAddrs }).(GetDdnsDdnsServerAddrArrayOutput)
}

// Generic DDNS server IP.
func (o LookupDdnsResultOutput) DdnsServerIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdnsResult) string { return v.DdnsServerIp }).(pulumi.StringOutput)
}

// DDNS Serial Number.
func (o LookupDdnsResultOutput) DdnsSn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdnsResult) string { return v.DdnsSn }).(pulumi.StringOutput)
}

// Time-to-live for DDNS packets.
func (o LookupDdnsResultOutput) DdnsTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDdnsResult) int { return v.DdnsTtl }).(pulumi.IntOutput)
}

// DDNS user name.
func (o LookupDdnsResultOutput) DdnsUsername() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdnsResult) string { return v.DdnsUsername }).(pulumi.StringOutput)
}

// Zone of your domain name (for example, DDNS.com).
func (o LookupDdnsResultOutput) DdnsZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdnsResult) string { return v.DdnsZone }).(pulumi.StringOutput)
}

// DDNS ID.
func (o LookupDdnsResultOutput) Ddnsid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDdnsResult) int { return v.Ddnsid }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDdnsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdnsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Monitored interface. The structure of `monitorInterface` block is documented below.
func (o LookupDdnsResultOutput) MonitorInterfaces() GetDdnsMonitorInterfaceArrayOutput {
	return o.ApplyT(func(v LookupDdnsResult) []GetDdnsMonitorInterface { return v.MonitorInterfaces }).(GetDdnsMonitorInterfaceArrayOutput)
}

// Address type of the DDNS server.
func (o LookupDdnsResultOutput) ServerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdnsResult) string { return v.ServerType }).(pulumi.StringOutput)
}

// Name of local certificate for SSL connections.
func (o LookupDdnsResultOutput) SslCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdnsResult) string { return v.SslCertificate }).(pulumi.StringOutput)
}

// DDNS update interval (60 - 2592000 sec, default = 300).
func (o LookupDdnsResultOutput) UpdateInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDdnsResult) int { return v.UpdateInterval }).(pulumi.IntOutput)
}

// Enable/disable use of public IP address.
func (o LookupDdnsResultOutput) UsePublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdnsResult) string { return v.UsePublicIp }).(pulumi.StringOutput)
}

func (o LookupDdnsResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDdnsResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDdnsResultOutput{})
}
