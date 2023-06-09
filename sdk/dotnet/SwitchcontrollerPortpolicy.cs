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
    /// Configure port policy to be applied on the managed FortiSwitch ports through NAC device. Applies to FortiOS Version `6.4.0,6.4.1,6.4.2,6.4.10,7.0.0`.
    /// 
    /// ## Import
    /// 
    /// SwitchController PortPolicy can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollerPortpolicy:SwitchcontrollerPortpolicy labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollerPortpolicy:SwitchcontrollerPortpolicy labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/switchcontrollerPortpolicy:SwitchcontrollerPortpolicy")]
    public partial class SwitchcontrollerPortpolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable bouncing (administratively bring the link down, up) of a switch port where this port policy is applied. Helps to clear and reassign VLAN from lldp-profile. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("bouncePortLink")]
        public Output<string> BouncePortLink { get; private set; } = null!;

        /// <summary>
        /// Description for the port policy.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// FortiLink interface for which this port policy belongs to.
        /// </summary>
        [Output("fortilink")]
        public Output<string> Fortilink { get; private set; } = null!;

        /// <summary>
        /// LLDP profile to be applied when using this port-policy.
        /// </summary>
        [Output("lldpProfile")]
        public Output<string> LldpProfile { get; private set; } = null!;

        /// <summary>
        /// 802.1x security policy to be applied when using this port-policy.
        /// </summary>
        [Output("n8021x")]
        public Output<string> N8021x { get; private set; } = null!;

        /// <summary>
        /// Port policy name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// QoS policy to be applied when using this port-policy.
        /// </summary>
        [Output("qosPolicy")]
        public Output<string> QosPolicy { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// VLAN policy to be applied when using this port-policy.
        /// </summary>
        [Output("vlanPolicy")]
        public Output<string> VlanPolicy { get; private set; } = null!;


        /// <summary>
        /// Create a SwitchcontrollerPortpolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SwitchcontrollerPortpolicy(string name, SwitchcontrollerPortpolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollerPortpolicy:SwitchcontrollerPortpolicy", name, args ?? new SwitchcontrollerPortpolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SwitchcontrollerPortpolicy(string name, Input<string> id, SwitchcontrollerPortpolicyState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollerPortpolicy:SwitchcontrollerPortpolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SwitchcontrollerPortpolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SwitchcontrollerPortpolicy Get(string name, Input<string> id, SwitchcontrollerPortpolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new SwitchcontrollerPortpolicy(name, id, state, options);
        }
    }

    public sealed class SwitchcontrollerPortpolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable bouncing (administratively bring the link down, up) of a switch port where this port policy is applied. Helps to clear and reassign VLAN from lldp-profile. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("bouncePortLink")]
        public Input<string>? BouncePortLink { get; set; }

        /// <summary>
        /// Description for the port policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// FortiLink interface for which this port policy belongs to.
        /// </summary>
        [Input("fortilink")]
        public Input<string>? Fortilink { get; set; }

        /// <summary>
        /// LLDP profile to be applied when using this port-policy.
        /// </summary>
        [Input("lldpProfile")]
        public Input<string>? LldpProfile { get; set; }

        /// <summary>
        /// 802.1x security policy to be applied when using this port-policy.
        /// </summary>
        [Input("n8021x")]
        public Input<string>? N8021x { get; set; }

        /// <summary>
        /// Port policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// QoS policy to be applied when using this port-policy.
        /// </summary>
        [Input("qosPolicy")]
        public Input<string>? QosPolicy { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// VLAN policy to be applied when using this port-policy.
        /// </summary>
        [Input("vlanPolicy")]
        public Input<string>? VlanPolicy { get; set; }

        public SwitchcontrollerPortpolicyArgs()
        {
        }
        public static new SwitchcontrollerPortpolicyArgs Empty => new SwitchcontrollerPortpolicyArgs();
    }

    public sealed class SwitchcontrollerPortpolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable bouncing (administratively bring the link down, up) of a switch port where this port policy is applied. Helps to clear and reassign VLAN from lldp-profile. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("bouncePortLink")]
        public Input<string>? BouncePortLink { get; set; }

        /// <summary>
        /// Description for the port policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// FortiLink interface for which this port policy belongs to.
        /// </summary>
        [Input("fortilink")]
        public Input<string>? Fortilink { get; set; }

        /// <summary>
        /// LLDP profile to be applied when using this port-policy.
        /// </summary>
        [Input("lldpProfile")]
        public Input<string>? LldpProfile { get; set; }

        /// <summary>
        /// 802.1x security policy to be applied when using this port-policy.
        /// </summary>
        [Input("n8021x")]
        public Input<string>? N8021x { get; set; }

        /// <summary>
        /// Port policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// QoS policy to be applied when using this port-policy.
        /// </summary>
        [Input("qosPolicy")]
        public Input<string>? QosPolicy { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// VLAN policy to be applied when using this port-policy.
        /// </summary>
        [Input("vlanPolicy")]
        public Input<string>? VlanPolicy { get; set; }

        public SwitchcontrollerPortpolicyState()
        {
        }
        public static new SwitchcontrollerPortpolicyState Empty => new SwitchcontrollerPortpolicyState();
    }
}
