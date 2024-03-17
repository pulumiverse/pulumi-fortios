// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on fortios system passwordpolicyguestadmin
func LookupPasswordpolicyguestadmin(ctx *pulumi.Context, args *LookupPasswordpolicyguestadminArgs, opts ...pulumi.InvokeOption) (*LookupPasswordpolicyguestadminResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPasswordpolicyguestadminResult
	err := ctx.Invoke("fortios:system/getPasswordpolicyguestadmin:getPasswordpolicyguestadmin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPasswordpolicyguestadmin.
type LookupPasswordpolicyguestadminArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getPasswordpolicyguestadmin.
type LookupPasswordpolicyguestadminResult struct {
	// Guest administrator to which this password policy applies.
	ApplyTo string `pulumi:"applyTo"`
	// Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled).
	Change4Characters string `pulumi:"change4Characters"`
	// Number of days after which passwords expire (1 - 999 days, default = 90).
	ExpireDay int `pulumi:"expireDay"`
	// Enable/disable password expiration.
	ExpireStatus string `pulumi:"expireStatus"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
	MinChangeCharacters int `pulumi:"minChangeCharacters"`
	// Minimum number of lowercase characters in password (0 - 128, default = 0).
	MinLowerCaseLetter int `pulumi:"minLowerCaseLetter"`
	// Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
	MinNonAlphanumeric int `pulumi:"minNonAlphanumeric"`
	// Minimum number of numeric characters in password (0 - 128, default = 0).
	MinNumber int `pulumi:"minNumber"`
	// Minimum number of uppercase characters in password (0 - 128, default = 0).
	MinUpperCaseLetter int `pulumi:"minUpperCaseLetter"`
	// Minimum password length (8 - 128, default = 8).
	MinimumLength int `pulumi:"minimumLength"`
	// Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides).
	ReusePassword string `pulumi:"reusePassword"`
	// Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.
	Status    string  `pulumi:"status"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupPasswordpolicyguestadminOutput(ctx *pulumi.Context, args LookupPasswordpolicyguestadminOutputArgs, opts ...pulumi.InvokeOption) LookupPasswordpolicyguestadminResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPasswordpolicyguestadminResult, error) {
			args := v.(LookupPasswordpolicyguestadminArgs)
			r, err := LookupPasswordpolicyguestadmin(ctx, &args, opts...)
			var s LookupPasswordpolicyguestadminResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPasswordpolicyguestadminResultOutput)
}

// A collection of arguments for invoking getPasswordpolicyguestadmin.
type LookupPasswordpolicyguestadminOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupPasswordpolicyguestadminOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPasswordpolicyguestadminArgs)(nil)).Elem()
}

// A collection of values returned by getPasswordpolicyguestadmin.
type LookupPasswordpolicyguestadminResultOutput struct{ *pulumi.OutputState }

func (LookupPasswordpolicyguestadminResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPasswordpolicyguestadminResult)(nil)).Elem()
}

func (o LookupPasswordpolicyguestadminResultOutput) ToLookupPasswordpolicyguestadminResultOutput() LookupPasswordpolicyguestadminResultOutput {
	return o
}

func (o LookupPasswordpolicyguestadminResultOutput) ToLookupPasswordpolicyguestadminResultOutputWithContext(ctx context.Context) LookupPasswordpolicyguestadminResultOutput {
	return o
}

// Guest administrator to which this password policy applies.
func (o LookupPasswordpolicyguestadminResultOutput) ApplyTo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPasswordpolicyguestadminResult) string { return v.ApplyTo }).(pulumi.StringOutput)
}

// Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled).
func (o LookupPasswordpolicyguestadminResultOutput) Change4Characters() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPasswordpolicyguestadminResult) string { return v.Change4Characters }).(pulumi.StringOutput)
}

// Number of days after which passwords expire (1 - 999 days, default = 90).
func (o LookupPasswordpolicyguestadminResultOutput) ExpireDay() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPasswordpolicyguestadminResult) int { return v.ExpireDay }).(pulumi.IntOutput)
}

// Enable/disable password expiration.
func (o LookupPasswordpolicyguestadminResultOutput) ExpireStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPasswordpolicyguestadminResult) string { return v.ExpireStatus }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPasswordpolicyguestadminResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPasswordpolicyguestadminResult) string { return v.Id }).(pulumi.StringOutput)
}

// Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
func (o LookupPasswordpolicyguestadminResultOutput) MinChangeCharacters() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPasswordpolicyguestadminResult) int { return v.MinChangeCharacters }).(pulumi.IntOutput)
}

// Minimum number of lowercase characters in password (0 - 128, default = 0).
func (o LookupPasswordpolicyguestadminResultOutput) MinLowerCaseLetter() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPasswordpolicyguestadminResult) int { return v.MinLowerCaseLetter }).(pulumi.IntOutput)
}

// Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
func (o LookupPasswordpolicyguestadminResultOutput) MinNonAlphanumeric() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPasswordpolicyguestadminResult) int { return v.MinNonAlphanumeric }).(pulumi.IntOutput)
}

// Minimum number of numeric characters in password (0 - 128, default = 0).
func (o LookupPasswordpolicyguestadminResultOutput) MinNumber() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPasswordpolicyguestadminResult) int { return v.MinNumber }).(pulumi.IntOutput)
}

// Minimum number of uppercase characters in password (0 - 128, default = 0).
func (o LookupPasswordpolicyguestadminResultOutput) MinUpperCaseLetter() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPasswordpolicyguestadminResult) int { return v.MinUpperCaseLetter }).(pulumi.IntOutput)
}

// Minimum password length (8 - 128, default = 8).
func (o LookupPasswordpolicyguestadminResultOutput) MinimumLength() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPasswordpolicyguestadminResult) int { return v.MinimumLength }).(pulumi.IntOutput)
}

// Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides).
func (o LookupPasswordpolicyguestadminResultOutput) ReusePassword() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPasswordpolicyguestadminResult) string { return v.ReusePassword }).(pulumi.StringOutput)
}

// Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.
func (o LookupPasswordpolicyguestadminResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPasswordpolicyguestadminResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupPasswordpolicyguestadminResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPasswordpolicyguestadminResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPasswordpolicyguestadminResultOutput{})
}
