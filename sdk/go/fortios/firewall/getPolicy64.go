// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall policy64
func LookupPolicy64(ctx *pulumi.Context, args *LookupPolicy64Args, opts ...pulumi.InvokeOption) (*LookupPolicy64Result, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupPolicy64Result
	err := ctx.Invoke("fortios:firewall/getPolicy64:getPolicy64", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicy64.
type LookupPolicy64Args struct {
	// Specify the policyid of the desired firewall policy64.
	Policyid int `pulumi:"policyid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getPolicy64.
type LookupPolicy64Result struct {
	// Policy action.
	Action string `pulumi:"action"`
	// Comment.
	Comments string `pulumi:"comments"`
	// Destination address name. The structure of `dstaddr` block is documented below.
	Dstaddrs []GetPolicy64Dstaddr `pulumi:"dstaddrs"`
	// Destination interface name.
	Dstintf string `pulumi:"dstintf"`
	// Enable/disable policy fixed port.
	Fixedport string `pulumi:"fixedport"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable/disable policy64 IP pool.
	Ippool string `pulumi:"ippool"`
	// Enable/disable policy log traffic.
	Logtraffic string `pulumi:"logtraffic"`
	// Record logs when a session starts and ends.
	LogtrafficStart string `pulumi:"logtrafficStart"`
	// IP pool name.
	Name string `pulumi:"name"`
	// Per-IP traffic shaper.
	PerIpShaper string `pulumi:"perIpShaper"`
	// Enable/disable permit any host in.
	PermitAnyHost string `pulumi:"permitAnyHost"`
	// Policy ID.
	Policyid int `pulumi:"policyid"`
	// Policy IP pool names. The structure of `poolname` block is documented below.
	Poolnames []GetPolicy64Poolname `pulumi:"poolnames"`
	// Schedule name.
	Schedule string `pulumi:"schedule"`
	// Service name. The structure of `service` block is documented below.
	Services []GetPolicy64Service `pulumi:"services"`
	// Source address name. The structure of `srcaddr` block is documented below.
	Srcaddrs []GetPolicy64Srcaddr `pulumi:"srcaddrs"`
	// Source interface name.
	Srcintf string `pulumi:"srcintf"`
	// Enable/disable policy status.
	Status string `pulumi:"status"`
	// TCP MSS value of receiver.
	TcpMssReceiver int `pulumi:"tcpMssReceiver"`
	// TCP MSS value of sender.
	TcpMssSender int `pulumi:"tcpMssSender"`
	// Traffic shaper.
	TrafficShaper string `pulumi:"trafficShaper"`
	// Reverse traffic shaper.
	TrafficShaperReverse string `pulumi:"trafficShaperReverse"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid      string  `pulumi:"uuid"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupPolicy64Output(ctx *pulumi.Context, args LookupPolicy64OutputArgs, opts ...pulumi.InvokeOption) LookupPolicy64ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicy64Result, error) {
			args := v.(LookupPolicy64Args)
			r, err := LookupPolicy64(ctx, &args, opts...)
			var s LookupPolicy64Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicy64ResultOutput)
}

// A collection of arguments for invoking getPolicy64.
type LookupPolicy64OutputArgs struct {
	// Specify the policyid of the desired firewall policy64.
	Policyid pulumi.IntInput `pulumi:"policyid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupPolicy64OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicy64Args)(nil)).Elem()
}

// A collection of values returned by getPolicy64.
type LookupPolicy64ResultOutput struct{ *pulumi.OutputState }

func (LookupPolicy64ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicy64Result)(nil)).Elem()
}

func (o LookupPolicy64ResultOutput) ToLookupPolicy64ResultOutput() LookupPolicy64ResultOutput {
	return o
}

func (o LookupPolicy64ResultOutput) ToLookupPolicy64ResultOutputWithContext(ctx context.Context) LookupPolicy64ResultOutput {
	return o
}

// Policy action.
func (o LookupPolicy64ResultOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicy64Result) string { return v.Action }).(pulumi.StringOutput)
}

// Comment.
func (o LookupPolicy64ResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicy64Result) string { return v.Comments }).(pulumi.StringOutput)
}

// Destination address name. The structure of `dstaddr` block is documented below.
func (o LookupPolicy64ResultOutput) Dstaddrs() GetPolicy64DstaddrArrayOutput {
	return o.ApplyT(func(v LookupPolicy64Result) []GetPolicy64Dstaddr { return v.Dstaddrs }).(GetPolicy64DstaddrArrayOutput)
}

// Destination interface name.
func (o LookupPolicy64ResultOutput) Dstintf() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicy64Result) string { return v.Dstintf }).(pulumi.StringOutput)
}

// Enable/disable policy fixed port.
func (o LookupPolicy64ResultOutput) Fixedport() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicy64Result) string { return v.Fixedport }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPolicy64ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicy64Result) string { return v.Id }).(pulumi.StringOutput)
}

// Enable/disable policy64 IP pool.
func (o LookupPolicy64ResultOutput) Ippool() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicy64Result) string { return v.Ippool }).(pulumi.StringOutput)
}

// Enable/disable policy log traffic.
func (o LookupPolicy64ResultOutput) Logtraffic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicy64Result) string { return v.Logtraffic }).(pulumi.StringOutput)
}

// Record logs when a session starts and ends.
func (o LookupPolicy64ResultOutput) LogtrafficStart() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicy64Result) string { return v.LogtrafficStart }).(pulumi.StringOutput)
}

// IP pool name.
func (o LookupPolicy64ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicy64Result) string { return v.Name }).(pulumi.StringOutput)
}

// Per-IP traffic shaper.
func (o LookupPolicy64ResultOutput) PerIpShaper() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicy64Result) string { return v.PerIpShaper }).(pulumi.StringOutput)
}

// Enable/disable permit any host in.
func (o LookupPolicy64ResultOutput) PermitAnyHost() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicy64Result) string { return v.PermitAnyHost }).(pulumi.StringOutput)
}

// Policy ID.
func (o LookupPolicy64ResultOutput) Policyid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPolicy64Result) int { return v.Policyid }).(pulumi.IntOutput)
}

// Policy IP pool names. The structure of `poolname` block is documented below.
func (o LookupPolicy64ResultOutput) Poolnames() GetPolicy64PoolnameArrayOutput {
	return o.ApplyT(func(v LookupPolicy64Result) []GetPolicy64Poolname { return v.Poolnames }).(GetPolicy64PoolnameArrayOutput)
}

// Schedule name.
func (o LookupPolicy64ResultOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicy64Result) string { return v.Schedule }).(pulumi.StringOutput)
}

// Service name. The structure of `service` block is documented below.
func (o LookupPolicy64ResultOutput) Services() GetPolicy64ServiceArrayOutput {
	return o.ApplyT(func(v LookupPolicy64Result) []GetPolicy64Service { return v.Services }).(GetPolicy64ServiceArrayOutput)
}

// Source address name. The structure of `srcaddr` block is documented below.
func (o LookupPolicy64ResultOutput) Srcaddrs() GetPolicy64SrcaddrArrayOutput {
	return o.ApplyT(func(v LookupPolicy64Result) []GetPolicy64Srcaddr { return v.Srcaddrs }).(GetPolicy64SrcaddrArrayOutput)
}

// Source interface name.
func (o LookupPolicy64ResultOutput) Srcintf() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicy64Result) string { return v.Srcintf }).(pulumi.StringOutput)
}

// Enable/disable policy status.
func (o LookupPolicy64ResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicy64Result) string { return v.Status }).(pulumi.StringOutput)
}

// TCP MSS value of receiver.
func (o LookupPolicy64ResultOutput) TcpMssReceiver() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPolicy64Result) int { return v.TcpMssReceiver }).(pulumi.IntOutput)
}

// TCP MSS value of sender.
func (o LookupPolicy64ResultOutput) TcpMssSender() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPolicy64Result) int { return v.TcpMssSender }).(pulumi.IntOutput)
}

// Traffic shaper.
func (o LookupPolicy64ResultOutput) TrafficShaper() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicy64Result) string { return v.TrafficShaper }).(pulumi.StringOutput)
}

// Reverse traffic shaper.
func (o LookupPolicy64ResultOutput) TrafficShaperReverse() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicy64Result) string { return v.TrafficShaperReverse }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o LookupPolicy64ResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicy64Result) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupPolicy64ResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicy64Result) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicy64ResultOutput{})
}
