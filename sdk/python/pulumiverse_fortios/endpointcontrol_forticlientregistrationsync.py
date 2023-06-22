# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['EndpointcontrolForticlientregistrationsyncArgs', 'EndpointcontrolForticlientregistrationsync']

@pulumi.input_type
class EndpointcontrolForticlientregistrationsyncArgs:
    def __init__(__self__, *,
                 peer_ip: pulumi.Input[str],
                 peer_name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EndpointcontrolForticlientregistrationsync resource.
        :param pulumi.Input[str] peer_ip: IP address of the peer FortiGate for endpoint license synchronization.
        :param pulumi.Input[str] peer_name: Peer name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "peer_ip", peer_ip)
        if peer_name is not None:
            pulumi.set(__self__, "peer_name", peer_name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="peerIp")
    def peer_ip(self) -> pulumi.Input[str]:
        """
        IP address of the peer FortiGate for endpoint license synchronization.
        """
        return pulumi.get(self, "peer_ip")

    @peer_ip.setter
    def peer_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "peer_ip", value)

    @property
    @pulumi.getter(name="peerName")
    def peer_name(self) -> Optional[pulumi.Input[str]]:
        """
        Peer name.
        """
        return pulumi.get(self, "peer_name")

    @peer_name.setter
    def peer_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_name", value)

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
class _EndpointcontrolForticlientregistrationsyncState:
    def __init__(__self__, *,
                 peer_ip: Optional[pulumi.Input[str]] = None,
                 peer_name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EndpointcontrolForticlientregistrationsync resources.
        :param pulumi.Input[str] peer_ip: IP address of the peer FortiGate for endpoint license synchronization.
        :param pulumi.Input[str] peer_name: Peer name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if peer_ip is not None:
            pulumi.set(__self__, "peer_ip", peer_ip)
        if peer_name is not None:
            pulumi.set(__self__, "peer_name", peer_name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="peerIp")
    def peer_ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP address of the peer FortiGate for endpoint license synchronization.
        """
        return pulumi.get(self, "peer_ip")

    @peer_ip.setter
    def peer_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_ip", value)

    @property
    @pulumi.getter(name="peerName")
    def peer_name(self) -> Optional[pulumi.Input[str]]:
        """
        Peer name.
        """
        return pulumi.get(self, "peer_name")

    @peer_name.setter
    def peer_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_name", value)

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


class EndpointcontrolForticlientregistrationsync(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 peer_ip: Optional[pulumi.Input[str]] = None,
                 peer_name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure FortiClient registration synchronization settings. Applies to FortiOS Version `<= 6.2.0`.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.EndpointcontrolForticlientregistrationsync("trname",
            peer_ip="1.1.1.1",
            peer_name="1")
        ```

        ## Import

        EndpointControl ForticlientRegistrationSync can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/endpointcontrolForticlientregistrationsync:EndpointcontrolForticlientregistrationsync labelname {{peer_name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/endpointcontrolForticlientregistrationsync:EndpointcontrolForticlientregistrationsync labelname {{peer_name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] peer_ip: IP address of the peer FortiGate for endpoint license synchronization.
        :param pulumi.Input[str] peer_name: Peer name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EndpointcontrolForticlientregistrationsyncArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure FortiClient registration synchronization settings. Applies to FortiOS Version `<= 6.2.0`.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.EndpointcontrolForticlientregistrationsync("trname",
            peer_ip="1.1.1.1",
            peer_name="1")
        ```

        ## Import

        EndpointControl ForticlientRegistrationSync can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/endpointcontrolForticlientregistrationsync:EndpointcontrolForticlientregistrationsync labelname {{peer_name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/endpointcontrolForticlientregistrationsync:EndpointcontrolForticlientregistrationsync labelname {{peer_name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param EndpointcontrolForticlientregistrationsyncArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndpointcontrolForticlientregistrationsyncArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 peer_ip: Optional[pulumi.Input[str]] = None,
                 peer_name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EndpointcontrolForticlientregistrationsyncArgs.__new__(EndpointcontrolForticlientregistrationsyncArgs)

            if peer_ip is None and not opts.urn:
                raise TypeError("Missing required property 'peer_ip'")
            __props__.__dict__["peer_ip"] = peer_ip
            __props__.__dict__["peer_name"] = peer_name
            __props__.__dict__["vdomparam"] = vdomparam
        super(EndpointcontrolForticlientregistrationsync, __self__).__init__(
            'fortios:index/endpointcontrolForticlientregistrationsync:EndpointcontrolForticlientregistrationsync',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            peer_ip: Optional[pulumi.Input[str]] = None,
            peer_name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'EndpointcontrolForticlientregistrationsync':
        """
        Get an existing EndpointcontrolForticlientregistrationsync resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] peer_ip: IP address of the peer FortiGate for endpoint license synchronization.
        :param pulumi.Input[str] peer_name: Peer name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EndpointcontrolForticlientregistrationsyncState.__new__(_EndpointcontrolForticlientregistrationsyncState)

        __props__.__dict__["peer_ip"] = peer_ip
        __props__.__dict__["peer_name"] = peer_name
        __props__.__dict__["vdomparam"] = vdomparam
        return EndpointcontrolForticlientregistrationsync(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="peerIp")
    def peer_ip(self) -> pulumi.Output[str]:
        """
        IP address of the peer FortiGate for endpoint license synchronization.
        """
        return pulumi.get(self, "peer_ip")

    @property
    @pulumi.getter(name="peerName")
    def peer_name(self) -> pulumi.Output[str]:
        """
        Peer name.
        """
        return pulumi.get(self, "peer_name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

