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
    /// Configure FSSO active directory servers for polling mode.
    /// 
    /// ## Import
    /// 
    /// User FssoPolling can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/fssopolling:Fssopolling labelname {{fosid}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/fssopolling:Fssopolling labelname {{fosid}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:user/fssopolling:Fssopolling")]
    public partial class Fssopolling : global::Pulumi.CustomResource
    {
        /// <summary>
        /// LDAP Group Info. The structure of `adgrp` block is documented below.
        /// </summary>
        [Output("adgrps")]
        public Output<ImmutableArray<Outputs.FssopollingAdgrp>> Adgrps { get; private set; } = null!;

        /// <summary>
        /// Default domain managed by this Active Directory server.
        /// </summary>
        [Output("defaultDomain")]
        public Output<string> DefaultDomain { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Active Directory server ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// LDAP server name used in LDAP connection strings.
        /// </summary>
        [Output("ldapServer")]
        public Output<string> LdapServer { get; private set; } = null!;

        /// <summary>
        /// Number of hours of logon history to keep, 0 means keep all history.
        /// </summary>
        [Output("logonHistory")]
        public Output<int> LogonHistory { get; private set; } = null!;

        /// <summary>
        /// Password required to log into this Active Directory server
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Polling frequency (every 1 to 30 seconds).
        /// </summary>
        [Output("pollingFrequency")]
        public Output<int> PollingFrequency { get; private set; } = null!;

        /// <summary>
        /// Port to communicate with this Active Directory server.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Host name or IP address of the Active Directory server.
        /// </summary>
        [Output("server")]
        public Output<string> Server { get; private set; } = null!;

        /// <summary>
        /// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("smbNtlmv1Auth")]
        public Output<string> SmbNtlmv1Auth { get; private set; } = null!;

        /// <summary>
        /// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("smbv1")]
        public Output<string> Smbv1 { get; private set; } = null!;

        /// <summary>
        /// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// User name required to log into this Active Directory server.
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Fssopolling resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Fssopolling(string name, FssopollingArgs args, CustomResourceOptions? options = null)
            : base("fortios:user/fssopolling:Fssopolling", name, args ?? new FssopollingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Fssopolling(string name, Input<string> id, FssopollingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:user/fssopolling:Fssopolling", name, state, MakeResourceOptions(options, id))
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
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Fssopolling resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Fssopolling Get(string name, Input<string> id, FssopollingState? state = null, CustomResourceOptions? options = null)
        {
            return new Fssopolling(name, id, state, options);
        }
    }

    public sealed class FssopollingArgs : global::Pulumi.ResourceArgs
    {
        [Input("adgrps")]
        private InputList<Inputs.FssopollingAdgrpArgs>? _adgrps;

        /// <summary>
        /// LDAP Group Info. The structure of `adgrp` block is documented below.
        /// </summary>
        public InputList<Inputs.FssopollingAdgrpArgs> Adgrps
        {
            get => _adgrps ?? (_adgrps = new InputList<Inputs.FssopollingAdgrpArgs>());
            set => _adgrps = value;
        }

        /// <summary>
        /// Default domain managed by this Active Directory server.
        /// </summary>
        [Input("defaultDomain")]
        public Input<string>? DefaultDomain { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Active Directory server ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// LDAP server name used in LDAP connection strings.
        /// </summary>
        [Input("ldapServer", required: true)]
        public Input<string> LdapServer { get; set; } = null!;

        /// <summary>
        /// Number of hours of logon history to keep, 0 means keep all history.
        /// </summary>
        [Input("logonHistory")]
        public Input<int>? LogonHistory { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password required to log into this Active Directory server
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

        /// <summary>
        /// Polling frequency (every 1 to 30 seconds).
        /// </summary>
        [Input("pollingFrequency")]
        public Input<int>? PollingFrequency { get; set; }

        /// <summary>
        /// Port to communicate with this Active Directory server.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Host name or IP address of the Active Directory server.
        /// </summary>
        [Input("server", required: true)]
        public Input<string> Server { get; set; } = null!;

        /// <summary>
        /// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("smbNtlmv1Auth")]
        public Input<string>? SmbNtlmv1Auth { get; set; }

        /// <summary>
        /// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("smbv1")]
        public Input<string>? Smbv1 { get; set; }

        /// <summary>
        /// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// User name required to log into this Active Directory server.
        /// </summary>
        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FssopollingArgs()
        {
        }
        public static new FssopollingArgs Empty => new FssopollingArgs();
    }

    public sealed class FssopollingState : global::Pulumi.ResourceArgs
    {
        [Input("adgrps")]
        private InputList<Inputs.FssopollingAdgrpGetArgs>? _adgrps;

        /// <summary>
        /// LDAP Group Info. The structure of `adgrp` block is documented below.
        /// </summary>
        public InputList<Inputs.FssopollingAdgrpGetArgs> Adgrps
        {
            get => _adgrps ?? (_adgrps = new InputList<Inputs.FssopollingAdgrpGetArgs>());
            set => _adgrps = value;
        }

        /// <summary>
        /// Default domain managed by this Active Directory server.
        /// </summary>
        [Input("defaultDomain")]
        public Input<string>? DefaultDomain { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Active Directory server ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// LDAP server name used in LDAP connection strings.
        /// </summary>
        [Input("ldapServer")]
        public Input<string>? LdapServer { get; set; }

        /// <summary>
        /// Number of hours of logon history to keep, 0 means keep all history.
        /// </summary>
        [Input("logonHistory")]
        public Input<int>? LogonHistory { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password required to log into this Active Directory server
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

        /// <summary>
        /// Polling frequency (every 1 to 30 seconds).
        /// </summary>
        [Input("pollingFrequency")]
        public Input<int>? PollingFrequency { get; set; }

        /// <summary>
        /// Port to communicate with this Active Directory server.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Host name or IP address of the Active Directory server.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("smbNtlmv1Auth")]
        public Input<string>? SmbNtlmv1Auth { get; set; }

        /// <summary>
        /// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("smbv1")]
        public Input<string>? Smbv1 { get; set; }

        /// <summary>
        /// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// User name required to log into this Active Directory server.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FssopollingState()
        {
        }
        public static new FssopollingState Empty => new FssopollingState();
    }
}
