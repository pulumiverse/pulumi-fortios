// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Certificate
{
    /// <summary>
    /// VPN certificate setting.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.Vpn.Certificate.Setting("trname", new()
    ///     {
    ///         CertnameDsa1024 = "Fortinet_SSL_DSA1024",
    ///         CertnameDsa2048 = "Fortinet_SSL_DSA2048",
    ///         CertnameEcdsa256 = "Fortinet_SSL_ECDSA256",
    ///         CertnameEcdsa384 = "Fortinet_SSL_ECDSA384",
    ///         CertnameRsa1024 = "Fortinet_SSL_RSA1024",
    ///         CertnameRsa2048 = "Fortinet_SSL_RSA2048",
    ///         CheckCaCert = "enable",
    ///         CheckCaChain = "disable",
    ///         CmpSaveExtraCerts = "disable",
    ///         CnMatch = "substring",
    ///         OcspOption = "server",
    ///         OcspStatus = "disable",
    ///         SslMinProtoVersion = "default",
    ///         StrictCrlCheck = "disable",
    ///         StrictOcspCheck = "disable",
    ///         SubjectMatch = "substring",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// VpnCertificate Setting can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:vpn/certificate/setting:Setting labelname VpnCertificateSetting
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:vpn/certificate/setting:Setting labelname VpnCertificateSetting
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:vpn/certificate/setting:Setting")]
    public partial class Setting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Number of days before a certificate expires to send a warning. Set to 0 to disable sending of the warning (0 - 100, default = 14).
        /// </summary>
        [Output("certExpireWarning")]
        public Output<int> CertExpireWarning { get; private set; } = null!;

        /// <summary>
        /// 1024 bit DSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Output("certnameDsa1024")]
        public Output<string> CertnameDsa1024 { get; private set; } = null!;

        /// <summary>
        /// 2048 bit DSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Output("certnameDsa2048")]
        public Output<string> CertnameDsa2048 { get; private set; } = null!;

        /// <summary>
        /// 256 bit ECDSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Output("certnameEcdsa256")]
        public Output<string> CertnameEcdsa256 { get; private set; } = null!;

        /// <summary>
        /// 384 bit ECDSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Output("certnameEcdsa384")]
        public Output<string> CertnameEcdsa384 { get; private set; } = null!;

        /// <summary>
        /// 521 bit ECDSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Output("certnameEcdsa521")]
        public Output<string> CertnameEcdsa521 { get; private set; } = null!;

        /// <summary>
        /// 253 bit EdDSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Output("certnameEd25519")]
        public Output<string> CertnameEd25519 { get; private set; } = null!;

        /// <summary>
        /// 456 bit EdDSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Output("certnameEd448")]
        public Output<string> CertnameEd448 { get; private set; } = null!;

        /// <summary>
        /// 1024 bit RSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Output("certnameRsa1024")]
        public Output<string> CertnameRsa1024 { get; private set; } = null!;

        /// <summary>
        /// 2048 bit RSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Output("certnameRsa2048")]
        public Output<string> CertnameRsa2048 { get; private set; } = null!;

        /// <summary>
        /// 4096 bit RSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Output("certnameRsa4096")]
        public Output<string> CertnameRsa4096 { get; private set; } = null!;

        /// <summary>
        /// Enable/disable verification of the user certificate and pass authentication if any CA in the chain is trusted (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("checkCaCert")]
        public Output<string> CheckCaCert { get; private set; } = null!;

        /// <summary>
        /// Enable/disable verification of the entire certificate chain and pass authentication only if the chain is complete and all of the CAs in the chain are trusted (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("checkCaChain")]
        public Output<string> CheckCaChain { get; private set; } = null!;

        /// <summary>
        /// Enable/disable server certificate key usage checking in CMP mode (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("cmpKeyUsageChecking")]
        public Output<string> CmpKeyUsageChecking { get; private set; } = null!;

        /// <summary>
        /// Enable/disable saving extra certificates in CMP mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("cmpSaveExtraCerts")]
        public Output<string> CmpSaveExtraCerts { get; private set; } = null!;

        /// <summary>
        /// When searching for a matching certificate, allow mutliple CN fields in certificate subject name (default = enable). Valid values: `disable`, `enable`.
        /// </summary>
        [Output("cnAllowMulti")]
        public Output<string> CnAllowMulti { get; private set; } = null!;

        /// <summary>
        /// When searching for a matching certificate, control how to find matches in the cn attribute of the certificate subject name. Valid values: `substring`, `value`.
        /// </summary>
        [Output("cnMatch")]
        public Output<string> CnMatch { get; private set; } = null!;

        /// <summary>
        /// CRL verification options. The structure of `crl_verification` block is documented below.
        /// </summary>
        [Output("crlVerification")]
        public Output<Outputs.SettingCrlVerification> CrlVerification { get; private set; } = null!;

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Output("interfaceSelectMethod")]
        public Output<string> InterfaceSelectMethod { get; private set; } = null!;

        /// <summary>
        /// Default OCSP server.
        /// </summary>
        [Output("ocspDefaultServer")]
        public Output<string> OcspDefaultServer { get; private set; } = null!;

        /// <summary>
        /// Specify whether the OCSP URL is from certificate or configured OCSP server. Valid values: `certificate`, `server`.
        /// </summary>
        [Output("ocspOption")]
        public Output<string> OcspOption { get; private set; } = null!;

        /// <summary>
        /// Enable/disable receiving certificates using the OCSP. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ocspStatus")]
        public Output<string> OcspStatus { get; private set; } = null!;

        /// <summary>
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        /// </summary>
        [Output("sslMinProtoVersion")]
        public Output<string> SslMinProtoVersion { get; private set; } = null!;

        /// <summary>
        /// Source IP address to use to communicate with the OCSP server.
        /// </summary>
        [Output("sslOcspSourceIp")]
        public Output<string> SslOcspSourceIp { get; private set; } = null!;

        /// <summary>
        /// Enable/disable strict mode CRL checking. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("strictCrlCheck")]
        public Output<string> StrictCrlCheck { get; private set; } = null!;

        /// <summary>
        /// Enable/disable strict mode OCSP checking. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("strictOcspCheck")]
        public Output<string> StrictOcspCheck { get; private set; } = null!;

        /// <summary>
        /// When searching for a matching certificate, control how to find matches in the certificate subject name. Valid values: `substring`, `value`.
        /// </summary>
        [Output("subjectMatch")]
        public Output<string> SubjectMatch { get; private set; } = null!;

        /// <summary>
        /// When searching for a matching certificate, control how to do RDN set matching with certificate subject name (default = subset). Valid values: `subset`, `superset`.
        /// </summary>
        [Output("subjectSet")]
        public Output<string> SubjectSet { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Setting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Setting(string name, SettingArgs args, CustomResourceOptions? options = null)
            : base("fortios:vpn/certificate/setting:Setting", name, args ?? new SettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Setting(string name, Input<string> id, SettingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:vpn/certificate/setting:Setting", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Setting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Setting Get(string name, Input<string> id, SettingState? state = null, CustomResourceOptions? options = null)
        {
            return new Setting(name, id, state, options);
        }
    }

    public sealed class SettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of days before a certificate expires to send a warning. Set to 0 to disable sending of the warning (0 - 100, default = 14).
        /// </summary>
        [Input("certExpireWarning")]
        public Input<int>? CertExpireWarning { get; set; }

        /// <summary>
        /// 1024 bit DSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameDsa1024", required: true)]
        public Input<string> CertnameDsa1024 { get; set; } = null!;

        /// <summary>
        /// 2048 bit DSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameDsa2048", required: true)]
        public Input<string> CertnameDsa2048 { get; set; } = null!;

        /// <summary>
        /// 256 bit ECDSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameEcdsa256", required: true)]
        public Input<string> CertnameEcdsa256 { get; set; } = null!;

        /// <summary>
        /// 384 bit ECDSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameEcdsa384", required: true)]
        public Input<string> CertnameEcdsa384 { get; set; } = null!;

        /// <summary>
        /// 521 bit ECDSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameEcdsa521")]
        public Input<string>? CertnameEcdsa521 { get; set; }

        /// <summary>
        /// 253 bit EdDSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameEd25519")]
        public Input<string>? CertnameEd25519 { get; set; }

        /// <summary>
        /// 456 bit EdDSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameEd448")]
        public Input<string>? CertnameEd448 { get; set; }

        /// <summary>
        /// 1024 bit RSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameRsa1024", required: true)]
        public Input<string> CertnameRsa1024 { get; set; } = null!;

        /// <summary>
        /// 2048 bit RSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameRsa2048", required: true)]
        public Input<string> CertnameRsa2048 { get; set; } = null!;

        /// <summary>
        /// 4096 bit RSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameRsa4096")]
        public Input<string>? CertnameRsa4096 { get; set; }

        /// <summary>
        /// Enable/disable verification of the user certificate and pass authentication if any CA in the chain is trusted (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("checkCaCert")]
        public Input<string>? CheckCaCert { get; set; }

        /// <summary>
        /// Enable/disable verification of the entire certificate chain and pass authentication only if the chain is complete and all of the CAs in the chain are trusted (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("checkCaChain")]
        public Input<string>? CheckCaChain { get; set; }

        /// <summary>
        /// Enable/disable server certificate key usage checking in CMP mode (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("cmpKeyUsageChecking")]
        public Input<string>? CmpKeyUsageChecking { get; set; }

        /// <summary>
        /// Enable/disable saving extra certificates in CMP mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("cmpSaveExtraCerts")]
        public Input<string>? CmpSaveExtraCerts { get; set; }

        /// <summary>
        /// When searching for a matching certificate, allow mutliple CN fields in certificate subject name (default = enable). Valid values: `disable`, `enable`.
        /// </summary>
        [Input("cnAllowMulti")]
        public Input<string>? CnAllowMulti { get; set; }

        /// <summary>
        /// When searching for a matching certificate, control how to find matches in the cn attribute of the certificate subject name. Valid values: `substring`, `value`.
        /// </summary>
        [Input("cnMatch")]
        public Input<string>? CnMatch { get; set; }

        /// <summary>
        /// CRL verification options. The structure of `crl_verification` block is documented below.
        /// </summary>
        [Input("crlVerification")]
        public Input<Inputs.SettingCrlVerificationArgs>? CrlVerification { get; set; }

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// Default OCSP server.
        /// </summary>
        [Input("ocspDefaultServer")]
        public Input<string>? OcspDefaultServer { get; set; }

        /// <summary>
        /// Specify whether the OCSP URL is from certificate or configured OCSP server. Valid values: `certificate`, `server`.
        /// </summary>
        [Input("ocspOption")]
        public Input<string>? OcspOption { get; set; }

        /// <summary>
        /// Enable/disable receiving certificates using the OCSP. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ocspStatus")]
        public Input<string>? OcspStatus { get; set; }

        /// <summary>
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        /// </summary>
        [Input("sslMinProtoVersion")]
        public Input<string>? SslMinProtoVersion { get; set; }

        /// <summary>
        /// Source IP address to use to communicate with the OCSP server.
        /// </summary>
        [Input("sslOcspSourceIp")]
        public Input<string>? SslOcspSourceIp { get; set; }

        /// <summary>
        /// Enable/disable strict mode CRL checking. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("strictCrlCheck")]
        public Input<string>? StrictCrlCheck { get; set; }

        /// <summary>
        /// Enable/disable strict mode OCSP checking. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("strictOcspCheck")]
        public Input<string>? StrictOcspCheck { get; set; }

        /// <summary>
        /// When searching for a matching certificate, control how to find matches in the certificate subject name. Valid values: `substring`, `value`.
        /// </summary>
        [Input("subjectMatch")]
        public Input<string>? SubjectMatch { get; set; }

        /// <summary>
        /// When searching for a matching certificate, control how to do RDN set matching with certificate subject name (default = subset). Valid values: `subset`, `superset`.
        /// </summary>
        [Input("subjectSet")]
        public Input<string>? SubjectSet { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SettingArgs()
        {
        }
        public static new SettingArgs Empty => new SettingArgs();
    }

    public sealed class SettingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of days before a certificate expires to send a warning. Set to 0 to disable sending of the warning (0 - 100, default = 14).
        /// </summary>
        [Input("certExpireWarning")]
        public Input<int>? CertExpireWarning { get; set; }

        /// <summary>
        /// 1024 bit DSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameDsa1024")]
        public Input<string>? CertnameDsa1024 { get; set; }

        /// <summary>
        /// 2048 bit DSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameDsa2048")]
        public Input<string>? CertnameDsa2048 { get; set; }

        /// <summary>
        /// 256 bit ECDSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameEcdsa256")]
        public Input<string>? CertnameEcdsa256 { get; set; }

        /// <summary>
        /// 384 bit ECDSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameEcdsa384")]
        public Input<string>? CertnameEcdsa384 { get; set; }

        /// <summary>
        /// 521 bit ECDSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameEcdsa521")]
        public Input<string>? CertnameEcdsa521 { get; set; }

        /// <summary>
        /// 253 bit EdDSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameEd25519")]
        public Input<string>? CertnameEd25519 { get; set; }

        /// <summary>
        /// 456 bit EdDSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameEd448")]
        public Input<string>? CertnameEd448 { get; set; }

        /// <summary>
        /// 1024 bit RSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameRsa1024")]
        public Input<string>? CertnameRsa1024 { get; set; }

        /// <summary>
        /// 2048 bit RSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameRsa2048")]
        public Input<string>? CertnameRsa2048 { get; set; }

        /// <summary>
        /// 4096 bit RSA key certificate for re-signing server certificates for SSL inspection.
        /// </summary>
        [Input("certnameRsa4096")]
        public Input<string>? CertnameRsa4096 { get; set; }

        /// <summary>
        /// Enable/disable verification of the user certificate and pass authentication if any CA in the chain is trusted (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("checkCaCert")]
        public Input<string>? CheckCaCert { get; set; }

        /// <summary>
        /// Enable/disable verification of the entire certificate chain and pass authentication only if the chain is complete and all of the CAs in the chain are trusted (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("checkCaChain")]
        public Input<string>? CheckCaChain { get; set; }

        /// <summary>
        /// Enable/disable server certificate key usage checking in CMP mode (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("cmpKeyUsageChecking")]
        public Input<string>? CmpKeyUsageChecking { get; set; }

        /// <summary>
        /// Enable/disable saving extra certificates in CMP mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("cmpSaveExtraCerts")]
        public Input<string>? CmpSaveExtraCerts { get; set; }

        /// <summary>
        /// When searching for a matching certificate, allow mutliple CN fields in certificate subject name (default = enable). Valid values: `disable`, `enable`.
        /// </summary>
        [Input("cnAllowMulti")]
        public Input<string>? CnAllowMulti { get; set; }

        /// <summary>
        /// When searching for a matching certificate, control how to find matches in the cn attribute of the certificate subject name. Valid values: `substring`, `value`.
        /// </summary>
        [Input("cnMatch")]
        public Input<string>? CnMatch { get; set; }

        /// <summary>
        /// CRL verification options. The structure of `crl_verification` block is documented below.
        /// </summary>
        [Input("crlVerification")]
        public Input<Inputs.SettingCrlVerificationGetArgs>? CrlVerification { get; set; }

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// Default OCSP server.
        /// </summary>
        [Input("ocspDefaultServer")]
        public Input<string>? OcspDefaultServer { get; set; }

        /// <summary>
        /// Specify whether the OCSP URL is from certificate or configured OCSP server. Valid values: `certificate`, `server`.
        /// </summary>
        [Input("ocspOption")]
        public Input<string>? OcspOption { get; set; }

        /// <summary>
        /// Enable/disable receiving certificates using the OCSP. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ocspStatus")]
        public Input<string>? OcspStatus { get; set; }

        /// <summary>
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        /// </summary>
        [Input("sslMinProtoVersion")]
        public Input<string>? SslMinProtoVersion { get; set; }

        /// <summary>
        /// Source IP address to use to communicate with the OCSP server.
        /// </summary>
        [Input("sslOcspSourceIp")]
        public Input<string>? SslOcspSourceIp { get; set; }

        /// <summary>
        /// Enable/disable strict mode CRL checking. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("strictCrlCheck")]
        public Input<string>? StrictCrlCheck { get; set; }

        /// <summary>
        /// Enable/disable strict mode OCSP checking. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("strictOcspCheck")]
        public Input<string>? StrictOcspCheck { get; set; }

        /// <summary>
        /// When searching for a matching certificate, control how to find matches in the certificate subject name. Valid values: `substring`, `value`.
        /// </summary>
        [Input("subjectMatch")]
        public Input<string>? SubjectMatch { get; set; }

        /// <summary>
        /// When searching for a matching certificate, control how to do RDN set matching with certificate subject name (default = subset). Valid values: `subset`, `superset`.
        /// </summary>
        [Input("subjectSet")]
        public Input<string>? SubjectSet { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SettingState()
        {
        }
        public static new SettingState Empty => new SettingState();
    }
}
