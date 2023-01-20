package test;

import java.util.HashMap;
import java.io.Serializable;

import cmd.GojoTest;

public class Test45 implements GojoTest<Test45.Model> {
    public static class Model implements Serializable {
        private HashMap<String, String> hm = new HashMap<>();
    }

    @Override
    public Test45.Model getnerateTestObject() {
        Model m = new Model();

        m.hm.put("1", null);
        return m;
    }
}