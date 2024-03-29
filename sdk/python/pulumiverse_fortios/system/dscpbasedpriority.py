# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DscpbasedpriorityArgs', 'Dscpbasedpriority']

@pulumi.input_type
class DscpbasedpriorityArgs:
    def __init__(__self__, *,
                 ds: Optional[pulumi.Input[int]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 priority: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Dscpbasedpriority resource.
        :param pulumi.Input[int] ds: DSCP(DiffServ) DS value (0 - 63).
        :param pulumi.Input[int] fosid: Item ID.
        :param pulumi.Input[str] priority: DSCP based priority level. Valid values: `low`, `medium`, `high`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if ds is not None:
            pulumi.set(__self__, "ds", ds)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def ds(self) -> Optional[pulumi.Input[int]]:
        """
        DSCP(DiffServ) DS value (0 - 63).
        """
        return pulumi.get(self, "ds")

    @ds.setter
    def ds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ds", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        Item ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[str]]:
        """
        DSCP based priority level. Valid values: `low`, `medium`, `high`.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "priority", value)

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
class _DscpbasedpriorityState:
    def __init__(__self__, *,
                 ds: Optional[pulumi.Input[int]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 priority: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Dscpbasedpriority resources.
        :param pulumi.Input[int] ds: DSCP(DiffServ) DS value (0 - 63).
        :param pulumi.Input[int] fosid: Item ID.
        :param pulumi.Input[str] priority: DSCP based priority level. Valid values: `low`, `medium`, `high`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if ds is not None:
            pulumi.set(__self__, "ds", ds)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def ds(self) -> Optional[pulumi.Input[int]]:
        """
        DSCP(DiffServ) DS value (0 - 63).
        """
        return pulumi.get(self, "ds")

    @ds.setter
    def ds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ds", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        Item ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[str]]:
        """
        DSCP based priority level. Valid values: `low`, `medium`, `high`.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "priority", value)

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


class Dscpbasedpriority(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ds: Optional[pulumi.Input[int]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 priority: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure DSCP based priority table.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.system.Dscpbasedpriority("trname",
            ds=1,
            fosid=1,
            priority="low")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        System DscpBasedPriority can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/dscpbasedpriority:Dscpbasedpriority labelname {{fosid}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/dscpbasedpriority:Dscpbasedpriority labelname {{fosid}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] ds: DSCP(DiffServ) DS value (0 - 63).
        :param pulumi.Input[int] fosid: Item ID.
        :param pulumi.Input[str] priority: DSCP based priority level. Valid values: `low`, `medium`, `high`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DscpbasedpriorityArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure DSCP based priority table.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.system.Dscpbasedpriority("trname",
            ds=1,
            fosid=1,
            priority="low")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        System DscpBasedPriority can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/dscpbasedpriority:Dscpbasedpriority labelname {{fosid}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/dscpbasedpriority:Dscpbasedpriority labelname {{fosid}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param DscpbasedpriorityArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DscpbasedpriorityArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ds: Optional[pulumi.Input[int]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 priority: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DscpbasedpriorityArgs.__new__(DscpbasedpriorityArgs)

            __props__.__dict__["ds"] = ds
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["priority"] = priority
            __props__.__dict__["vdomparam"] = vdomparam
        super(Dscpbasedpriority, __self__).__init__(
            'fortios:system/dscpbasedpriority:Dscpbasedpriority',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ds: Optional[pulumi.Input[int]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            priority: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Dscpbasedpriority':
        """
        Get an existing Dscpbasedpriority resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] ds: DSCP(DiffServ) DS value (0 - 63).
        :param pulumi.Input[int] fosid: Item ID.
        :param pulumi.Input[str] priority: DSCP based priority level. Valid values: `low`, `medium`, `high`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DscpbasedpriorityState.__new__(_DscpbasedpriorityState)

        __props__.__dict__["ds"] = ds
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["priority"] = priority
        __props__.__dict__["vdomparam"] = vdomparam
        return Dscpbasedpriority(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def ds(self) -> pulumi.Output[int]:
        """
        DSCP(DiffServ) DS value (0 - 63).
        """
        return pulumi.get(self, "ds")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        """
        Item ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[str]:
        """
        DSCP based priority level. Valid values: `low`, `medium`, `high`.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

