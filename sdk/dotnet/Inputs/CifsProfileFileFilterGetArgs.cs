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

    public sealed class CifsProfileFileFilterGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("entries")]
        private InputList<Inputs.CifsProfileFileFilterEntryGetArgs>? _entries;

        /// <summary>
        /// File filter entries. The structure of `entries` block is documented below.
        /// </summary>
        public InputList<Inputs.CifsProfileFileFilterEntryGetArgs> Entries
        {
            get => _entries ?? (_entries = new InputList<Inputs.CifsProfileFileFilterEntryGetArgs>());
            set => _entries = value;
        }

        /// <summary>
        /// Enable/disable file filter logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// Enable/disable file filter. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public CifsProfileFileFilterGetArgs()
        {
        }
        public static new CifsProfileFileFilterGetArgs Empty => new CifsProfileFileFilterGetArgs();
    }
}
