use strict;
use warnings;

use lib qw(..);
use JSON qw();
use File::Copy;


sub read_file {
    my ($filepath) = @_;
    my $data = do {
        open(my $json_fh, "<:encoding(UTF-8)", $filepath)
            or die("Can't open $filepath\": $!\n");
        local $/;
        <$json_fh>
    };
    return $data;
}

my $data = read_file("src/misc/schema_users.json");

# Parsing json
my $json = JSON->new;
my $parsed = $json->decode($data);

foreach (@{$parsed})
{
    print "[COPY] src/misc/schema_users.json to $_ \n";
    copy("src/schema.sql","../$_") or die "Copy failed: $!";
}
