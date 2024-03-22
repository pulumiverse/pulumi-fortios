# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetSysinfoResult',
    'AwaitableGetSysinfoResult',
    'get_sysinfo',
    'get_sysinfo_output',
]

@pulumi.output_type
class GetSysinfoResult:
    """
    A collection of values returned by getSysinfo.
    """
    def __init__(__self__, contact_info=None, description=None, engine_id=None, engine_id_type=None, id=None, location=None, status=None, trap_free_memory_threshold=None, trap_freeable_memory_threshold=None, trap_high_cpu_threshold=None, trap_log_full_threshold=None, trap_low_memory_threshold=None, vdomparam=None):
        if contact_info and not isinstance(contact_info, str):
            raise TypeError("Expected argument 'contact_info' to be a str")
        pulumi.set(__self__, "contact_info", contact_info)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if engine_id and not isinstance(engine_id, str):
            raise TypeError("Expected argument 'engine_id' to be a str")
        pulumi.set(__self__, "engine_id", engine_id)
        if engine_id_type and not isinstance(engine_id_type, str):
            raise TypeError("Expected argument 'engine_id_type' to be a str")
        pulumi.set(__self__, "engine_id_type", engine_id_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if trap_free_memory_threshold and not isinstance(trap_free_memory_threshold, int):
            raise TypeError("Expected argument 'trap_free_memory_threshold' to be a int")
        pulumi.set(__self__, "trap_free_memory_threshold", trap_free_memory_threshold)
        if trap_freeable_memory_threshold and not isinstance(trap_freeable_memory_threshold, int):
            raise TypeError("Expected argument 'trap_freeable_memory_threshold' to be a int")
        pulumi.set(__self__, "trap_freeable_memory_threshold", trap_freeable_memory_threshold)
        if trap_high_cpu_threshold and not isinstance(trap_high_cpu_threshold, int):
            raise TypeError("Expected argument 'trap_high_cpu_threshold' to be a int")
        pulumi.set(__self__, "trap_high_cpu_threshold", trap_high_cpu_threshold)
        if trap_log_full_threshold and not isinstance(trap_log_full_threshold, int):
            raise TypeError("Expected argument 'trap_log_full_threshold' to be a int")
        pulumi.set(__self__, "trap_log_full_threshold", trap_log_full_threshold)
        if trap_low_memory_threshold and not isinstance(trap_low_memory_threshold, int):
            raise TypeError("Expected argument 'trap_low_memory_threshold' to be a int")
        pulumi.set(__self__, "trap_low_memory_threshold", trap_low_memory_threshold)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="contactInfo")
    def contact_info(self) -> str:
        """
        Contact information.
        """
        return pulumi.get(self, "contact_info")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        System description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="engineId")
    def engine_id(self) -> str:
        """
        Local SNMP engineID string (maximum 24 characters).
        """
        return pulumi.get(self, "engine_id")

    @property
    @pulumi.getter(name="engineIdType")
    def engine_id_type(self) -> str:
        """
        Local SNMP engineID type (text/hex/mac).
        """
        return pulumi.get(self, "engine_id_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        System location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Enable/disable SNMP.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="trapFreeMemoryThreshold")
    def trap_free_memory_threshold(self) -> int:
        """
        Free memory usage when trap is sent.
        """
        return pulumi.get(self, "trap_free_memory_threshold")

    @property
    @pulumi.getter(name="trapFreeableMemoryThreshold")
    def trap_freeable_memory_threshold(self) -> int:
        """
        Freeable memory usage when trap is sent.
        """
        return pulumi.get(self, "trap_freeable_memory_threshold")

    @property
    @pulumi.getter(name="trapHighCpuThreshold")
    def trap_high_cpu_threshold(self) -> int:
        """
        CPU usage when trap is sent.
        """
        return pulumi.get(self, "trap_high_cpu_threshold")

    @property
    @pulumi.getter(name="trapLogFullThreshold")
    def trap_log_full_threshold(self) -> int:
        """
        Log disk usage when trap is sent.
        """
        return pulumi.get(self, "trap_log_full_threshold")

    @property
    @pulumi.getter(name="trapLowMemoryThreshold")
    def trap_low_memory_threshold(self) -> int:
        """
        Memory usage when trap is sent.
        """
        return pulumi.get(self, "trap_low_memory_threshold")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSysinfoResult(GetSysinfoResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSysinfoResult(
            contact_info=self.contact_info,
            description=self.description,
            engine_id=self.engine_id,
            engine_id_type=self.engine_id_type,
            id=self.id,
            location=self.location,
            status=self.status,
            trap_free_memory_threshold=self.trap_free_memory_threshold,
            trap_freeable_memory_threshold=self.trap_freeable_memory_threshold,
            trap_high_cpu_threshold=self.trap_high_cpu_threshold,
            trap_log_full_threshold=self.trap_log_full_threshold,
            trap_low_memory_threshold=self.trap_low_memory_threshold,
            vdomparam=self.vdomparam)


def get_sysinfo(vdomparam: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSysinfoResult:
    """
    Use this data source to get information on fortios systemsnmp sysinfo


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:system/snmp/getSysinfo:getSysinfo', __args__, opts=opts, typ=GetSysinfoResult).value

    return AwaitableGetSysinfoResult(
        contact_info=pulumi.get(__ret__, 'contact_info'),
        description=pulumi.get(__ret__, 'description'),
        engine_id=pulumi.get(__ret__, 'engine_id'),
        engine_id_type=pulumi.get(__ret__, 'engine_id_type'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        status=pulumi.get(__ret__, 'status'),
        trap_free_memory_threshold=pulumi.get(__ret__, 'trap_free_memory_threshold'),
        trap_freeable_memory_threshold=pulumi.get(__ret__, 'trap_freeable_memory_threshold'),
        trap_high_cpu_threshold=pulumi.get(__ret__, 'trap_high_cpu_threshold'),
        trap_log_full_threshold=pulumi.get(__ret__, 'trap_log_full_threshold'),
        trap_low_memory_threshold=pulumi.get(__ret__, 'trap_low_memory_threshold'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_sysinfo)
def get_sysinfo_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSysinfoResult]:
    """
    Use this data source to get information on fortios systemsnmp sysinfo


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
