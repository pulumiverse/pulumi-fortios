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

    public sealed class SpeedtestserverHostArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Server host ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Server host IPv4 address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Speed test host password.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Server host port number to communicate with client.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Speed test host user name.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public SpeedtestserverHostArgs()
        {
        }
        public static new SpeedtestserverHostArgs Empty => new SpeedtestserverHostArgs();
    }
}
