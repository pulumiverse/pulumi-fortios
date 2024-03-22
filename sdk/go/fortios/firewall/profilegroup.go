// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure profile groups.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/firewall"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firewall.NewProfilegroup(ctx, "trname", &firewall.ProfilegroupArgs{
//				ProfileProtocolOptions: pulumi.String("default"),
//				SslSshProfile:          pulumi.String("deep-inspection"),
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
// Firewall ProfileGroup can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/profilegroup:Profilegroup labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/profilegroup:Profilegroup labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Profilegroup struct {
	pulumi.CustomResourceState

	// Name of an existing Application list.
	ApplicationList pulumi.StringOutput `pulumi:"applicationList"`
	// Name of an existing Antivirus profile.
	AvProfile pulumi.StringOutput `pulumi:"avProfile"`
	// Name of an existing CASB profile.
	CasbProfile pulumi.StringOutput `pulumi:"casbProfile"`
	// Name of an existing CIFS profile.
	CifsProfile pulumi.StringOutput `pulumi:"cifsProfile"`
	// Name of an existing Diameter filter profile.
	DiameterFilterProfile pulumi.StringOutput `pulumi:"diameterFilterProfile"`
	// Name of an existing DLP profile.
	DlpProfile pulumi.StringOutput `pulumi:"dlpProfile"`
	// Name of an existing DLP sensor.
	DlpSensor pulumi.StringOutput `pulumi:"dlpSensor"`
	// Name of an existing DNS filter profile.
	DnsfilterProfile pulumi.StringOutput `pulumi:"dnsfilterProfile"`
	// Name of an existing email filter profile.
	EmailfilterProfile pulumi.StringOutput `pulumi:"emailfilterProfile"`
	// Name of an existing file-filter profile.
	FileFilterProfile pulumi.StringOutput `pulumi:"fileFilterProfile"`
	// Name of an existing ICAP profile.
	IcapProfile pulumi.StringOutput `pulumi:"icapProfile"`
	// Name of an existing IPS sensor.
	IpsSensor pulumi.StringOutput `pulumi:"ipsSensor"`
	// Name of an existing VoIP (ips) profile.
	IpsVoipFilter pulumi.StringOutput `pulumi:"ipsVoipFilter"`
	// Profile group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of an existing Protocol options profile.
	ProfileProtocolOptions pulumi.StringOutput `pulumi:"profileProtocolOptions"`
	// Name of an existing SCTP filter profile.
	SctpFilterProfile pulumi.StringOutput `pulumi:"sctpFilterProfile"`
	// Name of an existing Spam filter profile.
	SpamfilterProfile pulumi.StringOutput `pulumi:"spamfilterProfile"`
	// Name of an existing SSH filter profile.
	SshFilterProfile pulumi.StringOutput `pulumi:"sshFilterProfile"`
	// Name of an existing SSL SSH profile.
	SslSshProfile pulumi.StringOutput `pulumi:"sslSshProfile"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Name of an existing VideoFilter profile.
	VideofilterProfile pulumi.StringOutput `pulumi:"videofilterProfile"`
	// Name of an existing virtual-patch profile.
	VirtualPatchProfile pulumi.StringOutput `pulumi:"virtualPatchProfile"`
	// Name of an existing VoIP (voipd) profile.
	VoipProfile pulumi.StringOutput `pulumi:"voipProfile"`
	// Name of an existing Web application firewall profile.
	WafProfile pulumi.StringOutput `pulumi:"wafProfile"`
	// Name of an existing Web filter profile.
	WebfilterProfile pulumi.StringOutput `pulumi:"webfilterProfile"`
}

// NewProfilegroup registers a new resource with the given unique name, arguments, and options.
func NewProfilegroup(ctx *pulumi.Context,
	name string, args *ProfilegroupArgs, opts ...pulumi.ResourceOption) (*Profilegroup, error) {
	if args == nil {
		args = &ProfilegroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Profilegroup
	err := ctx.RegisterResource("fortios:firewall/profilegroup:Profilegroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfilegroup gets an existing Profilegroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfilegroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfilegroupState, opts ...pulumi.ResourceOption) (*Profilegroup, error) {
	var resource Profilegroup
	err := ctx.ReadResource("fortios:firewall/profilegroup:Profilegroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Profilegroup resources.
type profilegroupState struct {
	// Name of an existing Application list.
	ApplicationList *string `pulumi:"applicationList"`
	// Name of an existing Antivirus profile.
	AvProfile *string `pulumi:"avProfile"`
	// Name of an existing CASB profile.
	CasbProfile *string `pulumi:"casbProfile"`
	// Name of an existing CIFS profile.
	CifsProfile *string `pulumi:"cifsProfile"`
	// Name of an existing Diameter filter profile.
	DiameterFilterProfile *string `pulumi:"diameterFilterProfile"`
	// Name of an existing DLP profile.
	DlpProfile *string `pulumi:"dlpProfile"`
	// Name of an existing DLP sensor.
	DlpSensor *string `pulumi:"dlpSensor"`
	// Name of an existing DNS filter profile.
	DnsfilterProfile *string `pulumi:"dnsfilterProfile"`
	// Name of an existing email filter profile.
	EmailfilterProfile *string `pulumi:"emailfilterProfile"`
	// Name of an existing file-filter profile.
	FileFilterProfile *string `pulumi:"fileFilterProfile"`
	// Name of an existing ICAP profile.
	IcapProfile *string `pulumi:"icapProfile"`
	// Name of an existing IPS sensor.
	IpsSensor *string `pulumi:"ipsSensor"`
	// Name of an existing VoIP (ips) profile.
	IpsVoipFilter *string `pulumi:"ipsVoipFilter"`
	// Profile group name.
	Name *string `pulumi:"name"`
	// Name of an existing Protocol options profile.
	ProfileProtocolOptions *string `pulumi:"profileProtocolOptions"`
	// Name of an existing SCTP filter profile.
	SctpFilterProfile *string `pulumi:"sctpFilterProfile"`
	// Name of an existing Spam filter profile.
	SpamfilterProfile *string `pulumi:"spamfilterProfile"`
	// Name of an existing SSH filter profile.
	SshFilterProfile *string `pulumi:"sshFilterProfile"`
	// Name of an existing SSL SSH profile.
	SslSshProfile *string `pulumi:"sslSshProfile"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Name of an existing VideoFilter profile.
	VideofilterProfile *string `pulumi:"videofilterProfile"`
	// Name of an existing virtual-patch profile.
	VirtualPatchProfile *string `pulumi:"virtualPatchProfile"`
	// Name of an existing VoIP (voipd) profile.
	VoipProfile *string `pulumi:"voipProfile"`
	// Name of an existing Web application firewall profile.
	WafProfile *string `pulumi:"wafProfile"`
	// Name of an existing Web filter profile.
	WebfilterProfile *string `pulumi:"webfilterProfile"`
}

type ProfilegroupState struct {
	// Name of an existing Application list.
	ApplicationList pulumi.StringPtrInput
	// Name of an existing Antivirus profile.
	AvProfile pulumi.StringPtrInput
	// Name of an existing CASB profile.
	CasbProfile pulumi.StringPtrInput
	// Name of an existing CIFS profile.
	CifsProfile pulumi.StringPtrInput
	// Name of an existing Diameter filter profile.
	DiameterFilterProfile pulumi.StringPtrInput
	// Name of an existing DLP profile.
	DlpProfile pulumi.StringPtrInput
	// Name of an existing DLP sensor.
	DlpSensor pulumi.StringPtrInput
	// Name of an existing DNS filter profile.
	DnsfilterProfile pulumi.StringPtrInput
	// Name of an existing email filter profile.
	EmailfilterProfile pulumi.StringPtrInput
	// Name of an existing file-filter profile.
	FileFilterProfile pulumi.StringPtrInput
	// Name of an existing ICAP profile.
	IcapProfile pulumi.StringPtrInput
	// Name of an existing IPS sensor.
	IpsSensor pulumi.StringPtrInput
	// Name of an existing VoIP (ips) profile.
	IpsVoipFilter pulumi.StringPtrInput
	// Profile group name.
	Name pulumi.StringPtrInput
	// Name of an existing Protocol options profile.
	ProfileProtocolOptions pulumi.StringPtrInput
	// Name of an existing SCTP filter profile.
	SctpFilterProfile pulumi.StringPtrInput
	// Name of an existing Spam filter profile.
	SpamfilterProfile pulumi.StringPtrInput
	// Name of an existing SSH filter profile.
	SshFilterProfile pulumi.StringPtrInput
	// Name of an existing SSL SSH profile.
	SslSshProfile pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Name of an existing VideoFilter profile.
	VideofilterProfile pulumi.StringPtrInput
	// Name of an existing virtual-patch profile.
	VirtualPatchProfile pulumi.StringPtrInput
	// Name of an existing VoIP (voipd) profile.
	VoipProfile pulumi.StringPtrInput
	// Name of an existing Web application firewall profile.
	WafProfile pulumi.StringPtrInput
	// Name of an existing Web filter profile.
	WebfilterProfile pulumi.StringPtrInput
}

func (ProfilegroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*profilegroupState)(nil)).Elem()
}

type profilegroupArgs struct {
	// Name of an existing Application list.
	ApplicationList *string `pulumi:"applicationList"`
	// Name of an existing Antivirus profile.
	AvProfile *string `pulumi:"avProfile"`
	// Name of an existing CASB profile.
	CasbProfile *string `pulumi:"casbProfile"`
	// Name of an existing CIFS profile.
	CifsProfile *string `pulumi:"cifsProfile"`
	// Name of an existing Diameter filter profile.
	DiameterFilterProfile *string `pulumi:"diameterFilterProfile"`
	// Name of an existing DLP profile.
	DlpProfile *string `pulumi:"dlpProfile"`
	// Name of an existing DLP sensor.
	DlpSensor *string `pulumi:"dlpSensor"`
	// Name of an existing DNS filter profile.
	DnsfilterProfile *string `pulumi:"dnsfilterProfile"`
	// Name of an existing email filter profile.
	EmailfilterProfile *string `pulumi:"emailfilterProfile"`
	// Name of an existing file-filter profile.
	FileFilterProfile *string `pulumi:"fileFilterProfile"`
	// Name of an existing ICAP profile.
	IcapProfile *string `pulumi:"icapProfile"`
	// Name of an existing IPS sensor.
	IpsSensor *string `pulumi:"ipsSensor"`
	// Name of an existing VoIP (ips) profile.
	IpsVoipFilter *string `pulumi:"ipsVoipFilter"`
	// Profile group name.
	Name *string `pulumi:"name"`
	// Name of an existing Protocol options profile.
	ProfileProtocolOptions *string `pulumi:"profileProtocolOptions"`
	// Name of an existing SCTP filter profile.
	SctpFilterProfile *string `pulumi:"sctpFilterProfile"`
	// Name of an existing Spam filter profile.
	SpamfilterProfile *string `pulumi:"spamfilterProfile"`
	// Name of an existing SSH filter profile.
	SshFilterProfile *string `pulumi:"sshFilterProfile"`
	// Name of an existing SSL SSH profile.
	SslSshProfile *string `pulumi:"sslSshProfile"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Name of an existing VideoFilter profile.
	VideofilterProfile *string `pulumi:"videofilterProfile"`
	// Name of an existing virtual-patch profile.
	VirtualPatchProfile *string `pulumi:"virtualPatchProfile"`
	// Name of an existing VoIP (voipd) profile.
	VoipProfile *string `pulumi:"voipProfile"`
	// Name of an existing Web application firewall profile.
	WafProfile *string `pulumi:"wafProfile"`
	// Name of an existing Web filter profile.
	WebfilterProfile *string `pulumi:"webfilterProfile"`
}

// The set of arguments for constructing a Profilegroup resource.
type ProfilegroupArgs struct {
	// Name of an existing Application list.
	ApplicationList pulumi.StringPtrInput
	// Name of an existing Antivirus profile.
	AvProfile pulumi.StringPtrInput
	// Name of an existing CASB profile.
	CasbProfile pulumi.StringPtrInput
	// Name of an existing CIFS profile.
	CifsProfile pulumi.StringPtrInput
	// Name of an existing Diameter filter profile.
	DiameterFilterProfile pulumi.StringPtrInput
	// Name of an existing DLP profile.
	DlpProfile pulumi.StringPtrInput
	// Name of an existing DLP sensor.
	DlpSensor pulumi.StringPtrInput
	// Name of an existing DNS filter profile.
	DnsfilterProfile pulumi.StringPtrInput
	// Name of an existing email filter profile.
	EmailfilterProfile pulumi.StringPtrInput
	// Name of an existing file-filter profile.
	FileFilterProfile pulumi.StringPtrInput
	// Name of an existing ICAP profile.
	IcapProfile pulumi.StringPtrInput
	// Name of an existing IPS sensor.
	IpsSensor pulumi.StringPtrInput
	// Name of an existing VoIP (ips) profile.
	IpsVoipFilter pulumi.StringPtrInput
	// Profile group name.
	Name pulumi.StringPtrInput
	// Name of an existing Protocol options profile.
	ProfileProtocolOptions pulumi.StringPtrInput
	// Name of an existing SCTP filter profile.
	SctpFilterProfile pulumi.StringPtrInput
	// Name of an existing Spam filter profile.
	SpamfilterProfile pulumi.StringPtrInput
	// Name of an existing SSH filter profile.
	SshFilterProfile pulumi.StringPtrInput
	// Name of an existing SSL SSH profile.
	SslSshProfile pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Name of an existing VideoFilter profile.
	VideofilterProfile pulumi.StringPtrInput
	// Name of an existing virtual-patch profile.
	VirtualPatchProfile pulumi.StringPtrInput
	// Name of an existing VoIP (voipd) profile.
	VoipProfile pulumi.StringPtrInput
	// Name of an existing Web application firewall profile.
	WafProfile pulumi.StringPtrInput
	// Name of an existing Web filter profile.
	WebfilterProfile pulumi.StringPtrInput
}

func (ProfilegroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profilegroupArgs)(nil)).Elem()
}

type ProfilegroupInput interface {
	pulumi.Input

	ToProfilegroupOutput() ProfilegroupOutput
	ToProfilegroupOutputWithContext(ctx context.Context) ProfilegroupOutput
}

func (*Profilegroup) ElementType() reflect.Type {
	return reflect.TypeOf((**Profilegroup)(nil)).Elem()
}

func (i *Profilegroup) ToProfilegroupOutput() ProfilegroupOutput {
	return i.ToProfilegroupOutputWithContext(context.Background())
}

func (i *Profilegroup) ToProfilegroupOutputWithContext(ctx context.Context) ProfilegroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfilegroupOutput)
}

// ProfilegroupArrayInput is an input type that accepts ProfilegroupArray and ProfilegroupArrayOutput values.
// You can construct a concrete instance of `ProfilegroupArrayInput` via:
//
//	ProfilegroupArray{ ProfilegroupArgs{...} }
type ProfilegroupArrayInput interface {
	pulumi.Input

	ToProfilegroupArrayOutput() ProfilegroupArrayOutput
	ToProfilegroupArrayOutputWithContext(context.Context) ProfilegroupArrayOutput
}

type ProfilegroupArray []ProfilegroupInput

func (ProfilegroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Profilegroup)(nil)).Elem()
}

func (i ProfilegroupArray) ToProfilegroupArrayOutput() ProfilegroupArrayOutput {
	return i.ToProfilegroupArrayOutputWithContext(context.Background())
}

func (i ProfilegroupArray) ToProfilegroupArrayOutputWithContext(ctx context.Context) ProfilegroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfilegroupArrayOutput)
}

// ProfilegroupMapInput is an input type that accepts ProfilegroupMap and ProfilegroupMapOutput values.
// You can construct a concrete instance of `ProfilegroupMapInput` via:
//
//	ProfilegroupMap{ "key": ProfilegroupArgs{...} }
type ProfilegroupMapInput interface {
	pulumi.Input

	ToProfilegroupMapOutput() ProfilegroupMapOutput
	ToProfilegroupMapOutputWithContext(context.Context) ProfilegroupMapOutput
}

type ProfilegroupMap map[string]ProfilegroupInput

func (ProfilegroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Profilegroup)(nil)).Elem()
}

func (i ProfilegroupMap) ToProfilegroupMapOutput() ProfilegroupMapOutput {
	return i.ToProfilegroupMapOutputWithContext(context.Background())
}

func (i ProfilegroupMap) ToProfilegroupMapOutputWithContext(ctx context.Context) ProfilegroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfilegroupMapOutput)
}

type ProfilegroupOutput struct{ *pulumi.OutputState }

func (ProfilegroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Profilegroup)(nil)).Elem()
}

func (o ProfilegroupOutput) ToProfilegroupOutput() ProfilegroupOutput {
	return o
}

func (o ProfilegroupOutput) ToProfilegroupOutputWithContext(ctx context.Context) ProfilegroupOutput {
	return o
}

// Name of an existing Application list.
func (o ProfilegroupOutput) ApplicationList() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.ApplicationList }).(pulumi.StringOutput)
}

// Name of an existing Antivirus profile.
func (o ProfilegroupOutput) AvProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.AvProfile }).(pulumi.StringOutput)
}

// Name of an existing CASB profile.
func (o ProfilegroupOutput) CasbProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.CasbProfile }).(pulumi.StringOutput)
}

// Name of an existing CIFS profile.
func (o ProfilegroupOutput) CifsProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.CifsProfile }).(pulumi.StringOutput)
}

// Name of an existing Diameter filter profile.
func (o ProfilegroupOutput) DiameterFilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.DiameterFilterProfile }).(pulumi.StringOutput)
}

// Name of an existing DLP profile.
func (o ProfilegroupOutput) DlpProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.DlpProfile }).(pulumi.StringOutput)
}

// Name of an existing DLP sensor.
func (o ProfilegroupOutput) DlpSensor() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.DlpSensor }).(pulumi.StringOutput)
}

// Name of an existing DNS filter profile.
func (o ProfilegroupOutput) DnsfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.DnsfilterProfile }).(pulumi.StringOutput)
}

// Name of an existing email filter profile.
func (o ProfilegroupOutput) EmailfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.EmailfilterProfile }).(pulumi.StringOutput)
}

// Name of an existing file-filter profile.
func (o ProfilegroupOutput) FileFilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.FileFilterProfile }).(pulumi.StringOutput)
}

// Name of an existing ICAP profile.
func (o ProfilegroupOutput) IcapProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.IcapProfile }).(pulumi.StringOutput)
}

// Name of an existing IPS sensor.
func (o ProfilegroupOutput) IpsSensor() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.IpsSensor }).(pulumi.StringOutput)
}

// Name of an existing VoIP (ips) profile.
func (o ProfilegroupOutput) IpsVoipFilter() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.IpsVoipFilter }).(pulumi.StringOutput)
}

// Profile group name.
func (o ProfilegroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Name of an existing Protocol options profile.
func (o ProfilegroupOutput) ProfileProtocolOptions() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.ProfileProtocolOptions }).(pulumi.StringOutput)
}

// Name of an existing SCTP filter profile.
func (o ProfilegroupOutput) SctpFilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.SctpFilterProfile }).(pulumi.StringOutput)
}

// Name of an existing Spam filter profile.
func (o ProfilegroupOutput) SpamfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.SpamfilterProfile }).(pulumi.StringOutput)
}

// Name of an existing SSH filter profile.
func (o ProfilegroupOutput) SshFilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.SshFilterProfile }).(pulumi.StringOutput)
}

// Name of an existing SSL SSH profile.
func (o ProfilegroupOutput) SslSshProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.SslSshProfile }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ProfilegroupOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Name of an existing VideoFilter profile.
func (o ProfilegroupOutput) VideofilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.VideofilterProfile }).(pulumi.StringOutput)
}

// Name of an existing virtual-patch profile.
func (o ProfilegroupOutput) VirtualPatchProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.VirtualPatchProfile }).(pulumi.StringOutput)
}

// Name of an existing VoIP (voipd) profile.
func (o ProfilegroupOutput) VoipProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.VoipProfile }).(pulumi.StringOutput)
}

// Name of an existing Web application firewall profile.
func (o ProfilegroupOutput) WafProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.WafProfile }).(pulumi.StringOutput)
}

// Name of an existing Web filter profile.
func (o ProfilegroupOutput) WebfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Profilegroup) pulumi.StringOutput { return v.WebfilterProfile }).(pulumi.StringOutput)
}

type ProfilegroupArrayOutput struct{ *pulumi.OutputState }

func (ProfilegroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Profilegroup)(nil)).Elem()
}

func (o ProfilegroupArrayOutput) ToProfilegroupArrayOutput() ProfilegroupArrayOutput {
	return o
}

func (o ProfilegroupArrayOutput) ToProfilegroupArrayOutputWithContext(ctx context.Context) ProfilegroupArrayOutput {
	return o
}

func (o ProfilegroupArrayOutput) Index(i pulumi.IntInput) ProfilegroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Profilegroup {
		return vs[0].([]*Profilegroup)[vs[1].(int)]
	}).(ProfilegroupOutput)
}

type ProfilegroupMapOutput struct{ *pulumi.OutputState }

func (ProfilegroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Profilegroup)(nil)).Elem()
}

func (o ProfilegroupMapOutput) ToProfilegroupMapOutput() ProfilegroupMapOutput {
	return o
}

func (o ProfilegroupMapOutput) ToProfilegroupMapOutputWithContext(ctx context.Context) ProfilegroupMapOutput {
	return o
}

func (o ProfilegroupMapOutput) MapIndex(k pulumi.StringInput) ProfilegroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Profilegroup {
		return vs[0].(map[string]*Profilegroup)[vs[1].(string)]
	}).(ProfilegroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProfilegroupInput)(nil)).Elem(), &Profilegroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfilegroupArrayInput)(nil)).Elem(), ProfilegroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfilegroupMapInput)(nil)).Elem(), ProfilegroupMap{})
	pulumi.RegisterOutputType(ProfilegroupOutput{})
	pulumi.RegisterOutputType(ProfilegroupArrayOutput{})
	pulumi.RegisterOutputType(ProfilegroupMapOutput{})
}
