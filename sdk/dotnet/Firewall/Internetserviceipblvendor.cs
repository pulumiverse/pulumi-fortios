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
    /// IP blacklist vendor. Applies to FortiOS Version `&gt;= 6.2.4`.
    /// 
    /// ## Import
    /// 
    /// Firewall InternetServiceIpblVendor can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/internetserviceipblvendor:Internetserviceipblvendor labelname {{fosid}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/internetserviceipblvendor:Internetserviceipblvendor labelname {{fosid}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/internetserviceipblvendor:Internetserviceipblvendor")]
    public partial class Internetserviceipblvendor : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IP blacklist vendor ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// IP blacklist vendor name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Internetserviceipblvendor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Internetserviceipblvendor(string name, InternetserviceipblvendorArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/internetserviceipblvendor:Internetserviceipblvendor", name, args ?? new InternetserviceipblvendorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Internetserviceipblvendor(string name, Input<string> id, InternetserviceipblvendorState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/internetserviceipblvendor:Internetserviceipblvendor", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Internetserviceipblvendor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Internetserviceipblvendor Get(string name, Input<string> id, InternetserviceipblvendorState? state = null, CustomResourceOptions? options = null)
        {
            return new Internetserviceipblvendor(name, id, state, options);
        }
    }

    public sealed class InternetserviceipblvendorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP blacklist vendor ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// IP blacklist vendor name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public InternetserviceipblvendorArgs()
        {
        }
        public static new InternetserviceipblvendorArgs Empty => new InternetserviceipblvendorArgs();
    }

    public sealed class InternetserviceipblvendorState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP blacklist vendor ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// IP blacklist vendor name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public InternetserviceipblvendorState()
        {
        }
        public static new InternetserviceipblvendorState Empty => new InternetserviceipblvendorState();
    }
}
