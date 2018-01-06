package test;

import cmd.GojoTest;
import model.t7.Child;

public class Test7 implements GojoTest<Child> {
    @Override
    public Child getnerateTestObject() {
        Child c1 = new Child();

        return c1;
    }
}