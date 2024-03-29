// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Inputs
{

    public sealed class StandaloneclusterClusterPeerSessionSyncFilterCustomServiceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom service destination port range.
        /// </summary>
        [Input("dstPortRange")]
        public Input<string>? DstPortRange { get; set; }

        /// <summary>
        /// Custom service ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Custom service source port range.
        /// </summary>
        [Input("srcPortRange")]
        public Input<string>? SrcPortRange { get; set; }

        public StandaloneclusterClusterPeerSessionSyncFilterCustomServiceGetArgs()
        {
        }
        public static new StandaloneclusterClusterPeerSessionSyncFilterCustomServiceGetArgs Empty => new StandaloneclusterClusterPeerSessionSyncFilterCustomServiceGetArgs();
    }
}
