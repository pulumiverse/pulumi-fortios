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

    public sealed class SessionttlPortArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// End port number.
        /// </summary>
        [Input("endPort")]
        public Input<int>? EndPort { get; set; }

        /// <summary>
        /// Table entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Protocol (0 - 255).
        /// </summary>
        [Input("protocol")]
        public Input<int>? Protocol { get; set; }

        /// <summary>
        /// Start port number.
        /// </summary>
        [Input("startPort")]
        public Input<int>? StartPort { get; set; }

        /// <summary>
        /// Session timeout (TTL).
        /// </summary>
        [Input("timeout")]
        public Input<string>? Timeout { get; set; }

        public SessionttlPortArgs()
        {
        }
        public static new SessionttlPortArgs Empty => new SessionttlPortArgs();
    }
}
