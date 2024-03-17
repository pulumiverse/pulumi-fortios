// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Authentication
{
    /// <summary>
    /// Configure authentication setting.
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
    ///     var trname = new Fortios.Authentication.Setting("trname", new()
    ///     {
    ///         AuthHttps = "enable",
    ///         CaptivePortalIp = "0.0.0.0",
    ///         CaptivePortalIp6 = "::",
    ///         CaptivePortalPort = 7830,
    ///         CaptivePortalSslPort = 7831,
    ///         CaptivePortalType = "fqdn",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Authentication Setting can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:authentication/setting:Setting labelname AuthenticationSetting
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:authentication/setting:Setting labelname AuthenticationSetting
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:authentication/setting:Setting")]
    public partial class Setting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Active authentication method (scheme name).
        /// </summary>
        [Output("activeAuthScheme")]
        public Output<string> ActiveAuthScheme { get; private set; } = null!;

        /// <summary>
        /// Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("authHttps")]
        public Output<string> AuthHttps { get; private set; } = null!;

        /// <summary>
        /// Captive portal host name.
        /// </summary>
        [Output("captivePortal")]
        public Output<string> CaptivePortal { get; private set; } = null!;

        /// <summary>
        /// IPv6 captive portal host name.
        /// </summary>
        [Output("captivePortal6")]
        public Output<string> CaptivePortal6 { get; private set; } = null!;

        /// <summary>
        /// Captive portal IP address.
        /// </summary>
        [Output("captivePortalIp")]
        public Output<string> CaptivePortalIp { get; private set; } = null!;

        /// <summary>
        /// Captive portal IPv6 address.
        /// </summary>
        [Output("captivePortalIp6")]
        public Output<string> CaptivePortalIp6 { get; private set; } = null!;

        /// <summary>
        /// Captive portal port number (1 - 65535, default = 7830).
        /// </summary>
        [Output("captivePortalPort")]
        public Output<int> CaptivePortalPort { get; private set; } = null!;

        /// <summary>
        /// Captive portal SSL port number (1 - 65535, default = 7831).
        /// </summary>
        [Output("captivePortalSslPort")]
        public Output<int> CaptivePortalSslPort { get; private set; } = null!;

        /// <summary>
        /// Captive portal type. Valid values: `fqdn`, `ip`.
        /// </summary>
        [Output("captivePortalType")]
        public Output<string> CaptivePortalType { get; private set; } = null!;

        /// <summary>
        /// Enable/disable redirecting certificate authentication to HTTPS portal. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("certAuth")]
        public Output<string> CertAuth { get; private set; } = null!;

        /// <summary>
        /// Certificate captive portal host name.
        /// </summary>
        [Output("certCaptivePortal")]
        public Output<string> CertCaptivePortal { get; private set; } = null!;

        /// <summary>
        /// Certificate captive portal IP address.
        /// </summary>
        [Output("certCaptivePortalIp")]
        public Output<string> CertCaptivePortalIp { get; private set; } = null!;

        /// <summary>
        /// Certificate captive portal port number (1 - 65535, default = 7832).
        /// </summary>
        [Output("certCaptivePortalPort")]
        public Output<int> CertCaptivePortalPort { get; private set; } = null!;

        /// <summary>
        /// Persistent web portal cookie maximum age in minutes (30 - 10080 (1 week), default = 480 (8 hours)).
        /// </summary>
        [Output("cookieMaxAge")]
        public Output<int> CookieMaxAge { get; private set; } = null!;

        /// <summary>
        /// Refresh rate divider of persistent web portal cookie (default = 2). Refresh value = cookie-max-age/cookie-refresh-div.
        /// </summary>
        [Output("cookieRefreshDiv")]
        public Output<int> CookieRefreshDiv { get; private set; } = null!;

        /// <summary>
        /// Address range for the IP based device query. The structure of `dev_range` block is documented below.
        /// </summary>
        [Output("devRanges")]
        public Output<ImmutableArray<Outputs.SettingDevRange>> DevRanges { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Enable/disable persistent cookie on IP based web portal authentication (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ipAuthCookie")]
        public Output<string> IpAuthCookie { get; private set; } = null!;

        /// <summary>
        /// Enable/disable persistent cookie on web portal authentication (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("persistentCookie")]
        public Output<string> PersistentCookie { get; private set; } = null!;

        /// <summary>
        /// Single-Sign-On authentication method (scheme name).
        /// </summary>
        [Output("ssoAuthScheme")]
        public Output<string> SsoAuthScheme { get; private set; } = null!;

        /// <summary>
        /// Time of the last update.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// CA certificate used for client certificate verification. The structure of `user_cert_ca` block is documented below.
        /// </summary>
        [Output("userCertCas")]
        public Output<ImmutableArray<Outputs.SettingUserCertCa>> UserCertCas { get; private set; } = null!;

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
        public Setting(string name, SettingArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:authentication/setting:Setting", name, args ?? new SettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Setting(string name, Input<string> id, SettingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:authentication/setting:Setting", name, state, MakeResourceOptions(options, id))
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
        /// Active authentication method (scheme name).
        /// </summary>
        [Input("activeAuthScheme")]
        public Input<string>? ActiveAuthScheme { get; set; }

        /// <summary>
        /// Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authHttps")]
        public Input<string>? AuthHttps { get; set; }

        /// <summary>
        /// Captive portal host name.
        /// </summary>
        [Input("captivePortal")]
        public Input<string>? CaptivePortal { get; set; }

        /// <summary>
        /// IPv6 captive portal host name.
        /// </summary>
        [Input("captivePortal6")]
        public Input<string>? CaptivePortal6 { get; set; }

        /// <summary>
        /// Captive portal IP address.
        /// </summary>
        [Input("captivePortalIp")]
        public Input<string>? CaptivePortalIp { get; set; }

        /// <summary>
        /// Captive portal IPv6 address.
        /// </summary>
        [Input("captivePortalIp6")]
        public Input<string>? CaptivePortalIp6 { get; set; }

        /// <summary>
        /// Captive portal port number (1 - 65535, default = 7830).
        /// </summary>
        [Input("captivePortalPort")]
        public Input<int>? CaptivePortalPort { get; set; }

        /// <summary>
        /// Captive portal SSL port number (1 - 65535, default = 7831).
        /// </summary>
        [Input("captivePortalSslPort")]
        public Input<int>? CaptivePortalSslPort { get; set; }

        /// <summary>
        /// Captive portal type. Valid values: `fqdn`, `ip`.
        /// </summary>
        [Input("captivePortalType")]
        public Input<string>? CaptivePortalType { get; set; }

        /// <summary>
        /// Enable/disable redirecting certificate authentication to HTTPS portal. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("certAuth")]
        public Input<string>? CertAuth { get; set; }

        /// <summary>
        /// Certificate captive portal host name.
        /// </summary>
        [Input("certCaptivePortal")]
        public Input<string>? CertCaptivePortal { get; set; }

        /// <summary>
        /// Certificate captive portal IP address.
        /// </summary>
        [Input("certCaptivePortalIp")]
        public Input<string>? CertCaptivePortalIp { get; set; }

        /// <summary>
        /// Certificate captive portal port number (1 - 65535, default = 7832).
        /// </summary>
        [Input("certCaptivePortalPort")]
        public Input<int>? CertCaptivePortalPort { get; set; }

        /// <summary>
        /// Persistent web portal cookie maximum age in minutes (30 - 10080 (1 week), default = 480 (8 hours)).
        /// </summary>
        [Input("cookieMaxAge")]
        public Input<int>? CookieMaxAge { get; set; }

        /// <summary>
        /// Refresh rate divider of persistent web portal cookie (default = 2). Refresh value = cookie-max-age/cookie-refresh-div.
        /// </summary>
        [Input("cookieRefreshDiv")]
        public Input<int>? CookieRefreshDiv { get; set; }

        [Input("devRanges")]
        private InputList<Inputs.SettingDevRangeArgs>? _devRanges;

        /// <summary>
        /// Address range for the IP based device query. The structure of `dev_range` block is documented below.
        /// </summary>
        public InputList<Inputs.SettingDevRangeArgs> DevRanges
        {
            get => _devRanges ?? (_devRanges = new InputList<Inputs.SettingDevRangeArgs>());
            set => _devRanges = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable persistent cookie on IP based web portal authentication (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipAuthCookie")]
        public Input<string>? IpAuthCookie { get; set; }

        /// <summary>
        /// Enable/disable persistent cookie on web portal authentication (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("persistentCookie")]
        public Input<string>? PersistentCookie { get; set; }

        /// <summary>
        /// Single-Sign-On authentication method (scheme name).
        /// </summary>
        [Input("ssoAuthScheme")]
        public Input<string>? SsoAuthScheme { get; set; }

        /// <summary>
        /// Time of the last update.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        [Input("userCertCas")]
        private InputList<Inputs.SettingUserCertCaArgs>? _userCertCas;

        /// <summary>
        /// CA certificate used for client certificate verification. The structure of `user_cert_ca` block is documented below.
        /// </summary>
        public InputList<Inputs.SettingUserCertCaArgs> UserCertCas
        {
            get => _userCertCas ?? (_userCertCas = new InputList<Inputs.SettingUserCertCaArgs>());
            set => _userCertCas = value;
        }

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
        /// Active authentication method (scheme name).
        /// </summary>
        [Input("activeAuthScheme")]
        public Input<string>? ActiveAuthScheme { get; set; }

        /// <summary>
        /// Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authHttps")]
        public Input<string>? AuthHttps { get; set; }

        /// <summary>
        /// Captive portal host name.
        /// </summary>
        [Input("captivePortal")]
        public Input<string>? CaptivePortal { get; set; }

        /// <summary>
        /// IPv6 captive portal host name.
        /// </summary>
        [Input("captivePortal6")]
        public Input<string>? CaptivePortal6 { get; set; }

        /// <summary>
        /// Captive portal IP address.
        /// </summary>
        [Input("captivePortalIp")]
        public Input<string>? CaptivePortalIp { get; set; }

        /// <summary>
        /// Captive portal IPv6 address.
        /// </summary>
        [Input("captivePortalIp6")]
        public Input<string>? CaptivePortalIp6 { get; set; }

        /// <summary>
        /// Captive portal port number (1 - 65535, default = 7830).
        /// </summary>
        [Input("captivePortalPort")]
        public Input<int>? CaptivePortalPort { get; set; }

        /// <summary>
        /// Captive portal SSL port number (1 - 65535, default = 7831).
        /// </summary>
        [Input("captivePortalSslPort")]
        public Input<int>? CaptivePortalSslPort { get; set; }

        /// <summary>
        /// Captive portal type. Valid values: `fqdn`, `ip`.
        /// </summary>
        [Input("captivePortalType")]
        public Input<string>? CaptivePortalType { get; set; }

        /// <summary>
        /// Enable/disable redirecting certificate authentication to HTTPS portal. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("certAuth")]
        public Input<string>? CertAuth { get; set; }

        /// <summary>
        /// Certificate captive portal host name.
        /// </summary>
        [Input("certCaptivePortal")]
        public Input<string>? CertCaptivePortal { get; set; }

        /// <summary>
        /// Certificate captive portal IP address.
        /// </summary>
        [Input("certCaptivePortalIp")]
        public Input<string>? CertCaptivePortalIp { get; set; }

        /// <summary>
        /// Certificate captive portal port number (1 - 65535, default = 7832).
        /// </summary>
        [Input("certCaptivePortalPort")]
        public Input<int>? CertCaptivePortalPort { get; set; }

        /// <summary>
        /// Persistent web portal cookie maximum age in minutes (30 - 10080 (1 week), default = 480 (8 hours)).
        /// </summary>
        [Input("cookieMaxAge")]
        public Input<int>? CookieMaxAge { get; set; }

        /// <summary>
        /// Refresh rate divider of persistent web portal cookie (default = 2). Refresh value = cookie-max-age/cookie-refresh-div.
        /// </summary>
        [Input("cookieRefreshDiv")]
        public Input<int>? CookieRefreshDiv { get; set; }

        [Input("devRanges")]
        private InputList<Inputs.SettingDevRangeGetArgs>? _devRanges;

        /// <summary>
        /// Address range for the IP based device query. The structure of `dev_range` block is documented below.
        /// </summary>
        public InputList<Inputs.SettingDevRangeGetArgs> DevRanges
        {
            get => _devRanges ?? (_devRanges = new InputList<Inputs.SettingDevRangeGetArgs>());
            set => _devRanges = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable persistent cookie on IP based web portal authentication (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipAuthCookie")]
        public Input<string>? IpAuthCookie { get; set; }

        /// <summary>
        /// Enable/disable persistent cookie on web portal authentication (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("persistentCookie")]
        public Input<string>? PersistentCookie { get; set; }

        /// <summary>
        /// Single-Sign-On authentication method (scheme name).
        /// </summary>
        [Input("ssoAuthScheme")]
        public Input<string>? SsoAuthScheme { get; set; }

        /// <summary>
        /// Time of the last update.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        [Input("userCertCas")]
        private InputList<Inputs.SettingUserCertCaGetArgs>? _userCertCas;

        /// <summary>
        /// CA certificate used for client certificate verification. The structure of `user_cert_ca` block is documented below.
        /// </summary>
        public InputList<Inputs.SettingUserCertCaGetArgs> UserCertCas
        {
            get => _userCertCas ?? (_userCertCas = new InputList<Inputs.SettingUserCertCaGetArgs>());
            set => _userCertCas = value;
        }

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
