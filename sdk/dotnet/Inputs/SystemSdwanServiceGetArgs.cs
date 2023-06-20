// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class SystemSdwanServiceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
        /// </summary>
        [Input("addrMode")]
        public Input<string>? AddrMode { get; set; }

        /// <summary>
        /// Coefficient of reciprocal of available bidirectional bandwidth in the formula of custom-profile-1.
        /// </summary>
        [Input("bandwidthWeight")]
        public Input<int>? BandwidthWeight { get; set; }

        /// <summary>
        /// Enable/disable use of SD-WAN as default service. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("default")]
        public Input<string>? Default { get; set; }

        /// <summary>
        /// Enable/disable forward traffic DSCP tag. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dscpForward")]
        public Input<string>? DscpForward { get; set; }

        /// <summary>
        /// Forward traffic DSCP tag.
        /// </summary>
        [Input("dscpForwardTag")]
        public Input<string>? DscpForwardTag { get; set; }

        /// <summary>
        /// Enable/disable reverse traffic DSCP tag. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dscpReverse")]
        public Input<string>? DscpReverse { get; set; }

        /// <summary>
        /// Reverse traffic DSCP tag.
        /// </summary>
        [Input("dscpReverseTag")]
        public Input<string>? DscpReverseTag { get; set; }

        [Input("dst6s")]
        private InputList<Inputs.SystemSdwanServiceDst6GetArgs>? _dst6s;

        /// <summary>
        /// Destination address6 name. The structure of `dst6` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdwanServiceDst6GetArgs> Dst6s
        {
            get => _dst6s ?? (_dst6s = new InputList<Inputs.SystemSdwanServiceDst6GetArgs>());
            set => _dst6s = value;
        }

        /// <summary>
        /// Enable/disable negation of destination address match. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dstNegate")]
        public Input<string>? DstNegate { get; set; }

        [Input("dsts")]
        private InputList<Inputs.SystemSdwanServiceDstGetArgs>? _dsts;

        /// <summary>
        /// Destination address name. The structure of `dst` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdwanServiceDstGetArgs> Dsts
        {
            get => _dsts ?? (_dsts = new InputList<Inputs.SystemSdwanServiceDstGetArgs>());
            set => _dsts = value;
        }

        /// <summary>
        /// End destination port number.
        /// </summary>
        [Input("endPort")]
        public Input<int>? EndPort { get; set; }

        /// <summary>
        /// Enable/disable SD-WAN service gateway. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        [Input("groups")]
        private InputList<Inputs.SystemSdwanServiceGroupGetArgs>? _groups;

        /// <summary>
        /// User groups. The structure of `groups` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdwanServiceGroupGetArgs> Groups
        {
            get => _groups ?? (_groups = new InputList<Inputs.SystemSdwanServiceGroupGetArgs>());
            set => _groups = value;
        }

        /// <summary>
        /// Hash algorithm for selected priority members for load balance mode. Valid values: `round-robin`, `source-ip-based`, `source-dest-ip-based`, `inbandwidth`, `outbandwidth`, `bibandwidth`.
        /// </summary>
        [Input("hashMode")]
        public Input<string>? HashMode { get; set; }

        [Input("healthChecks")]
        private InputList<Inputs.SystemSdwanServiceHealthCheckGetArgs>? _healthChecks;

        /// <summary>
        /// Health check list. The structure of `health_check` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdwanServiceHealthCheckGetArgs> HealthChecks
        {
            get => _healthChecks ?? (_healthChecks = new InputList<Inputs.SystemSdwanServiceHealthCheckGetArgs>());
            set => _healthChecks = value;
        }

        /// <summary>
        /// Waiting period in seconds when switching from the back-up member to the primary member (0 - 10000000, default = 0).
        /// </summary>
        [Input("holdDownTime")]
        public Input<int>? HoldDownTime { get; set; }

        /// <summary>
        /// SD-WAN rule ID (1 - 4000).
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Enable/disable negation of input device match. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("inputDeviceNegate")]
        public Input<string>? InputDeviceNegate { get; set; }

        [Input("inputDevices")]
        private InputList<Inputs.SystemSdwanServiceInputDeviceGetArgs>? _inputDevices;

        /// <summary>
        /// Source interface name. The structure of `input_device` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdwanServiceInputDeviceGetArgs> InputDevices
        {
            get => _inputDevices ?? (_inputDevices = new InputList<Inputs.SystemSdwanServiceInputDeviceGetArgs>());
            set => _inputDevices = value;
        }

        [Input("inputZones")]
        private InputList<Inputs.SystemSdwanServiceInputZoneGetArgs>? _inputZones;

        /// <summary>
        /// Source input-zone name. The structure of `input_zone` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdwanServiceInputZoneGetArgs> InputZones
        {
            get => _inputZones ?? (_inputZones = new InputList<Inputs.SystemSdwanServiceInputZoneGetArgs>());
            set => _inputZones = value;
        }

        /// <summary>
        /// Enable/disable use of Internet service for application-based load balancing. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("internetService")]
        public Input<string>? InternetService { get; set; }

        [Input("internetServiceAppCtrlCategories")]
        private InputList<Inputs.SystemSdwanServiceInternetServiceAppCtrlCategoryGetArgs>? _internetServiceAppCtrlCategories;

        /// <summary>
        /// IDs of one or more application control categories. The structure of `internet_service_app_ctrl_category` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdwanServiceInternetServiceAppCtrlCategoryGetArgs> InternetServiceAppCtrlCategories
        {
            get => _internetServiceAppCtrlCategories ?? (_internetServiceAppCtrlCategories = new InputList<Inputs.SystemSdwanServiceInternetServiceAppCtrlCategoryGetArgs>());
            set => _internetServiceAppCtrlCategories = value;
        }

        [Input("internetServiceAppCtrlGroups")]
        private InputList<Inputs.SystemSdwanServiceInternetServiceAppCtrlGroupGetArgs>? _internetServiceAppCtrlGroups;

        /// <summary>
        /// Application control based Internet Service group list. The structure of `internet_service_app_ctrl_group` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdwanServiceInternetServiceAppCtrlGroupGetArgs> InternetServiceAppCtrlGroups
        {
            get => _internetServiceAppCtrlGroups ?? (_internetServiceAppCtrlGroups = new InputList<Inputs.SystemSdwanServiceInternetServiceAppCtrlGroupGetArgs>());
            set => _internetServiceAppCtrlGroups = value;
        }

        [Input("internetServiceAppCtrls")]
        private InputList<Inputs.SystemSdwanServiceInternetServiceAppCtrlGetArgs>? _internetServiceAppCtrls;

        /// <summary>
        /// Application control based Internet Service ID list. The structure of `internet_service_app_ctrl` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdwanServiceInternetServiceAppCtrlGetArgs> InternetServiceAppCtrls
        {
            get => _internetServiceAppCtrls ?? (_internetServiceAppCtrls = new InputList<Inputs.SystemSdwanServiceInternetServiceAppCtrlGetArgs>());
            set => _internetServiceAppCtrls = value;
        }

        [Input("internetServiceCustomGroups")]
        private InputList<Inputs.SystemSdwanServiceInternetServiceCustomGroupGetArgs>? _internetServiceCustomGroups;

        /// <summary>
        /// Custom Internet Service group list. The structure of `internet_service_custom_group` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdwanServiceInternetServiceCustomGroupGetArgs> InternetServiceCustomGroups
        {
            get => _internetServiceCustomGroups ?? (_internetServiceCustomGroups = new InputList<Inputs.SystemSdwanServiceInternetServiceCustomGroupGetArgs>());
            set => _internetServiceCustomGroups = value;
        }

        [Input("internetServiceCustoms")]
        private InputList<Inputs.SystemSdwanServiceInternetServiceCustomGetArgs>? _internetServiceCustoms;

        /// <summary>
        /// Custom Internet service name list. The structure of `internet_service_custom` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdwanServiceInternetServiceCustomGetArgs> InternetServiceCustoms
        {
            get => _internetServiceCustoms ?? (_internetServiceCustoms = new InputList<Inputs.SystemSdwanServiceInternetServiceCustomGetArgs>());
            set => _internetServiceCustoms = value;
        }

        [Input("internetServiceGroups")]
        private InputList<Inputs.SystemSdwanServiceInternetServiceGroupGetArgs>? _internetServiceGroups;

        /// <summary>
        /// Internet Service group list. The structure of `internet_service_group` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdwanServiceInternetServiceGroupGetArgs> InternetServiceGroups
        {
            get => _internetServiceGroups ?? (_internetServiceGroups = new InputList<Inputs.SystemSdwanServiceInternetServiceGroupGetArgs>());
            set => _internetServiceGroups = value;
        }

        [Input("internetServiceNames")]
        private InputList<Inputs.SystemSdwanServiceInternetServiceNameGetArgs>? _internetServiceNames;

        /// <summary>
        /// Internet service name list. The structure of `internet_service_name` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdwanServiceInternetServiceNameGetArgs> InternetServiceNames
        {
            get => _internetServiceNames ?? (_internetServiceNames = new InputList<Inputs.SystemSdwanServiceInternetServiceNameGetArgs>());
            set => _internetServiceNames = value;
        }

        /// <summary>
        /// Coefficient of jitter in the formula of custom-profile-1.
        /// </summary>
        [Input("jitterWeight")]
        public Input<int>? JitterWeight { get; set; }

        /// <summary>
        /// Coefficient of latency in the formula of custom-profile-1.
        /// </summary>
        [Input("latencyWeight")]
        public Input<int>? LatencyWeight { get; set; }

        /// <summary>
        /// Link cost factor. Valid values: `latency`, `jitter`, `packet-loss`, `inbandwidth`, `outbandwidth`, `bibandwidth`, `custom-profile-1`.
        /// </summary>
        [Input("linkCostFactor")]
        public Input<string>? LinkCostFactor { get; set; }

        /// <summary>
        /// Percentage threshold change of link cost values that will result in policy route regeneration (0 - 10000000, default = 10).
        /// </summary>
        [Input("linkCostThreshold")]
        public Input<int>? LinkCostThreshold { get; set; }

        /// <summary>
        /// Minimum number of members which meet SLA.
        /// </summary>
        [Input("minimumSlaMeetMembers")]
        public Input<int>? MinimumSlaMeetMembers { get; set; }

        /// <summary>
        /// Control how the SD-WAN rule sets the priority of interfaces in the SD-WAN. Valid values: `auto`, `manual`, `priority`, `sla`, `load-balance`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Service and service group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Coefficient of packet-loss in the formula of custom-profile-1.
        /// </summary>
        [Input("packetLossWeight")]
        public Input<int>? PacketLossWeight { get; set; }

        /// <summary>
        /// Enable/disable passive measurement based on the service criteria. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("passiveMeasurement")]
        public Input<string>? PassiveMeasurement { get; set; }

        [Input("priorityMembers")]
        private InputList<Inputs.SystemSdwanServicePriorityMemberGetArgs>? _priorityMembers;

        /// <summary>
        /// Member sequence number list. The structure of `priority_members` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdwanServicePriorityMemberGetArgs> PriorityMembers
        {
            get => _priorityMembers ?? (_priorityMembers = new InputList<Inputs.SystemSdwanServicePriorityMemberGetArgs>());
            set => _priorityMembers = value;
        }

        [Input("priorityZones")]
        private InputList<Inputs.SystemSdwanServicePriorityZoneGetArgs>? _priorityZones;

        /// <summary>
        /// Priority zone name list. The structure of `priority_zone` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdwanServicePriorityZoneGetArgs> PriorityZones
        {
            get => _priorityZones ?? (_priorityZones = new InputList<Inputs.SystemSdwanServicePriorityZoneGetArgs>());
            set => _priorityZones = value;
        }

        /// <summary>
        /// Protocol number.
        /// </summary>
        [Input("protocol")]
        public Input<int>? Protocol { get; set; }

        /// <summary>
        /// Quality grade.
        /// </summary>
        [Input("qualityLink")]
        public Input<int>? QualityLink { get; set; }

        /// <summary>
        /// Service role to work with neighbor. Valid values: `standalone`, `primary`, `secondary`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// IPv4 route map route-tag.
        /// </summary>
        [Input("routeTag")]
        public Input<int>? RouteTag { get; set; }

        /// <summary>
        /// Method to compare SLA value for SLA mode. Valid values: `order`, `number`.
        /// </summary>
        [Input("slaCompareMethod")]
        public Input<string>? SlaCompareMethod { get; set; }

        [Input("slas")]
        private InputList<Inputs.SystemSdwanServiceSlaGetArgs>? _slas;

        /// <summary>
        /// Service level agreement (SLA). The structure of `sla` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdwanServiceSlaGetArgs> Slas
        {
            get => _slas ?? (_slas = new InputList<Inputs.SystemSdwanServiceSlaGetArgs>());
            set => _slas = value;
        }

        [Input("src6s")]
        private InputList<Inputs.SystemSdwanServiceSrc6GetArgs>? _src6s;

        /// <summary>
        /// Source address6 name. The structure of `src6` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdwanServiceSrc6GetArgs> Src6s
        {
            get => _src6s ?? (_src6s = new InputList<Inputs.SystemSdwanServiceSrc6GetArgs>());
            set => _src6s = value;
        }

        /// <summary>
        /// Enable/disable negation of source address match. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("srcNegate")]
        public Input<string>? SrcNegate { get; set; }

        [Input("srcs")]
        private InputList<Inputs.SystemSdwanServiceSrcGetArgs>? _srcs;

        /// <summary>
        /// Source address name. The structure of `src` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdwanServiceSrcGetArgs> Srcs
        {
            get => _srcs ?? (_srcs = new InputList<Inputs.SystemSdwanServiceSrcGetArgs>());
            set => _srcs = value;
        }

        /// <summary>
        /// Enable/disable service when selected neighbor role is standalone while service role is not standalone. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("standaloneAction")]
        public Input<string>? StandaloneAction { get; set; }

        /// <summary>
        /// Start destination port number.
        /// </summary>
        [Input("startPort")]
        public Input<int>? StartPort { get; set; }

        /// <summary>
        /// Enable/disable SD-WAN service. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Method of selecting member if more than one meets the SLA.
        /// </summary>
        [Input("tieBreak")]
        public Input<string>? TieBreak { get; set; }

        /// <summary>
        /// Type of service bit pattern.
        /// </summary>
        [Input("tos")]
        public Input<string>? Tos { get; set; }

        /// <summary>
        /// Type of service evaluated bits.
        /// </summary>
        [Input("tosMask")]
        public Input<string>? TosMask { get; set; }

        /// <summary>
        /// Enable/disable use of ADVPN shortcut for quality comparison. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("useShortcutSla")]
        public Input<string>? UseShortcutSla { get; set; }

        [Input("users")]
        private InputList<Inputs.SystemSdwanServiceUserGetArgs>? _users;

        /// <summary>
        /// User name. The structure of `users` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdwanServiceUserGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.SystemSdwanServiceUserGetArgs>());
            set => _users = value;
        }

        public SystemSdwanServiceGetArgs()
        {
        }
        public static new SystemSdwanServiceGetArgs Empty => new SystemSdwanServiceGetArgs();
    }
}