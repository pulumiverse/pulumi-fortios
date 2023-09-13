# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['GlobalArgs', 'Global']

@pulumi.input_type
class GlobalArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 elasticbuffer: Optional[pulumi.Input[str]] = None,
                 hugepage_percentage: Optional[pulumi.Input[int]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['GlobalInterfaceArgs']]]] = None,
                 ipsec_offload: Optional[pulumi.Input[str]] = None,
                 mbufpool_percentage: Optional[pulumi.Input[int]] = None,
                 multiqueue: Optional[pulumi.Input[str]] = None,
                 per_session_accounting: Optional[pulumi.Input[str]] = None,
                 sleep_on_idle: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Global resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] elasticbuffer: Enable/disable elasticbuffer support for all DPDK ports. Valid values: `disable`, `enable`.
        :param pulumi.Input[int] hugepage_percentage: Percentage of main memory allocated to hugepages, which are available for DPDK operation.
        :param pulumi.Input[Sequence[pulumi.Input['GlobalInterfaceArgs']]] interfaces: Physical interfaces that enable DPDK. The structure of `interface` block is documented below.
        :param pulumi.Input[str] ipsec_offload: Enable/disable DPDK IPsec phase 2 offloading. Valid values: `disable`, `enable`.
        :param pulumi.Input[int] mbufpool_percentage: Percentage of main memory allocated to DPDK packet buffer.
        :param pulumi.Input[str] multiqueue: Enable/disable multi-queue RX/TX support for all DPDK ports. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] per_session_accounting: Enable/disable per-session accounting. Valid values: `disable`, `traffic-log-only`, `enable`.
        :param pulumi.Input[str] sleep_on_idle: Enable/disable sleep-on-idle support for all FDH engines. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] status: Enable/disable DPDK operation for the entire system. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if elasticbuffer is not None:
            pulumi.set(__self__, "elasticbuffer", elasticbuffer)
        if hugepage_percentage is not None:
            pulumi.set(__self__, "hugepage_percentage", hugepage_percentage)
        if interfaces is not None:
            pulumi.set(__self__, "interfaces", interfaces)
        if ipsec_offload is not None:
            pulumi.set(__self__, "ipsec_offload", ipsec_offload)
        if mbufpool_percentage is not None:
            pulumi.set(__self__, "mbufpool_percentage", mbufpool_percentage)
        if multiqueue is not None:
            pulumi.set(__self__, "multiqueue", multiqueue)
        if per_session_accounting is not None:
            pulumi.set(__self__, "per_session_accounting", per_session_accounting)
        if sleep_on_idle is not None:
            pulumi.set(__self__, "sleep_on_idle", sleep_on_idle)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def elasticbuffer(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable elasticbuffer support for all DPDK ports. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "elasticbuffer")

    @elasticbuffer.setter
    def elasticbuffer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "elasticbuffer", value)

    @property
    @pulumi.getter(name="hugepagePercentage")
    def hugepage_percentage(self) -> Optional[pulumi.Input[int]]:
        """
        Percentage of main memory allocated to hugepages, which are available for DPDK operation.
        """
        return pulumi.get(self, "hugepage_percentage")

    @hugepage_percentage.setter
    def hugepage_percentage(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hugepage_percentage", value)

    @property
    @pulumi.getter
    def interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GlobalInterfaceArgs']]]]:
        """
        Physical interfaces that enable DPDK. The structure of `interface` block is documented below.
        """
        return pulumi.get(self, "interfaces")

    @interfaces.setter
    def interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GlobalInterfaceArgs']]]]):
        pulumi.set(self, "interfaces", value)

    @property
    @pulumi.getter(name="ipsecOffload")
    def ipsec_offload(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable DPDK IPsec phase 2 offloading. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "ipsec_offload")

    @ipsec_offload.setter
    def ipsec_offload(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipsec_offload", value)

    @property
    @pulumi.getter(name="mbufpoolPercentage")
    def mbufpool_percentage(self) -> Optional[pulumi.Input[int]]:
        """
        Percentage of main memory allocated to DPDK packet buffer.
        """
        return pulumi.get(self, "mbufpool_percentage")

    @mbufpool_percentage.setter
    def mbufpool_percentage(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mbufpool_percentage", value)

    @property
    @pulumi.getter
    def multiqueue(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable multi-queue RX/TX support for all DPDK ports. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "multiqueue")

    @multiqueue.setter
    def multiqueue(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "multiqueue", value)

    @property
    @pulumi.getter(name="perSessionAccounting")
    def per_session_accounting(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable per-session accounting. Valid values: `disable`, `traffic-log-only`, `enable`.
        """
        return pulumi.get(self, "per_session_accounting")

    @per_session_accounting.setter
    def per_session_accounting(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "per_session_accounting", value)

    @property
    @pulumi.getter(name="sleepOnIdle")
    def sleep_on_idle(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable sleep-on-idle support for all FDH engines. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "sleep_on_idle")

    @sleep_on_idle.setter
    def sleep_on_idle(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sleep_on_idle", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable DPDK operation for the entire system. Valid values: `disable`, `enable`.
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
class _GlobalState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 elasticbuffer: Optional[pulumi.Input[str]] = None,
                 hugepage_percentage: Optional[pulumi.Input[int]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['GlobalInterfaceArgs']]]] = None,
                 ipsec_offload: Optional[pulumi.Input[str]] = None,
                 mbufpool_percentage: Optional[pulumi.Input[int]] = None,
                 multiqueue: Optional[pulumi.Input[str]] = None,
                 per_session_accounting: Optional[pulumi.Input[str]] = None,
                 sleep_on_idle: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Global resources.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] elasticbuffer: Enable/disable elasticbuffer support for all DPDK ports. Valid values: `disable`, `enable`.
        :param pulumi.Input[int] hugepage_percentage: Percentage of main memory allocated to hugepages, which are available for DPDK operation.
        :param pulumi.Input[Sequence[pulumi.Input['GlobalInterfaceArgs']]] interfaces: Physical interfaces that enable DPDK. The structure of `interface` block is documented below.
        :param pulumi.Input[str] ipsec_offload: Enable/disable DPDK IPsec phase 2 offloading. Valid values: `disable`, `enable`.
        :param pulumi.Input[int] mbufpool_percentage: Percentage of main memory allocated to DPDK packet buffer.
        :param pulumi.Input[str] multiqueue: Enable/disable multi-queue RX/TX support for all DPDK ports. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] per_session_accounting: Enable/disable per-session accounting. Valid values: `disable`, `traffic-log-only`, `enable`.
        :param pulumi.Input[str] sleep_on_idle: Enable/disable sleep-on-idle support for all FDH engines. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] status: Enable/disable DPDK operation for the entire system. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if elasticbuffer is not None:
            pulumi.set(__self__, "elasticbuffer", elasticbuffer)
        if hugepage_percentage is not None:
            pulumi.set(__self__, "hugepage_percentage", hugepage_percentage)
        if interfaces is not None:
            pulumi.set(__self__, "interfaces", interfaces)
        if ipsec_offload is not None:
            pulumi.set(__self__, "ipsec_offload", ipsec_offload)
        if mbufpool_percentage is not None:
            pulumi.set(__self__, "mbufpool_percentage", mbufpool_percentage)
        if multiqueue is not None:
            pulumi.set(__self__, "multiqueue", multiqueue)
        if per_session_accounting is not None:
            pulumi.set(__self__, "per_session_accounting", per_session_accounting)
        if sleep_on_idle is not None:
            pulumi.set(__self__, "sleep_on_idle", sleep_on_idle)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def elasticbuffer(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable elasticbuffer support for all DPDK ports. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "elasticbuffer")

    @elasticbuffer.setter
    def elasticbuffer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "elasticbuffer", value)

    @property
    @pulumi.getter(name="hugepagePercentage")
    def hugepage_percentage(self) -> Optional[pulumi.Input[int]]:
        """
        Percentage of main memory allocated to hugepages, which are available for DPDK operation.
        """
        return pulumi.get(self, "hugepage_percentage")

    @hugepage_percentage.setter
    def hugepage_percentage(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hugepage_percentage", value)

    @property
    @pulumi.getter
    def interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GlobalInterfaceArgs']]]]:
        """
        Physical interfaces that enable DPDK. The structure of `interface` block is documented below.
        """
        return pulumi.get(self, "interfaces")

    @interfaces.setter
    def interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GlobalInterfaceArgs']]]]):
        pulumi.set(self, "interfaces", value)

    @property
    @pulumi.getter(name="ipsecOffload")
    def ipsec_offload(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable DPDK IPsec phase 2 offloading. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "ipsec_offload")

    @ipsec_offload.setter
    def ipsec_offload(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipsec_offload", value)

    @property
    @pulumi.getter(name="mbufpoolPercentage")
    def mbufpool_percentage(self) -> Optional[pulumi.Input[int]]:
        """
        Percentage of main memory allocated to DPDK packet buffer.
        """
        return pulumi.get(self, "mbufpool_percentage")

    @mbufpool_percentage.setter
    def mbufpool_percentage(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mbufpool_percentage", value)

    @property
    @pulumi.getter
    def multiqueue(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable multi-queue RX/TX support for all DPDK ports. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "multiqueue")

    @multiqueue.setter
    def multiqueue(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "multiqueue", value)

    @property
    @pulumi.getter(name="perSessionAccounting")
    def per_session_accounting(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable per-session accounting. Valid values: `disable`, `traffic-log-only`, `enable`.
        """
        return pulumi.get(self, "per_session_accounting")

    @per_session_accounting.setter
    def per_session_accounting(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "per_session_accounting", value)

    @property
    @pulumi.getter(name="sleepOnIdle")
    def sleep_on_idle(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable sleep-on-idle support for all FDH engines. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "sleep_on_idle")

    @sleep_on_idle.setter
    def sleep_on_idle(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sleep_on_idle", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable DPDK operation for the entire system. Valid values: `disable`, `enable`.
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


class Global(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 elasticbuffer: Optional[pulumi.Input[str]] = None,
                 hugepage_percentage: Optional[pulumi.Input[int]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GlobalInterfaceArgs']]]]] = None,
                 ipsec_offload: Optional[pulumi.Input[str]] = None,
                 mbufpool_percentage: Optional[pulumi.Input[int]] = None,
                 multiqueue: Optional[pulumi.Input[str]] = None,
                 per_session_accounting: Optional[pulumi.Input[str]] = None,
                 sleep_on_idle: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure global DPDK options. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        Dpdk Global can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:dpdk/global:Global labelname DpdkGlobal
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:dpdk/global:Global labelname DpdkGlobal
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] elasticbuffer: Enable/disable elasticbuffer support for all DPDK ports. Valid values: `disable`, `enable`.
        :param pulumi.Input[int] hugepage_percentage: Percentage of main memory allocated to hugepages, which are available for DPDK operation.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GlobalInterfaceArgs']]]] interfaces: Physical interfaces that enable DPDK. The structure of `interface` block is documented below.
        :param pulumi.Input[str] ipsec_offload: Enable/disable DPDK IPsec phase 2 offloading. Valid values: `disable`, `enable`.
        :param pulumi.Input[int] mbufpool_percentage: Percentage of main memory allocated to DPDK packet buffer.
        :param pulumi.Input[str] multiqueue: Enable/disable multi-queue RX/TX support for all DPDK ports. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] per_session_accounting: Enable/disable per-session accounting. Valid values: `disable`, `traffic-log-only`, `enable`.
        :param pulumi.Input[str] sleep_on_idle: Enable/disable sleep-on-idle support for all FDH engines. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] status: Enable/disable DPDK operation for the entire system. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[GlobalArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure global DPDK options. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        Dpdk Global can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:dpdk/global:Global labelname DpdkGlobal
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:dpdk/global:Global labelname DpdkGlobal
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param GlobalArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GlobalArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 elasticbuffer: Optional[pulumi.Input[str]] = None,
                 hugepage_percentage: Optional[pulumi.Input[int]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GlobalInterfaceArgs']]]]] = None,
                 ipsec_offload: Optional[pulumi.Input[str]] = None,
                 mbufpool_percentage: Optional[pulumi.Input[int]] = None,
                 multiqueue: Optional[pulumi.Input[str]] = None,
                 per_session_accounting: Optional[pulumi.Input[str]] = None,
                 sleep_on_idle: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GlobalArgs.__new__(GlobalArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["elasticbuffer"] = elasticbuffer
            __props__.__dict__["hugepage_percentage"] = hugepage_percentage
            __props__.__dict__["interfaces"] = interfaces
            __props__.__dict__["ipsec_offload"] = ipsec_offload
            __props__.__dict__["mbufpool_percentage"] = mbufpool_percentage
            __props__.__dict__["multiqueue"] = multiqueue
            __props__.__dict__["per_session_accounting"] = per_session_accounting
            __props__.__dict__["sleep_on_idle"] = sleep_on_idle
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(Global, __self__).__init__(
            'fortios:dpdk/global:Global',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            elasticbuffer: Optional[pulumi.Input[str]] = None,
            hugepage_percentage: Optional[pulumi.Input[int]] = None,
            interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GlobalInterfaceArgs']]]]] = None,
            ipsec_offload: Optional[pulumi.Input[str]] = None,
            mbufpool_percentage: Optional[pulumi.Input[int]] = None,
            multiqueue: Optional[pulumi.Input[str]] = None,
            per_session_accounting: Optional[pulumi.Input[str]] = None,
            sleep_on_idle: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Global':
        """
        Get an existing Global resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] elasticbuffer: Enable/disable elasticbuffer support for all DPDK ports. Valid values: `disable`, `enable`.
        :param pulumi.Input[int] hugepage_percentage: Percentage of main memory allocated to hugepages, which are available for DPDK operation.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GlobalInterfaceArgs']]]] interfaces: Physical interfaces that enable DPDK. The structure of `interface` block is documented below.
        :param pulumi.Input[str] ipsec_offload: Enable/disable DPDK IPsec phase 2 offloading. Valid values: `disable`, `enable`.
        :param pulumi.Input[int] mbufpool_percentage: Percentage of main memory allocated to DPDK packet buffer.
        :param pulumi.Input[str] multiqueue: Enable/disable multi-queue RX/TX support for all DPDK ports. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] per_session_accounting: Enable/disable per-session accounting. Valid values: `disable`, `traffic-log-only`, `enable`.
        :param pulumi.Input[str] sleep_on_idle: Enable/disable sleep-on-idle support for all FDH engines. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] status: Enable/disable DPDK operation for the entire system. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GlobalState.__new__(_GlobalState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["elasticbuffer"] = elasticbuffer
        __props__.__dict__["hugepage_percentage"] = hugepage_percentage
        __props__.__dict__["interfaces"] = interfaces
        __props__.__dict__["ipsec_offload"] = ipsec_offload
        __props__.__dict__["mbufpool_percentage"] = mbufpool_percentage
        __props__.__dict__["multiqueue"] = multiqueue
        __props__.__dict__["per_session_accounting"] = per_session_accounting
        __props__.__dict__["sleep_on_idle"] = sleep_on_idle
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return Global(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def elasticbuffer(self) -> pulumi.Output[str]:
        """
        Enable/disable elasticbuffer support for all DPDK ports. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "elasticbuffer")

    @property
    @pulumi.getter(name="hugepagePercentage")
    def hugepage_percentage(self) -> pulumi.Output[int]:
        """
        Percentage of main memory allocated to hugepages, which are available for DPDK operation.
        """
        return pulumi.get(self, "hugepage_percentage")

    @property
    @pulumi.getter
    def interfaces(self) -> pulumi.Output[Optional[Sequence['outputs.GlobalInterface']]]:
        """
        Physical interfaces that enable DPDK. The structure of `interface` block is documented below.
        """
        return pulumi.get(self, "interfaces")

    @property
    @pulumi.getter(name="ipsecOffload")
    def ipsec_offload(self) -> pulumi.Output[str]:
        """
        Enable/disable DPDK IPsec phase 2 offloading. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "ipsec_offload")

    @property
    @pulumi.getter(name="mbufpoolPercentage")
    def mbufpool_percentage(self) -> pulumi.Output[int]:
        """
        Percentage of main memory allocated to DPDK packet buffer.
        """
        return pulumi.get(self, "mbufpool_percentage")

    @property
    @pulumi.getter
    def multiqueue(self) -> pulumi.Output[str]:
        """
        Enable/disable multi-queue RX/TX support for all DPDK ports. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "multiqueue")

    @property
    @pulumi.getter(name="perSessionAccounting")
    def per_session_accounting(self) -> pulumi.Output[str]:
        """
        Enable/disable per-session accounting. Valid values: `disable`, `traffic-log-only`, `enable`.
        """
        return pulumi.get(self, "per_session_accounting")

    @property
    @pulumi.getter(name="sleepOnIdle")
    def sleep_on_idle(self) -> pulumi.Output[str]:
        """
        Enable/disable sleep-on-idle support for all FDH engines. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "sleep_on_idle")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable DPDK operation for the entire system. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

