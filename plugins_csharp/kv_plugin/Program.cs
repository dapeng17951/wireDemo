using Grpc.Core;
using libPlugin;
using System.Net;

namespace kv_plugin
{
    internal class Program
    {
        async static Task Main(string[] args)
        {
            // go-plugin semantics depend on the Health Check service from gRPC
            var health = HealthService.Get();
            health.SetStatus("plugin", HealthStatus.Serving);

            // Build a server to host the plugin over gRPC
            var server = new Server
            {
                Ports = { { Plugin.ServiceHost, Plugin.ServicePort, ServerCredentials.Insecure } },
                Services = {
                    { HealthService.BindService(health) },
                    { KV.KV.BindService(new Plugin()) },
                },
            };

            server.Start();

            // Part of the go-plugin handshake:
            //  https://github.com/hashicorp/go-plugin/blob/master/docs/guide-plugin-write-non-go.md#4-output-handshake-information
            await Console.Out.WriteAsync($"1|1|tcp|{Plugin.ServiceHost}:{Plugin.ServicePort}|grpc\n");
            await Console.Out.FlushAsync();

            while (Console.Read() == -1)
                await Task.Delay(1000);

            await server.ShutdownAsync();
        }
    }
}
