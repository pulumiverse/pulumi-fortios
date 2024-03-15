// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Extendercontroller
{
    /// <summary>
    /// FortiExtender extender profile configuration. Applies to FortiOS Version `7.0.2,7.0.3,7.0.4,7.0.5,7.0.6,7.2.0`.
    /// 
    /// ## Import
    /// 
    /// ExtenderController ExtenderProfile can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:extendercontroller/extenderprofile:Extenderprofile labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:extendercontroller/extenderprofile:Extenderprofile labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:extendercontroller/extenderprofile:Extenderprofile")]
    public partial class Extenderprofile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
        /// </summary>
        [Output("allowaccess")]
        public Output<string> Allowaccess { get; private set; } = null!;

        /// <summary>
        /// FortiExtender LAN extension bandwidth limit (Mbps).
        /// </summary>
        [Output("bandwidthLimit")]
        public Output<int> BandwidthLimit { get; private set; } = null!;

        /// <summary>
        /// FortiExtender cellular configuration. The structure of `cellular` block is documented below.
        /// </summary>
        [Output("cellular")]
        public Output<Outputs.ExtenderprofileCellular> Cellular { get; private set; } = null!;

        /// <summary>
        /// Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("enforceBandwidth")]
        public Output<string> EnforceBandwidth { get; private set; } = null!;

        /// <summary>
        /// Extension option. Valid values: `wan-extension`, `lan-extension`.
        /// </summary>
        [Output("extension")]
        public Output<string> Extension { get; private set; } = null!;

        /// <summary>
        /// id
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// FortiExtender lan extension configuration. The structure of `lan_extension` block is documented below.
        /// </summary>
        [Output("lanExtension")]
        public Output<Outputs.ExtenderprofileLanExtension> LanExtension { get; private set; } = null!;

        /// <summary>
        /// Set the managed extender's administrator password.
        /// </summary>
        [Output("loginPassword")]
        public Output<string?> LoginPassword { get; private set; } = null!;

        /// <summary>
        /// Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
        /// </summary>
        [Output("loginPasswordChange")]
        public Output<string> LoginPasswordChange { get; private set; } = null!;

        /// <summary>
        /// Model. Valid values: `FX201E`, `FX211E`, `FX200F`, `FXA11F`, `FXE11F`, `FXA21F`, `FXE21F`, `FXA22F`, `FXE22F`, `FX212F`, `FX311F`, `FX312F`, `FX511F`, `FVG21F`, `FVA21F`, `FVG22F`, `FVA22F`, `FX04DA`.
        /// </summary>
        [Output("model")]
        public Output<string> Model { get; private set; } = null!;

        /// <summary>
        /// FortiExtender profile name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Extenderprofile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Extenderprofile(string name, ExtenderprofileArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:extendercontroller/extenderprofile:Extenderprofile", name, args ?? new ExtenderprofileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Extenderprofile(string name, Input<string> id, ExtenderprofileState? state = null, CustomResourceOptions? options = null)
            : base("fortios:extendercontroller/extenderprofile:Extenderprofile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Extenderprofile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Extenderprofile Get(string name, Input<string> id, ExtenderprofileState? state = null, CustomResourceOptions? options = null)
        {
            return new Extenderprofile(name, id, state, options);
        }
    }

    public sealed class ExtenderprofileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
        /// </summary>
        [Input("allowaccess")]
        public Input<string>? Allowaccess { get; set; }

        /// <summary>
        /// FortiExtender LAN extension bandwidth limit (Mbps).
        /// </summary>
        [Input("bandwidthLimit")]
        public Input<int>? BandwidthLimit { get; set; }

        /// <summary>
        /// FortiExtender cellular configuration. The structure of `cellular` block is documented below.
        /// </summary>
        [Input("cellular")]
        public Input<Inputs.ExtenderprofileCellularArgs>? Cellular { get; set; }

        /// <summary>
        /// Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("enforceBandwidth")]
        public Input<string>? EnforceBandwidth { get; set; }

        /// <summary>
        /// Extension option. Valid values: `wan-extension`, `lan-extension`.
        /// </summary>
        [Input("extension")]
        public Input<string>? Extension { get; set; }

        /// <summary>
        /// id
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// FortiExtender lan extension configuration. The structure of `lan_extension` block is documented below.
        /// </summary>
        [Input("lanExtension")]
        public Input<Inputs.ExtenderprofileLanExtensionArgs>? LanExtension { get; set; }

        /// <summary>
        /// Set the managed extender's administrator password.
        /// </summary>
        [Input("loginPassword")]
        public Input<string>? LoginPassword { get; set; }

        /// <summary>
        /// Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
        /// </summary>
        [Input("loginPasswordChange")]
        public Input<string>? LoginPasswordChange { get; set; }

        /// <summary>
        /// Model. Valid values: `FX201E`, `FX211E`, `FX200F`, `FXA11F`, `FXE11F`, `FXA21F`, `FXE21F`, `FXA22F`, `FXE22F`, `FX212F`, `FX311F`, `FX312F`, `FX511F`, `FVG21F`, `FVA21F`, `FVG22F`, `FVA22F`, `FX04DA`.
        /// </summary>
        [Input("model")]
        public Input<string>? Model { get; set; }

        /// <summary>
        /// FortiExtender profile name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ExtenderprofileArgs()
        {
        }
        public static new ExtenderprofileArgs Empty => new ExtenderprofileArgs();
    }

    public sealed class ExtenderprofileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
        /// </summary>
        [Input("allowaccess")]
        public Input<string>? Allowaccess { get; set; }

        /// <summary>
        /// FortiExtender LAN extension bandwidth limit (Mbps).
        /// </summary>
        [Input("bandwidthLimit")]
        public Input<int>? BandwidthLimit { get; set; }

        /// <summary>
        /// FortiExtender cellular configuration. The structure of `cellular` block is documented below.
        /// </summary>
        [Input("cellular")]
        public Input<Inputs.ExtenderprofileCellularGetArgs>? Cellular { get; set; }

        /// <summary>
        /// Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("enforceBandwidth")]
        public Input<string>? EnforceBandwidth { get; set; }

        /// <summary>
        /// Extension option. Valid values: `wan-extension`, `lan-extension`.
        /// </summary>
        [Input("extension")]
        public Input<string>? Extension { get; set; }

        /// <summary>
        /// id
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// FortiExtender lan extension configuration. The structure of `lan_extension` block is documented below.
        /// </summary>
        [Input("lanExtension")]
        public Input<Inputs.ExtenderprofileLanExtensionGetArgs>? LanExtension { get; set; }

        /// <summary>
        /// Set the managed extender's administrator password.
        /// </summary>
        [Input("loginPassword")]
        public Input<string>? LoginPassword { get; set; }

        /// <summary>
        /// Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
        /// </summary>
        [Input("loginPasswordChange")]
        public Input<string>? LoginPasswordChange { get; set; }

        /// <summary>
        /// Model. Valid values: `FX201E`, `FX211E`, `FX200F`, `FXA11F`, `FXE11F`, `FXA21F`, `FXE21F`, `FXA22F`, `FXE22F`, `FX212F`, `FX311F`, `FX312F`, `FX511F`, `FVG21F`, `FVA21F`, `FVG22F`, `FVA22F`, `FX04DA`.
        /// </summary>
        [Input("model")]
        public Input<string>? Model { get; set; }

        /// <summary>
        /// FortiExtender profile name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ExtenderprofileState()
        {
        }
        public static new ExtenderprofileState Empty => new ExtenderprofileState();
    }
}