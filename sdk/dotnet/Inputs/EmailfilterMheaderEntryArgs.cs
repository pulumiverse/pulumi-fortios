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

    public sealed class EmailfilterMheaderEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Mark spam or good. Valid values: `spam`, `clear`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Pattern for the header field body.
        /// </summary>
        [Input("fieldbody")]
        public Input<string>? Fieldbody { get; set; }

        /// <summary>
        /// Pattern for header field name.
        /// </summary>
        [Input("fieldname")]
        public Input<string>? Fieldname { get; set; }

        /// <summary>
        /// Mime header entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Wildcard pattern or regular expression. Valid values: `wildcard`, `regexp`.
        /// </summary>
        [Input("patternType")]
        public Input<string>? PatternType { get; set; }

        /// <summary>
        /// Enable/disable status. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public EmailfilterMheaderEntryArgs()
        {
        }
        public static new EmailfilterMheaderEntryArgs Empty => new EmailfilterMheaderEntryArgs();
    }
}
