// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on an fortios firewall proxyaddress
func LookupProxyaddress(ctx *pulumi.Context, args *LookupProxyaddressArgs, opts ...pulumi.InvokeOption) (*LookupProxyaddressResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProxyaddressResult
	err := ctx.Invoke("fortios:firewall/getProxyaddress:getProxyaddress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProxyaddress.
type LookupProxyaddressArgs struct {
	// Specify the name of the desired firewall proxyaddress.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getProxyaddress.
type LookupProxyaddressResult struct {
	// SaaS application. The structure of `application` block is documented below.
	Applications []GetProxyaddressApplication `pulumi:"applications"`
	// Case sensitivity in pattern.
	CaseSensitivity string `pulumi:"caseSensitivity"`
	// Tag category.
	Categories []GetProxyaddressCategory `pulumi:"categories"`
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color int `pulumi:"color"`
	// Optional comments.
	Comment string `pulumi:"comment"`
	// HTTP header regular expression.
	Header string `pulumi:"header"`
	// HTTP header group. The structure of `headerGroup` block is documented below.
	HeaderGroups []GetProxyaddressHeaderGroup `pulumi:"headerGroups"`
	// HTTP header.
	HeaderName string `pulumi:"headerName"`
	// Address object for the host.
	Host string `pulumi:"host"`
	// Host name as a regular expression.
	HostRegex string `pulumi:"hostRegex"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// HTTP request methods to be used.
	Method string `pulumi:"method"`
	// SaaS applicaton name.
	Name string `pulumi:"name"`
	// URL path as a regular expression.
	Path string `pulumi:"path"`
	// Match the query part of the URL as a regular expression.
	Query string `pulumi:"query"`
	// Enable/disable use of referrer field in the HTTP header to match the address.
	Referrer string `pulumi:"referrer"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []GetProxyaddressTagging `pulumi:"taggings"`
	// Proxy address type.
	Type string `pulumi:"type"`
	// Names of browsers to be used as user agent.
	Ua string `pulumi:"ua"`
	// Maximum version of the user agent specified in dotted notation. For example, use 120 with the ua field set to "chrome" to require Google Chrome's maximum version must be 120.
	UaMaxVer string `pulumi:"uaMaxVer"`
	// Minimum version of the user agent specified in dotted notation. For example, use 90.0.1 with the ua field set to "chrome" to require Google Chrome's minimum version must be 90.0.1.
	UaMinVer string `pulumi:"uaMinVer"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid      string  `pulumi:"uuid"`
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable visibility of the object in the GUI.
	Visibility string `pulumi:"visibility"`
}

func LookupProxyaddressOutput(ctx *pulumi.Context, args LookupProxyaddressOutputArgs, opts ...pulumi.InvokeOption) LookupProxyaddressResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProxyaddressResult, error) {
			args := v.(LookupProxyaddressArgs)
			r, err := LookupProxyaddress(ctx, &args, opts...)
			var s LookupProxyaddressResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProxyaddressResultOutput)
}

// A collection of arguments for invoking getProxyaddress.
type LookupProxyaddressOutputArgs struct {
	// Specify the name of the desired firewall proxyaddress.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupProxyaddressOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProxyaddressArgs)(nil)).Elem()
}

// A collection of values returned by getProxyaddress.
type LookupProxyaddressResultOutput struct{ *pulumi.OutputState }

func (LookupProxyaddressResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProxyaddressResult)(nil)).Elem()
}

func (o LookupProxyaddressResultOutput) ToLookupProxyaddressResultOutput() LookupProxyaddressResultOutput {
	return o
}

func (o LookupProxyaddressResultOutput) ToLookupProxyaddressResultOutputWithContext(ctx context.Context) LookupProxyaddressResultOutput {
	return o
}

// SaaS application. The structure of `application` block is documented below.
func (o LookupProxyaddressResultOutput) Applications() GetProxyaddressApplicationArrayOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) []GetProxyaddressApplication { return v.Applications }).(GetProxyaddressApplicationArrayOutput)
}

// Case sensitivity in pattern.
func (o LookupProxyaddressResultOutput) CaseSensitivity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) string { return v.CaseSensitivity }).(pulumi.StringOutput)
}

// Tag category.
func (o LookupProxyaddressResultOutput) Categories() GetProxyaddressCategoryArrayOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) []GetProxyaddressCategory { return v.Categories }).(GetProxyaddressCategoryArrayOutput)
}

// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
func (o LookupProxyaddressResultOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) int { return v.Color }).(pulumi.IntOutput)
}

// Optional comments.
func (o LookupProxyaddressResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) string { return v.Comment }).(pulumi.StringOutput)
}

// HTTP header regular expression.
func (o LookupProxyaddressResultOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) string { return v.Header }).(pulumi.StringOutput)
}

// HTTP header group. The structure of `headerGroup` block is documented below.
func (o LookupProxyaddressResultOutput) HeaderGroups() GetProxyaddressHeaderGroupArrayOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) []GetProxyaddressHeaderGroup { return v.HeaderGroups }).(GetProxyaddressHeaderGroupArrayOutput)
}

// HTTP header.
func (o LookupProxyaddressResultOutput) HeaderName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) string { return v.HeaderName }).(pulumi.StringOutput)
}

// Address object for the host.
func (o LookupProxyaddressResultOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) string { return v.Host }).(pulumi.StringOutput)
}

// Host name as a regular expression.
func (o LookupProxyaddressResultOutput) HostRegex() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) string { return v.HostRegex }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProxyaddressResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) string { return v.Id }).(pulumi.StringOutput)
}

// HTTP request methods to be used.
func (o LookupProxyaddressResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) string { return v.Method }).(pulumi.StringOutput)
}

// SaaS applicaton name.
func (o LookupProxyaddressResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) string { return v.Name }).(pulumi.StringOutput)
}

// URL path as a regular expression.
func (o LookupProxyaddressResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) string { return v.Path }).(pulumi.StringOutput)
}

// Match the query part of the URL as a regular expression.
func (o LookupProxyaddressResultOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) string { return v.Query }).(pulumi.StringOutput)
}

// Enable/disable use of referrer field in the HTTP header to match the address.
func (o LookupProxyaddressResultOutput) Referrer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) string { return v.Referrer }).(pulumi.StringOutput)
}

// Config object tagging. The structure of `tagging` block is documented below.
func (o LookupProxyaddressResultOutput) Taggings() GetProxyaddressTaggingArrayOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) []GetProxyaddressTagging { return v.Taggings }).(GetProxyaddressTaggingArrayOutput)
}

// Proxy address type.
func (o LookupProxyaddressResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) string { return v.Type }).(pulumi.StringOutput)
}

// Names of browsers to be used as user agent.
func (o LookupProxyaddressResultOutput) Ua() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) string { return v.Ua }).(pulumi.StringOutput)
}

// Maximum version of the user agent specified in dotted notation. For example, use 120 with the ua field set to "chrome" to require Google Chrome's maximum version must be 120.
func (o LookupProxyaddressResultOutput) UaMaxVer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) string { return v.UaMaxVer }).(pulumi.StringOutput)
}

// Minimum version of the user agent specified in dotted notation. For example, use 90.0.1 with the ua field set to "chrome" to require Google Chrome's minimum version must be 90.0.1.
func (o LookupProxyaddressResultOutput) UaMinVer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) string { return v.UaMinVer }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o LookupProxyaddressResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupProxyaddressResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable visibility of the object in the GUI.
func (o LookupProxyaddressResultOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyaddressResult) string { return v.Visibility }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProxyaddressResultOutput{})
}
