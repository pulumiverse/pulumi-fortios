// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure SSL VPN user bookmark.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/vpn"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpn.NewUserbookmark(ctx, "trname", &vpn.UserbookmarkArgs{
//				CustomLang: pulumi.String("big5"),
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
// VpnSslWeb UserBookmark can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:vpn/ssl/web/userbookmark:Userbookmark labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:vpn/ssl/web/userbookmark:Userbookmark labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Userbookmark struct {
	pulumi.CustomResourceState

	// Bookmark table. The structure of `bookmarks` block is documented below.
	Bookmarks UserbookmarkBookmarkArrayOutput `pulumi:"bookmarks"`
	// Personal language.
	CustomLang pulumi.StringOutput `pulumi:"customLang"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// User and group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewUserbookmark registers a new resource with the given unique name, arguments, and options.
func NewUserbookmark(ctx *pulumi.Context,
	name string, args *UserbookmarkArgs, opts ...pulumi.ResourceOption) (*Userbookmark, error) {
	if args == nil {
		args = &UserbookmarkArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Userbookmark
	err := ctx.RegisterResource("fortios:vpn/ssl/web/userbookmark:Userbookmark", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserbookmark gets an existing Userbookmark resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserbookmark(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserbookmarkState, opts ...pulumi.ResourceOption) (*Userbookmark, error) {
	var resource Userbookmark
	err := ctx.ReadResource("fortios:vpn/ssl/web/userbookmark:Userbookmark", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Userbookmark resources.
type userbookmarkState struct {
	// Bookmark table. The structure of `bookmarks` block is documented below.
	Bookmarks []UserbookmarkBookmark `pulumi:"bookmarks"`
	// Personal language.
	CustomLang *string `pulumi:"customLang"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// User and group name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type UserbookmarkState struct {
	// Bookmark table. The structure of `bookmarks` block is documented below.
	Bookmarks UserbookmarkBookmarkArrayInput
	// Personal language.
	CustomLang pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// User and group name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserbookmarkState) ElementType() reflect.Type {
	return reflect.TypeOf((*userbookmarkState)(nil)).Elem()
}

type userbookmarkArgs struct {
	// Bookmark table. The structure of `bookmarks` block is documented below.
	Bookmarks []UserbookmarkBookmark `pulumi:"bookmarks"`
	// Personal language.
	CustomLang *string `pulumi:"customLang"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// User and group name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Userbookmark resource.
type UserbookmarkArgs struct {
	// Bookmark table. The structure of `bookmarks` block is documented below.
	Bookmarks UserbookmarkBookmarkArrayInput
	// Personal language.
	CustomLang pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// User and group name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserbookmarkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userbookmarkArgs)(nil)).Elem()
}

type UserbookmarkInput interface {
	pulumi.Input

	ToUserbookmarkOutput() UserbookmarkOutput
	ToUserbookmarkOutputWithContext(ctx context.Context) UserbookmarkOutput
}

func (*Userbookmark) ElementType() reflect.Type {
	return reflect.TypeOf((**Userbookmark)(nil)).Elem()
}

func (i *Userbookmark) ToUserbookmarkOutput() UserbookmarkOutput {
	return i.ToUserbookmarkOutputWithContext(context.Background())
}

func (i *Userbookmark) ToUserbookmarkOutputWithContext(ctx context.Context) UserbookmarkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserbookmarkOutput)
}

// UserbookmarkArrayInput is an input type that accepts UserbookmarkArray and UserbookmarkArrayOutput values.
// You can construct a concrete instance of `UserbookmarkArrayInput` via:
//
//	UserbookmarkArray{ UserbookmarkArgs{...} }
type UserbookmarkArrayInput interface {
	pulumi.Input

	ToUserbookmarkArrayOutput() UserbookmarkArrayOutput
	ToUserbookmarkArrayOutputWithContext(context.Context) UserbookmarkArrayOutput
}

type UserbookmarkArray []UserbookmarkInput

func (UserbookmarkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Userbookmark)(nil)).Elem()
}

func (i UserbookmarkArray) ToUserbookmarkArrayOutput() UserbookmarkArrayOutput {
	return i.ToUserbookmarkArrayOutputWithContext(context.Background())
}

func (i UserbookmarkArray) ToUserbookmarkArrayOutputWithContext(ctx context.Context) UserbookmarkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserbookmarkArrayOutput)
}

// UserbookmarkMapInput is an input type that accepts UserbookmarkMap and UserbookmarkMapOutput values.
// You can construct a concrete instance of `UserbookmarkMapInput` via:
//
//	UserbookmarkMap{ "key": UserbookmarkArgs{...} }
type UserbookmarkMapInput interface {
	pulumi.Input

	ToUserbookmarkMapOutput() UserbookmarkMapOutput
	ToUserbookmarkMapOutputWithContext(context.Context) UserbookmarkMapOutput
}

type UserbookmarkMap map[string]UserbookmarkInput

func (UserbookmarkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Userbookmark)(nil)).Elem()
}

func (i UserbookmarkMap) ToUserbookmarkMapOutput() UserbookmarkMapOutput {
	return i.ToUserbookmarkMapOutputWithContext(context.Background())
}

func (i UserbookmarkMap) ToUserbookmarkMapOutputWithContext(ctx context.Context) UserbookmarkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserbookmarkMapOutput)
}

type UserbookmarkOutput struct{ *pulumi.OutputState }

func (UserbookmarkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Userbookmark)(nil)).Elem()
}

func (o UserbookmarkOutput) ToUserbookmarkOutput() UserbookmarkOutput {
	return o
}

func (o UserbookmarkOutput) ToUserbookmarkOutputWithContext(ctx context.Context) UserbookmarkOutput {
	return o
}

// Bookmark table. The structure of `bookmarks` block is documented below.
func (o UserbookmarkOutput) Bookmarks() UserbookmarkBookmarkArrayOutput {
	return o.ApplyT(func(v *Userbookmark) UserbookmarkBookmarkArrayOutput { return v.Bookmarks }).(UserbookmarkBookmarkArrayOutput)
}

// Personal language.
func (o UserbookmarkOutput) CustomLang() pulumi.StringOutput {
	return o.ApplyT(func(v *Userbookmark) pulumi.StringOutput { return v.CustomLang }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o UserbookmarkOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Userbookmark) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o UserbookmarkOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Userbookmark) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// User and group name.
func (o UserbookmarkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Userbookmark) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o UserbookmarkOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Userbookmark) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type UserbookmarkArrayOutput struct{ *pulumi.OutputState }

func (UserbookmarkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Userbookmark)(nil)).Elem()
}

func (o UserbookmarkArrayOutput) ToUserbookmarkArrayOutput() UserbookmarkArrayOutput {
	return o
}

func (o UserbookmarkArrayOutput) ToUserbookmarkArrayOutputWithContext(ctx context.Context) UserbookmarkArrayOutput {
	return o
}

func (o UserbookmarkArrayOutput) Index(i pulumi.IntInput) UserbookmarkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Userbookmark {
		return vs[0].([]*Userbookmark)[vs[1].(int)]
	}).(UserbookmarkOutput)
}

type UserbookmarkMapOutput struct{ *pulumi.OutputState }

func (UserbookmarkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Userbookmark)(nil)).Elem()
}

func (o UserbookmarkMapOutput) ToUserbookmarkMapOutput() UserbookmarkMapOutput {
	return o
}

func (o UserbookmarkMapOutput) ToUserbookmarkMapOutputWithContext(ctx context.Context) UserbookmarkMapOutput {
	return o
}

func (o UserbookmarkMapOutput) MapIndex(k pulumi.StringInput) UserbookmarkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Userbookmark {
		return vs[0].(map[string]*Userbookmark)[vs[1].(string)]
	}).(UserbookmarkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserbookmarkInput)(nil)).Elem(), &Userbookmark{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserbookmarkArrayInput)(nil)).Elem(), UserbookmarkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserbookmarkMapInput)(nil)).Elem(), UserbookmarkMap{})
	pulumi.RegisterOutputType(UserbookmarkOutput{})
	pulumi.RegisterOutputType(UserbookmarkArrayOutput{})
	pulumi.RegisterOutputType(UserbookmarkMapOutput{})
}
