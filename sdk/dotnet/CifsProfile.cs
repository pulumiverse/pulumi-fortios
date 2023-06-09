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
    /// Configure CIFS profile. Applies to FortiOS Version `6.2.4,6.2.6,6.4.0,6.4.1`.
    /// 
    /// ## Import
    /// 
    /// Cifs Profile can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/cifsProfile:CifsProfile labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/cifsProfile:CifsProfile labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/cifsProfile:CifsProfile")]
    public partial class CifsProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Domain for which to decrypt CIFS traffic.
        /// </summary>
        [Output("domainController")]
        public Output<string> DomainController { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// File filter. The structure of `file_filter` block is documented below.
        /// </summary>
        [Output("fileFilter")]
        public Output<Outputs.CifsProfileFileFilter> FileFilter { get; private set; } = null!;

        /// <summary>
        /// Profile name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
        /// </summary>
        [Output("serverCredentialType")]
        public Output<string> ServerCredentialType { get; private set; } = null!;

        /// <summary>
        /// Server keytab. The structure of `server_keytab` block is documented below.
        /// </summary>
        [Output("serverKeytabs")]
        public Output<ImmutableArray<Outputs.CifsProfileServerKeytab>> ServerKeytabs { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a CifsProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CifsProfile(string name, CifsProfileArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/cifsProfile:CifsProfile", name, args ?? new CifsProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CifsProfile(string name, Input<string> id, CifsProfileState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/cifsProfile:CifsProfile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CifsProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CifsProfile Get(string name, Input<string> id, CifsProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new CifsProfile(name, id, state, options);
        }
    }

    public sealed class CifsProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Domain for which to decrypt CIFS traffic.
        /// </summary>
        [Input("domainController")]
        public Input<string>? DomainController { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// File filter. The structure of `file_filter` block is documented below.
        /// </summary>
        [Input("fileFilter")]
        public Input<Inputs.CifsProfileFileFilterArgs>? FileFilter { get; set; }

        /// <summary>
        /// Profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
        /// </summary>
        [Input("serverCredentialType")]
        public Input<string>? ServerCredentialType { get; set; }

        [Input("serverKeytabs")]
        private InputList<Inputs.CifsProfileServerKeytabArgs>? _serverKeytabs;

        /// <summary>
        /// Server keytab. The structure of `server_keytab` block is documented below.
        /// </summary>
        public InputList<Inputs.CifsProfileServerKeytabArgs> ServerKeytabs
        {
            get => _serverKeytabs ?? (_serverKeytabs = new InputList<Inputs.CifsProfileServerKeytabArgs>());
            set => _serverKeytabs = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public CifsProfileArgs()
        {
        }
        public static new CifsProfileArgs Empty => new CifsProfileArgs();
    }

    public sealed class CifsProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Domain for which to decrypt CIFS traffic.
        /// </summary>
        [Input("domainController")]
        public Input<string>? DomainController { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// File filter. The structure of `file_filter` block is documented below.
        /// </summary>
        [Input("fileFilter")]
        public Input<Inputs.CifsProfileFileFilterGetArgs>? FileFilter { get; set; }

        /// <summary>
        /// Profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
        /// </summary>
        [Input("serverCredentialType")]
        public Input<string>? ServerCredentialType { get; set; }

        [Input("serverKeytabs")]
        private InputList<Inputs.CifsProfileServerKeytabGetArgs>? _serverKeytabs;

        /// <summary>
        /// Server keytab. The structure of `server_keytab` block is documented below.
        /// </summary>
        public InputList<Inputs.CifsProfileServerKeytabGetArgs> ServerKeytabs
        {
            get => _serverKeytabs ?? (_serverKeytabs = new InputList<Inputs.CifsProfileServerKeytabGetArgs>());
            set => _serverKeytabs = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public CifsProfileState()
        {
        }
        public static new CifsProfileState Empty => new CifsProfileState();
    }
}
