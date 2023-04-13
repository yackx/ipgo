## Abstraction vs low level

Why not simply manipulate “integers”? Why “floating point arithmetic” and different integers? After all, we mentionned the importance to abstract ourselves from the underlyting platform. Some languages only expose “integers” and “decimals”, but it comes with a substantial performance cost. Go integers types are closer to the hardware architecture. That is a trade-off the languages authors decided to do, based on the intent and purpose of the language, where high performance is key.

From a teaching perspective, this design choice gets a bit into our way as it clutters the explanations, at least at the begining. On the bright side, as far as learning goes, you are exposed to technical underlying details that would otherwise remain hidden, and you can already get a grasp at them.
