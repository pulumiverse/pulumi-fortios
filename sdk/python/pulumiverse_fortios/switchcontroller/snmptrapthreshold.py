# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SnmptrapthresholdArgs', 'Snmptrapthreshold']

@pulumi.input_type
class SnmptrapthresholdArgs:
    def __init__(__self__, *,
                 trap_high_cpu_threshold: Optional[pulumi.Input[int]] = None,
                 trap_log_full_threshold: Optional[pulumi.Input[int]] = None,
                 trap_low_memory_threshold: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Snmptrapthreshold resource.
        :param pulumi.Input[int] trap_high_cpu_threshold: CPU usage when trap is sent.
        :param pulumi.Input[int] trap_log_full_threshold: Log disk usage when trap is sent.
        :param pulumi.Input[int] trap_low_memory_threshold: Memory usage when trap is sent.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if trap_high_cpu_threshold is not None:
            pulumi.set(__self__, "trap_high_cpu_threshold", trap_high_cpu_threshold)
        if trap_log_full_threshold is not None:
            pulumi.set(__self__, "trap_log_full_threshold", trap_log_full_threshold)
        if trap_low_memory_threshold is not None:
            pulumi.set(__self__, "trap_low_memory_threshold", trap_low_memory_threshold)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="trapHighCpuThreshold")
    def trap_high_cpu_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        CPU usage when trap is sent.
        """
        return pulumi.get(self, "trap_high_cpu_threshold")

    @trap_high_cpu_threshold.setter
    def trap_high_cpu_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "trap_high_cpu_threshold", value)

    @property
    @pulumi.getter(name="trapLogFullThreshold")
    def trap_log_full_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Log disk usage when trap is sent.
        """
        return pulumi.get(self, "trap_log_full_threshold")

    @trap_log_full_threshold.setter
    def trap_log_full_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "trap_log_full_threshold", value)

    @property
    @pulumi.getter(name="trapLowMemoryThreshold")
    def trap_low_memory_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Memory usage when trap is sent.
        """
        return pulumi.get(self, "trap_low_memory_threshold")

    @trap_low_memory_threshold.setter
    def trap_low_memory_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "trap_low_memory_threshold", value)

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
class _SnmptrapthresholdState:
    def __init__(__self__, *,
                 trap_high_cpu_threshold: Optional[pulumi.Input[int]] = None,
                 trap_log_full_threshold: Optional[pulumi.Input[int]] = None,
                 trap_low_memory_threshold: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Snmptrapthreshold resources.
        :param pulumi.Input[int] trap_high_cpu_threshold: CPU usage when trap is sent.
        :param pulumi.Input[int] trap_log_full_threshold: Log disk usage when trap is sent.
        :param pulumi.Input[int] trap_low_memory_threshold: Memory usage when trap is sent.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if trap_high_cpu_threshold is not None:
            pulumi.set(__self__, "trap_high_cpu_threshold", trap_high_cpu_threshold)
        if trap_log_full_threshold is not None:
            pulumi.set(__self__, "trap_log_full_threshold", trap_log_full_threshold)
        if trap_low_memory_threshold is not None:
            pulumi.set(__self__, "trap_low_memory_threshold", trap_low_memory_threshold)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="trapHighCpuThreshold")
    def trap_high_cpu_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        CPU usage when trap is sent.
        """
        return pulumi.get(self, "trap_high_cpu_threshold")

    @trap_high_cpu_threshold.setter
    def trap_high_cpu_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "trap_high_cpu_threshold", value)

    @property
    @pulumi.getter(name="trapLogFullThreshold")
    def trap_log_full_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Log disk usage when trap is sent.
        """
        return pulumi.get(self, "trap_log_full_threshold")

    @trap_log_full_threshold.setter
    def trap_log_full_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "trap_log_full_threshold", value)

    @property
    @pulumi.getter(name="trapLowMemoryThreshold")
    def trap_low_memory_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Memory usage when trap is sent.
        """
        return pulumi.get(self, "trap_low_memory_threshold")

    @trap_low_memory_threshold.setter
    def trap_low_memory_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "trap_low_memory_threshold", value)

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


class Snmptrapthreshold(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 trap_high_cpu_threshold: Optional[pulumi.Input[int]] = None,
                 trap_log_full_threshold: Optional[pulumi.Input[int]] = None,
                 trap_low_memory_threshold: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure FortiSwitch SNMP trap threshold values globally. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        SwitchController SnmpTrapThreshold can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:switchcontroller/snmptrapthreshold:Snmptrapthreshold labelname SwitchControllerSnmpTrapThreshold
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:switchcontroller/snmptrapthreshold:Snmptrapthreshold labelname SwitchControllerSnmpTrapThreshold
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] trap_high_cpu_threshold: CPU usage when trap is sent.
        :param pulumi.Input[int] trap_log_full_threshold: Log disk usage when trap is sent.
        :param pulumi.Input[int] trap_low_memory_threshold: Memory usage when trap is sent.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SnmptrapthresholdArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure FortiSwitch SNMP trap threshold values globally. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        SwitchController SnmpTrapThreshold can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:switchcontroller/snmptrapthreshold:Snmptrapthreshold labelname SwitchControllerSnmpTrapThreshold
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:switchcontroller/snmptrapthreshold:Snmptrapthreshold labelname SwitchControllerSnmpTrapThreshold
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SnmptrapthresholdArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnmptrapthresholdArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 trap_high_cpu_threshold: Optional[pulumi.Input[int]] = None,
                 trap_log_full_threshold: Optional[pulumi.Input[int]] = None,
                 trap_low_memory_threshold: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SnmptrapthresholdArgs.__new__(SnmptrapthresholdArgs)

            __props__.__dict__["trap_high_cpu_threshold"] = trap_high_cpu_threshold
            __props__.__dict__["trap_log_full_threshold"] = trap_log_full_threshold
            __props__.__dict__["trap_low_memory_threshold"] = trap_low_memory_threshold
            __props__.__dict__["vdomparam"] = vdomparam
        super(Snmptrapthreshold, __self__).__init__(
            'fortios:switchcontroller/snmptrapthreshold:Snmptrapthreshold',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            trap_high_cpu_threshold: Optional[pulumi.Input[int]] = None,
            trap_log_full_threshold: Optional[pulumi.Input[int]] = None,
            trap_low_memory_threshold: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Snmptrapthreshold':
        """
        Get an existing Snmptrapthreshold resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] trap_high_cpu_threshold: CPU usage when trap is sent.
        :param pulumi.Input[int] trap_log_full_threshold: Log disk usage when trap is sent.
        :param pulumi.Input[int] trap_low_memory_threshold: Memory usage when trap is sent.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnmptrapthresholdState.__new__(_SnmptrapthresholdState)

        __props__.__dict__["trap_high_cpu_threshold"] = trap_high_cpu_threshold
        __props__.__dict__["trap_log_full_threshold"] = trap_log_full_threshold
        __props__.__dict__["trap_low_memory_threshold"] = trap_low_memory_threshold
        __props__.__dict__["vdomparam"] = vdomparam
        return Snmptrapthreshold(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="trapHighCpuThreshold")
    def trap_high_cpu_threshold(self) -> pulumi.Output[int]:
        """
        CPU usage when trap is sent.
        """
        return pulumi.get(self, "trap_high_cpu_threshold")

    @property
    @pulumi.getter(name="trapLogFullThreshold")
    def trap_log_full_threshold(self) -> pulumi.Output[int]:
        """
        Log disk usage when trap is sent.
        """
        return pulumi.get(self, "trap_log_full_threshold")

    @property
    @pulumi.getter(name="trapLowMemoryThreshold")
    def trap_low_memory_threshold(self) -> pulumi.Output[int]:
        """
        Memory usage when trap is sent.
        """
        return pulumi.get(self, "trap_low_memory_threshold")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

