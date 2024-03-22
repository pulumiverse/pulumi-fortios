// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System
{
    public static class GetCentralmanagement
    {
        /// <summary>
        /// Use this data source to get information on fortios system centralmanagement
        /// </summary>
        public static Task<GetCentralmanagementResult> InvokeAsync(GetCentralmanagementArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCentralmanagementResult>("fortios:system/getCentralmanagement:getCentralmanagement", args ?? new GetCentralmanagementArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system centralmanagement
        /// </summary>
        public static Output<GetCentralmanagementResult> Invoke(GetCentralmanagementInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCentralmanagementResult>("fortios:system/getCentralmanagement:getCentralmanagement", args ?? new GetCentralmanagementInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCentralmanagementArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetCentralmanagementArgs()
        {
        }
        public static new GetCentralmanagementArgs Empty => new GetCentralmanagementArgs();
    }

    public sealed class GetCentralmanagementInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetCentralmanagementInvokeArgs()
        {
        }
        public static new GetCentralmanagementInvokeArgs Empty => new GetCentralmanagementInvokeArgs();
    }


    [OutputType]
    public sealed class GetCentralmanagementResult
    {
        /// <summary>
        /// Enable/disable allowing the central management server to remotely monitor this FortiGate
        /// </summary>
        public readonly string AllowMonitor;
        /// <summary>
        /// Enable/disable allowing the central management server to push configuration changes to this FortiGate.
        /// </summary>
        public readonly string AllowPushConfiguration;
        /// <summary>
        /// Enable/disable allowing the central management server to push firmware updates to this FortiGate.
        /// </summary>
        public readonly string AllowPushFirmware;
        /// <summary>
        /// Enable/disable remotely upgrading the firmware on this FortiGate from the central management server.
        /// </summary>
        public readonly string AllowRemoteFirmwareUpgrade;
        /// <summary>
        /// CA certificate to be used by FGFM protocol.
        /// </summary>
        public readonly string CaCert;
        /// <summary>
        /// Encryption strength for communications between the FortiGate and central management.
        /// </summary>
        public readonly string EncAlgorithm;
        /// <summary>
        /// IP address or FQDN of the FortiManager.
        /// </summary>
        public readonly string Fmg;
        /// <summary>
        /// IPv4 source address that this FortiGate uses when communicating with FortiManager.
        /// </summary>
        public readonly string FmgSourceIp;
        /// <summary>
        /// IPv6 source address that this FortiGate uses when communicating with FortiManager.
        /// </summary>
        public readonly string FmgSourceIp6;
        /// <summary>
        /// Port used to communicate with FortiManager that is acting as a FortiGuard update server.
        /// </summary>
        public readonly string FmgUpdatePort;
        /// <summary>
        /// Override access profile.
        /// </summary>
        public readonly string FortigateCloudSsoDefaultProfile;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Enable/disable inclusion of public FortiGuard servers in the override server list.
        /// </summary>
        public readonly string IncludeDefaultServers;
        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        public readonly string Interface;
        /// <summary>
        /// Specify how to select outgoing interface to reach server.
        /// </summary>
        public readonly string InterfaceSelectMethod;
        /// <summary>
        /// Certificate to be used by FGFM protocol.
        /// </summary>
        public readonly string LocalCert;
        /// <summary>
        /// Central management mode.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// Enable/disable allowing the central management server to restore the configuration of this FortiGate.
        /// </summary>
        public readonly string ScheduleConfigRestore;
        /// <summary>
        /// Enable/disable allowing the central management server to restore the scripts stored on this FortiGate.
        /// </summary>
        public readonly string ScheduleScriptRestore;
        /// <summary>
        /// Serial number.
        /// </summary>
        public readonly string SerialNumber;
        /// <summary>
        /// Additional severs that the FortiGate can use for updates (for AV, IPS, updates) and ratings (for web filter and antispam ratings) servers. The structure of `server_list` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCentralmanagementServerListResult> ServerLists;
        /// <summary>
        /// Central management type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Virtual domain (VDOM) name to use when communicating with FortiManager.
        /// </summary>
        public readonly string Vdom;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetCentralmanagementResult(
            string allowMonitor,

            string allowPushConfiguration,

            string allowPushFirmware,

            string allowRemoteFirmwareUpgrade,

            string caCert,

            string encAlgorithm,

            string fmg,

            string fmgSourceIp,

            string fmgSourceIp6,

            string fmgUpdatePort,

            string fortigateCloudSsoDefaultProfile,

            string id,

            string includeDefaultServers,

            string @interface,

            string interfaceSelectMethod,

            string localCert,

            string mode,

            string scheduleConfigRestore,

            string scheduleScriptRestore,

            string serialNumber,

            ImmutableArray<Outputs.GetCentralmanagementServerListResult> serverLists,

            string type,

            string vdom,

            string? vdomparam)
        {
            AllowMonitor = allowMonitor;
            AllowPushConfiguration = allowPushConfiguration;
            AllowPushFirmware = allowPushFirmware;
            AllowRemoteFirmwareUpgrade = allowRemoteFirmwareUpgrade;
            CaCert = caCert;
            EncAlgorithm = encAlgorithm;
            Fmg = fmg;
            FmgSourceIp = fmgSourceIp;
            FmgSourceIp6 = fmgSourceIp6;
            FmgUpdatePort = fmgUpdatePort;
            FortigateCloudSsoDefaultProfile = fortigateCloudSsoDefaultProfile;
            Id = id;
            IncludeDefaultServers = includeDefaultServers;
            Interface = @interface;
            InterfaceSelectMethod = interfaceSelectMethod;
            LocalCert = localCert;
            Mode = mode;
            ScheduleConfigRestore = scheduleConfigRestore;
            ScheduleScriptRestore = scheduleScriptRestore;
            SerialNumber = serialNumber;
            ServerLists = serverLists;
            Type = type;
            Vdom = vdom;
            Vdomparam = vdomparam;
        }
    }
}
