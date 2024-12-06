# Who am I?

You know when whenever you need to check your public IP address for whatever reason? You know how you end-up on a crappy website with a crappy UI?

Well, this is a simple web service that returns your public IP address. And just your public IP address. Nothing more. Not one more byte.

It's about 20 lines of Go code, deployable on your Kubernetes with Helm. Docker, Podman, or even bare-metal will do just fine as well.

## How to use

You can test it on my public instance: https://ip.smyler.net (but there isn't much to test, it just returns your IP public address...).

It also works great with `curl`:

```sh
curl https://ip.smyler.net
```

And that's about it.

## What about IPv6?

IPv6 is supported by the application itself, but not by the Kubernetes node that runs my instance. So if you want it, you'll need to self-host (and you are lucky, that's the next section).

## What if I wan to self-host?

You can install the application fairly easily using Helm. My instance is deployed as follow:

```sh
# Run from the 'helm' directory
helm upgrade --install whoami . \
    --set ingress.enabled=true \
    --set 'ingress.hosts[0]=ip.smyler.net' \
    --set 'ingress.hosts[1]=ipv4.smyler.net' \
    --set 'ingress.hosts[2]=ipv6.smyler.net' \
    --set 'ingress.hosts[3]=whoami.smyler.net'
```

The different hosts in the ingress are meant to allow for more control over the IP version in use dependending on the host that gets accessed. For example, by only setting an `AAAA` record to resolve `ipv6.example.com`, accessing it will always return your public IPv6 address (or will fail if you do not have IPv6).

If your deployment is behind a proxy, it will need to set the `X-Forwarded-For` header or the server will show the proxy IP address. The end-application blindly trusts all proxies.
