# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetServerResult',
    'AwaitableGetServerResult',
    'get_server',
    'get_server_output',
]

@pulumi.output_type
class GetServerResult:
    """
    A collection of values returned by getServer.
    """
    def __init__(__self__, auto_configuration=None, auto_managed_status=None, conflicted_ip_timeout=None, ddns_auth=None, ddns_key=None, ddns_keyname=None, ddns_server_ip=None, ddns_ttl=None, ddns_update=None, ddns_update_override=None, ddns_zone=None, default_gateway=None, dhcp_settings_from_fortiipam=None, dns_server1=None, dns_server2=None, dns_server3=None, dns_server4=None, dns_service=None, domain=None, exclude_ranges=None, filename=None, forticlient_on_net_status=None, fosid=None, id=None, interface=None, ip_mode=None, ip_ranges=None, ipsec_lease_hold=None, lease_time=None, mac_acl_default_action=None, netmask=None, next_server=None, ntp_server1=None, ntp_server2=None, ntp_server3=None, ntp_service=None, options=None, reserved_addresses=None, server_type=None, status=None, tftp_servers=None, timezone=None, timezone_option=None, vci_match=None, vci_strings=None, vdomparam=None, wifi_ac1=None, wifi_ac2=None, wifi_ac3=None, wifi_ac_service=None, wins_server1=None, wins_server2=None):
        if auto_configuration and not isinstance(auto_configuration, str):
            raise TypeError("Expected argument 'auto_configuration' to be a str")
        pulumi.set(__self__, "auto_configuration", auto_configuration)
        if auto_managed_status and not isinstance(auto_managed_status, str):
            raise TypeError("Expected argument 'auto_managed_status' to be a str")
        pulumi.set(__self__, "auto_managed_status", auto_managed_status)
        if conflicted_ip_timeout and not isinstance(conflicted_ip_timeout, int):
            raise TypeError("Expected argument 'conflicted_ip_timeout' to be a int")
        pulumi.set(__self__, "conflicted_ip_timeout", conflicted_ip_timeout)
        if ddns_auth and not isinstance(ddns_auth, str):
            raise TypeError("Expected argument 'ddns_auth' to be a str")
        pulumi.set(__self__, "ddns_auth", ddns_auth)
        if ddns_key and not isinstance(ddns_key, str):
            raise TypeError("Expected argument 'ddns_key' to be a str")
        pulumi.set(__self__, "ddns_key", ddns_key)
        if ddns_keyname and not isinstance(ddns_keyname, str):
            raise TypeError("Expected argument 'ddns_keyname' to be a str")
        pulumi.set(__self__, "ddns_keyname", ddns_keyname)
        if ddns_server_ip and not isinstance(ddns_server_ip, str):
            raise TypeError("Expected argument 'ddns_server_ip' to be a str")
        pulumi.set(__self__, "ddns_server_ip", ddns_server_ip)
        if ddns_ttl and not isinstance(ddns_ttl, int):
            raise TypeError("Expected argument 'ddns_ttl' to be a int")
        pulumi.set(__self__, "ddns_ttl", ddns_ttl)
        if ddns_update and not isinstance(ddns_update, str):
            raise TypeError("Expected argument 'ddns_update' to be a str")
        pulumi.set(__self__, "ddns_update", ddns_update)
        if ddns_update_override and not isinstance(ddns_update_override, str):
            raise TypeError("Expected argument 'ddns_update_override' to be a str")
        pulumi.set(__self__, "ddns_update_override", ddns_update_override)
        if ddns_zone and not isinstance(ddns_zone, str):
            raise TypeError("Expected argument 'ddns_zone' to be a str")
        pulumi.set(__self__, "ddns_zone", ddns_zone)
        if default_gateway and not isinstance(default_gateway, str):
            raise TypeError("Expected argument 'default_gateway' to be a str")
        pulumi.set(__self__, "default_gateway", default_gateway)
        if dhcp_settings_from_fortiipam and not isinstance(dhcp_settings_from_fortiipam, str):
            raise TypeError("Expected argument 'dhcp_settings_from_fortiipam' to be a str")
        pulumi.set(__self__, "dhcp_settings_from_fortiipam", dhcp_settings_from_fortiipam)
        if dns_server1 and not isinstance(dns_server1, str):
            raise TypeError("Expected argument 'dns_server1' to be a str")
        pulumi.set(__self__, "dns_server1", dns_server1)
        if dns_server2 and not isinstance(dns_server2, str):
            raise TypeError("Expected argument 'dns_server2' to be a str")
        pulumi.set(__self__, "dns_server2", dns_server2)
        if dns_server3 and not isinstance(dns_server3, str):
            raise TypeError("Expected argument 'dns_server3' to be a str")
        pulumi.set(__self__, "dns_server3", dns_server3)
        if dns_server4 and not isinstance(dns_server4, str):
            raise TypeError("Expected argument 'dns_server4' to be a str")
        pulumi.set(__self__, "dns_server4", dns_server4)
        if dns_service and not isinstance(dns_service, str):
            raise TypeError("Expected argument 'dns_service' to be a str")
        pulumi.set(__self__, "dns_service", dns_service)
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if exclude_ranges and not isinstance(exclude_ranges, list):
            raise TypeError("Expected argument 'exclude_ranges' to be a list")
        pulumi.set(__self__, "exclude_ranges", exclude_ranges)
        if filename and not isinstance(filename, str):
            raise TypeError("Expected argument 'filename' to be a str")
        pulumi.set(__self__, "filename", filename)
        if forticlient_on_net_status and not isinstance(forticlient_on_net_status, str):
            raise TypeError("Expected argument 'forticlient_on_net_status' to be a str")
        pulumi.set(__self__, "forticlient_on_net_status", forticlient_on_net_status)
        if fosid and not isinstance(fosid, int):
            raise TypeError("Expected argument 'fosid' to be a int")
        pulumi.set(__self__, "fosid", fosid)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface and not isinstance(interface, str):
            raise TypeError("Expected argument 'interface' to be a str")
        pulumi.set(__self__, "interface", interface)
        if ip_mode and not isinstance(ip_mode, str):
            raise TypeError("Expected argument 'ip_mode' to be a str")
        pulumi.set(__self__, "ip_mode", ip_mode)
        if ip_ranges and not isinstance(ip_ranges, list):
            raise TypeError("Expected argument 'ip_ranges' to be a list")
        pulumi.set(__self__, "ip_ranges", ip_ranges)
        if ipsec_lease_hold and not isinstance(ipsec_lease_hold, int):
            raise TypeError("Expected argument 'ipsec_lease_hold' to be a int")
        pulumi.set(__self__, "ipsec_lease_hold", ipsec_lease_hold)
        if lease_time and not isinstance(lease_time, int):
            raise TypeError("Expected argument 'lease_time' to be a int")
        pulumi.set(__self__, "lease_time", lease_time)
        if mac_acl_default_action and not isinstance(mac_acl_default_action, str):
            raise TypeError("Expected argument 'mac_acl_default_action' to be a str")
        pulumi.set(__self__, "mac_acl_default_action", mac_acl_default_action)
        if netmask and not isinstance(netmask, str):
            raise TypeError("Expected argument 'netmask' to be a str")
        pulumi.set(__self__, "netmask", netmask)
        if next_server and not isinstance(next_server, str):
            raise TypeError("Expected argument 'next_server' to be a str")
        pulumi.set(__self__, "next_server", next_server)
        if ntp_server1 and not isinstance(ntp_server1, str):
            raise TypeError("Expected argument 'ntp_server1' to be a str")
        pulumi.set(__self__, "ntp_server1", ntp_server1)
        if ntp_server2 and not isinstance(ntp_server2, str):
            raise TypeError("Expected argument 'ntp_server2' to be a str")
        pulumi.set(__self__, "ntp_server2", ntp_server2)
        if ntp_server3 and not isinstance(ntp_server3, str):
            raise TypeError("Expected argument 'ntp_server3' to be a str")
        pulumi.set(__self__, "ntp_server3", ntp_server3)
        if ntp_service and not isinstance(ntp_service, str):
            raise TypeError("Expected argument 'ntp_service' to be a str")
        pulumi.set(__self__, "ntp_service", ntp_service)
        if options and not isinstance(options, list):
            raise TypeError("Expected argument 'options' to be a list")
        pulumi.set(__self__, "options", options)
        if reserved_addresses and not isinstance(reserved_addresses, list):
            raise TypeError("Expected argument 'reserved_addresses' to be a list")
        pulumi.set(__self__, "reserved_addresses", reserved_addresses)
        if server_type and not isinstance(server_type, str):
            raise TypeError("Expected argument 'server_type' to be a str")
        pulumi.set(__self__, "server_type", server_type)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tftp_servers and not isinstance(tftp_servers, list):
            raise TypeError("Expected argument 'tftp_servers' to be a list")
        pulumi.set(__self__, "tftp_servers", tftp_servers)
        if timezone and not isinstance(timezone, str):
            raise TypeError("Expected argument 'timezone' to be a str")
        pulumi.set(__self__, "timezone", timezone)
        if timezone_option and not isinstance(timezone_option, str):
            raise TypeError("Expected argument 'timezone_option' to be a str")
        pulumi.set(__self__, "timezone_option", timezone_option)
        if vci_match and not isinstance(vci_match, str):
            raise TypeError("Expected argument 'vci_match' to be a str")
        pulumi.set(__self__, "vci_match", vci_match)
        if vci_strings and not isinstance(vci_strings, list):
            raise TypeError("Expected argument 'vci_strings' to be a list")
        pulumi.set(__self__, "vci_strings", vci_strings)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)
        if wifi_ac1 and not isinstance(wifi_ac1, str):
            raise TypeError("Expected argument 'wifi_ac1' to be a str")
        pulumi.set(__self__, "wifi_ac1", wifi_ac1)
        if wifi_ac2 and not isinstance(wifi_ac2, str):
            raise TypeError("Expected argument 'wifi_ac2' to be a str")
        pulumi.set(__self__, "wifi_ac2", wifi_ac2)
        if wifi_ac3 and not isinstance(wifi_ac3, str):
            raise TypeError("Expected argument 'wifi_ac3' to be a str")
        pulumi.set(__self__, "wifi_ac3", wifi_ac3)
        if wifi_ac_service and not isinstance(wifi_ac_service, str):
            raise TypeError("Expected argument 'wifi_ac_service' to be a str")
        pulumi.set(__self__, "wifi_ac_service", wifi_ac_service)
        if wins_server1 and not isinstance(wins_server1, str):
            raise TypeError("Expected argument 'wins_server1' to be a str")
        pulumi.set(__self__, "wins_server1", wins_server1)
        if wins_server2 and not isinstance(wins_server2, str):
            raise TypeError("Expected argument 'wins_server2' to be a str")
        pulumi.set(__self__, "wins_server2", wins_server2)

    @property
    @pulumi.getter(name="autoConfiguration")
    def auto_configuration(self) -> str:
        """
        Enable/disable auto configuration.
        """
        return pulumi.get(self, "auto_configuration")

    @property
    @pulumi.getter(name="autoManagedStatus")
    def auto_managed_status(self) -> str:
        """
        Enable/disable use of this DHCP server once this interface has been assigned an IP address from FortiIPAM.
        """
        return pulumi.get(self, "auto_managed_status")

    @property
    @pulumi.getter(name="conflictedIpTimeout")
    def conflicted_ip_timeout(self) -> int:
        """
        Time in seconds to wait after a conflicted IP address is removed from the DHCP range before it can be reused.
        """
        return pulumi.get(self, "conflicted_ip_timeout")

    @property
    @pulumi.getter(name="ddnsAuth")
    def ddns_auth(self) -> str:
        """
        DDNS authentication mode.
        """
        return pulumi.get(self, "ddns_auth")

    @property
    @pulumi.getter(name="ddnsKey")
    def ddns_key(self) -> str:
        """
        DDNS update key (base 64 encoding).
        """
        return pulumi.get(self, "ddns_key")

    @property
    @pulumi.getter(name="ddnsKeyname")
    def ddns_keyname(self) -> str:
        """
        DDNS update key name.
        """
        return pulumi.get(self, "ddns_keyname")

    @property
    @pulumi.getter(name="ddnsServerIp")
    def ddns_server_ip(self) -> str:
        """
        DDNS server IP.
        """
        return pulumi.get(self, "ddns_server_ip")

    @property
    @pulumi.getter(name="ddnsTtl")
    def ddns_ttl(self) -> int:
        """
        TTL.
        """
        return pulumi.get(self, "ddns_ttl")

    @property
    @pulumi.getter(name="ddnsUpdate")
    def ddns_update(self) -> str:
        """
        Enable/disable DDNS update for DHCP.
        """
        return pulumi.get(self, "ddns_update")

    @property
    @pulumi.getter(name="ddnsUpdateOverride")
    def ddns_update_override(self) -> str:
        """
        Enable/disable DDNS update override for DHCP.
        """
        return pulumi.get(self, "ddns_update_override")

    @property
    @pulumi.getter(name="ddnsZone")
    def ddns_zone(self) -> str:
        """
        Zone of your domain name (ex. DDNS.com).
        """
        return pulumi.get(self, "ddns_zone")

    @property
    @pulumi.getter(name="defaultGateway")
    def default_gateway(self) -> str:
        """
        Default gateway IP address assigned by the DHCP server.
        """
        return pulumi.get(self, "default_gateway")

    @property
    @pulumi.getter(name="dhcpSettingsFromFortiipam")
    def dhcp_settings_from_fortiipam(self) -> str:
        """
        Enable/disable populating of DHCP server settings from FortiIPAM.
        """
        return pulumi.get(self, "dhcp_settings_from_fortiipam")

    @property
    @pulumi.getter(name="dnsServer1")
    def dns_server1(self) -> str:
        """
        DNS server 1.
        """
        return pulumi.get(self, "dns_server1")

    @property
    @pulumi.getter(name="dnsServer2")
    def dns_server2(self) -> str:
        """
        DNS server 2.
        """
        return pulumi.get(self, "dns_server2")

    @property
    @pulumi.getter(name="dnsServer3")
    def dns_server3(self) -> str:
        """
        DNS server 3.
        """
        return pulumi.get(self, "dns_server3")

    @property
    @pulumi.getter(name="dnsServer4")
    def dns_server4(self) -> str:
        """
        DNS server 4.
        """
        return pulumi.get(self, "dns_server4")

    @property
    @pulumi.getter(name="dnsService")
    def dns_service(self) -> str:
        """
        Options for assigning DNS servers to DHCP clients.
        """
        return pulumi.get(self, "dns_service")

    @property
    @pulumi.getter
    def domain(self) -> str:
        """
        Domain name suffix for the IP addresses that the DHCP server assigns to clients.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="excludeRanges")
    def exclude_ranges(self) -> Sequence['outputs.GetServerExcludeRangeResult']:
        """
        Exclude one or more ranges of IP addresses from being assigned to clients. The structure of `exclude_range` block is documented below.
        """
        return pulumi.get(self, "exclude_ranges")

    @property
    @pulumi.getter
    def filename(self) -> str:
        """
        Name of the boot file on the TFTP server.
        """
        return pulumi.get(self, "filename")

    @property
    @pulumi.getter(name="forticlientOnNetStatus")
    def forticlient_on_net_status(self) -> str:
        """
        Enable/disable FortiClient-On-Net service for this DHCP server.
        """
        return pulumi.get(self, "forticlient_on_net_status")

    @property
    @pulumi.getter
    def fosid(self) -> int:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interface(self) -> str:
        """
        DHCP server can assign IP configurations to clients connected to this interface.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter(name="ipMode")
    def ip_mode(self) -> str:
        """
        Method used to assign client IP.
        """
        return pulumi.get(self, "ip_mode")

    @property
    @pulumi.getter(name="ipRanges")
    def ip_ranges(self) -> Sequence['outputs.GetServerIpRangeResult']:
        """
        DHCP IP range configuration. The structure of `ip_range` block is documented below.
        """
        return pulumi.get(self, "ip_ranges")

    @property
    @pulumi.getter(name="ipsecLeaseHold")
    def ipsec_lease_hold(self) -> int:
        """
        DHCP over IPsec leases expire this many seconds after tunnel down (0 to disable forced-expiry).
        """
        return pulumi.get(self, "ipsec_lease_hold")

    @property
    @pulumi.getter(name="leaseTime")
    def lease_time(self) -> int:
        """
        Lease time in seconds, 0 means unlimited.
        """
        return pulumi.get(self, "lease_time")

    @property
    @pulumi.getter(name="macAclDefaultAction")
    def mac_acl_default_action(self) -> str:
        """
        MAC access control default action (allow or block assigning IP settings).
        """
        return pulumi.get(self, "mac_acl_default_action")

    @property
    @pulumi.getter
    def netmask(self) -> str:
        """
        Netmask assigned by the DHCP server.
        """
        return pulumi.get(self, "netmask")

    @property
    @pulumi.getter(name="nextServer")
    def next_server(self) -> str:
        """
        IP address of a server (for example, a TFTP sever) that DHCP clients can download a boot file from.
        """
        return pulumi.get(self, "next_server")

    @property
    @pulumi.getter(name="ntpServer1")
    def ntp_server1(self) -> str:
        """
        NTP server 1.
        """
        return pulumi.get(self, "ntp_server1")

    @property
    @pulumi.getter(name="ntpServer2")
    def ntp_server2(self) -> str:
        """
        NTP server 2.
        """
        return pulumi.get(self, "ntp_server2")

    @property
    @pulumi.getter(name="ntpServer3")
    def ntp_server3(self) -> str:
        """
        NTP server 3.
        """
        return pulumi.get(self, "ntp_server3")

    @property
    @pulumi.getter(name="ntpService")
    def ntp_service(self) -> str:
        """
        Options for assigning Network Time Protocol (NTP) servers to DHCP clients.
        """
        return pulumi.get(self, "ntp_service")

    @property
    @pulumi.getter
    def options(self) -> Sequence['outputs.GetServerOptionResult']:
        """
        DHCP options. The structure of `options` block is documented below.
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter(name="reservedAddresses")
    def reserved_addresses(self) -> Sequence['outputs.GetServerReservedAddressResult']:
        """
        Options for the DHCP server to assign IP settings to specific MAC addresses. The structure of `reserved_address` block is documented below.
        """
        return pulumi.get(self, "reserved_addresses")

    @property
    @pulumi.getter(name="serverType")
    def server_type(self) -> str:
        """
        DHCP server can be a normal DHCP server or an IPsec DHCP server.
        """
        return pulumi.get(self, "server_type")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Enable/disable this DHCP configuration.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="tftpServers")
    def tftp_servers(self) -> Sequence['outputs.GetServerTftpServerResult']:
        """
        TFTP server.
        """
        return pulumi.get(self, "tftp_servers")

    @property
    @pulumi.getter
    def timezone(self) -> str:
        """
        Select the time zone to be assigned to DHCP clients.
        """
        return pulumi.get(self, "timezone")

    @property
    @pulumi.getter(name="timezoneOption")
    def timezone_option(self) -> str:
        """
        Options for the DHCP server to set the client's time zone.
        """
        return pulumi.get(self, "timezone_option")

    @property
    @pulumi.getter(name="vciMatch")
    def vci_match(self) -> str:
        """
        Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this range.
        """
        return pulumi.get(self, "vci_match")

    @property
    @pulumi.getter(name="vciStrings")
    def vci_strings(self) -> Sequence['outputs.GetServerVciStringResult']:
        """
        VCI strings.
        """
        return pulumi.get(self, "vci_strings")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter(name="wifiAc1")
    def wifi_ac1(self) -> str:
        """
        WiFi Access Controller 1 IP address (DHCP option 138, RFC 5417).
        """
        return pulumi.get(self, "wifi_ac1")

    @property
    @pulumi.getter(name="wifiAc2")
    def wifi_ac2(self) -> str:
        """
        WiFi Access Controller 2 IP address (DHCP option 138, RFC 5417).
        """
        return pulumi.get(self, "wifi_ac2")

    @property
    @pulumi.getter(name="wifiAc3")
    def wifi_ac3(self) -> str:
        """
        WiFi Access Controller 3 IP address (DHCP option 138, RFC 5417).
        """
        return pulumi.get(self, "wifi_ac3")

    @property
    @pulumi.getter(name="wifiAcService")
    def wifi_ac_service(self) -> str:
        """
        Options for assigning WiFi Access Controllers to DHCP clients
        """
        return pulumi.get(self, "wifi_ac_service")

    @property
    @pulumi.getter(name="winsServer1")
    def wins_server1(self) -> str:
        """
        WINS server 1.
        """
        return pulumi.get(self, "wins_server1")

    @property
    @pulumi.getter(name="winsServer2")
    def wins_server2(self) -> str:
        """
        WINS server 2.
        """
        return pulumi.get(self, "wins_server2")


class AwaitableGetServerResult(GetServerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServerResult(
            auto_configuration=self.auto_configuration,
            auto_managed_status=self.auto_managed_status,
            conflicted_ip_timeout=self.conflicted_ip_timeout,
            ddns_auth=self.ddns_auth,
            ddns_key=self.ddns_key,
            ddns_keyname=self.ddns_keyname,
            ddns_server_ip=self.ddns_server_ip,
            ddns_ttl=self.ddns_ttl,
            ddns_update=self.ddns_update,
            ddns_update_override=self.ddns_update_override,
            ddns_zone=self.ddns_zone,
            default_gateway=self.default_gateway,
            dhcp_settings_from_fortiipam=self.dhcp_settings_from_fortiipam,
            dns_server1=self.dns_server1,
            dns_server2=self.dns_server2,
            dns_server3=self.dns_server3,
            dns_server4=self.dns_server4,
            dns_service=self.dns_service,
            domain=self.domain,
            exclude_ranges=self.exclude_ranges,
            filename=self.filename,
            forticlient_on_net_status=self.forticlient_on_net_status,
            fosid=self.fosid,
            id=self.id,
            interface=self.interface,
            ip_mode=self.ip_mode,
            ip_ranges=self.ip_ranges,
            ipsec_lease_hold=self.ipsec_lease_hold,
            lease_time=self.lease_time,
            mac_acl_default_action=self.mac_acl_default_action,
            netmask=self.netmask,
            next_server=self.next_server,
            ntp_server1=self.ntp_server1,
            ntp_server2=self.ntp_server2,
            ntp_server3=self.ntp_server3,
            ntp_service=self.ntp_service,
            options=self.options,
            reserved_addresses=self.reserved_addresses,
            server_type=self.server_type,
            status=self.status,
            tftp_servers=self.tftp_servers,
            timezone=self.timezone,
            timezone_option=self.timezone_option,
            vci_match=self.vci_match,
            vci_strings=self.vci_strings,
            vdomparam=self.vdomparam,
            wifi_ac1=self.wifi_ac1,
            wifi_ac2=self.wifi_ac2,
            wifi_ac3=self.wifi_ac3,
            wifi_ac_service=self.wifi_ac_service,
            wins_server1=self.wins_server1,
            wins_server2=self.wins_server2)


def get_server(fosid: Optional[int] = None,
               vdomparam: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServerResult:
    """
    Use this data source to get information on an fortios systemdhcp server


    :param int fosid: Specify the fosid of the desired systemdhcp server.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['fosid'] = fosid
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:systemdhcp/getServer:getServer', __args__, opts=opts, typ=GetServerResult).value

    return AwaitableGetServerResult(
        auto_configuration=__ret__.auto_configuration,
        auto_managed_status=__ret__.auto_managed_status,
        conflicted_ip_timeout=__ret__.conflicted_ip_timeout,
        ddns_auth=__ret__.ddns_auth,
        ddns_key=__ret__.ddns_key,
        ddns_keyname=__ret__.ddns_keyname,
        ddns_server_ip=__ret__.ddns_server_ip,
        ddns_ttl=__ret__.ddns_ttl,
        ddns_update=__ret__.ddns_update,
        ddns_update_override=__ret__.ddns_update_override,
        ddns_zone=__ret__.ddns_zone,
        default_gateway=__ret__.default_gateway,
        dhcp_settings_from_fortiipam=__ret__.dhcp_settings_from_fortiipam,
        dns_server1=__ret__.dns_server1,
        dns_server2=__ret__.dns_server2,
        dns_server3=__ret__.dns_server3,
        dns_server4=__ret__.dns_server4,
        dns_service=__ret__.dns_service,
        domain=__ret__.domain,
        exclude_ranges=__ret__.exclude_ranges,
        filename=__ret__.filename,
        forticlient_on_net_status=__ret__.forticlient_on_net_status,
        fosid=__ret__.fosid,
        id=__ret__.id,
        interface=__ret__.interface,
        ip_mode=__ret__.ip_mode,
        ip_ranges=__ret__.ip_ranges,
        ipsec_lease_hold=__ret__.ipsec_lease_hold,
        lease_time=__ret__.lease_time,
        mac_acl_default_action=__ret__.mac_acl_default_action,
        netmask=__ret__.netmask,
        next_server=__ret__.next_server,
        ntp_server1=__ret__.ntp_server1,
        ntp_server2=__ret__.ntp_server2,
        ntp_server3=__ret__.ntp_server3,
        ntp_service=__ret__.ntp_service,
        options=__ret__.options,
        reserved_addresses=__ret__.reserved_addresses,
        server_type=__ret__.server_type,
        status=__ret__.status,
        tftp_servers=__ret__.tftp_servers,
        timezone=__ret__.timezone,
        timezone_option=__ret__.timezone_option,
        vci_match=__ret__.vci_match,
        vci_strings=__ret__.vci_strings,
        vdomparam=__ret__.vdomparam,
        wifi_ac1=__ret__.wifi_ac1,
        wifi_ac2=__ret__.wifi_ac2,
        wifi_ac3=__ret__.wifi_ac3,
        wifi_ac_service=__ret__.wifi_ac_service,
        wins_server1=__ret__.wins_server1,
        wins_server2=__ret__.wins_server2)


@_utilities.lift_output_func(get_server)
def get_server_output(fosid: Optional[pulumi.Input[int]] = None,
                      vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServerResult]:
    """
    Use this data source to get information on an fortios systemdhcp server


    :param int fosid: Specify the fosid of the desired systemdhcp server.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
