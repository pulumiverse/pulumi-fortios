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
    public static class GetFirewallProxypolicy
    {
        /// <summary>
        /// Use this data source to get information on an fortios firewall proxypolicy
        /// </summary>
        public static Task<GetFirewallProxypolicyResult> InvokeAsync(GetFirewallProxypolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallProxypolicyResult>("fortios:index/getFirewallProxypolicy:getFirewallProxypolicy", args ?? new GetFirewallProxypolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios firewall proxypolicy
        /// </summary>
        public static Output<GetFirewallProxypolicyResult> Invoke(GetFirewallProxypolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallProxypolicyResult>("fortios:index/getFirewallProxypolicy:getFirewallProxypolicy", args ?? new GetFirewallProxypolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallProxypolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the policyid of the desired firewall proxypolicy.
        /// </summary>
        [Input("policyid", required: true)]
        public int Policyid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetFirewallProxypolicyArgs()
        {
        }
        public static new GetFirewallProxypolicyArgs Empty => new GetFirewallProxypolicyArgs();
    }

    public sealed class GetFirewallProxypolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the policyid of the desired firewall proxypolicy.
        /// </summary>
        [Input("policyid", required: true)]
        public Input<int> Policyid { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetFirewallProxypolicyInvokeArgs()
        {
        }
        public static new GetFirewallProxypolicyInvokeArgs Empty => new GetFirewallProxypolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallProxypolicyResult
    {
        /// <summary>
        /// IPv4 access proxy. The structure of `access_proxy` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallProxypolicyAccessProxyResult> AccessProxies;
        /// <summary>
        /// IPv6 access proxy. The structure of `access_proxy6` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallProxypolicyAccessProxy6Result> AccessProxy6s;
        /// <summary>
        /// Accept or deny traffic matching the policy parameters.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Name of an existing Application list.
        /// </summary>
        public readonly string ApplicationList;
        /// <summary>
        /// Name of an existing Antivirus profile.
        /// </summary>
        public readonly string AvProfile;
        /// <summary>
        /// Enable/disable block notification.
        /// </summary>
        public readonly string BlockNotification;
        /// <summary>
        /// Name of an existing CIFS profile.
        /// </summary>
        public readonly string CifsProfile;
        /// <summary>
        /// Optional comments.
        /// </summary>
        public readonly string Comments;
        /// <summary>
        /// Decrypted traffic mirror.
        /// </summary>
        public readonly string DecryptedTrafficMirror;
        /// <summary>
        /// When enabled, the ownership enforcement will be done at policy level.
        /// </summary>
        public readonly string DeviceOwnership;
        /// <summary>
        /// Web proxy disclaimer setting: by domain, policy, or user.
        /// </summary>
        public readonly string Disclaimer;
        /// <summary>
        /// Name of an existing DLP profile.
        /// </summary>
        public readonly string DlpProfile;
        /// <summary>
        /// Name of an existing DLP sensor.
        /// </summary>
        public readonly string DlpSensor;
        /// <summary>
        /// IPv6 destination address objects. The structure of `dstaddr6` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallProxypolicyDstaddr6Result> Dstaddr6s;
        /// <summary>
        /// When enabled, destination addresses match against any address EXCEPT the specified destination addresses.
        /// </summary>
        public readonly string DstaddrNegate;
        /// <summary>
        /// Destination address objects. The structure of `dstaddr` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallProxypolicyDstaddrResult> Dstaddrs;
        /// <summary>
        /// Destination interface names. The structure of `dstintf` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallProxypolicyDstintfResult> Dstintfs;
        /// <summary>
        /// Name of an existing email filter profile.
        /// </summary>
        public readonly string EmailfilterProfile;
        /// <summary>
        /// Name of an existing file-filter profile.
        /// </summary>
        public readonly string FileFilterProfile;
        /// <summary>
        /// Global web-based manager visible label.
        /// </summary>
        public readonly string GlobalLabel;
        /// <summary>
        /// Names of group objects. The structure of `groups` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallProxypolicyGroupResult> Groups;
        /// <summary>
        /// Enable/disable HTTP tunnel authentication.
        /// </summary>
        public readonly string HttpTunnelAuth;
        /// <summary>
        /// Name of an existing ICAP profile.
        /// </summary>
        public readonly string IcapProfile;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.
        /// </summary>
        public readonly string InternetService;
        /// <summary>
        /// Custom Internet Service group name. The structure of `internet_service_custom_group` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallProxypolicyInternetServiceCustomGroupResult> InternetServiceCustomGroups;
        /// <summary>
        /// Custom Internet Service name. The structure of `internet_service_custom` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallProxypolicyInternetServiceCustomResult> InternetServiceCustoms;
        /// <summary>
        /// Internet Service group name. The structure of `internet_service_group` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallProxypolicyInternetServiceGroupResult> InternetServiceGroups;
        /// <summary>
        /// Internet Service ID. The structure of `internet_service_id` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallProxypolicyInternetServiceIdResult> InternetServiceIds;
        /// <summary>
        /// Internet Service name. The structure of `internet_service_name` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallProxypolicyInternetServiceNameResult> InternetServiceNames;
        /// <summary>
        /// When enabled, Internet Services match against any internet service EXCEPT the selected Internet Service.
        /// </summary>
        public readonly string InternetServiceNegate;
        /// <summary>
        /// Name of an existing IPS sensor.
        /// </summary>
        public readonly string IpsSensor;
        /// <summary>
        /// VDOM-specific GUI visible label.
        /// </summary>
        public readonly string Label;
        /// <summary>
        /// Enable/disable logging traffic through the policy.
        /// </summary>
        public readonly string Logtraffic;
        /// <summary>
        /// Enable/disable policy log traffic start.
        /// </summary>
        public readonly string LogtrafficStart;
        /// <summary>
        /// Group name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Policy ID.
        /// </summary>
        public readonly int Policyid;
        /// <summary>
        /// Name of IP pool object. The structure of `poolname` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallProxypolicyPoolnameResult> Poolnames;
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
        /// Type of explicit proxy.
        /// </summary>
        public readonly string Proxy;
        /// <summary>
        /// Redirect URL for further explicit web proxy processing.
        /// </summary>
        public readonly string RedirectUrl;
        /// <summary>
        /// Authentication replacement message override group.
        /// </summary>
        public readonly string ReplacemsgOverrideGroup;
        /// <summary>
        /// Enable/disable scanning of connections to Botnet servers.
        /// </summary>
        public readonly string ScanBotnetConnections;
        /// <summary>
        /// Name of schedule object.
        /// </summary>
        public readonly string Schedule;
        /// <summary>
        /// Name of an existing SCTP filter profile.
        /// </summary>
        public readonly string SctpFilterProfile;
        /// <summary>
        /// When enabled, services match against any service EXCEPT the specified destination services.
        /// </summary>
        public readonly string ServiceNegate;
        /// <summary>
        /// Name of service objects. The structure of `service` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallProxypolicyServiceResult> Services;
        /// <summary>
        /// TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
        /// </summary>
        public readonly int SessionTtl;
        /// <summary>
        /// Name of an existing Spam filter profile.
        /// </summary>
        public readonly string SpamfilterProfile;
        /// <summary>
        /// IPv6 source address objects. The structure of `srcaddr6` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallProxypolicySrcaddr6Result> Srcaddr6s;
        /// <summary>
        /// When enabled, source addresses match against any address EXCEPT the specified source addresses.
        /// </summary>
        public readonly string SrcaddrNegate;
        /// <summary>
        /// Source address objects. The structure of `srcaddr` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallProxypolicySrcaddrResult> Srcaddrs;
        /// <summary>
        /// Source interface names. The structure of `srcintf` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallProxypolicySrcintfResult> Srcintfs;
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
        /// Enable/disable the active status of the policy.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Enable to use the IP address of the client to connect to the server.
        /// </summary>
        public readonly string Transparent;
        /// <summary>
        /// Names of user objects. The structure of `users` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallProxypolicyUserResult> Users;
        /// <summary>
        /// Enable the use of UTM profiles/sensors/lists.
        /// </summary>
        public readonly string UtmStatus;
        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        public readonly string Uuid;
        public readonly string? Vdomparam;
        /// <summary>
        /// Name of an existing VideoFilter profile.
        /// </summary>
        public readonly string VideofilterProfile;
        /// <summary>
        /// Name of an existing VoIP profile.
        /// </summary>
        public readonly string VoipProfile;
        /// <summary>
        /// Name of an existing Web application firewall profile.
        /// </summary>
        public readonly string WafProfile;
        /// <summary>
        /// Enable/disable web caching.
        /// </summary>
        public readonly string Webcache;
        /// <summary>
        /// Enable/disable web caching for HTTPS (Requires deep-inspection enabled in ssl-ssh-profile).
        /// </summary>
        public readonly string WebcacheHttps;
        /// <summary>
        /// Name of an existing Web filter profile.
        /// </summary>
        public readonly string WebfilterProfile;
        /// <summary>
        /// Web proxy forward server name.
        /// </summary>
        public readonly string WebproxyForwardServer;
        /// <summary>
        /// Name of web proxy profile.
        /// </summary>
        public readonly string WebproxyProfile;
        /// <summary>
        /// ZTNA EMS Tag names. The structure of `ztna_ems_tag` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallProxypolicyZtnaEmsTagResult> ZtnaEmsTags;
        /// <summary>
        /// ZTNA tag matching logic.
        /// </summary>
        public readonly string ZtnaTagsMatchLogic;

        [OutputConstructor]
        private GetFirewallProxypolicyResult(
            ImmutableArray<Outputs.GetFirewallProxypolicyAccessProxyResult> accessProxies,

            ImmutableArray<Outputs.GetFirewallProxypolicyAccessProxy6Result> accessProxy6s,

            string action,

            string applicationList,

            string avProfile,

            string blockNotification,

            string cifsProfile,

            string comments,

            string decryptedTrafficMirror,

            string deviceOwnership,

            string disclaimer,

            string dlpProfile,

            string dlpSensor,

            ImmutableArray<Outputs.GetFirewallProxypolicyDstaddr6Result> dstaddr6s,

            string dstaddrNegate,

            ImmutableArray<Outputs.GetFirewallProxypolicyDstaddrResult> dstaddrs,

            ImmutableArray<Outputs.GetFirewallProxypolicyDstintfResult> dstintfs,

            string emailfilterProfile,

            string fileFilterProfile,

            string globalLabel,

            ImmutableArray<Outputs.GetFirewallProxypolicyGroupResult> groups,

            string httpTunnelAuth,

            string icapProfile,

            string id,

            string internetService,

            ImmutableArray<Outputs.GetFirewallProxypolicyInternetServiceCustomGroupResult> internetServiceCustomGroups,

            ImmutableArray<Outputs.GetFirewallProxypolicyInternetServiceCustomResult> internetServiceCustoms,

            ImmutableArray<Outputs.GetFirewallProxypolicyInternetServiceGroupResult> internetServiceGroups,

            ImmutableArray<Outputs.GetFirewallProxypolicyInternetServiceIdResult> internetServiceIds,

            ImmutableArray<Outputs.GetFirewallProxypolicyInternetServiceNameResult> internetServiceNames,

            string internetServiceNegate,

            string ipsSensor,

            string label,

            string logtraffic,

            string logtrafficStart,

            string name,

            int policyid,

            ImmutableArray<Outputs.GetFirewallProxypolicyPoolnameResult> poolnames,

            string profileGroup,

            string profileProtocolOptions,

            string profileType,

            string proxy,

            string redirectUrl,

            string replacemsgOverrideGroup,

            string scanBotnetConnections,

            string schedule,

            string sctpFilterProfile,

            string serviceNegate,

            ImmutableArray<Outputs.GetFirewallProxypolicyServiceResult> services,

            int sessionTtl,

            string spamfilterProfile,

            ImmutableArray<Outputs.GetFirewallProxypolicySrcaddr6Result> srcaddr6s,

            string srcaddrNegate,

            ImmutableArray<Outputs.GetFirewallProxypolicySrcaddrResult> srcaddrs,

            ImmutableArray<Outputs.GetFirewallProxypolicySrcintfResult> srcintfs,

            string sshFilterProfile,

            string sshPolicyRedirect,

            string sslSshProfile,

            string status,

            string transparent,

            ImmutableArray<Outputs.GetFirewallProxypolicyUserResult> users,

            string utmStatus,

            string uuid,

            string? vdomparam,

            string videofilterProfile,

            string voipProfile,

            string wafProfile,

            string webcache,

            string webcacheHttps,

            string webfilterProfile,

            string webproxyForwardServer,

            string webproxyProfile,

            ImmutableArray<Outputs.GetFirewallProxypolicyZtnaEmsTagResult> ztnaEmsTags,

            string ztnaTagsMatchLogic)
        {
            AccessProxies = accessProxies;
            AccessProxy6s = accessProxy6s;
            Action = action;
            ApplicationList = applicationList;
            AvProfile = avProfile;
            BlockNotification = blockNotification;
            CifsProfile = cifsProfile;
            Comments = comments;
            DecryptedTrafficMirror = decryptedTrafficMirror;
            DeviceOwnership = deviceOwnership;
            Disclaimer = disclaimer;
            DlpProfile = dlpProfile;
            DlpSensor = dlpSensor;
            Dstaddr6s = dstaddr6s;
            DstaddrNegate = dstaddrNegate;
            Dstaddrs = dstaddrs;
            Dstintfs = dstintfs;
            EmailfilterProfile = emailfilterProfile;
            FileFilterProfile = fileFilterProfile;
            GlobalLabel = globalLabel;
            Groups = groups;
            HttpTunnelAuth = httpTunnelAuth;
            IcapProfile = icapProfile;
            Id = id;
            InternetService = internetService;
            InternetServiceCustomGroups = internetServiceCustomGroups;
            InternetServiceCustoms = internetServiceCustoms;
            InternetServiceGroups = internetServiceGroups;
            InternetServiceIds = internetServiceIds;
            InternetServiceNames = internetServiceNames;
            InternetServiceNegate = internetServiceNegate;
            IpsSensor = ipsSensor;
            Label = label;
            Logtraffic = logtraffic;
            LogtrafficStart = logtrafficStart;
            Name = name;
            Policyid = policyid;
            Poolnames = poolnames;
            ProfileGroup = profileGroup;
            ProfileProtocolOptions = profileProtocolOptions;
            ProfileType = profileType;
            Proxy = proxy;
            RedirectUrl = redirectUrl;
            ReplacemsgOverrideGroup = replacemsgOverrideGroup;
            ScanBotnetConnections = scanBotnetConnections;
            Schedule = schedule;
            SctpFilterProfile = sctpFilterProfile;
            ServiceNegate = serviceNegate;
            Services = services;
            SessionTtl = sessionTtl;
            SpamfilterProfile = spamfilterProfile;
            Srcaddr6s = srcaddr6s;
            SrcaddrNegate = srcaddrNegate;
            Srcaddrs = srcaddrs;
            Srcintfs = srcintfs;
            SshFilterProfile = sshFilterProfile;
            SshPolicyRedirect = sshPolicyRedirect;
            SslSshProfile = sslSshProfile;
            Status = status;
            Transparent = transparent;
            Users = users;
            UtmStatus = utmStatus;
            Uuid = uuid;
            Vdomparam = vdomparam;
            VideofilterProfile = videofilterProfile;
            VoipProfile = voipProfile;
            WafProfile = wafProfile;
            Webcache = webcache;
            WebcacheHttps = webcacheHttps;
            WebfilterProfile = webfilterProfile;
            WebproxyForwardServer = webproxyForwardServer;
            WebproxyProfile = webproxyProfile;
            ZtnaEmsTags = ztnaEmsTags;
            ZtnaTagsMatchLogic = ztnaTagsMatchLogic;
        }
    }
}
