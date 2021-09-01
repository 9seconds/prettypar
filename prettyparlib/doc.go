// prettyparlib is a library that format text strings and paragraphs.
//
// There are couple of utilities available: namely, fold and par.
// fold (https://linux.die.net/man/1/fold) wraps a text string by
// a given length. Unfortunately, it knows nothing, that paragraphs
// may have an additional context structure.
//
// par is more powerful that prettypar
// (https://bitbucket.org/amc-nicemice/par/src/master/) but way more
// complex at the same time. It knows about comments and text decorations
// but works poorly with lists. Also, its knowledge of text decorations
// drives me nuts. I haven't seen these ASCII-style text boxes for ages.
// But par tends to consider dots (.) as an element of that box and align
// based on that.
//
// So, par is very opionated. It knows how to format this guy
//
//     # Hello, this is an example
//     # paragraph full of
//     # some random thoughts.
//     #
//     # Another paragraph.
//     #
//     #     1. lala
//     #     2. blabla
//     #
//     #   * long-long list of something
//
// into this
//
//     # Hello, this is an example paragraph full of some random thoughts.
//     #
//     # Another paragraph.
//     #
//     #     1. lala
//     #     2. blabla
//     #
//     #   * long-long list of something
package prettyparlib
