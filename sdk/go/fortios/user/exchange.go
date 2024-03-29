// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package user

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure MS Exchange server entries. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// User Exchange can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:user/exchange:Exchange labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:user/exchange:Exchange labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Exchange struct {
	pulumi.CustomResourceState

	// Authentication security level used for the RPC protocol layer. Valid values: `connect`, `call`, `packet`, `integrity`, `privacy`.
	AuthLevel pulumi.StringOutput `pulumi:"authLevel"`
	// Authentication security type used for the RPC protocol layer. Valid values: `spnego`, `ntlm`, `kerberos`.
	AuthType pulumi.StringOutput `pulumi:"authType"`
	// Enable/disable automatic discovery of KDC IP addresses. Valid values: `enable`, `disable`.
	AutoDiscoverKdc pulumi.StringOutput `pulumi:"autoDiscoverKdc"`
	// Connection protocol used to connect to MS Exchange service. Valid values: `rpc-over-tcp`, `rpc-over-http`, `rpc-over-https`.
	ConnectProtocol pulumi.StringOutput `pulumi:"connectProtocol"`
	// MS Exchange server fully qualified domain name.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Authentication security type used for the HTTP transport. Valid values: `basic`, `ntlm`.
	HttpAuthType pulumi.StringOutput `pulumi:"httpAuthType"`
	// Server IPv4 address.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// KDC IPv4 addresses for Kerberos authentication. The structure of `kdcIp` block is documented below.
	KdcIps ExchangeKdcIpArrayOutput `pulumi:"kdcIps"`
	// MS Exchange server entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password for the specified username.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// MS Exchange server hostname.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting).
	SslMinProtoVersion pulumi.StringOutput `pulumi:"sslMinProtoVersion"`
	// User name used to sign in to the server. Must have proper permissions for service.
	Username pulumi.StringOutput `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewExchange registers a new resource with the given unique name, arguments, and options.
func NewExchange(ctx *pulumi.Context,
	name string, args *ExchangeArgs, opts ...pulumi.ResourceOption) (*Exchange, error) {
	if args == nil {
		args = &ExchangeArgs{}
	}

	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Exchange
	err := ctx.RegisterResource("fortios:user/exchange:Exchange", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExchange gets an existing Exchange resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExchange(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExchangeState, opts ...pulumi.ResourceOption) (*Exchange, error) {
	var resource Exchange
	err := ctx.ReadResource("fortios:user/exchange:Exchange", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Exchange resources.
type exchangeState struct {
	// Authentication security level used for the RPC protocol layer. Valid values: `connect`, `call`, `packet`, `integrity`, `privacy`.
	AuthLevel *string `pulumi:"authLevel"`
	// Authentication security type used for the RPC protocol layer. Valid values: `spnego`, `ntlm`, `kerberos`.
	AuthType *string `pulumi:"authType"`
	// Enable/disable automatic discovery of KDC IP addresses. Valid values: `enable`, `disable`.
	AutoDiscoverKdc *string `pulumi:"autoDiscoverKdc"`
	// Connection protocol used to connect to MS Exchange service. Valid values: `rpc-over-tcp`, `rpc-over-http`, `rpc-over-https`.
	ConnectProtocol *string `pulumi:"connectProtocol"`
	// MS Exchange server fully qualified domain name.
	DomainName *string `pulumi:"domainName"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Authentication security type used for the HTTP transport. Valid values: `basic`, `ntlm`.
	HttpAuthType *string `pulumi:"httpAuthType"`
	// Server IPv4 address.
	Ip *string `pulumi:"ip"`
	// KDC IPv4 addresses for Kerberos authentication. The structure of `kdcIp` block is documented below.
	KdcIps []ExchangeKdcIp `pulumi:"kdcIps"`
	// MS Exchange server entry name.
	Name *string `pulumi:"name"`
	// Password for the specified username.
	Password *string `pulumi:"password"`
	// MS Exchange server hostname.
	ServerName *string `pulumi:"serverName"`
	// Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting).
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// User name used to sign in to the server. Must have proper permissions for service.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ExchangeState struct {
	// Authentication security level used for the RPC protocol layer. Valid values: `connect`, `call`, `packet`, `integrity`, `privacy`.
	AuthLevel pulumi.StringPtrInput
	// Authentication security type used for the RPC protocol layer. Valid values: `spnego`, `ntlm`, `kerberos`.
	AuthType pulumi.StringPtrInput
	// Enable/disable automatic discovery of KDC IP addresses. Valid values: `enable`, `disable`.
	AutoDiscoverKdc pulumi.StringPtrInput
	// Connection protocol used to connect to MS Exchange service. Valid values: `rpc-over-tcp`, `rpc-over-http`, `rpc-over-https`.
	ConnectProtocol pulumi.StringPtrInput
	// MS Exchange server fully qualified domain name.
	DomainName pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Authentication security type used for the HTTP transport. Valid values: `basic`, `ntlm`.
	HttpAuthType pulumi.StringPtrInput
	// Server IPv4 address.
	Ip pulumi.StringPtrInput
	// KDC IPv4 addresses for Kerberos authentication. The structure of `kdcIp` block is documented below.
	KdcIps ExchangeKdcIpArrayInput
	// MS Exchange server entry name.
	Name pulumi.StringPtrInput
	// Password for the specified username.
	Password pulumi.StringPtrInput
	// MS Exchange server hostname.
	ServerName pulumi.StringPtrInput
	// Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting).
	SslMinProtoVersion pulumi.StringPtrInput
	// User name used to sign in to the server. Must have proper permissions for service.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ExchangeState) ElementType() reflect.Type {
	return reflect.TypeOf((*exchangeState)(nil)).Elem()
}

type exchangeArgs struct {
	// Authentication security level used for the RPC protocol layer. Valid values: `connect`, `call`, `packet`, `integrity`, `privacy`.
	AuthLevel *string `pulumi:"authLevel"`
	// Authentication security type used for the RPC protocol layer. Valid values: `spnego`, `ntlm`, `kerberos`.
	AuthType *string `pulumi:"authType"`
	// Enable/disable automatic discovery of KDC IP addresses. Valid values: `enable`, `disable`.
	AutoDiscoverKdc *string `pulumi:"autoDiscoverKdc"`
	// Connection protocol used to connect to MS Exchange service. Valid values: `rpc-over-tcp`, `rpc-over-http`, `rpc-over-https`.
	ConnectProtocol *string `pulumi:"connectProtocol"`
	// MS Exchange server fully qualified domain name.
	DomainName *string `pulumi:"domainName"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Authentication security type used for the HTTP transport. Valid values: `basic`, `ntlm`.
	HttpAuthType *string `pulumi:"httpAuthType"`
	// Server IPv4 address.
	Ip *string `pulumi:"ip"`
	// KDC IPv4 addresses for Kerberos authentication. The structure of `kdcIp` block is documented below.
	KdcIps []ExchangeKdcIp `pulumi:"kdcIps"`
	// MS Exchange server entry name.
	Name *string `pulumi:"name"`
	// Password for the specified username.
	Password *string `pulumi:"password"`
	// MS Exchange server hostname.
	ServerName *string `pulumi:"serverName"`
	// Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting).
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// User name used to sign in to the server. Must have proper permissions for service.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Exchange resource.
type ExchangeArgs struct {
	// Authentication security level used for the RPC protocol layer. Valid values: `connect`, `call`, `packet`, `integrity`, `privacy`.
	AuthLevel pulumi.StringPtrInput
	// Authentication security type used for the RPC protocol layer. Valid values: `spnego`, `ntlm`, `kerberos`.
	AuthType pulumi.StringPtrInput
	// Enable/disable automatic discovery of KDC IP addresses. Valid values: `enable`, `disable`.
	AutoDiscoverKdc pulumi.StringPtrInput
	// Connection protocol used to connect to MS Exchange service. Valid values: `rpc-over-tcp`, `rpc-over-http`, `rpc-over-https`.
	ConnectProtocol pulumi.StringPtrInput
	// MS Exchange server fully qualified domain name.
	DomainName pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Authentication security type used for the HTTP transport. Valid values: `basic`, `ntlm`.
	HttpAuthType pulumi.StringPtrInput
	// Server IPv4 address.
	Ip pulumi.StringPtrInput
	// KDC IPv4 addresses for Kerberos authentication. The structure of `kdcIp` block is documented below.
	KdcIps ExchangeKdcIpArrayInput
	// MS Exchange server entry name.
	Name pulumi.StringPtrInput
	// Password for the specified username.
	Password pulumi.StringPtrInput
	// MS Exchange server hostname.
	ServerName pulumi.StringPtrInput
	// Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting).
	SslMinProtoVersion pulumi.StringPtrInput
	// User name used to sign in to the server. Must have proper permissions for service.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ExchangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*exchangeArgs)(nil)).Elem()
}

type ExchangeInput interface {
	pulumi.Input

	ToExchangeOutput() ExchangeOutput
	ToExchangeOutputWithContext(ctx context.Context) ExchangeOutput
}

func (*Exchange) ElementType() reflect.Type {
	return reflect.TypeOf((**Exchange)(nil)).Elem()
}

func (i *Exchange) ToExchangeOutput() ExchangeOutput {
	return i.ToExchangeOutputWithContext(context.Background())
}

func (i *Exchange) ToExchangeOutputWithContext(ctx context.Context) ExchangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExchangeOutput)
}

// ExchangeArrayInput is an input type that accepts ExchangeArray and ExchangeArrayOutput values.
// You can construct a concrete instance of `ExchangeArrayInput` via:
//
//	ExchangeArray{ ExchangeArgs{...} }
type ExchangeArrayInput interface {
	pulumi.Input

	ToExchangeArrayOutput() ExchangeArrayOutput
	ToExchangeArrayOutputWithContext(context.Context) ExchangeArrayOutput
}

type ExchangeArray []ExchangeInput

func (ExchangeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Exchange)(nil)).Elem()
}

func (i ExchangeArray) ToExchangeArrayOutput() ExchangeArrayOutput {
	return i.ToExchangeArrayOutputWithContext(context.Background())
}

func (i ExchangeArray) ToExchangeArrayOutputWithContext(ctx context.Context) ExchangeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExchangeArrayOutput)
}

// ExchangeMapInput is an input type that accepts ExchangeMap and ExchangeMapOutput values.
// You can construct a concrete instance of `ExchangeMapInput` via:
//
//	ExchangeMap{ "key": ExchangeArgs{...} }
type ExchangeMapInput interface {
	pulumi.Input

	ToExchangeMapOutput() ExchangeMapOutput
	ToExchangeMapOutputWithContext(context.Context) ExchangeMapOutput
}

type ExchangeMap map[string]ExchangeInput

func (ExchangeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Exchange)(nil)).Elem()
}

func (i ExchangeMap) ToExchangeMapOutput() ExchangeMapOutput {
	return i.ToExchangeMapOutputWithContext(context.Background())
}

func (i ExchangeMap) ToExchangeMapOutputWithContext(ctx context.Context) ExchangeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExchangeMapOutput)
}

type ExchangeOutput struct{ *pulumi.OutputState }

func (ExchangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Exchange)(nil)).Elem()
}

func (o ExchangeOutput) ToExchangeOutput() ExchangeOutput {
	return o
}

func (o ExchangeOutput) ToExchangeOutputWithContext(ctx context.Context) ExchangeOutput {
	return o
}

// Authentication security level used for the RPC protocol layer. Valid values: `connect`, `call`, `packet`, `integrity`, `privacy`.
func (o ExchangeOutput) AuthLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *Exchange) pulumi.StringOutput { return v.AuthLevel }).(pulumi.StringOutput)
}

// Authentication security type used for the RPC protocol layer. Valid values: `spnego`, `ntlm`, `kerberos`.
func (o ExchangeOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v *Exchange) pulumi.StringOutput { return v.AuthType }).(pulumi.StringOutput)
}

// Enable/disable automatic discovery of KDC IP addresses. Valid values: `enable`, `disable`.
func (o ExchangeOutput) AutoDiscoverKdc() pulumi.StringOutput {
	return o.ApplyT(func(v *Exchange) pulumi.StringOutput { return v.AutoDiscoverKdc }).(pulumi.StringOutput)
}

// Connection protocol used to connect to MS Exchange service. Valid values: `rpc-over-tcp`, `rpc-over-http`, `rpc-over-https`.
func (o ExchangeOutput) ConnectProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Exchange) pulumi.StringOutput { return v.ConnectProtocol }).(pulumi.StringOutput)
}

// MS Exchange server fully qualified domain name.
func (o ExchangeOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Exchange) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o ExchangeOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Exchange) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o ExchangeOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Exchange) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Authentication security type used for the HTTP transport. Valid values: `basic`, `ntlm`.
func (o ExchangeOutput) HttpAuthType() pulumi.StringOutput {
	return o.ApplyT(func(v *Exchange) pulumi.StringOutput { return v.HttpAuthType }).(pulumi.StringOutput)
}

// Server IPv4 address.
func (o ExchangeOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *Exchange) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// KDC IPv4 addresses for Kerberos authentication. The structure of `kdcIp` block is documented below.
func (o ExchangeOutput) KdcIps() ExchangeKdcIpArrayOutput {
	return o.ApplyT(func(v *Exchange) ExchangeKdcIpArrayOutput { return v.KdcIps }).(ExchangeKdcIpArrayOutput)
}

// MS Exchange server entry name.
func (o ExchangeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Exchange) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Password for the specified username.
func (o ExchangeOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Exchange) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// MS Exchange server hostname.
func (o ExchangeOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v *Exchange) pulumi.StringOutput { return v.ServerName }).(pulumi.StringOutput)
}

// Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting).
func (o ExchangeOutput) SslMinProtoVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Exchange) pulumi.StringOutput { return v.SslMinProtoVersion }).(pulumi.StringOutput)
}

// User name used to sign in to the server. Must have proper permissions for service.
func (o ExchangeOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *Exchange) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ExchangeOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Exchange) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ExchangeArrayOutput struct{ *pulumi.OutputState }

func (ExchangeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Exchange)(nil)).Elem()
}

func (o ExchangeArrayOutput) ToExchangeArrayOutput() ExchangeArrayOutput {
	return o
}

func (o ExchangeArrayOutput) ToExchangeArrayOutputWithContext(ctx context.Context) ExchangeArrayOutput {
	return o
}

func (o ExchangeArrayOutput) Index(i pulumi.IntInput) ExchangeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Exchange {
		return vs[0].([]*Exchange)[vs[1].(int)]
	}).(ExchangeOutput)
}

type ExchangeMapOutput struct{ *pulumi.OutputState }

func (ExchangeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Exchange)(nil)).Elem()
}

func (o ExchangeMapOutput) ToExchangeMapOutput() ExchangeMapOutput {
	return o
}

func (o ExchangeMapOutput) ToExchangeMapOutputWithContext(ctx context.Context) ExchangeMapOutput {
	return o
}

func (o ExchangeMapOutput) MapIndex(k pulumi.StringInput) ExchangeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Exchange {
		return vs[0].(map[string]*Exchange)[vs[1].(string)]
	}).(ExchangeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExchangeInput)(nil)).Elem(), &Exchange{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExchangeArrayInput)(nil)).Elem(), ExchangeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExchangeMapInput)(nil)).Elem(), ExchangeMap{})
	pulumi.RegisterOutputType(ExchangeOutput{})
	pulumi.RegisterOutputType(ExchangeArrayOutput{})
	pulumi.RegisterOutputType(ExchangeMapOutput{})
}
