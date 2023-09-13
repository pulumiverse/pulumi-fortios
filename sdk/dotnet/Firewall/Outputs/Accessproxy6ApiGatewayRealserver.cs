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
    public sealed class Accessproxy6ApiGatewayRealserver
    {
        /// <summary>
        /// Type of address. Valid values: `ip`, `fqdn`.
        /// </summary>
        public readonly string? AddrType;
        /// <summary>
        /// Address or address group of the real server.
        /// </summary>
        public readonly string? Address;
        /// <summary>
        /// Wildcard domain name of the real server.
        /// </summary>
        public readonly string? Domain;
        /// <summary>
        /// Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? HealthCheck;
        /// <summary>
        /// Protocol of the health check monitor to use when polling to determine server's connectivity status. Valid values: `ping`, `http`, `tcp-connect`.
        /// </summary>
        public readonly string? HealthCheckProto;
        /// <summary>
        /// Enable/disable holddown timer. Server will be considered active and reachable once the holddown period has expired (30 seconds). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? HolddownInterval;
        /// <summary>
        /// HTTP server domain name in HTTP header.
        /// </summary>
        public readonly string? HttpHost;
        /// <summary>
        /// Real server ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// IPv6 address of the real server.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// Port for communicating with the real server.
        /// </summary>
        public readonly string? Mappedport;
        /// <summary>
        /// Port for communicating with the real server.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// Set access-proxy SSH client certificate profile.
        /// </summary>
        public readonly string? SshClientCert;
        /// <summary>
        /// Enable/disable SSH real server host key validation. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? SshHostKeyValidation;
        /// <summary>
        /// One or more server host key. The structure of `ssh_host_key` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.Accessproxy6ApiGatewayRealserverSshHostKey> SshHostKeys;
        /// <summary>
        /// Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active`, `standby`, `disable`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// TCP forwarding server type. Valid values: `tcp-forwarding`, `ssh`.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.
        /// </summary>
        public readonly int? Weight;

        [OutputConstructor]
        private Accessproxy6ApiGatewayRealserver(
            string? addrType,

            string? address,

            string? domain,

            string? healthCheck,

            string? healthCheckProto,

            string? holddownInterval,

            string? httpHost,

            int? id,

            string? ip,

            string? mappedport,

            int? port,

            string? sshClientCert,

            string? sshHostKeyValidation,

            ImmutableArray<Outputs.Accessproxy6ApiGatewayRealserverSshHostKey> sshHostKeys,

            string? status,

            string? type,

            int? weight)
        {
            AddrType = addrType;
            Address = address;
            Domain = domain;
            HealthCheck = healthCheck;
            HealthCheckProto = healthCheckProto;
            HolddownInterval = holddownInterval;
            HttpHost = httpHost;
            Id = id;
            Ip = ip;
            Mappedport = mappedport;
            Port = port;
            SshClientCert = sshClientCert;
            SshHostKeyValidation = sshHostKeyValidation;
            SshHostKeys = sshHostKeys;
            Status = status;
            Type = type;
            Weight = weight;
        }
    }
}
