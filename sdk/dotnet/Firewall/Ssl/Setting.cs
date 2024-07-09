// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Ssl
{
    /// <summary>
    /// SSL proxy settings.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.Firewall.Ssl.Setting("trname", new()
    ///     {
    ///         AbbreviateHandshake = "enable",
    ///         CertCacheCapacity = 200,
    ///         CertCacheTimeout = 10,
    ///         KxpQueueThreshold = 16,
    ///         NoMatchingCipherAction = "bypass",
    ///         ProxyConnectTimeout = 30,
    ///         SessionCacheCapacity = 500,
    ///         SessionCacheTimeout = 20,
    ///         SslDhBits = "2048",
    ///         SslQueueThreshold = 32,
    ///         SslSendEmptyFrags = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// FirewallSsl Setting can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/ssl/setting:Setting labelname FirewallSslSetting
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/ssl/setting:Setting labelname FirewallSslSetting
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/ssl/setting:Setting")]
    public partial class Setting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable use of SSL abbreviated handshake. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("abbreviateHandshake")]
        public Output<string> AbbreviateHandshake { get; private set; } = null!;

        /// <summary>
        /// Maximum capacity of the host certificate cache (0 - 500, default = 200).
        /// </summary>
        [Output("certCacheCapacity")]
        public Output<int> CertCacheCapacity { get; private set; } = null!;

        /// <summary>
        /// Time limit to keep certificate cache (1 - 120 min, default = 10).
        /// </summary>
        [Output("certCacheTimeout")]
        public Output<int> CertCacheTimeout { get; private set; } = null!;

        /// <summary>
        /// Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
        /// </summary>
        [Output("kxpQueueThreshold")]
        public Output<int> KxpQueueThreshold { get; private set; } = null!;

        /// <summary>
        /// Bypass or drop the connection when no matching cipher is found. Valid values: `bypass`, `drop`.
        /// </summary>
        [Output("noMatchingCipherAction")]
        public Output<string> NoMatchingCipherAction { get; private set; } = null!;

        /// <summary>
        /// Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
        /// </summary>
        [Output("proxyConnectTimeout")]
        public Output<int> ProxyConnectTimeout { get; private set; } = null!;

        /// <summary>
        /// Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
        /// </summary>
        [Output("sessionCacheCapacity")]
        public Output<int> SessionCacheCapacity { get; private set; } = null!;

        /// <summary>
        /// Time limit to keep SSL session state (1 - 60 min, default = 20).
        /// </summary>
        [Output("sessionCacheTimeout")]
        public Output<int> SessionCacheTimeout { get; private set; } = null!;

        /// <summary>
        /// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        /// </summary>
        [Output("sslDhBits")]
        public Output<string> SslDhBits { get; private set; } = null!;

        /// <summary>
        /// Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
        /// </summary>
        [Output("sslQueueThreshold")]
        public Output<int> SslQueueThreshold { get; private set; } = null!;

        /// <summary>
        /// Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("sslSendEmptyFrags")]
        public Output<string> SslSendEmptyFrags { get; private set; } = null!;

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
            : base("fortios:firewall/ssl/setting:Setting", name, args ?? new SettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Setting(string name, Input<string> id, SettingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/ssl/setting:Setting", name, state, MakeResourceOptions(options, id))
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
        /// Enable/disable use of SSL abbreviated handshake. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("abbreviateHandshake")]
        public Input<string>? AbbreviateHandshake { get; set; }

        /// <summary>
        /// Maximum capacity of the host certificate cache (0 - 500, default = 200).
        /// </summary>
        [Input("certCacheCapacity", required: true)]
        public Input<int> CertCacheCapacity { get; set; } = null!;

        /// <summary>
        /// Time limit to keep certificate cache (1 - 120 min, default = 10).
        /// </summary>
        [Input("certCacheTimeout", required: true)]
        public Input<int> CertCacheTimeout { get; set; } = null!;

        /// <summary>
        /// Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
        /// </summary>
        [Input("kxpQueueThreshold")]
        public Input<int>? KxpQueueThreshold { get; set; }

        /// <summary>
        /// Bypass or drop the connection when no matching cipher is found. Valid values: `bypass`, `drop`.
        /// </summary>
        [Input("noMatchingCipherAction", required: true)]
        public Input<string> NoMatchingCipherAction { get; set; } = null!;

        /// <summary>
        /// Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
        /// </summary>
        [Input("proxyConnectTimeout", required: true)]
        public Input<int> ProxyConnectTimeout { get; set; } = null!;

        /// <summary>
        /// Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
        /// </summary>
        [Input("sessionCacheCapacity", required: true)]
        public Input<int> SessionCacheCapacity { get; set; } = null!;

        /// <summary>
        /// Time limit to keep SSL session state (1 - 60 min, default = 20).
        /// </summary>
        [Input("sessionCacheTimeout", required: true)]
        public Input<int> SessionCacheTimeout { get; set; } = null!;

        /// <summary>
        /// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        /// </summary>
        [Input("sslDhBits", required: true)]
        public Input<string> SslDhBits { get; set; } = null!;

        /// <summary>
        /// Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
        /// </summary>
        [Input("sslQueueThreshold")]
        public Input<int>? SslQueueThreshold { get; set; }

        /// <summary>
        /// Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sslSendEmptyFrags", required: true)]
        public Input<string> SslSendEmptyFrags { get; set; } = null!;

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
        /// Enable/disable use of SSL abbreviated handshake. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("abbreviateHandshake")]
        public Input<string>? AbbreviateHandshake { get; set; }

        /// <summary>
        /// Maximum capacity of the host certificate cache (0 - 500, default = 200).
        /// </summary>
        [Input("certCacheCapacity")]
        public Input<int>? CertCacheCapacity { get; set; }

        /// <summary>
        /// Time limit to keep certificate cache (1 - 120 min, default = 10).
        /// </summary>
        [Input("certCacheTimeout")]
        public Input<int>? CertCacheTimeout { get; set; }

        /// <summary>
        /// Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
        /// </summary>
        [Input("kxpQueueThreshold")]
        public Input<int>? KxpQueueThreshold { get; set; }

        /// <summary>
        /// Bypass or drop the connection when no matching cipher is found. Valid values: `bypass`, `drop`.
        /// </summary>
        [Input("noMatchingCipherAction")]
        public Input<string>? NoMatchingCipherAction { get; set; }

        /// <summary>
        /// Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
        /// </summary>
        [Input("proxyConnectTimeout")]
        public Input<int>? ProxyConnectTimeout { get; set; }

        /// <summary>
        /// Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
        /// </summary>
        [Input("sessionCacheCapacity")]
        public Input<int>? SessionCacheCapacity { get; set; }

        /// <summary>
        /// Time limit to keep SSL session state (1 - 60 min, default = 20).
        /// </summary>
        [Input("sessionCacheTimeout")]
        public Input<int>? SessionCacheTimeout { get; set; }

        /// <summary>
        /// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        /// </summary>
        [Input("sslDhBits")]
        public Input<string>? SslDhBits { get; set; }

        /// <summary>
        /// Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
        /// </summary>
        [Input("sslQueueThreshold")]
        public Input<int>? SslQueueThreshold { get; set; }

        /// <summary>
        /// Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sslSendEmptyFrags")]
        public Input<string>? SslSendEmptyFrags { get; set; }

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
