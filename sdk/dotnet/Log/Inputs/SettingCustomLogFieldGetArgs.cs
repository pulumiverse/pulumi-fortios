// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Log.Inputs
{

    public sealed class SettingCustomLogFieldGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom log field.
        /// </summary>
        [Input("fieldId")]
        public Input<string>? FieldId { get; set; }

        public SettingCustomLogFieldGetArgs()
        {
        }
        public static new SettingCustomLogFieldGetArgs Empty => new SettingCustomLogFieldGetArgs();
    }
}