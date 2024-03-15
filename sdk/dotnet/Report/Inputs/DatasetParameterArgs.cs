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

    public sealed class DatasetParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Data type. Valid values: `text`, `integer`, `double`, `long-integer`, `date-time`.
        /// </summary>
        [Input("dataType")]
        public Input<string>? DataType { get; set; }

        /// <summary>
        /// Display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// SQL field name.
        /// </summary>
        [Input("field")]
        public Input<string>? Field { get; set; }

        /// <summary>
        /// Parameter ID (1 to number of columns in SQL result).
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        public DatasetParameterArgs()
        {
        }
        public static new DatasetParameterArgs Empty => new DatasetParameterArgs();
    }
}