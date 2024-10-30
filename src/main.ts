import { generateJsonResourceDescriptor, Identity } from "../lib/webfinger";

interface Context {
  IDENTITIES: string;
}

export default {
  async fetch(request: Request, env: Context): Promise<Response> {
		const url = new URL(request.url);
		if (url.pathname !== "/.well-known/webfinger") {
			return new Response("Not found", {
				status: 404
			});
		}

		const withPrefix = url.searchParams.get("resource");
		if (withPrefix === null) {
			return new Response("Bad request", {
				status: 400
			});
		}

		const resourceParts = withPrefix.split(":");
		if (resourceParts.length !== 2 || resourceParts[0] !== "acct") {
			return new Response("Bad request", {
				status: 400
			});
		}

		const identites = JSON.parse(env.IDENTITIES);
		const descriptor = generateJsonResourceDescriptor(identites);

		const resourceIdentifier = resourceParts[1];
		const resource = descriptor[resourceIdentifier];
    if (resource === undefined) {
      return new Response("Not found", {
        status: 404,
      });
    }

    return Response.json(resource);
  },
}
