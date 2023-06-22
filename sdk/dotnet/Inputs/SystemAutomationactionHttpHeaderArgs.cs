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

    public sealed class SystemAutomationactionHttpHeaderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Request header key.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Request header value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public SystemAutomationactionHttpHeaderArgs()
        {
        }
        public static new SystemAutomationactionHttpHeaderArgs Empty => new SystemAutomationactionHttpHeaderArgs();
    }
}
