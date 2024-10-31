interface Context {
  descriptors: string;
}

interface Descriptor {
  subject: string;
  links: {
    rel: string;
    href: string;
  }[];
}

export default {
  async fetch(request: Request, ctx: Context): Promise<Response> {
		const url = new URL(request.url);
		if (url.pathname !== "/.well-known/webfinger") {
			return new Response("Not found", { status: 404 });
		}

		const resource = url.searchParams.get("resource");
		if (!resource) {
			return new Response("Bad request", { status: 400 });
		}

		const descriptors: Descriptor[] = JSON.parse(ctx.descriptors);
		const descriptor = descriptors.find(descriptor => descriptor.subject === resource);
		return descriptor ? Response.json(descriptor) : new Response("Not found", { status: 404 });
  },
}
