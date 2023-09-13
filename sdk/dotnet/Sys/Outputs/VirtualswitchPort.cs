// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Outputs
{

    [OutputType]
    public sealed class VirtualswitchPort
    {
        /// <summary>
        /// Alias.
        /// </summary>
        public readonly string? Alias;
        /// <summary>
        /// Physical interface name.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private VirtualswitchPort(
            string? alias,

            string? name)
        {
            Alias = alias;
            Name = name;
        }
    }
}
