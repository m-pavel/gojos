package test;

import java.util.HashMap;

import cmd.GojoTest;

public class Test4 implements GojoTest<HashMap<String, String>> {
    @Override
    public HashMap<String, String> getnerateTestObject() {
        HashMap<String, String> hm = new HashMap<>();
        hm.put("1", "2");
        hm.put("2", "3");
        hm.put("4", "5");
        hm.put("6", null);
        hm.put(null, "7");
        return hm;
    }
}