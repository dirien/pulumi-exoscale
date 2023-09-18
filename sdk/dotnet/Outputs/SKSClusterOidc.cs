// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Exoscale.Outputs
{

    [OutputType]
    public sealed class SKSClusterOidc
    {
        /// <summary>
        /// The OpenID client ID.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// An OpenID JWT claim to use as the user's group.
        /// </summary>
        public readonly string? GroupsClaim;
        /// <summary>
        /// An OpenID prefix prepended to group claims.
        /// </summary>
        public readonly string? GroupsPrefix;
        /// <summary>
        /// The OpenID provider URL.
        /// </summary>
        public readonly string IssuerUrl;
        /// <summary>
        /// A map of key/value pairs that describes a required claim in the OpenID Token.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? RequiredClaim;
        /// <summary>
        /// An OpenID JWT claim to use as the user name.
        /// </summary>
        public readonly string? UsernameClaim;
        /// <summary>
        /// An OpenID prefix prepended to username claims.
        /// </summary>
        public readonly string? UsernamePrefix;

        [OutputConstructor]
        private SKSClusterOidc(
            string clientId,

            string? groupsClaim,

            string? groupsPrefix,

            string issuerUrl,

            ImmutableDictionary<string, string>? requiredClaim,

            string? usernameClaim,

            string? usernamePrefix)
        {
            ClientId = clientId;
            GroupsClaim = groupsClaim;
            GroupsPrefix = groupsPrefix;
            IssuerUrl = issuerUrl;
            RequiredClaim = requiredClaim;
            UsernameClaim = usernameClaim;
            UsernamePrefix = usernamePrefix;
        }
    }
}
