// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class SystemSpeedtestserverHost
    {
        /// <summary>
        /// Server host ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Server host IPv4 address.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// Speed test host password.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Server host port number to communicate with client.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// Speed test host user name.
        /// </summary>
        public readonly string? User;

        [OutputConstructor]
        private SystemSpeedtestserverHost(
            int? id,

            string? ip,

            string? password,

            int? port,

            string? user)
        {
            Id = id;
            Ip = ip;
            Password = password;
            Port = port;
            User = user;
        }
    }
}
