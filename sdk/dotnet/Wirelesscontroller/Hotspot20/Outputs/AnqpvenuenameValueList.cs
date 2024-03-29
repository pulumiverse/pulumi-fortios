// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Hotspot20.Outputs
{

    [OutputType]
    public sealed class AnqpvenuenameValueList
    {
        /// <summary>
        /// Value index.
        /// </summary>
        public readonly int? Index;
        /// <summary>
        /// Language code.
        /// </summary>
        public readonly string? Lang;
        /// <summary>
        /// Venue name value.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private AnqpvenuenameValueList(
            int? index,

            string? lang,

            string? value)
        {
            Index = index;
            Lang = lang;
            Value = value;
        }
    }
}
