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

    public sealed class VirtualwanlinkServiceInternetServiceAppCtrlGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application control based Internet Service group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public VirtualwanlinkServiceInternetServiceAppCtrlGroupArgs()
        {
        }
        public static new VirtualwanlinkServiceInternetServiceAppCtrlGroupArgs Empty => new VirtualwanlinkServiceInternetServiceAppCtrlGroupArgs();
    }
}