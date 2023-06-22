# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['CertificateCrlArgs', 'CertificateCrl']

@pulumi.input_type
class CertificateCrlArgs:
    def __init__(__self__, *,
                 crl: Optional[pulumi.Input[str]] = None,
                 http_url: Optional[pulumi.Input[str]] = None,
                 last_updated: Optional[pulumi.Input[int]] = None,
                 ldap_password: Optional[pulumi.Input[str]] = None,
                 ldap_server: Optional[pulumi.Input[str]] = None,
                 ldap_username: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 range: Optional[pulumi.Input[str]] = None,
                 scep_cert: Optional[pulumi.Input[str]] = None,
                 scep_url: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 update_interval: Optional[pulumi.Input[int]] = None,
                 update_vdom: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CertificateCrl resource.
        :param pulumi.Input[str] crl: Certificate Revocation List as a PEM file.
        :param pulumi.Input[str] http_url: HTTP server URL for CRL auto-update.
        :param pulumi.Input[int] last_updated: Time at which CRL was last updated.
        :param pulumi.Input[str] ldap_password: LDAP server user password.
        :param pulumi.Input[str] ldap_server: LDAP server name for CRL auto-update.
        :param pulumi.Input[str] ldap_username: LDAP server user name.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] range: Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
        :param pulumi.Input[str] scep_cert: Local certificate for SCEP communication for CRL auto-update.
        :param pulumi.Input[str] scep_url: SCEP server URL for CRL auto-update.
        :param pulumi.Input[str] source: Certificate source type.
        :param pulumi.Input[str] source_ip: Source IP address for communications to a HTTP or SCEP CA server.
        :param pulumi.Input[int] update_interval: Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
        :param pulumi.Input[str] update_vdom: VDOM for CRL update.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if crl is not None:
            pulumi.set(__self__, "crl", crl)
        if http_url is not None:
            pulumi.set(__self__, "http_url", http_url)
        if last_updated is not None:
            pulumi.set(__self__, "last_updated", last_updated)
        if ldap_password is not None:
            pulumi.set(__self__, "ldap_password", ldap_password)
        if ldap_server is not None:
            pulumi.set(__self__, "ldap_server", ldap_server)
        if ldap_username is not None:
            pulumi.set(__self__, "ldap_username", ldap_username)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if range is not None:
            pulumi.set(__self__, "range", range)
        if scep_cert is not None:
            pulumi.set(__self__, "scep_cert", scep_cert)
        if scep_url is not None:
            pulumi.set(__self__, "scep_url", scep_url)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if update_interval is not None:
            pulumi.set(__self__, "update_interval", update_interval)
        if update_vdom is not None:
            pulumi.set(__self__, "update_vdom", update_vdom)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def crl(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate Revocation List as a PEM file.
        """
        return pulumi.get(self, "crl")

    @crl.setter
    def crl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "crl", value)

    @property
    @pulumi.getter(name="httpUrl")
    def http_url(self) -> Optional[pulumi.Input[str]]:
        """
        HTTP server URL for CRL auto-update.
        """
        return pulumi.get(self, "http_url")

    @http_url.setter
    def http_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_url", value)

    @property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> Optional[pulumi.Input[int]]:
        """
        Time at which CRL was last updated.
        """
        return pulumi.get(self, "last_updated")

    @last_updated.setter
    def last_updated(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "last_updated", value)

    @property
    @pulumi.getter(name="ldapPassword")
    def ldap_password(self) -> Optional[pulumi.Input[str]]:
        """
        LDAP server user password.
        """
        return pulumi.get(self, "ldap_password")

    @ldap_password.setter
    def ldap_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_password", value)

    @property
    @pulumi.getter(name="ldapServer")
    def ldap_server(self) -> Optional[pulumi.Input[str]]:
        """
        LDAP server name for CRL auto-update.
        """
        return pulumi.get(self, "ldap_server")

    @ldap_server.setter
    def ldap_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_server", value)

    @property
    @pulumi.getter(name="ldapUsername")
    def ldap_username(self) -> Optional[pulumi.Input[str]]:
        """
        LDAP server user name.
        """
        return pulumi.get(self, "ldap_username")

    @ldap_username.setter
    def ldap_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_username", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def range(self) -> Optional[pulumi.Input[str]]:
        """
        Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
        """
        return pulumi.get(self, "range")

    @range.setter
    def range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "range", value)

    @property
    @pulumi.getter(name="scepCert")
    def scep_cert(self) -> Optional[pulumi.Input[str]]:
        """
        Local certificate for SCEP communication for CRL auto-update.
        """
        return pulumi.get(self, "scep_cert")

    @scep_cert.setter
    def scep_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scep_cert", value)

    @property
    @pulumi.getter(name="scepUrl")
    def scep_url(self) -> Optional[pulumi.Input[str]]:
        """
        SCEP server URL for CRL auto-update.
        """
        return pulumi.get(self, "scep_url")

    @scep_url.setter
    def scep_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scep_url", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate source type.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Source IP address for communications to a HTTP or SCEP CA server.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="updateInterval")
    def update_interval(self) -> Optional[pulumi.Input[int]]:
        """
        Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
        """
        return pulumi.get(self, "update_interval")

    @update_interval.setter
    def update_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "update_interval", value)

    @property
    @pulumi.getter(name="updateVdom")
    def update_vdom(self) -> Optional[pulumi.Input[str]]:
        """
        VDOM for CRL update.
        """
        return pulumi.get(self, "update_vdom")

    @update_vdom.setter
    def update_vdom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_vdom", value)

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
class _CertificateCrlState:
    def __init__(__self__, *,
                 crl: Optional[pulumi.Input[str]] = None,
                 http_url: Optional[pulumi.Input[str]] = None,
                 last_updated: Optional[pulumi.Input[int]] = None,
                 ldap_password: Optional[pulumi.Input[str]] = None,
                 ldap_server: Optional[pulumi.Input[str]] = None,
                 ldap_username: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 range: Optional[pulumi.Input[str]] = None,
                 scep_cert: Optional[pulumi.Input[str]] = None,
                 scep_url: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 update_interval: Optional[pulumi.Input[int]] = None,
                 update_vdom: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CertificateCrl resources.
        :param pulumi.Input[str] crl: Certificate Revocation List as a PEM file.
        :param pulumi.Input[str] http_url: HTTP server URL for CRL auto-update.
        :param pulumi.Input[int] last_updated: Time at which CRL was last updated.
        :param pulumi.Input[str] ldap_password: LDAP server user password.
        :param pulumi.Input[str] ldap_server: LDAP server name for CRL auto-update.
        :param pulumi.Input[str] ldap_username: LDAP server user name.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] range: Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
        :param pulumi.Input[str] scep_cert: Local certificate for SCEP communication for CRL auto-update.
        :param pulumi.Input[str] scep_url: SCEP server URL for CRL auto-update.
        :param pulumi.Input[str] source: Certificate source type.
        :param pulumi.Input[str] source_ip: Source IP address for communications to a HTTP or SCEP CA server.
        :param pulumi.Input[int] update_interval: Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
        :param pulumi.Input[str] update_vdom: VDOM for CRL update.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if crl is not None:
            pulumi.set(__self__, "crl", crl)
        if http_url is not None:
            pulumi.set(__self__, "http_url", http_url)
        if last_updated is not None:
            pulumi.set(__self__, "last_updated", last_updated)
        if ldap_password is not None:
            pulumi.set(__self__, "ldap_password", ldap_password)
        if ldap_server is not None:
            pulumi.set(__self__, "ldap_server", ldap_server)
        if ldap_username is not None:
            pulumi.set(__self__, "ldap_username", ldap_username)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if range is not None:
            pulumi.set(__self__, "range", range)
        if scep_cert is not None:
            pulumi.set(__self__, "scep_cert", scep_cert)
        if scep_url is not None:
            pulumi.set(__self__, "scep_url", scep_url)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if update_interval is not None:
            pulumi.set(__self__, "update_interval", update_interval)
        if update_vdom is not None:
            pulumi.set(__self__, "update_vdom", update_vdom)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def crl(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate Revocation List as a PEM file.
        """
        return pulumi.get(self, "crl")

    @crl.setter
    def crl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "crl", value)

    @property
    @pulumi.getter(name="httpUrl")
    def http_url(self) -> Optional[pulumi.Input[str]]:
        """
        HTTP server URL for CRL auto-update.
        """
        return pulumi.get(self, "http_url")

    @http_url.setter
    def http_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_url", value)

    @property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> Optional[pulumi.Input[int]]:
        """
        Time at which CRL was last updated.
        """
        return pulumi.get(self, "last_updated")

    @last_updated.setter
    def last_updated(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "last_updated", value)

    @property
    @pulumi.getter(name="ldapPassword")
    def ldap_password(self) -> Optional[pulumi.Input[str]]:
        """
        LDAP server user password.
        """
        return pulumi.get(self, "ldap_password")

    @ldap_password.setter
    def ldap_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_password", value)

    @property
    @pulumi.getter(name="ldapServer")
    def ldap_server(self) -> Optional[pulumi.Input[str]]:
        """
        LDAP server name for CRL auto-update.
        """
        return pulumi.get(self, "ldap_server")

    @ldap_server.setter
    def ldap_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_server", value)

    @property
    @pulumi.getter(name="ldapUsername")
    def ldap_username(self) -> Optional[pulumi.Input[str]]:
        """
        LDAP server user name.
        """
        return pulumi.get(self, "ldap_username")

    @ldap_username.setter
    def ldap_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_username", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def range(self) -> Optional[pulumi.Input[str]]:
        """
        Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
        """
        return pulumi.get(self, "range")

    @range.setter
    def range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "range", value)

    @property
    @pulumi.getter(name="scepCert")
    def scep_cert(self) -> Optional[pulumi.Input[str]]:
        """
        Local certificate for SCEP communication for CRL auto-update.
        """
        return pulumi.get(self, "scep_cert")

    @scep_cert.setter
    def scep_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scep_cert", value)

    @property
    @pulumi.getter(name="scepUrl")
    def scep_url(self) -> Optional[pulumi.Input[str]]:
        """
        SCEP server URL for CRL auto-update.
        """
        return pulumi.get(self, "scep_url")

    @scep_url.setter
    def scep_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scep_url", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate source type.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Source IP address for communications to a HTTP or SCEP CA server.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="updateInterval")
    def update_interval(self) -> Optional[pulumi.Input[int]]:
        """
        Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
        """
        return pulumi.get(self, "update_interval")

    @update_interval.setter
    def update_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "update_interval", value)

    @property
    @pulumi.getter(name="updateVdom")
    def update_vdom(self) -> Optional[pulumi.Input[str]]:
        """
        VDOM for CRL update.
        """
        return pulumi.get(self, "update_vdom")

    @update_vdom.setter
    def update_vdom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_vdom", value)

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


class CertificateCrl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 crl: Optional[pulumi.Input[str]] = None,
                 http_url: Optional[pulumi.Input[str]] = None,
                 last_updated: Optional[pulumi.Input[int]] = None,
                 ldap_password: Optional[pulumi.Input[str]] = None,
                 ldap_server: Optional[pulumi.Input[str]] = None,
                 ldap_username: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 range: Optional[pulumi.Input[str]] = None,
                 scep_cert: Optional[pulumi.Input[str]] = None,
                 scep_url: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 update_interval: Optional[pulumi.Input[int]] = None,
                 update_vdom: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Certificate Revocation List as a PEM file.

        ## Import

        Certificate Crl can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/certificateCrl:CertificateCrl labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/certificateCrl:CertificateCrl labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] crl: Certificate Revocation List as a PEM file.
        :param pulumi.Input[str] http_url: HTTP server URL for CRL auto-update.
        :param pulumi.Input[int] last_updated: Time at which CRL was last updated.
        :param pulumi.Input[str] ldap_password: LDAP server user password.
        :param pulumi.Input[str] ldap_server: LDAP server name for CRL auto-update.
        :param pulumi.Input[str] ldap_username: LDAP server user name.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] range: Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
        :param pulumi.Input[str] scep_cert: Local certificate for SCEP communication for CRL auto-update.
        :param pulumi.Input[str] scep_url: SCEP server URL for CRL auto-update.
        :param pulumi.Input[str] source: Certificate source type.
        :param pulumi.Input[str] source_ip: Source IP address for communications to a HTTP or SCEP CA server.
        :param pulumi.Input[int] update_interval: Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
        :param pulumi.Input[str] update_vdom: VDOM for CRL update.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[CertificateCrlArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Certificate Revocation List as a PEM file.

        ## Import

        Certificate Crl can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/certificateCrl:CertificateCrl labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/certificateCrl:CertificateCrl labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param CertificateCrlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CertificateCrlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 crl: Optional[pulumi.Input[str]] = None,
                 http_url: Optional[pulumi.Input[str]] = None,
                 last_updated: Optional[pulumi.Input[int]] = None,
                 ldap_password: Optional[pulumi.Input[str]] = None,
                 ldap_server: Optional[pulumi.Input[str]] = None,
                 ldap_username: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 range: Optional[pulumi.Input[str]] = None,
                 scep_cert: Optional[pulumi.Input[str]] = None,
                 scep_url: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 update_interval: Optional[pulumi.Input[int]] = None,
                 update_vdom: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CertificateCrlArgs.__new__(CertificateCrlArgs)

            __props__.__dict__["crl"] = crl
            __props__.__dict__["http_url"] = http_url
            __props__.__dict__["last_updated"] = last_updated
            __props__.__dict__["ldap_password"] = None if ldap_password is None else pulumi.Output.secret(ldap_password)
            __props__.__dict__["ldap_server"] = ldap_server
            __props__.__dict__["ldap_username"] = ldap_username
            __props__.__dict__["name"] = name
            __props__.__dict__["range"] = range
            __props__.__dict__["scep_cert"] = scep_cert
            __props__.__dict__["scep_url"] = scep_url
            __props__.__dict__["source"] = source
            __props__.__dict__["source_ip"] = source_ip
            __props__.__dict__["update_interval"] = update_interval
            __props__.__dict__["update_vdom"] = update_vdom
            __props__.__dict__["vdomparam"] = vdomparam
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["ldapPassword"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(CertificateCrl, __self__).__init__(
            'fortios:index/certificateCrl:CertificateCrl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            crl: Optional[pulumi.Input[str]] = None,
            http_url: Optional[pulumi.Input[str]] = None,
            last_updated: Optional[pulumi.Input[int]] = None,
            ldap_password: Optional[pulumi.Input[str]] = None,
            ldap_server: Optional[pulumi.Input[str]] = None,
            ldap_username: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            range: Optional[pulumi.Input[str]] = None,
            scep_cert: Optional[pulumi.Input[str]] = None,
            scep_url: Optional[pulumi.Input[str]] = None,
            source: Optional[pulumi.Input[str]] = None,
            source_ip: Optional[pulumi.Input[str]] = None,
            update_interval: Optional[pulumi.Input[int]] = None,
            update_vdom: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'CertificateCrl':
        """
        Get an existing CertificateCrl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] crl: Certificate Revocation List as a PEM file.
        :param pulumi.Input[str] http_url: HTTP server URL for CRL auto-update.
        :param pulumi.Input[int] last_updated: Time at which CRL was last updated.
        :param pulumi.Input[str] ldap_password: LDAP server user password.
        :param pulumi.Input[str] ldap_server: LDAP server name for CRL auto-update.
        :param pulumi.Input[str] ldap_username: LDAP server user name.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] range: Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
        :param pulumi.Input[str] scep_cert: Local certificate for SCEP communication for CRL auto-update.
        :param pulumi.Input[str] scep_url: SCEP server URL for CRL auto-update.
        :param pulumi.Input[str] source: Certificate source type.
        :param pulumi.Input[str] source_ip: Source IP address for communications to a HTTP or SCEP CA server.
        :param pulumi.Input[int] update_interval: Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
        :param pulumi.Input[str] update_vdom: VDOM for CRL update.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CertificateCrlState.__new__(_CertificateCrlState)

        __props__.__dict__["crl"] = crl
        __props__.__dict__["http_url"] = http_url
        __props__.__dict__["last_updated"] = last_updated
        __props__.__dict__["ldap_password"] = ldap_password
        __props__.__dict__["ldap_server"] = ldap_server
        __props__.__dict__["ldap_username"] = ldap_username
        __props__.__dict__["name"] = name
        __props__.__dict__["range"] = range
        __props__.__dict__["scep_cert"] = scep_cert
        __props__.__dict__["scep_url"] = scep_url
        __props__.__dict__["source"] = source
        __props__.__dict__["source_ip"] = source_ip
        __props__.__dict__["update_interval"] = update_interval
        __props__.__dict__["update_vdom"] = update_vdom
        __props__.__dict__["vdomparam"] = vdomparam
        return CertificateCrl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def crl(self) -> pulumi.Output[str]:
        """
        Certificate Revocation List as a PEM file.
        """
        return pulumi.get(self, "crl")

    @property
    @pulumi.getter(name="httpUrl")
    def http_url(self) -> pulumi.Output[str]:
        """
        HTTP server URL for CRL auto-update.
        """
        return pulumi.get(self, "http_url")

    @property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> pulumi.Output[int]:
        """
        Time at which CRL was last updated.
        """
        return pulumi.get(self, "last_updated")

    @property
    @pulumi.getter(name="ldapPassword")
    def ldap_password(self) -> pulumi.Output[Optional[str]]:
        """
        LDAP server user password.
        """
        return pulumi.get(self, "ldap_password")

    @property
    @pulumi.getter(name="ldapServer")
    def ldap_server(self) -> pulumi.Output[str]:
        """
        LDAP server name for CRL auto-update.
        """
        return pulumi.get(self, "ldap_server")

    @property
    @pulumi.getter(name="ldapUsername")
    def ldap_username(self) -> pulumi.Output[str]:
        """
        LDAP server user name.
        """
        return pulumi.get(self, "ldap_username")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def range(self) -> pulumi.Output[str]:
        """
        Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
        """
        return pulumi.get(self, "range")

    @property
    @pulumi.getter(name="scepCert")
    def scep_cert(self) -> pulumi.Output[str]:
        """
        Local certificate for SCEP communication for CRL auto-update.
        """
        return pulumi.get(self, "scep_cert")

    @property
    @pulumi.getter(name="scepUrl")
    def scep_url(self) -> pulumi.Output[str]:
        """
        SCEP server URL for CRL auto-update.
        """
        return pulumi.get(self, "scep_url")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[str]:
        """
        Certificate source type.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> pulumi.Output[str]:
        """
        Source IP address for communications to a HTTP or SCEP CA server.
        """
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter(name="updateInterval")
    def update_interval(self) -> pulumi.Output[int]:
        """
        Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
        """
        return pulumi.get(self, "update_interval")

    @property
    @pulumi.getter(name="updateVdom")
    def update_vdom(self) -> pulumi.Output[str]:
        """
        VDOM for CRL update.
        """
        return pulumi.get(self, "update_vdom")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

