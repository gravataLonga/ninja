#!/usr/bin/env sh

echo "======================="
echo "\n>>> SCRIPTING FIB(5);"

echo "### PHP"
time php ./fib.php 5
echo ""


echo "### GO"
time ./fib-go 5
echo ""


echo "### NINJA"
time ninja-lang -e "function fib(n) { if (n < 2) { return n; } return fib(n-1) + fib(n-2); };fib(5)"

echo "======================="
echo "\n>>> SCRIPTING FIB(10);"

echo "### PHP"
time php ./fib.php 10
echo ""

echo "### GO"
time ./fib-go 10
echo ""

echo "### NINJA"
time ninja-lang -e "function fib(n) { if (n < 2) { return n; } return fib(n-1) + fib(n-2); };fib(10)"

echo "======================="
echo "\n>>> SCRIPTING FIB(20);"

echo "### PHP"
time php ./fib.php 20
echo ""


echo "### GO"
time ./fib-go 20
echo ""

echo "### NINJA"
time ninja-lang -e "function fib(n) { if (n < 2) { return n; } return fib(n-1) + fib(n-2); };fib(20)"

echo "======================="
echo "\n>>> SCRIPTING FIB(25);"

echo "### PHP"
time php ./fib.php 25
echo ""


echo "### GO"
time ./fib-go 25
echo ""

echo "### NINJA"
time ninja-lang -e "function fib(n) { if (n < 2) { return n; } return fib(n-1) + fib(n-2); };fib(25)"

echo "======================="
echo "\n>>> SCRIPTING FIB(30);"

echo "### PHP"
time php ./fib.php 30
echo ""


echo "### GO"
time ./fib-go 30
echo ""

echo "### NINJA"
time ninja-lang -e "function fib(n) { if (n < 2) { return n; } return fib(n-1) + fib(n-2); };fib(30)"