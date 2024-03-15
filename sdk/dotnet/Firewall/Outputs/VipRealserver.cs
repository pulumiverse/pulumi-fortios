// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Outputs
{

    [OutputType]
    public sealed class VipRealserver
    {
        /// <summary>
        /// Dynamic address of the real server.
        /// </summary>
        public readonly string? Address;
        /// <summary>
        /// Only clients in this IP range can connect to this real server.
        /// </summary>
        public readonly string? ClientIp;
        /// <summary>
        /// Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable`, `enable`, `vip`.
        /// </summary>
        public readonly string? Healthcheck;
        /// <summary>
        /// Time in seconds that the health check monitor continues to monitor and unresponsive server that should be active.
        /// </summary>
        public readonly int? HolddownInterval;
        /// <summary>
        /// HTTP server domain name in HTTP header.
        /// </summary>
        public readonly string? HttpHost;
        /// <summary>
        /// Real server ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// IP address of the real server.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// Max number of active connections that can be directed to the real server. When reached, sessions are sent to other real servers.
        /// </summary>
        public readonly int? MaxConnections;
        /// <summary>
        /// Name of the health check monitor to use when polling to determine a virtual server's connectivity status.
        /// </summary>
        public readonly string? Monitor;
        /// <summary>
        /// Port for communicating with the real server. Required if port forwarding is enabled.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active`, `standby`, `disable`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Type of address. Valid values: `ip`, `address`.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.
        /// </summary>
        public readonly int? Weight;

        [OutputConstructor]
        private VipRealserver(
            string? address,

            string? clientIp,

            string? healthcheck,

            int? holddownInterval,

            string? httpHost,

            int? id,

            string? ip,

            int? maxConnections,

            string? monitor,

            int? port,

            string? status,

            string? type,

            int? weight)
        {
            Address = address;
            ClientIp = clientIp;
            Healthcheck = healthcheck;
            HolddownInterval = holddownInterval;
            HttpHost = httpHost;
            Id = id;
            Ip = ip;
            MaxConnections = maxConnections;
            Monitor = monitor;
            Port = port;
            Status = status;
            Type = type;
            Weight = weight;
        }
    }
}