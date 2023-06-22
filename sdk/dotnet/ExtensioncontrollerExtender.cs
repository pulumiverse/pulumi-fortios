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
    /// Extender controller configuration. Applies to FortiOS Version `&gt;= 7.2.1`.
    /// 
    /// ## Import
    /// 
    /// ExtensionController Extender can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/extensioncontrollerExtender:ExtensioncontrollerExtender labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/extensioncontrollerExtender:ExtensioncontrollerExtender labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/extensioncontrollerExtender:ExtensioncontrollerExtender")]
    public partial class ExtensioncontrollerExtender : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
        /// </summary>
        [Output("allowaccess")]
        public Output<string> Allowaccess { get; private set; } = null!;

        /// <summary>
        /// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
        /// </summary>
        [Output("authorized")]
        public Output<string> Authorized { get; private set; } = null!;

        /// <summary>
        /// FortiExtender LAN extension bandwidth limit (Mbps).
        /// </summary>
        [Output("bandwidthLimit")]
        public Output<int> BandwidthLimit { get; private set; } = null!;

        /// <summary>
        /// Description.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Device ID.
        /// </summary>
        [Output("deviceId")]
        public Output<int> DeviceId { get; private set; } = null!;

        /// <summary>
        /// Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("enforceBandwidth")]
        public Output<string> EnforceBandwidth { get; private set; } = null!;

        /// <summary>
        /// FortiExtender name.
        /// </summary>
        [Output("extName")]
        public Output<string> ExtName { get; private set; } = null!;

        /// <summary>
        /// Extension type for this FortiExtender. Valid values: `wan-extension`, `lan-extension`.
        /// </summary>
        [Output("extensionType")]
        public Output<string> ExtensionType { get; private set; } = null!;

        /// <summary>
        /// Enable/disable one-time automatic provisioning of the latest firmware version. Valid values: `disable`, `once`.
        /// </summary>
        [Output("firmwareProvisionLatest")]
        public Output<string> FirmwareProvisionLatest { get; private set; } = null!;

        /// <summary>
        /// FortiExtender serial number.
        /// </summary>
        [Output("fosid")]
        public Output<string> Fosid { get; private set; } = null!;

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
        /// FortiExtender entry name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable to override the extender profile management access configuration. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("overrideAllowaccess")]
        public Output<string> OverrideAllowaccess { get; private set; } = null!;

        /// <summary>
        /// Enable to override the extender profile enforce-bandwidth setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("overrideEnforceBandwidth")]
        public Output<string> OverrideEnforceBandwidth { get; private set; } = null!;

        /// <summary>
        /// Enable to override the extender profile login-password (administrator password) setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("overrideLoginPasswordChange")]
        public Output<string> OverrideLoginPasswordChange { get; private set; } = null!;

        /// <summary>
        /// FortiExtender profile configuration.
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
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// FortiExtender wan extension configuration. The structure of `wan_extension` block is documented below.
        /// </summary>
        [Output("wanExtension")]
        public Output<Outputs.ExtensioncontrollerExtenderWanExtension> WanExtension { get; private set; } = null!;


        /// <summary>
        /// Create a ExtensioncontrollerExtender resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExtensioncontrollerExtender(string name, ExtensioncontrollerExtenderArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/extensioncontrollerExtender:ExtensioncontrollerExtender", name, args ?? new ExtensioncontrollerExtenderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExtensioncontrollerExtender(string name, Input<string> id, ExtensioncontrollerExtenderState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/extensioncontrollerExtender:ExtensioncontrollerExtender", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ExtensioncontrollerExtender resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExtensioncontrollerExtender Get(string name, Input<string> id, ExtensioncontrollerExtenderState? state = null, CustomResourceOptions? options = null)
        {
            return new ExtensioncontrollerExtender(name, id, state, options);
        }
    }

    public sealed class ExtensioncontrollerExtenderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
        /// </summary>
        [Input("allowaccess")]
        public Input<string>? Allowaccess { get; set; }

        /// <summary>
        /// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
        /// </summary>
        [Input("authorized")]
        public Input<string>? Authorized { get; set; }

        /// <summary>
        /// FortiExtender LAN extension bandwidth limit (Mbps).
        /// </summary>
        [Input("bandwidthLimit")]
        public Input<int>? BandwidthLimit { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Device ID.
        /// </summary>
        [Input("deviceId")]
        public Input<int>? DeviceId { get; set; }

        /// <summary>
        /// Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("enforceBandwidth")]
        public Input<string>? EnforceBandwidth { get; set; }

        /// <summary>
        /// FortiExtender name.
        /// </summary>
        [Input("extName")]
        public Input<string>? ExtName { get; set; }

        /// <summary>
        /// Extension type for this FortiExtender. Valid values: `wan-extension`, `lan-extension`.
        /// </summary>
        [Input("extensionType")]
        public Input<string>? ExtensionType { get; set; }

        /// <summary>
        /// Enable/disable one-time automatic provisioning of the latest firmware version. Valid values: `disable`, `once`.
        /// </summary>
        [Input("firmwareProvisionLatest")]
        public Input<string>? FirmwareProvisionLatest { get; set; }

        /// <summary>
        /// FortiExtender serial number.
        /// </summary>
        [Input("fosid")]
        public Input<string>? Fosid { get; set; }

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
        /// FortiExtender entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable to override the extender profile management access configuration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("overrideAllowaccess")]
        public Input<string>? OverrideAllowaccess { get; set; }

        /// <summary>
        /// Enable to override the extender profile enforce-bandwidth setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("overrideEnforceBandwidth")]
        public Input<string>? OverrideEnforceBandwidth { get; set; }

        /// <summary>
        /// Enable to override the extender profile login-password (administrator password) setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("overrideLoginPasswordChange")]
        public Input<string>? OverrideLoginPasswordChange { get; set; }

        /// <summary>
        /// FortiExtender profile configuration.
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

        /// <summary>
        /// FortiExtender wan extension configuration. The structure of `wan_extension` block is documented below.
        /// </summary>
        [Input("wanExtension")]
        public Input<Inputs.ExtensioncontrollerExtenderWanExtensionArgs>? WanExtension { get; set; }

        public ExtensioncontrollerExtenderArgs()
        {
        }
        public static new ExtensioncontrollerExtenderArgs Empty => new ExtensioncontrollerExtenderArgs();
    }

    public sealed class ExtensioncontrollerExtenderState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
        /// </summary>
        [Input("allowaccess")]
        public Input<string>? Allowaccess { get; set; }

        /// <summary>
        /// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
        /// </summary>
        [Input("authorized")]
        public Input<string>? Authorized { get; set; }

        /// <summary>
        /// FortiExtender LAN extension bandwidth limit (Mbps).
        /// </summary>
        [Input("bandwidthLimit")]
        public Input<int>? BandwidthLimit { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Device ID.
        /// </summary>
        [Input("deviceId")]
        public Input<int>? DeviceId { get; set; }

        /// <summary>
        /// Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("enforceBandwidth")]
        public Input<string>? EnforceBandwidth { get; set; }

        /// <summary>
        /// FortiExtender name.
        /// </summary>
        [Input("extName")]
        public Input<string>? ExtName { get; set; }

        /// <summary>
        /// Extension type for this FortiExtender. Valid values: `wan-extension`, `lan-extension`.
        /// </summary>
        [Input("extensionType")]
        public Input<string>? ExtensionType { get; set; }

        /// <summary>
        /// Enable/disable one-time automatic provisioning of the latest firmware version. Valid values: `disable`, `once`.
        /// </summary>
        [Input("firmwareProvisionLatest")]
        public Input<string>? FirmwareProvisionLatest { get; set; }

        /// <summary>
        /// FortiExtender serial number.
        /// </summary>
        [Input("fosid")]
        public Input<string>? Fosid { get; set; }

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
        /// FortiExtender entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable to override the extender profile management access configuration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("overrideAllowaccess")]
        public Input<string>? OverrideAllowaccess { get; set; }

        /// <summary>
        /// Enable to override the extender profile enforce-bandwidth setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("overrideEnforceBandwidth")]
        public Input<string>? OverrideEnforceBandwidth { get; set; }

        /// <summary>
        /// Enable to override the extender profile login-password (administrator password) setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("overrideLoginPasswordChange")]
        public Input<string>? OverrideLoginPasswordChange { get; set; }

        /// <summary>
        /// FortiExtender profile configuration.
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

        /// <summary>
        /// FortiExtender wan extension configuration. The structure of `wan_extension` block is documented below.
        /// </summary>
        [Input("wanExtension")]
        public Input<Inputs.ExtensioncontrollerExtenderWanExtensionGetArgs>? WanExtension { get; set; }

        public ExtensioncontrollerExtenderState()
        {
        }
        public static new ExtensioncontrollerExtenderState Empty => new ExtensioncontrollerExtenderState();
    }
}
