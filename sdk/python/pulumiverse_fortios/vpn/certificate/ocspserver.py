# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['OcspserverArgs', 'Ocspserver']

@pulumi.input_type
class OcspserverArgs:
    def __init__(__self__, *,
                 cert: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secondary_cert: Optional[pulumi.Input[str]] = None,
                 secondary_url: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 unavail_action: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Ocspserver resource.
        :param pulumi.Input[str] cert: OCSP server certificate.
        :param pulumi.Input[str] name: OCSP server entry name.
        :param pulumi.Input[str] secondary_cert: Secondary OCSP server certificate.
        :param pulumi.Input[str] secondary_url: Secondary OCSP server URL.
        :param pulumi.Input[str] source_ip: Source IP address for communications to the OCSP server.
        :param pulumi.Input[str] unavail_action: Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
        :param pulumi.Input[str] url: OCSP server URL.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if cert is not None:
            pulumi.set(__self__, "cert", cert)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if secondary_cert is not None:
            pulumi.set(__self__, "secondary_cert", secondary_cert)
        if secondary_url is not None:
            pulumi.set(__self__, "secondary_url", secondary_url)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if unavail_action is not None:
            pulumi.set(__self__, "unavail_action", unavail_action)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def cert(self) -> Optional[pulumi.Input[str]]:
        """
        OCSP server certificate.
        """
        return pulumi.get(self, "cert")

    @cert.setter
    def cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        OCSP server entry name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="secondaryCert")
    def secondary_cert(self) -> Optional[pulumi.Input[str]]:
        """
        Secondary OCSP server certificate.
        """
        return pulumi.get(self, "secondary_cert")

    @secondary_cert.setter
    def secondary_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_cert", value)

    @property
    @pulumi.getter(name="secondaryUrl")
    def secondary_url(self) -> Optional[pulumi.Input[str]]:
        """
        Secondary OCSP server URL.
        """
        return pulumi.get(self, "secondary_url")

    @secondary_url.setter
    def secondary_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_url", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Source IP address for communications to the OCSP server.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="unavailAction")
    def unavail_action(self) -> Optional[pulumi.Input[str]]:
        """
        Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
        """
        return pulumi.get(self, "unavail_action")

    @unavail_action.setter
    def unavail_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unavail_action", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        OCSP server URL.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

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
class _OcspserverState:
    def __init__(__self__, *,
                 cert: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secondary_cert: Optional[pulumi.Input[str]] = None,
                 secondary_url: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 unavail_action: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Ocspserver resources.
        :param pulumi.Input[str] cert: OCSP server certificate.
        :param pulumi.Input[str] name: OCSP server entry name.
        :param pulumi.Input[str] secondary_cert: Secondary OCSP server certificate.
        :param pulumi.Input[str] secondary_url: Secondary OCSP server URL.
        :param pulumi.Input[str] source_ip: Source IP address for communications to the OCSP server.
        :param pulumi.Input[str] unavail_action: Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
        :param pulumi.Input[str] url: OCSP server URL.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if cert is not None:
            pulumi.set(__self__, "cert", cert)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if secondary_cert is not None:
            pulumi.set(__self__, "secondary_cert", secondary_cert)
        if secondary_url is not None:
            pulumi.set(__self__, "secondary_url", secondary_url)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if unavail_action is not None:
            pulumi.set(__self__, "unavail_action", unavail_action)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def cert(self) -> Optional[pulumi.Input[str]]:
        """
        OCSP server certificate.
        """
        return pulumi.get(self, "cert")

    @cert.setter
    def cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        OCSP server entry name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="secondaryCert")
    def secondary_cert(self) -> Optional[pulumi.Input[str]]:
        """
        Secondary OCSP server certificate.
        """
        return pulumi.get(self, "secondary_cert")

    @secondary_cert.setter
    def secondary_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_cert", value)

    @property
    @pulumi.getter(name="secondaryUrl")
    def secondary_url(self) -> Optional[pulumi.Input[str]]:
        """
        Secondary OCSP server URL.
        """
        return pulumi.get(self, "secondary_url")

    @secondary_url.setter
    def secondary_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_url", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Source IP address for communications to the OCSP server.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="unavailAction")
    def unavail_action(self) -> Optional[pulumi.Input[str]]:
        """
        Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
        """
        return pulumi.get(self, "unavail_action")

    @unavail_action.setter
    def unavail_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unavail_action", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        OCSP server URL.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

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


class Ocspserver(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cert: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secondary_cert: Optional[pulumi.Input[str]] = None,
                 secondary_url: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 unavail_action: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        OCSP server configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.vpn.certificate.Ocspserver("trname",
            cert="ACCVRAIZ1",
            source_ip="0.0.0.0",
            unavail_action="revoke",
            url="www.tetserv.com")
        ```

        ## Import

        VpnCertificate OcspServer can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:vpn/certificate/ocspserver:Ocspserver labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:vpn/certificate/ocspserver:Ocspserver labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cert: OCSP server certificate.
        :param pulumi.Input[str] name: OCSP server entry name.
        :param pulumi.Input[str] secondary_cert: Secondary OCSP server certificate.
        :param pulumi.Input[str] secondary_url: Secondary OCSP server URL.
        :param pulumi.Input[str] source_ip: Source IP address for communications to the OCSP server.
        :param pulumi.Input[str] unavail_action: Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
        :param pulumi.Input[str] url: OCSP server URL.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[OcspserverArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        OCSP server configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.vpn.certificate.Ocspserver("trname",
            cert="ACCVRAIZ1",
            source_ip="0.0.0.0",
            unavail_action="revoke",
            url="www.tetserv.com")
        ```

        ## Import

        VpnCertificate OcspServer can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:vpn/certificate/ocspserver:Ocspserver labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:vpn/certificate/ocspserver:Ocspserver labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param OcspserverArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OcspserverArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cert: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secondary_cert: Optional[pulumi.Input[str]] = None,
                 secondary_url: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 unavail_action: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OcspserverArgs.__new__(OcspserverArgs)

            __props__.__dict__["cert"] = cert
            __props__.__dict__["name"] = name
            __props__.__dict__["secondary_cert"] = secondary_cert
            __props__.__dict__["secondary_url"] = secondary_url
            __props__.__dict__["source_ip"] = source_ip
            __props__.__dict__["unavail_action"] = unavail_action
            __props__.__dict__["url"] = url
            __props__.__dict__["vdomparam"] = vdomparam
        super(Ocspserver, __self__).__init__(
            'fortios:vpn/certificate/ocspserver:Ocspserver',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cert: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            secondary_cert: Optional[pulumi.Input[str]] = None,
            secondary_url: Optional[pulumi.Input[str]] = None,
            source_ip: Optional[pulumi.Input[str]] = None,
            unavail_action: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Ocspserver':
        """
        Get an existing Ocspserver resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cert: OCSP server certificate.
        :param pulumi.Input[str] name: OCSP server entry name.
        :param pulumi.Input[str] secondary_cert: Secondary OCSP server certificate.
        :param pulumi.Input[str] secondary_url: Secondary OCSP server URL.
        :param pulumi.Input[str] source_ip: Source IP address for communications to the OCSP server.
        :param pulumi.Input[str] unavail_action: Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
        :param pulumi.Input[str] url: OCSP server URL.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OcspserverState.__new__(_OcspserverState)

        __props__.__dict__["cert"] = cert
        __props__.__dict__["name"] = name
        __props__.__dict__["secondary_cert"] = secondary_cert
        __props__.__dict__["secondary_url"] = secondary_url
        __props__.__dict__["source_ip"] = source_ip
        __props__.__dict__["unavail_action"] = unavail_action
        __props__.__dict__["url"] = url
        __props__.__dict__["vdomparam"] = vdomparam
        return Ocspserver(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cert(self) -> pulumi.Output[str]:
        """
        OCSP server certificate.
        """
        return pulumi.get(self, "cert")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        OCSP server entry name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="secondaryCert")
    def secondary_cert(self) -> pulumi.Output[str]:
        """
        Secondary OCSP server certificate.
        """
        return pulumi.get(self, "secondary_cert")

    @property
    @pulumi.getter(name="secondaryUrl")
    def secondary_url(self) -> pulumi.Output[str]:
        """
        Secondary OCSP server URL.
        """
        return pulumi.get(self, "secondary_url")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> pulumi.Output[str]:
        """
        Source IP address for communications to the OCSP server.
        """
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter(name="unavailAction")
    def unavail_action(self) -> pulumi.Output[str]:
        """
        Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
        """
        return pulumi.get(self, "unavail_action")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        OCSP server URL.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[str]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

