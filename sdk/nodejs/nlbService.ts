// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manage Exoscale [Network Load Balancer (NLB)](https://community.exoscale.com/documentation/compute/network-load-balancer/) Services.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumiverse/exoscale";
 *
 * const myNlb = new exoscale.Nlb("my_nlb", {
 *     zone: "ch-gva-2",
 *     name: "my-nlb",
 * });
 * const myNlbService = new exoscale.NlbService("my_nlb_service", {
 *     nlbId: myNlb.id,
 *     zone: myNlb.zone,
 *     name: "my-nlb-service",
 *     instancePoolId: myInstancePool.id,
 *     protocol: "tcp",
 *     port: 443,
 *     targetPort: 8443,
 *     strategy: "round-robin",
 *     healthchecks: [{
 *         mode: "https",
 *         port: 8443,
 *         uri: "/healthz",
 *         tlsSni: "example.net",
 *         interval: 5,
 *         timeout: 3,
 *         retries: 1,
 *     }],
 * });
 * ```
 *
 * Please refer to the examples
 * directory for complete configuration examples.
 *
 * ## Import
 *
 * An existing NLB service may be imported by `<nlb-ID>/<service-ID>@<zone>`:
 *
 * ```sh
 * $ pulumi import exoscale:index/nlbService:NlbService \
 * ```
 *
 *   exoscale_nlb_service.my_nlb_service \
 *
 *   f81d4fae-7dec-11d0-a765-00a0c91e6bf6/9ecc6b8b-73d4-4211-8ced-f7f29bb79524@ch-gva-2
 */
export class NlbService extends pulumi.CustomResource {
    /**
     * Get an existing NlbService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NlbServiceState, opts?: pulumi.CustomResourceOptions): NlbService {
        return new NlbService(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'exoscale:index/nlbService:NlbService';

    /**
     * Returns true if the given object is an instance of NlbService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NlbService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NlbService.__pulumiType;
    }

    /**
     * A free-form text describing the NLB service.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The service health checking configuration.
     */
    public readonly healthchecks!: pulumi.Output<outputs.NlbServiceHealthcheck[]>;
    /**
     * ❗ The exoscale*instance*pool (ID) to forward traffic to.
     */
    public readonly instancePoolId!: pulumi.Output<string>;
    /**
     * The NLB service name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ❗ The parent exoscale.Nlb ID.
     */
    public readonly nlbId!: pulumi.Output<string>;
    /**
     * The healthcheck port.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * The protocol (`tcp`|`udp`; default: `tcp`).
     */
    public readonly protocol!: pulumi.Output<string | undefined>;
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The strategy (`round-robin`|`source-hash`; default: `round-robin`).
     */
    public readonly strategy!: pulumi.Output<string | undefined>;
    /**
     * The (TCP/UDP) port to forward traffic to (on target instance pool members).
     */
    public readonly targetPort!: pulumi.Output<number>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a NlbService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NlbServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NlbServiceArgs | NlbServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NlbServiceState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["healthchecks"] = state ? state.healthchecks : undefined;
            resourceInputs["instancePoolId"] = state ? state.instancePoolId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nlbId"] = state ? state.nlbId : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["strategy"] = state ? state.strategy : undefined;
            resourceInputs["targetPort"] = state ? state.targetPort : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as NlbServiceArgs | undefined;
            if ((!args || args.healthchecks === undefined) && !opts.urn) {
                throw new Error("Missing required property 'healthchecks'");
            }
            if ((!args || args.instancePoolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instancePoolId'");
            }
            if ((!args || args.nlbId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nlbId'");
            }
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
            }
            if ((!args || args.targetPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetPort'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["healthchecks"] = args ? args.healthchecks : undefined;
            resourceInputs["instancePoolId"] = args ? args.instancePoolId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nlbId"] = args ? args.nlbId : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["strategy"] = args ? args.strategy : undefined;
            resourceInputs["targetPort"] = args ? args.targetPort : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NlbService.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NlbService resources.
 */
export interface NlbServiceState {
    /**
     * A free-form text describing the NLB service.
     */
    description?: pulumi.Input<string>;
    /**
     * The service health checking configuration.
     */
    healthchecks?: pulumi.Input<pulumi.Input<inputs.NlbServiceHealthcheck>[]>;
    /**
     * ❗ The exoscale*instance*pool (ID) to forward traffic to.
     */
    instancePoolId?: pulumi.Input<string>;
    /**
     * The NLB service name.
     */
    name?: pulumi.Input<string>;
    /**
     * ❗ The parent exoscale.Nlb ID.
     */
    nlbId?: pulumi.Input<string>;
    /**
     * The healthcheck port.
     */
    port?: pulumi.Input<number>;
    /**
     * The protocol (`tcp`|`udp`; default: `tcp`).
     */
    protocol?: pulumi.Input<string>;
    state?: pulumi.Input<string>;
    /**
     * The strategy (`round-robin`|`source-hash`; default: `round-robin`).
     */
    strategy?: pulumi.Input<string>;
    /**
     * The (TCP/UDP) port to forward traffic to (on target instance pool members).
     */
    targetPort?: pulumi.Input<number>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NlbService resource.
 */
export interface NlbServiceArgs {
    /**
     * A free-form text describing the NLB service.
     */
    description?: pulumi.Input<string>;
    /**
     * The service health checking configuration.
     */
    healthchecks: pulumi.Input<pulumi.Input<inputs.NlbServiceHealthcheck>[]>;
    /**
     * ❗ The exoscale*instance*pool (ID) to forward traffic to.
     */
    instancePoolId: pulumi.Input<string>;
    /**
     * The NLB service name.
     */
    name?: pulumi.Input<string>;
    /**
     * ❗ The parent exoscale.Nlb ID.
     */
    nlbId: pulumi.Input<string>;
    /**
     * The healthcheck port.
     */
    port: pulumi.Input<number>;
    /**
     * The protocol (`tcp`|`udp`; default: `tcp`).
     */
    protocol?: pulumi.Input<string>;
    /**
     * The strategy (`round-robin`|`source-hash`; default: `round-robin`).
     */
    strategy?: pulumi.Input<string>;
    /**
     * The (TCP/UDP) port to forward traffic to (on target instance pool members).
     */
    targetPort: pulumi.Input<number>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: pulumi.Input<string>;
}
