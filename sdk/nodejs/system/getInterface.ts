// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios system interface
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const sample1 = fortios.system.getInterface({
 *     name: "port1",
 * });
 * export const output1 = sample1.then(sample1 => sample1.ip);
 * ```
 */
export function getInterface(args: GetInterfaceArgs, opts?: pulumi.InvokeOptions): Promise<GetInterfaceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:system/getInterface:getInterface", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getInterface.
 */
export interface GetInterfaceArgs {
    /**
     * Specify the name of the desired system interface.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getInterface.
 */
export interface GetInterfaceResult {
    /**
     * PPPoE server name.
     */
    readonly acName: string;
    /**
     * Aggregate interface.
     */
    readonly aggregate: string;
    /**
     * Type of aggregation.
     */
    readonly aggregateType: string;
    /**
     * Frame distribution algorithm.
     */
    readonly algorithm: string;
    /**
     * Alias will be displayed with the interface name to make it easier to distinguish.
     */
    readonly alias: string;
    /**
     * Management access settings for the secondary IP address.
     */
    readonly allowaccess: string;
    /**
     * Enable/disable automatic registration of unknown FortiAP devices.
     */
    readonly apDiscover: string;
    /**
     * Enable/disable ARP forwarding.
     */
    readonly arpforward: string;
    /**
     * HTTPS server certificate.
     */
    readonly authCert: string;
    /**
     * Address of captive portal.
     */
    readonly authPortalAddr: string;
    /**
     * PPP authentication type to use.
     */
    readonly authType: string;
    /**
     * Enable/disable automatic authorization of dedicated Fortinet extension device on this interface.
     */
    readonly autoAuthExtensionDevice: string;
    /**
     * Bandwidth measure time
     */
    readonly bandwidthMeasureTime: number;
    /**
     * Bidirectional Forwarding Detection (BFD) settings.
     */
    readonly bfd: string;
    /**
     * BFD desired minimal transmit interval.
     */
    readonly bfdDesiredMinTx: number;
    /**
     * BFD detection multiplier.
     */
    readonly bfdDetectMult: number;
    /**
     * BFD required minimal receive interval.
     */
    readonly bfdRequiredMinRx: number;
    /**
     * Enable/disable broadcasting FortiClient discovery messages.
     */
    readonly broadcastForticlientDiscovery: string;
    /**
     * Enable/disable broadcast forwarding.
     */
    readonly broadcastForward: string;
    /**
     * Enable/disable captive portal.
     */
    readonly captivePortal: number;
    /**
     * CLI connection status.
     */
    readonly cliConnStatus: number;
    /**
     * DHCP client options. The structure of `clientOptions` block is documented below.
     */
    readonly clientOptions: outputs.system.GetInterfaceClientOption[];
    /**
     * Color of icon on the GUI.
     */
    readonly color: number;
    /**
     * Configure interface for single purpose.
     */
    readonly dedicatedTo: string;
    /**
     * default purdue level of device detected on this interface.
     */
    readonly defaultPurdueLevel: string;
    /**
     * Enable to get the gateway IP from the DHCP or PPPoE server.
     */
    readonly defaultgw: string;
    /**
     * Description.
     */
    readonly description: string;
    /**
     * MTU of detected peer (0 - 4294967295).
     */
    readonly detectedPeerMtu: number;
    /**
     * Protocols used to detect the server.
     */
    readonly detectprotocol: string;
    /**
     * Gateway's ping server for this IP.
     */
    readonly detectserver: string;
    /**
     * Device access list.
     */
    readonly deviceAccessList: string;
    /**
     * Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
     */
    readonly deviceIdentification: string;
    /**
     * Enable/disable active gathering of device identity information about the devices on the network connected to this interface.
     */
    readonly deviceIdentificationActiveScan: string;
    /**
     * Enable/disable inclusion of devices detected on this interface in network vulnerability scans.
     */
    readonly deviceNetscan: string;
    /**
     * Enable/disable passive gathering of user identity information about users on this interface.
     */
    readonly deviceUserIdentification: string;
    /**
     * Device Index.
     */
    readonly devindex: number;
    /**
     * Enable/disable setting of the broadcast flag in messages sent by the DHCP client (default = enable).
     */
    readonly dhcpBroadcastFlag: string;
    /**
     * Enable/disable addition of classless static routes retrieved from DHCP server.
     */
    readonly dhcpClasslessRouteAddition: string;
    /**
     * DHCP client identifier.
     */
    readonly dhcpClientIdentifier: string;
    /**
     * Enable/disable DHCP relay agent option.
     */
    readonly dhcpRelayAgentOption: string;
    /**
     * DHCP relay circuit ID.
     */
    readonly dhcpRelayCircuitId: string;
    /**
     * Specify outgoing interface to reach server.
     */
    readonly dhcpRelayInterface: string;
    /**
     * Specify how to select outgoing interface to reach server.
     */
    readonly dhcpRelayInterfaceSelectMethod: string;
    /**
     * DHCP relay IP address.
     */
    readonly dhcpRelayIp: string;
    /**
     * DHCP relay link selection.
     */
    readonly dhcpRelayLinkSelection: string;
    /**
     * Enable/disable sending DHCP request to all servers.
     */
    readonly dhcpRelayRequestAllServer: string;
    /**
     * Enable/disable allowing this interface to act as a DHCP relay.
     */
    readonly dhcpRelayService: string;
    /**
     * IP address used by the DHCP relay as its source IP.
     */
    readonly dhcpRelaySourceIp: string;
    /**
     * DHCP relay type (regular or IPsec).
     */
    readonly dhcpRelayType: string;
    /**
     * DHCP renew time in seconds (300-604800), 0 means use the renew time provided by the server.
     */
    readonly dhcpRenewTime: number;
    /**
     * Enable/disable DHCP smart relay.
     */
    readonly dhcpSmartRelay: string;
    /**
     * Configure DHCP server access list. The structure of `dhcpSnoopingServerList` block is documented below.
     */
    readonly dhcpSnoopingServerLists: outputs.system.GetInterfaceDhcpSnoopingServerList[];
    /**
     * Time in seconds to wait before retrying to start a PPPoE discovery, 0 means no timeout.
     */
    readonly discRetryTimeout: number;
    /**
     * Time in milliseconds to wait before sending a notification that this interface is down or disconnected.
     */
    readonly disconnectThreshold: number;
    /**
     * Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
     */
    readonly distance: number;
    /**
     * Enable/disable use DNS acquired by DHCP or PPPoE.
     */
    readonly dnsServerOverride: string;
    /**
     * DNS transport protocols.
     */
    readonly dnsServerProtocol: string;
    /**
     * Enable/disable drop fragment packets.
     */
    readonly dropFragment: string;
    /**
     * Enable/disable drop overlapped fragment packets.
     */
    readonly dropOverlappedFragment: string;
    /**
     * EAP CA certificate name.
     */
    readonly eapCaCert: string;
    /**
     * EAP identity.
     */
    readonly eapIdentity: string;
    /**
     * EAP method.
     */
    readonly eapMethod: string;
    /**
     * EAP password.
     */
    readonly eapPassword: string;
    /**
     * Enable/disable EAP-Supplicant.
     */
    readonly eapSupplicant: string;
    /**
     * EAP user certificate name.
     */
    readonly eapUserCert: string;
    /**
     * Outgoing traffic shaping profile.
     */
    readonly egressShapingProfile: string;
    /**
     * Enable/disable endpoint compliance enforcement.
     */
    readonly endpointCompliance: string;
    /**
     * Estimated maximum downstream bandwidth (kbps). Used to estimate link utilization.
     */
    readonly estimatedDownstreamBandwidth: number;
    /**
     * Estimated maximum upstream bandwidth (kbps). Used to estimate link utilization.
     */
    readonly estimatedUpstreamBandwidth: number;
    /**
     * Enable/disable the explicit FTP proxy on this interface.
     */
    readonly explicitFtpProxy: string;
    /**
     * Enable/disable the explicit web proxy on this interface.
     */
    readonly explicitWebProxy: string;
    /**
     * Enable/disable identifying the interface as an external interface (which usually means it's connected to the Internet).
     */
    readonly external: string;
    /**
     * Action on extender when interface fail .
     */
    readonly failActionOnExtender: string;
    /**
     * Names of the FortiGate interfaces from which the link failure alert is sent for this interface. The structure of `failAlertInterfaces` block is documented below.
     */
    readonly failAlertInterfaces: outputs.system.GetInterfaceFailAlertInterface[];
    /**
     * Select link-failed-signal or link-down method to alert about a failed link.
     */
    readonly failAlertMethod: string;
    /**
     * Enable/disable fail detection features for this interface.
     */
    readonly failDetect: string;
    /**
     * Options for detecting that this interface has failed.
     */
    readonly failDetectOption: string;
    /**
     * Enable/disable FortiHeartBeat (FortiTelemetry on GUI).
     */
    readonly fortiheartbeat: string;
    /**
     * Enable FortiLink to dedicate this interface to manage other Fortinet devices.
     */
    readonly fortilink: string;
    /**
     * fortilink split interface backup link.
     */
    readonly fortilinkBackupLink: number;
    /**
     * Protocol for FortiGate neighbor discovery.
     */
    readonly fortilinkNeighborDetect: string;
    /**
     * Enable/disable FortiLink split interface to connect member link to different FortiSwitch in stack for uplink redundancy.
     */
    readonly fortilinkSplitInterface: string;
    /**
     * Enable/disable FortiLink switch-stacking on this interface.
     */
    readonly fortilinkStacking: string;
    /**
     * Transparent mode forward domain.
     */
    readonly forwardDomain: number;
    /**
     * Configure forward error correction (FEC).
     */
    readonly forwardErrorCorrection: string;
    /**
     * Enable/disable detect gateway alive for first.
     */
    readonly gwdetect: string;
    /**
     * HA election priority for the PING server.
     */
    readonly haPriority: number;
    /**
     * Enable/disable ICMP accept redirect.
     */
    readonly icmpAcceptRedirect: string;
    /**
     * Enable/disable ICMP send redirect.
     */
    readonly icmpSendRedirect: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Enable/disable authentication for this interface.
     */
    readonly identAccept: string;
    /**
     * PPPoE auto disconnect after idle timeout seconds, 0 means no timeout.
     */
    readonly idleTimeout: number;
    /**
     * Configure IKE authentication SAML server.
     */
    readonly ikeSamlServer: string;
    /**
     * Bandwidth limit for incoming traffic (0 - 16776000 kbps), 0 means unlimited.
     */
    readonly inbandwidth: number;
    /**
     * Incoming traffic shaping profile.
     */
    readonly ingressShapingProfile: string;
    /**
     * Ingress Spillover threshold (0 - 16776000 kbps).
     */
    readonly ingressSpilloverThreshold: number;
    /**
     * Interface name.
     */
    readonly interface: string;
    /**
     * Implicitly created.
     */
    readonly internal: number;
    /**
     * Secondary IP address of the interface.
     */
    readonly ip: string;
    /**
     * Enable/disable automatic IP address assignment of this interface by FortiIPAM.
     */
    readonly ipManagedByFortiipam: string;
    /**
     * Enable/disable IP/MAC binding.
     */
    readonly ipmac: string;
    /**
     * Enable/disable the use of this interface as a one-armed sniffer.
     */
    readonly ipsSnifferMode: string;
    /**
     * Unnumbered IP used for PPPoE interfaces for which no unique local address is provided.
     */
    readonly ipunnumbered: string;
    /**
     * IPv6 of interface. The structure of `ipv6` block is documented below.
     */
    readonly ipv6s: outputs.system.GetInterfaceIpv6[];
    /**
     * Enable/disable l2 forwarding.
     */
    readonly l2forward: string;
    /**
     * LACP HA secondary member.
     */
    readonly lacpHaSecondary: string;
    /**
     * LACP HA slave.
     */
    readonly lacpHaSlave: string;
    /**
     * LACP mode.
     */
    readonly lacpMode: string;
    /**
     * How often the interface sends LACP messages.
     */
    readonly lacpSpeed: string;
    /**
     * Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
     */
    readonly lcpEchoInterval: number;
    /**
     * Maximum missed LCP echo messages before disconnect.
     */
    readonly lcpMaxEchoFails: number;
    /**
     * Number of milliseconds to wait before considering a link is up.
     */
    readonly linkUpDelay: number;
    /**
     * LLDP-MED network policy profile.
     */
    readonly lldpNetworkPolicy: string;
    /**
     * Enable/disable Link Layer Discovery Protocol (LLDP) reception.
     */
    readonly lldpReception: string;
    /**
     * Enable/disable Link Layer Discovery Protocol (LLDP) transmission.
     */
    readonly lldpTransmission: string;
    /**
     * Change the interface's MAC address.
     */
    readonly macaddr: string;
    /**
     * Available when FortiLink is enabled, used for managed devices through FortiLink interface. The structure of `managedDevice` block is documented below.
     */
    readonly managedDevices: outputs.system.GetInterfaceManagedDevice[];
    /**
     * Number of IP addresses to be allocated by FortiIPAM and used by this FortiGate unit's DHCP server settings.
     */
    readonly managedSubnetworkSize: string;
    /**
     * High Availability in-band management IP address of this interface.
     */
    readonly managementIp: string;
    /**
     * Measured downstream bandwidth (kbps).
     */
    readonly measuredDownstreamBandwidth: number;
    /**
     * Measured upstream bandwidth (kbps).
     */
    readonly measuredUpstreamBandwidth: number;
    /**
     * Select SFP media interface type
     */
    readonly mediatype: string;
    /**
     * Physical interfaces that belong to the aggregate or redundant interface. The structure of `member` block is documented below.
     */
    readonly members: outputs.system.GetInterfaceMember[];
    /**
     * Minimum number of aggregated ports that must be up.
     */
    readonly minLinks: number;
    /**
     * Action to take when less than the configured minimum number of links are active.
     */
    readonly minLinksDown: string;
    /**
     * Addressing mode (static, DHCP, PPPoE).
     */
    readonly mode: string;
    /**
     * Enable monitoring bandwidth on this interface.
     */
    readonly monitorBandwidth: string;
    /**
     * MTU value for this interface.
     */
    readonly mtu: number;
    /**
     * Enable to set a custom MTU for this interface.
     */
    readonly mtuOverride: string;
    /**
     * Tag name.
     */
    readonly name: string;
    /**
     * Enable/disable NDISC forwarding.
     */
    readonly ndiscforward: string;
    /**
     * Enable/disable NETBIOS forwarding.
     */
    readonly netbiosForward: string;
    /**
     * Enable/disable NetFlow on this interface and set the data that NetFlow collects (rx, tx, or both).
     */
    readonly netflowSampler: string;
    /**
     * Bandwidth limit for outgoing traffic (0 - 16776000 kbps).
     */
    readonly outbandwidth: number;
    /**
     * PPPoE Active Discovery Terminate (PADT) used to terminate sessions after an idle time.
     */
    readonly padtRetryTimeout: number;
    /**
     * PPPoE account's password.
     */
    readonly password: string;
    /**
     * PING server status.
     */
    readonly pingServStatus: number;
    /**
     * sFlow polling interval (1 - 255 sec).
     */
    readonly pollingInterval: number;
    /**
     * Enable/disable PPPoE unnumbered negotiation.
     */
    readonly pppoeUnnumberedNegotiate: string;
    /**
     * PPTP authentication type.
     */
    readonly pptpAuthType: string;
    /**
     * Enable/disable PPTP client.
     */
    readonly pptpClient: string;
    /**
     * PPTP password.
     */
    readonly pptpPassword: string;
    /**
     * PPTP server IP address.
     */
    readonly pptpServerIp: string;
    /**
     * Idle timer in minutes (0 for disabled).
     */
    readonly pptpTimeout: number;
    /**
     * PPTP user name.
     */
    readonly pptpUser: string;
    /**
     * Enable/disable preservation of session route when dirty.
     */
    readonly preserveSessionRoute: string;
    /**
     * Priority of the virtual router (1 - 255).
     */
    readonly priority: number;
    /**
     * Enable/disable fail back to higher priority port once recovered.
     */
    readonly priorityOverride: string;
    /**
     * Enable/disable proxy captive portal on this interface.
     */
    readonly proxyCaptivePortal: string;
    /**
     * IPv4 reachable time in milliseconds (30000 - 3600000, default = 30000).
     */
    readonly reachableTime: number;
    /**
     * Redundant interface.
     */
    readonly redundantInterface: string;
    /**
     * Remote IP address of tunnel.
     */
    readonly remoteIp: string;
    /**
     * Replacement message override group.
     */
    readonly replacemsgOverrideGroup: string;
    /**
     * RX ring size.
     */
    readonly ringRx: number;
    /**
     * TX ring size.
     */
    readonly ringTx: number;
    /**
     * Interface role.
     */
    readonly role: string;
    /**
     * Data that NetFlow collects (rx, tx, or both).
     */
    readonly sampleDirection: string;
    /**
     * sFlow sample rate (10 - 99999).
     */
    readonly sampleRate: number;
    /**
     * Enable monitoring or blocking connections to Botnet servers through this interface.
     */
    readonly scanBotnetConnections: string;
    /**
     * Enable/disable adding a secondary IP to this interface.
     */
    readonly secondaryIp: string;
    /**
     * Second IP address of interface. The structure of `secondaryip` block is documented below.
     */
    readonly secondaryips: outputs.system.GetInterfaceSecondaryip[];
    /**
     * Name of security-exempt-list.
     */
    readonly securityExemptList: string;
    /**
     * URL of external authentication logout server.
     */
    readonly securityExternalLogout: string;
    /**
     * URL of external authentication web server.
     */
    readonly securityExternalWeb: string;
    /**
     * User groups that can authenticate with the captive portal. The structure of `securityGroups` block is documented below.
     */
    readonly securityGroups: outputs.system.GetInterfaceSecurityGroup[];
    /**
     * Enable/disable MAC authentication bypass.
     */
    readonly securityMacAuthBypass: string;
    /**
     * Turn on captive portal authentication for this interface.
     */
    readonly securityMode: string;
    /**
     * URL redirection after disclaimer/authentication.
     */
    readonly securityRedirectUrl: string;
    /**
     * PPPoE service name.
     */
    readonly serviceName: string;
    /**
     * Enable/disable sFlow on this interface.
     */
    readonly sflowSampler: string;
    /**
     * Permanent SNMP Index of the interface.
     */
    readonly snmpIndex: number;
    /**
     * Interface speed. The default setting and the options available depend on the interface hardware.
     */
    readonly speed: string;
    /**
     * Egress Spillover threshold (0 - 16776000 kbps), 0 means unlimited.
     */
    readonly spilloverThreshold: number;
    /**
     * Enable/disable source IP check.
     */
    readonly srcCheck: string;
    /**
     * Enable/disable VRRP.
     */
    readonly status: string;
    /**
     * Enable/disable STP.
     */
    readonly stp: string;
    /**
     * Control STP behaviour on HA secondary.
     */
    readonly stpHaSecondary: string;
    /**
     * Enable/disable STP forwarding.
     */
    readonly stpforward: string;
    /**
     * Configure STP forwarding mode.
     */
    readonly stpforwardMode: string;
    /**
     * Enable to always send packets from this interface to a destination MAC address.
     */
    readonly subst: string;
    /**
     * Destination MAC address that all packets are sent to from this interface.
     */
    readonly substituteDstMac: string;
    /**
     * Initial create for switch-controller VLANs.
     */
    readonly swcFirstCreate: number;
    /**
     * Creation status for switch-controller VLANs.
     */
    readonly swcVlan: number;
    /**
     * Contained in switch.
     */
    readonly switch: string;
    /**
     * Block FortiSwitch port-to-port traffic.
     */
    readonly switchControllerAccessVlan: string;
    /**
     * Enable/disable FortiSwitch ARP inspection.
     */
    readonly switchControllerArpInspection: string;
    /**
     * Switch controller DHCP snooping.
     */
    readonly switchControllerDhcpSnooping: string;
    /**
     * Switch controller DHCP snooping option82.
     */
    readonly switchControllerDhcpSnoopingOption82: string;
    /**
     * Switch controller DHCP snooping verify MAC.
     */
    readonly switchControllerDhcpSnoopingVerifyMac: string;
    /**
     * Integrated FortiLink settings for managed FortiSwitch.
     */
    readonly switchControllerDynamic: string;
    /**
     * Interface's purpose when assigning traffic (read only).
     */
    readonly switchControllerFeature: string;
    /**
     * Switch controller IGMP snooping.
     */
    readonly switchControllerIgmpSnooping: string;
    /**
     * Switch controller IGMP snooping fast-leave.
     */
    readonly switchControllerIgmpSnoopingFastLeave: string;
    /**
     * Switch controller IGMP snooping proxy.
     */
    readonly switchControllerIgmpSnoopingProxy: string;
    /**
     * Enable/disable managed FortiSwitch IoT scanning.
     */
    readonly switchControllerIotScanning: string;
    /**
     * Limit the number of dynamic MAC addresses on this VLAN (1 - 128, 0 = no limit, default).
     */
    readonly switchControllerLearningLimit: number;
    /**
     * VLAN to use for FortiLink management purposes.
     */
    readonly switchControllerMgmtVlan: number;
    /**
     * Integrated NAC settings for managed FortiSwitch.
     */
    readonly switchControllerNac: string;
    /**
     * NetFlow collection and processing.
     */
    readonly switchControllerNetflowCollect: string;
    /**
     * Enable/disable managed FortiSwitch routing offload.
     */
    readonly switchControllerOffload: string;
    /**
     * Enable/disable managed FortiSwitch routing offload gateway.
     */
    readonly switchControllerOffloadGw: string;
    /**
     * IP for routing offload on FortiSwitch.
     */
    readonly switchControllerOffloadIp: string;
    /**
     * Stop Layer2 MAC learning and interception of BPDUs and other packets on this interface.
     */
    readonly switchControllerRspanMode: string;
    /**
     * Source IP address used in FortiLink over L3 connections.
     */
    readonly switchControllerSourceIp: string;
    /**
     * Switch controller traffic policy for the VLAN.
     */
    readonly switchControllerTrafficPolicy: string;
    /**
     * Define a system ID for the aggregate interface.
     */
    readonly systemId: string;
    /**
     * Method in which system ID is generated.
     */
    readonly systemIdType: string;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    readonly taggings: outputs.system.GetInterfaceTagging[];
    /**
     * TCP maximum segment size. 0 means do not change segment size.
     */
    readonly tcpMss: number;
    /**
     * Enable/disable VLAN trunk.
     */
    readonly trunk: string;
    /**
     * Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
     */
    readonly trustIp1: string;
    /**
     * Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
     */
    readonly trustIp2: string;
    /**
     * Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
     */
    readonly trustIp3: string;
    /**
     * Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
     */
    readonly trustIp61: string;
    /**
     * Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
     */
    readonly trustIp62: string;
    /**
     * Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
     */
    readonly trustIp63: string;
    /**
     * DHCP client option type.
     */
    readonly type: string;
    /**
     * Username of the PPPoE account, provided by your ISP.
     */
    readonly username: string;
    /**
     * Interface is in this virtual domain (VDOM).
     */
    readonly vdom: string;
    readonly vdomparam?: string;
    /**
     * Switch control interface VLAN ID.
     */
    readonly vindex: number;
    /**
     * Ethernet protocol of VLAN.
     */
    readonly vlanProtocol: string;
    /**
     * Enable/disable traffic forwarding between VLANs on this interface.
     */
    readonly vlanforward: string;
    /**
     * VLAN ID (1 - 4094).
     */
    readonly vlanid: number;
    /**
     * Virtual Routing Forwarding ID.
     */
    readonly vrf: number;
    /**
     * Enable/disable use of virtual MAC for VRRP.
     */
    readonly vrrpVirtualMac: string;
    /**
     * VRRP configuration. The structure of `vrrp` block is documented below.
     */
    readonly vrrps: outputs.system.GetInterfaceVrrp[];
    /**
     * Enable/disable WCCP on this interface. Used for encapsulated WCCP communication between WCCP clients and servers.
     */
    readonly wccp: string;
    /**
     * Default weight for static routes (if route has no weight configured).
     */
    readonly weight: number;
    /**
     * WINS server IP.
     */
    readonly winsIp: string;
}
/**
 * Use this data source to get information on an fortios system interface
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const sample1 = fortios.system.getInterface({
 *     name: "port1",
 * });
 * export const output1 = sample1.then(sample1 => sample1.ip);
 * ```
 */
export function getInterfaceOutput(args: GetInterfaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInterfaceResult> {
    return pulumi.output(args).apply((a: any) => getInterface(a, opts))
}

/**
 * A collection of arguments for invoking getInterface.
 */
export interface GetInterfaceOutputArgs {
    /**
     * Specify the name of the desired system interface.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
