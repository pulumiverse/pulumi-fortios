// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lldp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure LLDP network policy.
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
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := system.NewNetworkpolicy(ctx, "trname", &system.NetworkpolicyArgs{
//				Comment: pulumi.String("test"),
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
// SystemLldp NetworkPolicy can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/lldp/networkpolicy:Networkpolicy labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/lldp/networkpolicy:Networkpolicy labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Networkpolicy struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Guest. The structure of `guest` block is documented below.
	Guest NetworkpolicyGuestOutput `pulumi:"guest"`
	// Guest Voice Signaling. The structure of `guestVoiceSignaling` block is documented below.
	GuestVoiceSignaling NetworkpolicyGuestVoiceSignalingOutput `pulumi:"guestVoiceSignaling"`
	// LLDP network policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Softphone. The structure of `softphone` block is documented below.
	Softphone NetworkpolicySoftphoneOutput `pulumi:"softphone"`
	// Streaming Video. The structure of `streamingVideo` block is documented below.
	StreamingVideo NetworkpolicyStreamingVideoOutput `pulumi:"streamingVideo"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Video Conferencing. The structure of `videoConferencing` block is documented below.
	VideoConferencing NetworkpolicyVideoConferencingOutput `pulumi:"videoConferencing"`
	// Video Signaling. The structure of `videoSignaling` block is documented below.
	VideoSignaling NetworkpolicyVideoSignalingOutput `pulumi:"videoSignaling"`
	// Voice. The structure of `voice` block is documented below.
	Voice NetworkpolicyVoiceOutput `pulumi:"voice"`
	// Voice signaling. The structure of `voiceSignaling` block is documented below.
	VoiceSignaling NetworkpolicyVoiceSignalingOutput `pulumi:"voiceSignaling"`
}

// NewNetworkpolicy registers a new resource with the given unique name, arguments, and options.
func NewNetworkpolicy(ctx *pulumi.Context,
	name string, args *NetworkpolicyArgs, opts ...pulumi.ResourceOption) (*Networkpolicy, error) {
	if args == nil {
		args = &NetworkpolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Networkpolicy
	err := ctx.RegisterResource("fortios:system/lldp/networkpolicy:Networkpolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkpolicy gets an existing Networkpolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkpolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkpolicyState, opts ...pulumi.ResourceOption) (*Networkpolicy, error) {
	var resource Networkpolicy
	err := ctx.ReadResource("fortios:system/lldp/networkpolicy:Networkpolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Networkpolicy resources.
type networkpolicyState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Guest. The structure of `guest` block is documented below.
	Guest *NetworkpolicyGuest `pulumi:"guest"`
	// Guest Voice Signaling. The structure of `guestVoiceSignaling` block is documented below.
	GuestVoiceSignaling *NetworkpolicyGuestVoiceSignaling `pulumi:"guestVoiceSignaling"`
	// LLDP network policy name.
	Name *string `pulumi:"name"`
	// Softphone. The structure of `softphone` block is documented below.
	Softphone *NetworkpolicySoftphone `pulumi:"softphone"`
	// Streaming Video. The structure of `streamingVideo` block is documented below.
	StreamingVideo *NetworkpolicyStreamingVideo `pulumi:"streamingVideo"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Video Conferencing. The structure of `videoConferencing` block is documented below.
	VideoConferencing *NetworkpolicyVideoConferencing `pulumi:"videoConferencing"`
	// Video Signaling. The structure of `videoSignaling` block is documented below.
	VideoSignaling *NetworkpolicyVideoSignaling `pulumi:"videoSignaling"`
	// Voice. The structure of `voice` block is documented below.
	Voice *NetworkpolicyVoice `pulumi:"voice"`
	// Voice signaling. The structure of `voiceSignaling` block is documented below.
	VoiceSignaling *NetworkpolicyVoiceSignaling `pulumi:"voiceSignaling"`
}

type NetworkpolicyState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Guest. The structure of `guest` block is documented below.
	Guest NetworkpolicyGuestPtrInput
	// Guest Voice Signaling. The structure of `guestVoiceSignaling` block is documented below.
	GuestVoiceSignaling NetworkpolicyGuestVoiceSignalingPtrInput
	// LLDP network policy name.
	Name pulumi.StringPtrInput
	// Softphone. The structure of `softphone` block is documented below.
	Softphone NetworkpolicySoftphonePtrInput
	// Streaming Video. The structure of `streamingVideo` block is documented below.
	StreamingVideo NetworkpolicyStreamingVideoPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Video Conferencing. The structure of `videoConferencing` block is documented below.
	VideoConferencing NetworkpolicyVideoConferencingPtrInput
	// Video Signaling. The structure of `videoSignaling` block is documented below.
	VideoSignaling NetworkpolicyVideoSignalingPtrInput
	// Voice. The structure of `voice` block is documented below.
	Voice NetworkpolicyVoicePtrInput
	// Voice signaling. The structure of `voiceSignaling` block is documented below.
	VoiceSignaling NetworkpolicyVoiceSignalingPtrInput
}

func (NetworkpolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkpolicyState)(nil)).Elem()
}

type networkpolicyArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Guest. The structure of `guest` block is documented below.
	Guest *NetworkpolicyGuest `pulumi:"guest"`
	// Guest Voice Signaling. The structure of `guestVoiceSignaling` block is documented below.
	GuestVoiceSignaling *NetworkpolicyGuestVoiceSignaling `pulumi:"guestVoiceSignaling"`
	// LLDP network policy name.
	Name *string `pulumi:"name"`
	// Softphone. The structure of `softphone` block is documented below.
	Softphone *NetworkpolicySoftphone `pulumi:"softphone"`
	// Streaming Video. The structure of `streamingVideo` block is documented below.
	StreamingVideo *NetworkpolicyStreamingVideo `pulumi:"streamingVideo"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Video Conferencing. The structure of `videoConferencing` block is documented below.
	VideoConferencing *NetworkpolicyVideoConferencing `pulumi:"videoConferencing"`
	// Video Signaling. The structure of `videoSignaling` block is documented below.
	VideoSignaling *NetworkpolicyVideoSignaling `pulumi:"videoSignaling"`
	// Voice. The structure of `voice` block is documented below.
	Voice *NetworkpolicyVoice `pulumi:"voice"`
	// Voice signaling. The structure of `voiceSignaling` block is documented below.
	VoiceSignaling *NetworkpolicyVoiceSignaling `pulumi:"voiceSignaling"`
}

// The set of arguments for constructing a Networkpolicy resource.
type NetworkpolicyArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Guest. The structure of `guest` block is documented below.
	Guest NetworkpolicyGuestPtrInput
	// Guest Voice Signaling. The structure of `guestVoiceSignaling` block is documented below.
	GuestVoiceSignaling NetworkpolicyGuestVoiceSignalingPtrInput
	// LLDP network policy name.
	Name pulumi.StringPtrInput
	// Softphone. The structure of `softphone` block is documented below.
	Softphone NetworkpolicySoftphonePtrInput
	// Streaming Video. The structure of `streamingVideo` block is documented below.
	StreamingVideo NetworkpolicyStreamingVideoPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Video Conferencing. The structure of `videoConferencing` block is documented below.
	VideoConferencing NetworkpolicyVideoConferencingPtrInput
	// Video Signaling. The structure of `videoSignaling` block is documented below.
	VideoSignaling NetworkpolicyVideoSignalingPtrInput
	// Voice. The structure of `voice` block is documented below.
	Voice NetworkpolicyVoicePtrInput
	// Voice signaling. The structure of `voiceSignaling` block is documented below.
	VoiceSignaling NetworkpolicyVoiceSignalingPtrInput
}

func (NetworkpolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkpolicyArgs)(nil)).Elem()
}

type NetworkpolicyInput interface {
	pulumi.Input

	ToNetworkpolicyOutput() NetworkpolicyOutput
	ToNetworkpolicyOutputWithContext(ctx context.Context) NetworkpolicyOutput
}

func (*Networkpolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**Networkpolicy)(nil)).Elem()
}

func (i *Networkpolicy) ToNetworkpolicyOutput() NetworkpolicyOutput {
	return i.ToNetworkpolicyOutputWithContext(context.Background())
}

func (i *Networkpolicy) ToNetworkpolicyOutputWithContext(ctx context.Context) NetworkpolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkpolicyOutput)
}

// NetworkpolicyArrayInput is an input type that accepts NetworkpolicyArray and NetworkpolicyArrayOutput values.
// You can construct a concrete instance of `NetworkpolicyArrayInput` via:
//
//	NetworkpolicyArray{ NetworkpolicyArgs{...} }
type NetworkpolicyArrayInput interface {
	pulumi.Input

	ToNetworkpolicyArrayOutput() NetworkpolicyArrayOutput
	ToNetworkpolicyArrayOutputWithContext(context.Context) NetworkpolicyArrayOutput
}

type NetworkpolicyArray []NetworkpolicyInput

func (NetworkpolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Networkpolicy)(nil)).Elem()
}

func (i NetworkpolicyArray) ToNetworkpolicyArrayOutput() NetworkpolicyArrayOutput {
	return i.ToNetworkpolicyArrayOutputWithContext(context.Background())
}

func (i NetworkpolicyArray) ToNetworkpolicyArrayOutputWithContext(ctx context.Context) NetworkpolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkpolicyArrayOutput)
}

// NetworkpolicyMapInput is an input type that accepts NetworkpolicyMap and NetworkpolicyMapOutput values.
// You can construct a concrete instance of `NetworkpolicyMapInput` via:
//
//	NetworkpolicyMap{ "key": NetworkpolicyArgs{...} }
type NetworkpolicyMapInput interface {
	pulumi.Input

	ToNetworkpolicyMapOutput() NetworkpolicyMapOutput
	ToNetworkpolicyMapOutputWithContext(context.Context) NetworkpolicyMapOutput
}

type NetworkpolicyMap map[string]NetworkpolicyInput

func (NetworkpolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Networkpolicy)(nil)).Elem()
}

func (i NetworkpolicyMap) ToNetworkpolicyMapOutput() NetworkpolicyMapOutput {
	return i.ToNetworkpolicyMapOutputWithContext(context.Background())
}

func (i NetworkpolicyMap) ToNetworkpolicyMapOutputWithContext(ctx context.Context) NetworkpolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkpolicyMapOutput)
}

type NetworkpolicyOutput struct{ *pulumi.OutputState }

func (NetworkpolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Networkpolicy)(nil)).Elem()
}

func (o NetworkpolicyOutput) ToNetworkpolicyOutput() NetworkpolicyOutput {
	return o
}

func (o NetworkpolicyOutput) ToNetworkpolicyOutputWithContext(ctx context.Context) NetworkpolicyOutput {
	return o
}

// Comment.
func (o NetworkpolicyOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Networkpolicy) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o NetworkpolicyOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Networkpolicy) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Guest. The structure of `guest` block is documented below.
func (o NetworkpolicyOutput) Guest() NetworkpolicyGuestOutput {
	return o.ApplyT(func(v *Networkpolicy) NetworkpolicyGuestOutput { return v.Guest }).(NetworkpolicyGuestOutput)
}

// Guest Voice Signaling. The structure of `guestVoiceSignaling` block is documented below.
func (o NetworkpolicyOutput) GuestVoiceSignaling() NetworkpolicyGuestVoiceSignalingOutput {
	return o.ApplyT(func(v *Networkpolicy) NetworkpolicyGuestVoiceSignalingOutput { return v.GuestVoiceSignaling }).(NetworkpolicyGuestVoiceSignalingOutput)
}

// LLDP network policy name.
func (o NetworkpolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Networkpolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Softphone. The structure of `softphone` block is documented below.
func (o NetworkpolicyOutput) Softphone() NetworkpolicySoftphoneOutput {
	return o.ApplyT(func(v *Networkpolicy) NetworkpolicySoftphoneOutput { return v.Softphone }).(NetworkpolicySoftphoneOutput)
}

// Streaming Video. The structure of `streamingVideo` block is documented below.
func (o NetworkpolicyOutput) StreamingVideo() NetworkpolicyStreamingVideoOutput {
	return o.ApplyT(func(v *Networkpolicy) NetworkpolicyStreamingVideoOutput { return v.StreamingVideo }).(NetworkpolicyStreamingVideoOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o NetworkpolicyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Networkpolicy) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Video Conferencing. The structure of `videoConferencing` block is documented below.
func (o NetworkpolicyOutput) VideoConferencing() NetworkpolicyVideoConferencingOutput {
	return o.ApplyT(func(v *Networkpolicy) NetworkpolicyVideoConferencingOutput { return v.VideoConferencing }).(NetworkpolicyVideoConferencingOutput)
}

// Video Signaling. The structure of `videoSignaling` block is documented below.
func (o NetworkpolicyOutput) VideoSignaling() NetworkpolicyVideoSignalingOutput {
	return o.ApplyT(func(v *Networkpolicy) NetworkpolicyVideoSignalingOutput { return v.VideoSignaling }).(NetworkpolicyVideoSignalingOutput)
}

// Voice. The structure of `voice` block is documented below.
func (o NetworkpolicyOutput) Voice() NetworkpolicyVoiceOutput {
	return o.ApplyT(func(v *Networkpolicy) NetworkpolicyVoiceOutput { return v.Voice }).(NetworkpolicyVoiceOutput)
}

// Voice signaling. The structure of `voiceSignaling` block is documented below.
func (o NetworkpolicyOutput) VoiceSignaling() NetworkpolicyVoiceSignalingOutput {
	return o.ApplyT(func(v *Networkpolicy) NetworkpolicyVoiceSignalingOutput { return v.VoiceSignaling }).(NetworkpolicyVoiceSignalingOutput)
}

type NetworkpolicyArrayOutput struct{ *pulumi.OutputState }

func (NetworkpolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Networkpolicy)(nil)).Elem()
}

func (o NetworkpolicyArrayOutput) ToNetworkpolicyArrayOutput() NetworkpolicyArrayOutput {
	return o
}

func (o NetworkpolicyArrayOutput) ToNetworkpolicyArrayOutputWithContext(ctx context.Context) NetworkpolicyArrayOutput {
	return o
}

func (o NetworkpolicyArrayOutput) Index(i pulumi.IntInput) NetworkpolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Networkpolicy {
		return vs[0].([]*Networkpolicy)[vs[1].(int)]
	}).(NetworkpolicyOutput)
}

type NetworkpolicyMapOutput struct{ *pulumi.OutputState }

func (NetworkpolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Networkpolicy)(nil)).Elem()
}

func (o NetworkpolicyMapOutput) ToNetworkpolicyMapOutput() NetworkpolicyMapOutput {
	return o
}

func (o NetworkpolicyMapOutput) ToNetworkpolicyMapOutputWithContext(ctx context.Context) NetworkpolicyMapOutput {
	return o
}

func (o NetworkpolicyMapOutput) MapIndex(k pulumi.StringInput) NetworkpolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Networkpolicy {
		return vs[0].(map[string]*Networkpolicy)[vs[1].(string)]
	}).(NetworkpolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkpolicyInput)(nil)).Elem(), &Networkpolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkpolicyArrayInput)(nil)).Elem(), NetworkpolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkpolicyMapInput)(nil)).Elem(), NetworkpolicyMap{})
	pulumi.RegisterOutputType(NetworkpolicyOutput{})
	pulumi.RegisterOutputType(NetworkpolicyArrayOutput{})
	pulumi.RegisterOutputType(NetworkpolicyMapOutput{})
}
