// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.User
{
    /// <summary>
    /// Configure Fortinet Single Sign On (FSSO) agents.
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
    ///     var trname = new Fortios.User.Fsso("trname", new()
    ///     {
    ///         Port = 32381,
    ///         Port2 = 8000,
    ///         Port3 = 8000,
    ///         Port4 = 8000,
    ///         Port5 = 8000,
    ///         Server = "1.1.1.1",
    ///         SourceIp = "0.0.0.0",
    ///         SourceIp6 = "::",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// User Fsso can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/fsso:Fsso labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/fsso:Fsso labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:user/fsso:Fsso")]
    public partial class Fsso : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Interval in minutes within to fetch groups from FSSO server, or unset to disable.
        /// </summary>
        [Output("groupPollInterval")]
        public Output<int> GroupPollInterval { get; private set; } = null!;

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
        /// Enable/disable automatic fetching of groups from LDAP server. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ldapPoll")]
        public Output<string> LdapPoll { get; private set; } = null!;

        /// <summary>
        /// Filter used to fetch groups.
        /// </summary>
        [Output("ldapPollFilter")]
        public Output<string> LdapPollFilter { get; private set; } = null!;

        /// <summary>
        /// Interval in minutes within to fetch groups from LDAP server.
        /// </summary>
        [Output("ldapPollInterval")]
        public Output<int> LdapPollInterval { get; private set; } = null!;

        /// <summary>
        /// LDAP server to get group information.
        /// </summary>
        [Output("ldapServer")]
        public Output<string> LdapServer { get; private set; } = null!;

        /// <summary>
        /// Interval in minutes to keep logons after FSSO server down.
        /// </summary>
        [Output("logonTimeout")]
        public Output<int> LogonTimeout { get; private set; } = null!;

        /// <summary>
        /// Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Password of the first FSSO collector agent.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Password of the second FSSO collector agent.
        /// </summary>
        [Output("password2")]
        public Output<string?> Password2 { get; private set; } = null!;

        /// <summary>
        /// Password of the third FSSO collector agent.
        /// </summary>
        [Output("password3")]
        public Output<string?> Password3 { get; private set; } = null!;

        /// <summary>
        /// Password of the fourth FSSO collector agent.
        /// </summary>
        [Output("password4")]
        public Output<string?> Password4 { get; private set; } = null!;

        /// <summary>
        /// Password of the fifth FSSO collector agent.
        /// </summary>
        [Output("password5")]
        public Output<string?> Password5 { get; private set; } = null!;

        /// <summary>
        /// Port of the first FSSO collector agent.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Port of the second FSSO collector agent.
        /// </summary>
        [Output("port2")]
        public Output<int> Port2 { get; private set; } = null!;

        /// <summary>
        /// Port of the third FSSO collector agent.
        /// </summary>
        [Output("port3")]
        public Output<int> Port3 { get; private set; } = null!;

        /// <summary>
        /// Port of the fourth FSSO collector agent.
        /// </summary>
        [Output("port4")]
        public Output<int> Port4 { get; private set; } = null!;

        /// <summary>
        /// Port of the fifth FSSO collector agent.
        /// </summary>
        [Output("port5")]
        public Output<int> Port5 { get; private set; } = null!;

        /// <summary>
        /// Domain name or IP address of the first FSSO collector agent.
        /// </summary>
        [Output("server")]
        public Output<string> Server { get; private set; } = null!;

        /// <summary>
        /// Domain name or IP address of the second FSSO collector agent.
        /// </summary>
        [Output("server2")]
        public Output<string> Server2 { get; private set; } = null!;

        /// <summary>
        /// Domain name or IP address of the third FSSO collector agent.
        /// </summary>
        [Output("server3")]
        public Output<string> Server3 { get; private set; } = null!;

        /// <summary>
        /// Domain name or IP address of the fourth FSSO collector agent.
        /// </summary>
        [Output("server4")]
        public Output<string> Server4 { get; private set; } = null!;

        /// <summary>
        /// Domain name or IP address of the fifth FSSO collector agent.
        /// </summary>
        [Output("server5")]
        public Output<string> Server5 { get; private set; } = null!;

        /// <summary>
        /// Server Name Indication.
        /// </summary>
        [Output("sni")]
        public Output<string> Sni { get; private set; } = null!;

        /// <summary>
        /// Source IP for communications to FSSO agent.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// IPv6 source for communications to FSSO agent.
        /// </summary>
        [Output("sourceIp6")]
        public Output<string> SourceIp6 { get; private set; } = null!;

        /// <summary>
        /// Enable/disable use of SSL. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ssl")]
        public Output<string> Ssl { get; private set; } = null!;

        /// <summary>
        /// Enable/disable server host/IP verification. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("sslServerHostIpCheck")]
        public Output<string> SslServerHostIpCheck { get; private set; } = null!;

        /// <summary>
        /// Trusted server certificate or CA certificate.
        /// </summary>
        [Output("sslTrustedCert")]
        public Output<string> SslTrustedCert { get; private set; } = null!;

        /// <summary>
        /// Server type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// LDAP server to get user information.
        /// </summary>
        [Output("userInfoServer")]
        public Output<string> UserInfoServer { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Fsso resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Fsso(string name, FssoArgs args, CustomResourceOptions? options = null)
            : base("fortios:user/fsso:Fsso", name, args ?? new FssoArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Fsso(string name, Input<string> id, FssoState? state = null, CustomResourceOptions? options = null)
            : base("fortios:user/fsso:Fsso", name, state, MakeResourceOptions(options, id))
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
                    "password",
                    "password2",
                    "password3",
                    "password4",
                    "password5",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Fsso resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Fsso Get(string name, Input<string> id, FssoState? state = null, CustomResourceOptions? options = null)
        {
            return new Fsso(name, id, state, options);
        }
    }

    public sealed class FssoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Interval in minutes within to fetch groups from FSSO server, or unset to disable.
        /// </summary>
        [Input("groupPollInterval")]
        public Input<int>? GroupPollInterval { get; set; }

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
        /// Enable/disable automatic fetching of groups from LDAP server. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ldapPoll")]
        public Input<string>? LdapPoll { get; set; }

        /// <summary>
        /// Filter used to fetch groups.
        /// </summary>
        [Input("ldapPollFilter")]
        public Input<string>? LdapPollFilter { get; set; }

        /// <summary>
        /// Interval in minutes within to fetch groups from LDAP server.
        /// </summary>
        [Input("ldapPollInterval")]
        public Input<int>? LdapPollInterval { get; set; }

        /// <summary>
        /// LDAP server to get group information.
        /// </summary>
        [Input("ldapServer")]
        public Input<string>? LdapServer { get; set; }

        /// <summary>
        /// Interval in minutes to keep logons after FSSO server down.
        /// </summary>
        [Input("logonTimeout")]
        public Input<int>? LogonTimeout { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password of the first FSSO collector agent.
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

        [Input("password2")]
        private Input<string>? _password2;

        /// <summary>
        /// Password of the second FSSO collector agent.
        /// </summary>
        public Input<string>? Password2
        {
            get => _password2;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password2 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("password3")]
        private Input<string>? _password3;

        /// <summary>
        /// Password of the third FSSO collector agent.
        /// </summary>
        public Input<string>? Password3
        {
            get => _password3;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password3 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("password4")]
        private Input<string>? _password4;

        /// <summary>
        /// Password of the fourth FSSO collector agent.
        /// </summary>
        public Input<string>? Password4
        {
            get => _password4;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password4 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("password5")]
        private Input<string>? _password5;

        /// <summary>
        /// Password of the fifth FSSO collector agent.
        /// </summary>
        public Input<string>? Password5
        {
            get => _password5;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password5 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Port of the first FSSO collector agent.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Port of the second FSSO collector agent.
        /// </summary>
        [Input("port2")]
        public Input<int>? Port2 { get; set; }

        /// <summary>
        /// Port of the third FSSO collector agent.
        /// </summary>
        [Input("port3")]
        public Input<int>? Port3 { get; set; }

        /// <summary>
        /// Port of the fourth FSSO collector agent.
        /// </summary>
        [Input("port4")]
        public Input<int>? Port4 { get; set; }

        /// <summary>
        /// Port of the fifth FSSO collector agent.
        /// </summary>
        [Input("port5")]
        public Input<int>? Port5 { get; set; }

        /// <summary>
        /// Domain name or IP address of the first FSSO collector agent.
        /// </summary>
        [Input("server", required: true)]
        public Input<string> Server { get; set; } = null!;

        /// <summary>
        /// Domain name or IP address of the second FSSO collector agent.
        /// </summary>
        [Input("server2")]
        public Input<string>? Server2 { get; set; }

        /// <summary>
        /// Domain name or IP address of the third FSSO collector agent.
        /// </summary>
        [Input("server3")]
        public Input<string>? Server3 { get; set; }

        /// <summary>
        /// Domain name or IP address of the fourth FSSO collector agent.
        /// </summary>
        [Input("server4")]
        public Input<string>? Server4 { get; set; }

        /// <summary>
        /// Domain name or IP address of the fifth FSSO collector agent.
        /// </summary>
        [Input("server5")]
        public Input<string>? Server5 { get; set; }

        /// <summary>
        /// Server Name Indication.
        /// </summary>
        [Input("sni")]
        public Input<string>? Sni { get; set; }

        /// <summary>
        /// Source IP for communications to FSSO agent.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// IPv6 source for communications to FSSO agent.
        /// </summary>
        [Input("sourceIp6")]
        public Input<string>? SourceIp6 { get; set; }

        /// <summary>
        /// Enable/disable use of SSL. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ssl")]
        public Input<string>? Ssl { get; set; }

        /// <summary>
        /// Enable/disable server host/IP verification. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sslServerHostIpCheck")]
        public Input<string>? SslServerHostIpCheck { get; set; }

        /// <summary>
        /// Trusted server certificate or CA certificate.
        /// </summary>
        [Input("sslTrustedCert")]
        public Input<string>? SslTrustedCert { get; set; }

        /// <summary>
        /// Server type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// LDAP server to get user information.
        /// </summary>
        [Input("userInfoServer")]
        public Input<string>? UserInfoServer { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FssoArgs()
        {
        }
        public static new FssoArgs Empty => new FssoArgs();
    }

    public sealed class FssoState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Interval in minutes within to fetch groups from FSSO server, or unset to disable.
        /// </summary>
        [Input("groupPollInterval")]
        public Input<int>? GroupPollInterval { get; set; }

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
        /// Enable/disable automatic fetching of groups from LDAP server. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ldapPoll")]
        public Input<string>? LdapPoll { get; set; }

        /// <summary>
        /// Filter used to fetch groups.
        /// </summary>
        [Input("ldapPollFilter")]
        public Input<string>? LdapPollFilter { get; set; }

        /// <summary>
        /// Interval in minutes within to fetch groups from LDAP server.
        /// </summary>
        [Input("ldapPollInterval")]
        public Input<int>? LdapPollInterval { get; set; }

        /// <summary>
        /// LDAP server to get group information.
        /// </summary>
        [Input("ldapServer")]
        public Input<string>? LdapServer { get; set; }

        /// <summary>
        /// Interval in minutes to keep logons after FSSO server down.
        /// </summary>
        [Input("logonTimeout")]
        public Input<int>? LogonTimeout { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password of the first FSSO collector agent.
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

        [Input("password2")]
        private Input<string>? _password2;

        /// <summary>
        /// Password of the second FSSO collector agent.
        /// </summary>
        public Input<string>? Password2
        {
            get => _password2;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password2 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("password3")]
        private Input<string>? _password3;

        /// <summary>
        /// Password of the third FSSO collector agent.
        /// </summary>
        public Input<string>? Password3
        {
            get => _password3;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password3 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("password4")]
        private Input<string>? _password4;

        /// <summary>
        /// Password of the fourth FSSO collector agent.
        /// </summary>
        public Input<string>? Password4
        {
            get => _password4;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password4 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("password5")]
        private Input<string>? _password5;

        /// <summary>
        /// Password of the fifth FSSO collector agent.
        /// </summary>
        public Input<string>? Password5
        {
            get => _password5;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password5 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Port of the first FSSO collector agent.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Port of the second FSSO collector agent.
        /// </summary>
        [Input("port2")]
        public Input<int>? Port2 { get; set; }

        /// <summary>
        /// Port of the third FSSO collector agent.
        /// </summary>
        [Input("port3")]
        public Input<int>? Port3 { get; set; }

        /// <summary>
        /// Port of the fourth FSSO collector agent.
        /// </summary>
        [Input("port4")]
        public Input<int>? Port4 { get; set; }

        /// <summary>
        /// Port of the fifth FSSO collector agent.
        /// </summary>
        [Input("port5")]
        public Input<int>? Port5 { get; set; }

        /// <summary>
        /// Domain name or IP address of the first FSSO collector agent.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Domain name or IP address of the second FSSO collector agent.
        /// </summary>
        [Input("server2")]
        public Input<string>? Server2 { get; set; }

        /// <summary>
        /// Domain name or IP address of the third FSSO collector agent.
        /// </summary>
        [Input("server3")]
        public Input<string>? Server3 { get; set; }

        /// <summary>
        /// Domain name or IP address of the fourth FSSO collector agent.
        /// </summary>
        [Input("server4")]
        public Input<string>? Server4 { get; set; }

        /// <summary>
        /// Domain name or IP address of the fifth FSSO collector agent.
        /// </summary>
        [Input("server5")]
        public Input<string>? Server5 { get; set; }

        /// <summary>
        /// Server Name Indication.
        /// </summary>
        [Input("sni")]
        public Input<string>? Sni { get; set; }

        /// <summary>
        /// Source IP for communications to FSSO agent.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// IPv6 source for communications to FSSO agent.
        /// </summary>
        [Input("sourceIp6")]
        public Input<string>? SourceIp6 { get; set; }

        /// <summary>
        /// Enable/disable use of SSL. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ssl")]
        public Input<string>? Ssl { get; set; }

        /// <summary>
        /// Enable/disable server host/IP verification. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sslServerHostIpCheck")]
        public Input<string>? SslServerHostIpCheck { get; set; }

        /// <summary>
        /// Trusted server certificate or CA certificate.
        /// </summary>
        [Input("sslTrustedCert")]
        public Input<string>? SslTrustedCert { get; set; }

        /// <summary>
        /// Server type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// LDAP server to get user information.
        /// </summary>
        [Input("userInfoServer")]
        public Input<string>? UserInfoServer { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FssoState()
        {
        }
        public static new FssoState Empty => new FssoState();
    }
}
