## About Go

Before we roll up our sleeves, it is important to understand why the Go programming language was developed. Let’s have a look at a [quote](https://talks.golang.org/2012/splash.article)[^about-go-1] [[video](https://www.youtube.com/watch?v=bmZNaUcwBt4)] from [Rob Pike](https://en.wikipedia.org/wiki/Rob_Pike)[^about-go-2], one of its creators.

[^about-go-1]: Video: https://www.youtube.com/watch?v=bmZNaUcwBt4 "Go at Google: Language Design in the Service of Software Engineering" https://talks.golang.org/2012/splash.article

[^about-go-2]: Rob Pike https://en.wikipedia.org/wiki/Rob_Pike

> “Go was conceived as an answer to some of the problems we were seeing developing software infrastructure at  Google. The computing landscape today is almost unrelated to the environment in the languages being used, mostly C++, Java and Python have been created. The problems introduced by multicore processors, networked systems, massive computation clusters, and the web programming model  were being worked around rather than addressed head-on.”

> “Moreover, the scale has changed: todays’s server programs comprise ten of millions of lines of code, are worked on by hundred or even thousands of programmers, and are updated literally every day. To make matters worse, build time has stretched to many minutes, even hours, on large compilation clusters”.

The **language creators** all have brought major contributions to our field. [Robert Griesemer](https://de.wikipedia.org/wiki/Robert_Griesemer)[^about-go-3] has worked on the V8 Javascript engine and on the Java VM hotspot; Rob Pike was on the Unix team at the Bell Lab and created the first windowing system for Unix; [Ken Thompson](https://en.wikipedia.org/wiki/Ken_Thompson) designed and implemented the original Unix. Thompson and Pike the co-creator of UTF-8.

[^about-go-3]: Robert Griesemer https://de.wikipedia.org/wiki/Robert_Griesemer

Go has a very **simple programming model**. The authors purposefully avoided complex and advanced constructions. The mechanism called _exceptions_ found in many programming languages has not been included, which can be surprising at first. _Generics_ were added only recently. Go has a “less is better” approach and its maintainers are extremely wary not to add unnecessary features to the language.

As a result, you can learn Go in only a few days.

Some of the **cruft** accumulated by older languages has been omitted, like the need to specifically declare a data type, although a modern compiler should often infer the type.

Go also strongly favours certain **idioms**, and that’s an excellent thing to achieve higher consistency, and to avoid endless debates on personal tastes. The `go fmt` tool is a good illustration. It settles the debate about formatting.

It has a **garbage collector** language, so the programmer does not have manage memory allocation themselves.

It is built with **concurrent programming** in mind, with its *coroutines* and *channels*. **Testing** also has a key place with dedicated constructs.

Go compiles to **native code** and allows **cross compilation** to another platform. So for instance one can create a executable for GNU/Linux on a macOS machine.

The tools are well designed and well thought, and so is the language. There is a decent amount of libraries available. The **memory footprint** is extremely small. It is very capable to handle heavy loads.

It was designed to run “server-side” programs, and although solutions exist to plug Go to a GUI, its sweet spot is on the server side. Youtube relies amongst other things on a Google project called vitess, and [vitess is written in Go](https://blog.golang.org/fosdem14)[^about-go-4].

[^about-go-4]: FOSDEM 2014 https://blog.golang.org/fosdem14

From a career perspective, Go is also is a [sensible choice](https://survey.stackoverflow.co/2023/#technology-top-paying-technologies)[^so-top-paying].

[^so-top-paying]: Survey on StackOverflow https://insights.stackoverflow.com/survey/2017#top-paying-technologie
