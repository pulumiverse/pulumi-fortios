// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortianalyzer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Global FortiAnalyzer settings.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/log"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := log.NewSetting(ctx, "trname", &log.SettingArgs{
//				__changeIp:                pulumi.Int(0),
//				ConnTimeout:               pulumi.Int(10),
//				EncAlgorithm:              pulumi.String("high"),
//				FazType:                   pulumi.Int(1),
//				HmacAlgorithm:             pulumi.String("sha256"),
//				IpsArchive:                pulumi.String("enable"),
//				MgmtName:                  pulumi.String("FGh_Log1"),
//				MonitorFailureRetryPeriod: pulumi.Int(5),
//				MonitorKeepalivePeriod:    pulumi.Int(5),
//				Reliable:                  pulumi.String("disable"),
//				SslMinProtoVersion:        pulumi.String("default"),
//				Status:                    pulumi.String("disable"),
//				UploadInterval:            pulumi.String("daily"),
//				UploadOption:              pulumi.String("5-minute"),
//				UploadTime:                pulumi.String("00:59"),
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
// LogFortianalyzer Setting can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:log/fortianalyzer/setting:Setting labelname LogFortianalyzerSetting
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:log/fortianalyzer/setting:Setting labelname LogFortianalyzerSetting
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Setting struct {
	pulumi.CustomResourceState

	// Hidden attribute.
	__changeIp pulumi.IntOutput `pulumi:"__changeIp"`
	// Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
	AccessConfig pulumi.StringOutput `pulumi:"accessConfig"`
	// Alternate FortiAnalyzer.
	AltServer pulumi.StringOutput `pulumi:"altServer"`
	// Certificate used to communicate with FortiAnalyzer.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
	CertificateVerification pulumi.StringOutput `pulumi:"certificateVerification"`
	// FortiAnalyzer connection time-out in seconds (for status and log buffer).
	ConnTimeout pulumi.IntOutput `pulumi:"connTimeout"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable sending FortiAnalyzer log data with SSL encryption. Valid values: `high-medium`, `high`, `low`.
	EncAlgorithm pulumi.StringOutput `pulumi:"encAlgorithm"`
	// Enable/disable this FortiGate unit to fallback to the primary FortiAnalyzer when it is available. Valid values: `enable`, `disable`.
	FallbackToPrimary pulumi.StringOutput `pulumi:"fallbackToPrimary"`
	// Hidden setting index of FortiAnalyzer.
	FazType pulumi.IntOutput `pulumi:"fazType"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// FortiAnalyzer IPsec tunnel HMAC algorithm.
	HmacAlgorithm pulumi.StringOutput `pulumi:"hmacAlgorithm"`
	// Specify outgoing interface to reach server.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringOutput `pulumi:"interfaceSelectMethod"`
	// Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
	IpsArchive pulumi.StringOutput `pulumi:"ipsArchive"`
	// FortiAnalyzer maximum log rate in MBps (0 = unlimited).
	MaxLogRate pulumi.IntOutput `pulumi:"maxLogRate"`
	// Hidden management name of FortiAnalyzer.
	MgmtName pulumi.StringOutput `pulumi:"mgmtName"`
	// Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
	MonitorFailureRetryPeriod pulumi.IntOutput `pulumi:"monitorFailureRetryPeriod"`
	// Time between OFTP keepalives in seconds (for status and log buffer).
	MonitorKeepalivePeriod pulumi.IntOutput `pulumi:"monitorKeepalivePeriod"`
	// Preshared-key used for auto-authorization on FortiAnalyzer.
	PresharedKey pulumi.StringOutput `pulumi:"presharedKey"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringOutput `pulumi:"priority"`
	// Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Reliable pulumi.StringOutput `pulumi:"reliable"`
	// Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
	Serials SettingSerialArrayOutput `pulumi:"serials"`
	// The remote FortiAnalyzer.
	Server pulumi.StringOutput `pulumi:"server"`
	// Mandatory CA on FortiGate in certificate chain of server.
	ServerCertCa pulumi.StringOutput `pulumi:"serverCertCa"`
	// Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion pulumi.StringOutput `pulumi:"sslMinProtoVersion"`
	// Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Day of week (month) to upload logs.
	UploadDay pulumi.StringOutput `pulumi:"uploadDay"`
	// Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval pulumi.StringOutput `pulumi:"uploadInterval"`
	// Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption pulumi.StringOutput `pulumi:"uploadOption"`
	// Time to upload logs (hh:mm).
	UploadTime pulumi.StringOutput `pulumi:"uploadTime"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSetting registers a new resource with the given unique name, arguments, and options.
func NewSetting(ctx *pulumi.Context,
	name string, args *SettingArgs, opts ...pulumi.ResourceOption) (*Setting, error) {
	if args == nil {
		args = &SettingArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Setting
	err := ctx.RegisterResource("fortios:log/fortianalyzer/setting:Setting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSetting gets an existing Setting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SettingState, opts ...pulumi.ResourceOption) (*Setting, error) {
	var resource Setting
	err := ctx.ReadResource("fortios:log/fortianalyzer/setting:Setting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Setting resources.
type settingState struct {
	// Hidden attribute.
	__changeIp *int `pulumi:"__changeIp"`
	// Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
	AccessConfig *string `pulumi:"accessConfig"`
	// Alternate FortiAnalyzer.
	AltServer *string `pulumi:"altServer"`
	// Certificate used to communicate with FortiAnalyzer.
	Certificate *string `pulumi:"certificate"`
	// Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
	CertificateVerification *string `pulumi:"certificateVerification"`
	// FortiAnalyzer connection time-out in seconds (for status and log buffer).
	ConnTimeout *int `pulumi:"connTimeout"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable sending FortiAnalyzer log data with SSL encryption. Valid values: `high-medium`, `high`, `low`.
	EncAlgorithm *string `pulumi:"encAlgorithm"`
	// Enable/disable this FortiGate unit to fallback to the primary FortiAnalyzer when it is available. Valid values: `enable`, `disable`.
	FallbackToPrimary *string `pulumi:"fallbackToPrimary"`
	// Hidden setting index of FortiAnalyzer.
	FazType *int `pulumi:"fazType"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// FortiAnalyzer IPsec tunnel HMAC algorithm.
	HmacAlgorithm *string `pulumi:"hmacAlgorithm"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
	IpsArchive *string `pulumi:"ipsArchive"`
	// FortiAnalyzer maximum log rate in MBps (0 = unlimited).
	MaxLogRate *int `pulumi:"maxLogRate"`
	// Hidden management name of FortiAnalyzer.
	MgmtName *string `pulumi:"mgmtName"`
	// Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
	MonitorFailureRetryPeriod *int `pulumi:"monitorFailureRetryPeriod"`
	// Time between OFTP keepalives in seconds (for status and log buffer).
	MonitorKeepalivePeriod *int `pulumi:"monitorKeepalivePeriod"`
	// Preshared-key used for auto-authorization on FortiAnalyzer.
	PresharedKey *string `pulumi:"presharedKey"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority *string `pulumi:"priority"`
	// Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Reliable *string `pulumi:"reliable"`
	// Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
	Serials []SettingSerial `pulumi:"serials"`
	// The remote FortiAnalyzer.
	Server *string `pulumi:"server"`
	// Mandatory CA on FortiGate in certificate chain of server.
	ServerCertCa *string `pulumi:"serverCertCa"`
	// Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
	SourceIp *string `pulumi:"sourceIp"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Day of week (month) to upload logs.
	UploadDay *string `pulumi:"uploadDay"`
	// Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval *string `pulumi:"uploadInterval"`
	// Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption *string `pulumi:"uploadOption"`
	// Time to upload logs (hh:mm).
	UploadTime *string `pulumi:"uploadTime"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SettingState struct {
	// Hidden attribute.
	__changeIp pulumi.IntPtrInput
	// Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
	AccessConfig pulumi.StringPtrInput
	// Alternate FortiAnalyzer.
	AltServer pulumi.StringPtrInput
	// Certificate used to communicate with FortiAnalyzer.
	Certificate pulumi.StringPtrInput
	// Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
	CertificateVerification pulumi.StringPtrInput
	// FortiAnalyzer connection time-out in seconds (for status and log buffer).
	ConnTimeout pulumi.IntPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable sending FortiAnalyzer log data with SSL encryption. Valid values: `high-medium`, `high`, `low`.
	EncAlgorithm pulumi.StringPtrInput
	// Enable/disable this FortiGate unit to fallback to the primary FortiAnalyzer when it is available. Valid values: `enable`, `disable`.
	FallbackToPrimary pulumi.StringPtrInput
	// Hidden setting index of FortiAnalyzer.
	FazType pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// FortiAnalyzer IPsec tunnel HMAC algorithm.
	HmacAlgorithm pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
	IpsArchive pulumi.StringPtrInput
	// FortiAnalyzer maximum log rate in MBps (0 = unlimited).
	MaxLogRate pulumi.IntPtrInput
	// Hidden management name of FortiAnalyzer.
	MgmtName pulumi.StringPtrInput
	// Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
	MonitorFailureRetryPeriod pulumi.IntPtrInput
	// Time between OFTP keepalives in seconds (for status and log buffer).
	MonitorKeepalivePeriod pulumi.IntPtrInput
	// Preshared-key used for auto-authorization on FortiAnalyzer.
	PresharedKey pulumi.StringPtrInput
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringPtrInput
	// Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Reliable pulumi.StringPtrInput
	// Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
	Serials SettingSerialArrayInput
	// The remote FortiAnalyzer.
	Server pulumi.StringPtrInput
	// Mandatory CA on FortiGate in certificate chain of server.
	ServerCertCa pulumi.StringPtrInput
	// Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
	SourceIp pulumi.StringPtrInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion pulumi.StringPtrInput
	// Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Day of week (month) to upload logs.
	UploadDay pulumi.StringPtrInput
	// Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval pulumi.StringPtrInput
	// Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption pulumi.StringPtrInput
	// Time to upload logs (hh:mm).
	UploadTime pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*settingState)(nil)).Elem()
}

type settingArgs struct {
	// Hidden attribute.
	__changeIp *int `pulumi:"__changeIp"`
	// Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
	AccessConfig *string `pulumi:"accessConfig"`
	// Alternate FortiAnalyzer.
	AltServer *string `pulumi:"altServer"`
	// Certificate used to communicate with FortiAnalyzer.
	Certificate *string `pulumi:"certificate"`
	// Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
	CertificateVerification *string `pulumi:"certificateVerification"`
	// FortiAnalyzer connection time-out in seconds (for status and log buffer).
	ConnTimeout *int `pulumi:"connTimeout"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable sending FortiAnalyzer log data with SSL encryption. Valid values: `high-medium`, `high`, `low`.
	EncAlgorithm *string `pulumi:"encAlgorithm"`
	// Enable/disable this FortiGate unit to fallback to the primary FortiAnalyzer when it is available. Valid values: `enable`, `disable`.
	FallbackToPrimary *string `pulumi:"fallbackToPrimary"`
	// Hidden setting index of FortiAnalyzer.
	FazType *int `pulumi:"fazType"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// FortiAnalyzer IPsec tunnel HMAC algorithm.
	HmacAlgorithm *string `pulumi:"hmacAlgorithm"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
	IpsArchive *string `pulumi:"ipsArchive"`
	// FortiAnalyzer maximum log rate in MBps (0 = unlimited).
	MaxLogRate *int `pulumi:"maxLogRate"`
	// Hidden management name of FortiAnalyzer.
	MgmtName *string `pulumi:"mgmtName"`
	// Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
	MonitorFailureRetryPeriod *int `pulumi:"monitorFailureRetryPeriod"`
	// Time between OFTP keepalives in seconds (for status and log buffer).
	MonitorKeepalivePeriod *int `pulumi:"monitorKeepalivePeriod"`
	// Preshared-key used for auto-authorization on FortiAnalyzer.
	PresharedKey *string `pulumi:"presharedKey"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority *string `pulumi:"priority"`
	// Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Reliable *string `pulumi:"reliable"`
	// Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
	Serials []SettingSerial `pulumi:"serials"`
	// The remote FortiAnalyzer.
	Server *string `pulumi:"server"`
	// Mandatory CA on FortiGate in certificate chain of server.
	ServerCertCa *string `pulumi:"serverCertCa"`
	// Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
	SourceIp *string `pulumi:"sourceIp"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Day of week (month) to upload logs.
	UploadDay *string `pulumi:"uploadDay"`
	// Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval *string `pulumi:"uploadInterval"`
	// Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption *string `pulumi:"uploadOption"`
	// Time to upload logs (hh:mm).
	UploadTime *string `pulumi:"uploadTime"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Setting resource.
type SettingArgs struct {
	// Hidden attribute.
	__changeIp pulumi.IntPtrInput
	// Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
	AccessConfig pulumi.StringPtrInput
	// Alternate FortiAnalyzer.
	AltServer pulumi.StringPtrInput
	// Certificate used to communicate with FortiAnalyzer.
	Certificate pulumi.StringPtrInput
	// Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
	CertificateVerification pulumi.StringPtrInput
	// FortiAnalyzer connection time-out in seconds (for status and log buffer).
	ConnTimeout pulumi.IntPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable sending FortiAnalyzer log data with SSL encryption. Valid values: `high-medium`, `high`, `low`.
	EncAlgorithm pulumi.StringPtrInput
	// Enable/disable this FortiGate unit to fallback to the primary FortiAnalyzer when it is available. Valid values: `enable`, `disable`.
	FallbackToPrimary pulumi.StringPtrInput
	// Hidden setting index of FortiAnalyzer.
	FazType pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// FortiAnalyzer IPsec tunnel HMAC algorithm.
	HmacAlgorithm pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
	IpsArchive pulumi.StringPtrInput
	// FortiAnalyzer maximum log rate in MBps (0 = unlimited).
	MaxLogRate pulumi.IntPtrInput
	// Hidden management name of FortiAnalyzer.
	MgmtName pulumi.StringPtrInput
	// Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
	MonitorFailureRetryPeriod pulumi.IntPtrInput
	// Time between OFTP keepalives in seconds (for status and log buffer).
	MonitorKeepalivePeriod pulumi.IntPtrInput
	// Preshared-key used for auto-authorization on FortiAnalyzer.
	PresharedKey pulumi.StringPtrInput
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringPtrInput
	// Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Reliable pulumi.StringPtrInput
	// Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
	Serials SettingSerialArrayInput
	// The remote FortiAnalyzer.
	Server pulumi.StringPtrInput
	// Mandatory CA on FortiGate in certificate chain of server.
	ServerCertCa pulumi.StringPtrInput
	// Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
	SourceIp pulumi.StringPtrInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion pulumi.StringPtrInput
	// Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Day of week (month) to upload logs.
	UploadDay pulumi.StringPtrInput
	// Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval pulumi.StringPtrInput
	// Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption pulumi.StringPtrInput
	// Time to upload logs (hh:mm).
	UploadTime pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*settingArgs)(nil)).Elem()
}

type SettingInput interface {
	pulumi.Input

	ToSettingOutput() SettingOutput
	ToSettingOutputWithContext(ctx context.Context) SettingOutput
}

func (*Setting) ElementType() reflect.Type {
	return reflect.TypeOf((**Setting)(nil)).Elem()
}

func (i *Setting) ToSettingOutput() SettingOutput {
	return i.ToSettingOutputWithContext(context.Background())
}

func (i *Setting) ToSettingOutputWithContext(ctx context.Context) SettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingOutput)
}

// SettingArrayInput is an input type that accepts SettingArray and SettingArrayOutput values.
// You can construct a concrete instance of `SettingArrayInput` via:
//
//	SettingArray{ SettingArgs{...} }
type SettingArrayInput interface {
	pulumi.Input

	ToSettingArrayOutput() SettingArrayOutput
	ToSettingArrayOutputWithContext(context.Context) SettingArrayOutput
}

type SettingArray []SettingInput

func (SettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Setting)(nil)).Elem()
}

func (i SettingArray) ToSettingArrayOutput() SettingArrayOutput {
	return i.ToSettingArrayOutputWithContext(context.Background())
}

func (i SettingArray) ToSettingArrayOutputWithContext(ctx context.Context) SettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingArrayOutput)
}

// SettingMapInput is an input type that accepts SettingMap and SettingMapOutput values.
// You can construct a concrete instance of `SettingMapInput` via:
//
//	SettingMap{ "key": SettingArgs{...} }
type SettingMapInput interface {
	pulumi.Input

	ToSettingMapOutput() SettingMapOutput
	ToSettingMapOutputWithContext(context.Context) SettingMapOutput
}

type SettingMap map[string]SettingInput

func (SettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Setting)(nil)).Elem()
}

func (i SettingMap) ToSettingMapOutput() SettingMapOutput {
	return i.ToSettingMapOutputWithContext(context.Background())
}

func (i SettingMap) ToSettingMapOutputWithContext(ctx context.Context) SettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingMapOutput)
}

type SettingOutput struct{ *pulumi.OutputState }

func (SettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Setting)(nil)).Elem()
}

func (o SettingOutput) ToSettingOutput() SettingOutput {
	return o
}

func (o SettingOutput) ToSettingOutputWithContext(ctx context.Context) SettingOutput {
	return o
}

// Hidden attribute.
func (o SettingOutput) __changeIp() pulumi.IntOutput {
	return o.ApplyT(func(v *Setting) pulumi.IntOutput { return v.__changeIp }).(pulumi.IntOutput)
}

// Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
func (o SettingOutput) AccessConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.AccessConfig }).(pulumi.StringOutput)
}

// Alternate FortiAnalyzer.
func (o SettingOutput) AltServer() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.AltServer }).(pulumi.StringOutput)
}

// Certificate used to communicate with FortiAnalyzer.
func (o SettingOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
func (o SettingOutput) CertificateVerification() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.CertificateVerification }).(pulumi.StringOutput)
}

// FortiAnalyzer connection time-out in seconds (for status and log buffer).
func (o SettingOutput) ConnTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Setting) pulumi.IntOutput { return v.ConnTimeout }).(pulumi.IntOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SettingOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Enable/disable sending FortiAnalyzer log data with SSL encryption. Valid values: `high-medium`, `high`, `low`.
func (o SettingOutput) EncAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.EncAlgorithm }).(pulumi.StringOutput)
}

// Enable/disable this FortiGate unit to fallback to the primary FortiAnalyzer when it is available. Valid values: `enable`, `disable`.
func (o SettingOutput) FallbackToPrimary() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.FallbackToPrimary }).(pulumi.StringOutput)
}

// Hidden setting index of FortiAnalyzer.
func (o SettingOutput) FazType() pulumi.IntOutput {
	return o.ApplyT(func(v *Setting) pulumi.IntOutput { return v.FazType }).(pulumi.IntOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o SettingOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// FortiAnalyzer IPsec tunnel HMAC algorithm.
func (o SettingOutput) HmacAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.HmacAlgorithm }).(pulumi.StringOutput)
}

// Specify outgoing interface to reach server.
func (o SettingOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
func (o SettingOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

// Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
func (o SettingOutput) IpsArchive() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.IpsArchive }).(pulumi.StringOutput)
}

// FortiAnalyzer maximum log rate in MBps (0 = unlimited).
func (o SettingOutput) MaxLogRate() pulumi.IntOutput {
	return o.ApplyT(func(v *Setting) pulumi.IntOutput { return v.MaxLogRate }).(pulumi.IntOutput)
}

// Hidden management name of FortiAnalyzer.
func (o SettingOutput) MgmtName() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.MgmtName }).(pulumi.StringOutput)
}

// Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
func (o SettingOutput) MonitorFailureRetryPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *Setting) pulumi.IntOutput { return v.MonitorFailureRetryPeriod }).(pulumi.IntOutput)
}

// Time between OFTP keepalives in seconds (for status and log buffer).
func (o SettingOutput) MonitorKeepalivePeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *Setting) pulumi.IntOutput { return v.MonitorKeepalivePeriod }).(pulumi.IntOutput)
}

// Preshared-key used for auto-authorization on FortiAnalyzer.
func (o SettingOutput) PresharedKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.PresharedKey }).(pulumi.StringOutput)
}

// Set log transmission priority. Valid values: `default`, `low`.
func (o SettingOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.Priority }).(pulumi.StringOutput)
}

// Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
func (o SettingOutput) Reliable() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.Reliable }).(pulumi.StringOutput)
}

// Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
func (o SettingOutput) Serials() SettingSerialArrayOutput {
	return o.ApplyT(func(v *Setting) SettingSerialArrayOutput { return v.Serials }).(SettingSerialArrayOutput)
}

// The remote FortiAnalyzer.
func (o SettingOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

// Mandatory CA on FortiGate in certificate chain of server.
func (o SettingOutput) ServerCertCa() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.ServerCertCa }).(pulumi.StringOutput)
}

// Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
func (o SettingOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
func (o SettingOutput) SslMinProtoVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.SslMinProtoVersion }).(pulumi.StringOutput)
}

// Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
func (o SettingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Day of week (month) to upload logs.
func (o SettingOutput) UploadDay() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.UploadDay }).(pulumi.StringOutput)
}

// Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
func (o SettingOutput) UploadInterval() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.UploadInterval }).(pulumi.StringOutput)
}

// Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
func (o SettingOutput) UploadOption() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.UploadOption }).(pulumi.StringOutput)
}

// Time to upload logs (hh:mm).
func (o SettingOutput) UploadTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.UploadTime }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SettingArrayOutput struct{ *pulumi.OutputState }

func (SettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Setting)(nil)).Elem()
}

func (o SettingArrayOutput) ToSettingArrayOutput() SettingArrayOutput {
	return o
}

func (o SettingArrayOutput) ToSettingArrayOutputWithContext(ctx context.Context) SettingArrayOutput {
	return o
}

func (o SettingArrayOutput) Index(i pulumi.IntInput) SettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Setting {
		return vs[0].([]*Setting)[vs[1].(int)]
	}).(SettingOutput)
}

type SettingMapOutput struct{ *pulumi.OutputState }

func (SettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Setting)(nil)).Elem()
}

func (o SettingMapOutput) ToSettingMapOutput() SettingMapOutput {
	return o
}

func (o SettingMapOutput) ToSettingMapOutputWithContext(ctx context.Context) SettingMapOutput {
	return o
}

func (o SettingMapOutput) MapIndex(k pulumi.StringInput) SettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Setting {
		return vs[0].(map[string]*Setting)[vs[1].(string)]
	}).(SettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SettingInput)(nil)).Elem(), &Setting{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingArrayInput)(nil)).Elem(), SettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingMapInput)(nil)).Elem(), SettingMap{})
	pulumi.RegisterOutputType(SettingOutput{})
	pulumi.RegisterOutputType(SettingArrayOutput{})
	pulumi.RegisterOutputType(SettingMapOutput{})
}
