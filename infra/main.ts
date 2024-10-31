import * as pulumi from "@pulumi/pulumi";
import * as cloudflare from "@pulumi/cloudflare";
import * as std from "@pulumi/std";

const config = new pulumi.Config();

const script = new cloudflare.WorkersScript("script", {
    accountId: config.require("accountId"),
    name: "webfinger",
    content: std.file({
        input: "../dist/main.js",
    }).then(invoke => invoke.result),
    module: true,
});

const domain = new cloudflare.WorkersDomain("domain", {
    accountId: config.require("accountId"),
    hostname: config.require("zoneName"),
    service: script.name,
    zoneId: config.require("zoneId"),
});

new cloudflare.WorkersRoute("route", {
    zoneId: config.require("zoneId"),
    pattern: pulumi.interpolate`${domain.hostname}/.well-known/webfinger`,
    scriptName: script.name,
});

new cloudflare.WorkersSecret("descriptors", {
    accountId: config.require("accountId"),
    name: "descriptors",
    scriptName: script.name,
    secretText: config.require("descriptors"),
});
