// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Outputs
{

    [OutputType]
    public sealed class RoutemapRuleSetExtcommunitySoo
    {
        /// <summary>
        /// Community (format = AA:NN).
        /// </summary>
        public readonly string? Community;

        [OutputConstructor]
        private RoutemapRuleSetExtcommunitySoo(string? community)
        {
            Community = community;
        }
    }
}
