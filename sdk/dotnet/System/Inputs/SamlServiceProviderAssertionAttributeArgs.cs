// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Inputs
{

    public sealed class SamlServiceProviderAssertionAttributeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public SamlServiceProviderAssertionAttributeArgs()
        {
        }
        public static new SamlServiceProviderAssertionAttributeArgs Empty => new SamlServiceProviderAssertionAttributeArgs();
    }
}