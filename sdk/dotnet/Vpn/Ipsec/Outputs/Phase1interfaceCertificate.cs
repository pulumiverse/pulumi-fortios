// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Ipsec.Outputs
{

    [OutputType]
    public sealed class Phase1interfaceCertificate
    {
        /// <summary>
        /// Certificate name.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private Phase1interfaceCertificate(string? name)
        {
            Name = name;
        }
    }
}
