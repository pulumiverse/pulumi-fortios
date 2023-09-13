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
    'ProfileIcapHeader',
    'ProfileRespmodForwardRule',
    'ProfileRespmodForwardRuleHeaderGroup',
    'ProfileRespmodForwardRuleHttpRespStatusCode',
    'ServergroupServerList',
]

@pulumi.output_type
class ProfileIcapHeader(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "base64Encoding":
            suggest = "base64_encoding"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProfileIcapHeader. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProfileIcapHeader.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProfileIcapHeader.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 base64_encoding: Optional[str] = None,
                 content: Optional[str] = None,
                 id: Optional[int] = None,
                 name: Optional[str] = None):
        """
        :param str base64_encoding: Enable/disable use of base64 encoding of HTTP content. Valid values: `disable`, `enable`.
        :param str content: HTTP header content.
        :param int id: HTTP forwarded header ID.
        :param str name: HTTP forwarded header name.
        """
        if base64_encoding is not None:
            pulumi.set(__self__, "base64_encoding", base64_encoding)
        if content is not None:
            pulumi.set(__self__, "content", content)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)

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
    def id(self) -> Optional[int]:
        """
        HTTP forwarded header ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        HTTP forwarded header name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class ProfileRespmodForwardRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "headerGroups":
            suggest = "header_groups"
        elif key == "httpRespStatusCodes":
            suggest = "http_resp_status_codes"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProfileRespmodForwardRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProfileRespmodForwardRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProfileRespmodForwardRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action: Optional[str] = None,
                 header_groups: Optional[Sequence['outputs.ProfileRespmodForwardRuleHeaderGroup']] = None,
                 host: Optional[str] = None,
                 http_resp_status_codes: Optional[Sequence['outputs.ProfileRespmodForwardRuleHttpRespStatusCode']] = None,
                 name: Optional[str] = None):
        """
        :param str action: Action to be taken for ICAP server. Valid values: `forward`, `bypass`.
        :param Sequence['ProfileRespmodForwardRuleHeaderGroupArgs'] header_groups: HTTP header group. The structure of `header_group` block is documented below.
        :param str host: Address object for the host.
        :param Sequence['ProfileRespmodForwardRuleHttpRespStatusCodeArgs'] http_resp_status_codes: HTTP response status code. The structure of `http_resp_status_code` block is documented below.
        :param str name: Address name.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if header_groups is not None:
            pulumi.set(__self__, "header_groups", header_groups)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if http_resp_status_codes is not None:
            pulumi.set(__self__, "http_resp_status_codes", http_resp_status_codes)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def action(self) -> Optional[str]:
        """
        Action to be taken for ICAP server. Valid values: `forward`, `bypass`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="headerGroups")
    def header_groups(self) -> Optional[Sequence['outputs.ProfileRespmodForwardRuleHeaderGroup']]:
        """
        HTTP header group. The structure of `header_group` block is documented below.
        """
        return pulumi.get(self, "header_groups")

    @property
    @pulumi.getter
    def host(self) -> Optional[str]:
        """
        Address object for the host.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="httpRespStatusCodes")
    def http_resp_status_codes(self) -> Optional[Sequence['outputs.ProfileRespmodForwardRuleHttpRespStatusCode']]:
        """
        HTTP response status code. The structure of `http_resp_status_code` block is documented below.
        """
        return pulumi.get(self, "http_resp_status_codes")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Address name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class ProfileRespmodForwardRuleHeaderGroup(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "caseSensitivity":
            suggest = "case_sensitivity"
        elif key == "headerName":
            suggest = "header_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProfileRespmodForwardRuleHeaderGroup. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProfileRespmodForwardRuleHeaderGroup.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProfileRespmodForwardRuleHeaderGroup.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 case_sensitivity: Optional[str] = None,
                 header: Optional[str] = None,
                 header_name: Optional[str] = None,
                 id: Optional[int] = None):
        """
        :param str case_sensitivity: Enable/disable case sensitivity when matching header. Valid values: `disable`, `enable`.
        :param str header: HTTP header regular expression.
        :param str header_name: HTTP header.
        :param int id: ID.
        """
        if case_sensitivity is not None:
            pulumi.set(__self__, "case_sensitivity", case_sensitivity)
        if header is not None:
            pulumi.set(__self__, "header", header)
        if header_name is not None:
            pulumi.set(__self__, "header_name", header_name)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="caseSensitivity")
    def case_sensitivity(self) -> Optional[str]:
        """
        Enable/disable case sensitivity when matching header. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "case_sensitivity")

    @property
    @pulumi.getter
    def header(self) -> Optional[str]:
        """
        HTTP header regular expression.
        """
        return pulumi.get(self, "header")

    @property
    @pulumi.getter(name="headerName")
    def header_name(self) -> Optional[str]:
        """
        HTTP header.
        """
        return pulumi.get(self, "header_name")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        ID.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class ProfileRespmodForwardRuleHttpRespStatusCode(dict):
    def __init__(__self__, *,
                 code: Optional[int] = None):
        """
        :param int code: HTTP response status code.
        """
        if code is not None:
            pulumi.set(__self__, "code", code)

    @property
    @pulumi.getter
    def code(self) -> Optional[int]:
        """
        HTTP response status code.
        """
        return pulumi.get(self, "code")


@pulumi.output_type
class ServergroupServerList(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None,
                 weight: Optional[int] = None):
        """
        :param str name: ICAP server name.
        :param int weight: Optionally assign a weight of the ICAP server for weighted load balancing (1 - 100, default = 10)
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        ICAP server name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def weight(self) -> Optional[int]:
        """
        Optionally assign a weight of the ICAP server for weighted load balancing (1 - 100, default = 10)
        """
        return pulumi.get(self, "weight")


