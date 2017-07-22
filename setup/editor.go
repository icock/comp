package setup

/* Editors pursuit productivity in a wrong way.

To maximize productivity, editors provide a lot of keyboard shortcuts.
However, keyboards have limited numbers of keys.
And editors have three choices:


Mode Switch

Keyboard shortcuts of Vi is easy to type.
However, the mental cost of mode switch is high.


Two-Step Shortcuts

Emacs uses two-step keyboard shortcuts to avoid the hassle of mode switch.
However, they are slow to type.


More Than Two Keys

Shortcuts like `ctrl+alt+x` and `ctrl+alt+shift+y` are hard to type.


All Are Hard to Remember

All three choices above have a common disadvantage:

they are difficult to remember.


The Alternative Way

Just like typing speed should not be the bottleneck of programming,
keyboard shortcuts should not be the bottleneck of programming productivity too.

It is navigation, completion, and refactoring
that boost programming productivity mostly,
not keyboard shortcuts.

Forget about keyboard shortcuts, replace the editor with an IDE,
or configure the editor with a language framework
which provides assistant with navigation, completion, and refactoring of code.


The Only Keyboard Shortcut to Remember

IntelliJ introduces "Find Action" (`ctrl+shift+a`) in 2007,
and Sublime Text introduces a similar feature called "Command Palette" (`ctrl+shift+p`) in 2011,
which inherited by Atom and VS Code.
For editors and IDEs provideing such a feature,
it is the only keyboard shortcut need to remember.

If I find out that I am using an action so frequently
that I can afford the cost of remembering an additional keyboard shortcut,
I can always bind the action to a keyboard shortcut.
Since there won't be too many such keyboard shortcuts,
none of mode switch, two-step shortcuts, and more-than-two-key shortcuts is necessary.
*/
func EditorsSuck() {}