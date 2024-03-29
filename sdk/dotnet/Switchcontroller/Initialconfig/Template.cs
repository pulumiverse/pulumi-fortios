// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Initialconfig
{
    /// <summary>
    /// Configure template for auto-generated VLANs. Applies to FortiOS Version `&gt;= 6.4.0`.
    /// 
    /// ## Import
    /// 
    /// SwitchControllerInitialConfig Template can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/initialconfig/template:Template labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/initialconfig/template:Template labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:switchcontroller/initialconfig/template:Template")]
    public partial class Template : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Permitted types of management access to this interface. Valid values: `ping`, `https`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `radius-acct`, `probe-response`, `fabric`, `ftm`.
        /// </summary>
        [Output("allowaccess")]
        public Output<string> Allowaccess { get; private set; } = null!;

        /// <summary>
        /// Automatically allocate interface address and subnet block. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("autoIp")]
        public Output<string> AutoIp { get; private set; } = null!;

        /// <summary>
        /// Enable/disable a DHCP server on this interface. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dhcpServer")]
        public Output<string> DhcpServer { get; private set; } = null!;

        /// <summary>
        /// Interface IPv4 address and subnet mask.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// Initial config template name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Unique VLAN ID.
        /// </summary>
        [Output("vlanid")]
        public Output<int> Vlanid { get; private set; } = null!;


        /// <summary>
        /// Create a Template resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Template(string name, TemplateArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/initialconfig/template:Template", name, args ?? new TemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Template(string name, Input<string> id, TemplateState? state = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/initialconfig/template:Template", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Template resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Template Get(string name, Input<string> id, TemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new Template(name, id, state, options);
        }
    }

    public sealed class TemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Permitted types of management access to this interface. Valid values: `ping`, `https`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `radius-acct`, `probe-response`, `fabric`, `ftm`.
        /// </summary>
        [Input("allowaccess")]
        public Input<string>? Allowaccess { get; set; }

        /// <summary>
        /// Automatically allocate interface address and subnet block. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("autoIp")]
        public Input<string>? AutoIp { get; set; }

        /// <summary>
        /// Enable/disable a DHCP server on this interface. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dhcpServer")]
        public Input<string>? DhcpServer { get; set; }

        /// <summary>
        /// Interface IPv4 address and subnet mask.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Initial config template name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Unique VLAN ID.
        /// </summary>
        [Input("vlanid")]
        public Input<int>? Vlanid { get; set; }

        public TemplateArgs()
        {
        }
        public static new TemplateArgs Empty => new TemplateArgs();
    }

    public sealed class TemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Permitted types of management access to this interface. Valid values: `ping`, `https`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `radius-acct`, `probe-response`, `fabric`, `ftm`.
        /// </summary>
        [Input("allowaccess")]
        public Input<string>? Allowaccess { get; set; }

        /// <summary>
        /// Automatically allocate interface address and subnet block. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("autoIp")]
        public Input<string>? AutoIp { get; set; }

        /// <summary>
        /// Enable/disable a DHCP server on this interface. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dhcpServer")]
        public Input<string>? DhcpServer { get; set; }

        /// <summary>
        /// Interface IPv4 address and subnet mask.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Initial config template name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Unique VLAN ID.
        /// </summary>
        [Input("vlanid")]
        public Input<int>? Vlanid { get; set; }

        public TemplateState()
        {
        }
        public static new TemplateState Empty => new TemplateState();
    }
}
