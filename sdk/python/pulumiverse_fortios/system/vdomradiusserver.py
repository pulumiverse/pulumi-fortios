# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VdomradiusserverArgs', 'Vdomradiusserver']

@pulumi.input_type
class VdomradiusserverArgs:
    def __init__(__self__, *,
                 radius_server_vdom: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Vdomradiusserver resource.
        :param pulumi.Input[str] radius_server_vdom: Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
        :param pulumi.Input[str] name: Name of the VDOM that you are adding the RADIUS server to.
        :param pulumi.Input[str] status: Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "radius_server_vdom", radius_server_vdom)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="radiusServerVdom")
    def radius_server_vdom(self) -> pulumi.Input[str]:
        """
        Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
        """
        return pulumi.get(self, "radius_server_vdom")

    @radius_server_vdom.setter
    def radius_server_vdom(self, value: pulumi.Input[str]):
        pulumi.set(self, "radius_server_vdom", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the VDOM that you are adding the RADIUS server to.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `enable`, `disable`.
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
class _VdomradiusserverState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 radius_server_vdom: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Vdomradiusserver resources.
        :param pulumi.Input[str] name: Name of the VDOM that you are adding the RADIUS server to.
        :param pulumi.Input[str] radius_server_vdom: Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
        :param pulumi.Input[str] status: Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if radius_server_vdom is not None:
            pulumi.set(__self__, "radius_server_vdom", radius_server_vdom)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the VDOM that you are adding the RADIUS server to.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="radiusServerVdom")
    def radius_server_vdom(self) -> Optional[pulumi.Input[str]]:
        """
        Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
        """
        return pulumi.get(self, "radius_server_vdom")

    @radius_server_vdom.setter
    def radius_server_vdom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "radius_server_vdom", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `enable`, `disable`.
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


class Vdomradiusserver(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 radius_server_vdom: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure a RADIUS server to use as a RADIUS Single Sign On (RSSO) server for this VDOM.

        ## Import

        System VdomRadiusServer can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/vdomradiusserver:Vdomradiusserver labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/vdomradiusserver:Vdomradiusserver labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the VDOM that you are adding the RADIUS server to.
        :param pulumi.Input[str] radius_server_vdom: Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
        :param pulumi.Input[str] status: Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VdomradiusserverArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure a RADIUS server to use as a RADIUS Single Sign On (RSSO) server for this VDOM.

        ## Import

        System VdomRadiusServer can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/vdomradiusserver:Vdomradiusserver labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/vdomradiusserver:Vdomradiusserver labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param VdomradiusserverArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VdomradiusserverArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 radius_server_vdom: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VdomradiusserverArgs.__new__(VdomradiusserverArgs)

            __props__.__dict__["name"] = name
            if radius_server_vdom is None and not opts.urn:
                raise TypeError("Missing required property 'radius_server_vdom'")
            __props__.__dict__["radius_server_vdom"] = radius_server_vdom
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(Vdomradiusserver, __self__).__init__(
            'fortios:system/vdomradiusserver:Vdomradiusserver',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            radius_server_vdom: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Vdomradiusserver':
        """
        Get an existing Vdomradiusserver resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the VDOM that you are adding the RADIUS server to.
        :param pulumi.Input[str] radius_server_vdom: Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
        :param pulumi.Input[str] status: Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VdomradiusserverState.__new__(_VdomradiusserverState)

        __props__.__dict__["name"] = name
        __props__.__dict__["radius_server_vdom"] = radius_server_vdom
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return Vdomradiusserver(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the VDOM that you are adding the RADIUS server to.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="radiusServerVdom")
    def radius_server_vdom(self) -> pulumi.Output[str]:
        """
        Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
        """
        return pulumi.get(self, "radius_server_vdom")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

