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

    public sealed class StandaloneclusterClusterPeerSessionSyncFilterGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("customServices")]
        private InputList<Inputs.StandaloneclusterClusterPeerSessionSyncFilterCustomServiceGetArgs>? _customServices;

        /// <summary>
        /// Only sessions using these custom services are synchronized. Use source and destination port ranges to define these custom services. The structure of `custom_service` block is documented below.
        /// </summary>
        public InputList<Inputs.StandaloneclusterClusterPeerSessionSyncFilterCustomServiceGetArgs> CustomServices
        {
            get => _customServices ?? (_customServices = new InputList<Inputs.StandaloneclusterClusterPeerSessionSyncFilterCustomServiceGetArgs>());
            set => _customServices = value;
        }

        /// <summary>
        /// Only sessions to this IPv4 address are synchronized.
        /// </summary>
        [Input("dstaddr")]
        public Input<string>? Dstaddr { get; set; }

        /// <summary>
        /// Only sessions to this IPv6 address are synchronized.
        /// </summary>
        [Input("dstaddr6")]
        public Input<string>? Dstaddr6 { get; set; }

        /// <summary>
        /// Only sessions to this interface are synchronized.
        /// </summary>
        [Input("dstintf")]
        public Input<string>? Dstintf { get; set; }

        /// <summary>
        /// Only sessions from this IPv4 address are synchronized.
        /// </summary>
        [Input("srcaddr")]
        public Input<string>? Srcaddr { get; set; }

        /// <summary>
        /// Only sessions from this IPv6 address are synchronized.
        /// </summary>
        [Input("srcaddr6")]
        public Input<string>? Srcaddr6 { get; set; }

        /// <summary>
        /// Only sessions from this interface are synchronized.
        /// </summary>
        [Input("srcintf")]
        public Input<string>? Srcintf { get; set; }

        public StandaloneclusterClusterPeerSessionSyncFilterGetArgs()
        {
        }
        public static new StandaloneclusterClusterPeerSessionSyncFilterGetArgs Empty => new StandaloneclusterClusterPeerSessionSyncFilterGetArgs();
    }
}
