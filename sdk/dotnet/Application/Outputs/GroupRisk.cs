// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Application.Outputs
{

    [OutputType]
    public sealed class GroupRisk
    {
        /// <summary>
        /// Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).
        /// </summary>
        public readonly int? Level;

        [OutputConstructor]
        private GroupRisk(int? level)
        {
            Level = level;
        }
    }
}
