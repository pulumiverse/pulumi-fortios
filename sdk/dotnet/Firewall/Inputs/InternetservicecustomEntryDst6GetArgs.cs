// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Inputs
{

    public sealed class InternetservicecustomEntryDst6GetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Internet Service name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public InternetservicecustomEntryDst6GetArgs()
        {
        }
        public static new InternetservicecustomEntryDst6GetArgs Empty => new InternetservicecustomEntryDst6GetArgs();
    }
}