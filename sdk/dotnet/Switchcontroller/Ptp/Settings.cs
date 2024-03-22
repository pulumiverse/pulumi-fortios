// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Ptp
{
    /// <summary>
    /// Global PTP settings. Applies to FortiOS Version `6.4.2,6.4.10,6.4.11,6.4.12,6.4.13,6.4.14,7.0.0,7.0.1,7.0.2,7.0.3,7.0.4,7.0.5,7.0.6,7.0.7,7.0.8,7.0.9,7.0.10,7.0.11,7.0.12,7.0.13,7.2.0,7.2.1,7.2.2,7.2.3,7.2.4,7.2.6,7.4.0`.
    /// 
    /// ## Import
    /// 
    /// SwitchControllerPtp Settings can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/ptp/settings:Settings labelname SwitchControllerPtpSettings
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/ptp/settings:Settings labelname SwitchControllerPtpSettings
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:switchcontroller/ptp/settings:Settings")]
    public partial class Settings : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable PTP mode. Valid values: `disable`, `transparent-e2e`, `transparent-p2p`.
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Settings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Settings(string name, SettingsArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/ptp/settings:Settings", name, args ?? new SettingsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Settings(string name, Input<string> id, SettingsState? state = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/ptp/settings:Settings", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Settings resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Settings Get(string name, Input<string> id, SettingsState? state = null, CustomResourceOptions? options = null)
        {
            return new Settings(name, id, state, options);
        }
    }

    public sealed class SettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable PTP mode. Valid values: `disable`, `transparent-e2e`, `transparent-p2p`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SettingsArgs()
        {
        }
        public static new SettingsArgs Empty => new SettingsArgs();
    }

    public sealed class SettingsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable PTP mode. Valid values: `disable`, `transparent-e2e`, `transparent-p2p`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SettingsState()
        {
        }
        public static new SettingsState Empty => new SettingsState();
    }
}
