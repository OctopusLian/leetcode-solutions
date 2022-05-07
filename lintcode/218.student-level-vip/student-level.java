public class Student {
    /**
     * Declare two public attributes name(string) and score(int).
     */
    public String name;
    public int score;

    /**
     * Declare a constructor expect a name as a parameter.
     */
    public Student(String name) {
        this.name = name;
    }

    /**
     * Declare a public method `getLevel` to get the level(char) of this student.
     */
    public char getLevel() {
        if (score >= 90) {
            return 'A';
        } else if (score >= 80) {
            return 'B';
        } else if (score >= 60) {
            return 'C';
        } else {
            return 'D';
        }
    }
}