// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Inputs
{

    public sealed class ExtcommunitylistRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Permit or deny route-based operations, based on the route's EXTENDED COMMUNITY attribute. Valid values: `deny`, `permit`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Extended community specifications for matching a reserved extended community.
        /// </summary>
        [Input("match")]
        public Input<string>? Match { get; set; }

        /// <summary>
        /// Ordered list of EXTENDED COMMUNITY attributes as a regular expression.
        /// </summary>
        [Input("regexp")]
        public Input<string>? Regexp { get; set; }

        /// <summary>
        /// Type of extended community. Valid values: `rt`, `soo`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ExtcommunitylistRuleArgs()
        {
        }
        public static new ExtcommunitylistRuleArgs Empty => new ExtcommunitylistRuleArgs();
    }
}
