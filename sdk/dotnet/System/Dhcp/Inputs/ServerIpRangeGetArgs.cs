// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Dhcp.Inputs
{

    public sealed class ServerIpRangeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// End of IP range.
        /// </summary>
        [Input("endIp")]
        public Input<string>? EndIp { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Lease time in seconds, 0 means default lease time.
        /// </summary>
        [Input("leaseTime")]
        public Input<int>? LeaseTime { get; set; }

        /// <summary>
        /// Start of IP range.
        /// </summary>
        [Input("startIp")]
        public Input<string>? StartIp { get; set; }

        /// <summary>
        /// Enable/disable user class identifier (UCI) matching. When enabled only DHCP requests with a matching UCI are served with this range. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("uciMatch")]
        public Input<string>? UciMatch { get; set; }

        [Input("uciStrings")]
        private InputList<Inputs.ServerIpRangeUciStringGetArgs>? _uciStrings;

        /// <summary>
        /// One or more UCI strings in quotes separated by spaces. The structure of `uci_string` block is documented below.
        /// </summary>
        public InputList<Inputs.ServerIpRangeUciStringGetArgs> UciStrings
        {
            get => _uciStrings ?? (_uciStrings = new InputList<Inputs.ServerIpRangeUciStringGetArgs>());
            set => _uciStrings = value;
        }

        /// <summary>
        /// Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this range. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("vciMatch")]
        public Input<string>? VciMatch { get; set; }

        [Input("vciStrings")]
        private InputList<Inputs.ServerIpRangeVciStringGetArgs>? _vciStrings;

        /// <summary>
        /// One or more VCI strings in quotes separated by spaces. The structure of `vci_string` block is documented below.
        /// </summary>
        public InputList<Inputs.ServerIpRangeVciStringGetArgs> VciStrings
        {
            get => _vciStrings ?? (_vciStrings = new InputList<Inputs.ServerIpRangeVciStringGetArgs>());
            set => _vciStrings = value;
        }

        public ServerIpRangeGetArgs()
        {
        }
        public static new ServerIpRangeGetArgs Empty => new ServerIpRangeGetArgs();
    }
}
