// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Inputs
{

    public sealed class PolicySortStatePolicyListArgs : global::Pulumi.ResourceArgs
    {
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policyid")]
        public Input<string>? Policyid { get; set; }

        public PolicySortStatePolicyListArgs()
        {
        }
        public static new PolicySortStatePolicyListArgs Empty => new PolicySortStatePolicyListArgs();
    }
}
