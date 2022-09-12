use strict;
use warnings;

use FindBin;
use lib $FindBin::Bin;

require Generator;

SetCounter(1);
WriteDefine();

my $f = FormatString();

my $a = -1;
my $b = -1;
my $o = -1;

$a = LO();
$b = LO();
$o = Nand($a, $b);
printf("//Nand($f, $f) -> $f \n", $a, $b, $o);
Write();

$a = HI();
$b = LO();
$o = Nand($a, $b);
printf("//Nand($f, $f) -> $f \n", $a, $b, $o);
Write();

$a = LO();
$b = HI();
$o = Nand($a, $b);
printf("//Nand($f, $f) -> $f \n", $a, $b, $o);
Write();

$a = HI();
$b = HI();
$o = Nand($a, $b);
printf("//Nand($f, $f) -> $f \n", $a, $b, $o);
Write();

