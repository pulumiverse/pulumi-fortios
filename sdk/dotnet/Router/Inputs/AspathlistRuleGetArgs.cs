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

    public sealed class AspathlistRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Permit or deny route-based operations, based on the route's AS_PATH attribute. Valid values: `deny`, `permit`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Regular-expression to match the Border Gateway Protocol (BGP) AS paths.
        /// </summary>
        [Input("regexp")]
        public Input<string>? Regexp { get; set; }

        public AspathlistRuleGetArgs()
        {
        }
        public static new AspathlistRuleGetArgs Empty => new AspathlistRuleGetArgs();
    }
}
