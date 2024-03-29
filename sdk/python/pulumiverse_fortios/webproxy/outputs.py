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
    'ExplicitPacPolicy',
    'ExplicitPacPolicyDstaddr',
    'ExplicitPacPolicySrcaddr6',
    'ExplicitPacPolicySrcaddr',
    'ExplicitSecureWebProxyCert',
    'ForwardservergroupServerList',
    'GlobalLearnClientIpSrcaddr6',
    'GlobalLearnClientIpSrcaddr',
    'ProfileHeader',
    'ProfileHeaderDstaddr6',
    'ProfileHeaderDstaddr',
]

@pulumi.output_type
class ExplicitPacPolicy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "pacFileData":
            suggest = "pac_file_data"
        elif key == "pacFileName":
            suggest = "pac_file_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ExplicitPacPolicy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ExplicitPacPolicy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ExplicitPacPolicy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 comments: Optional[str] = None,
                 dstaddrs: Optional[Sequence['outputs.ExplicitPacPolicyDstaddr']] = None,
                 pac_file_data: Optional[str] = None,
                 pac_file_name: Optional[str] = None,
                 policyid: Optional[int] = None,
                 srcaddr6s: Optional[Sequence['outputs.ExplicitPacPolicySrcaddr6']] = None,
                 srcaddrs: Optional[Sequence['outputs.ExplicitPacPolicySrcaddr']] = None,
                 status: Optional[str] = None):
        """
        :param str comments: Optional comments.
        :param Sequence['ExplicitPacPolicyDstaddrArgs'] dstaddrs: Destination address objects. The structure of `dstaddr` block is documented below.
        :param str pac_file_data: PAC file contents enclosed in quotes (maximum of 256K bytes).
        :param str pac_file_name: Pac file name.
        :param int policyid: Policy ID.
        :param Sequence['ExplicitPacPolicySrcaddr6Args'] srcaddr6s: Source address6 objects. The structure of `srcaddr6` block is documented below.
        :param Sequence['ExplicitPacPolicySrcaddrArgs'] srcaddrs: Source address objects. The structure of `srcaddr` block is documented below.
        :param str status: Enable/disable policy. Valid values: `enable`, `disable`.
        """
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if dstaddrs is not None:
            pulumi.set(__self__, "dstaddrs", dstaddrs)
        if pac_file_data is not None:
            pulumi.set(__self__, "pac_file_data", pac_file_data)
        if pac_file_name is not None:
            pulumi.set(__self__, "pac_file_name", pac_file_name)
        if policyid is not None:
            pulumi.set(__self__, "policyid", policyid)
        if srcaddr6s is not None:
            pulumi.set(__self__, "srcaddr6s", srcaddr6s)
        if srcaddrs is not None:
            pulumi.set(__self__, "srcaddrs", srcaddrs)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def comments(self) -> Optional[str]:
        """
        Optional comments.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def dstaddrs(self) -> Optional[Sequence['outputs.ExplicitPacPolicyDstaddr']]:
        """
        Destination address objects. The structure of `dstaddr` block is documented below.
        """
        return pulumi.get(self, "dstaddrs")

    @property
    @pulumi.getter(name="pacFileData")
    def pac_file_data(self) -> Optional[str]:
        """
        PAC file contents enclosed in quotes (maximum of 256K bytes).
        """
        return pulumi.get(self, "pac_file_data")

    @property
    @pulumi.getter(name="pacFileName")
    def pac_file_name(self) -> Optional[str]:
        """
        Pac file name.
        """
        return pulumi.get(self, "pac_file_name")

    @property
    @pulumi.getter
    def policyid(self) -> Optional[int]:
        """
        Policy ID.
        """
        return pulumi.get(self, "policyid")

    @property
    @pulumi.getter
    def srcaddr6s(self) -> Optional[Sequence['outputs.ExplicitPacPolicySrcaddr6']]:
        """
        Source address6 objects. The structure of `srcaddr6` block is documented below.
        """
        return pulumi.get(self, "srcaddr6s")

    @property
    @pulumi.getter
    def srcaddrs(self) -> Optional[Sequence['outputs.ExplicitPacPolicySrcaddr']]:
        """
        Source address objects. The structure of `srcaddr` block is documented below.
        """
        return pulumi.get(self, "srcaddrs")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Enable/disable policy. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class ExplicitPacPolicyDstaddr(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Address name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Address name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class ExplicitPacPolicySrcaddr6(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Address name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Address name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class ExplicitPacPolicySrcaddr(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Address name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Address name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class ExplicitSecureWebProxyCert(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Certificate list.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Certificate list.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class ForwardservergroupServerList(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None,
                 weight: Optional[int] = None):
        """
        :param str name: Forward server name.
        :param int weight: Optionally assign a weight of the forwarding server for weighted load balancing (1 - 100, default = 10)
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Forward server name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def weight(self) -> Optional[int]:
        """
        Optionally assign a weight of the forwarding server for weighted load balancing (1 - 100, default = 10)
        """
        return pulumi.get(self, "weight")


@pulumi.output_type
class GlobalLearnClientIpSrcaddr6(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Address name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Address name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GlobalLearnClientIpSrcaddr(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Address name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Address name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class ProfileHeader(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "addOption":
            suggest = "add_option"
        elif key == "base64Encoding":
            suggest = "base64_encoding"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProfileHeader. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProfileHeader.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProfileHeader.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action: Optional[str] = None,
                 add_option: Optional[str] = None,
                 base64_encoding: Optional[str] = None,
                 content: Optional[str] = None,
                 dstaddr6s: Optional[Sequence['outputs.ProfileHeaderDstaddr6']] = None,
                 dstaddrs: Optional[Sequence['outputs.ProfileHeaderDstaddr']] = None,
                 id: Optional[int] = None,
                 name: Optional[str] = None,
                 protocol: Optional[str] = None):
        """
        :param str action: Action when the HTTP header is forwarded.
        :param str add_option: Configure options to append content to existing HTTP header or add new HTTP header. Valid values: `append`, `new-on-not-found`, `new`.
        :param str base64_encoding: Enable/disable use of base64 encoding of HTTP content. Valid values: `disable`, `enable`.
        :param str content: HTTP header content.
        :param Sequence['ProfileHeaderDstaddr6Args'] dstaddr6s: Destination address and address group names (IPv6). The structure of `dstaddr6` block is documented below.
        :param Sequence['ProfileHeaderDstaddrArgs'] dstaddrs: Destination address and address group names. The structure of `dstaddr` block is documented below.
        :param int id: HTTP forwarded header id.
        :param str name: HTTP forwarded header name.
        :param str protocol: Configure protocol(s) to take add-option action on (HTTP, HTTPS, or both). Valid values: `https`, `http`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if add_option is not None:
            pulumi.set(__self__, "add_option", add_option)
        if base64_encoding is not None:
            pulumi.set(__self__, "base64_encoding", base64_encoding)
        if content is not None:
            pulumi.set(__self__, "content", content)
        if dstaddr6s is not None:
            pulumi.set(__self__, "dstaddr6s", dstaddr6s)
        if dstaddrs is not None:
            pulumi.set(__self__, "dstaddrs", dstaddrs)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)

    @property
    @pulumi.getter
    def action(self) -> Optional[str]:
        """
        Action when the HTTP header is forwarded.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="addOption")
    def add_option(self) -> Optional[str]:
        """
        Configure options to append content to existing HTTP header or add new HTTP header. Valid values: `append`, `new-on-not-found`, `new`.
        """
        return pulumi.get(self, "add_option")

    @property
    @pulumi.getter(name="base64Encoding")
    def base64_encoding(self) -> Optional[str]:
        """
        Enable/disable use of base64 encoding of HTTP content. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "base64_encoding")

    @property
    @pulumi.getter
    def content(self) -> Optional[str]:
        """
        HTTP header content.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter
    def dstaddr6s(self) -> Optional[Sequence['outputs.ProfileHeaderDstaddr6']]:
        """
        Destination address and address group names (IPv6). The structure of `dstaddr6` block is documented below.
        """
        return pulumi.get(self, "dstaddr6s")

    @property
    @pulumi.getter
    def dstaddrs(self) -> Optional[Sequence['outputs.ProfileHeaderDstaddr']]:
        """
        Destination address and address group names. The structure of `dstaddr` block is documented below.
        """
        return pulumi.get(self, "dstaddrs")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        HTTP forwarded header id.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        HTTP forwarded header name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def protocol(self) -> Optional[str]:
        """
        Configure protocol(s) to take add-option action on (HTTP, HTTPS, or both). Valid values: `https`, `http`.
        """
        return pulumi.get(self, "protocol")


@pulumi.output_type
class ProfileHeaderDstaddr6(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Profile name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Profile name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class ProfileHeaderDstaddr(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Address name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Address name.
        """
        return pulumi.get(self, "name")


