# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['WebproxyDebugurlArgs', 'WebproxyDebugurl']

@pulumi.input_type
class WebproxyDebugurlArgs:
    def __init__(__self__, *,
                 url_pattern: pulumi.Input[str],
                 exact: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WebproxyDebugurl resource.
        :param pulumi.Input[str] url_pattern: URL exemption pattern.
        :param pulumi.Input[str] exact: Enable/disable matching the exact path. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: Debug URL name.
        :param pulumi.Input[str] status: Enable/disable this URL exemption. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "url_pattern", url_pattern)
        if exact is not None:
            pulumi.set(__self__, "exact", exact)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="urlPattern")
    def url_pattern(self) -> pulumi.Input[str]:
        """
        URL exemption pattern.
        """
        return pulumi.get(self, "url_pattern")

    @url_pattern.setter
    def url_pattern(self, value: pulumi.Input[str]):
        pulumi.set(self, "url_pattern", value)

    @property
    @pulumi.getter
    def exact(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable matching the exact path. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "exact")

    @exact.setter
    def exact(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exact", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Debug URL name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable this URL exemption. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

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
class _WebproxyDebugurlState:
    def __init__(__self__, *,
                 exact: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 url_pattern: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WebproxyDebugurl resources.
        :param pulumi.Input[str] exact: Enable/disable matching the exact path. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: Debug URL name.
        :param pulumi.Input[str] status: Enable/disable this URL exemption. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] url_pattern: URL exemption pattern.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if exact is not None:
            pulumi.set(__self__, "exact", exact)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if url_pattern is not None:
            pulumi.set(__self__, "url_pattern", url_pattern)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def exact(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable matching the exact path. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "exact")

    @exact.setter
    def exact(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exact", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Debug URL name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable this URL exemption. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="urlPattern")
    def url_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        URL exemption pattern.
        """
        return pulumi.get(self, "url_pattern")

    @url_pattern.setter
    def url_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url_pattern", value)

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


class WebproxyDebugurl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 exact: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 url_pattern: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure debug URL addresses.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.WebproxyDebugurl("trname",
            exact="enable",
            status="enable",
            url_pattern="/examples/servlet/*Servlet")
        ```

        ## Import

        WebProxy DebugUrl can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/webproxyDebugurl:WebproxyDebugurl labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/webproxyDebugurl:WebproxyDebugurl labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] exact: Enable/disable matching the exact path. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: Debug URL name.
        :param pulumi.Input[str] status: Enable/disable this URL exemption. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] url_pattern: URL exemption pattern.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebproxyDebugurlArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure debug URL addresses.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.WebproxyDebugurl("trname",
            exact="enable",
            status="enable",
            url_pattern="/examples/servlet/*Servlet")
        ```

        ## Import

        WebProxy DebugUrl can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/webproxyDebugurl:WebproxyDebugurl labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/webproxyDebugurl:WebproxyDebugurl labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WebproxyDebugurlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebproxyDebugurlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 exact: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 url_pattern: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebproxyDebugurlArgs.__new__(WebproxyDebugurlArgs)

            __props__.__dict__["exact"] = exact
            __props__.__dict__["name"] = name
            __props__.__dict__["status"] = status
            if url_pattern is None and not opts.urn:
                raise TypeError("Missing required property 'url_pattern'")
            __props__.__dict__["url_pattern"] = url_pattern
            __props__.__dict__["vdomparam"] = vdomparam
        super(WebproxyDebugurl, __self__).__init__(
            'fortios:index/webproxyDebugurl:WebproxyDebugurl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            exact: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            url_pattern: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WebproxyDebugurl':
        """
        Get an existing WebproxyDebugurl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] exact: Enable/disable matching the exact path. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: Debug URL name.
        :param pulumi.Input[str] status: Enable/disable this URL exemption. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] url_pattern: URL exemption pattern.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WebproxyDebugurlState.__new__(_WebproxyDebugurlState)

        __props__.__dict__["exact"] = exact
        __props__.__dict__["name"] = name
        __props__.__dict__["status"] = status
        __props__.__dict__["url_pattern"] = url_pattern
        __props__.__dict__["vdomparam"] = vdomparam
        return WebproxyDebugurl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def exact(self) -> pulumi.Output[str]:
        """
        Enable/disable matching the exact path. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "exact")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Debug URL name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable this URL exemption. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="urlPattern")
    def url_pattern(self) -> pulumi.Output[str]:
        """
        URL exemption pattern.
        """
        return pulumi.get(self, "url_pattern")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

