/* My Setup Sucks

I tried my best to choose the one that sucks less.
*/
package setup

/* I can not name a programming languages that does not suck.

All Dynamic Typed Languages Sucks

Except for tiny programs,
static typing helps programmers to detect errors,
and makes refactoring programs easy.
Dynamic typed languages do not have this important mechanism.

Dynamic typed languages also have other issues.

To name a few:


Lisp

Dynamic scoping. Period.


Some Later Dialacts of Lisp

Macros make debugging hard.


Lua

Everything is a table does not work if used as a general programming language.


Perl

Perl encourages to write unreadable code.


AWK

Worse than Perl.


Ruby

A cleanup of Perl.
Just like Perl, it may feel nature to their creators.
But I often surprised by it.
Also, the classes in standard library usually have too many methods.
Some methods are just aliases.
Some methods are just with slight difference.
Saving a few key strokes do not make up the pollution of auto completion list.


Python

Indentation sensitive syntax is error prone.
One careless tab will break the language.

Although Ruby is messy,
at least it admits sometimes there are different ways to do the same thing.
Python advertises itself as "there is only one obvious way to do one thing".
But there is Python 2 and Python 3.

Okay, since the support of Python 2 will end soon
(they always said so, but then they just extend the support period again and again),
let me just consider Python 3.

Even in Python 3 there are many ways to do one thing.
Take string formatting for example,
there are three ways,
C style printf format string,
template string  `"print({variable}", variable=value)`,
and the f-string `f"{expression}""`.
Another example is path and pathlib.


Julia

Julia supports all the following operators for comparison:

	> < >= ≥ <= ≤ == === ≡ != ≠ !==
	≢ ∈ ∉ ∋ ∌ ⊆ ⊈ ⊂ ⊄ ⊊ ∝ ∊ ∍ ∥ ∦
	∷ ∺ ∻ ∽ ∾ ≁ ≃ ≄ ≅ ≆ ≇ ≈ ≉ ≊
	≋ ≌ ≍ ≎ ≐ ≑ ≒ ≓ ≔ ≕ ≖ ≗ ≘ ≙
	≚ ≛ ≜ ≝ ≞ ≟ ≣ ≦ ≧ ≨ ≩ ≪ ≫ ≬ ≭
	≮ ≯ ≰ ≱ ≲ ≳ ≴ ≵ ≶ ≷ ≸ ≹ ≺ ≻ ≼ ≽
	≾ ≿ ⊀ ⊁ ⊃ ⊅ ⊇ ⊉ ⊋ ⊏ ⊐ ⊑ ⊒ ⊜
	⊩ ⊬ ⊮ ⊰ ⊱ ⊲ ⊳ ⊴ ⊵ ⊶ ⊷ ⋍ ⋐ ⋑ ⋕
	⋖ ⋗ ⋘ ⋙ ⋚ ⋛ ⋜ ⋝ ⋞ ⋟ ⋠ ⋡ ⋢ ⋣ ⋤ ⋥
	⋦ ⋧ ⋨ ⋩ ⋪ ⋫ ⋬ ⋭ ⋲ ⋳ ⋴ ⋵ ⋶ ⋷ ⋸ ⋹ ⋺ ⋻
	⋼ ⋽ ⋾ ⋿ ⟈ ⟉ ⟒ ⦷ ⧀ ⧁ ⧡ ⧣ ⧤ ⧥
	⩦ ⩧ ⩪ ⩫ ⩬ ⩭ ⩮ ⩯ ⩰ ⩱ ⩲ ⩳ ⩴ ⩵ ⩶
	⩷ ⩸ ⩹ ⩺ ⩻ ⩼ ⩽ ⩾ ⩿ ⪀ ⪁ ⪂ ⪃ ⪄ ⪅ ⪆
	⪇ ⪈ ⪉ ⪊ ⪋ ⪌ ⪍ ⪎ ⪏ ⪐ ⪑ ⪒ ⪓ ⪔ ⪕ ⪖
	⪗ ⪘ ⪙ ⪚ ⪛ ⪜ ⪝ ⪞ ⪟ ⪠ ⪡ ⪢ ⪣ ⪤ ⪥
	⪦ ⪧ ⪨ ⪩ ⪪ ⪫ ⪬ ⪭ ⪮ ⪯ ⪰ ⪱ ⪲ ⪳ ⪴ ⪵
	⪶ ⪷ ⪸ ⪹ ⪺ ⪻ ⪼ ⪽ ⪾ ⪿ ⫀ ⫁ ⫂ ⫃ ⫄
	⫅ ⫆ ⫇ ⫈ ⫉ ⫊ ⫋ ⫌ ⫍ ⫎ ⫏ ⫐ ⫑ ⫒ ⫓ ⫔
	⫕ ⫖ ⫗ ⫘ ⫙ ⫷ ⫸ ⫹ ⫺ ⊢ ⊣

To make things worse, some operators are same under certain conditions.
For example, in for loops and array comprehensions,
`=`, `in`, and `∈` is the same.

To make things even worse, all theses operators can de extended by users!

Like some dialects of Lisp, Julia also supports macros,
though its debugging information is better.

BTW, although as a language Julia sucks,
it shines when it is used as an advanced calculator.


PHP

The list of its flaws is too long to be list here.


JavaScript

It lacks basic mechanisms like modules.
Also, the semicolon auto insertion rules and type auto conversion rules are unnatural.
Here I am talking about modern JavaScript,
e.g. `const` and `let` instead of `var`,
`=>` instead of `function`, and so on.


Groovy

Groovy pretends that it has fix Java's Null issue
by returning null instead of throwing NullPointException.
Passing null everywhere does not fix the problem and makes debugging harder.

Groovy also does not fix Java's non first-class function issue.
Well, at least in Groovy a method can be converted to first-class
by prefixing it with `&`,
instead of wrapping in an anonymous class which implements a function interface.

And Groovy introduces new problems Java does not have.
I will give two examples.

One is an implementation issue.
For example, in Groovy 1, private fields and methods are not private.
And Groovy 2 still does not fix this.
Will Groovy 3 fix this?

The other is a design issue.
Java's checked exception is evil.
But just removing it is more evil.


Static Typed Languages May Suck Less

In above sections, I have listed issues of popular dynamic typed languages,
other than the dynamic typing itself.

Similarly, although static typed languages do not have the issue of dynamic typing,
they may have other issues.


Assembly Language

No abstractions.
Bound to a certain type of machines.


Fortran

To many confusing features to list.
It lacks abstractions, encouraging unmaintainable code.


C

Accessing memory out of the bound is hard to detect in C,
e.g. array index and pointer arithemetic.
Managing memory by hand is tedious and distracting for a general programing language.


C++

It tries to fix C's problems,
but in the end it brings in so much complexity,
which causes more troubles than improvements.


Java

Functions are not first-class, NullPointerException,
unnecessary differences between value type and reference type,
and generic type erasure.
To be fair, those issues, or similar ones also exist in C and C++.


C#

C# is similar to Java.
It fixes Java's generic type erause.
However, like Groovy, it just removes Java's checked exception.


Haskell

Just like many other static typed languages (C, C++, Java, C#, etc),
in most Haskell programs function signature is explicitly written out.
Thus HM type system's global type inference does not make any difference,
except that it makes programs harder to write.
To fix the limits of HM type system,
Haskell introduces concepts from more advanced type system,
which makes the language more complex, like C++.

Also, the unnecessary pursuit for pureness and lazyness
make code even harder to write.

To be fair, Haskell is not bad as an experimental languages for type systems,
just do not abuse it for practical projects.


Go

Similar to Java and C#, except that
Go fix the non first-class function problem of Java and C#,
and introduces new problems,
e.g. abusing multiple return values to represent errors,
zero values,
and no generics.


Kotlin

Unlike Groovy, it fix Java's NullPointerException.
Like Groovy, it just removes checked exception and use `::method`.
Also, Java's generic type erasure remains.

Kotlin borrows too many features from other languages,
making the language complex.


Swift

Like Kotlin, Swift fix the nullability problem with optional typing.

In Swift, functions need to declare if it `throws`,
without specifying the types of exceptions it may throw.
Thus the caller is forced to have a default cause,
i.e. `catch (Exception)` in Java.
If necessary, a function that may throw may be converted to an optional type.
This is a crippled balance between the soundness and verbosity of checked exceptions.

In Swift, only classes are reference types.
All collection types (arrays, sets, and dictionaries) are value type.
Thus despite the Swift documentation at apple.com stated that
the mutability of collections depends on
whether it is be assigned to a variable or const,
in fact all collection types are immutable.
Mutable collections returns a new collection under the hood.
In other words, Swift redefine mutability:
a mutable value is an immutable value assigned to a reassignable variable,
and it has methods to mimic a mutable value.
The implementation uses copy-on-write to reduce the cost of copying,
which may cause unpredictable slow down of the program.

To be fair, the two issues above may not cause serious problems in practical projects.
But Swift makes breaking changes on every versions,
which is a serious problems in practical projects.
*/
func ProgrammingLanguagesSuck() {}
