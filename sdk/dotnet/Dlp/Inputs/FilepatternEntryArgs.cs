// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Dlp.Inputs
{

    public sealed class FilepatternEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Select a file type.
        /// </summary>
        [Input("fileType")]
        public Input<string>? FileType { get; set; }

        /// <summary>
        /// Filter by file name pattern or by file type. Valid values: `pattern`, `type`.
        /// </summary>
        [Input("filterType")]
        public Input<string>? FilterType { get; set; }

        /// <summary>
        /// Add a file name pattern.
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        public FilepatternEntryArgs()
        {
        }
        public static new FilepatternEntryArgs Empty => new FilepatternEntryArgs();
    }
}
