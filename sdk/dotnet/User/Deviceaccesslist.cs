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
    /// Configure device access control lists. Applies to FortiOS Version `&lt;= 6.2.0`.
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
    ///     var trname = new Fortios.User.Deviceaccesslist("trname", new()
    ///     {
    ///         DefaultAction = "accept",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// User DeviceAccessList can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/deviceaccesslist:Deviceaccesslist labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/deviceaccesslist:Deviceaccesslist labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:user/deviceaccesslist:Deviceaccesslist")]
    public partial class Deviceaccesslist : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Accept or deny unknown/unspecified devices. Valid values: `accept`, `deny`.
        /// </summary>
        [Output("defaultAction")]
        public Output<string> DefaultAction { get; private set; } = null!;

        /// <summary>
        /// Device list. The structure of `device_list` block is documented below.
        /// </summary>
        [Output("deviceLists")]
        public Output<ImmutableArray<Outputs.DeviceaccesslistDeviceList>> DeviceLists { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Device access list name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Deviceaccesslist resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Deviceaccesslist(string name, DeviceaccesslistArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:user/deviceaccesslist:Deviceaccesslist", name, args ?? new DeviceaccesslistArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Deviceaccesslist(string name, Input<string> id, DeviceaccesslistState? state = null, CustomResourceOptions? options = null)
            : base("fortios:user/deviceaccesslist:Deviceaccesslist", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Deviceaccesslist resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Deviceaccesslist Get(string name, Input<string> id, DeviceaccesslistState? state = null, CustomResourceOptions? options = null)
        {
            return new Deviceaccesslist(name, id, state, options);
        }
    }

    public sealed class DeviceaccesslistArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Accept or deny unknown/unspecified devices. Valid values: `accept`, `deny`.
        /// </summary>
        [Input("defaultAction")]
        public Input<string>? DefaultAction { get; set; }

        [Input("deviceLists")]
        private InputList<Inputs.DeviceaccesslistDeviceListArgs>? _deviceLists;

        /// <summary>
        /// Device list. The structure of `device_list` block is documented below.
        /// </summary>
        public InputList<Inputs.DeviceaccesslistDeviceListArgs> DeviceLists
        {
            get => _deviceLists ?? (_deviceLists = new InputList<Inputs.DeviceaccesslistDeviceListArgs>());
            set => _deviceLists = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Device access list name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DeviceaccesslistArgs()
        {
        }
        public static new DeviceaccesslistArgs Empty => new DeviceaccesslistArgs();
    }

    public sealed class DeviceaccesslistState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Accept or deny unknown/unspecified devices. Valid values: `accept`, `deny`.
        /// </summary>
        [Input("defaultAction")]
        public Input<string>? DefaultAction { get; set; }

        [Input("deviceLists")]
        private InputList<Inputs.DeviceaccesslistDeviceListGetArgs>? _deviceLists;

        /// <summary>
        /// Device list. The structure of `device_list` block is documented below.
        /// </summary>
        public InputList<Inputs.DeviceaccesslistDeviceListGetArgs> DeviceLists
        {
            get => _deviceLists ?? (_deviceLists = new InputList<Inputs.DeviceaccesslistDeviceListGetArgs>());
            set => _deviceLists = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Device access list name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DeviceaccesslistState()
        {
        }
        public static new DeviceaccesslistState Empty => new DeviceaccesslistState();
    }
}
