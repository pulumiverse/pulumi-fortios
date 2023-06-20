// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall proxyaddrgrp
func LookupFirewallProxyaddrgrp(ctx *pulumi.Context, args *LookupFirewallProxyaddrgrpArgs, opts ...pulumi.InvokeOption) (*LookupFirewallProxyaddrgrpResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallProxyaddrgrpResult
	err := ctx.Invoke("fortios:index/getFirewallProxyaddrgrp:getFirewallProxyaddrgrp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallProxyaddrgrp.
type LookupFirewallProxyaddrgrpArgs struct {
	// Specify the name of the desired firewall proxyaddrgrp.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFirewallProxyaddrgrp.
type LookupFirewallProxyaddrgrpResult struct {
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color int `pulumi:"color"`
	// Optional comments.
	Comment string `pulumi:"comment"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Members of address group. The structure of `member` block is documented below.
	Members []GetFirewallProxyaddrgrpMember `pulumi:"members"`
	// Tag name.
	Name string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []GetFirewallProxyaddrgrpTagging `pulumi:"taggings"`
	// Source or destination address group type.
	Type string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid      string  `pulumi:"uuid"`
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable visibility of the object in the GUI.
	Visibility string `pulumi:"visibility"`
}

func LookupFirewallProxyaddrgrpOutput(ctx *pulumi.Context, args LookupFirewallProxyaddrgrpOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallProxyaddrgrpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallProxyaddrgrpResult, error) {
			args := v.(LookupFirewallProxyaddrgrpArgs)
			r, err := LookupFirewallProxyaddrgrp(ctx, &args, opts...)
			var s LookupFirewallProxyaddrgrpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallProxyaddrgrpResultOutput)
}

// A collection of arguments for invoking getFirewallProxyaddrgrp.
type LookupFirewallProxyaddrgrpOutputArgs struct {
	// Specify the name of the desired firewall proxyaddrgrp.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallProxyaddrgrpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallProxyaddrgrpArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallProxyaddrgrp.
type LookupFirewallProxyaddrgrpResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallProxyaddrgrpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallProxyaddrgrpResult)(nil)).Elem()
}

func (o LookupFirewallProxyaddrgrpResultOutput) ToLookupFirewallProxyaddrgrpResultOutput() LookupFirewallProxyaddrgrpResultOutput {
	return o
}

func (o LookupFirewallProxyaddrgrpResultOutput) ToLookupFirewallProxyaddrgrpResultOutputWithContext(ctx context.Context) LookupFirewallProxyaddrgrpResultOutput {
	return o
}

// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
func (o LookupFirewallProxyaddrgrpResultOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddrgrpResult) int { return v.Color }).(pulumi.IntOutput)
}

// Optional comments.
func (o LookupFirewallProxyaddrgrpResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddrgrpResult) string { return v.Comment }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallProxyaddrgrpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddrgrpResult) string { return v.Id }).(pulumi.StringOutput)
}

// Members of address group. The structure of `member` block is documented below.
func (o LookupFirewallProxyaddrgrpResultOutput) Members() GetFirewallProxyaddrgrpMemberArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddrgrpResult) []GetFirewallProxyaddrgrpMember { return v.Members }).(GetFirewallProxyaddrgrpMemberArrayOutput)
}

// Tag name.
func (o LookupFirewallProxyaddrgrpResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddrgrpResult) string { return v.Name }).(pulumi.StringOutput)
}

// Config object tagging. The structure of `tagging` block is documented below.
func (o LookupFirewallProxyaddrgrpResultOutput) Taggings() GetFirewallProxyaddrgrpTaggingArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddrgrpResult) []GetFirewallProxyaddrgrpTagging { return v.Taggings }).(GetFirewallProxyaddrgrpTaggingArrayOutput)
}

// Source or destination address group type.
func (o LookupFirewallProxyaddrgrpResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddrgrpResult) string { return v.Type }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o LookupFirewallProxyaddrgrpResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddrgrpResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupFirewallProxyaddrgrpResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddrgrpResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable visibility of the object in the GUI.
func (o LookupFirewallProxyaddrgrpResultOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddrgrpResult) string { return v.Visibility }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallProxyaddrgrpResultOutput{})
}