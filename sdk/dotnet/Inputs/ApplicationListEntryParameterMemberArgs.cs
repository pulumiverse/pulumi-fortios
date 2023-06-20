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

    public sealed class ApplicationListEntryParameterMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Parameter.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Parameter name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Parameter value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public ApplicationListEntryParameterMemberArgs()
        {
        }
        public static new ApplicationListEntryParameterMemberArgs Empty => new ApplicationListEntryParameterMemberArgs();
    }
}