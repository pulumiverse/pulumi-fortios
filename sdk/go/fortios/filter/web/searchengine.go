// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure web filter search engines.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/filter"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := filter.NewSearchengine(ctx, "trname", &filter.SearchengineArgs{
//				Charset:    pulumi.String("utf-8"),
//				Hostname:   pulumi.String("sg.eiwuc.com"),
//				Query:      pulumi.String("sc="),
//				Safesearch: pulumi.String("disable"),
//				Url:        pulumi.String("^\\/f"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Webfilter SearchEngine can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:filter/web/searchengine:Searchengine labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:filter/web/searchengine:Searchengine labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Searchengine struct {
	pulumi.CustomResourceState

	// Search engine charset. Valid values: `utf-8`, `gb2312`.
	Charset pulumi.StringOutput `pulumi:"charset"`
	// Hostname (regular expression).
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// Search engine name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Code used to prefix a query (must end with an equals character).
	Query pulumi.StringOutput `pulumi:"query"`
	// Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
	Safesearch pulumi.StringOutput `pulumi:"safesearch"`
	// Safe search parameter used in the URL.
	SafesearchStr pulumi.StringOutput `pulumi:"safesearchStr"`
	// URL (regular expression).
	Url pulumi.StringOutput `pulumi:"url"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSearchengine registers a new resource with the given unique name, arguments, and options.
func NewSearchengine(ctx *pulumi.Context,
	name string, args *SearchengineArgs, opts ...pulumi.ResourceOption) (*Searchengine, error) {
	if args == nil {
		args = &SearchengineArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Searchengine
	err := ctx.RegisterResource("fortios:filter/web/searchengine:Searchengine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSearchengine gets an existing Searchengine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSearchengine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SearchengineState, opts ...pulumi.ResourceOption) (*Searchengine, error) {
	var resource Searchengine
	err := ctx.ReadResource("fortios:filter/web/searchengine:Searchengine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Searchengine resources.
type searchengineState struct {
	// Search engine charset. Valid values: `utf-8`, `gb2312`.
	Charset *string `pulumi:"charset"`
	// Hostname (regular expression).
	Hostname *string `pulumi:"hostname"`
	// Search engine name.
	Name *string `pulumi:"name"`
	// Code used to prefix a query (must end with an equals character).
	Query *string `pulumi:"query"`
	// Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
	Safesearch *string `pulumi:"safesearch"`
	// Safe search parameter used in the URL.
	SafesearchStr *string `pulumi:"safesearchStr"`
	// URL (regular expression).
	Url *string `pulumi:"url"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SearchengineState struct {
	// Search engine charset. Valid values: `utf-8`, `gb2312`.
	Charset pulumi.StringPtrInput
	// Hostname (regular expression).
	Hostname pulumi.StringPtrInput
	// Search engine name.
	Name pulumi.StringPtrInput
	// Code used to prefix a query (must end with an equals character).
	Query pulumi.StringPtrInput
	// Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
	Safesearch pulumi.StringPtrInput
	// Safe search parameter used in the URL.
	SafesearchStr pulumi.StringPtrInput
	// URL (regular expression).
	Url pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SearchengineState) ElementType() reflect.Type {
	return reflect.TypeOf((*searchengineState)(nil)).Elem()
}

type searchengineArgs struct {
	// Search engine charset. Valid values: `utf-8`, `gb2312`.
	Charset *string `pulumi:"charset"`
	// Hostname (regular expression).
	Hostname *string `pulumi:"hostname"`
	// Search engine name.
	Name *string `pulumi:"name"`
	// Code used to prefix a query (must end with an equals character).
	Query *string `pulumi:"query"`
	// Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
	Safesearch *string `pulumi:"safesearch"`
	// Safe search parameter used in the URL.
	SafesearchStr *string `pulumi:"safesearchStr"`
	// URL (regular expression).
	Url *string `pulumi:"url"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Searchengine resource.
type SearchengineArgs struct {
	// Search engine charset. Valid values: `utf-8`, `gb2312`.
	Charset pulumi.StringPtrInput
	// Hostname (regular expression).
	Hostname pulumi.StringPtrInput
	// Search engine name.
	Name pulumi.StringPtrInput
	// Code used to prefix a query (must end with an equals character).
	Query pulumi.StringPtrInput
	// Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
	Safesearch pulumi.StringPtrInput
	// Safe search parameter used in the URL.
	SafesearchStr pulumi.StringPtrInput
	// URL (regular expression).
	Url pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SearchengineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*searchengineArgs)(nil)).Elem()
}

type SearchengineInput interface {
	pulumi.Input

	ToSearchengineOutput() SearchengineOutput
	ToSearchengineOutputWithContext(ctx context.Context) SearchengineOutput
}

func (*Searchengine) ElementType() reflect.Type {
	return reflect.TypeOf((**Searchengine)(nil)).Elem()
}

func (i *Searchengine) ToSearchengineOutput() SearchengineOutput {
	return i.ToSearchengineOutputWithContext(context.Background())
}

func (i *Searchengine) ToSearchengineOutputWithContext(ctx context.Context) SearchengineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SearchengineOutput)
}

// SearchengineArrayInput is an input type that accepts SearchengineArray and SearchengineArrayOutput values.
// You can construct a concrete instance of `SearchengineArrayInput` via:
//
//	SearchengineArray{ SearchengineArgs{...} }
type SearchengineArrayInput interface {
	pulumi.Input

	ToSearchengineArrayOutput() SearchengineArrayOutput
	ToSearchengineArrayOutputWithContext(context.Context) SearchengineArrayOutput
}

type SearchengineArray []SearchengineInput

func (SearchengineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Searchengine)(nil)).Elem()
}

func (i SearchengineArray) ToSearchengineArrayOutput() SearchengineArrayOutput {
	return i.ToSearchengineArrayOutputWithContext(context.Background())
}

func (i SearchengineArray) ToSearchengineArrayOutputWithContext(ctx context.Context) SearchengineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SearchengineArrayOutput)
}

// SearchengineMapInput is an input type that accepts SearchengineMap and SearchengineMapOutput values.
// You can construct a concrete instance of `SearchengineMapInput` via:
//
//	SearchengineMap{ "key": SearchengineArgs{...} }
type SearchengineMapInput interface {
	pulumi.Input

	ToSearchengineMapOutput() SearchengineMapOutput
	ToSearchengineMapOutputWithContext(context.Context) SearchengineMapOutput
}

type SearchengineMap map[string]SearchengineInput

func (SearchengineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Searchengine)(nil)).Elem()
}

func (i SearchengineMap) ToSearchengineMapOutput() SearchengineMapOutput {
	return i.ToSearchengineMapOutputWithContext(context.Background())
}

func (i SearchengineMap) ToSearchengineMapOutputWithContext(ctx context.Context) SearchengineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SearchengineMapOutput)
}

type SearchengineOutput struct{ *pulumi.OutputState }

func (SearchengineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Searchengine)(nil)).Elem()
}

func (o SearchengineOutput) ToSearchengineOutput() SearchengineOutput {
	return o
}

func (o SearchengineOutput) ToSearchengineOutputWithContext(ctx context.Context) SearchengineOutput {
	return o
}

// Search engine charset. Valid values: `utf-8`, `gb2312`.
func (o SearchengineOutput) Charset() pulumi.StringOutput {
	return o.ApplyT(func(v *Searchengine) pulumi.StringOutput { return v.Charset }).(pulumi.StringOutput)
}

// Hostname (regular expression).
func (o SearchengineOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *Searchengine) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// Search engine name.
func (o SearchengineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Searchengine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Code used to prefix a query (must end with an equals character).
func (o SearchengineOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v *Searchengine) pulumi.StringOutput { return v.Query }).(pulumi.StringOutput)
}

// Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
func (o SearchengineOutput) Safesearch() pulumi.StringOutput {
	return o.ApplyT(func(v *Searchengine) pulumi.StringOutput { return v.Safesearch }).(pulumi.StringOutput)
}

// Safe search parameter used in the URL.
func (o SearchengineOutput) SafesearchStr() pulumi.StringOutput {
	return o.ApplyT(func(v *Searchengine) pulumi.StringOutput { return v.SafesearchStr }).(pulumi.StringOutput)
}

// URL (regular expression).
func (o SearchengineOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Searchengine) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SearchengineOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Searchengine) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SearchengineArrayOutput struct{ *pulumi.OutputState }

func (SearchengineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Searchengine)(nil)).Elem()
}

func (o SearchengineArrayOutput) ToSearchengineArrayOutput() SearchengineArrayOutput {
	return o
}

func (o SearchengineArrayOutput) ToSearchengineArrayOutputWithContext(ctx context.Context) SearchengineArrayOutput {
	return o
}

func (o SearchengineArrayOutput) Index(i pulumi.IntInput) SearchengineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Searchengine {
		return vs[0].([]*Searchengine)[vs[1].(int)]
	}).(SearchengineOutput)
}

type SearchengineMapOutput struct{ *pulumi.OutputState }

func (SearchengineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Searchengine)(nil)).Elem()
}

func (o SearchengineMapOutput) ToSearchengineMapOutput() SearchengineMapOutput {
	return o
}

func (o SearchengineMapOutput) ToSearchengineMapOutputWithContext(ctx context.Context) SearchengineMapOutput {
	return o
}

func (o SearchengineMapOutput) MapIndex(k pulumi.StringInput) SearchengineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Searchengine {
		return vs[0].(map[string]*Searchengine)[vs[1].(string)]
	}).(SearchengineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SearchengineInput)(nil)).Elem(), &Searchengine{})
	pulumi.RegisterInputType(reflect.TypeOf((*SearchengineArrayInput)(nil)).Elem(), SearchengineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SearchengineMapInput)(nil)).Elem(), SearchengineMap{})
	pulumi.RegisterOutputType(SearchengineOutput{})
	pulumi.RegisterOutputType(SearchengineArrayOutput{})
	pulumi.RegisterOutputType(SearchengineMapOutput{})
}
