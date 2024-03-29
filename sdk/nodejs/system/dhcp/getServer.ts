// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Use this data source to get information on an fortios systemdhcp server
 */
export function getServer(args: GetServerArgs, opts?: pulumi.InvokeOptions): Promise<GetServerResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:system/dhcp/getServer:getServer", {
        "fosid": args.fosid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getServer.
 */
export interface GetServerArgs {
    /**
     * Specify the fosid of the desired systemdhcp server.
     */
    fosid: number;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getServer.
 */
export interface GetServerResult {
    /**
     * Enable/disable auto configuration.
     */
    readonly autoConfiguration: string;
    /**
     * Enable/disable use of this DHCP server once this interface has been assigned an IP address from FortiIPAM.
     */
    readonly autoManagedStatus: string;
    /**
     * Time in seconds to wait after a conflicted IP address is removed from the DHCP range before it can be reused.
     */
    readonly conflictedIpTimeout: number;
    /**
     * DDNS authentication mode.
     */
    readonly ddnsAuth: string;
    /**
     * DDNS update key (base 64 encoding).
     */
    readonly ddnsKey: string;
    /**
     * DDNS update key name.
     */
    readonly ddnsKeyname: string;
    /**
     * DDNS server IP.
     */
    readonly ddnsServerIp: string;
    /**
     * TTL.
     */
    readonly ddnsTtl: number;
    /**
     * Enable/disable DDNS update for DHCP.
     */
    readonly ddnsUpdate: string;
    /**
     * Enable/disable DDNS update override for DHCP.
     */
    readonly ddnsUpdateOverride: string;
    /**
     * Zone of your domain name (ex. DDNS.com).
     */
    readonly ddnsZone: string;
    /**
     * Default gateway IP address assigned by the DHCP server.
     */
    readonly defaultGateway: string;
    /**
     * Enable/disable populating of DHCP server settings from FortiIPAM.
     */
    readonly dhcpSettingsFromFortiipam: string;
    /**
     * DNS server 1.
     */
    readonly dnsServer1: string;
    /**
     * DNS server 2.
     */
    readonly dnsServer2: string;
    /**
     * DNS server 3.
     */
    readonly dnsServer3: string;
    /**
     * DNS server 4.
     */
    readonly dnsServer4: string;
    /**
     * Options for assigning DNS servers to DHCP clients.
     */
    readonly dnsService: string;
    /**
     * Domain name suffix for the IP addresses that the DHCP server assigns to clients.
     */
    readonly domain: string;
    /**
     * Exclude one or more ranges of IP addresses from being assigned to clients. The structure of `excludeRange` block is documented below.
     */
    readonly excludeRanges: outputs.system.dhcp.GetServerExcludeRange[];
    /**
     * Name of the boot file on the TFTP server.
     */
    readonly filename: string;
    /**
     * Enable/disable FortiClient-On-Net service for this DHCP server.
     */
    readonly forticlientOnNetStatus: string;
    /**
     * ID.
     */
    readonly fosid: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * DHCP server can assign IP configurations to clients connected to this interface.
     */
    readonly interface: string;
    /**
     * Method used to assign client IP.
     */
    readonly ipMode: string;
    /**
     * DHCP IP range configuration. The structure of `ipRange` block is documented below.
     */
    readonly ipRanges: outputs.system.dhcp.GetServerIpRange[];
    /**
     * DHCP over IPsec leases expire this many seconds after tunnel down (0 to disable forced-expiry).
     */
    readonly ipsecLeaseHold: number;
    /**
     * Lease time in seconds, 0 means default lease time.
     */
    readonly leaseTime: number;
    /**
     * MAC access control default action (allow or block assigning IP settings).
     */
    readonly macAclDefaultAction: string;
    /**
     * Netmask assigned by the DHCP server.
     */
    readonly netmask: string;
    /**
     * IP address of a server (for example, a TFTP sever) that DHCP clients can download a boot file from.
     */
    readonly nextServer: string;
    /**
     * NTP server 1.
     */
    readonly ntpServer1: string;
    /**
     * NTP server 2.
     */
    readonly ntpServer2: string;
    /**
     * NTP server 3.
     */
    readonly ntpServer3: string;
    /**
     * Options for assigning Network Time Protocol (NTP) servers to DHCP clients.
     */
    readonly ntpService: string;
    /**
     * DHCP options. The structure of `options` block is documented below.
     */
    readonly options: outputs.system.dhcp.GetServerOption[];
    /**
     * Relay agent IP.
     */
    readonly relayAgent: string;
    /**
     * Options for the DHCP server to assign IP settings to specific MAC addresses. The structure of `reservedAddress` block is documented below.
     */
    readonly reservedAddresses: outputs.system.dhcp.GetServerReservedAddress[];
    /**
     * DHCP server can be a normal DHCP server or an IPsec DHCP server.
     */
    readonly serverType: string;
    /**
     * Enable/disable shared subnet.
     */
    readonly sharedSubnet: string;
    /**
     * Enable/disable this DHCP configuration.
     */
    readonly status: string;
    /**
     * TFTP server.
     */
    readonly tftpServers: outputs.system.dhcp.GetServerTftpServer[];
    /**
     * Select the time zone to be assigned to DHCP clients.
     */
    readonly timezone: string;
    /**
     * Options for the DHCP server to set the client's time zone.
     */
    readonly timezoneOption: string;
    /**
     * Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this range.
     */
    readonly vciMatch: string;
    /**
     * VCI strings.
     */
    readonly vciStrings: outputs.system.dhcp.GetServerVciString[];
    readonly vdomparam?: string;
    /**
     * WiFi Access Controller 1 IP address (DHCP option 138, RFC 5417).
     */
    readonly wifiAc1: string;
    /**
     * WiFi Access Controller 2 IP address (DHCP option 138, RFC 5417).
     */
    readonly wifiAc2: string;
    /**
     * WiFi Access Controller 3 IP address (DHCP option 138, RFC 5417).
     */
    readonly wifiAc3: string;
    /**
     * Options for assigning WiFi Access Controllers to DHCP clients
     */
    readonly wifiAcService: string;
    /**
     * WINS server 1.
     */
    readonly winsServer1: string;
    /**
     * WINS server 2.
     */
    readonly winsServer2: string;
}
/**
 * Use this data source to get information on an fortios systemdhcp server
 */
export function getServerOutput(args: GetServerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServerResult> {
    return pulumi.output(args).apply((a: any) => getServer(a, opts))
}

/**
 * A collection of arguments for invoking getServer.
 */
export interface GetServerOutputArgs {
    /**
     * Specify the fosid of the desired systemdhcp server.
     */
    fosid: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
