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
    public sealed class GetBgpNeighborConditionalAdvertise6Result
    {
        /// <summary>
        /// Name of advertising route map.
        /// </summary>
        public readonly string AdvertiseRoutemap;
        /// <summary>
        /// Name of condition route map.
        /// </summary>
        public readonly string ConditionRoutemap;
        /// <summary>
        /// Type of condition.
        /// </summary>
        public readonly string ConditionType;

        [OutputConstructor]
        private GetBgpNeighborConditionalAdvertise6Result(
            string advertiseRoutemap,

            string conditionRoutemap,

            string conditionType)
        {
            AdvertiseRoutemap = advertiseRoutemap;
            ConditionRoutemap = conditionRoutemap;
            ConditionType = conditionType;
        }
    }
}
