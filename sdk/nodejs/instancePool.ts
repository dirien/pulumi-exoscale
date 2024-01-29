// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manage Exoscale [Instance Pools](https://community.exoscale.com/documentation/compute/instance-pools/).
 *
 * Corresponding data sources: exoscale_instance_pool, exoscale_instance_pool_list.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumi/exoscale";
 * import * as exoscale from "@pulumiverse/exoscale";
 *
 * const myTemplate = exoscale.getTemplate({
 *     zone: "ch-gva-2",
 *     name: "Linux Ubuntu 22.04 LTS 64-bit",
 * });
 * const myInstancePool = new exoscale.InstancePool("myInstancePool", {
 *     zone: "ch-gva-2",
 *     templateId: myTemplate.then(myTemplate => myTemplate.id),
 *     instanceType: "standard.medium",
 *     diskSize: 10,
 *     size: 3,
 * });
 * ```
 *
 * Please refer to the examples
 * directory for complete configuration examples.
 *
 * ## Import
 *
 * An existing instance pool may be imported by `<ID>@<zone>`
 *
 * ```sh
 *  $ pulumi import exoscale:index/instancePool:InstancePool \
 * ```
 *
 *  exoscale_instance_pool.my_instance_pool \
 *
 *  f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2
 */
export class InstancePool extends pulumi.CustomResource {
    /**
     * Get an existing InstancePool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstancePoolState, opts?: pulumi.CustomResourceOptions): InstancePool {
        return new InstancePool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'exoscale:index/instancePool:InstancePool';

    /**
     * Returns true if the given object is an instance of InstancePool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstancePool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstancePool.__pulumiType;
    }

    /**
     * A list of exoscale*anti*affinity_group (IDs; may only be set at creation time).
     */
    public readonly affinityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * A deploy target ID.
     */
    public readonly deployTargetId!: pulumi.Output<string | undefined>;
    /**
     * A free-form text describing the pool.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The managed instances disk size (GiB).
     */
    public readonly diskSize!: pulumi.Output<number>;
    /**
     * A list of exoscale*elastic*ip (IDs).
     */
    public readonly elasticIpIds!: pulumi.Output<string[] | undefined>;
    /**
     * The string used to prefix managed instances name (default: `pool`).
     */
    public readonly instancePrefix!: pulumi.Output<string | undefined>;
    /**
     * The managed compute instances type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * The list of managed instances. Structure is documented below.
     */
    public readonly instances!: pulumi.Output<outputs.InstancePoolInstance[]>;
    /**
     * Enable IPv6 on managed instances (boolean; default: `false`).
     */
    public readonly ipv6!: pulumi.Output<boolean | undefined>;
    /**
     * The exoscale*ssh*key (name) to authorize in the managed instances.
     */
    public readonly keyPair!: pulumi.Output<string | undefined>;
    /**
     * A map of key/value labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The instance name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of exoscale*private*network (IDs).
     */
    public readonly networkIds!: pulumi.Output<string[] | undefined>;
    /**
     * A list of exoscale*security*group (IDs).
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * The managed instances type. Please use the `instanceType` argument instead.
     *
     * @deprecated This attribute has been replaced by "instance_type".
     */
    public readonly serviceOffering!: pulumi.Output<string>;
    /**
     * The number of managed instances.
     */
    public readonly size!: pulumi.Output<number>;
    public readonly state!: pulumi.Output<string>;
    /**
     * The exoscale.getTemplate (ID) to use when creating the managed instances.
     */
    public readonly templateId!: pulumi.Output<string>;
    /**
     * [cloud-init](http://cloudinit.readthedocs.io/) configuration to apply to the managed instances.
     */
    public readonly userData!: pulumi.Output<string | undefined>;
    /**
     * The list of managed instances (IDs). Please use the `instances.*.id` attribute instead.
     *
     * @deprecated Use the instances exported attribute instead.
     */
    public readonly virtualMachines!: pulumi.Output<string[]>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a InstancePool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstancePoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstancePoolArgs | InstancePoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstancePoolState | undefined;
            resourceInputs["affinityGroupIds"] = state ? state.affinityGroupIds : undefined;
            resourceInputs["deployTargetId"] = state ? state.deployTargetId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["diskSize"] = state ? state.diskSize : undefined;
            resourceInputs["elasticIpIds"] = state ? state.elasticIpIds : undefined;
            resourceInputs["instancePrefix"] = state ? state.instancePrefix : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["instances"] = state ? state.instances : undefined;
            resourceInputs["ipv6"] = state ? state.ipv6 : undefined;
            resourceInputs["keyPair"] = state ? state.keyPair : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkIds"] = state ? state.networkIds : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["serviceOffering"] = state ? state.serviceOffering : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["templateId"] = state ? state.templateId : undefined;
            resourceInputs["userData"] = state ? state.userData : undefined;
            resourceInputs["virtualMachines"] = state ? state.virtualMachines : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstancePoolArgs | undefined;
            if ((!args || args.size === undefined) && !opts.urn) {
                throw new Error("Missing required property 'size'");
            }
            if ((!args || args.templateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateId'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["affinityGroupIds"] = args ? args.affinityGroupIds : undefined;
            resourceInputs["deployTargetId"] = args ? args.deployTargetId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["diskSize"] = args ? args.diskSize : undefined;
            resourceInputs["elasticIpIds"] = args ? args.elasticIpIds : undefined;
            resourceInputs["instancePrefix"] = args ? args.instancePrefix : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["instances"] = args ? args.instances : undefined;
            resourceInputs["ipv6"] = args ? args.ipv6 : undefined;
            resourceInputs["keyPair"] = args ? args.keyPair : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkIds"] = args ? args.networkIds : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["serviceOffering"] = args ? args.serviceOffering : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["templateId"] = args ? args.templateId : undefined;
            resourceInputs["userData"] = args ? args.userData : undefined;
            resourceInputs["virtualMachines"] = args ? args.virtualMachines : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstancePool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstancePool resources.
 */
export interface InstancePoolState {
    /**
     * A list of exoscale*anti*affinity_group (IDs; may only be set at creation time).
     */
    affinityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A deploy target ID.
     */
    deployTargetId?: pulumi.Input<string>;
    /**
     * A free-form text describing the pool.
     */
    description?: pulumi.Input<string>;
    /**
     * The managed instances disk size (GiB).
     */
    diskSize?: pulumi.Input<number>;
    /**
     * A list of exoscale*elastic*ip (IDs).
     */
    elasticIpIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The string used to prefix managed instances name (default: `pool`).
     */
    instancePrefix?: pulumi.Input<string>;
    /**
     * The managed compute instances type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The list of managed instances. Structure is documented below.
     */
    instances?: pulumi.Input<pulumi.Input<inputs.InstancePoolInstance>[]>;
    /**
     * Enable IPv6 on managed instances (boolean; default: `false`).
     */
    ipv6?: pulumi.Input<boolean>;
    /**
     * The exoscale*ssh*key (name) to authorize in the managed instances.
     */
    keyPair?: pulumi.Input<string>;
    /**
     * A map of key/value labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The instance name.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of exoscale*private*network (IDs).
     */
    networkIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of exoscale*security*group (IDs).
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The managed instances type. Please use the `instanceType` argument instead.
     *
     * @deprecated This attribute has been replaced by "instance_type".
     */
    serviceOffering?: pulumi.Input<string>;
    /**
     * The number of managed instances.
     */
    size?: pulumi.Input<number>;
    state?: pulumi.Input<string>;
    /**
     * The exoscale.getTemplate (ID) to use when creating the managed instances.
     */
    templateId?: pulumi.Input<string>;
    /**
     * [cloud-init](http://cloudinit.readthedocs.io/) configuration to apply to the managed instances.
     */
    userData?: pulumi.Input<string>;
    /**
     * The list of managed instances (IDs). Please use the `instances.*.id` attribute instead.
     *
     * @deprecated Use the instances exported attribute instead.
     */
    virtualMachines?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstancePool resource.
 */
export interface InstancePoolArgs {
    /**
     * A list of exoscale*anti*affinity_group (IDs; may only be set at creation time).
     */
    affinityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A deploy target ID.
     */
    deployTargetId?: pulumi.Input<string>;
    /**
     * A free-form text describing the pool.
     */
    description?: pulumi.Input<string>;
    /**
     * The managed instances disk size (GiB).
     */
    diskSize?: pulumi.Input<number>;
    /**
     * A list of exoscale*elastic*ip (IDs).
     */
    elasticIpIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The string used to prefix managed instances name (default: `pool`).
     */
    instancePrefix?: pulumi.Input<string>;
    /**
     * The managed compute instances type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The list of managed instances. Structure is documented below.
     */
    instances?: pulumi.Input<pulumi.Input<inputs.InstancePoolInstance>[]>;
    /**
     * Enable IPv6 on managed instances (boolean; default: `false`).
     */
    ipv6?: pulumi.Input<boolean>;
    /**
     * The exoscale*ssh*key (name) to authorize in the managed instances.
     */
    keyPair?: pulumi.Input<string>;
    /**
     * A map of key/value labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The instance name.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of exoscale*private*network (IDs).
     */
    networkIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of exoscale*security*group (IDs).
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The managed instances type. Please use the `instanceType` argument instead.
     *
     * @deprecated This attribute has been replaced by "instance_type".
     */
    serviceOffering?: pulumi.Input<string>;
    /**
     * The number of managed instances.
     */
    size: pulumi.Input<number>;
    state?: pulumi.Input<string>;
    /**
     * The exoscale.getTemplate (ID) to use when creating the managed instances.
     */
    templateId: pulumi.Input<string>;
    /**
     * [cloud-init](http://cloudinit.readthedocs.io/) configuration to apply to the managed instances.
     */
    userData?: pulumi.Input<string>;
    /**
     * The list of managed instances (IDs). Please use the `instances.*.id` attribute instead.
     *
     * @deprecated Use the instances exported attribute instead.
     */
    virtualMachines?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: pulumi.Input<string>;
}
