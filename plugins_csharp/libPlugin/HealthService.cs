using Grpc.Core;

namespace libPlugin
{
    public class HealthService
    {

        public static IHealthService Get() =>
            new HealthServiceImpl();

        public static ServerServiceDefinition BindService(IHealthService health) =>
            Grpc.Health.V1.Health.BindService((Grpc.Health.V1.Health.HealthBase)health);
    }

    public interface IHealthService
    {
        void ClearAll();
        void ClearStatus(string service);
        void SetStatus(string service, HealthStatus status);
    }

    public enum HealthStatus
    {
        Unknown = 0,
        Serving = 1,
        NotServing = 2,
    }
}
