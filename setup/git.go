package setup

/* Git UI sucks.

To be fair, the core of Git is not bad.

But its command line interface sucks.


Git Command Line Sucks

There is an xkcd commic for this:

https://xkcd.com/1597/

Unfortunately, Git GUI clients sucks too.


SourceTree

Need to login into Atlassian account before first startup.


GitEye

Add a lot of features to clutter the UI.
And most features are in a
"yes, we have this (although it is quite basic and not quite useful)" state.


GitHub Desktop

Featureless.
In fact it has fewer features than the GitHub website.


GitList

Viewer only.
Can only view files, commits, and stats.


GitKraken

Need to login via GitHub or GitKraken on first startup.


GitUp

- Minimal UI, powerful features.
- Extremely fast, particularly for large repositories.
- Open source.

But it is only available for macOS.


Workarounds

For some editors and IDEs,
combining several git plugins together actually make Git usable.
Not perfect.
But at least not as counter-intuitive like Git command line.
And they do not require registering an account.
Some tweaking of configurations may be necessary.

Alternatively, I can write aliases for Git.
In fact I have more than 50 git aliases defined in my git config.

However, to write such aliases,
I first need to understand the shit of Git UI.
Also I have to maintain them.

Instead of writing unmaintainable aliases,
I could have written extensions of Git.
libgit2 is written in C with bindings for most general programming languages.
There is also JGit for Java.
However, this requires much more effort than defining aliases.
*/
func GitSucks() {}
