## Abstraction vs low level

Why not simply manipulate “integers”? Why “floating point arithmetic” and different integers? After all, we mentionned the importance to abstract ourselves from the underlyting platform. Some languages only expose “integers” and “decimals”, but it comes with a performance cost. Go types are closer to the hardware architecture. That is a trade-off the language authors decided to do, based on its intent and purpose, where high performance is key.

From a teaching perspective, this design choice gets a bit into our way as it clutters the explanations, at least at the begining. On the bright side, as far as learning goes, the reader is exposed to technical underlying details that would otherwise remain hidden, and can already get a grasp at them.
