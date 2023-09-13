// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure protocol options.
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
//			_, err := firewall.NewProfileprotocoloptions(ctx, "trname", &firewall.ProfileprotocoloptionsArgs{
//				Dns: &firewall.ProfileprotocoloptionsDnsArgs{
//					Ports:  pulumi.Int(53),
//					Status: pulumi.String("enable"),
//				},
//				Ftp: &firewall.ProfileprotocoloptionsFtpArgs{
//					ComfortAmount:             pulumi.Int(1),
//					ComfortInterval:           pulumi.Int(10),
//					InspectAll:                pulumi.String("disable"),
//					Options:                   pulumi.String("splice"),
//					OversizeLimit:             pulumi.Int(10),
//					Ports:                     pulumi.Int(21),
//					ScanBzip2:                 pulumi.String("enable"),
//					Status:                    pulumi.String("enable"),
//					UncompressedNestLimit:     pulumi.Int(12),
//					UncompressedOversizeLimit: pulumi.Int(10),
//				},
//				Http: &firewall.ProfileprotocoloptionsHttpArgs{
//					BlockPageStatusCode:       pulumi.Int(403),
//					ComfortAmount:             pulumi.Int(1),
//					ComfortInterval:           pulumi.Int(10),
//					FortinetBar:               pulumi.String("disable"),
//					FortinetBarPort:           pulumi.Int(8011),
//					HttpPolicy:                pulumi.String("disable"),
//					InspectAll:                pulumi.String("disable"),
//					OversizeLimit:             pulumi.Int(10),
//					Ports:                     pulumi.Int(80),
//					RangeBlock:                pulumi.String("disable"),
//					RetryCount:                pulumi.Int(0),
//					ScanBzip2:                 pulumi.String("enable"),
//					Status:                    pulumi.String("enable"),
//					StreamingContentBypass:    pulumi.String("enable"),
//					StripXForwardedFor:        pulumi.String("disable"),
//					SwitchingProtocols:        pulumi.String("bypass"),
//					UncompressedNestLimit:     pulumi.Int(12),
//					UncompressedOversizeLimit: pulumi.Int(10),
//				},
//				Imap: &firewall.ProfileprotocoloptionsImapArgs{
//					InspectAll:                pulumi.String("disable"),
//					Options:                   pulumi.String("fragmail"),
//					OversizeLimit:             pulumi.Int(10),
//					Ports:                     pulumi.Int(143),
//					ScanBzip2:                 pulumi.String("enable"),
//					Status:                    pulumi.String("enable"),
//					UncompressedNestLimit:     pulumi.Int(12),
//					UncompressedOversizeLimit: pulumi.Int(10),
//				},
//				MailSignature: &firewall.ProfileprotocoloptionsMailSignatureArgs{
//					Status: pulumi.String("disable"),
//				},
//				Mapi: &firewall.ProfileprotocoloptionsMapiArgs{
//					Options:                   pulumi.String("fragmail"),
//					OversizeLimit:             pulumi.Int(10),
//					Ports:                     pulumi.Int(135),
//					ScanBzip2:                 pulumi.String("enable"),
//					Status:                    pulumi.String("enable"),
//					UncompressedNestLimit:     pulumi.Int(12),
//					UncompressedOversizeLimit: pulumi.Int(10),
//				},
//				Nntp: &firewall.ProfileprotocoloptionsNntpArgs{
//					InspectAll:                pulumi.String("disable"),
//					Options:                   pulumi.String("splice"),
//					OversizeLimit:             pulumi.Int(10),
//					Ports:                     pulumi.Int(119),
//					ScanBzip2:                 pulumi.String("enable"),
//					Status:                    pulumi.String("enable"),
//					UncompressedNestLimit:     pulumi.Int(12),
//					UncompressedOversizeLimit: pulumi.Int(10),
//				},
//				OversizeLog: pulumi.String("disable"),
//				Pop3: &firewall.ProfileprotocoloptionsPop3Args{
//					InspectAll:                pulumi.String("disable"),
//					Options:                   pulumi.String("fragmail"),
//					OversizeLimit:             pulumi.Int(10),
//					Ports:                     pulumi.Int(110),
//					ScanBzip2:                 pulumi.String("enable"),
//					Status:                    pulumi.String("enable"),
//					UncompressedNestLimit:     pulumi.Int(12),
//					UncompressedOversizeLimit: pulumi.Int(10),
//				},
//				RpcOverHttp: pulumi.String("disable"),
//				Smtp: &firewall.ProfileprotocoloptionsSmtpArgs{
//					InspectAll:                pulumi.String("disable"),
//					Options:                   pulumi.String("fragmail splice"),
//					OversizeLimit:             pulumi.Int(10),
//					Ports:                     pulumi.Int(25),
//					ScanBzip2:                 pulumi.String("enable"),
//					ServerBusy:                pulumi.String("disable"),
//					Status:                    pulumi.String("enable"),
//					UncompressedNestLimit:     pulumi.Int(12),
//					UncompressedOversizeLimit: pulumi.Int(10),
//				},
//				SwitchingProtocolsLog: pulumi.String("disable"),
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
// # Firewall ProfileProtocolOptions can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:firewall/profileprotocoloptions:Profileprotocoloptions labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:firewall/profileprotocoloptions:Profileprotocoloptions labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Profileprotocoloptions struct {
	pulumi.CustomResourceState

	// Configure CIFS protocol options. The structure of `cifs` block is documented below.
	Cifs ProfileprotocoloptionsCifsOutput `pulumi:"cifs"`
	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Configure DNS protocol options. The structure of `dns` block is documented below.
	Dns ProfileprotocoloptionsDnsOutput `pulumi:"dns"`
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet pulumi.StringOutput `pulumi:"featureSet"`
	// Configure FTP protocol options. The structure of `ftp` block is documented below.
	Ftp ProfileprotocoloptionsFtpOutput `pulumi:"ftp"`
	// Configure HTTP protocol options. The structure of `http` block is documented below.
	Http ProfileprotocoloptionsHttpOutput `pulumi:"http"`
	// Configure IMAP protocol options. The structure of `imap` block is documented below.
	Imap ProfileprotocoloptionsImapOutput `pulumi:"imap"`
	// Configure Mail signature. The structure of `mailSignature` block is documented below.
	MailSignature ProfileprotocoloptionsMailSignatureOutput `pulumi:"mailSignature"`
	// Configure MAPI protocol options. The structure of `mapi` block is documented below.
	Mapi ProfileprotocoloptionsMapiOutput `pulumi:"mapi"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Configure NNTP protocol options. The structure of `nntp` block is documented below.
	Nntp ProfileprotocoloptionsNntpOutput `pulumi:"nntp"`
	// Enable/disable logging for antivirus oversize file blocking. Valid values: `disable`, `enable`.
	OversizeLog pulumi.StringOutput `pulumi:"oversizeLog"`
	// Configure POP3 protocol options. The structure of `pop3` block is documented below.
	Pop3 ProfileprotocoloptionsPop3Output `pulumi:"pop3"`
	// Name of the replacement message group to be used
	ReplacemsgGroup pulumi.StringOutput `pulumi:"replacemsgGroup"`
	// Enable/disable inspection of RPC over HTTP. Valid values: `enable`, `disable`.
	RpcOverHttp pulumi.StringOutput `pulumi:"rpcOverHttp"`
	// Configure SMTP protocol options. The structure of `smtp` block is documented below.
	Smtp ProfileprotocoloptionsSmtpOutput `pulumi:"smtp"`
	// Configure SFTP and SCP protocol options. The structure of `ssh` block is documented below.
	Ssh ProfileprotocoloptionsSshOutput `pulumi:"ssh"`
	// Enable/disable logging for HTTP/HTTPS switching protocols. Valid values: `disable`, `enable`.
	SwitchingProtocolsLog pulumi.StringOutput `pulumi:"switchingProtocolsLog"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewProfileprotocoloptions registers a new resource with the given unique name, arguments, and options.
func NewProfileprotocoloptions(ctx *pulumi.Context,
	name string, args *ProfileprotocoloptionsArgs, opts ...pulumi.ResourceOption) (*Profileprotocoloptions, error) {
	if args == nil {
		args = &ProfileprotocoloptionsArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Profileprotocoloptions
	err := ctx.RegisterResource("fortios:firewall/profileprotocoloptions:Profileprotocoloptions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfileprotocoloptions gets an existing Profileprotocoloptions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfileprotocoloptions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileprotocoloptionsState, opts ...pulumi.ResourceOption) (*Profileprotocoloptions, error) {
	var resource Profileprotocoloptions
	err := ctx.ReadResource("fortios:firewall/profileprotocoloptions:Profileprotocoloptions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Profileprotocoloptions resources.
type profileprotocoloptionsState struct {
	// Configure CIFS protocol options. The structure of `cifs` block is documented below.
	Cifs *ProfileprotocoloptionsCifs `pulumi:"cifs"`
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Configure DNS protocol options. The structure of `dns` block is documented below.
	Dns *ProfileprotocoloptionsDns `pulumi:"dns"`
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet *string `pulumi:"featureSet"`
	// Configure FTP protocol options. The structure of `ftp` block is documented below.
	Ftp *ProfileprotocoloptionsFtp `pulumi:"ftp"`
	// Configure HTTP protocol options. The structure of `http` block is documented below.
	Http *ProfileprotocoloptionsHttp `pulumi:"http"`
	// Configure IMAP protocol options. The structure of `imap` block is documented below.
	Imap *ProfileprotocoloptionsImap `pulumi:"imap"`
	// Configure Mail signature. The structure of `mailSignature` block is documented below.
	MailSignature *ProfileprotocoloptionsMailSignature `pulumi:"mailSignature"`
	// Configure MAPI protocol options. The structure of `mapi` block is documented below.
	Mapi *ProfileprotocoloptionsMapi `pulumi:"mapi"`
	// Name.
	Name *string `pulumi:"name"`
	// Configure NNTP protocol options. The structure of `nntp` block is documented below.
	Nntp *ProfileprotocoloptionsNntp `pulumi:"nntp"`
	// Enable/disable logging for antivirus oversize file blocking. Valid values: `disable`, `enable`.
	OversizeLog *string `pulumi:"oversizeLog"`
	// Configure POP3 protocol options. The structure of `pop3` block is documented below.
	Pop3 *ProfileprotocoloptionsPop3 `pulumi:"pop3"`
	// Name of the replacement message group to be used
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// Enable/disable inspection of RPC over HTTP. Valid values: `enable`, `disable`.
	RpcOverHttp *string `pulumi:"rpcOverHttp"`
	// Configure SMTP protocol options. The structure of `smtp` block is documented below.
	Smtp *ProfileprotocoloptionsSmtp `pulumi:"smtp"`
	// Configure SFTP and SCP protocol options. The structure of `ssh` block is documented below.
	Ssh *ProfileprotocoloptionsSsh `pulumi:"ssh"`
	// Enable/disable logging for HTTP/HTTPS switching protocols. Valid values: `disable`, `enable`.
	SwitchingProtocolsLog *string `pulumi:"switchingProtocolsLog"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ProfileprotocoloptionsState struct {
	// Configure CIFS protocol options. The structure of `cifs` block is documented below.
	Cifs ProfileprotocoloptionsCifsPtrInput
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Configure DNS protocol options. The structure of `dns` block is documented below.
	Dns ProfileprotocoloptionsDnsPtrInput
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet pulumi.StringPtrInput
	// Configure FTP protocol options. The structure of `ftp` block is documented below.
	Ftp ProfileprotocoloptionsFtpPtrInput
	// Configure HTTP protocol options. The structure of `http` block is documented below.
	Http ProfileprotocoloptionsHttpPtrInput
	// Configure IMAP protocol options. The structure of `imap` block is documented below.
	Imap ProfileprotocoloptionsImapPtrInput
	// Configure Mail signature. The structure of `mailSignature` block is documented below.
	MailSignature ProfileprotocoloptionsMailSignaturePtrInput
	// Configure MAPI protocol options. The structure of `mapi` block is documented below.
	Mapi ProfileprotocoloptionsMapiPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Configure NNTP protocol options. The structure of `nntp` block is documented below.
	Nntp ProfileprotocoloptionsNntpPtrInput
	// Enable/disable logging for antivirus oversize file blocking. Valid values: `disable`, `enable`.
	OversizeLog pulumi.StringPtrInput
	// Configure POP3 protocol options. The structure of `pop3` block is documented below.
	Pop3 ProfileprotocoloptionsPop3PtrInput
	// Name of the replacement message group to be used
	ReplacemsgGroup pulumi.StringPtrInput
	// Enable/disable inspection of RPC over HTTP. Valid values: `enable`, `disable`.
	RpcOverHttp pulumi.StringPtrInput
	// Configure SMTP protocol options. The structure of `smtp` block is documented below.
	Smtp ProfileprotocoloptionsSmtpPtrInput
	// Configure SFTP and SCP protocol options. The structure of `ssh` block is documented below.
	Ssh ProfileprotocoloptionsSshPtrInput
	// Enable/disable logging for HTTP/HTTPS switching protocols. Valid values: `disable`, `enable`.
	SwitchingProtocolsLog pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ProfileprotocoloptionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileprotocoloptionsState)(nil)).Elem()
}

type profileprotocoloptionsArgs struct {
	// Configure CIFS protocol options. The structure of `cifs` block is documented below.
	Cifs *ProfileprotocoloptionsCifs `pulumi:"cifs"`
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Configure DNS protocol options. The structure of `dns` block is documented below.
	Dns *ProfileprotocoloptionsDns `pulumi:"dns"`
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet *string `pulumi:"featureSet"`
	// Configure FTP protocol options. The structure of `ftp` block is documented below.
	Ftp *ProfileprotocoloptionsFtp `pulumi:"ftp"`
	// Configure HTTP protocol options. The structure of `http` block is documented below.
	Http *ProfileprotocoloptionsHttp `pulumi:"http"`
	// Configure IMAP protocol options. The structure of `imap` block is documented below.
	Imap *ProfileprotocoloptionsImap `pulumi:"imap"`
	// Configure Mail signature. The structure of `mailSignature` block is documented below.
	MailSignature *ProfileprotocoloptionsMailSignature `pulumi:"mailSignature"`
	// Configure MAPI protocol options. The structure of `mapi` block is documented below.
	Mapi *ProfileprotocoloptionsMapi `pulumi:"mapi"`
	// Name.
	Name *string `pulumi:"name"`
	// Configure NNTP protocol options. The structure of `nntp` block is documented below.
	Nntp *ProfileprotocoloptionsNntp `pulumi:"nntp"`
	// Enable/disable logging for antivirus oversize file blocking. Valid values: `disable`, `enable`.
	OversizeLog *string `pulumi:"oversizeLog"`
	// Configure POP3 protocol options. The structure of `pop3` block is documented below.
	Pop3 *ProfileprotocoloptionsPop3 `pulumi:"pop3"`
	// Name of the replacement message group to be used
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// Enable/disable inspection of RPC over HTTP. Valid values: `enable`, `disable`.
	RpcOverHttp *string `pulumi:"rpcOverHttp"`
	// Configure SMTP protocol options. The structure of `smtp` block is documented below.
	Smtp *ProfileprotocoloptionsSmtp `pulumi:"smtp"`
	// Configure SFTP and SCP protocol options. The structure of `ssh` block is documented below.
	Ssh *ProfileprotocoloptionsSsh `pulumi:"ssh"`
	// Enable/disable logging for HTTP/HTTPS switching protocols. Valid values: `disable`, `enable`.
	SwitchingProtocolsLog *string `pulumi:"switchingProtocolsLog"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Profileprotocoloptions resource.
type ProfileprotocoloptionsArgs struct {
	// Configure CIFS protocol options. The structure of `cifs` block is documented below.
	Cifs ProfileprotocoloptionsCifsPtrInput
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Configure DNS protocol options. The structure of `dns` block is documented below.
	Dns ProfileprotocoloptionsDnsPtrInput
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet pulumi.StringPtrInput
	// Configure FTP protocol options. The structure of `ftp` block is documented below.
	Ftp ProfileprotocoloptionsFtpPtrInput
	// Configure HTTP protocol options. The structure of `http` block is documented below.
	Http ProfileprotocoloptionsHttpPtrInput
	// Configure IMAP protocol options. The structure of `imap` block is documented below.
	Imap ProfileprotocoloptionsImapPtrInput
	// Configure Mail signature. The structure of `mailSignature` block is documented below.
	MailSignature ProfileprotocoloptionsMailSignaturePtrInput
	// Configure MAPI protocol options. The structure of `mapi` block is documented below.
	Mapi ProfileprotocoloptionsMapiPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Configure NNTP protocol options. The structure of `nntp` block is documented below.
	Nntp ProfileprotocoloptionsNntpPtrInput
	// Enable/disable logging for antivirus oversize file blocking. Valid values: `disable`, `enable`.
	OversizeLog pulumi.StringPtrInput
	// Configure POP3 protocol options. The structure of `pop3` block is documented below.
	Pop3 ProfileprotocoloptionsPop3PtrInput
	// Name of the replacement message group to be used
	ReplacemsgGroup pulumi.StringPtrInput
	// Enable/disable inspection of RPC over HTTP. Valid values: `enable`, `disable`.
	RpcOverHttp pulumi.StringPtrInput
	// Configure SMTP protocol options. The structure of `smtp` block is documented below.
	Smtp ProfileprotocoloptionsSmtpPtrInput
	// Configure SFTP and SCP protocol options. The structure of `ssh` block is documented below.
	Ssh ProfileprotocoloptionsSshPtrInput
	// Enable/disable logging for HTTP/HTTPS switching protocols. Valid values: `disable`, `enable`.
	SwitchingProtocolsLog pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ProfileprotocoloptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileprotocoloptionsArgs)(nil)).Elem()
}

type ProfileprotocoloptionsInput interface {
	pulumi.Input

	ToProfileprotocoloptionsOutput() ProfileprotocoloptionsOutput
	ToProfileprotocoloptionsOutputWithContext(ctx context.Context) ProfileprotocoloptionsOutput
}

func (*Profileprotocoloptions) ElementType() reflect.Type {
	return reflect.TypeOf((**Profileprotocoloptions)(nil)).Elem()
}

func (i *Profileprotocoloptions) ToProfileprotocoloptionsOutput() ProfileprotocoloptionsOutput {
	return i.ToProfileprotocoloptionsOutputWithContext(context.Background())
}

func (i *Profileprotocoloptions) ToProfileprotocoloptionsOutputWithContext(ctx context.Context) ProfileprotocoloptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileprotocoloptionsOutput)
}

// ProfileprotocoloptionsArrayInput is an input type that accepts ProfileprotocoloptionsArray and ProfileprotocoloptionsArrayOutput values.
// You can construct a concrete instance of `ProfileprotocoloptionsArrayInput` via:
//
//	ProfileprotocoloptionsArray{ ProfileprotocoloptionsArgs{...} }
type ProfileprotocoloptionsArrayInput interface {
	pulumi.Input

	ToProfileprotocoloptionsArrayOutput() ProfileprotocoloptionsArrayOutput
	ToProfileprotocoloptionsArrayOutputWithContext(context.Context) ProfileprotocoloptionsArrayOutput
}

type ProfileprotocoloptionsArray []ProfileprotocoloptionsInput

func (ProfileprotocoloptionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Profileprotocoloptions)(nil)).Elem()
}

func (i ProfileprotocoloptionsArray) ToProfileprotocoloptionsArrayOutput() ProfileprotocoloptionsArrayOutput {
	return i.ToProfileprotocoloptionsArrayOutputWithContext(context.Background())
}

func (i ProfileprotocoloptionsArray) ToProfileprotocoloptionsArrayOutputWithContext(ctx context.Context) ProfileprotocoloptionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileprotocoloptionsArrayOutput)
}

// ProfileprotocoloptionsMapInput is an input type that accepts ProfileprotocoloptionsMap and ProfileprotocoloptionsMapOutput values.
// You can construct a concrete instance of `ProfileprotocoloptionsMapInput` via:
//
//	ProfileprotocoloptionsMap{ "key": ProfileprotocoloptionsArgs{...} }
type ProfileprotocoloptionsMapInput interface {
	pulumi.Input

	ToProfileprotocoloptionsMapOutput() ProfileprotocoloptionsMapOutput
	ToProfileprotocoloptionsMapOutputWithContext(context.Context) ProfileprotocoloptionsMapOutput
}

type ProfileprotocoloptionsMap map[string]ProfileprotocoloptionsInput

func (ProfileprotocoloptionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Profileprotocoloptions)(nil)).Elem()
}

func (i ProfileprotocoloptionsMap) ToProfileprotocoloptionsMapOutput() ProfileprotocoloptionsMapOutput {
	return i.ToProfileprotocoloptionsMapOutputWithContext(context.Background())
}

func (i ProfileprotocoloptionsMap) ToProfileprotocoloptionsMapOutputWithContext(ctx context.Context) ProfileprotocoloptionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileprotocoloptionsMapOutput)
}

type ProfileprotocoloptionsOutput struct{ *pulumi.OutputState }

func (ProfileprotocoloptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Profileprotocoloptions)(nil)).Elem()
}

func (o ProfileprotocoloptionsOutput) ToProfileprotocoloptionsOutput() ProfileprotocoloptionsOutput {
	return o
}

func (o ProfileprotocoloptionsOutput) ToProfileprotocoloptionsOutputWithContext(ctx context.Context) ProfileprotocoloptionsOutput {
	return o
}

// Configure CIFS protocol options. The structure of `cifs` block is documented below.
func (o ProfileprotocoloptionsOutput) Cifs() ProfileprotocoloptionsCifsOutput {
	return o.ApplyT(func(v *Profileprotocoloptions) ProfileprotocoloptionsCifsOutput { return v.Cifs }).(ProfileprotocoloptionsCifsOutput)
}

// Optional comments.
func (o ProfileprotocoloptionsOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profileprotocoloptions) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Configure DNS protocol options. The structure of `dns` block is documented below.
func (o ProfileprotocoloptionsOutput) Dns() ProfileprotocoloptionsDnsOutput {
	return o.ApplyT(func(v *Profileprotocoloptions) ProfileprotocoloptionsDnsOutput { return v.Dns }).(ProfileprotocoloptionsDnsOutput)
}

// Flow/proxy feature set. Valid values: `flow`, `proxy`.
func (o ProfileprotocoloptionsOutput) FeatureSet() pulumi.StringOutput {
	return o.ApplyT(func(v *Profileprotocoloptions) pulumi.StringOutput { return v.FeatureSet }).(pulumi.StringOutput)
}

// Configure FTP protocol options. The structure of `ftp` block is documented below.
func (o ProfileprotocoloptionsOutput) Ftp() ProfileprotocoloptionsFtpOutput {
	return o.ApplyT(func(v *Profileprotocoloptions) ProfileprotocoloptionsFtpOutput { return v.Ftp }).(ProfileprotocoloptionsFtpOutput)
}

// Configure HTTP protocol options. The structure of `http` block is documented below.
func (o ProfileprotocoloptionsOutput) Http() ProfileprotocoloptionsHttpOutput {
	return o.ApplyT(func(v *Profileprotocoloptions) ProfileprotocoloptionsHttpOutput { return v.Http }).(ProfileprotocoloptionsHttpOutput)
}

// Configure IMAP protocol options. The structure of `imap` block is documented below.
func (o ProfileprotocoloptionsOutput) Imap() ProfileprotocoloptionsImapOutput {
	return o.ApplyT(func(v *Profileprotocoloptions) ProfileprotocoloptionsImapOutput { return v.Imap }).(ProfileprotocoloptionsImapOutput)
}

// Configure Mail signature. The structure of `mailSignature` block is documented below.
func (o ProfileprotocoloptionsOutput) MailSignature() ProfileprotocoloptionsMailSignatureOutput {
	return o.ApplyT(func(v *Profileprotocoloptions) ProfileprotocoloptionsMailSignatureOutput { return v.MailSignature }).(ProfileprotocoloptionsMailSignatureOutput)
}

// Configure MAPI protocol options. The structure of `mapi` block is documented below.
func (o ProfileprotocoloptionsOutput) Mapi() ProfileprotocoloptionsMapiOutput {
	return o.ApplyT(func(v *Profileprotocoloptions) ProfileprotocoloptionsMapiOutput { return v.Mapi }).(ProfileprotocoloptionsMapiOutput)
}

// Name.
func (o ProfileprotocoloptionsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Profileprotocoloptions) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Configure NNTP protocol options. The structure of `nntp` block is documented below.
func (o ProfileprotocoloptionsOutput) Nntp() ProfileprotocoloptionsNntpOutput {
	return o.ApplyT(func(v *Profileprotocoloptions) ProfileprotocoloptionsNntpOutput { return v.Nntp }).(ProfileprotocoloptionsNntpOutput)
}

// Enable/disable logging for antivirus oversize file blocking. Valid values: `disable`, `enable`.
func (o ProfileprotocoloptionsOutput) OversizeLog() pulumi.StringOutput {
	return o.ApplyT(func(v *Profileprotocoloptions) pulumi.StringOutput { return v.OversizeLog }).(pulumi.StringOutput)
}

// Configure POP3 protocol options. The structure of `pop3` block is documented below.
func (o ProfileprotocoloptionsOutput) Pop3() ProfileprotocoloptionsPop3Output {
	return o.ApplyT(func(v *Profileprotocoloptions) ProfileprotocoloptionsPop3Output { return v.Pop3 }).(ProfileprotocoloptionsPop3Output)
}

// Name of the replacement message group to be used
func (o ProfileprotocoloptionsOutput) ReplacemsgGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Profileprotocoloptions) pulumi.StringOutput { return v.ReplacemsgGroup }).(pulumi.StringOutput)
}

// Enable/disable inspection of RPC over HTTP. Valid values: `enable`, `disable`.
func (o ProfileprotocoloptionsOutput) RpcOverHttp() pulumi.StringOutput {
	return o.ApplyT(func(v *Profileprotocoloptions) pulumi.StringOutput { return v.RpcOverHttp }).(pulumi.StringOutput)
}

// Configure SMTP protocol options. The structure of `smtp` block is documented below.
func (o ProfileprotocoloptionsOutput) Smtp() ProfileprotocoloptionsSmtpOutput {
	return o.ApplyT(func(v *Profileprotocoloptions) ProfileprotocoloptionsSmtpOutput { return v.Smtp }).(ProfileprotocoloptionsSmtpOutput)
}

// Configure SFTP and SCP protocol options. The structure of `ssh` block is documented below.
func (o ProfileprotocoloptionsOutput) Ssh() ProfileprotocoloptionsSshOutput {
	return o.ApplyT(func(v *Profileprotocoloptions) ProfileprotocoloptionsSshOutput { return v.Ssh }).(ProfileprotocoloptionsSshOutput)
}

// Enable/disable logging for HTTP/HTTPS switching protocols. Valid values: `disable`, `enable`.
func (o ProfileprotocoloptionsOutput) SwitchingProtocolsLog() pulumi.StringOutput {
	return o.ApplyT(func(v *Profileprotocoloptions) pulumi.StringOutput { return v.SwitchingProtocolsLog }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ProfileprotocoloptionsOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profileprotocoloptions) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ProfileprotocoloptionsArrayOutput struct{ *pulumi.OutputState }

func (ProfileprotocoloptionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Profileprotocoloptions)(nil)).Elem()
}

func (o ProfileprotocoloptionsArrayOutput) ToProfileprotocoloptionsArrayOutput() ProfileprotocoloptionsArrayOutput {
	return o
}

func (o ProfileprotocoloptionsArrayOutput) ToProfileprotocoloptionsArrayOutputWithContext(ctx context.Context) ProfileprotocoloptionsArrayOutput {
	return o
}

func (o ProfileprotocoloptionsArrayOutput) Index(i pulumi.IntInput) ProfileprotocoloptionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Profileprotocoloptions {
		return vs[0].([]*Profileprotocoloptions)[vs[1].(int)]
	}).(ProfileprotocoloptionsOutput)
}

type ProfileprotocoloptionsMapOutput struct{ *pulumi.OutputState }

func (ProfileprotocoloptionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Profileprotocoloptions)(nil)).Elem()
}

func (o ProfileprotocoloptionsMapOutput) ToProfileprotocoloptionsMapOutput() ProfileprotocoloptionsMapOutput {
	return o
}

func (o ProfileprotocoloptionsMapOutput) ToProfileprotocoloptionsMapOutputWithContext(ctx context.Context) ProfileprotocoloptionsMapOutput {
	return o
}

func (o ProfileprotocoloptionsMapOutput) MapIndex(k pulumi.StringInput) ProfileprotocoloptionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Profileprotocoloptions {
		return vs[0].(map[string]*Profileprotocoloptions)[vs[1].(string)]
	}).(ProfileprotocoloptionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileprotocoloptionsInput)(nil)).Elem(), &Profileprotocoloptions{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileprotocoloptionsArrayInput)(nil)).Elem(), ProfileprotocoloptionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileprotocoloptionsMapInput)(nil)).Elem(), ProfileprotocoloptionsMap{})
	pulumi.RegisterOutputType(ProfileprotocoloptionsOutput{})
	pulumi.RegisterOutputType(ProfileprotocoloptionsArrayOutput{})
	pulumi.RegisterOutputType(ProfileprotocoloptionsMapOutput{})
}
