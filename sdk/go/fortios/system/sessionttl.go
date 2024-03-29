// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure global session TTL timers for this FortiGate.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/system"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := system.NewSessionttl(ctx, "trname", &system.SessionttlArgs{
//				Default: pulumi.String("3600"),
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
// System SessionTtl can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/sessionttl:Sessionttl labelname SystemSessionTtl
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/sessionttl:Sessionttl labelname SystemSessionTtl
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Sessionttl struct {
	pulumi.CustomResourceState

	// Default timeout.
	Default pulumi.StringOutput `pulumi:"default"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Session TTL port. The structure of `port` block is documented below.
	Ports SessionttlPortArrayOutput `pulumi:"ports"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSessionttl registers a new resource with the given unique name, arguments, and options.
func NewSessionttl(ctx *pulumi.Context,
	name string, args *SessionttlArgs, opts ...pulumi.ResourceOption) (*Sessionttl, error) {
	if args == nil {
		args = &SessionttlArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Sessionttl
	err := ctx.RegisterResource("fortios:system/sessionttl:Sessionttl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSessionttl gets an existing Sessionttl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSessionttl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SessionttlState, opts ...pulumi.ResourceOption) (*Sessionttl, error) {
	var resource Sessionttl
	err := ctx.ReadResource("fortios:system/sessionttl:Sessionttl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Sessionttl resources.
type sessionttlState struct {
	// Default timeout.
	Default *string `pulumi:"default"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Session TTL port. The structure of `port` block is documented below.
	Ports []SessionttlPort `pulumi:"ports"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SessionttlState struct {
	// Default timeout.
	Default pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Session TTL port. The structure of `port` block is documented below.
	Ports SessionttlPortArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SessionttlState) ElementType() reflect.Type {
	return reflect.TypeOf((*sessionttlState)(nil)).Elem()
}

type sessionttlArgs struct {
	// Default timeout.
	Default *string `pulumi:"default"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Session TTL port. The structure of `port` block is documented below.
	Ports []SessionttlPort `pulumi:"ports"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Sessionttl resource.
type SessionttlArgs struct {
	// Default timeout.
	Default pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Session TTL port. The structure of `port` block is documented below.
	Ports SessionttlPortArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SessionttlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sessionttlArgs)(nil)).Elem()
}

type SessionttlInput interface {
	pulumi.Input

	ToSessionttlOutput() SessionttlOutput
	ToSessionttlOutputWithContext(ctx context.Context) SessionttlOutput
}

func (*Sessionttl) ElementType() reflect.Type {
	return reflect.TypeOf((**Sessionttl)(nil)).Elem()
}

func (i *Sessionttl) ToSessionttlOutput() SessionttlOutput {
	return i.ToSessionttlOutputWithContext(context.Background())
}

func (i *Sessionttl) ToSessionttlOutputWithContext(ctx context.Context) SessionttlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SessionttlOutput)
}

// SessionttlArrayInput is an input type that accepts SessionttlArray and SessionttlArrayOutput values.
// You can construct a concrete instance of `SessionttlArrayInput` via:
//
//	SessionttlArray{ SessionttlArgs{...} }
type SessionttlArrayInput interface {
	pulumi.Input

	ToSessionttlArrayOutput() SessionttlArrayOutput
	ToSessionttlArrayOutputWithContext(context.Context) SessionttlArrayOutput
}

type SessionttlArray []SessionttlInput

func (SessionttlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sessionttl)(nil)).Elem()
}

func (i SessionttlArray) ToSessionttlArrayOutput() SessionttlArrayOutput {
	return i.ToSessionttlArrayOutputWithContext(context.Background())
}

func (i SessionttlArray) ToSessionttlArrayOutputWithContext(ctx context.Context) SessionttlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SessionttlArrayOutput)
}

// SessionttlMapInput is an input type that accepts SessionttlMap and SessionttlMapOutput values.
// You can construct a concrete instance of `SessionttlMapInput` via:
//
//	SessionttlMap{ "key": SessionttlArgs{...} }
type SessionttlMapInput interface {
	pulumi.Input

	ToSessionttlMapOutput() SessionttlMapOutput
	ToSessionttlMapOutputWithContext(context.Context) SessionttlMapOutput
}

type SessionttlMap map[string]SessionttlInput

func (SessionttlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sessionttl)(nil)).Elem()
}

func (i SessionttlMap) ToSessionttlMapOutput() SessionttlMapOutput {
	return i.ToSessionttlMapOutputWithContext(context.Background())
}

func (i SessionttlMap) ToSessionttlMapOutputWithContext(ctx context.Context) SessionttlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SessionttlMapOutput)
}

type SessionttlOutput struct{ *pulumi.OutputState }

func (SessionttlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sessionttl)(nil)).Elem()
}

func (o SessionttlOutput) ToSessionttlOutput() SessionttlOutput {
	return o
}

func (o SessionttlOutput) ToSessionttlOutputWithContext(ctx context.Context) SessionttlOutput {
	return o
}

// Default timeout.
func (o SessionttlOutput) Default() pulumi.StringOutput {
	return o.ApplyT(func(v *Sessionttl) pulumi.StringOutput { return v.Default }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SessionttlOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sessionttl) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o SessionttlOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sessionttl) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Session TTL port. The structure of `port` block is documented below.
func (o SessionttlOutput) Ports() SessionttlPortArrayOutput {
	return o.ApplyT(func(v *Sessionttl) SessionttlPortArrayOutput { return v.Ports }).(SessionttlPortArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SessionttlOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sessionttl) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SessionttlArrayOutput struct{ *pulumi.OutputState }

func (SessionttlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sessionttl)(nil)).Elem()
}

func (o SessionttlArrayOutput) ToSessionttlArrayOutput() SessionttlArrayOutput {
	return o
}

func (o SessionttlArrayOutput) ToSessionttlArrayOutputWithContext(ctx context.Context) SessionttlArrayOutput {
	return o
}

func (o SessionttlArrayOutput) Index(i pulumi.IntInput) SessionttlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Sessionttl {
		return vs[0].([]*Sessionttl)[vs[1].(int)]
	}).(SessionttlOutput)
}

type SessionttlMapOutput struct{ *pulumi.OutputState }

func (SessionttlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sessionttl)(nil)).Elem()
}

func (o SessionttlMapOutput) ToSessionttlMapOutput() SessionttlMapOutput {
	return o
}

func (o SessionttlMapOutput) ToSessionttlMapOutputWithContext(ctx context.Context) SessionttlMapOutput {
	return o
}

func (o SessionttlMapOutput) MapIndex(k pulumi.StringInput) SessionttlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Sessionttl {
		return vs[0].(map[string]*Sessionttl)[vs[1].(string)]
	}).(SessionttlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SessionttlInput)(nil)).Elem(), &Sessionttl{})
	pulumi.RegisterInputType(reflect.TypeOf((*SessionttlArrayInput)(nil)).Elem(), SessionttlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SessionttlMapInput)(nil)).Elem(), SessionttlMap{})
	pulumi.RegisterOutputType(SessionttlOutput{})
	pulumi.RegisterOutputType(SessionttlArrayOutput{})
	pulumi.RegisterOutputType(SessionttlMapOutput{})
}
