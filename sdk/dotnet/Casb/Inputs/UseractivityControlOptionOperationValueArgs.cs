// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Casb.Inputs
{

    public sealed class UseractivityControlOptionOperationValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Operation value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public UseractivityControlOptionOperationValueArgs()
        {
        }
        public static new UseractivityControlOptionOperationValueArgs Empty => new UseractivityControlOptionOperationValueArgs();
    }
}
