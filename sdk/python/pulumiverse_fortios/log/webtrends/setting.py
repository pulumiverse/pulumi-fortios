# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['SettingArgs', 'Setting']

@pulumi.input_type
class SettingArgs:
    def __init__(__self__, *,
                 server: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Setting resource.
        :param pulumi.Input[str] server: Address of the remote WebTrends server.
        :param pulumi.Input[str] status: Enable/disable logging to WebTrends. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if server is not None:
            pulumi.set(__self__, "server", server)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        """
        Address of the remote WebTrends server.
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable logging to WebTrends. Valid values: `enable`, `disable`.
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
class _SettingState:
    def __init__(__self__, *,
                 server: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Setting resources.
        :param pulumi.Input[str] server: Address of the remote WebTrends server.
        :param pulumi.Input[str] status: Enable/disable logging to WebTrends. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if server is not None:
            pulumi.set(__self__, "server", server)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        """
        Address of the remote WebTrends server.
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable logging to WebTrends. Valid values: `enable`, `disable`.
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


class Setting(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Settings for WebTrends.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.log.webtrends.Setting("trname", status="disable")
        ```

        ## Import

        LogWebtrends Setting can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:log/webtrends/setting:Setting labelname LogWebtrendsSetting
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:log/webtrends/setting:Setting labelname LogWebtrendsSetting
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] server: Address of the remote WebTrends server.
        :param pulumi.Input[str] status: Enable/disable logging to WebTrends. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SettingArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Settings for WebTrends.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.log.webtrends.Setting("trname", status="disable")
        ```

        ## Import

        LogWebtrends Setting can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:log/webtrends/setting:Setting labelname LogWebtrendsSetting
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:log/webtrends/setting:Setting labelname LogWebtrendsSetting
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SettingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SettingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SettingArgs.__new__(SettingArgs)

            __props__.__dict__["server"] = server
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(Setting, __self__).__init__(
            'fortios:log/webtrends/setting:Setting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            server: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Setting':
        """
        Get an existing Setting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] server: Address of the remote WebTrends server.
        :param pulumi.Input[str] status: Enable/disable logging to WebTrends. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SettingState.__new__(_SettingState)

        __props__.__dict__["server"] = server
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return Setting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def server(self) -> pulumi.Output[str]:
        """
        Address of the remote WebTrends server.
        """
        return pulumi.get(self, "server")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable logging to WebTrends. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[str]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

