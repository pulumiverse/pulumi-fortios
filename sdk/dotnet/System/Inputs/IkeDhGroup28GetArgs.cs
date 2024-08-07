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

    public sealed class IkeDhGroup28GetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.
        /// </summary>
        [Input("keypairCache")]
        public Input<string>? KeypairCache { get; set; }

        /// <summary>
        /// Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
        /// </summary>
        [Input("keypairCount")]
        public Input<int>? KeypairCount { get; set; }

        /// <summary>
        /// Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        public IkeDhGroup28GetArgs()
        {
        }
        public static new IkeDhGroup28GetArgs Empty => new IkeDhGroup28GetArgs();
    }
}
