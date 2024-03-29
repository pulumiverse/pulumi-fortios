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
    'KmipserverServerList',
    'OcvpnForticlientAccess',
    'OcvpnForticlientAccessAuthGroup',
    'OcvpnForticlientAccessAuthGroupOverlay',
    'OcvpnOverlay',
    'OcvpnOverlaySubnet',
    'OcvpnWanInterface',
    'QkdCertificate',
]

@pulumi.output_type
class KmipserverServerList(dict):
    def __init__(__self__, *,
                 cert: Optional[str] = None,
                 id: Optional[int] = None,
                 port: Optional[int] = None,
                 server: Optional[str] = None,
                 status: Optional[str] = None):
        """
        :param str cert: Client certificate to use for connectivity to the KMIP server.
        :param int id: ID
        :param int port: KMIP server port.
        :param str server: KMIP server FQDN or IP address.
        :param str status: Enable/disable KMIP server. Valid values: `enable`, `disable`.
        """
        if cert is not None:
            pulumi.set(__self__, "cert", cert)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def cert(self) -> Optional[str]:
        """
        Client certificate to use for connectivity to the KMIP server.
        """
        return pulumi.get(self, "cert")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        ID
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def port(self) -> Optional[int]:
        """
        KMIP server port.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def server(self) -> Optional[str]:
        """
        KMIP server FQDN or IP address.
        """
        return pulumi.get(self, "server")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Enable/disable KMIP server. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class OcvpnForticlientAccess(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "authGroups":
            suggest = "auth_groups"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in OcvpnForticlientAccess. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        OcvpnForticlientAccess.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        OcvpnForticlientAccess.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 auth_groups: Optional[Sequence['outputs.OcvpnForticlientAccessAuthGroup']] = None,
                 psksecret: Optional[str] = None,
                 status: Optional[str] = None):
        """
        :param Sequence['OcvpnForticlientAccessAuthGroupArgs'] auth_groups: FortiClient user authentication groups. The structure of `auth_groups` block is documented below.
        :param str psksecret: Pre-shared secret for FortiClient PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).
        :param str status: Enable/disable FortiClient to access OCVPN networks. Valid values: `enable`, `disable`.
        """
        if auth_groups is not None:
            pulumi.set(__self__, "auth_groups", auth_groups)
        if psksecret is not None:
            pulumi.set(__self__, "psksecret", psksecret)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="authGroups")
    def auth_groups(self) -> Optional[Sequence['outputs.OcvpnForticlientAccessAuthGroup']]:
        """
        FortiClient user authentication groups. The structure of `auth_groups` block is documented below.
        """
        return pulumi.get(self, "auth_groups")

    @property
    @pulumi.getter
    def psksecret(self) -> Optional[str]:
        """
        Pre-shared secret for FortiClient PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).
        """
        return pulumi.get(self, "psksecret")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Enable/disable FortiClient to access OCVPN networks. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class OcvpnForticlientAccessAuthGroup(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "authGroup":
            suggest = "auth_group"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in OcvpnForticlientAccessAuthGroup. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        OcvpnForticlientAccessAuthGroup.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        OcvpnForticlientAccessAuthGroup.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 auth_group: Optional[str] = None,
                 name: Optional[str] = None,
                 overlays: Optional[Sequence['outputs.OcvpnForticlientAccessAuthGroupOverlay']] = None):
        """
        :param str auth_group: Authentication user group for FortiClient access.
        :param str name: Group name.
        :param Sequence['OcvpnForticlientAccessAuthGroupOverlayArgs'] overlays: OCVPN overlays to allow access to. The structure of `overlays` block is documented below.
        """
        if auth_group is not None:
            pulumi.set(__self__, "auth_group", auth_group)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if overlays is not None:
            pulumi.set(__self__, "overlays", overlays)

    @property
    @pulumi.getter(name="authGroup")
    def auth_group(self) -> Optional[str]:
        """
        Authentication user group for FortiClient access.
        """
        return pulumi.get(self, "auth_group")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Group name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def overlays(self) -> Optional[Sequence['outputs.OcvpnForticlientAccessAuthGroupOverlay']]:
        """
        OCVPN overlays to allow access to. The structure of `overlays` block is documented below.
        """
        return pulumi.get(self, "overlays")


@pulumi.output_type
class OcvpnForticlientAccessAuthGroupOverlay(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "overlayName":
            suggest = "overlay_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in OcvpnForticlientAccessAuthGroupOverlay. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        OcvpnForticlientAccessAuthGroupOverlay.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        OcvpnForticlientAccessAuthGroupOverlay.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 overlay_name: Optional[str] = None):
        """
        :param str overlay_name: Overlay name.
        """
        if overlay_name is not None:
            pulumi.set(__self__, "overlay_name", overlay_name)

    @property
    @pulumi.getter(name="overlayName")
    def overlay_name(self) -> Optional[str]:
        """
        Overlay name.
        """
        return pulumi.get(self, "overlay_name")


@pulumi.output_type
class OcvpnOverlay(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "assignIp":
            suggest = "assign_ip"
        elif key == "interOverlay":
            suggest = "inter_overlay"
        elif key == "ipv4EndIp":
            suggest = "ipv4_end_ip"
        elif key == "ipv4StartIp":
            suggest = "ipv4_start_ip"
        elif key == "overlayName":
            suggest = "overlay_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in OcvpnOverlay. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        OcvpnOverlay.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        OcvpnOverlay.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 assign_ip: Optional[str] = None,
                 id: Optional[int] = None,
                 inter_overlay: Optional[str] = None,
                 ipv4_end_ip: Optional[str] = None,
                 ipv4_start_ip: Optional[str] = None,
                 name: Optional[str] = None,
                 overlay_name: Optional[str] = None,
                 subnets: Optional[Sequence['outputs.OcvpnOverlaySubnet']] = None):
        """
        :param str assign_ip: Enable/disable client address assignment. Valid values: `enable`, `disable`.
        :param int id: ID.
        :param str inter_overlay: Allow or deny traffic from other overlays. Valid values: `allow`, `deny`.
        :param str ipv4_end_ip: End of client IPv4 range.
        :param str ipv4_start_ip: Start of client IPv4 range.
        :param str name: Overlay name.
        :param str overlay_name: Overlay name.
        :param Sequence['OcvpnOverlaySubnetArgs'] subnets: Internal subnets to register with OCVPN service. The structure of `subnets` block is documented below.
        """
        if assign_ip is not None:
            pulumi.set(__self__, "assign_ip", assign_ip)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if inter_overlay is not None:
            pulumi.set(__self__, "inter_overlay", inter_overlay)
        if ipv4_end_ip is not None:
            pulumi.set(__self__, "ipv4_end_ip", ipv4_end_ip)
        if ipv4_start_ip is not None:
            pulumi.set(__self__, "ipv4_start_ip", ipv4_start_ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if overlay_name is not None:
            pulumi.set(__self__, "overlay_name", overlay_name)
        if subnets is not None:
            pulumi.set(__self__, "subnets", subnets)

    @property
    @pulumi.getter(name="assignIp")
    def assign_ip(self) -> Optional[str]:
        """
        Enable/disable client address assignment. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "assign_ip")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="interOverlay")
    def inter_overlay(self) -> Optional[str]:
        """
        Allow or deny traffic from other overlays. Valid values: `allow`, `deny`.
        """
        return pulumi.get(self, "inter_overlay")

    @property
    @pulumi.getter(name="ipv4EndIp")
    def ipv4_end_ip(self) -> Optional[str]:
        """
        End of client IPv4 range.
        """
        return pulumi.get(self, "ipv4_end_ip")

    @property
    @pulumi.getter(name="ipv4StartIp")
    def ipv4_start_ip(self) -> Optional[str]:
        """
        Start of client IPv4 range.
        """
        return pulumi.get(self, "ipv4_start_ip")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Overlay name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="overlayName")
    def overlay_name(self) -> Optional[str]:
        """
        Overlay name.
        """
        return pulumi.get(self, "overlay_name")

    @property
    @pulumi.getter
    def subnets(self) -> Optional[Sequence['outputs.OcvpnOverlaySubnet']]:
        """
        Internal subnets to register with OCVPN service. The structure of `subnets` block is documented below.
        """
        return pulumi.get(self, "subnets")


@pulumi.output_type
class OcvpnOverlaySubnet(dict):
    def __init__(__self__, *,
                 id: Optional[int] = None,
                 interface: Optional[str] = None,
                 subnet: Optional[str] = None,
                 type: Optional[str] = None):
        """
        :param int id: ID.
        :param str interface: LAN interface.
        :param str subnet: IPv4 address and subnet mask.
        :param str type: Subnet type. Valid values: `subnet`, `interface`.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if subnet is not None:
            pulumi.set(__self__, "subnet", subnet)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interface(self) -> Optional[str]:
        """
        LAN interface.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def subnet(self) -> Optional[str]:
        """
        IPv4 address and subnet mask.
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Subnet type. Valid values: `subnet`, `interface`.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class OcvpnWanInterface(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Interface name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Interface name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class QkdCertificate(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Certificate name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Certificate name.
        """
        return pulumi.get(self, "name")


