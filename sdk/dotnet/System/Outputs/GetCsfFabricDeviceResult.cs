// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Outputs
{

    [OutputType]
    public sealed class GetCsfFabricDeviceResult
    {
        /// <summary>
        /// Device access token.
        /// </summary>
        public readonly string AccessToken;
        /// <summary>
        /// Device IP.
        /// </summary>
        public readonly string DeviceIp;
        /// <summary>
        /// Device type.
        /// </summary>
        public readonly string DeviceType;
        /// <summary>
        /// HTTPS port for fabric device.
        /// </summary>
        public readonly int HttpsPort;
        /// <summary>
        /// Device login name.
        /// </summary>
        public readonly string Login;
        /// <summary>
        /// Device name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Device login password.
        /// </summary>
        public readonly string Password;

        [OutputConstructor]
        private GetCsfFabricDeviceResult(
            string accessToken,

            string deviceIp,

            string deviceType,

            int httpsPort,

            string login,

            string name,

            string password)
        {
            AccessToken = accessToken;
            DeviceIp = deviceIp;
            DeviceType = deviceType;
            HttpsPort = httpsPort;
            Login = login;
            Name = name;
            Password = password;
        }
    }
}
