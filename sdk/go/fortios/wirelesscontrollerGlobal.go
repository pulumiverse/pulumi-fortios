// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure wireless controller global settings.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fortios.NewWirelesscontrollerGlobal(ctx, "trname", &fortios.WirelesscontrollerGlobalArgs{
//				ApLogServer:           pulumi.String("disable"),
//				ApLogServerIp:         pulumi.String("0.0.0.0"),
//				ApLogServerPort:       pulumi.Int(0),
//				ControlMessageOffload: pulumi.String("ebp-frame aeroscout-tag ap-list sta-list sta-cap-list stats aeroscout-mu"),
//				DataEthernetIi:        pulumi.String("disable"),
//				DiscoveryMcAddr:       pulumi.String("224.0.1.140"),
//				FiappEthType:          pulumi.Int(5252),
//				ImageDownload:         pulumi.String("enable"),
//				IpsecBaseIp:           pulumi.String("169.254.0.1"),
//				LinkAggregation:       pulumi.String("disable"),
//				MaxClients:            pulumi.Int(0),
//				MaxRetransmit:         pulumi.Int(3),
//				MeshEthType:           pulumi.Int(8755),
//				RogueScanMacAdjacency: pulumi.Int(7),
//				WtpShare:              pulumi.String("disable"),
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
// # WirelessController Global can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerGlobal:WirelesscontrollerGlobal labelname WirelessControllerGlobal
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerGlobal:WirelesscontrollerGlobal labelname WirelessControllerGlobal
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type WirelesscontrollerGlobal struct {
	pulumi.CustomResourceState

	// Enable/disable configuring APs or FortiAPs to send log messages to a syslog server (default = disable). Valid values: `enable`, `disable`.
	ApLogServer pulumi.StringOutput `pulumi:"apLogServer"`
	// IP address that APs or FortiAPs send log messages to.
	ApLogServerIp pulumi.StringOutput `pulumi:"apLogServerIp"`
	// Port that APs or FortiAPs send log messages to.
	ApLogServerPort pulumi.IntOutput `pulumi:"apLogServerPort"`
	// Configure CAPWAP control message data channel offload.
	ControlMessageOffload pulumi.StringOutput `pulumi:"controlMessageOffload"`
	// Configure the wireless controller to use Ethernet II or 802.3 frames with 802.3 data tunnel mode (default = disable). Valid values: `enable`, `disable`.
	DataEthernetIi pulumi.StringOutput `pulumi:"dataEthernetIi"`
	// Multicast IP address for AP discovery (default = 244.0.1.140).
	DiscoveryMcAddr pulumi.StringOutput `pulumi:"discoveryMcAddr"`
	// Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
	FiappEthType pulumi.IntOutput `pulumi:"fiappEthType"`
	// Enable/disable WTP image download at join time. Valid values: `enable`, `disable`.
	ImageDownload pulumi.StringOutput `pulumi:"imageDownload"`
	// Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
	IpsecBaseIp pulumi.StringOutput `pulumi:"ipsecBaseIp"`
	// Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable). Valid values: `enable`, `disable`.
	LinkAggregation pulumi.StringOutput `pulumi:"linkAggregation"`
	// Description of the location of the wireless controller.
	Location pulumi.StringOutput `pulumi:"location"`
	// Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
	MaxClients pulumi.IntOutput `pulumi:"maxClients"`
	// Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
	MaxRetransmit pulumi.IntOutput `pulumi:"maxRetransmit"`
	// Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
	MeshEthType pulumi.IntOutput `pulumi:"meshEthType"`
	// Interval in seconds between two WiFi network access control (NAC) checks (10 - 600, default = 120).
	NacInterval pulumi.IntOutput `pulumi:"nacInterval"`
	// Name of the wireless controller.
	Name pulumi.StringOutput `pulumi:"name"`
	// Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
	RogueScanMacAdjacency pulumi.IntOutput `pulumi:"rogueScanMacAdjacency"`
	// Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
	TunnelMode pulumi.StringOutput `pulumi:"tunnelMode"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable sharing of WTPs between VDOMs. Valid values: `enable`, `disable`.
	WtpShare pulumi.StringOutput `pulumi:"wtpShare"`
}

// NewWirelesscontrollerGlobal registers a new resource with the given unique name, arguments, and options.
func NewWirelesscontrollerGlobal(ctx *pulumi.Context,
	name string, args *WirelesscontrollerGlobalArgs, opts ...pulumi.ResourceOption) (*WirelesscontrollerGlobal, error) {
	if args == nil {
		args = &WirelesscontrollerGlobalArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WirelesscontrollerGlobal
	err := ctx.RegisterResource("fortios:index/wirelesscontrollerGlobal:WirelesscontrollerGlobal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelesscontrollerGlobal gets an existing WirelesscontrollerGlobal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelesscontrollerGlobal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelesscontrollerGlobalState, opts ...pulumi.ResourceOption) (*WirelesscontrollerGlobal, error) {
	var resource WirelesscontrollerGlobal
	err := ctx.ReadResource("fortios:index/wirelesscontrollerGlobal:WirelesscontrollerGlobal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelesscontrollerGlobal resources.
type wirelesscontrollerGlobalState struct {
	// Enable/disable configuring APs or FortiAPs to send log messages to a syslog server (default = disable). Valid values: `enable`, `disable`.
	ApLogServer *string `pulumi:"apLogServer"`
	// IP address that APs or FortiAPs send log messages to.
	ApLogServerIp *string `pulumi:"apLogServerIp"`
	// Port that APs or FortiAPs send log messages to.
	ApLogServerPort *int `pulumi:"apLogServerPort"`
	// Configure CAPWAP control message data channel offload.
	ControlMessageOffload *string `pulumi:"controlMessageOffload"`
	// Configure the wireless controller to use Ethernet II or 802.3 frames with 802.3 data tunnel mode (default = disable). Valid values: `enable`, `disable`.
	DataEthernetIi *string `pulumi:"dataEthernetIi"`
	// Multicast IP address for AP discovery (default = 244.0.1.140).
	DiscoveryMcAddr *string `pulumi:"discoveryMcAddr"`
	// Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
	FiappEthType *int `pulumi:"fiappEthType"`
	// Enable/disable WTP image download at join time. Valid values: `enable`, `disable`.
	ImageDownload *string `pulumi:"imageDownload"`
	// Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
	IpsecBaseIp *string `pulumi:"ipsecBaseIp"`
	// Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable). Valid values: `enable`, `disable`.
	LinkAggregation *string `pulumi:"linkAggregation"`
	// Description of the location of the wireless controller.
	Location *string `pulumi:"location"`
	// Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
	MaxClients *int `pulumi:"maxClients"`
	// Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
	MaxRetransmit *int `pulumi:"maxRetransmit"`
	// Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
	MeshEthType *int `pulumi:"meshEthType"`
	// Interval in seconds between two WiFi network access control (NAC) checks (10 - 600, default = 120).
	NacInterval *int `pulumi:"nacInterval"`
	// Name of the wireless controller.
	Name *string `pulumi:"name"`
	// Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
	RogueScanMacAdjacency *int `pulumi:"rogueScanMacAdjacency"`
	// Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
	TunnelMode *string `pulumi:"tunnelMode"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable sharing of WTPs between VDOMs. Valid values: `enable`, `disable`.
	WtpShare *string `pulumi:"wtpShare"`
}

type WirelesscontrollerGlobalState struct {
	// Enable/disable configuring APs or FortiAPs to send log messages to a syslog server (default = disable). Valid values: `enable`, `disable`.
	ApLogServer pulumi.StringPtrInput
	// IP address that APs or FortiAPs send log messages to.
	ApLogServerIp pulumi.StringPtrInput
	// Port that APs or FortiAPs send log messages to.
	ApLogServerPort pulumi.IntPtrInput
	// Configure CAPWAP control message data channel offload.
	ControlMessageOffload pulumi.StringPtrInput
	// Configure the wireless controller to use Ethernet II or 802.3 frames with 802.3 data tunnel mode (default = disable). Valid values: `enable`, `disable`.
	DataEthernetIi pulumi.StringPtrInput
	// Multicast IP address for AP discovery (default = 244.0.1.140).
	DiscoveryMcAddr pulumi.StringPtrInput
	// Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
	FiappEthType pulumi.IntPtrInput
	// Enable/disable WTP image download at join time. Valid values: `enable`, `disable`.
	ImageDownload pulumi.StringPtrInput
	// Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
	IpsecBaseIp pulumi.StringPtrInput
	// Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable). Valid values: `enable`, `disable`.
	LinkAggregation pulumi.StringPtrInput
	// Description of the location of the wireless controller.
	Location pulumi.StringPtrInput
	// Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
	MaxClients pulumi.IntPtrInput
	// Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
	MaxRetransmit pulumi.IntPtrInput
	// Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
	MeshEthType pulumi.IntPtrInput
	// Interval in seconds between two WiFi network access control (NAC) checks (10 - 600, default = 120).
	NacInterval pulumi.IntPtrInput
	// Name of the wireless controller.
	Name pulumi.StringPtrInput
	// Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
	RogueScanMacAdjacency pulumi.IntPtrInput
	// Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
	TunnelMode pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable sharing of WTPs between VDOMs. Valid values: `enable`, `disable`.
	WtpShare pulumi.StringPtrInput
}

func (WirelesscontrollerGlobalState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerGlobalState)(nil)).Elem()
}

type wirelesscontrollerGlobalArgs struct {
	// Enable/disable configuring APs or FortiAPs to send log messages to a syslog server (default = disable). Valid values: `enable`, `disable`.
	ApLogServer *string `pulumi:"apLogServer"`
	// IP address that APs or FortiAPs send log messages to.
	ApLogServerIp *string `pulumi:"apLogServerIp"`
	// Port that APs or FortiAPs send log messages to.
	ApLogServerPort *int `pulumi:"apLogServerPort"`
	// Configure CAPWAP control message data channel offload.
	ControlMessageOffload *string `pulumi:"controlMessageOffload"`
	// Configure the wireless controller to use Ethernet II or 802.3 frames with 802.3 data tunnel mode (default = disable). Valid values: `enable`, `disable`.
	DataEthernetIi *string `pulumi:"dataEthernetIi"`
	// Multicast IP address for AP discovery (default = 244.0.1.140).
	DiscoveryMcAddr *string `pulumi:"discoveryMcAddr"`
	// Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
	FiappEthType *int `pulumi:"fiappEthType"`
	// Enable/disable WTP image download at join time. Valid values: `enable`, `disable`.
	ImageDownload *string `pulumi:"imageDownload"`
	// Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
	IpsecBaseIp *string `pulumi:"ipsecBaseIp"`
	// Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable). Valid values: `enable`, `disable`.
	LinkAggregation *string `pulumi:"linkAggregation"`
	// Description of the location of the wireless controller.
	Location *string `pulumi:"location"`
	// Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
	MaxClients *int `pulumi:"maxClients"`
	// Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
	MaxRetransmit *int `pulumi:"maxRetransmit"`
	// Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
	MeshEthType *int `pulumi:"meshEthType"`
	// Interval in seconds between two WiFi network access control (NAC) checks (10 - 600, default = 120).
	NacInterval *int `pulumi:"nacInterval"`
	// Name of the wireless controller.
	Name *string `pulumi:"name"`
	// Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
	RogueScanMacAdjacency *int `pulumi:"rogueScanMacAdjacency"`
	// Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
	TunnelMode *string `pulumi:"tunnelMode"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable sharing of WTPs between VDOMs. Valid values: `enable`, `disable`.
	WtpShare *string `pulumi:"wtpShare"`
}

// The set of arguments for constructing a WirelesscontrollerGlobal resource.
type WirelesscontrollerGlobalArgs struct {
	// Enable/disable configuring APs or FortiAPs to send log messages to a syslog server (default = disable). Valid values: `enable`, `disable`.
	ApLogServer pulumi.StringPtrInput
	// IP address that APs or FortiAPs send log messages to.
	ApLogServerIp pulumi.StringPtrInput
	// Port that APs or FortiAPs send log messages to.
	ApLogServerPort pulumi.IntPtrInput
	// Configure CAPWAP control message data channel offload.
	ControlMessageOffload pulumi.StringPtrInput
	// Configure the wireless controller to use Ethernet II or 802.3 frames with 802.3 data tunnel mode (default = disable). Valid values: `enable`, `disable`.
	DataEthernetIi pulumi.StringPtrInput
	// Multicast IP address for AP discovery (default = 244.0.1.140).
	DiscoveryMcAddr pulumi.StringPtrInput
	// Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
	FiappEthType pulumi.IntPtrInput
	// Enable/disable WTP image download at join time. Valid values: `enable`, `disable`.
	ImageDownload pulumi.StringPtrInput
	// Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
	IpsecBaseIp pulumi.StringPtrInput
	// Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable). Valid values: `enable`, `disable`.
	LinkAggregation pulumi.StringPtrInput
	// Description of the location of the wireless controller.
	Location pulumi.StringPtrInput
	// Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
	MaxClients pulumi.IntPtrInput
	// Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
	MaxRetransmit pulumi.IntPtrInput
	// Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
	MeshEthType pulumi.IntPtrInput
	// Interval in seconds between two WiFi network access control (NAC) checks (10 - 600, default = 120).
	NacInterval pulumi.IntPtrInput
	// Name of the wireless controller.
	Name pulumi.StringPtrInput
	// Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
	RogueScanMacAdjacency pulumi.IntPtrInput
	// Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
	TunnelMode pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable sharing of WTPs between VDOMs. Valid values: `enable`, `disable`.
	WtpShare pulumi.StringPtrInput
}

func (WirelesscontrollerGlobalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerGlobalArgs)(nil)).Elem()
}

type WirelesscontrollerGlobalInput interface {
	pulumi.Input

	ToWirelesscontrollerGlobalOutput() WirelesscontrollerGlobalOutput
	ToWirelesscontrollerGlobalOutputWithContext(ctx context.Context) WirelesscontrollerGlobalOutput
}

func (*WirelesscontrollerGlobal) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelesscontrollerGlobal)(nil)).Elem()
}

func (i *WirelesscontrollerGlobal) ToWirelesscontrollerGlobalOutput() WirelesscontrollerGlobalOutput {
	return i.ToWirelesscontrollerGlobalOutputWithContext(context.Background())
}

func (i *WirelesscontrollerGlobal) ToWirelesscontrollerGlobalOutputWithContext(ctx context.Context) WirelesscontrollerGlobalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerGlobalOutput)
}

// WirelesscontrollerGlobalArrayInput is an input type that accepts WirelesscontrollerGlobalArray and WirelesscontrollerGlobalArrayOutput values.
// You can construct a concrete instance of `WirelesscontrollerGlobalArrayInput` via:
//
//	WirelesscontrollerGlobalArray{ WirelesscontrollerGlobalArgs{...} }
type WirelesscontrollerGlobalArrayInput interface {
	pulumi.Input

	ToWirelesscontrollerGlobalArrayOutput() WirelesscontrollerGlobalArrayOutput
	ToWirelesscontrollerGlobalArrayOutputWithContext(context.Context) WirelesscontrollerGlobalArrayOutput
}

type WirelesscontrollerGlobalArray []WirelesscontrollerGlobalInput

func (WirelesscontrollerGlobalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelesscontrollerGlobal)(nil)).Elem()
}

func (i WirelesscontrollerGlobalArray) ToWirelesscontrollerGlobalArrayOutput() WirelesscontrollerGlobalArrayOutput {
	return i.ToWirelesscontrollerGlobalArrayOutputWithContext(context.Background())
}

func (i WirelesscontrollerGlobalArray) ToWirelesscontrollerGlobalArrayOutputWithContext(ctx context.Context) WirelesscontrollerGlobalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerGlobalArrayOutput)
}

// WirelesscontrollerGlobalMapInput is an input type that accepts WirelesscontrollerGlobalMap and WirelesscontrollerGlobalMapOutput values.
// You can construct a concrete instance of `WirelesscontrollerGlobalMapInput` via:
//
//	WirelesscontrollerGlobalMap{ "key": WirelesscontrollerGlobalArgs{...} }
type WirelesscontrollerGlobalMapInput interface {
	pulumi.Input

	ToWirelesscontrollerGlobalMapOutput() WirelesscontrollerGlobalMapOutput
	ToWirelesscontrollerGlobalMapOutputWithContext(context.Context) WirelesscontrollerGlobalMapOutput
}

type WirelesscontrollerGlobalMap map[string]WirelesscontrollerGlobalInput

func (WirelesscontrollerGlobalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelesscontrollerGlobal)(nil)).Elem()
}

func (i WirelesscontrollerGlobalMap) ToWirelesscontrollerGlobalMapOutput() WirelesscontrollerGlobalMapOutput {
	return i.ToWirelesscontrollerGlobalMapOutputWithContext(context.Background())
}

func (i WirelesscontrollerGlobalMap) ToWirelesscontrollerGlobalMapOutputWithContext(ctx context.Context) WirelesscontrollerGlobalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerGlobalMapOutput)
}

type WirelesscontrollerGlobalOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerGlobalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelesscontrollerGlobal)(nil)).Elem()
}

func (o WirelesscontrollerGlobalOutput) ToWirelesscontrollerGlobalOutput() WirelesscontrollerGlobalOutput {
	return o
}

func (o WirelesscontrollerGlobalOutput) ToWirelesscontrollerGlobalOutputWithContext(ctx context.Context) WirelesscontrollerGlobalOutput {
	return o
}

// Enable/disable configuring APs or FortiAPs to send log messages to a syslog server (default = disable). Valid values: `enable`, `disable`.
func (o WirelesscontrollerGlobalOutput) ApLogServer() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.StringOutput { return v.ApLogServer }).(pulumi.StringOutput)
}

// IP address that APs or FortiAPs send log messages to.
func (o WirelesscontrollerGlobalOutput) ApLogServerIp() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.StringOutput { return v.ApLogServerIp }).(pulumi.StringOutput)
}

// Port that APs or FortiAPs send log messages to.
func (o WirelesscontrollerGlobalOutput) ApLogServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.IntOutput { return v.ApLogServerPort }).(pulumi.IntOutput)
}

// Configure CAPWAP control message data channel offload.
func (o WirelesscontrollerGlobalOutput) ControlMessageOffload() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.StringOutput { return v.ControlMessageOffload }).(pulumi.StringOutput)
}

// Configure the wireless controller to use Ethernet II or 802.3 frames with 802.3 data tunnel mode (default = disable). Valid values: `enable`, `disable`.
func (o WirelesscontrollerGlobalOutput) DataEthernetIi() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.StringOutput { return v.DataEthernetIi }).(pulumi.StringOutput)
}

// Multicast IP address for AP discovery (default = 244.0.1.140).
func (o WirelesscontrollerGlobalOutput) DiscoveryMcAddr() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.StringOutput { return v.DiscoveryMcAddr }).(pulumi.StringOutput)
}

// Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
func (o WirelesscontrollerGlobalOutput) FiappEthType() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.IntOutput { return v.FiappEthType }).(pulumi.IntOutput)
}

// Enable/disable WTP image download at join time. Valid values: `enable`, `disable`.
func (o WirelesscontrollerGlobalOutput) ImageDownload() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.StringOutput { return v.ImageDownload }).(pulumi.StringOutput)
}

// Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
func (o WirelesscontrollerGlobalOutput) IpsecBaseIp() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.StringOutput { return v.IpsecBaseIp }).(pulumi.StringOutput)
}

// Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable). Valid values: `enable`, `disable`.
func (o WirelesscontrollerGlobalOutput) LinkAggregation() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.StringOutput { return v.LinkAggregation }).(pulumi.StringOutput)
}

// Description of the location of the wireless controller.
func (o WirelesscontrollerGlobalOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
func (o WirelesscontrollerGlobalOutput) MaxClients() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.IntOutput { return v.MaxClients }).(pulumi.IntOutput)
}

// Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
func (o WirelesscontrollerGlobalOutput) MaxRetransmit() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.IntOutput { return v.MaxRetransmit }).(pulumi.IntOutput)
}

// Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
func (o WirelesscontrollerGlobalOutput) MeshEthType() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.IntOutput { return v.MeshEthType }).(pulumi.IntOutput)
}

// Interval in seconds between two WiFi network access control (NAC) checks (10 - 600, default = 120).
func (o WirelesscontrollerGlobalOutput) NacInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.IntOutput { return v.NacInterval }).(pulumi.IntOutput)
}

// Name of the wireless controller.
func (o WirelesscontrollerGlobalOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
func (o WirelesscontrollerGlobalOutput) RogueScanMacAdjacency() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.IntOutput { return v.RogueScanMacAdjacency }).(pulumi.IntOutput)
}

// Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
func (o WirelesscontrollerGlobalOutput) TunnelMode() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.StringOutput { return v.TunnelMode }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o WirelesscontrollerGlobalOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable sharing of WTPs between VDOMs. Valid values: `enable`, `disable`.
func (o WirelesscontrollerGlobalOutput) WtpShare() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerGlobal) pulumi.StringOutput { return v.WtpShare }).(pulumi.StringOutput)
}

type WirelesscontrollerGlobalArrayOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerGlobalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelesscontrollerGlobal)(nil)).Elem()
}

func (o WirelesscontrollerGlobalArrayOutput) ToWirelesscontrollerGlobalArrayOutput() WirelesscontrollerGlobalArrayOutput {
	return o
}

func (o WirelesscontrollerGlobalArrayOutput) ToWirelesscontrollerGlobalArrayOutputWithContext(ctx context.Context) WirelesscontrollerGlobalArrayOutput {
	return o
}

func (o WirelesscontrollerGlobalArrayOutput) Index(i pulumi.IntInput) WirelesscontrollerGlobalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelesscontrollerGlobal {
		return vs[0].([]*WirelesscontrollerGlobal)[vs[1].(int)]
	}).(WirelesscontrollerGlobalOutput)
}

type WirelesscontrollerGlobalMapOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerGlobalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelesscontrollerGlobal)(nil)).Elem()
}

func (o WirelesscontrollerGlobalMapOutput) ToWirelesscontrollerGlobalMapOutput() WirelesscontrollerGlobalMapOutput {
	return o
}

func (o WirelesscontrollerGlobalMapOutput) ToWirelesscontrollerGlobalMapOutputWithContext(ctx context.Context) WirelesscontrollerGlobalMapOutput {
	return o
}

func (o WirelesscontrollerGlobalMapOutput) MapIndex(k pulumi.StringInput) WirelesscontrollerGlobalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelesscontrollerGlobal {
		return vs[0].(map[string]*WirelesscontrollerGlobal)[vs[1].(string)]
	}).(WirelesscontrollerGlobalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerGlobalInput)(nil)).Elem(), &WirelesscontrollerGlobal{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerGlobalArrayInput)(nil)).Elem(), WirelesscontrollerGlobalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerGlobalMapInput)(nil)).Elem(), WirelesscontrollerGlobalMap{})
	pulumi.RegisterOutputType(WirelesscontrollerGlobalOutput{})
	pulumi.RegisterOutputType(WirelesscontrollerGlobalArrayOutput{})
	pulumi.RegisterOutputType(WirelesscontrollerGlobalMapOutput{})
}
