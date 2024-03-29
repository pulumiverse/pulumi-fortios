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

    public sealed class SflowCollectorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP addresses of the sFlow collectors that sFlow agents added to interfaces in this VDOM send sFlow datagrams to.
        /// </summary>
        [Input("collectorIp")]
        public Input<string>? CollectorIp { get; set; }

        /// <summary>
        /// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
        /// </summary>
        [Input("collectorPort")]
        public Input<int>? CollectorPort { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// Source IP address for sFlow agent.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        public SflowCollectorArgs()
        {
        }
        public static new SflowCollectorArgs Empty => new SflowCollectorArgs();
    }
}
