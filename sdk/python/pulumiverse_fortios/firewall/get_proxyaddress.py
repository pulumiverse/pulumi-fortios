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
    'GetProxyaddressResult',
    'AwaitableGetProxyaddressResult',
    'get_proxyaddress',
    'get_proxyaddress_output',
]

@pulumi.output_type
class GetProxyaddressResult:
    """
    A collection of values returned by getProxyaddress.
    """
    def __init__(__self__, applications=None, case_sensitivity=None, categories=None, color=None, comment=None, header=None, header_groups=None, header_name=None, host=None, host_regex=None, id=None, method=None, name=None, path=None, query=None, referrer=None, taggings=None, type=None, ua=None, ua_max_ver=None, ua_min_ver=None, uuid=None, vdomparam=None, visibility=None):
        if applications and not isinstance(applications, list):
            raise TypeError("Expected argument 'applications' to be a list")
        pulumi.set(__self__, "applications", applications)
        if case_sensitivity and not isinstance(case_sensitivity, str):
            raise TypeError("Expected argument 'case_sensitivity' to be a str")
        pulumi.set(__self__, "case_sensitivity", case_sensitivity)
        if categories and not isinstance(categories, list):
            raise TypeError("Expected argument 'categories' to be a list")
        pulumi.set(__self__, "categories", categories)
        if color and not isinstance(color, int):
            raise TypeError("Expected argument 'color' to be a int")
        pulumi.set(__self__, "color", color)
        if comment and not isinstance(comment, str):
            raise TypeError("Expected argument 'comment' to be a str")
        pulumi.set(__self__, "comment", comment)
        if header and not isinstance(header, str):
            raise TypeError("Expected argument 'header' to be a str")
        pulumi.set(__self__, "header", header)
        if header_groups and not isinstance(header_groups, list):
            raise TypeError("Expected argument 'header_groups' to be a list")
        pulumi.set(__self__, "header_groups", header_groups)
        if header_name and not isinstance(header_name, str):
            raise TypeError("Expected argument 'header_name' to be a str")
        pulumi.set(__self__, "header_name", header_name)
        if host and not isinstance(host, str):
            raise TypeError("Expected argument 'host' to be a str")
        pulumi.set(__self__, "host", host)
        if host_regex and not isinstance(host_regex, str):
            raise TypeError("Expected argument 'host_regex' to be a str")
        pulumi.set(__self__, "host_regex", host_regex)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if method and not isinstance(method, str):
            raise TypeError("Expected argument 'method' to be a str")
        pulumi.set(__self__, "method", method)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if query and not isinstance(query, str):
            raise TypeError("Expected argument 'query' to be a str")
        pulumi.set(__self__, "query", query)
        if referrer and not isinstance(referrer, str):
            raise TypeError("Expected argument 'referrer' to be a str")
        pulumi.set(__self__, "referrer", referrer)
        if taggings and not isinstance(taggings, list):
            raise TypeError("Expected argument 'taggings' to be a list")
        pulumi.set(__self__, "taggings", taggings)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if ua and not isinstance(ua, str):
            raise TypeError("Expected argument 'ua' to be a str")
        pulumi.set(__self__, "ua", ua)
        if ua_max_ver and not isinstance(ua_max_ver, str):
            raise TypeError("Expected argument 'ua_max_ver' to be a str")
        pulumi.set(__self__, "ua_max_ver", ua_max_ver)
        if ua_min_ver and not isinstance(ua_min_ver, str):
            raise TypeError("Expected argument 'ua_min_ver' to be a str")
        pulumi.set(__self__, "ua_min_ver", ua_min_ver)
        if uuid and not isinstance(uuid, str):
            raise TypeError("Expected argument 'uuid' to be a str")
        pulumi.set(__self__, "uuid", uuid)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        pulumi.set(__self__, "visibility", visibility)

    @property
    @pulumi.getter
    def applications(self) -> Sequence['outputs.GetProxyaddressApplicationResult']:
        """
        SaaS application. The structure of `application` block is documented below.
        """
        return pulumi.get(self, "applications")

    @property
    @pulumi.getter(name="caseSensitivity")
    def case_sensitivity(self) -> str:
        """
        Case sensitivity in pattern.
        """
        return pulumi.get(self, "case_sensitivity")

    @property
    @pulumi.getter
    def categories(self) -> Sequence['outputs.GetProxyaddressCategoryResult']:
        """
        Tag category.
        """
        return pulumi.get(self, "categories")

    @property
    @pulumi.getter
    def color(self) -> int:
        """
        Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def comment(self) -> str:
        """
        Optional comments.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def header(self) -> str:
        """
        HTTP header regular expression.
        """
        return pulumi.get(self, "header")

    @property
    @pulumi.getter(name="headerGroups")
    def header_groups(self) -> Sequence['outputs.GetProxyaddressHeaderGroupResult']:
        """
        HTTP header group. The structure of `header_group` block is documented below.
        """
        return pulumi.get(self, "header_groups")

    @property
    @pulumi.getter(name="headerName")
    def header_name(self) -> str:
        """
        HTTP header.
        """
        return pulumi.get(self, "header_name")

    @property
    @pulumi.getter
    def host(self) -> str:
        """
        Address object for the host.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="hostRegex")
    def host_regex(self) -> str:
        """
        Host name as a regular expression.
        """
        return pulumi.get(self, "host_regex")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def method(self) -> str:
        """
        HTTP request methods to be used.
        """
        return pulumi.get(self, "method")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        SaaS applicaton name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> str:
        """
        URL path as a regular expression.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def query(self) -> str:
        """
        Match the query part of the URL as a regular expression.
        """
        return pulumi.get(self, "query")

    @property
    @pulumi.getter
    def referrer(self) -> str:
        """
        Enable/disable use of referrer field in the HTTP header to match the address.
        """
        return pulumi.get(self, "referrer")

    @property
    @pulumi.getter
    def taggings(self) -> Sequence['outputs.GetProxyaddressTaggingResult']:
        """
        Config object tagging. The structure of `tagging` block is documented below.
        """
        return pulumi.get(self, "taggings")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Proxy address type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def ua(self) -> str:
        """
        Names of browsers to be used as user agent.
        """
        return pulumi.get(self, "ua")

    @property
    @pulumi.getter(name="uaMaxVer")
    def ua_max_ver(self) -> str:
        """
        Maximum version of the user agent specified in dotted notation. For example, use 120 with the ua field set to "chrome" to require Google Chrome's maximum version must be 120.
        """
        return pulumi.get(self, "ua_max_ver")

    @property
    @pulumi.getter(name="uaMinVer")
    def ua_min_ver(self) -> str:
        """
        Minimum version of the user agent specified in dotted notation. For example, use 90.0.1 with the ua field set to "chrome" to require Google Chrome's minimum version must be 90.0.1.
        """
        return pulumi.get(self, "ua_min_ver")

    @property
    @pulumi.getter
    def uuid(self) -> str:
        """
        Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def visibility(self) -> str:
        """
        Enable/disable visibility of the object in the GUI.
        """
        return pulumi.get(self, "visibility")


class AwaitableGetProxyaddressResult(GetProxyaddressResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProxyaddressResult(
            applications=self.applications,
            case_sensitivity=self.case_sensitivity,
            categories=self.categories,
            color=self.color,
            comment=self.comment,
            header=self.header,
            header_groups=self.header_groups,
            header_name=self.header_name,
            host=self.host,
            host_regex=self.host_regex,
            id=self.id,
            method=self.method,
            name=self.name,
            path=self.path,
            query=self.query,
            referrer=self.referrer,
            taggings=self.taggings,
            type=self.type,
            ua=self.ua,
            ua_max_ver=self.ua_max_ver,
            ua_min_ver=self.ua_min_ver,
            uuid=self.uuid,
            vdomparam=self.vdomparam,
            visibility=self.visibility)


def get_proxyaddress(name: Optional[str] = None,
                     vdomparam: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProxyaddressResult:
    """
    Use this data source to get information on an fortios firewall proxyaddress


    :param str name: Specify the name of the desired firewall proxyaddress.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:firewall/getProxyaddress:getProxyaddress', __args__, opts=opts, typ=GetProxyaddressResult).value

    return AwaitableGetProxyaddressResult(
        applications=pulumi.get(__ret__, 'applications'),
        case_sensitivity=pulumi.get(__ret__, 'case_sensitivity'),
        categories=pulumi.get(__ret__, 'categories'),
        color=pulumi.get(__ret__, 'color'),
        comment=pulumi.get(__ret__, 'comment'),
        header=pulumi.get(__ret__, 'header'),
        header_groups=pulumi.get(__ret__, 'header_groups'),
        header_name=pulumi.get(__ret__, 'header_name'),
        host=pulumi.get(__ret__, 'host'),
        host_regex=pulumi.get(__ret__, 'host_regex'),
        id=pulumi.get(__ret__, 'id'),
        method=pulumi.get(__ret__, 'method'),
        name=pulumi.get(__ret__, 'name'),
        path=pulumi.get(__ret__, 'path'),
        query=pulumi.get(__ret__, 'query'),
        referrer=pulumi.get(__ret__, 'referrer'),
        taggings=pulumi.get(__ret__, 'taggings'),
        type=pulumi.get(__ret__, 'type'),
        ua=pulumi.get(__ret__, 'ua'),
        ua_max_ver=pulumi.get(__ret__, 'ua_max_ver'),
        ua_min_ver=pulumi.get(__ret__, 'ua_min_ver'),
        uuid=pulumi.get(__ret__, 'uuid'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'),
        visibility=pulumi.get(__ret__, 'visibility'))


@_utilities.lift_output_func(get_proxyaddress)
def get_proxyaddress_output(name: Optional[pulumi.Input[str]] = None,
                            vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProxyaddressResult]:
    """
    Use this data source to get information on an fortios firewall proxyaddress


    :param str name: Specify the name of the desired firewall proxyaddress.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
