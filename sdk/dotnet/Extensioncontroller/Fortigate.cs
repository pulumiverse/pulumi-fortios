// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Extensioncontroller
{
    /// <summary>
    /// FortiGate controller configuration. Applies to FortiOS Version `&gt;= 7.2.1`.
    /// 
    /// ## Import
    /// 
    /// ExtensionController Fortigate can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:extensioncontroller/fortigate:Fortigate labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:extensioncontroller/fortigate:Fortigate labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:extensioncontroller/fortigate:Fortigate")]
    public partial class Fortigate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable FortiGate administration.
        /// </summary>
        [Output("authorized")]
        public Output<string> Authorized { get; private set; } = null!;

        /// <summary>
        /// Description.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// device-id
        /// </summary>
        [Output("deviceId")]
        public Output<int> DeviceId { get; private set; } = null!;

        /// <summary>
        /// FortiGate serial number.
        /// </summary>
        [Output("fosid")]
        public Output<string> Fosid { get; private set; } = null!;

        /// <summary>
        /// FortiGate hostname.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// FortiGate entry name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// FortiGate profile configuration.
        /// </summary>
        [Output("profile")]
        public Output<string> Profile { get; private set; } = null!;

        /// <summary>
        /// VDOM.
        /// </summary>
        [Output("vdom")]
        public Output<int> Vdom { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Fortigate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Fortigate(string name, FortigateArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:extensioncontroller/fortigate:Fortigate", name, args ?? new FortigateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Fortigate(string name, Input<string> id, FortigateState? state = null, CustomResourceOptions? options = null)
            : base("fortios:extensioncontroller/fortigate:Fortigate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Fortigate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Fortigate Get(string name, Input<string> id, FortigateState? state = null, CustomResourceOptions? options = null)
        {
            return new Fortigate(name, id, state, options);
        }
    }

    public sealed class FortigateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable FortiGate administration.
        /// </summary>
        [Input("authorized")]
        public Input<string>? Authorized { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// device-id
        /// </summary>
        [Input("deviceId")]
        public Input<int>? DeviceId { get; set; }

        /// <summary>
        /// FortiGate serial number.
        /// </summary>
        [Input("fosid")]
        public Input<string>? Fosid { get; set; }

        /// <summary>
        /// FortiGate hostname.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// FortiGate entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// FortiGate profile configuration.
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        /// <summary>
        /// VDOM.
        /// </summary>
        [Input("vdom")]
        public Input<int>? Vdom { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FortigateArgs()
        {
        }
        public static new FortigateArgs Empty => new FortigateArgs();
    }

    public sealed class FortigateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable FortiGate administration.
        /// </summary>
        [Input("authorized")]
        public Input<string>? Authorized { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// device-id
        /// </summary>
        [Input("deviceId")]
        public Input<int>? DeviceId { get; set; }

        /// <summary>
        /// FortiGate serial number.
        /// </summary>
        [Input("fosid")]
        public Input<string>? Fosid { get; set; }

        /// <summary>
        /// FortiGate hostname.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// FortiGate entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// FortiGate profile configuration.
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        /// <summary>
        /// VDOM.
        /// </summary>
        [Input("vdom")]
        public Input<int>? Vdom { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FortigateState()
        {
        }
        public static new FortigateState Empty => new FortigateState();
    }
}
