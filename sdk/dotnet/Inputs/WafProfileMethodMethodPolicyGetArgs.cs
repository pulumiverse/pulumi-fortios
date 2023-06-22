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

    public sealed class WafProfileMethodMethodPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Host address.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Allowed Methods. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`, `others`.
        /// </summary>
        [Input("allowedMethods")]
        public Input<string>? AllowedMethods { get; set; }

        /// <summary>
        /// HTTP method policy ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// URL pattern.
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        /// <summary>
        /// Enable/disable regular expression based pattern match. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("regex")]
        public Input<string>? Regex { get; set; }

        public WafProfileMethodMethodPolicyGetArgs()
        {
        }
        public static new WafProfileMethodMethodPolicyGetArgs Empty => new WafProfileMethodMethodPolicyGetArgs();
    }
}
