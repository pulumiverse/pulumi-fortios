// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Webproxy.Inputs
{

    public sealed class GlobalLearnClientIpSrcaddr6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GlobalLearnClientIpSrcaddr6Args()
        {
        }
        public static new GlobalLearnClientIpSrcaddr6Args Empty => new GlobalLearnClientIpSrcaddr6Args();
    }
}
