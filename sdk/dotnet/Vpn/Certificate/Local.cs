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
    /// Local keys and certificates.
    /// 
    /// ## Import
    /// 
    /// VpnCertificate Local can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:vpn/certificate/local:Local labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:vpn/certificate/local:Local labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:vpn/certificate/local:Local")]
    public partial class Local : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The URL for the ACME CA server (Let's Encrypt is the default provider).
        /// </summary>
        [Output("acmeCaUrl")]
        public Output<string> AcmeCaUrl { get; private set; } = null!;

        /// <summary>
        /// A valid domain that resolves to this Fortigate.
        /// </summary>
        [Output("acmeDomain")]
        public Output<string> AcmeDomain { get; private set; } = null!;

        /// <summary>
        /// Contact email address that is required by some CAs like LetsEncrypt.
        /// </summary>
        [Output("acmeEmail")]
        public Output<string> AcmeEmail { get; private set; } = null!;

        /// <summary>
        /// Beginning of the renewal window (in days before certificate expiration, 30 by default).
        /// </summary>
        [Output("acmeRenewWindow")]
        public Output<int> AcmeRenewWindow { get; private set; } = null!;

        /// <summary>
        /// Length of the RSA private key of the generated cert (Minimum 2048 bits).
        /// </summary>
        [Output("acmeRsaKeySize")]
        public Output<int> AcmeRsaKeySize { get; private set; } = null!;

        /// <summary>
        /// Number of days to wait before expiry of an updated local certificate is requested (0 = disabled).
        /// </summary>
        [Output("autoRegenerateDays")]
        public Output<int> AutoRegenerateDays { get; private set; } = null!;

        /// <summary>
        /// Number of days to wait before an expiry warning message is generated (0 = disabled).
        /// </summary>
        [Output("autoRegenerateDaysWarning")]
        public Output<int> AutoRegenerateDaysWarning { get; private set; } = null!;

        /// <summary>
        /// CA identifier of the CA server for signing via SCEP.
        /// </summary>
        [Output("caIdentifier")]
        public Output<string> CaIdentifier { get; private set; } = null!;

        /// <summary>
        /// PEM format certificate.
        /// </summary>
        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        /// <summary>
        /// Path location inside CMP server.
        /// </summary>
        [Output("cmpPath")]
        public Output<string> CmpPath { get; private set; } = null!;

        /// <summary>
        /// CMP auto-regeneration method. Valid values: `keyupate`, `renewal`.
        /// </summary>
        [Output("cmpRegenerationMethod")]
        public Output<string> CmpRegenerationMethod { get; private set; } = null!;

        /// <summary>
        /// 'ADDRESS:PORT' for CMP server.
        /// </summary>
        [Output("cmpServer")]
        public Output<string> CmpServer { get; private set; } = null!;

        /// <summary>
        /// CMP server certificate.
        /// </summary>
        [Output("cmpServerCert")]
        public Output<string> CmpServerCert { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comments")]
        public Output<string> Comments { get; private set; } = null!;

        /// <summary>
        /// Certificate Signing Request.
        /// </summary>
        [Output("csr")]
        public Output<string> Csr { get; private set; } = null!;

        /// <summary>
        /// Certificate enrollment protocol.
        /// </summary>
        [Output("enrollProtocol")]
        public Output<string> EnrollProtocol { get; private set; } = null!;

        /// <summary>
        /// Local ID the FortiGate uses for authentication as a VPN client.
        /// </summary>
        [Output("ikeLocalid")]
        public Output<string> IkeLocalid { get; private set; } = null!;

        /// <summary>
        /// IKE local ID type. Valid values: `asn1dn`, `fqdn`.
        /// </summary>
        [Output("ikeLocalidType")]
        public Output<string> IkeLocalidType { get; private set; } = null!;

        /// <summary>
        /// Time at which certificate was last updated.
        /// </summary>
        [Output("lastUpdated")]
        public Output<int> LastUpdated { get; private set; } = null!;

        /// <summary>
        /// Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Name encoding method for auto-regeneration. Valid values: `printable`, `utf8`.
        /// </summary>
        [Output("nameEncoding")]
        public Output<string> NameEncoding { get; private set; } = null!;

        /// <summary>
        /// Password as a PEM file.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// PEM format key, encrypted with a password.
        /// </summary>
        [Output("privateKey")]
        public Output<string> PrivateKey { get; private set; } = null!;

        /// <summary>
        /// Enable/disable retention of private key during SCEP renewal (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("privateKeyRetain")]
        public Output<string> PrivateKeyRetain { get; private set; } = null!;

        /// <summary>
        /// Either a global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
        /// </summary>
        [Output("range")]
        public Output<string> Range { get; private set; } = null!;

        /// <summary>
        /// SCEP server challenge password for auto-regeneration.
        /// </summary>
        [Output("scepPassword")]
        public Output<string?> ScepPassword { get; private set; } = null!;

        /// <summary>
        /// SCEP server URL.
        /// </summary>
        [Output("scepUrl")]
        public Output<string> ScepUrl { get; private set; } = null!;

        /// <summary>
        /// Certificate source type.
        /// </summary>
        [Output("source")]
        public Output<string> Source { get; private set; } = null!;

        /// <summary>
        /// Source IP address for communications to the SCEP server.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// Certificate Signing Request State.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Local resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Local(string name, LocalArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:vpn/certificate/local:Local", name, args ?? new LocalArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Local(string name, Input<string> id, LocalState? state = null, CustomResourceOptions? options = null)
            : base("fortios:vpn/certificate/local:Local", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
                AdditionalSecretOutputs =
                {
                    "certificate",
                    "password",
                    "privateKey",
                    "scepPassword",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Local resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Local Get(string name, Input<string> id, LocalState? state = null, CustomResourceOptions? options = null)
        {
            return new Local(name, id, state, options);
        }
    }

    public sealed class LocalArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL for the ACME CA server (Let's Encrypt is the default provider).
        /// </summary>
        [Input("acmeCaUrl")]
        public Input<string>? AcmeCaUrl { get; set; }

        /// <summary>
        /// A valid domain that resolves to this Fortigate.
        /// </summary>
        [Input("acmeDomain")]
        public Input<string>? AcmeDomain { get; set; }

        /// <summary>
        /// Contact email address that is required by some CAs like LetsEncrypt.
        /// </summary>
        [Input("acmeEmail")]
        public Input<string>? AcmeEmail { get; set; }

        /// <summary>
        /// Beginning of the renewal window (in days before certificate expiration, 30 by default).
        /// </summary>
        [Input("acmeRenewWindow")]
        public Input<int>? AcmeRenewWindow { get; set; }

        /// <summary>
        /// Length of the RSA private key of the generated cert (Minimum 2048 bits).
        /// </summary>
        [Input("acmeRsaKeySize")]
        public Input<int>? AcmeRsaKeySize { get; set; }

        /// <summary>
        /// Number of days to wait before expiry of an updated local certificate is requested (0 = disabled).
        /// </summary>
        [Input("autoRegenerateDays")]
        public Input<int>? AutoRegenerateDays { get; set; }

        /// <summary>
        /// Number of days to wait before an expiry warning message is generated (0 = disabled).
        /// </summary>
        [Input("autoRegenerateDaysWarning")]
        public Input<int>? AutoRegenerateDaysWarning { get; set; }

        /// <summary>
        /// CA identifier of the CA server for signing via SCEP.
        /// </summary>
        [Input("caIdentifier")]
        public Input<string>? CaIdentifier { get; set; }

        [Input("certificate")]
        private Input<string>? _certificate;

        /// <summary>
        /// PEM format certificate.
        /// </summary>
        public Input<string>? Certificate
        {
            get => _certificate;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _certificate = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Path location inside CMP server.
        /// </summary>
        [Input("cmpPath")]
        public Input<string>? CmpPath { get; set; }

        /// <summary>
        /// CMP auto-regeneration method. Valid values: `keyupate`, `renewal`.
        /// </summary>
        [Input("cmpRegenerationMethod")]
        public Input<string>? CmpRegenerationMethod { get; set; }

        /// <summary>
        /// 'ADDRESS:PORT' for CMP server.
        /// </summary>
        [Input("cmpServer")]
        public Input<string>? CmpServer { get; set; }

        /// <summary>
        /// CMP server certificate.
        /// </summary>
        [Input("cmpServerCert")]
        public Input<string>? CmpServerCert { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Certificate Signing Request.
        /// </summary>
        [Input("csr")]
        public Input<string>? Csr { get; set; }

        /// <summary>
        /// Certificate enrollment protocol.
        /// </summary>
        [Input("enrollProtocol")]
        public Input<string>? EnrollProtocol { get; set; }

        /// <summary>
        /// Local ID the FortiGate uses for authentication as a VPN client.
        /// </summary>
        [Input("ikeLocalid")]
        public Input<string>? IkeLocalid { get; set; }

        /// <summary>
        /// IKE local ID type. Valid values: `asn1dn`, `fqdn`.
        /// </summary>
        [Input("ikeLocalidType")]
        public Input<string>? IkeLocalidType { get; set; }

        /// <summary>
        /// Time at which certificate was last updated.
        /// </summary>
        [Input("lastUpdated")]
        public Input<int>? LastUpdated { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name encoding method for auto-regeneration. Valid values: `printable`, `utf8`.
        /// </summary>
        [Input("nameEncoding")]
        public Input<string>? NameEncoding { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password as a PEM file.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateKey")]
        private Input<string>? _privateKey;

        /// <summary>
        /// PEM format key, encrypted with a password.
        /// </summary>
        public Input<string>? PrivateKey
        {
            get => _privateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Enable/disable retention of private key during SCEP renewal (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("privateKeyRetain")]
        public Input<string>? PrivateKeyRetain { get; set; }

        /// <summary>
        /// Either a global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
        /// </summary>
        [Input("range")]
        public Input<string>? Range { get; set; }

        [Input("scepPassword")]
        private Input<string>? _scepPassword;

        /// <summary>
        /// SCEP server challenge password for auto-regeneration.
        /// </summary>
        public Input<string>? ScepPassword
        {
            get => _scepPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _scepPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// SCEP server URL.
        /// </summary>
        [Input("scepUrl")]
        public Input<string>? ScepUrl { get; set; }

        /// <summary>
        /// Certificate source type.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// Source IP address for communications to the SCEP server.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Certificate Signing Request State.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public LocalArgs()
        {
        }
        public static new LocalArgs Empty => new LocalArgs();
    }

    public sealed class LocalState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL for the ACME CA server (Let's Encrypt is the default provider).
        /// </summary>
        [Input("acmeCaUrl")]
        public Input<string>? AcmeCaUrl { get; set; }

        /// <summary>
        /// A valid domain that resolves to this Fortigate.
        /// </summary>
        [Input("acmeDomain")]
        public Input<string>? AcmeDomain { get; set; }

        /// <summary>
        /// Contact email address that is required by some CAs like LetsEncrypt.
        /// </summary>
        [Input("acmeEmail")]
        public Input<string>? AcmeEmail { get; set; }

        /// <summary>
        /// Beginning of the renewal window (in days before certificate expiration, 30 by default).
        /// </summary>
        [Input("acmeRenewWindow")]
        public Input<int>? AcmeRenewWindow { get; set; }

        /// <summary>
        /// Length of the RSA private key of the generated cert (Minimum 2048 bits).
        /// </summary>
        [Input("acmeRsaKeySize")]
        public Input<int>? AcmeRsaKeySize { get; set; }

        /// <summary>
        /// Number of days to wait before expiry of an updated local certificate is requested (0 = disabled).
        /// </summary>
        [Input("autoRegenerateDays")]
        public Input<int>? AutoRegenerateDays { get; set; }

        /// <summary>
        /// Number of days to wait before an expiry warning message is generated (0 = disabled).
        /// </summary>
        [Input("autoRegenerateDaysWarning")]
        public Input<int>? AutoRegenerateDaysWarning { get; set; }

        /// <summary>
        /// CA identifier of the CA server for signing via SCEP.
        /// </summary>
        [Input("caIdentifier")]
        public Input<string>? CaIdentifier { get; set; }

        [Input("certificate")]
        private Input<string>? _certificate;

        /// <summary>
        /// PEM format certificate.
        /// </summary>
        public Input<string>? Certificate
        {
            get => _certificate;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _certificate = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Path location inside CMP server.
        /// </summary>
        [Input("cmpPath")]
        public Input<string>? CmpPath { get; set; }

        /// <summary>
        /// CMP auto-regeneration method. Valid values: `keyupate`, `renewal`.
        /// </summary>
        [Input("cmpRegenerationMethod")]
        public Input<string>? CmpRegenerationMethod { get; set; }

        /// <summary>
        /// 'ADDRESS:PORT' for CMP server.
        /// </summary>
        [Input("cmpServer")]
        public Input<string>? CmpServer { get; set; }

        /// <summary>
        /// CMP server certificate.
        /// </summary>
        [Input("cmpServerCert")]
        public Input<string>? CmpServerCert { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Certificate Signing Request.
        /// </summary>
        [Input("csr")]
        public Input<string>? Csr { get; set; }

        /// <summary>
        /// Certificate enrollment protocol.
        /// </summary>
        [Input("enrollProtocol")]
        public Input<string>? EnrollProtocol { get; set; }

        /// <summary>
        /// Local ID the FortiGate uses for authentication as a VPN client.
        /// </summary>
        [Input("ikeLocalid")]
        public Input<string>? IkeLocalid { get; set; }

        /// <summary>
        /// IKE local ID type. Valid values: `asn1dn`, `fqdn`.
        /// </summary>
        [Input("ikeLocalidType")]
        public Input<string>? IkeLocalidType { get; set; }

        /// <summary>
        /// Time at which certificate was last updated.
        /// </summary>
        [Input("lastUpdated")]
        public Input<int>? LastUpdated { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name encoding method for auto-regeneration. Valid values: `printable`, `utf8`.
        /// </summary>
        [Input("nameEncoding")]
        public Input<string>? NameEncoding { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password as a PEM file.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateKey")]
        private Input<string>? _privateKey;

        /// <summary>
        /// PEM format key, encrypted with a password.
        /// </summary>
        public Input<string>? PrivateKey
        {
            get => _privateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Enable/disable retention of private key during SCEP renewal (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("privateKeyRetain")]
        public Input<string>? PrivateKeyRetain { get; set; }

        /// <summary>
        /// Either a global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
        /// </summary>
        [Input("range")]
        public Input<string>? Range { get; set; }

        [Input("scepPassword")]
        private Input<string>? _scepPassword;

        /// <summary>
        /// SCEP server challenge password for auto-regeneration.
        /// </summary>
        public Input<string>? ScepPassword
        {
            get => _scepPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _scepPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// SCEP server URL.
        /// </summary>
        [Input("scepUrl")]
        public Input<string>? ScepUrl { get; set; }

        /// <summary>
        /// Certificate source type.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// Source IP address for communications to the SCEP server.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Certificate Signing Request State.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public LocalState()
        {
        }
        public static new LocalState Empty => new LocalState();
    }
}
