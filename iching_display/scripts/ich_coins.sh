#!/bin/bash

# 3-coins method of divination

(( t1 = 6 + (RANDOM % 2) + (RANDOM % 2) + (RANDOM % 2) ))
(( t2 = 6 + (RANDOM % 2) + (RANDOM % 2) + (RANDOM % 2) ))
(( t3 = 6 + (RANDOM % 2) + (RANDOM % 2) + (RANDOM % 2) ))
(( t4 = 6 + (RANDOM % 2) + (RANDOM % 2) + (RANDOM % 2) ))
(( t5 = 6 + (RANDOM % 2) + (RANDOM % 2) + (RANDOM % 2) ))
(( t6 = 6 + (RANDOM % 2) + (RANDOM % 2) + (RANDOM % 2) ))


iching_display $t1$t2$t3$t4$t5$t6

