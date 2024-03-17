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

__all__ = ['FortilinksettingsArgs', 'Fortilinksettings']

@pulumi.input_type
class FortilinksettingsArgs:
    def __init__(__self__, *,
                 fortilink: Optional[pulumi.Input[str]] = None,
                 inactive_timer: Optional[pulumi.Input[int]] = None,
                 link_down_flush: Optional[pulumi.Input[str]] = None,
                 nac_ports: Optional[pulumi.Input['FortilinksettingsNacPortsArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Fortilinksettings resource.
        :param pulumi.Input[str] fortilink: FortiLink interface to which this fortilink-setting belongs.
        :param pulumi.Input[int] inactive_timer: Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
        :param pulumi.Input[str] link_down_flush: Clear NAC and dynamic devices on switch ports on link down event. Valid values: `disable`, `enable`.
        :param pulumi.Input['FortilinksettingsNacPortsArgs'] nac_ports: NAC specific configuration. The structure of `nac_ports` block is documented below.
        :param pulumi.Input[str] name: FortiLink settings name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if fortilink is not None:
            pulumi.set(__self__, "fortilink", fortilink)
        if inactive_timer is not None:
            pulumi.set(__self__, "inactive_timer", inactive_timer)
        if link_down_flush is not None:
            pulumi.set(__self__, "link_down_flush", link_down_flush)
        if nac_ports is not None:
            pulumi.set(__self__, "nac_ports", nac_ports)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def fortilink(self) -> Optional[pulumi.Input[str]]:
        """
        FortiLink interface to which this fortilink-setting belongs.
        """
        return pulumi.get(self, "fortilink")

    @fortilink.setter
    def fortilink(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fortilink", value)

    @property
    @pulumi.getter(name="inactiveTimer")
    def inactive_timer(self) -> Optional[pulumi.Input[int]]:
        """
        Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
        """
        return pulumi.get(self, "inactive_timer")

    @inactive_timer.setter
    def inactive_timer(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "inactive_timer", value)

    @property
    @pulumi.getter(name="linkDownFlush")
    def link_down_flush(self) -> Optional[pulumi.Input[str]]:
        """
        Clear NAC and dynamic devices on switch ports on link down event. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "link_down_flush")

    @link_down_flush.setter
    def link_down_flush(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "link_down_flush", value)

    @property
    @pulumi.getter(name="nacPorts")
    def nac_ports(self) -> Optional[pulumi.Input['FortilinksettingsNacPortsArgs']]:
        """
        NAC specific configuration. The structure of `nac_ports` block is documented below.
        """
        return pulumi.get(self, "nac_ports")

    @nac_ports.setter
    def nac_ports(self, value: Optional[pulumi.Input['FortilinksettingsNacPortsArgs']]):
        pulumi.set(self, "nac_ports", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        FortiLink settings name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
class _FortilinksettingsState:
    def __init__(__self__, *,
                 fortilink: Optional[pulumi.Input[str]] = None,
                 inactive_timer: Optional[pulumi.Input[int]] = None,
                 link_down_flush: Optional[pulumi.Input[str]] = None,
                 nac_ports: Optional[pulumi.Input['FortilinksettingsNacPortsArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Fortilinksettings resources.
        :param pulumi.Input[str] fortilink: FortiLink interface to which this fortilink-setting belongs.
        :param pulumi.Input[int] inactive_timer: Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
        :param pulumi.Input[str] link_down_flush: Clear NAC and dynamic devices on switch ports on link down event. Valid values: `disable`, `enable`.
        :param pulumi.Input['FortilinksettingsNacPortsArgs'] nac_ports: NAC specific configuration. The structure of `nac_ports` block is documented below.
        :param pulumi.Input[str] name: FortiLink settings name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if fortilink is not None:
            pulumi.set(__self__, "fortilink", fortilink)
        if inactive_timer is not None:
            pulumi.set(__self__, "inactive_timer", inactive_timer)
        if link_down_flush is not None:
            pulumi.set(__self__, "link_down_flush", link_down_flush)
        if nac_ports is not None:
            pulumi.set(__self__, "nac_ports", nac_ports)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def fortilink(self) -> Optional[pulumi.Input[str]]:
        """
        FortiLink interface to which this fortilink-setting belongs.
        """
        return pulumi.get(self, "fortilink")

    @fortilink.setter
    def fortilink(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fortilink", value)

    @property
    @pulumi.getter(name="inactiveTimer")
    def inactive_timer(self) -> Optional[pulumi.Input[int]]:
        """
        Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
        """
        return pulumi.get(self, "inactive_timer")

    @inactive_timer.setter
    def inactive_timer(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "inactive_timer", value)

    @property
    @pulumi.getter(name="linkDownFlush")
    def link_down_flush(self) -> Optional[pulumi.Input[str]]:
        """
        Clear NAC and dynamic devices on switch ports on link down event. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "link_down_flush")

    @link_down_flush.setter
    def link_down_flush(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "link_down_flush", value)

    @property
    @pulumi.getter(name="nacPorts")
    def nac_ports(self) -> Optional[pulumi.Input['FortilinksettingsNacPortsArgs']]:
        """
        NAC specific configuration. The structure of `nac_ports` block is documented below.
        """
        return pulumi.get(self, "nac_ports")

    @nac_ports.setter
    def nac_ports(self, value: Optional[pulumi.Input['FortilinksettingsNacPortsArgs']]):
        pulumi.set(self, "nac_ports", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        FortiLink settings name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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


class Fortilinksettings(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fortilink: Optional[pulumi.Input[str]] = None,
                 inactive_timer: Optional[pulumi.Input[int]] = None,
                 link_down_flush: Optional[pulumi.Input[str]] = None,
                 nac_ports: Optional[pulumi.Input[pulumi.InputType['FortilinksettingsNacPortsArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure integrated FortiLink settings for FortiSwitch. Applies to FortiOS Version `>= 7.0.1`.

        ## Import

        SwitchController FortilinkSettings can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:switchcontroller/fortilinksettings:Fortilinksettings labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:switchcontroller/fortilinksettings:Fortilinksettings labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fortilink: FortiLink interface to which this fortilink-setting belongs.
        :param pulumi.Input[int] inactive_timer: Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
        :param pulumi.Input[str] link_down_flush: Clear NAC and dynamic devices on switch ports on link down event. Valid values: `disable`, `enable`.
        :param pulumi.Input[pulumi.InputType['FortilinksettingsNacPortsArgs']] nac_ports: NAC specific configuration. The structure of `nac_ports` block is documented below.
        :param pulumi.Input[str] name: FortiLink settings name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FortilinksettingsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure integrated FortiLink settings for FortiSwitch. Applies to FortiOS Version `>= 7.0.1`.

        ## Import

        SwitchController FortilinkSettings can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:switchcontroller/fortilinksettings:Fortilinksettings labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:switchcontroller/fortilinksettings:Fortilinksettings labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FortilinksettingsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FortilinksettingsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fortilink: Optional[pulumi.Input[str]] = None,
                 inactive_timer: Optional[pulumi.Input[int]] = None,
                 link_down_flush: Optional[pulumi.Input[str]] = None,
                 nac_ports: Optional[pulumi.Input[pulumi.InputType['FortilinksettingsNacPortsArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FortilinksettingsArgs.__new__(FortilinksettingsArgs)

            __props__.__dict__["fortilink"] = fortilink
            __props__.__dict__["inactive_timer"] = inactive_timer
            __props__.__dict__["link_down_flush"] = link_down_flush
            __props__.__dict__["nac_ports"] = nac_ports
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
        super(Fortilinksettings, __self__).__init__(
            'fortios:switchcontroller/fortilinksettings:Fortilinksettings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            fortilink: Optional[pulumi.Input[str]] = None,
            inactive_timer: Optional[pulumi.Input[int]] = None,
            link_down_flush: Optional[pulumi.Input[str]] = None,
            nac_ports: Optional[pulumi.Input[pulumi.InputType['FortilinksettingsNacPortsArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Fortilinksettings':
        """
        Get an existing Fortilinksettings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fortilink: FortiLink interface to which this fortilink-setting belongs.
        :param pulumi.Input[int] inactive_timer: Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
        :param pulumi.Input[str] link_down_flush: Clear NAC and dynamic devices on switch ports on link down event. Valid values: `disable`, `enable`.
        :param pulumi.Input[pulumi.InputType['FortilinksettingsNacPortsArgs']] nac_ports: NAC specific configuration. The structure of `nac_ports` block is documented below.
        :param pulumi.Input[str] name: FortiLink settings name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FortilinksettingsState.__new__(_FortilinksettingsState)

        __props__.__dict__["fortilink"] = fortilink
        __props__.__dict__["inactive_timer"] = inactive_timer
        __props__.__dict__["link_down_flush"] = link_down_flush
        __props__.__dict__["nac_ports"] = nac_ports
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        return Fortilinksettings(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def fortilink(self) -> pulumi.Output[str]:
        """
        FortiLink interface to which this fortilink-setting belongs.
        """
        return pulumi.get(self, "fortilink")

    @property
    @pulumi.getter(name="inactiveTimer")
    def inactive_timer(self) -> pulumi.Output[int]:
        """
        Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
        """
        return pulumi.get(self, "inactive_timer")

    @property
    @pulumi.getter(name="linkDownFlush")
    def link_down_flush(self) -> pulumi.Output[str]:
        """
        Clear NAC and dynamic devices on switch ports on link down event. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "link_down_flush")

    @property
    @pulumi.getter(name="nacPorts")
    def nac_ports(self) -> pulumi.Output['outputs.FortilinksettingsNacPorts']:
        """
        NAC specific configuration. The structure of `nac_ports` block is documented below.
        """
        return pulumi.get(self, "nac_ports")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        FortiLink settings name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

