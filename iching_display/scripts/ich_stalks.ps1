# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
# Script to generate an I-Ching casting via the yarrow stalk method. 
# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
# P R O B A B I L I T I E S
#
# 6 - moving yin: 1 in 16 
# 7 - static yang: 5 in 16
# 8 - static yin: 7 in 16
# 9 - moving yang: 3 in 16
# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
function cast {
  switch (Get-Random -Maximum 16) {
    0                                     { 6 }
    { @(1,2,3,4,5) -contains $_ }         { 7 }
    { @(6,7,8,9,10,11,12) -contains $_ }  { 8 }
    { @(13,14,15) -contains $_ }          { 9 }
  } 
}

$t1 = cast
$t2 = cast
$t3 = cast
$t4 = cast
$t5 = cast
$t6 = cast

& iching_display $t1$t2$t3$t4$t5$t6

