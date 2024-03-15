# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing
# Export this package's modules as members:
from .customcommand import *
from .dynamicportpolicy import *
from .flowtracking import *
from .fortilinksettings import *
from .global_ import *
from .igmpsnooping import *
from .lldpprofile import *
from .lldpsettings import *
from .location import *
from .macsyncsettings import *
from .managedswitch import *
from .nacdevice import *
from .nacsettings import *
from .networkmonitorsettings import *
from .portpolicy import *
from .quarantine import *
from .remotelog import *
from .settings8021_x import *
from .sflow import *
from .snmpcommunity import *
from .snmpsysinfo import *
from .snmptrapthreshold import *
from .snmpuser import *
from .stormcontrol import *
from .stormcontrolpolicy import *
from .stpinstance import *
from .stpsettings import *
from .switchgroup import *
from .switchinterfacetag import *
from .switchlog import *
from .switchprofile import *
from .system import *
from .trafficpolicy import *
from .trafficsniffer import *
from .virtualportpool import *
from .vlan import *
from .vlanpolicy import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumiverse_fortios.switchcontroller.autoconfig as __autoconfig
    autoconfig = __autoconfig
    import pulumiverse_fortios.switchcontroller.initialconfig as __initialconfig
    initialconfig = __initialconfig
    import pulumiverse_fortios.switchcontroller.ptp as __ptp
    ptp = __ptp
    import pulumiverse_fortios.switchcontroller.qos as __qos
    qos = __qos
    import pulumiverse_fortios.switchcontroller.securitypolicy as __securitypolicy
    securitypolicy = __securitypolicy
else:
    autoconfig = _utilities.lazy_import('pulumiverse_fortios.switchcontroller.autoconfig')
    initialconfig = _utilities.lazy_import('pulumiverse_fortios.switchcontroller.initialconfig')
    ptp = _utilities.lazy_import('pulumiverse_fortios.switchcontroller.ptp')
    qos = _utilities.lazy_import('pulumiverse_fortios.switchcontroller.qos')
    securitypolicy = _utilities.lazy_import('pulumiverse_fortios.switchcontroller.securitypolicy')

