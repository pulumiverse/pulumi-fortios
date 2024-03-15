// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure an aggregate of IPsec tunnels.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/system"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/vpn"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			trname1Phase1interface, err := vpn.NewPhase1interface(ctx, "trname1Phase1interface", &vpn.Phase1interfaceArgs{
//				AcctVerify:             pulumi.String("disable"),
//				AddGwRoute:             pulumi.String("disable"),
//				AddRoute:               pulumi.String("enable"),
//				AssignIp:               pulumi.String("enable"),
//				AssignIpFrom:           pulumi.String("range"),
//				Authmethod:             pulumi.String("psk"),
//				AutoDiscoveryForwarder: pulumi.String("disable"),
//				AutoDiscoveryPsk:       pulumi.String("disable"),
//				AutoDiscoveryReceiver:  pulumi.String("disable"),
//				AutoDiscoverySender:    pulumi.String("disable"),
//				AutoNegotiate:          pulumi.String("enable"),
//				CertIdValidation:       pulumi.String("enable"),
//				ChildlessIke:           pulumi.String("disable"),
//				ClientAutoNegotiate:    pulumi.String("disable"),
//				ClientKeepAlive:        pulumi.String("disable"),
//				DefaultGw:              pulumi.String("0.0.0.0"),
//				DefaultGwPriority:      pulumi.Int(0),
//				Dhgrp:                  pulumi.String("14 5"),
//				DigitalSignatureAuth:   pulumi.String("disable"),
//				Distance:               pulumi.Int(15),
//				DnsMode:                pulumi.String("manual"),
//				Dpd:                    pulumi.String("on-demand"),
//				DpdRetrycount:          pulumi.Int(3),
//				DpdRetryinterval:       pulumi.String("20"),
//				Eap:                    pulumi.String("disable"),
//				EapIdentity:            pulumi.String("use-id-payload"),
//				EncapLocalGw4:          pulumi.String("0.0.0.0"),
//				EncapLocalGw6:          pulumi.String("::"),
//				EncapRemoteGw4:         pulumi.String("0.0.0.0"),
//				EncapRemoteGw6:         pulumi.String("::"),
//				Encapsulation:          pulumi.String("none"),
//				EncapsulationAddress:   pulumi.String("ike"),
//				EnforceUniqueId:        pulumi.String("disable"),
//				ExchangeInterfaceIp:    pulumi.String("disable"),
//				ExchangeIpAddr4:        pulumi.String("0.0.0.0"),
//				ExchangeIpAddr6:        pulumi.String("::"),
//				ForticlientEnforcement: pulumi.String("disable"),
//				Fragmentation:          pulumi.String("enable"),
//				FragmentationMtu:       pulumi.Int(1200),
//				GroupAuthentication:    pulumi.String("disable"),
//				HaSyncEspSeqno:         pulumi.String("enable"),
//				IdleTimeout:            pulumi.String("disable"),
//				IdleTimeoutinterval:    pulumi.Int(15),
//				IkeVersion:             pulumi.String("1"),
//				IncludeLocalLan:        pulumi.String("disable"),
//				Interface:              pulumi.String("port3"),
//				IpVersion:              pulumi.String("4"),
//				Ipv4DnsServer1:         pulumi.String("0.0.0.0"),
//				Ipv4DnsServer2:         pulumi.String("0.0.0.0"),
//				Ipv4DnsServer3:         pulumi.String("0.0.0.0"),
//				Ipv4EndIp:              pulumi.String("0.0.0.0"),
//				Ipv4Netmask:            pulumi.String("255.255.255.255"),
//				Ipv4StartIp:            pulumi.String("0.0.0.0"),
//				Ipv4WinsServer1:        pulumi.String("0.0.0.0"),
//				Ipv4WinsServer2:        pulumi.String("0.0.0.0"),
//				Ipv6DnsServer1:         pulumi.String("::"),
//				Ipv6DnsServer2:         pulumi.String("::"),
//				Ipv6DnsServer3:         pulumi.String("::"),
//				Ipv6EndIp:              pulumi.String("::"),
//				Ipv6Prefix:             pulumi.Int(128),
//				Ipv6StartIp:            pulumi.String("::"),
//				Keepalive:              pulumi.Int(10),
//				Keylife:                pulumi.Int(86400),
//				LocalGw:                pulumi.String("0.0.0.0"),
//				LocalGw6:               pulumi.String("::"),
//				LocalidType:            pulumi.String("auto"),
//				MeshSelectorType:       pulumi.String("disable"),
//				Mode:                   pulumi.String("main"),
//				ModeCfg:                pulumi.String("disable"),
//				MonitorHoldDownDelay:   pulumi.Int(0),
//				MonitorHoldDownTime:    pulumi.String("00:00"),
//				MonitorHoldDownType:    pulumi.String("immediate"),
//				MonitorHoldDownWeekday: pulumi.String("sunday"),
//				Nattraversal:           pulumi.String("enable"),
//				NegotiateTimeout:       pulumi.Int(30),
//				NetDevice:              pulumi.String("disable"),
//				PassiveMode:            pulumi.String("disable"),
//				Peertype:               pulumi.String("any"),
//				Ppk:                    pulumi.String("disable"),
//				Priority:               pulumi.Int(0),
//				Proposal:               pulumi.String("aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"),
//				Psksecret:              pulumi.String("eweeeeeeeecee"),
//				Reauth:                 pulumi.String("disable"),
//				Rekey:                  pulumi.String("enable"),
//				RemoteGw:               pulumi.String("2.2.2.2"),
//				RemoteGw6:              pulumi.String("::"),
//				RsaSignatureFormat:     pulumi.String("pkcs1"),
//				SavePassword:           pulumi.String("disable"),
//				SendCertChain:          pulumi.String("enable"),
//				SignatureHashAlg:       pulumi.String("sha2-512 sha2-384 sha2-256 sha1"),
//				SuiteB:                 pulumi.String("disable"),
//				TunnelSearch:           pulumi.String("selectors"),
//				Type:                   pulumi.String("static"),
//				UnitySupport:           pulumi.String("enable"),
//				WizardType:             pulumi.String("custom"),
//				Xauthtype:              pulumi.String("disable"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpn.NewPhase2interface(ctx, "trname1Phase2interface", &vpn.Phase2interfaceArgs{
//				AddRoute:               pulumi.String("phase1"),
//				AutoDiscoveryForwarder: pulumi.String("phase1"),
//				AutoDiscoverySender:    pulumi.String("phase1"),
//				AutoNegotiate:          pulumi.String("disable"),
//				DhcpIpsec:              pulumi.String("disable"),
//				Dhgrp:                  pulumi.String("14 5"),
//				DstAddrType:            pulumi.String("subnet"),
//				DstEndIp6:              pulumi.String("::"),
//				DstPort:                pulumi.Int(0),
//				DstSubnet:              pulumi.String("0.0.0.0 0.0.0.0"),
//				Encapsulation:          pulumi.String("tunnel-mode"),
//				Keepalive:              pulumi.String("disable"),
//				KeylifeType:            pulumi.String("seconds"),
//				Keylifekbs:             pulumi.Int(5120),
//				Keylifeseconds:         pulumi.Int(43200),
//				L2tp:                   pulumi.String("disable"),
//				Pfs:                    pulumi.String("enable"),
//				Phase1name:             trname1Phase1interface.Name,
//				Proposal:               pulumi.String("aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305"),
//				Protocol:               pulumi.Int(0),
//				Replay:                 pulumi.String("enable"),
//				RouteOverlap:           pulumi.String("use-new"),
//				SingleSource:           pulumi.String("disable"),
//				SrcAddrType:            pulumi.String("subnet"),
//				SrcEndIp6:              pulumi.String("::"),
//				SrcPort:                pulumi.Int(0),
//				SrcSubnet:              pulumi.String("0.0.0.0 0.0.0.0"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = system.NewIpsecaggregate(ctx, "trname", &system.IpsecaggregateArgs{
//				Algorithm: pulumi.String("round-robin"),
//				Members: system.IpsecaggregateMemberArray{
//					&system.IpsecaggregateMemberArgs{
//						TunnelName: trname1Phase1interface.Name,
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// System IpsecAggregate can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/ipsecaggregate:Ipsecaggregate labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/ipsecaggregate:Ipsecaggregate labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Ipsecaggregate struct {
	pulumi.CustomResourceState

	// Frame distribution algorithm.
	Algorithm pulumi.StringOutput `pulumi:"algorithm"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Member tunnels of the aggregate. The structure of `member` block is documented below.
	Members IpsecaggregateMemberArrayOutput `pulumi:"members"`
	// IPsec aggregate name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewIpsecaggregate registers a new resource with the given unique name, arguments, and options.
func NewIpsecaggregate(ctx *pulumi.Context,
	name string, args *IpsecaggregateArgs, opts ...pulumi.ResourceOption) (*Ipsecaggregate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ipsecaggregate
	err := ctx.RegisterResource("fortios:system/ipsecaggregate:Ipsecaggregate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpsecaggregate gets an existing Ipsecaggregate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpsecaggregate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpsecaggregateState, opts ...pulumi.ResourceOption) (*Ipsecaggregate, error) {
	var resource Ipsecaggregate
	err := ctx.ReadResource("fortios:system/ipsecaggregate:Ipsecaggregate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ipsecaggregate resources.
type ipsecaggregateState struct {
	// Frame distribution algorithm.
	Algorithm *string `pulumi:"algorithm"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Member tunnels of the aggregate. The structure of `member` block is documented below.
	Members []IpsecaggregateMember `pulumi:"members"`
	// IPsec aggregate name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type IpsecaggregateState struct {
	// Frame distribution algorithm.
	Algorithm pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Member tunnels of the aggregate. The structure of `member` block is documented below.
	Members IpsecaggregateMemberArrayInput
	// IPsec aggregate name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IpsecaggregateState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsecaggregateState)(nil)).Elem()
}

type ipsecaggregateArgs struct {
	// Frame distribution algorithm.
	Algorithm *string `pulumi:"algorithm"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Member tunnels of the aggregate. The structure of `member` block is documented below.
	Members []IpsecaggregateMember `pulumi:"members"`
	// IPsec aggregate name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Ipsecaggregate resource.
type IpsecaggregateArgs struct {
	// Frame distribution algorithm.
	Algorithm pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Member tunnels of the aggregate. The structure of `member` block is documented below.
	Members IpsecaggregateMemberArrayInput
	// IPsec aggregate name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IpsecaggregateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsecaggregateArgs)(nil)).Elem()
}

type IpsecaggregateInput interface {
	pulumi.Input

	ToIpsecaggregateOutput() IpsecaggregateOutput
	ToIpsecaggregateOutputWithContext(ctx context.Context) IpsecaggregateOutput
}

func (*Ipsecaggregate) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipsecaggregate)(nil)).Elem()
}

func (i *Ipsecaggregate) ToIpsecaggregateOutput() IpsecaggregateOutput {
	return i.ToIpsecaggregateOutputWithContext(context.Background())
}

func (i *Ipsecaggregate) ToIpsecaggregateOutputWithContext(ctx context.Context) IpsecaggregateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsecaggregateOutput)
}

// IpsecaggregateArrayInput is an input type that accepts IpsecaggregateArray and IpsecaggregateArrayOutput values.
// You can construct a concrete instance of `IpsecaggregateArrayInput` via:
//
//	IpsecaggregateArray{ IpsecaggregateArgs{...} }
type IpsecaggregateArrayInput interface {
	pulumi.Input

	ToIpsecaggregateArrayOutput() IpsecaggregateArrayOutput
	ToIpsecaggregateArrayOutputWithContext(context.Context) IpsecaggregateArrayOutput
}

type IpsecaggregateArray []IpsecaggregateInput

func (IpsecaggregateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipsecaggregate)(nil)).Elem()
}

func (i IpsecaggregateArray) ToIpsecaggregateArrayOutput() IpsecaggregateArrayOutput {
	return i.ToIpsecaggregateArrayOutputWithContext(context.Background())
}

func (i IpsecaggregateArray) ToIpsecaggregateArrayOutputWithContext(ctx context.Context) IpsecaggregateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsecaggregateArrayOutput)
}

// IpsecaggregateMapInput is an input type that accepts IpsecaggregateMap and IpsecaggregateMapOutput values.
// You can construct a concrete instance of `IpsecaggregateMapInput` via:
//
//	IpsecaggregateMap{ "key": IpsecaggregateArgs{...} }
type IpsecaggregateMapInput interface {
	pulumi.Input

	ToIpsecaggregateMapOutput() IpsecaggregateMapOutput
	ToIpsecaggregateMapOutputWithContext(context.Context) IpsecaggregateMapOutput
}

type IpsecaggregateMap map[string]IpsecaggregateInput

func (IpsecaggregateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipsecaggregate)(nil)).Elem()
}

func (i IpsecaggregateMap) ToIpsecaggregateMapOutput() IpsecaggregateMapOutput {
	return i.ToIpsecaggregateMapOutputWithContext(context.Background())
}

func (i IpsecaggregateMap) ToIpsecaggregateMapOutputWithContext(ctx context.Context) IpsecaggregateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsecaggregateMapOutput)
}

type IpsecaggregateOutput struct{ *pulumi.OutputState }

func (IpsecaggregateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipsecaggregate)(nil)).Elem()
}

func (o IpsecaggregateOutput) ToIpsecaggregateOutput() IpsecaggregateOutput {
	return o
}

func (o IpsecaggregateOutput) ToIpsecaggregateOutputWithContext(ctx context.Context) IpsecaggregateOutput {
	return o
}

// Frame distribution algorithm.
func (o IpsecaggregateOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipsecaggregate) pulumi.StringOutput { return v.Algorithm }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o IpsecaggregateOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipsecaggregate) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Member tunnels of the aggregate. The structure of `member` block is documented below.
func (o IpsecaggregateOutput) Members() IpsecaggregateMemberArrayOutput {
	return o.ApplyT(func(v *Ipsecaggregate) IpsecaggregateMemberArrayOutput { return v.Members }).(IpsecaggregateMemberArrayOutput)
}

// IPsec aggregate name.
func (o IpsecaggregateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipsecaggregate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o IpsecaggregateOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipsecaggregate) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type IpsecaggregateArrayOutput struct{ *pulumi.OutputState }

func (IpsecaggregateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipsecaggregate)(nil)).Elem()
}

func (o IpsecaggregateArrayOutput) ToIpsecaggregateArrayOutput() IpsecaggregateArrayOutput {
	return o
}

func (o IpsecaggregateArrayOutput) ToIpsecaggregateArrayOutputWithContext(ctx context.Context) IpsecaggregateArrayOutput {
	return o
}

func (o IpsecaggregateArrayOutput) Index(i pulumi.IntInput) IpsecaggregateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ipsecaggregate {
		return vs[0].([]*Ipsecaggregate)[vs[1].(int)]
	}).(IpsecaggregateOutput)
}

type IpsecaggregateMapOutput struct{ *pulumi.OutputState }

func (IpsecaggregateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipsecaggregate)(nil)).Elem()
}

func (o IpsecaggregateMapOutput) ToIpsecaggregateMapOutput() IpsecaggregateMapOutput {
	return o
}

func (o IpsecaggregateMapOutput) ToIpsecaggregateMapOutputWithContext(ctx context.Context) IpsecaggregateMapOutput {
	return o
}

func (o IpsecaggregateMapOutput) MapIndex(k pulumi.StringInput) IpsecaggregateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ipsecaggregate {
		return vs[0].(map[string]*Ipsecaggregate)[vs[1].(string)]
	}).(IpsecaggregateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpsecaggregateInput)(nil)).Elem(), &Ipsecaggregate{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsecaggregateArrayInput)(nil)).Elem(), IpsecaggregateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsecaggregateMapInput)(nil)).Elem(), IpsecaggregateMap{})
	pulumi.RegisterOutputType(IpsecaggregateOutput{})
	pulumi.RegisterOutputType(IpsecaggregateArrayOutput{})
	pulumi.RegisterOutputType(IpsecaggregateMapOutput{})
}