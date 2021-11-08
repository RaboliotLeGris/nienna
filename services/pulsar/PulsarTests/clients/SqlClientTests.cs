using System;
using NUnit.Framework;
using pulsar.clients;
using PulsarTests.helpers;

namespace PulsarTests.clients
{
    public class SqlClientTests
    {
        private string _sqlUri;

        [OneTimeSetUp]
        public void SetSqlUri()
        {
            this._sqlUri = Environment.GetEnvironmentVariable("TEST_IN_DOCKER") != null
                ? "Host=pg;Username=nienna;Password=nienna;Database=nienna"
                : "Host=localhost;Username=nienna;Password=nienna;Database=nienna";
            // Setup database
            DatabaseHelper databaseHelper = new DatabaseHelper(this._sqlUri);
        }
        
        [Test]
        public void ConnectToTheDatabaseSucceed()
        {
            SqlClient sqlClient = new SqlClient(this._sqlUri);
            sqlClient.Connect();
            Assert.Pass();
        }
    }
}