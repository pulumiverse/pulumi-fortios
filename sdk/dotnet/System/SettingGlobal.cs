// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System
{
    [FortiosResourceType("fortios:system/settingGlobal:SettingGlobal")]
    public partial class SettingGlobal : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable SCP over SSH
        /// </summary>
        [Output("adminScp")]
        public Output<string> AdminScp { get; private set; } = null!;

        /// <summary>
        /// Administrative access port for HTTPS.
        /// </summary>
        [Output("adminSport")]
        public Output<string> AdminSport { get; private set; } = null!;

        /// <summary>
        /// Administrative access port for SSH.
        /// </summary>
        [Output("adminSshPort")]
        public Output<string> AdminSshPort { get; private set; } = null!;

        /// <summary>
        /// Number of minutes before an idle administrator session time out.
        /// </summary>
        [Output("admintimeout")]
        public Output<string> Admintimeout { get; private set; } = null!;

        /// <summary>
        /// FortiGate unit's hostname.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// Number corresponding to your time zone from 00 to 86.
        /// </summary>
        [Output("timezone")]
        public Output<string> Timezone { get; private set; } = null!;


        /// <summary>
        /// Create a SettingGlobal resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SettingGlobal(string name, SettingGlobalArgs args, CustomResourceOptions? options = null)
            : base("fortios:system/settingGlobal:SettingGlobal", name, args ?? new SettingGlobalArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SettingGlobal(string name, Input<string> id, SettingGlobalState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/settingGlobal:SettingGlobal", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SettingGlobal resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SettingGlobal Get(string name, Input<string> id, SettingGlobalState? state = null, CustomResourceOptions? options = null)
        {
            return new SettingGlobal(name, id, state, options);
        }
    }

    public sealed class SettingGlobalArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable SCP over SSH
        /// </summary>
        [Input("adminScp")]
        public Input<string>? AdminScp { get; set; }

        /// <summary>
        /// Administrative access port for HTTPS.
        /// </summary>
        [Input("adminSport")]
        public Input<string>? AdminSport { get; set; }

        /// <summary>
        /// Administrative access port for SSH.
        /// </summary>
        [Input("adminSshPort")]
        public Input<string>? AdminSshPort { get; set; }

        /// <summary>
        /// Number of minutes before an idle administrator session time out.
        /// </summary>
        [Input("admintimeout")]
        public Input<string>? Admintimeout { get; set; }

        /// <summary>
        /// FortiGate unit's hostname.
        /// </summary>
        [Input("hostname", required: true)]
        public Input<string> Hostname { get; set; } = null!;

        /// <summary>
        /// Number corresponding to your time zone from 00 to 86.
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        public SettingGlobalArgs()
        {
        }
        public static new SettingGlobalArgs Empty => new SettingGlobalArgs();
    }

    public sealed class SettingGlobalState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable SCP over SSH
        /// </summary>
        [Input("adminScp")]
        public Input<string>? AdminScp { get; set; }

        /// <summary>
        /// Administrative access port for HTTPS.
        /// </summary>
        [Input("adminSport")]
        public Input<string>? AdminSport { get; set; }

        /// <summary>
        /// Administrative access port for SSH.
        /// </summary>
        [Input("adminSshPort")]
        public Input<string>? AdminSshPort { get; set; }

        /// <summary>
        /// Number of minutes before an idle administrator session time out.
        /// </summary>
        [Input("admintimeout")]
        public Input<string>? Admintimeout { get; set; }

        /// <summary>
        /// FortiGate unit's hostname.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// Number corresponding to your time zone from 00 to 86.
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        public SettingGlobalState()
        {
        }
        public static new SettingGlobalState Empty => new SettingGlobalState();
    }
}