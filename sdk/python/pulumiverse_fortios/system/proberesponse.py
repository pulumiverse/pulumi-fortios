# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ProberesponseArgs', 'Proberesponse']

@pulumi.input_type
class ProberesponseArgs:
    def __init__(__self__, *,
                 http_probe_value: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 security_mode: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 ttl_mode: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Proberesponse resource.
        :param pulumi.Input[str] http_probe_value: Value to respond to the monitoring server.
        :param pulumi.Input[str] mode: SLA response mode. Valid values: `none`, `http-probe`, `twamp`.
        :param pulumi.Input[str] password: Twamp respondor password in authentication mode
        :param pulumi.Input[int] port: Port number to response.
        :param pulumi.Input[str] security_mode: Twamp respondor security mode. Valid values: `none`, `authentication`.
        :param pulumi.Input[int] timeout: An inactivity timer for a twamp test session.
        :param pulumi.Input[str] ttl_mode: Mode for TWAMP packet TTL modification. Valid values: `reinit`, `decrease`, `retain`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if http_probe_value is not None:
            pulumi.set(__self__, "http_probe_value", http_probe_value)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if security_mode is not None:
            pulumi.set(__self__, "security_mode", security_mode)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if ttl_mode is not None:
            pulumi.set(__self__, "ttl_mode", ttl_mode)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="httpProbeValue")
    def http_probe_value(self) -> Optional[pulumi.Input[str]]:
        """
        Value to respond to the monitoring server.
        """
        return pulumi.get(self, "http_probe_value")

    @http_probe_value.setter
    def http_probe_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_probe_value", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        SLA response mode. Valid values: `none`, `http-probe`, `twamp`.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Twamp respondor password in authentication mode
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Port number to response.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="securityMode")
    def security_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Twamp respondor security mode. Valid values: `none`, `authentication`.
        """
        return pulumi.get(self, "security_mode")

    @security_mode.setter
    def security_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_mode", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        An inactivity timer for a twamp test session.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter(name="ttlMode")
    def ttl_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Mode for TWAMP packet TTL modification. Valid values: `reinit`, `decrease`, `retain`.
        """
        return pulumi.get(self, "ttl_mode")

    @ttl_mode.setter
    def ttl_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ttl_mode", value)

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
class _ProberesponseState:
    def __init__(__self__, *,
                 http_probe_value: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 security_mode: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 ttl_mode: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Proberesponse resources.
        :param pulumi.Input[str] http_probe_value: Value to respond to the monitoring server.
        :param pulumi.Input[str] mode: SLA response mode. Valid values: `none`, `http-probe`, `twamp`.
        :param pulumi.Input[str] password: Twamp respondor password in authentication mode
        :param pulumi.Input[int] port: Port number to response.
        :param pulumi.Input[str] security_mode: Twamp respondor security mode. Valid values: `none`, `authentication`.
        :param pulumi.Input[int] timeout: An inactivity timer for a twamp test session.
        :param pulumi.Input[str] ttl_mode: Mode for TWAMP packet TTL modification. Valid values: `reinit`, `decrease`, `retain`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if http_probe_value is not None:
            pulumi.set(__self__, "http_probe_value", http_probe_value)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if security_mode is not None:
            pulumi.set(__self__, "security_mode", security_mode)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if ttl_mode is not None:
            pulumi.set(__self__, "ttl_mode", ttl_mode)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="httpProbeValue")
    def http_probe_value(self) -> Optional[pulumi.Input[str]]:
        """
        Value to respond to the monitoring server.
        """
        return pulumi.get(self, "http_probe_value")

    @http_probe_value.setter
    def http_probe_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_probe_value", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        SLA response mode. Valid values: `none`, `http-probe`, `twamp`.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Twamp respondor password in authentication mode
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Port number to response.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="securityMode")
    def security_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Twamp respondor security mode. Valid values: `none`, `authentication`.
        """
        return pulumi.get(self, "security_mode")

    @security_mode.setter
    def security_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_mode", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        An inactivity timer for a twamp test session.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter(name="ttlMode")
    def ttl_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Mode for TWAMP packet TTL modification. Valid values: `reinit`, `decrease`, `retain`.
        """
        return pulumi.get(self, "ttl_mode")

    @ttl_mode.setter
    def ttl_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ttl_mode", value)

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


class Proberesponse(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 http_probe_value: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 security_mode: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 ttl_mode: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure system probe response.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.system.Proberesponse("trname",
            http_probe_value="OK",
            mode="none",
            port=8008,
            security_mode="none",
            timeout=300,
            ttl_mode="retain")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        System ProbeResponse can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/proberesponse:Proberesponse labelname SystemProbeResponse
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/proberesponse:Proberesponse labelname SystemProbeResponse
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] http_probe_value: Value to respond to the monitoring server.
        :param pulumi.Input[str] mode: SLA response mode. Valid values: `none`, `http-probe`, `twamp`.
        :param pulumi.Input[str] password: Twamp respondor password in authentication mode
        :param pulumi.Input[int] port: Port number to response.
        :param pulumi.Input[str] security_mode: Twamp respondor security mode. Valid values: `none`, `authentication`.
        :param pulumi.Input[int] timeout: An inactivity timer for a twamp test session.
        :param pulumi.Input[str] ttl_mode: Mode for TWAMP packet TTL modification. Valid values: `reinit`, `decrease`, `retain`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProberesponseArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure system probe response.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.system.Proberesponse("trname",
            http_probe_value="OK",
            mode="none",
            port=8008,
            security_mode="none",
            timeout=300,
            ttl_mode="retain")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        System ProbeResponse can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/proberesponse:Proberesponse labelname SystemProbeResponse
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/proberesponse:Proberesponse labelname SystemProbeResponse
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param ProberesponseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProberesponseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 http_probe_value: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 security_mode: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 ttl_mode: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProberesponseArgs.__new__(ProberesponseArgs)

            __props__.__dict__["http_probe_value"] = http_probe_value
            __props__.__dict__["mode"] = mode
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["port"] = port
            __props__.__dict__["security_mode"] = security_mode
            __props__.__dict__["timeout"] = timeout
            __props__.__dict__["ttl_mode"] = ttl_mode
            __props__.__dict__["vdomparam"] = vdomparam
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Proberesponse, __self__).__init__(
            'fortios:system/proberesponse:Proberesponse',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            http_probe_value: Optional[pulumi.Input[str]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            security_mode: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[int]] = None,
            ttl_mode: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Proberesponse':
        """
        Get an existing Proberesponse resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] http_probe_value: Value to respond to the monitoring server.
        :param pulumi.Input[str] mode: SLA response mode. Valid values: `none`, `http-probe`, `twamp`.
        :param pulumi.Input[str] password: Twamp respondor password in authentication mode
        :param pulumi.Input[int] port: Port number to response.
        :param pulumi.Input[str] security_mode: Twamp respondor security mode. Valid values: `none`, `authentication`.
        :param pulumi.Input[int] timeout: An inactivity timer for a twamp test session.
        :param pulumi.Input[str] ttl_mode: Mode for TWAMP packet TTL modification. Valid values: `reinit`, `decrease`, `retain`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProberesponseState.__new__(_ProberesponseState)

        __props__.__dict__["http_probe_value"] = http_probe_value
        __props__.__dict__["mode"] = mode
        __props__.__dict__["password"] = password
        __props__.__dict__["port"] = port
        __props__.__dict__["security_mode"] = security_mode
        __props__.__dict__["timeout"] = timeout
        __props__.__dict__["ttl_mode"] = ttl_mode
        __props__.__dict__["vdomparam"] = vdomparam
        return Proberesponse(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="httpProbeValue")
    def http_probe_value(self) -> pulumi.Output[str]:
        """
        Value to respond to the monitoring server.
        """
        return pulumi.get(self, "http_probe_value")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[str]:
        """
        SLA response mode. Valid values: `none`, `http-probe`, `twamp`.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        Twamp respondor password in authentication mode
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        Port number to response.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="securityMode")
    def security_mode(self) -> pulumi.Output[str]:
        """
        Twamp respondor security mode. Valid values: `none`, `authentication`.
        """
        return pulumi.get(self, "security_mode")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[int]:
        """
        An inactivity timer for a twamp test session.
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter(name="ttlMode")
    def ttl_mode(self) -> pulumi.Output[str]:
        """
        Mode for TWAMP packet TTL modification. Valid values: `reinit`, `decrease`, `retain`.
        """
        return pulumi.get(self, "ttl_mode")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

