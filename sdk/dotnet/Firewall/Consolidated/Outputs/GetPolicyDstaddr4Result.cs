// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Consolidated.Outputs
{

    [OutputType]
    public sealed class GetPolicyDstaddr4Result
    {
        /// <summary>
        /// Application group names.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetPolicyDstaddr4Result(string name)
        {
            Name = name;
        }
    }
}
