# Channels Example

This is a Go program containing functions that consume and produce channels. I
think this is potentially a nicer API shape than e.g. `filepath.WalkDir`
(running your function as a side-effect), `bufio.Scanner` (a special-purpose
iterator), `regexp.FindAll` (appending to and then returning a potentially large
slice), and so on. With this `chan Value` interface ‘hinge’ type, you can
relatively easily compose unbounded producer/consumer functions.

The point of this toy project is to show what it would take for Go code to look
(and more easily perform) a bit more like Unix shell/PowerShell pipelines.
(Still nested function calls instead of a linear pipeline, but the generic
composition aspect and the bounded memory consumption aspects are there.)

To me it seems like `chan` is the fundamental organizing principle of Go — Go
‘wants’ to be written this way, I sort of think? — and so there must be some
reason I don’t understand that explains why the standard library is not written
in a style similar to what I show here. (More generally, you’d want `chan
interface{}`, or `chan TaggedUnionThing`, I suppose. And perhaps that is why
it’s not standard. Still, it seems solveable...)
