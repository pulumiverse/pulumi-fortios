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
    public sealed class GetRoutemapRuleSetAspathResult
    {
        /// <summary>
        /// AS number (0 - 42949672). NOTE: Use quotes for repeating numbers, e.g.: "1 1 2"
        /// </summary>
        public readonly string As;

        [OutputConstructor]
        private GetRoutemapRuleSetAspathResult(string @as)
        {
            As = @as;
        }
    }
}
