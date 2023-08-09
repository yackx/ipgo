## Bits and bytes

It is a well-known fact that computers only understand "ones and zeros". The most basic unit of information that the computer in store is called a **bit**. A bit can hold two possible values: `1` or `0`. Actually, one and zero are merely conventions. Instead we could use `on` and `off`, or `true` and `false`. We'll stick to the long standing convention though.

Dealing with `1` or `0` exclusively would be extermely cumbersome, even for a computer processor. Bits are usually[^byte] grouped by **8** to form a **byte**.

[^byte]: It is actually more complex than that. Abate can technically be of any size. But the common definition is to use eight bits.

\begin{center}
\begin{tikzpicture}[every node/.style={draw, minimum size=1cm}]

% Define the byte as an array of values
\edef\byte{{1,1,1,0,0,1,0,1}}

% Draw the bits as nodes in a row
\foreach \i in {0,...,7} {
  \pgfmathsetmacro{\b}{\byte[\i]}
  \node at (\i,0) {\b};
  \pgfmathtruncatemacro{\bpos}{7-\i}
  \node[anchor=north] at (\i,-0.5) {\tiny bit \bpos};
}

\end{tikzpicture}
\end{center}

This gives us $2^8$ or $256$ possible combinations.

What does the above byte represents? It depends. It can represent a positive number (unsigned integer). To convert it to a decimal value, multiply each bit by $2^k$, with $k$ equal to the bit position. In our case:

\begin{center}
\begin{tikzpicture}[every node/.style={draw, minimum size=1cm}]

% Define the byte as an array of values
\edef\byte{{1,1,1,0,0,1,0,1}}

% Draw the bits as nodes in a row
\foreach \i in {0,...,7} {
  \pgfmathsetmacro{\b}{\byte[\i]}
  \node at (\i,0) {\b};
  \pgfmathtruncatemacro{\bpos}{7-\i}
  \node[anchor=north] at (\i,-1) {$2^\bpos$};
  \pgfmathtruncatemacro{\bweight}{2^(7-\i)}
  \node[anchor=north] at (\i,-2) {\bweight};
}

\end{tikzpicture}
\end{center}

$$1.2^7 + 1.2^6 + 1.2^5+ 0.2^4+ 0.2^3 + 1.2^2 + 0.2^1 + 1.2^0$$
$$= 1.128 + 1.64 + 1.32 + 0.16 + 0.8 + 1.4 + 0.2 + 1.1$$
$$= 128 + 64 + 32 + 4 + 1$$
$$= 229$$
