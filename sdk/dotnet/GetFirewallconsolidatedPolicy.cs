// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    public static class GetFirewallconsolidatedPolicy
    {
        /// <summary>
        /// Use this data source to get information on an fortios firewallconsolidated policy
        /// </summary>
        public static Task<GetFirewallconsolidatedPolicyResult> InvokeAsync(GetFirewallconsolidatedPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallconsolidatedPolicyResult>("fortios:index/getFirewallconsolidatedPolicy:getFirewallconsolidatedPolicy", args ?? new GetFirewallconsolidatedPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios firewallconsolidated policy
        /// </summary>
        public static Output<GetFirewallconsolidatedPolicyResult> Invoke(GetFirewallconsolidatedPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallconsolidatedPolicyResult>("fortios:index/getFirewallconsolidatedPolicy:getFirewallconsolidatedPolicy", args ?? new GetFirewallconsolidatedPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallconsolidatedPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the policyid of the desired firewallconsolidated policy.
        /// </summary>
        [Input("policyid", required: true)]
        public int Policyid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetFirewallconsolidatedPolicyArgs()
        {
        }
        public static new GetFirewallconsolidatedPolicyArgs Empty => new GetFirewallconsolidatedPolicyArgs();
    }

    public sealed class GetFirewallconsolidatedPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the policyid of the desired firewallconsolidated policy.
        /// </summary>
        [Input("policyid", required: true)]
        public Input<int> Policyid { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetFirewallconsolidatedPolicyInvokeArgs()
        {
        }
        public static new GetFirewallconsolidatedPolicyInvokeArgs Empty => new GetFirewallconsolidatedPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallconsolidatedPolicyResult
    {
        /// <summary>
        /// Policy action (allow/deny/ipsec).
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Application category ID list. The structure of `app_category` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyAppCategoryResult> AppCategories;
        /// <summary>
        /// Application group names. The structure of `app_group` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyAppGroupResult> AppGroups;
        /// <summary>
        /// Name of an existing Application list.
        /// </summary>
        public readonly string ApplicationList;
        /// <summary>
        /// Application ID list. The structure of `application` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyApplicationResult> Applications;
        /// <summary>
        /// Enable/disable policy traffic ASIC offloading.
        /// </summary>
        public readonly string AutoAsicOffload;
        /// <summary>
        /// Name of an existing Antivirus profile.
        /// </summary>
        public readonly string AvProfile;
        /// <summary>
        /// Enable exemption of some users from the captive portal.
        /// </summary>
        public readonly string CaptivePortalExempt;
        /// <summary>
        /// Name of an existing CIFS profile.
        /// </summary>
        public readonly string CifsProfile;
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comments;
        /// <summary>
        /// Enable to change packet's DiffServ values to the specified diffservcode-forward value.
        /// </summary>
        public readonly string DiffservForward;
        /// <summary>
        /// Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value.
        /// </summary>
        public readonly string DiffservReverse;
        /// <summary>
        /// Change packet's DiffServ to this value.
        /// </summary>
        public readonly string DiffservcodeForward;
        /// <summary>
        /// Change packet's reverse (reply) DiffServ to this value.
        /// </summary>
        public readonly string DiffservcodeRev;
        /// <summary>
        /// Name of an existing DLP sensor.
        /// </summary>
        public readonly string DlpSensor;
        /// <summary>
        /// Name of an existing DNS filter profile.
        /// </summary>
        public readonly string DnsfilterProfile;
        /// <summary>
        /// Destination IPv4 address name and address group names. The structure of `dstaddr4` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyDstaddr4Result> Dstaddr4s;
        /// <summary>
        /// Destination IPv6 address name and address group names. The structure of `dstaddr6` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyDstaddr6Result> Dstaddr6s;
        /// <summary>
        /// When enabled dstaddr specifies what the destination address must NOT be.
        /// </summary>
        public readonly string DstaddrNegate;
        /// <summary>
        /// Outgoing (egress) interface. The structure of `dstintf` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyDstintfResult> Dstintfs;
        /// <summary>
        /// Name of an existing email filter profile.
        /// </summary>
        public readonly string EmailfilterProfile;
        /// <summary>
        /// Enable to prevent source NAT from changing a session's source port.
        /// </summary>
        public readonly string Fixedport;
        /// <summary>
        /// Names of FSSO groups. The structure of `fsso_groups` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyFssoGroupResult> FssoGroups;
        /// <summary>
        /// Names of user groups that can authenticate with this policy. The structure of `groups` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyGroupResult> Groups;
        /// <summary>
        /// Redirect HTTP(S) traffic to matching transparent web proxy policy.
        /// </summary>
        public readonly string HttpPolicyRedirect;
        /// <summary>
        /// Name of an existing ICAP profile.
        /// </summary>
        public readonly string IcapProfile;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN.
        /// </summary>
        public readonly string Inbound;
        /// <summary>
        /// Policy inspection mode (Flow/proxy). Default is Flow mode.
        /// </summary>
        public readonly string InspectionMode;
        /// <summary>
        /// Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.
        /// </summary>
        public readonly string InternetService;
        /// <summary>
        /// Custom Internet Service group name. The structure of `internet_service_custom_group` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceCustomGroupResult> InternetServiceCustomGroups;
        /// <summary>
        /// Custom Internet Service name. The structure of `internet_service_custom` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceCustomResult> InternetServiceCustoms;
        /// <summary>
        /// Internet Service group name. The structure of `internet_service_group` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceGroupResult> InternetServiceGroups;
        /// <summary>
        /// Internet Service ID. The structure of `internet_service_id` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceIdResult> InternetServiceIds;
        /// <summary>
        /// Internet Service name. The structure of `internet_service_name` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceNameResult> InternetServiceNames;
        /// <summary>
        /// When enabled internet-service specifies what the service must NOT be.
        /// </summary>
        public readonly string InternetServiceNegate;
        /// <summary>
        /// Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.
        /// </summary>
        public readonly string InternetServiceSrc;
        /// <summary>
        /// Custom Internet Service source group name. The structure of `internet_service_src_custom_group` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceSrcCustomGroupResult> InternetServiceSrcCustomGroups;
        /// <summary>
        /// Custom Internet Service source name. The structure of `internet_service_src_custom` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceSrcCustomResult> InternetServiceSrcCustoms;
        /// <summary>
        /// Internet Service source group name. The structure of `internet_service_src_group` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceSrcGroupResult> InternetServiceSrcGroups;
        /// <summary>
        /// Internet Service source ID. The structure of `internet_service_src_id` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceSrcIdResult> InternetServiceSrcIds;
        /// <summary>
        /// Internet Service source name. The structure of `internet_service_src_name` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceSrcNameResult> InternetServiceSrcNames;
        /// <summary>
        /// When enabled internet-service-src specifies what the service must NOT be.
        /// </summary>
        public readonly string InternetServiceSrcNegate;
        /// <summary>
        /// Enable to use IP Pools for source NAT.
        /// </summary>
        public readonly string Ippool;
        /// <summary>
        /// Name of an existing IPS sensor.
        /// </summary>
        public readonly string IpsSensor;
        /// <summary>
        /// Enable or disable logging. Log all sessions or security profile sessions.
        /// </summary>
        public readonly string Logtraffic;
        /// <summary>
        /// Record logs when a session starts.
        /// </summary>
        public readonly string LogtrafficStart;
        /// <summary>
        /// Application group names.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Enable/disable source NAT.
        /// </summary>
        public readonly string Nat;
        /// <summary>
        /// Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN.
        /// </summary>
        public readonly string Outbound;
        /// <summary>
        /// Per-IP traffic shaper.
        /// </summary>
        public readonly string PerIpShaper;
        /// <summary>
        /// Policy ID.
        /// </summary>
        public readonly int Policyid;
        /// <summary>
        /// IPv4 pool names. The structure of `poolname4` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyPoolname4Result> Poolname4s;
        /// <summary>
        /// IPv6 pool names. The structure of `poolname6` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyPoolname6Result> Poolname6s;
        /// <summary>
        /// Name of profile group.
        /// </summary>
        public readonly string ProfileGroup;
        /// <summary>
        /// Name of an existing Protocol options profile.
        /// </summary>
        public readonly string ProfileProtocolOptions;
        /// <summary>
        /// Determine whether the firewall policy allows security profile groups or single profiles only.
        /// </summary>
        public readonly string ProfileType;
        /// <summary>
        /// Schedule name.
        /// </summary>
        public readonly string Schedule;
        /// <summary>
        /// When enabled service specifies what the service must NOT be.
        /// </summary>
        public readonly string ServiceNegate;
        /// <summary>
        /// Service and service group names. The structure of `service` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyServiceResult> Services;
        /// <summary>
        /// TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
        /// </summary>
        public readonly int SessionTtl;
        /// <summary>
        /// Name of an existing Spam filter profile.
        /// </summary>
        public readonly string SpamfilterProfile;
        /// <summary>
        /// Source IPv4 address name and address group names. The structure of `srcaddr4` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicySrcaddr4Result> Srcaddr4s;
        /// <summary>
        /// Source IPv6 address name and address group names. The structure of `srcaddr6` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicySrcaddr6Result> Srcaddr6s;
        /// <summary>
        /// When enabled srcaddr specifies what the source address must NOT be.
        /// </summary>
        public readonly string SrcaddrNegate;
        /// <summary>
        /// Incoming (ingress) interface. The structure of `srcintf` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicySrcintfResult> Srcintfs;
        /// <summary>
        /// Name of an existing SSH filter profile.
        /// </summary>
        public readonly string SshFilterProfile;
        /// <summary>
        /// Redirect SSH traffic to matching transparent proxy policy.
        /// </summary>
        public readonly string SshPolicyRedirect;
        /// <summary>
        /// Name of an existing SSL SSH profile.
        /// </summary>
        public readonly string SslSshProfile;
        /// <summary>
        /// Enable or disable this policy.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Receiver TCP maximum segment size (MSS).
        /// </summary>
        public readonly int TcpMssReceiver;
        /// <summary>
        /// Sender TCP maximum segment size (MSS).
        /// </summary>
        public readonly int TcpMssSender;
        /// <summary>
        /// Traffic shaper.
        /// </summary>
        public readonly string TrafficShaper;
        /// <summary>
        /// Reverse traffic shaper.
        /// </summary>
        public readonly string TrafficShaperReverse;
        /// <summary>
        /// URL category ID list. The structure of `url_category` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyUrlCategoryResult> UrlCategories;
        /// <summary>
        /// Names of individual users that can authenticate with this policy. The structure of `users` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallconsolidatedPolicyUserResult> Users;
        /// <summary>
        /// Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.
        /// </summary>
        public readonly string UtmStatus;
        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        public readonly string Uuid;
        public readonly string? Vdomparam;
        /// <summary>
        /// Name of an existing VoIP profile.
        /// </summary>
        public readonly string VoipProfile;
        /// <summary>
        /// Policy-based IPsec VPN: name of the IPsec VPN Phase 1.
        /// </summary>
        public readonly string Vpntunnel;
        /// <summary>
        /// Name of an existing Web application firewall profile.
        /// </summary>
        public readonly string WafProfile;
        /// <summary>
        /// Enable/disable WAN optimization.
        /// </summary>
        public readonly string Wanopt;
        /// <summary>
        /// WAN optimization auto-detection mode.
        /// </summary>
        public readonly string WanoptDetection;
        /// <summary>
        /// WAN optimization passive mode options. This option decides what IP address will be used to connect to server.
        /// </summary>
        public readonly string WanoptPassiveOpt;
        /// <summary>
        /// WAN optimization peer.
        /// </summary>
        public readonly string WanoptPeer;
        /// <summary>
        /// WAN optimization profile.
        /// </summary>
        public readonly string WanoptProfile;
        /// <summary>
        /// Enable/disable web cache.
        /// </summary>
        public readonly string Webcache;
        /// <summary>
        /// Enable/disable web cache for HTTPS.
        /// </summary>
        public readonly string WebcacheHttps;
        /// <summary>
        /// Name of an existing Web filter profile.
        /// </summary>
        public readonly string WebfilterProfile;
        /// <summary>
        /// Webproxy forward server name.
        /// </summary>
        public readonly string WebproxyForwardServer;
        /// <summary>
        /// Webproxy profile name.
        /// </summary>
        public readonly string WebproxyProfile;

        [OutputConstructor]
        private GetFirewallconsolidatedPolicyResult(
            string action,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyAppCategoryResult> appCategories,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyAppGroupResult> appGroups,

            string applicationList,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyApplicationResult> applications,

            string autoAsicOffload,

            string avProfile,

            string captivePortalExempt,

            string cifsProfile,

            string comments,

            string diffservForward,

            string diffservReverse,

            string diffservcodeForward,

            string diffservcodeRev,

            string dlpSensor,

            string dnsfilterProfile,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyDstaddr4Result> dstaddr4s,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyDstaddr6Result> dstaddr6s,

            string dstaddrNegate,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyDstintfResult> dstintfs,

            string emailfilterProfile,

            string fixedport,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyFssoGroupResult> fssoGroups,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyGroupResult> groups,

            string httpPolicyRedirect,

            string icapProfile,

            string id,

            string inbound,

            string inspectionMode,

            string internetService,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceCustomGroupResult> internetServiceCustomGroups,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceCustomResult> internetServiceCustoms,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceGroupResult> internetServiceGroups,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceIdResult> internetServiceIds,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceNameResult> internetServiceNames,

            string internetServiceNegate,

            string internetServiceSrc,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceSrcCustomGroupResult> internetServiceSrcCustomGroups,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceSrcCustomResult> internetServiceSrcCustoms,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceSrcGroupResult> internetServiceSrcGroups,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceSrcIdResult> internetServiceSrcIds,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyInternetServiceSrcNameResult> internetServiceSrcNames,

            string internetServiceSrcNegate,

            string ippool,

            string ipsSensor,

            string logtraffic,

            string logtrafficStart,

            string name,

            string nat,

            string outbound,

            string perIpShaper,

            int policyid,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyPoolname4Result> poolname4s,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyPoolname6Result> poolname6s,

            string profileGroup,

            string profileProtocolOptions,

            string profileType,

            string schedule,

            string serviceNegate,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyServiceResult> services,

            int sessionTtl,

            string spamfilterProfile,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicySrcaddr4Result> srcaddr4s,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicySrcaddr6Result> srcaddr6s,

            string srcaddrNegate,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicySrcintfResult> srcintfs,

            string sshFilterProfile,

            string sshPolicyRedirect,

            string sslSshProfile,

            string status,

            int tcpMssReceiver,

            int tcpMssSender,

            string trafficShaper,

            string trafficShaperReverse,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyUrlCategoryResult> urlCategories,

            ImmutableArray<Outputs.GetFirewallconsolidatedPolicyUserResult> users,

            string utmStatus,

            string uuid,

            string? vdomparam,

            string voipProfile,

            string vpntunnel,

            string wafProfile,

            string wanopt,

            string wanoptDetection,

            string wanoptPassiveOpt,

            string wanoptPeer,

            string wanoptProfile,

            string webcache,

            string webcacheHttps,

            string webfilterProfile,

            string webproxyForwardServer,

            string webproxyProfile)
        {
            Action = action;
            AppCategories = appCategories;
            AppGroups = appGroups;
            ApplicationList = applicationList;
            Applications = applications;
            AutoAsicOffload = autoAsicOffload;
            AvProfile = avProfile;
            CaptivePortalExempt = captivePortalExempt;
            CifsProfile = cifsProfile;
            Comments = comments;
            DiffservForward = diffservForward;
            DiffservReverse = diffservReverse;
            DiffservcodeForward = diffservcodeForward;
            DiffservcodeRev = diffservcodeRev;
            DlpSensor = dlpSensor;
            DnsfilterProfile = dnsfilterProfile;
            Dstaddr4s = dstaddr4s;
            Dstaddr6s = dstaddr6s;
            DstaddrNegate = dstaddrNegate;
            Dstintfs = dstintfs;
            EmailfilterProfile = emailfilterProfile;
            Fixedport = fixedport;
            FssoGroups = fssoGroups;
            Groups = groups;
            HttpPolicyRedirect = httpPolicyRedirect;
            IcapProfile = icapProfile;
            Id = id;
            Inbound = inbound;
            InspectionMode = inspectionMode;
            InternetService = internetService;
            InternetServiceCustomGroups = internetServiceCustomGroups;
            InternetServiceCustoms = internetServiceCustoms;
            InternetServiceGroups = internetServiceGroups;
            InternetServiceIds = internetServiceIds;
            InternetServiceNames = internetServiceNames;
            InternetServiceNegate = internetServiceNegate;
            InternetServiceSrc = internetServiceSrc;
            InternetServiceSrcCustomGroups = internetServiceSrcCustomGroups;
            InternetServiceSrcCustoms = internetServiceSrcCustoms;
            InternetServiceSrcGroups = internetServiceSrcGroups;
            InternetServiceSrcIds = internetServiceSrcIds;
            InternetServiceSrcNames = internetServiceSrcNames;
            InternetServiceSrcNegate = internetServiceSrcNegate;
            Ippool = ippool;
            IpsSensor = ipsSensor;
            Logtraffic = logtraffic;
            LogtrafficStart = logtrafficStart;
            Name = name;
            Nat = nat;
            Outbound = outbound;
            PerIpShaper = perIpShaper;
            Policyid = policyid;
            Poolname4s = poolname4s;
            Poolname6s = poolname6s;
            ProfileGroup = profileGroup;
            ProfileProtocolOptions = profileProtocolOptions;
            ProfileType = profileType;
            Schedule = schedule;
            ServiceNegate = serviceNegate;
            Services = services;
            SessionTtl = sessionTtl;
            SpamfilterProfile = spamfilterProfile;
            Srcaddr4s = srcaddr4s;
            Srcaddr6s = srcaddr6s;
            SrcaddrNegate = srcaddrNegate;
            Srcintfs = srcintfs;
            SshFilterProfile = sshFilterProfile;
            SshPolicyRedirect = sshPolicyRedirect;
            SslSshProfile = sslSshProfile;
            Status = status;
            TcpMssReceiver = tcpMssReceiver;
            TcpMssSender = tcpMssSender;
            TrafficShaper = trafficShaper;
            TrafficShaperReverse = trafficShaperReverse;
            UrlCategories = urlCategories;
            Users = users;
            UtmStatus = utmStatus;
            Uuid = uuid;
            Vdomparam = vdomparam;
            VoipProfile = voipProfile;
            Vpntunnel = vpntunnel;
            WafProfile = wafProfile;
            Wanopt = wanopt;
            WanoptDetection = wanoptDetection;
            WanoptPassiveOpt = wanoptPassiveOpt;
            WanoptPeer = wanoptPeer;
            WanoptProfile = wanoptProfile;
            Webcache = webcache;
            WebcacheHttps = webcacheHttps;
            WebfilterProfile = webfilterProfile;
            WebproxyForwardServer = webproxyForwardServer;
            WebproxyProfile = webproxyProfile;
        }
    }
}
