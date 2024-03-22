// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package user

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on an fortios user saml
func LookupSaml(ctx *pulumi.Context, args *LookupSamlArgs, opts ...pulumi.InvokeOption) (*LookupSamlResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSamlResult
	err := ctx.Invoke("fortios:user/getSaml:getSaml", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSaml.
type LookupSamlArgs struct {
	// Specify the name of the desired user saml.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSaml.
type LookupSamlResult struct {
	// Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable).
	AdfsClaim string `pulumi:"adfsClaim"`
	// URL to verify authentication.
	AuthUrl string `pulumi:"authUrl"`
	// Certificate to sign SAML messages.
	Cert string `pulumi:"cert"`
	// Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
	ClockTolerance int `pulumi:"clockTolerance"`
	// Digest Method Algorithm. (default = sha1).
	DigestMethod string `pulumi:"digestMethod"`
	// SP entity ID.
	EntityId string `pulumi:"entityId"`
	// Group claim in assertion statement.
	GroupClaimType string `pulumi:"groupClaimType"`
	// Group name in assertion statement.
	GroupName string `pulumi:"groupName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IDP Certificate name.
	IdpCert string `pulumi:"idpCert"`
	// IDP entity ID.
	IdpEntityId string `pulumi:"idpEntityId"`
	// IDP single logout url.
	IdpSingleLogoutUrl string `pulumi:"idpSingleLogoutUrl"`
	// IDP single sign-on URL.
	IdpSingleSignOnUrl string `pulumi:"idpSingleSignOnUrl"`
	// Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes).
	LimitRelaystate string `pulumi:"limitRelaystate"`
	// SAML server entry name.
	Name string `pulumi:"name"`
	// Enable/disable signalling of IDP to force user re-authentication (default = disable).
	Reauth string `pulumi:"reauth"`
	// SP single logout URL.
	SingleLogoutUrl string `pulumi:"singleLogoutUrl"`
	// SP single sign-on URL.
	SingleSignOnUrl string `pulumi:"singleSignOnUrl"`
	// User name claim in assertion statement.
	UserClaimType string `pulumi:"userClaimType"`
	// User name in assertion statement.
	UserName  string  `pulumi:"userName"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSamlOutput(ctx *pulumi.Context, args LookupSamlOutputArgs, opts ...pulumi.InvokeOption) LookupSamlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSamlResult, error) {
			args := v.(LookupSamlArgs)
			r, err := LookupSaml(ctx, &args, opts...)
			var s LookupSamlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSamlResultOutput)
}

// A collection of arguments for invoking getSaml.
type LookupSamlOutputArgs struct {
	// Specify the name of the desired user saml.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSamlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSamlArgs)(nil)).Elem()
}

// A collection of values returned by getSaml.
type LookupSamlResultOutput struct{ *pulumi.OutputState }

func (LookupSamlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSamlResult)(nil)).Elem()
}

func (o LookupSamlResultOutput) ToLookupSamlResultOutput() LookupSamlResultOutput {
	return o
}

func (o LookupSamlResultOutput) ToLookupSamlResultOutputWithContext(ctx context.Context) LookupSamlResultOutput {
	return o
}

// Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable).
func (o LookupSamlResultOutput) AdfsClaim() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSamlResult) string { return v.AdfsClaim }).(pulumi.StringOutput)
}

// URL to verify authentication.
func (o LookupSamlResultOutput) AuthUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSamlResult) string { return v.AuthUrl }).(pulumi.StringOutput)
}

// Certificate to sign SAML messages.
func (o LookupSamlResultOutput) Cert() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSamlResult) string { return v.Cert }).(pulumi.StringOutput)
}

// Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
func (o LookupSamlResultOutput) ClockTolerance() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSamlResult) int { return v.ClockTolerance }).(pulumi.IntOutput)
}

// Digest Method Algorithm. (default = sha1).
func (o LookupSamlResultOutput) DigestMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSamlResult) string { return v.DigestMethod }).(pulumi.StringOutput)
}

// SP entity ID.
func (o LookupSamlResultOutput) EntityId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSamlResult) string { return v.EntityId }).(pulumi.StringOutput)
}

// Group claim in assertion statement.
func (o LookupSamlResultOutput) GroupClaimType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSamlResult) string { return v.GroupClaimType }).(pulumi.StringOutput)
}

// Group name in assertion statement.
func (o LookupSamlResultOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSamlResult) string { return v.GroupName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSamlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSamlResult) string { return v.Id }).(pulumi.StringOutput)
}

// IDP Certificate name.
func (o LookupSamlResultOutput) IdpCert() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSamlResult) string { return v.IdpCert }).(pulumi.StringOutput)
}

// IDP entity ID.
func (o LookupSamlResultOutput) IdpEntityId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSamlResult) string { return v.IdpEntityId }).(pulumi.StringOutput)
}

// IDP single logout url.
func (o LookupSamlResultOutput) IdpSingleLogoutUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSamlResult) string { return v.IdpSingleLogoutUrl }).(pulumi.StringOutput)
}

// IDP single sign-on URL.
func (o LookupSamlResultOutput) IdpSingleSignOnUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSamlResult) string { return v.IdpSingleSignOnUrl }).(pulumi.StringOutput)
}

// Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes).
func (o LookupSamlResultOutput) LimitRelaystate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSamlResult) string { return v.LimitRelaystate }).(pulumi.StringOutput)
}

// SAML server entry name.
func (o LookupSamlResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSamlResult) string { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable signalling of IDP to force user re-authentication (default = disable).
func (o LookupSamlResultOutput) Reauth() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSamlResult) string { return v.Reauth }).(pulumi.StringOutput)
}

// SP single logout URL.
func (o LookupSamlResultOutput) SingleLogoutUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSamlResult) string { return v.SingleLogoutUrl }).(pulumi.StringOutput)
}

// SP single sign-on URL.
func (o LookupSamlResultOutput) SingleSignOnUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSamlResult) string { return v.SingleSignOnUrl }).(pulumi.StringOutput)
}

// User name claim in assertion statement.
func (o LookupSamlResultOutput) UserClaimType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSamlResult) string { return v.UserClaimType }).(pulumi.StringOutput)
}

// User name in assertion statement.
func (o LookupSamlResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSamlResult) string { return v.UserName }).(pulumi.StringOutput)
}

func (o LookupSamlResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSamlResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSamlResultOutput{})
}
