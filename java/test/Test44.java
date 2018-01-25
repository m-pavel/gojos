package test;

import java.util.HashMap;
import java.io.Serializable;

import cmd.GojoTest;

public class Test44 implements GojoTest<Test44.Model> {
    public static class Model implements Serializable {
         private HashMap<String, String> hm = new HashMap<>();
    }
    @Override
    public Test44.Model getnerateTestObject() {
        Model m = new Model();

        m.hm.put("1", "2");
        m.hm.put("2", "3");
        m.hm.put("4", "5");
        m.hm.put("6", null);
        m.hm.put(null, "7");
        return m;
    }
}