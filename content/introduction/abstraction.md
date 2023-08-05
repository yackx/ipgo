## Abstraction

In order to write programs, we need **abstractions**.

Consider a hard disk drive (HDD). It is an enclosure containing rapidly rotating disks or platters, each of which is coated with magnetic material in order to store data. Arms with magnetic heads move above the platters in order to read (and write) data that can represent an image, a song or a poem, or more generally anything we call a *file*.

![A hard drive needs an abstraction](content/introduction/laptop-hard-drive-exposed-wikipedia.jpg){ width=230px }

We could read the poem stored on the hard drive and display it on screen by sending commands to move the heads above the right platter on the hard drive, communicating with the CPU and the hardware directly. But the task would be incredibly difficult to achieve. Every program ever written would have to repeat the same operations, for every single data it needs to read, and for every possible disk geometries. This would be at best a tedious and error-prone task. Above all, it would be pointless.

Maybe the hard disk drive can be seen as _cyclinders_ and _sectors_ that form some logical organization. Which it does in reality. This is a first level of **abstraction**. But this is still not the abstraction we are looking for --- unless we are writing specific parts of an operating system, or a HDD controller.

For sure, we have a more casual purpose. Say we want to read the poem from the disk with as little knowledge about the underlying hardware as possible. We are interessted in retrieving the poem, not about the low-level hardware details.

Now consider the following program fragment:

```go
poem, err := ioutil.ReadFile("poem.txt")
defer poem.Close()
if e != nil {
    panic(e)
}
fmt.Print(string(poem))
```

There are several things that require an explanation. What is `err` or `panic` for instance. All in due time. For now, using a high-level programming language, we have *abstracted* the process of retrieving the information on the disk magnetic surface. In fact, the abstraction is layered: the Go compiler provides you with a first abstraction from the operating system, which in turn abstracts you for the processor, the memory and the hard disk. This is what we were looking for: a mean to **express ideas and concepts while hiding the underlying complexity and details**. It is still nice to know how a hard drive or an operating system work though. But suffice to say it is much more efficient and confortable to rely on the work done by others to access a storage medium, and focus on our program.

> _Note_: Nowadays hard disk drives (HDD) are superseded replaced by solid state drives (SSD) that have no mechaninal moving parts. Moving parts or not, you still want to be abstracted from the bare medium. By virtue of abstraction, our program above will still work, no matter the actual underlying physical media, HDD, SSD or other.
