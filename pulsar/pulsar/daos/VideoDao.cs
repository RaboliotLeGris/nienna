using System;
using pulsar.clients;

namespace pulsar.daos
{
    public class VideoDao
    {
        private ISqlClient _sqlClient;

        public VideoDao(ISqlClient sqlClient)
        {
            this._sqlClient = sqlClient;
        }

        public void UpdateStatus(string slug, string status)
        {
            this._sqlClient.Execute($"UPDATE videos SET status='{status}' WHERE slug='{slug}';");
        }
    }
}