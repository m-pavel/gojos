package test;

import java.io.Serializable;

import cmd.GojoTest;

public class PrimitivesTest implements GojoTest<PrimitivesTest.Primitives>{
    public static class Primitives implements Serializable {
        private byte byteVal;
        private boolean boolVal;
        private char charVal;
        private short shortVal;        
        private int intVal;
        private float floatVal;
        private long longVal;
        private double dblVal;
    }

    @Override
    public Primitives getnerateTestObject() {
        Primitives p = new Primitives();
        p.boolVal = true;
        p.byteVal = (byte) 55;
        p.charVal = 'A';
        p.dblVal = 325.55;
        p.floatVal = 11.7f;
        p.intVal = 42;
        p.longVal = 55555555555555L;
        p.shortVal = (short) 200;
        return p;
    }
}
