# gojos
Library to read Java serialized objects (Serializible via ObjectOutputStream) in Go.
Performs reading Java serialized object bytes into internal model and mapping to Go structure.

For example java class
``` 
 public static class Primitives implements Serializable {
        private byte byteVal;
        private boolean boolVal;
        private char charVal;
 }
```
Serialized with
```
    ObjectOutputStream oos = new ObjectOutputStream(...);
    oos.writeObject(primitives);
```

Can be extracted into GO structure

```
type Primitives struct {
    ByteVal  byte
    BoolVal  bool
    CharVal  byte
}	
```
More examples are in tests

