// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Bgp.Inputs
{

    public sealed class NeighborConditionalAdvertise6Args : global::Pulumi.ResourceArgs
    {
        [Input("advertiseRoutemap")]
        public Input<string>? AdvertiseRoutemap { get; set; }

        [Input("conditionRoutemap")]
        public Input<string>? ConditionRoutemap { get; set; }

        [Input("conditionType")]
        public Input<string>? ConditionType { get; set; }

        public NeighborConditionalAdvertise6Args()
        {
        }
        public static new NeighborConditionalAdvertise6Args Empty => new NeighborConditionalAdvertise6Args();
    }
}
