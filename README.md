Text generator
-----

This started as an example of how to use git for a small presentation I did
for my coworkers.  I think it's my first time stealing my own work for a
project.  It was educational, but useless.  Now it's a general-case text
generator... and still pretty useless.

I'm hoping to have somewhat reusable libraries as part of this project, but the
command-line interface is what I'm focusing on for now.

### Parsing templates

Text is read in from a template file, passed on the command line.  Text in
double-curly-braces is parsed, reading strings from a file of the same name in
the wordlist directory.  i.e., if your template has {{noun}} in it, a random
line from [wordlist directory]/noun.txt will be inserted in its place.

Word lists are always lowercased.  Even if your template has {{nOUn}}, the file
will still be noun.txt.  However, there are a few rules which are followed when
the case matches certain patterns:

- If the substitution name ("noun" in this case) is all-lowercase, the value
  pulled from the word list is left as-is
- If the substitution name is all-caps ("NOUN"), the value pulled from the word
  list will be all-caps
- If the substitution name's first letter is uppercase, but it isn't all caps,
  the first letter in the word list value will be uppercased.

Text will be parsed forever, until there are no occurrences of double curly
braces.  This can be awesome for building random stories procedurally, as you
can start with {{sentence-1}} and within sentence-1.txt, you can have various
ways of writing a similar opening sentence with other random variables:

    It was a {{adjective}} and stormy night.
    Lightning etched its way across the {{adjective}} sky.

You could do something similar with an entire story, creating a slightly
different version of the same general narrative.  If I get a better system in
place, I hope to be able to allow user input to fill in some entries (parts of
speech, for instance) while the template fills in other parts, creating THE
MOST AMAZING MADLIBS-LIKE EXPERIENCE EVER.

But this "parse all the curlies out" logic can also bite you - e.g., if a file
called "adjective.txt" exists and has {{adjective}} in it.  Don't do that.

### Word lists

The word lists must contain one "word" per line (`\n`).  They can technically
be as many words / characters as you like, they just have to be separated by
newlines for the parser to consider them separate entities.

The word lists are shuffled randomly when they're initialized, and words are
then "popped" off the end of the list.  This ensures that each item is never
used multiple times, which can get tedious and particularly embarrassing at
parties.

HOWEVER: if your template causes the same word list to be used more times than
there are items in it, it would be disastrous to have blank words puked out all
over the place.  Even the most forgiving of friends would mock you for such an
offense!  To avoid this horror, a wordlist that runs out of items is
replenished with its original list, each item carefully wiped off with a tissue
in the hopes that nobody notices they were ever used.  It's a bit disgusting,
really, but far better than the aforementioned mockery.

### Named parameters

In a longer story, you may want some pieces to be consistent.  For instance,
if your template were:

    {{boyname}}: Hey there, {{girlname}}!
    {{girlname}}: Oh, it's you, {{boyname}}.  You know I hate you, right?
    ** {{boyname}} starts sobbing uncontrollably.
    {{girlname}}: Stop being such a Nerdmaster.

You probably wouldn't want each occurrence of `{{boyname}}` and (at least in
*this* example) `{{girlname}}` to be different.  So you can use named
parameters!

    {{boyname->$boy}}: Hey there, {{girlname->$girl}}!
    {{$girl}}: Oh, it's you, {{$boy}}.  You know I hate you, right?
    ** {{$boy}} starts sobbing uncontrollably.
    {{$girl}}: Stop being such a Nerdmaster.

I call this "stabby named parameters".  It's a bit like Ruby lambdas, except I
won't ask you to lie about it being a good syntax.

Example
-----

Check out the source code and try out the example from a sweet, sweet weblib on
yours truly's website's games's page:

```bash
  git clone https://github.com/Nerdmaster/text-generator.git
  cd text-generator
  make
  ./bin/textgen examples/weblibs/prince.txt examples/weblibs/wordlists
```

Using a specific seed for reproducible results:

```bash
  ./bin/textgen examples/weblibs/prince.txt examples/weblibs/wordlists --seed 5
```

And passing in a specific value to override a wordlist entirely:

```bash
  ./bin/textgen examples/weblibs/prince.txt examples/weblibs/wordlists --value "malename:Johnny Five"
```

What about an example of more complexiousness?  Say, for instance, a
**procedurally-generated story**?!?  Look no further than the BRAND NEW
example, "story":

```bash
  ./bin/textgen examples/story/story.txt examples/story/wordlists/
```

You might get a story like this:

> Call me Nerdmaster.  Or Queequeg, I don't really care.  The point is, there
> was a fire inside my tree and it kept me warmer than if I were out in the
> angry rain, okay?!?  I was reading one of my favorite pop-up books, "The cat
> and the barn", when the fire went out.  I sobbed for hours, terrified, and
> screamed for my mommy.  But it was all just a dream!  Oh, how content I felt
> when I woke up!

But it could also be RADICALLY DIFFERENT, such as:

> Well, the weather outside was blue, but inside it was so happy that I just
> sat by the fire.  Then, the fire simply stopped.  There's no other way to put
> it.  I sat there in the cold, dark, angry room and figured I had no more then
> 3 seconds before I succumbed to insanity.  But then I woke up.  It had only
> been a dream!  Of course, I was still in the Nazi torture cat, but that's
> better than being in a cold, dark, silly room during a storm.

OMFG THE SAME GENERAL STORY WITH DIFFERENT SENTENCE STRUCTURE THIS IS AMAZING.

I hope you enjoy this {{adjective}} {{noun}} I've created for you.
