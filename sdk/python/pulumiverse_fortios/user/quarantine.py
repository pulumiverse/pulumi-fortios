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

__all__ = ['QuarantineArgs', 'Quarantine']

@pulumi.input_type
class QuarantineArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 firewall_groups: Optional[pulumi.Input[str]] = None,
                 quarantine: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input['QuarantineTargetArgs']]]] = None,
                 traffic_policy: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Quarantine resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] firewall_groups: Firewall address group which includes all quarantine MAC address.
        :param pulumi.Input[str] quarantine: Enable/disable quarantine. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input['QuarantineTargetArgs']]] targets: Quarantine entry to hold multiple MACs. The structure of `targets` block is documented below.
        :param pulumi.Input[str] traffic_policy: Traffic policy for quarantined MACs.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if firewall_groups is not None:
            pulumi.set(__self__, "firewall_groups", firewall_groups)
        if quarantine is not None:
            pulumi.set(__self__, "quarantine", quarantine)
        if targets is not None:
            pulumi.set(__self__, "targets", targets)
        if traffic_policy is not None:
            pulumi.set(__self__, "traffic_policy", traffic_policy)
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
    @pulumi.getter(name="firewallGroups")
    def firewall_groups(self) -> Optional[pulumi.Input[str]]:
        """
        Firewall address group which includes all quarantine MAC address.
        """
        return pulumi.get(self, "firewall_groups")

    @firewall_groups.setter
    def firewall_groups(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "firewall_groups", value)

    @property
    @pulumi.getter
    def quarantine(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable quarantine. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "quarantine")

    @quarantine.setter
    def quarantine(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "quarantine", value)

    @property
    @pulumi.getter
    def targets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['QuarantineTargetArgs']]]]:
        """
        Quarantine entry to hold multiple MACs. The structure of `targets` block is documented below.
        """
        return pulumi.get(self, "targets")

    @targets.setter
    def targets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['QuarantineTargetArgs']]]]):
        pulumi.set(self, "targets", value)

    @property
    @pulumi.getter(name="trafficPolicy")
    def traffic_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Traffic policy for quarantined MACs.
        """
        return pulumi.get(self, "traffic_policy")

    @traffic_policy.setter
    def traffic_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_policy", value)

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
class _QuarantineState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 firewall_groups: Optional[pulumi.Input[str]] = None,
                 quarantine: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input['QuarantineTargetArgs']]]] = None,
                 traffic_policy: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Quarantine resources.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] firewall_groups: Firewall address group which includes all quarantine MAC address.
        :param pulumi.Input[str] quarantine: Enable/disable quarantine. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input['QuarantineTargetArgs']]] targets: Quarantine entry to hold multiple MACs. The structure of `targets` block is documented below.
        :param pulumi.Input[str] traffic_policy: Traffic policy for quarantined MACs.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if firewall_groups is not None:
            pulumi.set(__self__, "firewall_groups", firewall_groups)
        if quarantine is not None:
            pulumi.set(__self__, "quarantine", quarantine)
        if targets is not None:
            pulumi.set(__self__, "targets", targets)
        if traffic_policy is not None:
            pulumi.set(__self__, "traffic_policy", traffic_policy)
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
    @pulumi.getter(name="firewallGroups")
    def firewall_groups(self) -> Optional[pulumi.Input[str]]:
        """
        Firewall address group which includes all quarantine MAC address.
        """
        return pulumi.get(self, "firewall_groups")

    @firewall_groups.setter
    def firewall_groups(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "firewall_groups", value)

    @property
    @pulumi.getter
    def quarantine(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable quarantine. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "quarantine")

    @quarantine.setter
    def quarantine(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "quarantine", value)

    @property
    @pulumi.getter
    def targets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['QuarantineTargetArgs']]]]:
        """
        Quarantine entry to hold multiple MACs. The structure of `targets` block is documented below.
        """
        return pulumi.get(self, "targets")

    @targets.setter
    def targets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['QuarantineTargetArgs']]]]):
        pulumi.set(self, "targets", value)

    @property
    @pulumi.getter(name="trafficPolicy")
    def traffic_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Traffic policy for quarantined MACs.
        """
        return pulumi.get(self, "traffic_policy")

    @traffic_policy.setter
    def traffic_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_policy", value)

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


class Quarantine(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 firewall_groups: Optional[pulumi.Input[str]] = None,
                 quarantine: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['QuarantineTargetArgs']]]]] = None,
                 traffic_policy: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure quarantine support.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.user.Quarantine("trname", quarantine="enable")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        User Quarantine can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:user/quarantine:Quarantine labelname UserQuarantine
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:user/quarantine:Quarantine labelname UserQuarantine
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] firewall_groups: Firewall address group which includes all quarantine MAC address.
        :param pulumi.Input[str] quarantine: Enable/disable quarantine. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['QuarantineTargetArgs']]]] targets: Quarantine entry to hold multiple MACs. The structure of `targets` block is documented below.
        :param pulumi.Input[str] traffic_policy: Traffic policy for quarantined MACs.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[QuarantineArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure quarantine support.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.user.Quarantine("trname", quarantine="enable")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        User Quarantine can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:user/quarantine:Quarantine labelname UserQuarantine
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:user/quarantine:Quarantine labelname UserQuarantine
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param QuarantineArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(QuarantineArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 firewall_groups: Optional[pulumi.Input[str]] = None,
                 quarantine: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['QuarantineTargetArgs']]]]] = None,
                 traffic_policy: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = QuarantineArgs.__new__(QuarantineArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["firewall_groups"] = firewall_groups
            __props__.__dict__["quarantine"] = quarantine
            __props__.__dict__["targets"] = targets
            __props__.__dict__["traffic_policy"] = traffic_policy
            __props__.__dict__["vdomparam"] = vdomparam
        super(Quarantine, __self__).__init__(
            'fortios:user/quarantine:Quarantine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            firewall_groups: Optional[pulumi.Input[str]] = None,
            quarantine: Optional[pulumi.Input[str]] = None,
            targets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['QuarantineTargetArgs']]]]] = None,
            traffic_policy: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Quarantine':
        """
        Get an existing Quarantine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] firewall_groups: Firewall address group which includes all quarantine MAC address.
        :param pulumi.Input[str] quarantine: Enable/disable quarantine. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['QuarantineTargetArgs']]]] targets: Quarantine entry to hold multiple MACs. The structure of `targets` block is documented below.
        :param pulumi.Input[str] traffic_policy: Traffic policy for quarantined MACs.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _QuarantineState.__new__(_QuarantineState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["firewall_groups"] = firewall_groups
        __props__.__dict__["quarantine"] = quarantine
        __props__.__dict__["targets"] = targets
        __props__.__dict__["traffic_policy"] = traffic_policy
        __props__.__dict__["vdomparam"] = vdomparam
        return Quarantine(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="firewallGroups")
    def firewall_groups(self) -> pulumi.Output[str]:
        """
        Firewall address group which includes all quarantine MAC address.
        """
        return pulumi.get(self, "firewall_groups")

    @property
    @pulumi.getter
    def quarantine(self) -> pulumi.Output[str]:
        """
        Enable/disable quarantine. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "quarantine")

    @property
    @pulumi.getter
    def targets(self) -> pulumi.Output[Optional[Sequence['outputs.QuarantineTarget']]]:
        """
        Quarantine entry to hold multiple MACs. The structure of `targets` block is documented below.
        """
        return pulumi.get(self, "targets")

    @property
    @pulumi.getter(name="trafficPolicy")
    def traffic_policy(self) -> pulumi.Output[str]:
        """
        Traffic policy for quarantined MACs.
        """
        return pulumi.get(self, "traffic_policy")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

