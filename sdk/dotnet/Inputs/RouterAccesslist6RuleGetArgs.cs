// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class RouterAccesslist6RuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Permit or deny this IP address and netmask prefix. Valid values: `permit`, `deny`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Enable/disable exact prefix match. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("exactMatch")]
        public Input<string>? ExactMatch { get; set; }

        /// <summary>
        /// Flags.
        /// </summary>
        [Input("flags")]
        public Input<int>? Flags { get; set; }

        /// <summary>
        /// Rule ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// IPv6 prefix to define regular filter criteria, such as "any" or subnets.
        /// </summary>
        [Input("prefix6")]
        public Input<string>? Prefix6 { get; set; }

        public RouterAccesslist6RuleGetArgs()
        {
        }
        public static new RouterAccesslist6RuleGetArgs Empty => new RouterAccesslist6RuleGetArgs();
    }
}
