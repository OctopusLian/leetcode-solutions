/**
 * This reference program is provided by @jiuzhang.com Copyright is reserved.
 * Please indicate the source for forwarding
 */

public class School {
    /*
     * Declare a private attribute *name* of type string.
     */
    private String name;

    /**
     * Declare a setter method `setName` which expect a parameter *name*.
     */
    public void setName(String name) {
        this.name = name;
    }

    /**
     * Declare a getter method `getName` which expect no parameter and return the
     * name of this school
     */
    public String getName() {
        return this.name;
    }
}