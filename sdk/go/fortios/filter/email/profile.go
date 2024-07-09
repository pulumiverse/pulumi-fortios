// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package email

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure Email Filter profiles. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// Emailfilter Profile can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:filter/email/profile:Profile labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:filter/email/profile:Profile labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Profile struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Enable/disable external Email inspection. Valid values: `enable`, `disable`.
	External pulumi.StringOutput `pulumi:"external"`
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet pulumi.StringOutput `pulumi:"featureSet"`
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter ProfileFileFilterOutput `pulumi:"fileFilter"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Gmail. The structure of `gmail` block is documented below.
	Gmail ProfileGmailOutput `pulumi:"gmail"`
	// IMAP. The structure of `imap` block is documented below.
	Imap ProfileImapOutput `pulumi:"imap"`
	// MAPI. The structure of `mapi` block is documented below.
	Mapi ProfileMapiOutput `pulumi:"mapi"`
	// MSN Hotmail. The structure of `msnHotmail` block is documented below.
	MsnHotmail ProfileMsnHotmailOutput `pulumi:"msnHotmail"`
	// Profile name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Options.
	Options pulumi.StringOutput `pulumi:"options"`
	// Other supported webmails. The structure of `otherWebmails` block is documented below.
	OtherWebmails ProfileOtherWebmailsOutput `pulumi:"otherWebmails"`
	// POP3. The structure of `pop3` block is documented below.
	Pop3 ProfilePop3Output `pulumi:"pop3"`
	// Replacement message group.
	ReplacemsgGroup pulumi.StringOutput `pulumi:"replacemsgGroup"`
	// SMTP. The structure of `smtp` block is documented below.
	Smtp ProfileSmtpOutput `pulumi:"smtp"`
	// Anti-spam block/allow list table ID.
	SpamBalTable pulumi.IntOutput `pulumi:"spamBalTable"`
	// Anti-spam black/white list table ID.
	SpamBwlTable pulumi.IntOutput `pulumi:"spamBwlTable"`
	// Anti-spam banned word table ID.
	SpamBwordTable pulumi.IntOutput `pulumi:"spamBwordTable"`
	// Spam banned word threshold.
	SpamBwordThreshold pulumi.IntOutput `pulumi:"spamBwordThreshold"`
	// Enable/disable spam filtering. Valid values: `enable`, `disable`.
	SpamFiltering pulumi.StringOutput `pulumi:"spamFiltering"`
	// Anti-spam IP trust table ID.
	SpamIptrustTable pulumi.IntOutput `pulumi:"spamIptrustTable"`
	// Enable/disable spam logging for email filtering. Valid values: `disable`, `enable`.
	SpamLog pulumi.StringOutput `pulumi:"spamLog"`
	// Enable/disable logging FortiGuard spam response. Valid values: `disable`, `enable`.
	SpamLogFortiguardResponse pulumi.StringOutput `pulumi:"spamLogFortiguardResponse"`
	// Anti-spam MIME header table ID.
	SpamMheaderTable pulumi.IntOutput `pulumi:"spamMheaderTable"`
	// Anti-spam DNSBL table ID.
	SpamRblTable pulumi.IntOutput `pulumi:"spamRblTable"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
	// Yahoo! Mail. The structure of `yahooMail` block is documented below.
	YahooMail ProfileYahooMailOutput `pulumi:"yahooMail"`
}

// NewProfile registers a new resource with the given unique name, arguments, and options.
func NewProfile(ctx *pulumi.Context,
	name string, args *ProfileArgs, opts ...pulumi.ResourceOption) (*Profile, error) {
	if args == nil {
		args = &ProfileArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Profile
	err := ctx.RegisterResource("fortios:filter/email/profile:Profile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfile gets an existing Profile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileState, opts ...pulumi.ResourceOption) (*Profile, error) {
	var resource Profile
	err := ctx.ReadResource("fortios:filter/email/profile:Profile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Profile resources.
type profileState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Enable/disable external Email inspection. Valid values: `enable`, `disable`.
	External *string `pulumi:"external"`
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet *string `pulumi:"featureSet"`
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter *ProfileFileFilter `pulumi:"fileFilter"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Gmail. The structure of `gmail` block is documented below.
	Gmail *ProfileGmail `pulumi:"gmail"`
	// IMAP. The structure of `imap` block is documented below.
	Imap *ProfileImap `pulumi:"imap"`
	// MAPI. The structure of `mapi` block is documented below.
	Mapi *ProfileMapi `pulumi:"mapi"`
	// MSN Hotmail. The structure of `msnHotmail` block is documented below.
	MsnHotmail *ProfileMsnHotmail `pulumi:"msnHotmail"`
	// Profile name.
	Name *string `pulumi:"name"`
	// Options.
	Options *string `pulumi:"options"`
	// Other supported webmails. The structure of `otherWebmails` block is documented below.
	OtherWebmails *ProfileOtherWebmails `pulumi:"otherWebmails"`
	// POP3. The structure of `pop3` block is documented below.
	Pop3 *ProfilePop3 `pulumi:"pop3"`
	// Replacement message group.
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// SMTP. The structure of `smtp` block is documented below.
	Smtp *ProfileSmtp `pulumi:"smtp"`
	// Anti-spam block/allow list table ID.
	SpamBalTable *int `pulumi:"spamBalTable"`
	// Anti-spam black/white list table ID.
	SpamBwlTable *int `pulumi:"spamBwlTable"`
	// Anti-spam banned word table ID.
	SpamBwordTable *int `pulumi:"spamBwordTable"`
	// Spam banned word threshold.
	SpamBwordThreshold *int `pulumi:"spamBwordThreshold"`
	// Enable/disable spam filtering. Valid values: `enable`, `disable`.
	SpamFiltering *string `pulumi:"spamFiltering"`
	// Anti-spam IP trust table ID.
	SpamIptrustTable *int `pulumi:"spamIptrustTable"`
	// Enable/disable spam logging for email filtering. Valid values: `disable`, `enable`.
	SpamLog *string `pulumi:"spamLog"`
	// Enable/disable logging FortiGuard spam response. Valid values: `disable`, `enable`.
	SpamLogFortiguardResponse *string `pulumi:"spamLogFortiguardResponse"`
	// Anti-spam MIME header table ID.
	SpamMheaderTable *int `pulumi:"spamMheaderTable"`
	// Anti-spam DNSBL table ID.
	SpamRblTable *int `pulumi:"spamRblTable"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Yahoo! Mail. The structure of `yahooMail` block is documented below.
	YahooMail *ProfileYahooMail `pulumi:"yahooMail"`
}

type ProfileState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Enable/disable external Email inspection. Valid values: `enable`, `disable`.
	External pulumi.StringPtrInput
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet pulumi.StringPtrInput
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter ProfileFileFilterPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Gmail. The structure of `gmail` block is documented below.
	Gmail ProfileGmailPtrInput
	// IMAP. The structure of `imap` block is documented below.
	Imap ProfileImapPtrInput
	// MAPI. The structure of `mapi` block is documented below.
	Mapi ProfileMapiPtrInput
	// MSN Hotmail. The structure of `msnHotmail` block is documented below.
	MsnHotmail ProfileMsnHotmailPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// Options.
	Options pulumi.StringPtrInput
	// Other supported webmails. The structure of `otherWebmails` block is documented below.
	OtherWebmails ProfileOtherWebmailsPtrInput
	// POP3. The structure of `pop3` block is documented below.
	Pop3 ProfilePop3PtrInput
	// Replacement message group.
	ReplacemsgGroup pulumi.StringPtrInput
	// SMTP. The structure of `smtp` block is documented below.
	Smtp ProfileSmtpPtrInput
	// Anti-spam block/allow list table ID.
	SpamBalTable pulumi.IntPtrInput
	// Anti-spam black/white list table ID.
	SpamBwlTable pulumi.IntPtrInput
	// Anti-spam banned word table ID.
	SpamBwordTable pulumi.IntPtrInput
	// Spam banned word threshold.
	SpamBwordThreshold pulumi.IntPtrInput
	// Enable/disable spam filtering. Valid values: `enable`, `disable`.
	SpamFiltering pulumi.StringPtrInput
	// Anti-spam IP trust table ID.
	SpamIptrustTable pulumi.IntPtrInput
	// Enable/disable spam logging for email filtering. Valid values: `disable`, `enable`.
	SpamLog pulumi.StringPtrInput
	// Enable/disable logging FortiGuard spam response. Valid values: `disable`, `enable`.
	SpamLogFortiguardResponse pulumi.StringPtrInput
	// Anti-spam MIME header table ID.
	SpamMheaderTable pulumi.IntPtrInput
	// Anti-spam DNSBL table ID.
	SpamRblTable pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Yahoo! Mail. The structure of `yahooMail` block is documented below.
	YahooMail ProfileYahooMailPtrInput
}

func (ProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileState)(nil)).Elem()
}

type profileArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Enable/disable external Email inspection. Valid values: `enable`, `disable`.
	External *string `pulumi:"external"`
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet *string `pulumi:"featureSet"`
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter *ProfileFileFilter `pulumi:"fileFilter"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Gmail. The structure of `gmail` block is documented below.
	Gmail *ProfileGmail `pulumi:"gmail"`
	// IMAP. The structure of `imap` block is documented below.
	Imap *ProfileImap `pulumi:"imap"`
	// MAPI. The structure of `mapi` block is documented below.
	Mapi *ProfileMapi `pulumi:"mapi"`
	// MSN Hotmail. The structure of `msnHotmail` block is documented below.
	MsnHotmail *ProfileMsnHotmail `pulumi:"msnHotmail"`
	// Profile name.
	Name *string `pulumi:"name"`
	// Options.
	Options *string `pulumi:"options"`
	// Other supported webmails. The structure of `otherWebmails` block is documented below.
	OtherWebmails *ProfileOtherWebmails `pulumi:"otherWebmails"`
	// POP3. The structure of `pop3` block is documented below.
	Pop3 *ProfilePop3 `pulumi:"pop3"`
	// Replacement message group.
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// SMTP. The structure of `smtp` block is documented below.
	Smtp *ProfileSmtp `pulumi:"smtp"`
	// Anti-spam block/allow list table ID.
	SpamBalTable *int `pulumi:"spamBalTable"`
	// Anti-spam black/white list table ID.
	SpamBwlTable *int `pulumi:"spamBwlTable"`
	// Anti-spam banned word table ID.
	SpamBwordTable *int `pulumi:"spamBwordTable"`
	// Spam banned word threshold.
	SpamBwordThreshold *int `pulumi:"spamBwordThreshold"`
	// Enable/disable spam filtering. Valid values: `enable`, `disable`.
	SpamFiltering *string `pulumi:"spamFiltering"`
	// Anti-spam IP trust table ID.
	SpamIptrustTable *int `pulumi:"spamIptrustTable"`
	// Enable/disable spam logging for email filtering. Valid values: `disable`, `enable`.
	SpamLog *string `pulumi:"spamLog"`
	// Enable/disable logging FortiGuard spam response. Valid values: `disable`, `enable`.
	SpamLogFortiguardResponse *string `pulumi:"spamLogFortiguardResponse"`
	// Anti-spam MIME header table ID.
	SpamMheaderTable *int `pulumi:"spamMheaderTable"`
	// Anti-spam DNSBL table ID.
	SpamRblTable *int `pulumi:"spamRblTable"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Yahoo! Mail. The structure of `yahooMail` block is documented below.
	YahooMail *ProfileYahooMail `pulumi:"yahooMail"`
}

// The set of arguments for constructing a Profile resource.
type ProfileArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Enable/disable external Email inspection. Valid values: `enable`, `disable`.
	External pulumi.StringPtrInput
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet pulumi.StringPtrInput
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter ProfileFileFilterPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Gmail. The structure of `gmail` block is documented below.
	Gmail ProfileGmailPtrInput
	// IMAP. The structure of `imap` block is documented below.
	Imap ProfileImapPtrInput
	// MAPI. The structure of `mapi` block is documented below.
	Mapi ProfileMapiPtrInput
	// MSN Hotmail. The structure of `msnHotmail` block is documented below.
	MsnHotmail ProfileMsnHotmailPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// Options.
	Options pulumi.StringPtrInput
	// Other supported webmails. The structure of `otherWebmails` block is documented below.
	OtherWebmails ProfileOtherWebmailsPtrInput
	// POP3. The structure of `pop3` block is documented below.
	Pop3 ProfilePop3PtrInput
	// Replacement message group.
	ReplacemsgGroup pulumi.StringPtrInput
	// SMTP. The structure of `smtp` block is documented below.
	Smtp ProfileSmtpPtrInput
	// Anti-spam block/allow list table ID.
	SpamBalTable pulumi.IntPtrInput
	// Anti-spam black/white list table ID.
	SpamBwlTable pulumi.IntPtrInput
	// Anti-spam banned word table ID.
	SpamBwordTable pulumi.IntPtrInput
	// Spam banned word threshold.
	SpamBwordThreshold pulumi.IntPtrInput
	// Enable/disable spam filtering. Valid values: `enable`, `disable`.
	SpamFiltering pulumi.StringPtrInput
	// Anti-spam IP trust table ID.
	SpamIptrustTable pulumi.IntPtrInput
	// Enable/disable spam logging for email filtering. Valid values: `disable`, `enable`.
	SpamLog pulumi.StringPtrInput
	// Enable/disable logging FortiGuard spam response. Valid values: `disable`, `enable`.
	SpamLogFortiguardResponse pulumi.StringPtrInput
	// Anti-spam MIME header table ID.
	SpamMheaderTable pulumi.IntPtrInput
	// Anti-spam DNSBL table ID.
	SpamRblTable pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Yahoo! Mail. The structure of `yahooMail` block is documented below.
	YahooMail ProfileYahooMailPtrInput
}

func (ProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileArgs)(nil)).Elem()
}

type ProfileInput interface {
	pulumi.Input

	ToProfileOutput() ProfileOutput
	ToProfileOutputWithContext(ctx context.Context) ProfileOutput
}

func (*Profile) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (i *Profile) ToProfileOutput() ProfileOutput {
	return i.ToProfileOutputWithContext(context.Background())
}

func (i *Profile) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileOutput)
}

// ProfileArrayInput is an input type that accepts ProfileArray and ProfileArrayOutput values.
// You can construct a concrete instance of `ProfileArrayInput` via:
//
//	ProfileArray{ ProfileArgs{...} }
type ProfileArrayInput interface {
	pulumi.Input

	ToProfileArrayOutput() ProfileArrayOutput
	ToProfileArrayOutputWithContext(context.Context) ProfileArrayOutput
}

type ProfileArray []ProfileInput

func (ProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Profile)(nil)).Elem()
}

func (i ProfileArray) ToProfileArrayOutput() ProfileArrayOutput {
	return i.ToProfileArrayOutputWithContext(context.Background())
}

func (i ProfileArray) ToProfileArrayOutputWithContext(ctx context.Context) ProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileArrayOutput)
}

// ProfileMapInput is an input type that accepts ProfileMap and ProfileMapOutput values.
// You can construct a concrete instance of `ProfileMapInput` via:
//
//	ProfileMap{ "key": ProfileArgs{...} }
type ProfileMapInput interface {
	pulumi.Input

	ToProfileMapOutput() ProfileMapOutput
	ToProfileMapOutputWithContext(context.Context) ProfileMapOutput
}

type ProfileMap map[string]ProfileInput

func (ProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Profile)(nil)).Elem()
}

func (i ProfileMap) ToProfileMapOutput() ProfileMapOutput {
	return i.ToProfileMapOutputWithContext(context.Background())
}

func (i ProfileMap) ToProfileMapOutputWithContext(ctx context.Context) ProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileMapOutput)
}

type ProfileOutput struct{ *pulumi.OutputState }

func (ProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (o ProfileOutput) ToProfileOutput() ProfileOutput {
	return o
}

func (o ProfileOutput) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return o
}

// Comment.
func (o ProfileOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Enable/disable external Email inspection. Valid values: `enable`, `disable`.
func (o ProfileOutput) External() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.External }).(pulumi.StringOutput)
}

// Flow/proxy feature set. Valid values: `flow`, `proxy`.
func (o ProfileOutput) FeatureSet() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.FeatureSet }).(pulumi.StringOutput)
}

// File filter. The structure of `fileFilter` block is documented below.
func (o ProfileOutput) FileFilter() ProfileFileFilterOutput {
	return o.ApplyT(func(v *Profile) ProfileFileFilterOutput { return v.FileFilter }).(ProfileFileFilterOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o ProfileOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Gmail. The structure of `gmail` block is documented below.
func (o ProfileOutput) Gmail() ProfileGmailOutput {
	return o.ApplyT(func(v *Profile) ProfileGmailOutput { return v.Gmail }).(ProfileGmailOutput)
}

// IMAP. The structure of `imap` block is documented below.
func (o ProfileOutput) Imap() ProfileImapOutput {
	return o.ApplyT(func(v *Profile) ProfileImapOutput { return v.Imap }).(ProfileImapOutput)
}

// MAPI. The structure of `mapi` block is documented below.
func (o ProfileOutput) Mapi() ProfileMapiOutput {
	return o.ApplyT(func(v *Profile) ProfileMapiOutput { return v.Mapi }).(ProfileMapiOutput)
}

// MSN Hotmail. The structure of `msnHotmail` block is documented below.
func (o ProfileOutput) MsnHotmail() ProfileMsnHotmailOutput {
	return o.ApplyT(func(v *Profile) ProfileMsnHotmailOutput { return v.MsnHotmail }).(ProfileMsnHotmailOutput)
}

// Profile name.
func (o ProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Options.
func (o ProfileOutput) Options() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Options }).(pulumi.StringOutput)
}

// Other supported webmails. The structure of `otherWebmails` block is documented below.
func (o ProfileOutput) OtherWebmails() ProfileOtherWebmailsOutput {
	return o.ApplyT(func(v *Profile) ProfileOtherWebmailsOutput { return v.OtherWebmails }).(ProfileOtherWebmailsOutput)
}

// POP3. The structure of `pop3` block is documented below.
func (o ProfileOutput) Pop3() ProfilePop3Output {
	return o.ApplyT(func(v *Profile) ProfilePop3Output { return v.Pop3 }).(ProfilePop3Output)
}

// Replacement message group.
func (o ProfileOutput) ReplacemsgGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.ReplacemsgGroup }).(pulumi.StringOutput)
}

// SMTP. The structure of `smtp` block is documented below.
func (o ProfileOutput) Smtp() ProfileSmtpOutput {
	return o.ApplyT(func(v *Profile) ProfileSmtpOutput { return v.Smtp }).(ProfileSmtpOutput)
}

// Anti-spam block/allow list table ID.
func (o ProfileOutput) SpamBalTable() pulumi.IntOutput {
	return o.ApplyT(func(v *Profile) pulumi.IntOutput { return v.SpamBalTable }).(pulumi.IntOutput)
}

// Anti-spam black/white list table ID.
func (o ProfileOutput) SpamBwlTable() pulumi.IntOutput {
	return o.ApplyT(func(v *Profile) pulumi.IntOutput { return v.SpamBwlTable }).(pulumi.IntOutput)
}

// Anti-spam banned word table ID.
func (o ProfileOutput) SpamBwordTable() pulumi.IntOutput {
	return o.ApplyT(func(v *Profile) pulumi.IntOutput { return v.SpamBwordTable }).(pulumi.IntOutput)
}

// Spam banned word threshold.
func (o ProfileOutput) SpamBwordThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *Profile) pulumi.IntOutput { return v.SpamBwordThreshold }).(pulumi.IntOutput)
}

// Enable/disable spam filtering. Valid values: `enable`, `disable`.
func (o ProfileOutput) SpamFiltering() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.SpamFiltering }).(pulumi.StringOutput)
}

// Anti-spam IP trust table ID.
func (o ProfileOutput) SpamIptrustTable() pulumi.IntOutput {
	return o.ApplyT(func(v *Profile) pulumi.IntOutput { return v.SpamIptrustTable }).(pulumi.IntOutput)
}

// Enable/disable spam logging for email filtering. Valid values: `disable`, `enable`.
func (o ProfileOutput) SpamLog() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.SpamLog }).(pulumi.StringOutput)
}

// Enable/disable logging FortiGuard spam response. Valid values: `disable`, `enable`.
func (o ProfileOutput) SpamLogFortiguardResponse() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.SpamLogFortiguardResponse }).(pulumi.StringOutput)
}

// Anti-spam MIME header table ID.
func (o ProfileOutput) SpamMheaderTable() pulumi.IntOutput {
	return o.ApplyT(func(v *Profile) pulumi.IntOutput { return v.SpamMheaderTable }).(pulumi.IntOutput)
}

// Anti-spam DNSBL table ID.
func (o ProfileOutput) SpamRblTable() pulumi.IntOutput {
	return o.ApplyT(func(v *Profile) pulumi.IntOutput { return v.SpamRblTable }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ProfileOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

// Yahoo! Mail. The structure of `yahooMail` block is documented below.
func (o ProfileOutput) YahooMail() ProfileYahooMailOutput {
	return o.ApplyT(func(v *Profile) ProfileYahooMailOutput { return v.YahooMail }).(ProfileYahooMailOutput)
}

type ProfileArrayOutput struct{ *pulumi.OutputState }

func (ProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Profile)(nil)).Elem()
}

func (o ProfileArrayOutput) ToProfileArrayOutput() ProfileArrayOutput {
	return o
}

func (o ProfileArrayOutput) ToProfileArrayOutputWithContext(ctx context.Context) ProfileArrayOutput {
	return o
}

func (o ProfileArrayOutput) Index(i pulumi.IntInput) ProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Profile {
		return vs[0].([]*Profile)[vs[1].(int)]
	}).(ProfileOutput)
}

type ProfileMapOutput struct{ *pulumi.OutputState }

func (ProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Profile)(nil)).Elem()
}

func (o ProfileMapOutput) ToProfileMapOutput() ProfileMapOutput {
	return o
}

func (o ProfileMapOutput) ToProfileMapOutputWithContext(ctx context.Context) ProfileMapOutput {
	return o
}

func (o ProfileMapOutput) MapIndex(k pulumi.StringInput) ProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Profile {
		return vs[0].(map[string]*Profile)[vs[1].(string)]
	}).(ProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileInput)(nil)).Elem(), &Profile{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileArrayInput)(nil)).Elem(), ProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileMapInput)(nil)).Elem(), ProfileMap{})
	pulumi.RegisterOutputType(ProfileOutput{})
	pulumi.RegisterOutputType(ProfileArrayOutput{})
	pulumi.RegisterOutputType(ProfileMapOutput{})
}
