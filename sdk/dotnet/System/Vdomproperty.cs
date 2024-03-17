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
    /// <summary>
    /// Configure VDOM property.
    /// 
    /// ## Import
    /// 
    /// System VdomProperty can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/vdomproperty:Vdomproperty labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/vdomproperty:Vdomproperty labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/vdomproperty:Vdomproperty")]
    public partial class Vdomproperty : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Maximum guaranteed number of firewall custom services.
        /// </summary>
        [Output("customService")]
        public Output<string> CustomService { get; private set; } = null!;

        /// <summary>
        /// Description.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Maximum guaranteed number of dial-up tunnels.
        /// </summary>
        [Output("dialupTunnel")]
        public Output<string> DialupTunnel { get; private set; } = null!;

        /// <summary>
        /// Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
        /// </summary>
        [Output("firewallAddress")]
        public Output<string> FirewallAddress { get; private set; } = null!;

        /// <summary>
        /// Maximum guaranteed number of firewall address groups (IPv4, IPv6).
        /// </summary>
        [Output("firewallAddrgrp")]
        public Output<string> FirewallAddrgrp { get; private set; } = null!;

        /// <summary>
        /// Maximum guaranteed number of firewall policies (IPv4, IPv6, policy46, policy64, DoS-policy4, DoS-policy6, multicast).
        /// </summary>
        [Output("firewallPolicy")]
        public Output<string> FirewallPolicy { get; private set; } = null!;

        /// <summary>
        /// Maximum guaranteed number of VPN IPsec phase 1 tunnels.
        /// </summary>
        [Output("ipsecPhase1")]
        public Output<string> IpsecPhase1 { get; private set; } = null!;

        /// <summary>
        /// Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
        /// </summary>
        [Output("ipsecPhase1Interface")]
        public Output<string> IpsecPhase1Interface { get; private set; } = null!;

        /// <summary>
        /// Maximum guaranteed number of VPN IPsec phase 2 tunnels.
        /// </summary>
        [Output("ipsecPhase2")]
        public Output<string> IpsecPhase2 { get; private set; } = null!;

        /// <summary>
        /// Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
        /// </summary>
        [Output("ipsecPhase2Interface")]
        public Output<string> IpsecPhase2Interface { get; private set; } = null!;

        /// <summary>
        /// Log disk quota in MB (range depends on how much disk space is available).
        /// </summary>
        [Output("logDiskQuota")]
        public Output<string> LogDiskQuota { get; private set; } = null!;

        /// <summary>
        /// VDOM name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Maximum guaranteed number of firewall one-time schedules.
        /// </summary>
        [Output("onetimeSchedule")]
        public Output<string> OnetimeSchedule { get; private set; } = null!;

        /// <summary>
        /// Maximum guaranteed number of concurrent proxy users.
        /// </summary>
        [Output("proxy")]
        public Output<string> Proxy { get; private set; } = null!;

        /// <summary>
        /// Maximum guaranteed number of firewall recurring schedules.
        /// </summary>
        [Output("recurringSchedule")]
        public Output<string> RecurringSchedule { get; private set; } = null!;

        /// <summary>
        /// Maximum guaranteed number of firewall service groups.
        /// </summary>
        [Output("serviceGroup")]
        public Output<string> ServiceGroup { get; private set; } = null!;

        /// <summary>
        /// Maximum guaranteed number of sessions.
        /// </summary>
        [Output("session")]
        public Output<string> Session { get; private set; } = null!;

        /// <summary>
        /// Permanent SNMP Index of the virtual domain (0 - 4294967295).
        /// </summary>
        [Output("snmpIndex")]
        public Output<int> SnmpIndex { get; private set; } = null!;

        /// <summary>
        /// Maximum guaranteed number of SSL-VPNs.
        /// </summary>
        [Output("sslvpn")]
        public Output<string> Sslvpn { get; private set; } = null!;

        /// <summary>
        /// Maximum guaranteed number of local users.
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;

        /// <summary>
        /// Maximum guaranteed number of user groups.
        /// </summary>
        [Output("userGroup")]
        public Output<string> UserGroup { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Vdomproperty resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vdomproperty(string name, VdompropertyArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/vdomproperty:Vdomproperty", name, args ?? new VdompropertyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Vdomproperty(string name, Input<string> id, VdompropertyState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/vdomproperty:Vdomproperty", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Vdomproperty resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vdomproperty Get(string name, Input<string> id, VdompropertyState? state = null, CustomResourceOptions? options = null)
        {
            return new Vdomproperty(name, id, state, options);
        }
    }

    public sealed class VdompropertyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum guaranteed number of firewall custom services.
        /// </summary>
        [Input("customService")]
        public Input<string>? CustomService { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Maximum guaranteed number of dial-up tunnels.
        /// </summary>
        [Input("dialupTunnel")]
        public Input<string>? DialupTunnel { get; set; }

        /// <summary>
        /// Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
        /// </summary>
        [Input("firewallAddress")]
        public Input<string>? FirewallAddress { get; set; }

        /// <summary>
        /// Maximum guaranteed number of firewall address groups (IPv4, IPv6).
        /// </summary>
        [Input("firewallAddrgrp")]
        public Input<string>? FirewallAddrgrp { get; set; }

        /// <summary>
        /// Maximum guaranteed number of firewall policies (IPv4, IPv6, policy46, policy64, DoS-policy4, DoS-policy6, multicast).
        /// </summary>
        [Input("firewallPolicy")]
        public Input<string>? FirewallPolicy { get; set; }

        /// <summary>
        /// Maximum guaranteed number of VPN IPsec phase 1 tunnels.
        /// </summary>
        [Input("ipsecPhase1")]
        public Input<string>? IpsecPhase1 { get; set; }

        /// <summary>
        /// Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
        /// </summary>
        [Input("ipsecPhase1Interface")]
        public Input<string>? IpsecPhase1Interface { get; set; }

        /// <summary>
        /// Maximum guaranteed number of VPN IPsec phase 2 tunnels.
        /// </summary>
        [Input("ipsecPhase2")]
        public Input<string>? IpsecPhase2 { get; set; }

        /// <summary>
        /// Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
        /// </summary>
        [Input("ipsecPhase2Interface")]
        public Input<string>? IpsecPhase2Interface { get; set; }

        /// <summary>
        /// Log disk quota in MB (range depends on how much disk space is available).
        /// </summary>
        [Input("logDiskQuota")]
        public Input<string>? LogDiskQuota { get; set; }

        /// <summary>
        /// VDOM name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Maximum guaranteed number of firewall one-time schedules.
        /// </summary>
        [Input("onetimeSchedule")]
        public Input<string>? OnetimeSchedule { get; set; }

        /// <summary>
        /// Maximum guaranteed number of concurrent proxy users.
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        /// <summary>
        /// Maximum guaranteed number of firewall recurring schedules.
        /// </summary>
        [Input("recurringSchedule")]
        public Input<string>? RecurringSchedule { get; set; }

        /// <summary>
        /// Maximum guaranteed number of firewall service groups.
        /// </summary>
        [Input("serviceGroup")]
        public Input<string>? ServiceGroup { get; set; }

        /// <summary>
        /// Maximum guaranteed number of sessions.
        /// </summary>
        [Input("session")]
        public Input<string>? Session { get; set; }

        /// <summary>
        /// Permanent SNMP Index of the virtual domain (0 - 4294967295).
        /// </summary>
        [Input("snmpIndex")]
        public Input<int>? SnmpIndex { get; set; }

        /// <summary>
        /// Maximum guaranteed number of SSL-VPNs.
        /// </summary>
        [Input("sslvpn")]
        public Input<string>? Sslvpn { get; set; }

        /// <summary>
        /// Maximum guaranteed number of local users.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        /// <summary>
        /// Maximum guaranteed number of user groups.
        /// </summary>
        [Input("userGroup")]
        public Input<string>? UserGroup { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public VdompropertyArgs()
        {
        }
        public static new VdompropertyArgs Empty => new VdompropertyArgs();
    }

    public sealed class VdompropertyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum guaranteed number of firewall custom services.
        /// </summary>
        [Input("customService")]
        public Input<string>? CustomService { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Maximum guaranteed number of dial-up tunnels.
        /// </summary>
        [Input("dialupTunnel")]
        public Input<string>? DialupTunnel { get; set; }

        /// <summary>
        /// Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
        /// </summary>
        [Input("firewallAddress")]
        public Input<string>? FirewallAddress { get; set; }

        /// <summary>
        /// Maximum guaranteed number of firewall address groups (IPv4, IPv6).
        /// </summary>
        [Input("firewallAddrgrp")]
        public Input<string>? FirewallAddrgrp { get; set; }

        /// <summary>
        /// Maximum guaranteed number of firewall policies (IPv4, IPv6, policy46, policy64, DoS-policy4, DoS-policy6, multicast).
        /// </summary>
        [Input("firewallPolicy")]
        public Input<string>? FirewallPolicy { get; set; }

        /// <summary>
        /// Maximum guaranteed number of VPN IPsec phase 1 tunnels.
        /// </summary>
        [Input("ipsecPhase1")]
        public Input<string>? IpsecPhase1 { get; set; }

        /// <summary>
        /// Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
        /// </summary>
        [Input("ipsecPhase1Interface")]
        public Input<string>? IpsecPhase1Interface { get; set; }

        /// <summary>
        /// Maximum guaranteed number of VPN IPsec phase 2 tunnels.
        /// </summary>
        [Input("ipsecPhase2")]
        public Input<string>? IpsecPhase2 { get; set; }

        /// <summary>
        /// Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
        /// </summary>
        [Input("ipsecPhase2Interface")]
        public Input<string>? IpsecPhase2Interface { get; set; }

        /// <summary>
        /// Log disk quota in MB (range depends on how much disk space is available).
        /// </summary>
        [Input("logDiskQuota")]
        public Input<string>? LogDiskQuota { get; set; }

        /// <summary>
        /// VDOM name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Maximum guaranteed number of firewall one-time schedules.
        /// </summary>
        [Input("onetimeSchedule")]
        public Input<string>? OnetimeSchedule { get; set; }

        /// <summary>
        /// Maximum guaranteed number of concurrent proxy users.
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        /// <summary>
        /// Maximum guaranteed number of firewall recurring schedules.
        /// </summary>
        [Input("recurringSchedule")]
        public Input<string>? RecurringSchedule { get; set; }

        /// <summary>
        /// Maximum guaranteed number of firewall service groups.
        /// </summary>
        [Input("serviceGroup")]
        public Input<string>? ServiceGroup { get; set; }

        /// <summary>
        /// Maximum guaranteed number of sessions.
        /// </summary>
        [Input("session")]
        public Input<string>? Session { get; set; }

        /// <summary>
        /// Permanent SNMP Index of the virtual domain (0 - 4294967295).
        /// </summary>
        [Input("snmpIndex")]
        public Input<int>? SnmpIndex { get; set; }

        /// <summary>
        /// Maximum guaranteed number of SSL-VPNs.
        /// </summary>
        [Input("sslvpn")]
        public Input<string>? Sslvpn { get; set; }

        /// <summary>
        /// Maximum guaranteed number of local users.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        /// <summary>
        /// Maximum guaranteed number of user groups.
        /// </summary>
        [Input("userGroup")]
        public Input<string>? UserGroup { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public VdompropertyState()
        {
        }
        public static new VdompropertyState Empty => new VdompropertyState();
    }
}
