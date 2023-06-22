# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetSystemCsfResult',
    'AwaitableGetSystemCsfResult',
    'get_system_csf',
    'get_system_csf_output',
]

@pulumi.output_type
class GetSystemCsfResult:
    """
    A collection of values returned by getSystemCsf.
    """
    def __init__(__self__, accept_auth_by_cert=None, authorization_request_type=None, certificate=None, configuration_sync=None, downstream_access=None, downstream_accprofile=None, fabric_connectors=None, fabric_devices=None, fabric_object_unification=None, fabric_workers=None, fixed_key=None, forticloud_account_enforcement=None, group_name=None, group_password=None, id=None, log_unification=None, management_ip=None, management_port=None, saml_configuration_sync=None, status=None, trusted_lists=None, upstream=None, upstream_ip=None, upstream_port=None, vdomparam=None):
        if accept_auth_by_cert and not isinstance(accept_auth_by_cert, str):
            raise TypeError("Expected argument 'accept_auth_by_cert' to be a str")
        pulumi.set(__self__, "accept_auth_by_cert", accept_auth_by_cert)
        if authorization_request_type and not isinstance(authorization_request_type, str):
            raise TypeError("Expected argument 'authorization_request_type' to be a str")
        pulumi.set(__self__, "authorization_request_type", authorization_request_type)
        if certificate and not isinstance(certificate, str):
            raise TypeError("Expected argument 'certificate' to be a str")
        pulumi.set(__self__, "certificate", certificate)
        if configuration_sync and not isinstance(configuration_sync, str):
            raise TypeError("Expected argument 'configuration_sync' to be a str")
        pulumi.set(__self__, "configuration_sync", configuration_sync)
        if downstream_access and not isinstance(downstream_access, str):
            raise TypeError("Expected argument 'downstream_access' to be a str")
        pulumi.set(__self__, "downstream_access", downstream_access)
        if downstream_accprofile and not isinstance(downstream_accprofile, str):
            raise TypeError("Expected argument 'downstream_accprofile' to be a str")
        pulumi.set(__self__, "downstream_accprofile", downstream_accprofile)
        if fabric_connectors and not isinstance(fabric_connectors, list):
            raise TypeError("Expected argument 'fabric_connectors' to be a list")
        pulumi.set(__self__, "fabric_connectors", fabric_connectors)
        if fabric_devices and not isinstance(fabric_devices, list):
            raise TypeError("Expected argument 'fabric_devices' to be a list")
        pulumi.set(__self__, "fabric_devices", fabric_devices)
        if fabric_object_unification and not isinstance(fabric_object_unification, str):
            raise TypeError("Expected argument 'fabric_object_unification' to be a str")
        pulumi.set(__self__, "fabric_object_unification", fabric_object_unification)
        if fabric_workers and not isinstance(fabric_workers, int):
            raise TypeError("Expected argument 'fabric_workers' to be a int")
        pulumi.set(__self__, "fabric_workers", fabric_workers)
        if fixed_key and not isinstance(fixed_key, str):
            raise TypeError("Expected argument 'fixed_key' to be a str")
        pulumi.set(__self__, "fixed_key", fixed_key)
        if forticloud_account_enforcement and not isinstance(forticloud_account_enforcement, str):
            raise TypeError("Expected argument 'forticloud_account_enforcement' to be a str")
        pulumi.set(__self__, "forticloud_account_enforcement", forticloud_account_enforcement)
        if group_name and not isinstance(group_name, str):
            raise TypeError("Expected argument 'group_name' to be a str")
        pulumi.set(__self__, "group_name", group_name)
        if group_password and not isinstance(group_password, str):
            raise TypeError("Expected argument 'group_password' to be a str")
        pulumi.set(__self__, "group_password", group_password)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if log_unification and not isinstance(log_unification, str):
            raise TypeError("Expected argument 'log_unification' to be a str")
        pulumi.set(__self__, "log_unification", log_unification)
        if management_ip and not isinstance(management_ip, str):
            raise TypeError("Expected argument 'management_ip' to be a str")
        pulumi.set(__self__, "management_ip", management_ip)
        if management_port and not isinstance(management_port, int):
            raise TypeError("Expected argument 'management_port' to be a int")
        pulumi.set(__self__, "management_port", management_port)
        if saml_configuration_sync and not isinstance(saml_configuration_sync, str):
            raise TypeError("Expected argument 'saml_configuration_sync' to be a str")
        pulumi.set(__self__, "saml_configuration_sync", saml_configuration_sync)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if trusted_lists and not isinstance(trusted_lists, list):
            raise TypeError("Expected argument 'trusted_lists' to be a list")
        pulumi.set(__self__, "trusted_lists", trusted_lists)
        if upstream and not isinstance(upstream, str):
            raise TypeError("Expected argument 'upstream' to be a str")
        pulumi.set(__self__, "upstream", upstream)
        if upstream_ip and not isinstance(upstream_ip, str):
            raise TypeError("Expected argument 'upstream_ip' to be a str")
        pulumi.set(__self__, "upstream_ip", upstream_ip)
        if upstream_port and not isinstance(upstream_port, int):
            raise TypeError("Expected argument 'upstream_port' to be a int")
        pulumi.set(__self__, "upstream_port", upstream_port)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="acceptAuthByCert")
    def accept_auth_by_cert(self) -> str:
        """
        Accept connections with unknown certificates and ask admin for approval.
        """
        return pulumi.get(self, "accept_auth_by_cert")

    @property
    @pulumi.getter(name="authorizationRequestType")
    def authorization_request_type(self) -> str:
        """
        Authorization request type.
        """
        return pulumi.get(self, "authorization_request_type")

    @property
    @pulumi.getter
    def certificate(self) -> str:
        """
        Certificate.
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="configurationSync")
    def configuration_sync(self) -> str:
        """
        Configuration sync mode.
        """
        return pulumi.get(self, "configuration_sync")

    @property
    @pulumi.getter(name="downstreamAccess")
    def downstream_access(self) -> str:
        """
        Enable/disable downstream device access to this device's configuration and data.
        """
        return pulumi.get(self, "downstream_access")

    @property
    @pulumi.getter(name="downstreamAccprofile")
    def downstream_accprofile(self) -> str:
        """
        Default access profile for requests from downstream devices.
        """
        return pulumi.get(self, "downstream_accprofile")

    @property
    @pulumi.getter(name="fabricConnectors")
    def fabric_connectors(self) -> Sequence['outputs.GetSystemCsfFabricConnectorResult']:
        """
        Fabric connector configuration. The structure of `fabric_connector` block is documented below.
        """
        return pulumi.get(self, "fabric_connectors")

    @property
    @pulumi.getter(name="fabricDevices")
    def fabric_devices(self) -> Sequence['outputs.GetSystemCsfFabricDeviceResult']:
        """
        Fabric device configuration. The structure of `fabric_device` block is documented below.
        """
        return pulumi.get(self, "fabric_devices")

    @property
    @pulumi.getter(name="fabricObjectUnification")
    def fabric_object_unification(self) -> str:
        """
        Fabric CMDB Object Unification
        """
        return pulumi.get(self, "fabric_object_unification")

    @property
    @pulumi.getter(name="fabricWorkers")
    def fabric_workers(self) -> int:
        """
        Number of worker processes for Security Fabric daemon.
        """
        return pulumi.get(self, "fabric_workers")

    @property
    @pulumi.getter(name="fixedKey")
    def fixed_key(self) -> str:
        """
        Auto-generated fixed key used when this device is the root. (Will automatically be generated if not set.)
        """
        return pulumi.get(self, "fixed_key")

    @property
    @pulumi.getter(name="forticloudAccountEnforcement")
    def forticloud_account_enforcement(self) -> str:
        """
        Fabric FortiCloud account unification.
        """
        return pulumi.get(self, "forticloud_account_enforcement")

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> str:
        """
        Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter(name="groupPassword")
    def group_password(self) -> str:
        """
        Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
        """
        return pulumi.get(self, "group_password")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="logUnification")
    def log_unification(self) -> str:
        """
        Enable/disable broadcast of discovery messages for log unification.
        """
        return pulumi.get(self, "log_unification")

    @property
    @pulumi.getter(name="managementIp")
    def management_ip(self) -> str:
        """
        Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
        """
        return pulumi.get(self, "management_ip")

    @property
    @pulumi.getter(name="managementPort")
    def management_port(self) -> int:
        """
        Overriding port for management connection (Overrides admin port).
        """
        return pulumi.get(self, "management_port")

    @property
    @pulumi.getter(name="samlConfigurationSync")
    def saml_configuration_sync(self) -> str:
        """
        SAML setting configuration synchronization.
        """
        return pulumi.get(self, "saml_configuration_sync")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Enable/disable Security Fabric.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="trustedLists")
    def trusted_lists(self) -> Sequence['outputs.GetSystemCsfTrustedListResult']:
        """
        Pre-authorized and blocked security fabric nodes. The structure of `trusted_list` block is documented below.
        """
        return pulumi.get(self, "trusted_lists")

    @property
    @pulumi.getter
    def upstream(self) -> str:
        """
        IP/FQDN of the FortiGate upstream from this FortiGate in the Security Fabric.
        """
        return pulumi.get(self, "upstream")

    @property
    @pulumi.getter(name="upstreamIp")
    def upstream_ip(self) -> str:
        """
        IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
        """
        return pulumi.get(self, "upstream_ip")

    @property
    @pulumi.getter(name="upstreamPort")
    def upstream_port(self) -> int:
        """
        The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
        """
        return pulumi.get(self, "upstream_port")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemCsfResult(GetSystemCsfResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemCsfResult(
            accept_auth_by_cert=self.accept_auth_by_cert,
            authorization_request_type=self.authorization_request_type,
            certificate=self.certificate,
            configuration_sync=self.configuration_sync,
            downstream_access=self.downstream_access,
            downstream_accprofile=self.downstream_accprofile,
            fabric_connectors=self.fabric_connectors,
            fabric_devices=self.fabric_devices,
            fabric_object_unification=self.fabric_object_unification,
            fabric_workers=self.fabric_workers,
            fixed_key=self.fixed_key,
            forticloud_account_enforcement=self.forticloud_account_enforcement,
            group_name=self.group_name,
            group_password=self.group_password,
            id=self.id,
            log_unification=self.log_unification,
            management_ip=self.management_ip,
            management_port=self.management_port,
            saml_configuration_sync=self.saml_configuration_sync,
            status=self.status,
            trusted_lists=self.trusted_lists,
            upstream=self.upstream,
            upstream_ip=self.upstream_ip,
            upstream_port=self.upstream_port,
            vdomparam=self.vdomparam)


def get_system_csf(vdomparam: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemCsfResult:
    """
    Use this data source to get information on fortios system csf


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemCsf:getSystemCsf', __args__, opts=opts, typ=GetSystemCsfResult).value

    return AwaitableGetSystemCsfResult(
        accept_auth_by_cert=__ret__.accept_auth_by_cert,
        authorization_request_type=__ret__.authorization_request_type,
        certificate=__ret__.certificate,
        configuration_sync=__ret__.configuration_sync,
        downstream_access=__ret__.downstream_access,
        downstream_accprofile=__ret__.downstream_accprofile,
        fabric_connectors=__ret__.fabric_connectors,
        fabric_devices=__ret__.fabric_devices,
        fabric_object_unification=__ret__.fabric_object_unification,
        fabric_workers=__ret__.fabric_workers,
        fixed_key=__ret__.fixed_key,
        forticloud_account_enforcement=__ret__.forticloud_account_enforcement,
        group_name=__ret__.group_name,
        group_password=__ret__.group_password,
        id=__ret__.id,
        log_unification=__ret__.log_unification,
        management_ip=__ret__.management_ip,
        management_port=__ret__.management_port,
        saml_configuration_sync=__ret__.saml_configuration_sync,
        status=__ret__.status,
        trusted_lists=__ret__.trusted_lists,
        upstream=__ret__.upstream,
        upstream_ip=__ret__.upstream_ip,
        upstream_port=__ret__.upstream_port,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_system_csf)
def get_system_csf_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemCsfResult]:
    """
    Use this data source to get information on fortios system csf


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
