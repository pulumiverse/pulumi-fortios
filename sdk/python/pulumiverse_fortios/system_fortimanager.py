# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemFortimanagerArgs', 'SystemFortimanager']

@pulumi.input_type
class SystemFortimanagerArgs:
    def __init__(__self__, *,
                 central_management: Optional[pulumi.Input[str]] = None,
                 central_mgmt_auto_backup: Optional[pulumi.Input[str]] = None,
                 central_mgmt_schedule_config_restore: Optional[pulumi.Input[str]] = None,
                 central_mgmt_schedule_script_restore: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ipsec: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemFortimanager resource.
        """
        if central_management is not None:
            pulumi.set(__self__, "central_management", central_management)
        if central_mgmt_auto_backup is not None:
            pulumi.set(__self__, "central_mgmt_auto_backup", central_mgmt_auto_backup)
        if central_mgmt_schedule_config_restore is not None:
            pulumi.set(__self__, "central_mgmt_schedule_config_restore", central_mgmt_schedule_config_restore)
        if central_mgmt_schedule_script_restore is not None:
            pulumi.set(__self__, "central_mgmt_schedule_script_restore", central_mgmt_schedule_script_restore)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if ipsec is not None:
            pulumi.set(__self__, "ipsec", ipsec)
        if vdom is not None:
            pulumi.set(__self__, "vdom", vdom)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="centralManagement")
    def central_management(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "central_management")

    @central_management.setter
    def central_management(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "central_management", value)

    @property
    @pulumi.getter(name="centralMgmtAutoBackup")
    def central_mgmt_auto_backup(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "central_mgmt_auto_backup")

    @central_mgmt_auto_backup.setter
    def central_mgmt_auto_backup(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "central_mgmt_auto_backup", value)

    @property
    @pulumi.getter(name="centralMgmtScheduleConfigRestore")
    def central_mgmt_schedule_config_restore(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "central_mgmt_schedule_config_restore")

    @central_mgmt_schedule_config_restore.setter
    def central_mgmt_schedule_config_restore(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "central_mgmt_schedule_config_restore", value)

    @property
    @pulumi.getter(name="centralMgmtScheduleScriptRestore")
    def central_mgmt_schedule_script_restore(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "central_mgmt_schedule_script_restore")

    @central_mgmt_schedule_script_restore.setter
    def central_mgmt_schedule_script_restore(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "central_mgmt_schedule_script_restore", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def ipsec(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ipsec")

    @ipsec.setter
    def ipsec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipsec", value)

    @property
    @pulumi.getter
    def vdom(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdom")

    @vdom.setter
    def vdom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdom", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _SystemFortimanagerState:
    def __init__(__self__, *,
                 central_management: Optional[pulumi.Input[str]] = None,
                 central_mgmt_auto_backup: Optional[pulumi.Input[str]] = None,
                 central_mgmt_schedule_config_restore: Optional[pulumi.Input[str]] = None,
                 central_mgmt_schedule_script_restore: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ipsec: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemFortimanager resources.
        """
        if central_management is not None:
            pulumi.set(__self__, "central_management", central_management)
        if central_mgmt_auto_backup is not None:
            pulumi.set(__self__, "central_mgmt_auto_backup", central_mgmt_auto_backup)
        if central_mgmt_schedule_config_restore is not None:
            pulumi.set(__self__, "central_mgmt_schedule_config_restore", central_mgmt_schedule_config_restore)
        if central_mgmt_schedule_script_restore is not None:
            pulumi.set(__self__, "central_mgmt_schedule_script_restore", central_mgmt_schedule_script_restore)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if ipsec is not None:
            pulumi.set(__self__, "ipsec", ipsec)
        if vdom is not None:
            pulumi.set(__self__, "vdom", vdom)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="centralManagement")
    def central_management(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "central_management")

    @central_management.setter
    def central_management(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "central_management", value)

    @property
    @pulumi.getter(name="centralMgmtAutoBackup")
    def central_mgmt_auto_backup(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "central_mgmt_auto_backup")

    @central_mgmt_auto_backup.setter
    def central_mgmt_auto_backup(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "central_mgmt_auto_backup", value)

    @property
    @pulumi.getter(name="centralMgmtScheduleConfigRestore")
    def central_mgmt_schedule_config_restore(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "central_mgmt_schedule_config_restore")

    @central_mgmt_schedule_config_restore.setter
    def central_mgmt_schedule_config_restore(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "central_mgmt_schedule_config_restore", value)

    @property
    @pulumi.getter(name="centralMgmtScheduleScriptRestore")
    def central_mgmt_schedule_script_restore(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "central_mgmt_schedule_script_restore")

    @central_mgmt_schedule_script_restore.setter
    def central_mgmt_schedule_script_restore(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "central_mgmt_schedule_script_restore", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def ipsec(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ipsec")

    @ipsec.setter
    def ipsec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipsec", value)

    @property
    @pulumi.getter
    def vdom(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdom")

    @vdom.setter
    def vdom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdom", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class SystemFortimanager(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 central_management: Optional[pulumi.Input[str]] = None,
                 central_mgmt_auto_backup: Optional[pulumi.Input[str]] = None,
                 central_mgmt_schedule_config_restore: Optional[pulumi.Input[str]] = None,
                 central_mgmt_schedule_script_restore: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ipsec: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure FortiManager. Applies to FortiOS Version `<= 7.0.1`.

        By design considerations, the feature is using the SystemCentralmanagement resource as documented below.

        ## Example

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.SystemCentralmanagement("trname",
            allow_monitor="enable",
            allow_push_configuration="enable",
            allow_push_firmware="enable",
            allow_remote_firmware_upgrade="enable",
            enc_algorithm="high",
            fmg="\\"192.168.52.177\\"",
            include_default_servers="enable",
            mode="normal",
            type="fortimanager",
            vdom="root")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SystemFortimanagerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure FortiManager. Applies to FortiOS Version `<= 7.0.1`.

        By design considerations, the feature is using the SystemCentralmanagement resource as documented below.

        ## Example

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.SystemCentralmanagement("trname",
            allow_monitor="enable",
            allow_push_configuration="enable",
            allow_push_firmware="enable",
            allow_remote_firmware_upgrade="enable",
            enc_algorithm="high",
            fmg="\\"192.168.52.177\\"",
            include_default_servers="enable",
            mode="normal",
            type="fortimanager",
            vdom="root")
        ```

        :param str resource_name: The name of the resource.
        :param SystemFortimanagerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemFortimanagerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 central_management: Optional[pulumi.Input[str]] = None,
                 central_mgmt_auto_backup: Optional[pulumi.Input[str]] = None,
                 central_mgmt_schedule_config_restore: Optional[pulumi.Input[str]] = None,
                 central_mgmt_schedule_script_restore: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ipsec: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemFortimanagerArgs.__new__(SystemFortimanagerArgs)

            __props__.__dict__["central_management"] = central_management
            __props__.__dict__["central_mgmt_auto_backup"] = central_mgmt_auto_backup
            __props__.__dict__["central_mgmt_schedule_config_restore"] = central_mgmt_schedule_config_restore
            __props__.__dict__["central_mgmt_schedule_script_restore"] = central_mgmt_schedule_script_restore
            __props__.__dict__["ip"] = ip
            __props__.__dict__["ipsec"] = ipsec
            __props__.__dict__["vdom"] = vdom
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemFortimanager, __self__).__init__(
            'fortios:index/systemFortimanager:SystemFortimanager',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            central_management: Optional[pulumi.Input[str]] = None,
            central_mgmt_auto_backup: Optional[pulumi.Input[str]] = None,
            central_mgmt_schedule_config_restore: Optional[pulumi.Input[str]] = None,
            central_mgmt_schedule_script_restore: Optional[pulumi.Input[str]] = None,
            ip: Optional[pulumi.Input[str]] = None,
            ipsec: Optional[pulumi.Input[str]] = None,
            vdom: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemFortimanager':
        """
        Get an existing SystemFortimanager resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemFortimanagerState.__new__(_SystemFortimanagerState)

        __props__.__dict__["central_management"] = central_management
        __props__.__dict__["central_mgmt_auto_backup"] = central_mgmt_auto_backup
        __props__.__dict__["central_mgmt_schedule_config_restore"] = central_mgmt_schedule_config_restore
        __props__.__dict__["central_mgmt_schedule_script_restore"] = central_mgmt_schedule_script_restore
        __props__.__dict__["ip"] = ip
        __props__.__dict__["ipsec"] = ipsec
        __props__.__dict__["vdom"] = vdom
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemFortimanager(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="centralManagement")
    def central_management(self) -> pulumi.Output[str]:
        return pulumi.get(self, "central_management")

    @property
    @pulumi.getter(name="centralMgmtAutoBackup")
    def central_mgmt_auto_backup(self) -> pulumi.Output[str]:
        return pulumi.get(self, "central_mgmt_auto_backup")

    @property
    @pulumi.getter(name="centralMgmtScheduleConfigRestore")
    def central_mgmt_schedule_config_restore(self) -> pulumi.Output[str]:
        return pulumi.get(self, "central_mgmt_schedule_config_restore")

    @property
    @pulumi.getter(name="centralMgmtScheduleScriptRestore")
    def central_mgmt_schedule_script_restore(self) -> pulumi.Output[str]:
        return pulumi.get(self, "central_mgmt_schedule_script_restore")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def ipsec(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ipsec")

    @property
    @pulumi.getter
    def vdom(self) -> pulumi.Output[str]:
        return pulumi.get(self, "vdom")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

