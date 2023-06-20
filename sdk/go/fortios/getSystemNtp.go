// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system ntp
func LookupSystemNtp(ctx *pulumi.Context, args *LookupSystemNtpArgs, opts ...pulumi.InvokeOption) (*LookupSystemNtpResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemNtpResult
	err := ctx.Invoke("fortios:index/getSystemNtp:getSystemNtp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemNtp.
type LookupSystemNtpArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemNtp.
type LookupSystemNtpResult struct {
	// Enable/disable MD5/SHA1 authentication.
	Authentication string `pulumi:"authentication"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Specify outgoing interface to reach server.
	Interfaces []GetSystemNtpInterface `pulumi:"interfaces"`
	// Key for MD5/SHA1 authentication.
	Key string `pulumi:"key"`
	// Key ID for authentication.
	KeyId int `pulumi:"keyId"`
	// Key type for authentication (MD5, SHA1).
	KeyType string `pulumi:"keyType"`
	// Configure the FortiGate to connect to any available third-party NTP server. The structure of `ntpserver` block is documented below.
	Ntpservers []GetSystemNtpNtpserver `pulumi:"ntpservers"`
	// Enable/disable setting the FortiGate system time by synchronizing with an NTP Server.
	Ntpsync string `pulumi:"ntpsync"`
	// Enable/disable FortiGate NTP Server Mode. Your FortiGate becomes an NTP server for other devices on your network. The FortiGate relays NTP requests to its configured NTP server.
	ServerMode string `pulumi:"serverMode"`
	// Source IP address for communication to the NTP server.
	SourceIp string `pulumi:"sourceIp"`
	// Source IPv6 address for communication to the NTP server.
	SourceIp6 string `pulumi:"sourceIp6"`
	// NTP synchronization interval (1 - 1440 min).
	Syncinterval int `pulumi:"syncinterval"`
	// Use the FortiGuard NTP server or any other available NTP Server.
	Type      string  `pulumi:"type"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSystemNtpOutput(ctx *pulumi.Context, args LookupSystemNtpOutputArgs, opts ...pulumi.InvokeOption) LookupSystemNtpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemNtpResult, error) {
			args := v.(LookupSystemNtpArgs)
			r, err := LookupSystemNtp(ctx, &args, opts...)
			var s LookupSystemNtpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemNtpResultOutput)
}

// A collection of arguments for invoking getSystemNtp.
type LookupSystemNtpOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemNtpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemNtpArgs)(nil)).Elem()
}

// A collection of values returned by getSystemNtp.
type LookupSystemNtpResultOutput struct{ *pulumi.OutputState }

func (LookupSystemNtpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemNtpResult)(nil)).Elem()
}

func (o LookupSystemNtpResultOutput) ToLookupSystemNtpResultOutput() LookupSystemNtpResultOutput {
	return o
}

func (o LookupSystemNtpResultOutput) ToLookupSystemNtpResultOutputWithContext(ctx context.Context) LookupSystemNtpResultOutput {
	return o
}

// Enable/disable MD5/SHA1 authentication.
func (o LookupSystemNtpResultOutput) Authentication() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNtpResult) string { return v.Authentication }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemNtpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNtpResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specify outgoing interface to reach server.
func (o LookupSystemNtpResultOutput) Interfaces() GetSystemNtpInterfaceArrayOutput {
	return o.ApplyT(func(v LookupSystemNtpResult) []GetSystemNtpInterface { return v.Interfaces }).(GetSystemNtpInterfaceArrayOutput)
}

// Key for MD5/SHA1 authentication.
func (o LookupSystemNtpResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNtpResult) string { return v.Key }).(pulumi.StringOutput)
}

// Key ID for authentication.
func (o LookupSystemNtpResultOutput) KeyId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemNtpResult) int { return v.KeyId }).(pulumi.IntOutput)
}

// Key type for authentication (MD5, SHA1).
func (o LookupSystemNtpResultOutput) KeyType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNtpResult) string { return v.KeyType }).(pulumi.StringOutput)
}

// Configure the FortiGate to connect to any available third-party NTP server. The structure of `ntpserver` block is documented below.
func (o LookupSystemNtpResultOutput) Ntpservers() GetSystemNtpNtpserverArrayOutput {
	return o.ApplyT(func(v LookupSystemNtpResult) []GetSystemNtpNtpserver { return v.Ntpservers }).(GetSystemNtpNtpserverArrayOutput)
}

// Enable/disable setting the FortiGate system time by synchronizing with an NTP Server.
func (o LookupSystemNtpResultOutput) Ntpsync() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNtpResult) string { return v.Ntpsync }).(pulumi.StringOutput)
}

// Enable/disable FortiGate NTP Server Mode. Your FortiGate becomes an NTP server for other devices on your network. The FortiGate relays NTP requests to its configured NTP server.
func (o LookupSystemNtpResultOutput) ServerMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNtpResult) string { return v.ServerMode }).(pulumi.StringOutput)
}

// Source IP address for communication to the NTP server.
func (o LookupSystemNtpResultOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNtpResult) string { return v.SourceIp }).(pulumi.StringOutput)
}

// Source IPv6 address for communication to the NTP server.
func (o LookupSystemNtpResultOutput) SourceIp6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNtpResult) string { return v.SourceIp6 }).(pulumi.StringOutput)
}

// NTP synchronization interval (1 - 1440 min).
func (o LookupSystemNtpResultOutput) Syncinterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemNtpResult) int { return v.Syncinterval }).(pulumi.IntOutput)
}

// Use the FortiGuard NTP server or any other available NTP Server.
func (o LookupSystemNtpResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNtpResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSystemNtpResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemNtpResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemNtpResultOutput{})
}