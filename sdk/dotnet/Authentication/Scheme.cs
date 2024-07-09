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
    /// Configure Authentication Schemes.
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
    ///     var trname3 = new Fortios.User.Fsso("trname3", new()
    ///     {
    ///         Port = 8000,
    ///         Port2 = 8000,
    ///         Port3 = 8000,
    ///         Port4 = 8000,
    ///         Port5 = 8000,
    ///         Server = "1.1.1.1",
    ///         SourceIp = "0.0.0.0",
    ///         SourceIp6 = "::",
    ///     });
    /// 
    ///     var trname = new Fortios.Authentication.Scheme("trname", new()
    ///     {
    ///         FssoAgentForNtlm = trname3.Name,
    ///         FssoGuest = "disable",
    ///         Method = "ntlm",
    ///         NegotiateNtlm = "enable",
    ///         RequireTfa = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Authentication Scheme can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:authentication/scheme:Scheme labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:authentication/scheme:Scheme labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:authentication/scheme:Scheme")]
    public partial class Scheme : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Domain controller setting.
        /// </summary>
        [Output("domainController")]
        public Output<string> DomainController { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// FSSO agent to use for NTLM authentication.
        /// </summary>
        [Output("fssoAgentForNtlm")]
        public Output<string> FssoAgentForNtlm { get; private set; } = null!;

        /// <summary>
        /// Enable/disable user fsso-guest authentication (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("fssoGuest")]
        public Output<string> FssoGuest { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Kerberos keytab setting.
        /// </summary>
        [Output("kerberosKeytab")]
        public Output<string> KerberosKeytab { get; private set; } = null!;

        /// <summary>
        /// Authentication methods (default = basic).
        /// </summary>
        [Output("method")]
        public Output<string> Method { get; private set; } = null!;

        /// <summary>
        /// Authentication scheme name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("negotiateNtlm")]
        public Output<string> NegotiateNtlm { get; private set; } = null!;

        /// <summary>
        /// Enable/disable two-factor authentication (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("requireTfa")]
        public Output<string> RequireTfa { get; private set; } = null!;

        /// <summary>
        /// SAML configuration.
        /// </summary>
        [Output("samlServer")]
        public Output<string> SamlServer { get; private set; } = null!;

        /// <summary>
        /// SAML authentication timeout in seconds.
        /// </summary>
        [Output("samlTimeout")]
        public Output<int> SamlTimeout { get; private set; } = null!;

        /// <summary>
        /// SSH CA name.
        /// </summary>
        [Output("sshCa")]
        public Output<string> SshCa { get; private set; } = null!;

        /// <summary>
        /// Enable/disable authentication with user certificate (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("userCert")]
        public Output<string> UserCert { get; private set; } = null!;

        /// <summary>
        /// Authentication server to contain user information; "local" (default) or "123" (for LDAP). The structure of `user_database` block is documented below.
        /// </summary>
        [Output("userDatabases")]
        public Output<ImmutableArray<Outputs.SchemeUserDatabase>> UserDatabases { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Scheme resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Scheme(string name, SchemeArgs args, CustomResourceOptions? options = null)
            : base("fortios:authentication/scheme:Scheme", name, args ?? new SchemeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Scheme(string name, Input<string> id, SchemeState? state = null, CustomResourceOptions? options = null)
            : base("fortios:authentication/scheme:Scheme", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Scheme resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Scheme Get(string name, Input<string> id, SchemeState? state = null, CustomResourceOptions? options = null)
        {
            return new Scheme(name, id, state, options);
        }
    }

    public sealed class SchemeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Domain controller setting.
        /// </summary>
        [Input("domainController")]
        public Input<string>? DomainController { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// FSSO agent to use for NTLM authentication.
        /// </summary>
        [Input("fssoAgentForNtlm")]
        public Input<string>? FssoAgentForNtlm { get; set; }

        /// <summary>
        /// Enable/disable user fsso-guest authentication (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fssoGuest")]
        public Input<string>? FssoGuest { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Kerberos keytab setting.
        /// </summary>
        [Input("kerberosKeytab")]
        public Input<string>? KerberosKeytab { get; set; }

        /// <summary>
        /// Authentication methods (default = basic).
        /// </summary>
        [Input("method", required: true)]
        public Input<string> Method { get; set; } = null!;

        /// <summary>
        /// Authentication scheme name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("negotiateNtlm")]
        public Input<string>? NegotiateNtlm { get; set; }

        /// <summary>
        /// Enable/disable two-factor authentication (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("requireTfa")]
        public Input<string>? RequireTfa { get; set; }

        /// <summary>
        /// SAML configuration.
        /// </summary>
        [Input("samlServer")]
        public Input<string>? SamlServer { get; set; }

        /// <summary>
        /// SAML authentication timeout in seconds.
        /// </summary>
        [Input("samlTimeout")]
        public Input<int>? SamlTimeout { get; set; }

        /// <summary>
        /// SSH CA name.
        /// </summary>
        [Input("sshCa")]
        public Input<string>? SshCa { get; set; }

        /// <summary>
        /// Enable/disable authentication with user certificate (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("userCert")]
        public Input<string>? UserCert { get; set; }

        [Input("userDatabases")]
        private InputList<Inputs.SchemeUserDatabaseArgs>? _userDatabases;

        /// <summary>
        /// Authentication server to contain user information; "local" (default) or "123" (for LDAP). The structure of `user_database` block is documented below.
        /// </summary>
        public InputList<Inputs.SchemeUserDatabaseArgs> UserDatabases
        {
            get => _userDatabases ?? (_userDatabases = new InputList<Inputs.SchemeUserDatabaseArgs>());
            set => _userDatabases = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SchemeArgs()
        {
        }
        public static new SchemeArgs Empty => new SchemeArgs();
    }

    public sealed class SchemeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Domain controller setting.
        /// </summary>
        [Input("domainController")]
        public Input<string>? DomainController { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// FSSO agent to use for NTLM authentication.
        /// </summary>
        [Input("fssoAgentForNtlm")]
        public Input<string>? FssoAgentForNtlm { get; set; }

        /// <summary>
        /// Enable/disable user fsso-guest authentication (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fssoGuest")]
        public Input<string>? FssoGuest { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Kerberos keytab setting.
        /// </summary>
        [Input("kerberosKeytab")]
        public Input<string>? KerberosKeytab { get; set; }

        /// <summary>
        /// Authentication methods (default = basic).
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        /// <summary>
        /// Authentication scheme name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("negotiateNtlm")]
        public Input<string>? NegotiateNtlm { get; set; }

        /// <summary>
        /// Enable/disable two-factor authentication (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("requireTfa")]
        public Input<string>? RequireTfa { get; set; }

        /// <summary>
        /// SAML configuration.
        /// </summary>
        [Input("samlServer")]
        public Input<string>? SamlServer { get; set; }

        /// <summary>
        /// SAML authentication timeout in seconds.
        /// </summary>
        [Input("samlTimeout")]
        public Input<int>? SamlTimeout { get; set; }

        /// <summary>
        /// SSH CA name.
        /// </summary>
        [Input("sshCa")]
        public Input<string>? SshCa { get; set; }

        /// <summary>
        /// Enable/disable authentication with user certificate (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("userCert")]
        public Input<string>? UserCert { get; set; }

        [Input("userDatabases")]
        private InputList<Inputs.SchemeUserDatabaseGetArgs>? _userDatabases;

        /// <summary>
        /// Authentication server to contain user information; "local" (default) or "123" (for LDAP). The structure of `user_database` block is documented below.
        /// </summary>
        public InputList<Inputs.SchemeUserDatabaseGetArgs> UserDatabases
        {
            get => _userDatabases ?? (_userDatabases = new InputList<Inputs.SchemeUserDatabaseGetArgs>());
            set => _userDatabases = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SchemeState()
        {
        }
        public static new SchemeState Empty => new SchemeState();
    }
}
