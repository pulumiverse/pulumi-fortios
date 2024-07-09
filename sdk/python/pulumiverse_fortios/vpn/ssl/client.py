# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['ClientArgs', 'Client']

@pulumi.input_type
class ClientArgs:
    def __init__(__self__, *,
                 certificate: Optional[pulumi.Input[str]] = None,
                 class_id: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 distance: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 ipv4_subnets: Optional[pulumi.Input[str]] = None,
                 ipv6_subnets: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 psk: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Client resource.
        :param pulumi.Input[str] certificate: Certificate to offer to SSL-VPN server if it requests one.
        :param pulumi.Input[int] class_id: Traffic class ID.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[int] distance: Distance for routes added by SSL-VPN (1 - 255).
        :param pulumi.Input[str] interface: SSL interface to send/receive traffic over.
        :param pulumi.Input[str] ipv4_subnets: IPv4 subnets that the client is protecting.
        :param pulumi.Input[str] ipv6_subnets: IPv6 subnets that the client is protecting.
        :param pulumi.Input[str] name: SSL-VPN tunnel name.
        :param pulumi.Input[str] peer: Authenticate peer's certificate with the peer/peergrp.
        :param pulumi.Input[int] port: SSL-VPN server port.
        :param pulumi.Input[int] priority: Priority for routes added by SSL-VPN. On FortiOS versions 7.0.1-7.0.3: 0 - 4294967295. On FortiOS versions >= 7.0.4: 1 - 65535.
        :param pulumi.Input[str] psk: Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
        :param pulumi.Input[str] realm: Realm name configured on SSL-VPN server.
        :param pulumi.Input[str] server: IPv4, IPv6 or DNS address of the SSL-VPN server.
        :param pulumi.Input[str] source_ip: IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
        :param pulumi.Input[str] status: Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] user: Username to offer to the peer to authenticate the client.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if class_id is not None:
            pulumi.set(__self__, "class_id", class_id)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if distance is not None:
            pulumi.set(__self__, "distance", distance)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if ipv4_subnets is not None:
            pulumi.set(__self__, "ipv4_subnets", ipv4_subnets)
        if ipv6_subnets is not None:
            pulumi.set(__self__, "ipv6_subnets", ipv6_subnets)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if peer is not None:
            pulumi.set(__self__, "peer", peer)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if psk is not None:
            pulumi.set(__self__, "psk", psk)
        if realm is not None:
            pulumi.set(__self__, "realm", realm)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if user is not None:
            pulumi.set(__self__, "user", user)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate to offer to SSL-VPN server if it requests one.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter(name="classId")
    def class_id(self) -> Optional[pulumi.Input[int]]:
        """
        Traffic class ID.
        """
        return pulumi.get(self, "class_id")

    @class_id.setter
    def class_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "class_id", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def distance(self) -> Optional[pulumi.Input[int]]:
        """
        Distance for routes added by SSL-VPN (1 - 255).
        """
        return pulumi.get(self, "distance")

    @distance.setter
    def distance(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "distance", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        SSL interface to send/receive traffic over.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="ipv4Subnets")
    def ipv4_subnets(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 subnets that the client is protecting.
        """
        return pulumi.get(self, "ipv4_subnets")

    @ipv4_subnets.setter
    def ipv4_subnets(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv4_subnets", value)

    @property
    @pulumi.getter(name="ipv6Subnets")
    def ipv6_subnets(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 subnets that the client is protecting.
        """
        return pulumi.get(self, "ipv6_subnets")

    @ipv6_subnets.setter
    def ipv6_subnets(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6_subnets", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        SSL-VPN tunnel name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def peer(self) -> Optional[pulumi.Input[str]]:
        """
        Authenticate peer's certificate with the peer/peergrp.
        """
        return pulumi.get(self, "peer")

    @peer.setter
    def peer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        SSL-VPN server port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        Priority for routes added by SSL-VPN. On FortiOS versions 7.0.1-7.0.3: 0 - 4294967295. On FortiOS versions >= 7.0.4: 1 - 65535.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def psk(self) -> Optional[pulumi.Input[str]]:
        """
        Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
        """
        return pulumi.get(self, "psk")

    @psk.setter
    def psk(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "psk", value)

    @property
    @pulumi.getter
    def realm(self) -> Optional[pulumi.Input[str]]:
        """
        Realm name configured on SSL-VPN server.
        """
        return pulumi.get(self, "realm")

    @realm.setter
    def realm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4, IPv6 or DNS address of the SSL-VPN server.
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        Username to offer to the peer to authenticate the client.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _ClientState:
    def __init__(__self__, *,
                 certificate: Optional[pulumi.Input[str]] = None,
                 class_id: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 distance: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 ipv4_subnets: Optional[pulumi.Input[str]] = None,
                 ipv6_subnets: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 psk: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Client resources.
        :param pulumi.Input[str] certificate: Certificate to offer to SSL-VPN server if it requests one.
        :param pulumi.Input[int] class_id: Traffic class ID.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[int] distance: Distance for routes added by SSL-VPN (1 - 255).
        :param pulumi.Input[str] interface: SSL interface to send/receive traffic over.
        :param pulumi.Input[str] ipv4_subnets: IPv4 subnets that the client is protecting.
        :param pulumi.Input[str] ipv6_subnets: IPv6 subnets that the client is protecting.
        :param pulumi.Input[str] name: SSL-VPN tunnel name.
        :param pulumi.Input[str] peer: Authenticate peer's certificate with the peer/peergrp.
        :param pulumi.Input[int] port: SSL-VPN server port.
        :param pulumi.Input[int] priority: Priority for routes added by SSL-VPN. On FortiOS versions 7.0.1-7.0.3: 0 - 4294967295. On FortiOS versions >= 7.0.4: 1 - 65535.
        :param pulumi.Input[str] psk: Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
        :param pulumi.Input[str] realm: Realm name configured on SSL-VPN server.
        :param pulumi.Input[str] server: IPv4, IPv6 or DNS address of the SSL-VPN server.
        :param pulumi.Input[str] source_ip: IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
        :param pulumi.Input[str] status: Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] user: Username to offer to the peer to authenticate the client.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if class_id is not None:
            pulumi.set(__self__, "class_id", class_id)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if distance is not None:
            pulumi.set(__self__, "distance", distance)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if ipv4_subnets is not None:
            pulumi.set(__self__, "ipv4_subnets", ipv4_subnets)
        if ipv6_subnets is not None:
            pulumi.set(__self__, "ipv6_subnets", ipv6_subnets)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if peer is not None:
            pulumi.set(__self__, "peer", peer)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if psk is not None:
            pulumi.set(__self__, "psk", psk)
        if realm is not None:
            pulumi.set(__self__, "realm", realm)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if user is not None:
            pulumi.set(__self__, "user", user)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate to offer to SSL-VPN server if it requests one.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter(name="classId")
    def class_id(self) -> Optional[pulumi.Input[int]]:
        """
        Traffic class ID.
        """
        return pulumi.get(self, "class_id")

    @class_id.setter
    def class_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "class_id", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def distance(self) -> Optional[pulumi.Input[int]]:
        """
        Distance for routes added by SSL-VPN (1 - 255).
        """
        return pulumi.get(self, "distance")

    @distance.setter
    def distance(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "distance", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        SSL interface to send/receive traffic over.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="ipv4Subnets")
    def ipv4_subnets(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 subnets that the client is protecting.
        """
        return pulumi.get(self, "ipv4_subnets")

    @ipv4_subnets.setter
    def ipv4_subnets(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv4_subnets", value)

    @property
    @pulumi.getter(name="ipv6Subnets")
    def ipv6_subnets(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 subnets that the client is protecting.
        """
        return pulumi.get(self, "ipv6_subnets")

    @ipv6_subnets.setter
    def ipv6_subnets(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6_subnets", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        SSL-VPN tunnel name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def peer(self) -> Optional[pulumi.Input[str]]:
        """
        Authenticate peer's certificate with the peer/peergrp.
        """
        return pulumi.get(self, "peer")

    @peer.setter
    def peer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        SSL-VPN server port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        Priority for routes added by SSL-VPN. On FortiOS versions 7.0.1-7.0.3: 0 - 4294967295. On FortiOS versions >= 7.0.4: 1 - 65535.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def psk(self) -> Optional[pulumi.Input[str]]:
        """
        Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
        """
        return pulumi.get(self, "psk")

    @psk.setter
    def psk(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "psk", value)

    @property
    @pulumi.getter
    def realm(self) -> Optional[pulumi.Input[str]]:
        """
        Realm name configured on SSL-VPN server.
        """
        return pulumi.get(self, "realm")

    @realm.setter
    def realm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4, IPv6 or DNS address of the SSL-VPN server.
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        Username to offer to the peer to authenticate the client.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class Client(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 class_id: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 distance: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 ipv4_subnets: Optional[pulumi.Input[str]] = None,
                 ipv6_subnets: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 psk: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        client Applies to FortiOS Version `>= 7.0.1`.

        ## Import

        VpnSsl Client can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:vpn/ssl/client:Client labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:vpn/ssl/client:Client labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate: Certificate to offer to SSL-VPN server if it requests one.
        :param pulumi.Input[int] class_id: Traffic class ID.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[int] distance: Distance for routes added by SSL-VPN (1 - 255).
        :param pulumi.Input[str] interface: SSL interface to send/receive traffic over.
        :param pulumi.Input[str] ipv4_subnets: IPv4 subnets that the client is protecting.
        :param pulumi.Input[str] ipv6_subnets: IPv6 subnets that the client is protecting.
        :param pulumi.Input[str] name: SSL-VPN tunnel name.
        :param pulumi.Input[str] peer: Authenticate peer's certificate with the peer/peergrp.
        :param pulumi.Input[int] port: SSL-VPN server port.
        :param pulumi.Input[int] priority: Priority for routes added by SSL-VPN. On FortiOS versions 7.0.1-7.0.3: 0 - 4294967295. On FortiOS versions >= 7.0.4: 1 - 65535.
        :param pulumi.Input[str] psk: Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
        :param pulumi.Input[str] realm: Realm name configured on SSL-VPN server.
        :param pulumi.Input[str] server: IPv4, IPv6 or DNS address of the SSL-VPN server.
        :param pulumi.Input[str] source_ip: IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
        :param pulumi.Input[str] status: Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] user: Username to offer to the peer to authenticate the client.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ClientArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        client Applies to FortiOS Version `>= 7.0.1`.

        ## Import

        VpnSsl Client can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:vpn/ssl/client:Client labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:vpn/ssl/client:Client labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param ClientArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClientArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 class_id: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 distance: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 ipv4_subnets: Optional[pulumi.Input[str]] = None,
                 ipv6_subnets: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 psk: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClientArgs.__new__(ClientArgs)

            __props__.__dict__["certificate"] = certificate
            __props__.__dict__["class_id"] = class_id
            __props__.__dict__["comment"] = comment
            __props__.__dict__["distance"] = distance
            __props__.__dict__["interface"] = interface
            __props__.__dict__["ipv4_subnets"] = ipv4_subnets
            __props__.__dict__["ipv6_subnets"] = ipv6_subnets
            __props__.__dict__["name"] = name
            __props__.__dict__["peer"] = peer
            __props__.__dict__["port"] = port
            __props__.__dict__["priority"] = priority
            __props__.__dict__["psk"] = psk
            __props__.__dict__["realm"] = realm
            __props__.__dict__["server"] = server
            __props__.__dict__["source_ip"] = source_ip
            __props__.__dict__["status"] = status
            __props__.__dict__["user"] = user
            __props__.__dict__["vdomparam"] = vdomparam
        super(Client, __self__).__init__(
            'fortios:vpn/ssl/client:Client',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            certificate: Optional[pulumi.Input[str]] = None,
            class_id: Optional[pulumi.Input[int]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            distance: Optional[pulumi.Input[int]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            ipv4_subnets: Optional[pulumi.Input[str]] = None,
            ipv6_subnets: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            peer: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            priority: Optional[pulumi.Input[int]] = None,
            psk: Optional[pulumi.Input[str]] = None,
            realm: Optional[pulumi.Input[str]] = None,
            server: Optional[pulumi.Input[str]] = None,
            source_ip: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            user: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Client':
        """
        Get an existing Client resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate: Certificate to offer to SSL-VPN server if it requests one.
        :param pulumi.Input[int] class_id: Traffic class ID.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[int] distance: Distance for routes added by SSL-VPN (1 - 255).
        :param pulumi.Input[str] interface: SSL interface to send/receive traffic over.
        :param pulumi.Input[str] ipv4_subnets: IPv4 subnets that the client is protecting.
        :param pulumi.Input[str] ipv6_subnets: IPv6 subnets that the client is protecting.
        :param pulumi.Input[str] name: SSL-VPN tunnel name.
        :param pulumi.Input[str] peer: Authenticate peer's certificate with the peer/peergrp.
        :param pulumi.Input[int] port: SSL-VPN server port.
        :param pulumi.Input[int] priority: Priority for routes added by SSL-VPN. On FortiOS versions 7.0.1-7.0.3: 0 - 4294967295. On FortiOS versions >= 7.0.4: 1 - 65535.
        :param pulumi.Input[str] psk: Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
        :param pulumi.Input[str] realm: Realm name configured on SSL-VPN server.
        :param pulumi.Input[str] server: IPv4, IPv6 or DNS address of the SSL-VPN server.
        :param pulumi.Input[str] source_ip: IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
        :param pulumi.Input[str] status: Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] user: Username to offer to the peer to authenticate the client.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClientState.__new__(_ClientState)

        __props__.__dict__["certificate"] = certificate
        __props__.__dict__["class_id"] = class_id
        __props__.__dict__["comment"] = comment
        __props__.__dict__["distance"] = distance
        __props__.__dict__["interface"] = interface
        __props__.__dict__["ipv4_subnets"] = ipv4_subnets
        __props__.__dict__["ipv6_subnets"] = ipv6_subnets
        __props__.__dict__["name"] = name
        __props__.__dict__["peer"] = peer
        __props__.__dict__["port"] = port
        __props__.__dict__["priority"] = priority
        __props__.__dict__["psk"] = psk
        __props__.__dict__["realm"] = realm
        __props__.__dict__["server"] = server
        __props__.__dict__["source_ip"] = source_ip
        __props__.__dict__["status"] = status
        __props__.__dict__["user"] = user
        __props__.__dict__["vdomparam"] = vdomparam
        return Client(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[str]:
        """
        Certificate to offer to SSL-VPN server if it requests one.
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="classId")
    def class_id(self) -> pulumi.Output[int]:
        """
        Traffic class ID.
        """
        return pulumi.get(self, "class_id")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def distance(self) -> pulumi.Output[int]:
        """
        Distance for routes added by SSL-VPN (1 - 255).
        """
        return pulumi.get(self, "distance")

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Output[str]:
        """
        SSL interface to send/receive traffic over.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter(name="ipv4Subnets")
    def ipv4_subnets(self) -> pulumi.Output[str]:
        """
        IPv4 subnets that the client is protecting.
        """
        return pulumi.get(self, "ipv4_subnets")

    @property
    @pulumi.getter(name="ipv6Subnets")
    def ipv6_subnets(self) -> pulumi.Output[str]:
        """
        IPv6 subnets that the client is protecting.
        """
        return pulumi.get(self, "ipv6_subnets")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        SSL-VPN tunnel name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def peer(self) -> pulumi.Output[str]:
        """
        Authenticate peer's certificate with the peer/peergrp.
        """
        return pulumi.get(self, "peer")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        SSL-VPN server port.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[int]:
        """
        Priority for routes added by SSL-VPN. On FortiOS versions 7.0.1-7.0.3: 0 - 4294967295. On FortiOS versions >= 7.0.4: 1 - 65535.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def psk(self) -> pulumi.Output[Optional[str]]:
        """
        Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
        """
        return pulumi.get(self, "psk")

    @property
    @pulumi.getter
    def realm(self) -> pulumi.Output[str]:
        """
        Realm name configured on SSL-VPN server.
        """
        return pulumi.get(self, "realm")

    @property
    @pulumi.getter
    def server(self) -> pulumi.Output[str]:
        """
        IPv4, IPv6 or DNS address of the SSL-VPN server.
        """
        return pulumi.get(self, "server")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> pulumi.Output[str]:
        """
        IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
        """
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[str]:
        """
        Username to offer to the peer to authenticate the client.
        """
        return pulumi.get(self, "user")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[str]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

