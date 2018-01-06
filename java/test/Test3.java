package test;

import java.util.HashMap;

import cmd.GojoTest;

public class Test3 implements GojoTest<HashMap<String, String>> {

    @Override
    public HashMap<String, String> getnerateTestObject() {
        HashMap<String, String> hm = new HashMap<>();
        hm.put("1", "2");
        return hm;
    }
}