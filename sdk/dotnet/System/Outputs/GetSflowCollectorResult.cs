// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Outputs
{

    [OutputType]
    public sealed class GetSflowCollectorResult
    {
        /// <summary>
        /// IP addresses of the sFlow collectors that sFlow agents added to interfaces in this VDOM send sFlow datagrams to.
        /// </summary>
        public readonly string CollectorIp;
        /// <summary>
        /// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
        /// </summary>
        public readonly int CollectorPort;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        public readonly string Interface;
        /// <summary>
        /// Specify how to select outgoing interface to reach server.
        /// </summary>
        public readonly string InterfaceSelectMethod;
        /// <summary>
        /// Source IP address for sFlow agent.
        /// </summary>
        public readonly string SourceIp;

        [OutputConstructor]
        private GetSflowCollectorResult(
            string collectorIp,

            int collectorPort,

            int id,

            string @interface,

            string interfaceSelectMethod,

            string sourceIp)
        {
            CollectorIp = collectorIp;
            CollectorPort = collectorPort;
            Id = id;
            Interface = @interface;
            InterfaceSelectMethod = interfaceSelectMethod;
            SourceIp = sourceIp;
        }
    }
}
