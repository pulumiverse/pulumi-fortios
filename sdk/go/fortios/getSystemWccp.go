// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system wccp
func LookupSystemWccp(ctx *pulumi.Context, args *LookupSystemWccpArgs, opts ...pulumi.InvokeOption) (*LookupSystemWccpResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemWccpResult
	err := ctx.Invoke("fortios:index/getSystemWccp:getSystemWccp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemWccp.
type LookupSystemWccpArgs struct {
	// Specify the serviceId of the desired system wccp.
	ServiceId string `pulumi:"serviceId"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemWccp.
type LookupSystemWccpResult struct {
	// Assignment bucket format for the WCCP cache engine.
	AssignmentBucketFormat string `pulumi:"assignmentBucketFormat"`
	// Assignment destination address mask.
	AssignmentDstaddrMask string `pulumi:"assignmentDstaddrMask"`
	// Hash key assignment preference.
	AssignmentMethod string `pulumi:"assignmentMethod"`
	// Assignment source address mask.
	AssignmentSrcaddrMask string `pulumi:"assignmentSrcaddrMask"`
	// Assignment of hash weight/ratio for the WCCP cache engine.
	AssignmentWeight int `pulumi:"assignmentWeight"`
	// Enable/disable MD5 authentication.
	Authentication string `pulumi:"authentication"`
	// Method used to forward traffic to the routers or to return to the cache engine.
	CacheEngineMethod string `pulumi:"cacheEngineMethod"`
	// IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
	CacheId string `pulumi:"cacheId"`
	// Method used to forward traffic to the cache servers.
	ForwardMethod string `pulumi:"forwardMethod"`
	// IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
	GroupAddress string `pulumi:"groupAddress"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Password for MD5 authentication.
	Password string `pulumi:"password"`
	// Service ports.
	Ports string `pulumi:"ports"`
	// Match method.
	PortsDefined string `pulumi:"portsDefined"`
	// Hash method.
	PrimaryHash string `pulumi:"primaryHash"`
	// Service priority.
	Priority int `pulumi:"priority"`
	// Service protocol.
	Protocol int `pulumi:"protocol"`
	// Method used to decline a redirected packet and return it to the FortiGate.
	ReturnMethod string `pulumi:"returnMethod"`
	// IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
	RouterId string `pulumi:"routerId"`
	// IP addresses of one or more WCCP routers.
	RouterList string `pulumi:"routerList"`
	// IP addresses and netmasks for up to four cache servers.
	ServerList string `pulumi:"serverList"`
	// Cache server type.
	ServerType string `pulumi:"serverType"`
	// Service ID.
	ServiceId string `pulumi:"serviceId"`
	// WCCP service type used by the cache server for logical interception and redirection of traffic.
	ServiceType string  `pulumi:"serviceType"`
	Vdomparam   *string `pulumi:"vdomparam"`
}

func LookupSystemWccpOutput(ctx *pulumi.Context, args LookupSystemWccpOutputArgs, opts ...pulumi.InvokeOption) LookupSystemWccpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemWccpResult, error) {
			args := v.(LookupSystemWccpArgs)
			r, err := LookupSystemWccp(ctx, &args, opts...)
			var s LookupSystemWccpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemWccpResultOutput)
}

// A collection of arguments for invoking getSystemWccp.
type LookupSystemWccpOutputArgs struct {
	// Specify the serviceId of the desired system wccp.
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemWccpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemWccpArgs)(nil)).Elem()
}

// A collection of values returned by getSystemWccp.
type LookupSystemWccpResultOutput struct{ *pulumi.OutputState }

func (LookupSystemWccpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemWccpResult)(nil)).Elem()
}

func (o LookupSystemWccpResultOutput) ToLookupSystemWccpResultOutput() LookupSystemWccpResultOutput {
	return o
}

func (o LookupSystemWccpResultOutput) ToLookupSystemWccpResultOutputWithContext(ctx context.Context) LookupSystemWccpResultOutput {
	return o
}

// Assignment bucket format for the WCCP cache engine.
func (o LookupSystemWccpResultOutput) AssignmentBucketFormat() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.AssignmentBucketFormat }).(pulumi.StringOutput)
}

// Assignment destination address mask.
func (o LookupSystemWccpResultOutput) AssignmentDstaddrMask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.AssignmentDstaddrMask }).(pulumi.StringOutput)
}

// Hash key assignment preference.
func (o LookupSystemWccpResultOutput) AssignmentMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.AssignmentMethod }).(pulumi.StringOutput)
}

// Assignment source address mask.
func (o LookupSystemWccpResultOutput) AssignmentSrcaddrMask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.AssignmentSrcaddrMask }).(pulumi.StringOutput)
}

// Assignment of hash weight/ratio for the WCCP cache engine.
func (o LookupSystemWccpResultOutput) AssignmentWeight() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) int { return v.AssignmentWeight }).(pulumi.IntOutput)
}

// Enable/disable MD5 authentication.
func (o LookupSystemWccpResultOutput) Authentication() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.Authentication }).(pulumi.StringOutput)
}

// Method used to forward traffic to the routers or to return to the cache engine.
func (o LookupSystemWccpResultOutput) CacheEngineMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.CacheEngineMethod }).(pulumi.StringOutput)
}

// IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
func (o LookupSystemWccpResultOutput) CacheId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.CacheId }).(pulumi.StringOutput)
}

// Method used to forward traffic to the cache servers.
func (o LookupSystemWccpResultOutput) ForwardMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.ForwardMethod }).(pulumi.StringOutput)
}

// IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
func (o LookupSystemWccpResultOutput) GroupAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.GroupAddress }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemWccpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.Id }).(pulumi.StringOutput)
}

// Password for MD5 authentication.
func (o LookupSystemWccpResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.Password }).(pulumi.StringOutput)
}

// Service ports.
func (o LookupSystemWccpResultOutput) Ports() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.Ports }).(pulumi.StringOutput)
}

// Match method.
func (o LookupSystemWccpResultOutput) PortsDefined() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.PortsDefined }).(pulumi.StringOutput)
}

// Hash method.
func (o LookupSystemWccpResultOutput) PrimaryHash() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.PrimaryHash }).(pulumi.StringOutput)
}

// Service priority.
func (o LookupSystemWccpResultOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) int { return v.Priority }).(pulumi.IntOutput)
}

// Service protocol.
func (o LookupSystemWccpResultOutput) Protocol() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) int { return v.Protocol }).(pulumi.IntOutput)
}

// Method used to decline a redirected packet and return it to the FortiGate.
func (o LookupSystemWccpResultOutput) ReturnMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.ReturnMethod }).(pulumi.StringOutput)
}

// IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
func (o LookupSystemWccpResultOutput) RouterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.RouterId }).(pulumi.StringOutput)
}

// IP addresses of one or more WCCP routers.
func (o LookupSystemWccpResultOutput) RouterList() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.RouterList }).(pulumi.StringOutput)
}

// IP addresses and netmasks for up to four cache servers.
func (o LookupSystemWccpResultOutput) ServerList() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.ServerList }).(pulumi.StringOutput)
}

// Cache server type.
func (o LookupSystemWccpResultOutput) ServerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.ServerType }).(pulumi.StringOutput)
}

// Service ID.
func (o LookupSystemWccpResultOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.ServiceId }).(pulumi.StringOutput)
}

// WCCP service type used by the cache server for logical interception and redirection of traffic.
func (o LookupSystemWccpResultOutput) ServiceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) string { return v.ServiceType }).(pulumi.StringOutput)
}

func (o LookupSystemWccpResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemWccpResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemWccpResultOutput{})
}
