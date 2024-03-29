// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on an fortios firewall internetservicecustomgroup
func LookupInternetservicecustomgroup(ctx *pulumi.Context, args *LookupInternetservicecustomgroupArgs, opts ...pulumi.InvokeOption) (*LookupInternetservicecustomgroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInternetservicecustomgroupResult
	err := ctx.Invoke("fortios:firewall/getInternetservicecustomgroup:getInternetservicecustomgroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInternetservicecustomgroup.
type LookupInternetservicecustomgroupArgs struct {
	// Specify the name of the desired firewall internetservicecustomgroup.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getInternetservicecustomgroup.
type LookupInternetservicecustomgroupResult struct {
	// Comment.
	Comment string `pulumi:"comment"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Custom Internet Service group members. The structure of `member` block is documented below.
	Members []GetInternetservicecustomgroupMember `pulumi:"members"`
	// Group member name.
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupInternetservicecustomgroupOutput(ctx *pulumi.Context, args LookupInternetservicecustomgroupOutputArgs, opts ...pulumi.InvokeOption) LookupInternetservicecustomgroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInternetservicecustomgroupResult, error) {
			args := v.(LookupInternetservicecustomgroupArgs)
			r, err := LookupInternetservicecustomgroup(ctx, &args, opts...)
			var s LookupInternetservicecustomgroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInternetservicecustomgroupResultOutput)
}

// A collection of arguments for invoking getInternetservicecustomgroup.
type LookupInternetservicecustomgroupOutputArgs struct {
	// Specify the name of the desired firewall internetservicecustomgroup.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupInternetservicecustomgroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInternetservicecustomgroupArgs)(nil)).Elem()
}

// A collection of values returned by getInternetservicecustomgroup.
type LookupInternetservicecustomgroupResultOutput struct{ *pulumi.OutputState }

func (LookupInternetservicecustomgroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInternetservicecustomgroupResult)(nil)).Elem()
}

func (o LookupInternetservicecustomgroupResultOutput) ToLookupInternetservicecustomgroupResultOutput() LookupInternetservicecustomgroupResultOutput {
	return o
}

func (o LookupInternetservicecustomgroupResultOutput) ToLookupInternetservicecustomgroupResultOutputWithContext(ctx context.Context) LookupInternetservicecustomgroupResultOutput {
	return o
}

// Comment.
func (o LookupInternetservicecustomgroupResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternetservicecustomgroupResult) string { return v.Comment }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInternetservicecustomgroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternetservicecustomgroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Custom Internet Service group members. The structure of `member` block is documented below.
func (o LookupInternetservicecustomgroupResultOutput) Members() GetInternetservicecustomgroupMemberArrayOutput {
	return o.ApplyT(func(v LookupInternetservicecustomgroupResult) []GetInternetservicecustomgroupMember { return v.Members }).(GetInternetservicecustomgroupMemberArrayOutput)
}

// Group member name.
func (o LookupInternetservicecustomgroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternetservicecustomgroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupInternetservicecustomgroupResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInternetservicecustomgroupResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInternetservicecustomgroupResultOutput{})
}
