// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Service.Inputs
{

    public sealed class CustomApplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application id.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        public CustomApplicationArgs()
        {
        }
        public static new CustomApplicationArgs Empty => new CustomApplicationArgs();
    }
}
