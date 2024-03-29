// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Inputs
{

    public sealed class KeychainKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Lifetime of received authentication key (format: hh:mm:ss day month year).
        /// </summary>
        [Input("acceptLifetime")]
        public Input<string>? AcceptLifetime { get; set; }

        /// <summary>
        /// Cryptographic algorithm.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// Key ID (0 - 2147483647).
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        [Input("keyString")]
        private Input<string>? _keyString;

        /// <summary>
        /// Password for the key. On FortiOS versions 6.2.0-7.0.0: max. = 35 characters. On FortiOS versions 7.0.1-7.0.3: maximum = 64 characters.
        /// </summary>
        public Input<string>? KeyString
        {
            get => _keyString;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _keyString = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Lifetime of sent authentication key (format: hh:mm:ss day month year).
        /// </summary>
        [Input("sendLifetime")]
        public Input<string>? SendLifetime { get; set; }

        public KeychainKeyArgs()
        {
        }
        public static new KeychainKeyArgs Empty => new KeychainKeyArgs();
    }
}
