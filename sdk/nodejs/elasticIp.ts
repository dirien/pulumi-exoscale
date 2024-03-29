// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manage Exoscale [Elastic IPs (EIP)](https://community.exoscale.com/documentation/compute/eip/).
 *
 * Corresponding data source: exoscale_elastic_ip.
 *
 * ## Example Usage
 *
 * *Unmanaged* EIPv4:
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumiverse/exoscale";
 *
 * const myElasticIp = new exoscale.ElasticIp("myElasticIp", {zone: "ch-gva-2"});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * *Managed* EIPv6:
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumiverse/exoscale";
 *
 * const myManagedElasticIp = new exoscale.ElasticIp("myManagedElasticIp", {
 *     addressFamily: "inet6",
 *     healthcheck: {
 *         interval: 5,
 *         mode: "https",
 *         port: 443,
 *         strikesFail: 3,
 *         strikesOk: 2,
 *         timeout: 3,
 *         tlsSni: "example.net",
 *         uri: "/health",
 *     },
 *     reverseDns: "example.net",
 *     zone: "ch-gva-2",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * Please refer to the examples
 * directory for complete configuration examples.
 *
 * ## Import
 *
 * An existing Elastic IP (EIP) may be imported by `<ID>@<zone>`:
 *
 * ```sh
 * $ pulumi import exoscale:index/elasticIp:ElasticIp \
 * ```
 *
 *   exoscale_elastic_ip.my_elastic_ip \
 *
 *   f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2
 */
export class ElasticIp extends pulumi.CustomResource {
    /**
     * Get an existing ElasticIp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ElasticIpState, opts?: pulumi.CustomResourceOptions): ElasticIp {
        return new ElasticIp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'exoscale:index/elasticIp:ElasticIp';

    /**
     * Returns true if the given object is an instance of ElasticIp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ElasticIp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ElasticIp.__pulumiType;
    }

    /**
     * ❗ The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
     */
    public readonly addressFamily!: pulumi.Output<string>;
    /**
     * The Elastic IP (EIP) CIDR.
     */
    public /*out*/ readonly cidr!: pulumi.Output<string>;
    /**
     * A free-form text describing the Elastic IP (EIP).
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Healthcheck configuration for *managed* EIPs. It can not be added to an existing *Unmanaged* EIP.
     */
    public readonly healthcheck!: pulumi.Output<outputs.ElasticIpHealthcheck>;
    /**
     * The Elastic IP (EIP) IPv4 or IPv6 address.
     */
    public /*out*/ readonly ipAddress!: pulumi.Output<string>;
    /**
     * A map of key/value labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Domain name for reverse DNS record.
     */
    public readonly reverseDns!: pulumi.Output<string | undefined>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a ElasticIp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ElasticIpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ElasticIpArgs | ElasticIpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ElasticIpState | undefined;
            resourceInputs["addressFamily"] = state ? state.addressFamily : undefined;
            resourceInputs["cidr"] = state ? state.cidr : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["healthcheck"] = state ? state.healthcheck : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["reverseDns"] = state ? state.reverseDns : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as ElasticIpArgs | undefined;
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["addressFamily"] = args ? args.addressFamily : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["healthcheck"] = args ? args.healthcheck : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["reverseDns"] = args ? args.reverseDns : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["cidr"] = undefined /*out*/;
            resourceInputs["ipAddress"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ElasticIp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ElasticIp resources.
 */
export interface ElasticIpState {
    /**
     * ❗ The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
     */
    addressFamily?: pulumi.Input<string>;
    /**
     * The Elastic IP (EIP) CIDR.
     */
    cidr?: pulumi.Input<string>;
    /**
     * A free-form text describing the Elastic IP (EIP).
     */
    description?: pulumi.Input<string>;
    /**
     * Healthcheck configuration for *managed* EIPs. It can not be added to an existing *Unmanaged* EIP.
     */
    healthcheck?: pulumi.Input<inputs.ElasticIpHealthcheck>;
    /**
     * The Elastic IP (EIP) IPv4 or IPv6 address.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * A map of key/value labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Domain name for reverse DNS record.
     */
    reverseDns?: pulumi.Input<string>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ElasticIp resource.
 */
export interface ElasticIpArgs {
    /**
     * ❗ The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
     */
    addressFamily?: pulumi.Input<string>;
    /**
     * A free-form text describing the Elastic IP (EIP).
     */
    description?: pulumi.Input<string>;
    /**
     * Healthcheck configuration for *managed* EIPs. It can not be added to an existing *Unmanaged* EIP.
     */
    healthcheck?: pulumi.Input<inputs.ElasticIpHealthcheck>;
    /**
     * A map of key/value labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Domain name for reverse DNS record.
     */
    reverseDns?: pulumi.Input<string>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: pulumi.Input<string>;
}
