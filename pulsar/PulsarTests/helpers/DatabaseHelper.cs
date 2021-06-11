using System;
using Npgsql;

namespace PulsarTests.helpers
{
    public class DatabaseHelper
    {
        private readonly string _uri;
        private NpgsqlConnection _conn;

        public DatabaseHelper(string uri)
        {
            this._uri = uri;
            this._conn = new NpgsqlConnection(this._uri);
            this._conn.Open();

            this.DeleteTables();
            this.SetTables();
        }

        private void DeleteTables()
        {
            try
            {
                var deleteTablesCmd = new NpgsqlCommand(@"
                    DROP TABLE videos CASCADE;
                    DROP TABLE users CASCADE;
                    DROP TABLE meta_info CASCADE;
                    DROP TYPE video_status CASCADE;
                ", this._conn);
                deleteTablesCmd.ExecuteNonQuery();
            }
            catch (Exception)
            {
                // noop
            }
        }
        
        private void SetTables()
        {
            var createTablesCmd = new NpgsqlCommand(Schema.Get(), this._conn);
            createTablesCmd.ExecuteNonQuery();
        }

        public void Reset()
        {
            var truncate = @"
                TRUNCATE TABLE videos CASCADE;
		        TRUNCATE TABLE users CASCADE;
            ";
            var defaultValues = "INSERT INTO users (username, hashpass) VALUES ('admin', '$2y$10$gXkbidnOSdZoydtOpvnsyeXpG1ZKXJ79gmn10MmsQIKpTi.hg9wfa');";

            var tx = this._conn.BeginTransaction();

            var truncateCmd = new NpgsqlCommand(truncate, this._conn, tx);
            truncateCmd.ExecuteNonQuery();
            var insertCmd = new NpgsqlCommand(defaultValues, this._conn, tx);
            insertCmd.ExecuteNonQuery();

            tx.Commit();
        }
    }
}