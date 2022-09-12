use strict;
use warnings;

use FindBin;
use lib $FindBin::Bin;

require Generator;

SetCounter(1);
WriteDefine();

my $f = FormatString();

my $d0    = -1;
my $d1    = -1;
my $d2    = -1;
my $clock = -1;
my $load  = -1;
my $clear = -1;
my $q0    = -1;
my $q1    = -1;
my $q2    = -1;

$d0    = HI();
$d1    = HI();
$d2    = HI();
$clock = CLOCK();
$load  = HI();
$clear = LO();
($q0, $q1, $q2) = Register3($d0, $d1, $d2, $clock, $load, $clear);
printf("//Register3($f, $f, $f, $f, $f, $f) -> $f, $f, $f, \n", $d0, $d1, $d2, $clock, $load, $clear, $q0, $q1, $q2);
Write();

