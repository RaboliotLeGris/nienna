using System;
using Npgsql;

namespace pulsar.clients
{
    public class SqlClient : ISqlClient
    {
        private readonly string _uri;
        private NpgsqlConnection _conn;

        public SqlClient(string uri)
        {
            this._uri = uri;
        }

        ~SqlClient()
        {
            this._conn.Close();
        }

        public ISqlClient Connect()
        {
            this._conn = new NpgsqlConnection(this._uri);
            this._conn.Open();
            return this;
        }
        
        // FIXME(RaboliotLeGris): Not a great way to do it, if the complexity of the SqlClient class increase, we might want to refactor this
        public void Execute(string cmd)
        {
            var command = new NpgsqlCommand(cmd, this._conn);
            var affectedRows = command.ExecuteNonQuery();
            Console.WriteLine("AFFECTED ROWS " + affectedRows);
        }
        
        
    }
}