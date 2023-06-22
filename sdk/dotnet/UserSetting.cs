// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    /// <summary>
    /// Configure user authentication setting.
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
    ///     var trname = new Fortios.UserSetting("trname", new()
    ///     {
    ///         AuthBlackoutTime = 0,
    ///         AuthCert = "Fortinet_Factory",
    ///         AuthHttpBasic = "disable",
    ///         AuthInvalidMax = 5,
    ///         AuthLockoutDuration = 0,
    ///         AuthLockoutThreshold = 3,
    ///         AuthOnDemand = "implicitly",
    ///         AuthPortalTimeout = 3,
    ///         AuthSecureHttp = "disable",
    ///         AuthSrcMac = "enable",
    ///         AuthSslAllowRenegotiation = "disable",
    ///         AuthTimeout = 5,
    ///         AuthTimeoutType = "idle-timeout",
    ///         AuthType = "http https ftp telnet",
    ///         RadiusSesTimeoutAct = "hard-timeout",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// User Setting can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/userSetting:UserSetting labelname UserSetting
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/userSetting:UserSetting labelname UserSetting
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/userSetting:UserSetting")]
    public partial class UserSetting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Time in seconds an IP address is denied access after failing to authenticate five times within one minute.
        /// </summary>
        [Output("authBlackoutTime")]
        public Output<int> AuthBlackoutTime { get; private set; } = null!;

        /// <summary>
        /// HTTPS CA certificate for policy authentication.
        /// </summary>
        [Output("authCaCert")]
        public Output<string> AuthCaCert { get; private set; } = null!;

        /// <summary>
        /// HTTPS server certificate for policy authentication.
        /// </summary>
        [Output("authCert")]
        public Output<string> AuthCert { get; private set; } = null!;

        /// <summary>
        /// Enable/disable use of HTTP basic authentication for identity-based firewall policies. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("authHttpBasic")]
        public Output<string> AuthHttpBasic { get; private set; } = null!;

        /// <summary>
        /// Maximum number of failed authentication attempts before the user is blocked.
        /// </summary>
        [Output("authInvalidMax")]
        public Output<int> AuthInvalidMax { get; private set; } = null!;

        /// <summary>
        /// Lockout period in seconds after too many login failures.
        /// </summary>
        [Output("authLockoutDuration")]
        public Output<int> AuthLockoutDuration { get; private set; } = null!;

        /// <summary>
        /// Maximum number of failed login attempts before login lockout is triggered.
        /// </summary>
        [Output("authLockoutThreshold")]
        public Output<int> AuthLockoutThreshold { get; private set; } = null!;

        /// <summary>
        /// Always/implicitly trigger firewall authentication on demand. Valid values: `always`, `implicitly`.
        /// </summary>
        [Output("authOnDemand")]
        public Output<string> AuthOnDemand { get; private set; } = null!;

        /// <summary>
        /// Time in minutes before captive portal user have to re-authenticate (1 - 30 min, default 3 min).
        /// </summary>
        [Output("authPortalTimeout")]
        public Output<int> AuthPortalTimeout { get; private set; } = null!;

        /// <summary>
        /// Set up non-standard ports for authentication with HTTP, HTTPS, FTP, and TELNET. The structure of `auth_ports` block is documented below.
        /// </summary>
        [Output("authPorts")]
        public Output<ImmutableArray<Outputs.UserSettingAuthPort>> AuthPorts { get; private set; } = null!;

        /// <summary>
        /// Enable/disable redirecting HTTP user authentication to more secure HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("authSecureHttp")]
        public Output<string> AuthSecureHttp { get; private set; } = null!;

        /// <summary>
        /// Enable/disable source MAC for user identity. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("authSrcMac")]
        public Output<string> AuthSrcMac { get; private set; } = null!;

        /// <summary>
        /// Allow/forbid SSL re-negotiation for HTTPS authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("authSslAllowRenegotiation")]
        public Output<string> AuthSslAllowRenegotiation { get; private set; } = null!;

        /// <summary>
        /// Maximum supported protocol version for SSL/TLS connections (default is no limit). Valid values: `sslv3`, `tlsv1`, `tlsv1-1`, `tlsv1-2`, `tlsv1-3`.
        /// </summary>
        [Output("authSslMaxProtoVersion")]
        public Output<string> AuthSslMaxProtoVersion { get; private set; } = null!;

        /// <summary>
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        /// </summary>
        [Output("authSslMinProtoVersion")]
        public Output<string> AuthSslMinProtoVersion { get; private set; } = null!;

        /// <summary>
        /// Set signature algorithms related to HTTPS authentication (affects TLS version &lt;= 1.2 only, default is to enable all). Valid values: `no-rsa-pss`, `all`.
        /// </summary>
        [Output("authSslSigalgs")]
        public Output<string> AuthSslSigalgs { get; private set; } = null!;

        /// <summary>
        /// Time in minutes before the firewall user authentication timeout requires the user to re-authenticate.
        /// </summary>
        [Output("authTimeout")]
        public Output<int> AuthTimeout { get; private set; } = null!;

        /// <summary>
        /// Control if authenticated users have to login again after a hard timeout, after an idle timeout, or after a session timeout. Valid values: `idle-timeout`, `hard-timeout`, `new-session`.
        /// </summary>
        [Output("authTimeoutType")]
        public Output<string> AuthTimeoutType { get; private set; } = null!;

        /// <summary>
        /// Supported firewall policy authentication protocols/methods. Valid values: `http`, `https`, `ftp`, `telnet`.
        /// </summary>
        [Output("authType")]
        public Output<string> AuthType { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Enable/disable per policy disclaimer. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("perPolicyDisclaimer")]
        public Output<string> PerPolicyDisclaimer { get; private set; } = null!;

        /// <summary>
        /// Set the RADIUS session timeout to a hard timeout or to ignore RADIUS server session timeouts. Valid values: `hard-timeout`, `ignore-timeout`.
        /// </summary>
        [Output("radiusSesTimeoutAct")]
        public Output<string> RadiusSesTimeoutAct { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a UserSetting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserSetting(string name, UserSettingArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/userSetting:UserSetting", name, args ?? new UserSettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserSetting(string name, Input<string> id, UserSettingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/userSetting:UserSetting", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserSetting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserSetting Get(string name, Input<string> id, UserSettingState? state = null, CustomResourceOptions? options = null)
        {
            return new UserSetting(name, id, state, options);
        }
    }

    public sealed class UserSettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time in seconds an IP address is denied access after failing to authenticate five times within one minute.
        /// </summary>
        [Input("authBlackoutTime")]
        public Input<int>? AuthBlackoutTime { get; set; }

        /// <summary>
        /// HTTPS CA certificate for policy authentication.
        /// </summary>
        [Input("authCaCert")]
        public Input<string>? AuthCaCert { get; set; }

        /// <summary>
        /// HTTPS server certificate for policy authentication.
        /// </summary>
        [Input("authCert")]
        public Input<string>? AuthCert { get; set; }

        /// <summary>
        /// Enable/disable use of HTTP basic authentication for identity-based firewall policies. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authHttpBasic")]
        public Input<string>? AuthHttpBasic { get; set; }

        /// <summary>
        /// Maximum number of failed authentication attempts before the user is blocked.
        /// </summary>
        [Input("authInvalidMax")]
        public Input<int>? AuthInvalidMax { get; set; }

        /// <summary>
        /// Lockout period in seconds after too many login failures.
        /// </summary>
        [Input("authLockoutDuration")]
        public Input<int>? AuthLockoutDuration { get; set; }

        /// <summary>
        /// Maximum number of failed login attempts before login lockout is triggered.
        /// </summary>
        [Input("authLockoutThreshold")]
        public Input<int>? AuthLockoutThreshold { get; set; }

        /// <summary>
        /// Always/implicitly trigger firewall authentication on demand. Valid values: `always`, `implicitly`.
        /// </summary>
        [Input("authOnDemand")]
        public Input<string>? AuthOnDemand { get; set; }

        /// <summary>
        /// Time in minutes before captive portal user have to re-authenticate (1 - 30 min, default 3 min).
        /// </summary>
        [Input("authPortalTimeout")]
        public Input<int>? AuthPortalTimeout { get; set; }

        [Input("authPorts")]
        private InputList<Inputs.UserSettingAuthPortArgs>? _authPorts;

        /// <summary>
        /// Set up non-standard ports for authentication with HTTP, HTTPS, FTP, and TELNET. The structure of `auth_ports` block is documented below.
        /// </summary>
        public InputList<Inputs.UserSettingAuthPortArgs> AuthPorts
        {
            get => _authPorts ?? (_authPorts = new InputList<Inputs.UserSettingAuthPortArgs>());
            set => _authPorts = value;
        }

        /// <summary>
        /// Enable/disable redirecting HTTP user authentication to more secure HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authSecureHttp")]
        public Input<string>? AuthSecureHttp { get; set; }

        /// <summary>
        /// Enable/disable source MAC for user identity. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authSrcMac")]
        public Input<string>? AuthSrcMac { get; set; }

        /// <summary>
        /// Allow/forbid SSL re-negotiation for HTTPS authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authSslAllowRenegotiation")]
        public Input<string>? AuthSslAllowRenegotiation { get; set; }

        /// <summary>
        /// Maximum supported protocol version for SSL/TLS connections (default is no limit). Valid values: `sslv3`, `tlsv1`, `tlsv1-1`, `tlsv1-2`, `tlsv1-3`.
        /// </summary>
        [Input("authSslMaxProtoVersion")]
        public Input<string>? AuthSslMaxProtoVersion { get; set; }

        /// <summary>
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        /// </summary>
        [Input("authSslMinProtoVersion")]
        public Input<string>? AuthSslMinProtoVersion { get; set; }

        /// <summary>
        /// Set signature algorithms related to HTTPS authentication (affects TLS version &lt;= 1.2 only, default is to enable all). Valid values: `no-rsa-pss`, `all`.
        /// </summary>
        [Input("authSslSigalgs")]
        public Input<string>? AuthSslSigalgs { get; set; }

        /// <summary>
        /// Time in minutes before the firewall user authentication timeout requires the user to re-authenticate.
        /// </summary>
        [Input("authTimeout")]
        public Input<int>? AuthTimeout { get; set; }

        /// <summary>
        /// Control if authenticated users have to login again after a hard timeout, after an idle timeout, or after a session timeout. Valid values: `idle-timeout`, `hard-timeout`, `new-session`.
        /// </summary>
        [Input("authTimeoutType")]
        public Input<string>? AuthTimeoutType { get; set; }

        /// <summary>
        /// Supported firewall policy authentication protocols/methods. Valid values: `http`, `https`, `ftp`, `telnet`.
        /// </summary>
        [Input("authType")]
        public Input<string>? AuthType { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable per policy disclaimer. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("perPolicyDisclaimer")]
        public Input<string>? PerPolicyDisclaimer { get; set; }

        /// <summary>
        /// Set the RADIUS session timeout to a hard timeout or to ignore RADIUS server session timeouts. Valid values: `hard-timeout`, `ignore-timeout`.
        /// </summary>
        [Input("radiusSesTimeoutAct")]
        public Input<string>? RadiusSesTimeoutAct { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public UserSettingArgs()
        {
        }
        public static new UserSettingArgs Empty => new UserSettingArgs();
    }

    public sealed class UserSettingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time in seconds an IP address is denied access after failing to authenticate five times within one minute.
        /// </summary>
        [Input("authBlackoutTime")]
        public Input<int>? AuthBlackoutTime { get; set; }

        /// <summary>
        /// HTTPS CA certificate for policy authentication.
        /// </summary>
        [Input("authCaCert")]
        public Input<string>? AuthCaCert { get; set; }

        /// <summary>
        /// HTTPS server certificate for policy authentication.
        /// </summary>
        [Input("authCert")]
        public Input<string>? AuthCert { get; set; }

        /// <summary>
        /// Enable/disable use of HTTP basic authentication for identity-based firewall policies. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authHttpBasic")]
        public Input<string>? AuthHttpBasic { get; set; }

        /// <summary>
        /// Maximum number of failed authentication attempts before the user is blocked.
        /// </summary>
        [Input("authInvalidMax")]
        public Input<int>? AuthInvalidMax { get; set; }

        /// <summary>
        /// Lockout period in seconds after too many login failures.
        /// </summary>
        [Input("authLockoutDuration")]
        public Input<int>? AuthLockoutDuration { get; set; }

        /// <summary>
        /// Maximum number of failed login attempts before login lockout is triggered.
        /// </summary>
        [Input("authLockoutThreshold")]
        public Input<int>? AuthLockoutThreshold { get; set; }

        /// <summary>
        /// Always/implicitly trigger firewall authentication on demand. Valid values: `always`, `implicitly`.
        /// </summary>
        [Input("authOnDemand")]
        public Input<string>? AuthOnDemand { get; set; }

        /// <summary>
        /// Time in minutes before captive portal user have to re-authenticate (1 - 30 min, default 3 min).
        /// </summary>
        [Input("authPortalTimeout")]
        public Input<int>? AuthPortalTimeout { get; set; }

        [Input("authPorts")]
        private InputList<Inputs.UserSettingAuthPortGetArgs>? _authPorts;

        /// <summary>
        /// Set up non-standard ports for authentication with HTTP, HTTPS, FTP, and TELNET. The structure of `auth_ports` block is documented below.
        /// </summary>
        public InputList<Inputs.UserSettingAuthPortGetArgs> AuthPorts
        {
            get => _authPorts ?? (_authPorts = new InputList<Inputs.UserSettingAuthPortGetArgs>());
            set => _authPorts = value;
        }

        /// <summary>
        /// Enable/disable redirecting HTTP user authentication to more secure HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authSecureHttp")]
        public Input<string>? AuthSecureHttp { get; set; }

        /// <summary>
        /// Enable/disable source MAC for user identity. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authSrcMac")]
        public Input<string>? AuthSrcMac { get; set; }

        /// <summary>
        /// Allow/forbid SSL re-negotiation for HTTPS authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authSslAllowRenegotiation")]
        public Input<string>? AuthSslAllowRenegotiation { get; set; }

        /// <summary>
        /// Maximum supported protocol version for SSL/TLS connections (default is no limit). Valid values: `sslv3`, `tlsv1`, `tlsv1-1`, `tlsv1-2`, `tlsv1-3`.
        /// </summary>
        [Input("authSslMaxProtoVersion")]
        public Input<string>? AuthSslMaxProtoVersion { get; set; }

        /// <summary>
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        /// </summary>
        [Input("authSslMinProtoVersion")]
        public Input<string>? AuthSslMinProtoVersion { get; set; }

        /// <summary>
        /// Set signature algorithms related to HTTPS authentication (affects TLS version &lt;= 1.2 only, default is to enable all). Valid values: `no-rsa-pss`, `all`.
        /// </summary>
        [Input("authSslSigalgs")]
        public Input<string>? AuthSslSigalgs { get; set; }

        /// <summary>
        /// Time in minutes before the firewall user authentication timeout requires the user to re-authenticate.
        /// </summary>
        [Input("authTimeout")]
        public Input<int>? AuthTimeout { get; set; }

        /// <summary>
        /// Control if authenticated users have to login again after a hard timeout, after an idle timeout, or after a session timeout. Valid values: `idle-timeout`, `hard-timeout`, `new-session`.
        /// </summary>
        [Input("authTimeoutType")]
        public Input<string>? AuthTimeoutType { get; set; }

        /// <summary>
        /// Supported firewall policy authentication protocols/methods. Valid values: `http`, `https`, `ftp`, `telnet`.
        /// </summary>
        [Input("authType")]
        public Input<string>? AuthType { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable per policy disclaimer. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("perPolicyDisclaimer")]
        public Input<string>? PerPolicyDisclaimer { get; set; }

        /// <summary>
        /// Set the RADIUS session timeout to a hard timeout or to ignore RADIUS server session timeouts. Valid values: `hard-timeout`, `ignore-timeout`.
        /// </summary>
        [Input("radiusSesTimeoutAct")]
        public Input<string>? RadiusSesTimeoutAct { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public UserSettingState()
        {
        }
        public static new UserSettingState Empty => new UserSettingState();
    }
}
