// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Ssl.Web.Inputs
{

    public sealed class UserbookmarkBookmarkFormDataGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public UserbookmarkBookmarkFormDataGetArgs()
        {
        }
        public static new UserbookmarkBookmarkFormDataGetArgs Empty => new UserbookmarkBookmarkFormDataGetArgs();
    }
}
