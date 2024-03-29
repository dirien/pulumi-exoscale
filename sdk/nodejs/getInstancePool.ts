// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Fetch Exoscale [Instance Pools](https://community.exoscale.com/documentation/compute/instance-pools/) data.
 *
 * Corresponding resource: exoscale_instance_pool.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumi/exoscale";
 *
 * const myInstancePool = exoscale.getInstancePool({
 *     zone: "ch-gva-2",
 *     name: "my-instance-pool",
 * });
 * export const myInstancePoolId = myInstancePool.then(myInstancePool => myInstancePool.id);
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * Please refer to the examples
 * directory for complete configuration examples.
 */
export function getInstancePool(args: GetInstancePoolArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancePoolResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("exoscale:index/getInstancePool:getInstancePool", {
        "id": args.id,
        "labels": args.labels,
        "name": args.name,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstancePool.
 */
export interface GetInstancePoolArgs {
    id?: string;
    /**
     * A map of key/value labels.
     */
    labels?: {[key: string]: string};
    name?: string;
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: string;
}

/**
 * A collection of values returned by getInstancePool.
 */
export interface GetInstancePoolResult {
    /**
     * The list of attached exoscale*anti*affinity_group (IDs).
     */
    readonly affinityGroupIds: string[];
    /**
     * The deploy target ID.
     */
    readonly deployTargetId: string;
    /**
     * The instance pool description.
     */
    readonly description: string;
    /**
     * The managed instances disk size.
     */
    readonly diskSize: number;
    /**
     * The list of attached exoscale*elastic*ip (IDs).
     */
    readonly elasticIpIds: string[];
    /**
     * The instance pool ID to match (conflicts with `name`).
     */
    readonly id?: string;
    /**
     * The string used to prefix the managed instances name.
     */
    readonly instancePrefix: string;
    /**
     * The managed instances type.
     */
    readonly instanceType: string;
    /**
     * The list of managed instances. Structure is documented below.
     */
    readonly instances: outputs.GetInstancePoolInstance[];
    /**
     * Whether IPv6 is enabled on managed instances.
     */
    readonly ipv6: boolean;
    /**
     * The exoscale*ssh*key (name) authorized on the managed instances.
     */
    readonly keyPair: string;
    /**
     * A map of key/value labels.
     */
    readonly labels?: {[key: string]: string};
    /**
     * The pool name to match (conflicts with `id`).
     */
    readonly name?: string;
    /**
     * The list of attached exoscale*private*network (IDs).
     */
    readonly networkIds: string[];
    /**
     * The list of attached exoscale*security*group (IDs).
     */
    readonly securityGroupIds: string[];
    /**
     * The number managed instances.
     */
    readonly size: number;
    /**
     * The pool state.
     */
    readonly state: string;
    /**
     * The managed instances exoscale.getTemplate ID.
     */
    readonly templateId: string;
    /**
     * [cloud-init](http://cloudinit.readthedocs.io/en/latest/) configuration.
     */
    readonly userData: string;
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    readonly zone: string;
}
/**
 * Fetch Exoscale [Instance Pools](https://community.exoscale.com/documentation/compute/instance-pools/) data.
 *
 * Corresponding resource: exoscale_instance_pool.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumi/exoscale";
 *
 * const myInstancePool = exoscale.getInstancePool({
 *     zone: "ch-gva-2",
 *     name: "my-instance-pool",
 * });
 * export const myInstancePoolId = myInstancePool.then(myInstancePool => myInstancePool.id);
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * Please refer to the examples
 * directory for complete configuration examples.
 */
export function getInstancePoolOutput(args: GetInstancePoolOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstancePoolResult> {
    return pulumi.output(args).apply((a: any) => getInstancePool(a, opts))
}

/**
 * A collection of arguments for invoking getInstancePool.
 */
export interface GetInstancePoolOutputArgs {
    id?: pulumi.Input<string>;
    /**
     * A map of key/value labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    name?: pulumi.Input<string>;
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: pulumi.Input<string>;
}
