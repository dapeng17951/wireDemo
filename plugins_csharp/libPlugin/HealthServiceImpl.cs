using System.Collections.Generic;
using System.Threading.Tasks;

using Grpc.Core;
using Grpc.Health.V1;


namespace libPlugin
{
    /// </summary>
    public class HealthServiceImpl : Grpc.Health.V1.Health.HealthBase, IHealthService
    {
        private readonly object myLock = new object();
        private readonly Dictionary<string, HealthCheckResponse.Types.ServingStatus> statusMap =
            new Dictionary<string, HealthCheckResponse.Types.ServingStatus>();

        // Implement the service interface, mapping to concrete implementation
        void IHealthService.ClearAll() => ClearAll();
        void IHealthService.ClearStatus(string service) => ClearStatus(service);
        void IHealthService.SetStatus(string service, HealthStatus status) =>
            SetStatus(service, (HealthCheckResponse.Types.ServingStatus)status);

        /// <summary>
        /// Sets the health status for given service.
        /// </summary>
        /// <param name="service">The service. Cannot be null.</param>
        /// <param name="status">the health status</param>
        public void SetStatus(string service, HealthCheckResponse.Types.ServingStatus status)
        {
            lock (myLock)
            {
                statusMap[service] = status;
            }
        }

        /// <summary>
        /// Clears health status for given service.
        /// </summary>
        /// <param name="service">The service. Cannot be null.</param>
        public void ClearStatus(string service)
        {
            lock (myLock)
            {
                statusMap.Remove(service);
            }
        }

        /// <summary>
        /// Clears statuses for all services.
        /// </summary>
        public void ClearAll()
        {
            lock (myLock)
            {
                statusMap.Clear();
            }
        }

        /// <summary>
        /// Performs a health status check.
        /// </summary>
        /// <param name="request">The check request.</param>
        /// <param name="context">The call context.</param>
        /// <returns>The asynchronous response.</returns>
        public override Task<HealthCheckResponse> Check(HealthCheckRequest request, ServerCallContext context)
        {
            lock (myLock)
            {
                var service = request.Service;

                HealthCheckResponse.Types.ServingStatus status;
                if (!statusMap.TryGetValue(service, out status))
                {
                    // TODO(jtattermusch): returning specific status from server handler is not supported yet.
                    throw new RpcException(new Status(StatusCode.NotFound, ""));
                }
                return Task.FromResult(new HealthCheckResponse { Status = status });
            }
        }
    }
}
