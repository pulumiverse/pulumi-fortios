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
    'SettingsAuthenticationRuleArgs',
    'SettingsAuthenticationRuleGroupArgs',
    'SettingsAuthenticationRuleSourceAddress6Args',
    'SettingsAuthenticationRuleSourceAddressArgs',
    'SettingsAuthenticationRuleSourceInterfaceArgs',
    'SettingsAuthenticationRuleUserArgs',
    'SettingsSourceAddress6Args',
    'SettingsSourceAddressArgs',
    'SettingsSourceInterfaceArgs',
    'SettingsTunnelIpPoolArgs',
    'SettingsTunnelIpv6PoolArgs',
]

@pulumi.input_type
class SettingsAuthenticationRuleArgs:
    def __init__(__self__, *,
                 auth: Optional[pulumi.Input[str]] = None,
                 cipher: Optional[pulumi.Input[str]] = None,
                 client_cert: Optional[pulumi.Input[str]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleGroupArgs']]]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 portal: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 source_address6_negate: Optional[pulumi.Input[str]] = None,
                 source_address6s: Optional[pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleSourceAddress6Args']]]] = None,
                 source_address_negate: Optional[pulumi.Input[str]] = None,
                 source_addresses: Optional[pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleSourceAddressArgs']]]] = None,
                 source_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleSourceInterfaceArgs']]]] = None,
                 user_peer: Optional[pulumi.Input[str]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleUserArgs']]]] = None):
        """
        :param pulumi.Input[str] auth: SSL VPN authentication method restriction.
        :param pulumi.Input[str] cipher: SSL VPN cipher strength. Valid values: `any`, `high`, `medium`.
        :param pulumi.Input[str] client_cert: Enable/disable SSL VPN client certificate restrictive. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleGroupArgs']]] groups: User groups. The structure of `groups` block is documented below.
        :param pulumi.Input[int] id: ID (0 - 4294967295).
        :param pulumi.Input[str] portal: SSL VPN portal.
        :param pulumi.Input[str] realm: SSL VPN realm.
        :param pulumi.Input[str] source_address6_negate: Enable/disable negated source IPv6 address match. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleSourceAddress6Args']]] source_address6s: IPv6 source address of incoming traffic. The structure of `source_address6` block is documented below.
        :param pulumi.Input[str] source_address_negate: Enable/disable negated source address match. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleSourceAddressArgs']]] source_addresses: Source address of incoming traffic. The structure of `source_address` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleSourceInterfaceArgs']]] source_interfaces: SSL VPN source interface of incoming traffic. The structure of `source_interface` block is documented below.
        :param pulumi.Input[str] user_peer: Name of user peer.
        :param pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleUserArgs']]] users: User name. The structure of `users` block is documented below.
        """
        if auth is not None:
            pulumi.set(__self__, "auth", auth)
        if cipher is not None:
            pulumi.set(__self__, "cipher", cipher)
        if client_cert is not None:
            pulumi.set(__self__, "client_cert", client_cert)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if portal is not None:
            pulumi.set(__self__, "portal", portal)
        if realm is not None:
            pulumi.set(__self__, "realm", realm)
        if source_address6_negate is not None:
            pulumi.set(__self__, "source_address6_negate", source_address6_negate)
        if source_address6s is not None:
            pulumi.set(__self__, "source_address6s", source_address6s)
        if source_address_negate is not None:
            pulumi.set(__self__, "source_address_negate", source_address_negate)
        if source_addresses is not None:
            pulumi.set(__self__, "source_addresses", source_addresses)
        if source_interfaces is not None:
            pulumi.set(__self__, "source_interfaces", source_interfaces)
        if user_peer is not None:
            pulumi.set(__self__, "user_peer", user_peer)
        if users is not None:
            pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter
    def auth(self) -> Optional[pulumi.Input[str]]:
        """
        SSL VPN authentication method restriction.
        """
        return pulumi.get(self, "auth")

    @auth.setter
    def auth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth", value)

    @property
    @pulumi.getter
    def cipher(self) -> Optional[pulumi.Input[str]]:
        """
        SSL VPN cipher strength. Valid values: `any`, `high`, `medium`.
        """
        return pulumi.get(self, "cipher")

    @cipher.setter
    def cipher(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cipher", value)

    @property
    @pulumi.getter(name="clientCert")
    def client_cert(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable SSL VPN client certificate restrictive. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "client_cert")

    @client_cert.setter
    def client_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_cert", value)

    @property
    @pulumi.getter
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleGroupArgs']]]]:
        """
        User groups. The structure of `groups` block is documented below.
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleGroupArgs']]]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        ID (0 - 4294967295).
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def portal(self) -> Optional[pulumi.Input[str]]:
        """
        SSL VPN portal.
        """
        return pulumi.get(self, "portal")

    @portal.setter
    def portal(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "portal", value)

    @property
    @pulumi.getter
    def realm(self) -> Optional[pulumi.Input[str]]:
        """
        SSL VPN realm.
        """
        return pulumi.get(self, "realm")

    @realm.setter
    def realm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm", value)

    @property
    @pulumi.getter(name="sourceAddress6Negate")
    def source_address6_negate(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable negated source IPv6 address match. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "source_address6_negate")

    @source_address6_negate.setter
    def source_address6_negate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_address6_negate", value)

    @property
    @pulumi.getter(name="sourceAddress6s")
    def source_address6s(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleSourceAddress6Args']]]]:
        """
        IPv6 source address of incoming traffic. The structure of `source_address6` block is documented below.
        """
        return pulumi.get(self, "source_address6s")

    @source_address6s.setter
    def source_address6s(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleSourceAddress6Args']]]]):
        pulumi.set(self, "source_address6s", value)

    @property
    @pulumi.getter(name="sourceAddressNegate")
    def source_address_negate(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable negated source address match. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "source_address_negate")

    @source_address_negate.setter
    def source_address_negate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_address_negate", value)

    @property
    @pulumi.getter(name="sourceAddresses")
    def source_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleSourceAddressArgs']]]]:
        """
        Source address of incoming traffic. The structure of `source_address` block is documented below.
        """
        return pulumi.get(self, "source_addresses")

    @source_addresses.setter
    def source_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleSourceAddressArgs']]]]):
        pulumi.set(self, "source_addresses", value)

    @property
    @pulumi.getter(name="sourceInterfaces")
    def source_interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleSourceInterfaceArgs']]]]:
        """
        SSL VPN source interface of incoming traffic. The structure of `source_interface` block is documented below.
        """
        return pulumi.get(self, "source_interfaces")

    @source_interfaces.setter
    def source_interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleSourceInterfaceArgs']]]]):
        pulumi.set(self, "source_interfaces", value)

    @property
    @pulumi.getter(name="userPeer")
    def user_peer(self) -> Optional[pulumi.Input[str]]:
        """
        Name of user peer.
        """
        return pulumi.get(self, "user_peer")

    @user_peer.setter
    def user_peer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_peer", value)

    @property
    @pulumi.getter
    def users(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleUserArgs']]]]:
        """
        User name. The structure of `users` block is documented below.
        """
        return pulumi.get(self, "users")

    @users.setter
    def users(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SettingsAuthenticationRuleUserArgs']]]]):
        pulumi.set(self, "users", value)


@pulumi.input_type
class SettingsAuthenticationRuleGroupArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Group name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Group name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class SettingsAuthenticationRuleSourceAddress6Args:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class SettingsAuthenticationRuleSourceAddressArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: IPv6 address name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 address name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class SettingsAuthenticationRuleSourceInterfaceArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Interface name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Interface name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class SettingsAuthenticationRuleUserArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: User name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        User name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class SettingsSourceAddress6Args:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class SettingsSourceAddressArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: IPv6 address name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 address name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class SettingsSourceInterfaceArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Interface name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Interface name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class SettingsTunnelIpPoolArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Address name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Address name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class SettingsTunnelIpv6PoolArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


