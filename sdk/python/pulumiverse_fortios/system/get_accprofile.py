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

__all__ = [
    'GetAccprofileResult',
    'AwaitableGetAccprofileResult',
    'get_accprofile',
    'get_accprofile_output',
]

@pulumi.output_type
class GetAccprofileResult:
    """
    A collection of values returned by getAccprofile.
    """
    def __init__(__self__, admintimeout=None, admintimeout_override=None, authgrp=None, cli_config=None, cli_diagnose=None, cli_exec=None, cli_get=None, cli_show=None, comments=None, ftviewgrp=None, fwgrp=None, fwgrp_permissions=None, id=None, loggrp=None, loggrp_permissions=None, name=None, netgrp=None, netgrp_permissions=None, scope=None, secfabgrp=None, sysgrp=None, sysgrp_permissions=None, system_diagnostics=None, system_execute_ssh=None, system_execute_telnet=None, utmgrp=None, utmgrp_permissions=None, vdomparam=None, vpngrp=None, wanoptgrp=None, wifi=None):
        if admintimeout and not isinstance(admintimeout, int):
            raise TypeError("Expected argument 'admintimeout' to be a int")
        pulumi.set(__self__, "admintimeout", admintimeout)
        if admintimeout_override and not isinstance(admintimeout_override, str):
            raise TypeError("Expected argument 'admintimeout_override' to be a str")
        pulumi.set(__self__, "admintimeout_override", admintimeout_override)
        if authgrp and not isinstance(authgrp, str):
            raise TypeError("Expected argument 'authgrp' to be a str")
        pulumi.set(__self__, "authgrp", authgrp)
        if cli_config and not isinstance(cli_config, str):
            raise TypeError("Expected argument 'cli_config' to be a str")
        pulumi.set(__self__, "cli_config", cli_config)
        if cli_diagnose and not isinstance(cli_diagnose, str):
            raise TypeError("Expected argument 'cli_diagnose' to be a str")
        pulumi.set(__self__, "cli_diagnose", cli_diagnose)
        if cli_exec and not isinstance(cli_exec, str):
            raise TypeError("Expected argument 'cli_exec' to be a str")
        pulumi.set(__self__, "cli_exec", cli_exec)
        if cli_get and not isinstance(cli_get, str):
            raise TypeError("Expected argument 'cli_get' to be a str")
        pulumi.set(__self__, "cli_get", cli_get)
        if cli_show and not isinstance(cli_show, str):
            raise TypeError("Expected argument 'cli_show' to be a str")
        pulumi.set(__self__, "cli_show", cli_show)
        if comments and not isinstance(comments, str):
            raise TypeError("Expected argument 'comments' to be a str")
        pulumi.set(__self__, "comments", comments)
        if ftviewgrp and not isinstance(ftviewgrp, str):
            raise TypeError("Expected argument 'ftviewgrp' to be a str")
        pulumi.set(__self__, "ftviewgrp", ftviewgrp)
        if fwgrp and not isinstance(fwgrp, str):
            raise TypeError("Expected argument 'fwgrp' to be a str")
        pulumi.set(__self__, "fwgrp", fwgrp)
        if fwgrp_permissions and not isinstance(fwgrp_permissions, list):
            raise TypeError("Expected argument 'fwgrp_permissions' to be a list")
        pulumi.set(__self__, "fwgrp_permissions", fwgrp_permissions)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if loggrp and not isinstance(loggrp, str):
            raise TypeError("Expected argument 'loggrp' to be a str")
        pulumi.set(__self__, "loggrp", loggrp)
        if loggrp_permissions and not isinstance(loggrp_permissions, list):
            raise TypeError("Expected argument 'loggrp_permissions' to be a list")
        pulumi.set(__self__, "loggrp_permissions", loggrp_permissions)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if netgrp and not isinstance(netgrp, str):
            raise TypeError("Expected argument 'netgrp' to be a str")
        pulumi.set(__self__, "netgrp", netgrp)
        if netgrp_permissions and not isinstance(netgrp_permissions, list):
            raise TypeError("Expected argument 'netgrp_permissions' to be a list")
        pulumi.set(__self__, "netgrp_permissions", netgrp_permissions)
        if scope and not isinstance(scope, str):
            raise TypeError("Expected argument 'scope' to be a str")
        pulumi.set(__self__, "scope", scope)
        if secfabgrp and not isinstance(secfabgrp, str):
            raise TypeError("Expected argument 'secfabgrp' to be a str")
        pulumi.set(__self__, "secfabgrp", secfabgrp)
        if sysgrp and not isinstance(sysgrp, str):
            raise TypeError("Expected argument 'sysgrp' to be a str")
        pulumi.set(__self__, "sysgrp", sysgrp)
        if sysgrp_permissions and not isinstance(sysgrp_permissions, list):
            raise TypeError("Expected argument 'sysgrp_permissions' to be a list")
        pulumi.set(__self__, "sysgrp_permissions", sysgrp_permissions)
        if system_diagnostics and not isinstance(system_diagnostics, str):
            raise TypeError("Expected argument 'system_diagnostics' to be a str")
        pulumi.set(__self__, "system_diagnostics", system_diagnostics)
        if system_execute_ssh and not isinstance(system_execute_ssh, str):
            raise TypeError("Expected argument 'system_execute_ssh' to be a str")
        pulumi.set(__self__, "system_execute_ssh", system_execute_ssh)
        if system_execute_telnet and not isinstance(system_execute_telnet, str):
            raise TypeError("Expected argument 'system_execute_telnet' to be a str")
        pulumi.set(__self__, "system_execute_telnet", system_execute_telnet)
        if utmgrp and not isinstance(utmgrp, str):
            raise TypeError("Expected argument 'utmgrp' to be a str")
        pulumi.set(__self__, "utmgrp", utmgrp)
        if utmgrp_permissions and not isinstance(utmgrp_permissions, list):
            raise TypeError("Expected argument 'utmgrp_permissions' to be a list")
        pulumi.set(__self__, "utmgrp_permissions", utmgrp_permissions)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)
        if vpngrp and not isinstance(vpngrp, str):
            raise TypeError("Expected argument 'vpngrp' to be a str")
        pulumi.set(__self__, "vpngrp", vpngrp)
        if wanoptgrp and not isinstance(wanoptgrp, str):
            raise TypeError("Expected argument 'wanoptgrp' to be a str")
        pulumi.set(__self__, "wanoptgrp", wanoptgrp)
        if wifi and not isinstance(wifi, str):
            raise TypeError("Expected argument 'wifi' to be a str")
        pulumi.set(__self__, "wifi", wifi)

    @property
    @pulumi.getter
    def admintimeout(self) -> int:
        """
        Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
        """
        return pulumi.get(self, "admintimeout")

    @property
    @pulumi.getter(name="admintimeoutOverride")
    def admintimeout_override(self) -> str:
        """
        Enable/disable overriding the global administrator idle timeout.
        """
        return pulumi.get(self, "admintimeout_override")

    @property
    @pulumi.getter
    def authgrp(self) -> str:
        """
        Administrator access to Users and Devices.
        """
        return pulumi.get(self, "authgrp")

    @property
    @pulumi.getter(name="cliConfig")
    def cli_config(self) -> str:
        """
        Enable/disable permission to run config commands.
        """
        return pulumi.get(self, "cli_config")

    @property
    @pulumi.getter(name="cliDiagnose")
    def cli_diagnose(self) -> str:
        """
        Enable/disable permission to run diagnostic commands.
        """
        return pulumi.get(self, "cli_diagnose")

    @property
    @pulumi.getter(name="cliExec")
    def cli_exec(self) -> str:
        """
        Enable/disable permission to run execute commands.
        """
        return pulumi.get(self, "cli_exec")

    @property
    @pulumi.getter(name="cliGet")
    def cli_get(self) -> str:
        """
        Enable/disable permission to run get commands.
        """
        return pulumi.get(self, "cli_get")

    @property
    @pulumi.getter(name="cliShow")
    def cli_show(self) -> str:
        """
        Enable/disable permission to run show commands.
        """
        return pulumi.get(self, "cli_show")

    @property
    @pulumi.getter
    def comments(self) -> str:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def ftviewgrp(self) -> str:
        """
        FortiView.
        """
        return pulumi.get(self, "ftviewgrp")

    @property
    @pulumi.getter
    def fwgrp(self) -> str:
        """
        Administrator access to the Firewall configuration.
        """
        return pulumi.get(self, "fwgrp")

    @property
    @pulumi.getter(name="fwgrpPermissions")
    def fwgrp_permissions(self) -> Sequence['outputs.GetAccprofileFwgrpPermissionResult']:
        """
        Custom firewall permission. The structure of `fwgrp_permission` block is documented below.
        """
        return pulumi.get(self, "fwgrp_permissions")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def loggrp(self) -> str:
        """
        Administrator access to Logging and Reporting including viewing log messages.
        """
        return pulumi.get(self, "loggrp")

    @property
    @pulumi.getter(name="loggrpPermissions")
    def loggrp_permissions(self) -> Sequence['outputs.GetAccprofileLoggrpPermissionResult']:
        """
        Custom Log & Report permission. The structure of `loggrp_permission` block is documented below.
        """
        return pulumi.get(self, "loggrp_permissions")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Profile name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def netgrp(self) -> str:
        """
        Network Configuration.
        """
        return pulumi.get(self, "netgrp")

    @property
    @pulumi.getter(name="netgrpPermissions")
    def netgrp_permissions(self) -> Sequence['outputs.GetAccprofileNetgrpPermissionResult']:
        """
        Custom network permission. The structure of `netgrp_permission` block is documented below.
        """
        return pulumi.get(self, "netgrp_permissions")

    @property
    @pulumi.getter
    def scope(self) -> str:
        """
        Scope of admin access: global or specific VDOM(s).
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def secfabgrp(self) -> str:
        """
        Security Fabric.
        """
        return pulumi.get(self, "secfabgrp")

    @property
    @pulumi.getter
    def sysgrp(self) -> str:
        """
        System Configuration.
        """
        return pulumi.get(self, "sysgrp")

    @property
    @pulumi.getter(name="sysgrpPermissions")
    def sysgrp_permissions(self) -> Sequence['outputs.GetAccprofileSysgrpPermissionResult']:
        """
        Custom system permission. The structure of `sysgrp_permission` block is documented below.
        """
        return pulumi.get(self, "sysgrp_permissions")

    @property
    @pulumi.getter(name="systemDiagnostics")
    def system_diagnostics(self) -> str:
        """
        Enable/disable permission to run system diagnostic commands.
        """
        return pulumi.get(self, "system_diagnostics")

    @property
    @pulumi.getter(name="systemExecuteSsh")
    def system_execute_ssh(self) -> str:
        """
        Enable/disable permission to execute SSH commands.
        """
        return pulumi.get(self, "system_execute_ssh")

    @property
    @pulumi.getter(name="systemExecuteTelnet")
    def system_execute_telnet(self) -> str:
        """
        Enable/disable permission to execute TELNET commands.
        """
        return pulumi.get(self, "system_execute_telnet")

    @property
    @pulumi.getter
    def utmgrp(self) -> str:
        """
        Administrator access to Security Profiles.
        """
        return pulumi.get(self, "utmgrp")

    @property
    @pulumi.getter(name="utmgrpPermissions")
    def utmgrp_permissions(self) -> Sequence['outputs.GetAccprofileUtmgrpPermissionResult']:
        """
        Custom Security Profile permissions. The structure of `utmgrp_permission` block is documented below.
        """
        return pulumi.get(self, "utmgrp_permissions")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vpngrp(self) -> str:
        """
        Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
        """
        return pulumi.get(self, "vpngrp")

    @property
    @pulumi.getter
    def wanoptgrp(self) -> str:
        """
        Administrator access to WAN Opt & Cache.
        """
        return pulumi.get(self, "wanoptgrp")

    @property
    @pulumi.getter
    def wifi(self) -> str:
        """
        Administrator access to the WiFi controller and Switch controller.
        """
        return pulumi.get(self, "wifi")


class AwaitableGetAccprofileResult(GetAccprofileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccprofileResult(
            admintimeout=self.admintimeout,
            admintimeout_override=self.admintimeout_override,
            authgrp=self.authgrp,
            cli_config=self.cli_config,
            cli_diagnose=self.cli_diagnose,
            cli_exec=self.cli_exec,
            cli_get=self.cli_get,
            cli_show=self.cli_show,
            comments=self.comments,
            ftviewgrp=self.ftviewgrp,
            fwgrp=self.fwgrp,
            fwgrp_permissions=self.fwgrp_permissions,
            id=self.id,
            loggrp=self.loggrp,
            loggrp_permissions=self.loggrp_permissions,
            name=self.name,
            netgrp=self.netgrp,
            netgrp_permissions=self.netgrp_permissions,
            scope=self.scope,
            secfabgrp=self.secfabgrp,
            sysgrp=self.sysgrp,
            sysgrp_permissions=self.sysgrp_permissions,
            system_diagnostics=self.system_diagnostics,
            system_execute_ssh=self.system_execute_ssh,
            system_execute_telnet=self.system_execute_telnet,
            utmgrp=self.utmgrp,
            utmgrp_permissions=self.utmgrp_permissions,
            vdomparam=self.vdomparam,
            vpngrp=self.vpngrp,
            wanoptgrp=self.wanoptgrp,
            wifi=self.wifi)


def get_accprofile(name: Optional[str] = None,
                   vdomparam: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccprofileResult:
    """
    Use this data source to get information on an fortios system accprofile


    :param str name: Specify the name of the desired system accprofile.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:system/getAccprofile:getAccprofile', __args__, opts=opts, typ=GetAccprofileResult).value

    return AwaitableGetAccprofileResult(
        admintimeout=pulumi.get(__ret__, 'admintimeout'),
        admintimeout_override=pulumi.get(__ret__, 'admintimeout_override'),
        authgrp=pulumi.get(__ret__, 'authgrp'),
        cli_config=pulumi.get(__ret__, 'cli_config'),
        cli_diagnose=pulumi.get(__ret__, 'cli_diagnose'),
        cli_exec=pulumi.get(__ret__, 'cli_exec'),
        cli_get=pulumi.get(__ret__, 'cli_get'),
        cli_show=pulumi.get(__ret__, 'cli_show'),
        comments=pulumi.get(__ret__, 'comments'),
        ftviewgrp=pulumi.get(__ret__, 'ftviewgrp'),
        fwgrp=pulumi.get(__ret__, 'fwgrp'),
        fwgrp_permissions=pulumi.get(__ret__, 'fwgrp_permissions'),
        id=pulumi.get(__ret__, 'id'),
        loggrp=pulumi.get(__ret__, 'loggrp'),
        loggrp_permissions=pulumi.get(__ret__, 'loggrp_permissions'),
        name=pulumi.get(__ret__, 'name'),
        netgrp=pulumi.get(__ret__, 'netgrp'),
        netgrp_permissions=pulumi.get(__ret__, 'netgrp_permissions'),
        scope=pulumi.get(__ret__, 'scope'),
        secfabgrp=pulumi.get(__ret__, 'secfabgrp'),
        sysgrp=pulumi.get(__ret__, 'sysgrp'),
        sysgrp_permissions=pulumi.get(__ret__, 'sysgrp_permissions'),
        system_diagnostics=pulumi.get(__ret__, 'system_diagnostics'),
        system_execute_ssh=pulumi.get(__ret__, 'system_execute_ssh'),
        system_execute_telnet=pulumi.get(__ret__, 'system_execute_telnet'),
        utmgrp=pulumi.get(__ret__, 'utmgrp'),
        utmgrp_permissions=pulumi.get(__ret__, 'utmgrp_permissions'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'),
        vpngrp=pulumi.get(__ret__, 'vpngrp'),
        wanoptgrp=pulumi.get(__ret__, 'wanoptgrp'),
        wifi=pulumi.get(__ret__, 'wifi'))


@_utilities.lift_output_func(get_accprofile)
def get_accprofile_output(name: Optional[pulumi.Input[str]] = None,
                          vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccprofileResult]:
    """
    Use this data source to get information on an fortios system accprofile


    :param str name: Specify the name of the desired system accprofile.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
