// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Inputs
{

    public sealed class AutomationactionEmailToArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Email address.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public AutomationactionEmailToArgs()
        {
        }
        public static new AutomationactionEmailToArgs Empty => new AutomationactionEmailToArgs();
    }
}
