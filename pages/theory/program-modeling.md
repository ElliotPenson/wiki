# Program Modeling

Information on this page is taken from _The Algorithm Design Manual_ by Steven
S. Skiena.

Modeling is the art of formulating your application in terms of precisely
described, well-understood problems. Proper modeling can eliminate the need to
design or even implement algorithms, by relating your application to what has
been done before.

Real-world applications involve real-world objects. Most algorithms, however,
are designed to work on rigorously defined _abstract_ structures. To exploit the
algorithms literature, you must learn to describe your problem abstractly, in
terms of procedures on fundamental structures.

- **Permutations** are arrangements, or orderings of items. Usually the object
  in question if your problem seeks an "arrangement," "tour," "ordering," or
  "sequence."
- **Subsets** are selects from a set of items. Usually the object in question if
  your problem seeks a "cluster," "collection," "committee," "group,"
  "packaging," or "selection."
- **Trees** are hierarchical relationships between items. Usually the object in
  question whenever your problem seeks a "hierarchy," "dominance relationship,"
  "ancestor/descendant relationship," or "taxonomy."
- **Graphs** represent relationships between arbitrary pairs of objects. Usually
  the object in question whenever you seek a "network," "circuit," "web," or
  "relationship."
- **Points** represent locations in some geometric space. Usually the object in
  question whenever your problems work on "sites," "positions," "date records,"
  or "locations."
- **Polygons** represent regions in some geometric spaces. Usually the object in
  question whenever you are working on "shapes," "regions," "configurations," or
  "boundaries."
- **Strings** represent sequences of characters or patterns. Usually the object
  in question whenever you are dealing with "text," "characters," "patterns," or
  "labels."

Learn to think recursively. Recursive structures occur everywhere in the
algorithmic world. Each of the abstract structures described above can be
thought about recursively; they are big things made of smaller things of the
same type. Each structure has operations (like _delete_) that produce new
versions of the same type.
