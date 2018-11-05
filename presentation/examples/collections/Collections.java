package main;

public class Main {

    public static void main() {
        
    }

    public static void printUsers(String filter) {
        users.stream()
            .filter(p -> p.getName().contains(filter))
            .sorted(Comparator.comparing(User::getName))
            .forEach(u -> System.out.println(u));
    }    
}

