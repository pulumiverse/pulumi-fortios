# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SystemArgs', 'System']

@pulumi.input_type
class SystemArgs:
    def __init__(__self__, *,
                 data_sync_interval: Optional[pulumi.Input[int]] = None,
                 dynamic_periodic_interval: Optional[pulumi.Input[int]] = None,
                 iot_holdoff: Optional[pulumi.Input[int]] = None,
                 iot_mac_idle: Optional[pulumi.Input[int]] = None,
                 iot_scan_interval: Optional[pulumi.Input[int]] = None,
                 iot_weight_threshold: Optional[pulumi.Input[int]] = None,
                 nac_periodic_interval: Optional[pulumi.Input[int]] = None,
                 parallel_process: Optional[pulumi.Input[int]] = None,
                 parallel_process_override: Optional[pulumi.Input[str]] = None,
                 tunnel_mode: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a System resource.
        :param pulumi.Input[int] data_sync_interval: Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
        :param pulumi.Input[int] dynamic_periodic_interval: Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
        :param pulumi.Input[int] iot_holdoff: MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
        :param pulumi.Input[int] iot_mac_idle: MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
        :param pulumi.Input[int] iot_scan_interval: IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
        :param pulumi.Input[int] iot_weight_threshold: MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
        :param pulumi.Input[int] nac_periodic_interval: Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
        :param pulumi.Input[int] parallel_process: Maximum number of parallel processes (1 - 300, default = 1).
        :param pulumi.Input[str] parallel_process_override: Enable/disable parallel process override. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] tunnel_mode: Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if data_sync_interval is not None:
            pulumi.set(__self__, "data_sync_interval", data_sync_interval)
        if dynamic_periodic_interval is not None:
            pulumi.set(__self__, "dynamic_periodic_interval", dynamic_periodic_interval)
        if iot_holdoff is not None:
            pulumi.set(__self__, "iot_holdoff", iot_holdoff)
        if iot_mac_idle is not None:
            pulumi.set(__self__, "iot_mac_idle", iot_mac_idle)
        if iot_scan_interval is not None:
            pulumi.set(__self__, "iot_scan_interval", iot_scan_interval)
        if iot_weight_threshold is not None:
            pulumi.set(__self__, "iot_weight_threshold", iot_weight_threshold)
        if nac_periodic_interval is not None:
            pulumi.set(__self__, "nac_periodic_interval", nac_periodic_interval)
        if parallel_process is not None:
            pulumi.set(__self__, "parallel_process", parallel_process)
        if parallel_process_override is not None:
            pulumi.set(__self__, "parallel_process_override", parallel_process_override)
        if tunnel_mode is not None:
            pulumi.set(__self__, "tunnel_mode", tunnel_mode)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="dataSyncInterval")
    def data_sync_interval(self) -> Optional[pulumi.Input[int]]:
        """
        Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
        """
        return pulumi.get(self, "data_sync_interval")

    @data_sync_interval.setter
    def data_sync_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "data_sync_interval", value)

    @property
    @pulumi.getter(name="dynamicPeriodicInterval")
    def dynamic_periodic_interval(self) -> Optional[pulumi.Input[int]]:
        """
        Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
        """
        return pulumi.get(self, "dynamic_periodic_interval")

    @dynamic_periodic_interval.setter
    def dynamic_periodic_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dynamic_periodic_interval", value)

    @property
    @pulumi.getter(name="iotHoldoff")
    def iot_holdoff(self) -> Optional[pulumi.Input[int]]:
        """
        MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
        """
        return pulumi.get(self, "iot_holdoff")

    @iot_holdoff.setter
    def iot_holdoff(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "iot_holdoff", value)

    @property
    @pulumi.getter(name="iotMacIdle")
    def iot_mac_idle(self) -> Optional[pulumi.Input[int]]:
        """
        MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
        """
        return pulumi.get(self, "iot_mac_idle")

    @iot_mac_idle.setter
    def iot_mac_idle(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "iot_mac_idle", value)

    @property
    @pulumi.getter(name="iotScanInterval")
    def iot_scan_interval(self) -> Optional[pulumi.Input[int]]:
        """
        IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
        """
        return pulumi.get(self, "iot_scan_interval")

    @iot_scan_interval.setter
    def iot_scan_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "iot_scan_interval", value)

    @property
    @pulumi.getter(name="iotWeightThreshold")
    def iot_weight_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
        """
        return pulumi.get(self, "iot_weight_threshold")

    @iot_weight_threshold.setter
    def iot_weight_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "iot_weight_threshold", value)

    @property
    @pulumi.getter(name="nacPeriodicInterval")
    def nac_periodic_interval(self) -> Optional[pulumi.Input[int]]:
        """
        Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
        """
        return pulumi.get(self, "nac_periodic_interval")

    @nac_periodic_interval.setter
    def nac_periodic_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "nac_periodic_interval", value)

    @property
    @pulumi.getter(name="parallelProcess")
    def parallel_process(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of parallel processes (1 - 300, default = 1).
        """
        return pulumi.get(self, "parallel_process")

    @parallel_process.setter
    def parallel_process(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "parallel_process", value)

    @property
    @pulumi.getter(name="parallelProcessOverride")
    def parallel_process_override(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable parallel process override. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "parallel_process_override")

    @parallel_process_override.setter
    def parallel_process_override(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parallel_process_override", value)

    @property
    @pulumi.getter(name="tunnelMode")
    def tunnel_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
        """
        return pulumi.get(self, "tunnel_mode")

    @tunnel_mode.setter
    def tunnel_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tunnel_mode", value)

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
class _SystemState:
    def __init__(__self__, *,
                 data_sync_interval: Optional[pulumi.Input[int]] = None,
                 dynamic_periodic_interval: Optional[pulumi.Input[int]] = None,
                 iot_holdoff: Optional[pulumi.Input[int]] = None,
                 iot_mac_idle: Optional[pulumi.Input[int]] = None,
                 iot_scan_interval: Optional[pulumi.Input[int]] = None,
                 iot_weight_threshold: Optional[pulumi.Input[int]] = None,
                 nac_periodic_interval: Optional[pulumi.Input[int]] = None,
                 parallel_process: Optional[pulumi.Input[int]] = None,
                 parallel_process_override: Optional[pulumi.Input[str]] = None,
                 tunnel_mode: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering System resources.
        :param pulumi.Input[int] data_sync_interval: Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
        :param pulumi.Input[int] dynamic_periodic_interval: Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
        :param pulumi.Input[int] iot_holdoff: MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
        :param pulumi.Input[int] iot_mac_idle: MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
        :param pulumi.Input[int] iot_scan_interval: IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
        :param pulumi.Input[int] iot_weight_threshold: MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
        :param pulumi.Input[int] nac_periodic_interval: Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
        :param pulumi.Input[int] parallel_process: Maximum number of parallel processes (1 - 300, default = 1).
        :param pulumi.Input[str] parallel_process_override: Enable/disable parallel process override. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] tunnel_mode: Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if data_sync_interval is not None:
            pulumi.set(__self__, "data_sync_interval", data_sync_interval)
        if dynamic_periodic_interval is not None:
            pulumi.set(__self__, "dynamic_periodic_interval", dynamic_periodic_interval)
        if iot_holdoff is not None:
            pulumi.set(__self__, "iot_holdoff", iot_holdoff)
        if iot_mac_idle is not None:
            pulumi.set(__self__, "iot_mac_idle", iot_mac_idle)
        if iot_scan_interval is not None:
            pulumi.set(__self__, "iot_scan_interval", iot_scan_interval)
        if iot_weight_threshold is not None:
            pulumi.set(__self__, "iot_weight_threshold", iot_weight_threshold)
        if nac_periodic_interval is not None:
            pulumi.set(__self__, "nac_periodic_interval", nac_periodic_interval)
        if parallel_process is not None:
            pulumi.set(__self__, "parallel_process", parallel_process)
        if parallel_process_override is not None:
            pulumi.set(__self__, "parallel_process_override", parallel_process_override)
        if tunnel_mode is not None:
            pulumi.set(__self__, "tunnel_mode", tunnel_mode)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="dataSyncInterval")
    def data_sync_interval(self) -> Optional[pulumi.Input[int]]:
        """
        Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
        """
        return pulumi.get(self, "data_sync_interval")

    @data_sync_interval.setter
    def data_sync_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "data_sync_interval", value)

    @property
    @pulumi.getter(name="dynamicPeriodicInterval")
    def dynamic_periodic_interval(self) -> Optional[pulumi.Input[int]]:
        """
        Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
        """
        return pulumi.get(self, "dynamic_periodic_interval")

    @dynamic_periodic_interval.setter
    def dynamic_periodic_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dynamic_periodic_interval", value)

    @property
    @pulumi.getter(name="iotHoldoff")
    def iot_holdoff(self) -> Optional[pulumi.Input[int]]:
        """
        MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
        """
        return pulumi.get(self, "iot_holdoff")

    @iot_holdoff.setter
    def iot_holdoff(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "iot_holdoff", value)

    @property
    @pulumi.getter(name="iotMacIdle")
    def iot_mac_idle(self) -> Optional[pulumi.Input[int]]:
        """
        MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
        """
        return pulumi.get(self, "iot_mac_idle")

    @iot_mac_idle.setter
    def iot_mac_idle(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "iot_mac_idle", value)

    @property
    @pulumi.getter(name="iotScanInterval")
    def iot_scan_interval(self) -> Optional[pulumi.Input[int]]:
        """
        IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
        """
        return pulumi.get(self, "iot_scan_interval")

    @iot_scan_interval.setter
    def iot_scan_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "iot_scan_interval", value)

    @property
    @pulumi.getter(name="iotWeightThreshold")
    def iot_weight_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
        """
        return pulumi.get(self, "iot_weight_threshold")

    @iot_weight_threshold.setter
    def iot_weight_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "iot_weight_threshold", value)

    @property
    @pulumi.getter(name="nacPeriodicInterval")
    def nac_periodic_interval(self) -> Optional[pulumi.Input[int]]:
        """
        Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
        """
        return pulumi.get(self, "nac_periodic_interval")

    @nac_periodic_interval.setter
    def nac_periodic_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "nac_periodic_interval", value)

    @property
    @pulumi.getter(name="parallelProcess")
    def parallel_process(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of parallel processes (1 - 300, default = 1).
        """
        return pulumi.get(self, "parallel_process")

    @parallel_process.setter
    def parallel_process(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "parallel_process", value)

    @property
    @pulumi.getter(name="parallelProcessOverride")
    def parallel_process_override(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable parallel process override. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "parallel_process_override")

    @parallel_process_override.setter
    def parallel_process_override(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parallel_process_override", value)

    @property
    @pulumi.getter(name="tunnelMode")
    def tunnel_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
        """
        return pulumi.get(self, "tunnel_mode")

    @tunnel_mode.setter
    def tunnel_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tunnel_mode", value)

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


class System(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_sync_interval: Optional[pulumi.Input[int]] = None,
                 dynamic_periodic_interval: Optional[pulumi.Input[int]] = None,
                 iot_holdoff: Optional[pulumi.Input[int]] = None,
                 iot_mac_idle: Optional[pulumi.Input[int]] = None,
                 iot_scan_interval: Optional[pulumi.Input[int]] = None,
                 iot_weight_threshold: Optional[pulumi.Input[int]] = None,
                 nac_periodic_interval: Optional[pulumi.Input[int]] = None,
                 parallel_process: Optional[pulumi.Input[int]] = None,
                 parallel_process_override: Optional[pulumi.Input[str]] = None,
                 tunnel_mode: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure system-wide switch controller settings.

        ## Import

        SwitchController System can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:switchcontroller/system:System labelname SwitchControllerSystem
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:switchcontroller/system:System labelname SwitchControllerSystem
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] data_sync_interval: Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
        :param pulumi.Input[int] dynamic_periodic_interval: Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
        :param pulumi.Input[int] iot_holdoff: MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
        :param pulumi.Input[int] iot_mac_idle: MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
        :param pulumi.Input[int] iot_scan_interval: IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
        :param pulumi.Input[int] iot_weight_threshold: MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
        :param pulumi.Input[int] nac_periodic_interval: Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
        :param pulumi.Input[int] parallel_process: Maximum number of parallel processes (1 - 300, default = 1).
        :param pulumi.Input[str] parallel_process_override: Enable/disable parallel process override. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] tunnel_mode: Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SystemArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure system-wide switch controller settings.

        ## Import

        SwitchController System can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:switchcontroller/system:System labelname SwitchControllerSystem
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:switchcontroller/system:System labelname SwitchControllerSystem
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_sync_interval: Optional[pulumi.Input[int]] = None,
                 dynamic_periodic_interval: Optional[pulumi.Input[int]] = None,
                 iot_holdoff: Optional[pulumi.Input[int]] = None,
                 iot_mac_idle: Optional[pulumi.Input[int]] = None,
                 iot_scan_interval: Optional[pulumi.Input[int]] = None,
                 iot_weight_threshold: Optional[pulumi.Input[int]] = None,
                 nac_periodic_interval: Optional[pulumi.Input[int]] = None,
                 parallel_process: Optional[pulumi.Input[int]] = None,
                 parallel_process_override: Optional[pulumi.Input[str]] = None,
                 tunnel_mode: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemArgs.__new__(SystemArgs)

            __props__.__dict__["data_sync_interval"] = data_sync_interval
            __props__.__dict__["dynamic_periodic_interval"] = dynamic_periodic_interval
            __props__.__dict__["iot_holdoff"] = iot_holdoff
            __props__.__dict__["iot_mac_idle"] = iot_mac_idle
            __props__.__dict__["iot_scan_interval"] = iot_scan_interval
            __props__.__dict__["iot_weight_threshold"] = iot_weight_threshold
            __props__.__dict__["nac_periodic_interval"] = nac_periodic_interval
            __props__.__dict__["parallel_process"] = parallel_process
            __props__.__dict__["parallel_process_override"] = parallel_process_override
            __props__.__dict__["tunnel_mode"] = tunnel_mode
            __props__.__dict__["vdomparam"] = vdomparam
        super(System, __self__).__init__(
            'fortios:switchcontroller/system:System',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            data_sync_interval: Optional[pulumi.Input[int]] = None,
            dynamic_periodic_interval: Optional[pulumi.Input[int]] = None,
            iot_holdoff: Optional[pulumi.Input[int]] = None,
            iot_mac_idle: Optional[pulumi.Input[int]] = None,
            iot_scan_interval: Optional[pulumi.Input[int]] = None,
            iot_weight_threshold: Optional[pulumi.Input[int]] = None,
            nac_periodic_interval: Optional[pulumi.Input[int]] = None,
            parallel_process: Optional[pulumi.Input[int]] = None,
            parallel_process_override: Optional[pulumi.Input[str]] = None,
            tunnel_mode: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'System':
        """
        Get an existing System resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] data_sync_interval: Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
        :param pulumi.Input[int] dynamic_periodic_interval: Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
        :param pulumi.Input[int] iot_holdoff: MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
        :param pulumi.Input[int] iot_mac_idle: MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
        :param pulumi.Input[int] iot_scan_interval: IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
        :param pulumi.Input[int] iot_weight_threshold: MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
        :param pulumi.Input[int] nac_periodic_interval: Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
        :param pulumi.Input[int] parallel_process: Maximum number of parallel processes (1 - 300, default = 1).
        :param pulumi.Input[str] parallel_process_override: Enable/disable parallel process override. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] tunnel_mode: Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemState.__new__(_SystemState)

        __props__.__dict__["data_sync_interval"] = data_sync_interval
        __props__.__dict__["dynamic_periodic_interval"] = dynamic_periodic_interval
        __props__.__dict__["iot_holdoff"] = iot_holdoff
        __props__.__dict__["iot_mac_idle"] = iot_mac_idle
        __props__.__dict__["iot_scan_interval"] = iot_scan_interval
        __props__.__dict__["iot_weight_threshold"] = iot_weight_threshold
        __props__.__dict__["nac_periodic_interval"] = nac_periodic_interval
        __props__.__dict__["parallel_process"] = parallel_process
        __props__.__dict__["parallel_process_override"] = parallel_process_override
        __props__.__dict__["tunnel_mode"] = tunnel_mode
        __props__.__dict__["vdomparam"] = vdomparam
        return System(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dataSyncInterval")
    def data_sync_interval(self) -> pulumi.Output[int]:
        """
        Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
        """
        return pulumi.get(self, "data_sync_interval")

    @property
    @pulumi.getter(name="dynamicPeriodicInterval")
    def dynamic_periodic_interval(self) -> pulumi.Output[int]:
        """
        Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
        """
        return pulumi.get(self, "dynamic_periodic_interval")

    @property
    @pulumi.getter(name="iotHoldoff")
    def iot_holdoff(self) -> pulumi.Output[int]:
        """
        MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
        """
        return pulumi.get(self, "iot_holdoff")

    @property
    @pulumi.getter(name="iotMacIdle")
    def iot_mac_idle(self) -> pulumi.Output[int]:
        """
        MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
        """
        return pulumi.get(self, "iot_mac_idle")

    @property
    @pulumi.getter(name="iotScanInterval")
    def iot_scan_interval(self) -> pulumi.Output[int]:
        """
        IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
        """
        return pulumi.get(self, "iot_scan_interval")

    @property
    @pulumi.getter(name="iotWeightThreshold")
    def iot_weight_threshold(self) -> pulumi.Output[int]:
        """
        MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
        """
        return pulumi.get(self, "iot_weight_threshold")

    @property
    @pulumi.getter(name="nacPeriodicInterval")
    def nac_periodic_interval(self) -> pulumi.Output[int]:
        """
        Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
        """
        return pulumi.get(self, "nac_periodic_interval")

    @property
    @pulumi.getter(name="parallelProcess")
    def parallel_process(self) -> pulumi.Output[int]:
        """
        Maximum number of parallel processes (1 - 300, default = 1).
        """
        return pulumi.get(self, "parallel_process")

    @property
    @pulumi.getter(name="parallelProcessOverride")
    def parallel_process_override(self) -> pulumi.Output[str]:
        """
        Enable/disable parallel process override. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "parallel_process_override")

    @property
    @pulumi.getter(name="tunnelMode")
    def tunnel_mode(self) -> pulumi.Output[str]:
        """
        Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
        """
        return pulumi.get(self, "tunnel_mode")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

