using System;
using Npgsql;
using NUnit.Framework;
using pulsar.clients;
using pulsar.daos;
using PulsarTests.helpers;

namespace PulsarTests.daos
{
    public class VideoDaoTests
    {
        private string _uri;
        private NpgsqlConnection _conn;
        private DatabaseHelper dbHelper;
        private ISqlClient sqlClient;

        [OneTimeSetUp]
        public void SetupDb()
        {
            this._uri = Environment.GetEnvironmentVariable("TEST_IN_DOCKER") != null
                ? "Host=db;Username=nienna;Password=nienna;Database=nienna"
                : "Host=localhost;Username=nienna;Password=nienna;Database=nienna";
            this._conn = new NpgsqlConnection(this._uri);
            this._conn.Open();
            this.sqlClient = new SqlClient(this._uri).Connect();
            this.dbHelper = new DatabaseHelper(this._uri);
        }

        [SetUp]
        public void ResetDb()
        {
            this.dbHelper.Reset();
        }

        [Test]
        public void UpdateVideoStatus()
        {
            var insertOneVideoCmd =
                new NpgsqlCommand(
                    "INSERT INTO videos (slug, uploader, title, description, status) VALUES ('SomeSlug', 1, 'Un titre', 'description', 'PROCESSING');",
                    this._conn);
            insertOneVideoCmd.ExecuteNonQuery();

            VideoDao videoDao = new VideoDao(new SqlClient(this._uri).Connect());
            videoDao.UpdateStatus("SomeSlug", "READY");

            var selectOneCmd = new NpgsqlCommand("SELECT status::TEXT from videos where slug='SomeSlug'", this._conn);
            var res = selectOneCmd.ExecuteReader();

            if (!res.HasRows)
            {
                Assert.Fail();
            }

            res.Read();
            Assert.AreEqual(res[0], "READY");
        }

        [Test]
        public void UpdateVideoWithBadStatusFail()
        {
            var insertOneVideoCmd =
                new NpgsqlCommand(
                    "INSERT INTO videos (slug, uploader, title, description, status) VALUES ('SomeSlug', 1, 'Un titre', 'description', 'PROCESSING');",
                    this._conn);
            insertOneVideoCmd.ExecuteNonQuery();

            VideoDao videoDao = new VideoDao(new SqlClient(this._uri).Connect());
            var ex = Assert.Throws<PostgresException>(delegate { videoDao.UpdateStatus("SomeSlug", "RANDOM_STATUS"); });
            Assert.AreEqual(ex.Message, "22P02: invalid input value for enum video_status: \"RANDOM_STATUS\"");
        }
    }
}