public class Visibility{
       public static void main(String[] args) {
        Tag tag = new Tag();

        new Thread(()->{
            try {
                // sleep is necessary
                Thread.sleep(1000);
                tag.flag = -1;
                System.out.println("flag was updated to -1");
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }).start();
        // main thread check flag in loop
        while(tag.flag !=-1){

        }
        System.out.println("main thread read flag is -1 successful");
    }
    static class Tag{
        // volatile int flag = 0;
        int flag =0;
    }
}