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

    public sealed class SystemSdwanDuplicationDstaddrArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address6 or address6 group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SystemSdwanDuplicationDstaddrArgs()
        {
        }
        public static new SystemSdwanDuplicationDstaddrArgs Empty => new SystemSdwanDuplicationDstaddrArgs();
    }
}