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

    public sealed class FirewallProxyaddrgrpTaggingTagGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Tag name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FirewallProxyaddrgrpTaggingTagGetArgs()
        {
        }
        public static new FirewallProxyaddrgrpTaggingTagGetArgs Empty => new FirewallProxyaddrgrpTaggingTagGetArgs();
    }
}
