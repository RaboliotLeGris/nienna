namespace pulsar.clients
{
    public interface ISqlClient
    {
        public ISqlClient Connect();
        public void Execute(string cmd);
    }
}