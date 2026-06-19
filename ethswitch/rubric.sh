#!/bin/sh

testOutput=$1

total=0
partial=0

passTest() {
    inc=$1
    testName=$2
    
    if grep -qF -- "--- PASS: $testName " "$testOutput"
    then
        total=$((total + inc))
        echo "$testName: passed ==> $inc of $inc point(s)"
    else
        if grep -qF -- "=== RUN   $testName" "$testOutput"
        then
            partial=$((partial + inc))
        fi
        echo "$testName: failed ==> 0 of $inc point(s)"
    fi
}

echo "*** Grading rubric ***"
passTest 1 TestBasicRequirements
passTest 1 TestReadWriteFrame
passTest 1 TestReadWriteFrame_Invalid
passTest 1 TestSwitch_Simple/discovery
passTest 1 TestSwitch_Simple/discovery/unicast
passTest 1 TestSwitch_Simple/discovery/unicast/regular
passTest 1 TestSwitch_Simple/discovery/unicast/drop_packets
passTest 1 TestSwitch_Simple/broadcast_frame
passTest 1 TestSwitch_Simple/bad_frame
passTest 1 TestSwitch_Simple

partialPoints=$(( partial / 2 ))
echo "Partial credit: $partialPoints points"
echo "Total: $(( total + partialPoints )) points (10 points possible)"
echo "**********************"
