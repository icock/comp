/* My Personal Coding Style

Coding style is important because readability of code is important.
In other words, good coding style improves readability.
This could be all I need to know about coding style.
However, sometimes readability is vague.
Thus I need other more concreate principles to pursuit better coding style.


Coding Style Should Be Personal

The common way is to stick to the most popular style in the community.
But I prefer my own way.

First, for most projects, if not all projects,
I am the main, if not only, reader of the code.

Second, popularity does not mean clearness.

Third, different people have different ways of thinking,
and different tastes.


Coding Style Should be Universal

Do not create individual coding style for different languages.
Instead, stick to one universal coding style across languages.

After all, now matter how different languages are,
they are all programming languages.

If one concept can improve readability in one programming language,
why cannot it be applied to another programming language?
For example, if I am satisfied with the if function in scheme:

	(if condition then-clause else-clause)

Then I can write the following code in C:

	if (condition) {
		then-clause;
	} else {
		else-clause;
	}

even though `if` is not a function in C.
The advantage of this code style is to avoid not considering all conditions.

If a coding style is not possible or very difficult to port to another language,
then I will check the following things:

First, a coding style rule should be a high level concept,
not a low level dogma restricted to certain languages.

An example of a low level rule is that
the opening and closing braces of a block should be on the same column.
It cannot be applied to languages such as
Lisp (not using braces at all),
Python (using indentation instead of braces),
and Go (forcing the opening brace on the same line).

The high level version of this rule is
opening and closing marks of structures of same semantic level
should be matched up consistently.

This high level concept applies to programming languages universally,
and it covers more conditions,
for example, elements in lists, etc.
Whether the opening and closing marks are visible or invisible,
and how the opening and closing marks match up visually,
is low level details to a specific language.

Second, another language may not have the mechanism to apply the style.
It is normal that different programming languages provide different mechanisms.
But most mechanisms are universal, across languages.
A few certain rules of a style may have difficult to find their room in a new language,
but the style as a whole should fit in the new language well.
Otherwise the new language is not general enough,
lacks of a lot of important mechanisms.

For example, I think it is a good idea to declare variables with an explicit type,
instead of a mark for inferred type like `auto` in C++.
But a lot of dynamic typed languages do not support static typing.
So this rule does not apply to those languages.
But it is still unnecessary to maintain separated styles
for static typed and dynamic typed languages.

Third, am I using some clever corners of the language?
Or some distinct and trivial features of the language?
If so, then rewrite it in a simple, explicit, and stupid way.
Do not assume readers are familiar with all corners of the language,
or distinct and trivial features absent in other languages.

For example, in some languages zero is false.
Thus I could have written:

	if (integer) {
		// when integer is not zero
	}

instead of the long, explicit, and stupid version in other languages:

	if (integer != 0) {
		// when integer is not zero
	}

But I prefer the latter, even in languages where zero is false.


Consistent and Automatic Formatting Is Preferred

Formatting is mostly low level, so it does not need too much attention.

For formatting, consistence and automation is more important.

For example, Go unifies formatting via `go fmt`.
Although some choices of `go fmt` are against my taste,
I still think `go fmt` is an advantage of Go.
`go fmt` format code for me, so I do not have to format it manually.
*/
package style