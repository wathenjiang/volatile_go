public class Student {
      String name;
      int age;
      // private constructor
      private Student(String name, int age) {
          this.name = name;
          this.age = age;
      }
      // use volatile
      private static volatile Student student;
      //double check null
      public static Student getStudent() {
          if (student == null) {
              synchronized (Student.class) {
                  if (student == null) {
                      student = new Student("hello world",12);
                  }
              }
          }
          return student;
      }
      public String getName() {
          return name;
      }
      public int getAge() {
          return age;
      }
  }
