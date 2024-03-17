import * as pulumi from "@pulumi/pulumi";
import * as fortios from "@pulumiverse/fortios";

class M1 extends pulumi.ComponentResource {
  constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
    super('components:index:M1', name, {}, opts)
    const trname1 = new fortios.firewall.Policy(
      `${name}-trname1`,
      {
        action: 'accept',
        name: 'policy1',
        policyid: 1,
        status: 'enable',
        dstaddrs: [
          {
            name: 'all',
          },
        ],
        dstintfs: [
          {
            name: 'port4',
          },
        ],
        services: [
          {
            name: 'HTTP',
          },
        ],
        srcaddrs: [
          {
            name: 'all',
          },
        ],
        srcintfs: [
          {
            name: 'port3',
          },
        ],
      },
      {
        parent: this,
      },
    )

    const trname2 = new fortios.firewall.Policy(
      `${name}-trname2`,
      {
        action: 'accept',
        name: 'policy2',
        policyid: 2,
        status: 'enable',
        dstaddrs: [
          {
            name: 'swscan.apple.com',
          },
        ],
        dstintfs: [
          {
            name: 'port2',
          },
        ],
        services: [
          {
            name: 'AFS3',
          },
        ],
        srcaddrs: [
          {
            name: 'myaddress',
          },
        ],
        srcintfs: [
          {
            name: 'port1',
          },
        ],
      },
      {
        parent: this,
      },
    )

    const trname3 = new fortios.firewall.Policy(
      `${name}-trname3`,
      {
        action: 'accept',
        name: 'policy3',
        policyid: 3,
        status: 'enable',
        dstaddrs: [
          {
            name: 'all',
          },
        ],
        dstintfs: [
          {
            name: 'port2',
          },
        ],
        services: [
          {
            name: 'ALL_TCP',
          },
        ],
        srcaddrs: [
          {
            name: 'myaddress',
          },
        ],
        srcintfs: [
          {
            name: 'port1',
          },
        ],
      },
      {
        parent: this,
      },
    )
  }
}

class M2 extends pulumi.ComponentResource {
  constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
    super('components:index:M2', name, {}, opts)
    const trname4 = new fortios.firewall.Policy(
      `${name}-trname4`,
      {
        action: 'accept',
        name: 'policy4',
        policyid: 4,
        status: 'enable',
        dstaddrs: [
          {
            name: 'swscan.apple.com',
          },
        ],
        dstintfs: [
          {
            name: 'port2',
          },
        ],
        services: [
          {
            name: 'ALL_ICMP',
          },
        ],
        srcaddrs: [
          {
            name: 'google-play',
          },
        ],
        srcintfs: [
          {
            name: 'port1',
          },
        ],
      },
      {
        parent: this,
      },
    )

    const trname5 = new fortios.firewall.Policy(
      `${name}-trname5`,
      {
        action: 'accept',
        name: 'policy5',
        policyid: 5,
        status: 'enable',
        dstaddrs: [
          {
            name: 'autoupdate.opera.com',
          },
        ],
        dstintfs: [
          {
            name: 'port3',
          },
        ],
        services: [
          {
            name: 'ALL_TCP',
          },
        ],
        srcaddrs: [
          {
            name: 'myaddress',
          },
        ],
        srcintfs: [
          {
            name: 'port1',
          },
        ],
      },
      {
        parent: this,
      },
    )

    const trname6 = new fortios.firewall.Policy(
      `${name}-trname6`,
      {
        action: 'accept',
        name: 'policy6',
        policyid: 6,
        status: 'enable',
        dstaddrs: [
          {
            name: 'all',
          },
        ],
        dstintfs: [
          {
            name: 'port2',
          },
        ],
        services: [
          {
            name: 'ALL_TCP',
          },
        ],
        srcaddrs: [
          {
            name: 'autoupdate.opera.com',
          },
        ],
        srcintfs: [
          {
            name: 'port2',
          },
        ],
      },
      {
        parent: this,
      },
    )
  }
}

class M3 extends pulumi.ComponentResource {
  constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
    super('components:index:M3', name, {}, opts)
    const trname7 = new fortios.firewall.Policy(
      `${name}-trname7`,
      {
        action: 'accept',
        name: 'policy7',
        policyid: 7,
        status: 'enable',
        dstaddrs: [
          {
            name: 'swscan.apple.com',
          },
        ],
        dstintfs: [
          {
            name: 'port2',
          },
        ],
        services: [
          {
            name: 'ALL_TCP',
          },
        ],
        srcaddrs: [
          {
            name: 'google-play',
          },
        ],
        srcintfs: [
          {
            name: 'port1',
          },
        ],
      },
      {
        parent: this,
      },
    )

    const trname8 = new fortios.firewall.Policy(
      `${name}-trname8`,
      {
        action: 'accept',
        name: 'policy8',
        policyid: 8,
        status: 'enable',
        dstaddrs: [
          {
            name: 'autoupdate.opera.com',
          },
        ],
        dstintfs: [
          {
            name: 'port3',
          },
        ],
        services: [
          {
            name: 'ALL_ICMP',
          },
        ],
        srcaddrs: [
          {
            name: 'myaddress',
          },
        ],
        srcintfs: [
          {
            name: 'port1',
          },
        ],
      },
      {
        parent: this,
      },
    )

    const trname9 = new fortios.firewall.Policy(
      `${name}-trname9`,
      {
        action: 'accept',
        name: 'policy9',
        policyid: 9,
        status: 'enable',
        dstaddrs: [
          {
            name: 'all',
          },
        ],
        dstintfs: [
          {
            name: 'port2',
          },
        ],
        services: [
          {
            name: 'ALL_ICMP',
          },
        ],
        srcaddrs: [
          {
            name: 'autoupdate.opera.com',
          },
        ],
        srcintfs: [
          {
            name: 'port2',
          },
        ],
      },
      {
        parent: this,
      },
    )
  }
}

/*****
 * MAIN
 */
const m1 = new M1('m1')
const m2 = new M2('m2', {
  dependsOn: [m1],
})
const m3 = new M3('m3', {
  dependsOn: [m2],
})
