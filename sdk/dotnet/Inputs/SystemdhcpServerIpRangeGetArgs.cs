// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class SystemdhcpServerIpRangeGetArgs : global::Pulumi.ResourceArgs
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
        /// Start of IP range.
        /// </summary>
        [Input("startIp")]
        public Input<string>? StartIp { get; set; }

        /// <summary>
        /// Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this range. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("vciMatch")]
        public Input<string>? VciMatch { get; set; }

        [Input("vciStrings")]
        private InputList<Inputs.SystemdhcpServerIpRangeVciStringGetArgs>? _vciStrings;

        /// <summary>
        /// One or more VCI strings in quotes separated by spaces. The structure of `vci_string` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemdhcpServerIpRangeVciStringGetArgs> VciStrings
        {
            get => _vciStrings ?? (_vciStrings = new InputList<Inputs.SystemdhcpServerIpRangeVciStringGetArgs>());
            set => _vciStrings = value;
        }

        public SystemdhcpServerIpRangeGetArgs()
        {
        }
        public static new SystemdhcpServerIpRangeGetArgs Empty => new SystemdhcpServerIpRangeGetArgs();
    }
}
