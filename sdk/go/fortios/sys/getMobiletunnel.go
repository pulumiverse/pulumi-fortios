// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system mobiletunnel
func LookupMobiletunnel(ctx *pulumi.Context, args *LookupMobiletunnelArgs, opts ...pulumi.InvokeOption) (*LookupMobiletunnelResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupMobiletunnelResult
	err := ctx.Invoke("fortios:sys/getMobiletunnel:getMobiletunnel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMobiletunnel.
type LookupMobiletunnelArgs struct {
	// Specify the name of the desired system mobiletunnel.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getMobiletunnel.
type LookupMobiletunnelResult struct {
	// Hash Algorithm (Keyed MD5).
	HashAlgorithm string `pulumi:"hashAlgorithm"`
	// Home IP address (Format: xxx.xxx.xxx.xxx).
	HomeAddress string `pulumi:"homeAddress"`
	// IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
	HomeAgent string `pulumi:"homeAgent"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
	Lifetime int `pulumi:"lifetime"`
	// NEMO authentication key.
	NMhaeKey string `pulumi:"nMhaeKey"`
	// NEMO authentication key type (ascii or base64).
	NMhaeKeyType string `pulumi:"nMhaeKeyType"`
	// NEMO authentication SPI (default: 256).
	NMhaeSpi int `pulumi:"nMhaeSpi"`
	// Tunnel name.
	Name string `pulumi:"name"`
	// NEMO network configuration. The structure of `network` block is documented below.
	Networks []GetMobiletunnelNetwork `pulumi:"networks"`
	// NMMO HA registration interval (5 - 300, default = 5).
	RegInterval int `pulumi:"regInterval"`
	// Maximum number of NMMO HA registration retries (1 to 30, default = 3).
	RegRetry int `pulumi:"regRetry"`
	// Time before lifetime expiraton to send NMMO HA re-registration (5 - 60, default = 60).
	RenewInterval int `pulumi:"renewInterval"`
	// Select the associated interface name from available options.
	RoamingInterface string `pulumi:"roamingInterface"`
	// Enable/disable this mobile tunnel.
	Status string `pulumi:"status"`
	// NEMO tunnnel mode (GRE tunnel).
	TunnelMode string  `pulumi:"tunnelMode"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

func LookupMobiletunnelOutput(ctx *pulumi.Context, args LookupMobiletunnelOutputArgs, opts ...pulumi.InvokeOption) LookupMobiletunnelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMobiletunnelResult, error) {
			args := v.(LookupMobiletunnelArgs)
			r, err := LookupMobiletunnel(ctx, &args, opts...)
			var s LookupMobiletunnelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMobiletunnelResultOutput)
}

// A collection of arguments for invoking getMobiletunnel.
type LookupMobiletunnelOutputArgs struct {
	// Specify the name of the desired system mobiletunnel.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupMobiletunnelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMobiletunnelArgs)(nil)).Elem()
}

// A collection of values returned by getMobiletunnel.
type LookupMobiletunnelResultOutput struct{ *pulumi.OutputState }

func (LookupMobiletunnelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMobiletunnelResult)(nil)).Elem()
}

func (o LookupMobiletunnelResultOutput) ToLookupMobiletunnelResultOutput() LookupMobiletunnelResultOutput {
	return o
}

func (o LookupMobiletunnelResultOutput) ToLookupMobiletunnelResultOutputWithContext(ctx context.Context) LookupMobiletunnelResultOutput {
	return o
}

// Hash Algorithm (Keyed MD5).
func (o LookupMobiletunnelResultOutput) HashAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobiletunnelResult) string { return v.HashAlgorithm }).(pulumi.StringOutput)
}

// Home IP address (Format: xxx.xxx.xxx.xxx).
func (o LookupMobiletunnelResultOutput) HomeAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobiletunnelResult) string { return v.HomeAddress }).(pulumi.StringOutput)
}

// IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
func (o LookupMobiletunnelResultOutput) HomeAgent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobiletunnelResult) string { return v.HomeAgent }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMobiletunnelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobiletunnelResult) string { return v.Id }).(pulumi.StringOutput)
}

// NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
func (o LookupMobiletunnelResultOutput) Lifetime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMobiletunnelResult) int { return v.Lifetime }).(pulumi.IntOutput)
}

// NEMO authentication key.
func (o LookupMobiletunnelResultOutput) NMhaeKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobiletunnelResult) string { return v.NMhaeKey }).(pulumi.StringOutput)
}

// NEMO authentication key type (ascii or base64).
func (o LookupMobiletunnelResultOutput) NMhaeKeyType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobiletunnelResult) string { return v.NMhaeKeyType }).(pulumi.StringOutput)
}

// NEMO authentication SPI (default: 256).
func (o LookupMobiletunnelResultOutput) NMhaeSpi() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMobiletunnelResult) int { return v.NMhaeSpi }).(pulumi.IntOutput)
}

// Tunnel name.
func (o LookupMobiletunnelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobiletunnelResult) string { return v.Name }).(pulumi.StringOutput)
}

// NEMO network configuration. The structure of `network` block is documented below.
func (o LookupMobiletunnelResultOutput) Networks() GetMobiletunnelNetworkArrayOutput {
	return o.ApplyT(func(v LookupMobiletunnelResult) []GetMobiletunnelNetwork { return v.Networks }).(GetMobiletunnelNetworkArrayOutput)
}

// NMMO HA registration interval (5 - 300, default = 5).
func (o LookupMobiletunnelResultOutput) RegInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMobiletunnelResult) int { return v.RegInterval }).(pulumi.IntOutput)
}

// Maximum number of NMMO HA registration retries (1 to 30, default = 3).
func (o LookupMobiletunnelResultOutput) RegRetry() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMobiletunnelResult) int { return v.RegRetry }).(pulumi.IntOutput)
}

// Time before lifetime expiraton to send NMMO HA re-registration (5 - 60, default = 60).
func (o LookupMobiletunnelResultOutput) RenewInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMobiletunnelResult) int { return v.RenewInterval }).(pulumi.IntOutput)
}

// Select the associated interface name from available options.
func (o LookupMobiletunnelResultOutput) RoamingInterface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobiletunnelResult) string { return v.RoamingInterface }).(pulumi.StringOutput)
}

// Enable/disable this mobile tunnel.
func (o LookupMobiletunnelResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobiletunnelResult) string { return v.Status }).(pulumi.StringOutput)
}

// NEMO tunnnel mode (GRE tunnel).
func (o LookupMobiletunnelResultOutput) TunnelMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobiletunnelResult) string { return v.TunnelMode }).(pulumi.StringOutput)
}

func (o LookupMobiletunnelResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMobiletunnelResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMobiletunnelResultOutput{})
}
