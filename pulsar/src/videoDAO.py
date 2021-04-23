class VideoDAO:
    conn = None

    def __init__(self, conn):
        self.conn = conn

    def update_status(self, slug, status):
        if self.conn is not None:
            cs = self.conn.cursor()
            cs.execute("UPDATE videos SET status = %s WHERE slug = %s;", (status, slug))
