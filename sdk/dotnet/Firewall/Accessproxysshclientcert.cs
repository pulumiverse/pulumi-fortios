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
    /// Configure Access Proxy SSH client certificate. Applies to FortiOS Version `&gt;= 7.0.1`.
    /// 
    /// ## Import
    /// 
    /// Firewall AccessProxySshClientCert can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/accessproxysshclientcert:Accessproxysshclientcert labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/accessproxysshclientcert:Accessproxysshclientcert labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/accessproxysshclientcert:Accessproxysshclientcert")]
    public partial class Accessproxysshclientcert : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the SSH server public key authentication CA.
        /// </summary>
        [Output("authCa")]
        public Output<string> AuthCa { get; private set; } = null!;

        /// <summary>
        /// Configure certificate extension for user certificate. The structure of `cert_extension` block is documented below.
        /// </summary>
        [Output("certExtensions")]
        public Output<ImmutableArray<Outputs.AccessproxysshclientcertCertExtension>> CertExtensions { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// SSH client certificate name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable appending permit-agent-forwarding certificate extension. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("permitAgentForwarding")]
        public Output<string> PermitAgentForwarding { get; private set; } = null!;

        /// <summary>
        /// Enable/disable appending permit-port-forwarding certificate extension. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("permitPortForwarding")]
        public Output<string> PermitPortForwarding { get; private set; } = null!;

        /// <summary>
        /// Enable/disable appending permit-pty certificate extension. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("permitPty")]
        public Output<string> PermitPty { get; private set; } = null!;

        /// <summary>
        /// Enable/disable appending permit-user-rc certificate extension. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("permitUserRc")]
        public Output<string> PermitUserRc { get; private set; } = null!;

        /// <summary>
        /// Enable/disable appending permit-x11-forwarding certificate extension. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("permitX11Forwarding")]
        public Output<string> PermitX11Forwarding { get; private set; } = null!;

        /// <summary>
        /// Enable/disable appending source-address certificate critical option. This option ensure certificate only accepted from FortiGate source address. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("sourceAddress")]
        public Output<string> SourceAddress { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Accessproxysshclientcert resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Accessproxysshclientcert(string name, AccessproxysshclientcertArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/accessproxysshclientcert:Accessproxysshclientcert", name, args ?? new AccessproxysshclientcertArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Accessproxysshclientcert(string name, Input<string> id, AccessproxysshclientcertState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/accessproxysshclientcert:Accessproxysshclientcert", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Accessproxysshclientcert resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Accessproxysshclientcert Get(string name, Input<string> id, AccessproxysshclientcertState? state = null, CustomResourceOptions? options = null)
        {
            return new Accessproxysshclientcert(name, id, state, options);
        }
    }

    public sealed class AccessproxysshclientcertArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the SSH server public key authentication CA.
        /// </summary>
        [Input("authCa")]
        public Input<string>? AuthCa { get; set; }

        [Input("certExtensions")]
        private InputList<Inputs.AccessproxysshclientcertCertExtensionArgs>? _certExtensions;

        /// <summary>
        /// Configure certificate extension for user certificate. The structure of `cert_extension` block is documented below.
        /// </summary>
        public InputList<Inputs.AccessproxysshclientcertCertExtensionArgs> CertExtensions
        {
            get => _certExtensions ?? (_certExtensions = new InputList<Inputs.AccessproxysshclientcertCertExtensionArgs>());
            set => _certExtensions = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// SSH client certificate name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable appending permit-agent-forwarding certificate extension. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("permitAgentForwarding")]
        public Input<string>? PermitAgentForwarding { get; set; }

        /// <summary>
        /// Enable/disable appending permit-port-forwarding certificate extension. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("permitPortForwarding")]
        public Input<string>? PermitPortForwarding { get; set; }

        /// <summary>
        /// Enable/disable appending permit-pty certificate extension. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("permitPty")]
        public Input<string>? PermitPty { get; set; }

        /// <summary>
        /// Enable/disable appending permit-user-rc certificate extension. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("permitUserRc")]
        public Input<string>? PermitUserRc { get; set; }

        /// <summary>
        /// Enable/disable appending permit-x11-forwarding certificate extension. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("permitX11Forwarding")]
        public Input<string>? PermitX11Forwarding { get; set; }

        /// <summary>
        /// Enable/disable appending source-address certificate critical option. This option ensure certificate only accepted from FortiGate source address. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sourceAddress")]
        public Input<string>? SourceAddress { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AccessproxysshclientcertArgs()
        {
        }
        public static new AccessproxysshclientcertArgs Empty => new AccessproxysshclientcertArgs();
    }

    public sealed class AccessproxysshclientcertState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the SSH server public key authentication CA.
        /// </summary>
        [Input("authCa")]
        public Input<string>? AuthCa { get; set; }

        [Input("certExtensions")]
        private InputList<Inputs.AccessproxysshclientcertCertExtensionGetArgs>? _certExtensions;

        /// <summary>
        /// Configure certificate extension for user certificate. The structure of `cert_extension` block is documented below.
        /// </summary>
        public InputList<Inputs.AccessproxysshclientcertCertExtensionGetArgs> CertExtensions
        {
            get => _certExtensions ?? (_certExtensions = new InputList<Inputs.AccessproxysshclientcertCertExtensionGetArgs>());
            set => _certExtensions = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// SSH client certificate name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable appending permit-agent-forwarding certificate extension. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("permitAgentForwarding")]
        public Input<string>? PermitAgentForwarding { get; set; }

        /// <summary>
        /// Enable/disable appending permit-port-forwarding certificate extension. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("permitPortForwarding")]
        public Input<string>? PermitPortForwarding { get; set; }

        /// <summary>
        /// Enable/disable appending permit-pty certificate extension. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("permitPty")]
        public Input<string>? PermitPty { get; set; }

        /// <summary>
        /// Enable/disable appending permit-user-rc certificate extension. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("permitUserRc")]
        public Input<string>? PermitUserRc { get; set; }

        /// <summary>
        /// Enable/disable appending permit-x11-forwarding certificate extension. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("permitX11Forwarding")]
        public Input<string>? PermitX11Forwarding { get; set; }

        /// <summary>
        /// Enable/disable appending source-address certificate critical option. This option ensure certificate only accepted from FortiGate source address. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sourceAddress")]
        public Input<string>? SourceAddress { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AccessproxysshclientcertState()
        {
        }
        public static new AccessproxysshclientcertState Empty => new AccessproxysshclientcertState();
    }
}
