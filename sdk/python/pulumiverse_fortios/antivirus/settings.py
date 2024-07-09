# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SettingsArgs', 'Settings']

@pulumi.input_type
class SettingsArgs:
    def __init__(__self__, *,
                 cache_clean_result: Optional[pulumi.Input[str]] = None,
                 cache_infected_result: Optional[pulumi.Input[str]] = None,
                 default_db: Optional[pulumi.Input[str]] = None,
                 grayware: Optional[pulumi.Input[str]] = None,
                 machine_learning_detection: Optional[pulumi.Input[str]] = None,
                 override_timeout: Optional[pulumi.Input[int]] = None,
                 use_extreme_db: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Settings resource.
        :param pulumi.Input[str] cache_clean_result: Enable/disable cache of clean scan results (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] cache_infected_result: Enable/disable cache of infected scan results (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] default_db: Select the AV database to be used for AV scanning. Valid values: `normal`, `extended`, `extreme`.
        :param pulumi.Input[str] grayware: Enable/disable grayware detection when an AntiVirus profile is applied to traffic. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] machine_learning_detection: Use machine learning based malware detection. Valid values: `enable`, `monitor`, `disable`.
        :param pulumi.Input[int] override_timeout: Override the large file scan timeout value in seconds (30 - 3600). Zero is the default value and is used to disable this command. When disabled, the daemon adjusts the large file scan timeout based on the file size.
        :param pulumi.Input[str] use_extreme_db: Enable/disable the use of Extreme AVDB. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if cache_clean_result is not None:
            pulumi.set(__self__, "cache_clean_result", cache_clean_result)
        if cache_infected_result is not None:
            pulumi.set(__self__, "cache_infected_result", cache_infected_result)
        if default_db is not None:
            pulumi.set(__self__, "default_db", default_db)
        if grayware is not None:
            pulumi.set(__self__, "grayware", grayware)
        if machine_learning_detection is not None:
            pulumi.set(__self__, "machine_learning_detection", machine_learning_detection)
        if override_timeout is not None:
            pulumi.set(__self__, "override_timeout", override_timeout)
        if use_extreme_db is not None:
            pulumi.set(__self__, "use_extreme_db", use_extreme_db)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="cacheCleanResult")
    def cache_clean_result(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable cache of clean scan results (default = enable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "cache_clean_result")

    @cache_clean_result.setter
    def cache_clean_result(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cache_clean_result", value)

    @property
    @pulumi.getter(name="cacheInfectedResult")
    def cache_infected_result(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable cache of infected scan results (default = enable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "cache_infected_result")

    @cache_infected_result.setter
    def cache_infected_result(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cache_infected_result", value)

    @property
    @pulumi.getter(name="defaultDb")
    def default_db(self) -> Optional[pulumi.Input[str]]:
        """
        Select the AV database to be used for AV scanning. Valid values: `normal`, `extended`, `extreme`.
        """
        return pulumi.get(self, "default_db")

    @default_db.setter
    def default_db(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_db", value)

    @property
    @pulumi.getter
    def grayware(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable grayware detection when an AntiVirus profile is applied to traffic. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "grayware")

    @grayware.setter
    def grayware(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grayware", value)

    @property
    @pulumi.getter(name="machineLearningDetection")
    def machine_learning_detection(self) -> Optional[pulumi.Input[str]]:
        """
        Use machine learning based malware detection. Valid values: `enable`, `monitor`, `disable`.
        """
        return pulumi.get(self, "machine_learning_detection")

    @machine_learning_detection.setter
    def machine_learning_detection(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "machine_learning_detection", value)

    @property
    @pulumi.getter(name="overrideTimeout")
    def override_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Override the large file scan timeout value in seconds (30 - 3600). Zero is the default value and is used to disable this command. When disabled, the daemon adjusts the large file scan timeout based on the file size.
        """
        return pulumi.get(self, "override_timeout")

    @override_timeout.setter
    def override_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "override_timeout", value)

    @property
    @pulumi.getter(name="useExtremeDb")
    def use_extreme_db(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable the use of Extreme AVDB. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "use_extreme_db")

    @use_extreme_db.setter
    def use_extreme_db(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "use_extreme_db", value)

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
class _SettingsState:
    def __init__(__self__, *,
                 cache_clean_result: Optional[pulumi.Input[str]] = None,
                 cache_infected_result: Optional[pulumi.Input[str]] = None,
                 default_db: Optional[pulumi.Input[str]] = None,
                 grayware: Optional[pulumi.Input[str]] = None,
                 machine_learning_detection: Optional[pulumi.Input[str]] = None,
                 override_timeout: Optional[pulumi.Input[int]] = None,
                 use_extreme_db: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Settings resources.
        :param pulumi.Input[str] cache_clean_result: Enable/disable cache of clean scan results (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] cache_infected_result: Enable/disable cache of infected scan results (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] default_db: Select the AV database to be used for AV scanning. Valid values: `normal`, `extended`, `extreme`.
        :param pulumi.Input[str] grayware: Enable/disable grayware detection when an AntiVirus profile is applied to traffic. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] machine_learning_detection: Use machine learning based malware detection. Valid values: `enable`, `monitor`, `disable`.
        :param pulumi.Input[int] override_timeout: Override the large file scan timeout value in seconds (30 - 3600). Zero is the default value and is used to disable this command. When disabled, the daemon adjusts the large file scan timeout based on the file size.
        :param pulumi.Input[str] use_extreme_db: Enable/disable the use of Extreme AVDB. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if cache_clean_result is not None:
            pulumi.set(__self__, "cache_clean_result", cache_clean_result)
        if cache_infected_result is not None:
            pulumi.set(__self__, "cache_infected_result", cache_infected_result)
        if default_db is not None:
            pulumi.set(__self__, "default_db", default_db)
        if grayware is not None:
            pulumi.set(__self__, "grayware", grayware)
        if machine_learning_detection is not None:
            pulumi.set(__self__, "machine_learning_detection", machine_learning_detection)
        if override_timeout is not None:
            pulumi.set(__self__, "override_timeout", override_timeout)
        if use_extreme_db is not None:
            pulumi.set(__self__, "use_extreme_db", use_extreme_db)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="cacheCleanResult")
    def cache_clean_result(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable cache of clean scan results (default = enable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "cache_clean_result")

    @cache_clean_result.setter
    def cache_clean_result(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cache_clean_result", value)

    @property
    @pulumi.getter(name="cacheInfectedResult")
    def cache_infected_result(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable cache of infected scan results (default = enable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "cache_infected_result")

    @cache_infected_result.setter
    def cache_infected_result(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cache_infected_result", value)

    @property
    @pulumi.getter(name="defaultDb")
    def default_db(self) -> Optional[pulumi.Input[str]]:
        """
        Select the AV database to be used for AV scanning. Valid values: `normal`, `extended`, `extreme`.
        """
        return pulumi.get(self, "default_db")

    @default_db.setter
    def default_db(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_db", value)

    @property
    @pulumi.getter
    def grayware(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable grayware detection when an AntiVirus profile is applied to traffic. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "grayware")

    @grayware.setter
    def grayware(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grayware", value)

    @property
    @pulumi.getter(name="machineLearningDetection")
    def machine_learning_detection(self) -> Optional[pulumi.Input[str]]:
        """
        Use machine learning based malware detection. Valid values: `enable`, `monitor`, `disable`.
        """
        return pulumi.get(self, "machine_learning_detection")

    @machine_learning_detection.setter
    def machine_learning_detection(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "machine_learning_detection", value)

    @property
    @pulumi.getter(name="overrideTimeout")
    def override_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Override the large file scan timeout value in seconds (30 - 3600). Zero is the default value and is used to disable this command. When disabled, the daemon adjusts the large file scan timeout based on the file size.
        """
        return pulumi.get(self, "override_timeout")

    @override_timeout.setter
    def override_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "override_timeout", value)

    @property
    @pulumi.getter(name="useExtremeDb")
    def use_extreme_db(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable the use of Extreme AVDB. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "use_extreme_db")

    @use_extreme_db.setter
    def use_extreme_db(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "use_extreme_db", value)

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


class Settings(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cache_clean_result: Optional[pulumi.Input[str]] = None,
                 cache_infected_result: Optional[pulumi.Input[str]] = None,
                 default_db: Optional[pulumi.Input[str]] = None,
                 grayware: Optional[pulumi.Input[str]] = None,
                 machine_learning_detection: Optional[pulumi.Input[str]] = None,
                 override_timeout: Optional[pulumi.Input[int]] = None,
                 use_extreme_db: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure AntiVirus settings.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.antivirus.Settings("trname",
            default_db="extended",
            grayware="enable")
        ```

        ## Import

        Antivirus Settings can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:antivirus/settings:Settings labelname AntivirusSettings
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:antivirus/settings:Settings labelname AntivirusSettings
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cache_clean_result: Enable/disable cache of clean scan results (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] cache_infected_result: Enable/disable cache of infected scan results (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] default_db: Select the AV database to be used for AV scanning. Valid values: `normal`, `extended`, `extreme`.
        :param pulumi.Input[str] grayware: Enable/disable grayware detection when an AntiVirus profile is applied to traffic. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] machine_learning_detection: Use machine learning based malware detection. Valid values: `enable`, `monitor`, `disable`.
        :param pulumi.Input[int] override_timeout: Override the large file scan timeout value in seconds (30 - 3600). Zero is the default value and is used to disable this command. When disabled, the daemon adjusts the large file scan timeout based on the file size.
        :param pulumi.Input[str] use_extreme_db: Enable/disable the use of Extreme AVDB. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SettingsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure AntiVirus settings.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.antivirus.Settings("trname",
            default_db="extended",
            grayware="enable")
        ```

        ## Import

        Antivirus Settings can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:antivirus/settings:Settings labelname AntivirusSettings
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:antivirus/settings:Settings labelname AntivirusSettings
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SettingsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SettingsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cache_clean_result: Optional[pulumi.Input[str]] = None,
                 cache_infected_result: Optional[pulumi.Input[str]] = None,
                 default_db: Optional[pulumi.Input[str]] = None,
                 grayware: Optional[pulumi.Input[str]] = None,
                 machine_learning_detection: Optional[pulumi.Input[str]] = None,
                 override_timeout: Optional[pulumi.Input[int]] = None,
                 use_extreme_db: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SettingsArgs.__new__(SettingsArgs)

            __props__.__dict__["cache_clean_result"] = cache_clean_result
            __props__.__dict__["cache_infected_result"] = cache_infected_result
            __props__.__dict__["default_db"] = default_db
            __props__.__dict__["grayware"] = grayware
            __props__.__dict__["machine_learning_detection"] = machine_learning_detection
            __props__.__dict__["override_timeout"] = override_timeout
            __props__.__dict__["use_extreme_db"] = use_extreme_db
            __props__.__dict__["vdomparam"] = vdomparam
        super(Settings, __self__).__init__(
            'fortios:antivirus/settings:Settings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cache_clean_result: Optional[pulumi.Input[str]] = None,
            cache_infected_result: Optional[pulumi.Input[str]] = None,
            default_db: Optional[pulumi.Input[str]] = None,
            grayware: Optional[pulumi.Input[str]] = None,
            machine_learning_detection: Optional[pulumi.Input[str]] = None,
            override_timeout: Optional[pulumi.Input[int]] = None,
            use_extreme_db: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Settings':
        """
        Get an existing Settings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cache_clean_result: Enable/disable cache of clean scan results (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] cache_infected_result: Enable/disable cache of infected scan results (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] default_db: Select the AV database to be used for AV scanning. Valid values: `normal`, `extended`, `extreme`.
        :param pulumi.Input[str] grayware: Enable/disable grayware detection when an AntiVirus profile is applied to traffic. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] machine_learning_detection: Use machine learning based malware detection. Valid values: `enable`, `monitor`, `disable`.
        :param pulumi.Input[int] override_timeout: Override the large file scan timeout value in seconds (30 - 3600). Zero is the default value and is used to disable this command. When disabled, the daemon adjusts the large file scan timeout based on the file size.
        :param pulumi.Input[str] use_extreme_db: Enable/disable the use of Extreme AVDB. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SettingsState.__new__(_SettingsState)

        __props__.__dict__["cache_clean_result"] = cache_clean_result
        __props__.__dict__["cache_infected_result"] = cache_infected_result
        __props__.__dict__["default_db"] = default_db
        __props__.__dict__["grayware"] = grayware
        __props__.__dict__["machine_learning_detection"] = machine_learning_detection
        __props__.__dict__["override_timeout"] = override_timeout
        __props__.__dict__["use_extreme_db"] = use_extreme_db
        __props__.__dict__["vdomparam"] = vdomparam
        return Settings(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cacheCleanResult")
    def cache_clean_result(self) -> pulumi.Output[str]:
        """
        Enable/disable cache of clean scan results (default = enable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "cache_clean_result")

    @property
    @pulumi.getter(name="cacheInfectedResult")
    def cache_infected_result(self) -> pulumi.Output[str]:
        """
        Enable/disable cache of infected scan results (default = enable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "cache_infected_result")

    @property
    @pulumi.getter(name="defaultDb")
    def default_db(self) -> pulumi.Output[str]:
        """
        Select the AV database to be used for AV scanning. Valid values: `normal`, `extended`, `extreme`.
        """
        return pulumi.get(self, "default_db")

    @property
    @pulumi.getter
    def grayware(self) -> pulumi.Output[str]:
        """
        Enable/disable grayware detection when an AntiVirus profile is applied to traffic. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "grayware")

    @property
    @pulumi.getter(name="machineLearningDetection")
    def machine_learning_detection(self) -> pulumi.Output[str]:
        """
        Use machine learning based malware detection. Valid values: `enable`, `monitor`, `disable`.
        """
        return pulumi.get(self, "machine_learning_detection")

    @property
    @pulumi.getter(name="overrideTimeout")
    def override_timeout(self) -> pulumi.Output[int]:
        """
        Override the large file scan timeout value in seconds (30 - 3600). Zero is the default value and is used to disable this command. When disabled, the daemon adjusts the large file scan timeout based on the file size.
        """
        return pulumi.get(self, "override_timeout")

    @property
    @pulumi.getter(name="useExtremeDb")
    def use_extreme_db(self) -> pulumi.Output[str]:
        """
        Enable/disable the use of Extreme AVDB. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "use_extreme_db")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

