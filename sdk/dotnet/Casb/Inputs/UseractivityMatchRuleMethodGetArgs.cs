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

    public sealed class UseractivityMatchRuleMethodGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User activity method.
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        public UseractivityMatchRuleMethodGetArgs()
        {
        }
        public static new UseractivityMatchRuleMethodGetArgs Empty => new UseractivityMatchRuleMethodGetArgs();
    }
}
