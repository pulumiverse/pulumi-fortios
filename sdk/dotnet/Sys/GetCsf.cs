// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys
{
    public static class GetCsf
    {
        /// <summary>
        /// Use this data source to get information on fortios system csf
        /// </summary>
        public static Task<GetCsfResult> InvokeAsync(GetCsfArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCsfResult>("fortios:sys/getCsf:getCsf", args ?? new GetCsfArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system csf
        /// </summary>
        public static Output<GetCsfResult> Invoke(GetCsfInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCsfResult>("fortios:sys/getCsf:getCsf", args ?? new GetCsfInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCsfArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetCsfArgs()
        {
        }
        public static new GetCsfArgs Empty => new GetCsfArgs();
    }

    public sealed class GetCsfInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetCsfInvokeArgs()
        {
        }
        public static new GetCsfInvokeArgs Empty => new GetCsfInvokeArgs();
    }


    [OutputType]
    public sealed class GetCsfResult
    {
        /// <summary>
        /// Accept connections with unknown certificates and ask admin for approval.
        /// </summary>
        public readonly string AcceptAuthByCert;
        /// <summary>
        /// Authorization request type.
        /// </summary>
        public readonly string AuthorizationRequestType;
        /// <summary>
        /// Certificate.
        /// </summary>
        public readonly string Certificate;
        /// <summary>
        /// Configuration sync mode.
        /// </summary>
        public readonly string ConfigurationSync;
        /// <summary>
        /// Enable/disable downstream device access to this device's configuration and data.
        /// </summary>
        public readonly string DownstreamAccess;
        /// <summary>
        /// Default access profile for requests from downstream devices.
        /// </summary>
        public readonly string DownstreamAccprofile;
        /// <summary>
        /// Fabric connector configuration. The structure of `fabric_connector` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCsfFabricConnectorResult> FabricConnectors;
        /// <summary>
        /// Fabric device configuration. The structure of `fabric_device` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCsfFabricDeviceResult> FabricDevices;
        /// <summary>
        /// Fabric CMDB Object Unification
        /// </summary>
        public readonly string FabricObjectUnification;
        /// <summary>
        /// Number of worker processes for Security Fabric daemon.
        /// </summary>
        public readonly int FabricWorkers;
        /// <summary>
        /// Auto-generated fixed key used when this device is the root. (Will automatically be generated if not set.)
        /// </summary>
        public readonly string FixedKey;
        /// <summary>
        /// Fabric FortiCloud account unification.
        /// </summary>
        public readonly string ForticloudAccountEnforcement;
        /// <summary>
        /// Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
        /// </summary>
        public readonly string GroupName;
        /// <summary>
        /// Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
        /// </summary>
        public readonly string GroupPassword;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Enable/disable broadcast of discovery messages for log unification.
        /// </summary>
        public readonly string LogUnification;
        /// <summary>
        /// Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
        /// </summary>
        public readonly string ManagementIp;
        /// <summary>
        /// Overriding port for management connection (Overrides admin port).
        /// </summary>
        public readonly int ManagementPort;
        /// <summary>
        /// SAML setting configuration synchronization.
        /// </summary>
        public readonly string SamlConfigurationSync;
        /// <summary>
        /// Enable/disable Security Fabric.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Pre-authorized and blocked security fabric nodes. The structure of `trusted_list` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCsfTrustedListResult> TrustedLists;
        /// <summary>
        /// IP/FQDN of the FortiGate upstream from this FortiGate in the Security Fabric.
        /// </summary>
        public readonly string Upstream;
        /// <summary>
        /// IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
        /// </summary>
        public readonly string UpstreamIp;
        /// <summary>
        /// The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
        /// </summary>
        public readonly int UpstreamPort;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetCsfResult(
            string acceptAuthByCert,

            string authorizationRequestType,

            string certificate,

            string configurationSync,

            string downstreamAccess,

            string downstreamAccprofile,

            ImmutableArray<Outputs.GetCsfFabricConnectorResult> fabricConnectors,

            ImmutableArray<Outputs.GetCsfFabricDeviceResult> fabricDevices,

            string fabricObjectUnification,

            int fabricWorkers,

            string fixedKey,

            string forticloudAccountEnforcement,

            string groupName,

            string groupPassword,

            string id,

            string logUnification,

            string managementIp,

            int managementPort,

            string samlConfigurationSync,

            string status,

            ImmutableArray<Outputs.GetCsfTrustedListResult> trustedLists,

            string upstream,

            string upstreamIp,

            int upstreamPort,

            string? vdomparam)
        {
            AcceptAuthByCert = acceptAuthByCert;
            AuthorizationRequestType = authorizationRequestType;
            Certificate = certificate;
            ConfigurationSync = configurationSync;
            DownstreamAccess = downstreamAccess;
            DownstreamAccprofile = downstreamAccprofile;
            FabricConnectors = fabricConnectors;
            FabricDevices = fabricDevices;
            FabricObjectUnification = fabricObjectUnification;
            FabricWorkers = fabricWorkers;
            FixedKey = fixedKey;
            ForticloudAccountEnforcement = forticloudAccountEnforcement;
            GroupName = groupName;
            GroupPassword = groupPassword;
            Id = id;
            LogUnification = logUnification;
            ManagementIp = managementIp;
            ManagementPort = managementPort;
            SamlConfigurationSync = samlConfigurationSync;
            Status = status;
            TrustedLists = trustedLists;
            Upstream = upstream;
            UpstreamIp = upstreamIp;
            UpstreamPort = upstreamPort;
            Vdomparam = vdomparam;
        }
    }
}
