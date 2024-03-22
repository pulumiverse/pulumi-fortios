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
    /// Configure Internet Services Extension.
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
    ///     var trname = new Fortios.Firewall.Internetserviceextension("trname", new()
    ///     {
    ///         Comment = "EIWE",
    ///         Fosid = 65536,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Firewall InternetServiceExtension can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/internetserviceextension:Internetserviceextension labelname {{fosid}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/internetserviceextension:Internetserviceextension labelname {{fosid}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/internetserviceextension:Internetserviceextension")]
    public partial class Internetserviceextension : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Disable entries in the Internet Service database. The structure of `disable_entry` block is documented below.
        /// </summary>
        [Output("disableEntries")]
        public Output<ImmutableArray<Outputs.InternetserviceextensionDisableEntry>> DisableEntries { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
        /// </summary>
        [Output("entries")]
        public Output<ImmutableArray<Outputs.InternetserviceextensionEntry>> Entries { get; private set; } = null!;

        /// <summary>
        /// Internet Service ID in the Internet Service database.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Internetserviceextension resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Internetserviceextension(string name, InternetserviceextensionArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/internetserviceextension:Internetserviceextension", name, args ?? new InternetserviceextensionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Internetserviceextension(string name, Input<string> id, InternetserviceextensionState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/internetserviceextension:Internetserviceextension", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Internetserviceextension resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Internetserviceextension Get(string name, Input<string> id, InternetserviceextensionState? state = null, CustomResourceOptions? options = null)
        {
            return new Internetserviceextension(name, id, state, options);
        }
    }

    public sealed class InternetserviceextensionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("disableEntries")]
        private InputList<Inputs.InternetserviceextensionDisableEntryArgs>? _disableEntries;

        /// <summary>
        /// Disable entries in the Internet Service database. The structure of `disable_entry` block is documented below.
        /// </summary>
        public InputList<Inputs.InternetserviceextensionDisableEntryArgs> DisableEntries
        {
            get => _disableEntries ?? (_disableEntries = new InputList<Inputs.InternetserviceextensionDisableEntryArgs>());
            set => _disableEntries = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("entries")]
        private InputList<Inputs.InternetserviceextensionEntryArgs>? _entries;

        /// <summary>
        /// Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
        /// </summary>
        public InputList<Inputs.InternetserviceextensionEntryArgs> Entries
        {
            get => _entries ?? (_entries = new InputList<Inputs.InternetserviceextensionEntryArgs>());
            set => _entries = value;
        }

        /// <summary>
        /// Internet Service ID in the Internet Service database.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public InternetserviceextensionArgs()
        {
        }
        public static new InternetserviceextensionArgs Empty => new InternetserviceextensionArgs();
    }

    public sealed class InternetserviceextensionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("disableEntries")]
        private InputList<Inputs.InternetserviceextensionDisableEntryGetArgs>? _disableEntries;

        /// <summary>
        /// Disable entries in the Internet Service database. The structure of `disable_entry` block is documented below.
        /// </summary>
        public InputList<Inputs.InternetserviceextensionDisableEntryGetArgs> DisableEntries
        {
            get => _disableEntries ?? (_disableEntries = new InputList<Inputs.InternetserviceextensionDisableEntryGetArgs>());
            set => _disableEntries = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("entries")]
        private InputList<Inputs.InternetserviceextensionEntryGetArgs>? _entries;

        /// <summary>
        /// Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
        /// </summary>
        public InputList<Inputs.InternetserviceextensionEntryGetArgs> Entries
        {
            get => _entries ?? (_entries = new InputList<Inputs.InternetserviceextensionEntryGetArgs>());
            set => _entries = value;
        }

        /// <summary>
        /// Internet Service ID in the Internet Service database.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public InternetserviceextensionState()
        {
        }
        public static new InternetserviceextensionState Empty => new InternetserviceextensionState();
    }
}
