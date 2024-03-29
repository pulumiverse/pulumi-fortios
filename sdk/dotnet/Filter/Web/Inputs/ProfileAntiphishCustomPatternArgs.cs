// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Web.Inputs
{

    public sealed class ProfileAntiphishCustomPatternArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Category that the pattern matches. Valid values: `username`, `password`.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Target pattern.
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        /// <summary>
        /// Pattern will be treated either as a regex pattern or literal string. Valid values: `regex`, `literal`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ProfileAntiphishCustomPatternArgs()
        {
        }
        public static new ProfileAntiphishCustomPatternArgs Empty => new ProfileAntiphishCustomPatternArgs();
    }
}
