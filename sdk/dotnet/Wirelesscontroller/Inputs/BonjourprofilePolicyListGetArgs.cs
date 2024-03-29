// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Inputs
{

    public sealed class BonjourprofilePolicyListGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// VLAN ID from which the Bonjour service is advertised (0 - 4094, default = 0).
        /// </summary>
        [Input("fromVlan")]
        public Input<string>? FromVlan { get; set; }

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Input("policyId")]
        public Input<int>? PolicyId { get; set; }

        /// <summary>
        /// Bonjour services for the VLAN connecting to the Bonjour network.
        /// </summary>
        [Input("services")]
        public Input<string>? Services { get; set; }

        /// <summary>
        /// VLAN ID to which the Bonjour service is made available (0 - 4094, default = all).
        /// </summary>
        [Input("toVlan")]
        public Input<string>? ToVlan { get; set; }

        public BonjourprofilePolicyListGetArgs()
        {
        }
        public static new BonjourprofilePolicyListGetArgs Empty => new BonjourprofilePolicyListGetArgs();
    }
}
