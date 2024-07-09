# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['StpArgs', 'Stp']

@pulumi.input_type
class StpArgs:
    def __init__(__self__, *,
                 forward_delay: Optional[pulumi.Input[int]] = None,
                 hello_time: Optional[pulumi.Input[int]] = None,
                 max_age: Optional[pulumi.Input[int]] = None,
                 max_hops: Optional[pulumi.Input[int]] = None,
                 switch_priority: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Stp resource.
        :param pulumi.Input[int] forward_delay: Forward delay (4 - 30 sec, default = 15).
        :param pulumi.Input[int] hello_time: Hello time (1 - 10 sec, default = 2).
        :param pulumi.Input[int] max_age: Maximum packet age (6 - 40 sec, default = 20).
        :param pulumi.Input[int] max_hops: Maximum number of hops (1 - 40, default = 20).
        :param pulumi.Input[str] switch_priority: STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344). Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if forward_delay is not None:
            pulumi.set(__self__, "forward_delay", forward_delay)
        if hello_time is not None:
            pulumi.set(__self__, "hello_time", hello_time)
        if max_age is not None:
            pulumi.set(__self__, "max_age", max_age)
        if max_hops is not None:
            pulumi.set(__self__, "max_hops", max_hops)
        if switch_priority is not None:
            pulumi.set(__self__, "switch_priority", switch_priority)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="forwardDelay")
    def forward_delay(self) -> Optional[pulumi.Input[int]]:
        """
        Forward delay (4 - 30 sec, default = 15).
        """
        return pulumi.get(self, "forward_delay")

    @forward_delay.setter
    def forward_delay(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "forward_delay", value)

    @property
    @pulumi.getter(name="helloTime")
    def hello_time(self) -> Optional[pulumi.Input[int]]:
        """
        Hello time (1 - 10 sec, default = 2).
        """
        return pulumi.get(self, "hello_time")

    @hello_time.setter
    def hello_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hello_time", value)

    @property
    @pulumi.getter(name="maxAge")
    def max_age(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum packet age (6 - 40 sec, default = 20).
        """
        return pulumi.get(self, "max_age")

    @max_age.setter
    def max_age(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_age", value)

    @property
    @pulumi.getter(name="maxHops")
    def max_hops(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of hops (1 - 40, default = 20).
        """
        return pulumi.get(self, "max_hops")

    @max_hops.setter
    def max_hops(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_hops", value)

    @property
    @pulumi.getter(name="switchPriority")
    def switch_priority(self) -> Optional[pulumi.Input[str]]:
        """
        STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344). Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`.
        """
        return pulumi.get(self, "switch_priority")

    @switch_priority.setter
    def switch_priority(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "switch_priority", value)

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
class _StpState:
    def __init__(__self__, *,
                 forward_delay: Optional[pulumi.Input[int]] = None,
                 hello_time: Optional[pulumi.Input[int]] = None,
                 max_age: Optional[pulumi.Input[int]] = None,
                 max_hops: Optional[pulumi.Input[int]] = None,
                 switch_priority: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Stp resources.
        :param pulumi.Input[int] forward_delay: Forward delay (4 - 30 sec, default = 15).
        :param pulumi.Input[int] hello_time: Hello time (1 - 10 sec, default = 2).
        :param pulumi.Input[int] max_age: Maximum packet age (6 - 40 sec, default = 20).
        :param pulumi.Input[int] max_hops: Maximum number of hops (1 - 40, default = 20).
        :param pulumi.Input[str] switch_priority: STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344). Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if forward_delay is not None:
            pulumi.set(__self__, "forward_delay", forward_delay)
        if hello_time is not None:
            pulumi.set(__self__, "hello_time", hello_time)
        if max_age is not None:
            pulumi.set(__self__, "max_age", max_age)
        if max_hops is not None:
            pulumi.set(__self__, "max_hops", max_hops)
        if switch_priority is not None:
            pulumi.set(__self__, "switch_priority", switch_priority)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="forwardDelay")
    def forward_delay(self) -> Optional[pulumi.Input[int]]:
        """
        Forward delay (4 - 30 sec, default = 15).
        """
        return pulumi.get(self, "forward_delay")

    @forward_delay.setter
    def forward_delay(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "forward_delay", value)

    @property
    @pulumi.getter(name="helloTime")
    def hello_time(self) -> Optional[pulumi.Input[int]]:
        """
        Hello time (1 - 10 sec, default = 2).
        """
        return pulumi.get(self, "hello_time")

    @hello_time.setter
    def hello_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hello_time", value)

    @property
    @pulumi.getter(name="maxAge")
    def max_age(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum packet age (6 - 40 sec, default = 20).
        """
        return pulumi.get(self, "max_age")

    @max_age.setter
    def max_age(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_age", value)

    @property
    @pulumi.getter(name="maxHops")
    def max_hops(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of hops (1 - 40, default = 20).
        """
        return pulumi.get(self, "max_hops")

    @max_hops.setter
    def max_hops(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_hops", value)

    @property
    @pulumi.getter(name="switchPriority")
    def switch_priority(self) -> Optional[pulumi.Input[str]]:
        """
        STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344). Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`.
        """
        return pulumi.get(self, "switch_priority")

    @switch_priority.setter
    def switch_priority(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "switch_priority", value)

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


class Stp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 forward_delay: Optional[pulumi.Input[int]] = None,
                 hello_time: Optional[pulumi.Input[int]] = None,
                 max_age: Optional[pulumi.Input[int]] = None,
                 max_hops: Optional[pulumi.Input[int]] = None,
                 switch_priority: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure Spanning Tree Protocol (STP). Applies to FortiOS Version `7.0.4`.

        ## Import

        System Stp can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/stp:Stp labelname SystemStp
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/stp:Stp labelname SystemStp
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] forward_delay: Forward delay (4 - 30 sec, default = 15).
        :param pulumi.Input[int] hello_time: Hello time (1 - 10 sec, default = 2).
        :param pulumi.Input[int] max_age: Maximum packet age (6 - 40 sec, default = 20).
        :param pulumi.Input[int] max_hops: Maximum number of hops (1 - 40, default = 20).
        :param pulumi.Input[str] switch_priority: STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344). Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[StpArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure Spanning Tree Protocol (STP). Applies to FortiOS Version `7.0.4`.

        ## Import

        System Stp can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/stp:Stp labelname SystemStp
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/stp:Stp labelname SystemStp
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param StpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 forward_delay: Optional[pulumi.Input[int]] = None,
                 hello_time: Optional[pulumi.Input[int]] = None,
                 max_age: Optional[pulumi.Input[int]] = None,
                 max_hops: Optional[pulumi.Input[int]] = None,
                 switch_priority: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StpArgs.__new__(StpArgs)

            __props__.__dict__["forward_delay"] = forward_delay
            __props__.__dict__["hello_time"] = hello_time
            __props__.__dict__["max_age"] = max_age
            __props__.__dict__["max_hops"] = max_hops
            __props__.__dict__["switch_priority"] = switch_priority
            __props__.__dict__["vdomparam"] = vdomparam
        super(Stp, __self__).__init__(
            'fortios:system/stp:Stp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            forward_delay: Optional[pulumi.Input[int]] = None,
            hello_time: Optional[pulumi.Input[int]] = None,
            max_age: Optional[pulumi.Input[int]] = None,
            max_hops: Optional[pulumi.Input[int]] = None,
            switch_priority: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Stp':
        """
        Get an existing Stp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] forward_delay: Forward delay (4 - 30 sec, default = 15).
        :param pulumi.Input[int] hello_time: Hello time (1 - 10 sec, default = 2).
        :param pulumi.Input[int] max_age: Maximum packet age (6 - 40 sec, default = 20).
        :param pulumi.Input[int] max_hops: Maximum number of hops (1 - 40, default = 20).
        :param pulumi.Input[str] switch_priority: STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344). Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StpState.__new__(_StpState)

        __props__.__dict__["forward_delay"] = forward_delay
        __props__.__dict__["hello_time"] = hello_time
        __props__.__dict__["max_age"] = max_age
        __props__.__dict__["max_hops"] = max_hops
        __props__.__dict__["switch_priority"] = switch_priority
        __props__.__dict__["vdomparam"] = vdomparam
        return Stp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="forwardDelay")
    def forward_delay(self) -> pulumi.Output[int]:
        """
        Forward delay (4 - 30 sec, default = 15).
        """
        return pulumi.get(self, "forward_delay")

    @property
    @pulumi.getter(name="helloTime")
    def hello_time(self) -> pulumi.Output[int]:
        """
        Hello time (1 - 10 sec, default = 2).
        """
        return pulumi.get(self, "hello_time")

    @property
    @pulumi.getter(name="maxAge")
    def max_age(self) -> pulumi.Output[int]:
        """
        Maximum packet age (6 - 40 sec, default = 20).
        """
        return pulumi.get(self, "max_age")

    @property
    @pulumi.getter(name="maxHops")
    def max_hops(self) -> pulumi.Output[int]:
        """
        Maximum number of hops (1 - 40, default = 20).
        """
        return pulumi.get(self, "max_hops")

    @property
    @pulumi.getter(name="switchPriority")
    def switch_priority(self) -> pulumi.Output[str]:
        """
        STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344). Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`.
        """
        return pulumi.get(self, "switch_priority")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[str]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

