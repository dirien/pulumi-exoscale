// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ComputeInstanceNetworkInterface {
    /**
     * The IPv4 address to request as static DHCP lease if the network interface is attached to a *managed* private network.
     */
    ipAddress?: string;
    /**
     * The exoscale.PrivateNetwork (ID) to attach to the instance.
     */
    networkId: string;
}

export interface DatabaseKafka {
    /**
     * Enable certificate-based authentication method.
     */
    enableCertAuth: boolean;
    /**
     * Enable Kafka Connect.
     */
    enableKafkaConnect?: boolean;
    /**
     * Enable Kafka REST.
     */
    enableKafkaRest?: boolean;
    /**
     * Enable SASL-based authentication method.
     */
    enableSaslAuth: boolean;
    /**
     * Enable Schema Registry.
     */
    enableSchemaRegistry?: boolean;
    /**
     * A list of CIDR blocks to allow incoming connections from.
     */
    ipFilters: string[];
    /**
     * Kafka Connect configuration settings in JSON format (`exo dbaas type show kafka --settings=kafka-connect` for reference).
     */
    kafkaConnectSettings: string;
    /**
     * Kafka REST configuration settings in JSON format (`exo dbaas type show kafka --settings=kafka-rest` for reference).
     */
    kafkaRestSettings: string;
    /**
     * Kafka configuration settings in JSON format (`exo dbaas type show kafka --settings=kafka` for reference).
     */
    kafkaSettings: string;
    /**
     * Schema Registry configuration settings in JSON format (`exo dbaas type show kafka --settings=schema-registry` for reference)
     */
    schemaRegistrySettings: string;
    /**
     * PostgreSQL major version (`exo dbaas type show pg` for reference; may only be set at creation time).
     */
    version: string;
}

export interface DatabaseMysql {
    /**
     * A custom administrator account password (may only be set at creation time).
     */
    adminPassword: string;
    /**
     * A custom administrator account username (may only be set at creation time).
     */
    adminUsername: string;
    /**
     * The automated backup schedule (`HH:MM`).
     */
    backupSchedule: string;
    /**
     * A list of CIDR blocks to allow incoming connections from.
     */
    ipFilters: string[];
    /**
     * MySQL configuration settings in JSON format (`exo dbaas type show mysql --settings=mysql` for reference).
     */
    mysqlSettings: string;
    /**
     * PostgreSQL major version (`exo dbaas type show pg` for reference; may only be set at creation time).
     */
    version: string;
}

export interface DatabaseOpensearch {
    dashboards?: outputs.DatabaseOpensearchDashboards;
    /**
     * Service name
     */
    forkFromService?: string;
    /**
     * Allows you to create glob style patterns and set a max number of indexes matching this pattern you want to keep. Creating indexes exceeding this value will cause the oldest one to get deleted. You could for example create a pattern looking like 'logs.?' and then create index logs.1, logs.2 etc, it will delete logs.1 once you create logs.6. Do note 'logs.?' does not apply to logs.10. Note: Setting maxIndexCount to 0 will do nothing and the pattern gets ignored.
     */
    indexPatterns?: outputs.DatabaseOpensearchIndexPattern[];
    /**
     * Template settings for all new indexes
     */
    indexTemplate?: outputs.DatabaseOpensearchIndexTemplate;
    /**
     * A list of CIDR blocks to allow incoming connections from.
     */
    ipFilters?: string[];
    /**
     * Aiven automation resets index.refresh_interval to default value for every index to be sure that indices are always visible to search. If it doesn't fit your case, you can disable this by setting up this flag to true.
     */
    keepIndexRefreshInterval?: boolean;
    /**
     * Maximum number of indexes to keep before deleting the oldest one (Minimum value is `0`)
     * * `dashboards`
     */
    maxIndexCount?: number;
    /**
     * -
     */
    recoveryBackupName?: string;
    settings?: string;
    /**
     * PostgreSQL major version (`exo dbaas type show pg` for reference; may only be set at creation time).
     */
    version: string;
}

export interface DatabaseOpensearchDashboards {
    /**
     * {Type -  schema.TypeBool, Optional -  true, Default -  true},
     */
    enabled?: boolean;
    /**
     * {Type -  schema.TypeInt, Optional -  true, Default -  128},
     */
    maxOldSpaceSize?: number;
    /**
     * {Type -  schema.TypeInt, Optional -  true, Default -  30000},
     * `settings` -  OpenSearch-specific settings, in json. e.g.`jsonencode({thread_pool_search_size: 64})`. Use `exo x get-dbaas-settings-opensearch` to get a list of available settings.
     */
    requestTimeout?: number;
}

export interface DatabaseOpensearchIndexPattern {
    /**
     * Maximum number of indexes to keep before deleting the oldest one (Minimum value is `0`)
     * * `dashboards`
     */
    maxIndexCount?: number;
    /**
     * fnmatch pattern
     */
    pattern?: string;
    /**
     * `alphabetical` or `creationDate`.
     */
    sortingAlgorithm?: string;
}

export interface DatabaseOpensearchIndexTemplate {
    /**
     * The maximum number of nested JSON objects that a single document can contain across all nested types. This limit helps to prevent out of memory errors when a document contains too many nested objects. (Default is 10000. Minimum value is `0`, maximum value is `100000`.)
     */
    mappingNestedObjectsLimit?: number;
    /**
     * The number of replicas each primary shard has. (Minimum value is `0`, maximum value is `29`)
     */
    numberOfReplicas?: number;
    /**
     * The number of primary shards that an index should have. (Minimum value is `1`, maximum value is `1024`.)
     */
    numberOfShards?: number;
}

export interface DatabasePg {
    /**
     * A custom administrator account password (may only be set at creation time).
     */
    adminPassword: string;
    /**
     * A custom administrator account username (may only be set at creation time).
     */
    adminUsername: string;
    /**
     * The automated backup schedule (`HH:MM`).
     */
    backupSchedule: string;
    /**
     * A list of CIDR blocks to allow incoming connections from.
     */
    ipFilters: string[];
    /**
     * PostgreSQL configuration settings in JSON format (`exo dbaas type show pg --settings=pg` for reference).
     */
    pgSettings: string;
    /**
     * PgBouncer configuration settings in JSON format (`exo dbaas type show pg --settings=pgbouncer` for reference).
     */
    pgbouncerSettings: string;
    /**
     * pglookout configuration settings in JSON format (`exo dbaas type show pg --settings=pglookout` for reference).
     */
    pglookoutSettings: string;
    /**
     * PostgreSQL major version (`exo dbaas type show pg` for reference; may only be set at creation time).
     */
    version: string;
}

export interface DatabaseRedis {
    /**
     * A list of CIDR blocks to allow incoming connections from.
     */
    ipFilters: string[];
    /**
     * Redis configuration settings in JSON format (`exo dbaas type show redis --settings=redis` for reference).
     */
    redisSettings: string;
}

export interface ElasticIPHealthcheck {
    /**
     * The healthcheck interval (seconds; must be between `5` and `300`; default: `10`).
     */
    interval?: number;
    /**
     * The healthcheck mode (`tcp`, `http` or `https`; may only be set at creation time).
     */
    mode: string;
    /**
     * The healthcheck target port (must be between `1` and `65535`).
     */
    port: number;
    /**
     * The number of failed healthcheck attempts before considering the target unhealthy (must be between `1` and `20`; default: `2`).
     */
    strikesFail?: number;
    /**
     * The number of successful healthcheck attempts before considering the target healthy (must be between `1` and `20`; default: `3`).
     */
    strikesOk?: number;
    /**
     * The time before considering a healthcheck probing failed (seconds; must be between `2` and `60`; default: `3`).
     */
    timeout?: number;
    /**
     * Disable TLS certificate verification for healthcheck in `https` mode (boolean; default: `false`).
     */
    tlsSkipVerify?: boolean;
    /**
     * The healthcheck server name to present with SNI in `https` mode.
     */
    tlsSni?: string;
    /**
     * The healthcheck target URI (required in `http(s)` modes).
     */
    uri?: string;
}

export interface GetComputeInstanceListInstance {
    antiAffinityGroupIds?: string[];
    createdAt: string;
    deployTargetId: string;
    diskSize: number;
    elasticIpIds: string[];
    id?: string;
    ipv6: boolean;
    ipv6Address: string;
    labels?: {[key: string]: string};
    managerId: string;
    managerType: string;
    name?: string;
    privateNetworkIds: string[];
    publicIpAddress: string;
    reverseDns: string;
    securityGroupIds: string[];
    sshKey: string;
    state: string;
    templateId: string;
    type: string;
    userData: string;
    /**
     * The Exoscale [Zone][zone] name.
     */
    zone: string;
}

export interface GetDomainRecordFilter {
    /**
     * A regular expression to match the record content.
     */
    contentRegex?: string;
    /**
     * The record ID to match.
     */
    id?: string;
    /**
     * The domain record name to match.
     */
    name?: string;
    /**
     * The record type to match.
     */
    recordType?: string;
}

export interface GetDomainRecordRecord {
    /**
     * The domain record content.
     */
    content?: string;
    /**
     * The exoscale.Domain name to match.
     */
    domain?: string;
    /**
     * The record ID to match.
     */
    id?: string;
    /**
     * The domain record name to match.
     */
    name?: string;
    /**
     * The record priority.
     */
    prio?: number;
    /**
     * The record type to match.
     */
    recordType?: string;
    /**
     * The record TTL.
     */
    ttl?: number;
}

export interface GetElasticIPHealthcheck {
    /**
     * The healthcheck interval in seconds.
     */
    interval: number;
    /**
     * The healthcheck mode.
     */
    mode: string;
    /**
     * The healthcheck target port.
     */
    port: number;
    /**
     * The number of failed healthcheck attempts before considering the target unhealthy.
     */
    strikesFail: number;
    /**
     * The number of successful healthcheck attempts before considering the target healthy.
     */
    strikesOk: number;
    /**
     * The time in seconds before considering a healthcheck probing failed.
     */
    timeout: number;
    /**
     * Disable TLS certificate verification for healthcheck in `https` mode.
     */
    tlsSkipVerify: boolean;
    /**
     * The healthcheck server name to present with SNI in `https` mode.
     */
    tlsSni: string;
    /**
     * The healthcheck URI.
     */
    uri: string;
}

export interface GetInstancePoolInstance {
    /**
     * The instance pool ID to match (conflicts with `name`).
     */
    id?: string;
    /**
     * The instance (main network interface) IPv6 address.
     */
    ipv6Address: string;
    /**
     * The pool name to match (conflicts with `id`).
     */
    name?: string;
    /**
     * The instance (main network interface) IPv4 address.
     */
    publicIpAddress: string;
}

export interface GetInstancePoolListPool {
    affinityGroupIds: string[];
    deployTargetId: string;
    description: string;
    diskSize: number;
    elasticIpIds: string[];
    id?: string;
    instancePrefix: string;
    instanceType: string;
    instances: outputs.GetInstancePoolListPoolInstance[];
    ipv6: boolean;
    keyPair: string;
    labels?: {[key: string]: string};
    name?: string;
    networkIds: string[];
    securityGroupIds: string[];
    size: number;
    state: string;
    templateId: string;
    userData: string;
    /**
     * The Exoscale [Zone][zone] name.
     */
    zone: string;
}

export interface GetInstancePoolListPoolInstance {
    id?: string;
    ipv6Address: string;
    name?: string;
    publicIpAddress: string;
}

export interface InstancePoolInstance {
    /**
     * The compute instance ID.
     */
    id?: string;
    /**
     * The instance (main network interface) IPv6 address.
     */
    ipv6Address: string;
    /**
     * The instance pool name.
     */
    name?: string;
    /**
     * The instance (main network interface) IPv4 address.
     */
    publicIpAddress: string;
}

export interface NLBServiceHealthcheck {
    /**
     * The healthcheck interval in seconds (default: `10`).
     */
    interval?: number;
    /**
     * The healthcheck mode (`tcp`|`http`|`https`; default: `tcp`).
     */
    mode?: string;
    /**
     * The healthcheck port.
     */
    port: number;
    /**
     * The healthcheck retries (default: `1`).
     */
    retries?: number;
    /**
     * The healthcheck timeout (seconds; default: `5`).
     */
    timeout?: number;
    /**
     * The healthcheck TLS SNI server name (only if `mode` is `https`).
     */
    tlsSni?: string;
    /**
     * The healthcheck URI (must be set only if `mode` is `http(s)`).
     */
    uri?: string;
}

export interface SKSClusterOidc {
    /**
     * The OpenID client ID.
     */
    clientId: string;
    /**
     * An OpenID JWT claim to use as the user's group.
     */
    groupsClaim?: string;
    /**
     * An OpenID prefix prepended to group claims.
     */
    groupsPrefix?: string;
    /**
     * The OpenID provider URL.
     */
    issuerUrl: string;
    /**
     * A map of key/value pairs that describes a required claim in the OpenID Token.
     */
    requiredClaim?: {[key: string]: string};
    /**
     * An OpenID JWT claim to use as the user name.
     */
    usernameClaim?: string;
    /**
     * An OpenID prefix prepended to username claims.
     */
    usernamePrefix?: string;
}

export interface SecurityGroupRulesEgress {
    /**
     * A list of (`INGRESS`) source / (`EGRESS`) destination IP subnet (in CIDR notation) to match.
     */
    cidrLists?: string[];
    /**
     * A free-form text describing the block.
     * * `icmpType`/`icmpCode` - An ICMP/ICMPv6 type/code to match.
     */
    description?: string;
    icmpCode?: number;
    icmpType?: number;
    ids: string[];
    /**
     * A list of ports or port ranges (`<start_port>-<end_port>`).
     */
    ports?: string[];
    /**
     * The network protocol to match (`TCP`, `UDP`, `ICMP`, `ICMPv6`, `AH`, `ESP`, `GRE`, `IPIP` or `ALL`).
     */
    protocol?: string;
    /**
     * A list of source (for ingress)/destination (for egress) identified by a security group.
     */
    userSecurityGroupLists?: string[];
}

export interface SecurityGroupRulesIngress {
    /**
     * A list of (`INGRESS`) source / (`EGRESS`) destination IP subnet (in CIDR notation) to match.
     */
    cidrLists?: string[];
    /**
     * A free-form text describing the block.
     * * `icmpType`/`icmpCode` - An ICMP/ICMPv6 type/code to match.
     */
    description?: string;
    icmpCode?: number;
    icmpType?: number;
    ids: string[];
    /**
     * A list of ports or port ranges (`<start_port>-<end_port>`).
     */
    ports?: string[];
    /**
     * The network protocol to match (`TCP`, `UDP`, `ICMP`, `ICMPv6`, `AH`, `ESP`, `GRE`, `IPIP` or `ALL`).
     */
    protocol?: string;
    /**
     * A list of source (for ingress)/destination (for egress) identified by a security group.
     */
    userSecurityGroupLists?: string[];
}

