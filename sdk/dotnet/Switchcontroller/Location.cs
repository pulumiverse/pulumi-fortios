// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller
{
    /// <summary>
    /// Configure FortiSwitch location services. Applies to FortiOS Version `&gt;= 6.2.4`.
    /// 
    /// ## Import
    /// 
    /// SwitchController Location can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/location:Location labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/location:Location labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:switchcontroller/location:Location")]
    public partial class Location : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Configure location civic address. The structure of `address_civic` block is documented below.
        /// </summary>
        [Output("addressCivic")]
        public Output<Outputs.LocationAddressCivic> AddressCivic { get; private set; } = null!;

        /// <summary>
        /// Configure location GPS coordinates. The structure of `coordinates` block is documented below.
        /// </summary>
        [Output("coordinates")]
        public Output<Outputs.LocationCoordinates> Coordinates { get; private set; } = null!;

        /// <summary>
        /// Configure location ELIN number. The structure of `elin_number` block is documented below.
        /// </summary>
        [Output("elinNumber")]
        public Output<Outputs.LocationElinNumber> ElinNumber { get; private set; } = null!;

        /// <summary>
        /// Unique location item name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Location resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Location(string name, LocationArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/location:Location", name, args ?? new LocationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Location(string name, Input<string> id, LocationState? state = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/location:Location", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Location resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Location Get(string name, Input<string> id, LocationState? state = null, CustomResourceOptions? options = null)
        {
            return new Location(name, id, state, options);
        }
    }

    public sealed class LocationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configure location civic address. The structure of `address_civic` block is documented below.
        /// </summary>
        [Input("addressCivic")]
        public Input<Inputs.LocationAddressCivicArgs>? AddressCivic { get; set; }

        /// <summary>
        /// Configure location GPS coordinates. The structure of `coordinates` block is documented below.
        /// </summary>
        [Input("coordinates")]
        public Input<Inputs.LocationCoordinatesArgs>? Coordinates { get; set; }

        /// <summary>
        /// Configure location ELIN number. The structure of `elin_number` block is documented below.
        /// </summary>
        [Input("elinNumber")]
        public Input<Inputs.LocationElinNumberArgs>? ElinNumber { get; set; }

        /// <summary>
        /// Unique location item name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public LocationArgs()
        {
        }
        public static new LocationArgs Empty => new LocationArgs();
    }

    public sealed class LocationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configure location civic address. The structure of `address_civic` block is documented below.
        /// </summary>
        [Input("addressCivic")]
        public Input<Inputs.LocationAddressCivicGetArgs>? AddressCivic { get; set; }

        /// <summary>
        /// Configure location GPS coordinates. The structure of `coordinates` block is documented below.
        /// </summary>
        [Input("coordinates")]
        public Input<Inputs.LocationCoordinatesGetArgs>? Coordinates { get; set; }

        /// <summary>
        /// Configure location ELIN number. The structure of `elin_number` block is documented below.
        /// </summary>
        [Input("elinNumber")]
        public Input<Inputs.LocationElinNumberGetArgs>? ElinNumber { get; set; }

        /// <summary>
        /// Unique location item name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public LocationState()
        {
        }
        public static new LocationState Empty => new LocationState();
    }
}
