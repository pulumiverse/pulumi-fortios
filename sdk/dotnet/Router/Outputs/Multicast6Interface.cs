// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Outputs
{

    [OutputType]
    public sealed class Multicast6Interface
    {
        /// <summary>
        /// Time before old neighbor information expires in seconds (1 - 65535, default = 105).
        /// </summary>
        public readonly int? HelloHoldtime;
        /// <summary>
        /// Interval between sending PIM hello messages in seconds (1 - 65535, default = 30).
        /// </summary>
        public readonly int? HelloInterval;
        /// <summary>
        /// Interface name.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private Multicast6Interface(
            int? helloHoldtime,

            int? helloInterval,

            string? name)
        {
            HelloHoldtime = helloHoldtime;
            HelloInterval = helloInterval;
            Name = name;
        }
    }
}
