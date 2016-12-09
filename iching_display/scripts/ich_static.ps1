# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
# Script to generate a static I-Ching casting (i.e., no moving lines).  
# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

$t1 = 7 + (Get-Random -Maximum 2)
$t2 = 7 + (Get-Random -Maximum 2)
$t3 = 7 + (Get-Random -Maximum 2)
$t4 = 7 + (Get-Random -Maximum 2)
$t5 = 7 + (Get-Random -Maximum 2)
$t6 = 7 + (Get-Random -Maximum 2)

& iching_display $t1$t2$t3$t4$t5$t6

