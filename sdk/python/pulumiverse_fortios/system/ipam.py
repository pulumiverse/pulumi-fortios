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

__all__ = ['IpamArgs', 'Ipam']

@pulumi.input_type
class IpamArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 pool_subnet: Optional[pulumi.Input[str]] = None,
                 pools: Optional[pulumi.Input[Sequence[pulumi.Input['IpamPoolArgs']]]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['IpamRuleArgs']]]] = None,
                 server_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Ipam resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] pool_subnet: Configure IPAM pool subnet, Class A - Class B subnet.
        :param pulumi.Input[Sequence[pulumi.Input['IpamPoolArgs']]] pools: Configure IPAM pools. The structure of `pools` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['IpamRuleArgs']]] rules: Configure IPAM allocation rules. The structure of `rules` block is documented below.
        :param pulumi.Input[str] server_type: Configure the type of IPAM server to use. Valid values: `cloud`, `fabric-root`.
        :param pulumi.Input[str] status: Enable/disable IP address management services. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if pool_subnet is not None:
            pulumi.set(__self__, "pool_subnet", pool_subnet)
        if pools is not None:
            pulumi.set(__self__, "pools", pools)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if server_type is not None:
            pulumi.set(__self__, "server_type", server_type)
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
    @pulumi.getter(name="poolSubnet")
    def pool_subnet(self) -> Optional[pulumi.Input[str]]:
        """
        Configure IPAM pool subnet, Class A - Class B subnet.
        """
        return pulumi.get(self, "pool_subnet")

    @pool_subnet.setter
    def pool_subnet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pool_subnet", value)

    @property
    @pulumi.getter
    def pools(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IpamPoolArgs']]]]:
        """
        Configure IPAM pools. The structure of `pools` block is documented below.
        """
        return pulumi.get(self, "pools")

    @pools.setter
    def pools(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IpamPoolArgs']]]]):
        pulumi.set(self, "pools", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IpamRuleArgs']]]]:
        """
        Configure IPAM allocation rules. The structure of `rules` block is documented below.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IpamRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter(name="serverType")
    def server_type(self) -> Optional[pulumi.Input[str]]:
        """
        Configure the type of IPAM server to use. Valid values: `cloud`, `fabric-root`.
        """
        return pulumi.get(self, "server_type")

    @server_type.setter
    def server_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_type", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable IP address management services. Valid values: `enable`, `disable`.
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
class _IpamState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 pool_subnet: Optional[pulumi.Input[str]] = None,
                 pools: Optional[pulumi.Input[Sequence[pulumi.Input['IpamPoolArgs']]]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['IpamRuleArgs']]]] = None,
                 server_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Ipam resources.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] pool_subnet: Configure IPAM pool subnet, Class A - Class B subnet.
        :param pulumi.Input[Sequence[pulumi.Input['IpamPoolArgs']]] pools: Configure IPAM pools. The structure of `pools` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['IpamRuleArgs']]] rules: Configure IPAM allocation rules. The structure of `rules` block is documented below.
        :param pulumi.Input[str] server_type: Configure the type of IPAM server to use. Valid values: `cloud`, `fabric-root`.
        :param pulumi.Input[str] status: Enable/disable IP address management services. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if pool_subnet is not None:
            pulumi.set(__self__, "pool_subnet", pool_subnet)
        if pools is not None:
            pulumi.set(__self__, "pools", pools)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if server_type is not None:
            pulumi.set(__self__, "server_type", server_type)
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
    @pulumi.getter(name="poolSubnet")
    def pool_subnet(self) -> Optional[pulumi.Input[str]]:
        """
        Configure IPAM pool subnet, Class A - Class B subnet.
        """
        return pulumi.get(self, "pool_subnet")

    @pool_subnet.setter
    def pool_subnet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pool_subnet", value)

    @property
    @pulumi.getter
    def pools(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IpamPoolArgs']]]]:
        """
        Configure IPAM pools. The structure of `pools` block is documented below.
        """
        return pulumi.get(self, "pools")

    @pools.setter
    def pools(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IpamPoolArgs']]]]):
        pulumi.set(self, "pools", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IpamRuleArgs']]]]:
        """
        Configure IPAM allocation rules. The structure of `rules` block is documented below.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IpamRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter(name="serverType")
    def server_type(self) -> Optional[pulumi.Input[str]]:
        """
        Configure the type of IPAM server to use. Valid values: `cloud`, `fabric-root`.
        """
        return pulumi.get(self, "server_type")

    @server_type.setter
    def server_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_type", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable IP address management services. Valid values: `enable`, `disable`.
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


class Ipam(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 pool_subnet: Optional[pulumi.Input[str]] = None,
                 pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpamPoolArgs']]]]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpamRuleArgs']]]]] = None,
                 server_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure IP address management services. Applies to FortiOS Version `>= 7.0.2`.

        ## Import

        System Ipam can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/ipam:Ipam labelname SystemIpam
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/ipam:Ipam labelname SystemIpam
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] pool_subnet: Configure IPAM pool subnet, Class A - Class B subnet.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpamPoolArgs']]]] pools: Configure IPAM pools. The structure of `pools` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpamRuleArgs']]]] rules: Configure IPAM allocation rules. The structure of `rules` block is documented below.
        :param pulumi.Input[str] server_type: Configure the type of IPAM server to use. Valid values: `cloud`, `fabric-root`.
        :param pulumi.Input[str] status: Enable/disable IP address management services. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[IpamArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure IP address management services. Applies to FortiOS Version `>= 7.0.2`.

        ## Import

        System Ipam can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/ipam:Ipam labelname SystemIpam
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/ipam:Ipam labelname SystemIpam
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param IpamArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpamArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 pool_subnet: Optional[pulumi.Input[str]] = None,
                 pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpamPoolArgs']]]]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpamRuleArgs']]]]] = None,
                 server_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpamArgs.__new__(IpamArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["pool_subnet"] = pool_subnet
            __props__.__dict__["pools"] = pools
            __props__.__dict__["rules"] = rules
            __props__.__dict__["server_type"] = server_type
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(Ipam, __self__).__init__(
            'fortios:system/ipam:Ipam',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            pool_subnet: Optional[pulumi.Input[str]] = None,
            pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpamPoolArgs']]]]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpamRuleArgs']]]]] = None,
            server_type: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Ipam':
        """
        Get an existing Ipam resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] pool_subnet: Configure IPAM pool subnet, Class A - Class B subnet.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpamPoolArgs']]]] pools: Configure IPAM pools. The structure of `pools` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpamRuleArgs']]]] rules: Configure IPAM allocation rules. The structure of `rules` block is documented below.
        :param pulumi.Input[str] server_type: Configure the type of IPAM server to use. Valid values: `cloud`, `fabric-root`.
        :param pulumi.Input[str] status: Enable/disable IP address management services. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpamState.__new__(_IpamState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["pool_subnet"] = pool_subnet
        __props__.__dict__["pools"] = pools
        __props__.__dict__["rules"] = rules
        __props__.__dict__["server_type"] = server_type
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return Ipam(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="poolSubnet")
    def pool_subnet(self) -> pulumi.Output[str]:
        """
        Configure IPAM pool subnet, Class A - Class B subnet.
        """
        return pulumi.get(self, "pool_subnet")

    @property
    @pulumi.getter
    def pools(self) -> pulumi.Output[Optional[Sequence['outputs.IpamPool']]]:
        """
        Configure IPAM pools. The structure of `pools` block is documented below.
        """
        return pulumi.get(self, "pools")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Optional[Sequence['outputs.IpamRule']]]:
        """
        Configure IPAM allocation rules. The structure of `rules` block is documented below.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="serverType")
    def server_type(self) -> pulumi.Output[str]:
        """
        Configure the type of IPAM server to use. Valid values: `cloud`, `fabric-root`.
        """
        return pulumi.get(self, "server_type")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable IP address management services. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

