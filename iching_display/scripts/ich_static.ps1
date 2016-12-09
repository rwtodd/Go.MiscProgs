# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
# Script to generate a static I-Ching casting (i.e., no moving lines).  
# Suitable for piping into iching_display.exe
# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

$t1 = 7 + (Get-Random -Maximum 2)
$t2 = 7 + (Get-Random -Maximum 2)
$t3 = 7 + (Get-Random -Maximum 2)
$t4 = 7 + (Get-Random -Maximum 2)
$t5 = 7 + (Get-Random -Maximum 2)
$t6 = 7 + (Get-Random -Maximum 2)

Write-Output $t1$t2$t3$t4$t5$t6

