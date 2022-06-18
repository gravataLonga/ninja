#!/usr/bin/env sh

echo "======================="
echo "\n>>> SCRIPTING FIB(5);"

echo "### PHP"
time php ./tests/fib.php 5
echo ""


echo "### GO"
time ./tests/fib-go 5
echo ""


echo "### NINJA"
time ./ninja-osx -e "function fib(n) { if (n < 2) { return n; } return fib(n-1) + fib(n-2); };fib(5)"

echo "======================="
echo "\n>>> SCRIPTING FIB(10);"

echo "### PHP"
time php ./tests/fib.php 10
echo ""

echo "### GO"
time ./tests/fib-go 10
echo ""

echo "### NINJA"
time ./ninja-osx -e "function fib(n) { if (n < 2) { return n; } return fib(n-1) + fib(n-2); };fib(10)"

echo "======================="
echo "\n>>> SCRIPTING FIB(20);"

echo "### PHP"
time php ./tests/fib.php 20
echo ""


echo "### GO"
time ./tests/fib-go 20
echo ""

echo "### NINJA"
time ./ninja-osx -e "function fib(n) { if (n < 2) { return n; } return fib(n-1) + fib(n-2); };fib(20)"