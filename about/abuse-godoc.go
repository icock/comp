/* Abuse godoc to write notes.


Why

The markup is unnoticeable, notes are written in plain text.

No need to write or configure a static site generator.
I can just use `go doc`.

Push it to Bitbucket, GitHub, Launchpad, etc,
and the doc is automatically available on godoc.org.

The style and layout is clean on godoc.org.
With JavaScript turnned on,
readers can use a few simple keyboard shortcuts
(type `?` for the list).
With JavaScript turnned off,
the page looks the same.
And the page is viewable on a console browser,
which supports none of JavaScript, CSS, and images.

As an alternative to a console browser,
readers can also clone the repository,
and view and search notes via the `go doc` command line.
Or access the plain text version of godoc.org:

    curl -H 'Accept: text/plain' https://godoc.org/github.com/icock/comp/about

Comments can be sent as a pull request or an issue.
This can be done via command line, GUI client, or web.
A lot of SCM hosting services allow editing files on the web directly.
If the hosting does not allow this,
readers can use an online IDE.

If I do not want to use godoc.org,
or any of GitHub, BitBucket, Launchpad, etc.,
I can also build the pages via `go doc` on my machine,
then deploy it on any static site hosting services.


Why Not

HTML is used to be a markup for linkable documents.
JavaScript is used to be a small language to add some decorative widgets to web site.
And now they have been abused to build complex GUI applications.

RSS is used to deliver news.
Soon it has been abused to deliver full text of articles.
Unlike one-time news,
articles may be modified later.
And RSS can either ignore announcing the updated version,
or redeliver the new version of whole article.
It does not have the mechanism to deliver the diff.

Editors are used for editing text files.
Later they have been abused as IDEs.

Things tend to be abused.

So why not?


Not Have to Be Go

It dose not have to be Go.
Other languages supporting API doc can be used.

However, a lot of documentation generation tools use a two-column layout by default,
which may not suitable for notes.
Thus it may require additional configuration or customization.
*/
package about
