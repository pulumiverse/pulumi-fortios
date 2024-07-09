# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['ScheduleArgs', 'Schedule']

@pulumi.input_type
class ScheduleArgs:
    def __init__(__self__, *,
                 frequency: pulumi.Input[str],
                 status: pulumi.Input[str],
                 time: pulumi.Input[str],
                 day: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Schedule resource.
        :param pulumi.Input[str] frequency: Update frequency.
        :param pulumi.Input[str] status: Enable/disable scheduled updates. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] time: Update time.
        :param pulumi.Input[str] day: Update day. Valid values: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "frequency", frequency)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "time", time)
        if day is not None:
            pulumi.set(__self__, "day", day)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def frequency(self) -> pulumi.Input[str]:
        """
        Update frequency.
        """
        return pulumi.get(self, "frequency")

    @frequency.setter
    def frequency(self, value: pulumi.Input[str]):
        pulumi.set(self, "frequency", value)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input[str]:
        """
        Enable/disable scheduled updates. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input[str]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def time(self) -> pulumi.Input[str]:
        """
        Update time.
        """
        return pulumi.get(self, "time")

    @time.setter
    def time(self, value: pulumi.Input[str]):
        pulumi.set(self, "time", value)

    @property
    @pulumi.getter
    def day(self) -> Optional[pulumi.Input[str]]:
        """
        Update day. Valid values: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
        """
        return pulumi.get(self, "day")

    @day.setter
    def day(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "day", value)

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
class _ScheduleState:
    def __init__(__self__, *,
                 day: Optional[pulumi.Input[str]] = None,
                 frequency: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 time: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Schedule resources.
        :param pulumi.Input[str] day: Update day. Valid values: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
        :param pulumi.Input[str] frequency: Update frequency.
        :param pulumi.Input[str] status: Enable/disable scheduled updates. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] time: Update time.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if day is not None:
            pulumi.set(__self__, "day", day)
        if frequency is not None:
            pulumi.set(__self__, "frequency", frequency)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if time is not None:
            pulumi.set(__self__, "time", time)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def day(self) -> Optional[pulumi.Input[str]]:
        """
        Update day. Valid values: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
        """
        return pulumi.get(self, "day")

    @day.setter
    def day(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "day", value)

    @property
    @pulumi.getter
    def frequency(self) -> Optional[pulumi.Input[str]]:
        """
        Update frequency.
        """
        return pulumi.get(self, "frequency")

    @frequency.setter
    def frequency(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "frequency", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable scheduled updates. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def time(self) -> Optional[pulumi.Input[str]]:
        """
        Update time.
        """
        return pulumi.get(self, "time")

    @time.setter
    def time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time", value)

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


class Schedule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 day: Optional[pulumi.Input[str]] = None,
                 frequency: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 time: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure update schedule.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.system.autoupdate.Schedule("trname",
            day="Monday",
            frequency="every",
            status="enable",
            time="02:60")
        ```

        ## Import

        SystemAutoupdate Schedule can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/autoupdate/schedule:Schedule labelname SystemAutoupdateSchedule
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/autoupdate/schedule:Schedule labelname SystemAutoupdateSchedule
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] day: Update day. Valid values: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
        :param pulumi.Input[str] frequency: Update frequency.
        :param pulumi.Input[str] status: Enable/disable scheduled updates. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] time: Update time.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ScheduleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure update schedule.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.system.autoupdate.Schedule("trname",
            day="Monday",
            frequency="every",
            status="enable",
            time="02:60")
        ```

        ## Import

        SystemAutoupdate Schedule can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/autoupdate/schedule:Schedule labelname SystemAutoupdateSchedule
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/autoupdate/schedule:Schedule labelname SystemAutoupdateSchedule
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param ScheduleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ScheduleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 day: Optional[pulumi.Input[str]] = None,
                 frequency: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 time: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ScheduleArgs.__new__(ScheduleArgs)

            __props__.__dict__["day"] = day
            if frequency is None and not opts.urn:
                raise TypeError("Missing required property 'frequency'")
            __props__.__dict__["frequency"] = frequency
            if status is None and not opts.urn:
                raise TypeError("Missing required property 'status'")
            __props__.__dict__["status"] = status
            if time is None and not opts.urn:
                raise TypeError("Missing required property 'time'")
            __props__.__dict__["time"] = time
            __props__.__dict__["vdomparam"] = vdomparam
        super(Schedule, __self__).__init__(
            'fortios:system/autoupdate/schedule:Schedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            day: Optional[pulumi.Input[str]] = None,
            frequency: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            time: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Schedule':
        """
        Get an existing Schedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] day: Update day. Valid values: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
        :param pulumi.Input[str] frequency: Update frequency.
        :param pulumi.Input[str] status: Enable/disable scheduled updates. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] time: Update time.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ScheduleState.__new__(_ScheduleState)

        __props__.__dict__["day"] = day
        __props__.__dict__["frequency"] = frequency
        __props__.__dict__["status"] = status
        __props__.__dict__["time"] = time
        __props__.__dict__["vdomparam"] = vdomparam
        return Schedule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def day(self) -> pulumi.Output[str]:
        """
        Update day. Valid values: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
        """
        return pulumi.get(self, "day")

    @property
    @pulumi.getter
    def frequency(self) -> pulumi.Output[str]:
        """
        Update frequency.
        """
        return pulumi.get(self, "frequency")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable scheduled updates. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def time(self) -> pulumi.Output[str]:
        """
        Update time.
        """
        return pulumi.get(self, "time")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[str]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

