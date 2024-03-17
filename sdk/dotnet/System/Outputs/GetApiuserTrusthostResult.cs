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
    public sealed class GetApiuserTrusthostResult
    {
        /// <summary>
        /// Table ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// IPv4 trusted host address.
        /// </summary>
        public readonly string Ipv4Trusthost;
        /// <summary>
        /// IPv6 trusted host address.
        /// </summary>
        public readonly string Ipv6Trusthost;
        /// <summary>
        /// Trusthost type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetApiuserTrusthostResult(
            int id,

            string ipv4Trusthost,

            string ipv6Trusthost,

            string type)
        {
            Id = id;
            Ipv4Trusthost = ipv4Trusthost;
            Ipv6Trusthost = ipv6Trusthost;
            Type = type;
        }
    }
}
