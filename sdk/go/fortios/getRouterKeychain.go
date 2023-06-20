// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios router keychain
func LookupRouterKeychain(ctx *pulumi.Context, args *LookupRouterKeychainArgs, opts ...pulumi.InvokeOption) (*LookupRouterKeychainResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterKeychainResult
	err := ctx.Invoke("fortios:index/getRouterKeychain:getRouterKeychain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouterKeychain.
type LookupRouterKeychainArgs struct {
	// Specify the name of the desired router keychain.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getRouterKeychain.
type LookupRouterKeychainResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Configuration method to edit key settings. The structure of `key` block is documented below.
	Keys []GetRouterKeychainKey `pulumi:"keys"`
	// Key-chain name.
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupRouterKeychainOutput(ctx *pulumi.Context, args LookupRouterKeychainOutputArgs, opts ...pulumi.InvokeOption) LookupRouterKeychainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterKeychainResult, error) {
			args := v.(LookupRouterKeychainArgs)
			r, err := LookupRouterKeychain(ctx, &args, opts...)
			var s LookupRouterKeychainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouterKeychainResultOutput)
}

// A collection of arguments for invoking getRouterKeychain.
type LookupRouterKeychainOutputArgs struct {
	// Specify the name of the desired router keychain.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterKeychainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterKeychainArgs)(nil)).Elem()
}

// A collection of values returned by getRouterKeychain.
type LookupRouterKeychainResultOutput struct{ *pulumi.OutputState }

func (LookupRouterKeychainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterKeychainResult)(nil)).Elem()
}

func (o LookupRouterKeychainResultOutput) ToLookupRouterKeychainResultOutput() LookupRouterKeychainResultOutput {
	return o
}

func (o LookupRouterKeychainResultOutput) ToLookupRouterKeychainResultOutputWithContext(ctx context.Context) LookupRouterKeychainResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterKeychainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterKeychainResult) string { return v.Id }).(pulumi.StringOutput)
}

// Configuration method to edit key settings. The structure of `key` block is documented below.
func (o LookupRouterKeychainResultOutput) Keys() GetRouterKeychainKeyArrayOutput {
	return o.ApplyT(func(v LookupRouterKeychainResult) []GetRouterKeychainKey { return v.Keys }).(GetRouterKeychainKeyArrayOutput)
}

// Key-chain name.
func (o LookupRouterKeychainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterKeychainResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRouterKeychainResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterKeychainResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterKeychainResultOutput{})
}