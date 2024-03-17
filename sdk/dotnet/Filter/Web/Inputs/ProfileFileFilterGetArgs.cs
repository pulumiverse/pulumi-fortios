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

    public sealed class ProfileFileFilterGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("entries")]
        private InputList<Inputs.ProfileFileFilterEntryGetArgs>? _entries;

        /// <summary>
        /// File filter entries. The structure of `entries` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileFileFilterEntryGetArgs> Entries
        {
            get => _entries ?? (_entries = new InputList<Inputs.ProfileFileFilterEntryGetArgs>());
            set => _entries = value;
        }

        /// <summary>
        /// Enable/disable file filter logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// Enable/disable file filter archive contents scan. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("scanArchiveContents")]
        public Input<string>? ScanArchiveContents { get; set; }

        /// <summary>
        /// Enable/disable file filter. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ProfileFileFilterGetArgs()
        {
        }
        public static new ProfileFileFilterGetArgs Empty => new ProfileFileFilterGetArgs();
    }
}
