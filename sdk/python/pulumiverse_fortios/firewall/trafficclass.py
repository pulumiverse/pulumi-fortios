# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TrafficclassArgs', 'Trafficclass']

@pulumi.input_type
class TrafficclassArgs:
    def __init__(__self__, *,
                 class_id: Optional[pulumi.Input[int]] = None,
                 class_name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Trafficclass resource.
        :param pulumi.Input[int] class_id: Class ID to be named.
        :param pulumi.Input[str] class_name: Define the name for this class-id.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if class_id is not None:
            pulumi.set(__self__, "class_id", class_id)
        if class_name is not None:
            pulumi.set(__self__, "class_name", class_name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="classId")
    def class_id(self) -> Optional[pulumi.Input[int]]:
        """
        Class ID to be named.
        """
        return pulumi.get(self, "class_id")

    @class_id.setter
    def class_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "class_id", value)

    @property
    @pulumi.getter(name="className")
    def class_name(self) -> Optional[pulumi.Input[str]]:
        """
        Define the name for this class-id.
        """
        return pulumi.get(self, "class_name")

    @class_name.setter
    def class_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "class_name", value)

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
class _TrafficclassState:
    def __init__(__self__, *,
                 class_id: Optional[pulumi.Input[int]] = None,
                 class_name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Trafficclass resources.
        :param pulumi.Input[int] class_id: Class ID to be named.
        :param pulumi.Input[str] class_name: Define the name for this class-id.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if class_id is not None:
            pulumi.set(__self__, "class_id", class_id)
        if class_name is not None:
            pulumi.set(__self__, "class_name", class_name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="classId")
    def class_id(self) -> Optional[pulumi.Input[int]]:
        """
        Class ID to be named.
        """
        return pulumi.get(self, "class_id")

    @class_id.setter
    def class_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "class_id", value)

    @property
    @pulumi.getter(name="className")
    def class_name(self) -> Optional[pulumi.Input[str]]:
        """
        Define the name for this class-id.
        """
        return pulumi.get(self, "class_name")

    @class_name.setter
    def class_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "class_name", value)

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


class Trafficclass(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 class_id: Optional[pulumi.Input[int]] = None,
                 class_name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure names for shaping classes. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        Firewall TrafficClass can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:firewall/trafficclass:Trafficclass labelname {{class_id}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:firewall/trafficclass:Trafficclass labelname {{class_id}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] class_id: Class ID to be named.
        :param pulumi.Input[str] class_name: Define the name for this class-id.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TrafficclassArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure names for shaping classes. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        Firewall TrafficClass can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:firewall/trafficclass:Trafficclass labelname {{class_id}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:firewall/trafficclass:Trafficclass labelname {{class_id}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param TrafficclassArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TrafficclassArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 class_id: Optional[pulumi.Input[int]] = None,
                 class_name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TrafficclassArgs.__new__(TrafficclassArgs)

            __props__.__dict__["class_id"] = class_id
            __props__.__dict__["class_name"] = class_name
            __props__.__dict__["vdomparam"] = vdomparam
        super(Trafficclass, __self__).__init__(
            'fortios:firewall/trafficclass:Trafficclass',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            class_id: Optional[pulumi.Input[int]] = None,
            class_name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Trafficclass':
        """
        Get an existing Trafficclass resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] class_id: Class ID to be named.
        :param pulumi.Input[str] class_name: Define the name for this class-id.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TrafficclassState.__new__(_TrafficclassState)

        __props__.__dict__["class_id"] = class_id
        __props__.__dict__["class_name"] = class_name
        __props__.__dict__["vdomparam"] = vdomparam
        return Trafficclass(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="classId")
    def class_id(self) -> pulumi.Output[int]:
        """
        Class ID to be named.
        """
        return pulumi.get(self, "class_id")

    @property
    @pulumi.getter(name="className")
    def class_name(self) -> pulumi.Output[str]:
        """
        Define the name for this class-id.
        """
        return pulumi.get(self, "class_name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

