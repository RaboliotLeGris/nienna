#!/usr/bin/env perl

use strict;
use warnings;

use lib qw(..);
use JSON qw();
use Digest::SHA qw(sha256_hex);

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

sub get_sha256 {
    my ($data) = @_;
    my $sha = sha256_hex($data);

    return $sha;
}

my $data = read_file("misc/schema_users.json");

# Parsing json
my $json = JSON->new;
my $parsed = $json->decode($data);

my $root_sql = read_file("schema.sql");
my $root_sha = get_sha256($root_sql);
print "root sha: $root_sha\n";

my $is_failing= 0;

foreach (@{$parsed})
{
    my $subfile = read_file("../$_");

    my $got_sha = get_sha256($subfile);
    if ($root_sha ne $got_sha) {
        print "[FAIL] - \"$_\" does not match root SQL schema (Got: $got_sha)\n";
        $is_failing = 1;
    }
}

if ($is_failing eq 0) {
    print "\n\nResult: SUCCESS \n"
} else {
    print "\n\nResult: FAILURE \n"
}

exit $is_failing;