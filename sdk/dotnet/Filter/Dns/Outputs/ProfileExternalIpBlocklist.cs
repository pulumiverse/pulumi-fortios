// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Dns.Outputs
{

    [OutputType]
    public sealed class ProfileExternalIpBlocklist
    {
        /// <summary>
        /// External domain block list name.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private ProfileExternalIpBlocklist(string? name)
        {
            Name = name;
        }
    }
}