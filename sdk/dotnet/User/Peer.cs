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
    /// Configure peer users.
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
    ///     var trname1 = new Fortios.User.Peer("trname1", new()
    ///     {
    ///         Ca = "EC-ACC",
    ///         CnType = "string",
    ///         LdapMode = "password",
    ///         MandatoryCaVerify = "enable",
    ///         TwoFactor = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// User Peer can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/peer:Peer labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/peer:Peer labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:user/peer:Peer")]
    public partial class Peer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the CA certificate as returned by the execute vpn certificate ca list command.
        /// </summary>
        [Output("ca")]
        public Output<string> Ca { get; private set; } = null!;

        /// <summary>
        /// Peer certificate common name.
        /// </summary>
        [Output("cn")]
        public Output<string> Cn { get; private set; } = null!;

        /// <summary>
        /// Peer certificate common name type. Valid values: `string`, `email`, `FQDN`, `ipv4`, `ipv6`.
        /// </summary>
        [Output("cnType")]
        public Output<string> CnType { get; private set; } = null!;

        /// <summary>
        /// Mode for LDAP peer authentication. Valid values: `password`, `principal-name`.
        /// </summary>
        [Output("ldapMode")]
        public Output<string> LdapMode { get; private set; } = null!;

        /// <summary>
        /// Password for LDAP server bind.
        /// </summary>
        [Output("ldapPassword")]
        public Output<string?> LdapPassword { get; private set; } = null!;

        /// <summary>
        /// Name of an LDAP server defined under the user ldap command. Performs client access rights check.
        /// </summary>
        [Output("ldapServer")]
        public Output<string> LdapServer { get; private set; } = null!;

        /// <summary>
        /// Username for LDAP server bind.
        /// </summary>
        [Output("ldapUsername")]
        public Output<string> LdapUsername { get; private set; } = null!;

        /// <summary>
        /// Determine what happens to the peer if the CA certificate is not installed. Disable to automatically consider the peer certificate as valid. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("mandatoryCaVerify")]
        public Output<string> MandatoryCaVerify { get; private set; } = null!;

        /// <summary>
        /// MFA mode for remote peer authentication/authorization. Valid values: `none`, `password`, `subject-identity`.
        /// </summary>
        [Output("mfaMode")]
        public Output<string> MfaMode { get; private set; } = null!;

        /// <summary>
        /// Unified password for remote authentication. This field may be left empty when RADIUS authentication is used, in which case the FortiGate will use the RADIUS username as a password.
        /// </summary>
        [Output("mfaPassword")]
        public Output<string?> MfaPassword { get; private set; } = null!;

        /// <summary>
        /// Name of a remote authenticator. Performs client access right check.
        /// </summary>
        [Output("mfaServer")]
        public Output<string> MfaServer { get; private set; } = null!;

        /// <summary>
        /// Unified username for remote authentication.
        /// </summary>
        [Output("mfaUsername")]
        public Output<string> MfaUsername { get; private set; } = null!;

        /// <summary>
        /// Peer name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Online Certificate Status Protocol (OCSP) server for certificate retrieval.
        /// </summary>
        [Output("ocspOverrideServer")]
        public Output<string> OcspOverrideServer { get; private set; } = null!;

        /// <summary>
        /// Peer's password used for two-factor authentication.
        /// </summary>
        [Output("passwd")]
        public Output<string?> Passwd { get; private set; } = null!;

        /// <summary>
        /// Peer certificate name constraints.
        /// </summary>
        [Output("subject")]
        public Output<string> Subject { get; private set; } = null!;

        /// <summary>
        /// Enable/disable two-factor authentication, applying certificate and password-based authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("twoFactor")]
        public Output<string> TwoFactor { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Peer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Peer(string name, PeerArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:user/peer:Peer", name, args ?? new PeerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Peer(string name, Input<string> id, PeerState? state = null, CustomResourceOptions? options = null)
            : base("fortios:user/peer:Peer", name, state, MakeResourceOptions(options, id))
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
                    "ldapPassword",
                    "passwd",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Peer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Peer Get(string name, Input<string> id, PeerState? state = null, CustomResourceOptions? options = null)
        {
            return new Peer(name, id, state, options);
        }
    }

    public sealed class PeerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the CA certificate as returned by the execute vpn certificate ca list command.
        /// </summary>
        [Input("ca")]
        public Input<string>? Ca { get; set; }

        /// <summary>
        /// Peer certificate common name.
        /// </summary>
        [Input("cn")]
        public Input<string>? Cn { get; set; }

        /// <summary>
        /// Peer certificate common name type. Valid values: `string`, `email`, `FQDN`, `ipv4`, `ipv6`.
        /// </summary>
        [Input("cnType")]
        public Input<string>? CnType { get; set; }

        /// <summary>
        /// Mode for LDAP peer authentication. Valid values: `password`, `principal-name`.
        /// </summary>
        [Input("ldapMode")]
        public Input<string>? LdapMode { get; set; }

        [Input("ldapPassword")]
        private Input<string>? _ldapPassword;

        /// <summary>
        /// Password for LDAP server bind.
        /// </summary>
        public Input<string>? LdapPassword
        {
            get => _ldapPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _ldapPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Name of an LDAP server defined under the user ldap command. Performs client access rights check.
        /// </summary>
        [Input("ldapServer")]
        public Input<string>? LdapServer { get; set; }

        /// <summary>
        /// Username for LDAP server bind.
        /// </summary>
        [Input("ldapUsername")]
        public Input<string>? LdapUsername { get; set; }

        /// <summary>
        /// Determine what happens to the peer if the CA certificate is not installed. Disable to automatically consider the peer certificate as valid. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("mandatoryCaVerify")]
        public Input<string>? MandatoryCaVerify { get; set; }

        /// <summary>
        /// MFA mode for remote peer authentication/authorization. Valid values: `none`, `password`, `subject-identity`.
        /// </summary>
        [Input("mfaMode")]
        public Input<string>? MfaMode { get; set; }

        /// <summary>
        /// Unified password for remote authentication. This field may be left empty when RADIUS authentication is used, in which case the FortiGate will use the RADIUS username as a password.
        /// </summary>
        [Input("mfaPassword")]
        public Input<string>? MfaPassword { get; set; }

        /// <summary>
        /// Name of a remote authenticator. Performs client access right check.
        /// </summary>
        [Input("mfaServer")]
        public Input<string>? MfaServer { get; set; }

        /// <summary>
        /// Unified username for remote authentication.
        /// </summary>
        [Input("mfaUsername")]
        public Input<string>? MfaUsername { get; set; }

        /// <summary>
        /// Peer name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Online Certificate Status Protocol (OCSP) server for certificate retrieval.
        /// </summary>
        [Input("ocspOverrideServer")]
        public Input<string>? OcspOverrideServer { get; set; }

        [Input("passwd")]
        private Input<string>? _passwd;

        /// <summary>
        /// Peer's password used for two-factor authentication.
        /// </summary>
        public Input<string>? Passwd
        {
            get => _passwd;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _passwd = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Peer certificate name constraints.
        /// </summary>
        [Input("subject")]
        public Input<string>? Subject { get; set; }

        /// <summary>
        /// Enable/disable two-factor authentication, applying certificate and password-based authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("twoFactor")]
        public Input<string>? TwoFactor { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public PeerArgs()
        {
        }
        public static new PeerArgs Empty => new PeerArgs();
    }

    public sealed class PeerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the CA certificate as returned by the execute vpn certificate ca list command.
        /// </summary>
        [Input("ca")]
        public Input<string>? Ca { get; set; }

        /// <summary>
        /// Peer certificate common name.
        /// </summary>
        [Input("cn")]
        public Input<string>? Cn { get; set; }

        /// <summary>
        /// Peer certificate common name type. Valid values: `string`, `email`, `FQDN`, `ipv4`, `ipv6`.
        /// </summary>
        [Input("cnType")]
        public Input<string>? CnType { get; set; }

        /// <summary>
        /// Mode for LDAP peer authentication. Valid values: `password`, `principal-name`.
        /// </summary>
        [Input("ldapMode")]
        public Input<string>? LdapMode { get; set; }

        [Input("ldapPassword")]
        private Input<string>? _ldapPassword;

        /// <summary>
        /// Password for LDAP server bind.
        /// </summary>
        public Input<string>? LdapPassword
        {
            get => _ldapPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _ldapPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Name of an LDAP server defined under the user ldap command. Performs client access rights check.
        /// </summary>
        [Input("ldapServer")]
        public Input<string>? LdapServer { get; set; }

        /// <summary>
        /// Username for LDAP server bind.
        /// </summary>
        [Input("ldapUsername")]
        public Input<string>? LdapUsername { get; set; }

        /// <summary>
        /// Determine what happens to the peer if the CA certificate is not installed. Disable to automatically consider the peer certificate as valid. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("mandatoryCaVerify")]
        public Input<string>? MandatoryCaVerify { get; set; }

        /// <summary>
        /// MFA mode for remote peer authentication/authorization. Valid values: `none`, `password`, `subject-identity`.
        /// </summary>
        [Input("mfaMode")]
        public Input<string>? MfaMode { get; set; }

        /// <summary>
        /// Unified password for remote authentication. This field may be left empty when RADIUS authentication is used, in which case the FortiGate will use the RADIUS username as a password.
        /// </summary>
        [Input("mfaPassword")]
        public Input<string>? MfaPassword { get; set; }

        /// <summary>
        /// Name of a remote authenticator. Performs client access right check.
        /// </summary>
        [Input("mfaServer")]
        public Input<string>? MfaServer { get; set; }

        /// <summary>
        /// Unified username for remote authentication.
        /// </summary>
        [Input("mfaUsername")]
        public Input<string>? MfaUsername { get; set; }

        /// <summary>
        /// Peer name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Online Certificate Status Protocol (OCSP) server for certificate retrieval.
        /// </summary>
        [Input("ocspOverrideServer")]
        public Input<string>? OcspOverrideServer { get; set; }

        [Input("passwd")]
        private Input<string>? _passwd;

        /// <summary>
        /// Peer's password used for two-factor authentication.
        /// </summary>
        public Input<string>? Passwd
        {
            get => _passwd;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _passwd = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Peer certificate name constraints.
        /// </summary>
        [Input("subject")]
        public Input<string>? Subject { get; set; }

        /// <summary>
        /// Enable/disable two-factor authentication, applying certificate and password-based authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("twoFactor")]
        public Input<string>? TwoFactor { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public PeerState()
        {
        }
        public static new PeerState Empty => new PeerState();
    }
}
