# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing
# Export this package's modules as members:
from .accprofile import *
from .acme import *
from .admin import *
from .admin_administrator import *
from .admin_profiles import *
from .affinityinterrupt import *
from .affinitypacketredistribution import *
from .alarm import *
from .alias import *
from .apiuser import *
from .apiuser_setting import *
from .arptable import *
from .autoinstall import *
from .automationaction import *
from .automationdestination import *
from .automationstitch import *
from .automationtrigger import *
from .autoscript import *
from .centralmanagement import *
from .clustersync import *
from .console import *
from .csf import *
from .customlanguage import *
from .ddns import *
from .dedicatedmgmt import *
from .dns import *
from .dns64 import *
from .dnsdatabase import *
from .dnsserver import *
from .dscpbasedpriority import *
from .emailserver import *
from .externalresource import *
from .federatedupgrade import *
from .fipscc import *
from .fm import *
from .fortiai import *
from .fortiguard import *
from .fortimanager import *
from .fortindr import *
from .fortisandbox import *
from .fssopolling import *
from .ftmpush import *
from .geneve import *
from .geoipcountry import *
from .geoipoverride import *
from .get_accprofile import *
from .get_accprofilelist import *
from .get_admin import *
from .get_adminlist import *
from .get_alias import *
from .get_aliaslist import *
from .get_apiuser import *
from .get_apiuserlist import *
from .get_arptable import *
from .get_arptablelist import *
from .get_autoinstall import *
from .get_automationaction import *
from .get_automationactionlist import *
from .get_automationdestination import *
from .get_automationdestinationlist import *
from .get_automationtrigger import *
from .get_automationtriggerlist import *
from .get_autoscript import *
from .get_autoscriptlist import *
from .get_centralmanagement import *
from .get_clustersync import *
from .get_clustersynclist import *
from .get_console import *
from .get_csf import *
from .get_ddns import *
from .get_ddnslist import *
from .get_dns import *
from .get_dnsdatabase import *
from .get_dnsdatabaselist import *
from .get_dnsserver import *
from .get_dnsserverlist import *
from .get_dscpbasedpriority import *
from .get_dscpbasedprioritylist import *
from .get_emailserver import *
from .get_externalresource import *
from .get_externalresourcelist import *
from .get_fipscc import *
from .get_fm import *
from .get_fortiguard import *
from .get_fortimanager import *
from .get_fortisandbox import *
from .get_fssopolling import *
from .get_ftmpush import *
from .get_global import *
from .get_gretunnel import *
from .get_gretunnellist import *
from .get_ha import *
from .get_hamonitor import *
from .get_interface import *
from .get_interfacelist import *
from .get_ipiptunnel import *
from .get_ipiptunnellist import *
from .get_ipv6neighborcache import *
from .get_ipv6neighborcachelist import *
from .get_ipv6tunnel import *
from .get_ipv6tunnellist import *
from .get_linkmonitor import *
from .get_linkmonitorlist import *
from .get_managementtunnel import *
from .get_mobiletunnel import *
from .get_mobiletunnellist import *
from .get_nat64 import *
from .get_ndproxy import *
from .get_netflow import *
from .get_networkvisibility import *
from .get_ntp import *
from .get_objecttagging import *
from .get_objecttagginglist import *
from .get_passwordpolicy import *
from .get_passwordpolicyguestadmin import *
from .get_pppoeinterface import *
from .get_pppoeinterfacelist import *
from .get_proberesponse import *
from .get_proxyarp import *
from .get_proxyarplist import *
from .get_replacemsggroup import *
from .get_replacemsggrouplist import *
from .get_replacemsgimage import *
from .get_replacemsgimagelist import *
from .get_resourcelimits import *
from .get_sdnconnector import *
from .get_sdnconnectorlist import *
from .get_sessionhelper import *
from .get_sessionhelperlist import *
from .get_sessionttl import *
from .get_sflow import *
from .get_sittunnel import *
from .get_sittunnellist import *
from .get_smsserver import *
from .get_smsserverlist import *
from .get_tosbasedpriority import *
from .get_tosbasedprioritylist import *
from .get_vdomexception import *
from .get_vdomexceptionlist import *
from .get_vdomnetflow import *
from .get_vdomsflow import *
from .get_virtualwanlink import *
from .get_vxlan import *
from .get_vxlanlist import *
from .get_wccp import *
from .get_wccplist import *
from .get_zone import *
from .get_zonelist import *
from .global_ import *
from .gretunnel import *
from .ha import *
from .hamonitor import *
from .ike import *
from .interface import *
from .ipam import *
from .ipiptunnel import *
from .ips import *
from .ipsecaggregate import *
from .ipsurlfilterdns import *
from .ipsurlfilterdns6 import *
from .ipv6neighborcache import *
from .ipv6tunnel import *
from .license_forticare import *
from .license_vdom import *
from .license_vm import *
from .linkmonitor import *
from .ltemodem import *
from .macaddresstable import *
from .managementtunnel import *
from .mobiletunnel import *
from .modem import *
from .nat64 import *
from .ndproxy import *
from .netflow import *
from .networkvisibility import *
from .npu import *
from .ntp import *
from .objecttagging import *
from .passwordpolicy import *
from .passwordpolicyguestadmin import *
from .physicalswitch import *
from .pppoeinterface import *
from .proberesponse import *
from .proxyarp import *
from .ptp import *
from .replacemsggroup import *
from .replacemsgimage import *
from .resourcelimits import *
from .saml import *
from .sdnconnector import *
from .sdwan import *
from .sessionhelper import *
from .sessionttl import *
from .setting_dns import *
from .setting_global import *
from .setting_ntp import *
from .settings import *
from .sflow import *
from .sittunnel import *
from .smsserver import *
from .speedtestschedule import *
from .speedtestserver import *
from .ssoadmin import *
from .ssoforticloudadmin import *
from .standalonecluster import *
from .storage import *
from .stp import *
from .switchinterface import *
from .tosbasedpriority import *
from .vdom import *
from .vdom_setting import *
from .vdomdns import *
from .vdomexception import *
from .vdomlink import *
from .vdomnetflow import *
from .vdomproperty import *
from .vdomradiusserver import *
from .vdomsflow import *
from .virtualswitch import *
from .virtualwanlink import *
from .virtualwirepair import *
from .vnetunnel import *
from .vxlan import *
from .wccp import *
from .zone import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumiverse_fortios.system.autoupdate as __autoupdate
    autoupdate = __autoupdate
    import pulumiverse_fortios.system.dhcp as __dhcp
    dhcp = __dhcp
    import pulumiverse_fortios.system.dhcp6 as __dhcp6
    dhcp6 = __dhcp6
    import pulumiverse_fortios.system.lldp as __lldp
    lldp = __lldp
    import pulumiverse_fortios.system.modem3g as __modem3g
    modem3g = __modem3g
    import pulumiverse_fortios.system.replacemsg as __replacemsg
    replacemsg = __replacemsg
    import pulumiverse_fortios.system.snmp as __snmp
    snmp = __snmp
else:
    autoupdate = _utilities.lazy_import('pulumiverse_fortios.system.autoupdate')
    dhcp = _utilities.lazy_import('pulumiverse_fortios.system.dhcp')
    dhcp6 = _utilities.lazy_import('pulumiverse_fortios.system.dhcp6')
    lldp = _utilities.lazy_import('pulumiverse_fortios.system.lldp')
    modem3g = _utilities.lazy_import('pulumiverse_fortios.system.modem3g')
    replacemsg = _utilities.lazy_import('pulumiverse_fortios.system.replacemsg')
    snmp = _utilities.lazy_import('pulumiverse_fortios.system.snmp')

