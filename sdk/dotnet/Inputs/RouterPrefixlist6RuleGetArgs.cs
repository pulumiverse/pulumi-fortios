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

    public sealed class RouterPrefixlist6RuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Permit or deny packets that match this rule. Valid values: `permit`, `deny`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Flags.
        /// </summary>
        [Input("flags")]
        public Input<int>? Flags { get; set; }

        /// <summary>
        /// Minimum prefix length to be matched (0 - 128).
        /// </summary>
        [Input("ge")]
        public Input<int>? Ge { get; set; }

        /// <summary>
        /// Rule ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Maximum prefix length to be matched (0 - 128).
        /// </summary>
        [Input("le")]
        public Input<int>? Le { get; set; }

        /// <summary>
        /// IPv6 prefix to define regular filter criteria, such as "any" or subnets.
        /// </summary>
        [Input("prefix6")]
        public Input<string>? Prefix6 { get; set; }

        public RouterPrefixlist6RuleGetArgs()
        {
        }
        public static new RouterPrefixlist6RuleGetArgs Empty => new RouterPrefixlist6RuleGetArgs();
    }
}
