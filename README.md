# WebFinger Cloudflare Worker

This project is a simple WebFinger implementation hosted on a Cloudflare Worker and deployed via Pulumi.

## Overview

This application handles requests to the `/.well-known/webfinger` endpoint, responding with appropriate resource descriptors, if available.
It follows the specification defined by [RFC 7033](https://datatracker.ietf.org/doc/html/rfc7033).

### Why?

This is useful for if you want to configure a [custom OIDC provider for Tailscale](https://tailscale.com/kb/1240/sso-custom-oidc), since you can run this
free of charge on Cloudflare.

## Prerequisites

- Node.js and npm
- Wrangler
- Pulumi CLI
- Cloudflare Account and API keys (configured in Pulumi)

#### Setting Pulumi Secrets

Set the `CLOUDFLARE_API_TOKEN` from the [Dashboard](https://dash.cloudflare.com/profile/api-tokens) with at least `Workers Routes:Edit`, `Workers Scripts:Edit`, and `Account Settings:Edit` privileges.

The application requires specific Pulumi secrets for deployment to Cloudflare. Ensure these are set before running the deployment:

- `accountId`: Your Cloudflare account ID.
- `zoneName`: The domain name associated with your Cloudflare account.
- `zoneId`: The zone ID for the domain you are deploying to.
- `descriptors`: The JSON-encoded array of WebFinger descriptors (allowing for more than one subject) as defined by 
[RFC 7033](https://datatracker.ietf.org/doc/html/rfc7033#section-3.1). Replace the `$.subject` and `$.links[0].href` as required.

You can set these secrets using the Pulumi CLI:

```sh
  pulumi config set accountId <your_account_id>
  pulumi config set zoneName <your_zone_name>
  pulumi config set zoneId <your_zone_id>
  pulumi config set descriptors '<your_descriptors_json>'
```

## Deployment

1. **Build**: Compile TypeScript to JavaScript using Wrangler.
```sh
npm run build
```
2. **Deploy**: Use Pulumi to deploy the Cloudflare worker and configure routing.
```sh
npm run deploy
```

