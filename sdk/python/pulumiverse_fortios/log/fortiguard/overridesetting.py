# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['OverridesettingArgs', 'Overridesetting']

@pulumi.input_type
class OverridesettingArgs:
    def __init__(__self__, *,
                 access_config: Optional[pulumi.Input[str]] = None,
                 max_log_rate: Optional[pulumi.Input[int]] = None,
                 override: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 upload_day: Optional[pulumi.Input[str]] = None,
                 upload_interval: Optional[pulumi.Input[str]] = None,
                 upload_option: Optional[pulumi.Input[str]] = None,
                 upload_time: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Overridesetting resource.
        :param pulumi.Input[str] access_config: Enable/disable FortiCloud access to configuration and data. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] max_log_rate: FortiCloud maximum log rate in MBps (0 = unlimited).
        :param pulumi.Input[str] override: Overriding FortiCloud settings for this VDOM or use global settings. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] priority: Set log transmission priority. Valid values: `default`, `low`.
        :param pulumi.Input[str] status: Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] upload_day: Day of week to roll logs.
        :param pulumi.Input[str] upload_interval: Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
        :param pulumi.Input[str] upload_option: Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
        :param pulumi.Input[str] upload_time: Time of day to roll logs (hh:mm).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if access_config is not None:
            pulumi.set(__self__, "access_config", access_config)
        if max_log_rate is not None:
            pulumi.set(__self__, "max_log_rate", max_log_rate)
        if override is not None:
            pulumi.set(__self__, "override", override)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if upload_day is not None:
            pulumi.set(__self__, "upload_day", upload_day)
        if upload_interval is not None:
            pulumi.set(__self__, "upload_interval", upload_interval)
        if upload_option is not None:
            pulumi.set(__self__, "upload_option", upload_option)
        if upload_time is not None:
            pulumi.set(__self__, "upload_time", upload_time)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="accessConfig")
    def access_config(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FortiCloud access to configuration and data. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "access_config")

    @access_config.setter
    def access_config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_config", value)

    @property
    @pulumi.getter(name="maxLogRate")
    def max_log_rate(self) -> Optional[pulumi.Input[int]]:
        """
        FortiCloud maximum log rate in MBps (0 = unlimited).
        """
        return pulumi.get(self, "max_log_rate")

    @max_log_rate.setter
    def max_log_rate(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_log_rate", value)

    @property
    @pulumi.getter
    def override(self) -> Optional[pulumi.Input[str]]:
        """
        Overriding FortiCloud settings for this VDOM or use global settings. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "override")

    @override.setter
    def override(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "override", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[str]]:
        """
        Set log transmission priority. Valid values: `default`, `low`.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="uploadDay")
    def upload_day(self) -> Optional[pulumi.Input[str]]:
        """
        Day of week to roll logs.
        """
        return pulumi.get(self, "upload_day")

    @upload_day.setter
    def upload_day(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "upload_day", value)

    @property
    @pulumi.getter(name="uploadInterval")
    def upload_interval(self) -> Optional[pulumi.Input[str]]:
        """
        Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
        """
        return pulumi.get(self, "upload_interval")

    @upload_interval.setter
    def upload_interval(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "upload_interval", value)

    @property
    @pulumi.getter(name="uploadOption")
    def upload_option(self) -> Optional[pulumi.Input[str]]:
        """
        Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
        """
        return pulumi.get(self, "upload_option")

    @upload_option.setter
    def upload_option(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "upload_option", value)

    @property
    @pulumi.getter(name="uploadTime")
    def upload_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time of day to roll logs (hh:mm).
        """
        return pulumi.get(self, "upload_time")

    @upload_time.setter
    def upload_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "upload_time", value)

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
class _OverridesettingState:
    def __init__(__self__, *,
                 access_config: Optional[pulumi.Input[str]] = None,
                 max_log_rate: Optional[pulumi.Input[int]] = None,
                 override: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 upload_day: Optional[pulumi.Input[str]] = None,
                 upload_interval: Optional[pulumi.Input[str]] = None,
                 upload_option: Optional[pulumi.Input[str]] = None,
                 upload_time: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Overridesetting resources.
        :param pulumi.Input[str] access_config: Enable/disable FortiCloud access to configuration and data. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] max_log_rate: FortiCloud maximum log rate in MBps (0 = unlimited).
        :param pulumi.Input[str] override: Overriding FortiCloud settings for this VDOM or use global settings. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] priority: Set log transmission priority. Valid values: `default`, `low`.
        :param pulumi.Input[str] status: Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] upload_day: Day of week to roll logs.
        :param pulumi.Input[str] upload_interval: Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
        :param pulumi.Input[str] upload_option: Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
        :param pulumi.Input[str] upload_time: Time of day to roll logs (hh:mm).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if access_config is not None:
            pulumi.set(__self__, "access_config", access_config)
        if max_log_rate is not None:
            pulumi.set(__self__, "max_log_rate", max_log_rate)
        if override is not None:
            pulumi.set(__self__, "override", override)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if upload_day is not None:
            pulumi.set(__self__, "upload_day", upload_day)
        if upload_interval is not None:
            pulumi.set(__self__, "upload_interval", upload_interval)
        if upload_option is not None:
            pulumi.set(__self__, "upload_option", upload_option)
        if upload_time is not None:
            pulumi.set(__self__, "upload_time", upload_time)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="accessConfig")
    def access_config(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FortiCloud access to configuration and data. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "access_config")

    @access_config.setter
    def access_config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_config", value)

    @property
    @pulumi.getter(name="maxLogRate")
    def max_log_rate(self) -> Optional[pulumi.Input[int]]:
        """
        FortiCloud maximum log rate in MBps (0 = unlimited).
        """
        return pulumi.get(self, "max_log_rate")

    @max_log_rate.setter
    def max_log_rate(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_log_rate", value)

    @property
    @pulumi.getter
    def override(self) -> Optional[pulumi.Input[str]]:
        """
        Overriding FortiCloud settings for this VDOM or use global settings. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "override")

    @override.setter
    def override(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "override", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[str]]:
        """
        Set log transmission priority. Valid values: `default`, `low`.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="uploadDay")
    def upload_day(self) -> Optional[pulumi.Input[str]]:
        """
        Day of week to roll logs.
        """
        return pulumi.get(self, "upload_day")

    @upload_day.setter
    def upload_day(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "upload_day", value)

    @property
    @pulumi.getter(name="uploadInterval")
    def upload_interval(self) -> Optional[pulumi.Input[str]]:
        """
        Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
        """
        return pulumi.get(self, "upload_interval")

    @upload_interval.setter
    def upload_interval(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "upload_interval", value)

    @property
    @pulumi.getter(name="uploadOption")
    def upload_option(self) -> Optional[pulumi.Input[str]]:
        """
        Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
        """
        return pulumi.get(self, "upload_option")

    @upload_option.setter
    def upload_option(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "upload_option", value)

    @property
    @pulumi.getter(name="uploadTime")
    def upload_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time of day to roll logs (hh:mm).
        """
        return pulumi.get(self, "upload_time")

    @upload_time.setter
    def upload_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "upload_time", value)

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


class Overridesetting(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_config: Optional[pulumi.Input[str]] = None,
                 max_log_rate: Optional[pulumi.Input[int]] = None,
                 override: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 upload_day: Optional[pulumi.Input[str]] = None,
                 upload_interval: Optional[pulumi.Input[str]] = None,
                 upload_option: Optional[pulumi.Input[str]] = None,
                 upload_time: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Override global FortiCloud logging settings for this VDOM.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.log.fortiguard.Overridesetting("trname",
            override="disable",
            status="disable",
            upload_interval="daily",
            upload_option="5-minute",
            upload_time="00:00")
        ```

        ## Import

        LogFortiguard OverrideSetting can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:log/fortiguard/overridesetting:Overridesetting labelname LogFortiguardOverrideSetting
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:log/fortiguard/overridesetting:Overridesetting labelname LogFortiguardOverrideSetting
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_config: Enable/disable FortiCloud access to configuration and data. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] max_log_rate: FortiCloud maximum log rate in MBps (0 = unlimited).
        :param pulumi.Input[str] override: Overriding FortiCloud settings for this VDOM or use global settings. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] priority: Set log transmission priority. Valid values: `default`, `low`.
        :param pulumi.Input[str] status: Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] upload_day: Day of week to roll logs.
        :param pulumi.Input[str] upload_interval: Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
        :param pulumi.Input[str] upload_option: Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
        :param pulumi.Input[str] upload_time: Time of day to roll logs (hh:mm).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[OverridesettingArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Override global FortiCloud logging settings for this VDOM.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.log.fortiguard.Overridesetting("trname",
            override="disable",
            status="disable",
            upload_interval="daily",
            upload_option="5-minute",
            upload_time="00:00")
        ```

        ## Import

        LogFortiguard OverrideSetting can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:log/fortiguard/overridesetting:Overridesetting labelname LogFortiguardOverrideSetting
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:log/fortiguard/overridesetting:Overridesetting labelname LogFortiguardOverrideSetting
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param OverridesettingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OverridesettingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_config: Optional[pulumi.Input[str]] = None,
                 max_log_rate: Optional[pulumi.Input[int]] = None,
                 override: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 upload_day: Optional[pulumi.Input[str]] = None,
                 upload_interval: Optional[pulumi.Input[str]] = None,
                 upload_option: Optional[pulumi.Input[str]] = None,
                 upload_time: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OverridesettingArgs.__new__(OverridesettingArgs)

            __props__.__dict__["access_config"] = access_config
            __props__.__dict__["max_log_rate"] = max_log_rate
            __props__.__dict__["override"] = override
            __props__.__dict__["priority"] = priority
            __props__.__dict__["status"] = status
            __props__.__dict__["upload_day"] = upload_day
            __props__.__dict__["upload_interval"] = upload_interval
            __props__.__dict__["upload_option"] = upload_option
            __props__.__dict__["upload_time"] = upload_time
            __props__.__dict__["vdomparam"] = vdomparam
        super(Overridesetting, __self__).__init__(
            'fortios:log/fortiguard/overridesetting:Overridesetting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_config: Optional[pulumi.Input[str]] = None,
            max_log_rate: Optional[pulumi.Input[int]] = None,
            override: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            upload_day: Optional[pulumi.Input[str]] = None,
            upload_interval: Optional[pulumi.Input[str]] = None,
            upload_option: Optional[pulumi.Input[str]] = None,
            upload_time: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Overridesetting':
        """
        Get an existing Overridesetting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_config: Enable/disable FortiCloud access to configuration and data. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] max_log_rate: FortiCloud maximum log rate in MBps (0 = unlimited).
        :param pulumi.Input[str] override: Overriding FortiCloud settings for this VDOM or use global settings. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] priority: Set log transmission priority. Valid values: `default`, `low`.
        :param pulumi.Input[str] status: Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] upload_day: Day of week to roll logs.
        :param pulumi.Input[str] upload_interval: Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
        :param pulumi.Input[str] upload_option: Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
        :param pulumi.Input[str] upload_time: Time of day to roll logs (hh:mm).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OverridesettingState.__new__(_OverridesettingState)

        __props__.__dict__["access_config"] = access_config
        __props__.__dict__["max_log_rate"] = max_log_rate
        __props__.__dict__["override"] = override
        __props__.__dict__["priority"] = priority
        __props__.__dict__["status"] = status
        __props__.__dict__["upload_day"] = upload_day
        __props__.__dict__["upload_interval"] = upload_interval
        __props__.__dict__["upload_option"] = upload_option
        __props__.__dict__["upload_time"] = upload_time
        __props__.__dict__["vdomparam"] = vdomparam
        return Overridesetting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessConfig")
    def access_config(self) -> pulumi.Output[str]:
        """
        Enable/disable FortiCloud access to configuration and data. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "access_config")

    @property
    @pulumi.getter(name="maxLogRate")
    def max_log_rate(self) -> pulumi.Output[int]:
        """
        FortiCloud maximum log rate in MBps (0 = unlimited).
        """
        return pulumi.get(self, "max_log_rate")

    @property
    @pulumi.getter
    def override(self) -> pulumi.Output[str]:
        """
        Overriding FortiCloud settings for this VDOM or use global settings. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "override")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[str]:
        """
        Set log transmission priority. Valid values: `default`, `low`.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="uploadDay")
    def upload_day(self) -> pulumi.Output[str]:
        """
        Day of week to roll logs.
        """
        return pulumi.get(self, "upload_day")

    @property
    @pulumi.getter(name="uploadInterval")
    def upload_interval(self) -> pulumi.Output[str]:
        """
        Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
        """
        return pulumi.get(self, "upload_interval")

    @property
    @pulumi.getter(name="uploadOption")
    def upload_option(self) -> pulumi.Output[str]:
        """
        Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
        """
        return pulumi.get(self, "upload_option")

    @property
    @pulumi.getter(name="uploadTime")
    def upload_time(self) -> pulumi.Output[str]:
        """
        Time of day to roll logs (hh:mm).
        """
        return pulumi.get(self, "upload_time")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[str]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

