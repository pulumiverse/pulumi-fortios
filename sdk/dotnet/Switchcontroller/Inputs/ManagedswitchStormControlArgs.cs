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

    public sealed class ManagedswitchStormControlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable storm control to drop broadcast traffic. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("broadcast")]
        public Input<string>? Broadcast { get; set; }

        /// <summary>
        /// Enable to override global FortiSwitch storm control settings for this FortiSwitch. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("localOverride")]
        public Input<string>? LocalOverride { get; set; }

        /// <summary>
        /// Rate in packets per second at which storm traffic is controlled (1 - 10000000, default = 500). Storm control drops excess traffic data rates beyond this threshold.
        /// </summary>
        [Input("rate")]
        public Input<int>? Rate { get; set; }

        /// <summary>
        /// Enable/disable storm control to drop unknown multicast traffic. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("unknownMulticast")]
        public Input<string>? UnknownMulticast { get; set; }

        /// <summary>
        /// Enable/disable storm control to drop unknown unicast traffic. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("unknownUnicast")]
        public Input<string>? UnknownUnicast { get; set; }

        public ManagedswitchStormControlArgs()
        {
        }
        public static new ManagedswitchStormControlArgs Empty => new ManagedswitchStormControlArgs();
    }
}
