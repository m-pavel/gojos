package cmd;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;
import java.net.URL;

public class TestGenerator {

    private static boolean debug = true;

    private File getFile(String path, Class<?> cls) {
        return new File(new File(path), cls.getSimpleName().toLowerCase() + ".bin");
    }

    private void generate(String path) throws IOException {
        URL url = this.getClass().getClassLoader().getResource("test");
        String[] files = new File(url.getPath()).list();
        for (String fname : files) {
            processFile(fname, path);
        }
    }

    private void processFile(String fname, String path) throws IOException {
        Object obj = null;
        try {
            Class<?> testClass = Class.forName("test." + fname.substring(0, fname.indexOf('.')));
            obj = testClass.newInstance();
        } catch (ClassNotFoundException | InstantiationException | IllegalAccessException e) {
            // ignore
        }
        if (obj instanceof GojoTest) {
            File file = getFile(path, obj.getClass());
            Object to = ((GojoTest<?>) obj).getnerateTestObject();
            ObjectOutputStream oos = new ObjectOutputStream(new FileOutputStream(file));
            oos.writeObject(to);
            oos.close();
            if (debug) {
                ObjectInputStream ois = new ObjectInputStream(new FileInputStream(file));
                Object robj = null;
                try {
                    robj = ois.readObject();
                } catch (ClassNotFoundException e) {
                    // ignore
                }
                ois.close();
                System.out.println(robj);
            }
        }
    }

    public static void main(String[] args) throws IOException {
        switch (args.length) {
            case 1:
                new TestGenerator().generate(args[0]);
                break;
            case 2:
                new TestGenerator().processFile(args[1], args[0]);
                break;
            default:
                throw new IllegalArgumentException("Provide output folder");

        }
    }
}
