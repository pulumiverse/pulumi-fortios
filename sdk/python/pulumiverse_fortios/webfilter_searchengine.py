# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['WebfilterSearchengineArgs', 'WebfilterSearchengine']

@pulumi.input_type
class WebfilterSearchengineArgs:
    def __init__(__self__, *,
                 charset: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query: Optional[pulumi.Input[str]] = None,
                 safesearch: Optional[pulumi.Input[str]] = None,
                 safesearch_str: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WebfilterSearchengine resource.
        :param pulumi.Input[str] charset: Search engine charset. Valid values: `utf-8`, `gb2312`.
        :param pulumi.Input[str] hostname: Hostname (regular expression).
        :param pulumi.Input[str] name: Search engine name.
        :param pulumi.Input[str] query: Code used to prefix a query (must end with an equals character).
        :param pulumi.Input[str] safesearch: Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
        :param pulumi.Input[str] safesearch_str: Safe search parameter used in the URL.
        :param pulumi.Input[str] url: URL (regular expression).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if charset is not None:
            pulumi.set(__self__, "charset", charset)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if query is not None:
            pulumi.set(__self__, "query", query)
        if safesearch is not None:
            pulumi.set(__self__, "safesearch", safesearch)
        if safesearch_str is not None:
            pulumi.set(__self__, "safesearch_str", safesearch_str)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def charset(self) -> Optional[pulumi.Input[str]]:
        """
        Search engine charset. Valid values: `utf-8`, `gb2312`.
        """
        return pulumi.get(self, "charset")

    @charset.setter
    def charset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charset", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        Hostname (regular expression).
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Search engine name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def query(self) -> Optional[pulumi.Input[str]]:
        """
        Code used to prefix a query (must end with an equals character).
        """
        return pulumi.get(self, "query")

    @query.setter
    def query(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "query", value)

    @property
    @pulumi.getter
    def safesearch(self) -> Optional[pulumi.Input[str]]:
        """
        Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
        """
        return pulumi.get(self, "safesearch")

    @safesearch.setter
    def safesearch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "safesearch", value)

    @property
    @pulumi.getter(name="safesearchStr")
    def safesearch_str(self) -> Optional[pulumi.Input[str]]:
        """
        Safe search parameter used in the URL.
        """
        return pulumi.get(self, "safesearch_str")

    @safesearch_str.setter
    def safesearch_str(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "safesearch_str", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        URL (regular expression).
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

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
class _WebfilterSearchengineState:
    def __init__(__self__, *,
                 charset: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query: Optional[pulumi.Input[str]] = None,
                 safesearch: Optional[pulumi.Input[str]] = None,
                 safesearch_str: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WebfilterSearchengine resources.
        :param pulumi.Input[str] charset: Search engine charset. Valid values: `utf-8`, `gb2312`.
        :param pulumi.Input[str] hostname: Hostname (regular expression).
        :param pulumi.Input[str] name: Search engine name.
        :param pulumi.Input[str] query: Code used to prefix a query (must end with an equals character).
        :param pulumi.Input[str] safesearch: Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
        :param pulumi.Input[str] safesearch_str: Safe search parameter used in the URL.
        :param pulumi.Input[str] url: URL (regular expression).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if charset is not None:
            pulumi.set(__self__, "charset", charset)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if query is not None:
            pulumi.set(__self__, "query", query)
        if safesearch is not None:
            pulumi.set(__self__, "safesearch", safesearch)
        if safesearch_str is not None:
            pulumi.set(__self__, "safesearch_str", safesearch_str)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def charset(self) -> Optional[pulumi.Input[str]]:
        """
        Search engine charset. Valid values: `utf-8`, `gb2312`.
        """
        return pulumi.get(self, "charset")

    @charset.setter
    def charset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charset", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        Hostname (regular expression).
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Search engine name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def query(self) -> Optional[pulumi.Input[str]]:
        """
        Code used to prefix a query (must end with an equals character).
        """
        return pulumi.get(self, "query")

    @query.setter
    def query(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "query", value)

    @property
    @pulumi.getter
    def safesearch(self) -> Optional[pulumi.Input[str]]:
        """
        Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
        """
        return pulumi.get(self, "safesearch")

    @safesearch.setter
    def safesearch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "safesearch", value)

    @property
    @pulumi.getter(name="safesearchStr")
    def safesearch_str(self) -> Optional[pulumi.Input[str]]:
        """
        Safe search parameter used in the URL.
        """
        return pulumi.get(self, "safesearch_str")

    @safesearch_str.setter
    def safesearch_str(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "safesearch_str", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        URL (regular expression).
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

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


class WebfilterSearchengine(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 charset: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query: Optional[pulumi.Input[str]] = None,
                 safesearch: Optional[pulumi.Input[str]] = None,
                 safesearch_str: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure web filter search engines.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.WebfilterSearchengine("trname",
            charset="utf-8",
            hostname="sg.eiwuc.com",
            query="sc=",
            safesearch="disable",
            url="^\\\\/f")
        ```

        ## Import

        Webfilter SearchEngine can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/webfilterSearchengine:WebfilterSearchengine labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/webfilterSearchengine:WebfilterSearchengine labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] charset: Search engine charset. Valid values: `utf-8`, `gb2312`.
        :param pulumi.Input[str] hostname: Hostname (regular expression).
        :param pulumi.Input[str] name: Search engine name.
        :param pulumi.Input[str] query: Code used to prefix a query (must end with an equals character).
        :param pulumi.Input[str] safesearch: Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
        :param pulumi.Input[str] safesearch_str: Safe search parameter used in the URL.
        :param pulumi.Input[str] url: URL (regular expression).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WebfilterSearchengineArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure web filter search engines.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.WebfilterSearchengine("trname",
            charset="utf-8",
            hostname="sg.eiwuc.com",
            query="sc=",
            safesearch="disable",
            url="^\\\\/f")
        ```

        ## Import

        Webfilter SearchEngine can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/webfilterSearchengine:WebfilterSearchengine labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/webfilterSearchengine:WebfilterSearchengine labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WebfilterSearchengineArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebfilterSearchengineArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 charset: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query: Optional[pulumi.Input[str]] = None,
                 safesearch: Optional[pulumi.Input[str]] = None,
                 safesearch_str: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebfilterSearchengineArgs.__new__(WebfilterSearchengineArgs)

            __props__.__dict__["charset"] = charset
            __props__.__dict__["hostname"] = hostname
            __props__.__dict__["name"] = name
            __props__.__dict__["query"] = query
            __props__.__dict__["safesearch"] = safesearch
            __props__.__dict__["safesearch_str"] = safesearch_str
            __props__.__dict__["url"] = url
            __props__.__dict__["vdomparam"] = vdomparam
        super(WebfilterSearchengine, __self__).__init__(
            'fortios:index/webfilterSearchengine:WebfilterSearchengine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            charset: Optional[pulumi.Input[str]] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            query: Optional[pulumi.Input[str]] = None,
            safesearch: Optional[pulumi.Input[str]] = None,
            safesearch_str: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WebfilterSearchengine':
        """
        Get an existing WebfilterSearchengine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] charset: Search engine charset. Valid values: `utf-8`, `gb2312`.
        :param pulumi.Input[str] hostname: Hostname (regular expression).
        :param pulumi.Input[str] name: Search engine name.
        :param pulumi.Input[str] query: Code used to prefix a query (must end with an equals character).
        :param pulumi.Input[str] safesearch: Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
        :param pulumi.Input[str] safesearch_str: Safe search parameter used in the URL.
        :param pulumi.Input[str] url: URL (regular expression).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WebfilterSearchengineState.__new__(_WebfilterSearchengineState)

        __props__.__dict__["charset"] = charset
        __props__.__dict__["hostname"] = hostname
        __props__.__dict__["name"] = name
        __props__.__dict__["query"] = query
        __props__.__dict__["safesearch"] = safesearch
        __props__.__dict__["safesearch_str"] = safesearch_str
        __props__.__dict__["url"] = url
        __props__.__dict__["vdomparam"] = vdomparam
        return WebfilterSearchengine(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def charset(self) -> pulumi.Output[str]:
        """
        Search engine charset. Valid values: `utf-8`, `gb2312`.
        """
        return pulumi.get(self, "charset")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        """
        Hostname (regular expression).
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Search engine name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def query(self) -> pulumi.Output[str]:
        """
        Code used to prefix a query (must end with an equals character).
        """
        return pulumi.get(self, "query")

    @property
    @pulumi.getter
    def safesearch(self) -> pulumi.Output[str]:
        """
        Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
        """
        return pulumi.get(self, "safesearch")

    @property
    @pulumi.getter(name="safesearchStr")
    def safesearch_str(self) -> pulumi.Output[str]:
        """
        Safe search parameter used in the URL.
        """
        return pulumi.get(self, "safesearch_str")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        URL (regular expression).
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

