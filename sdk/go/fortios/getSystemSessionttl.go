// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system sessionttl
func LookupSystemSessionttl(ctx *pulumi.Context, args *LookupSystemSessionttlArgs, opts ...pulumi.InvokeOption) (*LookupSystemSessionttlResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemSessionttlResult
	err := ctx.Invoke("fortios:index/getSystemSessionttl:getSystemSessionttl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemSessionttl.
type LookupSystemSessionttlArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemSessionttl.
type LookupSystemSessionttlResult struct {
	// Default timeout.
	Default string `pulumi:"default"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Session TTL port. The structure of `port` block is documented below.
	Ports     []GetSystemSessionttlPort `pulumi:"ports"`
	Vdomparam *string                   `pulumi:"vdomparam"`
}

func LookupSystemSessionttlOutput(ctx *pulumi.Context, args LookupSystemSessionttlOutputArgs, opts ...pulumi.InvokeOption) LookupSystemSessionttlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemSessionttlResult, error) {
			args := v.(LookupSystemSessionttlArgs)
			r, err := LookupSystemSessionttl(ctx, &args, opts...)
			var s LookupSystemSessionttlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemSessionttlResultOutput)
}

// A collection of arguments for invoking getSystemSessionttl.
type LookupSystemSessionttlOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemSessionttlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemSessionttlArgs)(nil)).Elem()
}

// A collection of values returned by getSystemSessionttl.
type LookupSystemSessionttlResultOutput struct{ *pulumi.OutputState }

func (LookupSystemSessionttlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemSessionttlResult)(nil)).Elem()
}

func (o LookupSystemSessionttlResultOutput) ToLookupSystemSessionttlResultOutput() LookupSystemSessionttlResultOutput {
	return o
}

func (o LookupSystemSessionttlResultOutput) ToLookupSystemSessionttlResultOutputWithContext(ctx context.Context) LookupSystemSessionttlResultOutput {
	return o
}

// Default timeout.
func (o LookupSystemSessionttlResultOutput) Default() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSessionttlResult) string { return v.Default }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemSessionttlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSessionttlResult) string { return v.Id }).(pulumi.StringOutput)
}

// Session TTL port. The structure of `port` block is documented below.
func (o LookupSystemSessionttlResultOutput) Ports() GetSystemSessionttlPortArrayOutput {
	return o.ApplyT(func(v LookupSystemSessionttlResult) []GetSystemSessionttlPort { return v.Ports }).(GetSystemSessionttlPortArrayOutput)
}

func (o LookupSystemSessionttlResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemSessionttlResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemSessionttlResultOutput{})
}