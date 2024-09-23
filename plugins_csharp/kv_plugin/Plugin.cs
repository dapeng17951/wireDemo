using Google.Protobuf;
using Grpc.Core;
using KV;
using libPlugin;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace kv_plugin
{
    internal class Plugin : KV.KV.KVBase
    {
        public const string ServiceHost = "localhost";
        public const int ServicePort = 1234;
        public const int AppProtoVersion = 1;

        public override async Task<Empty> Put(PutRequest request, ServerCallContext context)
        {
            var filename = $"kv_{request.Key}";
            await File.WriteAllTextAsync(filename,
                $"{request.Value.ToStringUtf8()}\n\nWritten from plugin-dotnet\n");

            return new Empty();
        }

        public override async Task<GetResponse> Get(GetRequest request, ServerCallContext context)
        {
            var filename = $"kv_{request.Key}";
            //return new GetResponse
            //{
            //    Value = ByteString.CopyFromUtf8(await File.ReadAllTextAsync(filename)),
            //};
            return new GetResponse
            {
                Value = ByteString.CopyFromUtf8(filename),
            };
        }        
    }
}
