/* Introduction to Programming

I have difficulties to remember things.
Thus I tend to pay special attention to survival mechanisms:

For a lot of GUI programs, it is `F1`.

For most command line programs, it is `man command` or `command -h`.

For a lot of problems,  it is web search engine.

I can get forgotten information provided that I remember the survival mechanism.

What is the survival mechanism to programming languages?

Mine is counting numbers.

In the following sections I will try to explain
how to recall forgotten programming languages based from counting numbers.

Let me Start Counting Numbers

If I remember how to count, I tend to also remember the concept of recursion.
The concept of recursion is more difficult to forget than the naming (recursion).

Just ask myself how to count:

    - Do I remember 0?
    - Yes.

    - Do I remember 1?
    - Yes, it is the number next to 0.

    - Do I remember 2?
    - Yes, it is the number next to 1.
      And since 1 is the number next to 0,
      thus 2 is the number next to the number next to 0.

    - Do I remember n?
    - Yes, it is the number next to n-1,
      thus the number next to the number next to n-2,
      thus the number next to ... the number next to 1.

The above is essentially recursion.

    - How to count n?
    - Err, I do not know yet.

    - Can I count 0?
    - Sure.

    - Can I count 1?
    - Yes. I already counted 0 in the above step,
      I just count one more time to get 1.

    - Can I count n if I already counted n-1?
    - Easy. Since I already counted n-1,
      I just count one more time to get n.

    - Then I can count n.
    - :-)

    - How can I express it in programming language?
    - I forget it.

    - Can I recall it?
    - I will try.


Technical Jargon: Peano axioms

Technical jargons are included as a personal remark
because I often forget the mapping of the jargon to the concept.


Give it Some Input

    - How can I express how to count in programming languages?
    - I forget it.

    - Let me go back. To express X in Y, I need to know what is X and Y.
    - Sure.

    - What is counting?
    - Given n, I pretend I have already counted n-1,
      and count one more time to get n.
      Then I count n-1 recursively.

    - What is programming language?
    - Instructions given to computers.

    - What is computers?
    - I cannot express it accurately.

    - Can I describe what a computer can do?
    - Yes. A computer is a machine.
      I give it some input,
      and it returns some output.

    - Does a computer have to been a machine?
    - No. I can give the input to me,
      and I can compute it by myself,
      then returns some output to me.

    - Maybe I am a machine.
    - Maybe.

    - So what do I know about a computer?
    - I know I can give it some input,
      and it computes the input,
      then returns some output.

    - What is programming language?
    - Instructions given to computers
      on how to compute some input.

    - No more?
    - Also instructions like
      "I give you this input,
       please compute it and returns the output."

    - How to count some input?
    - Given an input n, the computer need to count n-1 recursively,
      then count one more time,
      and returns.

    - How to express the above in programming language?
    - I forget it.

    - Maybe I have told computers how to count?
    - Did I?

    - Or maybe someone else have told computers how to count?
    - Maybe.

    - Assumes a computer know how to count, can I tell it to count some number?
    - I still do not recall it.

    - Is programming language a language?
    - Yes.

    - How can I do it in other language?
    - Count 42.

    - Can I just give that to a computer?
    - Only if a computer understands English.

    - Can I tell a computer how to read English?
    - I do not know.
      But I know it is much much harder than telling a computer how to count.

    - Why not make English simpler?
    - Good idea.

Let me examine the English sentence.

    Count 42.

To make it simple, I will try to remove things from it.

First I remove Count:

    42.

I am silly.

Then I remove 42:

    Count .

Hmm, it may mean counting forever.
But I do not what to count forever,
I want to count 42.

At least I can remove `.`?

    Count 42

Similarly, uppercase is not necessary:

    count 42

Can I remove the space?

    count42

No. I need something to separate function (`count`) and its input (`42`).

Really? Inputs are numbers. Functions are letters. Space is not needed.

It is absolutely necessary here.
But it makes a computer simple to understand it.
When a computer encountered a space, it knows function name ends,
and input is expected.

Making a line shorter does not always making it simpler.

    - Am I satisfied with the syntax now?
    - Yes.

    - But is it really a programming language?
    - Why is it not?

    - But it still looks similar to English?
    - It does not matter.

    - I cannot believe it is so easy to design a programming language.
    - Actually I recall some programming languages with this syntax:
      Shell script, Tcl, Haskell, Ruby.


Technical Jargon: function application


Give it a Name

Now let me go to define `count`.

Still starting from English:

    Define count as: given 0, returns 0; given n, returns count n - 1.

Change its syntax as I did in the previous section:

    define count as given 0 returns 0 given n returns count n - 1

It is hard to read.
Let me reformat it:

    define count as
        given 0
            returns 0
        given n
            returns count n - 1

Let me simplify it:

    count
        0 0
        n count n - 1

What does it mean, according to the previous section?

There is a function called define,
and I have given the following inputs to it:

    function: count
    first input: 0
    second input: 0
    third input: n
    ...
    last input: 1

This is not what I expect.

I need to add new rules.

There are line breaks and indentations in the above code.
I need to give them semantics.
They brings in new concepts, defining a function,
not applying a function as in the previous section.
So I will decide it later.

This line is still function application:

    count n - 1

It is still wrong:

    function: count
    first input: n
    second input: -
    third input: 1

    - How can I tell computer to evaluate n - 1 first?
    - How?

    - What is `n - 1`?
    - It is an arithmetic operation.

    - In arithmetic operations, how do I give priority to `n - 1`?
    - Use parans.

    - How do I tell computer to evaluate n - 1 first?
    - Use parans.

Now it is correct:

    count (n - 1)

Wait. I have introduced priority.
But have I introduced arithmetic operation?
I took it for granted that computers understanding `n - 1`.

    - So what does `n - 1` means?
    - It is an arithmetic operation.

    - Does computers understand arithmetic operations?
    - Yes? I just cannot image a computer cannot do arithmetic operations.

    - Do I have to introduce the syntax of arithmetic operations?
    - No.

    - If there is no syntax for arithmetic operations, how to express them?
    - Just use function application, e.g. `- n 1`.

    - Thus `-` is a function which takes two inputs.
    - Yes.

    - Can I do more?
    - Yes. `- n 1 2` and so on.

Thus, instead of introducing a syntax for arithmetic operations,
I extend the syntax of function application.
The syntax is simple and consistent:

    count (- n 1)

Rewrite the count definition:

    count
        0 0
        n count (- n 1)

Let me continue:

    n count (- n 1)

This means applying a function named n to two inputs:

    1. count
    2. (- n 1)

Again this is wrong.

    - How can I tell the computer `n` is not a function name?
    - I do not know.

    - What is `n`?
    - `n` represents a natural number which is not zero.

    - I mean what is `n` for `count`?
    - It is the input of `count`. Isn't it?

    - Where?
    - In `count (- n 1)`.

    - Is `n` still means an input in the first occurrence?
    - Yes.

    - Are they different?
    - I sse.
      The second n represents an input to be applied.
      But the first n is a place holder,
      which will be filled with a number in future application of count.

    - How can I mark it?
    - I have no idea yet. I am inclined to introduce another special mark.

    - Do I have to mark it?
    - Maybe not. There are indentations before the first n.

    - And how can I know that n is not equal to zero?
    - From the above line `0 0`.

    - What if I change the above line to `1 1`?
    - Then n means a number not equal to one.

    - So what does `0 0` mean?
    - It means if the input to count equals to 0,
      `count` returns 0.

    - And `n count (- n 1)`?
    - It means if the input to count is not equal to 0,
      `count` returns `count (- n 1)`,
      where `n` is binding to the input.

    - Have I finished defining `count`?
    - Yes.

    - Define a function counting a number n for m times.
    - Easy.

Such a function should be easy since I already defined `count`.

    count-many-times
        0 count n
        m count-many-times n (- m 1)

Err, how do a computer know it should match the second input against 0,
not the first input?

I need to also list the first input:

    count-many-times
        n 0 count n
        n m count-many-times n (- m 1)

The last line is ambiguous:

    one input to match: n
    two inputs to match: n m
    three inputs to match: n m count-many-times
    four inputs to match: n m count-many-times n

I need to tell the computer to group `n` and `m` together.

    - How?
    - Simple. I already introduced parans for priority.

Thus I have:

    count-many-times
        (n 0) count n
        (n m) count-many-times n (- m 1)


Technical Jargon: pattern matching


Layout

What if I mistyped a space breaking alignment of the layout:

    count-many-times
        (n 0) count n
      (n m) count-many-times n (- m 1)

Or

    count-many-times
        (n 0) count n
            (n m) count-many-times n (- m 1)

Or

    count-many-times
        (n 0)
            count n
        (n m) count-many-times
            n (- m 1)

It is prone to errors to use line breaks and indentations for syntax

Let me go back to English.

    Count 42.

Here `.` marks the end of the sentence.

    count 1. count 2. count 3.

However, `.` is ambiguous.
It is also used in fractions.

So I need to find something unambiguous.

In English, semicolons (`;`) have a similar semantics to periods (`.`).
Thus I can use semicolons.

    count 1; count 2; count 3;

But do I have to introduce a new mark (semicolon)?

Let me remove the semicolon:

    count 1 count 2 count 3

The semicolon way is to tell the computer where to break.
After breaking, `count 1`, `count 2`, and `count 3` are grouped correctly.

But I already have parans for grouping:

    (count 1) (count 2) (count 3)

Thus semicolons are not needed.

    (count 42)

Next I remove indentation:

    (count
        ((0) 0)
        ((n) (count (- n  1))))

It is ambiguous.
It is indistinguishable from the syntax of function application.

Thus I put `define` back:

    (define count
        ((0) 0)
        ((n) (count (- n 1))))


Technical Jargon: layout sensitive syntax


Be Explicit on Inputs

Given this function:

    (define f
        (((= (- n m) 42)) 0)
        ((n (< m 42)) 1) 2)

Can I know how many inputs `f` takes at a glance?

Can I know whether f is not defined for some pair of n and m at a glance?

My previous design is not sound.

For the first issue, I need to declare inputs and the matching explicitly:

    (define count (input)
        ((match input 0) 0)
        ((match input n) (count (- n 1))))

Previously, `input` is implicit, thus I have to bind the input a name `n`,
to give `n` to `count` recursively.
Here I already have input, so `n` is not needed:

    (define count (input)
        ((match input 0) 0)
        ((match input) (count (- input 1))))

The second `(match input)` is not necessary,
it just matches input against 0,
and evaluates `0` on success,
or `(count n (- n 1))` on failure.

    (define count (input)
        (((match input 0)
            0
            (count (- input 1)))))

What is `(match input 0)`?
It tests whether `input` equals 0.

    (define count (input)
        (((test (= input 0))
            0
            (count (- input 1)))))

The `test` function also solve the second issue.

We can change the name `input` back to `n` to indicate `count` accepts a number.

    (define count (n)
        (((test (= n 0))
            0
            (count (- n 1)))))


Technical Jargon: function parameters


The If Function

In the above code, `test` accepts a condition,
and returns a function.
The returning function accepts two inputs.
When the condition is true,
`test` returns a function that will evaluate its first input;
when the condition is false,
`test` returns a function that will evaluates its second input;

This is a bit complicated.
Let me flatten `test` a bit.
The new `test` function accept three inputs,
one condition, one then clause, and one else clause.
If the condition is true, `test` evaluates and returns the then clause;
If the condition is false, `test` evaluates and returns the else clause.

Let me rewrite the `count` function with `if`
(name of the new `test` function):

    (define count (n)
        (if (= n 0) 0
            (count (- n 1))))

Isn't it clearer?


Consistent Define

`define` seems useful.
I want to also use `define` to bind other values.

    (define zero 0)
    (define one 1)
    (define two (+ 1 1))

Now define have two usages:

First, as in previous section, `define` accepts a name,
a group of function parameters,
and the function body.

Second, as in this section, `define` accepts a name, and a value.

It is inconsistent to use one name for two different semantics.

Thus I can introduce a new name, for example, `variable`:

    (variable pi 3.14)

Alternatively, I can unify the two semantics into one,
i.e. introducing a new syntax
to represent function parameters and body as a value.

I choose the second approach.
If I can represent function declaration as a value,
I can pass it as a value,
for example, to another function,
without first binding it to a name.

The bare minimal syntax for a function declaration is:

    ((x y) body)

It indistinguishable to function application,
the above line can be read as:

    1. Apply `x` on `y`, which returns a function.
    2. Apply that function on `body`.

So I need a special mark.
Unlike `define`, it seems there is nothing similar in English.
And since I already introduced parans from math,
I go to math for some mark again.

Math uses a symbol like `->`.
But I want to use one letter.
Also, `-` and `>` are already used for arithmetic operations.
I want something clearly different.

`>` reveals direction, so it is clear that parameters goes first.
I try to rotate it, and get:

    v: already used as a latin letter
    <: already used as the less than operator
    ^: some systems use as pow operator

But there is a Greek letter Λ, and its lower case λ is distinguishable.

I can use backslash (`\`), which is similar to λ.
I can recall that Haskell uses backslash.
But in programming languages, `\` is often used for escaping.
Thus I will keep λ.
Well, I am lying.
The real reason behind it is the λ symbol is used for λ calculus.

    (λ (x y) body)

With the help of λ, I have a consistent define:

    (define zero 0)

    (define count
      (λ (n)
        (if (= n zero) zero
            (count (- n 1)))))


Technical Jargon: λ calculus, variable assignment


Functions with Two Inputs

The concept of λ is very powerful.
With the help of λ, I have already unified `define`.
Now I will use it to unify functions taking more than one inputs.

The `+` and `-` functions takes two inputs.
In fact they are not different to the `count` function which takes one input semantically.

Applying one function to two inputs is equivalent to
applying another function to one input,
which returns yet another function taking one input,
and be applied in turns.

For example, given a function defined as:

    (define plus
      (λ (m n)
         (+ m n)))

and applied as:

    (plus 3 4)

It is equivalent to:

    (define plus
      (λ (m)
        (λ (n)
         (+ m n))))

And I can apply it like this:

    ((plus 4) 5)

What does `(plus 4)` returns?

It returns a function:

        (λ (n)
         (+ 4 n))))

And functions with more than two inputs can also be rewritten as function with one input.

On the other side, I can rewrite a function with one input
that returns another function with another input,
to a function with two inputs.

In fact, in the previous section "The If Function",
I have changed the `test` function,
which takes one input (condition) and returns a function taking two inputs,
to the `if` function,
which takes three inputs (condition, then, else).


Technical Jargon: binary function, β-reduction


Make Define Optional

In the previous section, I introduced new semantics denoted by λ.
To keep the semantics of the language as minimal as possible.
I will try to make `define` optional.

`define` assigns some value to a name (variable),
so I do not need to write same thing repeatedly.

When applying a function,
I can just replace the variable name with its declaration:

For example, Given a function:

    (define inc
      (λ (n) (+ n 1)))

I can replace its application

    (inc 0)

with

    ((λ (n) (+ n 1)) 0)

But how do I get rid of `define` in the declaration.
In other words,
how can I bind the function declaration to name `succ` without `define`?

Let me review the function application:

    ((λ (n) (+ n 1)) 0)

Here 0 is bound to the name `n` during function application.

So I can bind the function declaration to a name
via another function with that name as its parameter
during the application of that function.

However, in the application of `succ`, 1 is returned.
I need to have the function declaration returned.
In other words, I need a function,
which takes a function declaration as input,
and returns the same declaration as output.

It seems simple.
I just need a function that returns its input back.

    (λ (x) x)

Now I can declare `succ` without `define`:

    ((λ (inc) inc)
     (λ (n) (+ n 1)))

And apply it like this:

    (((λ (inc) inc)
      (λ (n) (+ n 1))) 3)


Technical Jargon: fixed point, identity function


More on Making Define Optional

However, I have difficulties to use the above technique for `count`.
I can use (λ (count) count) to bind the function to `count` variable,
but the definition part also uses `count`.
That is, `count` refers to itself.

Let me first work a simpler example.

The simplest self reference function I can image is:

    (define f
     (λ (x)
        (f x)))

It continues to apply itself, endlessly.

Try to remove `define`:


    ((λ (f) f)
     (λ (x)
        (f x)))

The `f` in the last line is undefined.

What if `f` is defined else where,
then we can bind it to a new name, say `g`:

    (define f
     (λ (x)
        (f x)))

    ((λ (g) g)
     (λ (x)
        (f x)))

How can I bind `f` without using `define` elsewhere?
Again, via function parameters.
I pass `f` itself as an additional parameter,
so I can bind it to

    ((λ (f x)
        (f f x)))

And pass itself on application:

    ((λ (f x) (f f x))
        (λ (f x) (f f x))
        0)

Now bind it to a new name `g`:

    (λ (g)
       (λ (f x) (f f x))
        (λ (g)
          (λ (f x) (f f x))))

And apply it to a single input:

    ((λ (g)
        (λ (f x) (f f x))
         (λ (g)
           (λ (f x) (f f x))))
                (λ (f x) (f f x)
                3))

This is semantically equivalent to

    (g 3)

As mentioned before,
functions that are semantically equivalent can use different parameter names:

    (λ (x) x)
    (λ (y) y)

Thus we have this equivalent form of the above `g` function:

    (((λ (g)
      (λ (s y) (s s y)))
      (λ (t)
         (λ (f x) (f f x)))) (λ (h z) (h h z)) 3)

On the other side, I can reuse variable names:

    (define x
      (λ (x) x))

    (x 4)

Binding via function parameters is the same:

    (((λ (x) x)
      (λ (x) x)) 4)

Thus this is also equivalent:

    (((λ (f)
      (λ (f x) (f f x)))
      (λ (f)
         (λ (f x) (f f x)))) (λ (f x) (f f x)) 3)

Let me go back to the definition:

    ((λ (f)
     (λ (f x) (f f x)))
     (λ (f)
        (λ (f x) (f f x))))

I wrote the same function twice.
How to eliminate duplicated code?

If I have `define`, then ...
No, I do not want to use `define`.

I need a function, it receives one function expression as input,
and returns the input twice.

Definition of this function is straightforward:

    (λ (x) (x x))

Apply it:

    ((λ (s) (s s))
     (λ (f)
     (λ (f x)
        (f f x))))

Here I use the variable `s`, though I could have written it as:

    ((λ (x) (x x))
     (λ (f)
     (λ (f x) (f f x))))

But the `x` on the first line and `x` on the last line are different things.
Thus I use different variable names to make the code clearer.

On the other side,
although `f` on the second line and `f` on the third line are different variables,
they refer to the same function, so I use the same letter.

Compare the inner function:

    (λ (f x)
       (f f x)

with the original one:

    (λ (x)
       (f x))

Can I remove the additional parameter `f`?

Recall that binary function declaration:

    (λ (x y) ...)

can be rewritten as

    (λ (x)
       (λ (y) ...)

And its application

    (f x y)

can be rewritten as:

    ((f x) y)

Thus I rewrite this:

    (λ (f x)
       (f f x)

to this:

    (λ (f)
      (λ (x)
        ((f f) x)))

After the rewrite:

    ((λ (s) (s s))
     (λ (f)
        (λ (f)
           (λ (x)
              ((f f) x)))))

This is unnecessary:

    ...
    (λ (f)
       (λ (f)
       ...

It means there is a function, which takes an input `f`,
then returns another function which takes the same input `f`.
It is equivalent to directly pass `f` to the inner function.
I can safely eliminate the outer wrapping.

    ((λ (s) (s s))
     (λ (f)
        (λ (x)
           ((f f) x)))))

The last two lines look similar to the original definition:

    (λ (x)
       (f x))

Both `(f f)` and `f` are functions.

If I can replace `(f f)` with `g`,
then structure will be the same.

How?

    (λ (s) (s s)) (in)

Here it receives one input and duplicates it.
Now I want to do the opposite.

    (λ (in) in) (s s)

Thus I have:

    ((λ (s) (s s))
     (λ (f)
        ((λ (g)
            (λ (x)
               (g x))) (f f))))

This part is almost the original definition:

    (λ (g)
       (λ (x)
          (g x)))

Fact it out:

    (λ (h)
       ((λ (s) (s s))
        (λ (f)
           (h (f f)))))

Let me apply it to the original definition:

    (((λ (h)
       ((λ (s) (s s))
        (λ (f)
           (h (f f)))))
      (λ (f) (λ (x) (f x)))) 0)

Evaluate it to check if it does work, i.e. replace `h` with

    (λ (f) (λ (x) (f x)))

I have:

    ((λ (s) (s s))
     (λ (f)
        ((λ (f)
            (λ (x)
               (f x)))
        (f f))))

Expand it:

    ((λ (f)               ; 2
       ((λ (f)
           (λ (x)
              (f x)))
       (f f)))            ; 3
       (λ (f)             ; 1
          ((λ (f)
              (λ (x)
                 (f x)))
          (f f))))        ; 4

The evaluating order would be:

                   +----+
                   |    |
                   V    |
    1 -> 2 -> 3 -> 1 -> 4

In other words, the whole program would sank into `(f f)` and never came out.

The problem lies in this transformation:

    (λ (in) in) (s s)

`(s s)` is a function application,
and it will pass the second s to the first s function,
and evaluates the first s before binding to `in`.
However, if the definition of s itself repeats this self application `(s s)`,
then the program will sink into there and will never come out.

How to fix this issue?
I need to tell the computer to not evaluate `(s s)` for now,
deferring its evaluation until it has reached `((f f) x)`.

How can I do this?
I can introduce a new feature to mark something to be evaluate later.

Do I have to?

The above transformation looks similar to the transformation below:

    (λ (s) (s s)) (in)

Why the latter transformation does not have this issue?

Because `in` is a function declaration, not a function application.
Function body will not be evaluated on declaration,
and its evaluation will be deferred on application.
So I already have something to defer evaluation.
Therefore I do not need to introduce a new feature.

I can just wrap it into a declaration of another function:

    (λ (w)
        (f f))

So the evaluation of `(f f)` will be deferred until

    ((λ (w)
        (f f)) x)

However, the above wrapping changed the semantics.
In previous versions,
after evaluating `(f f)`, the returned function will be applied on `x`.
But on the wrapping version,
`x` will be passed to `w` and discarded before evaluating `(f f)`.
To maintain the semantics of `((f f) x)`,
I change the body of wrapping function to `((f f) w)`:

    (λ (w)
       ((f f) w))

Finally, I have a working version:

    (λ (h)
       ((λ (s) (s s))
        (λ (f)
          (h (λ (w) ((f f) w))))))

I can use it like this:

    ((λ (h)
       ((λ (s) (s s))
        (λ (f)
          (h (λ (w) (f f))))))
     (λ (x) (f x))
     0)

And I can apply it to the `count` function without editing code:

    ((λ (h)
       ((λ (s) (s s))
        (λ (f)
          (h (λ (w) (f f))))))
     (λ (n)
        (if (= n zero) zero
            (f (- n 1))))
     42)

Thus it is the version of

    (λ (x) x)

for recursive functions.


Technical Jargon: Y combinator, η-conversion


Number

Go back to the `count` definition:

    (define count
     (λ (n)
        (if (= n zero) zero
            (count (- n 1)))))

I have shown that something seems essential like `define`
turns out to be optional.

Can I make other things optional?

Let me examine them carefully.

First, there is 0 and 1.

Most humans have 10 fingers.
Thus 0-9 is naturally defined for humans.

Most machines do not have 10 fingers.
But electric units have one switch that can be on or off.
Thus 0 and 1 is naturally defined for machines.

Some humans do not have 10 fingers.
Thus 0-9 is not naturally defined for them.
But they do have one head that can nod or shake.
Thus 0 and 1 is also naturally defined for all humans.


Technical Jargon: Binary number


Denotation of Numbers

Although 0 and 1 itself is naturally defined,
but `0` and `1` is just a randomly selected denotation.
We can denote them by "one" and "two",
turning off and on the switch,
or shaking and nodding the head.
The denotation is endless, but the concept remains the same.

All denotations may have its own advantage under certain conditions.
For example, 0 and 1 is easy to write with a pencil.
On and off is easy to "write" for a machine,
just imaging how difficult it would be for a machine to denote 0 as turning on a circle of light bulbs.

But the problem of the denotation mentioned above has a disadvantage:
they do not reveal the relation between 0 and 1.
They are not saying 1 is next to 0,
or on is next to off,
or nodding is next to shaking.

The problems remains same for 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, etc.,
or 001, 010, 011, 100, 101, 110, 111, etc.

How to denote numbers revealing their relations?

Let's go back to the counting code above.

    (define count
     (λ (n)
        (if (= n zero) zero
            (count (- n 1)))))

Apply it to 0, i.e. give 0 as its input:

    (count 0)

since

    (= 0 0)

it just returns 0, i.e. gives 0 as its output.

Now let's apply it to 1, i.e. give 1 as its input:

    (count 1)

Let's check

    (= 1 0)

No, 1 is not equal to 0.
Thus `count` continues to the else branch:

    (count (- 1 1))

That is, `count` applies 0 (- 1 1) to `count` then returns 0.

Now let's apply 2, `count` applies 1 (- 2 1) to `count`,
which in turns applies 0 (- 1 1) to `count`,
then returns 0.

Generally, when applying `count` to n, `count` applies n-1 to `count` recursively,
until `(count 0)`.

In other words, given a number n, `count` applies itself n times to return 0.
That is counting.

Thus, we can use the times a function applies itself to represent numbers.

    (denote 0) ; ->  0
    (denote 1) ; -> (denote 0)
    (denote 2) ; -> (denote (denote 0))
    ; ...
    (denote n) ; -> (denote (denote ... (denote 0)))

The above denotation assumes two concepts:

First, 0, i.e. the base case.
The denotation `0` is randomly chosen, and I can replace it with `x` or any other symbol.
But the concept remains.

Second, `(denote ...)`, i.e. function application.
Again, I can use other denotation, e.g. `(f x)` or `denote(x)`.

Since I am using times of function application to represent numbers,
I cannot get rid of the concept of function application.
Also, function application is essential to the programming language discussed above.
But can I get rid of 0?

Let me try.

    (denote 0) ; -> 0

means I am using 0 to denote 0.

    (denote (0) 0)

And since I can use 0 to denote 0,
I can also use 'denote 0' to denote 0:

    (denote (0) (denote 0))

and that 'denote 0' can be expanded as using 0 to denote 0:

    (denote (0) ((denote (0) 0) 0)

In other words, we can use `(denote (0) 0)` to represent 0:

    (denote (0) 0) ; 0
    (denote (denote (0) 0)) ; 1
    (denote (denote (denote (0) 0))) ; 2
    ; ...

In the above representation, 0 does not have special meaning.
Also, `denote` does not have any special meaning like `count` mentioned above.
To make this clear, replace `0` with `x`, and `denote` with f:

    (f (x) x)           ; 0
    (f (f (x) x))       ; 1
    (f (f (f (x) x)))   ; 2
    ; ...

I use function application times for numbers.
But function application is like an action,
and numbers are things.
It is usually to use declaration to define things.
Thus I change the definition from application to declaration:

    (λ (f)
       (λ (x) x))          ; 0
    (λ (f)
       (λ (x) (f x)))      ; 1
    (λ (f)
       (λ (x) (f (f x))))  ; 2


Technical Jargon: Church encoding


Arithmetic operations

Next let me define `-`.

    How to define `-`?
    I forget it.

    It seems difficult. Do I really need it?
    No. In fact I just need (minus-one n).

    Then, how to define minus-one?
    I forget it too.

    Can I define `(minus-one 0)`?
    No. I forget about negative numbers.

    Can I define `(minus-one 1)`?
    Yes.  (λ (f) (λ (x) (f x)))

    Can I define (minus-one n) if I already defined (minus-one (minus-one n)).
    Only if I know how to remove one level of f.

    How to remove one level of λ declaration?
    I have no idea.

    What about the opposite?
    It seems easier.

I can change the `succ` function declared before:

    (define inc
    (λ (n)
       (+ n 1)))

to something like:

    (define inc
    (λ (n)
       ...))

What is the missing `...` part?
I am not sure.
But since the `succ` returns a number,
with the new definition of number,
It should be something like:

    (define succ
    (λ (n)
       (λ (f)
          (λ (x)
             (f ...)))

And this time, the `...` part should be the same thing in `n`:

    (λ (f)      ; n
       (λ (x)
          (...)))

How can I get that part from `n`?
Simple. Just apply it.

    ((n f) x)

Now fill the missing part to the new definition of `succ`:

    (define succ)
    (λ (n)
       (λ (f)
          (λ (x)
             (f ((n f) x))))))

The minus-one definition should have a similar structure:

    (define minus-one
    (λ (n)
       (λ (f)
          (λ (x)
             (...)))

The missing `...` should be the same thing in `n`, without the outmost `f`:


    (λ (f)      ; n
       (λ (x)
          (f ...)))

How can I get that part of `n`?
Also apply it.
But instead of just applying it to `f`,
I need to apply some function f satisfying

    (f something)

will be evaluated to

    something

I have already met this function a few times before:

    (λ (x) x)

This function may look too simple to be useful at first.
If I give the input to computer,
and the computer will returns the same input back as output,
why bother give the input to this function?
Why not just use the input directly?
However, it turns out this simple function is extremely useful.

    (λ (n)
       (λ (f)
          (λ (x)
             (f ((n (λ (v) v)) x)))))

Hmm, it turns out minus-one is not that easy.
Applying `n` on `(λ (v) v)` does remove the outmost `f`,
but also changes the semantics of the number.
Because now `f` must be `(λ (v) v)`,
but `f` can be any function in the original definition.

Since minus-one is hard, I will defer its definition later.
Maybe I will have better tools later.


Technical Jargon: λ calculus, higher order function, successor function


Boolean

Now I will implement the if function.

Let me go back to the meaning of the `if` function:

The `if` function takes three inputs,
one condition, one then clause, and one else clause.
If the condition is true, it evaluates and returns the then clause;
If the condition is false, it evaluates and returns the else clause.

So to define if, I first need to define true and false.

    - How?
    - I do not know.

    - Well, maybe numbers can be reused to represent true and false.
    - Probably not. Booleans and numbers have different semantics.

    - Well, at least I can define booleans in a way similar to numbers.
    - Good idea. I will define true and false as functions.

But what is the definition of true and false?

Re-read the meaning of `if`:

If the condition is true, it evaluates and returns the then clause;
If the condition is false, it evaluates and returns the else clause.

Reverse it, I have:

If it evaluates and returns the then clause, the condition is true;
If it evaluates and returns the else clause, the condition is false.

Here the definition of true and false occurs naturally after their motivation.
Both true and false will accept two inputs, then clause and else clause.
True will return the first input (then clause),
while false will return the second input (else clause).

    (define true
      (λ (then-clause else-clause) then-clause))

    (define false
      (λ (then-clause else-clause) else-clause))

This definition of true and false makes `if` trivial to write:

    (define if
      (λ (condition then-clause else-clause)
         (condition then-clause else-clause))

When applying `if` on condition, then-clause, else-clause,
both then-clause and else-clause will be evaluated on passage.
However, one and only one of these two clauses will be returned.
Thus evaluating both clauses may waste some time
when theses clauses are complex.

As an optimization, I can pass then and else clauses wrapped in function declaration,
a technique I have used before.


Technical Jargon: Church booleans


False and Zero

Review the definition of false:

    (define false
      (λ (then-clause else-clause) else-clause))

Convert it to the one parameter form:

    (define false
      (λ (then-clause)
        (λ (else-clause) else-clause)))

This is in fact the definition of zero!

To make this clear, replace `then-clause` with `f`, and `else-clause` with `x`:

    (define false
      (λ (f)
        (λ (x) x)))

So the definition of false is indistinguishable to the definition of zero,
which may cause confusion.

On the bright side, this coincidence makes testing if a number is zero very easy.

    (define zero?
      (λ (n)
        (n ... true)))

The function `zero?` accepts a number.

If the number is zero, then it is also false,
and can be used as false to return true.

If the number is not zero, then it will apply the number on `...`.

Here number accepts one input, but false accepts two.
Thus I use the one parameter form of false:

    (define false
      (λ (f)
        (λ (x) x)))

    ((false then-clause) else-clause) ; -> else-clause

Now I rewrite `zero?` with the new definition of false:

    (define zero?
      (λ (n)
         ((n ...)
            true)))

Try to fill it with false:

    (define zero?
      (λ (n)
         ((n false)
            true)))

If n is not zero, then apply `n` on false, taking number 3 as example:

    (λ (x)
        (false (false (false x))))

No matter what `x` is (in fact `x` is true),
and no mater how many levels of `false` nested,
according to the definition of false,
`(false x)` would always be `(λ (y) y)` until the outmost false:

    (λ (true)
       (λ (y) y))

That is, it returns false in the end when `n` is not zero.


More on Arithmetic Operations

With zero? defined, equal function is trivial to define if minus is defined:

    (define =
        (λ (m n)
            (zero? (- m n))))

And minus is also trivial to define with recursion if pred is defined:

    (define -
        (λ (m n)
            ((zero? n) m
             (succ (- m (pred n))))))

Plus is similar:

    (define +
        (λ (m n)
            ((zero? n) m
             (succ (+ m (pred n))))))

And multiplication:

    (define *
        (λ (m n)
           ((zero? n) zero
            (+ m (* m (pred n))))))

So I need to define `pred`.


The Predecessor

In the definition of `zero?`, I pass the false function to `n`,
and the false function has been called n times.

What if I pass a counter function to `n`,
which starts to count at -1?
Thus it will be counted n-1 times.

    (define counter
        (λ (c)
           (succ c)))

So I apply `n` on `counter` (as `f` parameter) and `-1` (as `x` parameter).
And `counter` will be called n times during the application,
by the definition of numbers.
And `counter` will returns `n-1` in the end.

But I do not know what is -1.
How to define -1?

Suppose I have defined -1, then its successor will be 0:

    (= (succ -1) 0)

By the definition of succ:

    (-1 f x)

evaluates to

    x

On the other side:

    (0 f x)

i.e.

    (((λ (f)
         (λ (x) x)) f x))

evaluates to

    x

Thus -1 and 0 is the same.

However, the counter can start count at the second call.
That is, it uses another variable to record if it is the first call.

    (define counter
        (λ (c first-call)
           (first-call (false c) (false (succ c)))))

This does not work.

First, both `(false c)` and `(false (succ c))` will be evaluated before returning.

Second, counter accepts two inputs,
but the definition of numbers requires it to accept only one input.

Again, I can wrap them in a function declaration
to defer evaluation and to combine them into one input.

    (define counting
        (boolean number)
            (λ (f)
               ...))

How to fill in the missing part?
If I just want to combine them one and defer evaluation later, I can just write:

     (define counting
        (boolean number)
            (λ (f)
               boolean number))

However, to check if it is the first call, and to get the number at the end,
I need a way to deconstruct it.
In other words, I also need to get boolean or number back.

Recall the definition of true and false,
then-clause and else-clause can be viewed as a pair of elements,
and true and false can be viewed as a way
to retrieve the first and second element accordingly.

Thus I have

    (define counting
        (boolean number)
            (λ (f)
               f boolean number))

And I can use it like this:

    (define example
        (counting true 0)

    (define first
        (example true))

    (define second
        (example false))

It turns out that this works against any values,
the first does not have to be a boolean,
and the second does not have to be a number.
So I swap the naming to make this clear:

    (define pair
        (λ (x y)
            (λ (f)
               f x y)))

    (define first
        (λ (some-pair)
            (some-pair true))

    (define first
        (λ (some-pair)
            (some-pair false))

Now I can rewrite the counter function:

    (define counter
        (λ (counting)
           ((first counting)
                (pair false 0) (pair false (... counting)))))

Not yet.
I need to define a version of `succ` for pairs.
Fortunately it is easy.
I just need to deconstruct the pair,
call `succ` on the second value,
and reconstruct a pair with the original first value and the increased second value:

    (define increment
        (λ (some-pair)
            (pair (first some-pair) (succ (second some-pair)))))

Now I can finish the definition of counter:

    (define counter
        (λ (counting)
           ((first counting)
            (pair false 0)
            (pair false (increment counting)))))

And finally the pred function:

    (define pred
        (λ (n)
           (second (n counter (pair true 0)))))

With the help of `succ` itself and other tools,
its definition turns out to be simple, like `succ`.


Conclusion

The simple count function is built with a few other functions,
and all these functions can be defined
from the concept of function declaration and application.

    (define count
        (λ (n)
           ((zero? n) zero
                      (count (pred n)))))

    function
    |
    |_ function with multiple inputs
    |  |_ (λ (x y) ...)
    |  |_ (λ (x (λ (y) ...)
    |
    |_ define
    |   |_ (λ (x) x)
    |   |_ (λ (h)
    |         ((λ (s) (s s))
    |         (λ (f)
    |            (h (λ (w) ((f f) w))))))
    |
    |_ booleans
    |   |_ true  (λ (x) ((λ (y) x)))
    |   |_ false (λ (x) ((λ (y) y)))
    |  +>_ zero? (λ (n) ((n false) true)))
    |  ||_ pair (λ (x y) (λ (f) f x y))
    |  |   |_ first (λ (p) (p true))
    |  |   |_ second (λ (p) (p false))
    |  |   |_ increment (λ (p) (pair (first p) (succ (second p))))
    |  |         |_ counter  (λ (c) ((first c)                 ^
    |  |            |                (pair false 0)            |
    |  |            |                (pair false               |
    |  |            |                      (increment c))))    |                    ^
    |  |            |                                          |
    |  |            |_ pred (λ (n)                             |
    |  |                       (second                         |
    |  |                       (n counter (pair true 0))))     |
    |_ numbers                                                 |
       |_ 0 (λ (f) (λ (x) x))                                  |
       |_ 1 (λ (f) (λ (x) (f x)))                              |
       |_ succ (λ (n)    --------------------------------------+
                  (λ (f)
                     (λ (x)
                     (f ((n f) x)))))

Isn't it elegant?
*/
package introduction