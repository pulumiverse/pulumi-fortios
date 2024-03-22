# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'SettingCustomLogFieldArgs',
    'ThreatweightApplicationArgs',
    'ThreatweightGeolocationArgs',
    'ThreatweightIpsArgs',
    'ThreatweightLevelArgs',
    'ThreatweightMalwareArgs',
    'ThreatweightWebArgs',
]

@pulumi.input_type
class SettingCustomLogFieldArgs:
    def __init__(__self__, *,
                 field_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] field_id: Custom log field.
        """
        if field_id is not None:
            pulumi.set(__self__, "field_id", field_id)

    @property
    @pulumi.getter(name="fieldId")
    def field_id(self) -> Optional[pulumi.Input[str]]:
        """
        Custom log field.
        """
        return pulumi.get(self, "field_id")

    @field_id.setter
    def field_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "field_id", value)


@pulumi.input_type
class ThreatweightApplicationArgs:
    def __init__(__self__, *,
                 category: Optional[pulumi.Input[int]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 level: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] category: Application category.
        :param pulumi.Input[int] id: Entry ID.
        :param pulumi.Input[str] level: Threat weight score for Application events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        if category is not None:
            pulumi.set(__self__, "category", category)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if level is not None:
            pulumi.set(__self__, "level", level)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[int]]:
        """
        Application category.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        Entry ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def level(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for Application events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "level")

    @level.setter
    def level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "level", value)


@pulumi.input_type
class ThreatweightGeolocationArgs:
    def __init__(__self__, *,
                 country: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 level: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] country: Country code.
        :param pulumi.Input[int] id: Entry ID.
        :param pulumi.Input[str] level: Threat weight score for Geolocation-based events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        if country is not None:
            pulumi.set(__self__, "country", country)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if level is not None:
            pulumi.set(__self__, "level", level)

    @property
    @pulumi.getter
    def country(self) -> Optional[pulumi.Input[str]]:
        """
        Country code.
        """
        return pulumi.get(self, "country")

    @country.setter
    def country(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "country", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        Entry ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def level(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for Geolocation-based events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "level")

    @level.setter
    def level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "level", value)


@pulumi.input_type
class ThreatweightIpsArgs:
    def __init__(__self__, *,
                 critical_severity: Optional[pulumi.Input[str]] = None,
                 high_severity: Optional[pulumi.Input[str]] = None,
                 info_severity: Optional[pulumi.Input[str]] = None,
                 low_severity: Optional[pulumi.Input[str]] = None,
                 medium_severity: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] critical_severity: Threat weight score for IPS critical severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] high_severity: Threat weight score for IPS high severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] info_severity: Threat weight score for IPS info severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] low_severity: Threat weight score for IPS low severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] medium_severity: Threat weight score for IPS medium severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        if critical_severity is not None:
            pulumi.set(__self__, "critical_severity", critical_severity)
        if high_severity is not None:
            pulumi.set(__self__, "high_severity", high_severity)
        if info_severity is not None:
            pulumi.set(__self__, "info_severity", info_severity)
        if low_severity is not None:
            pulumi.set(__self__, "low_severity", low_severity)
        if medium_severity is not None:
            pulumi.set(__self__, "medium_severity", medium_severity)

    @property
    @pulumi.getter(name="criticalSeverity")
    def critical_severity(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for IPS critical severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "critical_severity")

    @critical_severity.setter
    def critical_severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "critical_severity", value)

    @property
    @pulumi.getter(name="highSeverity")
    def high_severity(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for IPS high severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "high_severity")

    @high_severity.setter
    def high_severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "high_severity", value)

    @property
    @pulumi.getter(name="infoSeverity")
    def info_severity(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for IPS info severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "info_severity")

    @info_severity.setter
    def info_severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "info_severity", value)

    @property
    @pulumi.getter(name="lowSeverity")
    def low_severity(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for IPS low severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "low_severity")

    @low_severity.setter
    def low_severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "low_severity", value)

    @property
    @pulumi.getter(name="mediumSeverity")
    def medium_severity(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for IPS medium severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "medium_severity")

    @medium_severity.setter
    def medium_severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "medium_severity", value)


@pulumi.input_type
class ThreatweightLevelArgs:
    def __init__(__self__, *,
                 critical: Optional[pulumi.Input[int]] = None,
                 high: Optional[pulumi.Input[int]] = None,
                 low: Optional[pulumi.Input[int]] = None,
                 medium: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] critical: Critical level score value (1 - 100).
        :param pulumi.Input[int] high: High level score value (1 - 100).
        :param pulumi.Input[int] low: Low level score value (1 - 100).
        :param pulumi.Input[int] medium: Medium level score value (1 - 100).
        """
        if critical is not None:
            pulumi.set(__self__, "critical", critical)
        if high is not None:
            pulumi.set(__self__, "high", high)
        if low is not None:
            pulumi.set(__self__, "low", low)
        if medium is not None:
            pulumi.set(__self__, "medium", medium)

    @property
    @pulumi.getter
    def critical(self) -> Optional[pulumi.Input[int]]:
        """
        Critical level score value (1 - 100).
        """
        return pulumi.get(self, "critical")

    @critical.setter
    def critical(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "critical", value)

    @property
    @pulumi.getter
    def high(self) -> Optional[pulumi.Input[int]]:
        """
        High level score value (1 - 100).
        """
        return pulumi.get(self, "high")

    @high.setter
    def high(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "high", value)

    @property
    @pulumi.getter
    def low(self) -> Optional[pulumi.Input[int]]:
        """
        Low level score value (1 - 100).
        """
        return pulumi.get(self, "low")

    @low.setter
    def low(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "low", value)

    @property
    @pulumi.getter
    def medium(self) -> Optional[pulumi.Input[int]]:
        """
        Medium level score value (1 - 100).
        """
        return pulumi.get(self, "medium")

    @medium.setter
    def medium(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "medium", value)


@pulumi.input_type
class ThreatweightMalwareArgs:
    def __init__(__self__, *,
                 botnet_connection: Optional[pulumi.Input[str]] = None,
                 command_blocked: Optional[pulumi.Input[str]] = None,
                 content_disarm: Optional[pulumi.Input[str]] = None,
                 ems_threat_feed: Optional[pulumi.Input[str]] = None,
                 file_blocked: Optional[pulumi.Input[str]] = None,
                 fortiai: Optional[pulumi.Input[str]] = None,
                 fortindr: Optional[pulumi.Input[str]] = None,
                 fortisandbox: Optional[pulumi.Input[str]] = None,
                 fsa_high_risk: Optional[pulumi.Input[str]] = None,
                 fsa_malicious: Optional[pulumi.Input[str]] = None,
                 fsa_medium_risk: Optional[pulumi.Input[str]] = None,
                 inline_block: Optional[pulumi.Input[str]] = None,
                 malware_list: Optional[pulumi.Input[str]] = None,
                 mimefragmented: Optional[pulumi.Input[str]] = None,
                 oversized: Optional[pulumi.Input[str]] = None,
                 switch_proto: Optional[pulumi.Input[str]] = None,
                 virus_file_type_executable: Optional[pulumi.Input[str]] = None,
                 virus_infected: Optional[pulumi.Input[str]] = None,
                 virus_outbreak_prevention: Optional[pulumi.Input[str]] = None,
                 virus_scan_error: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] botnet_connection: Threat weight score for detected botnet connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] command_blocked: Threat weight score for blocked command detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] content_disarm: Threat weight score for virus (content disarm) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] ems_threat_feed: Threat weight score for virus (EMS threat feed) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] file_blocked: Threat weight score for blocked file detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] fortiai: Threat weight score for FortiAI-detected virus. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] fortindr: Threat weight score for FortiNDR-detected virus. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] fortisandbox: Threat weight score for FortiSandbox-detected virus. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] fsa_high_risk: Threat weight score for FortiSandbox high risk malware detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] fsa_malicious: Threat weight score for FortiSandbox malicious malware detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] fsa_medium_risk: Threat weight score for FortiSandbox medium risk malware detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] inline_block: Threat weight score for malware detected by inline block. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] malware_list: Threat weight score for virus (malware list) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] mimefragmented: Threat weight score for mimefragmented detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] oversized: Threat weight score for oversized file detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] switch_proto: Threat weight score for switch proto detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] virus_file_type_executable: Threat weight score for virus (filetype executable) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] virus_infected: Threat weight score for virus (infected) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] virus_outbreak_prevention: Threat weight score for virus (outbreak prevention) event. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] virus_scan_error: Threat weight score for virus (scan error) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        if botnet_connection is not None:
            pulumi.set(__self__, "botnet_connection", botnet_connection)
        if command_blocked is not None:
            pulumi.set(__self__, "command_blocked", command_blocked)
        if content_disarm is not None:
            pulumi.set(__self__, "content_disarm", content_disarm)
        if ems_threat_feed is not None:
            pulumi.set(__self__, "ems_threat_feed", ems_threat_feed)
        if file_blocked is not None:
            pulumi.set(__self__, "file_blocked", file_blocked)
        if fortiai is not None:
            pulumi.set(__self__, "fortiai", fortiai)
        if fortindr is not None:
            pulumi.set(__self__, "fortindr", fortindr)
        if fortisandbox is not None:
            pulumi.set(__self__, "fortisandbox", fortisandbox)
        if fsa_high_risk is not None:
            pulumi.set(__self__, "fsa_high_risk", fsa_high_risk)
        if fsa_malicious is not None:
            pulumi.set(__self__, "fsa_malicious", fsa_malicious)
        if fsa_medium_risk is not None:
            pulumi.set(__self__, "fsa_medium_risk", fsa_medium_risk)
        if inline_block is not None:
            pulumi.set(__self__, "inline_block", inline_block)
        if malware_list is not None:
            pulumi.set(__self__, "malware_list", malware_list)
        if mimefragmented is not None:
            pulumi.set(__self__, "mimefragmented", mimefragmented)
        if oversized is not None:
            pulumi.set(__self__, "oversized", oversized)
        if switch_proto is not None:
            pulumi.set(__self__, "switch_proto", switch_proto)
        if virus_file_type_executable is not None:
            pulumi.set(__self__, "virus_file_type_executable", virus_file_type_executable)
        if virus_infected is not None:
            pulumi.set(__self__, "virus_infected", virus_infected)
        if virus_outbreak_prevention is not None:
            pulumi.set(__self__, "virus_outbreak_prevention", virus_outbreak_prevention)
        if virus_scan_error is not None:
            pulumi.set(__self__, "virus_scan_error", virus_scan_error)

    @property
    @pulumi.getter(name="botnetConnection")
    def botnet_connection(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for detected botnet connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "botnet_connection")

    @botnet_connection.setter
    def botnet_connection(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "botnet_connection", value)

    @property
    @pulumi.getter(name="commandBlocked")
    def command_blocked(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for blocked command detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "command_blocked")

    @command_blocked.setter
    def command_blocked(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "command_blocked", value)

    @property
    @pulumi.getter(name="contentDisarm")
    def content_disarm(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for virus (content disarm) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "content_disarm")

    @content_disarm.setter
    def content_disarm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_disarm", value)

    @property
    @pulumi.getter(name="emsThreatFeed")
    def ems_threat_feed(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for virus (EMS threat feed) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "ems_threat_feed")

    @ems_threat_feed.setter
    def ems_threat_feed(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ems_threat_feed", value)

    @property
    @pulumi.getter(name="fileBlocked")
    def file_blocked(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for blocked file detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "file_blocked")

    @file_blocked.setter
    def file_blocked(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_blocked", value)

    @property
    @pulumi.getter
    def fortiai(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for FortiAI-detected virus. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "fortiai")

    @fortiai.setter
    def fortiai(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fortiai", value)

    @property
    @pulumi.getter
    def fortindr(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for FortiNDR-detected virus. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "fortindr")

    @fortindr.setter
    def fortindr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fortindr", value)

    @property
    @pulumi.getter
    def fortisandbox(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for FortiSandbox-detected virus. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "fortisandbox")

    @fortisandbox.setter
    def fortisandbox(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fortisandbox", value)

    @property
    @pulumi.getter(name="fsaHighRisk")
    def fsa_high_risk(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for FortiSandbox high risk malware detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "fsa_high_risk")

    @fsa_high_risk.setter
    def fsa_high_risk(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fsa_high_risk", value)

    @property
    @pulumi.getter(name="fsaMalicious")
    def fsa_malicious(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for FortiSandbox malicious malware detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "fsa_malicious")

    @fsa_malicious.setter
    def fsa_malicious(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fsa_malicious", value)

    @property
    @pulumi.getter(name="fsaMediumRisk")
    def fsa_medium_risk(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for FortiSandbox medium risk malware detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "fsa_medium_risk")

    @fsa_medium_risk.setter
    def fsa_medium_risk(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fsa_medium_risk", value)

    @property
    @pulumi.getter(name="inlineBlock")
    def inline_block(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for malware detected by inline block. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "inline_block")

    @inline_block.setter
    def inline_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "inline_block", value)

    @property
    @pulumi.getter(name="malwareList")
    def malware_list(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for virus (malware list) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "malware_list")

    @malware_list.setter
    def malware_list(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "malware_list", value)

    @property
    @pulumi.getter
    def mimefragmented(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for mimefragmented detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "mimefragmented")

    @mimefragmented.setter
    def mimefragmented(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mimefragmented", value)

    @property
    @pulumi.getter
    def oversized(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for oversized file detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "oversized")

    @oversized.setter
    def oversized(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "oversized", value)

    @property
    @pulumi.getter(name="switchProto")
    def switch_proto(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for switch proto detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "switch_proto")

    @switch_proto.setter
    def switch_proto(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "switch_proto", value)

    @property
    @pulumi.getter(name="virusFileTypeExecutable")
    def virus_file_type_executable(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for virus (filetype executable) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "virus_file_type_executable")

    @virus_file_type_executable.setter
    def virus_file_type_executable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virus_file_type_executable", value)

    @property
    @pulumi.getter(name="virusInfected")
    def virus_infected(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for virus (infected) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "virus_infected")

    @virus_infected.setter
    def virus_infected(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virus_infected", value)

    @property
    @pulumi.getter(name="virusOutbreakPrevention")
    def virus_outbreak_prevention(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for virus (outbreak prevention) event. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "virus_outbreak_prevention")

    @virus_outbreak_prevention.setter
    def virus_outbreak_prevention(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virus_outbreak_prevention", value)

    @property
    @pulumi.getter(name="virusScanError")
    def virus_scan_error(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for virus (scan error) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "virus_scan_error")

    @virus_scan_error.setter
    def virus_scan_error(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virus_scan_error", value)


@pulumi.input_type
class ThreatweightWebArgs:
    def __init__(__self__, *,
                 category: Optional[pulumi.Input[int]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 level: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] category: Threat weight score for web category filtering matches.
        :param pulumi.Input[int] id: Entry ID.
        :param pulumi.Input[str] level: Threat weight score for web category filtering matches. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        if category is not None:
            pulumi.set(__self__, "category", category)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if level is not None:
            pulumi.set(__self__, "level", level)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[int]]:
        """
        Threat weight score for web category filtering matches.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        Entry ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def level(self) -> Optional[pulumi.Input[str]]:
        """
        Threat weight score for web category filtering matches. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "level")

    @level.setter
    def level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "level", value)


