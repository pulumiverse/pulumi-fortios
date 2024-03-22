// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Hotspot20
{
    /// <summary>
    /// Configure online sign up (OSU) provider list.
    /// 
    /// ## Import
    /// 
    /// WirelessControllerHotspot20 H2QpOsuProvider can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:wirelesscontroller/hotspot20/h2qposuprovider:H2qposuprovider labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:wirelesscontroller/hotspot20/h2qposuprovider:H2qposuprovider labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:wirelesscontroller/hotspot20/h2qposuprovider:H2qposuprovider")]
    public partial class H2qposuprovider : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// OSU provider friendly name. The structure of `friendly_name` block is documented below.
        /// </summary>
        [Output("friendlyNames")]
        public Output<ImmutableArray<Outputs.H2qposuproviderFriendlyName>> FriendlyNames { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// OSU provider icon.
        /// </summary>
        [Output("icon")]
        public Output<string> Icon { get; private set; } = null!;

        /// <summary>
        /// OSU provider ID.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// OSU method list. Valid values: `oma-dm`, `soap-xml-spp`, `reserved`.
        /// </summary>
        [Output("osuMethod")]
        public Output<string> OsuMethod { get; private set; } = null!;

        /// <summary>
        /// OSU NAI.
        /// </summary>
        [Output("osuNai")]
        public Output<string> OsuNai { get; private set; } = null!;

        /// <summary>
        /// Server URI.
        /// </summary>
        [Output("serverUri")]
        public Output<string> ServerUri { get; private set; } = null!;

        /// <summary>
        /// OSU service name. The structure of `service_description` block is documented below.
        /// </summary>
        [Output("serviceDescriptions")]
        public Output<ImmutableArray<Outputs.H2qposuproviderServiceDescription>> ServiceDescriptions { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a H2qposuprovider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public H2qposuprovider(string name, H2qposuproviderArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/hotspot20/h2qposuprovider:H2qposuprovider", name, args ?? new H2qposuproviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private H2qposuprovider(string name, Input<string> id, H2qposuproviderState? state = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/hotspot20/h2qposuprovider:H2qposuprovider", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing H2qposuprovider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static H2qposuprovider Get(string name, Input<string> id, H2qposuproviderState? state = null, CustomResourceOptions? options = null)
        {
            return new H2qposuprovider(name, id, state, options);
        }
    }

    public sealed class H2qposuproviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("friendlyNames")]
        private InputList<Inputs.H2qposuproviderFriendlyNameArgs>? _friendlyNames;

        /// <summary>
        /// OSU provider friendly name. The structure of `friendly_name` block is documented below.
        /// </summary>
        public InputList<Inputs.H2qposuproviderFriendlyNameArgs> FriendlyNames
        {
            get => _friendlyNames ?? (_friendlyNames = new InputList<Inputs.H2qposuproviderFriendlyNameArgs>());
            set => _friendlyNames = value;
        }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// OSU provider icon.
        /// </summary>
        [Input("icon")]
        public Input<string>? Icon { get; set; }

        /// <summary>
        /// OSU provider ID.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// OSU method list. Valid values: `oma-dm`, `soap-xml-spp`, `reserved`.
        /// </summary>
        [Input("osuMethod")]
        public Input<string>? OsuMethod { get; set; }

        /// <summary>
        /// OSU NAI.
        /// </summary>
        [Input("osuNai")]
        public Input<string>? OsuNai { get; set; }

        /// <summary>
        /// Server URI.
        /// </summary>
        [Input("serverUri")]
        public Input<string>? ServerUri { get; set; }

        [Input("serviceDescriptions")]
        private InputList<Inputs.H2qposuproviderServiceDescriptionArgs>? _serviceDescriptions;

        /// <summary>
        /// OSU service name. The structure of `service_description` block is documented below.
        /// </summary>
        public InputList<Inputs.H2qposuproviderServiceDescriptionArgs> ServiceDescriptions
        {
            get => _serviceDescriptions ?? (_serviceDescriptions = new InputList<Inputs.H2qposuproviderServiceDescriptionArgs>());
            set => _serviceDescriptions = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public H2qposuproviderArgs()
        {
        }
        public static new H2qposuproviderArgs Empty => new H2qposuproviderArgs();
    }

    public sealed class H2qposuproviderState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("friendlyNames")]
        private InputList<Inputs.H2qposuproviderFriendlyNameGetArgs>? _friendlyNames;

        /// <summary>
        /// OSU provider friendly name. The structure of `friendly_name` block is documented below.
        /// </summary>
        public InputList<Inputs.H2qposuproviderFriendlyNameGetArgs> FriendlyNames
        {
            get => _friendlyNames ?? (_friendlyNames = new InputList<Inputs.H2qposuproviderFriendlyNameGetArgs>());
            set => _friendlyNames = value;
        }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// OSU provider icon.
        /// </summary>
        [Input("icon")]
        public Input<string>? Icon { get; set; }

        /// <summary>
        /// OSU provider ID.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// OSU method list. Valid values: `oma-dm`, `soap-xml-spp`, `reserved`.
        /// </summary>
        [Input("osuMethod")]
        public Input<string>? OsuMethod { get; set; }

        /// <summary>
        /// OSU NAI.
        /// </summary>
        [Input("osuNai")]
        public Input<string>? OsuNai { get; set; }

        /// <summary>
        /// Server URI.
        /// </summary>
        [Input("serverUri")]
        public Input<string>? ServerUri { get; set; }

        [Input("serviceDescriptions")]
        private InputList<Inputs.H2qposuproviderServiceDescriptionGetArgs>? _serviceDescriptions;

        /// <summary>
        /// OSU service name. The structure of `service_description` block is documented below.
        /// </summary>
        public InputList<Inputs.H2qposuproviderServiceDescriptionGetArgs> ServiceDescriptions
        {
            get => _serviceDescriptions ?? (_serviceDescriptions = new InputList<Inputs.H2qposuproviderServiceDescriptionGetArgs>());
            set => _serviceDescriptions = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public H2qposuproviderState()
        {
        }
        public static new H2qposuproviderState Empty => new H2qposuproviderState();
    }
}
