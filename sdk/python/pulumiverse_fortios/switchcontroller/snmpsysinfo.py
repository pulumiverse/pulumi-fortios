# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SnmpsysinfoArgs', 'Snmpsysinfo']

@pulumi.input_type
class SnmpsysinfoArgs:
    def __init__(__self__, *,
                 contact_info: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Snmpsysinfo resource.
        :param pulumi.Input[str] contact_info: Contact information.
        :param pulumi.Input[str] description: System description.
        :param pulumi.Input[str] engine_id: Local SNMP engine ID string (max 24 char).
        :param pulumi.Input[str] location: System location.
        :param pulumi.Input[str] status: Enable/disable SNMP. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if contact_info is not None:
            pulumi.set(__self__, "contact_info", contact_info)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if engine_id is not None:
            pulumi.set(__self__, "engine_id", engine_id)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="contactInfo")
    def contact_info(self) -> Optional[pulumi.Input[str]]:
        """
        Contact information.
        """
        return pulumi.get(self, "contact_info")

    @contact_info.setter
    def contact_info(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "contact_info", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        System description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="engineId")
    def engine_id(self) -> Optional[pulumi.Input[str]]:
        """
        Local SNMP engine ID string (max 24 char).
        """
        return pulumi.get(self, "engine_id")

    @engine_id.setter
    def engine_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_id", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        System location.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable SNMP. Valid values: `disable`, `enable`.
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
class _SnmpsysinfoState:
    def __init__(__self__, *,
                 contact_info: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Snmpsysinfo resources.
        :param pulumi.Input[str] contact_info: Contact information.
        :param pulumi.Input[str] description: System description.
        :param pulumi.Input[str] engine_id: Local SNMP engine ID string (max 24 char).
        :param pulumi.Input[str] location: System location.
        :param pulumi.Input[str] status: Enable/disable SNMP. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if contact_info is not None:
            pulumi.set(__self__, "contact_info", contact_info)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if engine_id is not None:
            pulumi.set(__self__, "engine_id", engine_id)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="contactInfo")
    def contact_info(self) -> Optional[pulumi.Input[str]]:
        """
        Contact information.
        """
        return pulumi.get(self, "contact_info")

    @contact_info.setter
    def contact_info(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "contact_info", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        System description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="engineId")
    def engine_id(self) -> Optional[pulumi.Input[str]]:
        """
        Local SNMP engine ID string (max 24 char).
        """
        return pulumi.get(self, "engine_id")

    @engine_id.setter
    def engine_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_id", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        System location.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable SNMP. Valid values: `disable`, `enable`.
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


class Snmpsysinfo(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 contact_info: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure FortiSwitch SNMP system information globally. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        SwitchController SnmpSysinfo can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:switchcontroller/snmpsysinfo:Snmpsysinfo labelname SwitchControllerSnmpSysinfo
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:switchcontroller/snmpsysinfo:Snmpsysinfo labelname SwitchControllerSnmpSysinfo
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] contact_info: Contact information.
        :param pulumi.Input[str] description: System description.
        :param pulumi.Input[str] engine_id: Local SNMP engine ID string (max 24 char).
        :param pulumi.Input[str] location: System location.
        :param pulumi.Input[str] status: Enable/disable SNMP. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SnmpsysinfoArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure FortiSwitch SNMP system information globally. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        SwitchController SnmpSysinfo can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:switchcontroller/snmpsysinfo:Snmpsysinfo labelname SwitchControllerSnmpSysinfo
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:switchcontroller/snmpsysinfo:Snmpsysinfo labelname SwitchControllerSnmpSysinfo
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SnmpsysinfoArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnmpsysinfoArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 contact_info: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SnmpsysinfoArgs.__new__(SnmpsysinfoArgs)

            __props__.__dict__["contact_info"] = contact_info
            __props__.__dict__["description"] = description
            __props__.__dict__["engine_id"] = engine_id
            __props__.__dict__["location"] = location
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(Snmpsysinfo, __self__).__init__(
            'fortios:switchcontroller/snmpsysinfo:Snmpsysinfo',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            contact_info: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            engine_id: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Snmpsysinfo':
        """
        Get an existing Snmpsysinfo resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] contact_info: Contact information.
        :param pulumi.Input[str] description: System description.
        :param pulumi.Input[str] engine_id: Local SNMP engine ID string (max 24 char).
        :param pulumi.Input[str] location: System location.
        :param pulumi.Input[str] status: Enable/disable SNMP. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnmpsysinfoState.__new__(_SnmpsysinfoState)

        __props__.__dict__["contact_info"] = contact_info
        __props__.__dict__["description"] = description
        __props__.__dict__["engine_id"] = engine_id
        __props__.__dict__["location"] = location
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return Snmpsysinfo(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="contactInfo")
    def contact_info(self) -> pulumi.Output[str]:
        """
        Contact information.
        """
        return pulumi.get(self, "contact_info")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        System description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="engineId")
    def engine_id(self) -> pulumi.Output[str]:
        """
        Local SNMP engine ID string (max 24 char).
        """
        return pulumi.get(self, "engine_id")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        System location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable SNMP. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[str]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

