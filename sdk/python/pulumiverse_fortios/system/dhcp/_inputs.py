# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'ServerExcludeRangeArgs',
    'ServerExcludeRangeVciStringArgs',
    'ServerIpRangeArgs',
    'ServerIpRangeVciStringArgs',
    'ServerOptionArgs',
    'ServerOptionVciStringArgs',
    'ServerReservedAddressArgs',
    'ServerTftpServerArgs',
    'ServerVciStringArgs',
]

@pulumi.input_type
class ServerExcludeRangeArgs:
    def __init__(__self__, *,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 start_ip: Optional[pulumi.Input[str]] = None,
                 vci_match: Optional[pulumi.Input[str]] = None,
                 vci_strings: Optional[pulumi.Input[Sequence[pulumi.Input['ServerExcludeRangeVciStringArgs']]]] = None):
        """
        :param pulumi.Input[str] end_ip: End of IP range.
        :param pulumi.Input[int] id: ID.
        :param pulumi.Input[str] start_ip: Start of IP range.
        :param pulumi.Input[str] vci_match: Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this range. Valid values: `disable`, `enable`.
        :param pulumi.Input[Sequence[pulumi.Input['ServerExcludeRangeVciStringArgs']]] vci_strings: One or more VCI strings in quotes separated by spaces. The structure of `vci_string` block is documented below.
        """
        if end_ip is not None:
            pulumi.set(__self__, "end_ip", end_ip)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if start_ip is not None:
            pulumi.set(__self__, "start_ip", start_ip)
        if vci_match is not None:
            pulumi.set(__self__, "vci_match", vci_match)
        if vci_strings is not None:
            pulumi.set(__self__, "vci_strings", vci_strings)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> Optional[pulumi.Input[str]]:
        """
        End of IP range.
        """
        return pulumi.get(self, "end_ip")

    @end_ip.setter
    def end_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_ip", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="startIp")
    def start_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Start of IP range.
        """
        return pulumi.get(self, "start_ip")

    @start_ip.setter
    def start_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_ip", value)

    @property
    @pulumi.getter(name="vciMatch")
    def vci_match(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this range. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "vci_match")

    @vci_match.setter
    def vci_match(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vci_match", value)

    @property
    @pulumi.getter(name="vciStrings")
    def vci_strings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServerExcludeRangeVciStringArgs']]]]:
        """
        One or more VCI strings in quotes separated by spaces. The structure of `vci_string` block is documented below.
        """
        return pulumi.get(self, "vci_strings")

    @vci_strings.setter
    def vci_strings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServerExcludeRangeVciStringArgs']]]]):
        pulumi.set(self, "vci_strings", value)


@pulumi.input_type
class ServerExcludeRangeVciStringArgs:
    def __init__(__self__, *,
                 vci_string: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] vci_string: VCI strings.
        """
        if vci_string is not None:
            pulumi.set(__self__, "vci_string", vci_string)

    @property
    @pulumi.getter(name="vciString")
    def vci_string(self) -> Optional[pulumi.Input[str]]:
        """
        VCI strings.
        """
        return pulumi.get(self, "vci_string")

    @vci_string.setter
    def vci_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vci_string", value)


@pulumi.input_type
class ServerIpRangeArgs:
    def __init__(__self__, *,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 start_ip: Optional[pulumi.Input[str]] = None,
                 vci_match: Optional[pulumi.Input[str]] = None,
                 vci_strings: Optional[pulumi.Input[Sequence[pulumi.Input['ServerIpRangeVciStringArgs']]]] = None):
        """
        :param pulumi.Input[str] end_ip: End of IP range.
        :param pulumi.Input[int] id: ID.
        :param pulumi.Input[str] start_ip: Start of IP range.
        :param pulumi.Input[str] vci_match: Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this range. Valid values: `disable`, `enable`.
        :param pulumi.Input[Sequence[pulumi.Input['ServerIpRangeVciStringArgs']]] vci_strings: One or more VCI strings in quotes separated by spaces. The structure of `vci_string` block is documented below.
        """
        if end_ip is not None:
            pulumi.set(__self__, "end_ip", end_ip)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if start_ip is not None:
            pulumi.set(__self__, "start_ip", start_ip)
        if vci_match is not None:
            pulumi.set(__self__, "vci_match", vci_match)
        if vci_strings is not None:
            pulumi.set(__self__, "vci_strings", vci_strings)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> Optional[pulumi.Input[str]]:
        """
        End of IP range.
        """
        return pulumi.get(self, "end_ip")

    @end_ip.setter
    def end_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_ip", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="startIp")
    def start_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Start of IP range.
        """
        return pulumi.get(self, "start_ip")

    @start_ip.setter
    def start_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_ip", value)

    @property
    @pulumi.getter(name="vciMatch")
    def vci_match(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this range. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "vci_match")

    @vci_match.setter
    def vci_match(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vci_match", value)

    @property
    @pulumi.getter(name="vciStrings")
    def vci_strings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServerIpRangeVciStringArgs']]]]:
        """
        One or more VCI strings in quotes separated by spaces. The structure of `vci_string` block is documented below.
        """
        return pulumi.get(self, "vci_strings")

    @vci_strings.setter
    def vci_strings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServerIpRangeVciStringArgs']]]]):
        pulumi.set(self, "vci_strings", value)


@pulumi.input_type
class ServerIpRangeVciStringArgs:
    def __init__(__self__, *,
                 vci_string: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] vci_string: VCI strings.
        """
        if vci_string is not None:
            pulumi.set(__self__, "vci_string", vci_string)

    @property
    @pulumi.getter(name="vciString")
    def vci_string(self) -> Optional[pulumi.Input[str]]:
        """
        VCI strings.
        """
        return pulumi.get(self, "vci_string")

    @vci_string.setter
    def vci_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vci_string", value)


@pulumi.input_type
class ServerOptionArgs:
    def __init__(__self__, *,
                 code: Optional[pulumi.Input[int]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 vci_match: Optional[pulumi.Input[str]] = None,
                 vci_strings: Optional[pulumi.Input[Sequence[pulumi.Input['ServerOptionVciStringArgs']]]] = None):
        """
        :param pulumi.Input[int] code: DHCP option code.
        :param pulumi.Input[int] id: ID.
        :param pulumi.Input[str] ip: DHCP option IPs.
        :param pulumi.Input[str] type: DHCP option type. Valid values: `hex`, `string`, `ip`, `fqdn`.
        :param pulumi.Input[str] value: DHCP option value.
        :param pulumi.Input[str] vci_match: Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this option. Valid values: `disable`, `enable`.
        :param pulumi.Input[Sequence[pulumi.Input['ServerOptionVciStringArgs']]] vci_strings: One or more VCI strings in quotes separated by spaces. The structure of `vci_string` block is documented below.
        """
        if code is not None:
            pulumi.set(__self__, "code", code)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)
        if vci_match is not None:
            pulumi.set(__self__, "vci_match", vci_match)
        if vci_strings is not None:
            pulumi.set(__self__, "vci_strings", vci_strings)

    @property
    @pulumi.getter
    def code(self) -> Optional[pulumi.Input[int]]:
        """
        DHCP option code.
        """
        return pulumi.get(self, "code")

    @code.setter
    def code(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "code", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        DHCP option IPs.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        DHCP option type. Valid values: `hex`, `string`, `ip`, `fqdn`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        DHCP option value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="vciMatch")
    def vci_match(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this option. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "vci_match")

    @vci_match.setter
    def vci_match(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vci_match", value)

    @property
    @pulumi.getter(name="vciStrings")
    def vci_strings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServerOptionVciStringArgs']]]]:
        """
        One or more VCI strings in quotes separated by spaces. The structure of `vci_string` block is documented below.
        """
        return pulumi.get(self, "vci_strings")

    @vci_strings.setter
    def vci_strings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServerOptionVciStringArgs']]]]):
        pulumi.set(self, "vci_strings", value)


@pulumi.input_type
class ServerOptionVciStringArgs:
    def __init__(__self__, *,
                 vci_string: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] vci_string: VCI strings.
        """
        if vci_string is not None:
            pulumi.set(__self__, "vci_string", vci_string)

    @property
    @pulumi.getter(name="vciString")
    def vci_string(self) -> Optional[pulumi.Input[str]]:
        """
        VCI strings.
        """
        return pulumi.get(self, "vci_string")

    @vci_string.setter
    def vci_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vci_string", value)


@pulumi.input_type
class ServerReservedAddressArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 circuit_id: Optional[pulumi.Input[str]] = None,
                 circuit_id_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 remote_id: Optional[pulumi.Input[str]] = None,
                 remote_id_type: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] action: Options for the DHCP server to configure the client with the reserved MAC address. Valid values: `assign`, `block`, `reserved`.
        :param pulumi.Input[str] circuit_id: Option 82 circuit-ID of the client that will get the reserved IP address.
        :param pulumi.Input[str] circuit_id_type: DHCP option type. Valid values: `hex`, `string`.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[int] id: ID.
        :param pulumi.Input[str] ip: IP address to be reserved for the MAC address.
        :param pulumi.Input[str] mac: MAC address of the client that will get the reserved IP address.
        :param pulumi.Input[str] remote_id: Option 82 remote-ID of the client that will get the reserved IP address.
        :param pulumi.Input[str] remote_id_type: DHCP option type. Valid values: `hex`, `string`.
        :param pulumi.Input[str] type: DHCP reserved-address type. Valid values: `mac`, `option82`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if circuit_id is not None:
            pulumi.set(__self__, "circuit_id", circuit_id)
        if circuit_id_type is not None:
            pulumi.set(__self__, "circuit_id_type", circuit_id_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if mac is not None:
            pulumi.set(__self__, "mac", mac)
        if remote_id is not None:
            pulumi.set(__self__, "remote_id", remote_id)
        if remote_id_type is not None:
            pulumi.set(__self__, "remote_id_type", remote_id_type)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Options for the DHCP server to configure the client with the reserved MAC address. Valid values: `assign`, `block`, `reserved`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="circuitId")
    def circuit_id(self) -> Optional[pulumi.Input[str]]:
        """
        Option 82 circuit-ID of the client that will get the reserved IP address.
        """
        return pulumi.get(self, "circuit_id")

    @circuit_id.setter
    def circuit_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "circuit_id", value)

    @property
    @pulumi.getter(name="circuitIdType")
    def circuit_id_type(self) -> Optional[pulumi.Input[str]]:
        """
        DHCP option type. Valid values: `hex`, `string`.
        """
        return pulumi.get(self, "circuit_id_type")

    @circuit_id_type.setter
    def circuit_id_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "circuit_id_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP address to be reserved for the MAC address.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def mac(self) -> Optional[pulumi.Input[str]]:
        """
        MAC address of the client that will get the reserved IP address.
        """
        return pulumi.get(self, "mac")

    @mac.setter
    def mac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac", value)

    @property
    @pulumi.getter(name="remoteId")
    def remote_id(self) -> Optional[pulumi.Input[str]]:
        """
        Option 82 remote-ID of the client that will get the reserved IP address.
        """
        return pulumi.get(self, "remote_id")

    @remote_id.setter
    def remote_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_id", value)

    @property
    @pulumi.getter(name="remoteIdType")
    def remote_id_type(self) -> Optional[pulumi.Input[str]]:
        """
        DHCP option type. Valid values: `hex`, `string`.
        """
        return pulumi.get(self, "remote_id_type")

    @remote_id_type.setter
    def remote_id_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_id_type", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        DHCP reserved-address type. Valid values: `mac`, `option82`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ServerTftpServerArgs:
    def __init__(__self__, *,
                 tftp_server: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] tftp_server: TFTP server.
        """
        if tftp_server is not None:
            pulumi.set(__self__, "tftp_server", tftp_server)

    @property
    @pulumi.getter(name="tftpServer")
    def tftp_server(self) -> Optional[pulumi.Input[str]]:
        """
        TFTP server.
        """
        return pulumi.get(self, "tftp_server")

    @tftp_server.setter
    def tftp_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tftp_server", value)


@pulumi.input_type
class ServerVciStringArgs:
    def __init__(__self__, *,
                 vci_string: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] vci_string: VCI strings.
        """
        if vci_string is not None:
            pulumi.set(__self__, "vci_string", vci_string)

    @property
    @pulumi.getter(name="vciString")
    def vci_string(self) -> Optional[pulumi.Input[str]]:
        """
        VCI strings.
        """
        return pulumi.get(self, "vci_string")

    @vci_string.setter
    def vci_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vci_string", value)


