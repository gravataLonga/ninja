# Fixing Scope of Variables  

## Giving following code  

```
var a = "global";
function () {
    function showA() {
        puts(a);
    }
    
    showA(); // Here we expecte to print "global"  
    var a = "local";
    showA(); // We we expecte to print "global" aswell but we got "local".  
}();
```  

## Solution: Lexical Scooping  

For each identifier, register how many "hops" away are value register on *Environment  