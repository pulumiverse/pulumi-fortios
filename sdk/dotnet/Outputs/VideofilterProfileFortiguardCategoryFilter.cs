// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class VideofilterProfileFortiguardCategoryFilter
    {
        /// <summary>
        /// VideoFilter action. Valid values: `allow`, `monitor`, `block`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// Category ID.
        /// </summary>
        public readonly int? CategoryId;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Enable/disable logging. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Log;

        [OutputConstructor]
        private VideofilterProfileFortiguardCategoryFilter(
            string? action,

            int? categoryId,

            int? id,

            string? log)
        {
            Action = action;
            CategoryId = categoryId;
            Id = id;
            Log = log;
        }
    }
}
