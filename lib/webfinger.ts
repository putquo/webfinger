export interface Identity {
  email: string;
  provider: string;
}

export interface Resource {
  subject: string;
  links: {
    rel: string;
    href: string;
  }[];
}

export type Descriptor = Record<string, Resource>;

export const generateJsonResourceDescriptor = (identities: Identity[]): Descriptor => {
  return identities.reduce((descriptor, identity) => ({
    ...descriptor,
    [identity.email]: {
      subject: `acct:${identity.email}`,
      links: [
        {
          rel: "http://openid.net/specs/connect/1.0/issuer",
          href: identity.provider.toString(),
        },
      ],
    } as Resource
  }), {} as Descriptor)
}

