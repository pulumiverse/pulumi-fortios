// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource supports modifying system global setting for FortiManager.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test1 = new fortios.fmg.SystemGlobal("test1", {
 *     adomMode: "advanced",
 *     adomStatus: "disable",
 *     fortianalyzerStatus: "disable",
 *     hostname: "FMG-VM64-test",
 *     timezone: "09",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class SystemGlobal extends pulumi.CustomResource {
    /**
     * Get an existing SystemGlobal resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemGlobalState, opts?: pulumi.CustomResourceOptions): SystemGlobal {
        return new SystemGlobal(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:fmg/systemGlobal:SystemGlobal';

    /**
     * Returns true if the given object is an instance of SystemGlobal.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemGlobal {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemGlobal.__pulumiType;
    }

    /**
     * Adom Mode.
     */
    public readonly adomMode!: pulumi.Output<string | undefined>;
    /**
     * Enable/Disable ADOM.
     */
    public readonly adomStatus!: pulumi.Output<string | undefined>;
    /**
     * Enable/Disable FortiAnalyzer feature.
     */
    public readonly fortianalyzerStatus!: pulumi.Output<string | undefined>;
    /**
     * Hostname.
     */
    public readonly hostname!: pulumi.Output<string | undefined>;
    /**
     * TimeZone. 00 - (GMT-12:00) Eniwetak, Kwajalein.01 - (GMT-11:00) Midway Island, Samoa.02 - (GMT-10:00) Hawaii.03 - (GMT-9:00) Alaska.04 - (GMT-8:00) Pacific Time (US & Canada).05 - (GMT-7:00) Arizona.06 - (GMT-7:00) Mountain Time (US & Canada).07 - (GMT-6:00) Central America.08 - (GMT-6:00) Central Time (US & Canada).09 - (GMT-6:00) Mexico City.10 - (GMT-6:00) Saskatchewan.11 - (GMT-5:00) Bogota, Lima, Quito.12 - (GMT-5:00) Eastern Time (US & Canada).13 - (GMT-5:00) Indiana (East).14 - (GMT-4:00) Atlantic Time (Canada).15 - (GMT-4:00) La Paz.16 - (GMT-4:00) Santiago.17 - (GMT-3:30) Newfoundland.18 - (GMT-3:00) Brasilia.19 - (GMT-3:00) Buenos Aires, Georgetown.20 - (GMT-3:00) Nuuk (Greenland).21 - (GMT-2:00) Mid-Atlantic.22 - (GMT-1:00) Azores.23 - (GMT-1:00) Cape Verde Is.24 - (GMT) Monrovia.25 - (GMT) Greenwich Mean Time:Dublin, Edinburgh, Lisbon, London.26 - (GMT+1:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna.27 - (GMT+1:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague.28 - (GMT+1:00) Brussels, Copenhagen, Madrid, Paris.29 - (GMT+1:00) Sarajevo, Skopje, Warsaw, Zagreb.30 - (GMT+1:00) West Central Africa.31 - (GMT+2:00) Athens, Sofia, Vilnius.32 - (GMT+2:00) Bucharest.33 - (GMT+2:00) Cairo.34 - (GMT+2:00) Harare, Pretoria.35 - (GMT+2:00) Helsinki, Riga,Tallinn.36 - (GMT+2:00) Jerusalem.37 - (GMT+3:00) Baghdad.38 - (GMT+3:00) Kuwait, Riyadh.39 - (GMT+3:00) St.Petersburg, Volgograd.40 - (GMT+3:00) Nairobi.41 - (GMT+3:30) Tehran.42 - (GMT+4:00) Abu Dhabi, Muscat.43 - (GMT+4:00) Baku.44 - (GMT+4:30) Kabul.45 - (GMT+5:00) Ekaterinburg.46 - (GMT+5:00) Islamabad, Karachi,Tashkent.47 - (GMT+5:30) Calcutta, Chennai, Mumbai, New Delhi.48 - (GMT+5:45) Kathmandu.49 - (GMT+6:00) Almaty, Novosibirsk.50 - (GMT+6:00) Astana, Dhaka.51 - (GMT+6:00) Sri Jayawardenapura.52 - (GMT+6:30) Rangoon.53 - (GMT+7:00) Bangkok, Hanoi, Jakarta.54 - (GMT+7:00) Krasnoyarsk.55 - (GMT+8:00) Beijing,ChongQing, HongKong,Urumqi.56 - (GMT+8:00) Irkutsk, Ulaanbaatar.57 - (GMT+8:00) Kuala Lumpur, Singapore.58 - (GMT+8:00) Perth.59 - (GMT+8:00) Taipei.60 - (GMT+9:00) Osaka, Sapporo, Tokyo, Seoul.61 - (GMT+9:00) Yakutsk.62 - (GMT+9:30) Adelaide.63 - (GMT+9:30) Darwin.64 - (GMT+10:00) Brisbane.65 - (GMT+10:00) Canberra, Melbourne, Sydney.66 - (GMT+10:00) Guam, Port Moresby.67 - (GMT+10:00) Hobart.68 - (GMT+10:00) Vladivostok.69 - (GMT+11:00) Magadan.70 - (GMT+11:00) Solomon Is., New Caledonia.71 - (GMT+12:00) Auckland, Wellington.72 - (GMT+12:00) Fiji, Kamchatka, Marshall Is.73 - (GMT+13:00) Nuku'alofa.74 - (GMT-4:30) Caracas.75 - (GMT+1:00) Namibia.76 - (GMT-5:00) Brazil-Acre.77 - (GMT-4:00) Brazil-West.78 - (GMT-3:00) Brazil-East.79 - (GMT-2:00) Brazil-DeNoronha.80 - (GMT+14:00) Kiritimati.81 - (GMT-7:00) Baja California Sur, Chihuahua.82 - (GMT+12:45) Chatham Islands.83 - (GMT+3:00) Minsk.84 - (GMT+13:00) Samoa.85 - (GMT+3:00) Istanbul.86 - (GMT-4:00) Paraguay.87 - (GMT) Casablanca.88 - (GMT+3:00) Moscow.89 - (GMT) Greenwich Mean Time.
     */
    public readonly timezone!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemGlobal resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemGlobalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemGlobalArgs | SystemGlobalState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemGlobalState | undefined;
            resourceInputs["adomMode"] = state ? state.adomMode : undefined;
            resourceInputs["adomStatus"] = state ? state.adomStatus : undefined;
            resourceInputs["fortianalyzerStatus"] = state ? state.fortianalyzerStatus : undefined;
            resourceInputs["hostname"] = state ? state.hostname : undefined;
            resourceInputs["timezone"] = state ? state.timezone : undefined;
        } else {
            const args = argsOrState as SystemGlobalArgs | undefined;
            resourceInputs["adomMode"] = args ? args.adomMode : undefined;
            resourceInputs["adomStatus"] = args ? args.adomStatus : undefined;
            resourceInputs["fortianalyzerStatus"] = args ? args.fortianalyzerStatus : undefined;
            resourceInputs["hostname"] = args ? args.hostname : undefined;
            resourceInputs["timezone"] = args ? args.timezone : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemGlobal.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemGlobal resources.
 */
export interface SystemGlobalState {
    /**
     * Adom Mode.
     */
    adomMode?: pulumi.Input<string>;
    /**
     * Enable/Disable ADOM.
     */
    adomStatus?: pulumi.Input<string>;
    /**
     * Enable/Disable FortiAnalyzer feature.
     */
    fortianalyzerStatus?: pulumi.Input<string>;
    /**
     * Hostname.
     */
    hostname?: pulumi.Input<string>;
    /**
     * TimeZone. 00 - (GMT-12:00) Eniwetak, Kwajalein.01 - (GMT-11:00) Midway Island, Samoa.02 - (GMT-10:00) Hawaii.03 - (GMT-9:00) Alaska.04 - (GMT-8:00) Pacific Time (US & Canada).05 - (GMT-7:00) Arizona.06 - (GMT-7:00) Mountain Time (US & Canada).07 - (GMT-6:00) Central America.08 - (GMT-6:00) Central Time (US & Canada).09 - (GMT-6:00) Mexico City.10 - (GMT-6:00) Saskatchewan.11 - (GMT-5:00) Bogota, Lima, Quito.12 - (GMT-5:00) Eastern Time (US & Canada).13 - (GMT-5:00) Indiana (East).14 - (GMT-4:00) Atlantic Time (Canada).15 - (GMT-4:00) La Paz.16 - (GMT-4:00) Santiago.17 - (GMT-3:30) Newfoundland.18 - (GMT-3:00) Brasilia.19 - (GMT-3:00) Buenos Aires, Georgetown.20 - (GMT-3:00) Nuuk (Greenland).21 - (GMT-2:00) Mid-Atlantic.22 - (GMT-1:00) Azores.23 - (GMT-1:00) Cape Verde Is.24 - (GMT) Monrovia.25 - (GMT) Greenwich Mean Time:Dublin, Edinburgh, Lisbon, London.26 - (GMT+1:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna.27 - (GMT+1:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague.28 - (GMT+1:00) Brussels, Copenhagen, Madrid, Paris.29 - (GMT+1:00) Sarajevo, Skopje, Warsaw, Zagreb.30 - (GMT+1:00) West Central Africa.31 - (GMT+2:00) Athens, Sofia, Vilnius.32 - (GMT+2:00) Bucharest.33 - (GMT+2:00) Cairo.34 - (GMT+2:00) Harare, Pretoria.35 - (GMT+2:00) Helsinki, Riga,Tallinn.36 - (GMT+2:00) Jerusalem.37 - (GMT+3:00) Baghdad.38 - (GMT+3:00) Kuwait, Riyadh.39 - (GMT+3:00) St.Petersburg, Volgograd.40 - (GMT+3:00) Nairobi.41 - (GMT+3:30) Tehran.42 - (GMT+4:00) Abu Dhabi, Muscat.43 - (GMT+4:00) Baku.44 - (GMT+4:30) Kabul.45 - (GMT+5:00) Ekaterinburg.46 - (GMT+5:00) Islamabad, Karachi,Tashkent.47 - (GMT+5:30) Calcutta, Chennai, Mumbai, New Delhi.48 - (GMT+5:45) Kathmandu.49 - (GMT+6:00) Almaty, Novosibirsk.50 - (GMT+6:00) Astana, Dhaka.51 - (GMT+6:00) Sri Jayawardenapura.52 - (GMT+6:30) Rangoon.53 - (GMT+7:00) Bangkok, Hanoi, Jakarta.54 - (GMT+7:00) Krasnoyarsk.55 - (GMT+8:00) Beijing,ChongQing, HongKong,Urumqi.56 - (GMT+8:00) Irkutsk, Ulaanbaatar.57 - (GMT+8:00) Kuala Lumpur, Singapore.58 - (GMT+8:00) Perth.59 - (GMT+8:00) Taipei.60 - (GMT+9:00) Osaka, Sapporo, Tokyo, Seoul.61 - (GMT+9:00) Yakutsk.62 - (GMT+9:30) Adelaide.63 - (GMT+9:30) Darwin.64 - (GMT+10:00) Brisbane.65 - (GMT+10:00) Canberra, Melbourne, Sydney.66 - (GMT+10:00) Guam, Port Moresby.67 - (GMT+10:00) Hobart.68 - (GMT+10:00) Vladivostok.69 - (GMT+11:00) Magadan.70 - (GMT+11:00) Solomon Is., New Caledonia.71 - (GMT+12:00) Auckland, Wellington.72 - (GMT+12:00) Fiji, Kamchatka, Marshall Is.73 - (GMT+13:00) Nuku'alofa.74 - (GMT-4:30) Caracas.75 - (GMT+1:00) Namibia.76 - (GMT-5:00) Brazil-Acre.77 - (GMT-4:00) Brazil-West.78 - (GMT-3:00) Brazil-East.79 - (GMT-2:00) Brazil-DeNoronha.80 - (GMT+14:00) Kiritimati.81 - (GMT-7:00) Baja California Sur, Chihuahua.82 - (GMT+12:45) Chatham Islands.83 - (GMT+3:00) Minsk.84 - (GMT+13:00) Samoa.85 - (GMT+3:00) Istanbul.86 - (GMT-4:00) Paraguay.87 - (GMT) Casablanca.88 - (GMT+3:00) Moscow.89 - (GMT) Greenwich Mean Time.
     */
    timezone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemGlobal resource.
 */
export interface SystemGlobalArgs {
    /**
     * Adom Mode.
     */
    adomMode?: pulumi.Input<string>;
    /**
     * Enable/Disable ADOM.
     */
    adomStatus?: pulumi.Input<string>;
    /**
     * Enable/Disable FortiAnalyzer feature.
     */
    fortianalyzerStatus?: pulumi.Input<string>;
    /**
     * Hostname.
     */
    hostname?: pulumi.Input<string>;
    /**
     * TimeZone. 00 - (GMT-12:00) Eniwetak, Kwajalein.01 - (GMT-11:00) Midway Island, Samoa.02 - (GMT-10:00) Hawaii.03 - (GMT-9:00) Alaska.04 - (GMT-8:00) Pacific Time (US & Canada).05 - (GMT-7:00) Arizona.06 - (GMT-7:00) Mountain Time (US & Canada).07 - (GMT-6:00) Central America.08 - (GMT-6:00) Central Time (US & Canada).09 - (GMT-6:00) Mexico City.10 - (GMT-6:00) Saskatchewan.11 - (GMT-5:00) Bogota, Lima, Quito.12 - (GMT-5:00) Eastern Time (US & Canada).13 - (GMT-5:00) Indiana (East).14 - (GMT-4:00) Atlantic Time (Canada).15 - (GMT-4:00) La Paz.16 - (GMT-4:00) Santiago.17 - (GMT-3:30) Newfoundland.18 - (GMT-3:00) Brasilia.19 - (GMT-3:00) Buenos Aires, Georgetown.20 - (GMT-3:00) Nuuk (Greenland).21 - (GMT-2:00) Mid-Atlantic.22 - (GMT-1:00) Azores.23 - (GMT-1:00) Cape Verde Is.24 - (GMT) Monrovia.25 - (GMT) Greenwich Mean Time:Dublin, Edinburgh, Lisbon, London.26 - (GMT+1:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna.27 - (GMT+1:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague.28 - (GMT+1:00) Brussels, Copenhagen, Madrid, Paris.29 - (GMT+1:00) Sarajevo, Skopje, Warsaw, Zagreb.30 - (GMT+1:00) West Central Africa.31 - (GMT+2:00) Athens, Sofia, Vilnius.32 - (GMT+2:00) Bucharest.33 - (GMT+2:00) Cairo.34 - (GMT+2:00) Harare, Pretoria.35 - (GMT+2:00) Helsinki, Riga,Tallinn.36 - (GMT+2:00) Jerusalem.37 - (GMT+3:00) Baghdad.38 - (GMT+3:00) Kuwait, Riyadh.39 - (GMT+3:00) St.Petersburg, Volgograd.40 - (GMT+3:00) Nairobi.41 - (GMT+3:30) Tehran.42 - (GMT+4:00) Abu Dhabi, Muscat.43 - (GMT+4:00) Baku.44 - (GMT+4:30) Kabul.45 - (GMT+5:00) Ekaterinburg.46 - (GMT+5:00) Islamabad, Karachi,Tashkent.47 - (GMT+5:30) Calcutta, Chennai, Mumbai, New Delhi.48 - (GMT+5:45) Kathmandu.49 - (GMT+6:00) Almaty, Novosibirsk.50 - (GMT+6:00) Astana, Dhaka.51 - (GMT+6:00) Sri Jayawardenapura.52 - (GMT+6:30) Rangoon.53 - (GMT+7:00) Bangkok, Hanoi, Jakarta.54 - (GMT+7:00) Krasnoyarsk.55 - (GMT+8:00) Beijing,ChongQing, HongKong,Urumqi.56 - (GMT+8:00) Irkutsk, Ulaanbaatar.57 - (GMT+8:00) Kuala Lumpur, Singapore.58 - (GMT+8:00) Perth.59 - (GMT+8:00) Taipei.60 - (GMT+9:00) Osaka, Sapporo, Tokyo, Seoul.61 - (GMT+9:00) Yakutsk.62 - (GMT+9:30) Adelaide.63 - (GMT+9:30) Darwin.64 - (GMT+10:00) Brisbane.65 - (GMT+10:00) Canberra, Melbourne, Sydney.66 - (GMT+10:00) Guam, Port Moresby.67 - (GMT+10:00) Hobart.68 - (GMT+10:00) Vladivostok.69 - (GMT+11:00) Magadan.70 - (GMT+11:00) Solomon Is., New Caledonia.71 - (GMT+12:00) Auckland, Wellington.72 - (GMT+12:00) Fiji, Kamchatka, Marshall Is.73 - (GMT+13:00) Nuku'alofa.74 - (GMT-4:30) Caracas.75 - (GMT+1:00) Namibia.76 - (GMT-5:00) Brazil-Acre.77 - (GMT-4:00) Brazil-West.78 - (GMT-3:00) Brazil-East.79 - (GMT-2:00) Brazil-DeNoronha.80 - (GMT+14:00) Kiritimati.81 - (GMT-7:00) Baja California Sur, Chihuahua.82 - (GMT+12:45) Chatham Islands.83 - (GMT+3:00) Minsk.84 - (GMT+13:00) Samoa.85 - (GMT+3:00) Istanbul.86 - (GMT-4:00) Paraguay.87 - (GMT) Casablanca.88 - (GMT+3:00) Moscow.89 - (GMT) Greenwich Mean Time.
     */
    timezone?: pulumi.Input<string>;
}
