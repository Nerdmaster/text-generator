# How to Contribute

## Reporting Issues

Submit a [Github issue](./issues) if you can't submit code.

But when you submit an issue, please remember to GoDEEP:

* **D**escribe the problem in words a non-metadata-expert can understand
* **E**xplain what is happening that you consider incorrect
* **E**xplain what you expected to have happen
* **P**rovide steps to reproduce what you perceive as the problem

## What's yours is mine

If you commit stuff, it better be stuff you don't mind giving up, because once
it's here, it's part of the project, and licensed as such.

The current license is [defined in the license file](./LICENSE.txt).  If that
file isn't there, the default is a very restrictive license that makes me the
sole owner of all code in this project.

If you commit with any obvious non-personal email address I'll probably reject
your commit.  Chances are it's not your property if you code it AT WORK.

## git-flow

I use [git-flow](https://github.com/nvie/gitflow).  If you want to contribute,
you'll use it, too.  Feature branches start with "feature/".  All pull requests
must be based on the "develop" branch.

If you make a branch off master, I may or may not accept it, but it definitely
won't have your name on it.

## Commit messages

I follow Tim Pope's guidelines for [good commit
practices](http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html).
You will, too, if you want your commit to not be rewritten by me.

In case the good Pope's site goes down, here’s a model git commit message:

```
Capitalized, short (50 chars or less) summary

More detailed explanatory text, if necessary.  Wrap it to about 72
characters or so.  In some contexts, the first line is treated as the
subject of an email and the rest of the text as the body.  The blank
line separating the summary from the body is critical (unless you omit
the body entirely); tools like rebase can get confused if you run the
two together.

Write your commit message in the imperative: "Fix bug" and not "Fixed bug"
or "Fixes bug."  This convention matches up with commit messages generated
by commands like git merge and git revert.

Further paragraphs come after blank lines.

- Bullet points are okay, too

- Typically a hyphen or asterisk is used for the bullet, followed by a
  single space, with blank lines in between, but conventions vary here

- Use a hanging indent
```

Don't like it?  Go elsewhere.

## Workflow

* Fork the repository
* Create a nice workspace.  If you use `go get`, this is a bit weird:
  * `go get go.nerdbucket.com/text`
  * `cd $GOPATH/src/go.nerdbucket.com/text`
  * `git remote rename origin upstream`
  * `git remote add origin git@github.com:YOURNAME/text.git`
  * `git checkout develop`
  * `git pull origin develop`
  * Don't blame me, this is due to the insanity of what is otherwise a decent
    language.  Go discourages local importing of packages, opting instead for
    URLs (namespaces, but they're URLs if you want `go get` to work).
* **OR** if you have GNU Make or similar, you can do this a bit more easily:
  * Grab the source however you like - `go get`, clone `https://github.com/Nerdmaster/text-generator`,
    or fork and clone your copy of the repo
  * Switch to the repo's directory
  * Create a project-specific GOPATH, e.g., `export GOPATH=/tmp/nerdgotext`
    * This is *usually* optional, but may be necessary in some edge cases, such
      as if you have a conflicting version of a dependency, or you already have
      the real `go.nerdbucket.com/text` package.  It can also be useful just
      to help keep your gopath cleaner.
  * Set up dependencies: `make alldeps`
* Create a topic branch *based on develop*
  * I like hyphens and all lowercase for branch names.  You want to contribute?
    Then you do, too.  `feature/foo-bar`, not `feature/FooBar`
* Make commits of logical units.
  * If you follow the lunacy of "one commit per pull request", and your PR is
    non-trivial, I will not only not accept your work, but I'll comment on it
    as if I'm going to accept it, wasting as much of your time as I possibly
    can.  I may even steal some of it and put it into the project while leaving
    your PR up in a state of perpetual "code review".
* Run `make format` - this runs `gofmt -l -w -s` on all files, ensuring I don't
  smack you upside the head with your whitespace mistakes
  * Yes, I hate some of their formatting choices, too, but it's a standard, and
    it's easy to adopt and use.
* If you created an issue, you can close it by including "Closes #issue" in
  your commit message. See [Github's blog post for more
  details](https://github.com/blog/1386-closing-issues-via-commit-messages)
  * I'd prefer you put this in the pull request and not the actual commit
    history if you don't have a good reason for it
* Make sure you have added tests if it seems appropriate to do so
* Run *all* the tests to assure nothing else was accidentally broken (`make test`)
