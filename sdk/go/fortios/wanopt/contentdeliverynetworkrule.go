// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wanopt

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure WAN optimization content delivery network rules.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/wanopt"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := wanopt.NewContentdeliverynetworkrule(ctx, "trname", &wanopt.ContentdeliverynetworkruleArgs{
//				Category: pulumi.String("vcache"),
//				HostDomainNameSuffixes: wanopt.ContentdeliverynetworkruleHostDomainNameSuffixArray{
//					&wanopt.ContentdeliverynetworkruleHostDomainNameSuffixArgs{
//						Name: pulumi.String("kaf.com"),
//					},
//				},
//				RequestCacheControl:  pulumi.String("disable"),
//				ResponseCacheControl: pulumi.String("disable"),
//				ResponseExpires:      pulumi.String("enable"),
//				Status:               pulumi.String("enable"),
//				TextResponseVcache:   pulumi.String("enable"),
//				Updateserver:         pulumi.String("disable"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Wanopt ContentDeliveryNetworkRule can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:wanopt/contentdeliverynetworkrule:Contentdeliverynetworkrule labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:wanopt/contentdeliverynetworkrule:Contentdeliverynetworkrule labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Contentdeliverynetworkrule struct {
	pulumi.CustomResourceState

	// Content delivery network rule category. Valid values: `vcache`, `youtube`.
	Category pulumi.StringOutput `pulumi:"category"`
	// Comment about this CDN-rule.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Suffix portion of the fully qualified domain name (eg. fortinet.com in "www.fortinet.com"). The structure of `hostDomainNameSuffix` block is documented below.
	HostDomainNameSuffixes ContentdeliverynetworkruleHostDomainNameSuffixArrayOutput `pulumi:"hostDomainNameSuffixes"`
	// Name of table.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable HTTP request cache control. Valid values: `enable`, `disable`.
	RequestCacheControl pulumi.StringOutput `pulumi:"requestCacheControl"`
	// Enable/disable HTTP response cache control. Valid values: `enable`, `disable`.
	ResponseCacheControl pulumi.StringOutput `pulumi:"responseCacheControl"`
	// Enable/disable HTTP response cache expires. Valid values: `enable`, `disable`.
	ResponseExpires pulumi.StringOutput `pulumi:"responseExpires"`
	// WAN optimization content delivery network rule entries. The structure of `rules` block is documented below.
	Rules ContentdeliverynetworkruleRuleArrayOutput `pulumi:"rules"`
	// Enable/disable WAN optimization content delivery network rules. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Enable/disable caching of text responses. Valid values: `enable`, `disable`.
	TextResponseVcache pulumi.StringOutput `pulumi:"textResponseVcache"`
	// Enable/disable update server. Valid values: `enable`, `disable`.
	Updateserver pulumi.StringOutput `pulumi:"updateserver"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewContentdeliverynetworkrule registers a new resource with the given unique name, arguments, and options.
func NewContentdeliverynetworkrule(ctx *pulumi.Context,
	name string, args *ContentdeliverynetworkruleArgs, opts ...pulumi.ResourceOption) (*Contentdeliverynetworkrule, error) {
	if args == nil {
		args = &ContentdeliverynetworkruleArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Contentdeliverynetworkrule
	err := ctx.RegisterResource("fortios:wanopt/contentdeliverynetworkrule:Contentdeliverynetworkrule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContentdeliverynetworkrule gets an existing Contentdeliverynetworkrule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContentdeliverynetworkrule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContentdeliverynetworkruleState, opts ...pulumi.ResourceOption) (*Contentdeliverynetworkrule, error) {
	var resource Contentdeliverynetworkrule
	err := ctx.ReadResource("fortios:wanopt/contentdeliverynetworkrule:Contentdeliverynetworkrule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Contentdeliverynetworkrule resources.
type contentdeliverynetworkruleState struct {
	// Content delivery network rule category. Valid values: `vcache`, `youtube`.
	Category *string `pulumi:"category"`
	// Comment about this CDN-rule.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Suffix portion of the fully qualified domain name (eg. fortinet.com in "www.fortinet.com"). The structure of `hostDomainNameSuffix` block is documented below.
	HostDomainNameSuffixes []ContentdeliverynetworkruleHostDomainNameSuffix `pulumi:"hostDomainNameSuffixes"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Enable/disable HTTP request cache control. Valid values: `enable`, `disable`.
	RequestCacheControl *string `pulumi:"requestCacheControl"`
	// Enable/disable HTTP response cache control. Valid values: `enable`, `disable`.
	ResponseCacheControl *string `pulumi:"responseCacheControl"`
	// Enable/disable HTTP response cache expires. Valid values: `enable`, `disable`.
	ResponseExpires *string `pulumi:"responseExpires"`
	// WAN optimization content delivery network rule entries. The structure of `rules` block is documented below.
	Rules []ContentdeliverynetworkruleRule `pulumi:"rules"`
	// Enable/disable WAN optimization content delivery network rules. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Enable/disable caching of text responses. Valid values: `enable`, `disable`.
	TextResponseVcache *string `pulumi:"textResponseVcache"`
	// Enable/disable update server. Valid values: `enable`, `disable`.
	Updateserver *string `pulumi:"updateserver"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ContentdeliverynetworkruleState struct {
	// Content delivery network rule category. Valid values: `vcache`, `youtube`.
	Category pulumi.StringPtrInput
	// Comment about this CDN-rule.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Suffix portion of the fully qualified domain name (eg. fortinet.com in "www.fortinet.com"). The structure of `hostDomainNameSuffix` block is documented below.
	HostDomainNameSuffixes ContentdeliverynetworkruleHostDomainNameSuffixArrayInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Enable/disable HTTP request cache control. Valid values: `enable`, `disable`.
	RequestCacheControl pulumi.StringPtrInput
	// Enable/disable HTTP response cache control. Valid values: `enable`, `disable`.
	ResponseCacheControl pulumi.StringPtrInput
	// Enable/disable HTTP response cache expires. Valid values: `enable`, `disable`.
	ResponseExpires pulumi.StringPtrInput
	// WAN optimization content delivery network rule entries. The structure of `rules` block is documented below.
	Rules ContentdeliverynetworkruleRuleArrayInput
	// Enable/disable WAN optimization content delivery network rules. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Enable/disable caching of text responses. Valid values: `enable`, `disable`.
	TextResponseVcache pulumi.StringPtrInput
	// Enable/disable update server. Valid values: `enable`, `disable`.
	Updateserver pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ContentdeliverynetworkruleState) ElementType() reflect.Type {
	return reflect.TypeOf((*contentdeliverynetworkruleState)(nil)).Elem()
}

type contentdeliverynetworkruleArgs struct {
	// Content delivery network rule category. Valid values: `vcache`, `youtube`.
	Category *string `pulumi:"category"`
	// Comment about this CDN-rule.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Suffix portion of the fully qualified domain name (eg. fortinet.com in "www.fortinet.com"). The structure of `hostDomainNameSuffix` block is documented below.
	HostDomainNameSuffixes []ContentdeliverynetworkruleHostDomainNameSuffix `pulumi:"hostDomainNameSuffixes"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Enable/disable HTTP request cache control. Valid values: `enable`, `disable`.
	RequestCacheControl *string `pulumi:"requestCacheControl"`
	// Enable/disable HTTP response cache control. Valid values: `enable`, `disable`.
	ResponseCacheControl *string `pulumi:"responseCacheControl"`
	// Enable/disable HTTP response cache expires. Valid values: `enable`, `disable`.
	ResponseExpires *string `pulumi:"responseExpires"`
	// WAN optimization content delivery network rule entries. The structure of `rules` block is documented below.
	Rules []ContentdeliverynetworkruleRule `pulumi:"rules"`
	// Enable/disable WAN optimization content delivery network rules. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Enable/disable caching of text responses. Valid values: `enable`, `disable`.
	TextResponseVcache *string `pulumi:"textResponseVcache"`
	// Enable/disable update server. Valid values: `enable`, `disable`.
	Updateserver *string `pulumi:"updateserver"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Contentdeliverynetworkrule resource.
type ContentdeliverynetworkruleArgs struct {
	// Content delivery network rule category. Valid values: `vcache`, `youtube`.
	Category pulumi.StringPtrInput
	// Comment about this CDN-rule.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Suffix portion of the fully qualified domain name (eg. fortinet.com in "www.fortinet.com"). The structure of `hostDomainNameSuffix` block is documented below.
	HostDomainNameSuffixes ContentdeliverynetworkruleHostDomainNameSuffixArrayInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Enable/disable HTTP request cache control. Valid values: `enable`, `disable`.
	RequestCacheControl pulumi.StringPtrInput
	// Enable/disable HTTP response cache control. Valid values: `enable`, `disable`.
	ResponseCacheControl pulumi.StringPtrInput
	// Enable/disable HTTP response cache expires. Valid values: `enable`, `disable`.
	ResponseExpires pulumi.StringPtrInput
	// WAN optimization content delivery network rule entries. The structure of `rules` block is documented below.
	Rules ContentdeliverynetworkruleRuleArrayInput
	// Enable/disable WAN optimization content delivery network rules. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Enable/disable caching of text responses. Valid values: `enable`, `disable`.
	TextResponseVcache pulumi.StringPtrInput
	// Enable/disable update server. Valid values: `enable`, `disable`.
	Updateserver pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ContentdeliverynetworkruleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contentdeliverynetworkruleArgs)(nil)).Elem()
}

type ContentdeliverynetworkruleInput interface {
	pulumi.Input

	ToContentdeliverynetworkruleOutput() ContentdeliverynetworkruleOutput
	ToContentdeliverynetworkruleOutputWithContext(ctx context.Context) ContentdeliverynetworkruleOutput
}

func (*Contentdeliverynetworkrule) ElementType() reflect.Type {
	return reflect.TypeOf((**Contentdeliverynetworkrule)(nil)).Elem()
}

func (i *Contentdeliverynetworkrule) ToContentdeliverynetworkruleOutput() ContentdeliverynetworkruleOutput {
	return i.ToContentdeliverynetworkruleOutputWithContext(context.Background())
}

func (i *Contentdeliverynetworkrule) ToContentdeliverynetworkruleOutputWithContext(ctx context.Context) ContentdeliverynetworkruleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentdeliverynetworkruleOutput)
}

// ContentdeliverynetworkruleArrayInput is an input type that accepts ContentdeliverynetworkruleArray and ContentdeliverynetworkruleArrayOutput values.
// You can construct a concrete instance of `ContentdeliverynetworkruleArrayInput` via:
//
//	ContentdeliverynetworkruleArray{ ContentdeliverynetworkruleArgs{...} }
type ContentdeliverynetworkruleArrayInput interface {
	pulumi.Input

	ToContentdeliverynetworkruleArrayOutput() ContentdeliverynetworkruleArrayOutput
	ToContentdeliverynetworkruleArrayOutputWithContext(context.Context) ContentdeliverynetworkruleArrayOutput
}

type ContentdeliverynetworkruleArray []ContentdeliverynetworkruleInput

func (ContentdeliverynetworkruleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Contentdeliverynetworkrule)(nil)).Elem()
}

func (i ContentdeliverynetworkruleArray) ToContentdeliverynetworkruleArrayOutput() ContentdeliverynetworkruleArrayOutput {
	return i.ToContentdeliverynetworkruleArrayOutputWithContext(context.Background())
}

func (i ContentdeliverynetworkruleArray) ToContentdeliverynetworkruleArrayOutputWithContext(ctx context.Context) ContentdeliverynetworkruleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentdeliverynetworkruleArrayOutput)
}

// ContentdeliverynetworkruleMapInput is an input type that accepts ContentdeliverynetworkruleMap and ContentdeliverynetworkruleMapOutput values.
// You can construct a concrete instance of `ContentdeliverynetworkruleMapInput` via:
//
//	ContentdeliverynetworkruleMap{ "key": ContentdeliverynetworkruleArgs{...} }
type ContentdeliverynetworkruleMapInput interface {
	pulumi.Input

	ToContentdeliverynetworkruleMapOutput() ContentdeliverynetworkruleMapOutput
	ToContentdeliverynetworkruleMapOutputWithContext(context.Context) ContentdeliverynetworkruleMapOutput
}

type ContentdeliverynetworkruleMap map[string]ContentdeliverynetworkruleInput

func (ContentdeliverynetworkruleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Contentdeliverynetworkrule)(nil)).Elem()
}

func (i ContentdeliverynetworkruleMap) ToContentdeliverynetworkruleMapOutput() ContentdeliverynetworkruleMapOutput {
	return i.ToContentdeliverynetworkruleMapOutputWithContext(context.Background())
}

func (i ContentdeliverynetworkruleMap) ToContentdeliverynetworkruleMapOutputWithContext(ctx context.Context) ContentdeliverynetworkruleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentdeliverynetworkruleMapOutput)
}

type ContentdeliverynetworkruleOutput struct{ *pulumi.OutputState }

func (ContentdeliverynetworkruleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Contentdeliverynetworkrule)(nil)).Elem()
}

func (o ContentdeliverynetworkruleOutput) ToContentdeliverynetworkruleOutput() ContentdeliverynetworkruleOutput {
	return o
}

func (o ContentdeliverynetworkruleOutput) ToContentdeliverynetworkruleOutputWithContext(ctx context.Context) ContentdeliverynetworkruleOutput {
	return o
}

// Content delivery network rule category. Valid values: `vcache`, `youtube`.
func (o ContentdeliverynetworkruleOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *Contentdeliverynetworkrule) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

// Comment about this CDN-rule.
func (o ContentdeliverynetworkruleOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Contentdeliverynetworkrule) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o ContentdeliverynetworkruleOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Contentdeliverynetworkrule) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o ContentdeliverynetworkruleOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Contentdeliverynetworkrule) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Suffix portion of the fully qualified domain name (eg. fortinet.com in "www.fortinet.com"). The structure of `hostDomainNameSuffix` block is documented below.
func (o ContentdeliverynetworkruleOutput) HostDomainNameSuffixes() ContentdeliverynetworkruleHostDomainNameSuffixArrayOutput {
	return o.ApplyT(func(v *Contentdeliverynetworkrule) ContentdeliverynetworkruleHostDomainNameSuffixArrayOutput {
		return v.HostDomainNameSuffixes
	}).(ContentdeliverynetworkruleHostDomainNameSuffixArrayOutput)
}

// Name of table.
func (o ContentdeliverynetworkruleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Contentdeliverynetworkrule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable HTTP request cache control. Valid values: `enable`, `disable`.
func (o ContentdeliverynetworkruleOutput) RequestCacheControl() pulumi.StringOutput {
	return o.ApplyT(func(v *Contentdeliverynetworkrule) pulumi.StringOutput { return v.RequestCacheControl }).(pulumi.StringOutput)
}

// Enable/disable HTTP response cache control. Valid values: `enable`, `disable`.
func (o ContentdeliverynetworkruleOutput) ResponseCacheControl() pulumi.StringOutput {
	return o.ApplyT(func(v *Contentdeliverynetworkrule) pulumi.StringOutput { return v.ResponseCacheControl }).(pulumi.StringOutput)
}

// Enable/disable HTTP response cache expires. Valid values: `enable`, `disable`.
func (o ContentdeliverynetworkruleOutput) ResponseExpires() pulumi.StringOutput {
	return o.ApplyT(func(v *Contentdeliverynetworkrule) pulumi.StringOutput { return v.ResponseExpires }).(pulumi.StringOutput)
}

// WAN optimization content delivery network rule entries. The structure of `rules` block is documented below.
func (o ContentdeliverynetworkruleOutput) Rules() ContentdeliverynetworkruleRuleArrayOutput {
	return o.ApplyT(func(v *Contentdeliverynetworkrule) ContentdeliverynetworkruleRuleArrayOutput { return v.Rules }).(ContentdeliverynetworkruleRuleArrayOutput)
}

// Enable/disable WAN optimization content delivery network rules. Valid values: `enable`, `disable`.
func (o ContentdeliverynetworkruleOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Contentdeliverynetworkrule) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Enable/disable caching of text responses. Valid values: `enable`, `disable`.
func (o ContentdeliverynetworkruleOutput) TextResponseVcache() pulumi.StringOutput {
	return o.ApplyT(func(v *Contentdeliverynetworkrule) pulumi.StringOutput { return v.TextResponseVcache }).(pulumi.StringOutput)
}

// Enable/disable update server. Valid values: `enable`, `disable`.
func (o ContentdeliverynetworkruleOutput) Updateserver() pulumi.StringOutput {
	return o.ApplyT(func(v *Contentdeliverynetworkrule) pulumi.StringOutput { return v.Updateserver }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ContentdeliverynetworkruleOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Contentdeliverynetworkrule) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ContentdeliverynetworkruleArrayOutput struct{ *pulumi.OutputState }

func (ContentdeliverynetworkruleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Contentdeliverynetworkrule)(nil)).Elem()
}

func (o ContentdeliverynetworkruleArrayOutput) ToContentdeliverynetworkruleArrayOutput() ContentdeliverynetworkruleArrayOutput {
	return o
}

func (o ContentdeliverynetworkruleArrayOutput) ToContentdeliverynetworkruleArrayOutputWithContext(ctx context.Context) ContentdeliverynetworkruleArrayOutput {
	return o
}

func (o ContentdeliverynetworkruleArrayOutput) Index(i pulumi.IntInput) ContentdeliverynetworkruleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Contentdeliverynetworkrule {
		return vs[0].([]*Contentdeliverynetworkrule)[vs[1].(int)]
	}).(ContentdeliverynetworkruleOutput)
}

type ContentdeliverynetworkruleMapOutput struct{ *pulumi.OutputState }

func (ContentdeliverynetworkruleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Contentdeliverynetworkrule)(nil)).Elem()
}

func (o ContentdeliverynetworkruleMapOutput) ToContentdeliverynetworkruleMapOutput() ContentdeliverynetworkruleMapOutput {
	return o
}

func (o ContentdeliverynetworkruleMapOutput) ToContentdeliverynetworkruleMapOutputWithContext(ctx context.Context) ContentdeliverynetworkruleMapOutput {
	return o
}

func (o ContentdeliverynetworkruleMapOutput) MapIndex(k pulumi.StringInput) ContentdeliverynetworkruleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Contentdeliverynetworkrule {
		return vs[0].(map[string]*Contentdeliverynetworkrule)[vs[1].(string)]
	}).(ContentdeliverynetworkruleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContentdeliverynetworkruleInput)(nil)).Elem(), &Contentdeliverynetworkrule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentdeliverynetworkruleArrayInput)(nil)).Elem(), ContentdeliverynetworkruleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentdeliverynetworkruleMapInput)(nil)).Elem(), ContentdeliverynetworkruleMap{})
	pulumi.RegisterOutputType(ContentdeliverynetworkruleOutput{})
	pulumi.RegisterOutputType(ContentdeliverynetworkruleArrayOutput{})
	pulumi.RegisterOutputType(ContentdeliverynetworkruleMapOutput{})
}
