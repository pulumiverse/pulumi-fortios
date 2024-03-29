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
    'CommunityHost',
    'CommunityHosts6',
    'CommunityVdom',
    'UserVdom',
    'GetCommunityHostResult',
    'GetCommunityHosts6Result',
    'GetCommunityVdomResult',
    'GetUserVdomResult',
]

@pulumi.output_type
class CommunityHost(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "haDirect":
            suggest = "ha_direct"
        elif key == "hostType":
            suggest = "host_type"
        elif key == "sourceIp":
            suggest = "source_ip"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CommunityHost. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CommunityHost.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CommunityHost.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 ha_direct: Optional[str] = None,
                 host_type: Optional[str] = None,
                 id: Optional[int] = None,
                 ip: Optional[str] = None,
                 source_ip: Optional[str] = None):
        """
        :param str ha_direct: Enable/disable direct management of HA cluster members. Valid values: `enable`, `disable`.
        :param str host_type: Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both. Valid values: `any`, `query`, `trap`.
        :param int id: Host6 entry ID.
        :param str ip: IPv4 address of the SNMP manager (host).
        :param str source_ip: Source IPv4 address for SNMP traps.
        """
        if ha_direct is not None:
            pulumi.set(__self__, "ha_direct", ha_direct)
        if host_type is not None:
            pulumi.set(__self__, "host_type", host_type)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)

    @property
    @pulumi.getter(name="haDirect")
    def ha_direct(self) -> Optional[str]:
        """
        Enable/disable direct management of HA cluster members. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ha_direct")

    @property
    @pulumi.getter(name="hostType")
    def host_type(self) -> Optional[str]:
        """
        Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both. Valid values: `any`, `query`, `trap`.
        """
        return pulumi.get(self, "host_type")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        Host6 entry ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip(self) -> Optional[str]:
        """
        IPv4 address of the SNMP manager (host).
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[str]:
        """
        Source IPv4 address for SNMP traps.
        """
        return pulumi.get(self, "source_ip")


@pulumi.output_type
class CommunityHosts6(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "haDirect":
            suggest = "ha_direct"
        elif key == "hostType":
            suggest = "host_type"
        elif key == "sourceIpv6":
            suggest = "source_ipv6"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CommunityHosts6. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CommunityHosts6.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CommunityHosts6.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 ha_direct: Optional[str] = None,
                 host_type: Optional[str] = None,
                 id: Optional[int] = None,
                 ipv6: Optional[str] = None,
                 source_ipv6: Optional[str] = None):
        """
        :param str ha_direct: Enable/disable direct management of HA cluster members. Valid values: `enable`, `disable`.
        :param str host_type: Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both. Valid values: `any`, `query`, `trap`.
        :param int id: Host6 entry ID.
        :param str ipv6: SNMP manager IPv6 address prefix.
        :param str source_ipv6: Source IPv6 address for SNMP traps.
        """
        if ha_direct is not None:
            pulumi.set(__self__, "ha_direct", ha_direct)
        if host_type is not None:
            pulumi.set(__self__, "host_type", host_type)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if ipv6 is not None:
            pulumi.set(__self__, "ipv6", ipv6)
        if source_ipv6 is not None:
            pulumi.set(__self__, "source_ipv6", source_ipv6)

    @property
    @pulumi.getter(name="haDirect")
    def ha_direct(self) -> Optional[str]:
        """
        Enable/disable direct management of HA cluster members. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ha_direct")

    @property
    @pulumi.getter(name="hostType")
    def host_type(self) -> Optional[str]:
        """
        Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both. Valid values: `any`, `query`, `trap`.
        """
        return pulumi.get(self, "host_type")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        Host6 entry ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ipv6(self) -> Optional[str]:
        """
        SNMP manager IPv6 address prefix.
        """
        return pulumi.get(self, "ipv6")

    @property
    @pulumi.getter(name="sourceIpv6")
    def source_ipv6(self) -> Optional[str]:
        """
        Source IPv6 address for SNMP traps.
        """
        return pulumi.get(self, "source_ipv6")


@pulumi.output_type
class CommunityVdom(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: VDOM name
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        VDOM name
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class UserVdom(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: VDOM name
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        VDOM name
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetCommunityHostResult(dict):
    def __init__(__self__, *,
                 ha_direct: str,
                 host_type: str,
                 id: int,
                 ip: str,
                 source_ip: str):
        """
        :param str ha_direct: Enable/disable direct management of HA cluster members.
        :param str host_type: Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both.
        :param int id: Host6 entry ID.
        :param str ip: IPv4 address of the SNMP manager (host).
        :param str source_ip: Source IPv4 address for SNMP traps.
        """
        pulumi.set(__self__, "ha_direct", ha_direct)
        pulumi.set(__self__, "host_type", host_type)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "source_ip", source_ip)

    @property
    @pulumi.getter(name="haDirect")
    def ha_direct(self) -> str:
        """
        Enable/disable direct management of HA cluster members.
        """
        return pulumi.get(self, "ha_direct")

    @property
    @pulumi.getter(name="hostType")
    def host_type(self) -> str:
        """
        Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both.
        """
        return pulumi.get(self, "host_type")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        Host6 entry ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip(self) -> str:
        """
        IPv4 address of the SNMP manager (host).
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> str:
        """
        Source IPv4 address for SNMP traps.
        """
        return pulumi.get(self, "source_ip")


@pulumi.output_type
class GetCommunityHosts6Result(dict):
    def __init__(__self__, *,
                 ha_direct: str,
                 host_type: str,
                 id: int,
                 ipv6: str,
                 source_ipv6: str):
        """
        :param str ha_direct: Enable/disable direct management of HA cluster members.
        :param str host_type: Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both.
        :param int id: Host6 entry ID.
        :param str ipv6: SNMP manager IPv6 address prefix.
        :param str source_ipv6: Source IPv6 address for SNMP traps.
        """
        pulumi.set(__self__, "ha_direct", ha_direct)
        pulumi.set(__self__, "host_type", host_type)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "ipv6", ipv6)
        pulumi.set(__self__, "source_ipv6", source_ipv6)

    @property
    @pulumi.getter(name="haDirect")
    def ha_direct(self) -> str:
        """
        Enable/disable direct management of HA cluster members.
        """
        return pulumi.get(self, "ha_direct")

    @property
    @pulumi.getter(name="hostType")
    def host_type(self) -> str:
        """
        Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both.
        """
        return pulumi.get(self, "host_type")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        Host6 entry ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ipv6(self) -> str:
        """
        SNMP manager IPv6 address prefix.
        """
        return pulumi.get(self, "ipv6")

    @property
    @pulumi.getter(name="sourceIpv6")
    def source_ipv6(self) -> str:
        """
        Source IPv6 address for SNMP traps.
        """
        return pulumi.get(self, "source_ipv6")


@pulumi.output_type
class GetCommunityVdomResult(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: VDOM name
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        VDOM name
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetUserVdomResult(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Specify the name of the desired systemsnmp user.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specify the name of the desired systemsnmp user.
        """
        return pulumi.get(self, "name")


