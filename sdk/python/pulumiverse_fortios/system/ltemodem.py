# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LtemodemArgs', 'Ltemodem']

@pulumi.input_type
class LtemodemArgs:
    def __init__(__self__, *,
                 apn: Optional[pulumi.Input[str]] = None,
                 authtype: Optional[pulumi.Input[str]] = None,
                 extra_init: Optional[pulumi.Input[str]] = None,
                 holddown_timer: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 modem_port: Optional[pulumi.Input[int]] = None,
                 passwd: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Ltemodem resource.
        :param pulumi.Input[str] apn: Login APN string for PDP-IP packet data calls.
        :param pulumi.Input[str] authtype: Authentication type for PDP-IP packet data calls. Valid values: `none`, `pap`, `chap`.
        :param pulumi.Input[str] extra_init: Extra initialization string for USB LTE/WIMAX devices.
        :param pulumi.Input[int] holddown_timer: Hold down timer (10 - 60 sec).
        :param pulumi.Input[str] interface: The interface that the modem is acting as a redundant interface for.
        :param pulumi.Input[str] mode: Modem operation mode. Valid values: `standalone`, `redundant`.
        :param pulumi.Input[int] modem_port: Modem port index (0 - 20).
        :param pulumi.Input[str] passwd: Authentication password for PDP-IP packet data calls.
        :param pulumi.Input[str] status: Enable/disable USB LTE/WIMAX device. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] username: Authentication username for PDP-IP packet data calls.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if apn is not None:
            pulumi.set(__self__, "apn", apn)
        if authtype is not None:
            pulumi.set(__self__, "authtype", authtype)
        if extra_init is not None:
            pulumi.set(__self__, "extra_init", extra_init)
        if holddown_timer is not None:
            pulumi.set(__self__, "holddown_timer", holddown_timer)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if modem_port is not None:
            pulumi.set(__self__, "modem_port", modem_port)
        if passwd is not None:
            pulumi.set(__self__, "passwd", passwd)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def apn(self) -> Optional[pulumi.Input[str]]:
        """
        Login APN string for PDP-IP packet data calls.
        """
        return pulumi.get(self, "apn")

    @apn.setter
    def apn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "apn", value)

    @property
    @pulumi.getter
    def authtype(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication type for PDP-IP packet data calls. Valid values: `none`, `pap`, `chap`.
        """
        return pulumi.get(self, "authtype")

    @authtype.setter
    def authtype(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authtype", value)

    @property
    @pulumi.getter(name="extraInit")
    def extra_init(self) -> Optional[pulumi.Input[str]]:
        """
        Extra initialization string for USB LTE/WIMAX devices.
        """
        return pulumi.get(self, "extra_init")

    @extra_init.setter
    def extra_init(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "extra_init", value)

    @property
    @pulumi.getter(name="holddownTimer")
    def holddown_timer(self) -> Optional[pulumi.Input[int]]:
        """
        Hold down timer (10 - 60 sec).
        """
        return pulumi.get(self, "holddown_timer")

    @holddown_timer.setter
    def holddown_timer(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "holddown_timer", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        The interface that the modem is acting as a redundant interface for.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        Modem operation mode. Valid values: `standalone`, `redundant`.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="modemPort")
    def modem_port(self) -> Optional[pulumi.Input[int]]:
        """
        Modem port index (0 - 20).
        """
        return pulumi.get(self, "modem_port")

    @modem_port.setter
    def modem_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "modem_port", value)

    @property
    @pulumi.getter
    def passwd(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication password for PDP-IP packet data calls.
        """
        return pulumi.get(self, "passwd")

    @passwd.setter
    def passwd(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "passwd", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable USB LTE/WIMAX device. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication username for PDP-IP packet data calls.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

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
class _LtemodemState:
    def __init__(__self__, *,
                 apn: Optional[pulumi.Input[str]] = None,
                 authtype: Optional[pulumi.Input[str]] = None,
                 extra_init: Optional[pulumi.Input[str]] = None,
                 holddown_timer: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 modem_port: Optional[pulumi.Input[int]] = None,
                 passwd: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Ltemodem resources.
        :param pulumi.Input[str] apn: Login APN string for PDP-IP packet data calls.
        :param pulumi.Input[str] authtype: Authentication type for PDP-IP packet data calls. Valid values: `none`, `pap`, `chap`.
        :param pulumi.Input[str] extra_init: Extra initialization string for USB LTE/WIMAX devices.
        :param pulumi.Input[int] holddown_timer: Hold down timer (10 - 60 sec).
        :param pulumi.Input[str] interface: The interface that the modem is acting as a redundant interface for.
        :param pulumi.Input[str] mode: Modem operation mode. Valid values: `standalone`, `redundant`.
        :param pulumi.Input[int] modem_port: Modem port index (0 - 20).
        :param pulumi.Input[str] passwd: Authentication password for PDP-IP packet data calls.
        :param pulumi.Input[str] status: Enable/disable USB LTE/WIMAX device. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] username: Authentication username for PDP-IP packet data calls.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if apn is not None:
            pulumi.set(__self__, "apn", apn)
        if authtype is not None:
            pulumi.set(__self__, "authtype", authtype)
        if extra_init is not None:
            pulumi.set(__self__, "extra_init", extra_init)
        if holddown_timer is not None:
            pulumi.set(__self__, "holddown_timer", holddown_timer)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if modem_port is not None:
            pulumi.set(__self__, "modem_port", modem_port)
        if passwd is not None:
            pulumi.set(__self__, "passwd", passwd)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def apn(self) -> Optional[pulumi.Input[str]]:
        """
        Login APN string for PDP-IP packet data calls.
        """
        return pulumi.get(self, "apn")

    @apn.setter
    def apn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "apn", value)

    @property
    @pulumi.getter
    def authtype(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication type for PDP-IP packet data calls. Valid values: `none`, `pap`, `chap`.
        """
        return pulumi.get(self, "authtype")

    @authtype.setter
    def authtype(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authtype", value)

    @property
    @pulumi.getter(name="extraInit")
    def extra_init(self) -> Optional[pulumi.Input[str]]:
        """
        Extra initialization string for USB LTE/WIMAX devices.
        """
        return pulumi.get(self, "extra_init")

    @extra_init.setter
    def extra_init(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "extra_init", value)

    @property
    @pulumi.getter(name="holddownTimer")
    def holddown_timer(self) -> Optional[pulumi.Input[int]]:
        """
        Hold down timer (10 - 60 sec).
        """
        return pulumi.get(self, "holddown_timer")

    @holddown_timer.setter
    def holddown_timer(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "holddown_timer", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        The interface that the modem is acting as a redundant interface for.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        Modem operation mode. Valid values: `standalone`, `redundant`.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="modemPort")
    def modem_port(self) -> Optional[pulumi.Input[int]]:
        """
        Modem port index (0 - 20).
        """
        return pulumi.get(self, "modem_port")

    @modem_port.setter
    def modem_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "modem_port", value)

    @property
    @pulumi.getter
    def passwd(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication password for PDP-IP packet data calls.
        """
        return pulumi.get(self, "passwd")

    @passwd.setter
    def passwd(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "passwd", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable USB LTE/WIMAX device. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication username for PDP-IP packet data calls.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

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


class Ltemodem(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apn: Optional[pulumi.Input[str]] = None,
                 authtype: Optional[pulumi.Input[str]] = None,
                 extra_init: Optional[pulumi.Input[str]] = None,
                 holddown_timer: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 modem_port: Optional[pulumi.Input[int]] = None,
                 passwd: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure USB LTE/WIMAX devices. Applies to FortiOS Version `7.0.4`.

        ## Import

        System LteModem can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/ltemodem:Ltemodem labelname SystemLteModem
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/ltemodem:Ltemodem labelname SystemLteModem
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] apn: Login APN string for PDP-IP packet data calls.
        :param pulumi.Input[str] authtype: Authentication type for PDP-IP packet data calls. Valid values: `none`, `pap`, `chap`.
        :param pulumi.Input[str] extra_init: Extra initialization string for USB LTE/WIMAX devices.
        :param pulumi.Input[int] holddown_timer: Hold down timer (10 - 60 sec).
        :param pulumi.Input[str] interface: The interface that the modem is acting as a redundant interface for.
        :param pulumi.Input[str] mode: Modem operation mode. Valid values: `standalone`, `redundant`.
        :param pulumi.Input[int] modem_port: Modem port index (0 - 20).
        :param pulumi.Input[str] passwd: Authentication password for PDP-IP packet data calls.
        :param pulumi.Input[str] status: Enable/disable USB LTE/WIMAX device. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] username: Authentication username for PDP-IP packet data calls.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[LtemodemArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure USB LTE/WIMAX devices. Applies to FortiOS Version `7.0.4`.

        ## Import

        System LteModem can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/ltemodem:Ltemodem labelname SystemLteModem
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/ltemodem:Ltemodem labelname SystemLteModem
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param LtemodemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LtemodemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apn: Optional[pulumi.Input[str]] = None,
                 authtype: Optional[pulumi.Input[str]] = None,
                 extra_init: Optional[pulumi.Input[str]] = None,
                 holddown_timer: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 modem_port: Optional[pulumi.Input[int]] = None,
                 passwd: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LtemodemArgs.__new__(LtemodemArgs)

            __props__.__dict__["apn"] = apn
            __props__.__dict__["authtype"] = authtype
            __props__.__dict__["extra_init"] = extra_init
            __props__.__dict__["holddown_timer"] = holddown_timer
            __props__.__dict__["interface"] = interface
            __props__.__dict__["mode"] = mode
            __props__.__dict__["modem_port"] = modem_port
            __props__.__dict__["passwd"] = passwd
            __props__.__dict__["status"] = status
            __props__.__dict__["username"] = username
            __props__.__dict__["vdomparam"] = vdomparam
        super(Ltemodem, __self__).__init__(
            'fortios:system/ltemodem:Ltemodem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            apn: Optional[pulumi.Input[str]] = None,
            authtype: Optional[pulumi.Input[str]] = None,
            extra_init: Optional[pulumi.Input[str]] = None,
            holddown_timer: Optional[pulumi.Input[int]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            modem_port: Optional[pulumi.Input[int]] = None,
            passwd: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Ltemodem':
        """
        Get an existing Ltemodem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] apn: Login APN string for PDP-IP packet data calls.
        :param pulumi.Input[str] authtype: Authentication type for PDP-IP packet data calls. Valid values: `none`, `pap`, `chap`.
        :param pulumi.Input[str] extra_init: Extra initialization string for USB LTE/WIMAX devices.
        :param pulumi.Input[int] holddown_timer: Hold down timer (10 - 60 sec).
        :param pulumi.Input[str] interface: The interface that the modem is acting as a redundant interface for.
        :param pulumi.Input[str] mode: Modem operation mode. Valid values: `standalone`, `redundant`.
        :param pulumi.Input[int] modem_port: Modem port index (0 - 20).
        :param pulumi.Input[str] passwd: Authentication password for PDP-IP packet data calls.
        :param pulumi.Input[str] status: Enable/disable USB LTE/WIMAX device. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] username: Authentication username for PDP-IP packet data calls.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LtemodemState.__new__(_LtemodemState)

        __props__.__dict__["apn"] = apn
        __props__.__dict__["authtype"] = authtype
        __props__.__dict__["extra_init"] = extra_init
        __props__.__dict__["holddown_timer"] = holddown_timer
        __props__.__dict__["interface"] = interface
        __props__.__dict__["mode"] = mode
        __props__.__dict__["modem_port"] = modem_port
        __props__.__dict__["passwd"] = passwd
        __props__.__dict__["status"] = status
        __props__.__dict__["username"] = username
        __props__.__dict__["vdomparam"] = vdomparam
        return Ltemodem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def apn(self) -> pulumi.Output[str]:
        """
        Login APN string for PDP-IP packet data calls.
        """
        return pulumi.get(self, "apn")

    @property
    @pulumi.getter
    def authtype(self) -> pulumi.Output[str]:
        """
        Authentication type for PDP-IP packet data calls. Valid values: `none`, `pap`, `chap`.
        """
        return pulumi.get(self, "authtype")

    @property
    @pulumi.getter(name="extraInit")
    def extra_init(self) -> pulumi.Output[str]:
        """
        Extra initialization string for USB LTE/WIMAX devices.
        """
        return pulumi.get(self, "extra_init")

    @property
    @pulumi.getter(name="holddownTimer")
    def holddown_timer(self) -> pulumi.Output[int]:
        """
        Hold down timer (10 - 60 sec).
        """
        return pulumi.get(self, "holddown_timer")

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Output[str]:
        """
        The interface that the modem is acting as a redundant interface for.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[str]:
        """
        Modem operation mode. Valid values: `standalone`, `redundant`.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter(name="modemPort")
    def modem_port(self) -> pulumi.Output[int]:
        """
        Modem port index (0 - 20).
        """
        return pulumi.get(self, "modem_port")

    @property
    @pulumi.getter
    def passwd(self) -> pulumi.Output[Optional[str]]:
        """
        Authentication password for PDP-IP packet data calls.
        """
        return pulumi.get(self, "passwd")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable USB LTE/WIMAX device. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        Authentication username for PDP-IP packet data calls.
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

