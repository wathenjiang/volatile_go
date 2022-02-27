
#!/bin/bash
trap "rm *.class;kill 0" EXIT
javac Visibility.java
java Visibility