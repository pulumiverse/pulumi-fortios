// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall
{
    /// <summary>
    /// Configure SSL/SSH protocol options.
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
    ///     var t1 = new Fortios.Firewall.Sslsshprofile("t1", new()
    ///     {
    ///         Ftps = new Fortios.Firewall.Inputs.SslsshprofileFtpsArgs
    ///         {
    ///             Ports = "990",
    ///         },
    ///         Https = new Fortios.Firewall.Inputs.SslsshprofileHttpsArgs
    ///         {
    ///             Ports = "443 127 422 392",
    ///         },
    ///         Imaps = new Fortios.Firewall.Inputs.SslsshprofileImapsArgs
    ///         {
    ///             Ports = "993 1123",
    ///         },
    ///         Pop3s = new Fortios.Firewall.Inputs.SslsshprofilePop3sArgs
    ///         {
    ///             Ports = "995",
    ///         },
    ///         Smtps = new Fortios.Firewall.Inputs.SslsshprofileSmtpsArgs
    ///         {
    ///             Ports = "465",
    ///         },
    ///         Ssl = new Fortios.Firewall.Inputs.SslsshprofileSslArgs
    ///         {
    ///             InspectAll = "disable",
    ///         },
    ///     });
    /// 
    ///     var t2 = new Fortios.Firewall.Sslsshprofile("t2", new()
    ///     {
    ///         Https = new Fortios.Firewall.Inputs.SslsshprofileHttpsArgs
    ///         {
    ///             Ports = "443",
    ///         },
    ///         Ssl = new Fortios.Firewall.Inputs.SslsshprofileSslArgs
    ///         {
    ///             InspectAll = "deep-inspection",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Firewall SslSshProfile can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/sslsshprofile:Sslsshprofile labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/sslsshprofile:Sslsshprofile labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/sslsshprofile:Sslsshprofile")]
    public partial class Sslsshprofile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable exempting servers by FortiGuard allowlist. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("allowlist")]
        public Output<string> Allowlist { get; private set; } = null!;

        /// <summary>
        /// Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blacklist. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("blockBlacklistedCertificates")]
        public Output<string> BlockBlacklistedCertificates { get; private set; } = null!;

        /// <summary>
        /// Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blocklist. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("blockBlocklistedCertificates")]
        public Output<string> BlockBlocklistedCertificates { get; private set; } = null!;

        /// <summary>
        /// CA certificate used by SSL Inspection.
        /// </summary>
        [Output("caname")]
        public Output<string> Caname { get; private set; } = null!;

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Configure DNS over TLS options. The structure of `dot` block is documented below.
        /// </summary>
        [Output("dot")]
        public Output<Outputs.SslsshprofileDot> Dot { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Configure FTPS options. The structure of `ftps` block is documented below.
        /// </summary>
        [Output("ftps")]
        public Output<Outputs.SslsshprofileFtps> Ftps { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Configure HTTPS options. The structure of `https` block is documented below.
        /// </summary>
        [Output("https")]
        public Output<Outputs.SslsshprofileHttps> Https { get; private set; } = null!;

        /// <summary>
        /// Configure IMAPS options. The structure of `imaps` block is documented below.
        /// </summary>
        [Output("imaps")]
        public Output<Outputs.SslsshprofileImaps> Imaps { get; private set; } = null!;

        /// <summary>
        /// Enable/disable inspection of MAPI over HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("mapiOverHttps")]
        public Output<string> MapiOverHttps { get; private set; } = null!;

        /// <summary>
        /// Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Configure POP3S options. The structure of `pop3s` block is documented below.
        /// </summary>
        [Output("pop3s")]
        public Output<Outputs.SslsshprofilePop3s> Pop3s { get; private set; } = null!;

        /// <summary>
        /// Enable/disable inspection of RPC over HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("rpcOverHttps")]
        public Output<string> RpcOverHttps { get; private set; } = null!;

        /// <summary>
        /// Certificate used by SSL Inspection to replace server certificate.
        /// </summary>
        [Output("serverCert")]
        public Output<string> ServerCert { get; private set; } = null!;

        /// <summary>
        /// Re-sign or replace the server's certificate. Valid values: `re-sign`, `replace`.
        /// </summary>
        [Output("serverCertMode")]
        public Output<string> ServerCertMode { get; private set; } = null!;

        /// <summary>
        /// Configure SMTPS options. The structure of `smtps` block is documented below.
        /// </summary>
        [Output("smtps")]
        public Output<Outputs.SslsshprofileSmtps> Smtps { get; private set; } = null!;

        /// <summary>
        /// Configure SSH options. The structure of `ssh` block is documented below.
        /// </summary>
        [Output("ssh")]
        public Output<Outputs.SslsshprofileSsh> Ssh { get; private set; } = null!;

        /// <summary>
        /// Configure SSL options. The structure of `ssl` block is documented below.
        /// </summary>
        [Output("ssl")]
        public Output<Outputs.SslsshprofileSsl> Ssl { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging SSL anomalies. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("sslAnomaliesLog")]
        public Output<string> SslAnomaliesLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging SSL anomalies. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("sslAnomalyLog")]
        public Output<string> SslAnomalyLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable IP based URL rating. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("sslExemptionIpRating")]
        public Output<string> SslExemptionIpRating { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging SSL exemptions. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("sslExemptionLog")]
        public Output<string> SslExemptionLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging SSL exemptions. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("sslExemptionsLog")]
        public Output<string> SslExemptionsLog { get; private set; } = null!;

        /// <summary>
        /// Servers to exempt from SSL inspection. The structure of `ssl_exempt` block is documented below.
        /// </summary>
        [Output("sslExempts")]
        public Output<ImmutableArray<Outputs.SslsshprofileSslExempt>> SslExempts { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging of TLS handshakes. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("sslHandshakeLog")]
        public Output<string> SslHandshakeLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging SSL negotiation. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("sslNegotiationLog")]
        public Output<string> SslNegotiationLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging of server certificate information. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("sslServerCertLog")]
        public Output<string> SslServerCertLog { get; private set; } = null!;

        /// <summary>
        /// SSL servers. The structure of `ssl_server` block is documented below.
        /// </summary>
        [Output("sslServers")]
        public Output<ImmutableArray<Outputs.SslsshprofileSslServer>> SslServers { get; private set; } = null!;

        /// <summary>
        /// Configure ALPN option. Valid values: `http1-1`, `http2`, `all`, `none`.
        /// </summary>
        [Output("supportedAlpn")]
        public Output<string> SupportedAlpn { get; private set; } = null!;

        /// <summary>
        /// Untrusted CA certificate used by SSL Inspection.
        /// </summary>
        [Output("untrustedCaname")]
        public Output<string> UntrustedCaname { get; private set; } = null!;

        /// <summary>
        /// Enable/disable the use of SSL server table for SSL offloading. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("useSslServer")]
        public Output<string> UseSslServer { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Enable/disable exempting servers by FortiGuard whitelist. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("whitelist")]
        public Output<string> Whitelist { get; private set; } = null!;


        /// <summary>
        /// Create a Sslsshprofile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Sslsshprofile(string name, SslsshprofileArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/sslsshprofile:Sslsshprofile", name, args ?? new SslsshprofileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Sslsshprofile(string name, Input<string> id, SslsshprofileState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/sslsshprofile:Sslsshprofile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Sslsshprofile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Sslsshprofile Get(string name, Input<string> id, SslsshprofileState? state = null, CustomResourceOptions? options = null)
        {
            return new Sslsshprofile(name, id, state, options);
        }
    }

    public sealed class SslsshprofileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable exempting servers by FortiGuard allowlist. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("allowlist")]
        public Input<string>? Allowlist { get; set; }

        /// <summary>
        /// Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blacklist. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("blockBlacklistedCertificates")]
        public Input<string>? BlockBlacklistedCertificates { get; set; }

        /// <summary>
        /// Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blocklist. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("blockBlocklistedCertificates")]
        public Input<string>? BlockBlocklistedCertificates { get; set; }

        /// <summary>
        /// CA certificate used by SSL Inspection.
        /// </summary>
        [Input("caname")]
        public Input<string>? Caname { get; set; }

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Configure DNS over TLS options. The structure of `dot` block is documented below.
        /// </summary>
        [Input("dot")]
        public Input<Inputs.SslsshprofileDotArgs>? Dot { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Configure FTPS options. The structure of `ftps` block is documented below.
        /// </summary>
        [Input("ftps")]
        public Input<Inputs.SslsshprofileFtpsArgs>? Ftps { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Configure HTTPS options. The structure of `https` block is documented below.
        /// </summary>
        [Input("https")]
        public Input<Inputs.SslsshprofileHttpsArgs>? Https { get; set; }

        /// <summary>
        /// Configure IMAPS options. The structure of `imaps` block is documented below.
        /// </summary>
        [Input("imaps")]
        public Input<Inputs.SslsshprofileImapsArgs>? Imaps { get; set; }

        /// <summary>
        /// Enable/disable inspection of MAPI over HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("mapiOverHttps")]
        public Input<string>? MapiOverHttps { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configure POP3S options. The structure of `pop3s` block is documented below.
        /// </summary>
        [Input("pop3s")]
        public Input<Inputs.SslsshprofilePop3sArgs>? Pop3s { get; set; }

        /// <summary>
        /// Enable/disable inspection of RPC over HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("rpcOverHttps")]
        public Input<string>? RpcOverHttps { get; set; }

        /// <summary>
        /// Certificate used by SSL Inspection to replace server certificate.
        /// </summary>
        [Input("serverCert")]
        public Input<string>? ServerCert { get; set; }

        /// <summary>
        /// Re-sign or replace the server's certificate. Valid values: `re-sign`, `replace`.
        /// </summary>
        [Input("serverCertMode")]
        public Input<string>? ServerCertMode { get; set; }

        /// <summary>
        /// Configure SMTPS options. The structure of `smtps` block is documented below.
        /// </summary>
        [Input("smtps")]
        public Input<Inputs.SslsshprofileSmtpsArgs>? Smtps { get; set; }

        /// <summary>
        /// Configure SSH options. The structure of `ssh` block is documented below.
        /// </summary>
        [Input("ssh")]
        public Input<Inputs.SslsshprofileSshArgs>? Ssh { get; set; }

        /// <summary>
        /// Configure SSL options. The structure of `ssl` block is documented below.
        /// </summary>
        [Input("ssl")]
        public Input<Inputs.SslsshprofileSslArgs>? Ssl { get; set; }

        /// <summary>
        /// Enable/disable logging SSL anomalies. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sslAnomaliesLog")]
        public Input<string>? SslAnomaliesLog { get; set; }

        /// <summary>
        /// Enable/disable logging SSL anomalies. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sslAnomalyLog")]
        public Input<string>? SslAnomalyLog { get; set; }

        /// <summary>
        /// Enable/disable IP based URL rating. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sslExemptionIpRating")]
        public Input<string>? SslExemptionIpRating { get; set; }

        /// <summary>
        /// Enable/disable logging SSL exemptions. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sslExemptionLog")]
        public Input<string>? SslExemptionLog { get; set; }

        /// <summary>
        /// Enable/disable logging SSL exemptions. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sslExemptionsLog")]
        public Input<string>? SslExemptionsLog { get; set; }

        [Input("sslExempts")]
        private InputList<Inputs.SslsshprofileSslExemptArgs>? _sslExempts;

        /// <summary>
        /// Servers to exempt from SSL inspection. The structure of `ssl_exempt` block is documented below.
        /// </summary>
        public InputList<Inputs.SslsshprofileSslExemptArgs> SslExempts
        {
            get => _sslExempts ?? (_sslExempts = new InputList<Inputs.SslsshprofileSslExemptArgs>());
            set => _sslExempts = value;
        }

        /// <summary>
        /// Enable/disable logging of TLS handshakes. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sslHandshakeLog")]
        public Input<string>? SslHandshakeLog { get; set; }

        /// <summary>
        /// Enable/disable logging SSL negotiation. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sslNegotiationLog")]
        public Input<string>? SslNegotiationLog { get; set; }

        /// <summary>
        /// Enable/disable logging of server certificate information. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sslServerCertLog")]
        public Input<string>? SslServerCertLog { get; set; }

        [Input("sslServers")]
        private InputList<Inputs.SslsshprofileSslServerArgs>? _sslServers;

        /// <summary>
        /// SSL servers. The structure of `ssl_server` block is documented below.
        /// </summary>
        public InputList<Inputs.SslsshprofileSslServerArgs> SslServers
        {
            get => _sslServers ?? (_sslServers = new InputList<Inputs.SslsshprofileSslServerArgs>());
            set => _sslServers = value;
        }

        /// <summary>
        /// Configure ALPN option. Valid values: `http1-1`, `http2`, `all`, `none`.
        /// </summary>
        [Input("supportedAlpn")]
        public Input<string>? SupportedAlpn { get; set; }

        /// <summary>
        /// Untrusted CA certificate used by SSL Inspection.
        /// </summary>
        [Input("untrustedCaname")]
        public Input<string>? UntrustedCaname { get; set; }

        /// <summary>
        /// Enable/disable the use of SSL server table for SSL offloading. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("useSslServer")]
        public Input<string>? UseSslServer { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable exempting servers by FortiGuard whitelist. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("whitelist")]
        public Input<string>? Whitelist { get; set; }

        public SslsshprofileArgs()
        {
        }
        public static new SslsshprofileArgs Empty => new SslsshprofileArgs();
    }

    public sealed class SslsshprofileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable exempting servers by FortiGuard allowlist. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("allowlist")]
        public Input<string>? Allowlist { get; set; }

        /// <summary>
        /// Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blacklist. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("blockBlacklistedCertificates")]
        public Input<string>? BlockBlacklistedCertificates { get; set; }

        /// <summary>
        /// Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blocklist. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("blockBlocklistedCertificates")]
        public Input<string>? BlockBlocklistedCertificates { get; set; }

        /// <summary>
        /// CA certificate used by SSL Inspection.
        /// </summary>
        [Input("caname")]
        public Input<string>? Caname { get; set; }

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Configure DNS over TLS options. The structure of `dot` block is documented below.
        /// </summary>
        [Input("dot")]
        public Input<Inputs.SslsshprofileDotGetArgs>? Dot { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Configure FTPS options. The structure of `ftps` block is documented below.
        /// </summary>
        [Input("ftps")]
        public Input<Inputs.SslsshprofileFtpsGetArgs>? Ftps { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Configure HTTPS options. The structure of `https` block is documented below.
        /// </summary>
        [Input("https")]
        public Input<Inputs.SslsshprofileHttpsGetArgs>? Https { get; set; }

        /// <summary>
        /// Configure IMAPS options. The structure of `imaps` block is documented below.
        /// </summary>
        [Input("imaps")]
        public Input<Inputs.SslsshprofileImapsGetArgs>? Imaps { get; set; }

        /// <summary>
        /// Enable/disable inspection of MAPI over HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("mapiOverHttps")]
        public Input<string>? MapiOverHttps { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configure POP3S options. The structure of `pop3s` block is documented below.
        /// </summary>
        [Input("pop3s")]
        public Input<Inputs.SslsshprofilePop3sGetArgs>? Pop3s { get; set; }

        /// <summary>
        /// Enable/disable inspection of RPC over HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("rpcOverHttps")]
        public Input<string>? RpcOverHttps { get; set; }

        /// <summary>
        /// Certificate used by SSL Inspection to replace server certificate.
        /// </summary>
        [Input("serverCert")]
        public Input<string>? ServerCert { get; set; }

        /// <summary>
        /// Re-sign or replace the server's certificate. Valid values: `re-sign`, `replace`.
        /// </summary>
        [Input("serverCertMode")]
        public Input<string>? ServerCertMode { get; set; }

        /// <summary>
        /// Configure SMTPS options. The structure of `smtps` block is documented below.
        /// </summary>
        [Input("smtps")]
        public Input<Inputs.SslsshprofileSmtpsGetArgs>? Smtps { get; set; }

        /// <summary>
        /// Configure SSH options. The structure of `ssh` block is documented below.
        /// </summary>
        [Input("ssh")]
        public Input<Inputs.SslsshprofileSshGetArgs>? Ssh { get; set; }

        /// <summary>
        /// Configure SSL options. The structure of `ssl` block is documented below.
        /// </summary>
        [Input("ssl")]
        public Input<Inputs.SslsshprofileSslGetArgs>? Ssl { get; set; }

        /// <summary>
        /// Enable/disable logging SSL anomalies. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sslAnomaliesLog")]
        public Input<string>? SslAnomaliesLog { get; set; }

        /// <summary>
        /// Enable/disable logging SSL anomalies. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sslAnomalyLog")]
        public Input<string>? SslAnomalyLog { get; set; }

        /// <summary>
        /// Enable/disable IP based URL rating. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sslExemptionIpRating")]
        public Input<string>? SslExemptionIpRating { get; set; }

        /// <summary>
        /// Enable/disable logging SSL exemptions. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sslExemptionLog")]
        public Input<string>? SslExemptionLog { get; set; }

        /// <summary>
        /// Enable/disable logging SSL exemptions. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sslExemptionsLog")]
        public Input<string>? SslExemptionsLog { get; set; }

        [Input("sslExempts")]
        private InputList<Inputs.SslsshprofileSslExemptGetArgs>? _sslExempts;

        /// <summary>
        /// Servers to exempt from SSL inspection. The structure of `ssl_exempt` block is documented below.
        /// </summary>
        public InputList<Inputs.SslsshprofileSslExemptGetArgs> SslExempts
        {
            get => _sslExempts ?? (_sslExempts = new InputList<Inputs.SslsshprofileSslExemptGetArgs>());
            set => _sslExempts = value;
        }

        /// <summary>
        /// Enable/disable logging of TLS handshakes. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sslHandshakeLog")]
        public Input<string>? SslHandshakeLog { get; set; }

        /// <summary>
        /// Enable/disable logging SSL negotiation. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sslNegotiationLog")]
        public Input<string>? SslNegotiationLog { get; set; }

        /// <summary>
        /// Enable/disable logging of server certificate information. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sslServerCertLog")]
        public Input<string>? SslServerCertLog { get; set; }

        [Input("sslServers")]
        private InputList<Inputs.SslsshprofileSslServerGetArgs>? _sslServers;

        /// <summary>
        /// SSL servers. The structure of `ssl_server` block is documented below.
        /// </summary>
        public InputList<Inputs.SslsshprofileSslServerGetArgs> SslServers
        {
            get => _sslServers ?? (_sslServers = new InputList<Inputs.SslsshprofileSslServerGetArgs>());
            set => _sslServers = value;
        }

        /// <summary>
        /// Configure ALPN option. Valid values: `http1-1`, `http2`, `all`, `none`.
        /// </summary>
        [Input("supportedAlpn")]
        public Input<string>? SupportedAlpn { get; set; }

        /// <summary>
        /// Untrusted CA certificate used by SSL Inspection.
        /// </summary>
        [Input("untrustedCaname")]
        public Input<string>? UntrustedCaname { get; set; }

        /// <summary>
        /// Enable/disable the use of SSL server table for SSL offloading. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("useSslServer")]
        public Input<string>? UseSslServer { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable exempting servers by FortiGuard whitelist. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("whitelist")]
        public Input<string>? Whitelist { get; set; }

        public SslsshprofileState()
        {
        }
        public static new SslsshprofileState Empty => new SslsshprofileState();
    }
}
