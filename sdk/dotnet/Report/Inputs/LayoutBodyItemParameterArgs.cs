// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Report.Inputs
{

    public sealed class LayoutBodyItemParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Field name that match field of parameters defined in dataset.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Value to replace corresponding field of parameters defined in dataset.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public LayoutBodyItemParameterArgs()
        {
        }
        public static new LayoutBodyItemParameterArgs Empty => new LayoutBodyItemParameterArgs();
    }
}
