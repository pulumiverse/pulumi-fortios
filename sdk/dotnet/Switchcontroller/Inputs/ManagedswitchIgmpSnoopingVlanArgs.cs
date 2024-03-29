// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Inputs
{

    public sealed class ManagedswitchIgmpSnoopingVlanArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IGMP snooping proxy for the VLAN interface. Valid values: `disable`, `enable`, `global`.
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        /// <summary>
        /// Enable/disable IGMP snooping querier for the VLAN interface. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("querier")]
        public Input<string>? Querier { get; set; }

        /// <summary>
        /// IGMP snooping querier address.
        /// </summary>
        [Input("querierAddr")]
        public Input<string>? QuerierAddr { get; set; }

        /// <summary>
        /// IGMP snooping querier version.
        /// 
        /// The `n802_1x_settings` block supports:
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        /// <summary>
        /// List of FortiSwitch VLANs.
        /// </summary>
        [Input("vlanName")]
        public Input<string>? VlanName { get; set; }

        public ManagedswitchIgmpSnoopingVlanArgs()
        {
        }
        public static new ManagedswitchIgmpSnoopingVlanArgs Empty => new ManagedswitchIgmpSnoopingVlanArgs();
    }
}
