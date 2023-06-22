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
    public sealed class SystemSdnconnectorForwardingRule
    {
        /// <summary>
        /// Forwarding rule name.
        /// </summary>
        public readonly string? RuleName;
        /// <summary>
        /// Target instance name.
        /// </summary>
        public readonly string? Target;

        [OutputConstructor]
        private SystemSdnconnectorForwardingRule(
            string? ruleName,

            string? target)
        {
            RuleName = ruleName;
            Target = target;
        }
    }
}
