// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Inputs
{

    public sealed class MpskprofileMpskGroupMpskKeyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// MPSK client limit type options. Valid values: `default`, `unlimited`, `specified`.
        /// </summary>
        [Input("concurrentClientLimitType")]
        public Input<string>? ConcurrentClientLimitType { get; set; }

        /// <summary>
        /// Number of clients that can connect using this pre-shared key (1 - 65535, default is 256).
        /// </summary>
        [Input("concurrentClients")]
        public Input<int>? ConcurrentClients { get; set; }

        /// <summary>
        /// Select the type of the key. Valid values: `wpa2-personal`, `wpa3-sae`.
        /// </summary>
        [Input("keyType")]
        public Input<string>? KeyType { get; set; }

        /// <summary>
        /// MAC address.
        /// </summary>
        [Input("mac")]
        public Input<string>? Mac { get; set; }

        [Input("mpskSchedules")]
        private InputList<Inputs.MpskprofileMpskGroupMpskKeyMpskScheduleGetArgs>? _mpskSchedules;

        /// <summary>
        /// Firewall schedule for MPSK passphrase. The passphrase will be effective only when at least one schedule is valid. The structure of `mpsk_schedules` block is documented below.
        /// </summary>
        public InputList<Inputs.MpskprofileMpskGroupMpskKeyMpskScheduleGetArgs> MpskSchedules
        {
            get => _mpskSchedules ?? (_mpskSchedules = new InputList<Inputs.MpskprofileMpskGroupMpskKeyMpskScheduleGetArgs>());
            set => _mpskSchedules = value;
        }

        /// <summary>
        /// Pre-shared key name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("passphrase")]
        private Input<string>? _passphrase;

        /// <summary>
        /// WPA Pre-shared key.
        /// </summary>
        public Input<string>? Passphrase
        {
            get => _passphrase;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _passphrase = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// WPA3 SAE password.
        /// </summary>
        [Input("saePassword")]
        public Input<string>? SaePassword { get; set; }

        /// <summary>
        /// Enable/disable WPA3 SAE-PK (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("saePk")]
        public Input<string>? SaePk { get; set; }

        /// <summary>
        /// Private key used for WPA3 SAE-PK authentication.
        /// </summary>
        [Input("saePrivateKey")]
        public Input<string>? SaePrivateKey { get; set; }

        public MpskprofileMpskGroupMpskKeyGetArgs()
        {
        }
        public static new MpskprofileMpskGroupMpskKeyGetArgs Empty => new MpskprofileMpskGroupMpskKeyGetArgs();
    }
}
