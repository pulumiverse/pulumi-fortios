// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on fortios system ftmpush
func LookupFtmpush(ctx *pulumi.Context, args *LookupFtmpushArgs, opts ...pulumi.InvokeOption) (*LookupFtmpushResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFtmpushResult
	err := ctx.Invoke("fortios:system/getFtmpush:getFtmpush", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFtmpush.
type LookupFtmpushArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFtmpush.
type LookupFtmpushResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IPv4 address or domain name of FortiToken Mobile push services server.
	Server string `pulumi:"server"`
	// Name of the server certificate to be used for SSL (default = Fortinet_Factory).
	ServerCert string `pulumi:"serverCert"`
	// IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).
	ServerIp string `pulumi:"serverIp"`
	// Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).
	ServerPort int `pulumi:"serverPort"`
	// Enable/disable the use of FortiToken Mobile push services.
	Status    string  `pulumi:"status"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupFtmpushOutput(ctx *pulumi.Context, args LookupFtmpushOutputArgs, opts ...pulumi.InvokeOption) LookupFtmpushResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFtmpushResult, error) {
			args := v.(LookupFtmpushArgs)
			r, err := LookupFtmpush(ctx, &args, opts...)
			var s LookupFtmpushResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFtmpushResultOutput)
}

// A collection of arguments for invoking getFtmpush.
type LookupFtmpushOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFtmpushOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFtmpushArgs)(nil)).Elem()
}

// A collection of values returned by getFtmpush.
type LookupFtmpushResultOutput struct{ *pulumi.OutputState }

func (LookupFtmpushResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFtmpushResult)(nil)).Elem()
}

func (o LookupFtmpushResultOutput) ToLookupFtmpushResultOutput() LookupFtmpushResultOutput {
	return o
}

func (o LookupFtmpushResultOutput) ToLookupFtmpushResultOutputWithContext(ctx context.Context) LookupFtmpushResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFtmpushResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFtmpushResult) string { return v.Id }).(pulumi.StringOutput)
}

// IPv4 address or domain name of FortiToken Mobile push services server.
func (o LookupFtmpushResultOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFtmpushResult) string { return v.Server }).(pulumi.StringOutput)
}

// Name of the server certificate to be used for SSL (default = Fortinet_Factory).
func (o LookupFtmpushResultOutput) ServerCert() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFtmpushResult) string { return v.ServerCert }).(pulumi.StringOutput)
}

// IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).
func (o LookupFtmpushResultOutput) ServerIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFtmpushResult) string { return v.ServerIp }).(pulumi.StringOutput)
}

// Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).
func (o LookupFtmpushResultOutput) ServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFtmpushResult) int { return v.ServerPort }).(pulumi.IntOutput)
}

// Enable/disable the use of FortiToken Mobile push services.
func (o LookupFtmpushResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFtmpushResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupFtmpushResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFtmpushResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFtmpushResultOutput{})
}