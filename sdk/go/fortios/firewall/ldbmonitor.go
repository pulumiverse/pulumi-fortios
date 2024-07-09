// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure server load balancing health monitors.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/firewall"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firewall.NewLdbmonitor(ctx, "trname", &firewall.LdbmonitorArgs{
//				HttpMaxRedirects: pulumi.Int(0),
//				Interval:         pulumi.Int(10),
//				Port:             pulumi.Int(0),
//				Retry:            pulumi.Int(3),
//				Timeout:          pulumi.Int(2),
//				Type:             pulumi.String("ping"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Firewall LdbMonitor can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/ldbmonitor:Ldbmonitor labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/ldbmonitor:Ldbmonitor labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Ldbmonitor struct {
	pulumi.CustomResourceState

	// Response IP expected from DNS server.
	DnsMatchIp pulumi.StringOutput `pulumi:"dnsMatchIp"`
	// Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
	DnsProtocol pulumi.StringOutput `pulumi:"dnsProtocol"`
	// Fully qualified domain name to resolve for the DNS probe.
	DnsRequestDomain pulumi.StringOutput `pulumi:"dnsRequestDomain"`
	// URL used to send a GET request to check the health of an HTTP server.
	HttpGet pulumi.StringOutput `pulumi:"httpGet"`
	// String to match the value expected in response to an HTTP-GET request.
	HttpMatch pulumi.StringOutput `pulumi:"httpMatch"`
	// The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
	HttpMaxRedirects pulumi.IntOutput `pulumi:"httpMaxRedirects"`
	// Time between health checks (default = 10). On FortiOS versions 6.2.0-7.0.15: 5 - 65635 sec. On FortiOS versions >= 7.2.0: 5 - 65535 sec.
	Interval pulumi.IntOutput `pulumi:"interval"`
	// Monitor name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (default = 0). On FortiOS versions 6.2.0-7.0.15: 0 - 65635. On FortiOS versions >= 7.2.0: 0 - 65535.
	Port pulumi.IntOutput `pulumi:"port"`
	// Number health check attempts before the server is considered down (1 - 255, default = 3).
	Retry pulumi.IntOutput `pulumi:"retry"`
	// Source IP for ldb-monitor.
	SrcIp pulumi.StringOutput `pulumi:"srcIp"`
	// Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
	Timeout pulumi.IntOutput `pulumi:"timeout"`
	// Select the Monitor type used by the health check monitor to check the health of the server. On FortiOS versions 6.2.0: PING | TCP | HTTP. On FortiOS versions 6.2.4-7.0.0: PING | TCP | HTTP | HTTPS. On FortiOS versions >= 7.0.1: PING | TCP | HTTP | HTTPS | DNS.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewLdbmonitor registers a new resource with the given unique name, arguments, and options.
func NewLdbmonitor(ctx *pulumi.Context,
	name string, args *LdbmonitorArgs, opts ...pulumi.ResourceOption) (*Ldbmonitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ldbmonitor
	err := ctx.RegisterResource("fortios:firewall/ldbmonitor:Ldbmonitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLdbmonitor gets an existing Ldbmonitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLdbmonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LdbmonitorState, opts ...pulumi.ResourceOption) (*Ldbmonitor, error) {
	var resource Ldbmonitor
	err := ctx.ReadResource("fortios:firewall/ldbmonitor:Ldbmonitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ldbmonitor resources.
type ldbmonitorState struct {
	// Response IP expected from DNS server.
	DnsMatchIp *string `pulumi:"dnsMatchIp"`
	// Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
	DnsProtocol *string `pulumi:"dnsProtocol"`
	// Fully qualified domain name to resolve for the DNS probe.
	DnsRequestDomain *string `pulumi:"dnsRequestDomain"`
	// URL used to send a GET request to check the health of an HTTP server.
	HttpGet *string `pulumi:"httpGet"`
	// String to match the value expected in response to an HTTP-GET request.
	HttpMatch *string `pulumi:"httpMatch"`
	// The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
	HttpMaxRedirects *int `pulumi:"httpMaxRedirects"`
	// Time between health checks (default = 10). On FortiOS versions 6.2.0-7.0.15: 5 - 65635 sec. On FortiOS versions >= 7.2.0: 5 - 65535 sec.
	Interval *int `pulumi:"interval"`
	// Monitor name.
	Name *string `pulumi:"name"`
	// Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (default = 0). On FortiOS versions 6.2.0-7.0.15: 0 - 65635. On FortiOS versions >= 7.2.0: 0 - 65535.
	Port *int `pulumi:"port"`
	// Number health check attempts before the server is considered down (1 - 255, default = 3).
	Retry *int `pulumi:"retry"`
	// Source IP for ldb-monitor.
	SrcIp *string `pulumi:"srcIp"`
	// Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
	Timeout *int `pulumi:"timeout"`
	// Select the Monitor type used by the health check monitor to check the health of the server. On FortiOS versions 6.2.0: PING | TCP | HTTP. On FortiOS versions 6.2.4-7.0.0: PING | TCP | HTTP | HTTPS. On FortiOS versions >= 7.0.1: PING | TCP | HTTP | HTTPS | DNS.
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type LdbmonitorState struct {
	// Response IP expected from DNS server.
	DnsMatchIp pulumi.StringPtrInput
	// Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
	DnsProtocol pulumi.StringPtrInput
	// Fully qualified domain name to resolve for the DNS probe.
	DnsRequestDomain pulumi.StringPtrInput
	// URL used to send a GET request to check the health of an HTTP server.
	HttpGet pulumi.StringPtrInput
	// String to match the value expected in response to an HTTP-GET request.
	HttpMatch pulumi.StringPtrInput
	// The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
	HttpMaxRedirects pulumi.IntPtrInput
	// Time between health checks (default = 10). On FortiOS versions 6.2.0-7.0.15: 5 - 65635 sec. On FortiOS versions >= 7.2.0: 5 - 65535 sec.
	Interval pulumi.IntPtrInput
	// Monitor name.
	Name pulumi.StringPtrInput
	// Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (default = 0). On FortiOS versions 6.2.0-7.0.15: 0 - 65635. On FortiOS versions >= 7.2.0: 0 - 65535.
	Port pulumi.IntPtrInput
	// Number health check attempts before the server is considered down (1 - 255, default = 3).
	Retry pulumi.IntPtrInput
	// Source IP for ldb-monitor.
	SrcIp pulumi.StringPtrInput
	// Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
	Timeout pulumi.IntPtrInput
	// Select the Monitor type used by the health check monitor to check the health of the server. On FortiOS versions 6.2.0: PING | TCP | HTTP. On FortiOS versions 6.2.4-7.0.0: PING | TCP | HTTP | HTTPS. On FortiOS versions >= 7.0.1: PING | TCP | HTTP | HTTPS | DNS.
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LdbmonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*ldbmonitorState)(nil)).Elem()
}

type ldbmonitorArgs struct {
	// Response IP expected from DNS server.
	DnsMatchIp *string `pulumi:"dnsMatchIp"`
	// Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
	DnsProtocol *string `pulumi:"dnsProtocol"`
	// Fully qualified domain name to resolve for the DNS probe.
	DnsRequestDomain *string `pulumi:"dnsRequestDomain"`
	// URL used to send a GET request to check the health of an HTTP server.
	HttpGet *string `pulumi:"httpGet"`
	// String to match the value expected in response to an HTTP-GET request.
	HttpMatch *string `pulumi:"httpMatch"`
	// The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
	HttpMaxRedirects *int `pulumi:"httpMaxRedirects"`
	// Time between health checks (default = 10). On FortiOS versions 6.2.0-7.0.15: 5 - 65635 sec. On FortiOS versions >= 7.2.0: 5 - 65535 sec.
	Interval *int `pulumi:"interval"`
	// Monitor name.
	Name *string `pulumi:"name"`
	// Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (default = 0). On FortiOS versions 6.2.0-7.0.15: 0 - 65635. On FortiOS versions >= 7.2.0: 0 - 65535.
	Port *int `pulumi:"port"`
	// Number health check attempts before the server is considered down (1 - 255, default = 3).
	Retry *int `pulumi:"retry"`
	// Source IP for ldb-monitor.
	SrcIp *string `pulumi:"srcIp"`
	// Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
	Timeout *int `pulumi:"timeout"`
	// Select the Monitor type used by the health check monitor to check the health of the server. On FortiOS versions 6.2.0: PING | TCP | HTTP. On FortiOS versions 6.2.4-7.0.0: PING | TCP | HTTP | HTTPS. On FortiOS versions >= 7.0.1: PING | TCP | HTTP | HTTPS | DNS.
	Type string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Ldbmonitor resource.
type LdbmonitorArgs struct {
	// Response IP expected from DNS server.
	DnsMatchIp pulumi.StringPtrInput
	// Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
	DnsProtocol pulumi.StringPtrInput
	// Fully qualified domain name to resolve for the DNS probe.
	DnsRequestDomain pulumi.StringPtrInput
	// URL used to send a GET request to check the health of an HTTP server.
	HttpGet pulumi.StringPtrInput
	// String to match the value expected in response to an HTTP-GET request.
	HttpMatch pulumi.StringPtrInput
	// The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
	HttpMaxRedirects pulumi.IntPtrInput
	// Time between health checks (default = 10). On FortiOS versions 6.2.0-7.0.15: 5 - 65635 sec. On FortiOS versions >= 7.2.0: 5 - 65535 sec.
	Interval pulumi.IntPtrInput
	// Monitor name.
	Name pulumi.StringPtrInput
	// Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (default = 0). On FortiOS versions 6.2.0-7.0.15: 0 - 65635. On FortiOS versions >= 7.2.0: 0 - 65535.
	Port pulumi.IntPtrInput
	// Number health check attempts before the server is considered down (1 - 255, default = 3).
	Retry pulumi.IntPtrInput
	// Source IP for ldb-monitor.
	SrcIp pulumi.StringPtrInput
	// Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
	Timeout pulumi.IntPtrInput
	// Select the Monitor type used by the health check monitor to check the health of the server. On FortiOS versions 6.2.0: PING | TCP | HTTP. On FortiOS versions 6.2.4-7.0.0: PING | TCP | HTTP | HTTPS. On FortiOS versions >= 7.0.1: PING | TCP | HTTP | HTTPS | DNS.
	Type pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LdbmonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ldbmonitorArgs)(nil)).Elem()
}

type LdbmonitorInput interface {
	pulumi.Input

	ToLdbmonitorOutput() LdbmonitorOutput
	ToLdbmonitorOutputWithContext(ctx context.Context) LdbmonitorOutput
}

func (*Ldbmonitor) ElementType() reflect.Type {
	return reflect.TypeOf((**Ldbmonitor)(nil)).Elem()
}

func (i *Ldbmonitor) ToLdbmonitorOutput() LdbmonitorOutput {
	return i.ToLdbmonitorOutputWithContext(context.Background())
}

func (i *Ldbmonitor) ToLdbmonitorOutputWithContext(ctx context.Context) LdbmonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LdbmonitorOutput)
}

// LdbmonitorArrayInput is an input type that accepts LdbmonitorArray and LdbmonitorArrayOutput values.
// You can construct a concrete instance of `LdbmonitorArrayInput` via:
//
//	LdbmonitorArray{ LdbmonitorArgs{...} }
type LdbmonitorArrayInput interface {
	pulumi.Input

	ToLdbmonitorArrayOutput() LdbmonitorArrayOutput
	ToLdbmonitorArrayOutputWithContext(context.Context) LdbmonitorArrayOutput
}

type LdbmonitorArray []LdbmonitorInput

func (LdbmonitorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ldbmonitor)(nil)).Elem()
}

func (i LdbmonitorArray) ToLdbmonitorArrayOutput() LdbmonitorArrayOutput {
	return i.ToLdbmonitorArrayOutputWithContext(context.Background())
}

func (i LdbmonitorArray) ToLdbmonitorArrayOutputWithContext(ctx context.Context) LdbmonitorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LdbmonitorArrayOutput)
}

// LdbmonitorMapInput is an input type that accepts LdbmonitorMap and LdbmonitorMapOutput values.
// You can construct a concrete instance of `LdbmonitorMapInput` via:
//
//	LdbmonitorMap{ "key": LdbmonitorArgs{...} }
type LdbmonitorMapInput interface {
	pulumi.Input

	ToLdbmonitorMapOutput() LdbmonitorMapOutput
	ToLdbmonitorMapOutputWithContext(context.Context) LdbmonitorMapOutput
}

type LdbmonitorMap map[string]LdbmonitorInput

func (LdbmonitorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ldbmonitor)(nil)).Elem()
}

func (i LdbmonitorMap) ToLdbmonitorMapOutput() LdbmonitorMapOutput {
	return i.ToLdbmonitorMapOutputWithContext(context.Background())
}

func (i LdbmonitorMap) ToLdbmonitorMapOutputWithContext(ctx context.Context) LdbmonitorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LdbmonitorMapOutput)
}

type LdbmonitorOutput struct{ *pulumi.OutputState }

func (LdbmonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ldbmonitor)(nil)).Elem()
}

func (o LdbmonitorOutput) ToLdbmonitorOutput() LdbmonitorOutput {
	return o
}

func (o LdbmonitorOutput) ToLdbmonitorOutputWithContext(ctx context.Context) LdbmonitorOutput {
	return o
}

// Response IP expected from DNS server.
func (o LdbmonitorOutput) DnsMatchIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Ldbmonitor) pulumi.StringOutput { return v.DnsMatchIp }).(pulumi.StringOutput)
}

// Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
func (o LdbmonitorOutput) DnsProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Ldbmonitor) pulumi.StringOutput { return v.DnsProtocol }).(pulumi.StringOutput)
}

// Fully qualified domain name to resolve for the DNS probe.
func (o LdbmonitorOutput) DnsRequestDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *Ldbmonitor) pulumi.StringOutput { return v.DnsRequestDomain }).(pulumi.StringOutput)
}

// URL used to send a GET request to check the health of an HTTP server.
func (o LdbmonitorOutput) HttpGet() pulumi.StringOutput {
	return o.ApplyT(func(v *Ldbmonitor) pulumi.StringOutput { return v.HttpGet }).(pulumi.StringOutput)
}

// String to match the value expected in response to an HTTP-GET request.
func (o LdbmonitorOutput) HttpMatch() pulumi.StringOutput {
	return o.ApplyT(func(v *Ldbmonitor) pulumi.StringOutput { return v.HttpMatch }).(pulumi.StringOutput)
}

// The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
func (o LdbmonitorOutput) HttpMaxRedirects() pulumi.IntOutput {
	return o.ApplyT(func(v *Ldbmonitor) pulumi.IntOutput { return v.HttpMaxRedirects }).(pulumi.IntOutput)
}

// Time between health checks (default = 10). On FortiOS versions 6.2.0-7.0.15: 5 - 65635 sec. On FortiOS versions >= 7.2.0: 5 - 65535 sec.
func (o LdbmonitorOutput) Interval() pulumi.IntOutput {
	return o.ApplyT(func(v *Ldbmonitor) pulumi.IntOutput { return v.Interval }).(pulumi.IntOutput)
}

// Monitor name.
func (o LdbmonitorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Ldbmonitor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (default = 0). On FortiOS versions 6.2.0-7.0.15: 0 - 65635. On FortiOS versions >= 7.2.0: 0 - 65535.
func (o LdbmonitorOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Ldbmonitor) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Number health check attempts before the server is considered down (1 - 255, default = 3).
func (o LdbmonitorOutput) Retry() pulumi.IntOutput {
	return o.ApplyT(func(v *Ldbmonitor) pulumi.IntOutput { return v.Retry }).(pulumi.IntOutput)
}

// Source IP for ldb-monitor.
func (o LdbmonitorOutput) SrcIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Ldbmonitor) pulumi.StringOutput { return v.SrcIp }).(pulumi.StringOutput)
}

// Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
func (o LdbmonitorOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Ldbmonitor) pulumi.IntOutput { return v.Timeout }).(pulumi.IntOutput)
}

// Select the Monitor type used by the health check monitor to check the health of the server. On FortiOS versions 6.2.0: PING | TCP | HTTP. On FortiOS versions 6.2.4-7.0.0: PING | TCP | HTTP | HTTPS. On FortiOS versions >= 7.0.1: PING | TCP | HTTP | HTTPS | DNS.
func (o LdbmonitorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Ldbmonitor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o LdbmonitorOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Ldbmonitor) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type LdbmonitorArrayOutput struct{ *pulumi.OutputState }

func (LdbmonitorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ldbmonitor)(nil)).Elem()
}

func (o LdbmonitorArrayOutput) ToLdbmonitorArrayOutput() LdbmonitorArrayOutput {
	return o
}

func (o LdbmonitorArrayOutput) ToLdbmonitorArrayOutputWithContext(ctx context.Context) LdbmonitorArrayOutput {
	return o
}

func (o LdbmonitorArrayOutput) Index(i pulumi.IntInput) LdbmonitorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ldbmonitor {
		return vs[0].([]*Ldbmonitor)[vs[1].(int)]
	}).(LdbmonitorOutput)
}

type LdbmonitorMapOutput struct{ *pulumi.OutputState }

func (LdbmonitorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ldbmonitor)(nil)).Elem()
}

func (o LdbmonitorMapOutput) ToLdbmonitorMapOutput() LdbmonitorMapOutput {
	return o
}

func (o LdbmonitorMapOutput) ToLdbmonitorMapOutputWithContext(ctx context.Context) LdbmonitorMapOutput {
	return o
}

func (o LdbmonitorMapOutput) MapIndex(k pulumi.StringInput) LdbmonitorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ldbmonitor {
		return vs[0].(map[string]*Ldbmonitor)[vs[1].(string)]
	}).(LdbmonitorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LdbmonitorInput)(nil)).Elem(), &Ldbmonitor{})
	pulumi.RegisterInputType(reflect.TypeOf((*LdbmonitorArrayInput)(nil)).Elem(), LdbmonitorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LdbmonitorMapInput)(nil)).Elem(), LdbmonitorMap{})
	pulumi.RegisterOutputType(LdbmonitorOutput{})
	pulumi.RegisterOutputType(LdbmonitorArrayOutput{})
	pulumi.RegisterOutputType(LdbmonitorMapOutput{})
}
