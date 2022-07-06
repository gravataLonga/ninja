<p align="center">

![ninja programming language](./ninja.svg)

</p>

[![Testing](https://github.com/gravataLonga/ninja/actions/workflows/main.yml/badge.svg)](https://github.com/gravataLonga/ninja/actions/workflows/main.yml)  

# Install  

## Homebrew  

```sh
brew tap gravatalonga/ninja-lang
brew install ninja-lang
```  

## YUM / RPM  

To enable, add the following file `/etc/yum.repos.d/ninja.repo`:

```sh
[ninja]
name=Ninja Programming Language
baseurl=https://yum.fury.io/gravatalonga/
enabled=1
gpgcheck=0
```  

Check if correctly created  

```
yum --disablerepo=* --enablerepo=ninja list available
```  

To install you only need run following command:  

```
yum install ninja-lang  
```  

## APT    

To configure apt access, create a following file `/etc/apt/sources.list.d/ninja.list` with content of :  

```  
deb [trusted=yes] https://apt.fury.io/gravatalonga/ /
```  

Or use this one line command:  

```
echo "deb [trusted=yes] https://apt.fury.io/gravatalonga/ /" > /etc/apt/sources.list.d/ninja.list
```  

and them you can install  

```
sudo apt install ninja-lang
```

## Manual Download    

Download from [github](https://github.com/gravataLonga/ninja/releases)  

## Manual Installation  

```sh  
git clone https://github.com/gravataLonga/ninja
cd ninja
go build -o ninja-lang
```  

# Documentation  

For more detail about language, you can check [here](https://ninja.jonathan.pt) (Still working in progress).  

# Demo  

Resolving katas you can check this repository  
https://adventofcode.com/2015

# Syntax  

## Variable  

`var <identifier> = <expression>;`  

Examples  

```
var a = 1;
var a1 = "Name";
var b = 2.0;
var c = a + 1;
var d = a + b;
var e = ++a;
var f = function () {};
var g = [1, 2, 3, "hello", function() {}];
var h = {"me":"Jonathan Fontes","age":1,"likes":["php","golang","ninja"]}
var i = a < b;
var j = true;
var k = !j;
var l = if (a) { "yes" } else { "no" };  

g[0] = 10;
g[4] = "new value"; // it will append to array.  
h["other"] = true;
"ola"[0] // print o  
```  

It's possible to reassign variable for example:  

```
var a = 1;
a = a + 1;  
puts(a);  
```

## Data Types Availables  

```
 /**
  * Booleans
  */
 
 true;
 false;

 /**
  * Integer
  */
 
 1;
 20000;
 
 /**
  * Floats
  */
  
 100.20;
 5.20;
  
 /**
  * Strings
  */
   
 "hello"
 "hello \t world \x02\x03"
   
 /**
  * array
  */
    
 [1, "a", true, function() {}]
    
 /**
  * Objects  
  */
     
 {"key":"value","arr":[],"other":{}}
    
```  

## Comments  

`// <...>` or `/* <...> */`  

Comments can start with double slash `//` ou multiple lines with `\* *\`  

## Functions  

`var <identifier> = function (<identifierarguments>?) { <statements> }`  
`function <identifier> (<identifierarguments>?) { <statements> }`

Functions is where power of language reside, it's a first-citizen function, which mean it can accept function as arguments
or returning function. We got two ways declaring functions, literal or block.  

```
function say(name) {
    puts("Hello: " + name);
}
```

Or  

```
var say = function(name) {
    puts("Hello: " + name);
}
```  

They are exactly same, but this is illegal:  

```
var say = function say(name) {
    puts("Hello: " + name);  
}
```  

### Builtin Functions  
There are severals builtin functions that you can use:  

1. **puts** - print at console  
2. **len** - get length of object  
3. **first** - get first item of array  
4. **last** - get last item of array  
5. **rest** - get items after first one  
6. **push** - add item to array  
7. **exit** - exit program  
8. **args** - get arguments passed to ninja programs  

```
var a = [1, 2, 3, 4];
puts(len(a)); // print 4

puts(len("Hello!")); // print 5  
```

```
var a = [1, 2, 3, 4];
puts(first(a)); // print 1
```

```
puts("Hello World"); // print in screen  
```

```
var a = [1, 2, 3, 4];
puts(last(a)); // print 4
```  

```
var a = [1, 2, 3, 4];
puts(rest(a)); // print [2, 3, 4]; (all but not first)  
```

```
var a = [1, 2, 3, 4];
puts(push(a, 5)); // print [1, 2, 3, 4, 5];
```

## Import  

You can import another ninja files.  

```
import "testing.ninja";  
var lib = import "mylib.ninja"; // return function() {};  
```  

## Operators && Logics Operators  

`<expression> <operator> <expression>`  

Logic's Operators  
```
10 < 10;
10 > 10;
10 == 10;
10 != 10;
10 <= 10;
10 >= 10;
10 && 10;
10 || 10;
!10;

```  

`<expression>? <operator> <expression>`  
Arithmetics Operators  

```  
1 + 1;      // SUM
1 - 1;      // SUBTRACT
1 / 1;      // DIVIDER
1 * 1;      // MULTIPLE
4 % 2;      // MOD
10 ** 0;    // POW
10 & 2;     // AND Bitwise operator
10 | 2;     // OR Bitwise operator 
10 ^ 2;     // XOR Bitwise operator 
10 << 2;    // Shift left (multiply each step)
10 >> 2;    // Shift right (divide each step)
++1;        // First increment and then return incremented value  
--1;        // First decrement and then return decremented value  
1++;        // First return value and then increment value
1--;        // First return value and then decrement value  
```  

## Data Structures  

### Array  

`var <identifier> = [<expressions>...]`  

```
var a = [1 + 1, 2, 4, function() {}, ["a", "b"]];  
```  


#### Delete index

```
delete a[0];  
```  

It will keep the order  

#### Add Key

```
a[5] = "hello";  
push(a, "anotherKey"); 

// push by empty braces
a[] = 6;  
```  



### Hash  

`var <identifier> = {<expression>:<expression>,....}`

```
var a = {"key":"hello","key" + "key":"hello2", "other":["nice", "other"], 2: true};  
```  

#### Delete Key    

```
delete a["key"];
```  

#### Add Key  

```
a["testing"] = "hello";  
```  

### Enum  

```
enum STATUS {
    case OK: true;
    case NOK: false;
}

enum RESPONSE {
    case OK: 200;
    case NOT_FOUND: 404;
    case ERR_MSG: "There are some errors"
    case TIME: 32.2 + 43.3;
    case TEST: if (true) { 0 } else { 1 };
}
```  


then you can use his values:  

```
puts(STATUS::OK);  
```

## Conditions  

`if (<condition>) { <consequence> } else { <alternative> }`  

```
if (true) {
    puts("Hello");
} else {
    puts("Yes");
}  
```  

> Note: a value is Truth if isn't null or false, 0 will evaulated like true.  


## Loop  

`for (<initial>?;<condition>?;<iteration>?) { <statements> }`  

> **None** statements are optionals    

```
var i = 0;
for(;i<=3;++i) {
    puts(i);
}

var a = [1, 2, 3];
for(var i = 0; i <= len(a)-1; ++i) {
    puts(a[i]);
}

for(var i = 0; i <= len(a)-1; i = i + 1) {
    puts(a[i]);
}

for(;;) {
    break;
}
```  

# Object Call  

We support object call in any of data type.  

## String  

Here a full of list of support object call for string:  

```
"ola".type();                               // "STRING"
"a,b,c".split(",");                         // ["a", "b", "c"];
"hello world".replace("world", "ninja");    // "hello ninja"
"hello world".contain("hello");             // TRUE
"hello world".index("Hello");               // 0
"hello world".upper();                      // "HELLO WORLD"
"HELLO WORLD".lower();                      // "hello world"
" hello world ".trim();                     // "hello world"
"1".int();                                  // 1
"1.1".float();                              // 1.1 
```  

## Integer  

Here a full of list of support objec call for integer:  

```
1.type();               // "INTEGER"
1.string();             // "1"
1.float();              // 1.0
var a = -1; a.abs();    // 1.0
```  

## Float  

Here a full of list of support objec call for integer:

```
1.0.type();             // "FLOAT"
1.0.string();           // "1.0"  
var a = -1.0; a.abs();  // 1.0
1.5.round();            // 2.0  
```  

## Boolean  

```
true.type();     // "BOOLEAN"  
```  


## Array   

```
[1, 2, 3].type();            // "ARRAY"
[1, 2, 3].length();          // 3
[1, 2, 3].joint(",");        // "1,2,3"
[1, 2, 3].push(4);           // return null, but underlie value of array was change to [1, 2, 3, 4]  
[1, 2, 3].pop();             // return 3 and underlie value of array was change to [1, 2] 
[1, 2, 3].shift();           // return 1 and underlie value of array was change to [2, 3]  
[1, 2, 3].slice(1);          // copy array with following elements [2, 3] 
[1, 2, 3].slice(1, 1);       // copy array with following elements [2] 
```

## Object   

```
{"a":1,"b":2}.type();       // "HASH"
{"a":1,"b":2}.keys();       // ["a", "b"];
{"a":1,"b":2}.has("a");     // true
```

## Keywords  

```
var true false function delete enum 
return if else for import break case
```  

## Examples  

1. Check tests  
2. Check resolved [katas](https://github.com/gravataLonga/ninja-lang-katas)  


## Tests  

```
go test -v -race ./...  
```

## Profiling && Performance

### Where CPU resources spend more time
Create svg graph:  
[interpreter result](https://github.com/google/pprof/blob/main/doc/README.md#interpreting-the-callgraph)
```
go test -bench=. -cpuprofile cpu.prof  
go tool pprof -svg cpu.prof > cpu.svg  
```  

### Cores

```
go test -bench=. -trace trace.out  
go tool trace trace.out
```  

### Test Race Condition

```  
go test -race
```  