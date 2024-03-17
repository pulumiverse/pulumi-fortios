// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wildcardfqdn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on an fortios firewallwildcardfqdn custom
func LookupCustom(ctx *pulumi.Context, args *LookupCustomArgs, opts ...pulumi.InvokeOption) (*LookupCustomResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCustomResult
	err := ctx.Invoke("fortios:firewall/wildcardfqdn/getCustom:getCustom", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCustom.
type LookupCustomArgs struct {
	// Specify the name of the desired firewallwildcardfqdn custom.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getCustom.
type LookupCustomResult struct {
	// GUI icon color.
	Color int `pulumi:"color"`
	// Comment.
	Comment string `pulumi:"comment"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Address name.
	Name string `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid      string  `pulumi:"uuid"`
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable address visibility.
	Visibility string `pulumi:"visibility"`
	// Wildcard FQDN.
	WildcardFqdn string `pulumi:"wildcardFqdn"`
}

func LookupCustomOutput(ctx *pulumi.Context, args LookupCustomOutputArgs, opts ...pulumi.InvokeOption) LookupCustomResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomResult, error) {
			args := v.(LookupCustomArgs)
			r, err := LookupCustom(ctx, &args, opts...)
			var s LookupCustomResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomResultOutput)
}

// A collection of arguments for invoking getCustom.
type LookupCustomOutputArgs struct {
	// Specify the name of the desired firewallwildcardfqdn custom.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupCustomOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomArgs)(nil)).Elem()
}

// A collection of values returned by getCustom.
type LookupCustomResultOutput struct{ *pulumi.OutputState }

func (LookupCustomResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomResult)(nil)).Elem()
}

func (o LookupCustomResultOutput) ToLookupCustomResultOutput() LookupCustomResultOutput {
	return o
}

func (o LookupCustomResultOutput) ToLookupCustomResultOutputWithContext(ctx context.Context) LookupCustomResultOutput {
	return o
}

// GUI icon color.
func (o LookupCustomResultOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCustomResult) int { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o LookupCustomResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomResult) string { return v.Comment }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCustomResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomResult) string { return v.Id }).(pulumi.StringOutput)
}

// Address name.
func (o LookupCustomResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomResult) string { return v.Name }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o LookupCustomResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupCustomResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable address visibility.
func (o LookupCustomResultOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomResult) string { return v.Visibility }).(pulumi.StringOutput)
}

// Wildcard FQDN.
func (o LookupCustomResultOutput) WildcardFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomResult) string { return v.WildcardFqdn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomResultOutput{})
}
