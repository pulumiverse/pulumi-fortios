// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Casb.Inputs
{

    public sealed class UseractivityMatchRuleDomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Domain list separated by space.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        public UseractivityMatchRuleDomainArgs()
        {
        }
        public static new UseractivityMatchRuleDomainArgs Empty => new UseractivityMatchRuleDomainArgs();
    }
}
