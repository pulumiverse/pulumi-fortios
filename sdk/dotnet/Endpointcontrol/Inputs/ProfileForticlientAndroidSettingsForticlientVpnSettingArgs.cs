// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Endpointcontrol.Inputs
{

    public sealed class ProfileForticlientAndroidSettingsForticlientVpnSettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication method. Valid values: `psk`, `certificate`.
        /// </summary>
        [Input("authMethod")]
        public Input<string>? AuthMethod { get; set; }

        /// <summary>
        /// VPN name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("presharedKey")]
        private Input<string>? _presharedKey;

        /// <summary>
        /// Pre-shared secret for PSK authentication.
        /// </summary>
        public Input<string>? PresharedKey
        {
            get => _presharedKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _presharedKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// IP address or FQDN of the remote VPN gateway.
        /// </summary>
        [Input("remoteGw")]
        public Input<string>? RemoteGw { get; set; }

        /// <summary>
        /// SSL VPN access port (1 - 65535).
        /// </summary>
        [Input("sslvpnAccessPort")]
        public Input<int>? SslvpnAccessPort { get; set; }

        /// <summary>
        /// Enable/disable requiring SSL VPN client certificate. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sslvpnRequireCertificate")]
        public Input<string>? SslvpnRequireCertificate { get; set; }

        /// <summary>
        /// VPN type (IPsec or SSL VPN). Valid values: `ipsec`, `ssl`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ProfileForticlientAndroidSettingsForticlientVpnSettingArgs()
        {
        }
        public static new ProfileForticlientAndroidSettingsForticlientVpnSettingArgs Empty => new ProfileForticlientAndroidSettingsForticlientVpnSettingArgs();
    }
}
